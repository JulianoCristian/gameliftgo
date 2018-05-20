#ifndef __WRAPPER_HPP__
#define __WRAPPER_HPP__

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    int IsSuccess;
    int ErrorType;
} OutcomeC;

typedef struct {
    const char *GameSessionID;
    const char *Name;
    const char *FleetID;
    int MaximumPlayerSessionCount;
    const char *Status;
    int GamePropertiesCount;
    const char **GamePropertiesKeys;
    const char **GamePropertiesValues;
    const char *IPAddress;
    int Port;
    const char *GameSessionData;
    const char *MatchmakerData;
    const char *DNSName;
} GameSessionC;

typedef struct {
    const char *PlayerSessionID;
    const char *GameSessionID;
    const char *FleetID;
    long CreationTime;
    long TerminationTime;
    const char *Status;
    const char *IPAddress;
    int Port;
    const char *PlayerData;
    const char *DNSName;
} PlayerSessionC;

typedef struct {
    const char *GameSessionID;
    int Limit;
    const char *NextToken;
    const char *PlayerID;
    const char *PlayerSessionID;
    const char *PlayerSessionStatusFilter;
} DescribePlayerSessionsRequestC;

typedef struct {
    const char *NextToken;
    int PlayerSessionsCount;
    PlayerSessionC *PlayerSessions;
} DescribePlayerSessionsResultC;

typedef struct {
    int IsSuccess;
    int ErrorType;
    DescribePlayerSessionsResultC Result;
} DescribePlayerSessionsOutcomeC;

OutcomeC InitSDK();

OutcomeC ProcessReady(int, int, int, int, const char **, int);

OutcomeC ProcessEnding();

OutcomeC ActivateGameSession();

OutcomeC TerminateGameSession();

OutcomeC AcceptPlayerSession(char *);

OutcomeC RemovePlayerSession(char *);

DescribePlayerSessionsOutcomeC DescribePlayerSessions(DescribePlayerSessionsRequestC);

OutcomeC Destroy();

extern void onStartGameSessionGo(int, GameSessionC gameSession);

extern void onProcessTerminateGo(int);

extern int onHealthCheckGo(int);

#ifdef __cplusplus
}
#endif

#endif