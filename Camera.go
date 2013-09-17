package assimp

//#cgo linux LDFLAGS: -L/usr/local/lib -lassimp -lstdc++
//#include <assimp/camera.h>
import "C"

type Camera C.struct_aiCamera

func (cam *Camera) Name() string {
	return C.GoStringN(&cam.mName.data[0], C.int(cam.mName.length))
}

func (cam *Camera) Position() Vector3 {
	return Vector3(cam.mPosition)
}

func (cam *Camera) Up() Vector3 {
	return Vector3(cam.mUp)
}

func (cam *Camera) LookAt() Vector3 {
	return Vector3(cam.mLookAt)
}

func (cam *Camera) HorizontalFov() float32 {
	return float32(cam.mHorizontalFOV)
}

func (cam *Camera) ClipPlaneNear() float32 {
	return float32(cam.mClipPlaneNear)
}

func (cam *Camera) ClipPlaneFar() float32 {
	return float32(cam.mClipPlaneFar)
}

func (cam *Camera) Aspect() float32 {
	return float32(cam.mAspect)
}
