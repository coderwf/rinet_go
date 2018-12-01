package log

import (
	"fmt"
	log "github.com/jeanphorn/log4go"
)

var rootLogger  = make(log.Logger)

var (
	Stdout  = "Stdout"
	None    = "None"
)

type Logger interface {
	AddLogPrefix(string)
	ClearLogPrefixes()
    Debug(string,...interface{})
	Info(string,...interface{})
	Warn(string,...interface{}) error
	Error(string,...interface{}) error
}

//将日志记录到某处
func LogTo(target string,levelName string){
	var writer log.LogWriter = nil
    switch target{
	case Stdout:
		writer = log.NewConsoleLogWriter()
	case None :
        //不采取操作
	default :
		writer = log.NewFileLogWriter(target,true,false)
	}

	//不是None
	if writer != nil {

	}//
}

type PrefixedLogger struct {
	*log.Logger
	prefix string
}

func NewPrefixLogger(prefixes ...string) *PrefixedLogger{
	var logger = &PrefixedLogger{&rootLogger,""}

	for _ , prefix := range prefixes{
        logger.AddLogPrefix(prefix)
	}//for
	return logger
}

func (pfl *PrefixedLogger) AddLogPrefix(prefix string){
	if len(pfl.prefix) >0 {
		pfl.prefix += "-"
	}//
	pfl.prefix += "[" +prefix + "]"
}

func (pfl *PrefixedLogger) ClearLogPrefixes(){
	pfl.prefix  = ""
}

func (pfl *PrefixedLogger) pfx(fmtstr string) interface{}{
	return fmt.Sprintf("%s %s",pfl.prefix,fmtstr)
}

func (pfl *PrefixedLogger) Debug(arg0 string,args ...interface{}){
	pfl.Logger.Debug(pfl.pfx(arg0),args...)
}

func (pfl *PrefixedLogger) Error(arg0 string,args ...interface{}) error {
	return pfl.Logger.Error(pfl.pfx(arg0),args...)
}

func (pfl *PrefixedLogger) Warn(arg0 string,args ...interface{}) error {
	return pfl.Logger.Warn(pfl.pfx(arg0),args...)
}

func (pfl *PrefixedLogger) Info(arg0 string,args ...interface{}){
	pfl.Logger.Info(pfl.pfx(arg0),args...)
}

func (pfl *PrefixedLogger) Trace(arg0 string,args ...interface{}){
	pfl.Logger.Trace(pfl.pfx(arg0),args...)
}

func (pfl *PrefixedLogger) Critical(arg0 string,args ...interface{}) error{
	return pfl.Logger.Critical(pfl.pfx(arg0),args...)
}

func (pfl *PrefixedLogger) Fine(arg0 string , args ...interface{}){
	pfl.Logger.Fine(pfl.pfx(arg0),args...)
}

func (pfl *PrefixedLogger) Finest(arg0 string , args ...interface{}){
	pfl.Logger.Finest(pfl.pfx(arg0),args...)
}


/*最简单的使用方式,正式代码中不要使用*/
func Debug(arg0 string , args ...interface{}){
	rootLogger.Debug(arg0,args...)
}

func Info(arg0 string , args ...interface{}){
	rootLogger.Info(arg0 , args...)
}

func Error(arg0 string ,args ...interface{}) error{
	return rootLogger.Error(arg0,args...)
}

func Warn(arg0 string , args ...interface{}) error{
	return rootLogger.Warn(arg0 , args...)
}

func Critical(arg0 string , args ...interface{}) error{
	return rootLogger.Critical(arg0 , args...)
}

func Trace(arg0 string , args ...interface{}) {
	rootLogger.Trace(arg0 , args...)
}
