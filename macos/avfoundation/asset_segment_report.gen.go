// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetSegmentReport] class.
var AssetSegmentReportClass = _AssetSegmentReportClass{objc.GetClass("AVAssetSegmentReport")}

type _AssetSegmentReportClass struct {
	objc.Class
}

// An interface definition for the [AssetSegmentReport] class.
type IAssetSegmentReport interface {
	objc.IObject
	SegmentType() AssetSegmentType
	TrackReports() []AssetSegmentTrackReport
}

// An object that provides information about segment data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport?language=objc
type AssetSegmentReport struct {
	objc.Object
}

func AssetSegmentReportFrom(ptr unsafe.Pointer) AssetSegmentReport {
	return AssetSegmentReport{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetSegmentReportClass) Alloc() AssetSegmentReport {
	rv := objc.Call[AssetSegmentReport](ac, objc.Sel("alloc"))
	return rv
}

func AssetSegmentReport_Alloc() AssetSegmentReport {
	return AssetSegmentReportClass.Alloc()
}

func (ac _AssetSegmentReportClass) New() AssetSegmentReport {
	rv := objc.Call[AssetSegmentReport](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetSegmentReport() AssetSegmentReport {
	return AssetSegmentReportClass.New()
}

func (a_ AssetSegmentReport) Init() AssetSegmentReport {
	rv := objc.Call[AssetSegmentReport](a_, objc.Sel("init"))
	return rv
}

// The type of segment data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/3546570-segmenttype?language=objc
func (a_ AssetSegmentReport) SegmentType() AssetSegmentType {
	rv := objc.Call[AssetSegmentType](a_, objc.Sel("segmentType"))
	return rv
}

// The reports for the segment’s track data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/3546571-trackreports?language=objc
func (a_ AssetSegmentReport) TrackReports() []AssetSegmentTrackReport {
	rv := objc.Call[[]AssetSegmentTrackReport](a_, objc.Sel("trackReports"))
	return rv
}
