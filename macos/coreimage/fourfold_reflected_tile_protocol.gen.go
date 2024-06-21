// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a fourfold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile?language=objc
type PFourfoldReflectedTile interface {
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
	SetAcuteAngle(value float32)
	HasSetAcuteAngle() bool

	// optional
	AcuteAngle() float32
	HasAcuteAngle() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// ensure impl type implements protocol interface
var _ PFourfoldReflectedTile = (*FourfoldReflectedTileObject)(nil)

// A concrete type for the [PFourfoldReflectedTile] protocol.
type FourfoldReflectedTileObject struct {
	objc.Object
}

func (f_ FourfoldReflectedTileObject) HasSetWidth() bool {
	return f_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228448-width?language=objc
func (f_ FourfoldReflectedTileObject) SetWidth(value float32) {
	objc.Call[objc.Void](f_, objc.Sel("setWidth:"), value)
}

func (f_ FourfoldReflectedTileObject) HasWidth() bool {
	return f_.RespondsToSelector(objc.Sel("width"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228448-width?language=objc
func (f_ FourfoldReflectedTileObject) Width() float32 {
	rv := objc.Call[float32](f_, objc.Sel("width"))
	return rv
}

func (f_ FourfoldReflectedTileObject) HasSetInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228447-inputimage?language=objc
func (f_ FourfoldReflectedTileObject) SetInputImage(value Image) {
	objc.Call[objc.Void](f_, objc.Sel("setInputImage:"), value)
}

func (f_ FourfoldReflectedTileObject) HasInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228447-inputimage?language=objc
func (f_ FourfoldReflectedTileObject) InputImage() Image {
	rv := objc.Call[Image](f_, objc.Sel("inputImage"))
	return rv
}

func (f_ FourfoldReflectedTileObject) HasSetAngle() bool {
	return f_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228445-angle?language=objc
func (f_ FourfoldReflectedTileObject) SetAngle(value float32) {
	objc.Call[objc.Void](f_, objc.Sel("setAngle:"), value)
}

func (f_ FourfoldReflectedTileObject) HasAngle() bool {
	return f_.RespondsToSelector(objc.Sel("angle"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228445-angle?language=objc
func (f_ FourfoldReflectedTileObject) Angle() float32 {
	rv := objc.Call[float32](f_, objc.Sel("angle"))
	return rv
}

func (f_ FourfoldReflectedTileObject) HasSetAcuteAngle() bool {
	return f_.RespondsToSelector(objc.Sel("setAcuteAngle:"))
}

// The primary angle for the repeating reflected tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228444-acuteangle?language=objc
func (f_ FourfoldReflectedTileObject) SetAcuteAngle(value float32) {
	objc.Call[objc.Void](f_, objc.Sel("setAcuteAngle:"), value)
}

func (f_ FourfoldReflectedTileObject) HasAcuteAngle() bool {
	return f_.RespondsToSelector(objc.Sel("acuteAngle"))
}

// The primary angle for the repeating reflected tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228444-acuteangle?language=objc
func (f_ FourfoldReflectedTileObject) AcuteAngle() float32 {
	rv := objc.Call[float32](f_, objc.Sel("acuteAngle"))
	return rv
}

func (f_ FourfoldReflectedTileObject) HasSetCenter() bool {
	return f_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228446-center?language=objc
func (f_ FourfoldReflectedTileObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setCenter:"), value)
}

func (f_ FourfoldReflectedTileObject) HasCenter() bool {
	return f_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228446-center?language=objc
func (f_ FourfoldReflectedTileObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("center"))
	return rv
}
