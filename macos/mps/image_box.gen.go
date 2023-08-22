// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageBox] class.
var ImageBoxClass = _ImageBoxClass{objc.GetClass("MPSImageBox")}

type _ImageBoxClass struct {
	objc.Class
}

// An interface definition for the [ImageBox] class.
type IImageBox interface {
	IUnaryImageKernel
	KernelHeight() uint
	KernelWidth() uint
}

// A filter that convolves an image with a given kernel of odd width and height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebox?language=objc
type ImageBox struct {
	UnaryImageKernel
}

func ImageBoxFrom(ptr unsafe.Pointer) ImageBox {
	return ImageBox{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageBox) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) ImageBox {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageBox](i_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

// Initializes a box filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebox/1618789-initwithdevice?language=objc
func NewImageBoxWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) ImageBox {
	instance := ImageBoxClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (ic _ImageBoxClass) Alloc() ImageBox {
	rv := objc.Call[ImageBox](ic, objc.Sel("alloc"))
	return rv
}

func ImageBox_Alloc() ImageBox {
	return ImageBoxClass.Alloc()
}

func (ic _ImageBoxClass) New() ImageBox {
	rv := objc.Call[ImageBox](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageBox() ImageBox {
	return ImageBoxClass.New()
}

func (i_ ImageBox) Init() ImageBox {
	rv := objc.Call[ImageBox](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageBox) InitWithDevice(device metal.PDevice) ImageBox {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageBox](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageBoxWithDevice(device metal.PDevice) ImageBox {
	instance := ImageBoxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageBox) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageBox {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageBox](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageBox_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageBox {
	instance := ImageBoxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The height of the filter window. Must be an odd number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebox/1618739-kernelheight?language=objc
func (i_ ImageBox) KernelHeight() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelHeight"))
	return rv
}

// The width of the filter window. Must be an odd number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebox/1618834-kernelwidth?language=objc
func (i_ ImageBox) KernelWidth() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelWidth"))
	return rv
}
