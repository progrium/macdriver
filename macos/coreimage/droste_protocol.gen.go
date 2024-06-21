// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste?language=objc
type PDroste interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetPeriodicity(value float32)
	HasSetPeriodicity() bool

	// optional
	Periodicity() float32
	HasPeriodicity() bool

	// optional
	SetRotation(value float32)
	HasSetRotation() bool

	// optional
	Rotation() float32
	HasRotation() bool

	// optional
	SetInsetPoint0(value coregraphics.Point)
	HasSetInsetPoint0() bool

	// optional
	InsetPoint0() coregraphics.Point
	HasInsetPoint0() bool

	// optional
	SetZoom(value float32)
	HasSetZoom() bool

	// optional
	Zoom() float32
	HasZoom() bool

	// optional
	SetStrands(value float32)
	HasSetStrands() bool

	// optional
	Strands() float32
	HasStrands() bool

	// optional
	SetInsetPoint1(value coregraphics.Point)
	HasSetInsetPoint1() bool

	// optional
	InsetPoint1() coregraphics.Point
	HasInsetPoint1() bool
}

// ensure impl type implements protocol interface
var _ PDroste = (*DrosteObject)(nil)

// A concrete type for the [PDroste] protocol.
type DrosteObject struct {
	objc.Object
}

func (d_ DrosteObject) HasSetInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600131-inputimage?language=objc
func (d_ DrosteObject) SetInputImage(value Image) {
	objc.Call[objc.Void](d_, objc.Sel("setInputImage:"), value)
}

func (d_ DrosteObject) HasInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600131-inputimage?language=objc
func (d_ DrosteObject) InputImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("inputImage"))
	return rv
}

func (d_ DrosteObject) HasSetPeriodicity() bool {
	return d_.RespondsToSelector(objc.Sel("setPeriodicity:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600134-periodicity?language=objc
func (d_ DrosteObject) SetPeriodicity(value float32) {
	objc.Call[objc.Void](d_, objc.Sel("setPeriodicity:"), value)
}

func (d_ DrosteObject) HasPeriodicity() bool {
	return d_.RespondsToSelector(objc.Sel("periodicity"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600134-periodicity?language=objc
func (d_ DrosteObject) Periodicity() float32 {
	rv := objc.Call[float32](d_, objc.Sel("periodicity"))
	return rv
}

func (d_ DrosteObject) HasSetRotation() bool {
	return d_.RespondsToSelector(objc.Sel("setRotation:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600135-rotation?language=objc
func (d_ DrosteObject) SetRotation(value float32) {
	objc.Call[objc.Void](d_, objc.Sel("setRotation:"), value)
}

func (d_ DrosteObject) HasRotation() bool {
	return d_.RespondsToSelector(objc.Sel("rotation"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600135-rotation?language=objc
func (d_ DrosteObject) Rotation() float32 {
	rv := objc.Call[float32](d_, objc.Sel("rotation"))
	return rv
}

func (d_ DrosteObject) HasSetInsetPoint0() bool {
	return d_.RespondsToSelector(objc.Sel("setInsetPoint0:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600132-insetpoint0?language=objc
func (d_ DrosteObject) SetInsetPoint0(value coregraphics.Point) {
	objc.Call[objc.Void](d_, objc.Sel("setInsetPoint0:"), value)
}

func (d_ DrosteObject) HasInsetPoint0() bool {
	return d_.RespondsToSelector(objc.Sel("insetPoint0"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600132-insetpoint0?language=objc
func (d_ DrosteObject) InsetPoint0() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](d_, objc.Sel("insetPoint0"))
	return rv
}

func (d_ DrosteObject) HasSetZoom() bool {
	return d_.RespondsToSelector(objc.Sel("setZoom:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600137-zoom?language=objc
func (d_ DrosteObject) SetZoom(value float32) {
	objc.Call[objc.Void](d_, objc.Sel("setZoom:"), value)
}

func (d_ DrosteObject) HasZoom() bool {
	return d_.RespondsToSelector(objc.Sel("zoom"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600137-zoom?language=objc
func (d_ DrosteObject) Zoom() float32 {
	rv := objc.Call[float32](d_, objc.Sel("zoom"))
	return rv
}

func (d_ DrosteObject) HasSetStrands() bool {
	return d_.RespondsToSelector(objc.Sel("setStrands:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600136-strands?language=objc
func (d_ DrosteObject) SetStrands(value float32) {
	objc.Call[objc.Void](d_, objc.Sel("setStrands:"), value)
}

func (d_ DrosteObject) HasStrands() bool {
	return d_.RespondsToSelector(objc.Sel("strands"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600136-strands?language=objc
func (d_ DrosteObject) Strands() float32 {
	rv := objc.Call[float32](d_, objc.Sel("strands"))
	return rv
}

func (d_ DrosteObject) HasSetInsetPoint1() bool {
	return d_.RespondsToSelector(objc.Sel("setInsetPoint1:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600133-insetpoint1?language=objc
func (d_ DrosteObject) SetInsetPoint1(value coregraphics.Point) {
	objc.Call[objc.Void](d_, objc.Sel("setInsetPoint1:"), value)
}

func (d_ DrosteObject) HasInsetPoint1() bool {
	return d_.RespondsToSelector(objc.Sel("insetPoint1"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidroste/3600133-insetpoint1?language=objc
func (d_ DrosteObject) InsetPoint1() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](d_, objc.Sel("insetPoint1"))
	return rv
}
