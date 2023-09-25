package main

import (
	"net/http"
	"score-socket/internal/controller/rest"
	"score-socket/internal/controller/socket"
)

func main() {

	http.HandleFunc("/list", rest.HandlerGet)

	http.HandleFunc("/ws/math", socket.Handler)

	http.ListenAndServe(":8080", nil)
}
