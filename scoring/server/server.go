package main

import (
	"context"
	"log"
	"net"
	"scoringservice/pb"
	"time"

	"google.golang.org/grpc"
)

const (
	port          = ":80"
	loggerAddress = "localhost:81"
)

type notificationserver struct {
	pb.UnimplementedNotificationServer
}

func main() {
	fmt.Println("In scoring server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotificationServer(s, &notificationserver{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *notificationserver) ShouldNotify(ctx context.Context, notificationRequest *pb.ExposureNotificationRequest) (*pb.ExposureNotificationResponse, error) {
	// Insert Business Logic Here
	l := pb.LogRequest{}
	err := shipToLogger(&l)
	// Retry logic
	if err != nil {
		log.Println()
	}
	res := &pb.ExposureNotificationResponse{}
	return res, nil
}