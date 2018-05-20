package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marchinram/gameliftgo"
	"github.com/marchinram/gameliftgo/example/tictactoe/game"
	"github.com/marchinram/gameliftgo/example/tictactoe/server"
	"github.com/marchinram/gameliftgo/example/tictactoe/util"
)

const (
	logToFile = false // set to false when running locally
)

var (
	gameServer server.GameServer
	terminate  = make(chan bool)
)

func onStartGameSession(gameSession gameliftgo.GameSession) {
	if err := gameliftgo.ActivateGameSession(); err != nil {
		log.Fatal(err)
	}
}

func onProcessTerminate() {
	terminate <- true
}

func onHealthCheck() bool {
	return true
}

func onPlayerConnected(playerSessionID string) (string, bool) {
	if err := gameliftgo.AcceptPlayerSession(playerSessionID); err != nil {
		log.Printf("AcceptPlayerSession err: %v", err)
		return "", false
	}
	request := gameliftgo.DescribePlayerSessionsRequest{
		PlayerSessionID: playerSessionID,
	}
	response, err := gameliftgo.DescribePlayerSessions(request)
	if err != nil {
		log.Printf("AcceptDescribePlayerSessionsPlayerSession err: %v", err)
		return "", false
	}
	if len(response.PlayerSessions) != 1 {
		log.Println("Error getting PlayerSession")
		return "", false
	}
	playerSession := response.PlayerSessions[0]
	return playerSession.PlayerID, true
}

func onPlayerDisconnected(playerSessionID string) {
	if err := gameliftgo.RemovePlayerSession(playerSessionID); err != nil {
		log.Printf("RemovePlayerSession err: %v", err)
	}
	// If count of players is 0 terminate
	terminate <- true
}

func onCommandReceived(game.Command) game.Message {
	// log.Println("Received command: ")
	return game.Message{}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logPaths := make([]string, 0)
	if logToFile {
		os.MkdirAll("./logs/gameLiftLogs", os.ModePerm)
		logPaths = append(logPaths, "./logs/gameLiftLogs")

		serverLogsDir := fmt.Sprintf("./logs/serverLogs/%d", os.Getpid())
		os.MkdirAll(serverLogsDir, os.ModePerm)
		file, _ := os.Create(fmt.Sprintf("%s/serverLog.txt", serverLogsDir))
		log.SetOutput(file)
	}

	port, err := util.GetPort(10000, 60000)
	if err != nil {
		log.Fatal(err)
	}
	if err := gameliftgo.InitSDK(); err != nil {
		log.Fatal(err)
	}
	if err := gameliftgo.ProcessReady(onStartGameSession, onProcessTerminate, onHealthCheck, port, logPaths); err != nil {
		log.Fatal(err)
	}

	gameServer = server.NewGameServer(onPlayerConnected, onPlayerDisconnected, onCommandReceived)
	gameServer.Listen(port)

	select {
	case <-terminate:
		break
	}

	gameServer.ShutdownServer()

	gameliftgo.ProcessEnding()
	gameliftgo.Destroy()
}
