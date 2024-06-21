// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion?language=objc
type PTorusLensDistortion interface {
	// optional
	SetWidth(value float32)
	HasSetWidth() bool

	// optional
	Width() float32
	HasWidth() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetRefraction(value float32)
	HasSetRefraction() bool

	// optional
	Refraction() float32
	HasRefraction() bool

	// optional
	SetRadius(value float32)
	HasSetRadius() bool

	// optional
	Radius() float32
	HasRadius() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// ensure impl type implements protocol interface
var _ PTorusLensDistortion = (*TorusLensDistortionObject)(nil)

// A concrete type for the [PTorusLensDistortion] protocol.
type TorusLensDistortionObject struct {
	objc.Object
}

func (t_ TorusLensDistortionObject) HasSetWidth() bool {
	return t_.RespondsToSelector(objc.Sel("setWidth:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600201-width?language=objc
func (t_ TorusLensDistortionObject) SetWidth(value float32) {
	objc.Call[objc.Void](t_, objc.Sel("setWidth:"), value)
}

func (t_ TorusLensDistortionObject) HasWidth() bool {
	return t_.RespondsToSelector(objc.Sel("width"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600201-width?language=objc
func (t_ TorusLensDistortionObject) Width() float32 {
	rv := objc.Call[float32](t_, objc.Sel("width"))
	return rv
}

func (t_ TorusLensDistortionObject) HasSetInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600198-inputimage?language=objc
func (t_ TorusLensDistortionObject) SetInputImage(value Image) {
	objc.Call[objc.Void](t_, objc.Sel("setInputImage:"), value)
}

func (t_ TorusLensDistortionObject) HasInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600198-inputimage?language=objc
func (t_ TorusLensDistortionObject) InputImage() Image {
	rv := objc.Call[Image](t_, objc.Sel("inputImage"))
	return rv
}

func (t_ TorusLensDistortionObject) HasSetRefraction() bool {
	return t_.RespondsToSelector(objc.Sel("setRefraction:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600200-refraction?language=objc
func (t_ TorusLensDistortionObject) SetRefraction(value float32) {
	objc.Call[objc.Void](t_, objc.Sel("setRefraction:"), value)
}

func (t_ TorusLensDistortionObject) HasRefraction() bool {
	return t_.RespondsToSelector(objc.Sel("refraction"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600200-refraction?language=objc
func (t_ TorusLensDistortionObject) Refraction() float32 {
	rv := objc.Call[float32](t_, objc.Sel("refraction"))
	return rv
}

func (t_ TorusLensDistortionObject) HasSetRadius() bool {
	return t_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600199-radius?language=objc
func (t_ TorusLensDistortionObject) SetRadius(value float32) {
	objc.Call[objc.Void](t_, objc.Sel("setRadius:"), value)
}

func (t_ TorusLensDistortionObject) HasRadius() bool {
	return t_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600199-radius?language=objc
func (t_ TorusLensDistortionObject) Radius() float32 {
	rv := objc.Call[float32](t_, objc.Sel("radius"))
	return rv
}

func (t_ TorusLensDistortionObject) HasSetCenter() bool {
	return t_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600197-center?language=objc
func (t_ TorusLensDistortionObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](t_, objc.Sel("setCenter:"), value)
}

func (t_ TorusLensDistortionObject) HasCenter() bool {
	return t_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citoruslensdistortion/3600197-center?language=objc
func (t_ TorusLensDistortionObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("center"))
	return rv
}
