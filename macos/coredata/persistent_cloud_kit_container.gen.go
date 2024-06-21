// Code generated by DarwinKit. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/cloudkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [PersistentCloudKitContainer] class.
var PersistentCloudKitContainerClass = _PersistentCloudKitContainerClass{objc.GetClass("NSPersistentCloudKitContainer")}

type _PersistentCloudKitContainerClass struct {
	objc.Class
}

// An interface definition for the [PersistentCloudKitContainer] class.
type IPersistentCloudKitContainer interface {
	IPersistentContainer
	CanDeleteRecordForManagedObjectWithID(objectID IManagedObjectID) bool
	RecordsForManagedObjectIDs(managedObjectIDs []IManagedObjectID) foundation.Dictionary
	RecordIDForManagedObjectID(managedObjectID IManagedObjectID) cloudkit.RecordID
	PersistUpdatedShareInPersistentStoreCompletion(share cloudkit.IShare, persistentStore IPersistentStore, completion func(persistedShare cloudkit.Share, persistedShareError foundation.Error))
	RecordForManagedObjectID(managedObjectID IManagedObjectID) cloudkit.Record
	FetchParticipantsMatchingLookupInfosIntoPersistentStoreCompletion(lookupInfos []cloudkit.IUserIdentityLookupInfo, persistentStore IPersistentStore, completion func(fetchedParticipants []cloudkit.ShareParticipant, fetchError foundation.Error))
	PurgeObjectsAndRecordsInZoneWithIDInPersistentStoreCompletion(zoneID cloudkit.IRecordZoneID, persistentStore IPersistentStore, completion func(purgedZoneID cloudkit.RecordZoneID, purgeError foundation.Error))
	AcceptShareInvitationsFromMetadataIntoPersistentStoreCompletion(metadata []cloudkit.IShareMetadata, persistentStore IPersistentStore, completion func(acceptedShareMetadatas []cloudkit.ShareMetadata, acceptOperationError foundation.Error))
	FetchSharesInPersistentStoreError(persistentStore IPersistentStore, error unsafe.Pointer) []cloudkit.Share
	InitializeCloudKitSchemaWithOptionsError(options PersistentCloudKitContainerSchemaInitializationOptions, error unsafe.Pointer) bool
	RecordIDsForManagedObjectIDs(managedObjectIDs []IManagedObjectID) foundation.Dictionary
	CanUpdateRecordForManagedObjectWithID(objectID IManagedObjectID) bool
	CanModifyManagedObjectsInStore(store IPersistentStore) bool
	FetchSharesMatchingObjectIDsError(objectIDs []IManagedObjectID, error unsafe.Pointer) foundation.Dictionary
}

// A container that encapsulates the Core Data stack in your app, and mirrors select persistent stores to a CloudKit private database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer?language=objc
type PersistentCloudKitContainer struct {
	PersistentContainer
}

func PersistentCloudKitContainerFrom(ptr unsafe.Pointer) PersistentCloudKitContainer {
	return PersistentCloudKitContainer{
		PersistentContainer: PersistentContainerFrom(ptr),
	}
}

func (pc _PersistentCloudKitContainerClass) Alloc() PersistentCloudKitContainer {
	rv := objc.Call[PersistentCloudKitContainer](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PersistentCloudKitContainerClass) New() PersistentCloudKitContainer {
	rv := objc.Call[PersistentCloudKitContainer](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentCloudKitContainer() PersistentCloudKitContainer {
	return PersistentCloudKitContainerClass.New()
}

func (p_ PersistentCloudKitContainer) Init() PersistentCloudKitContainer {
	rv := objc.Call[PersistentCloudKitContainer](p_, objc.Sel("init"))
	return rv
}

func (p_ PersistentCloudKitContainer) InitWithName(name string) PersistentCloudKitContainer {
	rv := objc.Call[PersistentCloudKitContainer](p_, objc.Sel("initWithName:"), name)
	return rv
}

// Creates a container with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1640557-initwithname?language=objc
func NewPersistentCloudKitContainerWithName(name string) PersistentCloudKitContainer {
	instance := PersistentCloudKitContainerClass.Alloc().InitWithName(name)
	instance.Autorelease()
	return instance
}

func (pc _PersistentCloudKitContainerClass) PersistentContainerWithName(name string) PersistentCloudKitContainer {
	rv := objc.Call[PersistentCloudKitContainer](pc, objc.Sel("persistentContainerWithName:"), name)
	return rv
}

// Initializes a new persistent container using the provided name for the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcontainer/1646295-persistentcontainerwithname?language=objc
func PersistentCloudKitContainer_PersistentContainerWithName(name string) PersistentCloudKitContainer {
	return PersistentCloudKitContainerClass.PersistentContainerWithName(name)
}

// Returns a Boolean value that indicates whether the user can delete the managed object’s underlying CloudKit record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3649646-candeleterecordformanagedobjectw?language=objc
func (p_ PersistentCloudKitContainer) CanDeleteRecordForManagedObjectWithID(objectID IManagedObjectID) bool {
	rv := objc.Call[bool](p_, objc.Sel("canDeleteRecordForManagedObjectWithID:"), objectID)
	return rv
}

// Returns a dictionary that contains the CloudKit records for the specified managed object IDs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3141671-recordsformanagedobjectids?language=objc
func (p_ PersistentCloudKitContainer) RecordsForManagedObjectIDs(managedObjectIDs []IManagedObjectID) foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](p_, objc.Sel("recordsForManagedObjectIDs:"), managedObjectIDs)
	return rv
}

// Returns the CloudKit record ID for the specified managed object ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3141669-recordidformanagedobjectid?language=objc
func (p_ PersistentCloudKitContainer) RecordIDForManagedObjectID(managedObjectID IManagedObjectID) cloudkit.RecordID {
	rv := objc.Call[cloudkit.RecordID](p_, objc.Sel("recordIDForManagedObjectID:"), managedObjectID)
	return rv
}

// Saves the share record and schedules it for export to iCloud. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3746832-persistupdatedshare?language=objc
func (p_ PersistentCloudKitContainer) PersistUpdatedShareInPersistentStoreCompletion(share cloudkit.IShare, persistentStore IPersistentStore, completion func(persistedShare cloudkit.Share, persistedShareError foundation.Error)) {
	objc.Call[objc.Void](p_, objc.Sel("persistUpdatedShare:inPersistentStore:completion:"), share, persistentStore, completion)
}

// Returns the CloudKit record for the specified managed object ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3141668-recordformanagedobjectid?language=objc
func (p_ PersistentCloudKitContainer) RecordForManagedObjectID(managedObjectID IManagedObjectID) cloudkit.Record {
	rv := objc.Call[cloudkit.Record](p_, objc.Sel("recordForManagedObjectID:"), managedObjectID)
	return rv
}

// Fetches all participants that match the specified critieria. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3746829-fetchparticipantsmatchinglookupi?language=objc
func (p_ PersistentCloudKitContainer) FetchParticipantsMatchingLookupInfosIntoPersistentStoreCompletion(lookupInfos []cloudkit.IUserIdentityLookupInfo, persistentStore IPersistentStore, completion func(fetchedParticipants []cloudkit.ShareParticipant, fetchError foundation.Error)) {
	objc.Call[objc.Void](p_, objc.Sel("fetchParticipantsMatchingLookupInfos:intoPersistentStore:completion:"), lookupInfos, persistentStore, completion)
}

// Deletes all CloudKit records in the specified record zone, along with their corresponding managed objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3746833-purgeobjectsandrecordsinzonewith?language=objc
func (p_ PersistentCloudKitContainer) PurgeObjectsAndRecordsInZoneWithIDInPersistentStoreCompletion(zoneID cloudkit.IRecordZoneID, persistentStore IPersistentStore, completion func(purgedZoneID cloudkit.RecordZoneID, purgeError foundation.Error)) {
	objc.Call[objc.Void](p_, objc.Sel("purgeObjectsAndRecordsInZoneWithID:inPersistentStore:completion:"), zoneID, persistentStore, completion)
}

// Accepts one or more invitations to participate in sharing using the specified metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3746828-acceptshareinvitationsfrommetada?language=objc
func (p_ PersistentCloudKitContainer) AcceptShareInvitationsFromMetadataIntoPersistentStoreCompletion(metadata []cloudkit.IShareMetadata, persistentStore IPersistentStore, completion func(acceptedShareMetadatas []cloudkit.ShareMetadata, acceptOperationError foundation.Error)) {
	objc.Call[objc.Void](p_, objc.Sel("acceptShareInvitationsFromMetadata:intoPersistentStore:completion:"), metadata, persistentStore, completion)
}

// Returns an array that contains all share records in the specified persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3746830-fetchsharesinpersistentstore?language=objc
func (p_ PersistentCloudKitContainer) FetchSharesInPersistentStoreError(persistentStore IPersistentStore, error unsafe.Pointer) []cloudkit.Share {
	rv := objc.Call[[]cloudkit.Share](p_, objc.Sel("fetchSharesInPersistentStore:error:"), persistentStore, error)
	return rv
}

// Creates the CloudKit schema for all stores in the container that manage a CloudKit database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3343548-initializecloudkitschemawithopti?language=objc
func (p_ PersistentCloudKitContainer) InitializeCloudKitSchemaWithOptionsError(options PersistentCloudKitContainerSchemaInitializationOptions, error unsafe.Pointer) bool {
	rv := objc.Call[bool](p_, objc.Sel("initializeCloudKitSchemaWithOptions:error:"), options, error)
	return rv
}

// Returns a dictionary that contains the CloudKit record IDs for the specified managed object IDs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3141670-recordidsformanagedobjectids?language=objc
func (p_ PersistentCloudKitContainer) RecordIDsForManagedObjectIDs(managedObjectIDs []IManagedObjectID) foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](p_, objc.Sel("recordIDsForManagedObjectIDs:"), managedObjectIDs)
	return rv
}

// Returns a Boolean value that indicates whether the user can modify the managed object’s underlying CloudKit record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3649648-canupdaterecordformanagedobjectw?language=objc
func (p_ PersistentCloudKitContainer) CanUpdateRecordForManagedObjectWithID(objectID IManagedObjectID) bool {
	rv := objc.Call[bool](p_, objc.Sel("canUpdateRecordForManagedObjectWithID:"), objectID)
	return rv
}

// Returns a Boolean value that indicates whether the user can modify the specified persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3649647-canmodifymanagedobjectsinstore?language=objc
func (p_ PersistentCloudKitContainer) CanModifyManagedObjectsInStore(store IPersistentStore) bool {
	rv := objc.Call[bool](p_, objc.Sel("canModifyManagedObjectsInStore:"), store)
	return rv
}

// Returns a dictionary that contains the share records that CloudKit associates with specified managed object IDs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/3746831-fetchsharesmatchingobjectids?language=objc
func (p_ PersistentCloudKitContainer) FetchSharesMatchingObjectIDsError(objectIDs []IManagedObjectID, error unsafe.Pointer) foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](p_, objc.Sel("fetchSharesMatchingObjectIDs:error:"), objectIDs, error)
	return rv
}
