// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNDilatedPoolingMaxGradient] class.
var CNNDilatedPoolingMaxGradientClass = _CNNDilatedPoolingMaxGradientClass{objc.GetClass("MPSCNNDilatedPoolingMaxGradient")}

type _CNNDilatedPoolingMaxGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNDilatedPoolingMaxGradient] class.
type ICNNDilatedPoolingMaxGradient interface {
	ICNNPoolingGradient
}

// A gradient dilated max pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradient?language=objc
type CNNDilatedPoolingMaxGradient struct {
	CNNPoolingGradient
}

func CNNDilatedPoolingMaxGradientFrom(ptr unsafe.Pointer) CNNDilatedPoolingMaxGradient {
	return CNNDilatedPoolingMaxGradient{
		CNNPoolingGradient: CNNPoolingGradientFrom(ptr),
	}
}

func (c_ CNNDilatedPoolingMaxGradient) InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) CNNDilatedPoolingMaxGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDilatedPoolingMaxGradient](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:"), po0, kernelWidth, kernelHeight, dilationRateX, dilationRateY, strideInPixelsX, strideInPixelsY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradient/2942349-initwithdevice?language=objc
func NewCNNDilatedPoolingMaxGradientWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) CNNDilatedPoolingMaxGradient {
	instance := CNNDilatedPoolingMaxGradientClass.Alloc().InitWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device, kernelWidth, kernelHeight, dilationRateX, dilationRateY, strideInPixelsX, strideInPixelsY)
	instance.Autorelease()
	return instance
}

func (cc _CNNDilatedPoolingMaxGradientClass) Alloc() CNNDilatedPoolingMaxGradient {
	rv := objc.Call[CNNDilatedPoolingMaxGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNDilatedPoolingMaxGradient_Alloc() CNNDilatedPoolingMaxGradient {
	return CNNDilatedPoolingMaxGradientClass.Alloc()
}

func (cc _CNNDilatedPoolingMaxGradientClass) New() CNNDilatedPoolingMaxGradient {
	rv := objc.Call[CNNDilatedPoolingMaxGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDilatedPoolingMaxGradient() CNNDilatedPoolingMaxGradient {
	return CNNDilatedPoolingMaxGradientClass.New()
}

func (c_ CNNDilatedPoolingMaxGradient) Init() CNNDilatedPoolingMaxGradient {
	rv := objc.Call[CNNDilatedPoolingMaxGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNDilatedPoolingMaxGradient) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNDilatedPoolingMaxGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDilatedPoolingMaxGradient](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942337-initwithdevice?language=objc
func NewCNNDilatedPoolingMaxGradientWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNDilatedPoolingMaxGradient {
	instance := CNNDilatedPoolingMaxGradientClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (c_ CNNDilatedPoolingMaxGradient) InitWithDevice(device metal.PDevice) CNNDilatedPoolingMaxGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDilatedPoolingMaxGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNDilatedPoolingMaxGradientWithDevice(device metal.PDevice) CNNDilatedPoolingMaxGradient {
	instance := CNNDilatedPoolingMaxGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNDilatedPoolingMaxGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNDilatedPoolingMaxGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDilatedPoolingMaxGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNDilatedPoolingMaxGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNDilatedPoolingMaxGradient {
	instance := CNNDilatedPoolingMaxGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
