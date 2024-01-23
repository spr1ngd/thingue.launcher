package controller

import (
	"fmt"
	"thingue-launcher/server/middleware"
)

type application struct {
	HttpCloseChanel   chan error
	GrpcStopChanel    chan error
	HttpServerRunning bool
	GrpcServerRunning bool
}

var Application = application{
	HttpCloseChanel: make(chan error),
	GrpcStopChanel:  make(chan error),
}

func (a *application) Start() error {
	middleware.InitGorm()
	err := middleware.ListenTcp()
	if err != nil {
		return err
	}
	middleware.CreateMux()
	middleware.CreateGRPCServer()
	middleware.CreateGinRouter()
	middleware.CreateHttpServer()
	go func() {
		a.GrpcServerRunning = true
		grpcServerStopErr := middleware.RunGRPCServer()
		a.GrpcServerRunning = false
		fmt.Println("Grpc关闭", grpcServerStopErr)
		a.GrpcStopChanel <- grpcServerStopErr
	}()
	go func() {
		a.HttpServerRunning = true
		httpServerClosedErr := middleware.RunHttpServer()
		a.HttpServerRunning = false
		fmt.Println("Http关闭", httpServerClosedErr)
		a.HttpCloseChanel <- httpServerClosedErr
	}()
	go middleware.ServeMux()
	return nil
}

func (a *application) Stop() error {
	middleware.StopGRPCServer()
	middleware.StopHttpServer()
	middleware.CloseMux()
	return middleware.CloseListenTcp()
}

func (a *application) StartMain() {
	middleware.InitGorm()
	err := middleware.ListenTcp()
	if err != nil {
		panic(err)
	}
	middleware.CreateMux()
	middleware.CreateGRPCServer()
	middleware.CreateGinRouter()
	middleware.CreateHttpServer()
	go func() {
		grpcServerStopErr := middleware.RunGRPCServer()
		fmt.Println("Grpc关闭", grpcServerStopErr)
	}()
	go func() {
		httpServerClosedErr := middleware.RunHttpServer()
		fmt.Println("Http关闭", httpServerClosedErr)
	}()
	middleware.ServeMux()
}
