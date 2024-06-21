// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a star-shine generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator?language=objc
type PStarShineGenerator interface {
	// optional
	SetCrossOpacity(value float32)
	HasSetCrossOpacity() bool

	// optional
	CrossOpacity() float32
	HasCrossOpacity() bool

	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() Color
	HasColor() bool

	// optional
	SetCrossAngle(value float32)
	HasSetCrossAngle() bool

	// optional
	CrossAngle() float32
	HasCrossAngle() bool

	// optional
	SetEpsilon(value float32)
	HasSetEpsilon() bool

	// optional
	Epsilon() float32
	HasEpsilon() bool

	// optional
	SetRadius(value float32)
	HasSetRadius() bool

	// optional
	Radius() float32
	HasRadius() bool

	// optional
	SetCrossScale(value float32)
	HasSetCrossScale() bool

	// optional
	CrossScale() float32
	HasCrossScale() bool

	// optional
	SetCrossWidth(value float32)
	HasSetCrossWidth() bool

	// optional
	CrossWidth() float32
	HasCrossWidth() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// ensure impl type implements protocol interface
var _ PStarShineGenerator = (*StarShineGeneratorObject)(nil)

// A concrete type for the [PStarShineGenerator] protocol.
type StarShineGeneratorObject struct {
	objc.Object
}

func (s_ StarShineGeneratorObject) HasSetCrossOpacity() bool {
	return s_.RespondsToSelector(objc.Sel("setCrossOpacity:"))
}

// The opacity of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228752-crossopacity?language=objc
func (s_ StarShineGeneratorObject) SetCrossOpacity(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setCrossOpacity:"), value)
}

func (s_ StarShineGeneratorObject) HasCrossOpacity() bool {
	return s_.RespondsToSelector(objc.Sel("crossOpacity"))
}

// The opacity of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228752-crossopacity?language=objc
func (s_ StarShineGeneratorObject) CrossOpacity() float32 {
	rv := objc.Call[float32](s_, objc.Sel("crossOpacity"))
	return rv
}

func (s_ StarShineGeneratorObject) HasSetColor() bool {
	return s_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color to use for the outer shell of the circular star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228750-color?language=objc
func (s_ StarShineGeneratorObject) SetColor(value Color) {
	objc.Call[objc.Void](s_, objc.Sel("setColor:"), value)
}

func (s_ StarShineGeneratorObject) HasColor() bool {
	return s_.RespondsToSelector(objc.Sel("color"))
}

// The color to use for the outer shell of the circular star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228750-color?language=objc
func (s_ StarShineGeneratorObject) Color() Color {
	rv := objc.Call[Color](s_, objc.Sel("color"))
	return rv
}

func (s_ StarShineGeneratorObject) HasSetCrossAngle() bool {
	return s_.RespondsToSelector(objc.Sel("setCrossAngle:"))
}

// The angle of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228751-crossangle?language=objc
func (s_ StarShineGeneratorObject) SetCrossAngle(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setCrossAngle:"), value)
}

func (s_ StarShineGeneratorObject) HasCrossAngle() bool {
	return s_.RespondsToSelector(objc.Sel("crossAngle"))
}

// The angle of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228751-crossangle?language=objc
func (s_ StarShineGeneratorObject) CrossAngle() float32 {
	rv := objc.Call[float32](s_, objc.Sel("crossAngle"))
	return rv
}

func (s_ StarShineGeneratorObject) HasSetEpsilon() bool {
	return s_.RespondsToSelector(objc.Sel("setEpsilon:"))
}

// The length of the cross spikes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228755-epsilon?language=objc
func (s_ StarShineGeneratorObject) SetEpsilon(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setEpsilon:"), value)
}

func (s_ StarShineGeneratorObject) HasEpsilon() bool {
	return s_.RespondsToSelector(objc.Sel("epsilon"))
}

// The length of the cross spikes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228755-epsilon?language=objc
func (s_ StarShineGeneratorObject) Epsilon() float32 {
	rv := objc.Call[float32](s_, objc.Sel("epsilon"))
	return rv
}

func (s_ StarShineGeneratorObject) HasSetRadius() bool {
	return s_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228756-radius?language=objc
func (s_ StarShineGeneratorObject) SetRadius(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setRadius:"), value)
}

func (s_ StarShineGeneratorObject) HasRadius() bool {
	return s_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228756-radius?language=objc
func (s_ StarShineGeneratorObject) Radius() float32 {
	rv := objc.Call[float32](s_, objc.Sel("radius"))
	return rv
}

func (s_ StarShineGeneratorObject) HasSetCrossScale() bool {
	return s_.RespondsToSelector(objc.Sel("setCrossScale:"))
}

// The size of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228753-crossscale?language=objc
func (s_ StarShineGeneratorObject) SetCrossScale(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setCrossScale:"), value)
}

func (s_ StarShineGeneratorObject) HasCrossScale() bool {
	return s_.RespondsToSelector(objc.Sel("crossScale"))
}

// The size of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228753-crossscale?language=objc
func (s_ StarShineGeneratorObject) CrossScale() float32 {
	rv := objc.Call[float32](s_, objc.Sel("crossScale"))
	return rv
}

func (s_ StarShineGeneratorObject) HasSetCrossWidth() bool {
	return s_.RespondsToSelector(objc.Sel("setCrossWidth:"))
}

// The width of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228754-crosswidth?language=objc
func (s_ StarShineGeneratorObject) SetCrossWidth(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setCrossWidth:"), value)
}

func (s_ StarShineGeneratorObject) HasCrossWidth() bool {
	return s_.RespondsToSelector(objc.Sel("crossWidth"))
}

// The width of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228754-crosswidth?language=objc
func (s_ StarShineGeneratorObject) CrossWidth() float32 {
	rv := objc.Call[float32](s_, objc.Sel("crossWidth"))
	return rv
}

func (s_ StarShineGeneratorObject) HasSetCenter() bool {
	return s_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228749-center?language=objc
func (s_ StarShineGeneratorObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setCenter:"), value)
}

func (s_ StarShineGeneratorObject) HasCenter() bool {
	return s_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228749-center?language=objc
func (s_ StarShineGeneratorObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](s_, objc.Sel("center"))
	return rv
}
