// Code generated by DarwinKit. DO NOT EDIT.

package avkit

import (
	"github.com/progrium/macdriver/macos/avfoundation"
	"github.com/progrium/macdriver/objc"
)

// The protocol that defines the methods you can implement to respond to capture view events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureviewdelegate?language=objc
type PCaptureViewDelegate interface {
	// optional
	CaptureViewStartRecordingToFileOutput(captureView CaptureView, fileOutput avfoundation.CaptureFileOutput)
	HasCaptureViewStartRecordingToFileOutput() bool
}

// A delegate implementation builder for the [PCaptureViewDelegate] protocol.
type CaptureViewDelegate struct {
	_CaptureViewStartRecordingToFileOutput func(captureView CaptureView, fileOutput avfoundation.CaptureFileOutput)
}

func (di *CaptureViewDelegate) HasCaptureViewStartRecordingToFileOutput() bool {
	return di._CaptureViewStartRecordingToFileOutput != nil
}

// Tells the delegate that the user has made a request to start a new recording. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureviewdelegate/1519138-captureview?language=objc
func (di *CaptureViewDelegate) SetCaptureViewStartRecordingToFileOutput(f func(captureView CaptureView, fileOutput avfoundation.CaptureFileOutput)) {
	di._CaptureViewStartRecordingToFileOutput = f
}

// Tells the delegate that the user has made a request to start a new recording. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureviewdelegate/1519138-captureview?language=objc
func (di *CaptureViewDelegate) CaptureViewStartRecordingToFileOutput(captureView CaptureView, fileOutput avfoundation.CaptureFileOutput) {
	di._CaptureViewStartRecordingToFileOutput(captureView, fileOutput)
}

// ensure impl type implements protocol interface
var _ PCaptureViewDelegate = (*CaptureViewDelegateObject)(nil)

// A concrete type for the [PCaptureViewDelegate] protocol.
type CaptureViewDelegateObject struct {
	objc.Object
}

func (c_ CaptureViewDelegateObject) HasCaptureViewStartRecordingToFileOutput() bool {
	return c_.RespondsToSelector(objc.Sel("captureView:startRecordingToFileOutput:"))
}

// Tells the delegate that the user has made a request to start a new recording. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureviewdelegate/1519138-captureview?language=objc
func (c_ CaptureViewDelegateObject) CaptureViewStartRecordingToFileOutput(captureView CaptureView, fileOutput avfoundation.CaptureFileOutput) {
	objc.Call[objc.Void](c_, objc.Sel("captureView:startRecordingToFileOutput:"), objc.Ptr(captureView), objc.Ptr(fileOutput))
}
