// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion?language=objc
type PDisplacementDistortion interface {
	// optional
	SetScale(value float32)
	HasSetScale() bool

	// optional
	Scale() float32
	HasScale() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetDisplacementImage(value Image)
	HasSetDisplacementImage() bool

	// optional
	DisplacementImage() Image
	HasDisplacementImage() bool
}

// ensure impl type implements protocol interface
var _ PDisplacementDistortion = (*DisplacementDistortionObject)(nil)

// A concrete type for the [PDisplacementDistortion] protocol.
type DisplacementDistortionObject struct {
	objc.Object
}

func (d_ DisplacementDistortionObject) HasSetScale() bool {
	return d_.RespondsToSelector(objc.Sel("setScale:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600129-scale?language=objc
func (d_ DisplacementDistortionObject) SetScale(value float32) {
	objc.Call[objc.Void](d_, objc.Sel("setScale:"), value)
}

func (d_ DisplacementDistortionObject) HasScale() bool {
	return d_.RespondsToSelector(objc.Sel("scale"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600129-scale?language=objc
func (d_ DisplacementDistortionObject) Scale() float32 {
	rv := objc.Call[float32](d_, objc.Sel("scale"))
	return rv
}

func (d_ DisplacementDistortionObject) HasSetInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600128-inputimage?language=objc
func (d_ DisplacementDistortionObject) SetInputImage(value Image) {
	objc.Call[objc.Void](d_, objc.Sel("setInputImage:"), value)
}

func (d_ DisplacementDistortionObject) HasInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600128-inputimage?language=objc
func (d_ DisplacementDistortionObject) InputImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("inputImage"))
	return rv
}

func (d_ DisplacementDistortionObject) HasSetDisplacementImage() bool {
	return d_.RespondsToSelector(objc.Sel("setDisplacementImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600127-displacementimage?language=objc
func (d_ DisplacementDistortionObject) SetDisplacementImage(value Image) {
	objc.Call[objc.Void](d_, objc.Sel("setDisplacementImage:"), value)
}

func (d_ DisplacementDistortionObject) HasDisplacementImage() bool {
	return d_.RespondsToSelector(objc.Sel("displacementImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisplacementdistortion/3600127-displacementimage?language=objc
func (d_ DisplacementDistortionObject) DisplacementImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("displacementImage"))
	return rv
}
