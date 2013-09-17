package assimp

//#cgo linux LDFLAGS: -L/usr/local/lib -lassimp -lstdc++
//#include <assimp/light.h>
import "C"

import "math"

type LightSourceType C.enum_aiLightSourceType

const (
	LightSource_Undefined   LightSourceType = C.aiLightSource_UNDEFINED
	LightSource_Directional LightSourceType = C.aiLightSource_DIRECTIONAL
	LightSource_Point       LightSourceType = C.aiLightSource_POINT
	LightSource_Spot        LightSourceType = C.aiLightSource_SPOT
)

type Light C.struct_aiLight

func (l *Light) Name() string {
	return C.GoStringN(&l.mName.data[0], C.int(l.mName.length))
}

func (l *Light) Type() LightSourceType {
	return LightSourceType(l.mType)
}

func (l *Light) Position() Vector3 {
	return Vector3(l.mPosition)
}

func (l *Light) Direction() Vector3 {
	return Vector3(l.mDirection)
}

func (l *Light) AttenuationConstant() float32 {
	return float32(l.mAttenuationConstant)
}

func (l *Light) AttenuationLinear() float32 {
	return float32(l.mAttenuationLinear)
}

func (l *Light) AttenuationQuadratic() float32 {
	return float32(l.mAttenuationQuadratic)
}

func (l *Light) ColorDiffuse() Color3 {
	return Color3(l.mColorDiffuse)
}

func (l *Light) ColorSpecular() Color3 {
	return Color3(l.mColorSpecular)
}

func (l *Light) ColorAmbient() Color3 {
	return Color3(l.mColorAmbient)
}

func (l *Light) AngleInnerCone() float32 {
	return float32(l.mAngleInnerCone)
}

func (l *Light) AngleOuterCone() float32 {
	return float32(l.mAngleOuterCone)
}

func NewLight() *Light {
	l := new(Light)
	l.mAttenuationConstant = 0
	l.mAttenuationLinear = 1
	l.mAttenuationQuadratic = 0
	l.mAngleInnerCone = math.Pi * 2
	l.mAngleOuterCone = math.Pi * 2
	return l
}
