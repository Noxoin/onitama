package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	gameRooms map[string]*GameRoom
)

func init() {
	http.HandleFunc("/socket/", clientHandler)
	tmpls["game"] = template.Must(template.ParseFiles(
    "client/templates/game.html", "client/templates/base.html"))
	gameRooms = make(map[string]*GameRoom)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	gameId := r.URL.Path[1:]
	if gameRooms[gameId] == nil {
		http.NotFound(w, r)
		return
	}
	tmpls["game"].ExecuteTemplate(w, "base", gameId)
}

func createGameHandler(w http.ResponseWriter, r *http.Request) {
	randId := GenerateRandomId()
	for gameRooms[randId] != nil {
		randId = GenerateRandomId()
	}
	gameRooms[randId] = NewGameRoom()
	http.Redirect(w, r, r.URL.Host + "/" + randId, http.StatusFound)
}

func clientHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, []byte("Welcome to sockets")); err != nil {
		return
	}
}
