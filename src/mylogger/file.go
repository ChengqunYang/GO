package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志相关代码

//newFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, 50000),
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

//关闭日志文件
func (f *FileLogger) close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

//根据指定的日志文件路径和文件名称打开日志文件
func (f *FileLogger) initFile() (err error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open err log file failed,err:%v\n", err)
		return err
	}
	//日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	//开启5个后台的goroutine来写日志
	/*for i := 0; i < 5; i++ {
		go f.writeLogBackground()
	}*/
	go f.writeLogBackground()
	return
}

//判断日志级别,决定是否要记录该日志
func (f *FileLogger) enable(loglevel LogLevel) bool {
	return loglevel >= f.Level
}

//检查日志大小,是否要分页
func (f *FileLogger) checkSize(file *os.File) bool {
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("get file info failed,err:", err)
		return false
	}
	//如果当前文件大小大于等于日志文件的最大值,就返回true
	size := stat.Size()
	return size >= f.maxFileSize

}

//切割日志文件的具体步骤
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	/*
		1. 关闭当前日志文件
		2. 备份一下rename    rename xx.log -> xx.log.bak20200210
		3. 打开一个新的日志文件
		4. 将打开的新日志文件对象赋给f.fileObj
	*/
	nowStr := time.Now().Format("2006010215040500")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("ge file info failed,err:", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name()) //拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)

	file.Close()
	os.Rename(logName, newLogName)
	//打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open new log file failed,err:", err)
		return nil, err
	}
	return fileObj, nil
}

//后台写日志
func (f *FileLogger) writeLogBackground() {
	for true {
		if f.checkSize(f.fileObj) {
			//需要切割日志文件
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan:
			//把日志先组合好
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.lineNo, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.level >= ERROR {
				//如果要记录的日志大于等于ERROR级别,我还要在err日志文件中再记录一遍
				if f.checkSize(f.errFileObj) {
					//需要切割日志文件
					newFile, err := f.splitFile(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			//取不到日志先休息500毫秒
			time.Sleep(time.Microsecond*500)
		}
	}
}

//写日志
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		//先把日志发送到通道中
		logTmp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			lineNo:    lineNo,
		}
		select {
		case f.logChan <- logTmp:
		default:
			//如果通道已经满了,放不进去,那就把日志扔掉,保证不阻塞
		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Waring(format string, a ...interface{}) {
	f.log(WARING, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
