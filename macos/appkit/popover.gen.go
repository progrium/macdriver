// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Popover] class.
var PopoverClass = _PopoverClass{objc.GetClass("NSPopover")}

type _PopoverClass struct {
	objc.Class
}

// An interface definition for the [Popover] class.
type IPopover interface {
	IResponder
	Close()
	ShowRelativeToRectOfViewPreferredEdge(positioningRect foundation.Rect, positioningView IView, preferredEdge foundation.RectEdge)
	PerformClose(sender objc.IObject) objc.Object
	Behavior() PopoverBehavior
	SetBehavior(value PopoverBehavior)
	IsShown() bool
	EffectiveAppearance() Appearance
	IsDetached() bool
	Delegate() PopoverDelegateWrapper
	SetDelegate(value PPopoverDelegate)
	SetDelegateObject(valueObject objc.IObject)
	ContentSize() foundation.Size
	SetContentSize(value foundation.Size)
	ContentViewController() ViewController
	SetContentViewController(value IViewController)
	PositioningRect() foundation.Rect
	SetPositioningRect(value foundation.Rect)
	Animates() bool
	SetAnimates(value bool)
	Appearance() Appearance
	SetAppearance(value IAppearance)
}

// A means to display additional content related to existing content on the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover?language=objc
type Popover struct {
	Responder
}

func PopoverFrom(ptr unsafe.Pointer) Popover {
	return Popover{
		Responder: ResponderFrom(ptr),
	}
}

func (p_ Popover) Init() Popover {
	rv := objc.Call[Popover](p_, objc.Sel("init"))
	return rv
}

func (pc _PopoverClass) Alloc() Popover {
	rv := objc.Call[Popover](pc, objc.Sel("alloc"))
	return rv
}

func Popover_Alloc() Popover {
	return PopoverClass.Alloc()
}

func (pc _PopoverClass) New() Popover {
	rv := objc.Call[Popover](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPopover() Popover {
	return PopoverClass.New()
}

// Forces the popover to close without consulting its delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1526823-close?language=objc
func (p_ Popover) Close() {
	objc.Call[objc.Void](p_, objc.Sel("close"))
}

// Shows the popover anchored to the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1532113-showrelativetorect?language=objc
func (p_ Popover) ShowRelativeToRectOfViewPreferredEdge(positioningRect foundation.Rect, positioningView IView, preferredEdge foundation.RectEdge) {
	objc.Call[objc.Void](p_, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), positioningRect, objc.Ptr(positioningView), preferredEdge)
}

// Attempts to close the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1534290-performclose?language=objc
func (p_ Popover) PerformClose(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("performClose:"), sender)
	return rv
}

// Specifies the behavior of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1533539-behavior?language=objc
func (p_ Popover) Behavior() PopoverBehavior {
	rv := objc.Call[PopoverBehavior](p_, objc.Sel("behavior"))
	return rv
}

// Specifies the behavior of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1533539-behavior?language=objc
func (p_ Popover) SetBehavior(value PopoverBehavior) {
	objc.Call[objc.Void](p_, objc.Sel("setBehavior:"), value)
}

// The display state of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1535120-shown?language=objc
func (p_ Popover) IsShown() bool {
	rv := objc.Call[bool](p_, objc.Sel("isShown"))
	return rv
}

// The appearance that will be used when the popover is displayed onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1526863-effectiveappearance?language=objc
func (p_ Popover) EffectiveAppearance() Appearance {
	rv := objc.Call[Appearance](p_, objc.Sel("effectiveAppearance"))
	return rv
}

// A Boolean value that indicates whether the window created by a popover's detachment is automatically created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1534278-detached?language=objc
func (p_ Popover) IsDetached() bool {
	rv := objc.Call[bool](p_, objc.Sel("isDetached"))
	return rv
}

// The delegate of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1526708-delegate?language=objc
func (p_ Popover) Delegate() PopoverDelegateWrapper {
	rv := objc.Call[PopoverDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// The delegate of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1526708-delegate?language=objc
func (p_ Popover) SetDelegate(value PPopoverDelegate) {
	po0 := objc.WrapAsProtocol("NSPopoverDelegate", value)
	objc.SetAssociatedObject(p_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), po0)
}

// The delegate of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1526708-delegate?language=objc
func (p_ Popover) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The content size of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1524677-contentsize?language=objc
func (p_ Popover) ContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](p_, objc.Sel("contentSize"))
	return rv
}

// The content size of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1524677-contentsize?language=objc
func (p_ Popover) SetContentSize(value foundation.Size) {
	objc.Call[objc.Void](p_, objc.Sel("setContentSize:"), value)
}

// The view controller that manages the content of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1526794-contentviewcontroller?language=objc
func (p_ Popover) ContentViewController() ViewController {
	rv := objc.Call[ViewController](p_, objc.Sel("contentViewController"))
	return rv
}

// The view controller that manages the content of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1526794-contentviewcontroller?language=objc
func (p_ Popover) SetContentViewController(value IViewController) {
	objc.Call[objc.Void](p_, objc.Sel("setContentViewController:"), objc.Ptr(value))
}

// The rectangle within the positioning view relative to which the popover should be positioned. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1526090-positioningrect?language=objc
func (p_ Popover) PositioningRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](p_, objc.Sel("positioningRect"))
	return rv
}

// The rectangle within the positioning view relative to which the popover should be positioned. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1526090-positioningrect?language=objc
func (p_ Popover) SetPositioningRect(value foundation.Rect) {
	objc.Call[objc.Void](p_, objc.Sel("setPositioningRect:"), value)
}

// Specifies if the popover is to be animated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1526527-animates?language=objc
func (p_ Popover) Animates() bool {
	rv := objc.Call[bool](p_, objc.Sel("animates"))
	return rv
}

// Specifies if the popover is to be animated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1526527-animates?language=objc
func (p_ Popover) SetAnimates(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAnimates:"), value)
}

// The appearance of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1529859-appearance?language=objc
func (p_ Popover) Appearance() Appearance {
	rv := objc.Call[Appearance](p_, objc.Sel("appearance"))
	return rv
}

// The appearance of the popover. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopover/1529859-appearance?language=objc
func (p_ Popover) SetAppearance(value IAppearance) {
	objc.Call[objc.Void](p_, objc.Sel("setAppearance:"), objc.Ptr(value))
}
