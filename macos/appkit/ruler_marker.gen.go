// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RulerMarker] class.
var RulerMarkerClass = _RulerMarkerClass{objc.GetClass("NSRulerMarker")}

type _RulerMarkerClass struct {
	objc.Class
}

// An interface definition for the [RulerMarker] class.
type IRulerMarker interface {
	objc.IObject
	DrawRect(rect foundation.Rect)
	TrackMouseAdding(mouseDownEvent IEvent, isAdding bool) bool
	ImageRectInRuler() foundation.Rect
	ThicknessRequiredInRuler() float64
	MarkerLocation() float64
	SetMarkerLocation(value float64)
	ImageOrigin() foundation.Point
	SetImageOrigin(value foundation.Point)
	Ruler() RulerView
	IsRemovable() bool
	SetRemovable(value bool)
	IsDragging() bool
	RepresentedObject() foundation.CopyingWrapper
	SetRepresentedObject(value foundation.PCopying)
	SetRepresentedObjectObject(valueObject objc.IObject)
	Image() Image
	SetImage(value IImage)
	IsMovable() bool
	SetMovable(value bool)
}

// A symbol on a ruler view, indicating a location for the graphics element it represents in the client of the ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker?language=objc
type RulerMarker struct {
	objc.Object
}

func RulerMarkerFrom(ptr unsafe.Pointer) RulerMarker {
	return RulerMarker{
		Object: objc.ObjectFrom(ptr),
	}
}

func (r_ RulerMarker) InitWithRulerViewMarkerLocationImageImageOrigin(ruler IRulerView, location float64, image IImage, imageOrigin foundation.Point) RulerMarker {
	rv := objc.Call[RulerMarker](r_, objc.Sel("initWithRulerView:markerLocation:image:imageOrigin:"), objc.Ptr(ruler), location, objc.Ptr(image), imageOrigin)
	return rv
}

// Initializes a newly allocated ruler marker, associating it with (but not adding it to) a specified ruler view and assigning the attributes provided. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496240-initwithrulerview?language=objc
func RulerMarker_InitWithRulerViewMarkerLocationImageImageOrigin(ruler IRulerView, location float64, image IImage, imageOrigin foundation.Point) RulerMarker {
	return RulerMarkerClass.Alloc().InitWithRulerViewMarkerLocationImageImageOrigin(ruler, location, image, imageOrigin)
}

func (rc _RulerMarkerClass) Alloc() RulerMarker {
	rv := objc.Call[RulerMarker](rc, objc.Sel("alloc"))
	return rv
}

func RulerMarker_Alloc() RulerMarker {
	return RulerMarkerClass.Alloc()
}

func (rc _RulerMarkerClass) New() RulerMarker {
	rv := objc.Call[RulerMarker](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRulerMarker() RulerMarker {
	return RulerMarkerClass.New()
}

func (r_ RulerMarker) Init() RulerMarker {
	rv := objc.Call[RulerMarker](r_, objc.Sel("init"))
	return rv
}

// Draws the receiver’s image that appears in the supplied rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496234-drawrect?language=objc
func (r_ RulerMarker) DrawRect(rect foundation.Rect) {
	objc.Call[objc.Void](r_, objc.Sel("drawRect:"), rect)
}

// Handles user manipulation of the receiver in its ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496248-trackmouse?language=objc
func (r_ RulerMarker) TrackMouseAdding(mouseDownEvent IEvent, isAdding bool) bool {
	rv := objc.Call[bool](r_, objc.Sel("trackMouse:adding:"), objc.Ptr(mouseDownEvent), isAdding)
	return rv
}

// The rectangle occupied by the receiver’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496249-imagerectinruler?language=objc
func (r_ RulerMarker) ImageRectInRuler() foundation.Rect {
	rv := objc.Call[foundation.Rect](r_, objc.Sel("imageRectInRuler"))
	return rv
}

// The amount of the receiver’s image that’s displayed above or to the left of the ruler view's baseline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496245-thicknessrequiredinruler?language=objc
func (r_ RulerMarker) ThicknessRequiredInRuler() float64 {
	rv := objc.Call[float64](r_, objc.Sel("thicknessRequiredInRuler"))
	return rv
}

// The location of the receiver in the coordinate system of the ruler view's client view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496255-markerlocation?language=objc
func (r_ RulerMarker) MarkerLocation() float64 {
	rv := objc.Call[float64](r_, objc.Sel("markerLocation"))
	return rv
}

// The location of the receiver in the coordinate system of the ruler view's client view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496255-markerlocation?language=objc
func (r_ RulerMarker) SetMarkerLocation(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setMarkerLocation:"), value)
}

// The point in the receiver’s image that is positioned at the receiver’s location on the ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496236-imageorigin?language=objc
func (r_ RulerMarker) ImageOrigin() foundation.Point {
	rv := objc.Call[foundation.Point](r_, objc.Sel("imageOrigin"))
	return rv
}

// The point in the receiver’s image that is positioned at the receiver’s location on the ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496236-imageorigin?language=objc
func (r_ RulerMarker) SetImageOrigin(value foundation.Point) {
	objc.Call[objc.Void](r_, objc.Sel("setImageOrigin:"), value)
}

// The receiver's ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496239-ruler?language=objc
func (r_ RulerMarker) Ruler() RulerView {
	rv := objc.Call[RulerView](r_, objc.Sel("ruler"))
	return rv
}

// A Boolean that indicates whether the user can remove the receiver from its ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496238-removable?language=objc
func (r_ RulerMarker) IsRemovable() bool {
	rv := objc.Call[bool](r_, objc.Sel("isRemovable"))
	return rv
}

// A Boolean that indicates whether the user can remove the receiver from its ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496238-removable?language=objc
func (r_ RulerMarker) SetRemovable(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setRemovable:"), value)
}

// A Boolean that indicates whether the receiver is being dragged. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496253-dragging?language=objc
func (r_ RulerMarker) IsDragging() bool {
	rv := objc.Call[bool](r_, objc.Sel("isDragging"))
	return rv
}

// The object the receiver represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496244-representedobject?language=objc
func (r_ RulerMarker) RepresentedObject() foundation.CopyingWrapper {
	rv := objc.Call[foundation.CopyingWrapper](r_, objc.Sel("representedObject"))
	return rv
}

// The object the receiver represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496244-representedobject?language=objc
func (r_ RulerMarker) SetRepresentedObject(value foundation.PCopying) {
	po0 := objc.WrapAsProtocol("NSCopying", value)
	objc.Call[objc.Void](r_, objc.Sel("setRepresentedObject:"), po0)
}

// The object the receiver represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496244-representedobject?language=objc
func (r_ RulerMarker) SetRepresentedObjectObject(valueObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setRepresentedObject:"), objc.Ptr(valueObject))
}

// The receiver’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496251-image?language=objc
func (r_ RulerMarker) Image() Image {
	rv := objc.Call[Image](r_, objc.Sel("image"))
	return rv
}

// The receiver’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496251-image?language=objc
func (r_ RulerMarker) SetImage(value IImage) {
	objc.Call[objc.Void](r_, objc.Sel("setImage:"), objc.Ptr(value))
}

// A Boolean that indicates whether the user can move the receiver in its ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496247-movable?language=objc
func (r_ RulerMarker) IsMovable() bool {
	rv := objc.Call[bool](r_, objc.Sel("isMovable"))
	return rv
}

// A Boolean that indicates whether the user can move the receiver in its ruler view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/1496247-movable?language=objc
func (r_ RulerMarker) SetMovable(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setMovable:"), value)
}
