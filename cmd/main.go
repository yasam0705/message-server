package main

import (
	"log"
	"micro-grpc/pkg/proto"
	"micro-grpc/pkg/serv"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := serv.MessagesServer{}

	proto.RegisterMessageServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
