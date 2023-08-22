// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNSpatialNormalizationGradientNode] class.
var CNNSpatialNormalizationGradientNodeClass = _CNNSpatialNormalizationGradientNodeClass{objc.GetClass("MPSCNNSpatialNormalizationGradientNode")}

type _CNNSpatialNormalizationGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNSpatialNormalizationGradientNode] class.
type ICNNSpatialNormalizationGradientNode interface {
	INNGradientFilterNode
	Beta() float64
	SetBeta(value float64)
	Delta() float64
	SetDelta(value float64)
	Alpha() float64
	SetAlpha(value float64)
	KernelHeight() uint
	SetKernelHeight(value uint)
	KernelWidth() uint
	SetKernelWidth(value uint)
}

// A representation of a gradient spatial normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode?language=objc
type CNNSpatialNormalizationGradientNode struct {
	NNGradientFilterNode
}

func CNNSpatialNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNSpatialNormalizationGradientNode {
	return CNNSpatialNormalizationGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNSpatialNormalizationGradientNode) InitWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelSize uint) CNNSpatialNormalizationGradientNode {
	rv := objc.Call[CNNSpatialNormalizationGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelSize:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948009-initwithsourcegradient?language=objc
func NewCNNSpatialNormalizationGradientNodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelSize uint) CNNSpatialNormalizationGradientNode {
	instance := CNNSpatialNormalizationGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient, sourceImage, gradientState, kernelSize)
	instance.Autorelease()
	return instance
}

func (cc _CNNSpatialNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelSize uint) CNNSpatialNormalizationGradientNode {
	rv := objc.Call[CNNSpatialNormalizationGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelSize:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2947978-nodewithsourcegradient?language=objc
func CNNSpatialNormalizationGradientNode_NodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelSize uint) CNNSpatialNormalizationGradientNode {
	return CNNSpatialNormalizationGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient, sourceImage, gradientState, kernelSize)
}

func (cc _CNNSpatialNormalizationGradientNodeClass) Alloc() CNNSpatialNormalizationGradientNode {
	rv := objc.Call[CNNSpatialNormalizationGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNSpatialNormalizationGradientNode_Alloc() CNNSpatialNormalizationGradientNode {
	return CNNSpatialNormalizationGradientNodeClass.Alloc()
}

func (cc _CNNSpatialNormalizationGradientNodeClass) New() CNNSpatialNormalizationGradientNode {
	rv := objc.Call[CNNSpatialNormalizationGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNSpatialNormalizationGradientNode() CNNSpatialNormalizationGradientNode {
	return CNNSpatialNormalizationGradientNodeClass.New()
}

func (c_ CNNSpatialNormalizationGradientNode) Init() CNNSpatialNormalizationGradientNode {
	rv := objc.Call[CNNSpatialNormalizationGradientNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948006-beta?language=objc
func (c_ CNNSpatialNormalizationGradientNode) Beta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948006-beta?language=objc
func (c_ CNNSpatialNormalizationGradientNode) SetBeta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBeta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2947968-delta?language=objc
func (c_ CNNSpatialNormalizationGradientNode) Delta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2947968-delta?language=objc
func (c_ CNNSpatialNormalizationGradientNode) SetDelta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setDelta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948027-alpha?language=objc
func (c_ CNNSpatialNormalizationGradientNode) Alpha() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948027-alpha?language=objc
func (c_ CNNSpatialNormalizationGradientNode) SetAlpha(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAlpha:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948020-kernelheight?language=objc
func (c_ CNNSpatialNormalizationGradientNode) KernelHeight() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948020-kernelheight?language=objc
func (c_ CNNSpatialNormalizationGradientNode) SetKernelHeight(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelHeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948013-kernelwidth?language=objc
func (c_ CNNSpatialNormalizationGradientNode) KernelWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalizationgradientnode/2948013-kernelwidth?language=objc
func (c_ CNNSpatialNormalizationGradientNode) SetKernelWidth(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelWidth:"), value)
}
