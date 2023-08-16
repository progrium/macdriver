// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScrollView] class.
var ScrollViewClass = _ScrollViewClass{objc.GetClass("NSScrollView")}

type _ScrollViewClass struct {
	objc.Class
}

// An interface definition for the [ScrollView] class.
type IScrollView interface {
	IView
	SetMagnificationCenteredAtPoint(magnification float64, point foundation.Point)
	FlashScrollers()
	AddFloatingSubviewForAxis(view IView, axis EventGestureAxis)
	MagnifyToFitRect(rect foundation.Rect)
	Tile()
	VerticalLineScroll() float64
	SetVerticalLineScroll(value float64)
	HorizontalScrollElasticity() ScrollElasticity
	SetHorizontalScrollElasticity(value ScrollElasticity)
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	HorizontalPageScroll() float64
	SetHorizontalPageScroll(value float64)
	RulersVisible() bool
	SetRulersVisible(value bool)
	DocumentCursor() Cursor
	SetDocumentCursor(value ICursor)
	HasVerticalRuler() bool
	SetHasVerticalRuler(value bool)
	MaxMagnification() float64
	SetMaxMagnification(value float64)
	ScrollsDynamically() bool
	SetScrollsDynamically(value bool)
	ContentView() ClipView
	SetContentView(value IClipView)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	FindBarPosition() ScrollViewFindBarPosition
	SetFindBarPosition(value ScrollViewFindBarPosition)
	HorizontalScroller() Scroller
	SetHorizontalScroller(value IScroller)
	ScrollerStyle() ScrollerStyle
	SetScrollerStyle(value ScrollerStyle)
	VerticalScrollElasticity() ScrollElasticity
	SetVerticalScrollElasticity(value ScrollElasticity)
	PageScroll() float64
	SetPageScroll(value float64)
	BorderType() BorderType
	SetBorderType(value BorderType)
	ContentInsets() foundation.EdgeInsets
	SetContentInsets(value foundation.EdgeInsets)
	Magnification() float64
	SetMagnification(value float64)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	ScrollerKnobStyle() ScrollerKnobStyle
	SetScrollerKnobStyle(value ScrollerKnobStyle)
	LineScroll() float64
	SetLineScroll(value float64)
	ContentSize() foundation.Size
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	AutohidesScrollers() bool
	SetAutohidesScrollers(value bool)
	UsesPredominantAxisScrolling() bool
	SetUsesPredominantAxisScrolling(value bool)
	DocumentVisibleRect() foundation.Rect
	ScrollerInsets() foundation.EdgeInsets
	SetScrollerInsets(value foundation.EdgeInsets)
	MinMagnification() float64
	SetMinMagnification(value float64)
	HasHorizontalRuler() bool
	SetHasHorizontalRuler(value bool)
	DocumentView() View
	SetDocumentView(value IView)
	HorizontalLineScroll() float64
	SetHorizontalLineScroll(value float64)
	VerticalScroller() Scroller
	SetVerticalScroller(value IScroller)
	VerticalPageScroll() float64
	SetVerticalPageScroll(value float64)
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
	VerticalRulerView() RulerView
	SetVerticalRulerView(value IRulerView)
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	HorizontalRulerView() RulerView
	SetHorizontalRulerView(value IRulerView)
}

// A view that displays a portion of a document view and provides scroll bars that allow the user to move the document view within the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview?language=objc
type ScrollView struct {
	View
}

func ScrollViewFrom(ptr unsafe.Pointer) ScrollView {
	return ScrollView{
		View: ViewFrom(ptr),
	}
}

func (s_ ScrollView) InitWithFrame(frameRect foundation.Rect) ScrollView {
	rv := objc.Call[ScrollView](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403450-initwithframe?language=objc
func ScrollView_InitWithFrame(frameRect foundation.Rect) ScrollView {
	return ScrollViewClass.Alloc().InitWithFrame(frameRect)
}

func (sc _ScrollViewClass) Alloc() ScrollView {
	rv := objc.Call[ScrollView](sc, objc.Sel("alloc"))
	return rv
}

func ScrollView_Alloc() ScrollView {
	return ScrollViewClass.Alloc()
}

func (sc _ScrollViewClass) New() ScrollView {
	rv := objc.Call[ScrollView](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScrollView() ScrollView {
	return ScrollViewClass.New()
}

func (s_ ScrollView) Init() ScrollView {
	rv := objc.Call[ScrollView](s_, objc.Sel("init"))
	return rv
}

// Magnify the content by the given amount and center the result on the given point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403459-setmagnification?language=objc
func (s_ ScrollView) SetMagnificationCenteredAtPoint(magnification float64, point foundation.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}

// Flash the overlay scroll bars. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403460-flashscrollers?language=objc
func (s_ ScrollView) FlashScrollers() {
	objc.Call[objc.Void](s_, objc.Sel("flashScrollers"))
}

// Adds a floating subview to the document view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403546-addfloatingsubview?language=objc
func (s_ ScrollView) AddFloatingSubviewForAxis(view IView, axis EventGestureAxis) {
	objc.Call[objc.Void](s_, objc.Sel("addFloatingSubview:forAxis:"), objc.Ptr(view), axis)
}

// Magnifies the content view proportionally such that the given rectangle fits centered in the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403508-magnifytofitrect?language=objc
func (s_ ScrollView) MagnifyToFitRect(rect foundation.Rect) {
	objc.Call[objc.Void](s_, objc.Sel("magnifyToFitRect:"), rect)
}

// Lays out the components of the receiver: the content view, the scrollers, and the ruler views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403509-tile?language=objc
func (s_ ScrollView) Tile() {
	objc.Call[objc.Void](s_, objc.Sel("tile"))
}

// Returns the frame size of a scroll view that contains a content view with the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403545-framesizeforcontentsize?language=objc
func (sc _ScrollViewClass) FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize foundation.Size, horizontalScrollerClass objc.IClass, verticalScrollerClass objc.IClass, type_ BorderType, controlSize ControlSize, scrollerStyle ScrollerStyle) foundation.Size {
	rv := objc.Call[foundation.Size](sc, objc.Sel("frameSizeForContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), cSize, objc.Ptr(horizontalScrollerClass), objc.Ptr(verticalScrollerClass), type_, controlSize, scrollerStyle)
	return rv
}

// Returns the frame size of a scroll view that contains a content view with the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403545-framesizeforcontentsize?language=objc
func ScrollView_FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize foundation.Size, horizontalScrollerClass objc.IClass, verticalScrollerClass objc.IClass, type_ BorderType, controlSize ControlSize, scrollerStyle ScrollerStyle) foundation.Size {
	return ScrollViewClass.FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
}

// The scroll view’s vertical line by line scroll amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403490-verticallinescroll?language=objc
func (s_ ScrollView) VerticalLineScroll() float64 {
	rv := objc.Call[float64](s_, objc.Sel("verticalLineScroll"))
	return rv
}

// The scroll view’s vertical line by line scroll amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403490-verticallinescroll?language=objc
func (s_ ScrollView) SetVerticalLineScroll(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setVerticalLineScroll:"), value)
}

// The scroll view’s horizontal scrolling elasticity mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403540-horizontalscrollelasticity?language=objc
func (s_ ScrollView) HorizontalScrollElasticity() ScrollElasticity {
	rv := objc.Call[ScrollElasticity](s_, objc.Sel("horizontalScrollElasticity"))
	return rv
}

// The scroll view’s horizontal scrolling elasticity mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403540-horizontalscrollelasticity?language=objc
func (s_ ScrollView) SetHorizontalScrollElasticity(value ScrollElasticity) {
	objc.Call[objc.Void](s_, objc.Sel("setHorizontalScrollElasticity:"), value)
}

// Allows the user to magnify the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403531-allowsmagnification?language=objc
func (s_ ScrollView) AllowsMagnification() bool {
	rv := objc.Call[bool](s_, objc.Sel("allowsMagnification"))
	return rv
}

// Allows the user to magnify the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403531-allowsmagnification?language=objc
func (s_ ScrollView) SetAllowsMagnification(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAllowsMagnification:"), value)
}

// The amount of the document view kept visible when scrolling horizontally page by page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403478-horizontalpagescroll?language=objc
func (s_ ScrollView) HorizontalPageScroll() float64 {
	rv := objc.Call[float64](s_, objc.Sel("horizontalPageScroll"))
	return rv
}

// The amount of the document view kept visible when scrolling horizontally page by page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403478-horizontalpagescroll?language=objc
func (s_ ScrollView) SetHorizontalPageScroll(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setHorizontalPageScroll:"), value)
}

// Returns the default class to be used for ruler objects in NSScrollViews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403492-rulerviewclass?language=objc
func (sc _ScrollViewClass) RulerViewClass() objc.Class {
	rv := objc.Call[objc.Class](sc, objc.Sel("rulerViewClass"))
	return rv
}

// Returns the default class to be used for ruler objects in NSScrollViews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403492-rulerviewclass?language=objc
func ScrollView_RulerViewClass() objc.Class {
	return ScrollViewClass.RulerViewClass()
}

// Returns the default class to be used for ruler objects in NSScrollViews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403492-rulerviewclass?language=objc
func (sc _ScrollViewClass) SetRulerViewClass(value objc.IClass) {
	objc.Call[objc.Void](sc, objc.Sel("setRulerViewClass:"), objc.Ptr(value))
}

// Returns the default class to be used for ruler objects in NSScrollViews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403492-rulerviewclass?language=objc
func ScrollView_SetRulerViewClass(value objc.IClass) {
	ScrollViewClass.SetRulerViewClass(value)
}

// A Boolean that indicates whether the scroll view displays its rulers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403445-rulersvisible?language=objc
func (s_ ScrollView) RulersVisible() bool {
	rv := objc.Call[bool](s_, objc.Sel("rulersVisible"))
	return rv
}

// A Boolean that indicates whether the scroll view displays its rulers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403445-rulersvisible?language=objc
func (s_ ScrollView) SetRulersVisible(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setRulersVisible:"), value)
}

// The content view’s document cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403446-documentcursor?language=objc
func (s_ ScrollView) DocumentCursor() Cursor {
	rv := objc.Call[Cursor](s_, objc.Sel("documentCursor"))
	return rv
}

// The content view’s document cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403446-documentcursor?language=objc
func (s_ ScrollView) SetDocumentCursor(value ICursor) {
	objc.Call[objc.Void](s_, objc.Sel("setDocumentCursor:"), objc.Ptr(value))
}

// A Boolean that indicates whether the scroll view keeps a vertical ruler object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403496-hasverticalruler?language=objc
func (s_ ScrollView) HasVerticalRuler() bool {
	rv := objc.Call[bool](s_, objc.Sel("hasVerticalRuler"))
	return rv
}

// A Boolean that indicates whether the scroll view keeps a vertical ruler object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403496-hasverticalruler?language=objc
func (s_ ScrollView) SetHasVerticalRuler(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setHasVerticalRuler:"), value)
}

// The maximum value to which the content can be magnified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403510-maxmagnification?language=objc
func (s_ ScrollView) MaxMagnification() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maxMagnification"))
	return rv
}

// The maximum value to which the content can be magnified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403510-maxmagnification?language=objc
func (s_ ScrollView) SetMaxMagnification(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMaxMagnification:"), value)
}

// A Boolean that indicates whether the scroll view redraws its document view while scrolling continuously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403519-scrollsdynamically?language=objc
func (s_ ScrollView) ScrollsDynamically() bool {
	rv := objc.Call[bool](s_, objc.Sel("scrollsDynamically"))
	return rv
}

// A Boolean that indicates whether the scroll view redraws its document view while scrolling continuously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403519-scrollsdynamically?language=objc
func (s_ ScrollView) SetScrollsDynamically(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setScrollsDynamically:"), value)
}

// The scroll view’s content view, the view that clips the document view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403547-contentview?language=objc
func (s_ ScrollView) ContentView() ClipView {
	rv := objc.Call[ClipView](s_, objc.Sel("contentView"))
	return rv
}

// The scroll view’s content view, the view that clips the document view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403547-contentview?language=objc
func (s_ ScrollView) SetContentView(value IClipView) {
	objc.Call[objc.Void](s_, objc.Sel("setContentView:"), objc.Ptr(value))
}

// A Boolean that indicates whether the scroll view has a vertical scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403491-hasverticalscroller?language=objc
func (s_ ScrollView) HasVerticalScroller() bool {
	rv := objc.Call[bool](s_, objc.Sel("hasVerticalScroller"))
	return rv
}

// A Boolean that indicates whether the scroll view has a vertical scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403491-hasverticalscroller?language=objc
func (s_ ScrollView) SetHasVerticalScroller(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setHasVerticalScroller:"), value)
}

// The position of the find bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403501-findbarposition?language=objc
func (s_ ScrollView) FindBarPosition() ScrollViewFindBarPosition {
	rv := objc.Call[ScrollViewFindBarPosition](s_, objc.Sel("findBarPosition"))
	return rv
}

// The position of the find bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403501-findbarposition?language=objc
func (s_ ScrollView) SetFindBarPosition(value ScrollViewFindBarPosition) {
	objc.Call[objc.Void](s_, objc.Sel("setFindBarPosition:"), value)
}

// The scroll view’s horizontal scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403447-horizontalscroller?language=objc
func (s_ ScrollView) HorizontalScroller() Scroller {
	rv := objc.Call[Scroller](s_, objc.Sel("horizontalScroller"))
	return rv
}

// The scroll view’s horizontal scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403447-horizontalscroller?language=objc
func (s_ ScrollView) SetHorizontalScroller(value IScroller) {
	objc.Call[objc.Void](s_, objc.Sel("setHorizontalScroller:"), objc.Ptr(value))
}

// The scroller style used by the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403520-scrollerstyle?language=objc
func (s_ ScrollView) ScrollerStyle() ScrollerStyle {
	rv := objc.Call[ScrollerStyle](s_, objc.Sel("scrollerStyle"))
	return rv
}

// The scroller style used by the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403520-scrollerstyle?language=objc
func (s_ ScrollView) SetScrollerStyle(value ScrollerStyle) {
	objc.Call[objc.Void](s_, objc.Sel("setScrollerStyle:"), value)
}

// The scroll view’s vertical scrolling elasticity mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403475-verticalscrollelasticity?language=objc
func (s_ ScrollView) VerticalScrollElasticity() ScrollElasticity {
	rv := objc.Call[ScrollElasticity](s_, objc.Sel("verticalScrollElasticity"))
	return rv
}

// The scroll view’s vertical scrolling elasticity mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403475-verticalscrollelasticity?language=objc
func (s_ ScrollView) SetVerticalScrollElasticity(value ScrollElasticity) {
	objc.Call[objc.Void](s_, objc.Sel("setVerticalScrollElasticity:"), value)
}

// The amount of the document view kept visible when scrolling page by page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403449-pagescroll?language=objc
func (s_ ScrollView) PageScroll() float64 {
	rv := objc.Call[float64](s_, objc.Sel("pageScroll"))
	return rv
}

// The amount of the document view kept visible when scrolling page by page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403449-pagescroll?language=objc
func (s_ ScrollView) SetPageScroll(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setPageScroll:"), value)
}

// A value that specifies the appearance of the scroll view’s border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403528-bordertype?language=objc
func (s_ ScrollView) BorderType() BorderType {
	rv := objc.Call[BorderType](s_, objc.Sel("borderType"))
	return rv
}

// A value that specifies the appearance of the scroll view’s border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403528-bordertype?language=objc
func (s_ ScrollView) SetBorderType(value BorderType) {
	objc.Call[objc.Void](s_, objc.Sel("setBorderType:"), value)
}

// The distance that the scroll view’s subviews are inset from the enclosing scroll view during tiling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403461-contentinsets?language=objc
func (s_ ScrollView) ContentInsets() foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](s_, objc.Sel("contentInsets"))
	return rv
}

// The distance that the scroll view’s subviews are inset from the enclosing scroll view during tiling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403461-contentinsets?language=objc
func (s_ ScrollView) SetContentInsets(value foundation.EdgeInsets) {
	objc.Call[objc.Void](s_, objc.Sel("setContentInsets:"), value)
}

// The amount by which the content is currently scaled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403497-magnification?language=objc
func (s_ ScrollView) Magnification() float64 {
	rv := objc.Call[float64](s_, objc.Sel("magnification"))
	return rv
}

// The amount by which the content is currently scaled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403497-magnification?language=objc
func (s_ ScrollView) SetMagnification(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMagnification:"), value)
}

// The color of the content view’s background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403473-backgroundcolor?language=objc
func (s_ ScrollView) BackgroundColor() Color {
	rv := objc.Call[Color](s_, objc.Sel("backgroundColor"))
	return rv
}

// The color of the content view’s background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403473-backgroundcolor?language=objc
func (s_ ScrollView) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// The knob style of scroll views that use the overlay scroller style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403544-scrollerknobstyle?language=objc
func (s_ ScrollView) ScrollerKnobStyle() ScrollerKnobStyle {
	rv := objc.Call[ScrollerKnobStyle](s_, objc.Sel("scrollerKnobStyle"))
	return rv
}

// The knob style of scroll views that use the overlay scroller style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403544-scrollerknobstyle?language=objc
func (s_ ScrollView) SetScrollerKnobStyle(value ScrollerKnobStyle) {
	objc.Call[objc.Void](s_, objc.Sel("setScrollerKnobStyle:"), value)
}

// The scroll view’s line by line scroll amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403454-linescroll?language=objc
func (s_ ScrollView) LineScroll() float64 {
	rv := objc.Call[float64](s_, objc.Sel("lineScroll"))
	return rv
}

// The scroll view’s line by line scroll amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403454-linescroll?language=objc
func (s_ ScrollView) SetLineScroll(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setLineScroll:"), value)
}

// The size of the scroll view’s content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403458-contentsize?language=objc
func (s_ ScrollView) ContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](s_, objc.Sel("contentSize"))
	return rv
}

// A Boolean that indicates whether the scroll view draws its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403474-drawsbackground?language=objc
func (s_ ScrollView) DrawsBackground() bool {
	rv := objc.Call[bool](s_, objc.Sel("drawsBackground"))
	return rv
}

// A Boolean that indicates whether the scroll view draws its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403474-drawsbackground?language=objc
func (s_ ScrollView) SetDrawsBackground(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setDrawsBackground:"), value)
}

// A Boolean that indicates whether the scroll view automatically hides its scroll bars when they are not needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403536-autohidesscrollers?language=objc
func (s_ ScrollView) AutohidesScrollers() bool {
	rv := objc.Call[bool](s_, objc.Sel("autohidesScrollers"))
	return rv
}

// A Boolean that indicates whether the scroll view automatically hides its scroll bars when they are not needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403536-autohidesscrollers?language=objc
func (s_ ScrollView) SetAutohidesScrollers(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAutohidesScrollers:"), value)
}

// A Boolean that indicates whether the scroll view uses a predominant scrolling axis for content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403489-usespredominantaxisscrolling?language=objc
func (s_ ScrollView) UsesPredominantAxisScrolling() bool {
	rv := objc.Call[bool](s_, objc.Sel("usesPredominantAxisScrolling"))
	return rv
}

// A Boolean that indicates whether the scroll view uses a predominant scrolling axis for content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403489-usespredominantaxisscrolling?language=objc
func (s_ ScrollView) SetUsesPredominantAxisScrolling(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setUsesPredominantAxisScrolling:"), value)
}

// The portion of the document view, in its own coordinate system, visible through the scroll view’s content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403466-documentvisiblerect?language=objc
func (s_ ScrollView) DocumentVisibleRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](s_, objc.Sel("documentVisibleRect"))
	return rv
}

// The distance the scrollers are inset from the edge of the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403529-scrollerinsets?language=objc
func (s_ ScrollView) ScrollerInsets() foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](s_, objc.Sel("scrollerInsets"))
	return rv
}

// The distance the scrollers are inset from the edge of the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403529-scrollerinsets?language=objc
func (s_ ScrollView) SetScrollerInsets(value foundation.EdgeInsets) {
	objc.Call[objc.Void](s_, objc.Sel("setScrollerInsets:"), value)
}

// The minimum value to which the content can be magnified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403524-minmagnification?language=objc
func (s_ ScrollView) MinMagnification() float64 {
	rv := objc.Call[float64](s_, objc.Sel("minMagnification"))
	return rv
}

// The minimum value to which the content can be magnified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403524-minmagnification?language=objc
func (s_ ScrollView) SetMinMagnification(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMinMagnification:"), value)
}

// A Boolean that indicates whether the scroll view keeps a horizontal ruler object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403457-hashorizontalruler?language=objc
func (s_ ScrollView) HasHorizontalRuler() bool {
	rv := objc.Call[bool](s_, objc.Sel("hasHorizontalRuler"))
	return rv
}

// A Boolean that indicates whether the scroll view keeps a horizontal ruler object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403457-hashorizontalruler?language=objc
func (s_ ScrollView) SetHasHorizontalRuler(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setHasHorizontalRuler:"), value)
}

// The view the scroll view scrolls within its content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403485-documentview?language=objc
func (s_ ScrollView) DocumentView() View {
	rv := objc.Call[View](s_, objc.Sel("documentView"))
	return rv
}

// The view the scroll view scrolls within its content view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403485-documentview?language=objc
func (s_ ScrollView) SetDocumentView(value IView) {
	objc.Call[objc.Void](s_, objc.Sel("setDocumentView:"), objc.Ptr(value))
}

// The scroll view’s horizontal line by line scroll amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403539-horizontallinescroll?language=objc
func (s_ ScrollView) HorizontalLineScroll() float64 {
	rv := objc.Call[float64](s_, objc.Sel("horizontalLineScroll"))
	return rv
}

// The scroll view’s horizontal line by line scroll amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403539-horizontallinescroll?language=objc
func (s_ ScrollView) SetHorizontalLineScroll(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setHorizontalLineScroll:"), value)
}

// The scroll view’s vertical scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403526-verticalscroller?language=objc
func (s_ ScrollView) VerticalScroller() Scroller {
	rv := objc.Call[Scroller](s_, objc.Sel("verticalScroller"))
	return rv
}

// The scroll view’s vertical scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403526-verticalscroller?language=objc
func (s_ ScrollView) SetVerticalScroller(value IScroller) {
	objc.Call[objc.Void](s_, objc.Sel("setVerticalScroller:"), objc.Ptr(value))
}

// The amount of the document view kept visible when scrolling vertically page by page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403512-verticalpagescroll?language=objc
func (s_ ScrollView) VerticalPageScroll() float64 {
	rv := objc.Call[float64](s_, objc.Sel("verticalPageScroll"))
	return rv
}

// The amount of the document view kept visible when scrolling vertically page by page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403512-verticalpagescroll?language=objc
func (s_ ScrollView) SetVerticalPageScroll(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setVerticalPageScroll:"), value)
}

// A Boolean that indicates whether the scroll view automatically adjusts its content insets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403502-automaticallyadjustscontentinset?language=objc
func (s_ ScrollView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.Call[bool](s_, objc.Sel("automaticallyAdjustsContentInsets"))
	return rv
}

// A Boolean that indicates whether the scroll view automatically adjusts its content insets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403502-automaticallyadjustscontentinset?language=objc
func (s_ ScrollView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAutomaticallyAdjustsContentInsets:"), value)
}

// The scroll view’s vertical ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403507-verticalrulerview?language=objc
func (s_ ScrollView) VerticalRulerView() RulerView {
	rv := objc.Call[RulerView](s_, objc.Sel("verticalRulerView"))
	return rv
}

// The scroll view’s vertical ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403507-verticalrulerview?language=objc
func (s_ ScrollView) SetVerticalRulerView(value IRulerView) {
	objc.Call[objc.Void](s_, objc.Sel("setVerticalRulerView:"), objc.Ptr(value))
}

// A Boolean that indicates whether the scroll view has a horizontal scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403530-hashorizontalscroller?language=objc
func (s_ ScrollView) HasHorizontalScroller() bool {
	rv := objc.Call[bool](s_, objc.Sel("hasHorizontalScroller"))
	return rv
}

// A Boolean that indicates whether the scroll view has a horizontal scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403530-hashorizontalscroller?language=objc
func (s_ ScrollView) SetHasHorizontalScroller(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setHasHorizontalScroller:"), value)
}

// The scroll view’s horizontal ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403498-horizontalrulerview?language=objc
func (s_ ScrollView) HorizontalRulerView() RulerView {
	rv := objc.Call[RulerView](s_, objc.Sel("horizontalRulerView"))
	return rv
}

// The scroll view’s horizontal ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/1403498-horizontalrulerview?language=objc
func (s_ ScrollView) SetHorizontalRulerView(value IRulerView) {
	objc.Call[objc.Void](s_, objc.Sel("setHorizontalRulerView:"), objc.Ptr(value))
}
