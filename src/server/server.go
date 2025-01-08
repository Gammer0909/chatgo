package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type Server struct {
	Messages     []string
	MessagesFile *os.File
	Users        []string

	Upgrader websocket.Upgrader
}

func NewServer(upgrader websocket.Upgrader) *Server {

	fileMsg, err := os.OpenFile("chat_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("An error occured opening the file, ", err)
		return nil
	}

	return &Server{
		Messages:     make([]string, 0),
		MessagesFile: fileMsg,
		Users:        make([]string, 0),

		Upgrader: upgrader,
	}

}

func (s *Server) HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	if !Contains(s.Users, r.UserAgent()) {
		s.Users = append(s.Users, r.UserAgent())
	}

	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connection established with ", r.RemoteAddr)

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("An error occurred: ", err)
			break
		}

		fmt.Printf("Client sent: %s\n", msg)
		_, err = s.MessagesFile.WriteString(string(msg) + "\n")
		if err != nil {
			log.Println("Failed to write message to file:", err)
			break
		}

		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Printf("Error writing message: %s\n", err)
			break
		}

	}

	s.Users = s.Users[:len(s.Users)-1]
}

func Contains(slice []string, target string) bool {

	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return true
		}
	}
	return false

}

func (s *Server) Close() {
	s.MessagesFile.Close()
}
