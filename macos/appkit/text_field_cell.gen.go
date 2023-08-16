// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextFieldCell] class.
var TextFieldCellClass = _TextFieldCellClass{objc.GetClass("NSTextFieldCell")}

type _TextFieldCellClass struct {
	objc.Class
}

// An interface definition for the [TextFieldCell] class.
type ITextFieldCell interface {
	IActionCell
	SetWantsNotificationForMarkedText(flag bool)
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	TextColor() Color
	SetTextColor(value IColor)
	PlaceholderString() string
	SetPlaceholderString(value string)
	BezelStyle() TextFieldBezelStyle
	SetBezelStyle(value TextFieldBezelStyle)
}

// An object that enhances the text display capabilities of a cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell?language=objc
type TextFieldCell struct {
	ActionCell
}

func TextFieldCellFrom(ptr unsafe.Pointer) TextFieldCell {
	return TextFieldCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

func (t_ TextFieldCell) InitTextCell(string_ string) TextFieldCell {
	rv := objc.Call[TextFieldCell](t_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Initializes a text field cell that displays the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1642278-inittextcell?language=objc
func TextFieldCell_InitTextCell(string_ string) TextFieldCell {
	return TextFieldCellClass.Alloc().InitTextCell(string_)
}

func (tc _TextFieldCellClass) Alloc() TextFieldCell {
	rv := objc.Call[TextFieldCell](tc, objc.Sel("alloc"))
	return rv
}

func TextFieldCell_Alloc() TextFieldCell {
	return TextFieldCellClass.Alloc()
}

func (tc _TextFieldCellClass) New() TextFieldCell {
	rv := objc.Call[TextFieldCell](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextFieldCell() TextFieldCell {
	return TextFieldCellClass.New()
}

func (t_ TextFieldCell) Init() TextFieldCell {
	rv := objc.Call[TextFieldCell](t_, objc.Sel("init"))
	return rv
}

func (t_ TextFieldCell) InitImageCell(image IImage) TextFieldCell {
	rv := objc.Call[TextFieldCell](t_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func TextFieldCell_InitImageCell(image IImage) TextFieldCell {
	return TextFieldCellClass.Alloc().InitImageCell(image)
}

// Directs the cell’s associated field editor to post text change notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447147-setwantsnotificationformarkedtex?language=objc
func (t_ TextFieldCell) SetWantsNotificationForMarkedText(flag bool) {
	objc.Call[objc.Void](t_, objc.Sel("setWantsNotificationForMarkedText:"), flag)
}

// An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447163-allowedinputsourcelocales?language=objc
func (t_ TextFieldCell) AllowedInputSourceLocales() []string {
	rv := objc.Call[[]string](t_, objc.Sel("allowedInputSourceLocales"))
	return rv
}

// An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447163-allowedinputsourcelocales?language=objc
func (t_ TextFieldCell) SetAllowedInputSourceLocales(value []string) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowedInputSourceLocales:"), value)
}

// The placeholder text for the cell, specified as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447153-placeholderattributedstring?language=objc
func (t_ TextFieldCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](t_, objc.Sel("placeholderAttributedString"))
	return rv
}

// The placeholder text for the cell, specified as an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447153-placeholderattributedstring?language=objc
func (t_ TextFieldCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.Call[objc.Void](t_, objc.Sel("setPlaceholderAttributedString:"), objc.Ptr(value))
}

// The color of the cell’s background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447159-backgroundcolor?language=objc
func (t_ TextFieldCell) BackgroundColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("backgroundColor"))
	return rv
}

// The color of the cell’s background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447159-backgroundcolor?language=objc
func (t_ TextFieldCell) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the cell draws its background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447155-drawsbackground?language=objc
func (t_ TextFieldCell) DrawsBackground() bool {
	rv := objc.Call[bool](t_, objc.Sel("drawsBackground"))
	return rv
}

// A Boolean value that indicates whether the cell draws its background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447155-drawsbackground?language=objc
func (t_ TextFieldCell) SetDrawsBackground(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setDrawsBackground:"), value)
}

// The color to use to draw the cell’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447151-textcolor?language=objc
func (t_ TextFieldCell) TextColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("textColor"))
	return rv
}

// The color to use to draw the cell’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447151-textcolor?language=objc
func (t_ TextFieldCell) SetTextColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setTextColor:"), objc.Ptr(value))
}

// The placeholder text for the cell, specified as a plain text string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447165-placeholderstring?language=objc
func (t_ TextFieldCell) PlaceholderString() string {
	rv := objc.Call[string](t_, objc.Sel("placeholderString"))
	return rv
}

// The placeholder text for the cell, specified as a plain text string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447165-placeholderstring?language=objc
func (t_ TextFieldCell) SetPlaceholderString(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setPlaceholderString:"), value)
}

// The bezel style to use when drawing the text field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447161-bezelstyle?language=objc
func (t_ TextFieldCell) BezelStyle() TextFieldBezelStyle {
	rv := objc.Call[TextFieldBezelStyle](t_, objc.Sel("bezelStyle"))
	return rv
}

// The bezel style to use when drawing the text field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1447161-bezelstyle?language=objc
func (t_ TextFieldCell) SetBezelStyle(value TextFieldBezelStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setBezelStyle:"), value)
}
