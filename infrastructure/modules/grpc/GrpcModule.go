package grpcModule

import (
	"go-boilerplate/config/environment"
	grpcServer "go-boilerplate/infrastructure/modules/grpc/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Module struct {
	// Server
	server   *grpc.Server
	listener net.Listener
}

func NewGrpcModule() *Module {
	return &Module{}
}

func (handler *Module) Configure() {
	var config = environment.Vars.Grpc

	// Server
	server := grpc.NewServer()
	grpcServer.UserServer.ConfigureServer(server)
	lis, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		panic("[GRPC] Listen: " + err.Error())
	}

	*handler = Module{server, lis}
}

func (handler *Module) Start() error {
	var config = environment.Vars.Grpc

	// Server
	log.Printf("[GRPC] listening on %s\n", config.Port)
	return handler.server.Serve(handler.listener)
}

func (handler *Module) Finish() {
	// Server
	handler.server.Stop()

	log.Printf("[GRPC] Finished\n")
}
