package logger

type Logger interface {
	Info(format string, args ...interface{})
	DBInfo(format string, args ...interface{})
	Error(format string, args ...interface{})
	DBError(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	DBFatal(format string, args ...interface{})
	Status(format string, args ...interface{})
	DBStatus(format string, args ...interface{})
	Request(format string, args ...interface{})
	Debug(format string, args ...interface{})
}


var Reset = "\033[0m" 
var Red = "\033[31m" 
var Yellow = "\033[33m" 
var Cyan = "\033[36m" 

var Green = "\033[32m"
var Blue = "\033[34m"

var Magenta = "\033[35m"
var Orange = "\033[38;5;208m"
var Pink = "\033[38;5;206m"

var Purple = "\033[38;5;141m"
var White = "\033[38;5;255m"

var logger Logger

func GetLogger() Logger {
	return logger
}

func SetLogger(l Logger) {
	logger = l
}