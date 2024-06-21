// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [NNAdditionNode] class.
var NNAdditionNodeClass = _NNAdditionNodeClass{objc.GetClass("MPSNNAdditionNode")}

type _NNAdditionNodeClass struct {
	objc.Class
}

// An interface definition for the [NNAdditionNode] class.
type INNAdditionNode interface {
	INNBinaryArithmeticNode
}

// A representation of an addition operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnadditionnode?language=objc
type NNAdditionNode struct {
	NNBinaryArithmeticNode
}

func NNAdditionNodeFrom(ptr unsafe.Pointer) NNAdditionNode {
	return NNAdditionNode{
		NNBinaryArithmeticNode: NNBinaryArithmeticNodeFrom(ptr),
	}
}

func (nc _NNAdditionNodeClass) Alloc() NNAdditionNode {
	rv := objc.Call[NNAdditionNode](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NNAdditionNodeClass) New() NNAdditionNode {
	rv := objc.Call[NNAdditionNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNAdditionNode() NNAdditionNode {
	return NNAdditionNodeClass.New()
}

func (n_ NNAdditionNode) Init() NNAdditionNode {
	rv := objc.Call[NNAdditionNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNAdditionNodeClass) NodeWithSources(sourceNodes []INNImageNode) NNAdditionNode {
	rv := objc.Call[NNAdditionNode](nc, objc.Sel("nodeWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890829-nodewithsources?language=objc
func NNAdditionNode_NodeWithSources(sourceNodes []INNImageNode) NNAdditionNode {
	return NNAdditionNodeClass.NodeWithSources(sourceNodes)
}

func (nc _NNAdditionNodeClass) NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNAdditionNode {
	rv := objc.Call[NNAdditionNode](nc, objc.Sel("nodeWithLeftSource:rightSource:"), left, right)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890830-nodewithleftsource?language=objc
func NNAdditionNode_NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNAdditionNode {
	return NNAdditionNodeClass.NodeWithLeftSourceRightSource(left, right)
}

func (n_ NNAdditionNode) InitWithSources(sourceNodes []INNImageNode) NNAdditionNode {
	rv := objc.Call[NNAdditionNode](n_, objc.Sel("initWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890820-initwithsources?language=objc
func NewNNAdditionNodeWithSources(sourceNodes []INNImageNode) NNAdditionNode {
	instance := NNAdditionNodeClass.Alloc().InitWithSources(sourceNodes)
	instance.Autorelease()
	return instance
}

func (n_ NNAdditionNode) InitWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNAdditionNode {
	rv := objc.Call[NNAdditionNode](n_, objc.Sel("initWithLeftSource:rightSource:"), left, right)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890825-initwithleftsource?language=objc
func NewNNAdditionNodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNAdditionNode {
	instance := NNAdditionNodeClass.Alloc().InitWithLeftSourceRightSource(left, right)
	instance.Autorelease()
	return instance
}
