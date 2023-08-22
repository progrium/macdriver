// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LSTMDescriptor] class.
var LSTMDescriptorClass = _LSTMDescriptorClass{objc.GetClass("MPSGraphLSTMDescriptor")}

type _LSTMDescriptorClass struct {
	objc.Class
}

// An interface definition for the [LSTMDescriptor] class.
type ILSTMDescriptor interface {
	objc.IObject
	Activation() RNNActivation
	SetActivation(value RNNActivation)
	Training() bool
	SetTraining(value bool)
	ProduceCell() bool
	SetProduceCell(value bool)
	CellGateActivation() RNNActivation
	SetCellGateActivation(value RNNActivation)
	OutputGateActivation() RNNActivation
	SetOutputGateActivation(value RNNActivation)
	ForgetGateLast() bool
	SetForgetGateLast(value bool)
	Bidirectional() bool
	SetBidirectional(value bool)
	ForgetGateActivation() RNNActivation
	SetForgetGateActivation(value RNNActivation)
	Reverse() bool
	SetReverse(value bool)
	InputGateActivation() RNNActivation
	SetInputGateActivation(value RNNActivation)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor?language=objc
type LSTMDescriptor struct {
	objc.Object
}

func LSTMDescriptorFrom(ptr unsafe.Pointer) LSTMDescriptor {
	return LSTMDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (lc _LSTMDescriptorClass) Descriptor() LSTMDescriptor {
	rv := objc.Call[LSTMDescriptor](lc, objc.Sel("descriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925439-descriptor?language=objc
func LSTMDescriptor_Descriptor() LSTMDescriptor {
	return LSTMDescriptorClass.Descriptor()
}

func (lc _LSTMDescriptorClass) Alloc() LSTMDescriptor {
	rv := objc.Call[LSTMDescriptor](lc, objc.Sel("alloc"))
	return rv
}

func LSTMDescriptor_Alloc() LSTMDescriptor {
	return LSTMDescriptorClass.Alloc()
}

func (lc _LSTMDescriptorClass) New() LSTMDescriptor {
	rv := objc.Call[LSTMDescriptor](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLSTMDescriptor() LSTMDescriptor {
	return LSTMDescriptorClass.New()
}

func (l_ LSTMDescriptor) Init() LSTMDescriptor {
	rv := objc.Call[LSTMDescriptor](l_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925436-activation?language=objc
func (l_ LSTMDescriptor) Activation() RNNActivation {
	rv := objc.Call[RNNActivation](l_, objc.Sel("activation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925436-activation?language=objc
func (l_ LSTMDescriptor) SetActivation(value RNNActivation) {
	objc.Call[objc.Void](l_, objc.Sel("setActivation:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925446-training?language=objc
func (l_ LSTMDescriptor) Training() bool {
	rv := objc.Call[bool](l_, objc.Sel("training"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925446-training?language=objc
func (l_ LSTMDescriptor) SetTraining(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setTraining:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925444-producecell?language=objc
func (l_ LSTMDescriptor) ProduceCell() bool {
	rv := objc.Call[bool](l_, objc.Sel("produceCell"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925444-producecell?language=objc
func (l_ LSTMDescriptor) SetProduceCell(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setProduceCell:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925438-cellgateactivation?language=objc
func (l_ LSTMDescriptor) CellGateActivation() RNNActivation {
	rv := objc.Call[RNNActivation](l_, objc.Sel("cellGateActivation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925438-cellgateactivation?language=objc
func (l_ LSTMDescriptor) SetCellGateActivation(value RNNActivation) {
	objc.Call[objc.Void](l_, objc.Sel("setCellGateActivation:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925443-outputgateactivation?language=objc
func (l_ LSTMDescriptor) OutputGateActivation() RNNActivation {
	rv := objc.Call[RNNActivation](l_, objc.Sel("outputGateActivation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925443-outputgateactivation?language=objc
func (l_ LSTMDescriptor) SetOutputGateActivation(value RNNActivation) {
	objc.Call[objc.Void](l_, objc.Sel("setOutputGateActivation:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925441-forgetgatelast?language=objc
func (l_ LSTMDescriptor) ForgetGateLast() bool {
	rv := objc.Call[bool](l_, objc.Sel("forgetGateLast"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925441-forgetgatelast?language=objc
func (l_ LSTMDescriptor) SetForgetGateLast(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setForgetGateLast:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925437-bidirectional?language=objc
func (l_ LSTMDescriptor) Bidirectional() bool {
	rv := objc.Call[bool](l_, objc.Sel("bidirectional"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925437-bidirectional?language=objc
func (l_ LSTMDescriptor) SetBidirectional(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setBidirectional:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925440-forgetgateactivation?language=objc
func (l_ LSTMDescriptor) ForgetGateActivation() RNNActivation {
	rv := objc.Call[RNNActivation](l_, objc.Sel("forgetGateActivation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925440-forgetgateactivation?language=objc
func (l_ LSTMDescriptor) SetForgetGateActivation(value RNNActivation) {
	objc.Call[objc.Void](l_, objc.Sel("setForgetGateActivation:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925445-reverse?language=objc
func (l_ LSTMDescriptor) Reverse() bool {
	rv := objc.Call[bool](l_, objc.Sel("reverse"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925445-reverse?language=objc
func (l_ LSTMDescriptor) SetReverse(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setReverse:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925442-inputgateactivation?language=objc
func (l_ LSTMDescriptor) InputGateActivation() RNNActivation {
	rv := objc.Call[RNNActivation](l_, objc.Sel("inputGateActivation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphlstmdescriptor/3925442-inputgateactivation?language=objc
func (l_ LSTMDescriptor) SetInputGateActivation(value RNNActivation) {
	objc.Call[objc.Void](l_, objc.Sel("setInputGateActivation:"), value)
}
