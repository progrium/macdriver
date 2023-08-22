// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AsynchronousVideoCompositionRequest] class.
var AsynchronousVideoCompositionRequestClass = _AsynchronousVideoCompositionRequestClass{objc.GetClass("AVAsynchronousVideoCompositionRequest")}

type _AsynchronousVideoCompositionRequestClass struct {
	objc.Class
}

// An interface definition for the [AsynchronousVideoCompositionRequest] class.
type IAsynchronousVideoCompositionRequest interface {
	objc.IObject
	SourceFrameByTrackID(trackID objc.IObject) corevideo.PixelBufferRef
	FinishWithComposedVideoFrame(composedVideoFrame corevideo.PixelBufferRef)
	SourceTimedMetadataByTrackID(trackID objc.IObject) TimedMetadataGroup
	SourceSampleBufferByTrackID(trackID objc.IObject) coremedia.SampleBufferRef
	FinishWithError(error foundation.IError)
	FinishCancelledRequest()
	VideoCompositionInstruction() objc.Object
	SourceSampleDataTrackIDs() []foundation.Number
	SourceTrackIDs() []foundation.Number
	CompositionTime() coremedia.Time
	RenderContext() VideoCompositionRenderContext
}

// An object that contains information a video compositor needs to render an output pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest?language=objc
type AsynchronousVideoCompositionRequest struct {
	objc.Object
}

func AsynchronousVideoCompositionRequestFrom(ptr unsafe.Pointer) AsynchronousVideoCompositionRequest {
	return AsynchronousVideoCompositionRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AsynchronousVideoCompositionRequestClass) Alloc() AsynchronousVideoCompositionRequest {
	rv := objc.Call[AsynchronousVideoCompositionRequest](ac, objc.Sel("alloc"))
	return rv
}

func AsynchronousVideoCompositionRequest_Alloc() AsynchronousVideoCompositionRequest {
	return AsynchronousVideoCompositionRequestClass.Alloc()
}

func (ac _AsynchronousVideoCompositionRequestClass) New() AsynchronousVideoCompositionRequest {
	rv := objc.Call[AsynchronousVideoCompositionRequest](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAsynchronousVideoCompositionRequest() AsynchronousVideoCompositionRequest {
	return AsynchronousVideoCompositionRequestClass.New()
}

func (a_ AsynchronousVideoCompositionRequest) Init() AsynchronousVideoCompositionRequest {
	rv := objc.Call[AsynchronousVideoCompositionRequest](a_, objc.Sel("init"))
	return rv
}

// Returns a source pixel buffer for the track that contains the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest/1390379-sourceframebytrackid?language=objc
func (a_ AsynchronousVideoCompositionRequest) SourceFrameByTrackID(trackID objc.IObject) corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](a_, objc.Sel("sourceFrameByTrackID:"), objc.Ptr(trackID))
	return rv
}

// Finishes the request to compose the frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest/1387450-finishwithcomposedvideoframe?language=objc
func (a_ AsynchronousVideoCompositionRequest) FinishWithComposedVideoFrame(composedVideoFrame corevideo.PixelBufferRef) {
	objc.Call[objc.Void](a_, objc.Sel("finishWithComposedVideoFrame:"), composedVideoFrame)
}

// Returns a source timed metadata group for the track that contains the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest/3750313-sourcetimedmetadatabytrackid?language=objc
func (a_ AsynchronousVideoCompositionRequest) SourceTimedMetadataByTrackID(trackID objc.IObject) TimedMetadataGroup {
	rv := objc.Call[TimedMetadataGroup](a_, objc.Sel("sourceTimedMetadataByTrackID:"), objc.Ptr(trackID))
	return rv
}

// Returns a source sample buffer for the track that contains the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest/3750311-sourcesamplebufferbytrackid?language=objc
func (a_ AsynchronousVideoCompositionRequest) SourceSampleBufferByTrackID(trackID objc.IObject) coremedia.SampleBufferRef {
	rv := objc.Call[coremedia.SampleBufferRef](a_, objc.Sel("sourceSampleBufferByTrackID:"), objc.Ptr(trackID))
	return rv
}

// Finishes the request with an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest/1390797-finishwitherror?language=objc
func (a_ AsynchronousVideoCompositionRequest) FinishWithError(error foundation.IError) {
	objc.Call[objc.Void](a_, objc.Sel("finishWithError:"), objc.Ptr(error))
}

// Cancels the request to compose a video frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest/1386261-finishcancelledrequest?language=objc
func (a_ AsynchronousVideoCompositionRequest) FinishCancelledRequest() {
	objc.Call[objc.Void](a_, objc.Sel("finishCancelledRequest"))
}

// A video composition instruction that indicates how to compose the frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest/1386672-videocompositioninstruction?language=objc
func (a_ AsynchronousVideoCompositionRequest) VideoCompositionInstruction() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("videoCompositionInstruction"))
	return rv
}

// The identifiers of tracks that contain source metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest/3750312-sourcesampledatatrackids?language=objc
func (a_ AsynchronousVideoCompositionRequest) SourceSampleDataTrackIDs() []foundation.Number {
	rv := objc.Call[[]foundation.Number](a_, objc.Sel("sourceSampleDataTrackIDs"))
	return rv
}

// The identifiers of tracks that contain source video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest/1388898-sourcetrackids?language=objc
func (a_ AsynchronousVideoCompositionRequest) SourceTrackIDs() []foundation.Number {
	rv := objc.Call[[]foundation.Number](a_, objc.Sel("sourceTrackIDs"))
	return rv
}

// A time for which to compose the frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest/1386888-compositiontime?language=objc
func (a_ AsynchronousVideoCompositionRequest) CompositionTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("compositionTime"))
	return rv
}

// The rendering context of the video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasynchronousvideocompositionrequest/1389112-rendercontext?language=objc
func (a_ AsynchronousVideoCompositionRequest) RenderContext() VideoCompositionRenderContext {
	rv := objc.Call[VideoCompositionRenderContext](a_, objc.Sel("renderContext"))
	return rv
}
