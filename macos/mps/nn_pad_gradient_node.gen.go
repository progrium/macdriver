// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [NNPadGradientNode] class.
var NNPadGradientNodeClass = _NNPadGradientNodeClass{objc.GetClass("MPSNNPadGradientNode")}

type _NNPadGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNPadGradientNode] class.
type INNPadGradientNode interface {
	INNGradientFilterNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadgradientnode?language=objc
type NNPadGradientNode struct {
	NNGradientFilterNode
}

func NNPadGradientNodeFrom(ptr unsafe.Pointer) NNPadGradientNode {
	return NNPadGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (n_ NNPadGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNPadGradientNode {
	rv := objc.Call[NNPadGradientNode](n_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadgradientnode/3037391-initwithsourcegradient?language=objc
func NewNNPadGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNPadGradientNode {
	instance := NNPadGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
	instance.Autorelease()
	return instance
}

func (nc _NNPadGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNPadGradientNode {
	rv := objc.Call[NNPadGradientNode](nc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadgradientnode/3037392-nodewithsourcegradient?language=objc
func NNPadGradientNode_NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNPadGradientNode {
	return NNPadGradientNodeClass.NodeWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
}

func (nc _NNPadGradientNodeClass) Alloc() NNPadGradientNode {
	rv := objc.Call[NNPadGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NNPadGradientNodeClass) New() NNPadGradientNode {
	rv := objc.Call[NNPadGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNPadGradientNode() NNPadGradientNode {
	return NNPadGradientNodeClass.New()
}

func (n_ NNPadGradientNode) Init() NNPadGradientNode {
	rv := objc.Call[NNPadGradientNode](n_, objc.Sel("init"))
	return rv
}
