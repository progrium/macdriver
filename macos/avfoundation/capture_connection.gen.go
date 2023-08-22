// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureConnection] class.
var CaptureConnectionClass = _CaptureConnectionClass{objc.GetClass("AVCaptureConnection")}

type _CaptureConnectionClass struct {
	objc.Class
}

// An interface definition for the [CaptureConnection] class.
type ICaptureConnection interface {
	objc.IObject
	IsVideoMirrored() bool
	SetVideoMirrored(value bool)
	AudioChannels() []CaptureAudioChannel
	VideoFieldMode() VideoFieldMode
	SetVideoFieldMode(value VideoFieldMode)
	IsVideoMirroringSupported() bool
	IsVideoMaxFrameDurationSupported() bool
	VideoPreviewLayer() CaptureVideoPreviewLayer
	IsActive() bool
	Output() CaptureOutput
	VideoMaxFrameDuration() coremedia.Time
	SetVideoMaxFrameDuration(value coremedia.Time)
	IsVideoFieldModeSupported() bool
	AutomaticallyAdjustsVideoMirroring() bool
	SetAutomaticallyAdjustsVideoMirroring(value bool)
	IsVideoMinFrameDurationSupported() bool
	IsEnabled() bool
	SetEnabled(value bool)
	InputPorts() []CaptureInputPort
	VideoMinFrameDuration() coremedia.Time
	SetVideoMinFrameDuration(value coremedia.Time)
}

// An object that represents a connection from a capture input to a capture output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection?language=objc
type CaptureConnection struct {
	objc.Object
}

func CaptureConnectionFrom(ptr unsafe.Pointer) CaptureConnection {
	return CaptureConnection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CaptureConnection) InitWithInputPortVideoPreviewLayer(port ICaptureInputPort, layer ICaptureVideoPreviewLayer) CaptureConnection {
	rv := objc.Call[CaptureConnection](c_, objc.Sel("initWithInputPort:videoPreviewLayer:"), objc.Ptr(port), objc.Ptr(layer))
	return rv
}

// Creates a capture connection that represents a connection between an input port and a video preview layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1385882-initwithinputport?language=objc
func NewCaptureConnectionWithInputPortVideoPreviewLayer(port ICaptureInputPort, layer ICaptureVideoPreviewLayer) CaptureConnection {
	instance := CaptureConnectionClass.Alloc().InitWithInputPortVideoPreviewLayer(port, layer)
	instance.Autorelease()
	return instance
}

func (cc _CaptureConnectionClass) ConnectionWithInputPortsOutput(ports []ICaptureInputPort, output ICaptureOutput) CaptureConnection {
	rv := objc.Call[CaptureConnection](cc, objc.Sel("connectionWithInputPorts:output:"), ports, objc.Ptr(output))
	return rv
}

// Returns a capture connection that represents a connection between multiple input ports and an output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1444473-connectionwithinputports?language=objc
func CaptureConnection_ConnectionWithInputPortsOutput(ports []ICaptureInputPort, output ICaptureOutput) CaptureConnection {
	return CaptureConnectionClass.ConnectionWithInputPortsOutput(ports, output)
}

func (c_ CaptureConnection) InitWithInputPortsOutput(ports []ICaptureInputPort, output ICaptureOutput) CaptureConnection {
	rv := objc.Call[CaptureConnection](c_, objc.Sel("initWithInputPorts:output:"), ports, objc.Ptr(output))
	return rv
}

// Creates a capture connection that represents a connection between multiple input ports and an output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1388896-initwithinputports?language=objc
func NewCaptureConnectionWithInputPortsOutput(ports []ICaptureInputPort, output ICaptureOutput) CaptureConnection {
	instance := CaptureConnectionClass.Alloc().InitWithInputPortsOutput(ports, output)
	instance.Autorelease()
	return instance
}

func (cc _CaptureConnectionClass) ConnectionWithInputPortVideoPreviewLayer(port ICaptureInputPort, layer ICaptureVideoPreviewLayer) CaptureConnection {
	rv := objc.Call[CaptureConnection](cc, objc.Sel("connectionWithInputPort:videoPreviewLayer:"), objc.Ptr(port), objc.Ptr(layer))
	return rv
}

// Returns a capture connection that represents a connection between an input port and a video preview layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1444495-connectionwithinputport?language=objc
func CaptureConnection_ConnectionWithInputPortVideoPreviewLayer(port ICaptureInputPort, layer ICaptureVideoPreviewLayer) CaptureConnection {
	return CaptureConnectionClass.ConnectionWithInputPortVideoPreviewLayer(port, layer)
}

func (cc _CaptureConnectionClass) Alloc() CaptureConnection {
	rv := objc.Call[CaptureConnection](cc, objc.Sel("alloc"))
	return rv
}

func CaptureConnection_Alloc() CaptureConnection {
	return CaptureConnectionClass.Alloc()
}

func (cc _CaptureConnectionClass) New() CaptureConnection {
	rv := objc.Call[CaptureConnection](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureConnection() CaptureConnection {
	return CaptureConnectionClass.New()
}

func (c_ CaptureConnection) Init() CaptureConnection {
	rv := objc.Call[CaptureConnection](c_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the connection horizontally flips the video flowing through it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1389172-videomirrored?language=objc
func (c_ CaptureConnection) IsVideoMirrored() bool {
	rv := objc.Call[bool](c_, objc.Sel("isVideoMirrored"))
	return rv
}

// A Boolean value that indicates whether the connection horizontally flips the video flowing through it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1389172-videomirrored?language=objc
func (c_ CaptureConnection) SetVideoMirrored(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setVideoMirrored:"), value)
}

// An array of audio channels that the connection provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1387519-audiochannels?language=objc
func (c_ CaptureConnection) AudioChannels() []CaptureAudioChannel {
	rv := objc.Call[[]CaptureAudioChannel](c_, objc.Sel("audioChannels"))
	return rv
}

// A setting that tells the connection how to interlace video flowing through it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1390559-videofieldmode?language=objc
func (c_ CaptureConnection) VideoFieldMode() VideoFieldMode {
	rv := objc.Call[VideoFieldMode](c_, objc.Sel("videoFieldMode"))
	return rv
}

// A setting that tells the connection how to interlace video flowing through it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1390559-videofieldmode?language=objc
func (c_ CaptureConnection) SetVideoFieldMode(value VideoFieldMode) {
	objc.Call[objc.Void](c_, objc.Sel("setVideoFieldMode:"), value)
}

// A Boolean value that indicates whether the connection supports video mirroring. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1387424-supportsvideomirroring?language=objc
func (c_ CaptureConnection) IsVideoMirroringSupported() bool {
	rv := objc.Call[bool](c_, objc.Sel("isVideoMirroringSupported"))
	return rv
}

// A Boolean value that indicates whether the connection supports a maximum frame duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1389158-supportsvideomaxframeduration?language=objc
func (c_ CaptureConnection) IsVideoMaxFrameDurationSupported() bool {
	rv := objc.Call[bool](c_, objc.Sel("isVideoMaxFrameDurationSupported"))
	return rv
}

// The video preview layer associated with the connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1390614-videopreviewlayer?language=objc
func (c_ CaptureConnection) VideoPreviewLayer() CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](c_, objc.Sel("videoPreviewLayer"))
	return rv
}

// Indicates whether the connection is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1387696-active?language=objc
func (c_ CaptureConnection) IsActive() bool {
	rv := objc.Call[bool](c_, objc.Sel("isActive"))
	return rv
}

// The connection’s output port, if applicable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1385571-output?language=objc
func (c_ CaptureConnection) Output() CaptureOutput {
	rv := objc.Call[CaptureOutput](c_, objc.Sel("output"))
	return rv
}

// The largest time interval the connection can apply between consecutive video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1390246-videomaxframeduration?language=objc
func (c_ CaptureConnection) VideoMaxFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("videoMaxFrameDuration"))
	return rv
}

// The largest time interval the connection can apply between consecutive video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1390246-videomaxframeduration?language=objc
func (c_ CaptureConnection) SetVideoMaxFrameDuration(value coremedia.Time) {
	objc.Call[objc.Void](c_, objc.Sel("setVideoMaxFrameDuration:"), value)
}

// A Boolean value that indicates whether the connection supports setting a video field mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1390836-supportsvideofieldmode?language=objc
func (c_ CaptureConnection) IsVideoFieldModeSupported() bool {
	rv := objc.Call[bool](c_, objc.Sel("isVideoFieldModeSupported"))
	return rv
}

// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1387082-automaticallyadjustsvideomirrori?language=objc
func (c_ CaptureConnection) AutomaticallyAdjustsVideoMirroring() bool {
	rv := objc.Call[bool](c_, objc.Sel("automaticallyAdjustsVideoMirroring"))
	return rv
}

// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1387082-automaticallyadjustsvideomirrori?language=objc
func (c_ CaptureConnection) SetAutomaticallyAdjustsVideoMirroring(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setAutomaticallyAdjustsVideoMirroring:"), value)
}

// A Boolean value that indicates whether the connection supports a minimum frame duration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1386978-supportsvideominframeduration?language=objc
func (c_ CaptureConnection) IsVideoMinFrameDurationSupported() bool {
	rv := objc.Call[bool](c_, objc.Sel("isVideoMinFrameDurationSupported"))
	return rv
}

// Turns the connection on and off. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1390131-enabled?language=objc
func (c_ CaptureConnection) IsEnabled() bool {
	rv := objc.Call[bool](c_, objc.Sel("isEnabled"))
	return rv
}

// Turns the connection on and off. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1390131-enabled?language=objc
func (c_ CaptureConnection) SetEnabled(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setEnabled:"), value)
}

// An array of the connection’s input ports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1388322-inputports?language=objc
func (c_ CaptureConnection) InputPorts() []CaptureInputPort {
	rv := objc.Call[[]CaptureInputPort](c_, objc.Sel("inputPorts"))
	return rv
}

// The smallest time interval the connection can apply between consecutive video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1388931-videominframeduration?language=objc
func (c_ CaptureConnection) VideoMinFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("videoMinFrameDuration"))
	return rv
}

// The smallest time interval the connection can apply between consecutive video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureconnection/1388931-videominframeduration?language=objc
func (c_ CaptureConnection) SetVideoMinFrameDuration(value coremedia.Time) {
	objc.Call[objc.Void](c_, objc.Sel("setVideoMinFrameDuration:"), value)
}
