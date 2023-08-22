// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Asset] class.
var AssetClass = _AssetClass{objc.GetClass("AVAsset")}

type _AssetClass struct {
	objc.Class
}

// An interface definition for the [Asset] class.
type IAsset interface {
	objc.IObject
	FindUnusedTrackIDWithCompletionHandler(completionHandler func(arg0 objc.Object, arg1 foundation.Error))
	CancelLoading()
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic, completionHandler func(arg0 []AssetTrack, arg1 foundation.Error))
	LoadMetadataForFormatCompletionHandler(format MetadataFormat, completionHandler func(arg0 []MetadataItem, arg1 foundation.Error))
	LoadTrackWithTrackIDCompletionHandler(trackID objc.IObject, completionHandler func(arg0 AssetTrack, arg1 foundation.Error))
	LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType, completionHandler func(arg0 []AssetTrack, arg1 foundation.Error))
	LoadChapterMetadataGroupsBestMatchingPreferredLanguagesCompletionHandler(preferredLanguages []string, completionHandler func(arg0 []TimedMetadataGroup, arg1 foundation.Error))
	LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic, completionHandler func(arg0 MediaSelectionGroup, arg1 foundation.Error))
	LoadChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeysCompletionHandler(locale foundation.ILocale, commonKeys []MetadataKey, completionHandler func(arg0 []TimedMetadataGroup, arg1 foundation.Error))
	MinimumTimeOffsetFromLive() coremedia.Time
	Lyrics() string
	Tracks() []AssetTrack
	IsReadable() bool
	AvailableMediaCharacteristicsWithMediaSelectionOptions() []MediaCharacteristic
	PreferredVolume() float64
	CanContainFragments() bool
	AvailableChapterLocales() []foundation.Locale
	IsComposable() bool
	ProvidesPreciseDurationAndTiming() bool
	Metadata() []MetadataItem
	PreferredRate() float64
	AvailableMetadataFormats() []MetadataFormat
	OverallDurationHint() coremedia.Time
	CreationDate() MetadataItem
	TrackGroups() []AssetTrackGroup
	ReferenceRestrictions() AssetReferenceRestrictions
	AllMediaSelections() []MediaSelection
	HasProtectedContent() bool
	Duration() coremedia.Time
	PreferredMediaSelection() MediaSelection
	IsCompatibleWithAirPlayVideo() bool
	IsExportable() bool
	ContainsFragments() bool
	PreferredTransform() coregraphics.AffineTransform
	IsPlayable() bool
	CommonMetadata() []MetadataItem
}

// An object that models timed audiovisual media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset?language=objc
type Asset struct {
	objc.Object
}

func AssetFrom(ptr unsafe.Pointer) Asset {
	return Asset{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetClass) AssetWithURL(URL foundation.IURL) Asset {
	rv := objc.Call[Asset](ac, objc.Sel("assetWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates an asset that models the media at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389943-assetwithurl?language=objc
func Asset_AssetWithURL(URL foundation.IURL) Asset {
	return AssetClass.AssetWithURL(URL)
}

func (ac _AssetClass) Alloc() Asset {
	rv := objc.Call[Asset](ac, objc.Sel("alloc"))
	return rv
}

func Asset_Alloc() Asset {
	return AssetClass.Alloc()
}

func (ac _AssetClass) New() Asset {
	rv := objc.Call[Asset](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAsset() Asset {
	return AssetClass.New()
}

func (a_ Asset) Init() Asset {
	rv := objc.Call[Asset](a_, objc.Sel("init"))
	return rv
}

// Loads an identifier that no other track in the asset uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/3746590-findunusedtrackidwithcompletionh?language=objc
func (a_ Asset) FindUnusedTrackIDWithCompletionHandler(completionHandler func(arg0 objc.Object, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("findUnusedTrackIDWithCompletionHandler:"), completionHandler)
}

// Cancels all pending requests to asynchronously load property values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1388722-cancelloading?language=objc
func (a_ Asset) CancelLoading() {
	objc.Call[objc.Void](a_, objc.Sel("cancelLoading"))
}

// Loads tracks that contain media of a specified characteristic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/3746530-loadtrackswithmediacharacteristi?language=objc
func (a_ Asset) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic, completionHandler func(arg0 []AssetTrack, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}

// Loads an array of metadata items that the asset contains for the specified format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/3746528-loadmetadataforformat?language=objc
func (a_ Asset) LoadMetadataForFormatCompletionHandler(format MetadataFormat, completionHandler func(arg0 []MetadataItem, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("loadMetadataForFormat:completionHandler:"), format, completionHandler)
}

// Loads a track that contains the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/3746529-loadtrackwithtrackid?language=objc
func (a_ Asset) LoadTrackWithTrackIDCompletionHandler(trackID objc.IObject, completionHandler func(arg0 AssetTrack, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("loadTrackWithTrackID:completionHandler:"), objc.Ptr(trackID), completionHandler)
}

// Loads tracks that contain media of a specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/3746531-loadtrackswithmediatype?language=objc
func (a_ Asset) LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType, completionHandler func(arg0 []AssetTrack, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("loadTracksWithMediaType:completionHandler:"), mediaType, completionHandler)
}

// Loads chapter metadata with a locale that best matches the list of preferred languages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/3746525-loadchaptermetadatagroupsbestmat?language=objc
func (a_ Asset) LoadChapterMetadataGroupsBestMatchingPreferredLanguagesCompletionHandler(preferredLanguages []string, completionHandler func(arg0 []TimedMetadataGroup, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("loadChapterMetadataGroupsBestMatchingPreferredLanguages:completionHandler:"), preferredLanguages, completionHandler)
}

// Loads a media selection group that contains one or more options with the specified media characteristic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/3746527-loadmediaselectiongroupformediac?language=objc
func (a_ Asset) LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic, completionHandler func(arg0 MediaSelectionGroup, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("loadMediaSelectionGroupForMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}

// Loads chapter metadata that contains the specified title locale and common keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/3746526-loadchaptermetadatagroupswithtit?language=objc
func (a_ Asset) LoadChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeysCompletionHandler(locale foundation.ILocale, commonKeys []MetadataKey, completionHandler func(arg0 []TimedMetadataGroup, arg1 foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("loadChapterMetadataGroupsWithTitleLocale:containingItemsWithCommonKeys:completionHandler:"), objc.Ptr(locale), commonKeys, completionHandler)
}

// A time value that indicates how closely playback follows the latest live stream content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/3197641-minimumtimeoffsetfromlive?language=objc
func (a_ Asset) MinimumTimeOffsetFromLive() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("minimumTimeOffsetFromLive"))
	return rv
}

// The lyrics of the asset in a language suitable for the current locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1388104-lyrics?language=objc
func (a_ Asset) Lyrics() string {
	rv := objc.Call[string](a_, objc.Sel("lyrics"))
	return rv
}

// The tracks an asset contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1387953-tracks?language=objc
func (a_ Asset) Tracks() []AssetTrack {
	rv := objc.Call[[]AssetTrack](a_, objc.Sel("tracks"))
	return rv
}

// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1390475-readable?language=objc
func (a_ Asset) IsReadable() bool {
	rv := objc.Call[bool](a_, objc.Sel("isReadable"))
	return rv
}

// An array of media characteristics for which a media selection option is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389433-availablemediacharacteristicswit?language=objc
func (a_ Asset) AvailableMediaCharacteristicsWithMediaSelectionOptions() []MediaCharacteristic {
	rv := objc.Call[[]MediaCharacteristic](a_, objc.Sel("availableMediaCharacteristicsWithMediaSelectionOptions"))
	return rv
}

// The asset’s volume preference for playing its audible media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1390457-preferredvolume?language=objc
func (a_ Asset) PreferredVolume() float64 {
	rv := objc.Call[float64](a_, objc.Sel("preferredVolume"))
	return rv
}

// A Boolean value that indicates whether you can extend the asset by fragments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389520-cancontainfragments?language=objc
func (a_ Asset) CanContainFragments() bool {
	rv := objc.Call[bool](a_, objc.Sel("canContainFragments"))
	return rv
}

// The locales of the asset’s chapter metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1388228-availablechapterlocales?language=objc
func (a_ Asset) AvailableChapterLocales() []foundation.Locale {
	rv := objc.Call[[]foundation.Locale](a_, objc.Sel("availableChapterLocales"))
	return rv
}

// A Boolean value that indicates whether you can use the asset as a segment of a composition track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1386129-composable?language=objc
func (a_ Asset) IsComposable() bool {
	rv := objc.Call[bool](a_, objc.Sel("isComposable"))
	return rv
}

// A Boolean value that indicates whether the asset provides precise duration and timing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1390850-providesprecisedurationandtiming?language=objc
func (a_ Asset) ProvidesPreciseDurationAndTiming() bool {
	rv := objc.Call[bool](a_, objc.Sel("providesPreciseDurationAndTiming"))
	return rv
}

// An array of metadata items for all metadata identifiers for which a value is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1386884-metadata?language=objc
func (a_ Asset) Metadata() []MetadataItem {
	rv := objc.Call[[]MetadataItem](a_, objc.Sel("metadata"))
	return rv
}

// The asset’s rate preference for playing its media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389912-preferredrate?language=objc
func (a_ Asset) PreferredRate() float64 {
	rv := objc.Call[float64](a_, objc.Sel("preferredRate"))
	return rv
}

// The metadata formats this asset contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1385823-availablemetadataformats?language=objc
func (a_ Asset) AvailableMetadataFormats() []MetadataFormat {
	rv := objc.Call[[]MetadataFormat](a_, objc.Sel("availableMetadataFormats"))
	return rv
}

// The total duration of fragments that currently exist, or may exist in the future. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/2715834-overalldurationhint?language=objc
func (a_ Asset) OverallDurationHint() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("overallDurationHint"))
	return rv
}

// A metadata item that indicates the asset’s creation date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1386342-creationdate?language=objc
func (a_ Asset) CreationDate() MetadataItem {
	rv := objc.Call[MetadataItem](a_, objc.Sel("creationDate"))
	return rv
}

// The track groups an asset contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1390697-trackgroups?language=objc
func (a_ Asset) TrackGroups() []AssetTrackGroup {
	rv := objc.Call[[]AssetTrackGroup](a_, objc.Sel("trackGroups"))
	return rv
}

// The restrictions that an asset places on how it resolves references to external media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1390489-referencerestrictions?language=objc
func (a_ Asset) ReferenceRestrictions() AssetReferenceRestrictions {
	rv := objc.Call[AssetReferenceRestrictions](a_, objc.Sel("referenceRestrictions"))
	return rv
}

// The array of available media selections for this asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/2890796-allmediaselections?language=objc
func (a_ Asset) AllMediaSelections() []MediaSelection {
	rv := objc.Call[[]MediaSelection](a_, objc.Sel("allMediaSelections"))
	return rv
}

// A Boolean value that indicates whether the asset contains protected content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389223-hasprotectedcontent?language=objc
func (a_ Asset) HasProtectedContent() bool {
	rv := objc.Call[bool](a_, objc.Sel("hasProtectedContent"))
	return rv
}

// A time value that indicates the asset’s duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389400-duration?language=objc
func (a_ Asset) Duration() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("duration"))
	return rv
}

// The default media selections for this asset’s media selection groups. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1386122-preferredmediaselection?language=objc
func (a_ Asset) PreferredMediaSelection() MediaSelection {
	rv := objc.Call[MediaSelection](a_, objc.Sel("preferredMediaSelection"))
	return rv
}

// A Boolean value that indicates whether the asset is compatible with AirPlay Video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1390333-compatiblewithairplayvideo?language=objc
func (a_ Asset) IsCompatibleWithAirPlayVideo() bool {
	rv := objc.Call[bool](a_, objc.Sel("isCompatibleWithAirPlayVideo"))
	return rv
}

// A Boolean value that indicates whether you can export this asset using an export session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389245-exportable?language=objc
func (a_ Asset) IsExportable() bool {
	rv := objc.Call[bool](a_, objc.Sel("isExportable"))
	return rv
}

// A Boolean value that indicates whether at least one movie fragment extends the asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1385589-containsfragments?language=objc
func (a_ Asset) ContainsFragments() bool {
	rv := objc.Call[bool](a_, objc.Sel("containsFragments"))
	return rv
}

// The asset’s transform preference to apply to its visual content during presentation or processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1385906-preferredtransform?language=objc
func (a_ Asset) PreferredTransform() coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](a_, objc.Sel("preferredTransform"))
	return rv
}

// A Boolean value that indicates whether the asset has playable content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1385974-playable?language=objc
func (a_ Asset) IsPlayable() bool {
	rv := objc.Call[bool](a_, objc.Sel("isPlayable"))
	return rv
}

// The metadata items an asset contains for common metadata identifiers that provide a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1390498-commonmetadata?language=objc
func (a_ Asset) CommonMetadata() []MetadataItem {
	rv := objc.Call[[]MetadataItem](a_, objc.Sel("commonMetadata"))
	return rv
}
