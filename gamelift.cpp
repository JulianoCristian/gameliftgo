#include "gamelift.h"

#define GAMELIFT_USE_STD 1
#include <aws/gamelift/server/GameLiftServerAPI.h>

OutcomeC InitSDK() {
    auto outcome = Aws::GameLift::Server::InitSDK();
    if (!outcome.IsSuccess()) {
        return {0, static_cast<int>(outcome.GetError().GetErrorType())};
    }
    return {1, 0};
}

OutcomeC ProcessReady(int onStartGameSessionCallback, int onProcessTerminateCallback, int onHealthCheckCallback, int port, const char **logPathsC, int logPathsCCount) {
    auto onStartGameSession = [onStartGameSessionCallback](Aws::GameLift::Server::Model::GameSession gameSession){
        auto gameProperties = gameSession.GetGameProperties();
        const char **gamePropertiesKeys = new const char*[gameProperties.size()];
        const char **gamePropertiesValues = new const char*[gameProperties.size()];

        int i = 0;
        for (auto it = gameProperties.begin(); it != gameProperties.end(); it++) {
            gamePropertiesKeys[i] = new char[32 + 1];
            memset((void *)gamePropertiesKeys[i], 0, 32 + 1);
            memcpy((void *)gamePropertiesKeys[i], (*it).GetKey().c_str(), 32 + 1);
            gamePropertiesValues[i] = new char[96 + 1];
            memset((void *)gamePropertiesValues[i], 0, 96 + 1);
            memcpy((void *)gamePropertiesValues[i], (*it).GetValue().c_str(), 96 + 1);
            i++;
        }

        GameSessionC gameSessionC;
        gameSessionC.GameSessionID = gameSession.GetGameSessionId().c_str();
        gameSessionC.Name = gameSession.GetName().c_str();
        gameSessionC.FleetID = gameSession.GetFleetId().c_str();
        
        gameSessionC.MaximumPlayerSessionCount = gameSession.GetMaximumPlayerSessionCount();
        
        auto status = Aws::GameLift::Server::Model::GameSessionStatusMapper::GetNameForGameSessionStatus(gameSession.GetStatus()).c_str();
        gameSessionC.Status = new char[1024 + 1];
        memset((void *)gameSessionC.Status, 0, 1024 + 1);
        memcpy((void *)gameSessionC.Status, status, 1024 + 1);
        
        gameSessionC.GamePropertiesCount = gameProperties.size();
        gameSessionC.GamePropertiesKeys = gamePropertiesKeys;
        gameSessionC.GamePropertiesValues = gamePropertiesValues;
        
        gameSessionC.IPAddress = gameSession.GetIpAddress().c_str();
        gameSessionC.Port = gameSession.GetPort();
        
        gameSessionC.GameSessionData = gameSession.GetGameSessionData().c_str();
        gameSessionC.MatchmakerData = gameSession.GetMatchmakerData().c_str();
        
        gameSessionC.DNSName = gameSession.GetDnsName().c_str();
        
        onStartGameSessionGo(onStartGameSessionCallback, gameSessionC);

        delete[] gameSessionC.Status;

        for (int i = 0; i < gameProperties.size(); i++) {
            delete gameSessionC.GamePropertiesKeys[i];
            delete gameSessionC.GamePropertiesValues[i];
        }
        delete[] gameSessionC.GamePropertiesKeys;
        delete[] gameSessionC.GamePropertiesValues;
    };

    auto onProcessTerminate = [onProcessTerminateCallback]() {
        onProcessTerminateGo(onProcessTerminateCallback);
    };

    auto onHealthCheck = [onHealthCheckCallback]() {
        return onHealthCheckGo(onHealthCheckCallback);
    };
    
    std::vector<std::string> logPaths;
    for (int i = 0; i < logPathsCCount; i++) {
        logPaths.push_back(logPathsC[i]);
    }
    
    Aws::GameLift::Server::ProcessParameters processReadyParameter(onStartGameSession,
                                                                   onProcessTerminate,
                                                                   onHealthCheck,
                                                                   port,
                                                                   Aws::GameLift::Server::LogParameters(logPaths));

    auto outcome = Aws::GameLift::Server::ProcessReady(processReadyParameter);
    if (!outcome.IsSuccess()) {
        return {0, static_cast<int>(outcome.GetError().GetErrorType())};
    }
    return {1, 0};
}

OutcomeC ProcessEnding() {
    auto outcome = Aws::GameLift::Server::ProcessEnding();
    if (!outcome.IsSuccess()) {
        return {0, static_cast<int>(outcome.GetError().GetErrorType())};
    }
    return {1, 0};
}

OutcomeC ActivateGameSession() {
    auto outcome = Aws::GameLift::Server::ActivateGameSession();
    if (!outcome.IsSuccess()) {
        return {0, static_cast<int>(outcome.GetError().GetErrorType())};
    }
    return {1, 0};
}

OutcomeC TerminateGameSession() {
    auto outcome = Aws::GameLift::Server::TerminateGameSession();
    if (!outcome.IsSuccess()) {
        return {0, static_cast<int>(outcome.GetError().GetErrorType())};
    }
    return {1, 0};
}

OutcomeC AcceptPlayerSession(char *playerSessionID) {
    auto outcome = Aws::GameLift::Server::AcceptPlayerSession(std::string(playerSessionID));
    if (!outcome.IsSuccess()) {
        return {0, static_cast<int>(outcome.GetError().GetErrorType())};
    }
    return {1, 0};
}

OutcomeC RemovePlayerSession(char *playerSessionID) {
    auto outcome = Aws::GameLift::Server::RemovePlayerSession(std::string(playerSessionID));
    if (!outcome.IsSuccess()) {
        return {0, static_cast<int>(outcome.GetError().GetErrorType())};
    }
    return {1, 0};
}

DescribePlayerSessionsOutcomeC DescribePlayerSessions(DescribePlayerSessionsRequestC requestC) {
    Aws::GameLift::Server::Model::DescribePlayerSessionsRequest request;
    request.SetGameSessionId(requestC.GameSessionID);
    request.SetLimit(requestC.Limit);
    request.SetNextToken(requestC.NextToken);
    request.SetPlayerId(requestC.PlayerID);
    request.SetPlayerSessionId(requestC.PlayerSessionID);
    request.SetPlayerSessionStatusFilter(requestC.PlayerSessionStatusFilter);
    auto outcome = Aws::GameLift::Server::DescribePlayerSessions(request);
    if (!outcome.IsSuccess()) {
        return {0, static_cast<int>(outcome.GetError().GetErrorType())};
    }

    auto playerSessions = outcome.GetResult().GetPlayerSessions();
    PlayerSessionC *playerSessionsC = new PlayerSessionC[playerSessions.size()];
    int i = 0;
    for (auto it = playerSessions.begin(); it != playerSessions.end(); it++) {
        PlayerSessionC playerSessionC;

        playerSessionC.PlayerSessionID = new char[1024 + 1];
        memset((void *)playerSessionC.PlayerSessionID, 0, 1024 + 1);
        memcpy((void *)playerSessionC.PlayerSessionID, (*it).GetPlayerSessionId().c_str(), 1024 + 1);

        playerSessionC.PlayerID = new char[1024 + 1];
        memset((void *)playerSessionC.PlayerID, 0, 1024 + 1);
        memcpy((void *)playerSessionC.PlayerID, (*it).GetPlayerId().c_str(), 1024 + 1);

        playerSessionC.GameSessionID = new char[1024 + 1];
        memset((void *)playerSessionC.GameSessionID, 0, 1024 + 1);
        memcpy((void *)playerSessionC.GameSessionID, (*it).GetGameSessionId().c_str(), 1024 + 1);
        
        playerSessionC.FleetID = new char[1024 + 1];
        memset((void *)playerSessionC.FleetID, 0, 1024 + 1);
        memcpy((void *)playerSessionC.FleetID, (*it).GetFleetId().c_str(), 1024 + 1);
        
        playerSessionC.CreationTime = (*it).GetCreationTime();
        playerSessionC.TerminationTime = (*it).GetTerminationTime();

        const char *status = Aws::GameLift::Server::Model::PlayerSessionStatusMapper::GetNameForPlayerSessionStatus((*it).GetStatus()).c_str();
        playerSessionC.Status = new char[1024 + 1];
        memset((void *)playerSessionC.Status, 0, 1024 + 1);
        memcpy((void *)playerSessionC.Status, status, 1024 + 1);
        
        playerSessionC.IPAddress = new char[16 + 1];
        memset((void *)playerSessionC.IPAddress, 0, 16 + 1);
        memcpy((void *)playerSessionC.IPAddress, (*it).GetIpAddress().c_str(), 16 + 1);
        
        playerSessionC.Port = (*it).GetPort();
        
        playerSessionC.PlayerData = new char[2048 + 1];
        memset((void *)playerSessionC.PlayerData, 0, 2048 + 1);
        memcpy((void *)playerSessionC.PlayerData, (*it).GetPlayerData().c_str(), 2048 + 1);
        
        playerSessionC.DNSName = new char[1024 + 1];
        memset((void *)playerSessionC.DNSName, 0, 1024 + 1);
        memcpy((void *)playerSessionC.DNSName, (*it).GetDnsName().c_str(), 1024 + 1);
        
        playerSessionsC[i++] = playerSessionC;
    }

    DescribePlayerSessionsResultC describePlayerSessionsResultC;
    memset(&describePlayerSessionsResultC.NextToken, 0, 1024 + 1);
    memcpy(&describePlayerSessionsResultC.NextToken, outcome.GetResult().GetNextToken().c_str(), 1024 + 1);
    describePlayerSessionsResultC.PlayerSessionsCount = playerSessions.size();
    describePlayerSessionsResultC.PlayerSessions = playerSessionsC;

    DescribePlayerSessionsOutcomeC describePlayerSessionsOutcomeC = {1, 0};
    describePlayerSessionsOutcomeC.Result = describePlayerSessionsResultC;
    return describePlayerSessionsOutcomeC;
}

OutcomeC Destroy() {
    auto outcome = Aws::GameLift::Server::Destroy();
    if (!outcome.IsSuccess()) {
        return {0, static_cast<int>(outcome.GetError().GetErrorType())};
    }
    return {1, 0};
}
