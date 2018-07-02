package main

import (
	"net"
	"os"

	pb "github.com/orvice/http-monitor-client/proto"
	"github.com/orvice/utils/env"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port   string
	logger *logrus.Logger
)

func main() {
	logger = logrus.New()
	port = env.Get("PORT")
	err := serveGrpc()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

func serveGrpc() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterHttpMonitorSrvServer(s, NewServer())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
