package logger

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/jameschz/go-base/lib/gutil"
)

type LogStruct struct {
	Prefix  string
	LogInfo []interface{}
}

var (
	_loggerFile *os.File
	// _loggerInfo *log.Logger
	// _loggerWarn *log.Logger
	// _loggerErr  *log.Logger

	_logger *log.Logger

	_DataLog chan LogStruct
)

func init() {

	_DataLog = make(chan LogStruct, 1024)

	go InitReceive()
}

func _openLoggerFile() error {

	date := time.Now().Format("20060102")
	logPath := gutil.GetRootPath() + "/log/logger_" + date + ".log"

	_, err := os.Open(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(logPath)
			if err != nil {
				log.Fatalln("> logger create file error:", err)

				logPath = gutil.GetRootPath() + "/log/logger.log"
			}
			file.Close()
		}
	}

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("> logger init error:", err)
		return err
	}
	_loggerFile = logFile

	return nil
}

func _closeLoggerFile() {
	if _loggerFile != nil {
		_loggerFile.Close()
	}
}

// Info : log info msg
func Info(msgs ...interface{}) {
	// if err := _openLoggerFile(); err == nil {
	// 	_loggerInfo = log.New(io.MultiWriter(_loggerFile), "[Info]", log.Ldate|log.Ltime|log.Lshortfile)
	// 	_loggerInfo.Println(msgs...)
	// 	_closeLoggerFile()
	// }

	_DataLog <- LogStruct{"[Info]", msgs}
}

// Warn : log warn msg
func Warn(msgs ...interface{}) {
	// if err := _openLoggerFile(); err == nil {
	// 	_loggerWarn = log.New(io.MultiWriter(_loggerFile), "[Warn]", log.Ldate|log.Ltime|log.Lshortfile)
	// 	_loggerWarn.Println(msgs...)
	// 	_closeLoggerFile()
	// }

	_DataLog <- LogStruct{"[Warn]", msgs}
}

// Error : log error msg
func Error(msgs ...interface{}) {
	// if err := _openLoggerFile(); err == nil {
	// 	_loggerErr = log.New(io.MultiWriter(_loggerFile), "[Error]", log.Ldate|log.Ltime|log.Lshortfile)
	// 	_loggerErr.Println(msgs...)
	// 	_closeLoggerFile()
	// }

	_DataLog <- LogStruct{"[Error]", msgs}
}

func WriteLog(logs []LogStruct) {

	if err := _openLoggerFile(); err == nil {
		_logger = log.New(io.MultiWriter(_loggerFile), "", log.Ldate|log.Ltime|log.Lshortfile)

		for _, log := range logs {

			_logger.SetPrefix(log.Prefix)
			_logger.Println(log.LogInfo...)
		}

		_closeLoggerFile()
	}
}

func InitReceive() {

	go func() {

		for keepGoing := true; keepGoing; {
			var batch []LogStruct
			expire := time.After(30 * time.Second)
			for {
				select {

				case dataLog := <-_DataLog:

					batch = append(batch, dataLog)
					if len(batch) == cap(_DataLog) {
						goto done
					}

				case <-expire:
					goto done
				}
			}

		done:
			if len(batch) > 0 {
				WriteLog(batch)
			}
		}
	}()
}
