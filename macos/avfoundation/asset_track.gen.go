// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetTrack] class.
var AssetTrackClass = _AssetTrackClass{objc.GetClass("AVAssetTrack")}

type _AssetTrackClass struct {
	objc.Class
}

// An interface definition for the [AssetTrack] class.
type IAssetTrack interface {
	objc.IObject
	MakeSampleCursorWithPresentationTimeStamp(presentationTimeStamp coremedia.Time) SampleCursor
	LoadSamplePresentationTimeForTrackTimeCompletionHandler(trackTime coremedia.Time, completionHandler func(arg0 coremedia.Time, arg1 foundation.Error))
	LoadMetadataForFormatCompletionHandler(format MetadataFormat, completionHandler func(arg0 []MetadataItem, arg1 foundation.Error))
	HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic) bool
	MakeSampleCursorAtFirstSampleInDecodeOrder() SampleCursor
	LoadAssociatedTracksOfTypeCompletionHandler(trackAssociationType TrackAssociationType, completionHandler func(arg0 []AssetTrack, arg1 foundation.Error))
	MakeSampleCursorAtLastSampleInDecodeOrder() SampleCursor
	LoadSegmentForTrackTimeCompletionHandler(trackTime coremedia.Time, completionHandler func(arg0 AssetTrackSegment, arg1 foundation.Error))
	ExtendedLanguageTag() string
	FormatDescriptions() []objc.Object
	TrackID() objc.Object
	PreferredVolume() float64
	HasAudioSampleDependencies() bool
	Metadata() []MetadataItem
	MediaType() MediaType
	RequiresFrameReordering() bool
	TotalSampleDataLength() int64
	Segments() []AssetTrackSegment
	IsDecodable() bool
	MinFrameDuration() coremedia.Time
	NaturalSize() coregraphics.Size
	AvailableMetadataFormats() []MetadataFormat
	EstimatedDataRate() float64
	TimeRange() coremedia.TimeRange
	NominalFrameRate() float64
	CanProvideSampleCursors() bool
	IsSelfContained() bool
	LanguageCode() string
	IsEnabled() bool
	NaturalTimeScale() coremedia.TimeScale
	AvailableTrackAssociationTypes() []TrackAssociationType
	Asset() Asset
	PreferredTransform() coregraphics.AffineTransform
	IsPlayable() bool
	CommonMetadata() []MetadataItem
}

// An object that models a track of media that an asset contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack?language=objc
type AssetTrack struct {
	objc.Object
}

func AssetTrackFrom(ptr unsafe.Pointer) AssetTrack {
	return AssetTrack{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetTrackClass) Alloc() AssetTrack {
	rv := objc.Call[AssetTrack](ac, objc.Sel("alloc"))
	return rv
}

func AssetTrack_Alloc() AssetTrack {
	return AssetTrackClass.Alloc()
}

func (ac _AssetTrackClass) New() AssetTrack {
	rv := objc.Call[AssetTrack](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetTrack() AssetTrack {
	return AssetTrackClass.New()
}

func (a_ AssetTrack) Init() AssetTrack {
	rv := objc.Call[AssetTrack](a_, objc.Sel("init"))
	return rv
}

// Creates a sample cursor and positions it at or near the specified presentation timestamp. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1390248-makesamplecursorwithpresentation?language=objc
func (a_ AssetTrack) MakeSampleCursorWithPresentationTimeStamp(presentationTimeStamp coremedia.Time) SampleCursor {
	rv := objc.Call[SampleCursor](a_, objc.Sel("makeSampleCursorWithPresentationTimeStamp:"), presentationTimeStamp)
	return rv
}

// Loads a sample presentation time that maps to the specified track time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/3746540-loadsamplepresentationtimefortra?language=objc
func (a_ AssetTrack) LoadSamplePresentationTimeForTrackTimeCompletionHandler(trackTime coremedia.Time, completionHandler func(arg0 coremedia.Time, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("loadSamplePresentationTimeForTrackTime:completionHandler:"), trackTime, completionHandler)
}

// Loads metadata items that a track contains for the specified format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/3746539-loadmetadataforformat?language=objc
func (a_ AssetTrack) LoadMetadataForFormatCompletionHandler(format MetadataFormat, completionHandler func(arg0 []MetadataItem, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("loadMetadataForFormat:completionHandler:"), format, completionHandler)
}

// Returns a Boolean value that indicates whether the track references media with the specified media characteristic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1385847-hasmediacharacteristic?language=objc
func (a_ AssetTrack) HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic) bool {
	rv := objc.Call[bool](a_, objc.Sel("hasMediaCharacteristic:"), mediaCharacteristic)
	return rv
}

// Creates a sample cursor and positions it at the track’s first media sample in decode order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1387226-makesamplecursoratfirstsampleind?language=objc
func (a_ AssetTrack) MakeSampleCursorAtFirstSampleInDecodeOrder() SampleCursor {
	rv := objc.Call[SampleCursor](a_, objc.Sel("makeSampleCursorAtFirstSampleInDecodeOrder"))
	return rv
}

// Loads associated tracks that have the specified association type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/3746538-loadassociatedtracksoftype?language=objc
func (a_ AssetTrack) LoadAssociatedTracksOfTypeCompletionHandler(trackAssociationType TrackAssociationType, completionHandler func(arg0 []AssetTrack, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("loadAssociatedTracksOfType:completionHandler:"), trackAssociationType, completionHandler)
}

// Creates a sample cursor and positions it at the track’s last media sample in decode order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1386014-makesamplecursoratlastsampleinde?language=objc
func (a_ AssetTrack) MakeSampleCursorAtLastSampleInDecodeOrder() SampleCursor {
	rv := objc.Call[SampleCursor](a_, objc.Sel("makeSampleCursorAtLastSampleInDecodeOrder"))
	return rv
}

// Loads a segment with a target time range that contains, or is closest to, the specified track time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/3746541-loadsegmentfortracktime?language=objc
func (a_ AssetTrack) LoadSegmentForTrackTimeCompletionHandler(trackTime coremedia.Time, completionHandler func(arg0 AssetTrackSegment, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("loadSegmentForTrackTime:completionHandler:"), trackTime, completionHandler)
}

// The language tag of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1389105-extendedlanguagetag?language=objc
func (a_ AssetTrack) ExtendedLanguageTag() string {
	rv := objc.Call[string](a_, objc.Sel("extendedLanguageTag"))
	return rv
}

// The format descriptions of the media samples that a track references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1386694-formatdescriptions?language=objc
func (a_ AssetTrack) FormatDescriptions() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("formatDescriptions"))
	return rv
}

// The persistent unique identifier for this track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1385799-trackid?language=objc
func (a_ AssetTrack) TrackID() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("trackID"))
	return rv
}

// The track’s volume preference for playing its audible media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1388832-preferredvolume?language=objc
func (a_ AssetTrack) PreferredVolume() float64 {
	rv := objc.Call[float64](a_, objc.Sel("preferredVolume"))
	return rv
}

// A Boolean value that indicates whether the track has sample dependencies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/3131265-hasaudiosampledependencies?language=objc
func (a_ AssetTrack) HasAudioSampleDependencies() bool {
	rv := objc.Call[bool](a_, objc.Sel("hasAudioSampleDependencies"))
	return rv
}

// An array of metadata items for all metadata identifiers that have a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1389054-metadata?language=objc
func (a_ AssetTrack) Metadata() []MetadataItem {
	rv := objc.Call[[]MetadataItem](a_, objc.Sel("metadata"))
	return rv
}

// The type of media that a track presents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1385741-mediatype?language=objc
func (a_ AssetTrack) MediaType() MediaType {
	rv := objc.Call[MediaType](a_, objc.Sel("mediaType"))
	return rv
}

// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1390844-requiresframereordering?language=objc
func (a_ AssetTrack) RequiresFrameReordering() bool {
	rv := objc.Call[bool](a_, objc.Sel("requiresFrameReordering"))
	return rv
}

// The total number of bytes of sample data the track requires. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1388900-totalsampledatalength?language=objc
func (a_ AssetTrack) TotalSampleDataLength() int64 {
	rv := objc.Call[int64](a_, objc.Sel("totalSampleDataLength"))
	return rv
}

// The time mappings from the track’s media samples to its timeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1390665-segments?language=objc
func (a_ AssetTrack) Segments() []AssetTrackSegment {
	rv := objc.Call[[]AssetTrackSegment](a_, objc.Sel("segments"))
	return rv
}

// A Boolean value that indicates whether the track is decodable in the current environment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/2887366-decodable?language=objc
func (a_ AssetTrack) IsDecodable() bool {
	rv := objc.Call[bool](a_, objc.Sel("isDecodable"))
	return rv
}

// The minimum duration of the track’s frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1388608-minframeduration?language=objc
func (a_ AssetTrack) MinFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("minFrameDuration"))
	return rv
}

// The natural dimensions of the media data that the track references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1387724-naturalsize?language=objc
func (a_ AssetTrack) NaturalSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](a_, objc.Sel("naturalSize"))
	return rv
}

// An array of metadata formats available for the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1385751-availablemetadataformats?language=objc
func (a_ AssetTrack) AvailableMetadataFormats() []MetadataFormat {
	rv := objc.Call[[]MetadataFormat](a_, objc.Sel("availableMetadataFormats"))
	return rv
}

// The estimated data rate, in bits per second, of the media that the track references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1389758-estimateddatarate?language=objc
func (a_ AssetTrack) EstimatedDataRate() float64 {
	rv := objc.Call[float64](a_, objc.Sel("estimatedDataRate"))
	return rv
}

// The time range of the track within the overall timeline of the asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1388335-timerange?language=objc
func (a_ AssetTrack) TimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](a_, objc.Sel("timeRange"))
	return rv
}

// The frame rate of the track, in frames per second. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1386182-nominalframerate?language=objc
func (a_ AssetTrack) NominalFrameRate() float64 {
	rv := objc.Call[float64](a_, objc.Sel("nominalFrameRate"))
	return rv
}

// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1386692-canprovidesamplecursors?language=objc
func (a_ AssetTrack) CanProvideSampleCursors() bool {
	rv := objc.Call[bool](a_, objc.Sel("canProvideSampleCursors"))
	return rv
}

// A Boolean value that indicates whether this track references sample data only within its container file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1387643-selfcontained?language=objc
func (a_ AssetTrack) IsSelfContained() bool {
	rv := objc.Call[bool](a_, objc.Sel("isSelfContained"))
	return rv
}

// The language code of the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1388627-languagecode?language=objc
func (a_ AssetTrack) LanguageCode() string {
	rv := objc.Call[string](a_, objc.Sel("languageCode"))
	return rv
}

// A Boolean value that indicates whether the track’s container enables it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1387546-enabled?language=objc
func (a_ AssetTrack) IsEnabled() bool {
	rv := objc.Call[bool](a_, objc.Sel("isEnabled"))
	return rv
}

// The natural time scale of the media that a track references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1389233-naturaltimescale?language=objc
func (a_ AssetTrack) NaturalTimeScale() coremedia.TimeScale {
	rv := objc.Call[coremedia.TimeScale](a_, objc.Sel("naturalTimeScale"))
	return rv
}

// An array of association types that the track uses to associate with other tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1388065-availabletrackassociationtypes?language=objc
func (a_ AssetTrack) AvailableTrackAssociationTypes() []TrackAssociationType {
	rv := objc.Call[[]TrackAssociationType](a_, objc.Sel("availableTrackAssociationTypes"))
	return rv
}

// The asset object that contains this track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1385611-asset?language=objc
func (a_ AssetTrack) Asset() Asset {
	rv := objc.Call[Asset](a_, objc.Sel("asset"))
	return rv
}

// The track’s transform preference to apply to its visual content during presentation or processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1389837-preferredtransform?language=objc
func (a_ AssetTrack) PreferredTransform() coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](a_, objc.Sel("preferredTransform"))
	return rv
}

// A Boolean value that indicates whether the track is playable in the current environment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1388276-playable?language=objc
func (a_ AssetTrack) IsPlayable() bool {
	rv := objc.Call[bool](a_, objc.Sel("isPlayable"))
	return rv
}

// An array of metadata items for all common metadata keys that have a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrack/1390832-commonmetadata?language=objc
func (a_ AssetTrack) CommonMetadata() []MetadataItem {
	rv := objc.Call[[]MetadataItem](a_, objc.Sel("commonMetadata"))
	return rv
}
