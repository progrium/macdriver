// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetSegmentTrackReport] class.
var AssetSegmentTrackReportClass = _AssetSegmentTrackReportClass{objc.GetClass("AVAssetSegmentTrackReport")}

type _AssetSegmentTrackReportClass struct {
	objc.Class
}

// An interface definition for the [AssetSegmentTrackReport] class.
type IAssetSegmentTrackReport interface {
	objc.IObject
	EarliestPresentationTimeStamp() coremedia.Time
	TrackID() objc.Object
	MediaType() MediaType
	FirstVideoSampleInformation() AssetSegmentReportSampleInformation
	Duration() coremedia.Time
}

// An object that provides information on a track in segment data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport?language=objc
type AssetSegmentTrackReport struct {
	objc.Object
}

func AssetSegmentTrackReportFrom(ptr unsafe.Pointer) AssetSegmentTrackReport {
	return AssetSegmentTrackReport{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetSegmentTrackReportClass) Alloc() AssetSegmentTrackReport {
	rv := objc.Call[AssetSegmentTrackReport](ac, objc.Sel("alloc"))
	return rv
}

func AssetSegmentTrackReport_Alloc() AssetSegmentTrackReport {
	return AssetSegmentTrackReportClass.Alloc()
}

func (ac _AssetSegmentTrackReportClass) New() AssetSegmentTrackReport {
	rv := objc.Call[AssetSegmentTrackReport](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetSegmentTrackReport() AssetSegmentTrackReport {
	return AssetSegmentTrackReportClass.New()
}

func (a_ AssetSegmentTrackReport) Init() AssetSegmentTrackReport {
	rv := objc.Call[AssetSegmentTrackReport](a_, objc.Sel("init"))
	return rv
}

// The earliest presentation timestamp (PTS) for this track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/3600083-earliestpresentationtimestamp?language=objc
func (a_ AssetSegmentTrackReport) EarliestPresentationTimeStamp() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("earliestPresentationTimeStamp"))
	return rv
}

// A persistent unique identifier for a track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/3546581-trackid?language=objc
func (a_ AssetSegmentTrackReport) TrackID() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("trackID"))
	return rv
}

// The type of media a track contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/3546580-mediatype?language=objc
func (a_ AssetSegmentTrackReport) MediaType() MediaType {
	rv := objc.Call[MediaType](a_, objc.Sel("mediaType"))
	return rv
}

// Information about the first video sample in a track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/3546579-firstvideosampleinformation?language=objc
func (a_ AssetSegmentTrackReport) FirstVideoSampleInformation() AssetSegmentReportSampleInformation {
	rv := objc.Call[AssetSegmentReportSampleInformation](a_, objc.Sel("firstVideoSampleInformation"))
	return rv
}

// The duration of a track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/3546578-duration?language=objc
func (a_ AssetSegmentTrackReport) Duration() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("duration"))
	return rv
}
