// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNGramMatrixCalculation] class.
var NNGramMatrixCalculationClass = _NNGramMatrixCalculationClass{objc.GetClass("MPSNNGramMatrixCalculation")}

type _NNGramMatrixCalculationClass struct {
	objc.Class
}

// An interface definition for the [NNGramMatrixCalculation] class.
type INNGramMatrixCalculation interface {
	ICNNKernel
	Alpha() float64
	SetAlpha(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculation?language=objc
type NNGramMatrixCalculation struct {
	CNNKernel
}

func NNGramMatrixCalculationFrom(ptr unsafe.Pointer) NNGramMatrixCalculation {
	return NNGramMatrixCalculation{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (n_ NNGramMatrixCalculation) InitWithDevice(device metal.PDevice) NNGramMatrixCalculation {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNGramMatrixCalculation](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculation/3114079-initwithdevice?language=objc
func NewNNGramMatrixCalculationWithDevice(device metal.PDevice) NNGramMatrixCalculation {
	instance := NNGramMatrixCalculationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNGramMatrixCalculation) InitWithDeviceAlpha(device metal.PDevice, alpha float64) NNGramMatrixCalculation {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNGramMatrixCalculation](n_, objc.Sel("initWithDevice:alpha:"), po0, alpha)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculation/3114080-initwithdevice?language=objc
func NewNNGramMatrixCalculationWithDeviceAlpha(device metal.PDevice, alpha float64) NNGramMatrixCalculation {
	instance := NNGramMatrixCalculationClass.Alloc().InitWithDeviceAlpha(device, alpha)
	instance.Autorelease()
	return instance
}

func (nc _NNGramMatrixCalculationClass) Alloc() NNGramMatrixCalculation {
	rv := objc.Call[NNGramMatrixCalculation](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NNGramMatrixCalculationClass) New() NNGramMatrixCalculation {
	rv := objc.Call[NNGramMatrixCalculation](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNGramMatrixCalculation() NNGramMatrixCalculation {
	return NNGramMatrixCalculationClass.New()
}

func (n_ NNGramMatrixCalculation) Init() NNGramMatrixCalculation {
	rv := objc.Call[NNGramMatrixCalculation](n_, objc.Sel("init"))
	return rv
}

func (n_ NNGramMatrixCalculation) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNGramMatrixCalculation {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNGramMatrixCalculation](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNGramMatrixCalculation_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNGramMatrixCalculation {
	instance := NNGramMatrixCalculationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculation/3114077-alpha?language=objc
func (n_ NNGramMatrixCalculation) Alpha() float64 {
	rv := objc.Call[float64](n_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculation/3114077-alpha?language=objc
func (n_ NNGramMatrixCalculation) SetAlpha(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setAlpha:"), value)
}
