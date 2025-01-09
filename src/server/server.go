package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Gammer0909/chatgo/src/common"
	"github.com/gorilla/websocket"
)

type Server struct {
	Messages     []string
	MessagesFile *os.File
	Users        []string
	Connections  []*websocket.Conn

	Upgrader websocket.Upgrader
}

func NewServer(upgrader websocket.Upgrader) *Server {

	t := time.Now().Format("2008-01-09 15:05")
	fileName := "./data/log_" + t + ".txt"

	if err := os.MkdirAll("./data", os.ModePerm); err != nil {
		log.Println("An error occurred creating the directory, ", err)
		return nil
	}

	fileMsg, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("An error occured opening the file, ", err)
		return nil
	}

	return &Server{
		Messages:     make([]string, 0),
		MessagesFile: fileMsg,
		Users:        make([]string, 0),
		Connections:  make([]*websocket.Conn, 0),

		Upgrader: upgrader,
	}

}

func (s *Server) HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	if !common.ContainsString(s.Users, r.UserAgent()) {
		s.Users = append(s.Users, r.UserAgent())
	}

	fmt.Println(s.Connections)

	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	if !common.ContainsConnection(s.Connections, conn) {
		s.Connections = append(s.Connections, conn)
	}

	defer func() {
		s.removeUser(r.UserAgent())
		s.removeConnection(conn)
		conn.Close()
	}()

	fmt.Println("Connection established with ", r.RemoteAddr)

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Printf("Unexpected close error: %v\n", err)
			} else {
				log.Println("Connection closed: ", err)
			}
			break
		}

		formattedMsg := fmt.Sprintf("[%s]: %s", r.UserAgent(), msg)
		_, err = s.MessagesFile.WriteString(string(formattedMsg))
		if err != nil {
			log.Println("Failed to write message to file:", err)
			break
		}

		s.broadcastMessage(messageType, msg, conn, r.UserAgent())

	}

}

func (s *Server) removeUser(userAgent string) {
	for i, user := range s.Users {
		if user == userAgent {
			s.Users = append(s.Users[:i], s.Users[i+1:]...)
			break
		}
	}
}

func (s *Server) removeConnection(conn *websocket.Conn) {
	for i, c := range s.Connections {
		if c == conn {
			s.Connections = append(s.Connections[:i], s.Connections[i+1:]...)
			break
		}
	}
}

func (s *Server) broadcastMessage(messageType int, message []byte, sender *websocket.Conn, userName string) {

	for _, conn := range s.Connections {
		formatMessage := fmt.Sprintf("[%s]: %s", userName, message)
		if conn == sender {
			formatMessage = fmt.Sprintf("[YOU]: %s", message)
			if err := conn.WriteMessage(messageType, []byte(formatMessage)); err != nil {
				log.Printf("Error broadcasting message: %s", err)
			}
			continue
		}
		if err := conn.WriteMessage(messageType, []byte(formatMessage)); err != nil {
			log.Printf("Error broadcasting message: %s", err)
		}
	}
}

func (s *Server) Close() {
	s.MessagesFile.Close()
}
