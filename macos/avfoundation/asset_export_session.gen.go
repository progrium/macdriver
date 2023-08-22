// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetExportSession] class.
var AssetExportSessionClass = _AssetExportSessionClass{objc.GetClass("AVAssetExportSession")}

type _AssetExportSessionClass struct {
	objc.Class
}

// An interface definition for the [AssetExportSession] class.
type IAssetExportSession interface {
	objc.IObject
	ExportAsynchronouslyWithCompletionHandler(handler func())
	DetermineCompatibleFileTypesWithCompletionHandler(handler func(compatibleFileTypes []FileType))
	EstimateOutputFileLengthWithCompletionHandler(handler func(estimatedOutputFileLength int64, error foundation.Error))
	CancelExport()
	EstimateMaximumDurationWithCompletionHandler(handler func(estimatedMaximumDuration coremedia.Time, error foundation.Error))
	Error() foundation.Error
	MetadataItemFilter() MetadataItemFilter
	SetMetadataItemFilter(value IMetadataItemFilter)
	VideoComposition() VideoComposition
	SetVideoComposition(value IVideoComposition)
	FileLengthLimit() int64
	SetFileLengthLimit(value int64)
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm)
	SupportedFileTypes() []FileType
	Metadata() []MetadataItem
	SetMetadata(value []IMetadataItem)
	PresetName() string
	CanPerformMultiplePassesOverSourceMediaData() bool
	SetCanPerformMultiplePassesOverSourceMediaData(value bool)
	CustomVideoCompositor() VideoCompositingWrapper
	OutputFileType() FileType
	SetOutputFileType(value FileType)
	ShouldOptimizeForNetworkUse() bool
	SetShouldOptimizeForNetworkUse(value bool)
	DirectoryForTemporaryFiles() foundation.URL
	SetDirectoryForTemporaryFiles(value foundation.IURL)
	TimeRange() coremedia.TimeRange
	SetTimeRange(value coremedia.TimeRange)
	AudioMix() AudioMix
	SetAudioMix(value IAudioMix)
	Progress() float64
	Status() AssetExportSessionStatus
	Asset() Asset
	OutputURL() foundation.URL
	SetOutputURL(value foundation.IURL)
}

// An object that exports assets in a format that you specify using an export preset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession?language=objc
type AssetExportSession struct {
	objc.Object
}

func AssetExportSessionFrom(ptr unsafe.Pointer) AssetExportSession {
	return AssetExportSession{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetExportSessionClass) ExportSessionWithAssetPresetName(asset IAsset, presetName string) AssetExportSession {
	rv := objc.Call[AssetExportSession](ac, objc.Sel("exportSessionWithAsset:presetName:"), objc.Ptr(asset), presetName)
	return rv
}

// Returns a new asset export session that uses the specified preset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1564246-exportsessionwithasset?language=objc
func AssetExportSession_ExportSessionWithAssetPresetName(asset IAsset, presetName string) AssetExportSession {
	return AssetExportSessionClass.ExportSessionWithAssetPresetName(asset, presetName)
}

func (a_ AssetExportSession) InitWithAssetPresetName(asset IAsset, presetName string) AssetExportSession {
	rv := objc.Call[AssetExportSession](a_, objc.Sel("initWithAsset:presetName:"), objc.Ptr(asset), presetName)
	return rv
}

// Creates an export session with a preset configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1389367-initwithasset?language=objc
func NewAssetExportSessionWithAssetPresetName(asset IAsset, presetName string) AssetExportSession {
	instance := AssetExportSessionClass.Alloc().InitWithAssetPresetName(asset, presetName)
	instance.Autorelease()
	return instance
}

func (ac _AssetExportSessionClass) Alloc() AssetExportSession {
	rv := objc.Call[AssetExportSession](ac, objc.Sel("alloc"))
	return rv
}

func AssetExportSession_Alloc() AssetExportSession {
	return AssetExportSessionClass.Alloc()
}

func (ac _AssetExportSessionClass) New() AssetExportSession {
	rv := objc.Call[AssetExportSession](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetExportSession() AssetExportSession {
	return AssetExportSessionClass.New()
}

func (a_ AssetExportSession) Init() AssetExportSession {
	rv := objc.Call[AssetExportSession](a_, objc.Sel("init"))
	return rv
}

// Starts the asynchronous execution of an export session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1388005-exportasynchronouslywithcompleti?language=objc
func (a_ AssetExportSession) ExportAsynchronouslyWithCompletionHandler(handler func()) {
	objc.Call[objc.Void](a_, objc.Sel("exportAsynchronouslyWithCompletionHandler:"), handler)
}

// Determines the output file types an asset export session supports writing in its current configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1387907-determinecompatiblefiletypeswith?language=objc
func (a_ AssetExportSession) DetermineCompatibleFileTypesWithCompletionHandler(handler func(compatibleFileTypes []FileType)) {
	objc.Call[objc.Void](a_, objc.Sel("determineCompatibleFileTypesWithCompletionHandler:"), handler)
}

// Starts estimating the output file length of the export while considering the asset, preset, and time range configuration of the export session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/3042921-estimateoutputfilelengthwithcomp?language=objc
func (a_ AssetExportSession) EstimateOutputFileLengthWithCompletionHandler(handler func(estimatedOutputFileLength int64, error foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("estimateOutputFileLengthWithCompletionHandler:"), handler)
}

// Determines an export preset’s compatibility to export the asset in a container of the output file type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1385821-determinecompatibilityofexportpr?language=objc
func (ac _AssetExportSessionClass) DetermineCompatibilityOfExportPresetWithAssetOutputFileTypeCompletionHandler(presetName string, asset IAsset, outputFileType FileType, handler func(compatible bool)) {
	objc.Call[objc.Void](ac, objc.Sel("determineCompatibilityOfExportPreset:withAsset:outputFileType:completionHandler:"), presetName, objc.Ptr(asset), outputFileType, handler)
}

// Determines an export preset’s compatibility to export the asset in a container of the output file type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1385821-determinecompatibilityofexportpr?language=objc
func AssetExportSession_DetermineCompatibilityOfExportPresetWithAssetOutputFileTypeCompletionHandler(presetName string, asset IAsset, outputFileType FileType, handler func(compatible bool)) {
	AssetExportSessionClass.DetermineCompatibilityOfExportPresetWithAssetOutputFileTypeCompletionHandler(presetName, asset, outputFileType, handler)
}

// Returns all available export preset names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1387150-allexportpresets?language=objc
func (ac _AssetExportSessionClass) AllExportPresets() []string {
	rv := objc.Call[[]string](ac, objc.Sel("allExportPresets"))
	return rv
}

// Returns all available export preset names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1387150-allexportpresets?language=objc
func AssetExportSession_AllExportPresets() []string {
	return AssetExportSessionClass.AllExportPresets()
}

// Cancels the execution of an export session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1387794-cancelexport?language=objc
func (a_ AssetExportSession) CancelExport() {
	objc.Call[objc.Void](a_, objc.Sel("cancelExport"))
}

// Starts estimating the maximum duration of the export while considering the asset, preset, and time range configuration of the export session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/3042920-estimatemaximumdurationwithcompl?language=objc
func (a_ AssetExportSession) EstimateMaximumDurationWithCompletionHandler(handler func(estimatedMaximumDuration coremedia.Time, error foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("estimateMaximumDurationWithCompletionHandler:"), handler)
}

// An optional error object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1385936-error?language=objc
func (a_ AssetExportSession) Error() foundation.Error {
	rv := objc.Call[foundation.Error](a_, objc.Sel("error"))
	return rv
}

// An object the export session uses to filter the metadata items it transfers to the output asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1390226-metadataitemfilter?language=objc
func (a_ AssetExportSession) MetadataItemFilter() MetadataItemFilter {
	rv := objc.Call[MetadataItemFilter](a_, objc.Sel("metadataItemFilter"))
	return rv
}

// An object the export session uses to filter the metadata items it transfers to the output asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1390226-metadataitemfilter?language=objc
func (a_ AssetExportSession) SetMetadataItemFilter(value IMetadataItemFilter) {
	objc.Call[objc.Void](a_, objc.Sel("setMetadataItemFilter:"), objc.Ptr(value))
}

// An optional object that provides instructions for how to composite frames of video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1389477-videocomposition?language=objc
func (a_ AssetExportSession) VideoComposition() VideoComposition {
	rv := objc.Call[VideoComposition](a_, objc.Sel("videoComposition"))
	return rv
}

// An optional object that provides instructions for how to composite frames of video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1389477-videocomposition?language=objc
func (a_ AssetExportSession) SetVideoComposition(value IVideoComposition) {
	objc.Call[objc.Void](a_, objc.Sel("setVideoComposition:"), objc.Ptr(value))
}

// The file length that the output of the session must not exceed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1622333-filelengthlimit?language=objc
func (a_ AssetExportSession) FileLengthLimit() int64 {
	rv := objc.Call[int64](a_, objc.Sel("fileLengthLimit"))
	return rv
}

// The file length that the output of the session must not exceed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1622333-filelengthlimit?language=objc
func (a_ AssetExportSession) SetFileLengthLimit(value int64) {
	objc.Call[objc.Void](a_, objc.Sel("setFileLengthLimit:"), value)
}

// A processing algorithm for managing audio pitch for scaled audio edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1385835-audiotimepitchalgorithm?language=objc
func (a_ AssetExportSession) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Call[AudioTimePitchAlgorithm](a_, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}

// A processing algorithm for managing audio pitch for scaled audio edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1385835-audiotimepitchalgorithm?language=objc
func (a_ AssetExportSession) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm) {
	objc.Call[objc.Void](a_, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}

// An array containing the types of files the session can write. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1388762-supportedfiletypes?language=objc
func (a_ AssetExportSession) SupportedFileTypes() []FileType {
	rv := objc.Call[[]FileType](a_, objc.Sel("supportedFileTypes"))
	return rv
}

// The metadata an export session writes to the output container file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1390453-metadata?language=objc
func (a_ AssetExportSession) Metadata() []MetadataItem {
	rv := objc.Call[[]MetadataItem](a_, objc.Sel("metadata"))
	return rv
}

// The metadata an export session writes to the output container file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1390453-metadata?language=objc
func (a_ AssetExportSession) SetMetadata(value []IMetadataItem) {
	objc.Call[objc.Void](a_, objc.Sel("setMetadata:"), value)
}

// The name of the preset that the asset export session uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1390467-presetname?language=objc
func (a_ AssetExportSession) PresetName() string {
	rv := objc.Call[string](a_, objc.Sel("presetName"))
	return rv
}

// A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1388797-canperformmultiplepassesoversour?language=objc
func (a_ AssetExportSession) CanPerformMultiplePassesOverSourceMediaData() bool {
	rv := objc.Call[bool](a_, objc.Sel("canPerformMultiplePassesOverSourceMediaData"))
	return rv
}

// A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1388797-canperformmultiplepassesoversour?language=objc
func (a_ AssetExportSession) SetCanPerformMultiplePassesOverSourceMediaData(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setCanPerformMultiplePassesOverSourceMediaData:"), value)
}

// An optional custom object to use when compositing video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1388288-customvideocompositor?language=objc
func (a_ AssetExportSession) CustomVideoCompositor() VideoCompositingWrapper {
	rv := objc.Call[VideoCompositingWrapper](a_, objc.Sel("customVideoCompositor"))
	return rv
}

// The file type of the output an asset export session writes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1387110-outputfiletype?language=objc
func (a_ AssetExportSession) OutputFileType() FileType {
	rv := objc.Call[FileType](a_, objc.Sel("outputFileType"))
	return rv
}

// The file type of the output an asset export session writes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1387110-outputfiletype?language=objc
func (a_ AssetExportSession) SetOutputFileType(value FileType) {
	objc.Call[objc.Void](a_, objc.Sel("setOutputFileType:"), value)
}

// A Boolean value that indicates whether to optimize the movie for network use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1390593-shouldoptimizefornetworkuse?language=objc
func (a_ AssetExportSession) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Call[bool](a_, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}

// A Boolean value that indicates whether to optimize the movie for network use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1390593-shouldoptimizefornetworkuse?language=objc
func (a_ AssetExportSession) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}

// A directory suitable to store temporary files that the export process generates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1388699-directoryfortemporaryfiles?language=objc
func (a_ AssetExportSession) DirectoryForTemporaryFiles() foundation.URL {
	rv := objc.Call[foundation.URL](a_, objc.Sel("directoryForTemporaryFiles"))
	return rv
}

// A directory suitable to store temporary files that the export process generates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1388699-directoryfortemporaryfiles?language=objc
func (a_ AssetExportSession) SetDirectoryForTemporaryFiles(value foundation.IURL) {
	objc.Call[objc.Void](a_, objc.Sel("setDirectoryForTemporaryFiles:"), objc.Ptr(value))
}

// The time range of the source asset to export. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1388728-timerange?language=objc
func (a_ AssetExportSession) TimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](a_, objc.Sel("timeRange"))
	return rv
}

// The time range of the source asset to export. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1388728-timerange?language=objc
func (a_ AssetExportSession) SetTimeRange(value coremedia.TimeRange) {
	objc.Call[objc.Void](a_, objc.Sel("setTimeRange:"), value)
}

// The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1388155-audiomix?language=objc
func (a_ AssetExportSession) AudioMix() AudioMix {
	rv := objc.Call[AudioMix](a_, objc.Sel("audioMix"))
	return rv
}

// The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1388155-audiomix?language=objc
func (a_ AssetExportSession) SetAudioMix(value IAudioMix) {
	objc.Call[objc.Void](a_, objc.Sel("setAudioMix:"), objc.Ptr(value))
}

// A value that indicates the progress of the export. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1387530-progress?language=objc
func (a_ AssetExportSession) Progress() float64 {
	rv := objc.Call[float64](a_, objc.Sel("progress"))
	return rv
}

// The status of the export session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1390528-status?language=objc
func (a_ AssetExportSession) Status() AssetExportSessionStatus {
	rv := objc.Call[AssetExportSessionStatus](a_, objc.Sel("status"))
	return rv
}

// An asset that a session exports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1385690-asset?language=objc
func (a_ AssetExportSession) Asset() Asset {
	rv := objc.Call[Asset](a_, objc.Sel("asset"))
	return rv
}

// A URL where an asset export session writes its output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1389970-outputurl?language=objc
func (a_ AssetExportSession) OutputURL() foundation.URL {
	rv := objc.Call[foundation.URL](a_, objc.Sel("outputURL"))
	return rv
}

// A URL where an asset export session writes its output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/1389970-outputurl?language=objc
func (a_ AssetExportSession) SetOutputURL(value foundation.IURL) {
	objc.Call[objc.Void](a_, objc.Sel("setOutputURL:"), objc.Ptr(value))
}
