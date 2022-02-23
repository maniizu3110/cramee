// server/server.go

package main

import (
	"cramee/chat/chat"
	"fmt"
	"io"
	"sync"
)

type Connection struct {
	conn chat.Chat_ChatServer
	send chan *chat.ChatMessage
	quit chan struct{}
}

func NewConnection(conn chat.Chat_ChatServer) *Connection {
	c := &Connection{
		conn: conn,
		send: make(chan *chat.ChatMessage),
		quit: make(chan struct{}),
	}
	go c.start()
	return c
}

func (c *Connection) Close() error {
	close(c.quit)
	close(c.send)
	return nil
}

func (c *Connection) Send(msg *chat.ChatMessage) {
	defer func() {
		recover()
	}()
	c.send <- msg
}

func (c *Connection) start() {
	running := true
	for running {
		select {
		case msg := <-c.send:
			c.conn.Send(msg)
		case <-c.quit:
			running = false
		}
	}
}

func (c *Connection) GetMessage(broadcast chan<- *chat.ChatMessage) error {
	for {
		msg, err := c.conn.Recv()

		if err == io.EOF {
			c.Close()
			return nil
		} else if err != nil {
			c.Close()
			return err
		}

		go func(msg *chat.ChatMessage) {
			fmt.Println("block?")
			select {
			case broadcast <- msg:
			case <-c.quit:
			}
		}(msg)
	}
}

type ChatServer struct {
	broadcast   chan *chat.ChatMessage
	quit        chan struct{}
	connections []*Connection
	connLock    sync.Mutex
}

func NewChatServer() *ChatServer {
	srv := &ChatServer{
		broadcast: make(chan *chat.ChatMessage),
		quit:      make(chan struct{}),
	}
	go srv.start()
	return srv
}

func (c *ChatServer) Close() error {
	close(c.quit)
	return nil
}

func (c *ChatServer) start() {
	running := true
	for running {
		select {
		case msg := <-c.broadcast:
			c.connLock.Lock()
			for _, v := range c.connections {
				go v.Send(msg)
			}
			c.connLock.Unlock()
		case <-c.quit:
			running = false
		}
	}
}

/*
chat.pb.go
type ChatServer interface {
	Chat(Chat_ChatServer) error
}
Chatを実装しておかないとmain.goのRegisterChatServerでサーバーを登録できない
*/
func (c *ChatServer) Chat(stream chat.Chat_ChatServer) error {
	conn := NewConnection(stream)
	c.connLock.Lock()
	c.connections = append(c.connections, conn)
	c.connLock.Unlock()
	err := conn.GetMessage(c.broadcast)
	if err != nil {
		return err
	}
	c.connLock.Lock()
	for i, v := range c.connections {
		if v == conn {
			c.connections = append(c.connections[:i], c.connections[i+1:]...)
		}
	}
	c.connLock.Unlock()
	return nil
}
