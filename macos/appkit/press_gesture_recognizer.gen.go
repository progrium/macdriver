// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PressGestureRecognizer] class.
var PressGestureRecognizerClass = _PressGestureRecognizerClass{objc.GetClass("NSPressGestureRecognizer")}

type _PressGestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [PressGestureRecognizer] class.
type IPressGestureRecognizer interface {
	IGestureRecognizer
	MinimumPressDuration() foundation.TimeInterval
	SetMinimumPressDuration(value foundation.TimeInterval)
	ButtonMask() uint
	SetButtonMask(value uint)
	NumberOfTouchesRequired() int
	SetNumberOfTouchesRequired(value int)
	AllowableMovement() float64
	SetAllowableMovement(value float64)
}

// A discrete gesture recognizer that tracks whether the user holds down a mouse button for a minimum amount of time before releasing it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer?language=objc
type PressGestureRecognizer struct {
	GestureRecognizer
}

func PressGestureRecognizerFrom(ptr unsafe.Pointer) PressGestureRecognizer {
	return PressGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}

func (pc _PressGestureRecognizerClass) Alloc() PressGestureRecognizer {
	rv := objc.Call[PressGestureRecognizer](pc, objc.Sel("alloc"))
	return rv
}

func PressGestureRecognizer_Alloc() PressGestureRecognizer {
	return PressGestureRecognizerClass.Alloc()
}

func (pc _PressGestureRecognizerClass) New() PressGestureRecognizer {
	rv := objc.Call[PressGestureRecognizer](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPressGestureRecognizer() PressGestureRecognizer {
	return PressGestureRecognizerClass.New()
}

func (p_ PressGestureRecognizer) Init() PressGestureRecognizer {
	rv := objc.Call[PressGestureRecognizer](p_, objc.Sel("init"))
	return rv
}

func (p_ PressGestureRecognizer) InitWithTargetAction(target objc.IObject, action objc.Selector) PressGestureRecognizer {
	rv := objc.Call[PressGestureRecognizer](p_, objc.Sel("initWithTarget:action:"), target, action)
	return rv
}

// Initializes the gesture recognizer with the specified target and action information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/1535012-initwithtarget?language=objc
func PressGestureRecognizer_InitWithTargetAction(target objc.IObject, action objc.Selector) PressGestureRecognizer {
	return PressGestureRecognizerClass.Alloc().InitWithTargetAction(target, action)
}

// The minimum time (in seconds) that the user must hold the mouse button in the view for a valid gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/1531726-minimumpressduration?language=objc
func (p_ PressGestureRecognizer) MinimumPressDuration() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](p_, objc.Sel("minimumPressDuration"))
	return rv
}

// The minimum time (in seconds) that the user must hold the mouse button in the view for a valid gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/1531726-minimumpressduration?language=objc
func (p_ PressGestureRecognizer) SetMinimumPressDuration(value foundation.TimeInterval) {
	objc.Call[objc.Void](p_, objc.Sel("setMinimumPressDuration:"), value)
}

// A bit mask of the buttons required to recognize this press. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/1534468-buttonmask?language=objc
func (p_ PressGestureRecognizer) ButtonMask() uint {
	rv := objc.Call[uint](p_, objc.Sel("buttonMask"))
	return rv
}

// A bit mask of the buttons required to recognize this press. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/1534468-buttonmask?language=objc
func (p_ PressGestureRecognizer) SetButtonMask(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setButtonMask:"), value)
}

// The number of necessary touches on a Touch Bar for the gesture recognizer to match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/2544818-numberoftouchesrequired?language=objc
func (p_ PressGestureRecognizer) NumberOfTouchesRequired() int {
	rv := objc.Call[int](p_, objc.Sel("numberOfTouchesRequired"))
	return rv
}

// The number of necessary touches on a Touch Bar for the gesture recognizer to match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/2544818-numberoftouchesrequired?language=objc
func (p_ PressGestureRecognizer) SetNumberOfTouchesRequired(value int) {
	objc.Call[objc.Void](p_, objc.Sel("setNumberOfTouchesRequired:"), value)
}

// The maximum movement of the mouse in the view before the gesture fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/1527495-allowablemovement?language=objc
func (p_ PressGestureRecognizer) AllowableMovement() float64 {
	rv := objc.Call[float64](p_, objc.Sel("allowableMovement"))
	return rv
}

// The maximum movement of the mouse in the view before the gesture fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/1527495-allowablemovement?language=objc
func (p_ PressGestureRecognizer) SetAllowableMovement(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setAllowableMovement:"), value)
}
