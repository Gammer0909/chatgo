package common

import "github.com/gorilla/websocket"

func ContainsConnection(slice []*websocket.Conn, target *websocket.Conn) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return true
		}
	}
	return false
}

func ContainsString(slice []string, target string) bool {

	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return true
		}
	}
	return false

}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
