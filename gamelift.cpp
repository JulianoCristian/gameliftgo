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
        
        onStartGameSessionGo(onStartGameSessionCallback, 
            const_cast<char*>(gameSession.GetGameSessionId().c_str()),
            const_cast<char*>(gameSession.GetName().c_str()),
            const_cast<char*>(gameSession.GetFleetId().c_str()),
            gameSession.GetMaximumPlayerSessionCount(),
            static_cast<int>(gameSession.GetStatus()),
            gameProperties.size(), 
            &gamePropertiesKeys[0],
            &gamePropertiesValues[0],
            const_cast<char*>(gameSession.GetIpAddress().c_str()),
            gameSession.GetPort(),
            const_cast<char*>(gameSession.GetGameSessionData().c_str()),
            const_cast<char*>(gameSession.GetMatchmakerData().c_str()),
            const_cast<char*>(gameSession.GetDnsName().c_str()));
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

// DescribePlayerSessions

int Destroy() {
    auto outcome = Aws::GameLift::Server::Destroy();
    if (!outcome.IsSuccess()) {
        return static_cast<int>(outcome.GetError().GetErrorType());
    }
    return -1;
}
