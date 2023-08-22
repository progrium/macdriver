// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetReaderOutput] class.
var AssetReaderOutputClass = _AssetReaderOutputClass{objc.GetClass("AVAssetReaderOutput")}

type _AssetReaderOutputClass struct {
	objc.Class
}

// An interface definition for the [AssetReaderOutput] class.
type IAssetReaderOutput interface {
	objc.IObject
	ResetForReadingTimeRanges(timeRanges []foundation.IValue)
	CopyNextSampleBuffer() coremedia.SampleBufferRef
	MarkConfigurationAsFinal()
	MediaType() MediaType
	SupportsRandomAccess() bool
	SetSupportsRandomAccess(value bool)
	AlwaysCopiesSampleData() bool
	SetAlwaysCopiesSampleData(value bool)
}

// An abstract class that defines the interface to read media samples from an asset reader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput?language=objc
type AssetReaderOutput struct {
	objc.Object
}

func AssetReaderOutputFrom(ptr unsafe.Pointer) AssetReaderOutput {
	return AssetReaderOutput{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetReaderOutputClass) Alloc() AssetReaderOutput {
	rv := objc.Call[AssetReaderOutput](ac, objc.Sel("alloc"))
	return rv
}

func AssetReaderOutput_Alloc() AssetReaderOutput {
	return AssetReaderOutputClass.Alloc()
}

func (ac _AssetReaderOutputClass) New() AssetReaderOutput {
	rv := objc.Call[AssetReaderOutput](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetReaderOutput() AssetReaderOutput {
	return AssetReaderOutputClass.New()
}

func (a_ AssetReaderOutput) Init() AssetReaderOutput {
	rv := objc.Call[AssetReaderOutput](a_, objc.Sel("init"))
	return rv
}

// Restarts reading with a new set of time ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/1388890-resetforreadingtimeranges?language=objc
func (a_ AssetReaderOutput) ResetForReadingTimeRanges(timeRanges []foundation.IValue) {
	objc.Call[objc.Void](a_, objc.Sel("resetForReadingTimeRanges:"), timeRanges)
}

// Copies the next sample buffer from the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/1385732-copynextsamplebuffer?language=objc
func (a_ AssetReaderOutput) CopyNextSampleBuffer() coremedia.SampleBufferRef {
	rv := objc.Call[coremedia.SampleBufferRef](a_, objc.Sel("copyNextSampleBuffer"))
	return rv
}

// Tells the output that it’s finished reconfiguring time ranges, and allows the asset reader to advance to a completed state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/1386974-markconfigurationasfinal?language=objc
func (a_ AssetReaderOutput) MarkConfigurationAsFinal() {
	objc.Call[objc.Void](a_, objc.Sel("markConfigurationAsFinal"))
}

// The media type of samples that the output reads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/1390880-mediatype?language=objc
func (a_ AssetReaderOutput) MediaType() MediaType {
	rv := objc.Call[MediaType](a_, objc.Sel("mediaType"))
	return rv
}

// A Boolean value that indicates whether the output supports reconfiguring the time ranges it reads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/1387927-supportsrandomaccess?language=objc
func (a_ AssetReaderOutput) SupportsRandomAccess() bool {
	rv := objc.Call[bool](a_, objc.Sel("supportsRandomAccess"))
	return rv
}

// A Boolean value that indicates whether the output supports reconfiguring the time ranges it reads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/1387927-supportsrandomaccess?language=objc
func (a_ AssetReaderOutput) SetSupportsRandomAccess(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setSupportsRandomAccess:"), value)
}

// A Boolean value that indicates whether the output vends copied sample data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/1389189-alwayscopiessampledata?language=objc
func (a_ AssetReaderOutput) AlwaysCopiesSampleData() bool {
	rv := objc.Call[bool](a_, objc.Sel("alwaysCopiesSampleData"))
	return rv
}

// A Boolean value that indicates whether the output vends copied sample data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/1389189-alwayscopiessampledata?language=objc
func (a_ AssetReaderOutput) SetAlwaysCopiesSampleData(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAlwaysCopiesSampleData:"), value)
}
