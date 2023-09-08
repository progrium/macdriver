// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayMultiaryKernel] class.
var NDArrayMultiaryKernelClass = _NDArrayMultiaryKernelClass{objc.GetClass("MPSNDArrayMultiaryKernel")}

type _NDArrayMultiaryKernelClass struct {
	objc.Class
}

// An interface definition for the [NDArrayMultiaryKernel] class.
type INDArrayMultiaryKernel interface {
	INDArrayMultiaryBase
	EncodeToCommandBufferSourceArraysDestinationArray(cmdBuf metal.PCommandBuffer, sourceArrays []INDArray, destination INDArray)
	EncodeToCommandBufferObjectSourceArraysDestinationArray(cmdBufObject objc.IObject, sourceArrays []INDArray, destination INDArray)
	EncodeToCommandBufferSourceArraysResultStateDestinationArray(cmdBuf metal.PCommandBuffer, sourceArrays []INDArray, outGradientState IState, destination INDArray)
	EncodeToCommandBufferObjectSourceArraysResultStateDestinationArray(cmdBufObject objc.IObject, sourceArrays []INDArray, outGradientState IState, destination INDArray)
	EncodeToCommandBufferSourceArrays(cmdBuf metal.PCommandBuffer, sourceArrays []INDArray) NDArray
	EncodeToCommandBufferObjectSourceArrays(cmdBufObject objc.IObject, sourceArrays []INDArray) NDArray
	EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary(cmdBuf metal.PCommandBuffer, sourceArrays []INDArray, outGradientState IState, outputStateIsTemporary bool) NDArray
	EncodeToCommandBufferObjectSourceArraysResultStateOutputStateIsTemporary(cmdBufObject objc.IObject, sourceArrays []INDArray, outGradientState IState, outputStateIsTemporary bool) NDArray
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel?language=objc
type NDArrayMultiaryKernel struct {
	NDArrayMultiaryBase
}

func NDArrayMultiaryKernelFrom(ptr unsafe.Pointer) NDArrayMultiaryKernel {
	return NDArrayMultiaryKernel{
		NDArrayMultiaryBase: NDArrayMultiaryBaseFrom(ptr),
	}
}

func (n_ NDArrayMultiaryKernel) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayMultiaryKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMultiaryKernel](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3175010-initwithdevice?language=objc
func NewNDArrayMultiaryKernelWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayMultiaryKernel {
	instance := NDArrayMultiaryKernelClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayMultiaryKernelClass) Alloc() NDArrayMultiaryKernel {
	rv := objc.Call[NDArrayMultiaryKernel](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NDArrayMultiaryKernelClass) New() NDArrayMultiaryKernel {
	rv := objc.Call[NDArrayMultiaryKernel](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayMultiaryKernel() NDArrayMultiaryKernel {
	return NDArrayMultiaryKernelClass.New()
}

func (n_ NDArrayMultiaryKernel) Init() NDArrayMultiaryKernel {
	rv := objc.Call[NDArrayMultiaryKernel](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayMultiaryKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayMultiaryKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMultiaryKernel](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayMultiaryKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayMultiaryKernel {
	instance := NDArrayMultiaryKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayMultiaryKernel) InitWithDevice(device metal.PDevice) NDArrayMultiaryKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMultiaryKernel](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNDArrayMultiaryKernelWithDevice(device metal.PDevice) NDArrayMultiaryKernel {
	instance := NDArrayMultiaryKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143526-encodetocommandbuffer?language=objc
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysDestinationArray(cmdBuf metal.PCommandBuffer, sourceArrays []INDArray, destination INDArray) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:sourceArrays:destinationArray:"), po0, sourceArrays, objc.Ptr(destination))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143526-encodetocommandbuffer?language=objc
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferObjectSourceArraysDestinationArray(cmdBufObject objc.IObject, sourceArrays []INDArray, destination INDArray) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:sourceArrays:destinationArray:"), objc.Ptr(cmdBufObject), sourceArrays, objc.Ptr(destination))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143527-encodetocommandbuffer?language=objc
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysResultStateDestinationArray(cmdBuf metal.PCommandBuffer, sourceArrays []INDArray, outGradientState IState, destination INDArray) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:sourceArrays:resultState:destinationArray:"), po0, sourceArrays, objc.Ptr(outGradientState), objc.Ptr(destination))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143527-encodetocommandbuffer?language=objc
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferObjectSourceArraysResultStateDestinationArray(cmdBufObject objc.IObject, sourceArrays []INDArray, outGradientState IState, destination INDArray) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:sourceArrays:resultState:destinationArray:"), objc.Ptr(cmdBufObject), sourceArrays, objc.Ptr(outGradientState), objc.Ptr(destination))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143525-encodetocommandbuffer?language=objc
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArrays(cmdBuf metal.PCommandBuffer, sourceArrays []INDArray) NDArray {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:sourceArrays:"), po0, sourceArrays)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143525-encodetocommandbuffer?language=objc
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferObjectSourceArrays(cmdBufObject objc.IObject, sourceArrays []INDArray) NDArray {
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:sourceArrays:"), objc.Ptr(cmdBufObject), sourceArrays)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143528-encodetocommandbuffer?language=objc
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary(cmdBuf metal.PCommandBuffer, sourceArrays []INDArray, outGradientState IState, outputStateIsTemporary bool) NDArray {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:sourceArrays:resultState:outputStateIsTemporary:"), po0, sourceArrays, objc.Ptr(outGradientState), outputStateIsTemporary)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143528-encodetocommandbuffer?language=objc
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferObjectSourceArraysResultStateOutputStateIsTemporary(cmdBufObject objc.IObject, sourceArrays []INDArray, outGradientState IState, outputStateIsTemporary bool) NDArray {
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:sourceArrays:resultState:outputStateIsTemporary:"), objc.Ptr(cmdBufObject), sourceArrays, objc.Ptr(outGradientState), outputStateIsTemporary)
	return rv
}
