// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a line overlay filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay?language=objc
type PLineOverlay interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetEdgeIntensity(value float32)
	HasSetEdgeIntensity() bool

	// optional
	EdgeIntensity() float32
	HasEdgeIntensity() bool

	// optional
	SetNRNoiseLevel(value float32)
	HasSetNRNoiseLevel() bool

	// optional
	NRNoiseLevel() float32
	HasNRNoiseLevel() bool

	// optional
	SetNRSharpness(value float32)
	HasSetNRSharpness() bool

	// optional
	NRSharpness() float32
	HasNRSharpness() bool

	// optional
	SetContrast(value float32)
	HasSetContrast() bool

	// optional
	Contrast() float32
	HasContrast() bool

	// optional
	SetThreshold(value float32)
	HasSetThreshold() bool

	// optional
	Threshold() float32
	HasThreshold() bool
}

// ensure impl type implements protocol interface
var _ PLineOverlay = (*LineOverlayObject)(nil)

// A concrete type for the [PLineOverlay] protocol.
type LineOverlayObject struct {
	objc.Object
}

func (l_ LineOverlayObject) HasSetInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228533-inputimage?language=objc
func (l_ LineOverlayObject) SetInputImage(value Image) {
	objc.Call[objc.Void](l_, objc.Sel("setInputImage:"), value)
}

func (l_ LineOverlayObject) HasInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228533-inputimage?language=objc
func (l_ LineOverlayObject) InputImage() Image {
	rv := objc.Call[Image](l_, objc.Sel("inputImage"))
	return rv
}

func (l_ LineOverlayObject) HasSetEdgeIntensity() bool {
	return l_.RespondsToSelector(objc.Sel("setEdgeIntensity:"))
}

// The accentuation factor of the Sobel gradient information when tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228532-edgeintensity?language=objc
func (l_ LineOverlayObject) SetEdgeIntensity(value float32) {
	objc.Call[objc.Void](l_, objc.Sel("setEdgeIntensity:"), value)
}

func (l_ LineOverlayObject) HasEdgeIntensity() bool {
	return l_.RespondsToSelector(objc.Sel("edgeIntensity"))
}

// The accentuation factor of the Sobel gradient information when tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228532-edgeintensity?language=objc
func (l_ LineOverlayObject) EdgeIntensity() float32 {
	rv := objc.Call[float32](l_, objc.Sel("edgeIntensity"))
	return rv
}

func (l_ LineOverlayObject) HasSetNRNoiseLevel() bool {
	return l_.RespondsToSelector(objc.Sel("setNRNoiseLevel:"))
}

// The noise level of the image, used with camera data, that's removed before tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228529-nrnoiselevel?language=objc
func (l_ LineOverlayObject) SetNRNoiseLevel(value float32) {
	objc.Call[objc.Void](l_, objc.Sel("setNRNoiseLevel:"), value)
}

func (l_ LineOverlayObject) HasNRNoiseLevel() bool {
	return l_.RespondsToSelector(objc.Sel("NRNoiseLevel"))
}

// The noise level of the image, used with camera data, that's removed before tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228529-nrnoiselevel?language=objc
func (l_ LineOverlayObject) NRNoiseLevel() float32 {
	rv := objc.Call[float32](l_, objc.Sel("NRNoiseLevel"))
	return rv
}

func (l_ LineOverlayObject) HasSetNRSharpness() bool {
	return l_.RespondsToSelector(objc.Sel("setNRSharpness:"))
}

// The amount of sharpening done when removing noise in the image before tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228530-nrsharpness?language=objc
func (l_ LineOverlayObject) SetNRSharpness(value float32) {
	objc.Call[objc.Void](l_, objc.Sel("setNRSharpness:"), value)
}

func (l_ LineOverlayObject) HasNRSharpness() bool {
	return l_.RespondsToSelector(objc.Sel("NRSharpness"))
}

// The amount of sharpening done when removing noise in the image before tracing the edges of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228530-nrsharpness?language=objc
func (l_ LineOverlayObject) NRSharpness() float32 {
	rv := objc.Call[float32](l_, objc.Sel("NRSharpness"))
	return rv
}

func (l_ LineOverlayObject) HasSetContrast() bool {
	return l_.RespondsToSelector(objc.Sel("setContrast:"))
}

// The amount of antialiasing to use on the edges produced by this filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228531-contrast?language=objc
func (l_ LineOverlayObject) SetContrast(value float32) {
	objc.Call[objc.Void](l_, objc.Sel("setContrast:"), value)
}

func (l_ LineOverlayObject) HasContrast() bool {
	return l_.RespondsToSelector(objc.Sel("contrast"))
}

// The amount of antialiasing to use on the edges produced by this filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228531-contrast?language=objc
func (l_ LineOverlayObject) Contrast() float32 {
	rv := objc.Call[float32](l_, objc.Sel("contrast"))
	return rv
}

func (l_ LineOverlayObject) HasSetThreshold() bool {
	return l_.RespondsToSelector(objc.Sel("setThreshold:"))
}

// A value that determines edge visibility. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228534-threshold?language=objc
func (l_ LineOverlayObject) SetThreshold(value float32) {
	objc.Call[objc.Void](l_, objc.Sel("setThreshold:"), value)
}

func (l_ LineOverlayObject) HasThreshold() bool {
	return l_.RespondsToSelector(objc.Sel("threshold"))
}

// A value that determines edge visibility. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineoverlay/3228534-threshold?language=objc
func (l_ LineOverlayObject) Threshold() float32 {
	rv := objc.Call[float32](l_, objc.Sel("threshold"))
	return rv
}
