// server/main.go

package main

import (
	"chat/chat"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	chatSrv := NewChatServer()
	chat.RegisterChatServer(server, chatSrv)
	fmt.Println("Boot chat server on port 8080.")
	log.Fatal(server.Serve(lst))
}
