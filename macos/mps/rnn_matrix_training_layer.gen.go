// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RNNMatrixTrainingLayer] class.
var RNNMatrixTrainingLayerClass = _RNNMatrixTrainingLayerClass{objc.GetClass("MPSRNNMatrixTrainingLayer")}

type _RNNMatrixTrainingLayerClass struct {
	objc.Class
}

// An interface definition for the [RNNMatrixTrainingLayer] class.
type IRNNMatrixTrainingLayer interface {
	IKernel
	CreateTemporaryWeightGradientMatricesDataTypeCommandBuffer(matricesOut foundation.IMutableArray, dataType DataType, commandBuffer metal.PCommandBuffer)
	CreateTemporaryWeightGradientMatricesDataTypeCommandBufferObject(matricesOut foundation.IMutableArray, dataType DataType, commandBufferObject objc.IObject)
	EncodeCopyWeightsToCommandBufferWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset(commandBuffer metal.PCommandBuffer, weights []IMatrix, matrixId RNNMatrixId, matrix IMatrix, copyFromWeightsToMatrix bool, matrixOffset metal.Origin)
	EncodeCopyWeightsToCommandBufferObjectWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset(commandBufferObject objc.IObject, weights []IMatrix, matrixId RNNMatrixId, matrix IMatrix, copyFromWeightsToMatrix bool, matrixOffset metal.Origin)
	EncodeForwardSequenceToCommandBufferSourceMatricesDestinationMatricesTrainingStatesWeights(commandBuffer metal.PCommandBuffer, sourceMatrices []IMatrix, destinationMatrices []IMatrix, trainingStates foundation.IMutableArray, weights []IMatrix)
	EncodeForwardSequenceToCommandBufferObjectSourceMatricesDestinationMatricesTrainingStatesWeights(commandBufferObject objc.IObject, sourceMatrices []IMatrix, destinationMatrices []IMatrix, trainingStates foundation.IMutableArray, weights []IMatrix)
	EncodeGradientSequenceToCommandBufferForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights(commandBuffer metal.PCommandBuffer, forwardSources []IMatrix, sourceGradients []IMatrix, destinationGradients []IMatrix, weightGradients []IMatrix, trainingStates []IRNNMatrixTrainingState, weights []IMatrix)
	EncodeGradientSequenceToCommandBufferObjectForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights(commandBufferObject objc.IObject, forwardSources []IMatrix, sourceGradients []IMatrix, destinationGradients []IMatrix, weightGradients []IMatrix, trainingStates []IRNNMatrixTrainingState, weights []IMatrix)
	CreateWeightMatrices(matricesOut foundation.IMutableArray)
	CreateWeightGradientMatricesDataType(matricesOut foundation.IMutableArray, dataType DataType)
	RecurrentOutputIsTemporary() bool
	SetRecurrentOutputIsTemporary(value bool)
	AccumulateWeightGradients() bool
	SetAccumulateWeightGradients(value bool)
	OutputFeatureChannels() uint
	InputFeatureChannels() uint
	TrainingStateIsTemporary() bool
	SetTrainingStateIsTemporary(value bool)
	StoreAllIntermediateStates() bool
	SetStoreAllIntermediateStates(value bool)
}

// A layer for training recurrent neural networks on Metal Performance Shaders matrices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer?language=objc
type RNNMatrixTrainingLayer struct {
	Kernel
}

func RNNMatrixTrainingLayerFrom(ptr unsafe.Pointer) RNNMatrixTrainingLayer {
	return RNNMatrixTrainingLayer{
		Kernel: KernelFrom(ptr),
	}
}

func (r_ RNNMatrixTrainingLayer) InitWithDeviceRnnDescriptorTrainableWeights(device metal.PDevice, rnnDescriptor IRNNDescriptor, trainableWeights foundation.IMutableArray) RNNMatrixTrainingLayer {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNMatrixTrainingLayer](r_, objc.Sel("initWithDevice:rnnDescriptor:trainableWeights:"), po0, objc.Ptr(rnnDescriptor), objc.Ptr(trainableWeights))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966794-initwithdevice?language=objc
func NewRNNMatrixTrainingLayerWithDeviceRnnDescriptorTrainableWeights(device metal.PDevice, rnnDescriptor IRNNDescriptor, trainableWeights foundation.IMutableArray) RNNMatrixTrainingLayer {
	instance := RNNMatrixTrainingLayerClass.Alloc().InitWithDeviceRnnDescriptorTrainableWeights(device, rnnDescriptor, trainableWeights)
	instance.Autorelease()
	return instance
}

func (r_ RNNMatrixTrainingLayer) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) RNNMatrixTrainingLayer {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNMatrixTrainingLayer](r_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966784-copywithzone?language=objc
func RNNMatrixTrainingLayer_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) RNNMatrixTrainingLayer {
	instance := RNNMatrixTrainingLayerClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (rc _RNNMatrixTrainingLayerClass) Alloc() RNNMatrixTrainingLayer {
	rv := objc.Call[RNNMatrixTrainingLayer](rc, objc.Sel("alloc"))
	return rv
}

func RNNMatrixTrainingLayer_Alloc() RNNMatrixTrainingLayer {
	return RNNMatrixTrainingLayerClass.Alloc()
}

func (rc _RNNMatrixTrainingLayerClass) New() RNNMatrixTrainingLayer {
	rv := objc.Call[RNNMatrixTrainingLayer](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRNNMatrixTrainingLayer() RNNMatrixTrainingLayer {
	return RNNMatrixTrainingLayerClass.New()
}

func (r_ RNNMatrixTrainingLayer) Init() RNNMatrixTrainingLayer {
	rv := objc.Call[RNNMatrixTrainingLayer](r_, objc.Sel("init"))
	return rv
}

func (r_ RNNMatrixTrainingLayer) InitWithDevice(device metal.PDevice) RNNMatrixTrainingLayer {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNMatrixTrainingLayer](r_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewRNNMatrixTrainingLayerWithDevice(device metal.PDevice) RNNMatrixTrainingLayer {
	instance := RNNMatrixTrainingLayerClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966785-createtemporaryweightgradientmat?language=objc
func (r_ RNNMatrixTrainingLayer) CreateTemporaryWeightGradientMatricesDataTypeCommandBuffer(matricesOut foundation.IMutableArray, dataType DataType, commandBuffer metal.PCommandBuffer) {
	po2 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](r_, objc.Sel("createTemporaryWeightGradientMatrices:dataType:commandBuffer:"), objc.Ptr(matricesOut), dataType, po2)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966785-createtemporaryweightgradientmat?language=objc
func (r_ RNNMatrixTrainingLayer) CreateTemporaryWeightGradientMatricesDataTypeCommandBufferObject(matricesOut foundation.IMutableArray, dataType DataType, commandBufferObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("createTemporaryWeightGradientMatrices:dataType:commandBuffer:"), objc.Ptr(matricesOut), dataType, objc.Ptr(commandBufferObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966788-encodecopyweightstocommandbuffer?language=objc
func (r_ RNNMatrixTrainingLayer) EncodeCopyWeightsToCommandBufferWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset(commandBuffer metal.PCommandBuffer, weights []IMatrix, matrixId RNNMatrixId, matrix IMatrix, copyFromWeightsToMatrix bool, matrixOffset metal.Origin) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](r_, objc.Sel("encodeCopyWeightsToCommandBuffer:weights:matrixId:matrix:copyFromWeightsToMatrix:matrixOffset:"), po0, weights, matrixId, objc.Ptr(matrix), copyFromWeightsToMatrix, matrixOffset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966788-encodecopyweightstocommandbuffer?language=objc
func (r_ RNNMatrixTrainingLayer) EncodeCopyWeightsToCommandBufferObjectWeightsMatrixIdMatrixCopyFromWeightsToMatrixMatrixOffset(commandBufferObject objc.IObject, weights []IMatrix, matrixId RNNMatrixId, matrix IMatrix, copyFromWeightsToMatrix bool, matrixOffset metal.Origin) {
	objc.Call[objc.Void](r_, objc.Sel("encodeCopyWeightsToCommandBuffer:weights:matrixId:matrix:copyFromWeightsToMatrix:matrixOffset:"), objc.Ptr(commandBufferObject), weights, matrixId, objc.Ptr(matrix), copyFromWeightsToMatrix, matrixOffset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966789-encodeforwardsequencetocommandbu?language=objc
func (r_ RNNMatrixTrainingLayer) EncodeForwardSequenceToCommandBufferSourceMatricesDestinationMatricesTrainingStatesWeights(commandBuffer metal.PCommandBuffer, sourceMatrices []IMatrix, destinationMatrices []IMatrix, trainingStates foundation.IMutableArray, weights []IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](r_, objc.Sel("encodeForwardSequenceToCommandBuffer:sourceMatrices:destinationMatrices:trainingStates:weights:"), po0, sourceMatrices, destinationMatrices, objc.Ptr(trainingStates), weights)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966789-encodeforwardsequencetocommandbu?language=objc
func (r_ RNNMatrixTrainingLayer) EncodeForwardSequenceToCommandBufferObjectSourceMatricesDestinationMatricesTrainingStatesWeights(commandBufferObject objc.IObject, sourceMatrices []IMatrix, destinationMatrices []IMatrix, trainingStates foundation.IMutableArray, weights []IMatrix) {
	objc.Call[objc.Void](r_, objc.Sel("encodeForwardSequenceToCommandBuffer:sourceMatrices:destinationMatrices:trainingStates:weights:"), objc.Ptr(commandBufferObject), sourceMatrices, destinationMatrices, objc.Ptr(trainingStates), weights)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966792-encodegradientsequencetocommandb?language=objc
func (r_ RNNMatrixTrainingLayer) EncodeGradientSequenceToCommandBufferForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights(commandBuffer metal.PCommandBuffer, forwardSources []IMatrix, sourceGradients []IMatrix, destinationGradients []IMatrix, weightGradients []IMatrix, trainingStates []IRNNMatrixTrainingState, weights []IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](r_, objc.Sel("encodeGradientSequenceToCommandBuffer:forwardSources:sourceGradients:destinationGradients:weightGradients:trainingStates:weights:"), po0, forwardSources, sourceGradients, destinationGradients, weightGradients, trainingStates, weights)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966792-encodegradientsequencetocommandb?language=objc
func (r_ RNNMatrixTrainingLayer) EncodeGradientSequenceToCommandBufferObjectForwardSourcesSourceGradientsDestinationGradientsWeightGradientsTrainingStatesWeights(commandBufferObject objc.IObject, forwardSources []IMatrix, sourceGradients []IMatrix, destinationGradients []IMatrix, weightGradients []IMatrix, trainingStates []IRNNMatrixTrainingState, weights []IMatrix) {
	objc.Call[objc.Void](r_, objc.Sel("encodeGradientSequenceToCommandBuffer:forwardSources:sourceGradients:destinationGradients:weightGradients:trainingStates:weights:"), objc.Ptr(commandBufferObject), forwardSources, sourceGradients, destinationGradients, weightGradients, trainingStates, weights)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966787-createweightmatrices?language=objc
func (r_ RNNMatrixTrainingLayer) CreateWeightMatrices(matricesOut foundation.IMutableArray) {
	objc.Call[objc.Void](r_, objc.Sel("createWeightMatrices:"), objc.Ptr(matricesOut))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966786-createweightgradientmatrices?language=objc
func (r_ RNNMatrixTrainingLayer) CreateWeightGradientMatricesDataType(matricesOut foundation.IMutableArray, dataType DataType) {
	objc.Call[objc.Void](r_, objc.Sel("createWeightGradientMatrices:dataType:"), objc.Ptr(matricesOut), dataType)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966797-recurrentoutputistemporary?language=objc
func (r_ RNNMatrixTrainingLayer) RecurrentOutputIsTemporary() bool {
	rv := objc.Call[bool](r_, objc.Sel("recurrentOutputIsTemporary"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966797-recurrentoutputistemporary?language=objc
func (r_ RNNMatrixTrainingLayer) SetRecurrentOutputIsTemporary(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setRecurrentOutputIsTemporary:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966783-accumulateweightgradients?language=objc
func (r_ RNNMatrixTrainingLayer) AccumulateWeightGradients() bool {
	rv := objc.Call[bool](r_, objc.Sel("accumulateWeightGradients"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966783-accumulateweightgradients?language=objc
func (r_ RNNMatrixTrainingLayer) SetAccumulateWeightGradients(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setAccumulateWeightGradients:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966796-outputfeaturechannels?language=objc
func (r_ RNNMatrixTrainingLayer) OutputFeatureChannels() uint {
	rv := objc.Call[uint](r_, objc.Sel("outputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966795-inputfeaturechannels?language=objc
func (r_ RNNMatrixTrainingLayer) InputFeatureChannels() uint {
	rv := objc.Call[uint](r_, objc.Sel("inputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966799-trainingstateistemporary?language=objc
func (r_ RNNMatrixTrainingLayer) TrainingStateIsTemporary() bool {
	rv := objc.Call[bool](r_, objc.Sel("trainingStateIsTemporary"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966799-trainingstateistemporary?language=objc
func (r_ RNNMatrixTrainingLayer) SetTrainingStateIsTemporary(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setTrainingStateIsTemporary:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966798-storeallintermediatestates?language=objc
func (r_ RNNMatrixTrainingLayer) StoreAllIntermediateStates() bool {
	rv := objc.Call[bool](r_, objc.Sel("storeAllIntermediateStates"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtraininglayer/2966798-storeallintermediatestates?language=objc
func (r_ RNNMatrixTrainingLayer) SetStoreAllIntermediateStates(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setStoreAllIntermediateStates:"), value)
}
