package socket

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/marchinram/gameliftgo/examples/tictactoe/game"
)

type ConnectHandlerFunc func(playerSessionID string) (string, bool)
type DisconnectHandlerFunc func(playerSessionID string, playerID string)
type CommandHandlerFunc func(command game.Command)

type Server struct {
	httpServer           http.Server
	onPlayerConnected    ConnectHandlerFunc
	onPlayerDisconnected DisconnectHandlerFunc
	onCommandReceived    CommandHandlerFunc
	connMap              map[string]*websocket.Conn
}

func (s *Server) Listen(port int) {
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
				log.Println(err)
				break
			}
			s.onCommandReceived(command)
		}

		delete(s.connMap, playerID)
		s.onPlayerDisconnected(playerSessionID, playerID)
	})

	s.httpServer = http.Server{Addr: fmt.Sprintf(":%d", port), Handler: router}
	go func() { log.Println(s.httpServer.ListenAndServe()) }()
	log.Printf("Listening on port %d\n", port)
}

func (s *Server) ShutdownServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}

func (s *Server) Write(message game.Message) {
	if message.SendMode == game.SendModeEveryone {
		for _, conn := range s.connMap {
			if err := conn.WriteJSON(message); err != nil {
				log.Println(err)
			}
		}
	} else {
		for _, recipientID := range message.RecipientIDs {
			for connID, conn := range s.connMap {
				if recipientID == connID {
					if err := conn.WriteJSON(message); err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}

func NewServer(onPlayerConnected ConnectHandlerFunc, onPlayerDisconnected DisconnectHandlerFunc, onCommandReceived CommandHandlerFunc) *Server {
	return &Server{
		onPlayerConnected:    onPlayerConnected,
		onPlayerDisconnected: onPlayerDisconnected,
		onCommandReceived:    onCommandReceived,
		connMap:              make(map[string]*websocket.Conn),
	}
}
