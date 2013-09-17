package assimp

//#cgo linux LDFLAGS: -L/usr/local/lib -lassimp -lstdc++
//#include <assimp/scene.h>
import "C"

import (
	"reflect"
	"unsafe"
)

type Node C.struct_aiNode

func (this *Node) Name() string {
	return C.GoString(&this.mName.data[0])
}

func (this *Node) Transformation() Matrix4x4 {
	return Matrix4x4(this.mTransformation)
}

func (this *Node) Parent() *Node {
	return (*Node)(this.mParent)
}

func (this *Node) NumChildren() int {
	return int(this.mNumChildren)
}

func (this *Node) Children() []*Node {
	if this.mNumChildren > 0 && this.mChildren != nil {
		var result []*Node
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumChildren)
		header.Len = int(this.mNumChildren)
		header.Data = uintptr(unsafe.Pointer(this.mChildren))
		return result
	} else {
		return nil
	}
}

func (this *Node) NumMeshes() int {
	return int(this.mNumMeshes)
}

func (this *Node) Meshes() []int32 {
	if this.mNumMeshes > 0 && this.mMeshes != nil {
		var result []int32
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumMeshes)
		header.Len = int(this.mNumMeshes)
		header.Data = uintptr(unsafe.Pointer(this.mMeshes))
		return result
	} else {
		return nil
	}
}

const (
	SceneFlags_Incomplete        = C.AI_SCENE_FLAGS_INCOMPLETE
	SceneFlags_Validated         = C.AI_SCENE_FLAGS_VALIDATED
	SceneFlags_ValidationWarning = C.AI_SCENE_FLAGS_VALIDATION_WARNING
	SceneFlags_NonVerboseFormat  = C.AI_SCENE_FLAGS_NON_VERBOSE_FORMAT
	SceneFlags_Terrain           = C.AI_SCENE_FLAGS_TERRAIN
)

type Scene C.struct_aiScene

func (this *Scene) Flags() uint {
	return uint(this.mFlags)
}

func (this *Scene) RootNode() *Node {
	return (*Node)(this.mRootNode)
}

func (this *Scene) NumMeshes() int {
	return int(this.mNumMeshes)
}

func (this *Scene) Meshes() []*Mesh {
	if this.mNumMeshes > 0 && this.mMeshes != nil {
		var result []*Mesh
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumMeshes)
		header.Len = int(this.mNumMeshes)
		header.Data = uintptr(unsafe.Pointer(this.mMeshes))
		return result
	} else {
		return nil
	}
}

func (this *Scene) NumMaterials() int {
	return int(this.mNumMaterials)
}

func (this *Scene) Materials() []*Material {
	if this.mNumMaterials > 0 && this.mMaterials != nil {
		var result []*Material
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumMaterials)
		header.Len = int(this.mNumMaterials)
		header.Data = uintptr(unsafe.Pointer(this.mMaterials))
		return result
	} else {
		return nil
	}
}

func (this *Scene) NumAnimations() int {
	return int(this.mNumAnimations)
}

func (this *Scene) Animations() []*Animation {
	if this.mNumAnimations > 0 && this.mAnimations != nil {
		var result []*Animation
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumAnimations)
		header.Len = int(this.mNumAnimations)
		header.Data = uintptr(unsafe.Pointer(this.mAnimations))
		return result
	} else {
		return nil
	}
}

func (this *Scene) NumTextures() int {
	return int(this.mNumTextures)
}

func (this *Scene) Textures() []*Texture {
	if this.mNumTextures > 0 && this.mTextures != nil {
		var result []*Texture
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumTextures)
		header.Len = int(this.mNumTextures)
		header.Data = uintptr(unsafe.Pointer(this.mTextures))
		return result
	} else {
		return nil
	}
}

func (this *Scene) NumLights() int {
	return int(this.mNumLights)
}

func (this *Scene) Lights() []*Light {
	if this.mNumLights > 0 && this.mLights != nil {
		var result []*Light
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumLights)
		header.Len = int(this.mNumLights)
		header.Data = uintptr(unsafe.Pointer(this.mLights))
		return result
	} else {
		return nil
	}
}

func (this *Scene) NumCameras() int {
	return int(this.mNumCameras)
}

func (this *Scene) Cameras() []*Camera {
	if this.mNumCameras > 0 && this.mCameras != nil {
		var result []*Camera
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(this.mNumCameras)
		header.Len = int(this.mNumCameras)
		header.Data = uintptr(unsafe.Pointer(this.mCameras))
		return result
	} else {
		return nil
	}
}
