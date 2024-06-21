// Code generated by DarwinKit. DO NOT EDIT.

package coredata

import (
	"github.com/progrium/darwinkit/objc"
)

// A delegate protocol that describes the methods that will be called by the associated fetched results controller when the fetch results have changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontrollerdelegate?language=objc
type PFetchedResultsControllerDelegate interface {
	// optional
	ControllerWillChangeContent(controller FetchedResultsController)
	HasControllerWillChangeContent() bool

	// optional
	ControllerDidChangeContent(controller FetchedResultsController)
	HasControllerDidChangeContent() bool
}

// A delegate implementation builder for the [PFetchedResultsControllerDelegate] protocol.
type FetchedResultsControllerDelegate struct {
	_ControllerWillChangeContent func(controller FetchedResultsController)
	_ControllerDidChangeContent  func(controller FetchedResultsController)
}

func (di *FetchedResultsControllerDelegate) HasControllerWillChangeContent() bool {
	return di._ControllerWillChangeContent != nil
}

// Notifies the receiver that the fetched results controller is about to start processing of one or more changes due to an add, remove, move, or update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontrollerdelegate/1622295-controllerwillchangecontent?language=objc
func (di *FetchedResultsControllerDelegate) SetControllerWillChangeContent(f func(controller FetchedResultsController)) {
	di._ControllerWillChangeContent = f
}

// Notifies the receiver that the fetched results controller is about to start processing of one or more changes due to an add, remove, move, or update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontrollerdelegate/1622295-controllerwillchangecontent?language=objc
func (di *FetchedResultsControllerDelegate) ControllerWillChangeContent(controller FetchedResultsController) {
	di._ControllerWillChangeContent(controller)
}
func (di *FetchedResultsControllerDelegate) HasControllerDidChangeContent() bool {
	return di._ControllerDidChangeContent != nil
}

// Notifies the receiver that the fetched results controller has completed processing of one or more changes due to an add, remove, move, or update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontrollerdelegate/1622290-controllerdidchangecontent?language=objc
func (di *FetchedResultsControllerDelegate) SetControllerDidChangeContent(f func(controller FetchedResultsController)) {
	di._ControllerDidChangeContent = f
}

// Notifies the receiver that the fetched results controller has completed processing of one or more changes due to an add, remove, move, or update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontrollerdelegate/1622290-controllerdidchangecontent?language=objc
func (di *FetchedResultsControllerDelegate) ControllerDidChangeContent(controller FetchedResultsController) {
	di._ControllerDidChangeContent(controller)
}

// ensure impl type implements protocol interface
var _ PFetchedResultsControllerDelegate = (*FetchedResultsControllerDelegateObject)(nil)

// A concrete type for the [PFetchedResultsControllerDelegate] protocol.
type FetchedResultsControllerDelegateObject struct {
	objc.Object
}

func (f_ FetchedResultsControllerDelegateObject) HasControllerWillChangeContent() bool {
	return f_.RespondsToSelector(objc.Sel("controllerWillChangeContent:"))
}

// Notifies the receiver that the fetched results controller is about to start processing of one or more changes due to an add, remove, move, or update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontrollerdelegate/1622295-controllerwillchangecontent?language=objc
func (f_ FetchedResultsControllerDelegateObject) ControllerWillChangeContent(controller FetchedResultsController) {
	objc.Call[objc.Void](f_, objc.Sel("controllerWillChangeContent:"), controller)
}

func (f_ FetchedResultsControllerDelegateObject) HasControllerDidChangeContent() bool {
	return f_.RespondsToSelector(objc.Sel("controllerDidChangeContent:"))
}

// Notifies the receiver that the fetched results controller has completed processing of one or more changes due to an add, remove, move, or update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontrollerdelegate/1622290-controllerdidchangecontent?language=objc
func (f_ FetchedResultsControllerDelegateObject) ControllerDidChangeContent(controller FetchedResultsController) {
	objc.Call[objc.Void](f_, objc.Sel("controllerDidChangeContent:"), controller)
}
