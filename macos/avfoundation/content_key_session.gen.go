// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContentKeySession] class.
var ContentKeySessionClass = _ContentKeySessionClass{objc.GetClass("AVContentKeySession")}

type _ContentKeySessionClass struct {
	objc.Class
}

// An interface definition for the [ContentKeySession] class.
type IContentKeySession interface {
	objc.IObject
	MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler(persistableContentKeyData []byte, handler func(secureTokenData []byte, error foundation.Error))
	RenewExpiringResponseDataForContentKeyRequest(contentKeyRequest IContentKeyRequest)
	AddContentKeyRecipient(recipient PContentKeyRecipient)
	AddContentKeyRecipientObject(recipientObject objc.IObject)
	SetDelegateQueue(delegate PContentKeySessionDelegate, delegateQueue dispatch.Queue)
	SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue)
	RemoveContentKeyRecipient(recipient PContentKeyRecipient)
	RemoveContentKeyRecipientObject(recipientObject objc.IObject)
	InvalidatePersistableContentKeyOptionsCompletionHandler(persistableContentKeyData []byte, options map[ContentKeySessionServerPlaybackContextOption]objc.IObject, handler func(secureTokenData []byte, error foundation.Error))
	InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler(appIdentifier []byte, options map[ContentKeySessionServerPlaybackContextOption]objc.IObject, handler func(secureTokenData []byte, error foundation.Error))
	Expire()
	ProcessContentKeyRequestWithIdentifierInitializationDataOptions(identifier objc.IObject, initializationData []byte, options map[string]objc.IObject)
	KeySystem() ContentKeySystem
	ContentProtectionSessionIdentifier() []byte
	Delegate() ContentKeySessionDelegateWrapper
	StorageURL() foundation.URL
	ContentKeyRecipients() []ContentKeyRecipientWrapper
	DelegateQueue() dispatch.Queue
}

// An object that creates and tracks decryption keys for media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession?language=objc
type ContentKeySession struct {
	objc.Object
}

func ContentKeySessionFrom(ptr unsafe.Pointer) ContentKeySession {
	return ContentKeySession{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContentKeySessionClass) ContentKeySessionWithKeySystem(keySystem ContentKeySystem) ContentKeySession {
	rv := objc.Call[ContentKeySession](cc, objc.Sel("contentKeySessionWithKeySystem:"), keySystem)
	return rv
}

// Creates a content key session to manage a collection of content decryption keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2887362-contentkeysessionwithkeysystem?language=objc
func ContentKeySession_ContentKeySessionWithKeySystem(keySystem ContentKeySystem) ContentKeySession {
	return ContentKeySessionClass.ContentKeySessionWithKeySystem(keySystem)
}

func (cc _ContentKeySessionClass) Alloc() ContentKeySession {
	rv := objc.Call[ContentKeySession](cc, objc.Sel("alloc"))
	return rv
}

func ContentKeySession_Alloc() ContentKeySession {
	return ContentKeySessionClass.Alloc()
}

func (cc _ContentKeySessionClass) New() ContentKeySession {
	rv := objc.Call[ContentKeySession](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContentKeySession() ContentKeySession {
	return ContentKeySessionClass.New()
}

func (c_ ContentKeySession) Init() ContentKeySession {
	rv := objc.Call[ContentKeySession](c_, objc.Sel("init"))
	return rv
}

// Removes expired session reports from storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799170-removependingexpiredsessionrepor?language=objc
func (cc _ContentKeySessionClass) RemovePendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL(expiredSessionReports [][]byte, appIdentifier []byte, storageURL foundation.IURL) {
	objc.Call[objc.Void](cc, objc.Sel("removePendingExpiredSessionReports:withAppIdentifier:storageDirectoryAtURL:"), expiredSessionReports, appIdentifier, objc.Ptr(storageURL))
}

// Removes expired session reports from storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799170-removependingexpiredsessionrepor?language=objc
func ContentKeySession_RemovePendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL(expiredSessionReports [][]byte, appIdentifier []byte, storageURL foundation.IURL) {
	ContentKeySessionClass.RemovePendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL(expiredSessionReports, appIdentifier, storageURL)
}

// Creates a secure server playback context that the client sends to the key server to get an expiration date for the given persistable content key data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2887476-makesecuretokenforexpirationdate?language=objc
func (c_ ContentKeySession) MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler(persistableContentKeyData []byte, handler func(secureTokenData []byte, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("makeSecureTokenForExpirationDateOfPersistableContentKey:completionHandler:"), persistableContentKeyData, handler)
}

// Tells the delegate that previously provided response data for a content key request is about to expire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799208-renewexpiringresponsedataforcont?language=objc
func (c_ ContentKeySession) RenewExpiringResponseDataForContentKeyRequest(contentKeyRequest IContentKeyRequest) {
	objc.Call[objc.Void](c_, objc.Sel("renewExpiringResponseDataForContentKeyRequest:"), objc.Ptr(contentKeyRequest))
}

// Tells the delegate that the specified recipient should have access to the decryption keys loaded with the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799174-addcontentkeyrecipient?language=objc
func (c_ ContentKeySession) AddContentKeyRecipient(recipient PContentKeyRecipient) {
	po0 := objc.WrapAsProtocol("AVContentKeyRecipient", recipient)
	objc.Call[objc.Void](c_, objc.Sel("addContentKeyRecipient:"), po0)
}

// Tells the delegate that the specified recipient should have access to the decryption keys loaded with the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799174-addcontentkeyrecipient?language=objc
func (c_ ContentKeySession) AddContentKeyRecipientObject(recipientObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("addContentKeyRecipient:"), objc.Ptr(recipientObject))
}

// Sets the session’s delegate object and the dispatch queue on which to call the delegate’s methods. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799164-setdelegate?language=objc
func (c_ ContentKeySession) SetDelegateQueue(delegate PContentKeySessionDelegate, delegateQueue dispatch.Queue) {
	po0 := objc.WrapAsProtocol("AVContentKeySessionDelegate", delegate)
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:queue:"), po0, delegateQueue)
}

// Sets the session’s delegate object and the dispatch queue on which to call the delegate’s methods. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799164-setdelegate?language=objc
func (c_ ContentKeySession) SetDelegateObjectQueue(delegateObject objc.IObject, delegateQueue dispatch.Queue) {
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:queue:"), objc.Ptr(delegateObject), delegateQueue)
}

// Tells the delegate to remove the specified recipient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799181-removecontentkeyrecipient?language=objc
func (c_ ContentKeySession) RemoveContentKeyRecipient(recipient PContentKeyRecipient) {
	po0 := objc.WrapAsProtocol("AVContentKeyRecipient", recipient)
	objc.Call[objc.Void](c_, objc.Sel("removeContentKeyRecipient:"), po0)
}

// Tells the delegate to remove the specified recipient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799181-removecontentkeyrecipient?language=objc
func (c_ ContentKeySession) RemoveContentKeyRecipientObject(recipientObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("removeContentKeyRecipient:"), objc.Ptr(recipientObject))
}

// Invalidates the persistable content key and creates a secure server playback context (SPC) to verify the outcome of an invalidation request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/3089138-invalidatepersistablecontentkey?language=objc
func (c_ ContentKeySession) InvalidatePersistableContentKeyOptionsCompletionHandler(persistableContentKeyData []byte, options map[ContentKeySessionServerPlaybackContextOption]objc.IObject, handler func(secureTokenData []byte, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("invalidatePersistableContentKey:options:completionHandler:"), persistableContentKeyData, options, handler)
}

// Returns the expired session reports for content key sessions created with the specified app identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799161-pendingexpiredsessionreportswith?language=objc
func (cc _ContentKeySessionClass) PendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL(appIdentifier []byte, storageURL foundation.IURL) [][]byte {
	rv := objc.Call[[][]byte](cc, objc.Sel("pendingExpiredSessionReportsWithAppIdentifier:storageDirectoryAtURL:"), appIdentifier, objc.Ptr(storageURL))
	return rv
}

// Returns the expired session reports for content key sessions created with the specified app identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799161-pendingexpiredsessionreportswith?language=objc
func ContentKeySession_PendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL(appIdentifier []byte, storageURL foundation.IURL) [][]byte {
	return ContentKeySessionClass.PendingExpiredSessionReportsWithAppIdentifierStorageDirectoryAtURL(appIdentifier, storageURL)
}

// Invalidates all of an app’s persistable content keys and creates a secure server playback context (SPC) to verify the outcome of an invalidation request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/3089137-invalidateallpersistablecontentk?language=objc
func (c_ ContentKeySession) InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler(appIdentifier []byte, options map[ContentKeySessionServerPlaybackContextOption]objc.IObject, handler func(secureTokenData []byte, error foundation.Error)) {
	objc.Call[objc.Void](c_, objc.Sel("invalidateAllPersistableContentKeysForApp:options:completionHandler:"), appIdentifier, options, handler)
}

// Tells the delegate that the session expired as the result of normal, intentional processes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799171-expire?language=objc
func (c_ ContentKeySession) Expire() {
	objc.Call[objc.Void](c_, objc.Sel("expire"))
}

// Tells the delegate to start loading the content decryption key with the specified identifier and initialization data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799180-processcontentkeyrequestwithiden?language=objc
func (c_ ContentKeySession) ProcessContentKeyRequestWithIdentifierInitializationDataOptions(identifier objc.IObject, initializationData []byte, options map[string]objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("processContentKeyRequestWithIdentifier:initializationData:options:"), identifier, initializationData, options)
}

// The type of key system used to retrieve keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799167-keysystem?language=objc
func (c_ ContentKeySession) KeySystem() ContentKeySystem {
	rv := objc.Call[ContentKeySystem](c_, objc.Sel("keySystem"))
	return rv
}

// The identifier for the current content protection session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799189-contentprotectionsessionidentifi?language=objc
func (c_ ContentKeySession) ContentProtectionSessionIdentifier() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("contentProtectionSessionIdentifier"))
	return rv
}

// The content key session’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799191-delegate?language=objc
func (c_ ContentKeySession) Delegate() ContentKeySessionDelegateWrapper {
	rv := objc.Call[ContentKeySessionDelegateWrapper](c_, objc.Sel("delegate"))
	return rv
}

// A URL that points to a writable storage directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799203-storageurl?language=objc
func (c_ ContentKeySession) StorageURL() foundation.URL {
	rv := objc.Call[foundation.URL](c_, objc.Sel("storageURL"))
	return rv
}

// An array of content key recipients. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799192-contentkeyrecipients?language=objc
func (c_ ContentKeySession) ContentKeyRecipients() []ContentKeyRecipientWrapper {
	rv := objc.Call[[]ContentKeyRecipientWrapper](c_, objc.Sel("contentKeyRecipients"))
	return rv
}

// The dispatch queue the session uses to invoke delegate callbacks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeysession/2799178-delegatequeue?language=objc
func (c_ ContentKeySession) DelegateQueue() dispatch.Queue {
	rv := objc.Call[dispatch.Queue](c_, objc.Sel("delegateQueue"))
	return rv
}
