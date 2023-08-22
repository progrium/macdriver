// AUTO-GENERATED CODE, DO NOT MODIFY

package avkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/avfoundation"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerView] class.
var PlayerViewClass = _PlayerViewClass{objc.GetClass("AVPlayerView")}

type _PlayerViewClass struct {
	objc.Class
}

// An interface definition for the [PlayerView] class.
type IPlayerView interface {
	appkit.IView
	BeginTrimmingWithCompletionHandler(handler func(result PlayerViewTrimResult))
	FlashChapterNumberChapterTitle(chapterNumber uint, chapterTitle string)
	VideoGravity() avfoundation.LayerVideoGravity
	SetVideoGravity(value avfoundation.LayerVideoGravity)
	ShowsFrameSteppingButtons() bool
	SetShowsFrameSteppingButtons(value bool)
	VideoBounds() foundation.Rect
	ContentOverlayView() appkit.View
	Player() avfoundation.Player
	SetPlayer(value avfoundation.IPlayer)
	ShowsSharingServiceButton() bool
	SetShowsSharingServiceButton(value bool)
	Delegate() PlayerViewDelegateWrapper
	SetDelegate(value PPlayerViewDelegate)
	SetDelegateObject(valueObject objc.IObject)
	IsReadyForDisplay() bool
	ShowsTimecodes() bool
	SetShowsTimecodes(value bool)
	AllowsPictureInPicturePlayback() bool
	SetAllowsPictureInPicturePlayback(value bool)
	ActionPopUpButtonMenu() appkit.Menu
	SetActionPopUpButtonMenu(value appkit.IMenu)
	UpdatesNowPlayingInfoCenter() bool
	SetUpdatesNowPlayingInfoCenter(value bool)
	CanBeginTrimming() bool
	ControlsStyle() PlayerViewControlsStyle
	SetControlsStyle(value PlayerViewControlsStyle)
	PictureInPictureDelegate() PlayerViewPictureInPictureDelegateWrapper
	SetPictureInPictureDelegate(value PPlayerViewPictureInPictureDelegate)
	SetPictureInPictureDelegateObject(valueObject objc.IObject)
	ShowsFullScreenToggleButton() bool
	SetShowsFullScreenToggleButton(value bool)
}

// A view that displays content from a player and presents a native user interface to control playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview?language=objc
type PlayerView struct {
	appkit.View
}

func PlayerViewFrom(ptr unsafe.Pointer) PlayerView {
	return PlayerView{
		View: appkit.ViewFrom(ptr),
	}
}

func (pc _PlayerViewClass) Alloc() PlayerView {
	rv := objc.Call[PlayerView](pc, objc.Sel("alloc"))
	return rv
}

func PlayerView_Alloc() PlayerView {
	return PlayerViewClass.Alloc()
}

func (pc _PlayerViewClass) New() PlayerView {
	rv := objc.Call[PlayerView](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerView() PlayerView {
	return PlayerViewClass.New()
}

func (p_ PlayerView) Init() PlayerView {
	rv := objc.Call[PlayerView](p_, objc.Sel("init"))
	return rv
}

func (p_ PlayerView) InitWithFrame(frameRect foundation.Rect) PlayerView {
	rv := objc.Call[PlayerView](p_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewPlayerViewWithFrame(frameRect foundation.Rect) PlayerView {
	instance := PlayerViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Puts the player view into trimming mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416570-begintrimmingwithcompletionhandl?language=objc
func (p_ PlayerView) BeginTrimmingWithCompletionHandler(handler func(result PlayerViewTrimResult)) {
	objc.Call[objc.Void](p_, objc.Sel("beginTrimmingWithCompletionHandler:"), handler)
}

// Displays the chapter number and title in the player view for a brief moment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416547-flashchapternumber?language=objc
func (p_ PlayerView) FlashChapterNumberChapterTitle(chapterNumber uint, chapterTitle string) {
	objc.Call[objc.Void](p_, objc.Sel("flashChapterNumber:chapterTitle:"), chapterNumber, chapterTitle)
}

// A value that determines how the player view displays video content within its bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416559-videogravity?language=objc
func (p_ PlayerView) VideoGravity() avfoundation.LayerVideoGravity {
	rv := objc.Call[avfoundation.LayerVideoGravity](p_, objc.Sel("videoGravity"))
	return rv
}

// A value that determines how the player view displays video content within its bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416559-videogravity?language=objc
func (p_ PlayerView) SetVideoGravity(value avfoundation.LayerVideoGravity) {
	objc.Call[objc.Void](p_, objc.Sel("setVideoGravity:"), value)
}

// A Boolean value that determines whether the player view displays frame stepping buttons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416567-showsframesteppingbuttons?language=objc
func (p_ PlayerView) ShowsFrameSteppingButtons() bool {
	rv := objc.Call[bool](p_, objc.Sel("showsFrameSteppingButtons"))
	return rv
}

// A Boolean value that determines whether the player view displays frame stepping buttons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416567-showsframesteppingbuttons?language=objc
func (p_ PlayerView) SetShowsFrameSteppingButtons(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setShowsFrameSteppingButtons:"), value)
}

// The current size and position of the video image that displays within the player view’s bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416549-videobounds?language=objc
func (p_ PlayerView) VideoBounds() foundation.Rect {
	rv := objc.Call[foundation.Rect](p_, objc.Sel("videoBounds"))
	return rv
}

// A view that adds additional custom views between the video content and the controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416573-contentoverlayview?language=objc
func (p_ PlayerView) ContentOverlayView() appkit.View {
	rv := objc.Call[appkit.View](p_, objc.Sel("contentOverlayView"))
	return rv
}

// The player instance that provides the media content for the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416539-player?language=objc
func (p_ PlayerView) Player() avfoundation.Player {
	rv := objc.Call[avfoundation.Player](p_, objc.Sel("player"))
	return rv
}

// The player instance that provides the media content for the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416539-player?language=objc
func (p_ PlayerView) SetPlayer(value avfoundation.IPlayer) {
	objc.Call[objc.Void](p_, objc.Sel("setPlayer:"), objc.Ptr(value))
}

// A Boolean value that determines whether the player view displays a sharing service button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416558-showssharingservicebutton?language=objc
func (p_ PlayerView) ShowsSharingServiceButton() bool {
	rv := objc.Call[bool](p_, objc.Sel("showsSharingServiceButton"))
	return rv
}

// A Boolean value that determines whether the player view displays a sharing service button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416558-showssharingservicebutton?language=objc
func (p_ PlayerView) SetShowsSharingServiceButton(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setShowsSharingServiceButton:"), value)
}

// The player view’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/3752984-delegate?language=objc
func (p_ PlayerView) Delegate() PlayerViewDelegateWrapper {
	rv := objc.Call[PlayerViewDelegateWrapper](p_, objc.Sel("delegate"))
	return rv
}

// The player view’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/3752984-delegate?language=objc
func (p_ PlayerView) SetDelegate(value PPlayerViewDelegate) {
	po0 := objc.WrapAsProtocol("AVPlayerViewDelegate", value)
	objc.SetAssociatedObject(p_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), po0)
}

// The player view’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/3752984-delegate?language=objc
func (p_ PlayerView) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// A Boolean value that indicates whether the current player item’s first video frame is ready for display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416556-readyfordisplay?language=objc
func (p_ PlayerView) IsReadyForDisplay() bool {
	rv := objc.Call[bool](p_, objc.Sel("isReadyForDisplay"))
	return rv
}

// A Boolean value that determines whether the player view displays timecodes, if available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/3174919-showstimecodes?language=objc
func (p_ PlayerView) ShowsTimecodes() bool {
	rv := objc.Call[bool](p_, objc.Sel("showsTimecodes"))
	return rv
}

// A Boolean value that determines whether the player view displays timecodes, if available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/3174919-showstimecodes?language=objc
func (p_ PlayerView) SetShowsTimecodes(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setShowsTimecodes:"), value)
}

// A Boolean value that determines whether the player view allows Picture in Picture playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/3172688-allowspictureinpictureplayback?language=objc
func (p_ PlayerView) AllowsPictureInPicturePlayback() bool {
	rv := objc.Call[bool](p_, objc.Sel("allowsPictureInPicturePlayback"))
	return rv
}

// A Boolean value that determines whether the player view allows Picture in Picture playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/3172688-allowspictureinpictureplayback?language=objc
func (p_ PlayerView) SetAllowsPictureInPicturePlayback(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setAllowsPictureInPicturePlayback:"), value)
}

// An action pop-up button menu that the player view displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416543-actionpopupbuttonmenu?language=objc
func (p_ PlayerView) ActionPopUpButtonMenu() appkit.Menu {
	rv := objc.Call[appkit.Menu](p_, objc.Sel("actionPopUpButtonMenu"))
	return rv
}

// An action pop-up button menu that the player view displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416543-actionpopupbuttonmenu?language=objc
func (p_ PlayerView) SetActionPopUpButtonMenu(value appkit.IMenu) {
	objc.Call[objc.Void](p_, objc.Sel("setActionPopUpButtonMenu:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the player view controller updates the Now Playing info center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/2876219-updatesnowplayinginfocenter?language=objc
func (p_ PlayerView) UpdatesNowPlayingInfoCenter() bool {
	rv := objc.Call[bool](p_, objc.Sel("updatesNowPlayingInfoCenter"))
	return rv
}

// A Boolean value that indicates whether the player view controller updates the Now Playing info center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/2876219-updatesnowplayinginfocenter?language=objc
func (p_ PlayerView) SetUpdatesNowPlayingInfoCenter(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setUpdatesNowPlayingInfoCenter:"), value)
}

// A Boolean value that indicates whether the player view can begin trimming. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416545-canbegintrimming?language=objc
func (p_ PlayerView) CanBeginTrimming() bool {
	rv := objc.Call[bool](p_, objc.Sel("canBeginTrimming"))
	return rv
}

// The player view’s controls style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416565-controlsstyle?language=objc
func (p_ PlayerView) ControlsStyle() PlayerViewControlsStyle {
	rv := objc.Call[PlayerViewControlsStyle](p_, objc.Sel("controlsStyle"))
	return rv
}

// The player view’s controls style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416565-controlsstyle?language=objc
func (p_ PlayerView) SetControlsStyle(value PlayerViewControlsStyle) {
	objc.Call[objc.Void](p_, objc.Sel("setControlsStyle:"), value)
}

// The Picture in Picture delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/3172689-pictureinpicturedelegate?language=objc
func (p_ PlayerView) PictureInPictureDelegate() PlayerViewPictureInPictureDelegateWrapper {
	rv := objc.Call[PlayerViewPictureInPictureDelegateWrapper](p_, objc.Sel("pictureInPictureDelegate"))
	return rv
}

// The Picture in Picture delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/3172689-pictureinpicturedelegate?language=objc
func (p_ PlayerView) SetPictureInPictureDelegate(value PPlayerViewPictureInPictureDelegate) {
	po0 := objc.WrapAsProtocol("AVPlayerViewPictureInPictureDelegate", value)
	objc.SetAssociatedObject(p_, objc.AssociationKey("setPictureInPictureDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](p_, objc.Sel("setPictureInPictureDelegate:"), po0)
}

// The Picture in Picture delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/3172689-pictureinpicturedelegate?language=objc
func (p_ PlayerView) SetPictureInPictureDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("setPictureInPictureDelegate:"), objc.Ptr(valueObject))
}

// A Boolean value that determines whether the player view displays a full-screen toggle button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416561-showsfullscreentogglebutton?language=objc
func (p_ PlayerView) ShowsFullScreenToggleButton() bool {
	rv := objc.Call[bool](p_, objc.Sel("showsFullScreenToggleButton"))
	return rv
}

// A Boolean value that determines whether the player view displays a full-screen toggle button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/1416561-showsfullscreentogglebutton?language=objc
func (p_ PlayerView) SetShowsFullScreenToggleButton(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setShowsFullScreenToggleButton:"), value)
}
