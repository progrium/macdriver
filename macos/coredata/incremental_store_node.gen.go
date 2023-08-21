// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [IncrementalStoreNode] class.
var IncrementalStoreNodeClass = _IncrementalStoreNodeClass{objc.GetClass("NSIncrementalStoreNode")}

type _IncrementalStoreNodeClass struct {
	objc.Class
}

// An interface definition for the [IncrementalStoreNode] class.
type IIncrementalStoreNode interface {
	objc.IObject
	ValueForPropertyDescription(prop IPropertyDescription) objc.Object
	UpdateWithValuesVersion(values map[string]objc.IObject, version uint64)
	Version() uint64
	ObjectID() ManagedObjectID
}

// A concrete class used to represent basic nodes in a Core Data incremental store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstorenode?language=objc
type IncrementalStoreNode struct {
	objc.Object
}

func IncrementalStoreNodeFrom(ptr unsafe.Pointer) IncrementalStoreNode {
	return IncrementalStoreNode{
		Object: objc.ObjectFrom(ptr),
	}
}

func (i_ IncrementalStoreNode) InitWithObjectIDWithValuesVersion(objectID IManagedObjectID, values map[string]objc.IObject, version uint64) IncrementalStoreNode {
	rv := objc.Call[IncrementalStoreNode](i_, objc.Sel("initWithObjectID:withValues:version:"), objc.Ptr(objectID), values, version)
	return rv
}

// Returns an object initialized with the given values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstorenode/1506321-initwithobjectid?language=objc
func NewIncrementalStoreNodeWithObjectIDWithValuesVersion(objectID IManagedObjectID, values map[string]objc.IObject, version uint64) IncrementalStoreNode {
	instance := IncrementalStoreNodeClass.Alloc().InitWithObjectIDWithValuesVersion(objectID, values, version)
	instance.Autorelease()
	return instance
}

func (ic _IncrementalStoreNodeClass) Alloc() IncrementalStoreNode {
	rv := objc.Call[IncrementalStoreNode](ic, objc.Sel("alloc"))
	return rv
}

func IncrementalStoreNode_Alloc() IncrementalStoreNode {
	return IncrementalStoreNodeClass.Alloc()
}

func (ic _IncrementalStoreNodeClass) New() IncrementalStoreNode {
	rv := objc.Call[IncrementalStoreNode](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewIncrementalStoreNode() IncrementalStoreNode {
	return IncrementalStoreNodeClass.New()
}

func (i_ IncrementalStoreNode) Init() IncrementalStoreNode {
	rv := objc.Call[IncrementalStoreNode](i_, objc.Sel("init"))
	return rv
}

// Returns the value for the given property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstorenode/1506442-valueforpropertydescription?language=objc
func (i_ IncrementalStoreNode) ValueForPropertyDescription(prop IPropertyDescription) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("valueForPropertyDescription:"), objc.Ptr(prop))
	return rv
}

// Update the values and version to reflect new data being saved to or loaded from the external store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstorenode/1506721-updatewithvalues?language=objc
func (i_ IncrementalStoreNode) UpdateWithValuesVersion(values map[string]objc.IObject, version uint64) {
	objc.Call[objc.Void](i_, objc.Sel("updateWithValues:version:"), values, version)
}

// The version of data in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstorenode/1506769-version?language=objc
func (i_ IncrementalStoreNode) Version() uint64 {
	rv := objc.Call[uint64](i_, objc.Sel("version"))
	return rv
}

// The object ID that identifies the data stored by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstorenode/1506827-objectid?language=objc
func (i_ IncrementalStoreNode) ObjectID() ManagedObjectID {
	rv := objc.Call[ManagedObjectID](i_, objc.Sel("objectID"))
	return rv
}
