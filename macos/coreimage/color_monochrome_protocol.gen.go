// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a color monochrome filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome?language=objc
type PColorMonochrome interface {
	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() Color
	HasColor() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetIntensity(value float32)
	HasSetIntensity() bool

	// optional
	Intensity() float32
	HasIntensity() bool
}

// ensure impl type implements protocol interface
var _ PColorMonochrome = (*ColorMonochromeObject)(nil)

// A concrete type for the [PColorMonochrome] protocol.
type ColorMonochromeObject struct {
	objc.Object
}

func (c_ ColorMonochromeObject) HasSetColor() bool {
	return c_.RespondsToSelector(objc.Sel("setColor:"))
}

// The monochrome color to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228167-color?language=objc
func (c_ ColorMonochromeObject) SetColor(value Color) {
	objc.Call[objc.Void](c_, objc.Sel("setColor:"), value)
}

func (c_ ColorMonochromeObject) HasColor() bool {
	return c_.RespondsToSelector(objc.Sel("color"))
}

// The monochrome color to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228167-color?language=objc
func (c_ ColorMonochromeObject) Color() Color {
	rv := objc.Call[Color](c_, objc.Sel("color"))
	return rv
}

func (c_ ColorMonochromeObject) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228168-inputimage?language=objc
func (c_ ColorMonochromeObject) SetInputImage(value Image) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), value)
}

func (c_ ColorMonochromeObject) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228168-inputimage?language=objc
func (c_ ColorMonochromeObject) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorMonochromeObject) HasSetIntensity() bool {
	return c_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the monochrome effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228169-intensity?language=objc
func (c_ ColorMonochromeObject) SetIntensity(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setIntensity:"), value)
}

func (c_ ColorMonochromeObject) HasIntensity() bool {
	return c_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the monochrome effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228169-intensity?language=objc
func (c_ ColorMonochromeObject) Intensity() float32 {
	rv := objc.Call[float32](c_, objc.Sel("intensity"))
	return rv
}
