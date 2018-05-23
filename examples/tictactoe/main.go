package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marchinram/gameliftgo"
)

const (
	logToFile = false // set to false when running locally
)

var (
	server    = NewServer(onPlayerConnected, onPlayerDisconnected, onCommandReceived)
	game      = NewGame()
	terminate = make(chan bool)
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

	if !game.AddPlayer(playerSession.PlayerID) {
		log.Printf("Error adding player %s, game is full", playerSession.PlayerID)
		return "", false
	}
	if err := gameliftgo.AcceptPlayerSession(playerSessionID); err != nil {
		log.Printf("AcceptPlayerSession err: %v", err)
		return "", false
	}

	return playerSession.PlayerID, true
}

func onPlayerDisconnected(playerSessionID string, playerID string) {
	if err := gameliftgo.RemovePlayerSession(playerSessionID); err != nil {
		log.Printf("RemovePlayerSession err: %v", err)
	}
	if !game.RemovePlayer(playerID) {
		log.Printf("Error adding player %s, not in game", playerID)
	}
	if game.IsGameEmpty() {
		terminate <- true
	}
}

func onCommandReceived(command Command) {
	game.HandleCommand(command)
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

	port, err := GetPort(10000, 60000)
	if err != nil {
		log.Fatal(err)
	}
	if err := gameliftgo.InitSDK(); err != nil {
		log.Fatal(err)
	}
	if err := gameliftgo.ProcessReady(onStartGameSession, onProcessTerminate, onHealthCheck, port, logPaths); err != nil {
		log.Fatal(err)
	}

	server.Listen(port)

	select {
	case <-terminate:
		break
	}

	server.ShutdownServer()

	gameliftgo.ProcessEnding()
	gameliftgo.Destroy()
}
