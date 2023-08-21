// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PICTImageRep] class.
var PICTImageRepClass = _PICTImageRepClass{objc.GetClass("NSPICTImageRep")}

type _PICTImageRepClass struct {
	objc.Class
}

// An interface definition for the [PICTImageRep] class.
type IPICTImageRep interface {
	IImageRep
	PICTRepresentation() []byte
	BoundingBox() foundation.Rect
}

// An object that renders an image from a PICT format data stream of version 1, version 2, and extended version 2. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspictimagerep?language=objc
type PICTImageRep struct {
	ImageRep
}

func PICTImageRepFrom(ptr unsafe.Pointer) PICTImageRep {
	return PICTImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}

func (p_ PICTImageRep) InitWithData(pictData []byte) PICTImageRep {
	rv := objc.Call[PICTImageRep](p_, objc.Sel("initWithData:"), pictData)
	return rv
}

// Returns a representation of an image from the specified data in the PICT file format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspictimagerep/1533954-initwithdata?language=objc
func NewPICTImageRepWithData(pictData []byte) PICTImageRep {
	instance := PICTImageRepClass.Alloc().InitWithData(pictData)
	instance.Autorelease()
	return instance
}

func (pc _PICTImageRepClass) ImageRepWithData(pictData []byte) PICTImageRep {
	rv := objc.Call[PICTImageRep](pc, objc.Sel("imageRepWithData:"), pictData)
	return rv
}

// Creates and returns a representation of an image from the specified data in the PICT file format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspictimagerep/1588725-imagerepwithdata?language=objc
func PICTImageRep_ImageRepWithData(pictData []byte) PICTImageRep {
	return PICTImageRepClass.ImageRepWithData(pictData)
}

func (pc _PICTImageRepClass) Alloc() PICTImageRep {
	rv := objc.Call[PICTImageRep](pc, objc.Sel("alloc"))
	return rv
}

func PICTImageRep_Alloc() PICTImageRep {
	return PICTImageRepClass.Alloc()
}

func (pc _PICTImageRepClass) New() PICTImageRep {
	rv := objc.Call[PICTImageRep](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPICTImageRep() PICTImageRep {
	return PICTImageRepClass.New()
}

func (p_ PICTImageRep) Init() PICTImageRep {
	rv := objc.Call[PICTImageRep](p_, objc.Sel("init"))
	return rv
}

// The image representation’s PICT data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspictimagerep/1524591-pictrepresentation?language=objc
func (p_ PICTImageRep) PICTRepresentation() []byte {
	rv := objc.Call[[]byte](p_, objc.Sel("PICTRepresentation"))
	return rv
}

// The rectangle that bounds the image representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspictimagerep/1524978-boundingbox?language=objc
func (p_ PICTImageRep) BoundingBox() foundation.Rect {
	rv := objc.Call[foundation.Rect](p_, objc.Sel("boundingBox"))
	return rv
}
