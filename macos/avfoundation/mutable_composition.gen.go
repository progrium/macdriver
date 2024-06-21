// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MutableComposition] class.
var MutableCompositionClass = _MutableCompositionClass{objc.GetClass("AVMutableComposition")}

type _MutableCompositionClass struct {
	objc.Class
}

// An interface definition for the [MutableComposition] class.
type IMutableComposition interface {
	IComposition
	RemoveTimeRange(timeRange coremedia.TimeRange)
	MutableTrackCompatibleWithTrack(track IAssetTrack) MutableCompositionTrack
	RemoveTrack(track ICompositionTrack)
	AddMutableTrackWithMediaTypePreferredTrackID(mediaType MediaType, preferredTrackID objc.IObject) MutableCompositionTrack
	ScaleTimeRangeToDuration(timeRange coremedia.TimeRange, duration coremedia.Time)
	InsertEmptyTimeRange(timeRange coremedia.TimeRange)
	SetNaturalSize(value coregraphics.Size)
}

// An object that you use to create a new composition from existing assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition?language=objc
type MutableComposition struct {
	Composition
}

func MutableCompositionFrom(ptr unsafe.Pointer) MutableComposition {
	return MutableComposition{
		Composition: CompositionFrom(ptr),
	}
}

func (mc _MutableCompositionClass) CompositionWithURLAssetInitializationOptions(URLAssetInitializationOptions map[string]objc.IObject) MutableComposition {
	rv := objc.Call[MutableComposition](mc, objc.Sel("compositionWithURLAssetInitializationOptions:"), URLAssetInitializationOptions)
	return rv
}

// Creates a mutable composition that uses the specified initialization options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/1390705-compositionwithurlassetinitializ?language=objc
func MutableComposition_CompositionWithURLAssetInitializationOptions(URLAssetInitializationOptions map[string]objc.IObject) MutableComposition {
	return MutableCompositionClass.CompositionWithURLAssetInitializationOptions(URLAssetInitializationOptions)
}

func (mc _MutableCompositionClass) Composition() MutableComposition {
	rv := objc.Call[MutableComposition](mc, objc.Sel("composition"))
	return rv
}

// Returns a new mutable composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/1495098-composition?language=objc
func MutableComposition_Composition() MutableComposition {
	return MutableCompositionClass.Composition()
}

func (mc _MutableCompositionClass) Alloc() MutableComposition {
	rv := objc.Call[MutableComposition](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MutableCompositionClass) New() MutableComposition {
	rv := objc.Call[MutableComposition](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableComposition() MutableComposition {
	return MutableCompositionClass.New()
}

func (m_ MutableComposition) Init() MutableComposition {
	rv := objc.Call[MutableComposition](m_, objc.Sel("init"))
	return rv
}

func (mc _MutableCompositionClass) AssetWithURL(URL foundation.IURL) MutableComposition {
	rv := objc.Call[MutableComposition](mc, objc.Sel("assetWithURL:"), URL)
	return rv
}

// Creates an asset that models the media at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389943-assetwithurl?language=objc
func MutableComposition_AssetWithURL(URL foundation.IURL) MutableComposition {
	return MutableCompositionClass.AssetWithURL(URL)
}

// Removes a specified time range from all tracks of the composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/1387768-removetimerange?language=objc
func (m_ MutableComposition) RemoveTimeRange(timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("removeTimeRange:"), timeRange)
}

// Returns a composition track into which you can insert any time range of the specified asset track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/1386662-mutabletrackcompatiblewithtrack?language=objc
func (m_ MutableComposition) MutableTrackCompatibleWithTrack(track IAssetTrack) MutableCompositionTrack {
	rv := objc.Call[MutableCompositionTrack](m_, objc.Sel("mutableTrackCompatibleWithTrack:"), track)
	return rv
}

// Removes a specified track from the composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/1386818-removetrack?language=objc
func (m_ MutableComposition) RemoveTrack(track ICompositionTrack) {
	objc.Call[objc.Void](m_, objc.Sel("removeTrack:"), track)
}

// Adds an empty track to a composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/1387601-addmutabletrackwithmediatype?language=objc
func (m_ MutableComposition) AddMutableTrackWithMediaTypePreferredTrackID(mediaType MediaType, preferredTrackID objc.IObject) MutableCompositionTrack {
	rv := objc.Call[MutableCompositionTrack](m_, objc.Sel("addMutableTrackWithMediaType:preferredTrackID:"), mediaType, preferredTrackID)
	return rv
}

// Changes the duration of all tracks in a given time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/1390549-scaletimerange?language=objc
func (m_ MutableComposition) ScaleTimeRangeToDuration(timeRange coremedia.TimeRange, duration coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}

// Adds or extends an empty time range within all tracks of the composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/1386710-insertemptytimerange?language=objc
func (m_ MutableComposition) InsertEmptyTimeRange(timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("insertEmptyTimeRange:"), timeRange)
}

// The encoded or authored size of the visual portion of the asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/1390424-naturalsize?language=objc
func (m_ MutableComposition) SetNaturalSize(value coregraphics.Size) {
	objc.Call[objc.Void](m_, objc.Sel("setNaturalSize:"), value)
}
