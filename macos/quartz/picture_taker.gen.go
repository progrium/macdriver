// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PictureTaker] class.
var PictureTakerClass = _PictureTakerClass{objc.GetClass("IKPictureTaker")}

type _PictureTakerClass struct {
	objc.Class
}

// An interface definition for the [PictureTaker] class.
type IPictureTaker interface {
	appkit.IPanel
	PopUpRecentsMenuForViewWithDelegateDidEndSelectorContextInfo(aView appkit.IView, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	InputImage() appkit.Image
	Mirroring() bool
	SetInputImage(image appkit.IImage)
	RunModal() int
	BeginPictureTakerWithDelegateDidEndSelectorContextInfo(delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	BeginPictureTakerSheetForWindowWithDelegateDidEndSelectorContextInfo(aWindow appkit.IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	SetMirroring(b bool)
	OutputImage() appkit.Image
}

// The IKPictureTaker class represents a panel that allows users to choose images by browsing the file system. The picture taker panel provides an Open Recent menu, supports image cropping, and supports taking snapshots from an iSight or other digital camera. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker?language=objc
type PictureTaker struct {
	appkit.Panel
}

func PictureTakerFrom(ptr unsafe.Pointer) PictureTaker {
	return PictureTaker{
		Panel: appkit.PanelFrom(ptr),
	}
}

func (pc _PictureTakerClass) Alloc() PictureTaker {
	rv := objc.Call[PictureTaker](pc, objc.Sel("alloc"))
	return rv
}

func PictureTaker_Alloc() PictureTaker {
	return PictureTakerClass.Alloc()
}

func (pc _PictureTakerClass) New() PictureTaker {
	rv := objc.Call[PictureTaker](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPictureTaker() PictureTaker {
	return PictureTakerClass.New()
}

func (p_ PictureTaker) Init() PictureTaker {
	rv := objc.Call[PictureTaker](p_, objc.Sel("init"))
	return rv
}

func (pc _PictureTakerClass) WindowWithContentViewController(contentViewController appkit.IViewController) PictureTaker {
	rv := objc.Call[PictureTaker](pc, objc.Sel("windowWithContentViewController:"), objc.Ptr(contentViewController))
	return rv
}

// Creates a titled window that contains the specified content view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc
func PictureTaker_WindowWithContentViewController(contentViewController appkit.IViewController) PictureTaker {
	return PictureTakerClass.WindowWithContentViewController(contentViewController)
}

func (p_ PictureTaker) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style appkit.WindowStyleMask, backingStoreType appkit.BackingStoreType, flag bool, screen appkit.IScreen) PictureTaker {
	rv := objc.Call[PictureTaker](p_, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.Ptr(screen))
	return rv
}

// Initializes an allocated window with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc
func NewPictureTakerWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style appkit.WindowStyleMask, backingStoreType appkit.BackingStoreType, flag bool, screen appkit.IScreen) PictureTaker {
	instance := PictureTakerClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
	instance.Autorelease()
	return instance
}

// Displays the Open Recent popup menu associated with the  picture taker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker/1504753-popuprecentsmenuforview?language=objc
func (p_ PictureTaker) PopUpRecentsMenuForViewWithDelegateDidEndSelectorContextInfo(aView appkit.IView, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](p_, objc.Sel("popUpRecentsMenuForView:withDelegate:didEndSelector:contextInfo:"), objc.Ptr(aView), delegate, didEndSelector, contextInfo)
}

// Returns the  input  image associated with the picture taker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker/1505220-inputimage?language=objc
func (p_ PictureTaker) InputImage() appkit.Image {
	rv := objc.Call[appkit.Image](p_, objc.Sel("inputImage"))
	return rv
}

// Returns whether video mirroring is enabled during snapshots. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker/1504121-mirroring?language=objc
func (p_ PictureTaker) Mirroring() bool {
	rv := objc.Call[bool](p_, objc.Sel("mirroring"))
	return rv
}

// Set the image input for the picture taker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker/1503724-setinputimage?language=objc
func (p_ PictureTaker) SetInputImage(image appkit.IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), objc.Ptr(image))
}

// Opens a modal picture taker dialog. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker/1503911-runmodal?language=objc
func (p_ PictureTaker) RunModal() int {
	rv := objc.Call[int](p_, objc.Sel("runModal"))
	return rv
}

// Opens a picture taker pane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker/1503448-beginpicturetakerwithdelegate?language=objc
func (p_ PictureTaker) BeginPictureTakerWithDelegateDidEndSelectorContextInfo(delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](p_, objc.Sel("beginPictureTakerWithDelegate:didEndSelector:contextInfo:"), delegate, didEndSelector, contextInfo)
}

// Opens a picture taker as a sheet whose parent is the specified window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker/1504302-beginpicturetakersheetforwindow?language=objc
func (p_ PictureTaker) BeginPictureTakerSheetForWindowWithDelegateDidEndSelectorContextInfo(aWindow appkit.IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](p_, objc.Sel("beginPictureTakerSheetForWindow:withDelegate:didEndSelector:contextInfo:"), objc.Ptr(aWindow), delegate, didEndSelector, contextInfo)
}

// Controls whether the receiver enables video mirroring during snapshots. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker/1504915-setmirroring?language=objc
func (p_ PictureTaker) SetMirroring(b bool) {
	objc.Call[objc.Void](p_, objc.Sel("setMirroring:"), b)
}

// Returns a shared IKPictureTaker instance, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker/1503805-picturetaker?language=objc
func (pc _PictureTakerClass) PictureTaker() PictureTaker {
	rv := objc.Call[PictureTaker](pc, objc.Sel("pictureTaker"))
	return rv
}

// Returns a shared IKPictureTaker instance, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker/1503805-picturetaker?language=objc
func PictureTaker_PictureTaker() PictureTaker {
	return PictureTakerClass.PictureTaker()
}

// Returns the edited image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikpicturetaker/1504563-outputimage?language=objc
func (p_ PictureTaker) OutputImage() appkit.Image {
	rv := objc.Call[appkit.Image](p_, objc.Sel("outputImage"))
	return rv
}