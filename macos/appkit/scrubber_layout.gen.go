// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrubberLayout] class.
var ScrubberLayoutClass = _ScrubberLayoutClass{objc.GetClass("NSScrubberLayout")}

type _ScrubberLayoutClass struct {
	objc.Class
}

// An interface definition for the [ScrubberLayout] class.
type IScrubberLayout interface {
	objc.IObject
	LayoutAttributesForItemsInRect(rect foundation.Rect) foundation.Set
	PrepareLayout()
	InvalidateLayout()
	LayoutAttributesForItemAtIndex(index int) ScrubberLayoutAttributes
	ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect foundation.Rect, toVisibleRect foundation.Rect) bool
	VisibleRect() foundation.Rect
	ShouldInvalidateLayoutForSelectionChange() bool
	ScrubberContentSize() foundation.Size
	Scrubber() Scrubber
	AutomaticallyMirrorsInRightToLeftLayout() bool
	ShouldInvalidateLayoutForHighlightChange() bool
}

// An abstract class that describes the layout of items within a scrubber control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout?language=objc
type ScrubberLayout struct {
	objc.Object
}

func ScrubberLayoutFrom(ptr unsafe.Pointer) ScrubberLayout {
	return ScrubberLayout{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ ScrubberLayout) Init() ScrubberLayout {
	rv := objc.Call[ScrubberLayout](s_, objc.Sel("init"))
	return rv
}

func (sc _ScrubberLayoutClass) Alloc() ScrubberLayout {
	rv := objc.Call[ScrubberLayout](sc, objc.Sel("alloc"))
	return rv
}

func ScrubberLayout_Alloc() ScrubberLayout {
	return ScrubberLayoutClass.Alloc()
}

func (sc _ScrubberLayoutClass) New() ScrubberLayout {
	rv := objc.Call[ScrubberLayout](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrubberLayout() ScrubberLayout {
	return ScrubberLayoutClass.New()
}

// The set of layout attributes for all items within the provided rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544636-layoutattributesforitemsinrect?language=objc
func (s_ ScrubberLayout) LayoutAttributesForItemsInRect(rect foundation.Rect) foundation.Set {
	rv := objc.Call[foundation.Set](s_, objc.Sel("layoutAttributesForItemsInRect:"), rect)
	return rv
}

// Gives you an opportunity to perform layout calculations when the scrubber's layout is invalidated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544632-preparelayout?language=objc
func (s_ ScrubberLayout) PrepareLayout() {
	objc.Call[objc.Void](s_, objc.Sel("prepareLayout"))
}

// Signals that the layout has been invalidated, and that the scrubber control should perform a new layout pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544626-invalidatelayout?language=objc
func (s_ ScrubberLayout) InvalidateLayout() {
	objc.Call[objc.Void](s_, objc.Sel("invalidateLayout"))
}

// The layout attributes for the item with the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544650-layoutattributesforitematindex?language=objc
func (s_ ScrubberLayout) LayoutAttributesForItemAtIndex(index int) ScrubberLayoutAttributes {
	rv := objc.Call[ScrubberLayoutAttributes](s_, objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return rv
}

// Determines whether the scrubber should refresh its layout in response to a change of its visible region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544633-shouldinvalidatelayoutforchangef?language=objc
func (s_ ScrubberLayout) ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect foundation.Rect, toVisibleRect foundation.Rect) bool {
	rv := objc.Call[bool](s_, objc.Sel("shouldInvalidateLayoutForChangeFromVisibleRect:toVisibleRect:"), fromVisibleRect, toVisibleRect)
	return rv
}

// A property containing a class that describes layout attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544628-layoutattributesclass?language=objc
func (sc _ScrubberLayoutClass) LayoutAttributesClass() objc.Class {
	rv := objc.Call[objc.Class](sc, objc.Sel("layoutAttributesClass"))
	return rv
}

// A property containing a class that describes layout attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544628-layoutattributesclass?language=objc
func ScrubberLayout_LayoutAttributesClass() objc.Class {
	return ScrubberLayoutClass.LayoutAttributesClass()
}

// The currently visible rectangle, in the coordinate space of the scrubber content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544637-visiblerect?language=objc
func (s_ ScrubberLayout) VisibleRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("visibleRect"))
	return rv
}

// Determines whether the scrubber should refresh its layout when the selection changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544631-shouldinvalidatelayoutforselecti?language=objc
func (s_ ScrubberLayout) ShouldInvalidateLayoutForSelectionChange() bool {
	rv := objc.Call[bool](s_, objc.Sel("shouldInvalidateLayoutForSelectionChange"))
	return rv
}

// The size required to contain all elements within the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544644-scrubbercontentsize?language=objc
func (s_ ScrubberLayout) ScrubberContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](s_, objc.Sel("scrubberContentSize"))
	return rv
}

// The scrubber control that this layout is assigned to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544638-scrubber?language=objc
func (s_ ScrubberLayout) Scrubber() Scrubber {
	rv := objc.Call[Scrubber](s_, objc.Sel("scrubber"))
	return rv
}

// Determines whether the scrubber mirrors its layout for right-to-left layouts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2646926-automaticallymirrorsinrighttolef?language=objc
func (s_ ScrubberLayout) AutomaticallyMirrorsInRightToLeftLayout() bool {
	rv := objc.Call[bool](s_, objc.Sel("automaticallyMirrorsInRightToLeftLayout"))
	return rv
}

// Determines whether the scrubber should refresh its layout when an item is highlighted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberlayout/2544639-shouldinvalidatelayoutforhighlig?language=objc
func (s_ ScrubberLayout) ShouldInvalidateLayoutForHighlightChange() bool {
	rv := objc.Call[bool](s_, objc.Sel("shouldInvalidateLayoutForHighlightChange"))
	return rv
}
