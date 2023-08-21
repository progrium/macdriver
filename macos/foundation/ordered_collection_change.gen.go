// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OrderedCollectionChange] class.
var OrderedCollectionChangeClass = _OrderedCollectionChangeClass{objc.GetClass("NSOrderedCollectionChange")}

type _OrderedCollectionChangeClass struct {
	objc.Class
}

// An interface definition for the [OrderedCollectionChange] class.
type IOrderedCollectionChange interface {
	objc.IObject
	AssociatedIndex() uint
	Object_() objc.Object
	Index() uint
	ChangeType() CollectionChangeType
}

// An object that represents an indexed change within an ordered collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange?language=objc
type OrderedCollectionChange struct {
	objc.Object
}

func OrderedCollectionChangeFrom(ptr unsafe.Pointer) OrderedCollectionChange {
	return OrderedCollectionChange{
		Object: objc.ObjectFrom(ptr),
	}
}

func (o_ OrderedCollectionChange) InitWithObjectTypeIndexAssociatedIndex(anObject objc.IObject, type_ CollectionChangeType, index uint, associatedIndex uint) OrderedCollectionChange {
	rv := objc.Call[OrderedCollectionChange](o_, objc.Sel("initWithObject:type:index:associatedIndex:"), objc.Ptr(anObject), type_, index, associatedIndex)
	return rv
}

// Creates a change object that represents inserting, removing, or moving an object from an ordered collection at a specific index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/3152180-initwithobject?language=objc
func NewOrderedCollectionChangeWithObjectTypeIndexAssociatedIndex(anObject objc.IObject, type_ CollectionChangeType, index uint, associatedIndex uint) OrderedCollectionChange {
	instance := OrderedCollectionChangeClass.Alloc().InitWithObjectTypeIndexAssociatedIndex(anObject, type_, index, associatedIndex)
	instance.Autorelease()
	return instance
}

func (oc _OrderedCollectionChangeClass) Alloc() OrderedCollectionChange {
	rv := objc.Call[OrderedCollectionChange](oc, objc.Sel("alloc"))
	return rv
}

func OrderedCollectionChange_Alloc() OrderedCollectionChange {
	return OrderedCollectionChangeClass.Alloc()
}

func (oc _OrderedCollectionChangeClass) New() OrderedCollectionChange {
	rv := objc.Call[OrderedCollectionChange](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOrderedCollectionChange() OrderedCollectionChange {
	return OrderedCollectionChangeClass.New()
}

func (o_ OrderedCollectionChange) Init() OrderedCollectionChange {
	rv := objc.Call[OrderedCollectionChange](o_, objc.Sel("init"))
	return rv
}

// Creates an change object that represents inserting or removing an object from an ordered collection at a specific index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/3152176-changewithobject?language=objc
func (oc _OrderedCollectionChangeClass) ChangeWithObjectTypeIndex(anObject objc.IObject, type_ CollectionChangeType, index uint) OrderedCollectionChange {
	rv := objc.Call[OrderedCollectionChange](oc, objc.Sel("changeWithObject:type:index:"), objc.Ptr(anObject), type_, index)
	return rv
}

// Creates an change object that represents inserting or removing an object from an ordered collection at a specific index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/3152176-changewithobject?language=objc
func OrderedCollectionChange_ChangeWithObjectTypeIndex(anObject objc.IObject, type_ CollectionChangeType, index uint) OrderedCollectionChange {
	return OrderedCollectionChangeClass.ChangeWithObjectTypeIndex(anObject, type_, index)
}

// When this property is set to a value other than [foundation/8f9], the receiver is one half of a move, and this value is the index of the change’s counterpart of the opposite type in the diff. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/3152174-associatedindex?language=objc
func (o_ OrderedCollectionChange) AssociatedIndex() uint {
	rv := objc.Call[uint](o_, objc.Sel("associatedIndex"))
	return rv
}

// An object the change inserts or removes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/3152181-object?language=objc
func (o_ OrderedCollectionChange) Object_() objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("object"))
	return rv
}

// The index location of the change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/3152178-index?language=objc
func (o_ OrderedCollectionChange) Index() uint {
	rv := objc.Call[uint](o_, objc.Sel("index"))
	return rv
}

// The type of change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/3152175-changetype?language=objc
func (o_ OrderedCollectionChange) ChangeType() CollectionChangeType {
	rv := objc.Call[CollectionChangeType](o_, objc.Sel("changeType"))
	return rv
}
