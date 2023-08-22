// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageThresholdTruncate] class.
var ImageThresholdTruncateClass = _ImageThresholdTruncateClass{objc.GetClass("MPSImageThresholdTruncate")}

type _ImageThresholdTruncateClass struct {
	objc.Class
}

// An interface definition for the [ImageThresholdTruncate] class.
type IImageThresholdTruncate interface {
	IUnaryImageKernel
	Transform() *float64
	ThresholdValue() float64
}

// A filter that clamps the return value to an upper specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate?language=objc
type ImageThresholdTruncate struct {
	UnaryImageKernel
}

func ImageThresholdTruncateFrom(ptr unsafe.Pointer) ImageThresholdTruncate {
	return ImageThresholdTruncate{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageThresholdTruncate) InitWithDeviceThresholdValueLinearGrayColorTransform(device metal.PDevice, thresholdValue float64, transform *float64) ImageThresholdTruncate {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdTruncate](i_, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), po0, thresholdValue, transform)
	return rv
}

// Initializes the kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618818-initwithdevice?language=objc
func NewImageThresholdTruncateWithDeviceThresholdValueLinearGrayColorTransform(device metal.PDevice, thresholdValue float64, transform *float64) ImageThresholdTruncate {
	instance := ImageThresholdTruncateClass.Alloc().InitWithDeviceThresholdValueLinearGrayColorTransform(device, thresholdValue, transform)
	instance.Autorelease()
	return instance
}

func (ic _ImageThresholdTruncateClass) Alloc() ImageThresholdTruncate {
	rv := objc.Call[ImageThresholdTruncate](ic, objc.Sel("alloc"))
	return rv
}

func ImageThresholdTruncate_Alloc() ImageThresholdTruncate {
	return ImageThresholdTruncateClass.Alloc()
}

func (ic _ImageThresholdTruncateClass) New() ImageThresholdTruncate {
	rv := objc.Call[ImageThresholdTruncate](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageThresholdTruncate() ImageThresholdTruncate {
	return ImageThresholdTruncateClass.New()
}

func (i_ ImageThresholdTruncate) Init() ImageThresholdTruncate {
	rv := objc.Call[ImageThresholdTruncate](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageThresholdTruncate) InitWithDevice(device metal.PDevice) ImageThresholdTruncate {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdTruncate](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageThresholdTruncateWithDevice(device metal.PDevice) ImageThresholdTruncate {
	instance := ImageThresholdTruncateClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageThresholdTruncate) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageThresholdTruncate {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdTruncate](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageThresholdTruncate_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageThresholdTruncate {
	instance := ImageThresholdTruncateClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The color transform used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618787-transform?language=objc
func (i_ ImageThresholdTruncate) Transform() *float64 {
	rv := objc.Call[*float64](i_, objc.Sel("transform"))
	return rv
}

// The threshold value used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618882-thresholdvalue?language=objc
func (i_ ImageThresholdTruncate) ThresholdValue() float64 {
	rv := objc.Call[float64](i_, objc.Sel("thresholdValue"))
	return rv
}
