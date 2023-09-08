// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBatchNormalizationState] class.
var CNNBatchNormalizationStateClass = _CNNBatchNormalizationStateClass{objc.GetClass("MPSCNNBatchNormalizationState")}

type _CNNBatchNormalizationStateClass struct {
	objc.Class
}

// An interface definition for the [CNNBatchNormalizationState] class.
type ICNNBatchNormalizationState interface {
	INNGradientState
	Mean() metal.BufferObject
	GradientForGamma() metal.BufferObject
	Reset()
	Beta() metal.BufferObject
	GradientForBeta() metal.BufferObject
	Gamma() metal.BufferObject
	Variance() metal.BufferObject
	BatchNormalization() CNNBatchNormalization
}

// An object that stores data required to execute batch normalization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate?language=objc
type CNNBatchNormalizationState struct {
	NNGradientState
}

func CNNBatchNormalizationStateFrom(ptr unsafe.Pointer) CNNBatchNormalizationState {
	return CNNBatchNormalizationState{
		NNGradientState: NNGradientStateFrom(ptr),
	}
}

func (cc _CNNBatchNormalizationStateClass) Alloc() CNNBatchNormalizationState {
	rv := objc.Call[CNNBatchNormalizationState](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNBatchNormalizationStateClass) New() CNNBatchNormalizationState {
	rv := objc.Call[CNNBatchNormalizationState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBatchNormalizationState() CNNBatchNormalizationState {
	return CNNBatchNormalizationStateClass.New()
}

func (c_ CNNBatchNormalizationState) Init() CNNBatchNormalizationState {
	rv := objc.Call[CNNBatchNormalizationState](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNBatchNormalizationStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNBatchNormalizationState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNBatchNormalizationState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNBatchNormalizationState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNBatchNormalizationState {
	return CNNBatchNormalizationStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

func (c_ CNNBatchNormalizationState) InitWithDeviceTextureDescriptor(device metal.PDevice, descriptor metal.ITextureDescriptor) CNNBatchNormalizationState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationState](c_, objc.Sel("initWithDevice:textureDescriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942400-initwithdevice?language=objc
func NewCNNBatchNormalizationStateWithDeviceTextureDescriptor(device metal.PDevice, descriptor metal.ITextureDescriptor) CNNBatchNormalizationState {
	instance := CNNBatchNormalizationStateClass.Alloc().InitWithDeviceTextureDescriptor(device, descriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNBatchNormalizationState) InitWithResource(resource metal.PResource) CNNBatchNormalizationState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNBatchNormalizationState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNBatchNormalizationStateWithResource(resource metal.PResource) CNNBatchNormalizationState {
	instance := CNNBatchNormalizationStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (c_ CNNBatchNormalizationState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNBatchNormalizationState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNBatchNormalizationStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNBatchNormalizationState {
	instance := CNNBatchNormalizationStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNBatchNormalizationState) InitWithResources(resources []metal.PResource) CNNBatchNormalizationState {
	rv := objc.Call[CNNBatchNormalizationState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNBatchNormalizationStateWithResources(resources []metal.PResource) CNNBatchNormalizationState {
	instance := CNNBatchNormalizationStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (cc _CNNBatchNormalizationStateClass) TemporaryStateWithCommandBufferBufferSize(cmdBuf metal.PCommandBuffer, bufferSize uint) CNNBatchNormalizationState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNBatchNormalizationState](cc, objc.Sel("temporaryStateWithCommandBuffer:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942391-temporarystatewithcommandbuffer?language=objc
func CNNBatchNormalizationState_TemporaryStateWithCommandBufferBufferSize(cmdBuf metal.PCommandBuffer, bufferSize uint) CNNBatchNormalizationState {
	return CNNBatchNormalizationStateClass.TemporaryStateWithCommandBufferBufferSize(cmdBuf, bufferSize)
}

func (cc _CNNBatchNormalizationStateClass) TemporaryStateWithCommandBuffer(cmdBuf metal.PCommandBuffer) CNNBatchNormalizationState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNBatchNormalizationState](cc, objc.Sel("temporaryStateWithCommandBuffer:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942393-temporarystatewithcommandbuffer?language=objc
func CNNBatchNormalizationState_TemporaryStateWithCommandBuffer(cmdBuf metal.PCommandBuffer) CNNBatchNormalizationState {
	return CNNBatchNormalizationStateClass.TemporaryStateWithCommandBuffer(cmdBuf)
}

func (cc _CNNBatchNormalizationStateClass) TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf metal.PCommandBuffer, descriptor metal.ITextureDescriptor) CNNBatchNormalizationState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNBatchNormalizationState](cc, objc.Sel("temporaryStateWithCommandBuffer:textureDescriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942395-temporarystatewithcommandbuffer?language=objc
func CNNBatchNormalizationState_TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf metal.PCommandBuffer, descriptor metal.ITextureDescriptor) CNNBatchNormalizationState {
	return CNNBatchNormalizationStateClass.TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf, descriptor)
}

func (c_ CNNBatchNormalizationState) InitWithDeviceResourceList(device metal.PDevice, resourceList IStateResourceList) CNNBatchNormalizationState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationState](c_, objc.Sel("initWithDevice:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947908-initwithdevice?language=objc
func NewCNNBatchNormalizationStateWithDeviceResourceList(device metal.PDevice, resourceList IStateResourceList) CNNBatchNormalizationState {
	instance := CNNBatchNormalizationStateClass.Alloc().InitWithDeviceResourceList(device, resourceList)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2942612-mean?language=objc
func (c_ CNNBatchNormalizationState) Mean() metal.BufferObject {
	rv := objc.Call[metal.BufferObject](c_, objc.Sel("mean"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951893-gradientforgamma?language=objc
func (c_ CNNBatchNormalizationState) GradientForGamma() metal.BufferObject {
	rv := objc.Call[metal.BufferObject](c_, objc.Sel("gradientForGamma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2942587-reset?language=objc
func (c_ CNNBatchNormalizationState) Reset() {
	objc.Call[objc.Void](c_, objc.Sel("reset"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951888-beta?language=objc
func (c_ CNNBatchNormalizationState) Beta() metal.BufferObject {
	rv := objc.Call[metal.BufferObject](c_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951890-gradientforbeta?language=objc
func (c_ CNNBatchNormalizationState) GradientForBeta() metal.BufferObject {
	rv := objc.Call[metal.BufferObject](c_, objc.Sel("gradientForBeta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951892-gamma?language=objc
func (c_ CNNBatchNormalizationState) Gamma() metal.BufferObject {
	rv := objc.Call[metal.BufferObject](c_, objc.Sel("gamma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2942603-variance?language=objc
func (c_ CNNBatchNormalizationState) Variance() metal.BufferObject {
	rv := objc.Call[metal.BufferObject](c_, objc.Sel("variance"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2953969-batchnormalization?language=objc
func (c_ CNNBatchNormalizationState) BatchNormalization() CNNBatchNormalization {
	rv := objc.Call[CNNBatchNormalization](c_, objc.Sel("batchNormalization"))
	return rv
}
