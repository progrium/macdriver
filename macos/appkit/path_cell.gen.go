// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PathCell] class.
var PathCellClass = _PathCellClass{objc.GetClass("NSPathCell")}

type _PathCellClass struct {
	objc.Class
}

// An interface definition for the [PathCell] class.
type IPathCell interface {
	IActionCell
	MouseEnteredWithFrameInView(event IEvent, frame foundation.Rect, view IView)
	PathComponentCellAtPointWithFrameInView(point foundation.Point, frame foundation.Rect, view IView) PathComponentCell
	RectOfPathComponentCellWithFrameInView(cell IPathComponentCell, frame foundation.Rect, view IView) foundation.Rect
	MouseExitedWithFrameInView(event IEvent, frame foundation.Rect, view IView)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	Delegate() PathCellDelegateWrapper
	SetDelegate(value PPathCellDelegate)
	SetDelegateObject(valueObject objc.IObject)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	PathComponentCells() []PathComponentCell
	SetPathComponentCells(value []IPathComponentCell)
	PathStyle() PathStyle
	SetPathStyle(value PathStyle)
	AllowedTypes() []string
	SetAllowedTypes(value []string)
	URL() foundation.URL
	SetURL(value foundation.IURL)
	PlaceholderString() string
	SetPlaceholderString(value string)
	ClickedPathComponentCell() PathComponentCell
}

// The user interface of a path control object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell?language=objc
type PathCell struct {
	ActionCell
}

func PathCellFrom(ptr unsafe.Pointer) PathCell {
	return PathCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

func (pc _PathCellClass) Alloc() PathCell {
	rv := objc.Call[PathCell](pc, objc.Sel("alloc"))
	return rv
}

func PathCell_Alloc() PathCell {
	return PathCellClass.Alloc()
}

func (pc _PathCellClass) New() PathCell {
	rv := objc.Call[PathCell](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPathCell() PathCell {
	return PathCellClass.New()
}

func (p_ PathCell) Init() PathCell {
	rv := objc.Call[PathCell](p_, objc.Sel("init"))
	return rv
}

func (p_ PathCell) InitImageCell(image IImage) PathCell {
	rv := objc.Call[PathCell](p_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func PathCell_InitImageCell(image IImage) PathCell {
	return PathCellClass.Alloc().InitImageCell(image)
}

func (p_ PathCell) InitTextCell(string_ string) PathCell {
	rv := objc.Call[PathCell](p_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1530851-inittextcell?language=objc
func PathCell_InitTextCell(string_ string) PathCell {
	return PathCellClass.Alloc().InitTextCell(string_)
}

// Displays the cell component over which the mouse is hovering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1526407-mouseentered?language=objc
func (p_ PathCell) MouseEnteredWithFrameInView(event IEvent, frame foundation.Rect, view IView) {
	objc.Call[objc.Void](p_, objc.Sel("mouseEntered:withFrame:inView:"), objc.Ptr(event), frame, objc.Ptr(view))
}

// Returns the cell located at the given point within the given frame of the given view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1528185-pathcomponentcellatpoint?language=objc
func (p_ PathCell) PathComponentCellAtPointWithFrameInView(point foundation.Point, frame foundation.Rect, view IView) PathComponentCell {
	rv := objc.Call[PathComponentCell](p_, objc.Sel("pathComponentCellAtPoint:withFrame:inView:"), point, frame, objc.Ptr(view))
	return rv
}

// Returns the current rectangle being displayed for a given path component cell, with respect to a given frame in a given view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1534143-rectofpathcomponentcell?language=objc
func (p_ PathCell) RectOfPathComponentCellWithFrameInView(cell IPathComponentCell, frame foundation.Rect, view IView) foundation.Rect {
	rv := objc.Call[foundation.Rect](p_, objc.Sel("rectOfPathComponentCell:withFrame:inView:"), objc.Ptr(cell), frame, objc.Ptr(view))
	return rv
}

// Hides the cell component over which the mouse is hovering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1525878-mouseexited?language=objc
func (p_ PathCell) MouseExitedWithFrameInView(event IEvent, frame foundation.Rect, view IView) {
	objc.Call[objc.Void](p_, objc.Sel("mouseExited:withFrame:inView:"), objc.Ptr(event), frame, objc.Ptr(view))
}

// Sets the receiver’s double-click action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1532554-doubleaction?language=objc
func (p_ PathCell) DoubleAction() objc.Selector {
	rv := objc.Call[objc.Selector](p_, objc.Sel("doubleAction"))
	return rv
}

// Sets the receiver’s double-click action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1532554-doubleaction?language=objc
func (p_ PathCell) SetDoubleAction(value objc.Selector) {
	objc.Call[objc.Void](p_, objc.Sel("setDoubleAction:"), value)
}

// Sets the receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1532834-delegate?language=objc
func (p_ PathCell) Delegate() PathCellDelegateWrapper {
	rv := objc.Call[PathCellDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// Sets the receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1532834-delegate?language=objc
func (p_ PathCell) SetDelegate(value PPathCellDelegate) {
	po0 := objc.WrapAsProtocol("NSPathCellDelegate", value)
	objc.SetAssociatedObject(p_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), po0)
}

// Sets the receiver’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1532834-delegate?language=objc
func (p_ PathCell) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// Sets the value of the placeholder attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1524552-placeholderattributedstring?language=objc
func (p_ PathCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](p_, objc.Sel("placeholderAttributedString"))
	return rv
}

// Sets the value of the placeholder attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1524552-placeholderattributedstring?language=objc
func (p_ PathCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.Call[objc.Void](p_, objc.Sel("setPlaceholderAttributedString:"), objc.Ptr(value))
}

// Returns the current background color of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1527481-backgroundcolor?language=objc
func (p_ PathCell) BackgroundColor() Color {
	rv := objc.Call[Color](p_, objc.Sel("backgroundColor"))
	return rv
}

// Returns the current background color of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1527481-backgroundcolor?language=objc
func (p_ PathCell) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](p_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// Sets the array of NSPathComponentCell objects currently being displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1529433-pathcomponentcells?language=objc
func (p_ PathCell) PathComponentCells() []PathComponentCell {
	rv := objc.Call[[]PathComponentCell](p_, objc.Sel("pathComponentCells"))
	return rv
}

// Sets the array of NSPathComponentCell objects currently being displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1529433-pathcomponentcells?language=objc
func (p_ PathCell) SetPathComponentCells(value []IPathComponentCell) {
	objc.Call[objc.Void](p_, objc.Sel("setPathComponentCells:"), value)
}

// Sets the receiver’s path style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1524249-pathstyle?language=objc
func (p_ PathCell) PathStyle() PathStyle {
	rv := objc.Call[PathStyle](p_, objc.Sel("pathStyle"))
	return rv
}

// Sets the receiver’s path style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1524249-pathstyle?language=objc
func (p_ PathCell) SetPathStyle(value PathStyle) {
	objc.Call[objc.Void](p_, objc.Sel("setPathStyle:"), value)
}

// Sets the component types allowed in the path when the cell is editable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1524305-allowedtypes?language=objc
func (p_ PathCell) AllowedTypes() []string {
	rv := objc.Call[[]string](p_, objc.Sel("allowedTypes"))
	return rv
}

// Sets the component types allowed in the path when the cell is editable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1524305-allowedtypes?language=objc
func (p_ PathCell) SetAllowedTypes(value []string) {
	objc.Call[objc.Void](p_, objc.Sel("setAllowedTypes:"), value)
}

// Returns the path displayed by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1525034-url?language=objc
func (p_ PathCell) URL() foundation.URL {
	rv := objc.Call[foundation.URL](p_, objc.Sel("URL"))
	return rv
}

// Returns the path displayed by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1525034-url?language=objc
func (p_ PathCell) SetURL(value foundation.IURL) {
	objc.Call[objc.Void](p_, objc.Sel("setURL:"), objc.Ptr(value))
}

// Returns the placeholder string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1531136-placeholderstring?language=objc
func (p_ PathCell) PlaceholderString() string {
	rv := objc.Call[string](p_, objc.Sel("placeholderString"))
	return rv
}

// Returns the placeholder string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1531136-placeholderstring?language=objc
func (p_ PathCell) SetPlaceholderString(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setPlaceholderString:"), value)
}

// Sets the value of the path displayed by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1524894-clickedpathcomponentcell?language=objc
func (p_ PathCell) ClickedPathComponentCell() PathComponentCell {
	rv := objc.Call[PathComponentCell](p_, objc.Sel("clickedPathComponentCell"))
	return rv
}

// Returns the class used to create pathComponentCell objects when automatically filling up the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1535369-pathcomponentcellclass?language=objc
func (pc _PathCellClass) PathComponentCellClass() objc.Class {
	rv := objc.Call[objc.Class](pc, objc.Sel("pathComponentCellClass"))
	return rv
}

// Returns the class used to create pathComponentCell objects when automatically filling up the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcell/1535369-pathcomponentcellclass?language=objc
func PathCell_PathComponentCellClass() objc.Class {
	return PathCellClass.PathComponentCellClass()
}
