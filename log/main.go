package main

import (
	"context"
	"log"
	"os"
	"time"

	"logRPC/pb"

	"google.golang.org/grpc"
)

const (
	address = "localhost:80"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewLoggerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	eventtimestamp := time.Now().String()

	pid := int32(os.Getpid())

	hostname, _ := os.Hostname()

	e := &pb.Event{
		Method:     "Test Method",
		RequestURI: "/test",
		RouteName:  "Test",
		Timestamp:  eventtimestamp,
		Pid:        pid,
		Hostname:   hostname,
	}

	r, err := c.CreateEvent(ctx, e)
	if err != nil {
		log.Fatalf("could not log: %v", err)
	}

	log.Printf("Event Logged: %t, at Time: %s", r.GetEventLogged(), r.GetTimestamp())
}
