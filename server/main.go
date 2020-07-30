package main

import (
	"context"
	"encoding/json"
	"log"
	"logRPC/pb"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":80"
)

type server struct {
	pb.UnimplementedLoggerServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoggerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CreateEvent(ctx context.Context, event *pb.Event) (*pb.EventResponse, error) {
	method := event.GetMethod()
	timestamp := event.GetTimestamp()
	saveToFile(ctx, event)
	log.Printf("Event Logged: %s, at timestamp: %s", method, timestamp)
	eventtimestamp := time.Now().String()
	return &pb.EventResponse{EventLogged: true, Timestamp: eventtimestamp}, nil
}

// SaveToFile saves to a json file
func saveToFile(ctx context.Context, e *pb.Event) (bool, error) {
	m, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("failed to marshal: %v", err)
		return false, err
	}
	// If the file doesn't exist, create it, or append to the file
	l, err := os.OpenFile("logtest.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	if _, err := l.Write(m); err != nil {
		log.Fatal(err)
		return false, err
	}
	if err := l.Close(); err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}
