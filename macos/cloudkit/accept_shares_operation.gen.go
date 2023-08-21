// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AcceptSharesOperation] class.
var AcceptSharesOperationClass = _AcceptSharesOperationClass{objc.GetClass("CKAcceptSharesOperation")}

type _AcceptSharesOperationClass struct {
	objc.Class
}

// An interface definition for the [AcceptSharesOperation] class.
type IAcceptSharesOperation interface {
	IOperation
	AcceptSharesCompletionBlock() func(operationError foundation.Error)
	SetAcceptSharesCompletionBlock(value func(operationError foundation.Error))
	ShareMetadatas() []ShareMetadata
	SetShareMetadatas(value []IShareMetadata)
	PerShareCompletionBlock() func(shareMetadata ShareMetadata, acceptedShare Share, error foundation.Error)
	SetPerShareCompletionBlock(value func(shareMetadata ShareMetadata, acceptedShare Share, error foundation.Error))
}

// An operation that confirms a user’s participation in a share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation?language=objc
type AcceptSharesOperation struct {
	Operation
}

func AcceptSharesOperationFrom(ptr unsafe.Pointer) AcceptSharesOperation {
	return AcceptSharesOperation{
		Operation: OperationFrom(ptr),
	}
}

func (a_ AcceptSharesOperation) InitWithShareMetadatas(shareMetadatas []IShareMetadata) AcceptSharesOperation {
	rv := objc.Call[AcceptSharesOperation](a_, objc.Sel("initWithShareMetadatas:"), shareMetadatas)
	return rv
}

// Creates an operation for accepting the specified shares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/1823506-initwithsharemetadatas?language=objc
func NewAcceptSharesOperationWithShareMetadatas(shareMetadatas []IShareMetadata) AcceptSharesOperation {
	instance := AcceptSharesOperationClass.Alloc().InitWithShareMetadatas(shareMetadatas)
	instance.Autorelease()
	return instance
}

func (a_ AcceptSharesOperation) Init() AcceptSharesOperation {
	rv := objc.Call[AcceptSharesOperation](a_, objc.Sel("init"))
	return rv
}

func (ac _AcceptSharesOperationClass) Alloc() AcceptSharesOperation {
	rv := objc.Call[AcceptSharesOperation](ac, objc.Sel("alloc"))
	return rv
}

func AcceptSharesOperation_Alloc() AcceptSharesOperation {
	return AcceptSharesOperationClass.Alloc()
}

func (ac _AcceptSharesOperationClass) New() AcceptSharesOperation {
	rv := objc.Call[AcceptSharesOperation](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAcceptSharesOperation() AcceptSharesOperation {
	return AcceptSharesOperationClass.New()
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/1640442-acceptsharescompletionblock?language=objc
func (a_ AcceptSharesOperation) AcceptSharesCompletionBlock() func(operationError foundation.Error) {
	rv := objc.Call[func(operationError foundation.Error)](a_, objc.Sel("acceptSharesCompletionBlock"))
	return rv
}

// The block to execute when the operation finishes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/1640442-acceptsharescompletionblock?language=objc
func (a_ AcceptSharesOperation) SetAcceptSharesCompletionBlock(value func(operationError foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("setAcceptSharesCompletionBlock:"), value)
}

// The share metadatas to process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/1823508-sharemetadatas?language=objc
func (a_ AcceptSharesOperation) ShareMetadatas() []ShareMetadata {
	rv := objc.Call[[]ShareMetadata](a_, objc.Sel("shareMetadatas"))
	return rv
}

// The share metadatas to process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/1823508-sharemetadatas?language=objc
func (a_ AcceptSharesOperation) SetShareMetadatas(value []IShareMetadata) {
	objc.Call[objc.Void](a_, objc.Sel("setShareMetadatas:"), value)
}

// The block to execute as CloudKit processes individual shares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/1640426-persharecompletionblock?language=objc
func (a_ AcceptSharesOperation) PerShareCompletionBlock() func(shareMetadata ShareMetadata, acceptedShare Share, error foundation.Error) {
	rv := objc.Call[func(shareMetadata ShareMetadata, acceptedShare Share, error foundation.Error)](a_, objc.Sel("perShareCompletionBlock"))
	return rv
}

// The block to execute as CloudKit processes individual shares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckacceptsharesoperation/1640426-persharecompletionblock?language=objc
func (a_ AcceptSharesOperation) SetPerShareCompletionBlock(value func(shareMetadata ShareMetadata, acceptedShare Share, error foundation.Error)) {
	objc.Call[objc.Void](a_, objc.Sel("setPerShareCompletionBlock:"), value)
}
