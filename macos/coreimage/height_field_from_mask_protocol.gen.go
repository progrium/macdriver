// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a height-field-from-mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciheightfieldfrommask?language=objc
type PHeightFieldFromMask interface {
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
}

// ensure impl type implements protocol interface
var _ PHeightFieldFromMask = (*HeightFieldFromMaskObject)(nil)

// A concrete type for the [PHeightFieldFromMask] protocol.
type HeightFieldFromMaskObject struct {
	objc.Object
}

func (h_ HeightFieldFromMaskObject) HasSetInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciheightfieldfrommask/3228487-inputimage?language=objc
func (h_ HeightFieldFromMaskObject) SetInputImage(value Image) {
	objc.Call[objc.Void](h_, objc.Sel("setInputImage:"), value)
}

func (h_ HeightFieldFromMaskObject) HasInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciheightfieldfrommask/3228487-inputimage?language=objc
func (h_ HeightFieldFromMaskObject) InputImage() Image {
	rv := objc.Call[Image](h_, objc.Sel("inputImage"))
	return rv
}

func (h_ HeightFieldFromMaskObject) HasSetRadius() bool {
	return h_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The length of the height-field transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciheightfieldfrommask/3228488-radius?language=objc
func (h_ HeightFieldFromMaskObject) SetRadius(value float32) {
	objc.Call[objc.Void](h_, objc.Sel("setRadius:"), value)
}

func (h_ HeightFieldFromMaskObject) HasRadius() bool {
	return h_.RespondsToSelector(objc.Sel("radius"))
}

// The length of the height-field transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciheightfieldfrommask/3228488-radius?language=objc
func (h_ HeightFieldFromMaskObject) Radius() float32 {
	rv := objc.Call[float32](h_, objc.Sel("radius"))
	return rv
}
