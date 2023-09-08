// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureOutput] class.
var CaptureOutputClass = _CaptureOutputClass{objc.GetClass("AVCaptureOutput")}

type _CaptureOutputClass struct {
	objc.Class
}

// An interface definition for the [CaptureOutput] class.
type ICaptureOutput interface {
	objc.IObject
	TransformedMetadataObjectForMetadataObjectConnection(metadataObject IMetadataObject, connection ICaptureConnection) MetadataObject
	RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates coregraphics.Rect) coregraphics.Rect
	MetadataOutputRectOfInterestForRect(rectInOutputCoordinates coregraphics.Rect) coregraphics.Rect
	ConnectionWithMediaType(mediaType MediaType) CaptureConnection
	Connections() []CaptureConnection
}

// An abstract superclass for objects that provide media output destinations for a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput?language=objc
type CaptureOutput struct {
	objc.Object
}

func CaptureOutputFrom(ptr unsafe.Pointer) CaptureOutput {
	return CaptureOutput{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureOutputClass) Alloc() CaptureOutput {
	rv := objc.Call[CaptureOutput](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CaptureOutputClass) New() CaptureOutput {
	rv := objc.Call[CaptureOutput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureOutput() CaptureOutput {
	return CaptureOutputClass.New()
}

func (c_ CaptureOutput) Init() CaptureOutput {
	rv := objc.Call[CaptureOutput](c_, objc.Sel("init"))
	return rv
}

// Converts a metadata object’s visual properties to layer coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/1616310-transformedmetadataobjectformeta?language=objc
func (c_ CaptureOutput) TransformedMetadataObjectForMetadataObjectConnection(metadataObject IMetadataObject, connection ICaptureConnection) MetadataObject {
	rv := objc.Call[MetadataObject](c_, objc.Sel("transformedMetadataObjectForMetadataObject:connection:"), objc.Ptr(metadataObject), objc.Ptr(connection))
	return rv
}

// Converts a rectangle in the coordinate system used for metadata outputs to one in the capture output object’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/1616311-rectformetadataoutputrectofinter?language=objc
func (c_ CaptureOutput) RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates coregraphics.Rect) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](c_, objc.Sel("rectForMetadataOutputRectOfInterest:"), rectInMetadataOutputCoordinates)
	return rv
}

// Converts a rectangle in the capture output object’s coordinate system to one in the coordinate system used for metadata outputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/1616304-metadataoutputrectofinterestforr?language=objc
func (c_ CaptureOutput) MetadataOutputRectOfInterestForRect(rectInOutputCoordinates coregraphics.Rect) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](c_, objc.Sel("metadataOutputRectOfInterestForRect:"), rectInOutputCoordinates)
	return rv
}

// Returns the first connection with an input port of a specified media type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/1389574-connectionwithmediatype?language=objc
func (c_ CaptureOutput) ConnectionWithMediaType(mediaType MediaType) CaptureConnection {
	rv := objc.Call[CaptureConnection](c_, objc.Sel("connectionWithMediaType:"), mediaType)
	return rv
}

// The capture output object’s connections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureoutput/1386239-connections?language=objc
func (c_ CaptureOutput) Connections() []CaptureConnection {
	rv := objc.Call[[]CaptureConnection](c_, objc.Sel("connections"))
	return rv
}
