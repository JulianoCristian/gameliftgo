package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marchinram/gameliftgo"
	"github.com/marchinram/gameliftgo/examples/tictactoe/game"
	"github.com/marchinram/gameliftgo/examples/tictactoe/socket"
	"github.com/marchinram/gameliftgo/examples/tictactoe/util"
)

const (
	logToFile = false // set to false when running locally
)

var (
	server    *socket.Server
	tictactoe *game.TicTacToeGame
	terminate = make(chan bool)
)

func onStartGameSession(gameSession gameliftgo.GameSession) {
	if err := gameliftgo.ActivateGameSession(); err != nil {
		log.Print(err)
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
		log.Printf("DescribePlayerSessions error: %v", err)
		return "", false
	}
	if len(response.PlayerSessions) != 1 {
		log.Println("DescribePlayerSessions error: empty result")
		return "", false
	}
	playerSession := response.PlayerSessions[0]
	if err := gameliftgo.AcceptPlayerSession(playerSessionID); err != nil {
		log.Printf("AcceptPlayerSession error: %v", err)
		return "", false
	}
	return playerSession.PlayerID, true
}

func onPlayerDisconnected(playerSessionID string, playerID string) {
	if err := gameliftgo.RemovePlayerSession(playerSessionID); err != nil {
		log.Printf("RemovePlayerSession error: %v", err)
	}
	if tictactoe.IsGameEmpty() {
		terminate <- true
	}
}

func onCommandReceived(command game.Command) {
	tictactoe.HandleCommand(command)
}

func init() {
	server = socket.NewServer(onPlayerConnected, onPlayerDisconnected, onCommandReceived)
	tictactoe = game.NewTicTacToeGame(server)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logPaths := []string{"./logs"}
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

	server.Listen(port)

	select {
	case <-terminate:
		break
	}

	server.ShutdownServer()

	gameliftgo.ProcessEnding()
	gameliftgo.Destroy()
}
