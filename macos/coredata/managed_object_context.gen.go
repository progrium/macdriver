// Code generated by DarwinKit. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ManagedObjectContext] class.
var ManagedObjectContextClass = _ManagedObjectContextClass{objc.GetClass("NSManagedObjectContext")}

type _ManagedObjectContextClass struct {
	objc.Class
}

// An interface definition for the [ManagedObjectContext] class.
type IManagedObjectContext interface {
	objc.IObject
	AssignObjectToPersistentStore(object objc.IObject, store IPersistentStore)
	Rollback()
	ObserveValueForKeyPathOfObjectChangeContext(keyPath string, object objc.IObject, change map[string]objc.IObject, context unsafe.Pointer)
	ExecuteRequestError(request IPersistentStoreRequest, error foundation.IError) PersistentStoreResult
	RefreshObjectMergeChanges(object IManagedObject, flag bool)
	ObjectWithID(objectID IManagedObjectID) ManagedObject
	SetQueryGenerationFromTokenError(generation IQueryGenerationToken, error foundation.IError) bool
	DetectConflictsForObject(object IManagedObject)
	CountForFetchRequestError(request IFetchRequest, error foundation.IError) uint
	ExecuteFetchRequestError(request IFetchRequest, error foundation.IError) []objc.Object
	ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty(fault IManagedObject, oid IManagedObjectID, property IPropertyDescription) bool
	RefreshAllObjects()
	DeleteObject(object IManagedObject)
	ExistingObjectWithIDError(objectID IManagedObjectID, error foundation.IError) ManagedObject
	ObtainPermanentIDsForObjectsError(objects []IManagedObject, error foundation.IError) bool
	Undo()
	InsertObject(object IManagedObject)
	Save(error foundation.IError) bool
	Reset()
	ObjectRegisteredForID(objectID IManagedObjectID) ManagedObject
	MergeChangesFromContextDidSaveNotification(notification foundation.INotification)
	PerformBlock(block func())
	PerformBlockAndWait(block func())
	Redo()
	ProcessPendingChanges()
	HasChanges() bool
	ConcurrencyType() ManagedObjectContextConcurrencyType
	RegisteredObjects() foundation.Set
	RetainsRegisteredObjects() bool
	SetRetainsRegisteredObjects(value bool)
	MergePolicy() objc.Object
	SetMergePolicy(value objc.IObject)
	TransactionAuthor() string
	SetTransactionAuthor(value string)
	StalenessInterval() foundation.TimeInterval
	SetStalenessInterval(value foundation.TimeInterval)
	QueryGenerationToken() QueryGenerationToken
	InsertedObjects() foundation.Set
	PersistentStoreCoordinator() PersistentStoreCoordinator
	SetPersistentStoreCoordinator(value IPersistentStoreCoordinator)
	UpdatedObjects() foundation.Set
	Name() string
	SetName(value string)
	UserInfo() foundation.MutableDictionary
	ShouldDeleteInaccessibleFaults() bool
	SetShouldDeleteInaccessibleFaults(value bool)
	PropagatesDeletesAtEndOfEvent() bool
	SetPropagatesDeletesAtEndOfEvent(value bool)
	UndoManager() foundation.UndoManager
	SetUndoManager(value foundation.IUndoManager)
	ParentContext() ManagedObjectContext
	SetParentContext(value IManagedObjectContext)
	AutomaticallyMergesChangesFromParent() bool
	SetAutomaticallyMergesChangesFromParent(value bool)
	DeletedObjects() foundation.Set
}

// An object space to manipulate and track changes to managed objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext?language=objc
type ManagedObjectContext struct {
	objc.Object
}

func ManagedObjectContextFrom(ptr unsafe.Pointer) ManagedObjectContext {
	return ManagedObjectContext{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _ManagedObjectContextClass) Alloc() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](mc, objc.Sel("alloc"))
	return rv
}

func (mc _ManagedObjectContextClass) New() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewManagedObjectContext() ManagedObjectContext {
	return ManagedObjectContextClass.New()
}

func (m_ ManagedObjectContext) Init() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](m_, objc.Sel("init"))
	return rv
}

// Specifies the store in which a newly inserted object will be saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506436-assignobject?language=objc
func (m_ ManagedObjectContext) AssignObjectToPersistentStore(object objc.IObject, store IPersistentStore) {
	objc.Call[objc.Void](m_, objc.Sel("assignObject:toPersistentStore:"), object, objc.Ptr(store))
}

// Removes everything from the undo stack, discards all insertions and deletions, and restores updated objects to their last committed values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506942-rollback?language=objc
func (m_ ManagedObjectContext) Rollback() {
	objc.Call[objc.Void](m_, objc.Sel("rollback"))
}

// Allows a context that has registered as an observer of a value to be notified of a change to that value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506849-observevalueforkeypath?language=objc
func (m_ ManagedObjectContext) ObserveValueForKeyPathOfObjectChangeContext(keyPath string, object objc.IObject, change map[string]objc.IObject, context unsafe.Pointer) {
	objc.Call[objc.Void](m_, objc.Sel("observeValueForKeyPath:ofObject:change:context:"), keyPath, object, change, context)
}

// Passes a request to the persistent store without affecting the contents of the managed object context, and returns a persistent store result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506834-executerequest?language=objc
func (m_ ManagedObjectContext) ExecuteRequestError(request IPersistentStoreRequest, error foundation.IError) PersistentStoreResult {
	rv := objc.Call[PersistentStoreResult](m_, objc.Sel("executeRequest:error:"), objc.Ptr(request), objc.Ptr(error))
	return rv
}

// Updates the persistent properties of a managed object to use the latest values from the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506224-refreshobject?language=objc
func (m_ ManagedObjectContext) RefreshObjectMergeChanges(object IManagedObject, flag bool) {
	objc.Call[objc.Void](m_, objc.Sel("refreshObject:mergeChanges:"), objc.Ptr(object), flag)
}

// Returns either an existing object from the context or a fault that represents that object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506197-objectwithid?language=objc
func (m_ ManagedObjectContext) ObjectWithID(objectID IManagedObjectID) ManagedObject {
	rv := objc.Call[ManagedObject](m_, objc.Sel("objectWithID:"), objc.Ptr(objectID))
	return rv
}

// Handles changes from other processes or from a serialized state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506546-mergechangesfromremotecontextsav?language=objc
func (mc _ManagedObjectContextClass) MergeChangesFromRemoteContextSaveIntoContexts(changeNotificationData foundation.Dictionary, contexts []IManagedObjectContext) {
	objc.Call[objc.Void](mc, objc.Sel("mergeChangesFromRemoteContextSave:intoContexts:"), changeNotificationData, contexts)
}

// Handles changes from other processes or from a serialized state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506546-mergechangesfromremotecontextsav?language=objc
func ManagedObjectContext_MergeChangesFromRemoteContextSaveIntoContexts(changeNotificationData foundation.Dictionary, contexts []IManagedObjectContext) {
	ManagedObjectContextClass.MergeChangesFromRemoteContextSaveIntoContexts(changeNotificationData, contexts)
}

// Sets the query generation this context should use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1640469-setquerygenerationfromtoken?language=objc
func (m_ ManagedObjectContext) SetQueryGenerationFromTokenError(generation IQueryGenerationToken, error foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("setQueryGenerationFromToken:error:"), objc.Ptr(generation), objc.Ptr(error))
	return rv
}

// Marks an object for conflict detection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506843-detectconflictsforobject?language=objc
func (m_ ManagedObjectContext) DetectConflictsForObject(object IManagedObject) {
	objc.Call[objc.Void](m_, objc.Sel("detectConflictsForObject:"), objc.Ptr(object))
}

// Returns the number of objects the specified request fetches when it executes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506868-countforfetchrequest?language=objc
func (m_ ManagedObjectContext) CountForFetchRequestError(request IFetchRequest, error foundation.IError) uint {
	rv := objc.Call[uint](m_, objc.Sel("countForFetchRequest:error:"), objc.Ptr(request), objc.Ptr(error))
	return rv
}

// Returns an array of objects that meet the criteria of the specified fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506672-executefetchrequest?language=objc
func (m_ ManagedObjectContext) ExecuteFetchRequestError(request IFetchRequest, error foundation.IError) []objc.Object {
	rv := objc.Call[[]objc.Object](m_, objc.Sel("executeFetchRequest:error:"), objc.Ptr(request), objc.Ptr(error))
	return rv
}

// Creates a log of the inaccessible fault. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506810-shouldhandleinaccessiblefault?language=objc
func (m_ ManagedObjectContext) ShouldHandleInaccessibleFaultForObjectIDTriggeredByProperty(fault IManagedObject, oid IManagedObjectID, property IPropertyDescription) bool {
	rv := objc.Call[bool](m_, objc.Sel("shouldHandleInaccessibleFault:forObjectID:triggeredByProperty:"), objc.Ptr(fault), objc.Ptr(oid), objc.Ptr(property))
	return rv
}

// Refreshes all of the registered managed objects in the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506217-refreshallobjects?language=objc
func (m_ ManagedObjectContext) RefreshAllObjects() {
	objc.Call[objc.Void](m_, objc.Sel("refreshAllObjects"))
}

// Specifies an object that should be removed from its persistent store when changes are committed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506847-deleteobject?language=objc
func (m_ ManagedObjectContext) DeleteObject(object IManagedObject) {
	objc.Call[objc.Void](m_, objc.Sel("deleteObject:"), objc.Ptr(object))
}

// Returns an existing object from either the context or the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506686-existingobjectwithid?language=objc
func (m_ ManagedObjectContext) ExistingObjectWithIDError(objectID IManagedObjectID, error foundation.IError) ManagedObject {
	rv := objc.Call[ManagedObject](m_, objc.Sel("existingObjectWithID:error:"), objc.Ptr(objectID), objc.Ptr(error))
	return rv
}

// Converts to permanent IDs the object IDs of the objects in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506793-obtainpermanentidsforobjects?language=objc
func (m_ ManagedObjectContext) ObtainPermanentIDsForObjectsError(objects []IManagedObject, error foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("obtainPermanentIDsForObjects:error:"), objects, objc.Ptr(error))
	return rv
}

// Sends an undo message to the context’s undo manager, asking it to reverse the latest uncommitted changes applied to objects in the object graph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506969-undo?language=objc
func (m_ ManagedObjectContext) Undo() {
	objc.Call[objc.Void](m_, objc.Sel("undo"))
}

// Registers an object to be inserted in the context’s persistent store the next time changes are saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506794-insertobject?language=objc
func (m_ ManagedObjectContext) InsertObject(object IManagedObject) {
	objc.Call[objc.Void](m_, objc.Sel("insertObject:"), objc.Ptr(object))
}

// Attempts to commit unsaved changes to registered objects to the context’s parent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506866-save?language=objc
func (m_ ManagedObjectContext) Save(error foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("save:"), objc.Ptr(error))
	return rv
}

// Returns the context to its base state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506807-reset?language=objc
func (m_ ManagedObjectContext) Reset() {
	objc.Call[objc.Void](m_, objc.Sel("reset"))
}

// Returns an object that exists in the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506789-objectregisteredforid?language=objc
func (m_ ManagedObjectContext) ObjectRegisteredForID(objectID IManagedObjectID) ManagedObject {
	rv := objc.Call[ManagedObject](m_, objc.Sel("objectRegisteredForID:"), objc.Ptr(objectID))
	return rv
}

// Merges the changes specified in a given notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506606-mergechangesfromcontextdidsaveno?language=objc
func (m_ ManagedObjectContext) MergeChangesFromContextDidSaveNotification(notification foundation.INotification) {
	objc.Call[objc.Void](m_, objc.Sel("mergeChangesFromContextDidSaveNotification:"), objc.Ptr(notification))
}

// Asynchronously performs the specified block on the context’s queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506578-performblock?language=objc
func (m_ ManagedObjectContext) PerformBlock(block func()) {
	objc.Call[objc.Void](m_, objc.Sel("performBlock:"), block)
}

// Synchronously performs the specified block on the context’s queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506364-performblockandwait?language=objc
func (m_ ManagedObjectContext) PerformBlockAndWait(block func()) {
	objc.Call[objc.Void](m_, objc.Sel("performBlockAndWait:"), block)
}

// Sends a redo message to the context’s undo manager, asking it to reverse the latest undo operation applied to objects in the object graph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506688-redo?language=objc
func (m_ ManagedObjectContext) Redo() {
	objc.Call[objc.Void](m_, objc.Sel("redo"))
}

// Forces the context to process changes to the object graph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506661-processpendingchanges?language=objc
func (m_ ManagedObjectContext) ProcessPendingChanges() {
	objc.Call[objc.Void](m_, objc.Sel("processPendingChanges"))
}

// A Boolean value that indicates whether the context has uncommitted changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506954-haschanges?language=objc
func (m_ ManagedObjectContext) HasChanges() bool {
	rv := objc.Call[bool](m_, objc.Sel("hasChanges"))
	return rv
}

// The concurrency type for the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506792-concurrencytype?language=objc
func (m_ ManagedObjectContext) ConcurrencyType() ManagedObjectContextConcurrencyType {
	rv := objc.Call[ManagedObjectContextConcurrencyType](m_, objc.Sel("concurrencyType"))
	return rv
}

// The set of registered managed objects in the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506493-registeredobjects?language=objc
func (m_ ManagedObjectContext) RegisteredObjects() foundation.Set {
	rv := objc.Call[foundation.Set](m_, objc.Sel("registeredObjects"))
	return rv
}

// A Boolean value that indicates whether the context keeps strong references to all registered managed objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506290-retainsregisteredobjects?language=objc
func (m_ ManagedObjectContext) RetainsRegisteredObjects() bool {
	rv := objc.Call[bool](m_, objc.Sel("retainsRegisteredObjects"))
	return rv
}

// A Boolean value that indicates whether the context keeps strong references to all registered managed objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506290-retainsregisteredobjects?language=objc
func (m_ ManagedObjectContext) SetRetainsRegisteredObjects(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setRetainsRegisteredObjects:"), value)
}

// The merge policy of the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506490-mergepolicy?language=objc
func (m_ ManagedObjectContext) MergePolicy() objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("mergePolicy"))
	return rv
}

// The merge policy of the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506490-mergepolicy?language=objc
func (m_ ManagedObjectContext) SetMergePolicy(value objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setMergePolicy:"), value)
}

// The author for the context that is used as an identifier in persistent history transactions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/2892348-transactionauthor?language=objc
func (m_ ManagedObjectContext) TransactionAuthor() string {
	rv := objc.Call[string](m_, objc.Sel("transactionAuthor"))
	return rv
}

// The author for the context that is used as an identifier in persistent history transactions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/2892348-transactionauthor?language=objc
func (m_ ManagedObjectContext) SetTransactionAuthor(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setTransactionAuthor:"), value)
}

// The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506875-stalenessinterval?language=objc
func (m_ ManagedObjectContext) StalenessInterval() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](m_, objc.Sel("stalenessInterval"))
	return rv
}

// The maximum length of time that may have elapsed since the store previously fetched data before fulfilling a fault issues a new fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506875-stalenessinterval?language=objc
func (m_ ManagedObjectContext) SetStalenessInterval(value foundation.TimeInterval) {
	objc.Call[objc.Void](m_, objc.Sel("setStalenessInterval:"), value)
}

// Returns the token associated with the query generation currently in use by this context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1640477-querygenerationtoken?language=objc
func (m_ ManagedObjectContext) QueryGenerationToken() QueryGenerationToken {
	rv := objc.Call[QueryGenerationToken](m_, objc.Sel("queryGenerationToken"))
	return rv
}

// The set of objects that have been inserted into the context but not yet saved in a persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506192-insertedobjects?language=objc
func (m_ ManagedObjectContext) InsertedObjects() foundation.Set {
	rv := objc.Call[foundation.Set](m_, objc.Sel("insertedObjects"))
	return rv
}

// The persistent store coordinator of the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506618-persistentstorecoordinator?language=objc
func (m_ ManagedObjectContext) PersistentStoreCoordinator() PersistentStoreCoordinator {
	rv := objc.Call[PersistentStoreCoordinator](m_, objc.Sel("persistentStoreCoordinator"))
	return rv
}

// The persistent store coordinator of the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506618-persistentstorecoordinator?language=objc
func (m_ ManagedObjectContext) SetPersistentStoreCoordinator(value IPersistentStoreCoordinator) {
	objc.Call[objc.Void](m_, objc.Sel("setPersistentStoreCoordinator:"), objc.Ptr(value))
}

// The set of objects registered with the context that have uncommitted changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506985-updatedobjects?language=objc
func (m_ ManagedObjectContext) UpdatedObjects() foundation.Set {
	rv := objc.Call[foundation.Set](m_, objc.Sel("updatedObjects"))
	return rv
}

// The developer-provided name of the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506231-name?language=objc
func (m_ ManagedObjectContext) Name() string {
	rv := objc.Call[string](m_, objc.Sel("name"))
	return rv
}

// The developer-provided name of the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506231-name?language=objc
func (m_ ManagedObjectContext) SetName(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setName:"), value)
}

// The user information for the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506740-userinfo?language=objc
func (m_ ManagedObjectContext) UserInfo() foundation.MutableDictionary {
	rv := objc.Call[foundation.MutableDictionary](m_, objc.Sel("userInfo"))
	return rv
}

// A Boolean value that determines whether the context turns inaccessible faults into deleted objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506221-shoulddeleteinaccessiblefaults?language=objc
func (m_ ManagedObjectContext) ShouldDeleteInaccessibleFaults() bool {
	rv := objc.Call[bool](m_, objc.Sel("shouldDeleteInaccessibleFaults"))
	return rv
}

// A Boolean value that determines whether the context turns inaccessible faults into deleted objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506221-shoulddeleteinaccessiblefaults?language=objc
func (m_ ManagedObjectContext) SetShouldDeleteInaccessibleFaults(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setShouldDeleteInaccessibleFaults:"), value)
}

// A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506228-propagatesdeletesatendofevent?language=objc
func (m_ ManagedObjectContext) PropagatesDeletesAtEndOfEvent() bool {
	rv := objc.Call[bool](m_, objc.Sel("propagatesDeletesAtEndOfEvent"))
	return rv
}

// A Boolean value that indicates whether the context propagates deletes at the end of the event in which a change was made. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506228-propagatesdeletesatendofevent?language=objc
func (m_ ManagedObjectContext) SetPropagatesDeletesAtEndOfEvent(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setPropagatesDeletesAtEndOfEvent:"), value)
}

// The object that provides undo support for the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506663-undomanager?language=objc
func (m_ ManagedObjectContext) UndoManager() foundation.UndoManager {
	rv := objc.Call[foundation.UndoManager](m_, objc.Sel("undoManager"))
	return rv
}

// The object that provides undo support for the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506663-undomanager?language=objc
func (m_ ManagedObjectContext) SetUndoManager(value foundation.IUndoManager) {
	objc.Call[objc.Void](m_, objc.Sel("setUndoManager:"), objc.Ptr(value))
}

// The parent of the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506529-parentcontext?language=objc
func (m_ ManagedObjectContext) ParentContext() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](m_, objc.Sel("parentContext"))
	return rv
}

// The parent of the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506529-parentcontext?language=objc
func (m_ ManagedObjectContext) SetParentContext(value IManagedObjectContext) {
	objc.Call[objc.Void](m_, objc.Sel("setParentContext:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1845237-automaticallymergeschangesfrompa?language=objc
func (m_ ManagedObjectContext) AutomaticallyMergesChangesFromParent() bool {
	rv := objc.Call[bool](m_, objc.Sel("automaticallyMergesChangesFromParent"))
	return rv
}

// A Boolean value that indicates whether the context automatically merges changes saved to its persistent store coordinator or parent context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1845237-automaticallymergeschangesfrompa?language=objc
func (m_ ManagedObjectContext) SetAutomaticallyMergesChangesFromParent(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAutomaticallyMergesChangesFromParent:"), value)
}

// The set of objects that will be removed from their persistent store during the next save operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsmanagedobjectcontext/1506699-deletedobjects?language=objc
func (m_ ManagedObjectContext) DeletedObjects() foundation.Set {
	rv := objc.Call[foundation.Set](m_, objc.Sel("deletedObjects"))
	return rv
}
