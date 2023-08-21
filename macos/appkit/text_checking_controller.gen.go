// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextCheckingController] class.
var TextCheckingControllerClass = _TextCheckingControllerClass{objc.GetClass("NSTextCheckingController")}

type _TextCheckingControllerClass struct {
	objc.Class
}

// An interface definition for the [TextCheckingController] class.
type ITextCheckingController interface {
	objc.IObject
	InsertedTextInRange(range_ foundation.Range)
	OrderFrontSubstitutionsPanel(sender objc.IObject)
	CheckTextInRangeTypesOptions(range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject)
	ShowGuessPanel(sender objc.IObject)
	CheckTextInDocument(sender objc.IObject)
	IgnoreSpelling(sender objc.IObject)
	CheckTextInSelection(sender objc.IObject)
	UpdateCandidates()
	ConsiderTextCheckingForRange(range_ foundation.Range)
	DidChangeTextInRange(range_ foundation.Range)
	DidChangeSelectedRange()
	MenuAtIndexClickedOnSelectionEffectiveRange(location uint, clickedOnSelection bool, effectiveRange foundation.RangePointer) Menu
	ValidAnnotations() []foundation.AttributedStringKey
	Invalidate()
	ChangeSpelling(sender objc.IObject)
	CheckSpelling(sender objc.IObject)
	SpellCheckerDocumentTag() int
	SetSpellCheckerDocumentTag(value int)
	Client() TextCheckingClientWrapper
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller?language=objc
type TextCheckingController struct {
	objc.Object
}

func TextCheckingControllerFrom(ptr unsafe.Pointer) TextCheckingController {
	return TextCheckingController{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextCheckingController) InitWithClient(client PTextCheckingClient) TextCheckingController {
	po0 := objc.WrapAsProtocol("NSTextCheckingClient", client)
	rv := objc.Call[TextCheckingController](t_, objc.Sel("initWithClient:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242753-initwithclient?language=objc
func NewTextCheckingControllerWithClient(client PTextCheckingClient) TextCheckingController {
	instance := TextCheckingControllerClass.Alloc().InitWithClient(client)
	instance.Autorelease()
	return instance
}

func (tc _TextCheckingControllerClass) Alloc() TextCheckingController {
	rv := objc.Call[TextCheckingController](tc, objc.Sel("alloc"))
	return rv
}

func TextCheckingController_Alloc() TextCheckingController {
	return TextCheckingControllerClass.Alloc()
}

func (tc _TextCheckingControllerClass) New() TextCheckingController {
	rv := objc.Call[TextCheckingController](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextCheckingController() TextCheckingController {
	return TextCheckingControllerClass.New()
}

func (t_ TextCheckingController) Init() TextCheckingController {
	rv := objc.Call[TextCheckingController](t_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242754-insertedtextinrange?language=objc
func (t_ TextCheckingController) InsertedTextInRange(range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("insertedTextInRange:"), range_)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242757-orderfrontsubstitutionspanel?language=objc
func (t_ TextCheckingController) OrderFrontSubstitutionsPanel(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("orderFrontSubstitutionsPanel:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242746-checktextinrange?language=objc
func (t_ TextCheckingController) CheckTextInRangeTypesOptions(range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("checkTextInRange:types:options:"), range_, checkingTypes, options)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242758-showguesspanel?language=objc
func (t_ TextCheckingController) ShowGuessPanel(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("showGuessPanel:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242745-checktextindocument?language=objc
func (t_ TextCheckingController) CheckTextInDocument(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("checkTextInDocument:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242752-ignorespelling?language=objc
func (t_ TextCheckingController) IgnoreSpelling(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("ignoreSpelling:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242747-checktextinselection?language=objc
func (t_ TextCheckingController) CheckTextInSelection(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("checkTextInSelection:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242760-updatecandidates?language=objc
func (t_ TextCheckingController) UpdateCandidates() {
	objc.Call[objc.Void](t_, objc.Sel("updateCandidates"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242749-considertextcheckingforrange?language=objc
func (t_ TextCheckingController) ConsiderTextCheckingForRange(range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("considerTextCheckingForRange:"), range_)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242751-didchangetextinrange?language=objc
func (t_ TextCheckingController) DidChangeTextInRange(range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("didChangeTextInRange:"), range_)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242750-didchangeselectedrange?language=objc
func (t_ TextCheckingController) DidChangeSelectedRange() {
	objc.Call[objc.Void](t_, objc.Sel("didChangeSelectedRange"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242756-menuatindex?language=objc
func (t_ TextCheckingController) MenuAtIndexClickedOnSelectionEffectiveRange(location uint, clickedOnSelection bool, effectiveRange foundation.RangePointer) Menu {
	rv := objc.Call[Menu](t_, objc.Sel("menuAtIndex:clickedOnSelection:effectiveRange:"), location, clickedOnSelection, effectiveRange)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242761-validannotations?language=objc
func (t_ TextCheckingController) ValidAnnotations() []foundation.AttributedStringKey {
	rv := objc.Call[[]foundation.AttributedStringKey](t_, objc.Sel("validAnnotations"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242755-invalidate?language=objc
func (t_ TextCheckingController) Invalidate() {
	objc.Call[objc.Void](t_, objc.Sel("invalidate"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242743-changespelling?language=objc
func (t_ TextCheckingController) ChangeSpelling(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("changeSpelling:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242744-checkspelling?language=objc
func (t_ TextCheckingController) CheckSpelling(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("checkSpelling:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242759-spellcheckerdocumenttag?language=objc
func (t_ TextCheckingController) SpellCheckerDocumentTag() int {
	rv := objc.Call[int](t_, objc.Sel("spellCheckerDocumentTag"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242759-spellcheckerdocumenttag?language=objc
func (t_ TextCheckingController) SetSpellCheckerDocumentTag(value int) {
	objc.Call[objc.Void](t_, objc.Sel("setSpellCheckerDocumentTag:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/3242748-client?language=objc
func (t_ TextCheckingController) Client() TextCheckingClientWrapper {
	rv := objc.Call[TextCheckingClientWrapper](t_, objc.Sel("client"))
	return rv
}
