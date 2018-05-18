#include "gamelift.h"

#define GAMELIFT_USE_STD 1
#include <aws/gamelift/server/GameLiftServerAPI.h>

int InitSDK() {
    auto outcome = Aws::GameLift::Server::InitSDK();
    if (!outcome.IsSuccess()) {
        return static_cast<int>(outcome.GetError().GetErrorType());
    }
    return -1;
}

int ProcessReady(int onStartGameSessionCallback, int onProcessTerminateCallback, int onHealthCheckCallback, int port) {
    auto onStartGameSession = [onStartGameSessionCallback](Aws::GameLift::Server::Model::GameSession gameSession){
        auto gameProperties = gameSession.GetGameProperties();
        std::vector<char *> gamePropertiesKeys;
        for (auto it = gameProperties.begin(); it != gameProperties.end(); it++) {
            gamePropertiesKeys.push_back(const_cast<char*>((*it).GetKey().c_str()));
        }
        std::vector<char *> gamePropertiesValues;
        for (auto it = gameProperties.begin(); it != gameProperties.end(); it++) {
            gamePropertiesValues.push_back(const_cast<char*>((*it).GetValue().c_str()));
        }

        GameSessionC gameSessionOut;
        gameSessionOut.GameSessionID = gameSession.GetGameSessionId().c_str();
        gameSessionOut.Name = gameSession.GetName().c_str();
        gameSessionOut.FleetID = gameSession.GetFleetId().c_str();
        gameSessionOut.MaximumPlayerSessionCount = gameSession.GetMaximumPlayerSessionCount();
        gameSessionOut.Status = static_cast<int>(gameSession.GetStatus());
        gameSessionOut.GamePropertiesCount = gameProperties.size();
        gameSessionOut.GamePropertiesKeys = &gamePropertiesKeys[0];
        gameSessionOut.GamePropertiesValues = &gamePropertiesValues[0];
        gameSessionOut.IPAddress = gameSession.GetIpAddress().c_str();
        gameSessionOut.Port = gameSession.GetPort();
        gameSessionOut.GameSessionData = gameSession.GetGameSessionData().c_str();
        gameSessionOut.MatchmakerData = gameSession.GetMatchmakerData().c_str();
        gameSessionOut.DNSName = gameSession.GetDnsName().c_str();
        onStartGameSessionGo(onStartGameSessionCallback, gameSessionOut);
    };

    auto onProcessTerminate = [onProcessTerminateCallback]() {
        onProcessTerminateGo(onProcessTerminateCallback);
    };

    auto onHealthCheck = [onHealthCheckCallback]() {
        return onHealthCheckGo(onHealthCheckCallback);
    };
    
    std::vector<std::string> logPaths;

    Aws::GameLift::Server::ProcessParameters processReadyParameter(onStartGameSession,
                                                                   onProcessTerminate,
                                                                   onHealthCheck,
                                                                   port,
                                                                   Aws::GameLift::Server::LogParameters(logPaths));

    auto outcome = Aws::GameLift::Server::ProcessReady(processReadyParameter);
    if (!outcome.IsSuccess()) {
        return static_cast<int>(outcome.GetError().GetErrorType());
    }
    return -1;
}

int ProcessEnding() {
    auto outcome = Aws::GameLift::Server::ProcessEnding();
    if (!outcome.IsSuccess()) {
        return static_cast<int>(outcome.GetError().GetErrorType());
    }
    return -1;
}

int ActivateGameSession() {
    auto outcome = Aws::GameLift::Server::ActivateGameSession();
    if (!outcome.IsSuccess()) {
        return static_cast<int>(outcome.GetError().GetErrorType());
    }
    return -1;
}

int TerminateGameSession() {
    auto outcome = Aws::GameLift::Server::TerminateGameSession();
    if (!outcome.IsSuccess()) {
        return static_cast<int>(outcome.GetError().GetErrorType());
    }
    return -1;
}

int AcceptPlayerSession(char *playerSessionId) {
    auto outcome = Aws::GameLift::Server::AcceptPlayerSession(std::string(playerSessionId));
    if (!outcome.IsSuccess()) {
        return static_cast<int>(outcome.GetError().GetErrorType());
    }
    return -1;
}

int RemovePlayerSession(char *playerSessionId) {
    auto outcome = Aws::GameLift::Server::RemovePlayerSession(std::string(playerSessionId));
    if (!outcome.IsSuccess()) {
        return static_cast<int>(outcome.GetError().GetErrorType());
    }
    return -1;
}

int DescribePlayerSessions(char *playerSessionId) {
    Aws::GameLift::Server::Model::DescribePlayerSessionsRequest request;
    request.SetPlayerSessionId(playerSessionId);
    auto outcome = Aws::GameLift::Server::DescribePlayerSessions(request);
    if (!outcome.IsSuccess()) {
        return static_cast<int>(outcome.GetError().GetErrorType());
    }
    return -1;
}

int Destroy() {
    auto outcome = Aws::GameLift::Server::Destroy();
    if (!outcome.IsSuccess()) {
        return static_cast<int>(outcome.GetError().GetErrorType());
    }
    return -1;
}
