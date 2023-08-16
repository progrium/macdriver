// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScriptCoercionHandler] class.
var ScriptCoercionHandlerClass = _ScriptCoercionHandlerClass{objc.GetClass("NSScriptCoercionHandler")}

type _ScriptCoercionHandlerClass struct {
	objc.Class
}

// An interface definition for the [ScriptCoercionHandler] class.
type IScriptCoercionHandler interface {
	objc.IObject
	CoerceValueToClass(value objc.IObject, toClass objc.IClass) objc.Object
	RegisterCoercerSelectorToConvertFromClassToClass(coercer objc.IObject, selector objc.Selector, fromClass objc.IClass, toClass objc.IClass)
}

// A mechanism for converting one kind of scripting data to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcoercionhandler?language=objc
type ScriptCoercionHandler struct {
	objc.Object
}

func ScriptCoercionHandlerFrom(ptr unsafe.Pointer) ScriptCoercionHandler {
	return ScriptCoercionHandler{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _ScriptCoercionHandlerClass) Alloc() ScriptCoercionHandler {
	rv := objc.Call[ScriptCoercionHandler](sc, objc.Sel("alloc"))
	return rv
}

func ScriptCoercionHandler_Alloc() ScriptCoercionHandler {
	return ScriptCoercionHandlerClass.Alloc()
}

func (sc _ScriptCoercionHandlerClass) New() ScriptCoercionHandler {
	rv := objc.Call[ScriptCoercionHandler](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScriptCoercionHandler() ScriptCoercionHandler {
	return ScriptCoercionHandlerClass.New()
}

func (s_ ScriptCoercionHandler) Init() ScriptCoercionHandler {
	rv := objc.Call[ScriptCoercionHandler](s_, objc.Sel("init"))
	return rv
}

// Returns an object of a given class representing a given value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcoercionhandler/1412034-coercevalue?language=objc
func (s_ ScriptCoercionHandler) CoerceValueToClass(value objc.IObject, toClass objc.IClass) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("coerceValue:toClass:"), value, objc.Ptr(toClass))
	return rv
}

// Returns the shared NSScriptCoercionHandler for the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcoercionhandler/1411703-sharedcoercionhandler?language=objc
func (sc _ScriptCoercionHandlerClass) SharedCoercionHandler() ScriptCoercionHandler {
	rv := objc.Call[ScriptCoercionHandler](sc, objc.Sel("sharedCoercionHandler"))
	return rv
}

// Returns the shared NSScriptCoercionHandler for the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcoercionhandler/1411703-sharedcoercionhandler?language=objc
func ScriptCoercionHandler_SharedCoercionHandler() ScriptCoercionHandler {
	return ScriptCoercionHandlerClass.SharedCoercionHandler()
}

// Registers a given object (typically a class) to handle coercions (conversions) from one given class to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcoercionhandler/1413218-registercoercer?language=objc
func (s_ ScriptCoercionHandler) RegisterCoercerSelectorToConvertFromClassToClass(coercer objc.IObject, selector objc.Selector, fromClass objc.IClass, toClass objc.IClass) {
	objc.Call[objc.Void](s_, objc.Sel("registerCoercer:selector:toConvertFromClass:toClass:"), coercer, selector, objc.Ptr(fromClass), objc.Ptr(toClass))
}
