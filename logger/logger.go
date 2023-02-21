package logger

import (
	"log"
	"os"
)

const flags int = log.Ltime | log.Lmicroseconds | log.Lshortfile
var infoLogger 	*log.Logger = log.New(os.Stdout, "INFO:  ", flags)
var	warnLogger 	*log.Logger = log.New(os.Stdout, "WARN:  ", flags)
var errorLogger *log.Logger = log.New(os.Stdout, "ERROR: ", flags)

func Info(v...any) {
	infoLogger.Println(v...)
}

func Warn(v...any) {
	warnLogger.Println(v...)
}

func Error(v...any) {
	errorLogger.Println(v...)
}
