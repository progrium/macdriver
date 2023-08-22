// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingNode] class.
var CNNPoolingNodeClass = _CNNPoolingNodeClass{objc.GetClass("MPSCNNPoolingNode")}

type _CNNPoolingNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingNode] class.
type ICNNPoolingNode interface {
	INNFilterNode
	StrideInPixelsY() uint
	StrideInPixelsX() uint
	KernelHeight() uint
	KernelWidth() uint
}

// A representation of a MPS CNN pooling kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode?language=objc
type CNNPoolingNode struct {
	NNFilterNode
}

func CNNPoolingNodeFrom(ptr unsafe.Pointer) CNNPoolingNode {
	return CNNPoolingNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNPoolingNodeClass) NodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingNode {
	rv := objc.Call[CNNPoolingNode](cc, objc.Sel("nodeWithSource:filterSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866508-nodewithsource?language=objc
func CNNPoolingNode_NodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingNode {
	return CNNPoolingNodeClass.NodeWithSourceFilterSize(sourceNode, size)
}

func (c_ CNNPoolingNode) InitWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingNode {
	rv := objc.Call[CNNPoolingNode](c_, objc.Sel("initWithSource:filterSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2866488-initwithsource?language=objc
func NewCNNPoolingNodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNPoolingNode {
	instance := CNNPoolingNodeClass.Alloc().InitWithSourceFilterSize(sourceNode, size)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingNodeClass) Alloc() CNNPoolingNode {
	rv := objc.Call[CNNPoolingNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingNode_Alloc() CNNPoolingNode {
	return CNNPoolingNodeClass.Alloc()
}

func (cc _CNNPoolingNodeClass) New() CNNPoolingNode {
	rv := objc.Call[CNNPoolingNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingNode() CNNPoolingNode {
	return CNNPoolingNodeClass.New()
}

func (c_ CNNPoolingNode) Init() CNNPoolingNode {
	rv := objc.Call[CNNPoolingNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993004-strideinpixelsy?language=objc
func (c_ CNNPoolingNode) StrideInPixelsY() uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInPixelsY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993003-strideinpixelsx?language=objc
func (c_ CNNPoolingNode) StrideInPixelsX() uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInPixelsX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993001-kernelheight?language=objc
func (c_ CNNPoolingNode) KernelHeight() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingnode/2993002-kernelwidth?language=objc
func (c_ CNNPoolingNode) KernelWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelWidth"))
	return rv
}
