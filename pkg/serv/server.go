package serv

import (
	"context"
	"fmt"
	"micro-grpc/pkg/proto"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"
)

// var PriorityTable = map[string][]*http.Request{
// 	"low":    make([]*http.Request, 0),
// 	"medium": make([]*http.Request, 0),
// 	"high":   make([]*http.Request, 0),
// }

var (
	LowList    = make([]*http.Request, 0)
	MediumList = make([]*http.Request, 0)
	HighList   = make([]*http.Request, 0)
)

var URL = "https://api.telegram.org/bot1822246375:AAFBs9rUJ1wHJpweTlFHSOPuVXUfJQoKpTc/"

type MessagesServer struct {
	proto.UnimplementedMessageServiceServer
}

func (ms MessagesServer) SendChannel(ctx context.Context, m *proto.Mes) (*emptypb.Empty, error) {
	client := http.Client{}
	fmt.Println("Channel")
	err := createRequest(&client, URL, "-1001652337843", m.GetText(), m.GetPriority())

	return &emptypb.Empty{}, err
}

func (ms MessagesServer) SendGroupChat(ctx context.Context, m *proto.Mes) (*emptypb.Empty, error) {
	client := http.Client{}
	fmt.Println("Group")
	err := createRequest(&client, URL, "-1001317290790", m.GetText(), m.GetPriority())

	return &emptypb.Empty{}, err
}

func createRequest(cl *http.Client, url, chatId, text, priority string) error {
	post_url := fmt.Sprintf("%s%schat_id=%s&text=%s", url, "sendMessage?", chatId, text)

	req, err := http.NewRequest("POST", post_url, nil)
	if err != nil {
		return err
	}

	if priority == "high" {
		HighList = append(HighList, req)
	} else if priority == "medium" {
		MediumList = append(HighList, req)
	} else if priority == "low" {
		LowList = append(LowList, req)
	}

	return err
}
