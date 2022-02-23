// client/client.go

package main

import (
	"bufio"
	"chat/chat"
	"fmt"
	"io"
	"log"
	"os"
)

var waitc = make(chan struct{})

func argsValidate() bool {
	if len(os.Args) != 3 {
		fmt.Println("第一引数: URL, 第二引数: ユーザー名が必要です。")
		return false
	}
	return true
}

func streamRecv(stream chat.Chat_ChatClient) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			close(waitc)
			return
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg.User + ": " + msg.Message)
	}
}

func connectEstablish(stream chat.Chat_ChatClient) {
	fmt.Println("connection was established. Please input \"quit\" or press \"ctrl+c\" to stop.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "quit" {
			err := stream.CloseSend()
			if err != nil {
				log.Fatal(err)
			}
			break
		}
		err := stream.Send(&chat.ChatMessage{
			User:    os.Args[2],
			Message: msg,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	<-waitc
}
