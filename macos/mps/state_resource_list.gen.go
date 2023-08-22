// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StateResourceList] class.
var StateResourceListClass = _StateResourceListClass{objc.GetClass("MPSStateResourceList")}

type _StateResourceListClass struct {
	objc.Class
}

// An interface definition for the [StateResourceList] class.
type IStateResourceList interface {
	objc.IObject
	AppendTexture(descriptor metal.ITextureDescriptor)
	AppendBuffer(size uint)
}

// An interface for objects that define resources for Metal Performance Shaders state containers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcelist?language=objc
type StateResourceList struct {
	objc.Object
}

func StateResourceListFrom(ptr unsafe.Pointer) StateResourceList {
	return StateResourceList{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _StateResourceListClass) ResourceList() StateResourceList {
	rv := objc.Call[StateResourceList](sc, objc.Sel("resourceList"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcelist/2947904-resourcelist?language=objc
func StateResourceList_ResourceList() StateResourceList {
	return StateResourceListClass.ResourceList()
}

func (sc _StateResourceListClass) ResourceListWithTextureDescriptors(d metal.ITextureDescriptor, args ...any) StateResourceList {
	rv := objc.Call[StateResourceList](sc, objc.Sel("resourceListWithTextureDescriptors:"), append([]any{objc.Ptr(d)}, args...)...)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcelist/2947890-resourcelistwithtexturedescripto?language=objc
func StateResourceList_ResourceListWithTextureDescriptors(d metal.ITextureDescriptor, args ...any) StateResourceList {
	return StateResourceListClass.ResourceListWithTextureDescriptors(d, args...)
}

func (s_ StateResourceList) Init() StateResourceList {
	rv := objc.Call[StateResourceList](s_, objc.Sel("init"))
	return rv
}

func (sc _StateResourceListClass) ResourceListWithBufferSizes(firstSize uint, args ...any) StateResourceList {
	rv := objc.Call[StateResourceList](sc, objc.Sel("resourceListWithBufferSizes:"), append([]any{firstSize}, args...)...)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcelist/2947903-resourcelistwithbuffersizes?language=objc
func StateResourceList_ResourceListWithBufferSizes(firstSize uint, args ...any) StateResourceList {
	return StateResourceListClass.ResourceListWithBufferSizes(firstSize, args...)
}

func (sc _StateResourceListClass) Alloc() StateResourceList {
	rv := objc.Call[StateResourceList](sc, objc.Sel("alloc"))
	return rv
}

func StateResourceList_Alloc() StateResourceList {
	return StateResourceListClass.Alloc()
}

func (sc _StateResourceListClass) New() StateResourceList {
	rv := objc.Call[StateResourceList](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStateResourceList() StateResourceList {
	return StateResourceListClass.New()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcelist/2947894-appendtexture?language=objc
func (s_ StateResourceList) AppendTexture(descriptor metal.ITextureDescriptor) {
	objc.Call[objc.Void](s_, objc.Sel("appendTexture:"), objc.Ptr(descriptor))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcelist/2947905-appendbuffer?language=objc
func (s_ StateResourceList) AppendBuffer(size uint) {
	objc.Call[objc.Void](s_, objc.Sel("appendBuffer:"), size)
}
