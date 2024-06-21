// Code generated by DarwinKit. DO NOT EDIT.

package cloudkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [Operation] class.
var OperationClass = _OperationClass{objc.GetClass("CKOperation")}

type _OperationClass struct {
	objc.Class
}

// An interface definition for the [Operation] class.
type IOperation interface {
	foundation.IOperation
	OperationID() OperationID
	LongLivedOperationWasPersistedBlock() func()
	SetLongLivedOperationWasPersistedBlock(value func())
	Configuration() OperationConfiguration
	SetConfiguration(value IOperationConfiguration)
	Group() OperationGroup
	SetGroup(value IOperationGroup)
}

// The abstract base class for all operations that execute in a database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation?language=objc
type Operation struct {
	foundation.Operation
}

func OperationFrom(ptr unsafe.Pointer) Operation {
	return Operation{
		Operation: foundation.OperationFrom(ptr),
	}
}

func (o_ Operation) Init() Operation {
	rv := objc.Call[Operation](o_, objc.Sel("init"))
	return rv
}

func (oc _OperationClass) Alloc() Operation {
	rv := objc.Call[Operation](oc, objc.Sel("alloc"))
	return rv
}

func (oc _OperationClass) New() Operation {
	rv := objc.Call[Operation](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOperation() Operation {
	return OperationClass.New()
}

// A unique identifier for a long-lived operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/1452362-operationid?language=objc
func (o_ Operation) OperationID() OperationID {
	rv := objc.Call[OperationID](o_, objc.Sel("operationID"))
	return rv
}

// The block to execute when the server begins to store callbacks for the long-lived operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/1452366-longlivedoperationwaspersistedbl?language=objc
func (o_ Operation) LongLivedOperationWasPersistedBlock() func() {
	rv := objc.Call[func()](o_, objc.Sel("longLivedOperationWasPersistedBlock"))
	return rv
}

// The block to execute when the server begins to store callbacks for the long-lived operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/1452366-longlivedoperationwaspersistedbl?language=objc
func (o_ Operation) SetLongLivedOperationWasPersistedBlock(value func()) {
	objc.Call[objc.Void](o_, objc.Sel("setLongLivedOperationWasPersistedBlock:"), value)
}

// The operation’s configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/2866213-configuration?language=objc
func (o_ Operation) Configuration() OperationConfiguration {
	rv := objc.Call[OperationConfiguration](o_, objc.Sel("configuration"))
	return rv
}

// The operation’s configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/2866213-configuration?language=objc
func (o_ Operation) SetConfiguration(value IOperationConfiguration) {
	objc.Call[objc.Void](o_, objc.Sel("setConfiguration:"), value)
}

// The operation’s group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/2866228-group?language=objc
func (o_ Operation) Group() OperationGroup {
	rv := objc.Call[OperationGroup](o_, objc.Sel("group"))
	return rv
}

// The operation’s group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckoperation/2866228-group?language=objc
func (o_ Operation) SetGroup(value IOperationGroup) {
	objc.Call[objc.Void](o_, objc.Sel("setGroup:"), value)
}
