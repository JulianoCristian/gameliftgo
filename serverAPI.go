package gameliftgo

// #cgo CXXFLAGS: -I${SRCDIR}/GameLift_02_15_2018/include -std=c++11
// #cgo LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/linux-amd64/libaws-cpp-sdk-gamelift-server.a
// #cgo LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/linux-amd64/libprotobuf.a
// #cgo LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/linux-amd64/libsioclient.a
// #cgo LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/linux-amd64/libboost_system.a
// #include "wrapper.h"
import "C"
import "errors"

// TODO Improved error handling, user err not bool,

func InitSDK() error {
	if C.InitSDK() == 0 {
		return errors.New("InitSDK Error")
	}
	return nil
}

func ProcessReady(onStartGameSession func(GameSession), onProcessTerminate func(), onHealthCheck func() bool, port int) error {
	onStartGameSessionCallback := register(onStartGameSession)
	onProcessTerminateCallback := register(onProcessTerminate)
	onHealthCheckCallback := register(onHealthCheck)
	if C.ProcessReady(C.int(onStartGameSessionCallback), C.int(onProcessTerminateCallback), C.int(onHealthCheckCallback), C.int(port)) == 0 {
		return errors.New("ProcessReady Error")
	}
	return nil
}

func ProcessEnding() bool {
	unregisterAll()
	return C.ProcessEnding() == 1
}

func ActivateGameSession() bool {
	return C.ActivateGameSession() == 1
}

//export onStartGameSessionGo
func onStartGameSessionGo(onStartGameSessionCallback C.int, name *C.char) {
	callback := lookup(int(onStartGameSessionCallback)).(func(GameSession))
	callback(GameSession{
		Name: C.GoString(name),
	})
}

//export onProcessTerminateGo
func onProcessTerminateGo(onStartGameSessionCallback C.int) {
	callback := lookup(int(onStartGameSessionCallback)).(func())
	callback()
}

//export onHealthCheckGo
func onHealthCheckGo(onHealthCheckCallback C.int) C.int {
	callback := lookup(int(onHealthCheckCallback)).(func() bool)
	if callback() {
		return C.int(1)
	}
	return C.int(0)
}
