// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MediaSelection] class.
var MediaSelectionClass = _MediaSelectionClass{objc.GetClass("AVMediaSelection")}

type _MediaSelectionClass struct {
	objc.Class
}

// An interface definition for the [MediaSelection] class.
type IMediaSelection interface {
	objc.IObject
	MediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup(mediaSelectionGroup IMediaSelectionGroup) bool
	SelectedMediaOptionInMediaSelectionGroup(mediaSelectionGroup IMediaSelectionGroup) MediaSelectionOption
	Asset() Asset
}

// An object that represents a complete rendition of media selection options on an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselection?language=objc
type MediaSelection struct {
	objc.Object
}

func MediaSelectionFrom(ptr unsafe.Pointer) MediaSelection {
	return MediaSelection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MediaSelectionClass) Alloc() MediaSelection {
	rv := objc.Call[MediaSelection](mc, objc.Sel("alloc"))
	return rv
}

func MediaSelection_Alloc() MediaSelection {
	return MediaSelectionClass.Alloc()
}

func (mc _MediaSelectionClass) New() MediaSelection {
	rv := objc.Call[MediaSelection](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMediaSelection() MediaSelection {
	return MediaSelectionClass.New()
}

func (m_ MediaSelection) Init() MediaSelection {
	rv := objc.Call[MediaSelection](m_, objc.Sel("init"))
	return rv
}

// Indicates whether the specified media selection group is subject to automatic media selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselection/1386716-mediaselectioncriteriacanbeappli?language=objc
func (m_ MediaSelection) MediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup(mediaSelectionGroup IMediaSelectionGroup) bool {
	rv := objc.Call[bool](m_, objc.Sel("mediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup:"), objc.Ptr(mediaSelectionGroup))
	return rv
}

// Returns the media selection option that’s currently selected in the specified group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselection/1389197-selectedmediaoptioninmediaselect?language=objc
func (m_ MediaSelection) SelectedMediaOptionInMediaSelectionGroup(mediaSelectionGroup IMediaSelectionGroup) MediaSelectionOption {
	rv := objc.Call[MediaSelectionOption](m_, objc.Sel("selectedMediaOptionInMediaSelectionGroup:"), objc.Ptr(mediaSelectionGroup))
	return rv
}

// The asset associated with the media selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselection/1390874-asset?language=objc
func (m_ MediaSelection) Asset() Asset {
	rv := objc.Call[Asset](m_, objc.Sel("asset"))
	return rv
}
