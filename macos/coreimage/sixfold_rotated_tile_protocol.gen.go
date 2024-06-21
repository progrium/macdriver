// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a sixfold rotated tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldrotatedtile?language=objc
type PSixfoldRotatedTile interface {
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
	SetAngle(value float32)
	HasSetAngle() bool

	// optional
	Angle() float32
	HasAngle() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// ensure impl type implements protocol interface
var _ PSixfoldRotatedTile = (*SixfoldRotatedTileObject)(nil)

// A concrete type for the [PSixfoldRotatedTile] protocol.
type SixfoldRotatedTileObject struct {
	objc.Object
}

func (s_ SixfoldRotatedTileObject) HasSetWidth() bool {
	return s_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldrotatedtile/3228721-width?language=objc
func (s_ SixfoldRotatedTileObject) SetWidth(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setWidth:"), value)
}

func (s_ SixfoldRotatedTileObject) HasWidth() bool {
	return s_.RespondsToSelector(objc.Sel("width"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldrotatedtile/3228721-width?language=objc
func (s_ SixfoldRotatedTileObject) Width() float32 {
	rv := objc.Call[float32](s_, objc.Sel("width"))
	return rv
}

func (s_ SixfoldRotatedTileObject) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldrotatedtile/3228720-inputimage?language=objc
func (s_ SixfoldRotatedTileObject) SetInputImage(value Image) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), value)
}

func (s_ SixfoldRotatedTileObject) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldrotatedtile/3228720-inputimage?language=objc
func (s_ SixfoldRotatedTileObject) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}

func (s_ SixfoldRotatedTileObject) HasSetAngle() bool {
	return s_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldrotatedtile/3228718-angle?language=objc
func (s_ SixfoldRotatedTileObject) SetAngle(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setAngle:"), value)
}

func (s_ SixfoldRotatedTileObject) HasAngle() bool {
	return s_.RespondsToSelector(objc.Sel("angle"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldrotatedtile/3228718-angle?language=objc
func (s_ SixfoldRotatedTileObject) Angle() float32 {
	rv := objc.Call[float32](s_, objc.Sel("angle"))
	return rv
}

func (s_ SixfoldRotatedTileObject) HasSetCenter() bool {
	return s_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldrotatedtile/3228719-center?language=objc
func (s_ SixfoldRotatedTileObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setCenter:"), value)
}

func (s_ SixfoldRotatedTileObject) HasCenter() bool {
	return s_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldrotatedtile/3228719-center?language=objc
func (s_ SixfoldRotatedTileObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](s_, objc.Sel("center"))
	return rv
}
