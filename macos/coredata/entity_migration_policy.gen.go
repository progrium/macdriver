// Code generated by DarwinKit. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [EntityMigrationPolicy] class.
var EntityMigrationPolicyClass = _EntityMigrationPolicyClass{objc.GetClass("NSEntityMigrationPolicy")}

type _EntityMigrationPolicyClass struct {
	objc.Class
}

// An interface definition for the [EntityMigrationPolicy] class.
type IEntityMigrationPolicy interface {
	objc.IObject
	PerformCustomValidationForEntityMappingManagerError(mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool
	CreateDestinationInstancesForSourceInstanceEntityMappingManagerError(sInstance IManagedObject, mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool
	CreateRelationshipsForDestinationInstanceEntityMappingManagerError(dInstance IManagedObject, mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool
	BeginEntityMappingManagerError(mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool
	EndEntityMappingManagerError(mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool
	EndRelationshipCreationForEntityMappingManagerError(mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool
	EndInstanceCreationForEntityMappingManagerError(mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool
}

// A policy instance that customizes the migration process for an entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymigrationpolicy?language=objc
type EntityMigrationPolicy struct {
	objc.Object
}

func EntityMigrationPolicyFrom(ptr unsafe.Pointer) EntityMigrationPolicy {
	return EntityMigrationPolicy{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _EntityMigrationPolicyClass) Alloc() EntityMigrationPolicy {
	rv := objc.Call[EntityMigrationPolicy](ec, objc.Sel("alloc"))
	return rv
}

func (ec _EntityMigrationPolicyClass) New() EntityMigrationPolicy {
	rv := objc.Call[EntityMigrationPolicy](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewEntityMigrationPolicy() EntityMigrationPolicy {
	return EntityMigrationPolicyClass.New()
}

func (e_ EntityMigrationPolicy) Init() EntityMigrationPolicy {
	rv := objc.Call[EntityMigrationPolicy](e_, objc.Sel("init"))
	return rv
}

// Provides the option to perform custom validation on migrated objects during the validation stage of the entity migration policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymigrationpolicy/1423791-performcustomvalidationforentity?language=objc
func (e_ EntityMigrationPolicy) PerformCustomValidationForEntityMappingManagerError(mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool {
	rv := objc.Call[bool](e_, objc.Sel("performCustomValidationForEntityMapping:manager:error:"), mapping, manager, error)
	return rv
}

// Creates the destination instance(s) for a given source instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymigrationpolicy/1423801-createdestinationinstancesforsou?language=objc
func (e_ EntityMigrationPolicy) CreateDestinationInstancesForSourceInstanceEntityMappingManagerError(sInstance IManagedObject, mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool {
	rv := objc.Call[bool](e_, objc.Sel("createDestinationInstancesForSourceInstance:entityMapping:manager:error:"), sInstance, mapping, manager, error)
	return rv
}

// Constructs the relationships between the newly-created destination instances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymigrationpolicy/1423783-createrelationshipsfordestinatio?language=objc
func (e_ EntityMigrationPolicy) CreateRelationshipsForDestinationInstanceEntityMappingManagerError(dInstance IManagedObject, mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool {
	rv := objc.Call[bool](e_, objc.Sel("createRelationshipsForDestinationInstance:entityMapping:manager:error:"), dInstance, mapping, manager, error)
	return rv
}

// Sets up state information before the start of a given entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymigrationpolicy/1423785-beginentitymapping?language=objc
func (e_ EntityMigrationPolicy) BeginEntityMappingManagerError(mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool {
	rv := objc.Call[bool](e_, objc.Sel("beginEntityMapping:manager:error:"), mapping, manager, error)
	return rv
}

// Performs cleanup at the end of the migration, from any phase of the mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymigrationpolicy/1423787-endentitymapping?language=objc
func (e_ EntityMigrationPolicy) EndEntityMappingManagerError(mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool {
	rv := objc.Call[bool](e_, objc.Sel("endEntityMapping:manager:error:"), mapping, manager, error)
	return rv
}

// Indicates the end of the relationship creation stage for the specified entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymigrationpolicy/1423793-endrelationshipcreationforentity?language=objc
func (e_ EntityMigrationPolicy) EndRelationshipCreationForEntityMappingManagerError(mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool {
	rv := objc.Call[bool](e_, objc.Sel("endRelationshipCreationForEntityMapping:manager:error:"), mapping, manager, error)
	return rv
}

// Indicates the end of the instance creation stage for the specified entity mapping, and the precursor to the next migration stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymigrationpolicy/1423805-endinstancecreationforentitymapp?language=objc
func (e_ EntityMigrationPolicy) EndInstanceCreationForEntityMappingManagerError(mapping IEntityMapping, manager IMigrationManager, error unsafe.Pointer) bool {
	rv := objc.Call[bool](e_, objc.Sel("endInstanceCreationForEntityMapping:manager:error:"), mapping, manager, error)
	return rv
}
