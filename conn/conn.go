package conn

import (
	"fmt"
	"io"
	"net"
	"rinet_go/log"
	"rinet_go/util"
	"sync"
)

type LoggedConn struct {
	log.Logger
	id         string
	net.Conn
}

func (lg *LoggedConn) Id() string{
	return fmt.Sprintf("%s:%s",lg.id,lg.Conn.RemoteAddr())
}

type Proxy struct {
	log.Logger
	proxyUrl   string
	bindUrl    string
	net.Listener
}

func (p *Proxy) HandleAccept(accepts *int,exit chan<- bool){
	for {
		rawConn , err := p.Accept()
		if err != nil{
			p.Warn("Proxy %s to %s closed",p.bindUrl,p.proxyUrl)
			break
		}//if
		*accepts += 1
		p.Debug("New connection from %s",rawConn.RemoteAddr())
		wrappedConn := WrapConn(rawConn)
		go p.HandleConnection(wrappedConn)
	}	//for
	exit <- true
}


func WrapConn(conn net.Conn) *LoggedConn{
	if conn == nil{
		return nil
	}//
	id  := util.RandId(5)
	return &LoggedConn{log.NewPrefixedLogger("Conn"),
		id,conn}
}

func Dial(proxyUrl string,retries int) (rawConn net.Conn , err error){
	for i:= 0 ;i<retries ;i++{
		rawConn , err = net.Dial("tcp",proxyUrl)
		if err == nil{
			return
		}//if
	}//for
	return
} //Dial


func (p *Proxy) HandleConnection(_from *LoggedConn){
	//dial
	rawConn , err := Dial(p.proxyUrl,2)
	if err != nil {
		p.Warn("New proxy to %s failed after %d tries",p.proxyUrl,2)
		_from.Close()
		return
	}//
	p.Debug("New proxy to %s",p.proxyUrl)
	_to  := WrapConn(rawConn)
	Join(_from,_to) //开始互相传输数据
}//


func Join(_from *LoggedConn,_to *LoggedConn) (int64 , int64){
	var _fromBytes   int64
	var _toBytes     int64
	var err          error
	var wait         sync.WaitGroup
	defer _from.Close()
	defer _to.Close()

	pipe        := func(from *LoggedConn,to *LoggedConn,copiedBytes *int64) {
		defer wait.Done()
		*copiedBytes, err = io.Copy(to,from)
		if err != nil{
			from.Warn("Copied %d bytes from %s before failing with error %v",*copiedBytes,from.Id(),err)
			return
		}//if
		from.Info("Copied %d bytes from %s",*copiedBytes,from.Id())
	}//pipe

	wait.Add(2)
	go pipe(_from,_to,&_fromBytes)
	go pipe (_to,_from,&_toBytes)

	wait.Wait()
	return _fromBytes,_toBytes
}

func Listen(bindHost string,bindPort int,proxyHost string,proxyPort int) (*Proxy,error){
	bindUrl  := fmt.Sprintf("%v:%v",bindHost,bindPort)
	proxyUrl := fmt.Sprintf("%v:%v",proxyHost,proxyPort)
	log.Debug("bindUrl %s proxyUrl %s",bindUrl,proxyUrl)
	listener , err := net.Listen("tcp",bindUrl)
	if err != nil{
		log.Warn("Listen %s filed with error %v",bindUrl,err)
		return nil , err
	}//
	logger   := log.NewPrefixedLogger("Proxy")
	proxy    := Proxy{logger,proxyUrl,bindUrl,listener}
	return &proxy , nil
}
