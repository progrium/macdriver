// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageFindKeypoints] class.
var ImageFindKeypointsClass = _ImageFindKeypointsClass{objc.GetClass("MPSImageFindKeypoints")}

type _ImageFindKeypointsClass struct {
	objc.Class
}

// An interface definition for the [ImageFindKeypoints] class.
type IImageFindKeypoints interface {
	IKernel
	EncodeToCommandBufferSourceTextureRegionsNumberOfRegionsKeypointCountBufferKeypointCountBufferOffsetKeypointDataBufferKeypointDataBufferOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, regions *metal.Region, numberOfRegions uint, keypointCountBuffer metal.PBuffer, keypointCountBufferOffset uint, keypointDataBuffer metal.PBuffer, keypointDataBufferOffset uint)
	EncodeToCommandBufferObjectSourceTextureObjectRegionsNumberOfRegionsKeypointCountBufferObjectKeypointCountBufferOffsetKeypointDataBufferObjectKeypointDataBufferOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, regions *metal.Region, numberOfRegions uint, keypointCountBufferObject objc.IObject, keypointCountBufferOffset uint, keypointDataBufferObject objc.IObject, keypointDataBufferOffset uint)
	KeypointRangeInfo() ImageKeypointRangeInfo
}

// A kernel that is used to find a list of keypoints. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints?language=objc
type ImageFindKeypoints struct {
	Kernel
}

func ImageFindKeypointsFrom(ptr unsafe.Pointer) ImageFindKeypoints {
	return ImageFindKeypoints{
		Kernel: KernelFrom(ptr),
	}
}

func (i_ ImageFindKeypoints) InitWithDeviceInfo(device metal.PDevice, info *ImageKeypointRangeInfo) ImageFindKeypoints {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageFindKeypoints](i_, objc.Sel("initWithDevice:info:"), po0, info)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873309-initwithdevice?language=objc
func NewImageFindKeypointsWithDeviceInfo(device metal.PDevice, info *ImageKeypointRangeInfo) ImageFindKeypoints {
	instance := ImageFindKeypointsClass.Alloc().InitWithDeviceInfo(device, info)
	instance.Autorelease()
	return instance
}

func (ic _ImageFindKeypointsClass) Alloc() ImageFindKeypoints {
	rv := objc.Call[ImageFindKeypoints](ic, objc.Sel("alloc"))
	return rv
}

func (ic _ImageFindKeypointsClass) New() ImageFindKeypoints {
	rv := objc.Call[ImageFindKeypoints](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageFindKeypoints() ImageFindKeypoints {
	return ImageFindKeypointsClass.New()
}

func (i_ ImageFindKeypoints) Init() ImageFindKeypoints {
	rv := objc.Call[ImageFindKeypoints](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageFindKeypoints) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageFindKeypoints {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageFindKeypoints](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageFindKeypoints_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageFindKeypoints {
	instance := ImageFindKeypointsClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (i_ ImageFindKeypoints) InitWithDevice(device metal.PDevice) ImageFindKeypoints {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageFindKeypoints](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewImageFindKeypointsWithDevice(device metal.PDevice) ImageFindKeypoints {
	instance := ImageFindKeypointsClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873307-encodetocommandbuffer?language=objc
func (i_ ImageFindKeypoints) EncodeToCommandBufferSourceTextureRegionsNumberOfRegionsKeypointCountBufferKeypointCountBufferOffsetKeypointDataBufferKeypointDataBufferOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, regions *metal.Region, numberOfRegions uint, keypointCountBuffer metal.PBuffer, keypointCountBufferOffset uint, keypointDataBuffer metal.PBuffer, keypointDataBufferOffset uint) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", source)
	po4 := objc.WrapAsProtocol("MTLBuffer", keypointCountBuffer)
	po6 := objc.WrapAsProtocol("MTLBuffer", keypointDataBuffer)
	objc.Call[objc.Void](i_, objc.Sel("encodeToCommandBuffer:sourceTexture:regions:numberOfRegions:keypointCountBuffer:keypointCountBufferOffset:keypointDataBuffer:keypointDataBufferOffset:"), po0, po1, regions, numberOfRegions, po4, keypointCountBufferOffset, po6, keypointDataBufferOffset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873307-encodetocommandbuffer?language=objc
func (i_ ImageFindKeypoints) EncodeToCommandBufferObjectSourceTextureObjectRegionsNumberOfRegionsKeypointCountBufferObjectKeypointCountBufferOffsetKeypointDataBufferObjectKeypointDataBufferOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, regions *metal.Region, numberOfRegions uint, keypointCountBufferObject objc.IObject, keypointCountBufferOffset uint, keypointDataBufferObject objc.IObject, keypointDataBufferOffset uint) {
	objc.Call[objc.Void](i_, objc.Sel("encodeToCommandBuffer:sourceTexture:regions:numberOfRegions:keypointCountBuffer:keypointCountBufferOffset:keypointDataBuffer:keypointDataBufferOffset:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceObject), regions, numberOfRegions, objc.Ptr(keypointCountBufferObject), keypointCountBufferOffset, objc.Ptr(keypointDataBufferObject), keypointDataBufferOffset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873310-keypointrangeinfo?language=objc
func (i_ ImageFindKeypoints) KeypointRangeInfo() ImageKeypointRangeInfo {
	rv := objc.Call[ImageKeypointRangeInfo](i_, objc.Sel("keypointRangeInfo"))
	return rv
}
