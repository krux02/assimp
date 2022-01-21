package assimp

//#cgo linux LDFLAGS: -L/usr/local/lib -lassimp -lstdc++
//#include <assimp/mesh.h>
//#include <string.h>
import "C"

import "unsafe"
import "reflect"

const (
	MaxFaceIndices           = C.AI_MAX_FACE_INDICES
	MaxBoneWeights           = C.AI_MAX_BONE_WEIGHTS
	MaxVertices              = C.AI_MAX_VERTICES
	MaxFaces                 = C.AI_MAX_FACES
	MaxNumberOfColorSets     = C.AI_MAX_NUMBER_OF_COLOR_SETS
	MaxNumberOfTextureCoords = C.AI_MAX_NUMBER_OF_TEXTURECOORDS
)

type Face C.struct_aiFace

func (f *Face) NumIndices() uint32 {
	return uint32(f.mNumIndices)
}

func (f *Face) IndicesPtr() *uint32 {
	return (*uint32)(f.mIndices)
}

func (f *Face) CopyIndices() []uint32 {
	indices := make([]uint32, f.mNumIndices)
	if unsafe.Sizeof(C.uint(0)) != unsafe.Sizeof(uint32(0)) {
		panic("wrong size")
	}
	size := uint32(unsafe.Sizeof(C.uint(0))) * f.NumIndices()
	C.memcpy(unsafe.Pointer(&indices[0]), unsafe.Pointer(f.mIndices), C.size_t(size))
	return indices
}

type VertexWeight C.struct_aiVertexWeight

func (w VertexWeight) VertexId() uint {
	return uint(w.mVertexId)
}

func (w VertexWeight) Weight() float32 {
	return float32(w.mWeight)
}

type Bone C.struct_aiBone

func (b *Bone) Name() string {
	return C.GoStringN(&b.mName.data[0], C.int(b.mName.length))
}

func (b *Bone) NumWeights() uint {
	return uint(b.mNumWeights)
}

func (b *Bone) WeightsPtr() *VertexWeight {
	return (*VertexWeight)(b.mWeights)
}

func (b *Bone) Weights() []VertexWeight {
	var result []VertexWeight
	header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
	header.Cap = int(b.NumWeights())
	header.Len = int(b.NumWeights())
	header.Data = uintptr(unsafe.Pointer(b.mWeights))
	return result
}

func (b *Bone) OffsetMatrix() Matrix4x4 {
	return Matrix4x4(b.mOffsetMatrix)
}

type PrimitiveType C.enum_aiPrimitiveType

const (
	PrimitiveType_Point    PrimitiveType = C.aiPrimitiveType_POINT
	PrimitiveType_Line     PrimitiveType = C.aiPrimitiveType_LINE
	PrimitiveType_Triangle PrimitiveType = C.aiPrimitiveType_TRIANGLE
	PrimitiveType_Polygon  PrimitiveType = C.aiPrimitiveType_POLYGON
)

type AnimMesh C.struct_aiAnimMesh

func (mesh *AnimMesh) VerticesPtr() *Vector3 {
	return (*Vector3)(mesh.mVertices)
}

func (mesh *AnimMesh) Vertices() []Vector3 {
	if mesh.HasPositions() {
		var result []Vector3
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mVertices))
		return result
	} else {
		return nil
	}
}

func (mesh *AnimMesh) NormalsPtr() *Vector3 {
	return (*Vector3)(mesh.mNormals)
}

func (mesh *AnimMesh) Normals() []Vector3 {
	if mesh.HasNormals() {
		var result []Vector3
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mNormals))
		return result
	} else {
		return nil
	}
}

func (mesh *AnimMesh) TangentsPtr() *Vector3 {
	return (*Vector3)(mesh.mTangents)
}

func (mesh *AnimMesh) Tangents() []Vector3 {
	if mesh.HasTangentsAndBitangents() {
		var result []Vector3
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mTangents))
		return result
	} else {
		return nil
	}
}

func (mesh *AnimMesh) BitangentsPtr() *Vector3 {
	return (*Vector3)(mesh.mBitangents)
}

func (mesh *AnimMesh) Bitangents() []Vector3 {
	if mesh.HasTangentsAndBitangents() {
		var result []Vector3
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mBitangents))
		return result
	} else {
		return nil
	}
}

func (mesh *AnimMesh) ColorsPtr(pIndex int) *Color4 {
	return (*Color4)(mesh.mColors[pIndex])
}

func (mesh *AnimMesh) Colors(pIndex int) []Color4 {
	if mesh.HasVertexColors(pIndex) {
		var result []Color4
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mColors[pIndex]))
		return result
	} else {
		return nil
	}
}

func (mesh *AnimMesh) TextureCoordsPtr(pIndex int) *Vector3 {
	return (*Vector3)(mesh.mTextureCoords[pIndex])
}

func (mesh *AnimMesh) TextureCoords(pIndex int) []Color4 {
	if mesh.HasTextureCoords(pIndex) {
		var result []Color4
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mTextureCoords[pIndex]))
		return result
	} else {
		return nil
	}
}

/** Check whether the anim mesh overrides the vertex positions
 *  of its host mesh*/
func (mesh *AnimMesh) HasPositions() bool {
	return mesh.mVertices != nil
}

/** Check whether the anim mesh overrides the vertex normals
 *  of its host mesh*/
func (mesh *AnimMesh) HasNormals() bool {
	return mesh.mNormals != nil
}

/** Check whether the anim mesh overrides the vertex tangents
 *  and bitangents of its host mesh. As for aiMesh,
 *  tangents and bitangents always go together. */
func (mesh *AnimMesh) HasTangentsAndBitangents() bool {
	return mesh.mTangents != nil
}

/** Check whether the anim mesh overrides a particular
 * set of vertex colors on his host mesh.
 *  @param pIndex 0<index<AI_MAX_NUMBER_OF_COLOR_SETS */
func (mesh *AnimMesh) HasVertexColors(pIndex int) bool {
	if pIndex >= MaxNumberOfColorSets {
		return false
	} else {
		return mesh.mColors[pIndex] != nil
	}
}

/** Check whether the anim mesh overrides a particular
 * set of texture coordinates on his host mesh.
 *  @param pIndex 0<index<AI_MAX_NUMBER_OF_TEXTURECOORDS */
func (mesh *AnimMesh) HasTextureCoords(pIndex int) bool {
	if pIndex >= MaxNumberOfTextureCoords {
		return false
	} else {
		return mesh.mTextureCoords[pIndex] != nil
	}
}

type Mesh C.struct_aiMesh

func (mesh *Mesh) PrimitiveTypes() int {
	return int(mesh.mPrimitiveTypes)
}

func (mesh *Mesh) NumVertices() int {
	return int(mesh.mNumVertices)
}

func (mesh *Mesh) NumFaces() int {
	return int(mesh.mNumFaces)
}

func (mesh *Mesh) Vertices() []Vector3 {
	if mesh.mVertices != nil {
		var result []Vector3
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mVertices))
		return result
	} else {
		return nil
	}
}

func (mesh *Mesh) Normals() []Vector3 {
	if mesh.mNormals != nil {
		var result []Vector3
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mNormals))
		return result
	} else {
		return nil
	}
}

func (mesh *Mesh) Tangents() []Vector3 {
	if mesh.mTangents != nil {
		var result []Vector3
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mTangents))
		return result
	} else {
		return nil
	}
}

func (mesh *Mesh) Bitangents() []Vector3 {
	if mesh.mBitangents != nil {
		var result []Vector3
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mBitangents))
		return result
	} else {
		return nil
	}
}

func (mesh *Mesh) TextureCoords(pIndex int) []Vector3 {
	if pIndex < MaxNumberOfTextureCoords && mesh.mTextureCoords[pIndex] != nil {
		var result []Vector3
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mTextureCoords[pIndex]))
		return result
	} else {
		return nil
	}
}

func (mesh *Mesh) Colors(pIndex int) []Color4 {
	if pIndex < MaxNumberOfColorSets && mesh.mColors[pIndex] != nil {
		var result []Color4
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumVertices)
		header.Len = int(mesh.mNumVertices)
		header.Data = uintptr(unsafe.Pointer(mesh.mColors[pIndex]))
		return result
	} else {
		return nil
	}

}

func (mesh *Mesh) NumUvComponents(pIndex int) int {
	return int(mesh.mNumUVComponents[pIndex])
}

func (mesh *Mesh) Faces() []Face {
	if mesh.mFaces != nil {
		var result []Face
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumFaces)
		header.Len = int(mesh.mNumFaces)
		header.Data = uintptr(unsafe.Pointer(mesh.mFaces))
		return result
	} else {
		return nil
	}
}

func (mesh *Mesh) NumBones() int {
	return int(mesh.mNumBones)
}

func (mesh *Mesh) Bones() []*Bone {
	if mesh.mBones != nil {
		var result []*Bone
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumBones)
		header.Len = int(mesh.mNumBones)
		header.Data = uintptr(unsafe.Pointer(mesh.mBones))
		return result
	} else {
		return nil
	}
}

func (mesh *Mesh) MaterialIndex() int {
	return int(mesh.mMaterialIndex)
}

func (mesh *Mesh) Name() string {
	return C.GoStringN(&mesh.mName.data[0], C.int(mesh.mName.length))
}

func (mesh *Mesh) NumAnimMeshes() int {
	return int(mesh.mNumAnimMeshes)
}

func (mesh *Mesh) AnimMeshes() []*AnimMesh {
	if mesh.mAnimMeshes != nil {
		var result []*AnimMesh
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mesh.mNumAnimMeshes)
		header.Len = int(mesh.mNumAnimMeshes)
		header.Data = uintptr(unsafe.Pointer(mesh.mAnimMeshes))
		return result
	} else {
		return nil
	}
}
