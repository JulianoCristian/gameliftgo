package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/marchinram/gameliftgo"
)

var (
	srv      = &http.Server{Addr: ":8080"}
	addr     = flag.String("addr", "localhost:8080", "server address")
	upgrader = websocket.Upgrader{}
)

func tictactoe(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func onStartGameSession(gameSession gameliftgo.GameSession) {
	log.Printf("onStartGameSession")
	gameliftgo.ActivateGameSession()
}

func onProcessTerminate() {
	log.Println("onProcessTerminate")
	srv.Close()
	gameliftgo.ProcessEnding()
}

func onHealthCheck() bool {
	log.Println("onHealthCheck")
	return true
}

func main() {
	if err := gameliftgo.InitSDK(); err != nil {
		log.Fatal(err)
	}
	if err := gameliftgo.ProcessReady(onStartGameSession, onProcessTerminate, onHealthCheck, 8080); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", tictactoe)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Path
		message = strings.TrimPrefix(message, "/")
		message = "Hello " + message
		w.Write([]byte(message))
	})
	log.Fatal(srv.ListenAndServe())
}
