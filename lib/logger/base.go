package logger

import (
	"go-base/lib/util"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	_logger_file *os.File
	_logger_info *log.Logger
	_logger_warn *log.Logger
	_logger_err  *log.Logger
)

func _openLoggerFile() error {
	logPath := util.GetRootPath() + "/log/logger.log"
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("> logger init error:", err)
		return err
	}
	_logger_file = logFile
	return nil
}

func _closeLoggerFile() {
	if _logger_file != nil {
		_logger_file.Close()
	}
}

func Info(msgs ...interface{}) {
	if err := _openLoggerFile(); err == nil {
		_logger_info = log.New(io.MultiWriter(_logger_file), "[Info]", log.Ldate|log.Ltime|log.Lshortfile)
		_logger_info.Println(msgs...)
		_closeLoggerFile()
	}
}

func Warn(msgs ...interface{}) {
	if err := _openLoggerFile(); err == nil {
		fmt.Println("ssss")
		_logger_warn = log.New(io.MultiWriter(_logger_file), "[Warn]", log.Ldate|log.Ltime|log.Lshortfile)
		_logger_warn.Println(msgs...)
		_closeLoggerFile()
	}
}

func Error(msgs ...interface{}) {
	if err := _openLoggerFile(); err == nil {
		_logger_err = log.New(io.MultiWriter(_logger_file), "[Error]", log.Ldate|log.Ltime|log.Lshortfile)
		_logger_err.Println(msgs...)
		_closeLoggerFile()
	}
}
