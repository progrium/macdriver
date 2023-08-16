// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DictionaryController] class.
var DictionaryControllerClass = _DictionaryControllerClass{objc.GetClass("NSDictionaryController")}

type _DictionaryControllerClass struct {
	objc.Class
}

// An interface definition for the [DictionaryController] class.
type IDictionaryController interface {
	IArrayController
	ExcludedKeys() []string
	SetExcludedKeys(value []string)
	LocalizedKeyTable() string
	SetLocalizedKeyTable(value string)
	IncludedKeys() []string
	SetIncludedKeys(value []string)
	InitialValue() objc.Object
	SetInitialValue(value objc.IObject)
	InitialKey() string
	SetInitialKey(value string)
	LocalizedKeyDictionary() map[string]string
	SetLocalizedKeyDictionary(value map[string]string)
}

// A bindings-compatible controller that manages the display and editing of a dictionary of key-value pairs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller?language=objc
type DictionaryController struct {
	ArrayController
}

func DictionaryControllerFrom(ptr unsafe.Pointer) DictionaryController {
	return DictionaryController{
		ArrayController: ArrayControllerFrom(ptr),
	}
}

func (dc _DictionaryControllerClass) Alloc() DictionaryController {
	rv := objc.Call[DictionaryController](dc, objc.Sel("alloc"))
	return rv
}

func DictionaryController_Alloc() DictionaryController {
	return DictionaryControllerClass.Alloc()
}

func (dc _DictionaryControllerClass) New() DictionaryController {
	rv := objc.Call[DictionaryController](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDictionaryController() DictionaryController {
	return DictionaryControllerClass.New()
}

func (d_ DictionaryController) Init() DictionaryController {
	rv := objc.Call[DictionaryController](d_, objc.Sel("init"))
	return rv
}

func (d_ DictionaryController) InitWithContent(content objc.IObject) DictionaryController {
	rv := objc.Call[DictionaryController](d_, objc.Sel("initWithContent:"), content)
	return rv
}

// Initializes and returns an NSObjectController object with the given content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1529422-initwithcontent?language=objc
func DictionaryController_InitWithContent(content objc.IObject) DictionaryController {
	return DictionaryControllerClass.Alloc().InitWithContent(content)
}

// The key names that are never displayed in the user interface items bound to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1535356-excludedkeys?language=objc
func (d_ DictionaryController) ExcludedKeys() []string {
	rv := objc.Call[[]string](d_, objc.Sel("excludedKeys"))
	return rv
}

// The key names that are never displayed in the user interface items bound to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1535356-excludedkeys?language=objc
func (d_ DictionaryController) SetExcludedKeys(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setExcludedKeys:"), value)
}

// the strings file used to localize key names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1524332-localizedkeytable?language=objc
func (d_ DictionaryController) LocalizedKeyTable() string {
	rv := objc.Call[string](d_, objc.Sel("localizedKeyTable"))
	return rv
}

// the strings file used to localize key names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1524332-localizedkeytable?language=objc
func (d_ DictionaryController) SetLocalizedKeyTable(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setLocalizedKeyTable:"), value)
}

// The key names that are represented by a key-value pair, even if they are not present in the receiver’s content dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1525172-includedkeys?language=objc
func (d_ DictionaryController) IncludedKeys() []string {
	rv := objc.Call[[]string](d_, objc.Sel("includedKeys"))
	return rv
}

// The key names that are represented by a key-value pair, even if they are not present in the receiver’s content dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1525172-includedkeys?language=objc
func (d_ DictionaryController) SetIncludedKeys(value []string) {
	objc.Call[objc.Void](d_, objc.Sel("setIncludedKeys:"), value)
}

// The string used as the initial value for a newly inserted item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1535925-initialvalue?language=objc
func (d_ DictionaryController) InitialValue() objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("initialValue"))
	return rv
}

// The string used as the initial value for a newly inserted item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1535925-initialvalue?language=objc
func (d_ DictionaryController) SetInitialValue(value objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("setInitialValue:"), value)
}

// The string used as the initial key name for a newly inserted item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1530444-initialkey?language=objc
func (d_ DictionaryController) InitialKey() string {
	rv := objc.Call[string](d_, objc.Sel("initialKey"))
	return rv
}

// The string used as the initial key name for a newly inserted item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1530444-initialkey?language=objc
func (d_ DictionaryController) SetInitialKey(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setInitialKey:"), value)
}

// The localized key names that are displayed by the receiver in place of the key names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1532944-localizedkeydictionary?language=objc
func (d_ DictionaryController) LocalizedKeyDictionary() map[string]string {
	rv := objc.Call[map[string]string](d_, objc.Sel("localizedKeyDictionary"))
	return rv
}

// The localized key names that are displayed by the receiver in place of the key names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontroller/1532944-localizedkeydictionary?language=objc
func (d_ DictionaryController) SetLocalizedKeyDictionary(value map[string]string) {
	objc.Call[objc.Void](d_, objc.Sel("setLocalizedKeyDictionary:"), value)
}
