// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Event] class.
var EventClass = _EventClass{objc.GetClass("NSEvent")}

type _EventClass struct {
	objc.Class
}

// An interface definition for the [Event] class.
type IEvent interface {
	objc.IObject
	CoalescedTouchesForTouch(touch ITouch) []Touch
	TrackSwipeEventWithOptionsDampenAmountThresholdMinMaxUsingHandler(options EventSwipeTrackingOptions, minDampenThreshold float64, maxDampenThreshold float64, trackingHandler func(gestureAmount float64, phase EventPhase, isComplete bool, stop *bool))
	AllTouches() foundation.Set
	CharactersByApplyingModifiers(modifiers EventModifierFlags) string
	LocationInNode(node objc.IObject) coregraphics.Point
	TouchesForView(view IView) foundation.Set
	TouchesMatchingPhaseInView(phase TouchPhase, view IView) foundation.Set
	Pressure() float64
	AbsoluteZ() int
	SystemTabletID() uint
	TabletID() uint
	DeviceID() uint
	PointingDeviceType() PointingDeviceType
	Data1() int
	ButtonNumber() int
	EventRef() unsafe.Pointer
	DeltaZ() float64
	StageTransition() float64
	Tilt() foundation.Point
	Characters() string
	UserData() unsafe.Pointer
	PointingDeviceSerialNumber() uint
	VendorDefined() objc.Object
	EventNumber() int
	MomentumPhase() EventPhase
	TrackingArea() TrackingArea
	UniqueID() int64
	Rotation() float64
	PointingDeviceID() uint
	AssociatedEventsMask() EventMask
	Magnification() float64
	CGEvent() coregraphics.EventRef
	Subtype() EventSubtype
	DeltaX() float64
	VendorID() uint
	TrackingNumber() int
	ButtonMask() EventButtonMask
	VendorPointingDeviceType() uint
	DeltaY() float64
	KeyCode() int
	WindowNumber() int
	IsDirectionInvertedFromDevice() bool
	IsEnteringProximity() bool
	LocationInWindow() foundation.Point
	Timestamp() foundation.TimeInterval
	CapabilityMask() uint
	AbsoluteX() int
	Type() EventType
	ScrollingDeltaX() float64
	Window() Window
	IsARepeat() bool
	Data2() int
	TangentialPressure() float64
	Stage() int
	HasPreciseScrollingDeltas() bool
	ScrollingDeltaY() float64
	AbsoluteY() int
	ClickCount() int
	Phase() EventPhase
	CharactersIgnoringModifiers() string
	PressureBehavior() PressureBehavior
}

// An object that contains information about an input action, such as a mouse click or a key press. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent?language=objc
type Event struct {
	objc.Object
}

func EventFrom(ptr unsafe.Pointer) Event {
	return Event{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _EventClass) Alloc() Event {
	rv := objc.Call[Event](ec, objc.Sel("alloc"))
	return rv
}

func Event_Alloc() Event {
	return EventClass.Alloc()
}

func (ec _EventClass) New() Event {
	rv := objc.Call[Event](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewEvent() Event {
	return EventClass.New()
}

func (e_ Event) Init() Event {
	rv := objc.Call[Event](e_, objc.Sel("init"))
	return rv
}

// Creates and returns a new event object for a Carbon event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528021-eventwitheventref?language=objc
func (ec _EventClass) EventWithEventRef(eventRef unsafe.Pointer) Event {
	rv := objc.Call[Event](ec, objc.Sel("eventWithEventRef:"), eventRef)
	return rv
}

// Creates and returns a new event object for a Carbon event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528021-eventwitheventref?language=objc
func Event_EventWithEventRef(eventRef unsafe.Pointer) Event {
	return EventClass.EventWithEventRef(eventRef)
}

// Returns all of the touch objects associated with the specified main touch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/2646916-coalescedtouchesfortouch?language=objc
func (e_ Event) CoalescedTouchesForTouch(touch ITouch) []Touch {
	rv := objc.Call[[]Touch](e_, objc.Sel("coalescedTouchesForTouch:"), objc.Ptr(touch))
	return rv
}

// Allows tracking and user interface feedback of scroll wheel events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533300-trackswipeeventwithoptions?language=objc
func (e_ Event) TrackSwipeEventWithOptionsDampenAmountThresholdMinMaxUsingHandler(options EventSwipeTrackingOptions, minDampenThreshold float64, maxDampenThreshold float64, trackingHandler func(gestureAmount float64, phase EventPhase, isComplete bool, stop *bool)) {
	objc.Call[objc.Void](e_, objc.Sel("trackSwipeEventWithOptions:dampenAmountThresholdMin:max:usingHandler:"), options, minDampenThreshold, maxDampenThreshold, trackingHandler)
}

// Begins generating periodic events for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1526044-startperiodiceventsafterdelay?language=objc
func (ec _EventClass) StartPeriodicEventsAfterDelayWithPeriod(delay foundation.TimeInterval, period foundation.TimeInterval) {
	objc.Call[objc.Void](ec, objc.Sel("startPeriodicEventsAfterDelay:withPeriod:"), delay, period)
}

// Begins generating periodic events for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1526044-startperiodiceventsafterdelay?language=objc
func Event_StartPeriodicEventsAfterDelayWithPeriod(delay foundation.TimeInterval, period foundation.TimeInterval) {
	EventClass.StartPeriodicEventsAfterDelayWithPeriod(delay, period)
}

// Returns all touch objects associated with the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/2646917-alltouches?language=objc
func (e_ Event) AllTouches() foundation.Set {
	rv := objc.Call[foundation.Set](e_, objc.Sel("allTouches"))
	return rv
}

// Creates and returns a new event object that describes a tracking-rectangle or cursor-update event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1535383-enterexiteventwithtype?language=objc
func (ec _EventClass) EnterExitEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberTrackingNumberUserData(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, enum int, tNum int, data unsafe.Pointer) Event {
	rv := objc.Call[Event](ec, objc.Sel("enterExitEventWithType:location:modifierFlags:timestamp:windowNumber:context:eventNumber:trackingNumber:userData:"), type_, location, flags, time, wNum, objc.Ptr(unusedPassNil), enum, tNum, data)
	return rv
}

// Creates and returns a new event object that describes a tracking-rectangle or cursor-update event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1535383-enterexiteventwithtype?language=objc
func Event_EnterExitEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberTrackingNumberUserData(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, enum int, tNum int, data unsafe.Pointer) Event {
	return EventClass.EnterExitEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberTrackingNumberUserData(type_, location, flags, time, wNum, unusedPassNil, enum, tNum, data)
}

// Returns the new characters that result if you apply the specified modifier keys to the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/3242717-charactersbyapplyingmodifiers?language=objc
func (e_ Event) CharactersByApplyingModifiers(modifiers EventModifierFlags) string {
	rv := objc.Call[string](e_, objc.Sel("charactersByApplyingModifiers:"), modifiers)
	return rv
}

// Removes the specified event monitor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533709-removemonitor?language=objc
func (ec _EventClass) RemoveMonitor(eventMonitor objc.IObject) {
	objc.Call[objc.Void](ec, objc.Sel("removeMonitor:"), eventMonitor)
}

// Removes the specified event monitor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533709-removemonitor?language=objc
func Event_RemoveMonitor(eventMonitor objc.IObject) {
	EventClass.RemoveMonitor(eventMonitor)
}

// Returns the location of the receiver in the coordinate system of the given node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1483105-locationinnode?language=objc
func (e_ Event) LocationInNode(node objc.IObject) coregraphics.Point {
	rv := objc.Call[coregraphics.Point](e_, objc.Sel("locationInNode:"), objc.Ptr(node))
	return rv
}

// Returns the touch objects from the event that belong to the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/2646918-touchesforview?language=objc
func (e_ Event) TouchesForView(view IView) foundation.Set {
	rv := objc.Call[foundation.Set](e_, objc.Sel("touchesForView:"), objc.Ptr(view))
	return rv
}

// Stops generating periodic events for the current thread and discards any periodic events remaining in the queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533746-stopperiodicevents?language=objc
func (ec _EventClass) StopPeriodicEvents() {
	objc.Call[objc.Void](ec, objc.Sel("stopPeriodicEvents"))
}

// Stops generating periodic events for the current thread and discards any periodic events remaining in the queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533746-stopperiodicevents?language=objc
func Event_StopPeriodicEvents() {
	EventClass.StopPeriodicEvents()
}

// Creates and returns an event object for a Core Graphics event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1526299-eventwithcgevent?language=objc
func (ec _EventClass) EventWithCGEvent(cgEvent coregraphics.EventRef) Event {
	rv := objc.Call[Event](ec, objc.Sel("eventWithCGEvent:"), cgEvent)
	return rv
}

// Creates and returns an event object for a Core Graphics event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1526299-eventwithcgevent?language=objc
func Event_EventWithCGEvent(cgEvent coregraphics.EventRef) Event {
	return EventClass.EventWithCGEvent(cgEvent)
}

// Returns the touch objects associated with the specified phase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1530950-touchesmatchingphase?language=objc
func (e_ Event) TouchesMatchingPhaseInView(phase TouchPhase, view IView) foundation.Set {
	rv := objc.Call[foundation.Set](e_, objc.Sel("touchesMatchingPhase:inView:"), phase, objc.Ptr(view))
	return rv
}

// Installs an event monitor that receives copies of events the system posts to other applications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1535472-addglobalmonitorforeventsmatchin?language=objc
func (ec _EventClass) AddGlobalMonitorForEventsMatchingMaskHandler(mask EventMask, block func(event Event)) objc.Object {
	rv := objc.Call[objc.Object](ec, objc.Sel("addGlobalMonitorForEventsMatchingMask:handler:"), mask, block)
	return rv
}

// Installs an event monitor that receives copies of events the system posts to other applications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1535472-addglobalmonitorforeventsmatchin?language=objc
func Event_AddGlobalMonitorForEventsMatchingMaskHandler(mask EventMask, block func(event Event)) objc.Object {
	return EventClass.AddGlobalMonitorForEventsMatchingMaskHandler(mask, block)
}

// Installs an event monitor that receives copies of events the system posts to this app prior to their dispatch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1534971-addlocalmonitorforeventsmatching?language=objc
func (ec _EventClass) AddLocalMonitorForEventsMatchingMaskHandler(mask EventMask, block func(event Event) Event) objc.Object {
	rv := objc.Call[objc.Object](ec, objc.Sel("addLocalMonitorForEventsMatchingMask:handler:"), mask, block)
	return rv
}

// Installs an event monitor that receives copies of events the system posts to this app prior to their dispatch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1534971-addlocalmonitorforeventsmatching?language=objc
func Event_AddLocalMonitorForEventsMatchingMaskHandler(mask EventMask, block func(event Event) Event) objc.Object {
	return EventClass.AddLocalMonitorForEventsMatchingMaskHandler(mask, block)
}

// Creates and returns a new event object that describes a custom event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1530010-othereventwithtype?language=objc
func (ec _EventClass) OtherEventWithTypeLocationModifierFlagsTimestampWindowNumberContextSubtypeData1Data2(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, subtype int, d1 int, d2 int) Event {
	rv := objc.Call[Event](ec, objc.Sel("otherEventWithType:location:modifierFlags:timestamp:windowNumber:context:subtype:data1:data2:"), type_, location, flags, time, wNum, objc.Ptr(unusedPassNil), subtype, d1, d2)
	return rv
}

// Creates and returns a new event object that describes a custom event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1530010-othereventwithtype?language=objc
func Event_OtherEventWithTypeLocationModifierFlagsTimestampWindowNumberContextSubtypeData1Data2(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, subtype int, d1 int, d2 int) Event {
	return EventClass.OtherEventWithTypeLocationModifierFlagsTimestampWindowNumberContextSubtypeData1Data2(type_, location, flags, time, wNum, unusedPassNil, subtype, d1, d2)
}

// Creates and returns a new event object that describes a mouse-down, -up, -moved, or -dragged event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1532495-mouseeventwithtype?language=objc
func (ec _EventClass) MouseEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberClickCountPressure(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, enum int, cNum int, pressure float64) Event {
	rv := objc.Call[Event](ec, objc.Sel("mouseEventWithType:location:modifierFlags:timestamp:windowNumber:context:eventNumber:clickCount:pressure:"), type_, location, flags, time, wNum, objc.Ptr(unusedPassNil), enum, cNum, pressure)
	return rv
}

// Creates and returns a new event object that describes a mouse-down, -up, -moved, or -dragged event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1532495-mouseeventwithtype?language=objc
func Event_MouseEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberClickCountPressure(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, enum int, cNum int, pressure float64) Event {
	return EventClass.MouseEventWithTypeLocationModifierFlagsTimestampWindowNumberContextEventNumberClickCountPressure(type_, location, flags, time, wNum, unusedPassNil, enum, cNum, pressure)
}

// Creates and returns a new event object that describes a key event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533943-keyeventwithtype?language=objc
func (ec _EventClass) KeyEventWithTypeLocationModifierFlagsTimestampWindowNumberContextCharactersCharactersIgnoringModifiersIsARepeatKeyCode(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, keys string, ukeys string, flag bool, code int) Event {
	rv := objc.Call[Event](ec, objc.Sel("keyEventWithType:location:modifierFlags:timestamp:windowNumber:context:characters:charactersIgnoringModifiers:isARepeat:keyCode:"), type_, location, flags, time, wNum, objc.Ptr(unusedPassNil), keys, ukeys, flag, code)
	return rv
}

// Creates and returns a new event object that describes a key event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533943-keyeventwithtype?language=objc
func Event_KeyEventWithTypeLocationModifierFlagsTimestampWindowNumberContextCharactersCharactersIgnoringModifiersIsARepeatKeyCode(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, keys string, ukeys string, flag bool, code int) Event {
	return EventClass.KeyEventWithTypeLocationModifierFlagsTimestampWindowNumberContextCharactersCharactersIgnoringModifiersIsARepeatKeyCode(type_, location, flags, time, wNum, unusedPassNil, keys, ukeys, flag, code)
}

// A normalized value that indicates the degree of pressure applied to an appropriate input device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1534543-pressure?language=objc
func (e_ Event) Pressure() float64 {
	rv := objc.Call[float64](e_, objc.Sel("pressure"))
	return rv
}

// The absolute z coordinate of pointing device on its tablet at full tablet resolution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1532154-absolutez?language=objc
func (e_ Event) AbsoluteZ() int {
	rv := objc.Call[int](e_, objc.Sel("absoluteZ"))
	return rv
}

// The index of the tablet device connected to the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528299-systemtabletid?language=objc
func (e_ Event) SystemTabletID() uint {
	rv := objc.Call[uint](e_, objc.Sel("systemTabletID"))
	return rv
}

// The USB model identifier of the tablet device associated with this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1527003-tabletid?language=objc
func (e_ Event) TabletID() uint {
	rv := objc.Call[uint](e_, objc.Sel("tabletID"))
	return rv
}

// A special identifier the system matches against tablet-pointer and tablet-proximity events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1530014-deviceid?language=objc
func (e_ Event) DeviceID() uint {
	rv := objc.Call[uint](e_, objc.Sel("deviceID"))
	return rv
}

// The maximum number of seconds in which a second mouse click must occur for an event to be a double-click event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528384-doubleclickinterval?language=objc
func (ec _EventClass) DoubleClickInterval() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](ec, objc.Sel("doubleClickInterval"))
	return rv
}

// The maximum number of seconds in which a second mouse click must occur for an event to be a double-click event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528384-doubleclickinterval?language=objc
func Event_DoubleClickInterval() foundation.TimeInterval {
	return EventClass.DoubleClickInterval()
}

// The kind of pointing device associated with this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1535573-pointingdevicetype?language=objc
func (e_ Event) PointingDeviceType() PointingDeviceType {
	rv := objc.Call[PointingDeviceType](e_, objc.Sel("pointingDeviceType"))
	return rv
}

// Additional data associated with this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528289-data1?language=objc
func (e_ Event) Data1() int {
	rv := objc.Call[int](e_, objc.Sel("data1"))
	return rv
}

// A Boolean value that indicates whether the system coalesces mouse movement events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/2870068-mousecoalescingenabled?language=objc
func (ec _EventClass) MouseCoalescingEnabled() bool {
	rv := objc.Call[bool](ec, objc.Sel("mouseCoalescingEnabled"))
	return rv
}

// A Boolean value that indicates whether the system coalesces mouse movement events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/2870068-mousecoalescingenabled?language=objc
func Event_MouseCoalescingEnabled() bool {
	return EventClass.MouseCoalescingEnabled()
}

// A Boolean value that indicates whether the system coalesces mouse movement events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/2870068-mousecoalescingenabled?language=objc
func (ec _EventClass) SetMouseCoalescingEnabled(value bool) {
	objc.Call[objc.Void](ec, objc.Sel("setMouseCoalescingEnabled:"), value)
}

// A Boolean value that indicates whether the system coalesces mouse movement events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/2870068-mousecoalescingenabled?language=objc
func Event_SetMouseCoalescingEnabled(value bool) {
	EventClass.SetMouseCoalescingEnabled(value)
}

// The currently pressed modifier keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1535211-modifierflags?language=objc
func (ec _EventClass) ModifierFlags() EventModifierFlags {
	rv := objc.Call[EventModifierFlags](ec, objc.Sel("modifierFlags"))
	return rv
}

// The currently pressed modifier keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1535211-modifierflags?language=objc
func Event_ModifierFlags() EventModifierFlags {
	return EventClass.ModifierFlags()
}

// The button number for a mouse event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1527828-buttonnumber?language=objc
func (e_ Event) ButtonNumber() int {
	rv := objc.Call[int](e_, objc.Sel("buttonNumber"))
	return rv
}

// An opaque Carbon type associated with this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1525143-eventref?language=objc
func (e_ Event) EventRef() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](e_, objc.Sel("eventRef"))
	return rv
}

// The z-coordinate change for a scroll wheel, mouse-move, or mouse-drag event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1531528-deltaz?language=objc
func (e_ Event) DeltaZ() float64 {
	rv := objc.Call[float64](e_, objc.Sel("deltaZ"))
	return rv
}

// The transition value for the stage of a pressure gesture event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1526739-stagetransition?language=objc
func (e_ Event) StageTransition() float64 {
	rv := objc.Call[float64](e_, objc.Sel("stageTransition"))
	return rv
}

// The scaled tilt values of the pointing device that generated this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1534226-tilt?language=objc
func (e_ Event) Tilt() foundation.Point {
	rv := objc.Call[foundation.Point](e_, objc.Sel("tilt"))
	return rv
}

// The characters associated with a key-up or key-down event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1534183-characters?language=objc
func (e_ Event) Characters() string {
	rv := objc.Call[string](e_, objc.Sel("characters"))
	return rv
}

// The data associated with a mouse-tracking event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1526810-userdata?language=objc
func (e_ Event) UserData() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](e_, objc.Sel("userData"))
	return rv
}

// The vendor-assigned serial number of a pointing device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533420-pointingdeviceserialnumber?language=objc
func (e_ Event) PointingDeviceSerialNumber() uint {
	rv := objc.Call[uint](e_, objc.Sel("pointingDeviceSerialNumber"))
	return rv
}

// An array of three vendor-defined number objects associated with a pointing-type event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1530551-vendordefined?language=objc
func (e_ Event) VendorDefined() objc.Object {
	rv := objc.Call[objc.Object](e_, objc.Sel("vendorDefined"))
	return rv
}

// The counter value of the latest mouse or tracking-rectangle event object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1535220-eventnumber?language=objc
func (e_ Event) EventNumber() int {
	rv := objc.Call[int](e_, objc.Sel("eventNumber"))
	return rv
}

// The number of seconds someone must hold down a key to generate key-repeat events after the initial delay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1526076-keyrepeatinterval?language=objc
func (ec _EventClass) KeyRepeatInterval() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](ec, objc.Sel("keyRepeatInterval"))
	return rv
}

// The number of seconds someone must hold down a key to generate key-repeat events after the initial delay. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1526076-keyrepeatinterval?language=objc
func Event_KeyRepeatInterval() foundation.TimeInterval {
	return EventClass.KeyRepeatInterval()
}

// The number of seconds someone must hold down a key before the first key repeat event occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1530832-keyrepeatdelay?language=objc
func (ec _EventClass) KeyRepeatDelay() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](ec, objc.Sel("keyRepeatDelay"))
	return rv
}

// The number of seconds someone must hold down a key before the first key repeat event occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1530832-keyrepeatdelay?language=objc
func Event_KeyRepeatDelay() foundation.TimeInterval {
	return EventClass.KeyRepeatDelay()
}

// The momentum phase for a scroll or flick gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1525439-momentumphase?language=objc
func (e_ Event) MomentumPhase() EventPhase {
	rv := objc.Call[EventPhase](e_, objc.Sel("momentumPhase"))
	return rv
}

// The tracking area for the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1534800-trackingarea?language=objc
func (e_ Event) TrackingArea() TrackingArea {
	rv := objc.Call[TrackingArea](e_, objc.Sel("trackingArea"))
	return rv
}

// The unique identifier of the pointing device that generated this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1535813-uniqueid?language=objc
func (e_ Event) UniqueID() int64 {
	rv := objc.Call[int64](e_, objc.Sel("uniqueID"))
	return rv
}

// The rotation in degrees of the tablet pointing device associated with this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1526249-rotation?language=objc
func (e_ Event) Rotation() float64 {
	rv := objc.Call[float64](e_, objc.Sel("rotation"))
	return rv
}

// The index of the pointing device currently in proximity with the tablet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528818-pointingdeviceid?language=objc
func (e_ Event) PointingDeviceID() uint {
	rv := objc.Call[uint](e_, objc.Sel("pointingDeviceID"))
	return rv
}

// The associated events mask of a mouse event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1529610-associatedeventsmask?language=objc
func (e_ Event) AssociatedEventsMask() EventMask {
	rv := objc.Call[EventMask](e_, objc.Sel("associatedEventsMask"))
	return rv
}

// The amount of change to add to a magnification gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1531642-magnification?language=objc
func (e_ Event) Magnification() float64 {
	rv := objc.Call[float64](e_, objc.Sel("magnification"))
	return rv
}

// The Core Graphics event object corresponding to this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1530429-cgevent?language=objc
func (e_ Event) CGEvent() coregraphics.EventRef {
	rv := objc.Call[coregraphics.EventRef](e_, objc.Sel("CGEvent"))
	return rv
}

// Reports the current mouse position in screen coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533380-mouselocation?language=objc
func (ec _EventClass) MouseLocation() foundation.Point {
	rv := objc.Call[foundation.Point](ec, objc.Sel("mouseLocation"))
	return rv
}

// Reports the current mouse position in screen coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533380-mouselocation?language=objc
func Event_MouseLocation() foundation.Point {
	return EventClass.MouseLocation()
}

// The event’s subtype. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1527726-subtype?language=objc
func (e_ Event) Subtype() EventSubtype {
	rv := objc.Call[EventSubtype](e_, objc.Sel("subtype"))
	return rv
}

// The x-coordinate change for scroll wheel, mouse-move, mouse-drag, and swipe events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1534871-deltax?language=objc
func (e_ Event) DeltaX() float64 {
	rv := objc.Call[float64](e_, objc.Sel("deltaX"))
	return rv
}

// The vendor identifier of the tablet associated with the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1525177-vendorid?language=objc
func (e_ Event) VendorID() uint {
	rv := objc.Call[uint](e_, objc.Sel("vendorID"))
	return rv
}

// The indices of the currently pressed mouse buttons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1527943-pressedmousebuttons?language=objc
func (ec _EventClass) PressedMouseButtons() uint {
	rv := objc.Call[uint](ec, objc.Sel("pressedMouseButtons"))
	return rv
}

// The indices of the currently pressed mouse buttons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1527943-pressedmousebuttons?language=objc
func Event_PressedMouseButtons() uint {
	return EventClass.PressedMouseButtons()
}

// The identifier of a mouse-tracking event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533974-trackingnumber?language=objc
func (e_ Event) TrackingNumber() int {
	rv := objc.Call[int](e_, objc.Sel("trackingNumber"))
	return rv
}

// A bit mask identifying the buttons pressed for a tablet event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1535428-buttonmask?language=objc
func (e_ Event) ButtonMask() EventButtonMask {
	rv := objc.Call[EventButtonMask](e_, objc.Sel("buttonMask"))
	return rv
}

// A coded bit field whose set bits indicate the type of pointing device (within a vendor selection) associated with the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1527736-vendorpointingdevicetype?language=objc
func (e_ Event) VendorPointingDeviceType() uint {
	rv := objc.Call[uint](e_, objc.Sel("vendorPointingDeviceType"))
	return rv
}

// The y-coordinate change for scroll wheel, mouse-move, mouse-drag, and swipe events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1534158-deltay?language=objc
func (e_ Event) DeltaY() float64 {
	rv := objc.Call[float64](e_, objc.Sel("deltaY"))
	return rv
}

// The virtual code for the key associated with the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1534513-keycode?language=objc
func (e_ Event) KeyCode() int {
	rv := objc.Call[int](e_, objc.Sel("keyCode"))
	return rv
}

// A Boolean value that indicates whether to track fluid swipe gestures using scroll events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/2870067-swipetrackingfromscrolleventsena?language=objc
func (ec _EventClass) SwipeTrackingFromScrollEventsEnabled() bool {
	rv := objc.Call[bool](ec, objc.Sel("swipeTrackingFromScrollEventsEnabled"))
	return rv
}

// A Boolean value that indicates whether to track fluid swipe gestures using scroll events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/2870067-swipetrackingfromscrolleventsena?language=objc
func Event_SwipeTrackingFromScrollEventsEnabled() bool {
	return EventClass.SwipeTrackingFromScrollEventsEnabled()
}

// The identifier for the window device associated with the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1531361-windownumber?language=objc
func (e_ Event) WindowNumber() int {
	rv := objc.Call[int](e_, objc.Sel("windowNumber"))
	return rv
}

// A Boolean value that indicates whether the user has changed the device inversion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1525151-directioninvertedfromdevice?language=objc
func (e_ Event) IsDirectionInvertedFromDevice() bool {
	rv := objc.Call[bool](e_, objc.Sel("isDirectionInvertedFromDevice"))
	return rv
}

// A Boolean value that indicates whether a pointing device is entering or leaving the proximity of its tablet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1531702-enteringproximity?language=objc
func (e_ Event) IsEnteringProximity() bool {
	rv := objc.Call[bool](e_, objc.Sel("isEnteringProximity"))
	return rv
}

// The event location in the base coordinate system of the associated window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1529068-locationinwindow?language=objc
func (e_ Event) LocationInWindow() foundation.Point {
	rv := objc.Call[foundation.Point](e_, objc.Sel("locationInWindow"))
	return rv
}

// The time when the event occurred in seconds since system startup. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528239-timestamp?language=objc
func (e_ Event) Timestamp() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](e_, objc.Sel("timestamp"))
	return rv
}

// A mask that indicates the capabilities of the tablet device that generated this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1534648-capabilitymask?language=objc
func (e_ Event) CapabilityMask() uint {
	rv := objc.Call[uint](e_, objc.Sel("capabilityMask"))
	return rv
}

// The absolute x coordinate of a pointing device on its tablet at full tablet resolution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1530617-absolutex?language=objc
func (e_ Event) AbsoluteX() int {
	rv := objc.Call[int](e_, objc.Sel("absoluteX"))
	return rv
}

// The event’s type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528439-type?language=objc
func (e_ Event) Type() EventType {
	rv := objc.Call[EventType](e_, objc.Sel("type"))
	return rv
}

// The scroll wheel’s horizontal delta. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1524505-scrollingdeltax?language=objc
func (e_ Event) ScrollingDeltaX() float64 {
	rv := objc.Call[float64](e_, objc.Sel("scrollingDeltaX"))
	return rv
}

// The window object associated with the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1530808-window?language=objc
func (e_ Event) Window() Window {
	rv := objc.Call[Window](e_, objc.Sel("window"))
	return rv
}

// A Boolean value that indicates whether the key event is a repeat. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528049-arepeat?language=objc
func (e_ Event) IsARepeat() bool {
	rv := objc.Call[bool](e_, objc.Sel("isARepeat"))
	return rv
}

// Additional data associated with this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528647-data2?language=objc
func (e_ Event) Data2() int {
	rv := objc.Call[int](e_, objc.Sel("data2"))
	return rv
}

// The tangential pressure on the device that generated this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1525959-tangentialpressure?language=objc
func (e_ Event) TangentialPressure() float64 {
	rv := objc.Call[float64](e_, objc.Sel("tangentialPressure"))
	return rv
}

// A value that indicates the stage of a pressure gesture event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1527242-stage?language=objc
func (e_ Event) Stage() int {
	rv := objc.Call[int](e_, objc.Sel("stage"))
	return rv
}

// A Boolean value that indicates whether precise scrolling deltas are available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1525758-hasprecisescrollingdeltas?language=objc
func (e_ Event) HasPreciseScrollingDeltas() bool {
	rv := objc.Call[bool](e_, objc.Sel("hasPreciseScrollingDeltas"))
	return rv
}

// The scroll wheel’s vertical delta. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1535387-scrollingdeltay?language=objc
func (e_ Event) ScrollingDeltaY() float64 {
	rv := objc.Call[float64](e_, objc.Sel("scrollingDeltaY"))
	return rv
}

// The absolute y coordinate of a pointing device on its tablet at full tablet resolution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528904-absolutey?language=objc
func (e_ Event) AbsoluteY() int {
	rv := objc.Call[int](e_, objc.Sel("absoluteY"))
	return rv
}

// The number of mouse clicks associated with a mouse-down or mouse-up event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1528200-clickcount?language=objc
func (e_ Event) ClickCount() int {
	rv := objc.Call[int](e_, objc.Sel("clickCount"))
	return rv
}

// The phase of a gesture event, such as a magnify, scroll, or pressure change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1533550-phase?language=objc
func (e_ Event) Phase() EventPhase {
	rv := objc.Call[EventPhase](e_, objc.Sel("phase"))
	return rv
}

// The characters generated by a key event as if no modifier key (except for Shift) applies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1524605-charactersignoringmodifiers?language=objc
func (e_ Event) CharactersIgnoringModifiers() string {
	rv := objc.Call[string](e_, objc.Sel("charactersIgnoringModifiers"))
	return rv
}

// The behavior and progression for a pressure event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsevent/1531392-pressurebehavior?language=objc
func (e_ Event) PressureBehavior() PressureBehavior {
	rv := objc.Call[PressureBehavior](e_, objc.Sel("pressureBehavior"))
	return rv
}
