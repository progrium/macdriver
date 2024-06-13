// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a morphology maximum filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologymaximum?language=objc
type PMorphologyMaximum interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool
}

// ensure impl type implements protocol interface
var _ PMorphologyMaximum = (*MorphologyMaximumObject)(nil)

// A concrete type for the [PMorphologyMaximum] protocol.
type MorphologyMaximumObject struct {
	objc.Object
}

func (m_ MorphologyMaximumObject) HasSetInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologymaximum/3228577-inputimage?language=objc
func (m_ MorphologyMaximumObject) SetInputImage(value Image) {
	objc.Call[objc.Void](m_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (m_ MorphologyMaximumObject) HasInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologymaximum/3228577-inputimage?language=objc
func (m_ MorphologyMaximumObject) InputImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("inputImage"))
	return rv
}

func (m_ MorphologyMaximumObject) HasSetRadius() bool {
	return m_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the circular morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologymaximum/3228578-radius?language=objc
func (m_ MorphologyMaximumObject) SetRadius(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setRadius:"), value)
}

func (m_ MorphologyMaximumObject) HasRadius() bool {
	return m_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the circular morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologymaximum/3228578-radius?language=objc
func (m_ MorphologyMaximumObject) Radius() float64 {
	rv := objc.Call[float64](m_, objc.Sel("radius"))
	return rv
}