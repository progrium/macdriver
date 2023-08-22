// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixUnaryKernel] class.
var MatrixUnaryKernelClass = _MatrixUnaryKernelClass{objc.GetClass("MPSMatrixUnaryKernel")}

type _MatrixUnaryKernelClass struct {
	objc.Class
}

// An interface definition for the [MatrixUnaryKernel] class.
type IMatrixUnaryKernel interface {
	IKernel
	ResultMatrixOrigin() metal.Origin
	SetResultMatrixOrigin(value metal.Origin)
	BatchStart() uint
	SetBatchStart(value uint)
	BatchSize() uint
	SetBatchSize(value uint)
	SourceMatrixOrigin() metal.Origin
	SetSourceMatrixOrigin(value metal.Origin)
}

// A kernel that consumes one matrix and produces one matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel?language=objc
type MatrixUnaryKernel struct {
	Kernel
}

func MatrixUnaryKernelFrom(ptr unsafe.Pointer) MatrixUnaryKernel {
	return MatrixUnaryKernel{
		Kernel: KernelFrom(ptr),
	}
}

func (mc _MatrixUnaryKernelClass) Alloc() MatrixUnaryKernel {
	rv := objc.Call[MatrixUnaryKernel](mc, objc.Sel("alloc"))
	return rv
}

func MatrixUnaryKernel_Alloc() MatrixUnaryKernel {
	return MatrixUnaryKernelClass.Alloc()
}

func (mc _MatrixUnaryKernelClass) New() MatrixUnaryKernel {
	rv := objc.Call[MatrixUnaryKernel](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixUnaryKernel() MatrixUnaryKernel {
	return MatrixUnaryKernelClass.New()
}

func (m_ MatrixUnaryKernel) Init() MatrixUnaryKernel {
	rv := objc.Call[MatrixUnaryKernel](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixUnaryKernel) InitWithDevice(device metal.PDevice) MatrixUnaryKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixUnaryKernel](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixUnaryKernelWithDevice(device metal.PDevice) MatrixUnaryKernel {
	instance := MatrixUnaryKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixUnaryKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixUnaryKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixUnaryKernel](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixUnaryKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixUnaryKernel {
	instance := MatrixUnaryKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867150-resultmatrixorigin?language=objc
func (m_ MatrixUnaryKernel) ResultMatrixOrigin() metal.Origin {
	rv := objc.Call[metal.Origin](m_, objc.Sel("resultMatrixOrigin"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867150-resultmatrixorigin?language=objc
func (m_ MatrixUnaryKernel) SetResultMatrixOrigin(value metal.Origin) {
	objc.Call[objc.Void](m_, objc.Sel("setResultMatrixOrigin:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2866990-batchstart?language=objc
func (m_ MatrixUnaryKernel) BatchStart() uint {
	rv := objc.Call[uint](m_, objc.Sel("batchStart"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2866990-batchstart?language=objc
func (m_ MatrixUnaryKernel) SetBatchStart(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setBatchStart:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867118-batchsize?language=objc
func (m_ MatrixUnaryKernel) BatchSize() uint {
	rv := objc.Call[uint](m_, objc.Sel("batchSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867118-batchsize?language=objc
func (m_ MatrixUnaryKernel) SetBatchSize(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setBatchSize:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867053-sourcematrixorigin?language=objc
func (m_ MatrixUnaryKernel) SourceMatrixOrigin() metal.Origin {
	rv := objc.Call[metal.Origin](m_, objc.Sel("sourceMatrixOrigin"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixunarykernel/2867053-sourcematrixorigin?language=objc
func (m_ MatrixUnaryKernel) SetSourceMatrixOrigin(value metal.Origin) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceMatrixOrigin:"), value)
}
