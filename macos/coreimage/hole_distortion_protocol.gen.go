// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion?language=objc
type PHoleDistortion interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

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
var _ PHoleDistortion = (*HoleDistortionObject)(nil)

// A concrete type for the [PHoleDistortion] protocol.
type HoleDistortionObject struct {
	objc.Object
}

func (h_ HoleDistortionObject) HasSetInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600168-inputimage?language=objc
func (h_ HoleDistortionObject) SetInputImage(value Image) {
	objc.Call[objc.Void](h_, objc.Sel("setInputImage:"), value)
}

func (h_ HoleDistortionObject) HasInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600168-inputimage?language=objc
func (h_ HoleDistortionObject) InputImage() Image {
	rv := objc.Call[Image](h_, objc.Sel("inputImage"))
	return rv
}

func (h_ HoleDistortionObject) HasSetRadius() bool {
	return h_.RespondsToSelector(objc.Sel("setRadius:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600169-radius?language=objc
func (h_ HoleDistortionObject) SetRadius(value float32) {
	objc.Call[objc.Void](h_, objc.Sel("setRadius:"), value)
}

func (h_ HoleDistortionObject) HasRadius() bool {
	return h_.RespondsToSelector(objc.Sel("radius"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600169-radius?language=objc
func (h_ HoleDistortionObject) Radius() float32 {
	rv := objc.Call[float32](h_, objc.Sel("radius"))
	return rv
}

func (h_ HoleDistortionObject) HasSetCenter() bool {
	return h_.RespondsToSelector(objc.Sel("setCenter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600167-center?language=objc
func (h_ HoleDistortionObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](h_, objc.Sel("setCenter:"), value)
}

func (h_ HoleDistortionObject) HasCenter() bool {
	return h_.RespondsToSelector(objc.Sel("center"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciholedistortion/3600167-center?language=objc
func (h_ HoleDistortionObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](h_, objc.Sel("center"))
	return rv
}
