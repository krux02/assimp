#include "Import.h"
#include <assimp/cimport.h>

void setLogStreamCallback(C_STRUCT aiLogStream* ls){
	ls->callback = (aiLogStreamCallback)goLogStreamCallback;
}

void CallLogStream(struct aiLogStream* stream, char* message) {
	stream->callback(message, stream->user);
}
