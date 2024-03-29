// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNOptimizerStochasticGradientDescent] class.
var NNOptimizerStochasticGradientDescentClass = _NNOptimizerStochasticGradientDescentClass{objc.GetClass("MPSNNOptimizerStochasticGradientDescent")}

type _NNOptimizerStochasticGradientDescentClass struct {
	objc.Class
}

// An interface definition for the [NNOptimizerStochasticGradientDescent] class.
type INNOptimizerStochasticGradientDescent interface {
	INNOptimizer
	EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferObjectBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState(commandBufferObject objc.IObject, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, resultValuesMatrix IMatrix)
	EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, resultValuesMatrix IMatrix)
	EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferObjectBatchNormalizationStateInputMomentumVectorsResultState(commandBufferObject objc.IObject, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, resultState ICNNNormalizationGammaAndBetaState)
	EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState(commandBuffer metal.PCommandBuffer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferObjectConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState(commandBufferObject objc.IObject, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState)
	EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector(commandBuffer metal.PCommandBuffer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, resultValuesVector IVector)
	EncodeToCommandBufferObjectInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector(commandBufferObject objc.IObject, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, resultValuesVector IVector)
	MomentumScale() float64
	UseNestrovMomentum() bool
	UseNesterovMomentum() bool
}

// An optimization layer that performs a gradient descent with an optional momentum update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent?language=objc
type NNOptimizerStochasticGradientDescent struct {
	NNOptimizer
}

func NNOptimizerStochasticGradientDescentFrom(ptr unsafe.Pointer) NNOptimizerStochasticGradientDescent {
	return NNOptimizerStochasticGradientDescent{
		NNOptimizer: NNOptimizerFrom(ptr),
	}
}

func (n_ NNOptimizerStochasticGradientDescent) InitWithDeviceLearningRate(device metal.PDevice, learningRate float64) NNOptimizerStochasticGradientDescent {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerStochasticGradientDescent](n_, objc.Sel("initWithDevice:learningRate:"), po0, learningRate)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966741-initwithdevice?language=objc
func NewNNOptimizerStochasticGradientDescentWithDeviceLearningRate(device metal.PDevice, learningRate float64) NNOptimizerStochasticGradientDescent {
	instance := NNOptimizerStochasticGradientDescentClass.Alloc().InitWithDeviceLearningRate(device, learningRate)
	instance.Autorelease()
	return instance
}

func (n_ NNOptimizerStochasticGradientDescent) InitWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor(device metal.PDevice, momentumScale float64, useNestrovMomentum bool, optimizerDescriptor INNOptimizerDescriptor) NNOptimizerStochasticGradientDescent {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerStochasticGradientDescent](n_, objc.Sel("initWithDevice:momentumScale:useNestrovMomentum:optimizerDescriptor:"), po0, momentumScale, useNestrovMomentum, objc.Ptr(optimizerDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966742-initwithdevice?language=objc
func NewNNOptimizerStochasticGradientDescentWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor(device metal.PDevice, momentumScale float64, useNestrovMomentum bool, optimizerDescriptor INNOptimizerDescriptor) NNOptimizerStochasticGradientDescent {
	instance := NNOptimizerStochasticGradientDescentClass.Alloc().InitWithDeviceMomentumScaleUseNestrovMomentumOptimizerDescriptor(device, momentumScale, useNestrovMomentum, optimizerDescriptor)
	instance.Autorelease()
	return instance
}

func (n_ NNOptimizerStochasticGradientDescent) InitWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor(device metal.PDevice, momentumScale float64, useNesterovMomentum bool, optimizerDescriptor INNOptimizerDescriptor) NNOptimizerStochasticGradientDescent {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerStochasticGradientDescent](n_, objc.Sel("initWithDevice:momentumScale:useNesterovMomentum:optimizerDescriptor:"), po0, momentumScale, useNesterovMomentum, objc.Ptr(optimizerDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3675591-initwithdevice?language=objc
func NewNNOptimizerStochasticGradientDescentWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor(device metal.PDevice, momentumScale float64, useNesterovMomentum bool, optimizerDescriptor INNOptimizerDescriptor) NNOptimizerStochasticGradientDescent {
	instance := NNOptimizerStochasticGradientDescentClass.Alloc().InitWithDeviceMomentumScaleUseNesterovMomentumOptimizerDescriptor(device, momentumScale, useNesterovMomentum, optimizerDescriptor)
	instance.Autorelease()
	return instance
}

func (nc _NNOptimizerStochasticGradientDescentClass) Alloc() NNOptimizerStochasticGradientDescent {
	rv := objc.Call[NNOptimizerStochasticGradientDescent](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NNOptimizerStochasticGradientDescentClass) New() NNOptimizerStochasticGradientDescent {
	rv := objc.Call[NNOptimizerStochasticGradientDescent](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNOptimizerStochasticGradientDescent() NNOptimizerStochasticGradientDescent {
	return NNOptimizerStochasticGradientDescentClass.New()
}

func (n_ NNOptimizerStochasticGradientDescent) Init() NNOptimizerStochasticGradientDescent {
	rv := objc.Call[NNOptimizerStochasticGradientDescent](n_, objc.Sel("init"))
	return rv
}

func (n_ NNOptimizerStochasticGradientDescent) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizerStochasticGradientDescent {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerStochasticGradientDescent](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNOptimizerStochasticGradientDescent_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizerStochasticGradientDescent {
	instance := NNOptimizerStochasticGradientDescentClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (n_ NNOptimizerStochasticGradientDescent) InitWithDevice(device metal.PDevice) NNOptimizerStochasticGradientDescent {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerStochasticGradientDescent](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNNOptimizerStochasticGradientDescentWithDevice(device metal.PDevice) NNOptimizerStochasticGradientDescent {
	instance := NNOptimizerStochasticGradientDescentClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3013785-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:resultState:"), po0, objc.Ptr(batchNormalizationGradientState), objc.Ptr(batchNormalizationSourceState), inputMomentumVectors, objc.Ptr(resultState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3013785-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferObjectBatchNormalizationGradientStateBatchNormalizationSourceStateInputMomentumVectorsResultState(commandBufferObject objc.IObject, batchNormalizationGradientState ICNNBatchNormalizationState, batchNormalizationSourceState ICNNBatchNormalizationState, inputMomentumVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationGradientState:batchNormalizationSourceState:inputMomentumVectors:resultState:"), objc.Ptr(commandBufferObject), objc.Ptr(batchNormalizationGradientState), objc.Ptr(batchNormalizationSourceState), inputMomentumVectors, objc.Ptr(resultState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3197828-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, resultValuesMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:resultValuesMatrix:"), po0, objc.Ptr(inputGradientMatrix), objc.Ptr(inputValuesMatrix), objc.Ptr(inputMomentumMatrix), objc.Ptr(resultValuesMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3197828-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputMomentumMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputMomentumMatrix IMatrix, resultValuesMatrix IMatrix) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputMomentumMatrix:resultValuesMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(inputGradientMatrix), objc.Ptr(inputValuesMatrix), objc.Ptr(inputMomentumMatrix), objc.Ptr(resultValuesMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3019336-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferBatchNormalizationStateInputMomentumVectorsResultState(commandBuffer metal.PCommandBuffer, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:resultState:"), po0, objc.Ptr(batchNormalizationState), inputMomentumVectors, objc.Ptr(resultState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3019336-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferObjectBatchNormalizationStateInputMomentumVectorsResultState(commandBufferObject objc.IObject, batchNormalizationState ICNNBatchNormalizationState, inputMomentumVectors []IVector, resultState ICNNNormalizationGammaAndBetaState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:batchNormalizationState:inputMomentumVectors:resultState:"), objc.Ptr(commandBufferObject), objc.Ptr(batchNormalizationState), inputMomentumVectors, objc.Ptr(resultState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3013786-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState(commandBuffer metal.PCommandBuffer, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:resultState:"), po0, objc.Ptr(convolutionGradientState), objc.Ptr(convolutionSourceState), inputMomentumVectors, objc.Ptr(resultState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3013786-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferObjectConvolutionGradientStateConvolutionSourceStateInputMomentumVectorsResultState(commandBufferObject objc.IObject, convolutionGradientState ICNNConvolutionGradientState, convolutionSourceState ICNNConvolutionWeightsAndBiasesState, inputMomentumVectors []IVector, resultState ICNNConvolutionWeightsAndBiasesState) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:convolutionGradientState:convolutionSourceState:inputMomentumVectors:resultState:"), objc.Ptr(commandBufferObject), objc.Ptr(convolutionGradientState), objc.Ptr(convolutionSourceState), inputMomentumVectors, objc.Ptr(resultState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966740-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector(commandBuffer metal.PCommandBuffer, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, resultValuesVector IVector) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:resultValuesVector:"), po0, objc.Ptr(inputGradientVector), objc.Ptr(inputValuesVector), objc.Ptr(inputMomentumVector), objc.Ptr(resultValuesVector))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966740-encodetocommandbuffer?language=objc
func (n_ NNOptimizerStochasticGradientDescent) EncodeToCommandBufferObjectInputGradientVectorInputValuesVectorInputMomentumVectorResultValuesVector(commandBufferObject objc.IObject, inputGradientVector IVector, inputValuesVector IVector, inputMomentumVector IVector, resultValuesVector IVector) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientVector:inputValuesVector:inputMomentumVector:resultValuesVector:"), objc.Ptr(commandBufferObject), objc.Ptr(inputGradientVector), objc.Ptr(inputValuesVector), objc.Ptr(inputMomentumVector), objc.Ptr(resultValuesVector))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966743-momentumscale?language=objc
func (n_ NNOptimizerStochasticGradientDescent) MomentumScale() float64 {
	rv := objc.Call[float64](n_, objc.Sel("momentumScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/2966744-usenestrovmomentum?language=objc
func (n_ NNOptimizerStochasticGradientDescent) UseNestrovMomentum() bool {
	rv := objc.Call[bool](n_, objc.Sel("useNestrovMomentum"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerstochasticgradientdescent/3675592-usenesterovmomentum?language=objc
func (n_ NNOptimizerStochasticGradientDescent) UseNesterovMomentum() bool {
	rv := objc.Call[bool](n_, objc.Sel("useNesterovMomentum"))
	return rv
}
