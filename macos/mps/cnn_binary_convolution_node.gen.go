// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBinaryConvolutionNode] class.
var CNNBinaryConvolutionNodeClass = _CNNBinaryConvolutionNodeClass{objc.GetClass("MPSCNNBinaryConvolutionNode")}

type _CNNBinaryConvolutionNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNBinaryConvolutionNode] class.
type ICNNBinaryConvolutionNode interface {
	ICNNConvolutionNode
}

// A representation of a convolution kernel with binary weights and an input image using binary approximations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutionnode?language=objc
type CNNBinaryConvolutionNode struct {
	CNNConvolutionNode
}

func CNNBinaryConvolutionNodeFrom(ptr unsafe.Pointer) CNNBinaryConvolutionNode {
	return CNNBinaryConvolutionNode{
		CNNConvolutionNode: CNNConvolutionNodeFrom(ptr),
	}
}

func (cc _CNNBinaryConvolutionNodeClass) NodeWithSourceWeightsScaleValueTypeFlags(sourceNode INNImageNode, weights PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryConvolutionNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNBinaryConvolutionNode](cc, objc.Sel("nodeWithSource:weights:scaleValue:type:flags:"), objc.Ptr(sourceNode), po1, scaleValue, type_, flags)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutionnode/2866487-nodewithsource?language=objc
func CNNBinaryConvolutionNode_NodeWithSourceWeightsScaleValueTypeFlags(sourceNode INNImageNode, weights PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryConvolutionNode {
	return CNNBinaryConvolutionNodeClass.NodeWithSourceWeightsScaleValueTypeFlags(sourceNode, weights, scaleValue, type_, flags)
}

func (c_ CNNBinaryConvolutionNode) InitWithSourceWeightsScaleValueTypeFlags(sourceNode INNImageNode, weights PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryConvolutionNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNBinaryConvolutionNode](c_, objc.Sel("initWithSource:weights:scaleValue:type:flags:"), objc.Ptr(sourceNode), po1, scaleValue, type_, flags)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutionnode/2866509-initwithsource?language=objc
func NewCNNBinaryConvolutionNodeWithSourceWeightsScaleValueTypeFlags(sourceNode INNImageNode, weights PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryConvolutionNode {
	instance := CNNBinaryConvolutionNodeClass.Alloc().InitWithSourceWeightsScaleValueTypeFlags(sourceNode, weights, scaleValue, type_, flags)
	instance.Autorelease()
	return instance
}

func (cc _CNNBinaryConvolutionNodeClass) Alloc() CNNBinaryConvolutionNode {
	rv := objc.Call[CNNBinaryConvolutionNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNBinaryConvolutionNode_Alloc() CNNBinaryConvolutionNode {
	return CNNBinaryConvolutionNodeClass.Alloc()
}

func (cc _CNNBinaryConvolutionNodeClass) New() CNNBinaryConvolutionNode {
	rv := objc.Call[CNNBinaryConvolutionNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBinaryConvolutionNode() CNNBinaryConvolutionNode {
	return CNNBinaryConvolutionNodeClass.New()
}

func (c_ CNNBinaryConvolutionNode) Init() CNNBinaryConvolutionNode {
	rv := objc.Call[CNNBinaryConvolutionNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNBinaryConvolutionNodeClass) NodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNBinaryConvolutionNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNBinaryConvolutionNode](cc, objc.Sel("nodeWithSource:weights:"), objc.Ptr(sourceNode), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866436-nodewithsource?language=objc
func CNNBinaryConvolutionNode_NodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNBinaryConvolutionNode {
	return CNNBinaryConvolutionNodeClass.NodeWithSourceWeights(sourceNode, weights)
}

func (c_ CNNBinaryConvolutionNode) InitWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNBinaryConvolutionNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNBinaryConvolutionNode](c_, objc.Sel("initWithSource:weights:"), objc.Ptr(sourceNode), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866470-initwithsource?language=objc
func NewCNNBinaryConvolutionNodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNBinaryConvolutionNode {
	instance := CNNBinaryConvolutionNodeClass.Alloc().InitWithSourceWeights(sourceNode, weights)
	instance.Autorelease()
	return instance
}