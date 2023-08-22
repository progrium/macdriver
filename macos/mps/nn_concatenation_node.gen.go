// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNConcatenationNode] class.
var NNConcatenationNodeClass = _NNConcatenationNodeClass{objc.GetClass("MPSNNConcatenationNode")}

type _NNConcatenationNodeClass struct {
	objc.Class
}

// An interface definition for the [NNConcatenationNode] class.
type INNConcatenationNode interface {
	INNFilterNode
}

// A representation of the results from one or more kernels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationnode?language=objc
type NNConcatenationNode struct {
	NNFilterNode
}

func NNConcatenationNodeFrom(ptr unsafe.Pointer) NNConcatenationNode {
	return NNConcatenationNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (nc _NNConcatenationNodeClass) NodeWithSources(sourceNodes []INNImageNode) NNConcatenationNode {
	rv := objc.Call[NNConcatenationNode](nc, objc.Sel("nodeWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationnode/2866432-nodewithsources?language=objc
func NNConcatenationNode_NodeWithSources(sourceNodes []INNImageNode) NNConcatenationNode {
	return NNConcatenationNodeClass.NodeWithSources(sourceNodes)
}

func (n_ NNConcatenationNode) InitWithSources(sourceNodes []INNImageNode) NNConcatenationNode {
	rv := objc.Call[NNConcatenationNode](n_, objc.Sel("initWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationnode/2866423-initwithsources?language=objc
func NewNNConcatenationNodeWithSources(sourceNodes []INNImageNode) NNConcatenationNode {
	instance := NNConcatenationNodeClass.Alloc().InitWithSources(sourceNodes)
	instance.Autorelease()
	return instance
}

func (nc _NNConcatenationNodeClass) Alloc() NNConcatenationNode {
	rv := objc.Call[NNConcatenationNode](nc, objc.Sel("alloc"))
	return rv
}

func NNConcatenationNode_Alloc() NNConcatenationNode {
	return NNConcatenationNodeClass.Alloc()
}

func (nc _NNConcatenationNodeClass) New() NNConcatenationNode {
	rv := objc.Call[NNConcatenationNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNConcatenationNode() NNConcatenationNode {
	return NNConcatenationNodeClass.New()
}

func (n_ NNConcatenationNode) Init() NNConcatenationNode {
	rv := objc.Call[NNConcatenationNode](n_, objc.Sel("init"))
	return rv
}
