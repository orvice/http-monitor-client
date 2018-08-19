package main

import (
	"context"
	"net/http"
	"time"

	pb "github.com/orvice/http-monitor-client/proto"
)

func NewServer() pb.HttpMonitorSrvServer {
	return new(Server)
}

type Server struct {
}

func (s *Server) Send(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	res := new(pb.Response)
	statusCode, err := httpGet(in.Uri, in.Timeout)
	if err != nil {
		logger.Infof("req %v error %v", in, err)
		res.Result = pb.ResultCode_Fail
		return res, nil
	}
	logger.Infof("req %v statusCode %d", in, statusCode)
	res.StatusCode = int32(statusCode)
	res.Result = pb.ResultCode_Success
	return res, nil
}

func httpGet(uri string, timeout int64) (int, error) {
	tm := time.Second * 10
	if timeout > 0 {
		tm = time.Millisecond * time.Duration(timeout)
	}
	client := http.Client{
		Timeout: tm,
	}
	resp, err := client.Get(uri)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}
