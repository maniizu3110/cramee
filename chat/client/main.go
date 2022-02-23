// client/main.go

package main

import (
	"context"
	"cramee/chat/chat"
	"log"
	"os"

	"google.golang.org/grpc"
)

func init() {
	log.SetPrefix("Client: ")
}

func main() {
	ok := argsValidate()
	if !ok {
		return
	}
	ctx := context.Background()
	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := chat.NewChatClient(conn)
	stream, err := c.Chat(ctx)
	if err != nil {
		log.Fatal(err)
	}
	go streamRecv(stream)
	connectEstablish(stream)
}
