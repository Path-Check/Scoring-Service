package main

import (
	"context"
	"log"
	"net"
	"scoringservice/pb"

	"google.golang.org/grpc"
)

type LoggerServer struct {
	pb.UnimplementedLoggerServer
}

// Logger starts a server which then ships the logs to a storage bucket
func Logger() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoggerServer(s, &LoggerServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// LogEvent saves the event to a json file then stores it into a storage bucket
func (s *LoggerServer) LogEvent(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	res := &pb.LogResponse{}
	return res, nil
}
