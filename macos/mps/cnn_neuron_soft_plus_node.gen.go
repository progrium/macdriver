// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronSoftPlusNode] class.
var CNNNeuronSoftPlusNodeClass = _CNNNeuronSoftPlusNodeClass{objc.GetClass("MPSCNNNeuronSoftPlusNode")}

type _CNNNeuronSoftPlusNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronSoftPlusNode] class.
type ICNNNeuronSoftPlusNode interface {
	ICNNNeuronNode
}

// A representation of a parametric softplus neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode?language=objc
type CNNNeuronSoftPlusNode struct {
	CNNNeuronNode
}

func CNNNeuronSoftPlusNodeFrom(ptr unsafe.Pointer) CNNNeuronSoftPlusNode {
	return CNNNeuronSoftPlusNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (c_ CNNNeuronSoftPlusNode) InitWithSourceAB(sourceNode INNImageNode, a float64, b float64) CNNNeuronSoftPlusNode {
	rv := objc.Call[CNNNeuronSoftPlusNode](c_, objc.Sel("initWithSource:a:b:"), objc.Ptr(sourceNode), a, b)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2866413-initwithsource?language=objc
func NewCNNNeuronSoftPlusNodeWithSourceAB(sourceNode INNImageNode, a float64, b float64) CNNNeuronSoftPlusNode {
	instance := CNNNeuronSoftPlusNodeClass.Alloc().InitWithSourceAB(sourceNode, a, b)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronSoftPlusNodeClass) NodeWithSourceAB(sourceNode INNImageNode, a float64, b float64) CNNNeuronSoftPlusNode {
	rv := objc.Call[CNNNeuronSoftPlusNode](cc, objc.Sel("nodeWithSource:a:b:"), objc.Ptr(sourceNode), a, b)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2866450-nodewithsource?language=objc
func CNNNeuronSoftPlusNode_NodeWithSourceAB(sourceNode INNImageNode, a float64, b float64) CNNNeuronSoftPlusNode {
	return CNNNeuronSoftPlusNodeClass.NodeWithSourceAB(sourceNode, a, b)
}

func (c_ CNNNeuronSoftPlusNode) InitWithSource(sourceNode INNImageNode) CNNNeuronSoftPlusNode {
	rv := objc.Call[CNNNeuronSoftPlusNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2921457-initwithsource?language=objc
func NewCNNNeuronSoftPlusNodeWithSource(sourceNode INNImageNode) CNNNeuronSoftPlusNode {
	instance := CNNNeuronSoftPlusNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronSoftPlusNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronSoftPlusNode {
	rv := objc.Call[CNNNeuronSoftPlusNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplusnode/2921449-nodewithsource?language=objc
func CNNNeuronSoftPlusNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronSoftPlusNode {
	return CNNNeuronSoftPlusNodeClass.NodeWithSource(sourceNode)
}

func (cc _CNNNeuronSoftPlusNodeClass) Alloc() CNNNeuronSoftPlusNode {
	rv := objc.Call[CNNNeuronSoftPlusNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNNeuronSoftPlusNodeClass) New() CNNNeuronSoftPlusNode {
	rv := objc.Call[CNNNeuronSoftPlusNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronSoftPlusNode() CNNNeuronSoftPlusNode {
	return CNNNeuronSoftPlusNodeClass.New()
}

func (c_ CNNNeuronSoftPlusNode) Init() CNNNeuronSoftPlusNode {
	rv := objc.Call[CNNNeuronSoftPlusNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronSoftPlusNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronSoftPlusNode {
	rv := objc.Call[CNNNeuronSoftPlusNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronSoftPlusNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronSoftPlusNode {
	return CNNNeuronSoftPlusNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
