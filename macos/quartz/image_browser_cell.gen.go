// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageBrowserCell] class.
var ImageBrowserCellClass = _ImageBrowserCellClass{objc.GetClass("IKImageBrowserCell")}

type _ImageBrowserCellClass struct {
	objc.Class
}

// An interface definition for the [ImageBrowserCell] class.
type IImageBrowserCell interface {
	objc.IObject
	ImageAlignment() appkit.ImageAlignment
	Opacity() float64
	IndexOfRepresentedItem() uint
	IsSelected() bool
	ImageContainerFrame() foundation.Rect
	CellState() ImageBrowserCellState
	RepresentedItem() objc.Object
	ImageFrame() foundation.Rect
	TitleFrame() foundation.Rect
	Frame() foundation.Rect
	ImageBrowserView() ImageBrowserView
	SubtitleFrame() foundation.Rect
	LayerForType(type_ string) quartzcore.Layer
	SelectionFrame() foundation.Rect
}

// A class that is used to display a cell conforming to the [quartz/imagekit/ikimagebrowseritem_protocol] in an IKImageBrowserView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell?language=objc
type ImageBrowserCell struct {
	objc.Object
}

func ImageBrowserCellFrom(ptr unsafe.Pointer) ImageBrowserCell {
	return ImageBrowserCell{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _ImageBrowserCellClass) Alloc() ImageBrowserCell {
	rv := objc.Call[ImageBrowserCell](ic, objc.Sel("alloc"))
	return rv
}

func ImageBrowserCell_Alloc() ImageBrowserCell {
	return ImageBrowserCellClass.Alloc()
}

func (ic _ImageBrowserCellClass) New() ImageBrowserCell {
	rv := objc.Call[ImageBrowserCell](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageBrowserCell() ImageBrowserCell {
	return ImageBrowserCellClass.New()
}

func (i_ ImageBrowserCell) Init() ImageBrowserCell {
	rv := objc.Call[ImageBrowserCell](i_, objc.Sel("init"))
	return rv
}

// Returns the position of the cell’s image in the frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500042-imagealignment?language=objc
func (i_ ImageBrowserCell) ImageAlignment() appkit.ImageAlignment {
	rv := objc.Call[appkit.ImageAlignment](i_, objc.Sel("imageAlignment"))
	return rv
}

// Returns the opacity of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500060-opacity?language=objc
func (i_ ImageBrowserCell) Opacity() float64 {
	rv := objc.Call[float64](i_, objc.Sel("opacity"))
	return rv
}

// Returns the index of the receiver’s represented object in the datasource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500047-indexofrepresenteditem?language=objc
func (i_ ImageBrowserCell) IndexOfRepresentedItem() uint {
	rv := objc.Call[uint](i_, objc.Sel("indexOfRepresentedItem"))
	return rv
}

// Returns whether the cell is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500051-isselected?language=objc
func (i_ ImageBrowserCell) IsSelected() bool {
	rv := objc.Call[bool](i_, objc.Sel("isSelected"))
	return rv
}

// Returns the receiver’s image container frame rectangle, which defines the position of the container of the thumbnail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500064-imagecontainerframe?language=objc
func (i_ ImageBrowserCell) ImageContainerFrame() foundation.Rect {
	rv := objc.Call[foundation.Rect](i_, objc.Sel("imageContainerFrame"))
	return rv
}

// Returns the current cell state of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500062-cellstate?language=objc
func (i_ ImageBrowserCell) CellState() ImageBrowserCellState {
	rv := objc.Call[ImageBrowserCellState](i_, objc.Sel("cellState"))
	return rv
}

// Returns the receiver’s represented object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500071-representeditem?language=objc
func (i_ ImageBrowserCell) RepresentedItem() objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("representedItem"))
	return rv
}

// Returns the receiver’s image frame rectangle, which defines the position of the thumbnail in its IKImageBrowserView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500052-imageframe?language=objc
func (i_ ImageBrowserCell) ImageFrame() foundation.Rect {
	rv := objc.Call[foundation.Rect](i_, objc.Sel("imageFrame"))
	return rv
}

// Returns the receiver’s title frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500069-titleframe?language=objc
func (i_ ImageBrowserCell) TitleFrame() foundation.Rect {
	rv := objc.Call[foundation.Rect](i_, objc.Sel("titleFrame"))
	return rv
}

// Returns the receiver’s frame rectangle, which defines its position in its IKImageBrowserView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500044-frame?language=objc
func (i_ ImageBrowserCell) Frame() foundation.Rect {
	rv := objc.Call[foundation.Rect](i_, objc.Sel("frame"))
	return rv
}

// Returns the view the receiver uses to display the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500049-imagebrowserview?language=objc
func (i_ ImageBrowserCell) ImageBrowserView() ImageBrowserView {
	rv := objc.Call[ImageBrowserView](i_, objc.Sel("imageBrowserView"))
	return rv
}

// Returns the receiver’s subtitle frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500074-subtitleframe?language=objc
func (i_ ImageBrowserCell) SubtitleFrame() foundation.Rect {
	rv := objc.Call[foundation.Rect](i_, objc.Sel("subtitleFrame"))
	return rv
}

// Returns a layer for the specified position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500045-layerfortype?language=objc
func (i_ ImageBrowserCell) LayerForType(type_ string) quartzcore.Layer {
	rv := objc.Call[quartzcore.Layer](i_, objc.Sel("layerForType:"), type_)
	return rv
}

// Returns the receiver’s selection frame rectangle, which defines the position of the selection rectangle in its IKImageBrowserView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimagebrowsercell/1500070-selectionframe?language=objc
func (i_ ImageBrowserCell) SelectionFrame() foundation.Rect {
	rv := objc.Call[foundation.Rect](i_, objc.Sel("selectionFrame"))
	return rv
}
