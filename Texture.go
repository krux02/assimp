package assimp

//#cgo linux LDFLAGS: -L/usr/local/lib -lassimp -lstdc++
//#include <assimp/scene.h>
import "C"

type Texel C.struct_aiTexel

type Texture C.struct_aiTexture

func (t *Texture) Width() uint32 {
	return uint32(t.mWidth)
}

func (t *Texture) Height() uint32 {
	return uint32(t.mHeight)
}

func (t *Texture) Data() *Texel {
	return (*Texel)(t.pcData)
}
