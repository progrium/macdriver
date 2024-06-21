// Code generated by DarwinKit. DO NOT EDIT.

package fileprovider

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A File Provider extension in which the system replicates the contents on disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension?language=objc
type PFileProviderReplicatedExtension interface {
	// optional
	CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler(itemTemplate objc.Object, fields FileProviderItemFields, url foundation.URL, options FileProviderCreateItemOptions, request FileProviderRequest, completionHandler func(createdItem objc.Object, stillPendingFields FileProviderItemFields, shouldFetchContent bool, error foundation.Error)) foundation.Progress
	HasCreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler() bool

	// optional
	ImportDidFinishWithCompletionHandler(completionHandler func())
	HasImportDidFinishWithCompletionHandler() bool

	// optional
	MaterializedItemsDidChangeWithCompletionHandler(completionHandler func())
	HasMaterializedItemsDidChangeWithCompletionHandler() bool

	// optional
	DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler(identifier FileProviderItemIdentifier, version FileProviderItemVersion, options FileProviderDeleteItemOptions, request FileProviderRequest, completionHandler func(arg0 foundation.Error)) foundation.Progress
	HasDeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler() bool

	// optional
	ItemForIdentifierRequestCompletionHandler(identifier FileProviderItemIdentifier, request FileProviderRequest, completionHandler func(arg0 objc.Object, arg1 foundation.Error)) foundation.Progress
	HasItemForIdentifierRequestCompletionHandler() bool

	// optional
	InitWithDomain(domain FileProviderDomain) objc.Object
	HasInitWithDomain() bool

	// optional
	PendingItemsDidChangeWithCompletionHandler(completionHandler func())
	HasPendingItemsDidChangeWithCompletionHandler() bool

	// optional
	Invalidate()
	HasInvalidate() bool

	// optional
	ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler(item objc.Object, version FileProviderItemVersion, changedFields FileProviderItemFields, newContents foundation.URL, options FileProviderModifyItemOptions, request FileProviderRequest, completionHandler func(item objc.Object, stillPendingFields FileProviderItemFields, shouldFetchContent bool, error foundation.Error)) foundation.Progress
	HasModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler() bool

	// optional
	FetchContentsForItemWithIdentifierVersionRequestCompletionHandler(itemIdentifier FileProviderItemIdentifier, requestedVersion FileProviderItemVersion, request FileProviderRequest, completionHandler func(fileContents foundation.URL, item objc.Object, error foundation.Error)) foundation.Progress
	HasFetchContentsForItemWithIdentifierVersionRequestCompletionHandler() bool
}

// ensure impl type implements protocol interface
var _ PFileProviderReplicatedExtension = (*FileProviderReplicatedExtensionObject)(nil)

// A concrete type for the [PFileProviderReplicatedExtension] protocol.
type FileProviderReplicatedExtensionObject struct {
	objc.Object
}

func (f_ FileProviderReplicatedExtensionObject) HasCreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("createItemBasedOnTemplate:fields:contents:options:request:completionHandler:"))
}

// Tells the file provider to create or import an item based on a template. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3656549-createitembasedontemplate?language=objc
func (f_ FileProviderReplicatedExtensionObject) CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler(itemTemplate objc.Object, fields FileProviderItemFields, url foundation.URL, options FileProviderCreateItemOptions, request FileProviderRequest, completionHandler func(createdItem objc.Object, stillPendingFields FileProviderItemFields, shouldFetchContent bool, error foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("createItemBasedOnTemplate:fields:contents:options:request:completionHandler:"), itemTemplate, fields, url, options, request, completionHandler)
	return rv
}

func (f_ FileProviderReplicatedExtensionObject) HasImportDidFinishWithCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("importDidFinishWithCompletionHandler:"))
}

// Tells the File Provider extension that the system finished importing items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3553304-importdidfinishwithcompletionhan?language=objc
func (f_ FileProviderReplicatedExtensionObject) ImportDidFinishWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](f_, objc.Sel("importDidFinishWithCompletionHandler:"), completionHandler)
}

func (f_ FileProviderReplicatedExtensionObject) HasMaterializedItemsDidChangeWithCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("materializedItemsDidChangeWithCompletionHandler:"))
}

// Tells the file provider that the set of materialized items changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3553308-materializeditemsdidchangewithco?language=objc
func (f_ FileProviderReplicatedExtensionObject) MaterializedItemsDidChangeWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](f_, objc.Sel("materializedItemsDidChangeWithCompletionHandler:"), completionHandler)
}

func (f_ FileProviderReplicatedExtensionObject) HasDeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("deleteItemWithIdentifier:baseVersion:options:request:completionHandler:"))
}

// Tells the file provider to delete an item forever. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3656550-deleteitemwithidentifier?language=objc
func (f_ FileProviderReplicatedExtensionObject) DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler(identifier FileProviderItemIdentifier, version FileProviderItemVersion, options FileProviderDeleteItemOptions, request FileProviderRequest, completionHandler func(arg0 foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("deleteItemWithIdentifier:baseVersion:options:request:completionHandler:"), identifier, version, options, request, completionHandler)
	return rv
}

func (f_ FileProviderReplicatedExtensionObject) HasItemForIdentifierRequestCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("itemForIdentifier:request:completionHandler:"))
}

// Asks the file provider for the metadata of the provided item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3656551-itemforidentifier?language=objc
func (f_ FileProviderReplicatedExtensionObject) ItemForIdentifierRequestCompletionHandler(identifier FileProviderItemIdentifier, request FileProviderRequest, completionHandler func(arg0 objc.Object, arg1 foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("itemForIdentifier:request:completionHandler:"), identifier, request, completionHandler)
	return rv
}

func (f_ FileProviderReplicatedExtensionObject) HasInitWithDomain() bool {
	return f_.RespondsToSelector(objc.Sel("initWithDomain:"))
}

// Creates an instance of the file provider for the specified domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3553305-initwithdomain?language=objc
func (f_ FileProviderReplicatedExtensionObject) InitWithDomain(domain FileProviderDomain) objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("initWithDomain:"), domain)
	return rv
}

func (f_ FileProviderReplicatedExtensionObject) HasPendingItemsDidChangeWithCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("pendingItemsDidChangeWithCompletionHandler:"))
}

// Tells the file provider extension that the set of pending items has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3727821-pendingitemsdidchangewithcomplet?language=objc
func (f_ FileProviderReplicatedExtensionObject) PendingItemsDidChangeWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](f_, objc.Sel("pendingItemsDidChangeWithCompletionHandler:"), completionHandler)
}

func (f_ FileProviderReplicatedExtensionObject) HasInvalidate() bool {
	return f_.RespondsToSelector(objc.Sel("invalidate"))
}

// Tells the file provider to perform any necessary cleanup so that the system can deallocate it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3553306-invalidate?language=objc
func (f_ FileProviderReplicatedExtensionObject) Invalidate() {
	objc.Call[objc.Void](f_, objc.Sel("invalidate"))
}

func (f_ FileProviderReplicatedExtensionObject) HasModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("modifyItem:baseVersion:changedFields:contents:options:request:completionHandler:"))
}

// Tells the file provider that an item’s content or metadata changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3656552-modifyitem?language=objc
func (f_ FileProviderReplicatedExtensionObject) ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler(item objc.Object, version FileProviderItemVersion, changedFields FileProviderItemFields, newContents foundation.URL, options FileProviderModifyItemOptions, request FileProviderRequest, completionHandler func(item objc.Object, stillPendingFields FileProviderItemFields, shouldFetchContent bool, error foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("modifyItem:baseVersion:changedFields:contents:options:request:completionHandler:"), item, version, changedFields, newContents, options, request, completionHandler)
	return rv
}

func (f_ FileProviderReplicatedExtensionObject) HasFetchContentsForItemWithIdentifierVersionRequestCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("fetchContentsForItemWithIdentifier:version:request:completionHandler:"))
}

// Tells the file provider to download the requested item from remote storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderreplicatedextension/3553303-fetchcontentsforitemwithidentifi?language=objc
func (f_ FileProviderReplicatedExtensionObject) FetchContentsForItemWithIdentifierVersionRequestCompletionHandler(itemIdentifier FileProviderItemIdentifier, requestedVersion FileProviderItemVersion, request FileProviderRequest, completionHandler func(fileContents foundation.URL, item objc.Object, error foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("fetchContentsForItemWithIdentifier:version:request:completionHandler:"), itemIdentifier, requestedVersion, request, completionHandler)
	return rv
}
