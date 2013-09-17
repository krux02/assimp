#pragma once

#include <assimp/cimport.h>
#include <assimp/scene.h>
//#include <assimp/cfileio.h>

extern void goLogStreamCallback(char* message, char* user);

void setLogStreamCallback(struct aiLogStream* stream);

void CallLogStream(struct aiLogStream* stream, char* message);