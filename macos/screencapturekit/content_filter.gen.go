// Code generated by DarwinKit. DO NOT EDIT.

package screencapturekit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContentFilter] class.
var ContentFilterClass = _ContentFilterClass{objc.GetClass("SCContentFilter")}

type _ContentFilterClass struct {
	objc.Class
}

// An interface definition for the [ContentFilter] class.
type IContentFilter interface {
	objc.IObject
}

// An object that filters the content a stream captures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter?language=objc
type ContentFilter struct {
	objc.Object
}

func ContentFilterFrom(ptr unsafe.Pointer) ContentFilter {
	return ContentFilter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ ContentFilter) InitWithDisplayIncludingWindows(display IDisplay, includedWindows []IWindow) ContentFilter {
	rv := objc.Call[ContentFilter](c_, objc.Sel("initWithDisplay:includingWindows:"), objc.Ptr(display), includedWindows)
	return rv
}

// Creates a filter that captures only specific windows from a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/3919808-initwithdisplay?language=objc
func NewContentFilterWithDisplayIncludingWindows(display IDisplay, includedWindows []IWindow) ContentFilter {
	instance := ContentFilterClass.Alloc().InitWithDisplayIncludingWindows(display, includedWindows)
	instance.Autorelease()
	return instance
}

func (c_ ContentFilter) InitWithDesktopIndependentWindow(window IWindow) ContentFilter {
	rv := objc.Call[ContentFilter](c_, objc.Sel("initWithDesktopIndependentWindow:"), objc.Ptr(window))
	return rv
}

// Creates a filter that captures only the specified window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentfilter/3919804-initwithdesktopindependentwindow?language=objc
func NewContentFilterWithDesktopIndependentWindow(window IWindow) ContentFilter {
	instance := ContentFilterClass.Alloc().InitWithDesktopIndependentWindow(window)
	instance.Autorelease()
	return instance
}

func (cc _ContentFilterClass) Alloc() ContentFilter {
	rv := objc.Call[ContentFilter](cc, objc.Sel("alloc"))
	return rv
}

func (cc _ContentFilterClass) New() ContentFilter {
	rv := objc.Call[ContentFilter](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContentFilter() ContentFilter {
	return ContentFilterClass.New()
}

func (c_ ContentFilter) Init() ContentFilter {
	rv := objc.Call[ContentFilter](c_, objc.Sel("init"))
	return rv
}
