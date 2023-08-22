// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixSoftMaxGradient] class.
var MatrixSoftMaxGradientClass = _MatrixSoftMaxGradientClass{objc.GetClass("MPSMatrixSoftMaxGradient")}

type _MatrixSoftMaxGradientClass struct {
	objc.Class
}

// An interface definition for the [MatrixSoftMaxGradient] class.
type IMatrixSoftMaxGradient interface {
	IMatrixBinaryKernel
	EncodeToCommandBufferGradientMatrixForwardOutputMatrixResultMatrix(commandBuffer metal.PCommandBuffer, gradientMatrix IMatrix, forwardOutputMatrix IMatrix, resultMatrix IMatrix)
	EncodeToCommandBufferObjectGradientMatrixForwardOutputMatrixResultMatrix(commandBufferObject objc.IObject, gradientMatrix IMatrix, forwardOutputMatrix IMatrix, resultMatrix IMatrix)
	SourceRows() uint
	SetSourceRows(value uint)
	SourceColumns() uint
	SetSourceColumns(value uint)
}

// A gradient softmax kernel that operates on matrices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient?language=objc
type MatrixSoftMaxGradient struct {
	MatrixBinaryKernel
}

func MatrixSoftMaxGradientFrom(ptr unsafe.Pointer) MatrixSoftMaxGradient {
	return MatrixSoftMaxGradient{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}

func (m_ MatrixSoftMaxGradient) InitWithDevice(device metal.PDevice) MatrixSoftMaxGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSoftMaxGradient](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966654-initwithdevice?language=objc
func NewMatrixSoftMaxGradientWithDevice(device metal.PDevice) MatrixSoftMaxGradient {
	instance := MatrixSoftMaxGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixSoftMaxGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSoftMaxGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSoftMaxGradient](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966651-copywithzone?language=objc
func MatrixSoftMaxGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSoftMaxGradient {
	instance := MatrixSoftMaxGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixSoftMaxGradientClass) Alloc() MatrixSoftMaxGradient {
	rv := objc.Call[MatrixSoftMaxGradient](mc, objc.Sel("alloc"))
	return rv
}

func MatrixSoftMaxGradient_Alloc() MatrixSoftMaxGradient {
	return MatrixSoftMaxGradientClass.Alloc()
}

func (mc _MatrixSoftMaxGradientClass) New() MatrixSoftMaxGradient {
	rv := objc.Call[MatrixSoftMaxGradient](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixSoftMaxGradient() MatrixSoftMaxGradient {
	return MatrixSoftMaxGradientClass.New()
}

func (m_ MatrixSoftMaxGradient) Init() MatrixSoftMaxGradient {
	rv := objc.Call[MatrixSoftMaxGradient](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966652-encodetocommandbuffer?language=objc
func (m_ MatrixSoftMaxGradient) EncodeToCommandBufferGradientMatrixForwardOutputMatrixResultMatrix(commandBuffer metal.PCommandBuffer, gradientMatrix IMatrix, forwardOutputMatrix IMatrix, resultMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:gradientMatrix:forwardOutputMatrix:resultMatrix:"), po0, objc.Ptr(gradientMatrix), objc.Ptr(forwardOutputMatrix), objc.Ptr(resultMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966652-encodetocommandbuffer?language=objc
func (m_ MatrixSoftMaxGradient) EncodeToCommandBufferObjectGradientMatrixForwardOutputMatrixResultMatrix(commandBufferObject objc.IObject, gradientMatrix IMatrix, forwardOutputMatrix IMatrix, resultMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:gradientMatrix:forwardOutputMatrix:resultMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(gradientMatrix), objc.Ptr(forwardOutputMatrix), objc.Ptr(resultMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966656-sourcerows?language=objc
func (m_ MatrixSoftMaxGradient) SourceRows() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceRows"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966656-sourcerows?language=objc
func (m_ MatrixSoftMaxGradient) SetSourceRows(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceRows:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966655-sourcecolumns?language=objc
func (m_ MatrixSoftMaxGradient) SourceColumns() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceColumns"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966655-sourcecolumns?language=objc
func (m_ MatrixSoftMaxGradient) SetSourceColumns(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceColumns:"), value)
}
