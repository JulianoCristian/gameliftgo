#include "wrapper.h"

#define GAMELIFT_USE_STD 1
#include <aws/gamelift/server/GameLiftServerAPI.h>

int InitSDK() {
    return Aws::GameLift::Server::InitSDK().IsSuccess();
}

int ProcessReady(int onStartGameSessionCallback, int onProcessTerminateCallback, int onHealthCheckCallback, int port) {
    auto onStartGameSession = [onStartGameSessionCallback](Aws::GameLift::Server::Model::GameSession gameSession){
        onStartGameSessionGo(onStartGameSessionCallback, const_cast<char*>(gameSession.GetName().c_str()));
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

    auto readyOutcome = Aws::GameLift::Server::ProcessReady(processReadyParameter);

    return readyOutcome.IsSuccess();
}

int ProcessEnding() {
    return Aws::GameLift::Server::ProcessEnding().IsSuccess();
}

int ActivateGameSession() {
    return Aws::GameLift::Server::ActivateGameSession().IsSuccess();
}
