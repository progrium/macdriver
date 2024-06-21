// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a page-curl-with-shadow transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition?language=objc
type PPageCurlWithShadowTransition interface {
	// optional
	SetShadowExtent(value coregraphics.Rect)
	HasSetShadowExtent() bool

	// optional
	ShadowExtent() coregraphics.Rect
	HasShadowExtent() bool

	// optional
	SetShadowAmount(value float32)
	HasSetShadowAmount() bool

	// optional
	ShadowAmount() float32
	HasShadowAmount() bool

	// optional
	SetRadius(value float32)
	HasSetRadius() bool

	// optional
	Radius() float32
	HasRadius() bool

	// optional
	SetExtent(value coregraphics.Rect)
	HasSetExtent() bool

	// optional
	Extent() coregraphics.Rect
	HasExtent() bool

	// optional
	SetAngle(value float32)
	HasSetAngle() bool

	// optional
	Angle() float32
	HasAngle() bool

	// optional
	SetShadowSize(value float32)
	HasSetShadowSize() bool

	// optional
	ShadowSize() float32
	HasShadowSize() bool

	// optional
	SetBacksideImage(value Image)
	HasSetBacksideImage() bool

	// optional
	BacksideImage() Image
	HasBacksideImage() bool
}

// ensure impl type implements protocol interface
var _ PPageCurlWithShadowTransition = (*PageCurlWithShadowTransitionObject)(nil)

// A concrete type for the [PPageCurlWithShadowTransition] protocol.
type PageCurlWithShadowTransitionObject struct {
	objc.Object
}

func (p_ PageCurlWithShadowTransitionObject) HasSetShadowExtent() bool {
	return p_.RespondsToSelector(objc.Sel("setShadowExtent:"))
}

// The rectagular portion of input image that casts a shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228629-shadowextent?language=objc
func (p_ PageCurlWithShadowTransitionObject) SetShadowExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](p_, objc.Sel("setShadowExtent:"), value)
}

func (p_ PageCurlWithShadowTransitionObject) HasShadowExtent() bool {
	return p_.RespondsToSelector(objc.Sel("shadowExtent"))
}

// The rectagular portion of input image that casts a shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228629-shadowextent?language=objc
func (p_ PageCurlWithShadowTransitionObject) ShadowExtent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](p_, objc.Sel("shadowExtent"))
	return rv
}

func (p_ PageCurlWithShadowTransitionObject) HasSetShadowAmount() bool {
	return p_.RespondsToSelector(objc.Sel("setShadowAmount:"))
}

// The strength of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228628-shadowamount?language=objc
func (p_ PageCurlWithShadowTransitionObject) SetShadowAmount(value float32) {
	objc.Call[objc.Void](p_, objc.Sel("setShadowAmount:"), value)
}

func (p_ PageCurlWithShadowTransitionObject) HasShadowAmount() bool {
	return p_.RespondsToSelector(objc.Sel("shadowAmount"))
}

// The strength of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228628-shadowamount?language=objc
func (p_ PageCurlWithShadowTransitionObject) ShadowAmount() float32 {
	rv := objc.Call[float32](p_, objc.Sel("shadowAmount"))
	return rv
}

func (p_ PageCurlWithShadowTransitionObject) HasSetRadius() bool {
	return p_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the curl. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228627-radius?language=objc
func (p_ PageCurlWithShadowTransitionObject) SetRadius(value float32) {
	objc.Call[objc.Void](p_, objc.Sel("setRadius:"), value)
}

func (p_ PageCurlWithShadowTransitionObject) HasRadius() bool {
	return p_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the curl. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228627-radius?language=objc
func (p_ PageCurlWithShadowTransitionObject) Radius() float32 {
	rv := objc.Call[float32](p_, objc.Sel("radius"))
	return rv
}

func (p_ PageCurlWithShadowTransitionObject) HasSetExtent() bool {
	return p_.RespondsToSelector(objc.Sel("setExtent:"))
}

// The extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228626-extent?language=objc
func (p_ PageCurlWithShadowTransitionObject) SetExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](p_, objc.Sel("setExtent:"), value)
}

func (p_ PageCurlWithShadowTransitionObject) HasExtent() bool {
	return p_.RespondsToSelector(objc.Sel("extent"))
}

// The extent of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228626-extent?language=objc
func (p_ PageCurlWithShadowTransitionObject) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](p_, objc.Sel("extent"))
	return rv
}

func (p_ PageCurlWithShadowTransitionObject) HasSetAngle() bool {
	return p_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the curling page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228624-angle?language=objc
func (p_ PageCurlWithShadowTransitionObject) SetAngle(value float32) {
	objc.Call[objc.Void](p_, objc.Sel("setAngle:"), value)
}

func (p_ PageCurlWithShadowTransitionObject) HasAngle() bool {
	return p_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the curling page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228624-angle?language=objc
func (p_ PageCurlWithShadowTransitionObject) Angle() float32 {
	rv := objc.Call[float32](p_, objc.Sel("angle"))
	return rv
}

func (p_ PageCurlWithShadowTransitionObject) HasSetShadowSize() bool {
	return p_.RespondsToSelector(objc.Sel("setShadowSize:"))
}

// The maximum size, in pixels, of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228630-shadowsize?language=objc
func (p_ PageCurlWithShadowTransitionObject) SetShadowSize(value float32) {
	objc.Call[objc.Void](p_, objc.Sel("setShadowSize:"), value)
}

func (p_ PageCurlWithShadowTransitionObject) HasShadowSize() bool {
	return p_.RespondsToSelector(objc.Sel("shadowSize"))
}

// The maximum size, in pixels, of the shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228630-shadowsize?language=objc
func (p_ PageCurlWithShadowTransitionObject) ShadowSize() float32 {
	rv := objc.Call[float32](p_, objc.Sel("shadowSize"))
	return rv
}

func (p_ PageCurlWithShadowTransitionObject) HasSetBacksideImage() bool {
	return p_.RespondsToSelector(objc.Sel("setBacksideImage:"))
}

// The image that appears on the back of the source image as the page curls to reveal the target image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228625-backsideimage?language=objc
func (p_ PageCurlWithShadowTransitionObject) SetBacksideImage(value Image) {
	objc.Call[objc.Void](p_, objc.Sel("setBacksideImage:"), value)
}

func (p_ PageCurlWithShadowTransitionObject) HasBacksideImage() bool {
	return p_.RespondsToSelector(objc.Sel("backsideImage"))
}

// The image that appears on the back of the source image as the page curls to reveal the target image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipagecurlwithshadowtransition/3228625-backsideimage?language=objc
func (p_ PageCurlWithShadowTransitionObject) BacksideImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("backsideImage"))
	return rv
}
