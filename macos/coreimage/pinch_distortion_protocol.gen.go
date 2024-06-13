// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion?language=objc
type PPinchDistortion interface {
	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

	// optional
	SetScale(value float64)
	HasSetScale() bool

	// optional
	Scale() float64
	HasScale() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PPinchDistortion = (*PinchDistortionObject)(nil)

// A concrete type for the [PPinchDistortion] protocol.
type PinchDistortionObject struct {
	objc.Object
}

func (p_ PinchDistortionObject) HasSetRadius() bool {
	return p_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600189-radius?language=objc
func (p_ PinchDistortionObject) SetRadius(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setRadius:"), value)
}

func (p_ PinchDistortionObject) HasRadius() bool {
	return p_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600189-radius?language=objc
func (p_ PinchDistortionObject) Radius() float64 {
	rv := objc.Call[float64](p_, objc.Sel("radius"))
	return rv
}

func (p_ PinchDistortionObject) HasSetScale() bool {
	return p_.RespondsToSelector(objc.Sel("setScale:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600190-scale?language=objc
func (p_ PinchDistortionObject) SetScale(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setScale:"), value)
}

func (p_ PinchDistortionObject) HasScale() bool {
	return p_.RespondsToSelector(objc.Sel("scale"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600190-scale?language=objc
func (p_ PinchDistortionObject) Scale() float64 {
	rv := objc.Call[float64](p_, objc.Sel("scale"))
	return rv
}

func (p_ PinchDistortionObject) HasSetCenter() bool {
	return p_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600187-center?language=objc
func (p_ PinchDistortionObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](p_, objc.Sel("setCenter:"), value)
}

func (p_ PinchDistortionObject) HasCenter() bool {
	return p_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600187-center?language=objc
func (p_ PinchDistortionObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](p_, objc.Sel("center"))
	return rv
}

func (p_ PinchDistortionObject) HasSetInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600188-inputimage?language=objc
func (p_ PinchDistortionObject) SetInputImage(value Image) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (p_ PinchDistortionObject) HasInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipinchdistortion/3600188-inputimage?language=objc
func (p_ PinchDistortionObject) InputImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("inputImage"))
	return rv
}