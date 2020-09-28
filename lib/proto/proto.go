package proto

import (
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
