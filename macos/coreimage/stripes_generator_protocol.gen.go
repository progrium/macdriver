// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a stripes generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator?language=objc
type PStripesGenerator interface {
	// optional
	SetWidth(value float32)
	HasSetWidth() bool

	// optional
	Width() float32
	HasWidth() bool

	// optional
	SetSharpness(value float32)
	HasSetSharpness() bool

	// optional
	Sharpness() float32
	HasSharpness() bool

	// optional
	SetColor0(value Color)
	HasSetColor0() bool

	// optional
	Color0() Color
	HasColor0() bool

	// optional
	SetColor1(value Color)
	HasSetColor1() bool

	// optional
	Color1() Color
	HasColor1() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// ensure impl type implements protocol interface
var _ PStripesGenerator = (*StripesGeneratorObject)(nil)

// A concrete type for the [PStripesGenerator] protocol.
type StripesGeneratorObject struct {
	objc.Object
}

func (s_ StripesGeneratorObject) HasSetWidth() bool {
	return s_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a stripe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228765-width?language=objc
func (s_ StripesGeneratorObject) SetWidth(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setWidth:"), value)
}

func (s_ StripesGeneratorObject) HasWidth() bool {
	return s_.RespondsToSelector(objc.Sel("width"))
}

// The width of a stripe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228765-width?language=objc
func (s_ StripesGeneratorObject) Width() float32 {
	rv := objc.Call[float32](s_, objc.Sel("width"))
	return rv
}

func (s_ StripesGeneratorObject) HasSetSharpness() bool {
	return s_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The sharpness of the stripe pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228764-sharpness?language=objc
func (s_ StripesGeneratorObject) SetSharpness(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setSharpness:"), value)
}

func (s_ StripesGeneratorObject) HasSharpness() bool {
	return s_.RespondsToSelector(objc.Sel("sharpness"))
}

// The sharpness of the stripe pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228764-sharpness?language=objc
func (s_ StripesGeneratorObject) Sharpness() float32 {
	rv := objc.Call[float32](s_, objc.Sel("sharpness"))
	return rv
}

func (s_ StripesGeneratorObject) HasSetColor0() bool {
	return s_.RespondsToSelector(objc.Sel("setColor0:"))
}

// A color to use for the odd stripes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228762-color0?language=objc
func (s_ StripesGeneratorObject) SetColor0(value Color) {
	objc.Call[objc.Void](s_, objc.Sel("setColor0:"), value)
}

func (s_ StripesGeneratorObject) HasColor0() bool {
	return s_.RespondsToSelector(objc.Sel("color0"))
}

// A color to use for the odd stripes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228762-color0?language=objc
func (s_ StripesGeneratorObject) Color0() Color {
	rv := objc.Call[Color](s_, objc.Sel("color0"))
	return rv
}

func (s_ StripesGeneratorObject) HasSetColor1() bool {
	return s_.RespondsToSelector(objc.Sel("setColor1:"))
}

// A color to use for the even stripes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228763-color1?language=objc
func (s_ StripesGeneratorObject) SetColor1(value Color) {
	objc.Call[objc.Void](s_, objc.Sel("setColor1:"), value)
}

func (s_ StripesGeneratorObject) HasColor1() bool {
	return s_.RespondsToSelector(objc.Sel("color1"))
}

// A color to use for the even stripes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228763-color1?language=objc
func (s_ StripesGeneratorObject) Color1() Color {
	rv := objc.Call[Color](s_, objc.Sel("color1"))
	return rv
}

func (s_ StripesGeneratorObject) HasSetCenter() bool {
	return s_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the stripe pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228761-center?language=objc
func (s_ StripesGeneratorObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setCenter:"), value)
}

func (s_ StripesGeneratorObject) HasCenter() bool {
	return s_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the stripe pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228761-center?language=objc
func (s_ StripesGeneratorObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](s_, objc.Sel("center"))
	return rv
}
