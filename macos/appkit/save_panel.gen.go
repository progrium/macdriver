// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/uti"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SavePanel] class.
var SavePanelClass = _SavePanelClass{objc.GetClass("NSSavePanel")}

type _SavePanelClass struct {
	objc.Class
}

// An interface definition for the [SavePanel] class.
type ISavePanel interface {
	IPanel
	Ok(sender objc.IObject) objc.Object
	RunModal() ModalResponse
	ValidateVisibleColumns()
	BeginWithCompletionHandler(handler func(result ModalResponse))
	BeginSheetModalForWindowCompletionHandler(window IWindow, handler func(result ModalResponse))
	Cancel(sender objc.IObject) objc.Object
	NameFieldLabel() string
	SetNameFieldLabel(value string)
	AllowsOtherFileTypes() bool
	SetAllowsOtherFileTypes(value bool)
	TagNames() []string
	SetTagNames(value []string)
	CanSelectHiddenExtension() bool
	SetCanSelectHiddenExtension(value bool)
	IsExpanded() bool
	Prompt() string
	SetPrompt(value string)
	DirectoryURL() foundation.URL
	SetDirectoryURL(value foundation.IURL)
	CanCreateDirectories() bool
	SetCanCreateDirectories(value bool)
	ShowsTagField() bool
	SetShowsTagField(value bool)
	AccessoryView() View
	SetAccessoryView(value IView)
	URL() foundation.URL
	ShowsHiddenFiles() bool
	SetShowsHiddenFiles(value bool)
	AllowedContentTypes() []uti.Type
	SetAllowedContentTypes(value []uti.IType)
	TreatsFilePackagesAsDirectories() bool
	SetTreatsFilePackagesAsDirectories(value bool)
	IsExtensionHidden() bool
	SetExtensionHidden(value bool)
	NameFieldStringValue() string
	SetNameFieldStringValue(value string)
	Message() string
	SetMessage(value string)
}

// A panel that prompts the user for information about where to save a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel?language=objc
type SavePanel struct {
	Panel
}

func SavePanelFrom(ptr unsafe.Pointer) SavePanel {
	return SavePanel{
		Panel: PanelFrom(ptr),
	}
}

func (sc _SavePanelClass) Alloc() SavePanel {
	rv := objc.Call[SavePanel](sc, objc.Sel("alloc"))
	return rv
}

func SavePanel_Alloc() SavePanel {
	return SavePanelClass.Alloc()
}

func (sc _SavePanelClass) New() SavePanel {
	rv := objc.Call[SavePanel](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSavePanel() SavePanel {
	return SavePanelClass.New()
}

func (s_ SavePanel) Init() SavePanel {
	rv := objc.Call[SavePanel](s_, objc.Sel("init"))
	return rv
}

func (sc _SavePanelClass) WindowWithContentViewController(contentViewController IViewController) SavePanel {
	rv := objc.Call[SavePanel](sc, objc.Sel("windowWithContentViewController:"), objc.Ptr(contentViewController))
	return rv
}

// Creates a titled window that contains the specified content view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc
func SavePanel_WindowWithContentViewController(contentViewController IViewController) SavePanel {
	return SavePanelClass.WindowWithContentViewController(contentViewController)
}

func (s_ SavePanel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) SavePanel {
	rv := objc.Call[SavePanel](s_, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.Ptr(screen))
	return rv
}

// Initializes an allocated window with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc
func SavePanel_InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) SavePanel {
	return SavePanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
}

// The action method that the panel calls when the user clicks the OK button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1535364-ok?language=objc
func (s_ SavePanel) Ok(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("ok:"), sender)
	return rv
}

// Displays the panel and begins its event loop with the current working (or last-selected) directory as the default starting point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1525357-runmodal?language=objc
func (s_ SavePanel) RunModal() ModalResponse {
	rv := objc.Call[ModalResponse](s_, objc.Sel("runModal"))
	return rv
}

// Creates a new Save panel and initializes it with default information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1539016-savepanel?language=objc
func (sc _SavePanelClass) SavePanel() SavePanel {
	rv := objc.Call[SavePanel](sc, objc.Sel("savePanel"))
	return rv
}

// Creates a new Save panel and initializes it with default information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1539016-savepanel?language=objc
func SavePanel_SavePanel() SavePanel {
	return SavePanelClass.SavePanel()
}

// Validates and reloads the browser columns visible in the panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1526381-validatevisiblecolumns?language=objc
func (s_ SavePanel) ValidateVisibleColumns() {
	objc.Call[objc.Void](s_, objc.Sel("validateVisibleColumns"))
}

// Presents the panel as a modeless window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1527007-beginwithcompletionhandler?language=objc
func (s_ SavePanel) BeginWithCompletionHandler(handler func(result ModalResponse)) {
	objc.Call[objc.Void](s_, objc.Sel("beginWithCompletionHandler:"), handler)
}

// Presents the panel as a sheet modal to the specified window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1535870-beginsheetmodalforwindow?language=objc
func (s_ SavePanel) BeginSheetModalForWindowCompletionHandler(window IWindow, handler func(result ModalResponse)) {
	objc.Call[objc.Void](s_, objc.Sel("beginSheetModalForWindow:completionHandler:"), objc.Ptr(window), handler)
}

// The action method that the panel calls when the user clicks the Cancel button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1534357-cancel?language=objc
func (s_ SavePanel) Cancel(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("cancel:"), sender)
	return rv
}

// The label text displayed in front of the filename text field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1535411-namefieldlabel?language=objc
func (s_ SavePanel) NameFieldLabel() string {
	rv := objc.Call[string](s_, objc.Sel("nameFieldLabel"))
	return rv
}

// The label text displayed in front of the filename text field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1535411-namefieldlabel?language=objc
func (s_ SavePanel) SetNameFieldLabel(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setNameFieldLabel:"), value)
}

// A Boolean value that indicates whether the panel allows the user to save files with a filename extension that’s not in the list of allowed types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1526960-allowsotherfiletypes?language=objc
func (s_ SavePanel) AllowsOtherFileTypes() bool {
	rv := objc.Call[bool](s_, objc.Sel("allowsOtherFileTypes"))
	return rv
}

// A Boolean value that indicates whether the panel allows the user to save files with a filename extension that’s not in the list of allowed types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1526960-allowsotherfiletypes?language=objc
func (s_ SavePanel) SetAllowsOtherFileTypes(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAllowsOtherFileTypes:"), value)
}

// The tag names that you want to include on a saved file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1535928-tagnames?language=objc
func (s_ SavePanel) TagNames() []string {
	rv := objc.Call[[]string](s_, objc.Sel("tagNames"))
	return rv
}

// The tag names that you want to include on a saved file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1535928-tagnames?language=objc
func (s_ SavePanel) SetTagNames(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setTagNames:"), value)
}

// A Boolean value that indicates whether the panel displays UI for hiding or showing filename extensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1535360-canselecthiddenextension?language=objc
func (s_ SavePanel) CanSelectHiddenExtension() bool {
	rv := objc.Call[bool](s_, objc.Sel("canSelectHiddenExtension"))
	return rv
}

// A Boolean value that indicates whether the panel displays UI for hiding or showing filename extensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1535360-canselecthiddenextension?language=objc
func (s_ SavePanel) SetCanSelectHiddenExtension(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setCanSelectHiddenExtension:"), value)
}

// A Boolean value that indicates whether whether the panel is expanded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1534515-expanded?language=objc
func (s_ SavePanel) IsExpanded() bool {
	rv := objc.Call[bool](s_, objc.Sel("isExpanded"))
	return rv
}

// The text to display in the default button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1525227-prompt?language=objc
func (s_ SavePanel) Prompt() string {
	rv := objc.Call[string](s_, objc.Sel("prompt"))
	return rv
}

// The text to display in the default button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1525227-prompt?language=objc
func (s_ SavePanel) SetPrompt(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setPrompt:"), value)
}

// The current directory shown in the panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1531279-directoryurl?language=objc
func (s_ SavePanel) DirectoryURL() foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("directoryURL"))
	return rv
}

// The current directory shown in the panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1531279-directoryurl?language=objc
func (s_ SavePanel) SetDirectoryURL(value foundation.IURL) {
	objc.Call[objc.Void](s_, objc.Sel("setDirectoryURL:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the panel displays UI for creating directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1532626-cancreatedirectories?language=objc
func (s_ SavePanel) CanCreateDirectories() bool {
	rv := objc.Call[bool](s_, objc.Sel("canCreateDirectories"))
	return rv
}

// A Boolean value that indicates whether the panel displays UI for creating directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1532626-cancreatedirectories?language=objc
func (s_ SavePanel) SetCanCreateDirectories(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setCanCreateDirectories:"), value)
}

// A Boolean value that indicates whether the panel displays the Tags field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1525589-showstagfield?language=objc
func (s_ SavePanel) ShowsTagField() bool {
	rv := objc.Call[bool](s_, objc.Sel("showsTagField"))
	return rv
}

// A Boolean value that indicates whether the panel displays the Tags field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1525589-showstagfield?language=objc
func (s_ SavePanel) SetShowsTagField(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setShowsTagField:"), value)
}

// The custom accessory view for the current app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1525544-accessoryview?language=objc
func (s_ SavePanel) AccessoryView() View {
	rv := objc.Call[View](s_, objc.Sel("accessoryView"))
	return rv
}

// The custom accessory view for the current app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1525544-accessoryview?language=objc
func (s_ SavePanel) SetAccessoryView(value IView) {
	objc.Call[objc.Void](s_, objc.Sel("setAccessoryView:"), objc.Ptr(value))
}

// A URL that contains the fully specified location of the targeted file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1534384-url?language=objc
func (s_ SavePanel) URL() foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("URL"))
	return rv
}

// A Boolean value that indicates whether the panel displays files that are normally hidden from the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1524285-showshiddenfiles?language=objc
func (s_ SavePanel) ShowsHiddenFiles() bool {
	rv := objc.Call[bool](s_, objc.Sel("showsHiddenFiles"))
	return rv
}

// A Boolean value that indicates whether the panel displays files that are normally hidden from the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1524285-showshiddenfiles?language=objc
func (s_ SavePanel) SetShowsHiddenFiles(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setShowsHiddenFiles:"), value)
}

// An array of types that specify the files types to which you can save. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/3566857-allowedcontenttypes?language=objc
func (s_ SavePanel) AllowedContentTypes() []uti.Type {
	rv := objc.Call[[]uti.Type](s_, objc.Sel("allowedContentTypes"))
	return rv
}

// An array of types that specify the files types to which you can save. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/3566857-allowedcontenttypes?language=objc
func (s_ SavePanel) SetAllowedContentTypes(value []uti.IType) {
	objc.Call[objc.Void](s_, objc.Sel("setAllowedContentTypes:"), value)
}

// A Boolean value that indicates whether the panel displays file packages as directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1529384-treatsfilepackagesasdirectories?language=objc
func (s_ SavePanel) TreatsFilePackagesAsDirectories() bool {
	rv := objc.Call[bool](s_, objc.Sel("treatsFilePackagesAsDirectories"))
	return rv
}

// A Boolean value that indicates whether the panel displays file packages as directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1529384-treatsfilepackagesasdirectories?language=objc
func (s_ SavePanel) SetTreatsFilePackagesAsDirectories(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setTreatsFilePackagesAsDirectories:"), value)
}

// A Boolean value that indicates whether to display filename extensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1529267-extensionhidden?language=objc
func (s_ SavePanel) IsExtensionHidden() bool {
	rv := objc.Call[bool](s_, objc.Sel("isExtensionHidden"))
	return rv
}

// A Boolean value that indicates whether to display filename extensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1529267-extensionhidden?language=objc
func (s_ SavePanel) SetExtensionHidden(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setExtensionHidden:"), value)
}

// The user-editable filename currently shown in the name field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1529299-namefieldstringvalue?language=objc
func (s_ SavePanel) NameFieldStringValue() string {
	rv := objc.Call[string](s_, objc.Sel("nameFieldStringValue"))
	return rv
}

// The user-editable filename currently shown in the name field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1529299-namefieldstringvalue?language=objc
func (s_ SavePanel) SetNameFieldStringValue(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setNameFieldStringValue:"), value)
}

// The message text displayed in the panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1528581-message?language=objc
func (s_ SavePanel) Message() string {
	rv := objc.Call[string](s_, objc.Sel("message"))
	return rv
}

// The message text displayed in the panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/1528581-message?language=objc
func (s_ SavePanel) SetMessage(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setMessage:"), value)
}
