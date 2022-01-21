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
type Color4 C.struct_aiColor4D
type Quaternion C.struct_aiQuaternion
type Matrix3x3 C.struct_aiMatrix3x3
type Matrix4x4 C.struct_aiMatrix4x4

func (v *Vector2) X() float32 {
	return float32(v.x)
}

func (v *Vector2) Y() float32 {
	return float32(v.y)
}

func (v *Vector2) Values() [2]float32 {
	return [2]float32{float32(v.x), float32(v.y)}
}

func (v *Vector3) X() float32 {
	return float32(v.x)
}

func (v *Vector3) Y() float32 {
	return float32(v.y)
}

func (v *Vector3) Z() float32 {
	return float32(v.z)
}

func (v *Vector3) Values() [3]float32 {
	return [3]float32{float32(v.x), float32(v.y), float32(v.z)}
}

func (v *Quaternion) W() float32 {
	return float32(v.w)
}

func (v *Quaternion) X() float32 {
	return float32(v.x)
}

func (v *Quaternion) Y() float32 {
	return float32(v.y)
}

func (v *Quaternion) Z() float32 {
	return float32(v.z)
}

// order w,x,y,z
func (q *Quaternion) Values() [4]float32 {
	return [4]float32{float32(q.w), float32(q.x), float32(q.y), float32(q.z)}
}

func (c *Color3) R() float32 {
	return float32(c.r)
}

func (c *Color3) G() float32 {
	return float32(c.g)
}

func (c *Color3) B() float32 {
	return float32(c.b)
}

func (c *Color3) Values() [3]float32 {
	return [3]float32{float32(c.r), float32(c.g), float32(c.b)}
}

func (c *Color4) R() float32 {
	return float32(c.r)
}

func (c *Color4) G() float32 {
	return float32(c.g)
}

func (c *Color4) B() float32 {
	return float32(c.b)
}

func (c *Color4) A() float32 {
	return float32(c.a)
}

func (c *Color4) Values() [4]float32 {
	return [4]float32{float32(c.r), float32(c.g), float32(c.b), float32(c.a)}
}

func (m *Matrix3x3) Values() [3][3]float32 {
	return [3][3]float32{
		[3]float32{float32(m.a1), float32(m.a2), float32(m.a3)},
		[3]float32{float32(m.b1), float32(m.b2), float32(m.b3)},
		[3]float32{float32(m.c1), float32(m.c2), float32(m.c3)},
	}
}

func (m *Matrix4x4) Values() [4][4]float32 {
	return [4][4]float32{
		[4]float32{float32(m.a1), float32(m.a2), float32(m.a3), float32(m.a4)},
		[4]float32{float32(m.b1), float32(m.b2), float32(m.b3), float32(m.b4)},
		[4]float32{float32(m.c1), float32(m.c2), float32(m.c3), float32(m.c4)},
		[4]float32{float32(m.d1), float32(m.d2), float32(m.d3), float32(m.d4)},
	}
}
