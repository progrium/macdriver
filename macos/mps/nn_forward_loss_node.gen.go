// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNForwardLossNode] class.
var NNForwardLossNodeClass = _NNForwardLossNodeClass{objc.GetClass("MPSNNForwardLossNode")}

type _NNForwardLossNodeClass struct {
	objc.Class
}

// An interface definition for the [NNForwardLossNode] class.
type INNForwardLossNode interface {
	INNFilterNode
	NumberOfClasses() uint
	Weight() float64
	PropertyCallBack() NNLossCallbackWrapper
	SetPropertyCallBack(value PNNLossCallback)
	SetPropertyCallBackObject(valueObject objc.IObject)
	Epsilon() float64
	Delta() float64
	ReductionType() CNNReductionType
	ReduceAcrossBatch() bool
	LossType() CNNLossType
	LabelSmoothing() float64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode?language=objc
type NNForwardLossNode struct {
	NNFilterNode
}

func NNForwardLossNodeFrom(ptr unsafe.Pointer) NNForwardLossNode {
	return NNForwardLossNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (nc _NNForwardLossNodeClass) NodeWithSourcesLossDescriptor(sourceNodes []INNImageNode, descriptor ICNNLossDescriptor) NNForwardLossNode {
	rv := objc.Call[NNForwardLossNode](nc, objc.Sel("nodeWithSources:lossDescriptor:"), sourceNodes, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131839-nodewithsources?language=objc
func NNForwardLossNode_NodeWithSourcesLossDescriptor(sourceNodes []INNImageNode, descriptor ICNNLossDescriptor) NNForwardLossNode {
	return NNForwardLossNodeClass.NodeWithSourcesLossDescriptor(sourceNodes, descriptor)
}

func (nc _NNForwardLossNodeClass) NodeWithSourceLabelsLossDescriptor(source INNImageNode, labels INNImageNode, descriptor ICNNLossDescriptor) NNForwardLossNode {
	rv := objc.Call[NNForwardLossNode](nc, objc.Sel("nodeWithSource:labels:lossDescriptor:"), objc.Ptr(source), objc.Ptr(labels), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131837-nodewithsource?language=objc
func NNForwardLossNode_NodeWithSourceLabelsLossDescriptor(source INNImageNode, labels INNImageNode, descriptor ICNNLossDescriptor) NNForwardLossNode {
	return NNForwardLossNodeClass.NodeWithSourceLabelsLossDescriptor(source, labels, descriptor)
}

func (n_ NNForwardLossNode) InitWithSourcesLossDescriptor(sourceNodes []INNImageNode, descriptor ICNNLossDescriptor) NNForwardLossNode {
	rv := objc.Call[NNForwardLossNode](n_, objc.Sel("initWithSources:lossDescriptor:"), sourceNodes, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131834-initwithsources?language=objc
func NewNNForwardLossNodeWithSourcesLossDescriptor(sourceNodes []INNImageNode, descriptor ICNNLossDescriptor) NNForwardLossNode {
	instance := NNForwardLossNodeClass.Alloc().InitWithSourcesLossDescriptor(sourceNodes, descriptor)
	instance.Autorelease()
	return instance
}

func (n_ NNForwardLossNode) InitWithSourceLabelsLossDescriptor(source INNImageNode, labels INNImageNode, descriptor ICNNLossDescriptor) NNForwardLossNode {
	rv := objc.Call[NNForwardLossNode](n_, objc.Sel("initWithSource:labels:lossDescriptor:"), objc.Ptr(source), objc.Ptr(labels), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131832-initwithsource?language=objc
func NewNNForwardLossNodeWithSourceLabelsLossDescriptor(source INNImageNode, labels INNImageNode, descriptor ICNNLossDescriptor) NNForwardLossNode {
	instance := NNForwardLossNodeClass.Alloc().InitWithSourceLabelsLossDescriptor(source, labels, descriptor)
	instance.Autorelease()
	return instance
}

func (nc _NNForwardLossNodeClass) Alloc() NNForwardLossNode {
	rv := objc.Call[NNForwardLossNode](nc, objc.Sel("alloc"))
	return rv
}

func NNForwardLossNode_Alloc() NNForwardLossNode {
	return NNForwardLossNodeClass.Alloc()
}

func (nc _NNForwardLossNodeClass) New() NNForwardLossNode {
	rv := objc.Call[NNForwardLossNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNForwardLossNode() NNForwardLossNode {
	return NNForwardLossNodeClass.New()
}

func (n_ NNForwardLossNode) Init() NNForwardLossNode {
	rv := objc.Call[NNForwardLossNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131840-numberofclasses?language=objc
func (n_ NNForwardLossNode) NumberOfClasses() uint {
	rv := objc.Call[uint](n_, objc.Sel("numberOfClasses"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131843-weight?language=objc
func (n_ NNForwardLossNode) Weight() float64 {
	rv := objc.Call[float64](n_, objc.Sel("weight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131841-propertycallback?language=objc
func (n_ NNForwardLossNode) PropertyCallBack() NNLossCallbackWrapper {
	rv := objc.Call[NNLossCallbackWrapper](n_, objc.Sel("propertyCallBack"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131841-propertycallback?language=objc
func (n_ NNForwardLossNode) SetPropertyCallBack(value PNNLossCallback) {
	po0 := objc.WrapAsProtocol("MPSNNLossCallback", value)
	objc.Call[objc.Void](n_, objc.Sel("setPropertyCallBack:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131841-propertycallback?language=objc
func (n_ NNForwardLossNode) SetPropertyCallBackObject(valueObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setPropertyCallBack:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131827-epsilon?language=objc
func (n_ NNForwardLossNode) Epsilon() float64 {
	rv := objc.Call[float64](n_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131826-delta?language=objc
func (n_ NNForwardLossNode) Delta() float64 {
	rv := objc.Call[float64](n_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131842-reductiontype?language=objc
func (n_ NNForwardLossNode) ReductionType() CNNReductionType {
	rv := objc.Call[CNNReductionType](n_, objc.Sel("reductionType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3547987-reduceacrossbatch?language=objc
func (n_ NNForwardLossNode) ReduceAcrossBatch() bool {
	rv := objc.Call[bool](n_, objc.Sel("reduceAcrossBatch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131836-losstype?language=objc
func (n_ NNForwardLossNode) LossType() CNNLossType {
	rv := objc.Call[CNNLossType](n_, objc.Sel("lossType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131835-labelsmoothing?language=objc
func (n_ NNForwardLossNode) LabelSmoothing() float64 {
	rv := objc.Call[float64](n_, objc.Sel("labelSmoothing"))
	return rv
}