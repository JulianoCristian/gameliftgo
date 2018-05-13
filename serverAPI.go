package gameliftgo

// #cgo CXXFLAGS: -I${SRCDIR}/GameLift_02_15_2018/include -std=c++11
// #cgo darwin LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/darwin-amd64/libaws-cpp-sdk-gamelift-server.a
// #cgo darwin LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/darwin-amd64/libprotobuf.a
// #cgo darwin LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/darwin-amd64/libsioclient.a
// #cgo darwin LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/darwin-amd64/libboost_system.a
// #cgo linux LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/linux-amd64/libaws-cpp-sdk-gamelift-server.a
// #cgo linux LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/linux-amd64/libprotobuf.a
// #cgo linux LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/linux-amd64/libsioclient.a
// #cgo linux LDFLAGS: ${SRCDIR}/GameLift_02_15_2018/lib/linux-amd64/libboost_system.a
// #include <stdlib.h>
// #include "gamelift.h"
import "C"
import (
	"unsafe"
)

func InitSDK() error {
	if errorType := C.InitSDK(); errorType >= 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(errorType)}
	}
	return nil
}

func ProcessReady(onStartGameSession func(GameSession), onProcessTerminate func(), onHealthCheck func() bool, port int) error {
	onStartGameSessionCallback := C.int(register(onStartGameSession))
	onProcessTerminateCallback := C.int(register(onProcessTerminate))
	onHealthCheckCallback := C.int(register(onHealthCheck))
	if errorType := C.ProcessReady(onStartGameSessionCallback, onProcessTerminateCallback, onHealthCheckCallback, C.int(port)); errorType >= 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(errorType)}
	}
	return nil
}

func ProcessEnding() error {
	unregisterAll()
	if errorType := C.ProcessEnding(); errorType >= 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(errorType)}
	}
	return nil
}

func ActivateGameSession() error {
	if errorType := C.ActivateGameSession(); errorType >= 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(errorType)}
	}
	return nil
}

func TerminateGameSession() error {
	if errorType := C.TerminateGameSession(); errorType >= 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(errorType)}
	}
	return nil
}

func AcceptPlayerSession(playerSessionID string) error {
	cPlayerSessionID := C.CString(playerSessionID)
	defer C.free(unsafe.Pointer(cPlayerSessionID))
	if errorType := C.AcceptPlayerSession(cPlayerSessionID); errorType >= 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(errorType)}
	}
	return nil
}

func RemovePlayerSession(playerSessionID string) error {
	cPlayerSessionID := C.CString(playerSessionID)
	defer C.free(unsafe.Pointer(cPlayerSessionID))
	if errorType := C.RemovePlayerSession(cPlayerSessionID); errorType >= 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(errorType)}
	}
	return nil
}

// DescribePlayerSessions

func Destroy() error {
	if errorType := C.Destroy(); errorType >= 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(errorType)}
	}
	return nil
}

//export onStartGameSessionGo
func onStartGameSessionGo(onStartGameSessionCallback C.int, gameSessionID *C.char, name *C.char, fleetID *C.char,
	maximumPlayerSessionCount C.int, status C.int, gamePropertiesSize C.int, gamePropertiesKeys **C.char, gamePropertiesValues **C.char,
	ipAddress *C.char, port C.int, gameSessionData *C.char, matchmakerData *C.char, dnsName *C.char) {
	gameProperties := make(map[string]string)
	size := unsafe.Sizeof(int(0))
	for i := 0; i < int(gamePropertiesSize); i++ {
		key := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(gamePropertiesKeys)) + size*uintptr(i)))
		val := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(gamePropertiesValues)) + size*uintptr(i)))
		gameProperties[C.GoString(key)] = C.GoString(val)
	}
	if callback := lookup(int(onStartGameSessionCallback)).(func(GameSession)); callback != nil {
		callback(GameSession{
			GameSessionID:             C.GoString(gameSessionID),
			Name:                      C.GoString(name),
			FleetID:                   C.GoString(fleetID),
			MaximumPlayerSessionCount: int(maximumPlayerSessionCount),
			Status:          GameSessionStatus(status),
			GameProperties:  gameProperties,
			IPAddress:       C.GoString(ipAddress),
			Port:            int(port),
			GameSessionData: C.GoString(gameSessionData),
			MatchmakerData:  C.GoString(matchmakerData),
			DNSName:         C.GoString(dnsName),
		})
	}
}

//export onProcessTerminateGo
func onProcessTerminateGo(onProcessTerminateCallback C.int) {
	if callback := lookup(int(onProcessTerminateCallback)).(func() bool); callback != nil {
		callback()
	}
}

//export onHealthCheckGo
func onHealthCheckGo(onHealthCheckCallback C.int) C.int {
	if callback := lookup(int(onHealthCheckCallback)).(func() bool); callback != nil {
		if callback() {
			return C.int(1)
		}
		return C.int(0)
	}
	return C.int(1)
}
