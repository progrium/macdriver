// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextFinder] class.
var TextFinderClass = _TextFinderClass{objc.GetClass("NSTextFinder")}

type _TextFinderClass struct {
	objc.Class
}

// An interface definition for the [TextFinder] class.
type ITextFinder interface {
	objc.IObject
	PerformAction(op TextFinderAction)
	CancelFindIndicator()
	NoteClientStringWillChange()
	ValidateAction(op TextFinderAction) bool
	IncrementalMatchRanges() []foundation.Value
	IsIncrementalSearchingEnabled() bool
	SetIncrementalSearchingEnabled(value bool)
	FindIndicatorNeedsUpdate() bool
	SetFindIndicatorNeedsUpdate(value bool)
	FindBarContainer() TextFinderBarContainerWrapper
	SetFindBarContainer(value PTextFinderBarContainer)
	SetFindBarContainerObject(valueObject objc.IObject)
	IncrementalSearchingShouldDimContentView() bool
	SetIncrementalSearchingShouldDimContentView(value bool)
	Client() TextFinderClientWrapper
	SetClient(value PTextFinderClient)
	SetClientObject(valueObject objc.IObject)
}

// An optional search-and-replace find interface inside a view, usually a scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder?language=objc
type TextFinder struct {
	objc.Object
}

func TextFinderFrom(ptr unsafe.Pointer) TextFinder {
	return TextFinder{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextFinder) Init() TextFinder {
	rv := objc.Call[TextFinder](t_, objc.Sel("init"))
	return rv
}

func (tc _TextFinderClass) Alloc() TextFinder {
	rv := objc.Call[TextFinder](tc, objc.Sel("alloc"))
	return rv
}

func TextFinder_Alloc() TextFinder {
	return TextFinderClass.Alloc()
}

func (tc _TextFinderClass) New() TextFinder {
	rv := objc.Call[TextFinder](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextFinder() TextFinder {
	return TextFinderClass.New()
}

// Performs the specified text finding action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1526414-performaction?language=objc
func (t_ TextFinder) PerformAction(op TextFinderAction) {
	objc.Call[objc.Void](t_, objc.Sel("performAction:"), op)
}

// Cancels the find indicator immediately. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1525467-cancelfindindicator?language=objc
func (t_ TextFinder) CancelFindIndicator() {
	objc.Call[objc.Void](t_, objc.Sel("cancelFindIndicator"))
}

// Invoke this method when the searched content will change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1534184-noteclientstringwillchange?language=objc
func (t_ TextFinder) NoteClientStringWillChange() {
	objc.Call[objc.Void](t_, objc.Sel("noteClientStringWillChange"))
}

// Allows validation of the find action before performing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1527203-validateaction?language=objc
func (t_ TextFinder) ValidateAction(op TextFinderAction) bool {
	rv := objc.Call[bool](t_, objc.Sel("validateAction:"), op)
	return rv
}

// Override this method to draw custom highlighting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1526120-drawincrementalmatchhighlightinr?language=objc
func (tc _TextFinderClass) DrawIncrementalMatchHighlightInRect(rect foundation.Rect) {
	objc.Call[objc.Void](tc, objc.Sel("drawIncrementalMatchHighlightInRect:"), rect)
}

// Override this method to draw custom highlighting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1526120-drawincrementalmatchhighlightinr?language=objc
func TextFinder_DrawIncrementalMatchHighlightInRect(rect foundation.Rect) {
	TextFinderClass.DrawIncrementalMatchHighlightInRect(rect)
}

// Array of incremental search matches posted on the main queue, which have been found during a background search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1528304-incrementalmatchranges?language=objc
func (t_ TextFinder) IncrementalMatchRanges() []foundation.Value {
	rv := objc.Call[[]foundation.Value](t_, objc.Sel("incrementalMatchRanges"))
	return rv
}

// Determines if incremental searching is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1535849-incrementalsearchingenabled?language=objc
func (t_ TextFinder) IsIncrementalSearchingEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isIncrementalSearchingEnabled"))
	return rv
}

// Determines if incremental searching is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1535849-incrementalsearchingenabled?language=objc
func (t_ TextFinder) SetIncrementalSearchingEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setIncrementalSearchingEnabled:"), value)
}

// Invoke to specify that the find indicator needs updating when not contained within a scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1534431-findindicatorneedsupdate?language=objc
func (t_ TextFinder) FindIndicatorNeedsUpdate() bool {
	rv := objc.Call[bool](t_, objc.Sel("findIndicatorNeedsUpdate"))
	return rv
}

// Invoke to specify that the find indicator needs updating when not contained within a scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1534431-findindicatorneedsupdate?language=objc
func (t_ TextFinder) SetFindIndicatorNeedsUpdate(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setFindIndicatorNeedsUpdate:"), value)
}

// Specifies the find bar container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1526748-findbarcontainer?language=objc
func (t_ TextFinder) FindBarContainer() TextFinderBarContainerWrapper {
	rv := objc.Call[TextFinderBarContainerWrapper](t_, objc.Sel("findBarContainer"))
	return rv
}

// Specifies the find bar container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1526748-findbarcontainer?language=objc
func (t_ TextFinder) SetFindBarContainer(value PTextFinderBarContainer) {
	po0 := objc.WrapAsProtocol("NSTextFinderBarContainer", value)
	objc.Call[objc.Void](t_, objc.Sel("setFindBarContainer:"), po0)
}

// Specifies the find bar container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1526748-findbarcontainer?language=objc
func (t_ TextFinder) SetFindBarContainerObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setFindBarContainer:"), objc.Ptr(valueObject))
}

// Determines the type of incremental search feedback to be presented [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1528196-incrementalsearchingshoulddimcon?language=objc
func (t_ TextFinder) IncrementalSearchingShouldDimContentView() bool {
	rv := objc.Call[bool](t_, objc.Sel("incrementalSearchingShouldDimContentView"))
	return rv
}

// Determines the type of incremental search feedback to be presented [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1528196-incrementalsearchingshoulddimcon?language=objc
func (t_ TextFinder) SetIncrementalSearchingShouldDimContentView(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setIncrementalSearchingShouldDimContentView:"), value)
}

// The object that provides the target search string, find bar location, and feedback methods. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1533813-client?language=objc
func (t_ TextFinder) Client() TextFinderClientWrapper {
	rv := objc.Call[TextFinderClientWrapper](t_, objc.Sel("client"))
	return rv
}

// The object that provides the target search string, find bar location, and feedback methods. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1533813-client?language=objc
func (t_ TextFinder) SetClient(value PTextFinderClient) {
	po0 := objc.WrapAsProtocol("NSTextFinderClient", value)
	objc.Call[objc.Void](t_, objc.Sel("setClient:"), po0)
}

// The object that provides the target search string, find bar location, and feedback methods. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinder/1533813-client?language=objc
func (t_ TextFinder) SetClientObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setClient:"), objc.Ptr(valueObject))
}
