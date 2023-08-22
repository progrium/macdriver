// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RNNImageInferenceLayer] class.
var RNNImageInferenceLayerClass = _RNNImageInferenceLayerClass{objc.GetClass("MPSRNNImageInferenceLayer")}

type _RNNImageInferenceLayerClass struct {
	objc.Class
}

// An interface definition for the [RNNImageInferenceLayer] class.
type IRNNImageInferenceLayer interface {
	ICNNKernel
	EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardImagesDestinationBackwardImages(commandBuffer metal.PCommandBuffer, sourceSequence []IImage, destinationForwardImages []IImage, destinationBackwardImages []IImage)
	EncodeBidirectionalSequenceToCommandBufferObjectSourceSequenceDestinationForwardImagesDestinationBackwardImages(commandBufferObject objc.IObject, sourceSequence []IImage, destinationForwardImages []IImage, destinationBackwardImages []IImage)
	EncodeSequenceToCommandBufferSourceImagesDestinationImagesRecurrentInputStateRecurrentOutputStates(commandBuffer metal.PCommandBuffer, sourceImages []IImage, destinationImages []IImage, recurrentInputState IRNNRecurrentImageState, recurrentOutputStates foundation.IMutableArray)
	EncodeSequenceToCommandBufferObjectSourceImagesDestinationImagesRecurrentInputStateRecurrentOutputStates(commandBufferObject objc.IObject, sourceImages []IImage, destinationImages []IImage, recurrentInputState IRNNRecurrentImageState, recurrentOutputStates foundation.IMutableArray)
	RecurrentOutputIsTemporary() bool
	SetRecurrentOutputIsTemporary(value bool)
	OutputFeatureChannels() uint
	InputFeatureChannels() uint
	BidirectionalCombineMode() RNNBidirectionalCombineMode
	SetBidirectionalCombineMode(value RNNBidirectionalCombineMode)
	StoreAllIntermediateStates() bool
	SetStoreAllIntermediateStates(value bool)
	NumberOfLayers() uint
}

// A recurrent neural network layer for inference on Metal Performance Shaders images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer?language=objc
type RNNImageInferenceLayer struct {
	CNNKernel
}

func RNNImageInferenceLayerFrom(ptr unsafe.Pointer) RNNImageInferenceLayer {
	return RNNImageInferenceLayer{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (r_ RNNImageInferenceLayer) InitWithDeviceRnnDescriptor(device metal.PDevice, rnnDescriptor IRNNDescriptor) RNNImageInferenceLayer {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNImageInferenceLayer](r_, objc.Sel("initWithDevice:rnnDescriptor:"), po0, objc.Ptr(rnnDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865691-initwithdevice?language=objc
func NewRNNImageInferenceLayerWithDeviceRnnDescriptor(device metal.PDevice, rnnDescriptor IRNNDescriptor) RNNImageInferenceLayer {
	instance := RNNImageInferenceLayerClass.Alloc().InitWithDeviceRnnDescriptor(device, rnnDescriptor)
	instance.Autorelease()
	return instance
}

func (r_ RNNImageInferenceLayer) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) RNNImageInferenceLayer {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNImageInferenceLayer](r_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865728-copywithzone?language=objc
func RNNImageInferenceLayer_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) RNNImageInferenceLayer {
	instance := RNNImageInferenceLayerClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (rc _RNNImageInferenceLayerClass) Alloc() RNNImageInferenceLayer {
	rv := objc.Call[RNNImageInferenceLayer](rc, objc.Sel("alloc"))
	return rv
}

func RNNImageInferenceLayer_Alloc() RNNImageInferenceLayer {
	return RNNImageInferenceLayerClass.Alloc()
}

func (rc _RNNImageInferenceLayerClass) New() RNNImageInferenceLayer {
	rv := objc.Call[RNNImageInferenceLayer](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRNNImageInferenceLayer() RNNImageInferenceLayer {
	return RNNImageInferenceLayerClass.New()
}

func (r_ RNNImageInferenceLayer) Init() RNNImageInferenceLayer {
	rv := objc.Call[RNNImageInferenceLayer](r_, objc.Sel("init"))
	return rv
}

func (r_ RNNImageInferenceLayer) InitWithDevice(device metal.PDevice) RNNImageInferenceLayer {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNImageInferenceLayer](r_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewRNNImageInferenceLayerWithDevice(device metal.PDevice) RNNImageInferenceLayer {
	instance := RNNImageInferenceLayerClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865693-encodebidirectionalsequencetocom?language=objc
func (r_ RNNImageInferenceLayer) EncodeBidirectionalSequenceToCommandBufferSourceSequenceDestinationForwardImagesDestinationBackwardImages(commandBuffer metal.PCommandBuffer, sourceSequence []IImage, destinationForwardImages []IImage, destinationBackwardImages []IImage) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](r_, objc.Sel("encodeBidirectionalSequenceToCommandBuffer:sourceSequence:destinationForwardImages:destinationBackwardImages:"), po0, sourceSequence, destinationForwardImages, destinationBackwardImages)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865693-encodebidirectionalsequencetocom?language=objc
func (r_ RNNImageInferenceLayer) EncodeBidirectionalSequenceToCommandBufferObjectSourceSequenceDestinationForwardImagesDestinationBackwardImages(commandBufferObject objc.IObject, sourceSequence []IImage, destinationForwardImages []IImage, destinationBackwardImages []IImage) {
	objc.Call[objc.Void](r_, objc.Sel("encodeBidirectionalSequenceToCommandBuffer:sourceSequence:destinationForwardImages:destinationBackwardImages:"), objc.Ptr(commandBufferObject), sourceSequence, destinationForwardImages, destinationBackwardImages)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865717-encodesequencetocommandbuffer?language=objc
func (r_ RNNImageInferenceLayer) EncodeSequenceToCommandBufferSourceImagesDestinationImagesRecurrentInputStateRecurrentOutputStates(commandBuffer metal.PCommandBuffer, sourceImages []IImage, destinationImages []IImage, recurrentInputState IRNNRecurrentImageState, recurrentOutputStates foundation.IMutableArray) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](r_, objc.Sel("encodeSequenceToCommandBuffer:sourceImages:destinationImages:recurrentInputState:recurrentOutputStates:"), po0, sourceImages, destinationImages, objc.Ptr(recurrentInputState), objc.Ptr(recurrentOutputStates))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865717-encodesequencetocommandbuffer?language=objc
func (r_ RNNImageInferenceLayer) EncodeSequenceToCommandBufferObjectSourceImagesDestinationImagesRecurrentInputStateRecurrentOutputStates(commandBufferObject objc.IObject, sourceImages []IImage, destinationImages []IImage, recurrentInputState IRNNRecurrentImageState, recurrentOutputStates foundation.IMutableArray) {
	objc.Call[objc.Void](r_, objc.Sel("encodeSequenceToCommandBuffer:sourceImages:destinationImages:recurrentInputState:recurrentOutputStates:"), objc.Ptr(commandBufferObject), sourceImages, destinationImages, objc.Ptr(recurrentInputState), objc.Ptr(recurrentOutputStates))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865749-recurrentoutputistemporary?language=objc
func (r_ RNNImageInferenceLayer) RecurrentOutputIsTemporary() bool {
	rv := objc.Call[bool](r_, objc.Sel("recurrentOutputIsTemporary"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865749-recurrentoutputistemporary?language=objc
func (r_ RNNImageInferenceLayer) SetRecurrentOutputIsTemporary(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setRecurrentOutputIsTemporary:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2890140-outputfeaturechannels?language=objc
func (r_ RNNImageInferenceLayer) OutputFeatureChannels() uint {
	rv := objc.Call[uint](r_, objc.Sel("outputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2890141-inputfeaturechannels?language=objc
func (r_ RNNImageInferenceLayer) InputFeatureChannels() uint {
	rv := objc.Call[uint](r_, objc.Sel("inputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865737-bidirectionalcombinemode?language=objc
func (r_ RNNImageInferenceLayer) BidirectionalCombineMode() RNNBidirectionalCombineMode {
	rv := objc.Call[RNNBidirectionalCombineMode](r_, objc.Sel("bidirectionalCombineMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865737-bidirectionalcombinemode?language=objc
func (r_ RNNImageInferenceLayer) SetBidirectionalCombineMode(value RNNBidirectionalCombineMode) {
	objc.Call[objc.Void](r_, objc.Sel("setBidirectionalCombineMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865706-storeallintermediatestates?language=objc
func (r_ RNNImageInferenceLayer) StoreAllIntermediateStates() bool {
	rv := objc.Call[bool](r_, objc.Sel("storeAllIntermediateStates"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865706-storeallintermediatestates?language=objc
func (r_ RNNImageInferenceLayer) SetStoreAllIntermediateStates(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setStoreAllIntermediateStates:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnimageinferencelayer/2865697-numberoflayers?language=objc
func (r_ RNNImageInferenceLayer) NumberOfLayers() uint {
	rv := objc.Call[uint](r_, objc.Sel("numberOfLayers"))
	return rv
}
