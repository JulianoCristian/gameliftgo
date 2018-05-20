package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/marchinram/gameliftgo/example/tictactoe/game"
)

type ConnectHandlerFunc func(playerSessionID string) (string, bool)
type DisconnectHandlerFunc func(playerSessionID string)
type CommandHandlerFunc func(game.Command) game.Message

type GameServer interface {
	Listen(port int)
	ShutdownServer()
}

type gameServer struct {
	httpServer           http.Server
	onPlayerConnected    ConnectHandlerFunc
	onPlayerDisconnected DisconnectHandlerFunc
	onCommandReceived    func(game.Command) game.Message
	connMap              map[string]*websocket.Conn
}

func (s *gameServer) Listen(port int) {
	router := mux.NewRouter()

	router.HandleFunc("/game/{pSessID}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		playerSessionID, ok := vars["pSessID"]
		if !ok {
			log.Println("Missing player session ID")
			return
		}
		playerID, ok := s.onPlayerConnected(playerSessionID)
		if !ok {
			return
		}

		upgrader := &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true }, // remove in production
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade error: ", err)
			return
		}
		defer conn.Close()

		s.connMap[playerID] = conn

		for {
			var command game.Command
			if err := conn.ReadJSON(&command); err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					log.Println(err)
				}
				break
			}
			message := s.onCommandReceived(command)
			if err := conn.WriteJSON(message); err != nil {
				log.Println(err)
				break
			}
		}

		delete(s.connMap, playerID)
		s.onPlayerDisconnected(playerSessionID)
	})

	s.httpServer = http.Server{Addr: fmt.Sprintf(":%d", port), Handler: router}
	go func() { log.Println(s.httpServer.ListenAndServe()) }()
	log.Printf("Listening on port %d\n", port)
}

func (s *gameServer) ShutdownServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}

func NewGameServer(onPlayerConnected ConnectHandlerFunc, onPlayerDisconnected DisconnectHandlerFunc, onCommandReceived CommandHandlerFunc) GameServer {
	return &gameServer{
		onPlayerConnected:    onPlayerConnected,
		onPlayerDisconnected: onPlayerDisconnected,
		onCommandReceived:    onCommandReceived,
		connMap:              make(map[string]*websocket.Conn),
	}
}
