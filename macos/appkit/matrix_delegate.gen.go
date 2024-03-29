// Code generated by DarwinKit. DO NOT EDIT.

package appkit

// The NSMatrixDelegate protocol defines the optional methods implemented by delegates of NSMatrix objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrixdelegate?language=objc
type PMatrixDelegate interface {
	PControlTextEditingDelegate
}

// A delegate implementation builder for the [PMatrixDelegate] protocol.
type MatrixDelegate struct {
	ControlTextEditingDelegate
}

// ensure impl type implements protocol interface
var _ PMatrixDelegate = (*MatrixDelegateObject)(nil)

// A concrete type for the [PMatrixDelegate] protocol.
type MatrixDelegateObject struct {
	ControlTextEditingDelegateObject
}
