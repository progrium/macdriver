// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableVideoCompositionLayerInstruction] class.
var MutableVideoCompositionLayerInstructionClass = _MutableVideoCompositionLayerInstructionClass{objc.GetClass("AVMutableVideoCompositionLayerInstruction")}

type _MutableVideoCompositionLayerInstructionClass struct {
	objc.Class
}

// An interface definition for the [MutableVideoCompositionLayerInstruction] class.
type IMutableVideoCompositionLayerInstruction interface {
	IVideoCompositionLayerInstruction
	SetCropRectangleRampFromStartCropRectangleToEndCropRectangleTimeRange(startCropRectangle coregraphics.Rect, endCropRectangle coregraphics.Rect, timeRange coremedia.TimeRange)
	SetOpacityRampFromStartOpacityToEndOpacityTimeRange(startOpacity float64, endOpacity float64, timeRange coremedia.TimeRange)
	SetCropRectangleAtTime(cropRectangle coregraphics.Rect, time coremedia.Time)
	SetOpacityAtTime(opacity float64, time coremedia.Time)
	SetTransformRampFromStartTransformToEndTransformTimeRange(startTransform coregraphics.AffineTransform, endTransform coregraphics.AffineTransform, timeRange coremedia.TimeRange)
	SetTransformAtTime(transform coregraphics.AffineTransform, time coremedia.Time)
	SetTrackID(value objc.IObject)
}

// An object used to modify the transform, cropping, and opacity ramps applied to a given track in a mutable composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositionlayerinstruction?language=objc
type MutableVideoCompositionLayerInstruction struct {
	VideoCompositionLayerInstruction
}

func MutableVideoCompositionLayerInstructionFrom(ptr unsafe.Pointer) MutableVideoCompositionLayerInstruction {
	return MutableVideoCompositionLayerInstruction{
		VideoCompositionLayerInstruction: VideoCompositionLayerInstructionFrom(ptr),
	}
}

func (mc _MutableVideoCompositionLayerInstructionClass) VideoCompositionLayerInstruction() MutableVideoCompositionLayerInstruction {
	rv := objc.Call[MutableVideoCompositionLayerInstruction](mc, objc.Sel("videoCompositionLayerInstruction"))
	return rv
}

// Returns a new mutable video composition layer instruction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositionlayerinstruction/1519717-videocompositionlayerinstruction?language=objc
func MutableVideoCompositionLayerInstruction_VideoCompositionLayerInstruction() MutableVideoCompositionLayerInstruction {
	return MutableVideoCompositionLayerInstructionClass.VideoCompositionLayerInstruction()
}

func (mc _MutableVideoCompositionLayerInstructionClass) Alloc() MutableVideoCompositionLayerInstruction {
	rv := objc.Call[MutableVideoCompositionLayerInstruction](mc, objc.Sel("alloc"))
	return rv
}

func MutableVideoCompositionLayerInstruction_Alloc() MutableVideoCompositionLayerInstruction {
	return MutableVideoCompositionLayerInstructionClass.Alloc()
}

func (mc _MutableVideoCompositionLayerInstructionClass) New() MutableVideoCompositionLayerInstruction {
	rv := objc.Call[MutableVideoCompositionLayerInstruction](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableVideoCompositionLayerInstruction() MutableVideoCompositionLayerInstruction {
	return MutableVideoCompositionLayerInstructionClass.New()
}

func (m_ MutableVideoCompositionLayerInstruction) Init() MutableVideoCompositionLayerInstruction {
	rv := objc.Call[MutableVideoCompositionLayerInstruction](m_, objc.Sel("init"))
	return rv
}

// Sets a crop rectangle ramp to apply during the specified time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositionlayerinstruction/1385677-setcroprectanglerampfromstartcro?language=objc
func (m_ MutableVideoCompositionLayerInstruction) SetCropRectangleRampFromStartCropRectangleToEndCropRectangleTimeRange(startCropRectangle coregraphics.Rect, endCropRectangle coregraphics.Rect, timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("setCropRectangleRampFromStartCropRectangle:toEndCropRectangle:timeRange:"), startCropRectangle, endCropRectangle, timeRange)
}

// Sets an opacity ramp to apply during a specified time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositionlayerinstruction/1387532-setopacityrampfromstartopacity?language=objc
func (m_ MutableVideoCompositionLayerInstruction) SetOpacityRampFromStartOpacityToEndOpacityTimeRange(startOpacity float64, endOpacity float64, timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("setOpacityRampFromStartOpacity:toEndOpacity:timeRange:"), startOpacity, endOpacity, timeRange)
}

// Sets the crop rectangle  value at a time within the time range of the instruction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositionlayerinstruction/1387402-setcroprectangle?language=objc
func (m_ MutableVideoCompositionLayerInstruction) SetCropRectangleAtTime(cropRectangle coregraphics.Rect, time coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("setCropRectangle:atTime:"), cropRectangle, time)
}

// Sets the opacity value at a specific time within the time range of the instruction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositionlayerinstruction/1390758-setopacity?language=objc
func (m_ MutableVideoCompositionLayerInstruction) SetOpacityAtTime(opacity float64, time coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("setOpacity:atTime:"), opacity, time)
}

// Sets a transform ramp to apply during a given time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositionlayerinstruction/1388192-settransformrampfromstarttransfo?language=objc
func (m_ MutableVideoCompositionLayerInstruction) SetTransformRampFromStartTransformToEndTransformTimeRange(startTransform coregraphics.AffineTransform, endTransform coregraphics.AffineTransform, timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("setTransformRampFromStartTransform:toEndTransform:timeRange:"), startTransform, endTransform, timeRange)
}

// Sets the transform value at a time within the time range of the instruction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositionlayerinstruction/1390899-settransform?language=objc
func (m_ MutableVideoCompositionLayerInstruction) SetTransformAtTime(transform coregraphics.AffineTransform, time coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("setTransform:atTime:"), transform, time)
}

// The track identifier of the source track to which the compositor applies the instruction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablevideocompositionlayerinstruction/1387222-trackid?language=objc
func (m_ MutableVideoCompositionLayerInstruction) SetTrackID(value objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setTrackID:"), objc.Ptr(value))
}
