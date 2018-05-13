package gameliftgo

type GameLiftErrorType int

const (
	ErrAlreadyInitialized           GameLiftErrorType = iota
	ErrFleetMismatch                                  = iota
	ErrGameliftClientNotInitialized                   = iota
	ErrGameliftServerNotInitialized                   = iota
	ErrGameSessionEndedFailed                         = iota
	ErrGameSessionNotReady                            = iota
	ErrGameSessionReadyFailed                         = iota
	ErrInitializationMismatch                         = iota
	ErrNotInitialized                                 = iota
	ErrNoTargetAliasIDSet                             = iota
	ErrNoTargetFleetSet                               = iota
	ErrProcessEndingFailed                            = iota
	ErrProcessNotActive                               = iota
	ErrProcessNotReady                                = iota
	ErrProcessReadyFailed                             = iota
	ErrSDKVersionDetectionFailed                      = iota
	ErrServiceCallFailed                              = iota
	ErrSTXCallFailed                                  = iota
	ErrSTXInitializationFailed                        = iota
	ErrUnexpectedPlayerSession                        = iota
)

type GameLiftError struct {
	ErrorType GameLiftErrorType
}

func (e *GameLiftError) Error() string {
	switch e.ErrorType {
	case ErrAlreadyInitialized:
		return "GameLift has already been initialized. You must call Destroy() before reinitializing the client or server."
	case ErrFleetMismatch:
		return "The Target fleet does not match the request fleet. Make sure GameSessions and PlayerSessions belong to your target fleet."
	case ErrGameliftClientNotInitialized:
		return "The GameLift client has not been initialized. Please call SetTargetFleet or SetTArgetAliasId."
	case ErrGameliftServerNotInitialized:
		return "The GameLift server has not been initialized. Please call InitSDK."
	case ErrGameSessionEndedFailed:
		return "The GameSessionEnded invocation failed."
	case ErrGameSessionNotReady:
		return "The Game session associated with this server was not activated."
	case ErrGameSessionReadyFailed:
		return "The GameSessionReady invocation failed."
	case ErrInitializationMismatch:
		return "The current call does not match the initialization state. Client calls require a call to Client::Initialize(), and Server calls require Server::Initialize(). Only one may be active at a time."
	case ErrNotInitialized:
		return "GameLift has not been initialized! You must call Client::Initialize() or Server::InitSDK() before making GameLift calls."
	case ErrNoTargetAliasIDSet:
		return "The aliasId has not been set. Clients should call SetTargetAliasId() before making calls that require an alias."
	case ErrNoTargetFleetSet:
		return "The target fleet has not been set. Clients should call SetTargetFleet() before making calls that require a fleet."
	case ErrProcessEndingFailed:
		return "ProcessEnding call to GameLift failed."
	case ErrProcessNotActive:
		return "The process has not yet been activated."
	case ErrProcessNotReady:
		return "The process has not yet been activated by calling ProcessReady(). Processes in standby cannot receive StartGameSession callbacks."
	case ErrProcessReadyFailed:
		return "ProcessReady call to GameLift failed."
	case ErrSDKVersionDetectionFailed:
		return "Could not detect SDK version."
	case ErrServiceCallFailed:
		return "An AWS service call has failed. See the root cause error for more information."
	case ErrSTXCallFailed:
		return "An internal call to the STX server backend component has failed."
	case ErrSTXInitializationFailed:
		return "The STX server backend component has failed to initialize."
	case ErrUnexpectedPlayerSession:
		return "The player session was not expected by the server. Clients wishing to connect to a server must obtain a PlayerSessionID from GameLift by creating a player session on the desired server's game instance."
	}
	return "An unexpected error has occurred."
}
