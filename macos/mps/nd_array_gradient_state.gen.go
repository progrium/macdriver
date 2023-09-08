// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayGradientState] class.
var NDArrayGradientStateClass = _NDArrayGradientStateClass{objc.GetClass("MPSNDArrayGradientState")}

type _NDArrayGradientStateClass struct {
	objc.Class
}

// An interface definition for the [NDArrayGradientState] class.
type INDArrayGradientState interface {
	IState
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraygradientstate?language=objc
type NDArrayGradientState struct {
	State
}

func NDArrayGradientStateFrom(ptr unsafe.Pointer) NDArrayGradientState {
	return NDArrayGradientState{
		State: StateFrom(ptr),
	}
}

func (nc _NDArrayGradientStateClass) Alloc() NDArrayGradientState {
	rv := objc.Call[NDArrayGradientState](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NDArrayGradientStateClass) New() NDArrayGradientState {
	rv := objc.Call[NDArrayGradientState](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayGradientState() NDArrayGradientState {
	return NDArrayGradientStateClass.New()
}

func (n_ NDArrayGradientState) Init() NDArrayGradientState {
	rv := objc.Call[NDArrayGradientState](n_, objc.Sel("init"))
	return rv
}

func (nc _NDArrayGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NDArrayGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[NDArrayGradientState](nc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func NDArrayGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NDArrayGradientState {
	return NDArrayGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

func (n_ NDArrayGradientState) InitWithDeviceTextureDescriptor(device metal.PDevice, descriptor metal.ITextureDescriptor) NDArrayGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGradientState](n_, objc.Sel("initWithDevice:textureDescriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942400-initwithdevice?language=objc
func NewNDArrayGradientStateWithDeviceTextureDescriptor(device metal.PDevice, descriptor metal.ITextureDescriptor) NDArrayGradientState {
	instance := NDArrayGradientStateClass.Alloc().InitWithDeviceTextureDescriptor(device, descriptor)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGradientState) InitWithResource(resource metal.PResource) NDArrayGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[NDArrayGradientState](n_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewNDArrayGradientStateWithResource(resource metal.PResource) NDArrayGradientState {
	instance := NDArrayGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NDArrayGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGradientState](n_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewNDArrayGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NDArrayGradientState {
	instance := NDArrayGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGradientState) InitWithResources(resources []metal.PResource) NDArrayGradientState {
	rv := objc.Call[NDArrayGradientState](n_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewNDArrayGradientStateWithResources(resources []metal.PResource) NDArrayGradientState {
	instance := NDArrayGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayGradientStateClass) TemporaryStateWithCommandBufferBufferSize(cmdBuf metal.PCommandBuffer, bufferSize uint) NDArrayGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[NDArrayGradientState](nc, objc.Sel("temporaryStateWithCommandBuffer:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942391-temporarystatewithcommandbuffer?language=objc
func NDArrayGradientState_TemporaryStateWithCommandBufferBufferSize(cmdBuf metal.PCommandBuffer, bufferSize uint) NDArrayGradientState {
	return NDArrayGradientStateClass.TemporaryStateWithCommandBufferBufferSize(cmdBuf, bufferSize)
}

func (nc _NDArrayGradientStateClass) TemporaryStateWithCommandBuffer(cmdBuf metal.PCommandBuffer) NDArrayGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[NDArrayGradientState](nc, objc.Sel("temporaryStateWithCommandBuffer:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942393-temporarystatewithcommandbuffer?language=objc
func NDArrayGradientState_TemporaryStateWithCommandBuffer(cmdBuf metal.PCommandBuffer) NDArrayGradientState {
	return NDArrayGradientStateClass.TemporaryStateWithCommandBuffer(cmdBuf)
}

func (nc _NDArrayGradientStateClass) TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf metal.PCommandBuffer, descriptor metal.ITextureDescriptor) NDArrayGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[NDArrayGradientState](nc, objc.Sel("temporaryStateWithCommandBuffer:textureDescriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942395-temporarystatewithcommandbuffer?language=objc
func NDArrayGradientState_TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf metal.PCommandBuffer, descriptor metal.ITextureDescriptor) NDArrayGradientState {
	return NDArrayGradientStateClass.TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf, descriptor)
}

func (n_ NDArrayGradientState) InitWithDeviceResourceList(device metal.PDevice, resourceList IStateResourceList) NDArrayGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGradientState](n_, objc.Sel("initWithDevice:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947908-initwithdevice?language=objc
func NewNDArrayGradientStateWithDeviceResourceList(device metal.PDevice, resourceList IStateResourceList) NDArrayGradientState {
	instance := NDArrayGradientStateClass.Alloc().InitWithDeviceResourceList(device, resourceList)
	instance.Autorelease()
	return instance
}
