package middleware

import (
	"errors"
	"github.com/soheilhy/cmux"
	"net"
	"thingue-launcher/common/logger"
	"thingue-launcher/common/provider"
)

var listener net.Listener
var grpcL net.Listener
var httpL net.Listener
var cmx cmux.CMux

func ListenTcp() error {
	lis, err := net.Listen("tcp", provider.AppConfig.LocalServer.BindAddr)
	if err != nil {
		return err
	}
	listener = lis
	return nil
}

func CloseListenTcp() error {
	if listener == nil {
		return errors.New("端口未在监听")
	}
	_ = listener.Close()
	listener = nil
	return nil
}

// CreateMux 创建多路复用器
func CreateMux() {
	cmx = cmux.New(listener)
	//grpcL = cmx.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	grpcL = cmx.Match(cmux.HTTP2())
	httpL = cmx.Match(cmux.HTTP1Fast())
}

func ServeMux() {
	err := cmx.Serve()
	logger.Zap.Error(err)
}

func CloseMux() {
	cmx.Close()
}
