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
	"time"
	"unsafe"
)

func InitSDK() error {
	if outcome := C.InitSDK(); outcome.IsSuccess == 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(outcome.ErrorType)}
	}
	return nil
}

func ProcessReady(onStartGameSession func(GameSession), onProcessTerminate func(), onHealthCheck func() bool, port int, logPaths []string) error {
	onStartGameSessionCallback := C.int(register(onStartGameSession))
	onProcessTerminateCallback := C.int(register(onProcessTerminate))
	onHealthCheckCallback := C.int(register(onHealthCheck))
	logPathsC := C.malloc(C.size_t(len(logPaths)) * C.size_t(unsafe.Sizeof(uintptr(0))))
	logPathsCArr := (*[1<<30 - 1]*C.char)(logPathsC)
	for index, logPath := range logPaths {
		logPathsCArr[index] = C.CString(logPath)
	}
	if outcome := C.ProcessReady(onStartGameSessionCallback, onProcessTerminateCallback, onHealthCheckCallback, C.int(port), (**C.char)(unsafe.Pointer(logPathsC)), C.int(len(logPaths))); outcome.IsSuccess == 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(outcome.ErrorType)}
	}
	for index := range logPaths {
		C.free(unsafe.Pointer(logPathsCArr[index]))
	}
	C.free(logPathsC)
	return nil
}

func ProcessEnding() error {
	unregisterAll()
	if outcome := C.ProcessEnding(); outcome.IsSuccess == 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(outcome.ErrorType)}
	}
	return nil
}

func ActivateGameSession() error {
	if outcome := C.ActivateGameSession(); outcome.IsSuccess == 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(outcome.ErrorType)}
	}
	return nil
}

func TerminateGameSession() error {
	if outcome := C.TerminateGameSession(); outcome.IsSuccess == 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(outcome.ErrorType)}
	}
	return nil
}

func AcceptPlayerSession(playerSessionID string) error {
	cPlayerSessionID := C.CString(playerSessionID)
	defer C.free(unsafe.Pointer(cPlayerSessionID))
	if outcome := C.AcceptPlayerSession(cPlayerSessionID); outcome.IsSuccess == 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(outcome.ErrorType)}
	}
	return nil
}

func RemovePlayerSession(playerSessionID string) error {
	cPlayerSessionID := C.CString(playerSessionID)
	defer C.free(unsafe.Pointer(cPlayerSessionID))
	if outcome := C.RemovePlayerSession(cPlayerSessionID); outcome.IsSuccess == 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(outcome.ErrorType)}
	}
	return nil
}

func DescribePlayerSessions(request DescribePlayerSessionsRequest) (*DescribePlayerSessionsResult, error) {
	cGameSessionID := C.CString(request.GameSessionID)
	defer C.free(unsafe.Pointer(cGameSessionID))
	cNextToken := C.CString(request.NextToken)
	defer C.free(unsafe.Pointer(cNextToken))
	cPlayerID := C.CString(request.PlayerID)
	defer C.free(unsafe.Pointer(cPlayerID))
	cPlayerSessionID := C.CString(request.PlayerSessionID)
	defer C.free(unsafe.Pointer(cPlayerSessionID))
	cPlayerSessionStatusFilter := C.CString(request.PlayerSessionStatusFilter)
	defer C.free(unsafe.Pointer(cPlayerSessionStatusFilter))
	requestC := C.DescribePlayerSessionsRequestC{
		GameSessionID:             cGameSessionID,
		Limit:                     C.int(request.Limit),
		NextToken:                 cNextToken,
		PlayerID:                  cPlayerID,
		PlayerSessionID:           cPlayerSessionID,
		PlayerSessionStatusFilter: cPlayerSessionStatusFilter,
	}
	outcome := C.DescribePlayerSessions(requestC)
	if outcome.IsSuccess == 0 {
		return nil, &GameLiftError{ErrorType: GameLiftErrorType(outcome.ErrorType)}
	}
	playerSessions := make([]PlayerSession, 0)
	size := unsafe.Sizeof(C.PlayerSessionC{})
	for i := 0; i < int(outcome.Result.PlayerSessionsCount); i++ {
		playerSessionC := *(*C.PlayerSessionC)(unsafe.Pointer(uintptr(unsafe.Pointer(outcome.Result.PlayerSessions)) + size*uintptr(i)))
		playerSession := PlayerSession{
			PlayerSessionID: C.GoString(playerSessionC.PlayerSessionID),
			GameSessionID:   C.GoString(playerSessionC.GameSessionID),
			FleetID:         C.GoString(playerSessionC.FleetID),
			CreationTime:    time.Unix(int64(playerSessionC.CreationTime/1000), 0),
			TerminationTime: time.Unix(int64(playerSessionC.TerminationTime/1000), 0),
			Status:          PlayerSessionStatus(C.GoString(playerSessionC.Status)),
			IPAddress:       C.GoString(playerSessionC.IPAddress),
			Port:            int(playerSessionC.Port),
			PlayerData:      C.GoString(playerSessionC.PlayerData),
			DNSName:         C.GoString(playerSessionC.DNSName),
		}
		playerSessions = append(playerSessions, playerSession)
		C.free(unsafe.Pointer(playerSessionC.PlayerSessionID))
		C.free(unsafe.Pointer(playerSessionC.GameSessionID))
		C.free(unsafe.Pointer(playerSessionC.FleetID))
		C.free(unsafe.Pointer(playerSessionC.Status))
		C.free(unsafe.Pointer(playerSessionC.IPAddress))
		C.free(unsafe.Pointer(playerSessionC.PlayerData))
		C.free(unsafe.Pointer(playerSessionC.DNSName))
	}
	result := &DescribePlayerSessionsResult{
		NextToken:      C.GoString(outcome.Result.NextToken),
		PlayerSessions: playerSessions,
	}
	C.free(unsafe.Pointer(outcome.Result.NextToken))
	C.free(unsafe.Pointer(outcome.Result.PlayerSessions))
	return result, nil
}

func Destroy() error {
	if outcome := C.Destroy(); outcome.IsSuccess == 0 {
		return &GameLiftError{ErrorType: GameLiftErrorType(outcome.ErrorType)}
	}
	return nil
}

//export onStartGameSessionGo
func onStartGameSessionGo(onStartGameSessionCallback C.int, gameSession C.GameSessionC) {
	gameProperties := make(map[string]string)
	size := unsafe.Sizeof(int(0))
	for i := 0; i < int(gameSession.GamePropertiesCount); i++ {
		key := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(gameSession.GamePropertiesKeys)) + size*uintptr(i)))
		val := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(gameSession.GamePropertiesValues)) + size*uintptr(i)))
		gameProperties[C.GoString(key)] = C.GoString(val)
	}
	if callback := lookup(int(onStartGameSessionCallback)).(func(GameSession)); callback != nil {
		callback(GameSession{
			GameSessionID:             C.GoString(gameSession.GameSessionID),
			Name:                      C.GoString(gameSession.Name),
			FleetID:                   C.GoString(gameSession.FleetID),
			MaximumPlayerSessionCount: int(gameSession.MaximumPlayerSessionCount),
			Status:          GameSessionStatus(C.GoString(gameSession.Status)),
			GameProperties:  gameProperties,
			IPAddress:       C.GoString(gameSession.IPAddress),
			Port:            int(gameSession.Port),
			GameSessionData: C.GoString(gameSession.GameSessionData),
			MatchmakerData:  C.GoString(gameSession.MatchmakerData),
			DNSName:         C.GoString(gameSession.DNSName),
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
