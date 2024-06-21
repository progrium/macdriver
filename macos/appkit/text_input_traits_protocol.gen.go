// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits?language=objc
type PTextInputTraits interface {
	// optional
	SetTextCompletionType(value TextInputTraitType)
	HasSetTextCompletionType() bool

	// optional
	TextCompletionType() TextInputTraitType
	HasTextCompletionType() bool

	// optional
	SetSmartQuotesType(value TextInputTraitType)
	HasSetSmartQuotesType() bool

	// optional
	SmartQuotesType() TextInputTraitType
	HasSmartQuotesType() bool

	// optional
	SetGrammarCheckingType(value TextInputTraitType)
	HasSetGrammarCheckingType() bool

	// optional
	GrammarCheckingType() TextInputTraitType
	HasGrammarCheckingType() bool

	// optional
	SetSpellCheckingType(value TextInputTraitType)
	HasSetSpellCheckingType() bool

	// optional
	SpellCheckingType() TextInputTraitType
	HasSpellCheckingType() bool

	// optional
	SetAutocorrectionType(value TextInputTraitType)
	HasSetAutocorrectionType() bool

	// optional
	AutocorrectionType() TextInputTraitType
	HasAutocorrectionType() bool

	// optional
	SetSmartDashesType(value TextInputTraitType)
	HasSetSmartDashesType() bool

	// optional
	SmartDashesType() TextInputTraitType
	HasSmartDashesType() bool

	// optional
	SetSmartInsertDeleteType(value TextInputTraitType)
	HasSetSmartInsertDeleteType() bool

	// optional
	SmartInsertDeleteType() TextInputTraitType
	HasSmartInsertDeleteType() bool

	// optional
	SetLinkDetectionType(value TextInputTraitType)
	HasSetLinkDetectionType() bool

	// optional
	LinkDetectionType() TextInputTraitType
	HasLinkDetectionType() bool

	// optional
	SetTextReplacementType(value TextInputTraitType)
	HasSetTextReplacementType() bool

	// optional
	TextReplacementType() TextInputTraitType
	HasTextReplacementType() bool

	// optional
	SetDataDetectionType(value TextInputTraitType)
	HasSetDataDetectionType() bool

	// optional
	DataDetectionType() TextInputTraitType
	HasDataDetectionType() bool
}

// ensure impl type implements protocol interface
var _ PTextInputTraits = (*TextInputTraitsObject)(nil)

// A concrete type for the [PTextInputTraits] protocol.
type TextInputTraitsObject struct {
	objc.Object
}

func (t_ TextInputTraitsObject) HasSetTextCompletionType() bool {
	return t_.RespondsToSelector(objc.Sel("setTextCompletionType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242740-textcompletiontype?language=objc
func (t_ TextInputTraitsObject) SetTextCompletionType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setTextCompletionType:"), value)
}

func (t_ TextInputTraitsObject) HasTextCompletionType() bool {
	return t_.RespondsToSelector(objc.Sel("textCompletionType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242740-textcompletiontype?language=objc
func (t_ TextInputTraitsObject) TextCompletionType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("textCompletionType"))
	return rv
}

func (t_ TextInputTraitsObject) HasSetSmartQuotesType() bool {
	return t_.RespondsToSelector(objc.Sel("setSmartQuotesType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242738-smartquotestype?language=objc
func (t_ TextInputTraitsObject) SetSmartQuotesType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setSmartQuotesType:"), value)
}

func (t_ TextInputTraitsObject) HasSmartQuotesType() bool {
	return t_.RespondsToSelector(objc.Sel("smartQuotesType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242738-smartquotestype?language=objc
func (t_ TextInputTraitsObject) SmartQuotesType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("smartQuotesType"))
	return rv
}

func (t_ TextInputTraitsObject) HasSetGrammarCheckingType() bool {
	return t_.RespondsToSelector(objc.Sel("setGrammarCheckingType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242734-grammarcheckingtype?language=objc
func (t_ TextInputTraitsObject) SetGrammarCheckingType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setGrammarCheckingType:"), value)
}

func (t_ TextInputTraitsObject) HasGrammarCheckingType() bool {
	return t_.RespondsToSelector(objc.Sel("grammarCheckingType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242734-grammarcheckingtype?language=objc
func (t_ TextInputTraitsObject) GrammarCheckingType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("grammarCheckingType"))
	return rv
}

func (t_ TextInputTraitsObject) HasSetSpellCheckingType() bool {
	return t_.RespondsToSelector(objc.Sel("setSpellCheckingType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242739-spellcheckingtype?language=objc
func (t_ TextInputTraitsObject) SetSpellCheckingType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setSpellCheckingType:"), value)
}

func (t_ TextInputTraitsObject) HasSpellCheckingType() bool {
	return t_.RespondsToSelector(objc.Sel("spellCheckingType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242739-spellcheckingtype?language=objc
func (t_ TextInputTraitsObject) SpellCheckingType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("spellCheckingType"))
	return rv
}

func (t_ TextInputTraitsObject) HasSetAutocorrectionType() bool {
	return t_.RespondsToSelector(objc.Sel("setAutocorrectionType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242732-autocorrectiontype?language=objc
func (t_ TextInputTraitsObject) SetAutocorrectionType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setAutocorrectionType:"), value)
}

func (t_ TextInputTraitsObject) HasAutocorrectionType() bool {
	return t_.RespondsToSelector(objc.Sel("autocorrectionType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242732-autocorrectiontype?language=objc
func (t_ TextInputTraitsObject) AutocorrectionType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("autocorrectionType"))
	return rv
}

func (t_ TextInputTraitsObject) HasSetSmartDashesType() bool {
	return t_.RespondsToSelector(objc.Sel("setSmartDashesType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242736-smartdashestype?language=objc
func (t_ TextInputTraitsObject) SetSmartDashesType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setSmartDashesType:"), value)
}

func (t_ TextInputTraitsObject) HasSmartDashesType() bool {
	return t_.RespondsToSelector(objc.Sel("smartDashesType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242736-smartdashestype?language=objc
func (t_ TextInputTraitsObject) SmartDashesType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("smartDashesType"))
	return rv
}

func (t_ TextInputTraitsObject) HasSetSmartInsertDeleteType() bool {
	return t_.RespondsToSelector(objc.Sel("setSmartInsertDeleteType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242737-smartinsertdeletetype?language=objc
func (t_ TextInputTraitsObject) SetSmartInsertDeleteType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setSmartInsertDeleteType:"), value)
}

func (t_ TextInputTraitsObject) HasSmartInsertDeleteType() bool {
	return t_.RespondsToSelector(objc.Sel("smartInsertDeleteType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242737-smartinsertdeletetype?language=objc
func (t_ TextInputTraitsObject) SmartInsertDeleteType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("smartInsertDeleteType"))
	return rv
}

func (t_ TextInputTraitsObject) HasSetLinkDetectionType() bool {
	return t_.RespondsToSelector(objc.Sel("setLinkDetectionType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242735-linkdetectiontype?language=objc
func (t_ TextInputTraitsObject) SetLinkDetectionType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setLinkDetectionType:"), value)
}

func (t_ TextInputTraitsObject) HasLinkDetectionType() bool {
	return t_.RespondsToSelector(objc.Sel("linkDetectionType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242735-linkdetectiontype?language=objc
func (t_ TextInputTraitsObject) LinkDetectionType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("linkDetectionType"))
	return rv
}

func (t_ TextInputTraitsObject) HasSetTextReplacementType() bool {
	return t_.RespondsToSelector(objc.Sel("setTextReplacementType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242741-textreplacementtype?language=objc
func (t_ TextInputTraitsObject) SetTextReplacementType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setTextReplacementType:"), value)
}

func (t_ TextInputTraitsObject) HasTextReplacementType() bool {
	return t_.RespondsToSelector(objc.Sel("textReplacementType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242741-textreplacementtype?language=objc
func (t_ TextInputTraitsObject) TextReplacementType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("textReplacementType"))
	return rv
}

func (t_ TextInputTraitsObject) HasSetDataDetectionType() bool {
	return t_.RespondsToSelector(objc.Sel("setDataDetectionType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242733-datadetectiontype?language=objc
func (t_ TextInputTraitsObject) SetDataDetectionType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setDataDetectionType:"), value)
}

func (t_ TextInputTraitsObject) HasDataDetectionType() bool {
	return t_.RespondsToSelector(objc.Sel("dataDetectionType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242733-datadetectiontype?language=objc
func (t_ TextInputTraitsObject) DataDetectionType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("dataDetectionType"))
	return rv
}
