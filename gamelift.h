#ifndef __WRAPPER_HPP__
#define __WRAPPER_HPP__

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    const char *GameSessionID;
    const char *Name;
    const char *FleetID;
    int MaximumPlayerSessionCount;
    int Status;
    int GamePropertiesCount;
    char **GamePropertiesKeys;
    char **GamePropertiesValues;
    const char *IPAddress;
    int Port;
    const char *GameSessionData;
    const char *MatchmakerData;
    const char *DNSName;
} GameSessionC;

int InitSDK();

int ProcessReady(int, int, int, int);

int ProcessEnding();

int ActivateGameSession();

int TerminateGameSession();

int AcceptPlayerSession(char *);

int RemovePlayerSession(char *);

int DescribePlayerSessions(char *);

int Destroy();

extern void onStartGameSessionGo(int, GameSessionC gameSession);

extern void onProcessTerminateGo(int);

extern int onHealthCheckGo(int);

#ifdef __cplusplus
}
#endif

#endif