package main

import (
	"log"
	"net"
	"scoringservice/pb"

	"google.golang.org/grpc"
)

type loggerserver struct {
	pb.UnimplementedLoggerServer
}

func Logger() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoggerServer(s, &loggerserver{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *loggerserver) LogEvent(req *pb.LogRequest) (res *pb.LogResponse) {

}
