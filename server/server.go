package server

import (
	"os"
	"rinet_go/conn"
	"rinet_go/log"
	"strconv"
)

func Main(){
	var bindHost   string
	var bindPort   int
	var proxyHost  string
	var proxyPort  int
	var err        error
	log.LogTo(log.STDOUT,log.DEBUG)
	defer log.Close()
	if len(os.Args) != 4 && len(os.Args) != 5 {
		log.Warn("Invalid args")
		return
	}//

	if len(os.Args) == 5 {
		bindHost       = os.Args[1]
		bindPort , err = strconv.Atoi(os.Args[2])
		proxyHost      = os.Args[3]
		proxyPort, err = strconv.Atoi(os.Args[4])
	}else{
		bindHost       = ""
		bindPort , err = strconv.Atoi(os.Args[1])
		proxyHost      = os.Args[2]
		proxyPort, err = strconv.Atoi(os.Args[3])
	}//else

	if err!= nil{
		log.Debug("Input error %v",err)
		return
	}//if

	proxy ,err := conn.Listen(bindHost,bindPort,proxyHost,proxyPort)
	if err != nil{
		log.Debug("Listen error %v",err)
		return
	}//if
	accepts := 0
	exit  := make(chan bool)
	go proxy.HandleAccept(&accepts,exit)
	<-exit
	log.Info("Handle %d connection before stopped",accepts)
}