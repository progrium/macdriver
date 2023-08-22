// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageArithmetic] class.
var ImageArithmeticClass = _ImageArithmeticClass{objc.GetClass("MPSImageArithmetic")}

type _ImageArithmeticClass struct {
	objc.Class
}

// An interface definition for the [ImageArithmetic] class.
type IImageArithmetic interface {
	IBinaryImageKernel
	SecondaryScale() float64
	SetSecondaryScale(value float64)
	PrimaryStrideInPixels() metal.Size
	SetPrimaryStrideInPixels(value metal.Size)
	SecondaryStrideInPixels() metal.Size
	SetSecondaryStrideInPixels(value metal.Size)
	MaximumValue() float64
	SetMaximumValue(value float64)
	PrimaryScale() float64
	SetPrimaryScale(value float64)
	MinimumValue() float64
	SetMinimumValue(value float64)
	Bias() float64
	SetBias(value float64)
}

// Base class for basic arithmetic nodes [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic?language=objc
type ImageArithmetic struct {
	BinaryImageKernel
}

func ImageArithmeticFrom(ptr unsafe.Pointer) ImageArithmetic {
	return ImageArithmetic{
		BinaryImageKernel: BinaryImageKernelFrom(ptr),
	}
}

func (ic _ImageArithmeticClass) Alloc() ImageArithmetic {
	rv := objc.Call[ImageArithmetic](ic, objc.Sel("alloc"))
	return rv
}

func ImageArithmetic_Alloc() ImageArithmetic {
	return ImageArithmeticClass.Alloc()
}

func (ic _ImageArithmeticClass) New() ImageArithmetic {
	rv := objc.Call[ImageArithmetic](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageArithmetic() ImageArithmetic {
	return ImageArithmeticClass.New()
}

func (i_ ImageArithmetic) Init() ImageArithmetic {
	rv := objc.Call[ImageArithmetic](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageArithmetic) InitWithDevice(device metal.PDevice) ImageArithmetic {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageArithmetic](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsbinaryimagekernel/2866331-initwithdevice?language=objc
func NewImageArithmeticWithDevice(device metal.PDevice) ImageArithmetic {
	instance := ImageArithmeticClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageArithmetic) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageArithmetic {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageArithmetic](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageArithmetic_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageArithmetic {
	instance := ImageArithmeticClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866601-secondaryscale?language=objc
func (i_ ImageArithmetic) SecondaryScale() float64 {
	rv := objc.Call[float64](i_, objc.Sel("secondaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866601-secondaryscale?language=objc
func (i_ ImageArithmetic) SetSecondaryScale(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setSecondaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2889864-primarystrideinpixels?language=objc
func (i_ ImageArithmetic) PrimaryStrideInPixels() metal.Size {
	rv := objc.Call[metal.Size](i_, objc.Sel("primaryStrideInPixels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2889864-primarystrideinpixels?language=objc
func (i_ ImageArithmetic) SetPrimaryStrideInPixels(value metal.Size) {
	objc.Call[objc.Void](i_, objc.Sel("setPrimaryStrideInPixels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2889865-secondarystrideinpixels?language=objc
func (i_ ImageArithmetic) SecondaryStrideInPixels() metal.Size {
	rv := objc.Call[metal.Size](i_, objc.Sel("secondaryStrideInPixels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2889865-secondarystrideinpixels?language=objc
func (i_ ImageArithmetic) SetSecondaryStrideInPixels(value metal.Size) {
	objc.Call[objc.Void](i_, objc.Sel("setSecondaryStrideInPixels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2942356-maximumvalue?language=objc
func (i_ ImageArithmetic) MaximumValue() float64 {
	rv := objc.Call[float64](i_, objc.Sel("maximumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2942356-maximumvalue?language=objc
func (i_ ImageArithmetic) SetMaximumValue(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setMaximumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866602-primaryscale?language=objc
func (i_ ImageArithmetic) PrimaryScale() float64 {
	rv := objc.Call[float64](i_, objc.Sel("primaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866602-primaryscale?language=objc
func (i_ ImageArithmetic) SetPrimaryScale(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setPrimaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2942357-minimumvalue?language=objc
func (i_ ImageArithmetic) MinimumValue() float64 {
	rv := objc.Call[float64](i_, objc.Sel("minimumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2942357-minimumvalue?language=objc
func (i_ ImageArithmetic) SetMinimumValue(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setMinimumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866609-bias?language=objc
func (i_ ImageArithmetic) Bias() float64 {
	rv := objc.Call[float64](i_, objc.Sel("bias"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagearithmetic/2866609-bias?language=objc
func (i_ ImageArithmetic) SetBias(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setBias:"), value)
}
