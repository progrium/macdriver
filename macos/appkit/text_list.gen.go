// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextList] class.
var TextListClass = _TextListClass{objc.GetClass("NSTextList")}

type _TextListClass struct {
	objc.Class
}

// An interface definition for the [TextList] class.
type ITextList interface {
	objc.IObject
	MarkerForItemNumber(itemNumber int) string
	ListOptions() TextListOptions
	StartingItemNumber() int
	SetStartingItemNumber(value int)
	MarkerFormat() TextListMarkerFormat
}

// A section of text that forms a single list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist?language=objc
type TextList struct {
	objc.Object
}

func TextListFrom(ptr unsafe.Pointer) TextList {
	return TextList{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextList) InitWithMarkerFormatOptions(markerFormat TextListMarkerFormat, options uint) TextList {
	rv := objc.Call[TextList](t_, objc.Sel("initWithMarkerFormat:options:"), markerFormat, options)
	return rv
}

// Returns an initialized text list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/1526123-initwithmarkerformat?language=objc
func NewTextListWithMarkerFormatOptions(markerFormat TextListMarkerFormat, options uint) TextList {
	instance := TextListClass.Alloc().InitWithMarkerFormatOptions(markerFormat, options)
	instance.Autorelease()
	return instance
}

func (tc _TextListClass) Alloc() TextList {
	rv := objc.Call[TextList](tc, objc.Sel("alloc"))
	return rv
}

func TextList_Alloc() TextList {
	return TextListClass.Alloc()
}

func (tc _TextListClass) New() TextList {
	rv := objc.Call[TextList](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextList() TextList {
	return TextListClass.New()
}

func (t_ TextList) Init() TextList {
	rv := objc.Call[TextList](t_, objc.Sel("init"))
	return rv
}

// Returns the computed value for a specific ordinal position in the list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/1532739-markerforitemnumber?language=objc
func (t_ TextList) MarkerForItemNumber(itemNumber int) string {
	rv := objc.Call[string](t_, objc.Sel("markerForItemNumber:"), itemNumber)
	return rv
}

// Returns the list options mask value of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/1533519-listoptions?language=objc
func (t_ TextList) ListOptions() TextListOptions {
	rv := objc.Call[TextListOptions](t_, objc.Sel("listOptions"))
	return rv
}

// Sets the starting item number for the text list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/1528597-startingitemnumber?language=objc
func (t_ TextList) StartingItemNumber() int {
	rv := objc.Call[int](t_, objc.Sel("startingItemNumber"))
	return rv
}

// Sets the starting item number for the text list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/1528597-startingitemnumber?language=objc
func (t_ TextList) SetStartingItemNumber(value int) {
	objc.Call[objc.Void](t_, objc.Sel("setStartingItemNumber:"), value)
}

// Returns the marker format string used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlist/1533865-markerformat?language=objc
func (t_ TextList) MarkerFormat() TextListMarkerFormat {
	rv := objc.Call[TextListMarkerFormat](t_, objc.Sel("markerFormat"))
	return rv
}
