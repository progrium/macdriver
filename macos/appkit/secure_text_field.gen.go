// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SecureTextField] class.
var SecureTextFieldClass = _SecureTextFieldClass{objc.GetClass("NSSecureTextField")}

type _SecureTextFieldClass struct {
	objc.Class
}

// An interface definition for the [SecureTextField] class.
type ISecureTextField interface {
	ITextField
}

// A text field that hides the typed text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssecuretextfield?language=objc
type SecureTextField struct {
	TextField
}

func SecureTextFieldFrom(ptr unsafe.Pointer) SecureTextField {
	return SecureTextField{
		TextField: TextFieldFrom(ptr),
	}
}

func (sc _SecureTextFieldClass) Alloc() SecureTextField {
	rv := objc.Call[SecureTextField](sc, objc.Sel("alloc"))
	return rv
}

func SecureTextField_Alloc() SecureTextField {
	return SecureTextFieldClass.Alloc()
}

func (sc _SecureTextFieldClass) New() SecureTextField {
	rv := objc.Call[SecureTextField](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSecureTextField() SecureTextField {
	return SecureTextFieldClass.New()
}

func (s_ SecureTextField) Init() SecureTextField {
	rv := objc.Call[SecureTextField](s_, objc.Sel("init"))
	return rv
}

func (sc _SecureTextFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) SecureTextField {
	rv := objc.Call[SecureTextField](sc, objc.Sel("labelWithAttributedString:"), objc.Ptr(attributedStringValue))
	return rv
}

// Creates a text field for use as a static label that displays styled text, doesn’t wrap, and doesn’t have selectable text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644658-labelwithattributedstring?language=objc
func SecureTextField_LabelWithAttributedString(attributedStringValue foundation.IAttributedString) SecureTextField {
	return SecureTextFieldClass.LabelWithAttributedString(attributedStringValue)
}

func (sc _SecureTextFieldClass) LabelWithString(stringValue string) SecureTextField {
	rv := objc.Call[SecureTextField](sc, objc.Sel("labelWithString:"), stringValue)
	return rv
}

// Initializes a text field for use as a static label that uses the system default font, doesn’t wrap, and doesn’t have selectable text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644377-labelwithstring?language=objc
func SecureTextField_LabelWithString(stringValue string) SecureTextField {
	return SecureTextFieldClass.LabelWithString(stringValue)
}

func (sc _SecureTextFieldClass) WrappingLabelWithString(stringValue string) SecureTextField {
	rv := objc.Call[SecureTextField](sc, objc.Sel("wrappingLabelWithString:"), stringValue)
	return rv
}

// Initializes a text field for use as a multiline static label with selectable text that uses the system default font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644543-wrappinglabelwithstring?language=objc
func SecureTextField_WrappingLabelWithString(stringValue string) SecureTextField {
	return SecureTextFieldClass.WrappingLabelWithString(stringValue)
}

func (sc _SecureTextFieldClass) TextFieldWithString(stringValue string) SecureTextField {
	rv := objc.Call[SecureTextField](sc, objc.Sel("textFieldWithString:"), stringValue)
	return rv
}

// Initializes a single-line editable text field for user input using the system default font and standard visual appearance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfield/1644706-textfieldwithstring?language=objc
func SecureTextField_TextFieldWithString(stringValue string) SecureTextField {
	return SecureTextFieldClass.TextFieldWithString(stringValue)
}

func (s_ SecureTextField) InitWithFrame(frameRect foundation.Rect) SecureTextField {
	rv := objc.Call[SecureTextField](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func NewSecureTextFieldWithFrame(frameRect foundation.Rect) SecureTextField {
	instance := SecureTextFieldClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}
