// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MatrixBatchNormalization] class.
var MatrixBatchNormalizationClass = _MatrixBatchNormalizationClass{objc.GetClass("MPSMatrixBatchNormalization")}

type _MatrixBatchNormalizationClass struct {
	objc.Class
}

// An interface definition for the [MatrixBatchNormalization] class.
type IMatrixBatchNormalization interface {
	IMatrixUnaryKernel
	NeuronParameterA() float32
	NeuronType() CNNNeuronType
	SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32)
	NeuronParameterB() float32
	EncodeToCommandBufferInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultMatrix IMatrix)
	EncodeToCommandBufferObjectInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultMatrix IMatrix)
	NeuronParameterC() float32
	ComputeStatistics() bool
	SetComputeStatistics(value bool)
	Epsilon() float32
	SetEpsilon(value float32)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
}

// A batch normalization kernel that operates on matrices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization?language=objc
type MatrixBatchNormalization struct {
	MatrixUnaryKernel
}

func MatrixBatchNormalizationFrom(ptr unsafe.Pointer) MatrixBatchNormalization {
	return MatrixBatchNormalization{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}

func (m_ MatrixBatchNormalization) InitWithDevice(device metal.PDevice) MatrixBatchNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixBatchNormalization](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980735-initwithdevice?language=objc
func NewMatrixBatchNormalizationWithDevice(device metal.PDevice) MatrixBatchNormalization {
	instance := MatrixBatchNormalizationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixBatchNormalization) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixBatchNormalization {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixBatchNormalization](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980731-copywithzone?language=objc
func MatrixBatchNormalization_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixBatchNormalization {
	instance := MatrixBatchNormalizationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixBatchNormalizationClass) Alloc() MatrixBatchNormalization {
	rv := objc.Call[MatrixBatchNormalization](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MatrixBatchNormalizationClass) New() MatrixBatchNormalization {
	rv := objc.Call[MatrixBatchNormalization](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixBatchNormalization() MatrixBatchNormalization {
	return MatrixBatchNormalizationClass.New()
}

func (m_ MatrixBatchNormalization) Init() MatrixBatchNormalization {
	rv := objc.Call[MatrixBatchNormalization](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980736-neuronparametera?language=objc
func (m_ MatrixBatchNormalization) NeuronParameterA() float32 {
	rv := objc.Call[float32](m_, objc.Sel("neuronParameterA"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980739-neurontype?language=objc
func (m_ MatrixBatchNormalization) NeuronType() CNNNeuronType {
	rv := objc.Call[CNNNeuronType](m_, objc.Sel("neuronType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980740-setneurontype?language=objc
func (m_ MatrixBatchNormalization) SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Call[objc.Void](m_, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980737-neuronparameterb?language=objc
func (m_ MatrixBatchNormalization) NeuronParameterB() float32 {
	rv := objc.Call[float32](m_, objc.Sel("neuronParameterB"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980732-encodetocommandbuffer?language=objc
func (m_ MatrixBatchNormalization) EncodeToCommandBufferInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:meanVector:varianceVector:gammaVector:betaVector:resultMatrix:"), po0, inputMatrix, meanVector, varianceVector, gammaVector, betaVector, resultMatrix)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980732-encodetocommandbuffer?language=objc
func (m_ MatrixBatchNormalization) EncodeToCommandBufferObjectInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:meanVector:varianceVector:gammaVector:betaVector:resultMatrix:"), commandBufferObject, inputMatrix, meanVector, varianceVector, gammaVector, betaVector, resultMatrix)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980738-neuronparameterc?language=objc
func (m_ MatrixBatchNormalization) NeuronParameterC() float32 {
	rv := objc.Call[float32](m_, objc.Sel("neuronParameterC"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980730-computestatistics?language=objc
func (m_ MatrixBatchNormalization) ComputeStatistics() bool {
	rv := objc.Call[bool](m_, objc.Sel("computeStatistics"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980730-computestatistics?language=objc
func (m_ MatrixBatchNormalization) SetComputeStatistics(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setComputeStatistics:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980733-epsilon?language=objc
func (m_ MatrixBatchNormalization) Epsilon() float32 {
	rv := objc.Call[float32](m_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980733-epsilon?language=objc
func (m_ MatrixBatchNormalization) SetEpsilon(value float32) {
	objc.Call[objc.Void](m_, objc.Sel("setEpsilon:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980742-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixBatchNormalization) SourceNumberOfFeatureVectors() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980742-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixBatchNormalization) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980741-sourceinputfeaturechannels?language=objc
func (m_ MatrixBatchNormalization) SourceInputFeatureChannels() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceInputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980741-sourceinputfeaturechannels?language=objc
func (m_ MatrixBatchNormalization) SetSourceInputFeatureChannels(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceInputFeatureChannels:"), value)
}
