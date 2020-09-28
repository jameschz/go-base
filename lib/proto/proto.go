package proto

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strconv"

	proto_base "github.com/golang/protobuf/proto"
)

// Message : message interface
type Message interface {
	proto_base.Message
}

// SendMessage : send message to conn
func SendMessage(conn net.Conn, msg Message) (err error) {
	var (
		sendBytes []byte
		sendBuf   [4]byte
		readLen   int
	)
	// protobuf pack msg
	if sendBytes, err = proto_base.Marshal(msg); err != nil {
		return err
	}
	// get msg length
	binary.BigEndian.PutUint32(sendBuf[:4], uint32(len(sendBytes)))
	// send msg length
	if readLen, err = conn.Write(sendBuf[:4]); readLen != 4 && err != nil {
		if readLen == 0 {
			return errors.New("proto.SendMessage length is zero")
		}
		return err
	}
	// send msg data
	if readLen, err = conn.Write(sendBytes); err != nil {
		if readLen == 0 {
			return errors.New("proto.SendMessage write error")
		}
		return err
	}
	return nil
}

// ReadMessage : read message from conn to buf
func ReadMessage(conn net.Conn, buf []byte, msg Message) (err error) {
	var (
		pkgLen  uint32
		readLen int
	)
	// read msg length
	if _, err = conn.Read(buf[:4]); err != nil {
		return errors.New("proto.ReadMessage read error")
	}
	// get msg length
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	fmt.Println("> pkgLen : ", strconv.FormatInt(int64(pkgLen), 10))
	// get msg data
	if readLen, err = conn.Read(buf[:pkgLen]); readLen != int(pkgLen) || err != nil {
		if err == nil {
			return errors.New("proto.ReadMessage length error")
		}
		return err
	}
	// protobuf unpack msg
	if err = proto_base.Unmarshal(buf[:pkgLen], msg); err != nil {
		return err
	}
	return nil
}

// SendPBC : send pbc message to conn
func SendPBC(conn net.Conn, msg Message) (err error) {
	var (
		msgBuf    []byte
		msgLen    []byte
		pStatus   []byte
		pType     []byte
		sendLen   []byte
		sendBytes []byte
		readLen   int
	)
	// protobuf pack msg
	if msgBuf, err = proto_base.Marshal(msg); err != nil {
		return err
	}
	// set msg status
	pStatus = make([]byte, 2)
	binary.LittleEndian.PutUint16(pStatus, uint16(104))
	// set msg type
	pType = make([]byte, 2)
	binary.LittleEndian.PutUint16(pType, uint16(2))
	// set msg length
	msgLen = make([]byte, 4)
	binary.LittleEndian.PutUint16(msgLen, uint16(len(msgBuf)))
	// set send length
	sendLen = make([]byte, 2)
	binary.LittleEndian.PutUint16(sendLen, uint16(len(msgBuf)+8))
	// build msg bytes
	sendBytes = bytes.Join([][]byte{sendLen, pStatus, pType, msgLen, msgBuf}, []byte{})
	fmt.Println("> proto.SendPBC : sendBytes ========== ", sendBytes)
	// send msg data
	if readLen, err = conn.Write(sendBytes); err != nil {
		if readLen == 0 {
			return errors.New("proto.SendPBC write error")
		}
		return err
	}
	return nil
}

// ReadPBC : read pbc message from conn to buf
func ReadPBC(conn net.Conn, buf []byte, msg Message) (err error) {
	var (
		pStatus   uint16
		pType     uint16
		pLen      uint32
		readLen   int
		readBytes []byte
	)
	// read msg length
	if readLen, err = conn.Read(buf); err != nil {
		return errors.New("proto.SendPBC read error")
	}
	if readLen <= 10 {
		return errors.New("proto.SendPBC msg length error")
	}
	readBytes = buf[:readLen]
	fmt.Println("> proto.ReadPBC : readBytes ========== ", readBytes)
	// get msg len
	pStatus = binary.LittleEndian.Uint16(readBytes[2:4])
	pType = binary.LittleEndian.Uint16(readBytes[4:6])
	pLen = binary.LittleEndian.Uint32(readBytes[6:10])
	fmt.Println("> proto.ReadPBC : readVars ========== ", readLen, pStatus, pType, pLen)
	// protobuf unpack msg (escape 6 bytes for PBC)
	if err = proto_base.Unmarshal(readBytes[10:readLen], msg); err != nil {
		return err
	}
	return nil
}
