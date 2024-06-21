// Code generated by DarwinKit. DO NOT EDIT.

package quartz

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/macos/coreimage"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/quartzcore"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ImageView] class.
var ImageViewClass = _ImageViewClass{objc.GetClass("IKImageView")}

type _ImageViewClass struct {
	objc.Class
}

// An interface definition for the [ImageView] class.
type IImageView interface {
	appkit.IView
	ConvertViewPointToImagePoint(viewPoint foundation.Point) foundation.Point
	ZoomImageToFit(sender objc.IObject) objc.Object
	RotateImageRight(sender objc.IObject) objc.Object
	ZoomImageToActualSize(sender objc.IObject) objc.Object
	ImageSize() foundation.Size
	RotateImageLeft(sender objc.IObject) objc.Object
	ConvertViewRectToImageRect(viewRect foundation.Rect) foundation.Rect
	ConvertImagePointToViewPoint(imagePoint foundation.Point) foundation.Point
	SetImageImageProperties(image coregraphics.ImageRef, metaData foundation.Dictionary)
	ScrollToPoint(point foundation.Point)
	Crop(sender objc.IObject) objc.Object
	ZoomIn(sender objc.IObject) objc.Object
	OverlayForType(layerType string) quartzcore.Layer
	SetOverlayForType(layer quartzcore.ILayer, layerType string)
	ConvertImageRectToViewRect(imageRect foundation.Rect) foundation.Rect
	ImageProperties() foundation.Dictionary
	ScrollToRect(rect foundation.Rect)
	FlipImageVertical(sender objc.IObject) objc.Object
	ZoomOut(sender objc.IObject) objc.Object
	SetRotationAngleCenterPoint(rotationAngle float64, centerPoint foundation.Point)
	FlipImageHorizontal(sender objc.IObject) objc.Object
	SetImageZoomFactorCenterPoint(zoomFactor float64, centerPoint foundation.Point)
	Image() coregraphics.ImageRef
	ZoomImageToRect(rect foundation.Rect)
	SetImageWithURL(url foundation.IURL)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	Editable() bool
	SetEditable(value bool)
	CurrentToolMode() string
	SetCurrentToolMode(value string)
	Autoresizes() bool
	SetAutoresizes(value bool)
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
	BackgroundColor() appkit.Color
	SetBackgroundColor(value appkit.IColor)
	ImageCorrection() coreimage.Filter
	SetImageCorrection(value coreimage.IFilter)
	AutohidesScrollers() bool
	SetAutohidesScrollers(value bool)
	SupportsDragAndDrop() bool
	SetSupportsDragAndDrop(value bool)
	RotationAngle() float64
	SetRotationAngle(value float64)
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	ZoomFactor() float64
	SetZoomFactor(value float64)
	DoubleClickOpensImageEditPanel() bool
	SetDoubleClickOpensImageEditPanel(value bool)
}

// The IKImageView class provides an efficient way to display images in a view while at the same time supporting a number of image editing operations such as rotating, zooming, and cropping.  It supports drag and drop for the NSFilenamesPboardType flavor so that the user can drag an image to the view. If possible, image rendering uses hardware acceleration to achieve optimal performance. The IKImageView class is implemented as a subclass of NSView. Similar to NSImageView, the IKImageView class is used to display a single image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview?language=objc
type ImageView struct {
	appkit.View
}

func ImageViewFrom(ptr unsafe.Pointer) ImageView {
	return ImageView{
		View: appkit.ViewFrom(ptr),
	}
}

func (ic _ImageViewClass) Alloc() ImageView {
	rv := objc.Call[ImageView](ic, objc.Sel("alloc"))
	return rv
}

func (ic _ImageViewClass) New() ImageView {
	rv := objc.Call[ImageView](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageView() ImageView {
	return ImageViewClass.New()
}

func (i_ ImageView) Init() ImageView {
	rv := objc.Call[ImageView](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageView) InitWithFrame(frameRect foundation.Rect) ImageView {
	rv := objc.Call[ImageView](i_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewImageViewWithFrame(frameRect foundation.Rect) ImageView {
	instance := ImageViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Converts an image view coordinate to an image coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503743-convertviewpointtoimagepoint?language=objc
func (i_ ImageView) ConvertViewPointToImagePoint(viewPoint foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](i_, objc.Sel("convertViewPointToImagePoint:"), viewPoint)
	return rv
}

// Zooms the image so that it fits in the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504450-zoomimagetofit?language=objc
func (i_ ImageView) ZoomImageToFit(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("zoomImageToFit:"), sender)
	return rv
}

// Rotates the image right (clockwise). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503427-rotateimageright?language=objc
func (i_ ImageView) RotateImageRight(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("rotateImageRight:"), sender)
	return rv
}

// Zooms the image so that it is displayed using its true size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504415-zoomimagetoactualsize?language=objc
func (i_ ImageView) ZoomImageToActualSize(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("zoomImageToActualSize:"), sender)
	return rv
}

// Returns the size of the image in the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504005-imagesize?language=objc
func (i_ ImageView) ImageSize() foundation.Size {
	rv := objc.Call[foundation.Size](i_, objc.Sel("imageSize"))
	return rv
}

// Rotates the image left (counter-clockwise). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503769-rotateimageleft?language=objc
func (i_ ImageView) RotateImageLeft(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("rotateImageLeft:"), sender)
	return rv
}

// Converts an image view rectangle to an image rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504464-convertviewrecttoimagerect?language=objc
func (i_ ImageView) ConvertViewRectToImageRect(viewRect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](i_, objc.Sel("convertViewRectToImageRect:"), viewRect)
	return rv
}

// Converts an image coordinate to an image view coordinate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504633-convertimagepointtoviewpoint?language=objc
func (i_ ImageView) ConvertImagePointToViewPoint(imagePoint foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](i_, objc.Sel("convertImagePointToViewPoint:"), imagePoint)
	return rv
}

// Sets the image to display in an image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503915-setimage?language=objc
func (i_ ImageView) SetImageImageProperties(image coregraphics.ImageRef, metaData foundation.Dictionary) {
	objc.Call[objc.Void](i_, objc.Sel("setImage:imageProperties:"), image, metaData)
}

// Scrolls the view to the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503677-scrolltopoint?language=objc
func (i_ ImageView) ScrollToPoint(point foundation.Point) {
	objc.Call[objc.Void](i_, objc.Sel("scrollToPoint:"), point)
}

// Crops the image using the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503855-crop?language=objc
func (i_ ImageView) Crop(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("crop:"), sender)
	return rv
}

// Zooms the image in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503800-zoomin?language=objc
func (i_ ImageView) ZoomIn(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("zoomIn:"), sender)
	return rv
}

// Returns the Core Animation layer associated with a layer type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504236-overlayfortype?language=objc
func (i_ ImageView) OverlayForType(layerType string) quartzcore.Layer {
	rv := objc.Call[quartzcore.Layer](i_, objc.Sel("overlayForType:"), layerType)
	return rv
}

// Sets an overlay type for a Core Animation layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504283-setoverlay?language=objc
func (i_ ImageView) SetOverlayForType(layer quartzcore.ILayer, layerType string) {
	objc.Call[objc.Void](i_, objc.Sel("setOverlay:forType:"), layer, layerType)
}

// Converts an image rectangle to an image view rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504409-convertimagerecttoviewrect?language=objc
func (i_ ImageView) ConvertImageRectToViewRect(imageRect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](i_, objc.Sel("convertImageRectToViewRect:"), imageRect)
	return rv
}

// Returns the metadata for the image in the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503845-imageproperties?language=objc
func (i_ ImageView) ImageProperties() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](i_, objc.Sel("imageProperties"))
	return rv
}

// Scrolls the view so that it includes the provided rectangular area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504366-scrolltorect?language=objc
func (i_ ImageView) ScrollToRect(rect foundation.Rect) {
	objc.Call[objc.Void](i_, objc.Sel("scrollToRect:"), rect)
}

// Flips an image along the vertical axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503836-flipimagevertical?language=objc
func (i_ ImageView) FlipImageVertical(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("flipImageVertical:"), sender)
	return rv
}

// Zooms the image out. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503436-zoomout?language=objc
func (i_ ImageView) ZoomOut(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("zoomOut:"), sender)
	return rv
}

// Sets the rotation angle at the provided origin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503535-setrotationangle?language=objc
func (i_ ImageView) SetRotationAngleCenterPoint(rotationAngle float64, centerPoint foundation.Point) {
	objc.Call[objc.Void](i_, objc.Sel("setRotationAngle:centerPoint:"), rotationAngle, centerPoint)
}

// Flips an image along the horizontal axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1505282-flipimagehorizontal?language=objc
func (i_ ImageView) FlipImageHorizontal(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("flipImageHorizontal:"), sender)
	return rv
}

// Sets the zoom factor at the provided origin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503762-setimagezoomfactor?language=objc
func (i_ ImageView) SetImageZoomFactorCenterPoint(zoomFactor float64, centerPoint foundation.Point) {
	objc.Call[objc.Void](i_, objc.Sel("setImageZoomFactor:centerPoint:"), zoomFactor, centerPoint)
}

// Returns the image associated with the view, after any image corrections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504371-image?language=objc
func (i_ ImageView) Image() coregraphics.ImageRef {
	rv := objc.Call[coregraphics.ImageRef](i_, objc.Sel("image"))
	return rv
}

// Zooms the image so that it fits in the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504460-zoomimagetorect?language=objc
func (i_ ImageView) ZoomImageToRect(rect foundation.Rect) {
	objc.Call[objc.Void](i_, objc.Sel("zoomImageToRect:"), rect)
}

// Initializes an image view with the image specified by a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1505314-setimagewithurl?language=objc
func (i_ ImageView) SetImageWithURL(url foundation.IURL) {
	objc.Call[objc.Void](i_, objc.Sel("setImageWithURL:"), url)
}

// Specifies the vertical scroll bar state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1505256-hasverticalscroller?language=objc
func (i_ ImageView) HasVerticalScroller() bool {
	rv := objc.Call[bool](i_, objc.Sel("hasVerticalScroller"))
	return rv
}

// Specifies the vertical scroll bar state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1505256-hasverticalscroller?language=objc
func (i_ ImageView) SetHasVerticalScroller(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setHasVerticalScroller:"), value)
}

// Specifies the editable state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1505302-editable?language=objc
func (i_ ImageView) Editable() bool {
	rv := objc.Call[bool](i_, objc.Sel("editable"))
	return rv
}

// Specifies the editable state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1505302-editable?language=objc
func (i_ ImageView) SetEditable(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setEditable:"), value)
}

// Specifies the current tool mode for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503607-currenttoolmode?language=objc
func (i_ ImageView) CurrentToolMode() string {
	rv := objc.Call[string](i_, objc.Sel("currentToolMode"))
	return rv
}

// Specifies the current tool mode for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503607-currenttoolmode?language=objc
func (i_ ImageView) SetCurrentToolMode(value string) {
	objc.Call[objc.Void](i_, objc.Sel("setCurrentToolMode:"), value)
}

// Specifies the automatic resizing state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503709-autoresizes?language=objc
func (i_ ImageView) Autoresizes() bool {
	rv := objc.Call[bool](i_, objc.Sel("autoresizes"))
	return rv
}

// Specifies the automatic resizing state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503709-autoresizes?language=objc
func (i_ ImageView) SetAutoresizes(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setAutoresizes:"), value)
}

// Specifies the delegate object of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504032-delegate?language=objc
func (i_ ImageView) Delegate() objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("delegate"))
	return rv
}

// Specifies the delegate object of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504032-delegate?language=objc
func (i_ ImageView) SetDelegate(value objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("setDelegate:"), value)
}

// Specifies the background color for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503567-backgroundcolor?language=objc
func (i_ ImageView) BackgroundColor() appkit.Color {
	rv := objc.Call[appkit.Color](i_, objc.Sel("backgroundColor"))
	return rv
}

// Specifies the background color for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503567-backgroundcolor?language=objc
func (i_ ImageView) SetBackgroundColor(value appkit.IColor) {
	objc.Call[objc.Void](i_, objc.Sel("setBackgroundColor:"), value)
}

// Specifies a Core Image filter for image correction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503698-imagecorrection?language=objc
func (i_ ImageView) ImageCorrection() coreimage.Filter {
	rv := objc.Call[coreimage.Filter](i_, objc.Sel("imageCorrection"))
	return rv
}

// Specifies a Core Image filter for image correction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503698-imagecorrection?language=objc
func (i_ ImageView) SetImageCorrection(value coreimage.IFilter) {
	objc.Call[objc.Void](i_, objc.Sel("setImageCorrection:"), value)
}

// Specifies the automatic-hiding scroll bar state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503469-autohidesscrollers?language=objc
func (i_ ImageView) AutohidesScrollers() bool {
	rv := objc.Call[bool](i_, objc.Sel("autohidesScrollers"))
	return rv
}

// Specifies the automatic-hiding scroll bar state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503469-autohidesscrollers?language=objc
func (i_ ImageView) SetAutohidesScrollers(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setAutohidesScrollers:"), value)
}

// Specifies the drag-and-drop support state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504887-supportsdraganddrop?language=objc
func (i_ ImageView) SupportsDragAndDrop() bool {
	rv := objc.Call[bool](i_, objc.Sel("supportsDragAndDrop"))
	return rv
}

// Specifies the drag-and-drop support state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504887-supportsdraganddrop?language=objc
func (i_ ImageView) SetSupportsDragAndDrop(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setSupportsDragAndDrop:"), value)
}

// Specifies the rotation angle for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504691-rotationangle?language=objc
func (i_ ImageView) RotationAngle() float64 {
	rv := objc.Call[float64](i_, objc.Sel("rotationAngle"))
	return rv
}

// Specifies the rotation angle for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504691-rotationangle?language=objc
func (i_ ImageView) SetRotationAngle(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setRotationAngle:"), value)
}

// Specifies the horizontal scroll bar state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503565-hashorizontalscroller?language=objc
func (i_ ImageView) HasHorizontalScroller() bool {
	rv := objc.Call[bool](i_, objc.Sel("hasHorizontalScroller"))
	return rv
}

// Specifies the horizontal scroll bar state for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1503565-hashorizontalscroller?language=objc
func (i_ ImageView) SetHasHorizontalScroller(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setHasHorizontalScroller:"), value)
}

// Specifies the zoom factor for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504496-zoomfactor?language=objc
func (i_ ImageView) ZoomFactor() float64 {
	rv := objc.Call[float64](i_, objc.Sel("zoomFactor"))
	return rv
}

// Specifies the zoom factor for the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504496-zoomfactor?language=objc
func (i_ ImageView) SetZoomFactor(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setZoomFactor:"), value)
}

// Specifies the image-opening state of the editing pane in the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504092-doubleclickopensimageeditpanel?language=objc
func (i_ ImageView) DoubleClickOpensImageEditPanel() bool {
	rv := objc.Call[bool](i_, objc.Sel("doubleClickOpensImageEditPanel"))
	return rv
}

// Specifies the image-opening state of the editing pane in the image view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageview/1504092-doubleclickopensimageeditpanel?language=objc
func (i_ ImageView) SetDoubleClickOpensImageEditPanel(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setDoubleClickOpensImageEditPanel:"), value)
}
