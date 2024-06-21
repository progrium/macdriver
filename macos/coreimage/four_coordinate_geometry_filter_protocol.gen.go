// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a geometry adjustment filters that requires four coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter?language=objc
type PFourCoordinateGeometryFilter interface {
	// optional
	SetBottomRight(value coregraphics.Point)
	HasSetBottomRight() bool

	// optional
	BottomRight() coregraphics.Point
	HasBottomRight() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetBottomLeft(value coregraphics.Point)
	HasSetBottomLeft() bool

	// optional
	BottomLeft() coregraphics.Point
	HasBottomLeft() bool

	// optional
	SetTopRight(value coregraphics.Point)
	HasSetTopRight() bool

	// optional
	TopRight() coregraphics.Point
	HasTopRight() bool

	// optional
	SetTopLeft(value coregraphics.Point)
	HasSetTopLeft() bool

	// optional
	TopLeft() coregraphics.Point
	HasTopLeft() bool
}

// ensure impl type implements protocol interface
var _ PFourCoordinateGeometryFilter = (*FourCoordinateGeometryFilterObject)(nil)

// A concrete type for the [PFourCoordinateGeometryFilter] protocol.
type FourCoordinateGeometryFilterObject struct {
	objc.Object
}

func (f_ FourCoordinateGeometryFilterObject) HasSetBottomRight() bool {
	return f_.RespondsToSelector(objc.Sel("setBottomRight:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338732-bottomright?language=objc
func (f_ FourCoordinateGeometryFilterObject) SetBottomRight(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setBottomRight:"), value)
}

func (f_ FourCoordinateGeometryFilterObject) HasBottomRight() bool {
	return f_.RespondsToSelector(objc.Sel("bottomRight"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338732-bottomright?language=objc
func (f_ FourCoordinateGeometryFilterObject) BottomRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("bottomRight"))
	return rv
}

func (f_ FourCoordinateGeometryFilterObject) HasSetInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338733-inputimage?language=objc
func (f_ FourCoordinateGeometryFilterObject) SetInputImage(value Image) {
	objc.Call[objc.Void](f_, objc.Sel("setInputImage:"), value)
}

func (f_ FourCoordinateGeometryFilterObject) HasInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338733-inputimage?language=objc
func (f_ FourCoordinateGeometryFilterObject) InputImage() Image {
	rv := objc.Call[Image](f_, objc.Sel("inputImage"))
	return rv
}

func (f_ FourCoordinateGeometryFilterObject) HasSetBottomLeft() bool {
	return f_.RespondsToSelector(objc.Sel("setBottomLeft:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338731-bottomleft?language=objc
func (f_ FourCoordinateGeometryFilterObject) SetBottomLeft(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setBottomLeft:"), value)
}

func (f_ FourCoordinateGeometryFilterObject) HasBottomLeft() bool {
	return f_.RespondsToSelector(objc.Sel("bottomLeft"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338731-bottomleft?language=objc
func (f_ FourCoordinateGeometryFilterObject) BottomLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("bottomLeft"))
	return rv
}

func (f_ FourCoordinateGeometryFilterObject) HasSetTopRight() bool {
	return f_.RespondsToSelector(objc.Sel("setTopRight:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338735-topright?language=objc
func (f_ FourCoordinateGeometryFilterObject) SetTopRight(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setTopRight:"), value)
}

func (f_ FourCoordinateGeometryFilterObject) HasTopRight() bool {
	return f_.RespondsToSelector(objc.Sel("topRight"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338735-topright?language=objc
func (f_ FourCoordinateGeometryFilterObject) TopRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("topRight"))
	return rv
}

func (f_ FourCoordinateGeometryFilterObject) HasSetTopLeft() bool {
	return f_.RespondsToSelector(objc.Sel("setTopLeft:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338734-topleft?language=objc
func (f_ FourCoordinateGeometryFilterObject) SetTopLeft(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setTopLeft:"), value)
}

func (f_ FourCoordinateGeometryFilterObject) HasTopLeft() bool {
	return f_.RespondsToSelector(objc.Sel("topLeft"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338734-topleft?language=objc
func (f_ FourCoordinateGeometryFilterObject) TopLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("topLeft"))
	return rv
}
