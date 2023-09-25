package socket

import (
	"log"
	"net/http"
	"score-socket/internal/service"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func Handler(resp http.ResponseWriter, req *http.Request) {
	// Upgrade http request to socket connection
	cSocket, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer cSocket.Close()

	matchID := req.URL.Query().Get("id")

	id, _ := strconv.Atoi(matchID)

	service.Add(id, cSocket)

}
