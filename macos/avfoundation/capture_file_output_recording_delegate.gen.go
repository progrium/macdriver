// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// Methods for responding to events that occur while recording captured media to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate?language=objc
type PCaptureFileOutputRecordingDelegate interface {
	// optional
	CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection, error foundation.Error)
	HasCaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError() bool

	// optional
	CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)
	HasCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections() bool

	// optional
	CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)
	HasCaptureOutputDidStartRecordingToOutputFileAtURLFromConnections() bool

	// optional
	CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)
	HasCaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections() bool

	// optional
	CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(output CaptureFileOutput, outputFileURL foundation.URL, connections []CaptureConnection, error foundation.Error)
	HasCaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError() bool
}

// A delegate implementation builder for the [PCaptureFileOutputRecordingDelegate] protocol.
type CaptureFileOutputRecordingDelegate struct {
	_CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError func(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection, error foundation.Error)
	_CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections        func(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)
	_CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections        func(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)
	_CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections       func(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)
	_CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError  func(output CaptureFileOutput, outputFileURL foundation.URL, connections []CaptureConnection, error foundation.Error)
}

func (di *CaptureFileOutputRecordingDelegate) HasCaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError() bool {
	return di._CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError != nil
}

// Informs the delegate when the output will stop writing new samples to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1390625-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) SetCaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError(f func(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection, error foundation.Error)) {
	di._CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError = f
}

// Informs the delegate when the output will stop writing new samples to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1390625-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection, error foundation.Error) {
	di._CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError(output, fileURL, connections, error)
}
func (di *CaptureFileOutputRecordingDelegate) HasCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections() bool {
	return di._CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections != nil
}

// Called whenever the output is recording to a file and successfully pauses the recording at the request of a client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1388838-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) SetCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(f func(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)) {
	di._CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections = f
}

// Called whenever the output is recording to a file and successfully pauses the recording at the request of a client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1388838-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection) {
	di._CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output, fileURL, connections)
}
func (di *CaptureFileOutputRecordingDelegate) HasCaptureOutputDidStartRecordingToOutputFileAtURLFromConnections() bool {
	return di._CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections != nil
}

// Informs the delegate when the output has started writing to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1387301-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) SetCaptureOutputDidStartRecordingToOutputFileAtURLFromConnections(f func(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)) {
	di._CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections = f
}

// Informs the delegate when the output has started writing to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1387301-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection) {
	di._CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections(output, fileURL, connections)
}
func (di *CaptureFileOutputRecordingDelegate) HasCaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections() bool {
	return di._CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections != nil
}

// Called whenever the output, at the request of the client, successfully resumes a file recording that was paused. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1387653-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) SetCaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections(f func(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection)) {
	di._CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections = f
}

// Called whenever the output, at the request of the client, successfully resumes a file recording that was paused. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1387653-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection) {
	di._CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections(output, fileURL, connections)
}
func (di *CaptureFileOutputRecordingDelegate) HasCaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError() bool {
	return di._CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError != nil
}

// Informs the delegate when all pending data has been written to an output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1390612-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) SetCaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(f func(output CaptureFileOutput, outputFileURL foundation.URL, connections []CaptureConnection, error foundation.Error)) {
	di._CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError = f
}

// Informs the delegate when all pending data has been written to an output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1390612-captureoutput?language=objc
func (di *CaptureFileOutputRecordingDelegate) CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(output CaptureFileOutput, outputFileURL foundation.URL, connections []CaptureConnection, error foundation.Error) {
	di._CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(output, outputFileURL, connections, error)
}

// ensure impl type implements protocol interface
var _ PCaptureFileOutputRecordingDelegate = (*CaptureFileOutputRecordingDelegateObject)(nil)

// A concrete type for the [PCaptureFileOutputRecordingDelegate] protocol.
type CaptureFileOutputRecordingDelegateObject struct {
	objc.Object
}

func (c_ CaptureFileOutputRecordingDelegateObject) HasCaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError() bool {
	return c_.RespondsToSelector(objc.Sel("captureOutput:willFinishRecordingToOutputFileAtURL:fromConnections:error:"))
}

// Informs the delegate when the output will stop writing new samples to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1390625-captureoutput?language=objc
func (c_ CaptureFileOutputRecordingDelegateObject) CaptureOutputWillFinishRecordingToOutputFileAtURLFromConnectionsError(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection, error foundation.Error) {
	objc.Call[objc.Void](c_, objc.Sel("captureOutput:willFinishRecordingToOutputFileAtURL:fromConnections:error:"), objc.Ptr(output), objc.Ptr(fileURL), connections, objc.Ptr(error))
}

func (c_ CaptureFileOutputRecordingDelegateObject) HasCaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections() bool {
	return c_.RespondsToSelector(objc.Sel("captureOutput:didPauseRecordingToOutputFileAtURL:fromConnections:"))
}

// Called whenever the output is recording to a file and successfully pauses the recording at the request of a client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1388838-captureoutput?language=objc
func (c_ CaptureFileOutputRecordingDelegateObject) CaptureOutputDidPauseRecordingToOutputFileAtURLFromConnections(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection) {
	objc.Call[objc.Void](c_, objc.Sel("captureOutput:didPauseRecordingToOutputFileAtURL:fromConnections:"), objc.Ptr(output), objc.Ptr(fileURL), connections)
}

func (c_ CaptureFileOutputRecordingDelegateObject) HasCaptureOutputDidStartRecordingToOutputFileAtURLFromConnections() bool {
	return c_.RespondsToSelector(objc.Sel("captureOutput:didStartRecordingToOutputFileAtURL:fromConnections:"))
}

// Informs the delegate when the output has started writing to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1387301-captureoutput?language=objc
func (c_ CaptureFileOutputRecordingDelegateObject) CaptureOutputDidStartRecordingToOutputFileAtURLFromConnections(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection) {
	objc.Call[objc.Void](c_, objc.Sel("captureOutput:didStartRecordingToOutputFileAtURL:fromConnections:"), objc.Ptr(output), objc.Ptr(fileURL), connections)
}

func (c_ CaptureFileOutputRecordingDelegateObject) HasCaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections() bool {
	return c_.RespondsToSelector(objc.Sel("captureOutput:didResumeRecordingToOutputFileAtURL:fromConnections:"))
}

// Called whenever the output, at the request of the client, successfully resumes a file recording that was paused. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1387653-captureoutput?language=objc
func (c_ CaptureFileOutputRecordingDelegateObject) CaptureOutputDidResumeRecordingToOutputFileAtURLFromConnections(output CaptureFileOutput, fileURL foundation.URL, connections []CaptureConnection) {
	objc.Call[objc.Void](c_, objc.Sel("captureOutput:didResumeRecordingToOutputFileAtURL:fromConnections:"), objc.Ptr(output), objc.Ptr(fileURL), connections)
}

func (c_ CaptureFileOutputRecordingDelegateObject) HasCaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError() bool {
	return c_.RespondsToSelector(objc.Sel("captureOutput:didFinishRecordingToOutputFileAtURL:fromConnections:error:"))
}

// Informs the delegate when all pending data has been written to an output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputrecordingdelegate/1390612-captureoutput?language=objc
func (c_ CaptureFileOutputRecordingDelegateObject) CaptureOutputDidFinishRecordingToOutputFileAtURLFromConnectionsError(output CaptureFileOutput, outputFileURL foundation.URL, connections []CaptureConnection, error foundation.Error) {
	objc.Call[objc.Void](c_, objc.Sel("captureOutput:didFinishRecordingToOutputFileAtURL:fromConnections:error:"), objc.Ptr(output), objc.Ptr(outputFileURL), connections, objc.Ptr(error))
}
