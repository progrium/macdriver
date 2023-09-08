// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BufferLayoutDescriptor] class.
var BufferLayoutDescriptorClass = _BufferLayoutDescriptorClass{objc.GetClass("MTLBufferLayoutDescriptor")}

type _BufferLayoutDescriptorClass struct {
	objc.Class
}

// An interface definition for the [BufferLayoutDescriptor] class.
type IBufferLayoutDescriptor interface {
	objc.IObject
	StepFunction() StepFunction
	SetStepFunction(value StepFunction)
	StepRate() uint
	SetStepRate(value uint)
	Stride() uint
	SetStride(value uint)
}

// An object that configures how a function fetches input data for an attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptor?language=objc
type BufferLayoutDescriptor struct {
	objc.Object
}

func BufferLayoutDescriptorFrom(ptr unsafe.Pointer) BufferLayoutDescriptor {
	return BufferLayoutDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BufferLayoutDescriptorClass) Alloc() BufferLayoutDescriptor {
	rv := objc.Call[BufferLayoutDescriptor](bc, objc.Sel("alloc"))
	return rv
}

func (bc _BufferLayoutDescriptorClass) New() BufferLayoutDescriptor {
	rv := objc.Call[BufferLayoutDescriptor](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBufferLayoutDescriptor() BufferLayoutDescriptor {
	return BufferLayoutDescriptorClass.New()
}

func (b_ BufferLayoutDescriptor) Init() BufferLayoutDescriptor {
	rv := objc.Call[BufferLayoutDescriptor](b_, objc.Sel("init"))
	return rv
}

// Determines which aspect of command execution triggers data fetches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptor/2097182-stepfunction?language=objc
func (b_ BufferLayoutDescriptor) StepFunction() StepFunction {
	rv := objc.Call[StepFunction](b_, objc.Sel("stepFunction"))
	return rv
}

// Determines which aspect of command execution triggers data fetches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptor/2097182-stepfunction?language=objc
func (b_ BufferLayoutDescriptor) SetStepFunction(value StepFunction) {
	objc.Call[objc.Void](b_, objc.Sel("setStepFunction:"), value)
}

// The rate at which data is fetched by the step function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptor/2097164-steprate?language=objc
func (b_ BufferLayoutDescriptor) StepRate() uint {
	rv := objc.Call[uint](b_, objc.Sel("stepRate"))
	return rv
}

// The rate at which data is fetched by the step function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptor/2097164-steprate?language=objc
func (b_ BufferLayoutDescriptor) SetStepRate(value uint) {
	objc.Call[objc.Void](b_, objc.Sel("setStepRate:"), value)
}

// The number of bytes between the first byte of two consecutive entries in a buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptor/2097190-stride?language=objc
func (b_ BufferLayoutDescriptor) Stride() uint {
	rv := objc.Call[uint](b_, objc.Sel("stride"))
	return rv
}

// The number of bytes between the first byte of two consecutive entries in a buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptor/2097190-stride?language=objc
func (b_ BufferLayoutDescriptor) SetStride(value uint) {
	objc.Call[objc.Void](b_, objc.Sel("setStride:"), value)
}
