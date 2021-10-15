package main

import (
	"log"
	"micro-grpc/pkg/proto"
	"micro-grpc/pkg/serv"
	"net"
	"net/http"
	"time"

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
	go DoRequest(&http.Client{})
	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}

}

func DoRequest(cl *http.Client) {
	for {
		for i, v := range serv.HighList {
			cl.Do(v)
			serv.HighList = append(serv.HighList[:i], serv.HighList[i+1:]...)
			time.Sleep(5 * time.Second)
		}
		for i, v := range serv.MediumList {
			cl.Do(v)
			serv.MediumList = append(serv.MediumList[:i], serv.MediumList[i+1:]...)
			time.Sleep(5 * time.Second)
		}
		for i, v := range serv.LowList {
			cl.Do(v)
			serv.LowList = append(serv.LowList[:i], serv.LowList[i+1:]...)
			time.Sleep(5 * time.Second)
		}
	}
}
