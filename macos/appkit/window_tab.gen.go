// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WindowTab] class.
var WindowTabClass = _WindowTabClass{objc.GetClass("NSWindowTab")}

type _WindowTabClass struct {
	objc.Class
}

// An interface definition for the [WindowTab] class.
type IWindowTab interface {
	objc.IObject
	ToolTip() string
	SetToolTip(value string)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	AccessoryView() View
	SetAccessoryView(value IView)
	Title() string
	SetTitle(value string)
}

// A tab associated with a window that is part of a tabbing group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab?language=objc
type WindowTab struct {
	objc.Object
}

func WindowTabFrom(ptr unsafe.Pointer) WindowTab {
	return WindowTab{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WindowTabClass) Alloc() WindowTab {
	rv := objc.Call[WindowTab](wc, objc.Sel("alloc"))
	return rv
}

func WindowTab_Alloc() WindowTab {
	return WindowTabClass.Alloc()
}

func (wc _WindowTabClass) New() WindowTab {
	rv := objc.Call[WindowTab](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWindowTab() WindowTab {
	return WindowTabClass.New()
}

func (w_ WindowTab) Init() WindowTab {
	rv := objc.Call[WindowTab](w_, objc.Sel("init"))
	return rv
}

// The tooltip for this window tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/2869744-tooltip?language=objc
func (w_ WindowTab) ToolTip() string {
	rv := objc.Call[string](w_, objc.Sel("toolTip"))
	return rv
}

// The tooltip for this window tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/2869744-tooltip?language=objc
func (w_ WindowTab) SetToolTip(value string) {
	objc.Call[objc.Void](w_, objc.Sel("setToolTip:"), value)
}

// The title for the window tab, specified as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/2869743-attributedtitle?language=objc
func (w_ WindowTab) AttributedTitle() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](w_, objc.Sel("attributedTitle"))
	return rv
}

// The title for the window tab, specified as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/2869743-attributedtitle?language=objc
func (w_ WindowTab) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Call[objc.Void](w_, objc.Sel("setAttributedTitle:"), objc.Ptr(value))
}

// An optional accessory view for the tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/2869745-accessoryview?language=objc
func (w_ WindowTab) AccessoryView() View {
	rv := objc.Call[View](w_, objc.Sel("accessoryView"))
	return rv
}

// An optional accessory view for the tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/2869745-accessoryview?language=objc
func (w_ WindowTab) SetAccessoryView(value IView) {
	objc.Call[objc.Void](w_, objc.Sel("setAccessoryView:"), objc.Ptr(value))
}

// The title for the window tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/2869746-title?language=objc
func (w_ WindowTab) Title() string {
	rv := objc.Call[string](w_, objc.Sel("title"))
	return rv
}

// The title for the window tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtab/2869746-title?language=objc
func (w_ WindowTab) SetTitle(value string) {
	objc.Call[objc.Void](w_, objc.Sel("setTitle:"), value)
}
