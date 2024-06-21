// Code generated by DarwinKit. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [PersistentCloudKitContainerEventResult] class.
var PersistentCloudKitContainerEventResultClass = _PersistentCloudKitContainerEventResultClass{objc.GetClass("NSPersistentCloudKitContainerEventResult")}

type _PersistentCloudKitContainerEventResultClass struct {
	objc.Class
}

// An interface definition for the [PersistentCloudKitContainerEventResult] class.
type IPersistentCloudKitContainerEventResult interface {
	IPersistentStoreResult
	Result() objc.Object
	ResultType() PersistentCloudKitContainerEventResultType
}

// The result of a request to fetch persistent CloudKit container events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventresult?language=objc
type PersistentCloudKitContainerEventResult struct {
	PersistentStoreResult
}

func PersistentCloudKitContainerEventResultFrom(ptr unsafe.Pointer) PersistentCloudKitContainerEventResult {
	return PersistentCloudKitContainerEventResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}

func (pc _PersistentCloudKitContainerEventResultClass) Alloc() PersistentCloudKitContainerEventResult {
	rv := objc.Call[PersistentCloudKitContainerEventResult](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PersistentCloudKitContainerEventResultClass) New() PersistentCloudKitContainerEventResult {
	rv := objc.Call[PersistentCloudKitContainerEventResult](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentCloudKitContainerEventResult() PersistentCloudKitContainerEventResult {
	return PersistentCloudKitContainerEventResultClass.New()
}

func (p_ PersistentCloudKitContainerEventResult) Init() PersistentCloudKitContainerEventResult {
	rv := objc.Call[PersistentCloudKitContainerEventResult](p_, objc.Sel("init"))
	return rv
}

// The result of the persistent CloudKit container event request, which the result type determines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventresult/3618821-result?language=objc
func (p_ PersistentCloudKitContainerEventResult) Result() objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("result"))
	return rv
}

// The type of result that the CloudKit container event fetch request returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventresult/3618822-resulttype?language=objc
func (p_ PersistentCloudKitContainerEventResult) ResultType() PersistentCloudKitContainerEventResultType {
	rv := objc.Call[PersistentCloudKitContainerEventResultType](p_, objc.Sel("resultType"))
	return rv
}
