// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixVectorMultiplication] class.
var MatrixVectorMultiplicationClass = _MatrixVectorMultiplicationClass{objc.GetClass("MPSMatrixVectorMultiplication")}

type _MatrixVectorMultiplicationClass struct {
	objc.Class
}

// An interface definition for the [MatrixVectorMultiplication] class.
type IMatrixVectorMultiplication interface {
	IMatrixBinaryKernel
	EncodeToCommandBufferInputMatrixInputVectorResultVector(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, inputVector IVector, resultVector IVector)
	EncodeToCommandBufferObjectInputMatrixInputVectorResultVector(commandBufferObject objc.IObject, inputMatrix IMatrix, inputVector IVector, resultVector IVector)
}

// A matrix-vector multiplication kernel [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication?language=objc
type MatrixVectorMultiplication struct {
	MatrixBinaryKernel
}

func MatrixVectorMultiplicationFrom(ptr unsafe.Pointer) MatrixVectorMultiplication {
	return MatrixVectorMultiplication{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}

func (m_ MatrixVectorMultiplication) InitWithDeviceRowsColumns(device metal.PDevice, rows uint, columns uint) MatrixVectorMultiplication {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixVectorMultiplication](m_, objc.Sel("initWithDevice:rows:columns:"), po0, rows, columns)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication/2909035-initwithdevice?language=objc
func NewMatrixVectorMultiplicationWithDeviceRowsColumns(device metal.PDevice, rows uint, columns uint) MatrixVectorMultiplication {
	instance := MatrixVectorMultiplicationClass.Alloc().InitWithDeviceRowsColumns(device, rows, columns)
	instance.Autorelease()
	return instance
}

func (mc _MatrixVectorMultiplicationClass) Alloc() MatrixVectorMultiplication {
	rv := objc.Call[MatrixVectorMultiplication](mc, objc.Sel("alloc"))
	return rv
}

func MatrixVectorMultiplication_Alloc() MatrixVectorMultiplication {
	return MatrixVectorMultiplicationClass.Alloc()
}

func (mc _MatrixVectorMultiplicationClass) New() MatrixVectorMultiplication {
	rv := objc.Call[MatrixVectorMultiplication](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixVectorMultiplication() MatrixVectorMultiplication {
	return MatrixVectorMultiplicationClass.New()
}

func (m_ MatrixVectorMultiplication) Init() MatrixVectorMultiplication {
	rv := objc.Call[MatrixVectorMultiplication](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixVectorMultiplication) InitWithDevice(device metal.PDevice) MatrixVectorMultiplication {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixVectorMultiplication](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixVectorMultiplicationWithDevice(device metal.PDevice) MatrixVectorMultiplication {
	instance := MatrixVectorMultiplicationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixVectorMultiplication) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixVectorMultiplication {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixVectorMultiplication](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixVectorMultiplication_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixVectorMultiplication {
	instance := MatrixVectorMultiplicationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication/2873084-encodetocommandbuffer?language=objc
func (m_ MatrixVectorMultiplication) EncodeToCommandBufferInputMatrixInputVectorResultVector(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, inputVector IVector, resultVector IVector) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:inputVector:resultVector:"), po0, objc.Ptr(inputMatrix), objc.Ptr(inputVector), objc.Ptr(resultVector))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixvectormultiplication/2873084-encodetocommandbuffer?language=objc
func (m_ MatrixVectorMultiplication) EncodeToCommandBufferObjectInputMatrixInputVectorResultVector(commandBufferObject objc.IObject, inputMatrix IMatrix, inputVector IVector, resultVector IVector) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:inputVector:resultVector:"), objc.Ptr(commandBufferObject), objc.Ptr(inputMatrix), objc.Ptr(inputVector), objc.Ptr(resultVector))
}
