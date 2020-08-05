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

func (s *notificationserver) ShouldNotify(ctx context.Context, notificationRequest *pb.ExposureNotificationRequest) (res *pb.ExposureNotificationResponse, error) {
	// Insert Business Logic Here
	l := pb.LogRequest{}
	err := shipToLogger(&l)
	// Retry logic
	if err != nil {
		log.Println()
	}
	return
}

func shipToLogger(req *pb.LogRequest) (res *pb.LogResponse) {
	conn, err := grpc.Dial(loggerAddress, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewLoggerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.LogEvent(ctx, req)
	if err != nil {
		log.Fatalf("could not log: %v", err)
	}
}
