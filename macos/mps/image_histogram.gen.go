// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/kernel"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageHistogram] class.
var ImageHistogramClass = _ImageHistogramClass{objc.GetClass("MPSImageHistogram")}

type _ImageHistogramClass struct {
	objc.Class
}

// An interface definition for the [ImageHistogram] class.
type IImageHistogram interface {
	IKernel
	HistogramSizeForSourceFormat(sourceFormat metal.PixelFormat) uint
	EncodeToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, histogram metal.PBuffer, histogramOffset uint)
	EncodeToCommandBufferObjectSourceTextureObjectHistogramObjectHistogramOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, histogramObject objc.IObject, histogramOffset uint)
	HistogramInfo() ImageHistogramInfo
	MinPixelThresholdValue() kernel.Vector_float4
	SetMinPixelThresholdValue(value kernel.Vector_float4)
	ZeroHistogram() bool
	SetZeroHistogram(value bool)
	ClipRectSource() metal.Region
	SetClipRectSource(value metal.Region)
}

// A filter that computes the histogram of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram?language=objc
type ImageHistogram struct {
	Kernel
}

func ImageHistogramFrom(ptr unsafe.Pointer) ImageHistogram {
	return ImageHistogram{
		Kernel: KernelFrom(ptr),
	}
}

func (i_ ImageHistogram) InitWithDeviceHistogramInfo(device metal.PDevice, histogramInfo *ImageHistogramInfo) ImageHistogram {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageHistogram](i_, objc.Sel("initWithDevice:histogramInfo:"), po0, histogramInfo)
	return rv
}

// Initializes a histogram with specific information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618910-initwithdevice?language=objc
func NewImageHistogramWithDeviceHistogramInfo(device metal.PDevice, histogramInfo *ImageHistogramInfo) ImageHistogram {
	instance := ImageHistogramClass.Alloc().InitWithDeviceHistogramInfo(device, histogramInfo)
	instance.Autorelease()
	return instance
}

func (ic _ImageHistogramClass) Alloc() ImageHistogram {
	rv := objc.Call[ImageHistogram](ic, objc.Sel("alloc"))
	return rv
}

func ImageHistogram_Alloc() ImageHistogram {
	return ImageHistogramClass.Alloc()
}

func (ic _ImageHistogramClass) New() ImageHistogram {
	rv := objc.Call[ImageHistogram](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageHistogram() ImageHistogram {
	return ImageHistogramClass.New()
}

func (i_ ImageHistogram) Init() ImageHistogram {
	rv := objc.Call[ImageHistogram](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageHistogram) InitWithDevice(device metal.PDevice) ImageHistogram {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageHistogram](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewImageHistogramWithDevice(device metal.PDevice) ImageHistogram {
	instance := ImageHistogramClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageHistogram) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageHistogram {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageHistogram](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageHistogram_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageHistogram {
	instance := ImageHistogramClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The amount of space the histogram will take up in the output buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618839-histogramsizeforsourceformat?language=objc
func (i_ ImageHistogram) HistogramSizeForSourceFormat(sourceFormat metal.PixelFormat) uint {
	rv := objc.Call[uint](i_, objc.Sel("histogramSizeForSourceFormat:"), sourceFormat)
	return rv
}

// Encodes the filter to a command buffer using a compute command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618853-encodetocommandbuffer?language=objc
func (i_ ImageHistogram) EncodeToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, histogram metal.PBuffer, histogramOffset uint) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", source)
	po2 := objc.WrapAsProtocol("MTLBuffer", histogram)
	objc.Call[objc.Void](i_, objc.Sel("encodeToCommandBuffer:sourceTexture:histogram:histogramOffset:"), po0, po1, po2, histogramOffset)
}

// Encodes the filter to a command buffer using a compute command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618853-encodetocommandbuffer?language=objc
func (i_ ImageHistogram) EncodeToCommandBufferObjectSourceTextureObjectHistogramObjectHistogramOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, histogramObject objc.IObject, histogramOffset uint) {
	objc.Call[objc.Void](i_, objc.Sel("encodeToCommandBuffer:sourceTexture:histogram:histogramOffset:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceObject), objc.Ptr(histogramObject), histogramOffset)
}

// A structure describing the histogram content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618844-histograminfo?language=objc
func (i_ ImageHistogram) HistogramInfo() ImageHistogramInfo {
	rv := objc.Call[ImageHistogramInfo](i_, objc.Sel("histogramInfo"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/2867008-minpixelthresholdvalue?language=objc
func (i_ ImageHistogram) MinPixelThresholdValue() kernel.Vector_float4 {
	rv := objc.Call[kernel.Vector_float4](i_, objc.Sel("minPixelThresholdValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/2867008-minpixelthresholdvalue?language=objc
func (i_ ImageHistogram) SetMinPixelThresholdValue(value kernel.Vector_float4) {
	objc.Call[objc.Void](i_, objc.Sel("setMinPixelThresholdValue:"), value)
}

// Determines whether to zero-initialize the histogram results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618891-zerohistogram?language=objc
func (i_ ImageHistogram) ZeroHistogram() bool {
	rv := objc.Call[bool](i_, objc.Sel("zeroHistogram"))
	return rv
}

// Determines whether to zero-initialize the histogram results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618891-zerohistogram?language=objc
func (i_ ImageHistogram) SetZeroHistogram(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setZeroHistogram:"), value)
}

// The source rectangle to use when reading data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618765-cliprectsource?language=objc
func (i_ ImageHistogram) ClipRectSource() metal.Region {
	rv := objc.Call[metal.Region](i_, objc.Sel("clipRectSource"))
	return rv
}

// The source rectangle to use when reading data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618765-cliprectsource?language=objc
func (i_ ImageHistogram) SetClipRectSource(value metal.Region) {
	objc.Call[objc.Void](i_, objc.Sel("setClipRectSource:"), value)
}
