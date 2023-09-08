// Code generated by DarwinKit. DO NOT EDIT.

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionDevice] class.
var ExtensionDeviceClass = _ExtensionDeviceClass{objc.GetClass("CMIOExtensionDevice")}

type _ExtensionDeviceClass struct {
	objc.Class
}

// An interface definition for the [ExtensionDevice] class.
type IExtensionDevice interface {
	objc.IObject
	AddStreamError(stream IExtensionStream, outError foundation.IError) bool
	RemoveStreamError(stream IExtensionStream, outError foundation.IError) bool
	NotifyPropertiesChanged(propertyStates map[ExtensionProperty]IExtensionPropertyState)
	LocalizedName() string
	Streams() []ExtensionStream
	LegacyDeviceID() string
	DeviceID() foundation.UUID
	Source() ExtensionDeviceSourceObject
}

// An object that represents a physical or virtual device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice?language=objc
type ExtensionDevice struct {
	objc.Object
}

func ExtensionDeviceFrom(ptr unsafe.Pointer) ExtensionDevice {
	return ExtensionDevice{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ExtensionDeviceClass) DeviceWithLocalizedNameDeviceIDLegacyDeviceIDSource(localizedName string, deviceID foundation.IUUID, legacyDeviceID string, source PExtensionDeviceSource) ExtensionDevice {
	po3 := objc.WrapAsProtocol("CMIOExtensionDeviceSource", source)
	rv := objc.Call[ExtensionDevice](ec, objc.Sel("deviceWithLocalizedName:deviceID:legacyDeviceID:source:"), localizedName, objc.Ptr(deviceID), legacyDeviceID, po3)
	return rv
}

// Returns a new extension device with an optional legacy device identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915828-devicewithlocalizedname?language=objc
func ExtensionDevice_DeviceWithLocalizedNameDeviceIDLegacyDeviceIDSource(localizedName string, deviceID foundation.IUUID, legacyDeviceID string, source PExtensionDeviceSource) ExtensionDevice {
	return ExtensionDeviceClass.DeviceWithLocalizedNameDeviceIDLegacyDeviceIDSource(localizedName, deviceID, legacyDeviceID, source)
}

func (e_ ExtensionDevice) InitWithLocalizedNameDeviceIDLegacyDeviceIDSource(localizedName string, deviceID foundation.IUUID, legacyDeviceID string, source PExtensionDeviceSource) ExtensionDevice {
	po3 := objc.WrapAsProtocol("CMIOExtensionDeviceSource", source)
	rv := objc.Call[ExtensionDevice](e_, objc.Sel("initWithLocalizedName:deviceID:legacyDeviceID:source:"), localizedName, objc.Ptr(deviceID), legacyDeviceID, po3)
	return rv
}

// Creates an extension device with an optional legacy device identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915830-initwithlocalizedname?language=objc
func NewExtensionDeviceWithLocalizedNameDeviceIDLegacyDeviceIDSource(localizedName string, deviceID foundation.IUUID, legacyDeviceID string, source PExtensionDeviceSource) ExtensionDevice {
	instance := ExtensionDeviceClass.Alloc().InitWithLocalizedNameDeviceIDLegacyDeviceIDSource(localizedName, deviceID, legacyDeviceID, source)
	instance.Autorelease()
	return instance
}

func (ec _ExtensionDeviceClass) Alloc() ExtensionDevice {
	rv := objc.Call[ExtensionDevice](ec, objc.Sel("alloc"))
	return rv
}

func (ec _ExtensionDeviceClass) New() ExtensionDevice {
	rv := objc.Call[ExtensionDevice](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionDevice() ExtensionDevice {
	return ExtensionDeviceClass.New()
}

func (e_ ExtensionDevice) Init() ExtensionDevice {
	rv := objc.Call[ExtensionDevice](e_, objc.Sel("init"))
	return rv
}

// Adds a stream to a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915826-addstream?language=objc
func (e_ ExtensionDevice) AddStreamError(stream IExtensionStream, outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("addStream:error:"), objc.Ptr(stream), objc.Ptr(outError))
	return rv
}

// Removes a stream from the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915835-removestream?language=objc
func (e_ ExtensionDevice) RemoveStreamError(stream IExtensionStream, outError foundation.IError) bool {
	rv := objc.Call[bool](e_, objc.Sel("removeStream:error:"), objc.Ptr(stream), objc.Ptr(outError))
	return rv
}

// Notifies clients of property changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915834-notifypropertieschanged?language=objc
func (e_ ExtensionDevice) NotifyPropertiesChanged(propertyStates map[ExtensionProperty]IExtensionPropertyState) {
	objc.Call[objc.Void](e_, objc.Sel("notifyPropertiesChanged:"), propertyStates)
}

// A localized name for a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915833-localizedname?language=objc
func (e_ ExtensionDevice) LocalizedName() string {
	rv := objc.Call[string](e_, objc.Sel("localizedName"))
	return rv
}

// An array of media streams attached to this device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915837-streams?language=objc
func (e_ ExtensionDevice) Streams() []ExtensionStream {
	rv := objc.Call[[]ExtensionStream](e_, objc.Sel("streams"))
	return rv
}

// A legacy device identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915832-legacydeviceid?language=objc
func (e_ ExtensionDevice) LegacyDeviceID() string {
	rv := objc.Call[string](e_, objc.Sel("legacyDeviceID"))
	return rv
}

// A universally unique device identifier value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915827-deviceid?language=objc
func (e_ ExtensionDevice) DeviceID() foundation.UUID {
	rv := objc.Call[foundation.UUID](e_, objc.Sel("deviceID"))
	return rv
}

// A source object for a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensiondevice/3915836-source?language=objc
func (e_ ExtensionDevice) Source() ExtensionDeviceSourceObject {
	rv := objc.Call[ExtensionDeviceSourceObject](e_, objc.Sel("source"))
	return rv
}
