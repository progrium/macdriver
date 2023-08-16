// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VisualEffectView] class.
var VisualEffectViewClass = _VisualEffectViewClass{objc.GetClass("NSVisualEffectView")}

type _VisualEffectViewClass struct {
	objc.Class
}

// An interface definition for the [VisualEffectView] class.
type IVisualEffectView interface {
	IView
	State() VisualEffectState
	SetState(value VisualEffectState)
	IsEmphasized() bool
	SetEmphasized(value bool)
	MaskImage() Image
	SetMaskImage(value IImage)
	InteriorBackgroundStyle() BackgroundStyle
	BlendingMode() VisualEffectBlendingMode
	SetBlendingMode(value VisualEffectBlendingMode)
	Material() VisualEffectMaterial
	SetMaterial(value VisualEffectMaterial)
}

// A view that adds translucency and vibrancy effects to the views in your interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview?language=objc
type VisualEffectView struct {
	View
}

func VisualEffectViewFrom(ptr unsafe.Pointer) VisualEffectView {
	return VisualEffectView{
		View: ViewFrom(ptr),
	}
}

func (vc _VisualEffectViewClass) Alloc() VisualEffectView {
	rv := objc.Call[VisualEffectView](vc, objc.Sel("alloc"))
	return rv
}

func VisualEffectView_Alloc() VisualEffectView {
	return VisualEffectViewClass.Alloc()
}

func (vc _VisualEffectViewClass) New() VisualEffectView {
	rv := objc.Call[VisualEffectView](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVisualEffectView() VisualEffectView {
	return VisualEffectViewClass.New()
}

func (v_ VisualEffectView) Init() VisualEffectView {
	rv := objc.Call[VisualEffectView](v_, objc.Sel("init"))
	return rv
}

func (v_ VisualEffectView) InitWithFrame(frameRect foundation.Rect) VisualEffectView {
	rv := objc.Call[VisualEffectView](v_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func VisualEffectView_InitWithFrame(frameRect foundation.Rect) VisualEffectView {
	return VisualEffectViewClass.Alloc().InitWithFrame(frameRect)
}

// A value that indicates whether a view has a visual effect applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/1532403-state?language=objc
func (v_ VisualEffectView) State() VisualEffectState {
	rv := objc.Call[VisualEffectState](v_, objc.Sel("state"))
	return rv
}

// A value that indicates whether a view has a visual effect applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/1532403-state?language=objc
func (v_ VisualEffectView) SetState(value VisualEffectState) {
	objc.Call[objc.Void](v_, objc.Sel("setState:"), value)
}

// A Boolean value indicating whether to emphasize the look of the material. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/1644721-emphasized?language=objc
func (v_ VisualEffectView) IsEmphasized() bool {
	rv := objc.Call[bool](v_, objc.Sel("isEmphasized"))
	return rv
}

// A Boolean value indicating whether to emphasize the look of the material. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/1644721-emphasized?language=objc
func (v_ VisualEffectView) SetEmphasized(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setEmphasized:"), value)
}

// An image whose alpha channel masks the visual effect view's material. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/1535318-maskimage?language=objc
func (v_ VisualEffectView) MaskImage() Image {
	rv := objc.Call[Image](v_, objc.Sel("maskImage"))
	return rv
}

// An image whose alpha channel masks the visual effect view's material. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/1535318-maskimage?language=objc
func (v_ VisualEffectView) SetMaskImage(value IImage) {
	objc.Call[objc.Void](v_, objc.Sel("setMaskImage:"), objc.Ptr(value))
}

// The view’s interior background style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/1529998-interiorbackgroundstyle?language=objc
func (v_ VisualEffectView) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.Call[BackgroundStyle](v_, objc.Sel("interiorBackgroundStyle"))
	return rv
}

// A value indicating how the view’s contents blend with the surrounding content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/1535468-blendingmode?language=objc
func (v_ VisualEffectView) BlendingMode() VisualEffectBlendingMode {
	rv := objc.Call[VisualEffectBlendingMode](v_, objc.Sel("blendingMode"))
	return rv
}

// A value indicating how the view’s contents blend with the surrounding content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/1535468-blendingmode?language=objc
func (v_ VisualEffectView) SetBlendingMode(value VisualEffectBlendingMode) {
	objc.Call[objc.Void](v_, objc.Sel("setBlendingMode:"), value)
}

// The material shown by the visual effect view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/1526623-material?language=objc
func (v_ VisualEffectView) Material() VisualEffectMaterial {
	rv := objc.Call[VisualEffectMaterial](v_, objc.Sel("material"))
	return rv
}

// The material shown by the visual effect view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsvisualeffectview/1526623-material?language=objc
func (v_ VisualEffectView) SetMaterial(value VisualEffectMaterial) {
	objc.Call[objc.Void](v_, objc.Sel("setMaterial:"), value)
}
