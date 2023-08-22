// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayBinarySecondaryGradientKernel] class.
var NDArrayBinarySecondaryGradientKernelClass = _NDArrayBinarySecondaryGradientKernelClass{objc.GetClass("MPSNDArrayBinarySecondaryGradientKernel")}

type _NDArrayBinarySecondaryGradientKernelClass struct {
	objc.Class
}

// An interface definition for the [NDArrayBinarySecondaryGradientKernel] class.
type INDArrayBinarySecondaryGradientKernel interface {
	INDArrayMultiaryGradientKernel
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf metal.PCommandBuffer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) NDArray
	EncodeToCommandBufferObjectPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBufObject objc.IObject, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) NDArray
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarysecondarygradientkernel?language=objc
type NDArrayBinarySecondaryGradientKernel struct {
	NDArrayMultiaryGradientKernel
}

func NDArrayBinarySecondaryGradientKernelFrom(ptr unsafe.Pointer) NDArrayBinarySecondaryGradientKernel {
	return NDArrayBinarySecondaryGradientKernel{
		NDArrayMultiaryGradientKernel: NDArrayMultiaryGradientKernelFrom(ptr),
	}
}

func (n_ NDArrayBinarySecondaryGradientKernel) InitWithDevice(device metal.PDevice) NDArrayBinarySecondaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayBinarySecondaryGradientKernel](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarysecondarygradientkernel/3143519-initwithdevice?language=objc
func NewNDArrayBinarySecondaryGradientKernelWithDevice(device metal.PDevice) NDArrayBinarySecondaryGradientKernel {
	instance := NDArrayBinarySecondaryGradientKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayBinarySecondaryGradientKernelClass) Alloc() NDArrayBinarySecondaryGradientKernel {
	rv := objc.Call[NDArrayBinarySecondaryGradientKernel](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayBinarySecondaryGradientKernel_Alloc() NDArrayBinarySecondaryGradientKernel {
	return NDArrayBinarySecondaryGradientKernelClass.Alloc()
}

func (nc _NDArrayBinarySecondaryGradientKernelClass) New() NDArrayBinarySecondaryGradientKernel {
	rv := objc.Call[NDArrayBinarySecondaryGradientKernel](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayBinarySecondaryGradientKernel() NDArrayBinarySecondaryGradientKernel {
	return NDArrayBinarySecondaryGradientKernelClass.New()
}

func (n_ NDArrayBinarySecondaryGradientKernel) Init() NDArrayBinarySecondaryGradientKernel {
	rv := objc.Call[NDArrayBinarySecondaryGradientKernel](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayBinarySecondaryGradientKernel) InitWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayBinarySecondaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayBinarySecondaryGradientKernel](n_, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), po0, count, sourceGradientIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143524-initwithdevice?language=objc
func NewNDArrayBinarySecondaryGradientKernelWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayBinarySecondaryGradientKernel {
	instance := NDArrayBinarySecondaryGradientKernelClass.Alloc().InitWithDeviceSourceCountSourceGradientIndex(device, count, sourceGradientIndex)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayBinarySecondaryGradientKernel) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayBinarySecondaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayBinarySecondaryGradientKernel](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131741-initwithdevice?language=objc
func NewNDArrayBinarySecondaryGradientKernelWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayBinarySecondaryGradientKernel {
	instance := NDArrayBinarySecondaryGradientKernelClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayBinarySecondaryGradientKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayBinarySecondaryGradientKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayBinarySecondaryGradientKernel](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayBinarySecondaryGradientKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayBinarySecondaryGradientKernel {
	instance := NDArrayBinarySecondaryGradientKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarysecondarygradientkernel/3143517-encodetocommandbuffer?language=objc
func (n_ NDArrayBinarySecondaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf metal.PCommandBuffer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) NDArray {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:"), po0, objc.Ptr(primarySourceArray), objc.Ptr(secondarySourceArray), objc.Ptr(gradient), objc.Ptr(state))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarysecondarygradientkernel/3143517-encodetocommandbuffer?language=objc
func (n_ NDArrayBinarySecondaryGradientKernel) EncodeToCommandBufferObjectPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBufObject objc.IObject, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) NDArray {
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:"), objc.Ptr(cmdBufObject), objc.Ptr(primarySourceArray), objc.Ptr(secondarySourceArray), objc.Ptr(gradient), objc.Ptr(state))
	return rv
}
