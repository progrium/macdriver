// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"github.com/progrium/darwinkit/objc"
)

// An object you use to synchronize access to Metal resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlevent?language=objc
type PEvent interface {
	// optional
	Device() DeviceObject
	HasDevice() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool
}

// ensure impl type implements protocol interface
var _ PEvent = (*EventObject)(nil)

// A concrete type for the [PEvent] protocol.
type EventObject struct {
	objc.Object
}

func (e_ EventObject) HasDevice() bool {
	return e_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlevent/2966571-device?language=objc
func (e_ EventObject) Device() DeviceObject {
	rv := objc.Call[DeviceObject](e_, objc.Sel("device"))
	return rv
}

func (e_ EventObject) HasSetLabel() bool {
	return e_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlevent/2966572-label?language=objc
func (e_ EventObject) SetLabel(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setLabel:"), value)
}

func (e_ EventObject) HasLabel() bool {
	return e_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlevent/2966572-label?language=objc
func (e_ EventObject) Label() string {
	rv := objc.Call[string](e_, objc.Sel("label"))
	return rv
}
