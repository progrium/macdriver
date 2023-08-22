// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/coreimage"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FilterUIView] class.
var FilterUIViewClass = _FilterUIViewClass{objc.GetClass("IKFilterUIView")}

type _FilterUIViewClass struct {
	objc.Class
}

// An interface definition for the [FilterUIView] class.
type IFilterUIView interface {
	appkit.IView
	ObjectController() appkit.ObjectController
	Filter() coreimage.Filter
	InitWithFrameFilter(frameRect foundation.Rect, inFilter coreimage.IFilter) objc.Object
}

// The IKFilterUIView class provides a view that contains input parameter controls for a Core Image filter (CIFilter). You need to use this class when providing a user interface for a custom filter. The class creates a view that has an object controller for the given filter. It also retains the filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilteruiview?language=objc
type FilterUIView struct {
	appkit.View
}

func FilterUIViewFrom(ptr unsafe.Pointer) FilterUIView {
	return FilterUIView{
		View: appkit.ViewFrom(ptr),
	}
}

func (fc _FilterUIViewClass) Alloc() FilterUIView {
	rv := objc.Call[FilterUIView](fc, objc.Sel("alloc"))
	return rv
}

func FilterUIView_Alloc() FilterUIView {
	return FilterUIViewClass.Alloc()
}

func (fc _FilterUIViewClass) New() FilterUIView {
	rv := objc.Call[FilterUIView](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFilterUIView() FilterUIView {
	return FilterUIViewClass.New()
}

func (f_ FilterUIView) Init() FilterUIView {
	rv := objc.Call[FilterUIView](f_, objc.Sel("init"))
	return rv
}

func (f_ FilterUIView) InitWithFrame(frameRect foundation.Rect) FilterUIView {
	rv := objc.Call[FilterUIView](f_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewFilterUIViewWithFrame(frameRect foundation.Rect) FilterUIView {
	instance := FilterUIViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Returns the object controller for the bindings between the filter and its view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilteruiview/1504206-objectcontroller?language=objc
func (f_ FilterUIView) ObjectController() appkit.ObjectController {
	rv := objc.Call[appkit.ObjectController](f_, objc.Sel("objectController"))
	return rv
}

// Returns the Core Image filter associated with the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilteruiview/1504622-filter?language=objc
func (f_ FilterUIView) Filter() coreimage.Filter {
	rv := objc.Call[coreimage.Filter](f_, objc.Sel("filter"))
	return rv
}

// Initializes a view that contains controls for the input parameters of a filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilteruiview/1504139-initwithframe?language=objc
func (f_ FilterUIView) InitWithFrameFilter(frameRect foundation.Rect, inFilter coreimage.IFilter) objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("initWithFrame:filter:"), frameRect, objc.Ptr(inFilter))
	return rv
}

// Creates a view that contains controls for the input parameters of a filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilteruiview/1504872-viewwithframe?language=objc
func (fc _FilterUIViewClass) ViewWithFrameFilter(frameRect foundation.Rect, inFilter coreimage.IFilter) objc.Object {
	rv := objc.Call[objc.Object](fc, objc.Sel("viewWithFrame:filter:"), frameRect, objc.Ptr(inFilter))
	return rv
}

// Creates a view that contains controls for the input parameters of a filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilteruiview/1504872-viewwithframe?language=objc
func FilterUIView_ViewWithFrameFilter(frameRect foundation.Rect, inFilter coreimage.IFilter) objc.Object {
	return FilterUIViewClass.ViewWithFrameFilter(frameRect, inFilter)
}
