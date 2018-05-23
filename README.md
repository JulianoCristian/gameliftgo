# gameliftgo
Go wrapper for GameLift Server SDK

```
import "github.com/marchinram/gameliftgo"

func onStartGameSession(gameSession gameliftgo.GameSession) {
  if err := gameliftgo.ActivateGameSession(); err != nil {
    log.Print(err)
  }
}

func onProcessTerminate() {
}

func onHealthCheck() bool {
  return true
}

func main() {  
  if err := gameliftgo.InitSDK(); err != nil {
    log.Fatal(err)
  }
  
  port := 8080
  logPaths := []string{"./logs"}
  if err := gameliftgo.ProcessReady(onStartGameSession, onProcessTerminate, onHealthCheck, port, logPaths); err != nil {
    log.Fatal(err)
  }
  
  ... Run server loop ...
  
  gameliftgo.ProcessEnding()
  gameliftgo.Destroy()
}
```
