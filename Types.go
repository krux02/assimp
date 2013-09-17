package assimp

//#cgo linux LDFLAGS: -L/usr/local/lib -lassimp -lstdc++
//#include <assimp/types.h>
import "C"

type Plane C.struct_aiPlane
type Ray C.struct_aiRay
type Color3 C.struct_aiColor3D

type Return C.enum_aiReturn

const (
	Return_Success     Return = C.aiReturn_SUCCESS
	Return_Failure     Return = C.aiReturn_FAILURE
	Return_OutOfMemory Return = C.aiReturn_OUTOFMEMORY
)

type Origin C.enum_aiOrigin

const (
	Origin_Set Origin = C.aiOrigin_SET
	Origin_Cur Origin = C.aiOrigin_CUR
	Origin_End Origin = C.aiOrigin_END
)

type DefaultLogStream C.enum_aiDefaultLogStream

const (
	DefaultLogStream_File     DefaultLogStream = C.aiDefaultLogStream_FILE
	DefaultLogStream_StdOut   DefaultLogStream = C.aiDefaultLogStream_STDOUT
	DefaultLogStream_StdErr   DefaultLogStream = C.aiDefaultLogStream_STDERR
	DefaultLogStream_Debugger DefaultLogStream = C.aiDefaultLogStream_DEBUGGER
)

type MemoryInfo C.struct_aiMemoryInfo

func (mi *MemoryInfo) Textures() uint {
	return uint(mi.textures)
}

func (mi *MemoryInfo) Materials() uint {
	return uint(mi.materials)
}

func (mi *MemoryInfo) Meshes() uint {
	return uint(mi.meshes)
}

func (mi *MemoryInfo) Nodes() uint {
	return uint(mi.nodes)
}
func (mi *MemoryInfo) Animations() uint {
	return uint(mi.animations)
}
func (mi *MemoryInfo) Cameras() uint {
	return uint(mi.cameras)
}

func (mi *MemoryInfo) Lights() uint {
	return uint(mi.lights)
}

func (mi *MemoryInfo) Total() uint {
	return uint(mi.total)
}

type Vector2 C.struct_aiVector2D
type Vector3 C.struct_aiVector3D
type Color4  C.struct_aiColor4D
type Quaternion C.struct_aiQuaternion
type Matrix3x3 C.struct_aiMatrix3x3
type Matrix4x4 C.struct_aiMatrix4x4
