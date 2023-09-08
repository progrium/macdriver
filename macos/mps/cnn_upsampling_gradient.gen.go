// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNUpsamplingGradient] class.
var CNNUpsamplingGradientClass = _CNNUpsamplingGradientClass{objc.GetClass("MPSCNNUpsamplingGradient")}

type _CNNUpsamplingGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNUpsamplingGradient] class.
type ICNNUpsamplingGradient interface {
	ICNNGradientKernel
	ScaleFactorX() float64
	ScaleFactorY() float64
}

// A gradient filter that upsamples an existing Metal Performance Shaders image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplinggradient?language=objc
type CNNUpsamplingGradient struct {
	CNNGradientKernel
}

func CNNUpsamplingGradientFrom(ptr unsafe.Pointer) CNNUpsamplingGradient {
	return CNNUpsamplingGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (cc _CNNUpsamplingGradientClass) Alloc() CNNUpsamplingGradient {
	rv := objc.Call[CNNUpsamplingGradient](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNUpsamplingGradientClass) New() CNNUpsamplingGradient {
	rv := objc.Call[CNNUpsamplingGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNUpsamplingGradient() CNNUpsamplingGradient {
	return CNNUpsamplingGradientClass.New()
}

func (c_ CNNUpsamplingGradient) Init() CNNUpsamplingGradient {
	rv := objc.Call[CNNUpsamplingGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNUpsamplingGradient) InitWithDevice(device metal.PDevice) CNNUpsamplingGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNUpsamplingGradientWithDevice(device metal.PDevice) CNNUpsamplingGradient {
	instance := CNNUpsamplingGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNUpsamplingGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsamplingGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNUpsamplingGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsamplingGradient {
	instance := CNNUpsamplingGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplinggradient/2942630-scalefactorx?language=objc
func (c_ CNNUpsamplingGradient) ScaleFactorX() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplinggradient/2942628-scalefactory?language=objc
func (c_ CNNUpsamplingGradient) ScaleFactorY() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorY"))
	return rv
}
