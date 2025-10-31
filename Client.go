package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("❌ Could not connect to server:", err)
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("✅ Connected to chat server. Type your message below.")
	fmt.Println("Type 'exit' to quit.\n")

	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("👋 Exiting chat.")
			break
		}

		var reply []string
		err = client.Call("ChatServer.SendMessage", text, &reply)
		if err != nil {
			log.Println("⚠️ Server connection failed:", err)
			break
		}

		fmt.Println("\n💬 Chat History:")
		for _, msg := range reply {
			fmt.Println("-", msg)
		}
		fmt.Println()
	}
}
