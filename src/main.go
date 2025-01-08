package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Gammer0909/chatgo/src/client"
	"github.com/Gammer0909/chatgo/src/server"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Expected Usage: chatgo [server|client]")
		return
	}
	if os.Args[1] == "server" {
		server := server.NewServer(upgrader)
		http.HandleFunc("/ws", server.HandleWebSocket)
		log.Fatal(http.ListenAndServe(":8080", nil))
		server.Close()
	} else if os.Args[1] == "client" {
		client := client.NewClient(os.Args[2])
		err := client.Connect("localhost:8080")
		if err != nil {
			fmt.Println("An error occurred: ", err)
			return
		}

		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			sent, _ := reader.ReadString('\n')
			if sent == "quit" {
				break
			}
			err := client.SendMessage(sent)
			if err != nil {
				fmt.Println("An error occurred: ", err)
				return
			}

			response, err := client.ReceiveMessage()
			if err != nil {
				fmt.Println("An error occurred: ", err)
				return
			}
			fmt.Println("[SERVER]: ", response)

		}

		client.Close()
		return
	}

	fmt.Println("Expected Usage: chatgo [server|client]")
}
