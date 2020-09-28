package logger

import (
	"github.com/jameschz/go-base/lib/util"
	"io"
	"log"
	"os"
)

var (
	_loggerFile *os.File
	_loggerInfo *log.Logger
	_loggerWarn *log.Logger
	_loggerErr  *log.Logger
)

func _openLoggerFile() error {
	logPath := util.GetRootPath() + "/log/logger.log"
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
	if err := _openLoggerFile(); err == nil {
		_loggerInfo = log.New(io.MultiWriter(_loggerFile), "[Info]", log.Ldate|log.Ltime|log.Lshortfile)
		_loggerInfo.Println(msgs...)
		_closeLoggerFile()
	}
}

// Warn : log warn msg
func Warn(msgs ...interface{}) {
	if err := _openLoggerFile(); err == nil {
		_loggerWarn = log.New(io.MultiWriter(_loggerFile), "[Warn]", log.Ldate|log.Ltime|log.Lshortfile)
		_loggerWarn.Println(msgs...)
		_closeLoggerFile()
	}
}

// Error : log error msg
func Error(msgs ...interface{}) {
	if err := _openLoggerFile(); err == nil {
		_loggerErr = log.New(io.MultiWriter(_loggerFile), "[Error]", log.Ldate|log.Ltime|log.Lshortfile)
		_loggerErr.Println(msgs...)
		_closeLoggerFile()
	}
}
