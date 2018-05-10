#ifndef __WRAPPER_HPP__
#define __WRAPPER_HPP__

#ifdef __cplusplus
extern "C" {
#endif

int InitSDK();

int ProcessReady(int, int, int, int);

int ProcessEnding();

int ActivateGameSession();

extern void onStartGameSessionGo(int, char *);

extern void onProcessTerminateGo(int);

extern int onHealthCheckGo(int);

#ifdef __cplusplus
}
#endif

#endif