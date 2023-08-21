// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Document] class.
var DocumentClass = _DocumentClass{objc.GetClass("NSDocument")}

type _DocumentClass struct {
	objc.Class
}

// An interface definition for the [Document] class.
type IDocument interface {
	objc.IObject
	LockWithCompletionHandler(completionHandler func(arg0 foundation.Error))
	RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo(saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer)
	FileWrapperOfTypeError(typeName string, outError foundation.IError) foundation.FileWrapper
	PrintDocument(sender objc.IObject) objc.Object
	RemoveWindowController(windowController IWindowController)
	RevertDocumentToSaved(sender objc.IObject) objc.Object
	DuplicateAndReturnError(outError foundation.IError) Document
	ShowWindows()
	WillPresentError(error foundation.IError) foundation.Error
	ContinueAsynchronousWorkOnMainThreadUsingBlock(block func())
	SetWindow(window IWindow)
	WillNotPresentError(error foundation.IError)
	DataOfTypeError(typeName string, outError foundation.IError) []byte
	ScheduleAutosaving()
	PrintOperationWithSettingsError(printSettings map[PrintInfoAttributeKey]objc.IObject, outError foundation.IError) PrintOperation
	ContinueActivityUsingBlock(block func())
	InvalidateRestorableState()
	HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object
	SaveDocumentAs(sender objc.IObject) objc.Object
	LockDocument(sender objc.IObject) objc.Object
	Close()
	ShouldChangePrintInfo(newPrintInfo IPrintInfo) bool
	HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object
	PreparePageLayout(pageLayout IPageLayout) bool
	PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block func(arg0 func()))
	LockDocumentWithCompletionHandler(completionHandler func(didLock bool))
	UnblockUserInteraction()
	SaveDocument(sender objc.IObject) objc.Object
	WindowControllerWillLoadNib(windowController IWindowController)
	MoveDocumentToUbiquityContainer(sender objc.IObject) objc.Object
	PerformSynchronousFileAccessUsingBlock(block func())
	WriteSafelyToURLOfTypeForSaveOperationError(url foundation.IURL, typeName string, saveOperation SaveOperationType, outError foundation.IError) bool
	SaveDocumentTo(sender objc.IObject) objc.Object
	BrowseDocumentVersions(sender objc.IObject) objc.Object
	SaveDocumentToPDF(sender objc.IObject) objc.Object
	WritableTypesForSaveOperation(saveOperation SaveOperationType) []string
	DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objc.IObject, didDuplicateSelector objc.Selector, contextInfo unsafe.Pointer)
	ReadFromURLOfTypeError(url foundation.IURL, typeName string, outError foundation.IError) bool
	AutosaveWithImplicitCancellabilityCompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler func(errorOrNil foundation.Error))
	DefaultDraftName() string
	AddWindowController(windowController IWindowController)
	HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object
	RestoreStateWithCoder(coder foundation.ICoder)
	SetDisplayName(displayNameOrNil string)
	CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo(delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer)
	RestoreDocumentWindowWithIdentifierStateCompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(arg0 Window, arg1 foundation.Error))
	ValidateUserInterfaceItem(item PValidatedUserInterfaceItem) bool
	ValidateUserInterfaceItemObject(itemObject objc.IObject) bool
	ReadFromDataOfTypeError(data []byte, typeName string, outError foundation.IError) bool
	RunPageLayout(sender objc.IObject) objc.Object
	UpdateUserActivityState(activity foundation.IUserActivity)
	UnlockWithCompletionHandler(completionHandler func(arg0 foundation.Error))
	MoveDocumentWithCompletionHandler(completionHandler func(didMove bool))
	PerformAsynchronousFileAccessUsingBlock(block func(arg0 func()))
	CanAsynchronouslyWriteToURLOfTypeForSaveOperation(url foundation.IURL, typeName string, saveOperation SaveOperationType) bool
	ChangeCountTokenForSaveOperation(saveOperation SaveOperationType) objc.Object
	FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError foundation.IError) map[string]objc.Object
	UnlockDocumentWithCompletionHandler(completionHandler func(didUnlock bool))
	FileNameExtensionForTypeSaveOperation(typeName string, saveOperation SaveOperationType) string
	RenameDocument(sender objc.IObject) objc.Object
	ShareDocumentWithSharingServiceCompletionHandler(sharingService ISharingService, completionHandler func(success bool))
	ReadFromFileWrapperOfTypeError(fileWrapper foundation.IFileWrapper, typeName string, outError foundation.IError) bool
	WriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError foundation.IError) bool
	StopBrowsingVersionsWithCompletionHandler(completionHandler func())
	PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo(printSettings map[PrintInfoAttributeKey]objc.IObject, showPrintPanel bool, delegate objc.IObject, didPrintSelector objc.Selector, contextInfo unsafe.Pointer)
	DuplicateDocument(sender objc.IObject) objc.Object
	MakeWindowControllers()
	UnlockDocument(sender objc.IObject) objc.Object
	RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation IPrintOperation, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer)
	SaveDocumentWithDelegateDidSaveSelectorContextInfo(delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer)
	RevertToContentsOfURLOfTypeError(url foundation.IURL, typeName string, outError foundation.IError) bool
	CheckAutosavingSafetyAndReturnError(outError foundation.IError) bool
	WindowControllerDidLoadNib(windowController IWindowController)
	ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo(windowController IWindowController, delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer)
	MoveToURLCompletionHandler(url foundation.IURL, completionHandler func(arg0 foundation.Error))
	AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objc.IObject, didAutosaveSelector objc.Selector, contextInfo unsafe.Pointer)
	PrepareSavePanel(savePanel ISavePanel) bool
	PresentError(error foundation.IError) bool
	MoveDocument(sender objc.IObject) objc.Object
	UpdateChangeCountWithTokenForSaveOperation(changeCountToken objc.IObject, saveOperation SaveOperationType)
	EncodeRestorableStateWithCoder(coder foundation.ICoder)
	UpdateChangeCount(change DocumentChangeType)
	PrepareSharingServicePicker(sharingServicePicker ISharingServicePicker)
	PrintInfo() PrintInfo
	SetPrintInfo(value IPrintInfo)
	ObjectSpecifier() foundation.ScriptObjectSpecifier
	IsDraft() bool
	SetDraft(value bool)
	FileNameExtensionWasHiddenInLastRunSavePanel() bool
	UndoManager() foundation.UndoManager
	SetUndoManager(value foundation.IUndoManager)
	AutosavedContentsFileURL() foundation.URL
	SetAutosavedContentsFileURL(value foundation.IURL)
	ShouldRunSavePanelWithAccessoryView() bool
	LastComponentOfFileName() string
	SetLastComponentOfFileName(value string)
	KeepBackupFile() bool
	BackupFileURL() foundation.URL
	FileURL() foundation.URL
	SetFileURL(value foundation.IURL)
	FileTypeFromLastRunSavePanel() string
	WindowControllers() []WindowController
	IsEntireFileLoaded() bool
	IsBrowsingVersions() bool
	IsInViewingMode() bool
	PDFPrintOperation() PrintOperation
	HasUndoManager() bool
	SetHasUndoManager(value bool)
	FileType() string
	SetFileType(value string)
	IsDocumentEdited() bool
	IsLocked() bool
	FileModificationDate() foundation.Date
	SetFileModificationDate(value foundation.IDate)
	WindowNibName() NibName
	AutosavingIsImplicitlyCancellable() bool
	AllowsDocumentSharing() bool
	HasUnautosavedChanges() bool
	AutosavingFileType() string
	WindowForSheet() Window
	DisplayName() string
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.IUserActivity)
}

// An abstract class that defines the interface for macOS documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument?language=objc
type Document struct {
	objc.Object
}

func DocumentFrom(ptr unsafe.Pointer) Document {
	return Document{
		Object: objc.ObjectFrom(ptr),
	}
}

func (d_ Document) InitWithTypeError(typeName string, outError foundation.IError) Document {
	rv := objc.Call[Document](d_, objc.Sel("initWithType:error:"), typeName, objc.Ptr(outError))
	return rv
}

// Initializes a document of a specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515159-initwithtype?language=objc
func NewDocumentWithTypeError(typeName string, outError foundation.IError) Document {
	instance := DocumentClass.Alloc().InitWithTypeError(typeName, outError)
	instance.Autorelease()
	return instance
}

func (d_ Document) Init() Document {
	rv := objc.Call[Document](d_, objc.Sel("init"))
	return rv
}

func (d_ Document) InitForURLWithContentsOfURLOfTypeError(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError foundation.IError) Document {
	rv := objc.Call[Document](d_, objc.Sel("initForURL:withContentsOfURL:ofType:error:"), objc.Ptr(urlOrNil), objc.Ptr(contentsURL), typeName, objc.Ptr(outError))
	return rv
}

// Initializes a document with the specified contents, and places the resulting document's file at the designated location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515041-initforurl?language=objc
func NewDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError foundation.IError) Document {
	instance := DocumentClass.Alloc().InitForURLWithContentsOfURLOfTypeError(urlOrNil, contentsURL, typeName, outError)
	instance.Autorelease()
	return instance
}

func (dc _DocumentClass) Alloc() Document {
	rv := objc.Call[Document](dc, objc.Sel("alloc"))
	return rv
}

func Document_Alloc() Document {
	return DocumentClass.Alloc()
}

func (dc _DocumentClass) New() Document {
	rv := objc.Call[Document](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDocument() Document {
	return DocumentClass.New()
}

// Prevents the user from making changes to the document's file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515189-lockwithcompletionhandler?language=objc
func (d_ Document) LockWithCompletionHandler(completionHandler func(arg0 foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("lockWithCompletionHandler:"), completionHandler)
}

// Presents a modal Save panel to the user, then tries to save the document if the user approves the operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515180-runmodalsavepanelforsaveoperatio?language=objc
func (d_ Document) RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo(saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](d_, objc.Sel("runModalSavePanelForSaveOperation:delegate:didSaveSelector:contextInfo:"), saveOperation, delegate, didSaveSelector, contextInfo)
}

// Returns a Boolean value that indicates whether the document can read and write the data natively. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515072-isnativetype?language=objc
func (dc _DocumentClass) IsNativeType(type_ string) bool {
	rv := objc.Call[bool](dc, objc.Sel("isNativeType:"), type_)
	return rv
}

// Returns a Boolean value that indicates whether the document can read and write the data natively. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515072-isnativetype?language=objc
func Document_IsNativeType(type_ string) bool {
	return DocumentClass.IsNativeType(type_)
}

// Returns the classes that support secure coding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/3762522-allowedclassesforrestorablestate?language=objc
func (dc _DocumentClass) AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	rv := objc.Call[[]objc.Class](dc, objc.Sel("allowedClassesForRestorableStateKeyPath:"), keyPath)
	return rv
}

// Returns the classes that support secure coding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/3762522-allowedclassesforrestorablestate?language=objc
func Document_AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	return DocumentClass.AllowedClassesForRestorableStateKeyPath(keyPath)
}

// Creates and returns a file wrapper that contains the contents of the document, formatted to the specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515089-filewrapperoftype?language=objc
func (d_ Document) FileWrapperOfTypeError(typeName string, outError foundation.IError) foundation.FileWrapper {
	rv := objc.Call[foundation.FileWrapper](d_, objc.Sel("fileWrapperOfType:error:"), typeName, objc.Ptr(outError))
	return rv
}

// Prints the receiver in response to the user choosing the Print menu command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515154-printdocument?language=objc
func (d_ Document) PrintDocument(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("printDocument:"), sender)
	return rv
}

// Removes the specified window controller from the receiver’s array of window controllers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515242-removewindowcontroller?language=objc
func (d_ Document) RemoveWindowController(windowController IWindowController) {
	objc.Call[objc.Void](d_, objc.Sel("removeWindowController:"), objc.Ptr(windowController))
}

// The action of the File menu item Revert in a document-based app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515059-revertdocumenttosaved?language=objc
func (d_ Document) RevertDocumentToSaved(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("revertDocumentToSaved:"), sender)
	return rv
}

// Creates a new document whose contents are the same as the receiver and returns an error object if unsuccessful. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515201-duplicateandreturnerror?language=objc
func (d_ Document) DuplicateAndReturnError(outError foundation.IError) Document {
	rv := objc.Call[Document](d_, objc.Sel("duplicateAndReturnError:"), objc.Ptr(outError))
	return rv
}

// Displays all of the document’s windows, bringing them to the front and making them main or key as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515049-showwindows?language=objc
func (d_ Document) ShowWindows() {
	objc.Call[objc.Void](d_, objc.Sel("showWindows"))
}

// Called when the receiver is about to present an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515229-willpresenterror?language=objc
func (d_ Document) WillPresentError(error foundation.IError) foundation.Error {
	rv := objc.Call[foundation.Error](d_, objc.Sel("willPresentError:"), objc.Ptr(error))
	return rv
}

// Invokes the passed-in block on the main thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515069-continueasynchronousworkonmainth?language=objc
func (d_ Document) ContinueAsynchronousWorkOnMainThreadUsingBlock(block func()) {
	objc.Call[objc.Void](d_, objc.Sel("continueAsynchronousWorkOnMainThreadUsingBlock:"), block)
}

// Sets the window outlet of this document to the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515217-setwindow?language=objc
func (d_ Document) SetWindow(window IWindow) {
	objc.Call[objc.Void](d_, objc.Sel("setWindow:"), objc.Ptr(window))
}

// Confirms that the error object is not to be presented to the user and the error cannot be recovered from, so cleanup can be done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515188-willnotpresenterror?language=objc
func (d_ Document) WillNotPresentError(error foundation.IError) {
	objc.Call[objc.Void](d_, objc.Sel("willNotPresentError:"), objc.Ptr(error))
}

// Creates and returns a data object that contains the contents of the document, formatted to a specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515205-dataoftype?language=objc
func (d_ Document) DataOfTypeError(typeName string, outError foundation.IError) []byte {
	rv := objc.Call[[]byte](d_, objc.Sel("dataOfType:error:"), typeName, objc.Ptr(outError))
	return rv
}

// Schedules periodic autosaving for the purpose of crash protection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515119-scheduleautosaving?language=objc
func (d_ Document) ScheduleAutosaving() {
	objc.Call[objc.Void](d_, objc.Sel("scheduleAutosaving"))
}

// Creates and returns a print operation for the document's contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515070-printoperationwithsettings?language=objc
func (d_ Document) PrintOperationWithSettingsError(printSettings map[PrintInfoAttributeKey]objc.IObject, outError foundation.IError) PrintOperation {
	rv := objc.Call[PrintOperation](d_, objc.Sel("printOperationWithSettings:error:"), printSettings, objc.Ptr(outError))
	return rv
}

// Continues to perform the task for a user activity object using a different block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515151-continueactivityusingblock?language=objc
func (d_ Document) ContinueActivityUsingBlock(block func()) {
	objc.Call[objc.Void](d_, objc.Sel("continueActivityUsingBlock:"), block)
}

// Marks the document’s interface-related state as dirty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1526250-invalidaterestorablestate?language=objc
func (d_ Document) InvalidateRestorableState() {
	objc.Call[objc.Void](d_, objc.Sel("invalidateRestorableState"))
}

// Handles the Save AppleScript command by attempting to save the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1500138-handlesavescriptcommand?language=objc
func (d_ Document) HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("handleSaveScriptCommand:"), objc.Ptr(command))
	return rv
}

// The action method invoked in the receiver as first responder when the user chooses the Save As menu command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515171-savedocumentas?language=objc
func (d_ Document) SaveDocumentAs(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("saveDocumentAs:"), sender)
	return rv
}

// Locks the document in response to the user choosing the Lock menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515218-lockdocument?language=objc
func (d_ Document) LockDocument(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("lockDocument:"), sender)
	return rv
}

// Closes all of the document's windows and removes the document from its document controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515237-close?language=objc
func (d_ Document) Close() {
	objc.Call[objc.Void](d_, objc.Sel("close"))
}

// Returns a Boolean value that indicates whether the document allows changes to the default printing information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515243-shouldchangeprintinfo?language=objc
func (d_ Document) ShouldChangePrintInfo(newPrintInfo IPrintInfo) bool {
	rv := objc.Call[bool](d_, objc.Sel("shouldChangePrintInfo:"), objc.Ptr(newPrintInfo))
	return rv
}

// Handles the Close AppleScript command by attempting to close the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1500136-handleclosescriptcommand?language=objc
func (d_ Document) HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("handleCloseScriptCommand:"), objc.Ptr(command))
	return rv
}

// Adds document-specific content to the Page Layout panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515169-preparepagelayout?language=objc
func (d_ Document) PreparePageLayout(pageLayout IPageLayout) bool {
	rv := objc.Call[bool](d_, objc.Sel("preparePageLayout:"), objc.Ptr(pageLayout))
	return rv
}

// Waits for any work scheduled by previous invocations of this method to complete, then invokes the passed-in block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515066-performactivitywithsynchronouswa?language=objc
func (d_ Document) PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block func(arg0 func())) {
	objc.Call[objc.Void](d_, objc.Sel("performActivityWithSynchronousWaiting:usingBlock:"), waitSynchronously, block)
}

// Prevents the user from making further changes to the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515233-lockdocumentwithcompletionhandle?language=objc
func (d_ Document) LockDocumentWithCompletionHandler(completionHandler func(didLock bool)) {
	objc.Call[objc.Void](d_, objc.Sel("lockDocumentWithCompletionHandler:"), completionHandler)
}

// Unblocks the main thread during asynchronous saving. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515238-unblockuserinteraction?language=objc
func (d_ Document) UnblockUserInteraction() {
	objc.Call[objc.Void](d_, objc.Sel("unblockUserInteraction"))
}

// The action method invoked in the receiver as first responder when the user chooses the Save menu command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515147-savedocument?language=objc
func (d_ Document) SaveDocument(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("saveDocument:"), sender)
	return rv
}

// Called before one of the document's window controllers loads its nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515116-windowcontrollerwillloadnib?language=objc
func (d_ Document) WindowControllerWillLoadNib(windowController IWindowController) {
	objc.Call[objc.Void](d_, objc.Sel("windowControllerWillLoadNib:"), objc.Ptr(windowController))
}

// Moves the document to the user’s iCloud storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515210-movedocumenttoubiquitycontainer?language=objc
func (d_ Document) MoveDocumentToUbiquityContainer(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("moveDocumentToUbiquityContainer:"), sender)
	return rv
}

// Waits for any scheduled file access to complete, then invokes the passed-in block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515227-performsynchronousfileaccessusin?language=objc
func (d_ Document) PerformSynchronousFileAccessUsingBlock(block func()) {
	objc.Call[objc.Void](d_, objc.Sel("performSynchronousFileAccessUsingBlock:"), block)
}

// Writes the contents of the document to a file or file package located by a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515150-writesafelytourl?language=objc
func (d_ Document) WriteSafelyToURLOfTypeForSaveOperationError(url foundation.IURL, typeName string, saveOperation SaveOperationType, outError foundation.IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("writeSafelyToURL:ofType:forSaveOperation:error:"), objc.Ptr(url), typeName, saveOperation, objc.Ptr(outError))
	return rv
}

// The action method invoked in the receiver as first responder when the user chooses the Save To menu command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515208-savedocumentto?language=objc
func (d_ Document) SaveDocumentTo(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("saveDocumentTo:"), sender)
	return rv
}

// Opens the Versions browser in the document’s main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515193-browsedocumentversions?language=objc
func (d_ Document) BrowseDocumentVersions(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("browseDocumentVersions:"), sender)
	return rv
}

// Exports a PDF representation of the document’s current contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515176-savedocumenttopdf?language=objc
func (d_ Document) SaveDocumentToPDF(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("saveDocumentToPDF:"), sender)
	return rv
}

// Returns the names of the types to which this document can be saved for a specified kind of save operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515186-writabletypesforsaveoperation?language=objc
func (d_ Document) WritableTypesForSaveOperation(saveOperation SaveOperationType) []string {
	rv := objc.Call[[]string](d_, objc.Sel("writableTypesForSaveOperation:"), saveOperation)
	return rv
}

// Creates a new document whose contents are the same as the current document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515133-duplicatedocumentwithdelegate?language=objc
func (d_ Document) DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objc.IObject, didDuplicateSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](d_, objc.Sel("duplicateDocumentWithDelegate:didDuplicateSelector:contextInfo:"), delegate, didDuplicateSelector, contextInfo)
}

// Sets the contents of this document by reading from a file or file package, of a specified type, located by a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515144-readfromurl?language=objc
func (d_ Document) ReadFromURLOfTypeError(url foundation.IURL, typeName string, outError foundation.IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("readFromURL:ofType:error:"), objc.Ptr(url), typeName, objc.Ptr(outError))
	return rv
}

// Autosaves the document’s contents to an appropriate file-system location, as needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515096-autosavewithimplicitcancellabili?language=objc
func (d_ Document) AutosaveWithImplicitCancellabilityCompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler func(errorOrNil foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("autosaveWithImplicitCancellability:completionHandler:"), autosavingIsImplicitlyCancellable, completionHandler)
}

// Returns the default draft name for the document subclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515245-defaultdraftname?language=objc
func (d_ Document) DefaultDraftName() string {
	rv := objc.Call[string](d_, objc.Sel("defaultDraftName"))
	return rv
}

// Adds the specified window controller to the current document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515179-addwindowcontroller?language=objc
func (d_ Document) AddWindowController(windowController IWindowController) {
	objc.Call[objc.Void](d_, objc.Sel("addWindowController:"), objc.Ptr(windowController))
}

// Handles the Print AppleScript command by attempting to print the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1500135-handleprintscriptcommand?language=objc
func (d_ Document) HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("handlePrintScriptCommand:"), objc.Ptr(command))
	return rv
}

// Restores the interface-related state of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1526237-restorestatewithcoder?language=objc
func (d_ Document) RestoreStateWithCoder(coder foundation.ICoder) {
	objc.Call[objc.Void](d_, objc.Sel("restoreStateWithCoder:"), objc.Ptr(coder))
}

// Sets the name of this document for presentation to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515143-setdisplayname?language=objc
func (d_ Document) SetDisplayName(displayNameOrNil string) {
	objc.Call[objc.Void](d_, objc.Sel("setDisplayName:"), displayNameOrNil)
}

// Determines whether to close the document, prompting the user as needed to choose a course of action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515206-canclosedocumentwithdelegate?language=objc
func (d_ Document) CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo(delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](d_, objc.Sel("canCloseDocumentWithDelegate:shouldCloseSelector:contextInfo:"), delegate, shouldCloseSelector, contextInfo)
}

// Restores a window that was associated with a document, after that document is reopened. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1524586-restoredocumentwindowwithidentif?language=objc
func (d_ Document) RestoreDocumentWindowWithIdentifierStateCompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(arg0 Window, arg1 foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("restoreDocumentWindowWithIdentifier:state:completionHandler:"), identifier, objc.Ptr(state), completionHandler)
}

// Validates the specified user interface item that the receiver manages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515190-validateuserinterfaceitem?language=objc
func (d_ Document) ValidateUserInterfaceItem(item PValidatedUserInterfaceItem) bool {
	po0 := objc.WrapAsProtocol("NSValidatedUserInterfaceItem", item)
	rv := objc.Call[bool](d_, objc.Sel("validateUserInterfaceItem:"), po0)
	return rv
}

// Validates the specified user interface item that the receiver manages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515190-validateuserinterfaceitem?language=objc
func (d_ Document) ValidateUserInterfaceItemObject(itemObject objc.IObject) bool {
	rv := objc.Call[bool](d_, objc.Sel("validateUserInterfaceItem:"), objc.Ptr(itemObject))
	return rv
}

// Sets the contents of this document by reading from data of a specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515198-readfromdata?language=objc
func (d_ Document) ReadFromDataOfTypeError(data []byte, typeName string, outError foundation.IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("readFromData:ofType:error:"), data, typeName, objc.Ptr(outError))
	return rv
}

// The action method invoked in the receiver as first responder when the user chooses the Page Setup menu command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515140-runpagelayout?language=objc
func (d_ Document) RunPageLayout(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("runPageLayout:"), sender)
	return rv
}

// Updates the state of the given user activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1529014-updateuseractivitystate?language=objc
func (d_ Document) UpdateUserActivityState(activity foundation.IUserActivity) {
	objc.Call[objc.Void](d_, objc.Sel("updateUserActivityState:"), objc.Ptr(activity))
}

// Allows the user to make modifications to the document's file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515131-unlockwithcompletionhandler?language=objc
func (d_ Document) UnlockWithCompletionHandler(completionHandler func(arg0 foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("unlockWithCompletionHandler:"), completionHandler)
}

// Moves the document to a user-selected location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515043-movedocumentwithcompletionhandle?language=objc
func (d_ Document) MoveDocumentWithCompletionHandler(completionHandler func(didMove bool)) {
	objc.Call[objc.Void](d_, objc.Sel("moveDocumentWithCompletionHandler:"), completionHandler)
}

// Waits for any scheduled file access to complete but without blocking the main thread, then invokes the passed-in block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515124-performasynchronousfileaccessusi?language=objc
func (d_ Document) PerformAsynchronousFileAccessUsingBlock(block func(arg0 func())) {
	objc.Call[objc.Void](d_, objc.Sel("performAsynchronousFileAccessUsingBlock:"), block)
}

// Returns whether the receiver can concurrently write to a file or file package located by a URL, that is formatted for a specific type, for a specific kind of save operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515177-canasynchronouslywritetourl?language=objc
func (d_ Document) CanAsynchronouslyWriteToURLOfTypeForSaveOperation(url foundation.IURL, typeName string, saveOperation SaveOperationType) bool {
	rv := objc.Call[bool](d_, objc.Sel("canAsynchronouslyWriteToURL:ofType:forSaveOperation:"), objc.Ptr(url), typeName, saveOperation)
	return rv
}

// Returns an object that encapsulates the current record of document changes at the beginning of a save operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515129-changecounttokenforsaveoperation?language=objc
func (d_ Document) ChangeCountTokenForSaveOperation(saveOperation SaveOperationType) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("changeCountTokenForSaveOperation:"), saveOperation)
	return rv
}

// Returns the attributes to write to the file or file package at the specified URL, and targeting the specified type of save operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515062-fileattributestowritetourl?language=objc
func (d_ Document) FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError foundation.IError) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](d_, objc.Sel("fileAttributesToWriteToURL:ofType:forSaveOperation:originalContentsURL:error:"), objc.Ptr(url), typeName, saveOperation, objc.Ptr(absoluteOriginalContentsURL), objc.Ptr(outError))
	return rv
}

// Allows the user to make modifications to the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515248-unlockdocumentwithcompletionhand?language=objc
func (d_ Document) UnlockDocumentWithCompletionHandler(completionHandler func(didUnlock bool)) {
	objc.Call[objc.Void](d_, objc.Sel("unlockDocumentWithCompletionHandler:"), completionHandler)
}

// Returns a filename extension that can be appended to a base filename, for a specified file type and kind of save operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515074-filenameextensionfortype?language=objc
func (d_ Document) FileNameExtensionForTypeSaveOperation(typeName string, saveOperation SaveOperationType) string {
	rv := objc.Call[string](d_, objc.Sel("fileNameExtensionForType:saveOperation:"), typeName, saveOperation)
	return rv
}

// Renames the current document in response to the user choosing the Rename menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515231-renamedocument?language=objc
func (d_ Document) RenameDocument(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("renameDocument:"), sender)
	return rv
}

// Share the document's file using the specified sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/2902309-sharedocumentwithsharingservice?language=objc
func (d_ Document) ShareDocumentWithSharingServiceCompletionHandler(sharingService ISharingService, completionHandler func(success bool)) {
	objc.Call[objc.Void](d_, objc.Sel("shareDocumentWithSharingService:completionHandler:"), objc.Ptr(sharingService), completionHandler)
}

// Sets the contents of this document by reading from a file wrapper of a specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515044-readfromfilewrapper?language=objc
func (d_ Document) ReadFromFileWrapperOfTypeError(fileWrapper foundation.IFileWrapper, typeName string, outError foundation.IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("readFromFileWrapper:ofType:error:"), objc.Ptr(fileWrapper), typeName, objc.Ptr(outError))
	return rv
}

// Writes the contents of the document to a file or file package located by a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515203-writetourl?language=objc
func (d_ Document) WriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.IURL, typeName string, saveOperation SaveOperationType, absoluteOriginalContentsURL foundation.IURL, outError foundation.IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("writeToURL:ofType:forSaveOperation:originalContentsURL:error:"), objc.Ptr(url), typeName, saveOperation, objc.Ptr(absoluteOriginalContentsURL), objc.Ptr(outError))
	return rv
}

// Dismiss the Versions browser for the current document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/2177312-stopbrowsingversionswithcompleti?language=objc
func (d_ Document) StopBrowsingVersionsWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](d_, objc.Sel("stopBrowsingVersionsWithCompletionHandler:"), completionHandler)
}

// Prints the document's contents, optionally displaying a print panel to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515058-printdocumentwithsettings?language=objc
func (d_ Document) PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo(printSettings map[PrintInfoAttributeKey]objc.IObject, showPrintPanel bool, delegate objc.IObject, didPrintSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](d_, objc.Sel("printDocumentWithSettings:showPrintPanel:delegate:didPrintSelector:contextInfo:"), printSettings, showPrintPanel, delegate, didPrintSelector, contextInfo)
}

// Creates a copy of the receiving document in response to the user choosing Duplicate from the File menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515226-duplicatedocument?language=objc
func (d_ Document) DuplicateDocument(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("duplicateDocument:"), sender)
	return rv
}

// Creates the window controller objects that the document uses to display its content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515220-makewindowcontrollers?language=objc
func (d_ Document) MakeWindowControllers() {
	objc.Call[objc.Void](d_, objc.Sel("makeWindowControllers"))
}

// Unlocks the document in response to the user choosing the Unlock menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515068-unlockdocument?language=objc
func (d_ Document) UnlockDocument(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("unlockDocument:"), sender)
	return rv
}

// Runs the specified print operation modally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515234-runmodalprintoperation?language=objc
func (d_ Document) RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation IPrintOperation, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](d_, objc.Sel("runModalPrintOperation:delegate:didRunSelector:contextInfo:"), objc.Ptr(printOperation), delegate, didRunSelector, contextInfo)
}

// Saves the document and delivers the results to the provided delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515048-savedocumentwithdelegate?language=objc
func (d_ Document) SaveDocumentWithDelegateDidSaveSelectorContextInfo(delegate objc.IObject, didSaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](d_, objc.Sel("saveDocumentWithDelegate:didSaveSelector:contextInfo:"), delegate, didSaveSelector, contextInfo)
}

// Discards all unsaved document modifications and replaces the document’s contents by reading a file or file package located by a URL of a specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515122-reverttocontentsofurl?language=objc
func (d_ Document) RevertToContentsOfURLOfTypeError(url foundation.IURL, typeName string, outError foundation.IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("revertToContentsOfURL:ofType:error:"), objc.Ptr(url), typeName, objc.Ptr(outError))
	return rv
}

// Returns a Boolean value that indicates whether it is safe to autosave document changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515061-checkautosavingsafetyandreturner?language=objc
func (d_ Document) CheckAutosavingSafetyAndReturnError(outError foundation.IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("checkAutosavingSafetyAndReturnError:"), objc.Ptr(outError))
	return rv
}

// Called after one of the document's window controllers loads its nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515221-windowcontrollerdidloadnib?language=objc
func (d_ Document) WindowControllerDidLoadNib(windowController IWindowController) {
	objc.Call[objc.Void](d_, objc.Sel("windowControllerDidLoadNib:"), objc.Ptr(windowController))
}

// Determines whether the system should close the document and its associated window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515247-shouldclosewindowcontroller?language=objc
func (d_ Document) ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo(windowController IWindowController, delegate objc.IObject, shouldCloseSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](d_, objc.Sel("shouldCloseWindowController:delegate:shouldCloseSelector:contextInfo:"), objc.Ptr(windowController), delegate, shouldCloseSelector, contextInfo)
}

// Moves the document’s file to the given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515057-movetourl?language=objc
func (d_ Document) MoveToURLCompletionHandler(url foundation.IURL, completionHandler func(arg0 foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("moveToURL:completionHandler:"), objc.Ptr(url), completionHandler)
}

// Autosaves the document’s contents to an appropriate location in the file system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515225-autosavedocumentwithdelegate?language=objc
func (d_ Document) AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objc.IObject, didAutosaveSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](d_, objc.Sel("autosaveDocumentWithDelegate:didAutosaveSelector:contextInfo:"), delegate, didAutosaveSelector, contextInfo)
}

// Tells the document to customize the specified Save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515094-preparesavepanel?language=objc
func (d_ Document) PrepareSavePanel(savePanel ISavePanel) bool {
	rv := objc.Call[bool](d_, objc.Sel("prepareSavePanel:"), objc.Ptr(savePanel))
	return rv
}

// Returns a Boolean value that indicates whether the receiver reads multiple documents of the given type concurrently. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515216-canconcurrentlyreaddocumentsofty?language=objc
func (dc _DocumentClass) CanConcurrentlyReadDocumentsOfType(typeName string) bool {
	rv := objc.Call[bool](dc, objc.Sel("canConcurrentlyReadDocumentsOfType:"), typeName)
	return rv
}

// Returns a Boolean value that indicates whether the receiver reads multiple documents of the given type concurrently. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515216-canconcurrentlyreaddocumentsofty?language=objc
func Document_CanConcurrentlyReadDocumentsOfType(typeName string) bool {
	return DocumentClass.CanConcurrentlyReadDocumentsOfType(typeName)
}

// Presents an error alert to the user as a modal panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515184-presenterror?language=objc
func (d_ Document) PresentError(error foundation.IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("presentError:"), objc.Ptr(error))
	return rv
}

// Moves the document to a new location in response to the user choosing the Move To… menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515118-movedocument?language=objc
func (d_ Document) MoveDocument(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("moveDocument:"), sender)
	return rv
}

// Updates the document's change count settings after a successful save operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515083-updatechangecountwithtoken?language=objc
func (d_ Document) UpdateChangeCountWithTokenForSaveOperation(changeCountToken objc.IObject, saveOperation SaveOperationType) {
	objc.Call[objc.Void](d_, objc.Sel("updateChangeCountWithToken:forSaveOperation:"), changeCountToken, saveOperation)
}

// Saves the interface-related state of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1526257-encoderestorablestatewithcoder?language=objc
func (d_ Document) EncodeRestorableStateWithCoder(coder foundation.ICoder) {
	objc.Call[objc.Void](d_, objc.Sel("encodeRestorableStateWithCoder:"), objc.Ptr(coder))
}

// Updates the receiver’s change count according to the given change type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515223-updatechangecount?language=objc
func (d_ Document) UpdateChangeCount(change DocumentChangeType) {
	objc.Call[objc.Void](d_, objc.Sel("updateChangeCount:"), change)
}

// Perform any custom setup associated with a sharing service picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/2902326-preparesharingservicepicker?language=objc
func (d_ Document) PrepareSharingServicePicker(sharingServicePicker ISharingServicePicker) {
	objc.Call[objc.Void](d_, objc.Sel("prepareSharingServicePicker:"), objc.Ptr(sharingServicePicker))
}

// The printing information associated with the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515163-printinfo?language=objc
func (d_ Document) PrintInfo() PrintInfo {
	rv := objc.Call[PrintInfo](d_, objc.Sel("printInfo"))
	return rv
}

// The printing information associated with the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515163-printinfo?language=objc
func (d_ Document) SetPrintInfo(value IPrintInfo) {
	objc.Call[objc.Void](d_, objc.Sel("setPrintInfo:"), objc.Ptr(value))
}

// Returns the object specifier that represents the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1500134-objectspecifier?language=objc
func (d_ Document) ObjectSpecifier() foundation.ScriptObjectSpecifier {
	rv := objc.Call[foundation.ScriptObjectSpecifier](d_, objc.Sel("objectSpecifier"))
	return rv
}

// A Boolean value that indicates whether the document is a draft that the user has not yet saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515065-draft?language=objc
func (d_ Document) IsDraft() bool {
	rv := objc.Call[bool](d_, objc.Sel("isDraft"))
	return rv
}

// A Boolean value that indicates whether the document is a draft that the user has not yet saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515065-draft?language=objc
func (d_ Document) SetDraft(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setDraft:"), value)
}

// A Boolean value that indicates whether the user chose to hide the document's filename extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515092-filenameextensionwashiddeninlast?language=objc
func (d_ Document) FileNameExtensionWasHiddenInLastRunSavePanel() bool {
	rv := objc.Call[bool](d_, objc.Sel("fileNameExtensionWasHiddenInLastRunSavePanel"))
	return rv
}

// The object that the document uses to support undo/redo operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515166-undomanager?language=objc
func (d_ Document) UndoManager() foundation.UndoManager {
	rv := objc.Call[foundation.UndoManager](d_, objc.Sel("undoManager"))
	return rv
}

// The object that the document uses to support undo/redo operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515166-undomanager?language=objc
func (d_ Document) SetUndoManager(value foundation.IUndoManager) {
	objc.Call[objc.Void](d_, objc.Sel("setUndoManager:"), objc.Ptr(value))
}

// The location of the most recently autosaved document contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515232-autosavedcontentsfileurl?language=objc
func (d_ Document) AutosavedContentsFileURL() foundation.URL {
	rv := objc.Call[foundation.URL](d_, objc.Sel("autosavedContentsFileURL"))
	return rv
}

// The location of the most recently autosaved document contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515232-autosavedcontentsfileurl?language=objc
func (d_ Document) SetAutosavedContentsFileURL(value foundation.IURL) {
	objc.Call[objc.Void](d_, objc.Sel("setAutosavedContentsFileURL:"), objc.Ptr(value))
}

// Returns whether the document object stores its contents in the user's iCloud document storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515085-usesubiquitousstorage?language=objc
func (dc _DocumentClass) UsesUbiquitousStorage() bool {
	rv := objc.Call[bool](dc, objc.Sel("usesUbiquitousStorage"))
	return rv
}

// Returns whether the document object stores its contents in the user's iCloud document storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515085-usesubiquitousstorage?language=objc
func Document_UsesUbiquitousStorage() bool {
	return DocumentClass.UsesUbiquitousStorage()
}

// A Boolean value that indicates whether the document’s Save panel displays a list of supported writable document types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515183-shouldrunsavepanelwithaccessoryv?language=objc
func (d_ Document) ShouldRunSavePanelWithAccessoryView() bool {
	rv := objc.Call[bool](d_, objc.Sel("shouldRunSavePanelWithAccessoryView"))
	return rv
}

// The name of the document seen by the user in AppleScript. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1500132-lastcomponentoffilename?language=objc
func (d_ Document) LastComponentOfFileName() string {
	rv := objc.Call[string](d_, objc.Sel("lastComponentOfFileName"))
	return rv
}

// The name of the document seen by the user in AppleScript. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1500132-lastcomponentoffilename?language=objc
func (d_ Document) SetLastComponentOfFileName(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setLastComponentOfFileName:"), value)
}

// A Boolean value that indicates whether the document archives previously saved versions of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515060-keepbackupfile?language=objc
func (d_ Document) KeepBackupFile() bool {
	rv := objc.Call[bool](d_, objc.Sel("keepBackupFile"))
	return rv
}

// A Boolean value that indicates whether the document subclass supports autosaving of drafts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515109-autosavesdrafts?language=objc
func (dc _DocumentClass) AutosavesDrafts() bool {
	rv := objc.Call[bool](dc, objc.Sel("autosavesDrafts"))
	return rv
}

// A Boolean value that indicates whether the document subclass supports autosaving of drafts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515109-autosavesdrafts?language=objc
func Document_AutosavesDrafts() bool {
	return DocumentClass.AutosavesDrafts()
}

// The URL for the document’s backup file that was created during an autosave operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515200-backupfileurl?language=objc
func (d_ Document) BackupFileURL() foundation.URL {
	rv := objc.Call[foundation.URL](d_, objc.Sel("backupFileURL"))
	return rv
}

// Returns an array of key paths that represent the restorable attributes of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1526232-restorablestatekeypaths?language=objc
func (dc _DocumentClass) RestorableStateKeyPaths() []string {
	rv := objc.Call[[]string](dc, objc.Sel("restorableStateKeyPaths"))
	return rv
}

// Returns an array of key paths that represent the restorable attributes of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1526232-restorablestatekeypaths?language=objc
func Document_RestorableStateKeyPaths() []string {
	return DocumentClass.RestorableStateKeyPaths()
}

// The location of the document’s on-disk representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515038-fileurl?language=objc
func (d_ Document) FileURL() foundation.URL {
	rv := objc.Call[foundation.URL](d_, objc.Sel("fileURL"))
	return rv
}

// The location of the document’s on-disk representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515038-fileurl?language=objc
func (d_ Document) SetFileURL(value foundation.IURL) {
	objc.Call[objc.Void](d_, objc.Sel("setFileURL:"), objc.Ptr(value))
}

// The file type that was last selected in the Save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515240-filetypefromlastrunsavepanel?language=objc
func (d_ Document) FileTypeFromLastRunSavePanel() string {
	rv := objc.Call[string](d_, objc.Sel("fileTypeFromLastRunSavePanel"))
	return rv
}

// The document’s current window controllers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515156-windowcontrollers?language=objc
func (d_ Document) WindowControllers() []WindowController {
	rv := objc.Call[[]WindowController](d_, objc.Sel("windowControllers"))
	return rv
}

// A Boolean value that indicates whether the document’s file is completely loaded into memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515053-entirefileloaded?language=objc
func (d_ Document) IsEntireFileLoaded() bool {
	rv := objc.Call[bool](d_, objc.Sel("isEntireFileLoaded"))
	return rv
}

// A Boolean value that indicates whether the document is currently displaying the Versions browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/2177310-browsingversions?language=objc
func (d_ Document) IsBrowsingVersions() bool {
	rv := objc.Call[bool](d_, objc.Sel("isBrowsingVersions"))
	return rv
}

// A Boolean value that indicates whether the document is in read-only mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515086-inviewingmode?language=objc
func (d_ Document) IsInViewingMode() bool {
	rv := objc.Call[bool](d_, objc.Sel("isInViewingMode"))
	return rv
}

// A print operation you can use to create a PDF representation of the document’s current contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515246-pdfprintoperation?language=objc
func (d_ Document) PDFPrintOperation() PrintOperation {
	rv := objc.Call[PrintOperation](d_, objc.Sel("PDFPrintOperation"))
	return rv
}

// Returns whether the document subclass supports version management. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515114-preservesversions?language=objc
func (dc _DocumentClass) PreservesVersions() bool {
	rv := objc.Call[bool](dc, objc.Sel("preservesVersions"))
	return rv
}

// Returns whether the document subclass supports version management. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515114-preservesversions?language=objc
func Document_PreservesVersions() bool {
	return DocumentClass.PreservesVersions()
}

// A Boolean value that indicates whether the document owns an undo manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515103-hasundomanager?language=objc
func (d_ Document) HasUndoManager() bool {
	rv := objc.Call[bool](d_, objc.Sel("hasUndoManager"))
	return rv
}

// A Boolean value that indicates whether the document owns an undo manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515103-hasundomanager?language=objc
func (d_ Document) SetHasUndoManager(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setHasUndoManager:"), value)
}

// Returns the types of data the receiver can read natively and any types filterable to that native type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515104-readabletypes?language=objc
func (dc _DocumentClass) ReadableTypes() []string {
	rv := objc.Call[[]string](dc, objc.Sel("readableTypes"))
	return rv
}

// Returns the types of data the receiver can read natively and any types filterable to that native type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515104-readabletypes?language=objc
func Document_ReadableTypes() []string {
	return DocumentClass.ReadableTypes()
}

// The name of the document type, as specified in the app’s information property-list file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515121-filetype?language=objc
func (d_ Document) FileType() string {
	rv := objc.Call[string](d_, objc.Sel("fileType"))
	return rv
}

// The name of the document type, as specified in the app’s information property-list file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515121-filetype?language=objc
func (d_ Document) SetFileType(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setFileType:"), value)
}

// A Boolean value that indicates whether the document has unsaved changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515091-documentedited?language=objc
func (d_ Document) IsDocumentEdited() bool {
	rv := objc.Call[bool](d_, objc.Sel("isDocumentEdited"))
	return rv
}

// A Boolean value that indicates whether or not the file can be written to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515212-locked?language=objc
func (d_ Document) IsLocked() bool {
	rv := objc.Call[bool](d_, objc.Sel("isLocked"))
	return rv
}

// The last-known modification date of the document’s on-disk representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515039-filemodificationdate?language=objc
func (d_ Document) FileModificationDate() foundation.Date {
	rv := objc.Call[foundation.Date](d_, objc.Sel("fileModificationDate"))
	return rv
}

// The last-known modification date of the document’s on-disk representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515039-filemodificationdate?language=objc
func (d_ Document) SetFileModificationDate(value foundation.IDate) {
	objc.Call[objc.Void](d_, objc.Sel("setFileModificationDate:"), objc.Ptr(value))
}

// The name of the document’s sole nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515174-windownibname?language=objc
func (d_ Document) WindowNibName() NibName {
	rv := objc.Call[NibName](d_, objc.Sel("windowNibName"))
	return rv
}

// Returns a Boolean value that indicates whether you can cancel an in-progress autosave operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515149-autosavingisimplicitlycancellabl?language=objc
func (d_ Document) AutosavingIsImplicitlyCancellable() bool {
	rv := objc.Call[bool](d_, objc.Sel("autosavingIsImplicitlyCancellable"))
	return rv
}

// A Boolean value that indicates whether the document is shareable from the standard Share menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/2902303-allowsdocumentsharing?language=objc
func (d_ Document) AllowsDocumentSharing() bool {
	rv := objc.Call[bool](d_, objc.Sel("allowsDocumentSharing"))
	return rv
}

// Returns whether the receiver supports autosaving in place. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515106-autosavesinplace?language=objc
func (dc _DocumentClass) AutosavesInPlace() bool {
	rv := objc.Call[bool](dc, objc.Sel("autosavesInPlace"))
	return rv
}

// Returns whether the receiver supports autosaving in place. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515106-autosavesinplace?language=objc
func Document_AutosavesInPlace() bool {
	return DocumentClass.AutosavesInPlace()
}

// Returns the types of data the receiver can write natively and any types filterable to that native type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515236-writabletypes?language=objc
func (dc _DocumentClass) WritableTypes() []string {
	rv := objc.Call[[]string](dc, objc.Sel("writableTypes"))
	return rv
}

// Returns the types of data the receiver can write natively and any types filterable to that native type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515236-writabletypes?language=objc
func Document_WritableTypes() []string {
	return DocumentClass.WritableTypes()
}

// A Boolean value that indicates whether the document has changes that have not been autosaved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515079-hasunautosavedchanges?language=objc
func (d_ Document) HasUnautosavedChanges() bool {
	rv := objc.Call[bool](d_, objc.Sel("hasUnautosavedChanges"))
	return rv
}

// Returns the document type to use for an autosave operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515136-autosavingfiletype?language=objc
func (d_ Document) AutosavingFileType() string {
	rv := objc.Call[string](d_, objc.Sel("autosavingFileType"))
	return rv
}

// Returns the document window to use as the parent of a document-modal sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515064-windowforsheet?language=objc
func (d_ Document) WindowForSheet() Window {
	rv := objc.Call[Window](d_, objc.Sel("windowForSheet"))
	return rv
}

// The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515077-displayname?language=objc
func (d_ Document) DisplayName() string {
	rv := objc.Call[string](d_, objc.Sel("displayName"))
	return rv
}

// An object that encapsulates a user activity supported by this document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1526106-useractivity?language=objc
func (d_ Document) UserActivity() foundation.UserActivity {
	rv := objc.Call[foundation.UserActivity](d_, objc.Sel("userActivity"))
	return rv
}

// An object that encapsulates a user activity supported by this document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1526106-useractivity?language=objc
func (d_ Document) SetUserActivity(value foundation.IUserActivity) {
	objc.Call[objc.Void](d_, objc.Sel("setUserActivity:"), objc.Ptr(value))
}
