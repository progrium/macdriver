// Code generated by DarwinKit. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [PersistentHistoryChange] class.
var PersistentHistoryChangeClass = _PersistentHistoryChangeClass{objc.GetClass("NSPersistentHistoryChange")}

type _PersistentHistoryChangeClass struct {
	objc.Class
}

// An interface definition for the [PersistentHistoryChange] class.
type IPersistentHistoryChange interface {
	objc.IObject
	Transaction() PersistentHistoryTransaction
	UpdatedProperties() foundation.Set
	ChangedObjectID() ManagedObjectID
	Tombstone() foundation.Dictionary
	ChangeType() PersistentHistoryChangeType
	ChangeID() int64
}

// A change representing the insertion, update, or deletion of a managed object in the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange?language=objc
type PersistentHistoryChange struct {
	objc.Object
}

func PersistentHistoryChangeFrom(ptr unsafe.Pointer) PersistentHistoryChange {
	return PersistentHistoryChange{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PersistentHistoryChangeClass) Alloc() PersistentHistoryChange {
	rv := objc.Call[PersistentHistoryChange](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PersistentHistoryChangeClass) New() PersistentHistoryChange {
	rv := objc.Call[PersistentHistoryChange](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentHistoryChange() PersistentHistoryChange {
	return PersistentHistoryChangeClass.New()
}

func (p_ PersistentHistoryChange) Init() PersistentHistoryChange {
	rv := objc.Call[PersistentHistoryChange](p_, objc.Sel("init"))
	return rv
}

// Requests an entity description for the managed object type affected by the change using the provided context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/3240590-entitydescriptionwithcontext?language=objc
func (pc _PersistentHistoryChangeClass) EntityDescriptionWithContext(context IManagedObjectContext) EntityDescription {
	rv := objc.Call[EntityDescription](pc, objc.Sel("entityDescriptionWithContext:"), context)
	return rv
}

// Requests an entity description for the managed object type affected by the change using the provided context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/3240590-entitydescriptionwithcontext?language=objc
func PersistentHistoryChange_EntityDescriptionWithContext(context IManagedObjectContext) EntityDescription {
	return PersistentHistoryChangeClass.EntityDescriptionWithContext(context)
}

// The persistent history transaction containing this change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/2892567-transaction?language=objc
func (p_ PersistentHistoryChange) Transaction() PersistentHistoryTransaction {
	rv := objc.Call[PersistentHistoryTransaction](p_, objc.Sel("transaction"))
	return rv
}

// A fetch request that has the persistent history change as the entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/3240591-fetchrequest?language=objc
func (pc _PersistentHistoryChangeClass) FetchRequest() FetchRequest {
	rv := objc.Call[FetchRequest](pc, objc.Sel("fetchRequest"))
	return rv
}

// A fetch request that has the persistent history change as the entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/3240591-fetchrequest?language=objc
func PersistentHistoryChange_FetchRequest() FetchRequest {
	return PersistentHistoryChangeClass.FetchRequest()
}

// The set of properties that were updated on the managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/2892568-updatedproperties?language=objc
func (p_ PersistentHistoryChange) UpdatedProperties() foundation.Set {
	rv := objc.Call[foundation.Set](p_, objc.Sel("updatedProperties"))
	return rv
}

// The identifier of the managed object that changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/2892570-changedobjectid?language=objc
func (p_ PersistentHistoryChange) ChangedObjectID() ManagedObjectID {
	rv := objc.Call[ManagedObjectID](p_, objc.Sel("changedObjectID"))
	return rv
}

// A dictionary of attributes marked for preservation after deletion, and their values when deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/2892564-tombstone?language=objc
func (p_ PersistentHistoryChange) Tombstone() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](p_, objc.Sel("tombstone"))
	return rv
}

// The entity description of the persistent history change entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/3240589-entitydescription?language=objc
func (pc _PersistentHistoryChangeClass) EntityDescription() EntityDescription {
	rv := objc.Call[EntityDescription](pc, objc.Sel("entityDescription"))
	return rv
}

// The entity description of the persistent history change entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/3240589-entitydescription?language=objc
func PersistentHistoryChange_EntityDescription() EntityDescription {
	return PersistentHistoryChangeClass.EntityDescription()
}

// The type of change to the managed object in the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/2892563-changetype?language=objc
func (p_ PersistentHistoryChange) ChangeType() PersistentHistoryChangeType {
	rv := objc.Call[PersistentHistoryChangeType](p_, objc.Sel("changeType"))
	return rv
}

// The change’s numeric identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/2892569-changeid?language=objc
func (p_ PersistentHistoryChange) ChangeID() int64 {
	rv := objc.Call[int64](p_, objc.Sel("changeID"))
	return rv
}
