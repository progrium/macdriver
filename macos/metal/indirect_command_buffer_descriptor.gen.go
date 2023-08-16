// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [IndirectCommandBufferDescriptor] class.
var IndirectCommandBufferDescriptorClass = _IndirectCommandBufferDescriptorClass{objc.GetClass("MTLIndirectCommandBufferDescriptor")}

type _IndirectCommandBufferDescriptorClass struct {
	objc.Class
}

// An interface definition for the [IndirectCommandBufferDescriptor] class.
type IIndirectCommandBufferDescriptor interface {
	objc.IObject
	InheritBuffers() bool
	SetInheritBuffers(value bool)
	MaxVertexBufferBindCount() uint
	SetMaxVertexBufferBindCount(value uint)
	InheritPipelineState() bool
	SetInheritPipelineState(value bool)
	MaxFragmentBufferBindCount() uint
	SetMaxFragmentBufferBindCount(value uint)
	CommandTypes() IndirectCommandType
	SetCommandTypes(value IndirectCommandType)
	MaxKernelBufferBindCount() uint
	SetMaxKernelBufferBindCount(value uint)
}

// A configuration you create to customize an indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor?language=objc
type IndirectCommandBufferDescriptor struct {
	objc.Object
}

func IndirectCommandBufferDescriptorFrom(ptr unsafe.Pointer) IndirectCommandBufferDescriptor {
	return IndirectCommandBufferDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _IndirectCommandBufferDescriptorClass) Alloc() IndirectCommandBufferDescriptor {
	rv := objc.Call[IndirectCommandBufferDescriptor](ic, objc.Sel("alloc"))
	return rv
}

func IndirectCommandBufferDescriptor_Alloc() IndirectCommandBufferDescriptor {
	return IndirectCommandBufferDescriptorClass.Alloc()
}

func (ic _IndirectCommandBufferDescriptorClass) New() IndirectCommandBufferDescriptor {
	rv := objc.Call[IndirectCommandBufferDescriptor](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewIndirectCommandBufferDescriptor() IndirectCommandBufferDescriptor {
	return IndirectCommandBufferDescriptorClass.New()
}

func (i_ IndirectCommandBufferDescriptor) Init() IndirectCommandBufferDescriptor {
	rv := objc.Call[IndirectCommandBufferDescriptor](i_, objc.Sel("init"))
	return rv
}

// A Boolean value that determines where commands in the indirect command buffer get their buffer arguments from when you execute them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2967432-inheritbuffers?language=objc
func (i_ IndirectCommandBufferDescriptor) InheritBuffers() bool {
	rv := objc.Call[bool](i_, objc.Sel("inheritBuffers"))
	return rv
}

// A Boolean value that determines where commands in the indirect command buffer get their buffer arguments from when you execute them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2967432-inheritbuffers?language=objc
func (i_ IndirectCommandBufferDescriptor) SetInheritBuffers(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setInheritBuffers:"), value)
}

// The maximum number of buffers that you can set per command for the vertex stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2976451-maxvertexbufferbindcount?language=objc
func (i_ IndirectCommandBufferDescriptor) MaxVertexBufferBindCount() uint {
	rv := objc.Call[uint](i_, objc.Sel("maxVertexBufferBindCount"))
	return rv
}

// The maximum number of buffers that you can set per command for the vertex stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2976451-maxvertexbufferbindcount?language=objc
func (i_ IndirectCommandBufferDescriptor) SetMaxVertexBufferBindCount(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setMaxVertexBufferBindCount:"), value)
}

// A Boolean value that determines where commands in the indirect command buffer get their pipeline state from when you execute them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2966584-inheritpipelinestate?language=objc
func (i_ IndirectCommandBufferDescriptor) InheritPipelineState() bool {
	rv := objc.Call[bool](i_, objc.Sel("inheritPipelineState"))
	return rv
}

// A Boolean value that determines where commands in the indirect command buffer get their pipeline state from when you execute them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2966584-inheritpipelinestate?language=objc
func (i_ IndirectCommandBufferDescriptor) SetInheritPipelineState(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setInheritPipelineState:"), value)
}

// The maximum number of buffers that you can set per command for the fragment stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2976449-maxfragmentbufferbindcount?language=objc
func (i_ IndirectCommandBufferDescriptor) MaxFragmentBufferBindCount() uint {
	rv := objc.Call[uint](i_, objc.Sel("maxFragmentBufferBindCount"))
	return rv
}

// The maximum number of buffers that you can set per command for the fragment stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2976449-maxfragmentbufferbindcount?language=objc
func (i_ IndirectCommandBufferDescriptor) SetMaxFragmentBufferBindCount(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setMaxFragmentBufferBindCount:"), value)
}

// The set of command types that you can encode into the indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2966582-commandtypes?language=objc
func (i_ IndirectCommandBufferDescriptor) CommandTypes() IndirectCommandType {
	rv := objc.Call[IndirectCommandType](i_, objc.Sel("commandTypes"))
	return rv
}

// The set of command types that you can encode into the indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2966582-commandtypes?language=objc
func (i_ IndirectCommandBufferDescriptor) SetCommandTypes(value IndirectCommandType) {
	objc.Call[objc.Void](i_, objc.Sel("setCommandTypes:"), value)
}

// The maximum number of buffers that you can set per command for the compute kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2976450-maxkernelbufferbindcount?language=objc
func (i_ IndirectCommandBufferDescriptor) MaxKernelBufferBindCount() uint {
	rv := objc.Call[uint](i_, objc.Sel("maxKernelBufferBindCount"))
	return rv
}

// The maximum number of buffers that you can set per command for the compute kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandbufferdescriptor/2976450-maxkernelbufferbindcount?language=objc
func (i_ IndirectCommandBufferDescriptor) SetMaxKernelBufferBindCount(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setMaxKernelBufferBindCount:"), value)
}
