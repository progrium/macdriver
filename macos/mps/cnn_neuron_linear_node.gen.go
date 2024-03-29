// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronLinearNode] class.
var CNNNeuronLinearNodeClass = _CNNNeuronLinearNodeClass{objc.GetClass("MPSCNNNeuronLinearNode")}

type _CNNNeuronLinearNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronLinearNode] class.
type ICNNNeuronLinearNode interface {
	ICNNNeuronNode
}

// A representation of a linear neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlinearnode?language=objc
type CNNNeuronLinearNode struct {
	CNNNeuronNode
}

func CNNNeuronLinearNodeFrom(ptr unsafe.Pointer) CNNNeuronLinearNode {
	return CNNNeuronLinearNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (c_ CNNNeuronLinearNode) InitWithSourceAB(sourceNode INNImageNode, a float64, b float64) CNNNeuronLinearNode {
	rv := objc.Call[CNNNeuronLinearNode](c_, objc.Sel("initWithSource:a:b:"), objc.Ptr(sourceNode), a, b)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlinearnode/2866495-initwithsource?language=objc
func NewCNNNeuronLinearNodeWithSourceAB(sourceNode INNImageNode, a float64, b float64) CNNNeuronLinearNode {
	instance := CNNNeuronLinearNodeClass.Alloc().InitWithSourceAB(sourceNode, a, b)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronLinearNode) InitWithSource(sourceNode INNImageNode) CNNNeuronLinearNode {
	rv := objc.Call[CNNNeuronLinearNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlinearnode/2921456-initwithsource?language=objc
func NewCNNNeuronLinearNodeWithSource(sourceNode INNImageNode) CNNNeuronLinearNode {
	instance := CNNNeuronLinearNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronLinearNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronLinearNode {
	rv := objc.Call[CNNNeuronLinearNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlinearnode/2921450-nodewithsource?language=objc
func CNNNeuronLinearNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronLinearNode {
	return CNNNeuronLinearNodeClass.NodeWithSource(sourceNode)
}

func (cc _CNNNeuronLinearNodeClass) NodeWithSourceAB(sourceNode INNImageNode, a float64, b float64) CNNNeuronLinearNode {
	rv := objc.Call[CNNNeuronLinearNode](cc, objc.Sel("nodeWithSource:a:b:"), objc.Ptr(sourceNode), a, b)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlinearnode/2866445-nodewithsource?language=objc
func CNNNeuronLinearNode_NodeWithSourceAB(sourceNode INNImageNode, a float64, b float64) CNNNeuronLinearNode {
	return CNNNeuronLinearNodeClass.NodeWithSourceAB(sourceNode, a, b)
}

func (cc _CNNNeuronLinearNodeClass) Alloc() CNNNeuronLinearNode {
	rv := objc.Call[CNNNeuronLinearNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNNeuronLinearNodeClass) New() CNNNeuronLinearNode {
	rv := objc.Call[CNNNeuronLinearNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronLinearNode() CNNNeuronLinearNode {
	return CNNNeuronLinearNodeClass.New()
}

func (c_ CNNNeuronLinearNode) Init() CNNNeuronLinearNode {
	rv := objc.Call[CNNNeuronLinearNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronLinearNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronLinearNode {
	rv := objc.Call[CNNNeuronLinearNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronLinearNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronLinearNode {
	return CNNNeuronLinearNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
