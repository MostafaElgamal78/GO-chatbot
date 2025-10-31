package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

type ChatServer struct {
	messages []string
	mu       sync.Mutex
}

func (s *ChatServer) SendMessage(msg string, reply *[]string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages = append(s.messages, msg)
	*reply = append([]string(nil), s.messages...)
	return nil
}

func (s *ChatServer) GetHistory(_ struct{}, reply *[]string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	*reply = append([]string(nil), s.messages...)
	return nil
}

func main() {
	server := new(ChatServer)
	err := rpc.Register(server)
	if err != nil {
		log.Fatal("Error registering RPC server:", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	fmt.Println("🚀 Chat server running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
