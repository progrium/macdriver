// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OrderedCollectionDifference] class.
var OrderedCollectionDifferenceClass = _OrderedCollectionDifferenceClass{objc.GetClass("NSOrderedCollectionDifference")}

type _OrderedCollectionDifferenceClass struct {
	objc.Class
}

// An interface definition for the [OrderedCollectionDifference] class.
type IOrderedCollectionDifference interface {
	objc.IObject
	DifferenceByTransformingChangesWithBlock(block func(arg0 OrderedCollectionChange) OrderedCollectionChange) OrderedCollectionDifference
	Removals() []OrderedCollectionChange
	Insertions() []OrderedCollectionChange
	HasChanges() bool
}

// An object representing the difference between two ordered collections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference?language=objc
type OrderedCollectionDifference struct {
	objc.Object
}

func OrderedCollectionDifferenceFrom(ptr unsafe.Pointer) OrderedCollectionDifference {
	return OrderedCollectionDifference{
		Object: objc.ObjectFrom(ptr),
	}
}

func (o_ OrderedCollectionDifference) InitWithChanges(changes []IOrderedCollectionChange) OrderedCollectionDifference {
	rv := objc.Call[OrderedCollectionDifference](o_, objc.Sel("initWithChanges:"), changes)
	return rv
}

// Creates an ordered collection difference using an array of ordered collection changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/3152185-initwithchanges?language=objc
func NewOrderedCollectionDifferenceWithChanges(changes []IOrderedCollectionChange) OrderedCollectionDifference {
	instance := OrderedCollectionDifferenceClass.Alloc().InitWithChanges(changes)
	instance.Autorelease()
	return instance
}

func (o_ OrderedCollectionDifference) InverseDifference() OrderedCollectionDifference {
	rv := objc.Call[OrderedCollectionDifference](o_, objc.Sel("inverseDifference"))
	return rv
}

// Calculate the difference between two objects in the reverse direction of comparison. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/3180114-inversedifference?language=objc
func OrderedCollectionDifference_InverseDifference() OrderedCollectionDifference {
	instance := OrderedCollectionDifferenceClass.Alloc().InverseDifference()
	instance.Autorelease()
	return instance
}

func (o_ OrderedCollectionDifference) InitWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjects(inserts IIndexSet, insertedObjects []objc.IObject, removes IIndexSet, removedObjects []objc.IObject) OrderedCollectionDifference {
	rv := objc.Call[OrderedCollectionDifference](o_, objc.Sel("initWithInsertIndexes:insertedObjects:removeIndexes:removedObjects:"), objc.Ptr(inserts), insertedObjects, objc.Ptr(removes), removedObjects)
	return rv
}

// Creates an ordered collection difference from arrays of inserted and removed objects with corresponding sets of indices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/3152186-initwithinsertindexes?language=objc
func NewOrderedCollectionDifferenceWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjects(inserts IIndexSet, insertedObjects []objc.IObject, removes IIndexSet, removedObjects []objc.IObject) OrderedCollectionDifference {
	instance := OrderedCollectionDifferenceClass.Alloc().InitWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjects(inserts, insertedObjects, removes, removedObjects)
	instance.Autorelease()
	return instance
}

func (oc _OrderedCollectionDifferenceClass) Alloc() OrderedCollectionDifference {
	rv := objc.Call[OrderedCollectionDifference](oc, objc.Sel("alloc"))
	return rv
}

func OrderedCollectionDifference_Alloc() OrderedCollectionDifference {
	return OrderedCollectionDifferenceClass.Alloc()
}

func (oc _OrderedCollectionDifferenceClass) New() OrderedCollectionDifference {
	rv := objc.Call[OrderedCollectionDifference](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOrderedCollectionDifference() OrderedCollectionDifference {
	return OrderedCollectionDifferenceClass.New()
}

func (o_ OrderedCollectionDifference) Init() OrderedCollectionDifference {
	rv := objc.Call[OrderedCollectionDifference](o_, objc.Sel("init"))
	return rv
}

// Create a new ordered collection difference by mapping over this difference’s members, processing the change objects with the block provided. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/3152183-differencebytransformingchangesw?language=objc
func (o_ OrderedCollectionDifference) DifferenceByTransformingChangesWithBlock(block func(arg0 OrderedCollectionChange) OrderedCollectionChange) OrderedCollectionDifference {
	rv := objc.Call[OrderedCollectionDifference](o_, objc.Sel("differenceByTransformingChangesWithBlock:"), block)
	return rv
}

// A collection of removal change objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/3152189-removals?language=objc
func (o_ OrderedCollectionDifference) Removals() []OrderedCollectionChange {
	rv := objc.Call[[]OrderedCollectionChange](o_, objc.Sel("removals"))
	return rv
}

// A collection of insertion change objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/3152188-insertions?language=objc
func (o_ OrderedCollectionDifference) Insertions() []OrderedCollectionChange {
	rv := objc.Call[[]OrderedCollectionChange](o_, objc.Sel("insertions"))
	return rv
}

// A Boolean value that indicates if the difference has changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/3152184-haschanges?language=objc
func (o_ OrderedCollectionDifference) HasChanges() bool {
	rv := objc.Call[bool](o_, objc.Sel("hasChanges"))
	return rv
}
