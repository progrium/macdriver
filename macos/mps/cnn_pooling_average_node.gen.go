// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingAverageNode] class.
var CNNPoolingAverageNodeClass = _CNNPoolingAverageNodeClass{objc.GetClass("MPSCNNPoolingAverageNode")}

type _CNNPoolingAverageNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingAverageNode] class.
type ICNNPoolingAverageNode interface {
	ICNNPoolingNode
}

// A representation of an average pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragenode?language=objc
type CNNPoolingAverageNode struct {
	CNNPoolingNode
}

func CNNPoolingAverageNodeFrom(ptr unsafe.Pointer) CNNPoolingAverageNode {
	return CNNPoolingAverageNode{
		CNNPoolingNode: CNNPoolingNodeFrom(ptr),
	}
}

func (cc _CNNPoolingAverageNodeClass) Alloc() CNNPoolingAverageNode {
	rv := objc.Call[CNNPoolingAverageNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingAverageNode_Alloc() CNNPoolingAverageNode {
	return CNNPoolingAverageNodeClass.Alloc()
}

func (cc _CNNPoolingAverageNodeClass) New() CNNPoolingAverageNode {
	rv := objc.Call[CNNPoolingAverageNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingAverageNode() CNNPoolingAverageNode {
	return CNNPoolingAverageNodeClass.New()
}

func (c_ CNNPoolingAverageNode) Init() CNNPoolingAverageNode {
	rv := objc.Call[CNNPoolingAverageNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNPoolingAverageNodeClass) NodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingAverageNode {
	rv := objc.Call[CNNPoolingAverageNode](cc, objc.Sel("nodeWithSource:filterSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866508-nodewithsource?language=objc
func CNNPoolingAverageNode_NodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingAverageNode {
	return CNNPoolingAverageNodeClass.NodeWithSourceFilterSize(sourceNode, size)
}

func (c_ CNNPoolingAverageNode) InitWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingAverageNode {
	rv := objc.Call[CNNPoolingAverageNode](c_, objc.Sel("initWithSource:filterSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866488-initwithsource?language=objc
func NewCNNPoolingAverageNodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingAverageNode {
	instance := CNNPoolingAverageNodeClass.Alloc().InitWithSourceFilterSize(sourceNode, size)
	instance.Autorelease()
	return instance
}
