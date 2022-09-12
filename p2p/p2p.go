package p2p

import (
	"fmt"
	"net/http"

	"github.com/WoodoCoin/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	initPeer(conn, "x", "x")
}

func AddPeer(address, port string) {
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws", address, port), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}
