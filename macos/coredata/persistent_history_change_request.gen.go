// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentHistoryChangeRequest] class.
var PersistentHistoryChangeRequestClass = _PersistentHistoryChangeRequestClass{objc.GetClass("NSPersistentHistoryChangeRequest")}

type _PersistentHistoryChangeRequestClass struct {
	objc.Class
}

// An interface definition for the [PersistentHistoryChangeRequest] class.
type IPersistentHistoryChangeRequest interface {
	IPersistentStoreRequest
	FetchRequest() FetchRequest
	SetFetchRequest(value IFetchRequest)
	Token() PersistentHistoryToken
	ResultType() PersistentHistoryResultType
	SetResultType(value PersistentHistoryResultType)
}

// A request to fetch or purge persistent history. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest?language=objc
type PersistentHistoryChangeRequest struct {
	PersistentStoreRequest
}

func PersistentHistoryChangeRequestFrom(ptr unsafe.Pointer) PersistentHistoryChangeRequest {
	return PersistentHistoryChangeRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

func (pc _PersistentHistoryChangeRequestClass) DeleteHistoryBeforeDate(date foundation.IDate) PersistentHistoryChangeRequest {
	rv := objc.Call[PersistentHistoryChangeRequest](pc, objc.Sel("deleteHistoryBeforeDate:"), objc.Ptr(date))
	return rv
}

// Purges history older than a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/2892338-deletehistorybeforedate?language=objc
func PersistentHistoryChangeRequest_DeleteHistoryBeforeDate(date foundation.IDate) PersistentHistoryChangeRequest {
	return PersistentHistoryChangeRequestClass.DeleteHistoryBeforeDate(date)
}

func (pc _PersistentHistoryChangeRequestClass) FetchHistoryAfterTransaction(transaction IPersistentHistoryTransaction) PersistentHistoryChangeRequest {
	rv := objc.Call[PersistentHistoryChangeRequest](pc, objc.Sel("fetchHistoryAfterTransaction:"), objc.Ptr(transaction))
	return rv
}

// Retrieves history since a given transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/2892344-fetchhistoryaftertransaction?language=objc
func PersistentHistoryChangeRequest_FetchHistoryAfterTransaction(transaction IPersistentHistoryTransaction) PersistentHistoryChangeRequest {
	return PersistentHistoryChangeRequestClass.FetchHistoryAfterTransaction(transaction)
}

func (pc _PersistentHistoryChangeRequestClass) FetchHistoryWithFetchRequest(fetchRequest IFetchRequest) PersistentHistoryChangeRequest {
	rv := objc.Call[PersistentHistoryChangeRequest](pc, objc.Sel("fetchHistoryWithFetchRequest:"), objc.Ptr(fetchRequest))
	return rv
}

// Retrieves history based on a fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/3240592-fetchhistorywithfetchrequest?language=objc
func PersistentHistoryChangeRequest_FetchHistoryWithFetchRequest(fetchRequest IFetchRequest) PersistentHistoryChangeRequest {
	return PersistentHistoryChangeRequestClass.FetchHistoryWithFetchRequest(fetchRequest)
}

func (pc _PersistentHistoryChangeRequestClass) DeleteHistoryBeforeTransaction(transaction IPersistentHistoryTransaction) PersistentHistoryChangeRequest {
	rv := objc.Call[PersistentHistoryChangeRequest](pc, objc.Sel("deleteHistoryBeforeTransaction:"), objc.Ptr(transaction))
	return rv
}

// Purges history older than a given transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/2892345-deletehistorybeforetransaction?language=objc
func PersistentHistoryChangeRequest_DeleteHistoryBeforeTransaction(transaction IPersistentHistoryTransaction) PersistentHistoryChangeRequest {
	return PersistentHistoryChangeRequestClass.DeleteHistoryBeforeTransaction(transaction)
}

func (pc _PersistentHistoryChangeRequestClass) FetchHistoryAfterDate(date foundation.IDate) PersistentHistoryChangeRequest {
	rv := objc.Call[PersistentHistoryChangeRequest](pc, objc.Sel("fetchHistoryAfterDate:"), objc.Ptr(date))
	return rv
}

// Retrieves history since a given date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/2892339-fetchhistoryafterdate?language=objc
func PersistentHistoryChangeRequest_FetchHistoryAfterDate(date foundation.IDate) PersistentHistoryChangeRequest {
	return PersistentHistoryChangeRequestClass.FetchHistoryAfterDate(date)
}

func (pc _PersistentHistoryChangeRequestClass) FetchHistoryAfterToken(token IPersistentHistoryToken) PersistentHistoryChangeRequest {
	rv := objc.Call[PersistentHistoryChangeRequest](pc, objc.Sel("fetchHistoryAfterToken:"), objc.Ptr(token))
	return rv
}

// Retrieves the request history after a given token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/2892337-fetchhistoryaftertoken?language=objc
func PersistentHistoryChangeRequest_FetchHistoryAfterToken(token IPersistentHistoryToken) PersistentHistoryChangeRequest {
	return PersistentHistoryChangeRequestClass.FetchHistoryAfterToken(token)
}

func (pc _PersistentHistoryChangeRequestClass) DeleteHistoryBeforeToken(token IPersistentHistoryToken) PersistentHistoryChangeRequest {
	rv := objc.Call[PersistentHistoryChangeRequest](pc, objc.Sel("deleteHistoryBeforeToken:"), objc.Ptr(token))
	return rv
}

// Purges history older than that defined by a given token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/2892340-deletehistorybeforetoken?language=objc
func PersistentHistoryChangeRequest_DeleteHistoryBeforeToken(token IPersistentHistoryToken) PersistentHistoryChangeRequest {
	return PersistentHistoryChangeRequestClass.DeleteHistoryBeforeToken(token)
}

func (pc _PersistentHistoryChangeRequestClass) Alloc() PersistentHistoryChangeRequest {
	rv := objc.Call[PersistentHistoryChangeRequest](pc, objc.Sel("alloc"))
	return rv
}

func PersistentHistoryChangeRequest_Alloc() PersistentHistoryChangeRequest {
	return PersistentHistoryChangeRequestClass.Alloc()
}

func (pc _PersistentHistoryChangeRequestClass) New() PersistentHistoryChangeRequest {
	rv := objc.Call[PersistentHistoryChangeRequest](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentHistoryChangeRequest() PersistentHistoryChangeRequest {
	return PersistentHistoryChangeRequestClass.New()
}

func (p_ PersistentHistoryChangeRequest) Init() PersistentHistoryChangeRequest {
	rv := objc.Call[PersistentHistoryChangeRequest](p_, objc.Sel("init"))
	return rv
}

// The specified fetch request, when retrieving history. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/3240593-fetchrequest?language=objc
func (p_ PersistentHistoryChangeRequest) FetchRequest() FetchRequest {
	rv := objc.Call[FetchRequest](p_, objc.Sel("fetchRequest"))
	return rv
}

// The specified fetch request, when retrieving history. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/3240593-fetchrequest?language=objc
func (p_ PersistentHistoryChangeRequest) SetFetchRequest(value IFetchRequest) {
	objc.Call[objc.Void](p_, objc.Sel("setFetchRequest:"), objc.Ptr(value))
}

// The specified token, when retrieving history defined by a token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/2892342-token?language=objc
func (p_ PersistentHistoryChangeRequest) Token() PersistentHistoryToken {
	rv := objc.Call[PersistentHistoryToken](p_, objc.Sel("token"))
	return rv
}

// The type of result that this request returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/2892341-resulttype?language=objc
func (p_ PersistentHistoryChangeRequest) ResultType() PersistentHistoryResultType {
	rv := objc.Call[PersistentHistoryResultType](p_, objc.Sel("resultType"))
	return rv
}

// The type of result that this request returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychangerequest/2892341-resulttype?language=objc
func (p_ PersistentHistoryChangeRequest) SetResultType(value PersistentHistoryResultType) {
	objc.Call[objc.Void](p_, objc.Sel("setResultType:"), value)
}
