package log

import (
	lg4 "github.com/alecthomas/log4go"
)

/*全局的logger,可配置*/
var rootLogger = make(lg4.Logger)

/*只需要用到四个日志级别即可*/
const (
	DEBUG    =  lg4.DEBUG
	INFO     =  lg4.INFO
	WARN     =  lg4.WARNING
	ERROR    =  lg4.ERROR
)

const (
	STDOUT   = "stdout"
	NONE     = "none"
)

/*日志输出到何处,日志登记设置,如果不设置则使用lg4默认配置*/
func LogTo(target string,level lg4.Level){

	var writer  lg4.LogWriter  = nil

	switch target {
	case STDOUT:
		//输出到屏幕
		writer  = lg4.NewConsoleLogWriter()
	case NONE :
		//不配置日志
	default:
		writer  = lg4.NewFileLogWriter(target,true)
	}//switch
	if writer != nil{
		rootLogger.AddFilter("log",level,writer)
	}//if
}

type Logger interface {
	AddPrefix(string)
	ClearPrefixes()
	Debug(arg0 string,args ...interface{})
	Info(arg0 string,args ...interface{})
	Warn(arg0 string,args ...interface{}) error
	Error(arg0 string,args ...interface{}) error
}

/*带自定义前缀的logger*/
type PrefixedLogger struct {
	prefix        string
	logger        *lg4.Logger // 封装lg4的logger
}

func (pf *PrefixedLogger) AddPrefix(prefix string){
    pf.prefix += "[" + prefix + "]"
}

func (pf *PrefixedLogger) ClearPrefixes(){
	pf.prefix  = ""
}

func (pf *PrefixedLogger) pfx(fmtStr string) string{
	return pf.prefix + " " + fmtStr
}
func (pf *PrefixedLogger) Debug(arg0 string,args ...interface{}){
	pf.logger.Debug(pf.pfx(arg0),args...)
}

func (pf *PrefixedLogger) Info(arg0 string,args ...interface{}){
	pf.logger.Info(pf.pfx(arg0),args...)
}

func (pf *PrefixedLogger) Warn(arg0 string,args ...interface{}) error{
    return pf.logger.Warn(pf.pfx(arg0),args...)
}

func (pf *PrefixedLogger) Error(arg0 string,args ...interface{}) error{
    return pf.logger.Error(pf.pfx(arg0),args...)
}

func NewPrefixedLogger(prefixes ...string) Logger{
    logger := &PrefixedLogger{"",&rootLogger}
    for _ , prefix := range prefixes{
    	logger.AddPrefix(prefix)
	}//for
	return logger
}//

/*不带prefix*/
func Debug(arg0 string,args ...interface{}){
	rootLogger.Debug(arg0,args...)
}

func Info(arg0 string,args ...interface{}){
	rootLogger.Info(arg0,args...)
}

func Warn(arg0 string,args ...interface{}) error{
    return rootLogger.Warn(arg0,args...)
}

func Error(arg0 string,args ...interface{}) error{
    return rootLogger.Error(arg0,args...)
}

func Close(){
	rootLogger.Close()
}