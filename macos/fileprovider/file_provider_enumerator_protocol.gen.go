// Code generated by DarwinKit. DO NOT EDIT.

package fileprovider

import (
	"github.com/progrium/darwinkit/objc"
)

// A protocol for enumerating items and changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerator?language=objc
type PFileProviderEnumerator interface {
	// optional
	EnumerateChangesForObserverFromSyncAnchor(observer FileProviderChangeObserverObject, syncAnchor FileProviderSyncAnchor)
	HasEnumerateChangesForObserverFromSyncAnchor() bool

	// optional
	EnumerateItemsForObserverStartingAtPage(observer FileProviderEnumerationObserverObject, page FileProviderPage)
	HasEnumerateItemsForObserverStartingAtPage() bool

	// optional
	CurrentSyncAnchorWithCompletionHandler(completionHandler func(currentAnchor FileProviderSyncAnchor))
	HasCurrentSyncAnchorWithCompletionHandler() bool

	// optional
	Invalidate()
	HasInvalidate() bool
}

// ensure impl type implements protocol interface
var _ PFileProviderEnumerator = (*FileProviderEnumeratorObject)(nil)

// A concrete type for the [PFileProviderEnumerator] protocol.
type FileProviderEnumeratorObject struct {
	objc.Object
}

func (f_ FileProviderEnumeratorObject) HasEnumerateChangesForObserverFromSyncAnchor() bool {
	return f_.RespondsToSelector(objc.Sel("enumerateChangesForObserver:fromSyncAnchor:"))
}

// Requests the next batch of changes after the specified sync anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerator/2879608-enumeratechangesforobserver?language=objc
func (f_ FileProviderEnumeratorObject) EnumerateChangesForObserverFromSyncAnchor(observer FileProviderChangeObserverObject, syncAnchor FileProviderSyncAnchor) {
	po0 := objc.WrapAsProtocol("NSFileProviderChangeObserver", observer)
	objc.Call[objc.Void](f_, objc.Sel("enumerateChangesForObserver:fromSyncAnchor:"), po0, syncAnchor)
}

func (f_ FileProviderEnumeratorObject) HasEnumerateItemsForObserverStartingAtPage() bool {
	return f_.RespondsToSelector(objc.Sel("enumerateItemsForObserver:startingAtPage:"))
}

// Requests the next batch of items, starting at the specified page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerator/2879613-enumerateitemsforobserver?language=objc
func (f_ FileProviderEnumeratorObject) EnumerateItemsForObserverStartingAtPage(observer FileProviderEnumerationObserverObject, page FileProviderPage) {
	po0 := objc.WrapAsProtocol("NSFileProviderEnumerationObserver", observer)
	objc.Call[objc.Void](f_, objc.Sel("enumerateItemsForObserver:startingAtPage:"), po0, page)
}

func (f_ FileProviderEnumeratorObject) HasCurrentSyncAnchorWithCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("currentSyncAnchorWithCompletionHandler:"))
}

// Returns the current sync anchor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerator/2890930-currentsyncanchorwithcompletionh?language=objc
func (f_ FileProviderEnumeratorObject) CurrentSyncAnchorWithCompletionHandler(completionHandler func(currentAnchor FileProviderSyncAnchor)) {
	objc.Call[objc.Void](f_, objc.Sel("currentSyncAnchorWithCompletionHandler:"), completionHandler)
}

func (f_ FileProviderEnumeratorObject) HasInvalidate() bool {
	return f_.RespondsToSelector(objc.Sel("invalidate"))
}

// Stops the enumeration of items and changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderenumerator/2879609-invalidate?language=objc
func (f_ FileProviderEnumeratorObject) Invalidate() {
	objc.Call[objc.Void](f_, objc.Sel("invalidate"))
}
