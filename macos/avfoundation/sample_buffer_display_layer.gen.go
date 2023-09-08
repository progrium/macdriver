// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SampleBufferDisplayLayer] class.
var SampleBufferDisplayLayerClass = _SampleBufferDisplayLayerClass{objc.GetClass("AVSampleBufferDisplayLayer")}

type _SampleBufferDisplayLayerClass struct {
	objc.Class
}

// An interface definition for the [SampleBufferDisplayLayer] class.
type ISampleBufferDisplayLayer interface {
	quartzcore.ILayer
	PreventsCapture() bool
	SetPreventsCapture(value bool)
	ControlTimebase() coremedia.TimebaseRef
	SetControlTimebase(value coremedia.TimebaseRef)
	OutputObscuredDueToInsufficientExternalProtection() bool
	VideoGravity() LayerVideoGravity
	SetVideoGravity(value LayerVideoGravity)
	PreventsDisplaySleepDuringVideoPlayback() bool
	SetPreventsDisplaySleepDuringVideoPlayback(value bool)
}

// An object that displays compressed or uncompressed video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer?language=objc
type SampleBufferDisplayLayer struct {
	quartzcore.Layer
}

func SampleBufferDisplayLayerFrom(ptr unsafe.Pointer) SampleBufferDisplayLayer {
	return SampleBufferDisplayLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}

func (sc _SampleBufferDisplayLayerClass) Alloc() SampleBufferDisplayLayer {
	rv := objc.Call[SampleBufferDisplayLayer](sc, objc.Sel("alloc"))
	return rv
}

func (sc _SampleBufferDisplayLayerClass) New() SampleBufferDisplayLayer {
	rv := objc.Call[SampleBufferDisplayLayer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSampleBufferDisplayLayer() SampleBufferDisplayLayer {
	return SampleBufferDisplayLayerClass.New()
}

func (s_ SampleBufferDisplayLayer) Init() SampleBufferDisplayLayer {
	rv := objc.Call[SampleBufferDisplayLayer](s_, objc.Sel("init"))
	return rv
}

func (s_ SampleBufferDisplayLayer) ModelLayer() SampleBufferDisplayLayer {
	rv := objc.Call[SampleBufferDisplayLayer](s_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func SampleBufferDisplayLayer_ModelLayer() SampleBufferDisplayLayer {
	instance := SampleBufferDisplayLayerClass.Alloc().ModelLayer()
	instance.Autorelease()
	return instance
}

func (s_ SampleBufferDisplayLayer) PresentationLayer() SampleBufferDisplayLayer {
	rv := objc.Call[SampleBufferDisplayLayer](s_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func SampleBufferDisplayLayer_PresentationLayer() SampleBufferDisplayLayer {
	instance := SampleBufferDisplayLayerClass.Alloc().PresentationLayer()
	instance.Autorelease()
	return instance
}

func (s_ SampleBufferDisplayLayer) InitWithLayer(layer objc.IObject) SampleBufferDisplayLayer {
	rv := objc.Call[SampleBufferDisplayLayer](s_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func NewSampleBufferDisplayLayerWithLayer(layer objc.IObject) SampleBufferDisplayLayer {
	instance := SampleBufferDisplayLayerClass.Alloc().InitWithLayer(layer)
	instance.Autorelease()
	return instance
}

func (sc _SampleBufferDisplayLayerClass) Layer() SampleBufferDisplayLayer {
	rv := objc.Call[SampleBufferDisplayLayer](sc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func SampleBufferDisplayLayer_Layer() SampleBufferDisplayLayer {
	return SampleBufferDisplayLayerClass.Layer()
}

// A Boolean value that indicates whether the layer protects against screen capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/3081651-preventscapture?language=objc
func (s_ SampleBufferDisplayLayer) PreventsCapture() bool {
	rv := objc.Call[bool](s_, objc.Sel("preventsCapture"))
	return rv
}

// A Boolean value that indicates whether the layer protects against screen capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/3081651-preventscapture?language=objc
func (s_ SampleBufferDisplayLayer) SetPreventsCapture(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setPreventsCapture:"), value)
}

// The layer's control timebase, which governs how timestamps are interpreted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/1390569-controltimebase?language=objc
func (s_ SampleBufferDisplayLayer) ControlTimebase() coremedia.TimebaseRef {
	rv := objc.Call[coremedia.TimebaseRef](s_, objc.Sel("controlTimebase"))
	return rv
}

// The layer's control timebase, which governs how timestamps are interpreted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/1390569-controltimebase?language=objc
func (s_ SampleBufferDisplayLayer) SetControlTimebase(value coremedia.TimebaseRef) {
	objc.Call[objc.Void](s_, objc.Sel("setControlTimebase:"), value)
}

// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/3726154-outputobscuredduetoinsufficiente?language=objc
func (s_ SampleBufferDisplayLayer) OutputObscuredDueToInsufficientExternalProtection() bool {
	rv := objc.Call[bool](s_, objc.Sel("outputObscuredDueToInsufficientExternalProtection"))
	return rv
}

// A string defining how the video is displayed within the bounds rect of a sample buffer display layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/1387625-videogravity?language=objc
func (s_ SampleBufferDisplayLayer) VideoGravity() LayerVideoGravity {
	rv := objc.Call[LayerVideoGravity](s_, objc.Sel("videoGravity"))
	return rv
}

// A string defining how the video is displayed within the bounds rect of a sample buffer display layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/1387625-videogravity?language=objc
func (s_ SampleBufferDisplayLayer) SetVideoGravity(value LayerVideoGravity) {
	objc.Call[objc.Void](s_, objc.Sel("setVideoGravity:"), value)
}

// A Boolean value that indicates whether the layer prevents the system from sleeping during video playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/3088800-preventsdisplaysleepduringvideop?language=objc
func (s_ SampleBufferDisplayLayer) PreventsDisplaySleepDuringVideoPlayback() bool {
	rv := objc.Call[bool](s_, objc.Sel("preventsDisplaySleepDuringVideoPlayback"))
	return rv
}

// A Boolean value that indicates whether the layer prevents the system from sleeping during video playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayer/3088800-preventsdisplaysleepduringvideop?language=objc
func (s_ SampleBufferDisplayLayer) SetPreventsDisplaySleepDuringVideoPlayback(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setPreventsDisplaySleepDuringVideoPlayback:"), value)
}
