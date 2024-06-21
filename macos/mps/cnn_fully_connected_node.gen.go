// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNFullyConnectedNode] class.
var CNNFullyConnectedNodeClass = _CNNFullyConnectedNodeClass{objc.GetClass("MPSCNNFullyConnectedNode")}

type _CNNFullyConnectedNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNFullyConnectedNode] class.
type ICNNFullyConnectedNode interface {
	ICNNConvolutionNode
}

// A representation of a fully connected convolution layer, also known as an inner product layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectednode?language=objc
type CNNFullyConnectedNode struct {
	CNNConvolutionNode
}

func CNNFullyConnectedNodeFrom(ptr unsafe.Pointer) CNNFullyConnectedNode {
	return CNNFullyConnectedNode{
		CNNConvolutionNode: CNNConvolutionNodeFrom(ptr),
	}
}

func (cc _CNNFullyConnectedNodeClass) NodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNFullyConnectedNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNFullyConnectedNode](cc, objc.Sel("nodeWithSource:weights:"), sourceNode, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectednode/2866458-nodewithsource?language=objc
func CNNFullyConnectedNode_NodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNFullyConnectedNode {
	return CNNFullyConnectedNodeClass.NodeWithSourceWeights(sourceNode, weights)
}

func (c_ CNNFullyConnectedNode) InitWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNFullyConnectedNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNFullyConnectedNode](c_, objc.Sel("initWithSource:weights:"), sourceNode, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectednode/2866412-initwithsource?language=objc
func NewCNNFullyConnectedNodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNFullyConnectedNode {
	instance := CNNFullyConnectedNodeClass.Alloc().InitWithSourceWeights(sourceNode, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNFullyConnectedNodeClass) Alloc() CNNFullyConnectedNode {
	rv := objc.Call[CNNFullyConnectedNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNFullyConnectedNodeClass) New() CNNFullyConnectedNode {
	rv := objc.Call[CNNFullyConnectedNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNFullyConnectedNode() CNNFullyConnectedNode {
	return CNNFullyConnectedNodeClass.New()
}

func (c_ CNNFullyConnectedNode) Init() CNNFullyConnectedNode {
	rv := objc.Call[CNNFullyConnectedNode](c_, objc.Sel("init"))
	return rv
}
