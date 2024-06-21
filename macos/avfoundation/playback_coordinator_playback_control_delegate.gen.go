// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"github.com/progrium/darwinkit/objc"
)

// A protocol that defines the method to implement to respond to playback commands from the playback coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinatorplaybackcontroldelegate?language=objc
type PPlaybackCoordinatorPlaybackControlDelegate interface {
	// optional
	PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator DelegatingPlaybackCoordinator, playCommand DelegatingPlaybackCoordinatorPlayCommand, completionHandler func())
	HasPlaybackCoordinatorDidIssuePlayCommandCompletionHandler() bool
}

// A delegate implementation builder for the [PPlaybackCoordinatorPlaybackControlDelegate] protocol.
type PlaybackCoordinatorPlaybackControlDelegate struct {
	_PlaybackCoordinatorDidIssuePlayCommandCompletionHandler func(coordinator DelegatingPlaybackCoordinator, playCommand DelegatingPlaybackCoordinatorPlayCommand, completionHandler func())
}

func (di *PlaybackCoordinatorPlaybackControlDelegate) HasPlaybackCoordinatorDidIssuePlayCommandCompletionHandler() bool {
	return di._PlaybackCoordinatorDidIssuePlayCommandCompletionHandler != nil
}

// Tells the delegate to match the playback rate to that of the group when the rate is nonzero. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinatorplaybackcontroldelegate/3750296-playbackcoordinator?language=objc
func (di *PlaybackCoordinatorPlaybackControlDelegate) SetPlaybackCoordinatorDidIssuePlayCommandCompletionHandler(f func(coordinator DelegatingPlaybackCoordinator, playCommand DelegatingPlaybackCoordinatorPlayCommand, completionHandler func())) {
	di._PlaybackCoordinatorDidIssuePlayCommandCompletionHandler = f
}

// Tells the delegate to match the playback rate to that of the group when the rate is nonzero. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinatorplaybackcontroldelegate/3750296-playbackcoordinator?language=objc
func (di *PlaybackCoordinatorPlaybackControlDelegate) PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator DelegatingPlaybackCoordinator, playCommand DelegatingPlaybackCoordinatorPlayCommand, completionHandler func()) {
	di._PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator, playCommand, completionHandler)
}

// ensure impl type implements protocol interface
var _ PPlaybackCoordinatorPlaybackControlDelegate = (*PlaybackCoordinatorPlaybackControlDelegateObject)(nil)

// A concrete type for the [PPlaybackCoordinatorPlaybackControlDelegate] protocol.
type PlaybackCoordinatorPlaybackControlDelegateObject struct {
	objc.Object
}

func (p_ PlaybackCoordinatorPlaybackControlDelegateObject) HasPlaybackCoordinatorDidIssuePlayCommandCompletionHandler() bool {
	return p_.RespondsToSelector(objc.Sel("playbackCoordinator:didIssuePlayCommand:completionHandler:"))
}

// Tells the delegate to match the playback rate to that of the group when the rate is nonzero. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinatorplaybackcontroldelegate/3750296-playbackcoordinator?language=objc
func (p_ PlaybackCoordinatorPlaybackControlDelegateObject) PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator DelegatingPlaybackCoordinator, playCommand DelegatingPlaybackCoordinatorPlayCommand, completionHandler func()) {
	objc.Call[objc.Void](p_, objc.Sel("playbackCoordinator:didIssuePlayCommand:completionHandler:"), coordinator, playCommand, completionHandler)
}
