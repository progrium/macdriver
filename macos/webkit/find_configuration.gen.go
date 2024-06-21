// Code generated by DarwinKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [FindConfiguration] class.
var FindConfigurationClass = _FindConfigurationClass{objc.GetClass("WKFindConfiguration")}

type _FindConfigurationClass struct {
	objc.Class
}

// An interface definition for the [FindConfiguration] class.
type IFindConfiguration interface {
	objc.IObject
	Backwards() bool
	SetBackwards(value bool)
	Wraps() bool
	SetWraps(value bool)
	CaseSensitive() bool
	SetCaseSensitive(value bool)
}

// The configuration parameters to use when searching the contents of the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindconfiguration?language=objc
type FindConfiguration struct {
	objc.Object
}

func FindConfigurationFrom(ptr unsafe.Pointer) FindConfiguration {
	return FindConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FindConfigurationClass) Alloc() FindConfiguration {
	rv := objc.Call[FindConfiguration](fc, objc.Sel("alloc"))
	return rv
}

func (fc _FindConfigurationClass) New() FindConfiguration {
	rv := objc.Call[FindConfiguration](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFindConfiguration() FindConfiguration {
	return FindConfigurationClass.New()
}

func (f_ FindConfiguration) Init() FindConfiguration {
	rv := objc.Call[FindConfiguration](f_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates the search direction, relative to the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindconfiguration/3516854-backwards?language=objc
func (f_ FindConfiguration) Backwards() bool {
	rv := objc.Call[bool](f_, objc.Sel("backwards"))
	return rv
}

// A Boolean value that indicates the search direction, relative to the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindconfiguration/3516854-backwards?language=objc
func (f_ FindConfiguration) SetBackwards(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setBackwards:"), value)
}

// A Boolean value that indicates whether the search wraps around to the other side of the page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindconfiguration/3516856-wraps?language=objc
func (f_ FindConfiguration) Wraps() bool {
	rv := objc.Call[bool](f_, objc.Sel("wraps"))
	return rv
}

// A Boolean value that indicates whether the search wraps around to the other side of the page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindconfiguration/3516856-wraps?language=objc
func (f_ FindConfiguration) SetWraps(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setWraps:"), value)
}

// A Boolean value that indicates whether to consider case when matching the search string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindconfiguration/3516855-casesensitive?language=objc
func (f_ FindConfiguration) CaseSensitive() bool {
	rv := objc.Call[bool](f_, objc.Sel("caseSensitive"))
	return rv
}

// A Boolean value that indicates whether to consider case when matching the search string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindconfiguration/3516855-casesensitive?language=objc
func (f_ FindConfiguration) SetCaseSensitive(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setCaseSensitive:"), value)
}
