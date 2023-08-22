// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataFaceObject] class.
var MetadataFaceObjectClass = _MetadataFaceObjectClass{objc.GetClass("AVMetadataFaceObject")}

type _MetadataFaceObjectClass struct {
	objc.Class
}

// An interface definition for the [MetadataFaceObject] class.
type IMetadataFaceObject interface {
	IMetadataObject
	HasRollAngle() bool
	YawAngle() float64
	RollAngle() float64
	HasYawAngle() bool
	FaceID() int
}

// Face information detected by a metadata capture output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject?language=objc
type MetadataFaceObject struct {
	MetadataObject
}

func MetadataFaceObjectFrom(ptr unsafe.Pointer) MetadataFaceObject {
	return MetadataFaceObject{
		MetadataObject: MetadataObjectFrom(ptr),
	}
}

func (mc _MetadataFaceObjectClass) Alloc() MetadataFaceObject {
	rv := objc.Call[MetadataFaceObject](mc, objc.Sel("alloc"))
	return rv
}

func MetadataFaceObject_Alloc() MetadataFaceObject {
	return MetadataFaceObjectClass.Alloc()
}

func (mc _MetadataFaceObjectClass) New() MetadataFaceObject {
	rv := objc.Call[MetadataFaceObject](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataFaceObject() MetadataFaceObject {
	return MetadataFaceObjectClass.New()
}

func (m_ MetadataFaceObject) Init() MetadataFaceObject {
	rv := objc.Call[MetadataFaceObject](m_, objc.Sel("init"))
	return rv
}

// A Boolean value indicating whether there is a valid roll angle associated with the face. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/1386866-hasrollangle?language=objc
func (m_ MetadataFaceObject) HasRollAngle() bool {
	rv := objc.Call[bool](m_, objc.Sel("hasRollAngle"))
	return rv
}

// The yaw angle of the face specified in degrees. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/1386517-yawangle?language=objc
func (m_ MetadataFaceObject) YawAngle() float64 {
	rv := objc.Call[float64](m_, objc.Sel("yawAngle"))
	return rv
}

// The roll angle of the face specified in degrees. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/1389110-rollangle?language=objc
func (m_ MetadataFaceObject) RollAngle() float64 {
	rv := objc.Call[float64](m_, objc.Sel("rollAngle"))
	return rv
}

// A Boolean value indicating whether there is a valid yaw angle associated with the face. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/1385888-hasyawangle?language=objc
func (m_ MetadataFaceObject) HasYawAngle() bool {
	rv := objc.Call[bool](m_, objc.Sel("hasYawAngle"))
	return rv
}

// The unique ID for this face metadata object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/1386945-faceid?language=objc
func (m_ MetadataFaceObject) FaceID() int {
	rv := objc.Call[int](m_, objc.Sel("faceID"))
	return rv
}
