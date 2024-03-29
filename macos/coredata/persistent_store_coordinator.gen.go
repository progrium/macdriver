// Code generated by DarwinKit. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentStoreCoordinator] class.
var PersistentStoreCoordinatorClass = _PersistentStoreCoordinatorClass{objc.GetClass("NSPersistentStoreCoordinator")}

type _PersistentStoreCoordinatorClass struct {
	objc.Class
}

// An interface definition for the [PersistentStoreCoordinator] class.
type IPersistentStoreCoordinator interface {
	objc.IObject
	SetURLForPersistentStore(url foundation.IURL, store IPersistentStore) bool
	PerformBlockAndWait(block func())
	PerformBlock(block func())
	RemovePersistentStoreError(store IPersistentStore, error foundation.IError) bool
	CurrentPersistentHistoryTokenFromStores(stores []objc.IObject) PersistentHistoryToken
	ManagedObjectIDForURIRepresentation(url foundation.IURL) ManagedObjectID
	MetadataForPersistentStore(store IPersistentStore) map[string]objc.Object
	AddPersistentStoreWithDescriptionCompletionHandler(storeDescription IPersistentStoreDescription, block func(arg0 PersistentStoreDescription, arg1 foundation.Error))
	SetMetadataForPersistentStore(metadata map[string]objc.IObject, store IPersistentStore)
	AddPersistentStoreWithTypeConfigurationURLOptionsError(storeType string, configuration string, storeURL foundation.IURL, options foundation.Dictionary, error foundation.IError) PersistentStore
	PersistentStoreForURL(URL foundation.IURL) PersistentStore
	ExecuteRequestWithContextError(request IPersistentStoreRequest, context IManagedObjectContext, error foundation.IError) objc.Object
	URLForPersistentStore(store IPersistentStore) foundation.URL
	Name() string
	SetName(value string)
	PersistentStores() []PersistentStore
	ManagedObjectModel() ManagedObjectModel
}

// An object that enables an app’s contexts and the underlying persistent stores to work together. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator?language=objc
type PersistentStoreCoordinator struct {
	objc.Object
}

func PersistentStoreCoordinatorFrom(ptr unsafe.Pointer) PersistentStoreCoordinator {
	return PersistentStoreCoordinator{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PersistentStoreCoordinator) InitWithManagedObjectModel(model IManagedObjectModel) PersistentStoreCoordinator {
	rv := objc.Call[PersistentStoreCoordinator](p_, objc.Sel("initWithManagedObjectModel:"), objc.Ptr(model))
	return rv
}

// Creates a persistent store coordinator with the specified managed object model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468895-initwithmanagedobjectmodel?language=objc
func NewPersistentStoreCoordinatorWithManagedObjectModel(model IManagedObjectModel) PersistentStoreCoordinator {
	instance := PersistentStoreCoordinatorClass.Alloc().InitWithManagedObjectModel(model)
	instance.Autorelease()
	return instance
}

func (pc _PersistentStoreCoordinatorClass) Alloc() PersistentStoreCoordinator {
	rv := objc.Call[PersistentStoreCoordinator](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PersistentStoreCoordinatorClass) New() PersistentStoreCoordinator {
	rv := objc.Call[PersistentStoreCoordinator](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentStoreCoordinator() PersistentStoreCoordinator {
	return PersistentStoreCoordinatorClass.New()
}

func (p_ PersistentStoreCoordinator) Init() PersistentStoreCoordinator {
	rv := objc.Call[PersistentStoreCoordinator](p_, objc.Sel("init"))
	return rv
}

// Changes the location of the specified persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468858-seturl?language=objc
func (p_ PersistentStoreCoordinator) SetURLForPersistentStore(url foundation.IURL, store IPersistentStore) bool {
	rv := objc.Call[bool](p_, objc.Sel("setURL:forPersistentStore:"), objc.Ptr(url), objc.Ptr(store))
	return rv
}

// Executes the provided closure on the coordinator’s queue and waits for it to finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468862-performblockandwait?language=objc
func (p_ PersistentStoreCoordinator) PerformBlockAndWait(block func()) {
	objc.Call[objc.Void](p_, objc.Sel("performBlockAndWait:"), block)
}

// Updates the metadata of a specific type of persistent store at the provided location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468893-setmetadata?language=objc
func (pc _PersistentStoreCoordinatorClass) SetMetadataForPersistentStoreOfTypeURLOptionsError(metadata map[string]objc.IObject, storeType string, url foundation.IURL, options foundation.Dictionary, error foundation.IError) bool {
	rv := objc.Call[bool](pc, objc.Sel("setMetadata:forPersistentStoreOfType:URL:options:error:"), metadata, storeType, objc.Ptr(url), options, objc.Ptr(error))
	return rv
}

// Updates the metadata of a specific type of persistent store at the provided location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468893-setmetadata?language=objc
func PersistentStoreCoordinator_SetMetadataForPersistentStoreOfTypeURLOptionsError(metadata map[string]objc.IObject, storeType string, url foundation.IURL, options foundation.Dictionary, error foundation.IError) bool {
	return PersistentStoreCoordinatorClass.SetMetadataForPersistentStoreOfTypeURLOptionsError(metadata, storeType, url, options, error)
}

// Executes the provided closure asynchronously on the coordinator’s queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468794-performblock?language=objc
func (p_ PersistentStoreCoordinator) PerformBlock(block func()) {
	objc.Call[objc.Void](p_, objc.Sel("performBlock:"), block)
}

// Removes the specified persistent store from the coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468907-removepersistentstore?language=objc
func (p_ PersistentStoreCoordinator) RemovePersistentStoreError(store IPersistentStore, error foundation.IError) bool {
	rv := objc.Call[bool](p_, objc.Sel("removePersistentStore:error:"), objc.Ptr(store), objc.Ptr(error))
	return rv
}

// Returns a single persistent history token representing all of the specified stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/3325497-currentpersistenthistorytokenfro?language=objc
func (p_ PersistentStoreCoordinator) CurrentPersistentHistoryTokenFromStores(stores []objc.IObject) PersistentHistoryToken {
	rv := objc.Call[PersistentHistoryToken](p_, objc.Sel("currentPersistentHistoryTokenFromStores:"), stores)
	return rv
}

// Returns the object identifier for the specified URI representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468882-managedobjectidforurirepresentat?language=objc
func (p_ PersistentStoreCoordinator) ManagedObjectIDForURIRepresentation(url foundation.IURL) ManagedObjectID {
	rv := objc.Call[ManagedObjectID](p_, objc.Sel("managedObjectIDForURIRepresentation:"), objc.Ptr(url))
	return rv
}

// Returns the metadata of the specified persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468911-metadataforpersistentstore?language=objc
func (p_ PersistentStoreCoordinator) MetadataForPersistentStore(store IPersistentStore) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](p_, objc.Sel("metadataForPersistentStore:"), objc.Ptr(store))
	return rv
}

// Adds a persistent store using the provided description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1640556-addpersistentstorewithdescriptio?language=objc
func (p_ PersistentStoreCoordinator) AddPersistentStoreWithDescriptionCompletionHandler(storeDescription IPersistentStoreDescription, block func(arg0 PersistentStoreDescription, arg1 foundation.Error)) {
	objc.Call[objc.Void](p_, objc.Sel("addPersistentStoreWithDescription:completionHandler:"), objc.Ptr(storeDescription), block)
}

// Updates the metadata for the specified persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468899-setmetadata?language=objc
func (p_ PersistentStoreCoordinator) SetMetadataForPersistentStore(metadata map[string]objc.IObject, store IPersistentStore) {
	objc.Call[objc.Void](p_, objc.Sel("setMetadata:forPersistentStore:"), metadata, objc.Ptr(store))
}

// Adds a specific type of persistent store at the provided location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468860-addpersistentstorewithtype?language=objc
func (p_ PersistentStoreCoordinator) AddPersistentStoreWithTypeConfigurationURLOptionsError(storeType string, configuration string, storeURL foundation.IURL, options foundation.Dictionary, error foundation.IError) PersistentStore {
	rv := objc.Call[PersistentStore](p_, objc.Sel("addPersistentStoreWithType:configuration:URL:options:error:"), storeType, configuration, objc.Ptr(storeURL), options, objc.Ptr(error))
	return rv
}

// Returns the persistent store for the specified file URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468824-persistentstoreforurl?language=objc
func (p_ PersistentStoreCoordinator) PersistentStoreForURL(URL foundation.IURL) PersistentStore {
	rv := objc.Call[PersistentStore](p_, objc.Sel("persistentStoreForURL:"), objc.Ptr(URL))
	return rv
}

// Returns the metadata of a specific type of persistent store at the provided location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468778-metadataforpersistentstoreoftype?language=objc
func (pc _PersistentStoreCoordinatorClass) MetadataForPersistentStoreOfTypeURLOptionsError(storeType string, url foundation.IURL, options foundation.Dictionary, error foundation.IError) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](pc, objc.Sel("metadataForPersistentStoreOfType:URL:options:error:"), storeType, objc.Ptr(url), options, objc.Ptr(error))
	return rv
}

// Returns the metadata of a specific type of persistent store at the provided location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468778-metadataforpersistentstoreoftype?language=objc
func PersistentStoreCoordinator_MetadataForPersistentStoreOfTypeURLOptionsError(storeType string, url foundation.IURL, options foundation.Dictionary, error foundation.IError) map[string]objc.Object {
	return PersistentStoreCoordinatorClass.MetadataForPersistentStoreOfTypeURLOptionsError(storeType, url, options, error)
}

// Executes the specified request on each of the coordinator’s persistent stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468872-executerequest?language=objc
func (p_ PersistentStoreCoordinator) ExecuteRequestWithContextError(request IPersistentStoreRequest, context IManagedObjectContext, error foundation.IError) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("executeRequest:withContext:error:"), objc.Ptr(request), objc.Ptr(context), objc.Ptr(error))
	return rv
}

// Returns the location of the provided persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468865-urlforpersistentstore?language=objc
func (p_ PersistentStoreCoordinator) URLForPersistentStore(store IPersistentStore) foundation.URL {
	rv := objc.Call[foundation.URL](p_, objc.Sel("URLForPersistentStore:"), objc.Ptr(store))
	return rv
}

// The coordinator’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468929-name?language=objc
func (p_ PersistentStoreCoordinator) Name() string {
	rv := objc.Call[string](p_, objc.Sel("name"))
	return rv
}

// The coordinator’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468929-name?language=objc
func (p_ PersistentStoreCoordinator) SetName(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setName:"), value)
}

// The coordinator’s registered store types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468870-registeredstoretypes?language=objc
func (pc _PersistentStoreCoordinatorClass) RegisteredStoreTypes() map[string]foundation.Value {
	rv := objc.Call[map[string]foundation.Value](pc, objc.Sel("registeredStoreTypes"))
	return rv
}

// The coordinator’s registered store types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468870-registeredstoretypes?language=objc
func PersistentStoreCoordinator_RegisteredStoreTypes() map[string]foundation.Value {
	return PersistentStoreCoordinatorClass.RegisteredStoreTypes()
}

// The coordinator’s persistent stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468790-persistentstores?language=objc
func (p_ PersistentStoreCoordinator) PersistentStores() []PersistentStore {
	rv := objc.Call[[]PersistentStore](p_, objc.Sel("persistentStores"))
	return rv
}

// The coordinator’s managed object model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/1468834-managedobjectmodel?language=objc
func (p_ PersistentStoreCoordinator) ManagedObjectModel() ManagedObjectModel {
	rv := objc.Call[ManagedObjectModel](p_, objc.Sel("managedObjectModel"))
	return rv
}
