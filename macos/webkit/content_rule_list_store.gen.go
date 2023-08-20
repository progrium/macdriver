// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContentRuleListStore] class.
var ContentRuleListStoreClass = _ContentRuleListStoreClass{objc.GetClass("WKContentRuleListStore")}

type _ContentRuleListStoreClass struct {
	objc.Class
}

// An interface definition for the [ContentRuleListStore] class.
type IContentRuleListStore interface {
	objc.IObject
	LookUpContentRuleListForIdentifierCompletionHandler(identifier string, completionHandler func(arg0 ContentRuleList, arg1 foundation.Error))
	GetAvailableContentRuleListIdentifiers(completionHandler func(arg0 []string))
	RemoveContentRuleListForIdentifierCompletionHandler(identifier string, completionHandler func(arg0 foundation.Error))
	CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler(identifier string, encodedContentRuleList string, completionHandler func(arg0 ContentRuleList, arg1 foundation.Error))
}

// An object that contains the rules for how to load and filter content in the web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentruleliststore?language=objc
type ContentRuleListStore struct {
	objc.Object
}

func ContentRuleListStoreFrom(ptr unsafe.Pointer) ContentRuleListStore {
	return ContentRuleListStore{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContentRuleListStoreClass) DefaultStore() ContentRuleListStore {
	rv := objc.Call[ContentRuleListStore](cc, objc.Sel("defaultStore"))
	return rv
}

// Returns the default content rule list store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentruleliststore/2902760-defaultstore?language=objc
func ContentRuleListStore_DefaultStore() ContentRuleListStore {
	return ContentRuleListStoreClass.DefaultStore()
}

func (cc _ContentRuleListStoreClass) StoreWithURL(url foundation.IURL) ContentRuleListStore {
	rv := objc.Call[ContentRuleListStore](cc, objc.Sel("storeWithURL:"), objc.Ptr(url))
	return rv
}

// Creates a new content rule list store in the specified directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentruleliststore/2902757-storewithurl?language=objc
func ContentRuleListStore_StoreWithURL(url foundation.IURL) ContentRuleListStore {
	return ContentRuleListStoreClass.StoreWithURL(url)
}

func (cc _ContentRuleListStoreClass) Alloc() ContentRuleListStore {
	rv := objc.Call[ContentRuleListStore](cc, objc.Sel("alloc"))
	return rv
}

func ContentRuleListStore_Alloc() ContentRuleListStore {
	return ContentRuleListStoreClass.Alloc()
}

func (cc _ContentRuleListStoreClass) New() ContentRuleListStore {
	rv := objc.Call[ContentRuleListStore](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContentRuleListStore() ContentRuleListStore {
	return ContentRuleListStoreClass.New()
}

func (c_ ContentRuleListStore) Init() ContentRuleListStore {
	rv := objc.Call[ContentRuleListStore](c_, objc.Sel("init"))
	return rv
}

// Searches asynchronously for a specific rule list in the data store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentruleliststore/2902747-lookupcontentrulelistforidentifi?language=objc
func (c_ ContentRuleListStore) LookUpContentRuleListForIdentifierCompletionHandler(identifier string, completionHandler func(arg0 ContentRuleList, arg1 foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("lookUpContentRuleListForIdentifier:completionHandler:"), identifier, completionHandler)
}

// Fetches the identifiers for all rule lists in the store asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentruleliststore/2902759-getavailablecontentrulelistident?language=objc
func (c_ ContentRuleListStore) GetAvailableContentRuleListIdentifiers(completionHandler func(arg0 []string)) {
	objc.Call[objc.Void](c_, objc.Sel("getAvailableContentRuleListIdentifiers:"), completionHandler)
}

// Removes a rule list from the current data store asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentruleliststore/2902764-removecontentrulelistforidentifi?language=objc
func (c_ ContentRuleListStore) RemoveContentRuleListForIdentifierCompletionHandler(identifier string, completionHandler func(arg0 foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("removeContentRuleListForIdentifier:completionHandler:"), identifier, completionHandler)
}

// Compiles the specified JSON content into a new rule list and adds it to the current data store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkcontentruleliststore/2902761-compilecontentrulelistforidentif?language=objc
func (c_ ContentRuleListStore) CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler(identifier string, encodedContentRuleList string, completionHandler func(arg0 ContentRuleList, arg1 foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("compileContentRuleListForIdentifier:encodedContentRuleList:completionHandler:"), identifier, encodedContentRuleList, completionHandler)
}