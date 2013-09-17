package assimp

/*
#cgo linux LDFLAGS: -L/usr/local/lib -lassimp -lstdc++
#include <assimp/material.h>
#include <stdlib.h>
*/
import "C"

import (
	"reflect"
	"unsafe"
)

type TextureOp C.enum_aiTextureOp

const (
	TextureOp_Multiply  TextureOp = C.aiTextureOp_Multiply
	TextureOp_Add       TextureOp = C.aiTextureOp_Add
	TextureOp_Subtract  TextureOp = C.aiTextureOp_Subtract
	TextureOp_Divide    TextureOp = C.aiTextureOp_Divide
	TextureOp_SmoothAdd TextureOp = C.aiTextureOp_SmoothAdd
	TextureOp_SignedAdd TextureOp = C.aiTextureOp_SignedAdd
)

type TextureMapMode C.enum_aiTextureMapMode

const (
	TextureMapMode_Wrap   TextureMapMode = C.aiTextureMapMode_Wrap
	TextureMapMode_Clamp  TextureMapMode = C.aiTextureMapMode_Clamp
	TextureMapMode_Decal  TextureMapMode = C.aiTextureMapMode_Decal
	TextureMapMode_Mirror TextureMapMode = C.aiTextureMapMode_Mirror
)

type TextureMapping C.enum_aiTextureMapping

const (
	TextureMapping_Uv       TextureMapping = C.aiTextureMapping_UV
	TextureMapping_Spere    TextureMapping = C.aiTextureMapping_SPHERE
	TextureMapping_Cylinder TextureMapping = C.aiTextureMapping_CYLINDER
	TextureMapping_Box      TextureMapping = C.aiTextureMapping_BOX
	TextureMapping_Plane    TextureMapping = C.aiTextureMapping_PLANE
	TextureMapping_Other    TextureMapping = C.aiTextureMapping_OTHER
)

type TextureType C.enum_aiTextureType

const (
	TextureMapping_None         TextureMapping = C.aiTextureType_NONE
	TextureMapping_Diffuse      TextureMapping = C.aiTextureType_DIFFUSE
	TextureMapping_Specular     TextureMapping = C.aiTextureType_SPECULAR
	TextureMapping_Ambient      TextureMapping = C.aiTextureType_AMBIENT
	TextureMapping_Emissive     TextureMapping = C.aiTextureType_EMISSIVE
	TextureMapping_Height       TextureMapping = C.aiTextureType_HEIGHT
	TextureMapping_Normals      TextureMapping = C.aiTextureType_NORMALS
	TextureMapping_Shininess    TextureMapping = C.aiTextureType_SHININESS
	TextureMapping_Opacity      TextureMapping = C.aiTextureType_OPACITY
	TextureMapping_Displacement TextureMapping = C.aiTextureType_DISPLACEMENT
	TextureMapping_Lightmap     TextureMapping = C.aiTextureType_LIGHTMAP
	TextureMapping_Reflection   TextureMapping = C.aiTextureType_REFLECTION
	TextureMapping_Unknown      TextureMapping = C.aiTextureType_UNKNOWN
)

const TextureTypeMax = C.AI_TEXTURE_TYPE_MAX

type ShadingMode C.enum_aiShadingMode

const (
	ShadingMode_Flat         ShadingMode = C.aiShadingMode_Flat
	ShadingMode_Gouraud      ShadingMode = C.aiShadingMode_Gouraud
	ShadingMode_Phong        ShadingMode = C.aiShadingMode_Phong
	ShadingMode_Blinn        ShadingMode = C.aiShadingMode_Blinn
	ShadingMode_Tonn         ShadingMode = C.aiShadingMode_Toon
	ShadingMode_OrenNayar    ShadingMode = C.aiShadingMode_OrenNayar
	ShadingMode_Minnaert     ShadingMode = C.aiShadingMode_Minnaert
	ShadingMode_CookTorrance ShadingMode = C.aiShadingMode_CookTorrance
	ShadingMode_NoShading    ShadingMode = C.aiShadingMode_NoShading
	ShadingMode_Fresnel      ShadingMode = C.aiShadingMode_Fresnel
)

type TextureFlags C.enum_aiTextureFlags

const (
	TextureFlags_Invert      TextureFlags = C.aiTextureFlags_Invert
	TextureFlags_UseAlpha    TextureFlags = C.aiTextureFlags_UseAlpha
	TextureFlags_IgnoreAlpha TextureFlags = C.aiTextureFlags_IgnoreAlpha
)

type BlendMode C.enum_aiBlendMode

const (
	BlendMode_Default  BlendMode = C.aiBlendMode_Default
	BlendMode_Additive BlendMode = C.aiBlendMode_Additive
)

type UvTransform C.struct_aiUVTransform

func (tf UvTransform) Translation() Vector2 {
	return Vector2(tf.mTranslation)
}

func (tf UvTransform) Scaling() Vector2 {
	return Vector2(tf.mScaling)
}

func (tf UvTransform) Rotation() float32 {
	return float32(tf.mRotation)
}

type PropertyTypeInfo C.enum_aiPropertyTypeInfo

const (
	PTI_Float   PropertyTypeInfo = C.aiPTI_Float
	PTI_String  PropertyTypeInfo = C.aiPTI_String
	PTI_Integer PropertyTypeInfo = C.aiPTI_Integer
	PTI_Buffer  PropertyTypeInfo = C.aiPTI_Buffer
)

type MaterialProperty C.struct_aiMaterialProperty

func (mp *MaterialProperty) Key() string {
	return C.GoStringN(&mp.mKey.data[0], C.int(mp.mKey.length))
}

func (mp *MaterialProperty) Semantic() int {
	return int(mp.mSemantic)
}

func (mp *MaterialProperty) Index() int {
	return int(mp.mIndex)
}

func (mp *MaterialProperty) DataLength() int {
	return int(mp.mDataLength)
}

func (mp *MaterialProperty) Type() PropertyTypeInfo {
	return PropertyTypeInfo(mp.mType)
}

func (mp *MaterialProperty) Data() []byte {
	if mp.mData != nil {
		var result []byte
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(mp.mDataLength)
		header.Len = int(mp.mDataLength)
		header.Data = uintptr(unsafe.Pointer(mp.mData))
		return result
	} else {
		return nil
	}
}

type Material C.struct_aiMaterial

func (m *Material) NumProperties() int {
	return int(m.mNumProperties)
}

func (m *Material) NumAllocated() int {
	return int(m.mNumAllocated)
}

func (m *Material) Properties() []MaterialProperty {
	if m.mProperties != nil {
		var result []MaterialProperty
		header := (*reflect.SliceHeader)(unsafe.Pointer(&result))
		header.Cap = int(m.mNumProperties)
		header.Len = int(m.mNumProperties)
		header.Data = uintptr(unsafe.Pointer(m.mProperties))
		return result
	} else {
		return nil
	}
}

type MatKey string

const (
	MatKey_Name                  MatKey = "?mat.name"
	MatKey_TwoSidid              MatKey = "$mat.twosided"
	MatKey_ShadingModel          MatKey = "$mat.shadingm"
	MatKey_EnableWireframe       MatKey = "$mat.wireframe"
	MatKey_BlendFunc             MatKey = "$mat.blend"
	MatKey_Opacity               MatKey = "$mat.opacity"
	MatKey_BumpScaling           MatKey = "$mat.bumpscaling"
	MatKey_Shininess             MatKey = "$mat.shininess"
	MatKey_Reflectivity          MatKey = "$mat.reflectivity"
	MatKey_ShininessStregth      MatKey = "$mat.shinpercent"
	MatKey_Refracti              MatKey = "$mat.refracti"
	MatKey_ColorDiffuse          MatKey = "$clr.diffuse"
	MatKey_ColorAmbient          MatKey = "$clr.ambient"
	MatKey_ColorSpecular         MatKey = "$clr.specular"
	MatKey_ColorEmissive         MatKey = "$clr.emissive"
	MatKey_ColorTransparent      MatKey = "$clr.transparent"
	MatKey_ColorReflective       MatKey = "$clr.reflective"
	MatKey_GlobalBackgroundImage MatKey = "?bg.global"
)

const (
	MatKey_Texture      MatKey = C._AI_MATKEY_TEXTURE_BASE
	MatKey_UvwSrc       MatKey = C._AI_MATKEY_UVWSRC_BASE
	MatKey_TexOp        MatKey = C._AI_MATKEY_TEXOP_BASE
	MatKey_Mapping      MatKey = C._AI_MATKEY_MAPPING_BASE
	MatKey_TexBlend     MatKey = C._AI_MATKEY_TEXBLEND_BASE
	MatKey_MappingModeU MatKey = C._AI_MATKEY_MAPPINGMODE_U_BASE
	MatKey_MappingModeV MatKey = C._AI_MATKEY_MAPPINGMODE_V_BASE
	MatKey_TexMapAxis   MatKey = C._AI_MATKEY_TEXMAP_AXIS_BASE
	MatKey_UvTransform  MatKey = C._AI_MATKEY_UVTRANSFORM_BASE
	MatKey_TexFlags     MatKey = C._AI_MATKEY_TEXFLAGS_BASE
)

func (mk MatKey) constString() *C.char {
	header := (*reflect.StringHeader)(unsafe.Pointer(&mk))
	return (*C.char)(unsafe.Pointer(header.Data))
}

func (m *Material) GetMaterialProperty(key MatKey, typ TextureType, textureIndex int) (*MaterialProperty, Return) {
	pKey := key.constString()
	var pPropOut *C.struct_aiMaterialProperty
	ret := C.aiGetMaterialProperty((*C.struct_aiMaterial)(m), pKey, C.uint(typ), C.uint(textureIndex), &pPropOut)

	return (*MaterialProperty)(pPropOut), Return(ret)
}

func (m *Material) GetMaterialFloatArray(key MatKey, typ TextureType, textureIndex int, pOut []float32) Return {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&pOut))
	ret := C.aiGetMaterialFloatArray((*C.struct_aiMaterial)(m), key.constString(), C.uint(typ), C.uint(textureIndex), (*C.float)(&pOut[0]), (*C.uint)(unsafe.Pointer((&header.Len))))
	return Return(ret)
}

func (m *Material) GetMaterialFloat(key MatKey, typ TextureType, textureIndex int) (float32, Return) {
	var f float32
	ret := C.aiGetMaterialFloatArray((*C.struct_aiMaterial)(m), key.constString(), C.uint(typ), C.uint(textureIndex), (*C.float)(&f), nil)
	return f, Return(ret)
}

func (m *Material) GetMaterialInteger(key MatKey, typ TextureType, textureIndex int) (int, Return) {
	var i C.int
	ret := C.aiGetMaterialIntegerArray((*C.struct_aiMaterial)(m), key.constString(), C.uint(typ), C.uint(textureIndex), &i, nil)
	return int(i), Return(ret)
}

func (m *Material) GetMaterialIntegerArray(key MatKey, typ TextureType, textureIndex int, pOut []int32) Return {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&pOut))
	ret := C.aiGetMaterialIntegerArray((*C.struct_aiMaterial)(m), key.constString(), C.uint(typ), C.uint(textureIndex), (*C.int)(&pOut[0]), (*C.uint)(unsafe.Pointer((&header.Len))))
	return Return(ret)
}

func (m *Material) GetMaterialColor(key MatKey, typ TextureType, textureIndex int) (Color4, Return) {
	var color C.struct_aiColor4D
	ret := C.aiGetMaterialColor((*C.struct_aiMaterial)(m), key.constString(), C.uint(typ), C.uint(textureIndex), &color)
	return Color4(color), Return(ret)
}

func (m *Material) GetMaterialString(key MatKey, typ TextureType, textureIndex int) (string,Return) {
	var str C.struct_aiString
	ret := C.aiGetMaterialString((*C.struct_aiMaterial)(m), key.constString(), C.uint(typ), C.uint(textureIndex), &str)
	return C.GoString(&str.data[0]),Return(ret)
}

func (m *Material) GetMaterialTextureCount(typ TextureType) int {
	ret := C.aiGetMaterialTextureCount((*C.struct_aiMaterial)(m), C.enum_aiTextureType(typ))
	return int(ret)
}

func (m *Material) GetMaterialTexture(typ TextureType, textureIndex int) (string, TextureMapping, int, float32, TextureOp, TextureMapMode, uint, Return) {
	var mat *C.struct_aiMaterial = (*C.struct_aiMaterial)(m)
	var typ_ C.enum_aiTextureType = C.enum_aiTextureType(typ)
	var index C.uint = C.uint(textureIndex)

	var path C.struct_aiString
	var mapping C.enum_aiTextureMapping
	var uvindex C.uint
	var blend C.float
	var op C.enum_aiTextureOp
	var mapmode C.enum_aiTextureMapMode
	var flags C.uint
	ret := C.aiGetMaterialTexture(mat, typ_, index, &path, &mapping, &uvindex, &blend, &op, &mapmode, &flags)
	return C.GoString(&path.data[0]), TextureMapping(mapping), int(uvindex), float32(blend), TextureOp(op), TextureMapMode(mapmode), uint(flags), Return(ret)
}
