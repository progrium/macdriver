// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNLocalCorrelation] class.
var NNLocalCorrelationClass = _NNLocalCorrelationClass{objc.GetClass("MPSNNLocalCorrelation")}

type _NNLocalCorrelationClass struct {
	objc.Class
}

// An interface definition for the [NNLocalCorrelation] class.
type INNLocalCorrelation interface {
	INNReduceBinary
	StrideInX() uint
	SetStrideInX(value uint)
	StrideInY() uint
	SetStrideInY(value uint)
	WindowInX() uint
	SetWindowInX(value uint)
	WindowInY() uint
	SetWindowInY(value uint)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation?language=objc
type NNLocalCorrelation struct {
	NNReduceBinary
}

func NNLocalCorrelationFrom(ptr unsafe.Pointer) NNLocalCorrelation {
	return NNLocalCorrelation{
		NNReduceBinary: NNReduceBinaryFrom(ptr),
	}
}

func (n_ NNLocalCorrelation) InitWithDevice(device metal.PDevice) NNLocalCorrelation {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNLocalCorrelation](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131875-initwithdevice?language=objc
func NewNNLocalCorrelationWithDevice(device metal.PDevice) NNLocalCorrelation {
	instance := NNLocalCorrelationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNLocalCorrelationClass) Alloc() NNLocalCorrelation {
	rv := objc.Call[NNLocalCorrelation](nc, objc.Sel("alloc"))
	return rv
}

func NNLocalCorrelation_Alloc() NNLocalCorrelation {
	return NNLocalCorrelationClass.Alloc()
}

func (nc _NNLocalCorrelationClass) New() NNLocalCorrelation {
	rv := objc.Call[NNLocalCorrelation](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNLocalCorrelation() NNLocalCorrelation {
	return NNLocalCorrelationClass.New()
}

func (n_ NNLocalCorrelation) Init() NNLocalCorrelation {
	rv := objc.Call[NNLocalCorrelation](n_, objc.Sel("init"))
	return rv
}

func (n_ NNLocalCorrelation) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNLocalCorrelation {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNLocalCorrelation](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNLocalCorrelation_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNLocalCorrelation {
	instance := NNLocalCorrelationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131877-strideinx?language=objc
func (n_ NNLocalCorrelation) StrideInX() uint {
	rv := objc.Call[uint](n_, objc.Sel("strideInX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131877-strideinx?language=objc
func (n_ NNLocalCorrelation) SetStrideInX(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setStrideInX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131878-strideiny?language=objc
func (n_ NNLocalCorrelation) StrideInY() uint {
	rv := objc.Call[uint](n_, objc.Sel("strideInY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131878-strideiny?language=objc
func (n_ NNLocalCorrelation) SetStrideInY(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setStrideInY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131879-windowinx?language=objc
func (n_ NNLocalCorrelation) WindowInX() uint {
	rv := objc.Call[uint](n_, objc.Sel("windowInX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131879-windowinx?language=objc
func (n_ NNLocalCorrelation) SetWindowInX(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setWindowInX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131880-windowiny?language=objc
func (n_ NNLocalCorrelation) WindowInY() uint {
	rv := objc.Call[uint](n_, objc.Sel("windowInY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlocalcorrelation/3131880-windowiny?language=objc
func (n_ NNLocalCorrelation) SetWindowInY(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setWindowInY:"), value)
}
