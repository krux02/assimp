package assimp

//#cgo linux LDFLAGS: -L/usr/local/lib -lassimp -lstdc++
//#include <stdlib.h>
//#include "Import.h"
import "C"

import (
	"reflect"
	"unsafe"
)

//export goLogStreamCallback
func goLogStreamCallback(message *C.char, user *C.char) {
	ls := *(*LogStream)(unsafe.Pointer(user))
	str := C.GoString(message)
	ls.Log(str)
}

type LogStream interface {
	Log(message string)
}

type PropertyStore C.struct_aiPropertyStore

func ImportFile(file string, flags uint) *Scene {
	pFile := C.CString(file)
	defer C.free(unsafe.Pointer(pFile))
	return (*Scene)(C.aiImportFile(pFile, C.uint(flags)))
}

func ImportFileEx(file string, flags uint)

//garbage collection prevention
type logStream C.struct_aiLogStream

var logStreams []*logStream

func (log *logStream) Log(message string) {
	str := C.CString(message)
	defer C.free(unsafe.Pointer(str))
	C.CallLogStream((*C.struct_aiLogStream)(log), str)
}

func ImportFileFromMemory(data []byte, flags uint, hint string) *Scene {
	pBuffer := (*C.char)(unsafe.Pointer(((*reflect.SliceHeader)(unsafe.Pointer(&data))).Data))
	pLength := C.uint(len(data))
	pFlags := C.uint(flags)
	pHint := C.CString(hint)
	defer C.free(unsafe.Pointer(pHint))
	return (*Scene)(C.aiImportFileFromMemory(pBuffer, pLength, pFlags, pHint))
}

func ImportFileFromMemoryWithProperties(data []byte, flags uint, hint string, props *PropertyStore) *Scene {
	pBuffer := (*C.char)(unsafe.Pointer(((*reflect.SliceHeader)(unsafe.Pointer(&data))).Data))
	pLength := C.uint(len(data))
	pFlags := C.uint(flags)
	pHint := C.CString(hint)
	defer C.free(unsafe.Pointer(pHint))
	pProps := (*C.struct_aiPropertyStore)(props)
	return (*Scene)(C.aiImportFileFromMemoryWithProperties(pBuffer, pLength, pFlags, pHint, pProps))
}

func (scene *Scene) ApplyPostProcessing(flags uint) *Scene {
	pScene := (*C.struct_aiScene)(scene)
	pFlags := C.uint(flags)

	return (*Scene)(C.aiApplyPostProcessing(pScene, pFlags))
}

func GetPredefinedLogStream(streams DefaultLogStream, file string) LogStream {
	pStreams := C.enum_aiDefaultLogStream(streams)
	pFile := C.CString(file)
	defer C.free(unsafe.Pointer(pFile))
	ls := logStream(C.aiGetPredefinedLogStream(pStreams, pFile))
	return &ls
}

func AttachLogStream(stream LogStream) {

	if ls, ok := stream.(*logStream); ok {
		// if the stream is a wrapped C-compatible aiLogStream unwrap it
		C.aiAttachLogStream((*C.struct_aiLogStream)(ls))
	} else {
		// this i a log Stream implemented in go and needs to be wrapped to be C-compatible
		ls := new(logStream)
		ls.user = (*C.char)(unsafe.Pointer(&stream))
		C.setLogStreamCallback((*C.struct_aiLogStream)(ls))
		logStreams = append(logStreams, ls)
		C.aiAttachLogStream((*C.struct_aiLogStream)(ls))
	}
}

func EnableVerboseLogging(d bool) {
	if d {
		C.aiEnableVerboseLogging(C.AI_TRUE)
	} else {
		C.aiEnableVerboseLogging(C.AI_FALSE)
	}
}

func DetachLogStream(stream LogStream) Return {
	if ls, ok := stream.(*logStream); ok {
		// if the stream is a wrapped C-compatible aiLogStream unwrap it
		return Return(C.aiDetachLogStream((*C.struct_aiLogStream)(ls)))
	} else {
		for i, ls := range logStreams {
			ls2 := *(*LogStream)(unsafe.Pointer(ls.user))
			if ls2 == stream {
				logStreams = append(logStreams[:i], logStreams[i+1:]...)
				return Return(C.aiDetachLogStream((*C.struct_aiLogStream)(ls)))
			}
		}
		return Return_Failure
	}
}

func DetachAllLogStreams() {
	C.aiDetachAllLogStreams()
	logStreams = make([]*logStream, 0)
}

func (scene *Scene) ReleaseImport() {
	C.aiReleaseImport((*C.struct_aiScene)(scene))
}

func GetErrorString() string {
	return C.GoString(C.aiGetErrorString())
}

func IsExtensionSupported(extension string) bool {
	szExtension := C.CString(extension)
	defer C.free(unsafe.Pointer(szExtension))
	switch C.aiIsExtensionSupported(szExtension) {
	case C.AI_FALSE:
		return false
	default:
		return true
	}
}

func GetExtensionList() string {
	var str C.struct_aiString
	C.aiGetExtensionList(&str)
	return C.GoString(&str.data[0])
}

func (scene *Scene) GetMemoryRequirements() MemoryInfo {
	var meminfo C.struct_aiMemoryInfo
	C.aiGetMemoryRequirements((*C.struct_aiScene)(scene), &meminfo)
	return MemoryInfo(meminfo)
}

func CreatePropertyStore() *PropertyStore {
	return (*PropertyStore)(C.aiCreatePropertyStore())
}

func (p *PropertyStore) Release() {
	C.aiReleasePropertyStore((*C.struct_aiPropertyStore)(p))
}

func (p *PropertyStore) SetPropertyInteger(name string, value int) {
	szName := C.CString(name)
	defer C.free(unsafe.Pointer(szName))
	C.aiSetImportPropertyInteger((*C.struct_aiPropertyStore)(p), szName, C.int(value))
}

func (p *PropertyStore) SetPropertyFloat(name string, value float32) {
	szName := C.CString(name)
	defer C.free(unsafe.Pointer(szName))
	C.aiSetImportPropertyFloat((*C.struct_aiPropertyStore)(p), szName, C.float(value))
}

func (p *PropertyStore) SetPropertyString(name string, value string) {
	szName := C.CString(name)
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(szName))
	defer C.free(unsafe.Pointer(cValue))

	var st C.struct_aiString
	C.strncpy(&st.data[0], cValue, C.MAXLEN)
	C.aiSetImportPropertyString((*C.struct_aiPropertyStore)(p), szName, &st)
}

func (quat *Quaternion) FromMatrix(mat *Matrix3x3) {
	C.aiCreateQuaternionFromMatrix((*C.struct_aiQuaternion)(quat), (*C.struct_aiMatrix3x3)(mat))
}

func (mat *Matrix4x4) Decompose(scaling *Vector3, rotation *Quaternion, position *Vector3) {
	C.aiDecomposeMatrix(
		(*C.struct_aiMatrix4x4)(mat),
		(*C.struct_aiVector3D)(scaling),
		(*C.struct_aiQuaternion)(rotation),
		(*C.struct_aiVector3D)(position),
	)
}

func (mat *Matrix4x4) Transpose() {
	C.aiTransposeMatrix4((*C.struct_aiMatrix4x4)(mat))
}

func (mat *Matrix3x3) Transpose() {
	C.aiTransposeMatrix3((*C.struct_aiMatrix3x3)(mat))
}

func (vec *Vector3) TransformByMatrix3(mat *Matrix3x3) {
	C.aiTransformVecByMatrix3(
		(*C.struct_aiVector3D)(vec),
		(*C.struct_aiMatrix3x3)(mat),
	)
}

func (vec *Vector3) TransformByMatrix4(mat *Matrix4x4) {
	C.aiTransformVecByMatrix4(
		(*C.struct_aiVector3D)(vec),
		(*C.struct_aiMatrix4x4)(mat),
	)
}

func (dst *Matrix4x4) Multiply(src *Matrix4x4) {
	C.aiMultiplyMatrix4(
		(*C.struct_aiMatrix4x4)(dst),
		(*C.struct_aiMatrix4x4)(src),
	)
}

func (dst *Matrix3x3) Multiply(src *Matrix3x3) {
	C.aiMultiplyMatrix3(
		(*C.struct_aiMatrix3x3)(dst),
		(*C.struct_aiMatrix3x3)(src),
	)
}

func (dst *Matrix3x3) Identity() {
	C.aiIdentityMatrix3((*C.struct_aiMatrix3x3)(dst))
}

func (dst *Matrix4x4) Identity() {
	C.aiIdentityMatrix4((*C.struct_aiMatrix4x4)(dst))
}
