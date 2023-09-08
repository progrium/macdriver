// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Stepper] class.
var StepperClass = _StepperClass{objc.GetClass("NSStepper")}

type _StepperClass struct {
	objc.Class
}

// An interface definition for the [Stepper] class.
type IStepper interface {
	IControl
	ValueWraps() bool
	SetValueWraps(value bool)
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	Increment() float64
	SetIncrement(value float64)
	Autorepeat() bool
	SetAutorepeat(value bool)
}

// An interface with up and down arrow buttons for incrementing or decrementing a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper?language=objc
type Stepper struct {
	Control
}

func StepperFrom(ptr unsafe.Pointer) Stepper {
	return Stepper{
		Control: ControlFrom(ptr),
	}
}

func (sc _StepperClass) Alloc() Stepper {
	rv := objc.Call[Stepper](sc, objc.Sel("alloc"))
	return rv
}

func (sc _StepperClass) New() Stepper {
	rv := objc.Call[Stepper](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStepper() Stepper {
	return StepperClass.New()
}

func (s_ Stepper) Init() Stepper {
	rv := objc.Call[Stepper](s_, objc.Sel("init"))
	return rv
}

func (s_ Stepper) InitWithFrame(frameRect foundation.Rect) Stepper {
	rv := objc.Call[Stepper](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func NewStepperWithFrame(frameRect foundation.Rect) Stepper {
	instance := StepperClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// A Boolean value that indicates whether the stepper wraps around the minimum and maximum values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/1523580-valuewraps?language=objc
func (s_ Stepper) ValueWraps() bool {
	rv := objc.Call[bool](s_, objc.Sel("valueWraps"))
	return rv
}

// A Boolean value that indicates whether the stepper wraps around the minimum and maximum values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/1523580-valuewraps?language=objc
func (s_ Stepper) SetValueWraps(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setValueWraps:"), value)
}

// The stepper’s maximum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/1523578-maxvalue?language=objc
func (s_ Stepper) MaxValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maxValue"))
	return rv
}

// The stepper’s maximum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/1523578-maxvalue?language=objc
func (s_ Stepper) SetMaxValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMaxValue:"), value)
}

// The stepper’s minimum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/1523569-minvalue?language=objc
func (s_ Stepper) MinValue() float64 {
	rv := objc.Call[float64](s_, objc.Sel("minValue"))
	return rv
}

// The stepper’s minimum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/1523569-minvalue?language=objc
func (s_ Stepper) SetMinValue(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMinValue:"), value)
}

// The amount by which the receiver changes with each increment or decrement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/1523573-increment?language=objc
func (s_ Stepper) Increment() float64 {
	rv := objc.Call[float64](s_, objc.Sel("increment"))
	return rv
}

// The amount by which the receiver changes with each increment or decrement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/1523573-increment?language=objc
func (s_ Stepper) SetIncrement(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setIncrement:"), value)
}

// A Boolean value that indicates how the stepper responds to mouse events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/1523571-autorepeat?language=objc
func (s_ Stepper) Autorepeat() bool {
	rv := objc.Call[bool](s_, objc.Sel("autorepeat"))
	return rv
}

// A Boolean value that indicates how the stepper responds to mouse events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/1523571-autorepeat?language=objc
func (s_ Stepper) SetAutorepeat(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAutorepeat:"), value)
}
