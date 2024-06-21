// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a switch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityswitch?language=objc
type PAccessibilitySwitch interface {
	// optional
	AccessibilityPerformDecrement() bool
	HasAccessibilityPerformDecrement() bool

	// optional
	AccessibilityPerformIncrement() bool
	HasAccessibilityPerformIncrement() bool

	// optional
	AccessibilityValue() string
	HasAccessibilityValue() bool
}

// ensure impl type implements protocol interface
var _ PAccessibilitySwitch = (*AccessibilitySwitchObject)(nil)

// A concrete type for the [PAccessibilitySwitch] protocol.
type AccessibilitySwitchObject struct {
	objc.Object
}

func (a_ AccessibilitySwitchObject) HasAccessibilityPerformDecrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformDecrement"))
}

// Decrements the switch’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityswitch/1528290-accessibilityperformdecrement?language=objc
func (a_ AccessibilitySwitchObject) AccessibilityPerformDecrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformDecrement"))
	return rv
}

func (a_ AccessibilitySwitchObject) HasAccessibilityPerformIncrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformIncrement"))
}

// Increments the switch’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityswitch/1533985-accessibilityperformincrement?language=objc
func (a_ AccessibilitySwitchObject) AccessibilityPerformIncrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

func (a_ AccessibilitySwitchObject) HasAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValue"))
}

// Returns the switch’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityswitch/1533946-accessibilityvalue?language=objc
func (a_ AccessibilitySwitchObject) AccessibilityValue() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityValue"))
	return rv
}
