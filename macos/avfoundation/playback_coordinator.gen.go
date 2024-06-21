// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [PlaybackCoordinator] class.
var PlaybackCoordinatorClass = _PlaybackCoordinatorClass{objc.GetClass("AVPlaybackCoordinator")}

type _PlaybackCoordinatorClass struct {
	objc.Class
}

// An interface definition for the [PlaybackCoordinator] class.
type IPlaybackCoordinator interface {
	objc.IObject
	ExpectedItemTimeAtHostTime(hostClockTime coremedia.Time) coremedia.Time
	SetParticipantLimitForWaitingOutSuspensionsWithReason(participantLimit int, reason CoordinatedPlaybackSuspensionReason)
	BeginSuspensionForReason(suspensionReason CoordinatedPlaybackSuspensionReason) CoordinatedPlaybackSuspension
	ParticipantLimitForWaitingOutSuspensionsWithReason(reason CoordinatedPlaybackSuspensionReason) int
	PauseSnapsToMediaTimeOfOriginator() bool
	SetPauseSnapsToMediaTimeOfOriginator(value bool)
	OtherParticipants() []CoordinatedPlaybackParticipant
	SuspensionReasons() []CoordinatedPlaybackSuspensionReason
	SuspensionReasonsThatTriggerWaiting() []CoordinatedPlaybackSuspensionReason
	SetSuspensionReasonsThatTriggerWaiting(value []CoordinatedPlaybackSuspensionReason)
}

// An object that coordinates the playback of players in a connected group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator?language=objc
type PlaybackCoordinator struct {
	objc.Object
}

func PlaybackCoordinatorFrom(ptr unsafe.Pointer) PlaybackCoordinator {
	return PlaybackCoordinator{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlaybackCoordinatorClass) Alloc() PlaybackCoordinator {
	rv := objc.Call[PlaybackCoordinator](pc, objc.Sel("alloc"))
	return rv
}

func (pc _PlaybackCoordinatorClass) New() PlaybackCoordinator {
	rv := objc.Call[PlaybackCoordinator](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlaybackCoordinator() PlaybackCoordinator {
	return PlaybackCoordinatorClass.New()
}

func (p_ PlaybackCoordinator) Init() PlaybackCoordinator {
	rv := objc.Call[PlaybackCoordinator](p_, objc.Sel("init"))
	return rv
}

// Returns a time in the current item’s timeline that the coordinator expects to play at the specified host time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/3750285-expecteditemtimeathosttime?language=objc
func (p_ PlaybackCoordinator) ExpectedItemTimeAtHostTime(hostClockTime coremedia.Time) coremedia.Time {
	rv := objc.Call[coremedia.Time](p_, objc.Sel("expectedItemTimeAtHostTime:"), hostClockTime)
	return rv
}

// Sets a limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/3750289-setparticipantlimit?language=objc
func (p_ PlaybackCoordinator) SetParticipantLimitForWaitingOutSuspensionsWithReason(participantLimit int, reason CoordinatedPlaybackSuspensionReason) {
	objc.Call[objc.Void](p_, objc.Sel("setParticipantLimit:forWaitingOutSuspensionsWithReason:"), participantLimit, reason)
}

// Tells the coordinator to stop sending playback commands temporarily when the playback object disconnects from the group activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/3750284-beginsuspensionforreason?language=objc
func (p_ PlaybackCoordinator) BeginSuspensionForReason(suspensionReason CoordinatedPlaybackSuspensionReason) CoordinatedPlaybackSuspension {
	rv := objc.Call[CoordinatedPlaybackSuspension](p_, objc.Sel("beginSuspensionForReason:"), suspensionReason)
	return rv
}

// Returns the limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/3750287-participantlimitforwaitingoutsus?language=objc
func (p_ PlaybackCoordinator) ParticipantLimitForWaitingOutSuspensionsWithReason(reason CoordinatedPlaybackSuspensionReason) int {
	rv := objc.Call[int](p_, objc.Sel("participantLimitForWaitingOutSuspensionsWithReason:"), reason)
	return rv
}

// A Boolean value that indicates whether participants mirror the originator’s stop time when they pause. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/3750288-pausesnapstomediatimeoforiginato?language=objc
func (p_ PlaybackCoordinator) PauseSnapsToMediaTimeOfOriginator() bool {
	rv := objc.Call[bool](p_, objc.Sel("pauseSnapsToMediaTimeOfOriginator"))
	return rv
}

// A Boolean value that indicates whether participants mirror the originator’s stop time when they pause. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/3750288-pausesnapstomediatimeoforiginato?language=objc
func (p_ PlaybackCoordinator) SetPauseSnapsToMediaTimeOfOriginator(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setPauseSnapsToMediaTimeOfOriginator:"), value)
}

// The identifiers of the other participants in a group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/3750286-otherparticipants?language=objc
func (p_ PlaybackCoordinator) OtherParticipants() []CoordinatedPlaybackParticipant {
	rv := objc.Call[[]CoordinatedPlaybackParticipant](p_, objc.Sel("otherParticipants"))
	return rv
}

// The reasons a coordinator is currently unable to participate in a group playback activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/3750290-suspensionreasons?language=objc
func (p_ PlaybackCoordinator) SuspensionReasons() []CoordinatedPlaybackSuspensionReason {
	rv := objc.Call[[]CoordinatedPlaybackSuspensionReason](p_, objc.Sel("suspensionReasons"))
	return rv
}

// The reasons that cause a coordinator to suspend playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/3750291-suspensionreasonsthattriggerwait?language=objc
func (p_ PlaybackCoordinator) SuspensionReasonsThatTriggerWaiting() []CoordinatedPlaybackSuspensionReason {
	rv := objc.Call[[]CoordinatedPlaybackSuspensionReason](p_, objc.Sel("suspensionReasonsThatTriggerWaiting"))
	return rv
}

// The reasons that cause a coordinator to suspend playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/3750291-suspensionreasonsthattriggerwait?language=objc
func (p_ PlaybackCoordinator) SetSuspensionReasonsThatTriggerWaiting(value []CoordinatedPlaybackSuspensionReason) {
	objc.Call[objc.Void](p_, objc.Sel("setSuspensionReasonsThatTriggerWaiting:"), value)
}
