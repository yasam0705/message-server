package serv

import (
	"context"
	"fmt"
	"micro-grpc/pkg/proto"
	"net/http"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

var priorityTable = map[string][]*http.Request{
	"low":    make([]*http.Request, 0),
	"medium": make([]*http.Request, 0),
	"high":   make([]*http.Request, 0),
}

var URL = "https://api.telegram.org/bot1822246375:AAFBs9rUJ1wHJpweTlFHSOPuVXUfJQoKpTc/"

type MessagesServer struct {
	proto.UnimplementedMessageServiceServer
}

func (ms MessagesServer) SendChannel(ctx context.Context, m *proto.Mes) (*emptypb.Empty, error) {
	client := http.Client{}
	fmt.Println("Channel")
	err := createRequest(&client, URL, "-1001652337843", m.GetText(), m.GetPriority())
	doRequest(&client)

	return &emptypb.Empty{}, err
}

func (ms MessagesServer) SendGroupChat(ctx context.Context, m *proto.Mes) (*emptypb.Empty, error) {
	client := http.Client{}
	fmt.Println("Group")
	err := createRequest(&client, URL, "-1001317290790", m.GetText(), m.GetPriority())
	doRequest(&client)

	return &emptypb.Empty{}, err
}

func createRequest(cl *http.Client, url, chatId, text, priority string) error {
	post_url := fmt.Sprintf("%s%schat_id=%s&text=%s", url, "sendMessage?", chatId, text)

	req, err := http.NewRequest("POST", post_url, nil)
	if err != nil {
		return err
	}

	priorityTable[priority] = append(priorityTable[priority], req)

	return err
}

func doRequest(cl *http.Client) {
	time.AfterFunc(5*time.Second, func() {
		for i, v := range priorityTable["high"] {
			cl.Do(v)
			priorityTable["high"] = append(priorityTable["high"][:i], priorityTable["high"][i+1:]...)
			return
		}
		for i, v := range priorityTable["medium"] {
			cl.Do(v)
			priorityTable["medium"] = append(priorityTable["medium"][:i], priorityTable["medium"][i+1:]...)
			return
		}
		for i, v := range priorityTable["low"] {
			cl.Do(v)
			priorityTable["low"] = append(priorityTable["low"][:i], priorityTable["low"][i+1:]...)
			return
		}
	})
}
