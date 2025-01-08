package client

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct {
	Username string
	Conn     *websocket.Conn
}

func NewClient(username string) *Client {
	return &Client{
		Username: username,
	}
}

func (c *Client) Connect(serverURL string) error {
	u := url.URL{Scheme: "ws", Host: serverURL, Path: "/ws"}

	headers := http.Header{}
	headers.Set("User-Agent", c.Username)

	fmt.Println("User-Agent: ", headers.Values("User-Agent"))

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	if err != nil {
		return err
	}

	c.Conn = conn
	return nil
}

func (c *Client) SendMessage(message string) error {
	err := c.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ReceiveMessage() (string, error) {
	_, message, err := c.Conn.ReadMessage()
	if err != nil {
		return "", err
	}

	return string(message), nil
}

func (c *Client) Close() {
	if c.Conn != nil {
		c.Conn.Close()
	}
}
