#ifndef __WRAPPER_HPP__
#define __WRAPPER_HPP__

#ifdef __cplusplus
extern "C" {
#endif

int InitSDK();

int ProcessReady(int, int, int, int);

int ProcessEnding();

int ActivateGameSession();

int TerminateGameSession();

int AcceptPlayerSession(char *);

int RemovePlayerSession(char *);

// DescribePlayerSessions

int Destroy();

extern void onStartGameSessionGo(int, char *, char *, char *, int, int, int, char **, char**, char *, int, char *, char *, char*);

extern void onProcessTerminateGo(int);

extern int onHealthCheckGo(int);

#ifdef __cplusplus
}
#endif

#endif