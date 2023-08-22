// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingL2NormNode] class.
var CNNPoolingL2NormNodeClass = _CNNPoolingL2NormNodeClass{objc.GetClass("MPSCNNPoolingL2NormNode")}

type _CNNPoolingL2NormNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingL2NormNode] class.
type ICNNPoolingL2NormNode interface {
	ICNNPoolingNode
}

// A representation of a L2-norm pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingl2normnode?language=objc
type CNNPoolingL2NormNode struct {
	CNNPoolingNode
}

func CNNPoolingL2NormNodeFrom(ptr unsafe.Pointer) CNNPoolingL2NormNode {
	return CNNPoolingL2NormNode{
		CNNPoolingNode: CNNPoolingNodeFrom(ptr),
	}
}

func (cc _CNNPoolingL2NormNodeClass) Alloc() CNNPoolingL2NormNode {
	rv := objc.Call[CNNPoolingL2NormNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingL2NormNode_Alloc() CNNPoolingL2NormNode {
	return CNNPoolingL2NormNodeClass.Alloc()
}

func (cc _CNNPoolingL2NormNodeClass) New() CNNPoolingL2NormNode {
	rv := objc.Call[CNNPoolingL2NormNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingL2NormNode() CNNPoolingL2NormNode {
	return CNNPoolingL2NormNodeClass.New()
}

func (c_ CNNPoolingL2NormNode) Init() CNNPoolingL2NormNode {
	rv := objc.Call[CNNPoolingL2NormNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNPoolingL2NormNodeClass) NodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingL2NormNode {
	rv := objc.Call[CNNPoolingL2NormNode](cc, objc.Sel("nodeWithSource:filterSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866508-nodewithsource?language=objc
func CNNPoolingL2NormNode_NodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingL2NormNode {
	return CNNPoolingL2NormNodeClass.NodeWithSourceFilterSize(sourceNode, size)
}

func (c_ CNNPoolingL2NormNode) InitWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingL2NormNode {
	rv := objc.Call[CNNPoolingL2NormNode](c_, objc.Sel("initWithSource:filterSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866488-initwithsource?language=objc
func NewCNNPoolingL2NormNodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingL2NormNode {
	instance := CNNPoolingL2NormNodeClass.Alloc().InitWithSourceFilterSize(sourceNode, size)
	instance.Autorelease()
	return instance
}
