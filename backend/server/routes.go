package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"leet-guys/gamelogic"
)

func Serve() {
	mux := routes()
	log.Println("listening and serving on 0.0.0.0:8080")
	http.ListenAndServe("0.0.0.0:8080", mux)
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("../frontend/dist/"))
	mux.Handle("/", fileServer)

	hub := gamelogic.NewHub()
	mux.HandleFunc("/ws", hub.ServeWs)

	return mux
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
