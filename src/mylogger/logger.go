package mylogger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

//自定义一个日志库
/*
	需求分析
	1. 支持往不同的地方输出日志
	2. 日志分级别
		1. Debug
		2. Trace
		3. Info
		4. Waring
		5. Error
		6. Fatal
		7. UNKNOWN
	3. 日志要支持开关控制,比如说开发的时候什么级别都能输出,但是上线之后就只有INFO级别往下的才能输出
	4. 完整的日志记录要包含有时间,行号,文件名,日志级别,日志信息
	5. 日志文件要切割
		1. 按照文件大小切割
			1. 每次记录日志前都判断一下当前写的这个文件的文件大小
		2. 按日期切割
*/
type LogLevel uint16

//Logger 接口
type Logger interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Info(format string, a ...interface{})
	Waring(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARING
	ERROR
	FATAL
)

//Logger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

//FileLogger 文件日志结构体
type FileLogger struct {
	Level       LogLevel
	filePath    string //日志文件保存的位置
	fileName    string //日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64 // 最大文件大小,按文件大小切割
	logChan     chan *logMsg
}
type logMsg struct {
	level     LogLevel
	msg       string
	funcName   string
	fileName  string
	timestamp string
	lineNo    int
}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "waring":
		return WARING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别!")
		return UNKNOWN, err
	}
}
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARING:
		return "WARING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "UNKNOWN"
}

//获取行号等信息
func getInfo(n int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(0) //参数表示显示的行数是调用者还是他本身的位置
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
