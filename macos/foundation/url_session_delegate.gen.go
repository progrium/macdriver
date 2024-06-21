// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"github.com/progrium/darwinkit/objc"
)

// A protocol that defines methods that URL session instances call on their delegates to handle session-level events, like session life cycle changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondelegate?language=objc
type PURLSessionDelegate interface {
	// optional
	URLSessionDidFinishEventsForBackgroundURLSession(session URLSession)
	HasURLSessionDidFinishEventsForBackgroundURLSession() bool

	// optional
	URLSessionDidBecomeInvalidWithError(session URLSession, error Error)
	HasURLSessionDidBecomeInvalidWithError() bool
}

// A delegate implementation builder for the [PURLSessionDelegate] protocol.
type URLSessionDelegate struct {
	_URLSessionDidFinishEventsForBackgroundURLSession func(session URLSession)
	_URLSessionDidBecomeInvalidWithError              func(session URLSession, error Error)
}

func (di *URLSessionDelegate) HasURLSessionDidFinishEventsForBackgroundURLSession() bool {
	return di._URLSessionDidFinishEventsForBackgroundURLSession != nil
}

// Tells the delegate that all messages enqueued for a session have been delivered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondelegate/1617185-urlsessiondidfinisheventsforback?language=objc
func (di *URLSessionDelegate) SetURLSessionDidFinishEventsForBackgroundURLSession(f func(session URLSession)) {
	di._URLSessionDidFinishEventsForBackgroundURLSession = f
}

// Tells the delegate that all messages enqueued for a session have been delivered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondelegate/1617185-urlsessiondidfinisheventsforback?language=objc
func (di *URLSessionDelegate) URLSessionDidFinishEventsForBackgroundURLSession(session URLSession) {
	di._URLSessionDidFinishEventsForBackgroundURLSession(session)
}
func (di *URLSessionDelegate) HasURLSessionDidBecomeInvalidWithError() bool {
	return di._URLSessionDidBecomeInvalidWithError != nil
}

// Tells the URL session that the session has been invalidated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondelegate/1407776-urlsession?language=objc
func (di *URLSessionDelegate) SetURLSessionDidBecomeInvalidWithError(f func(session URLSession, error Error)) {
	di._URLSessionDidBecomeInvalidWithError = f
}

// Tells the URL session that the session has been invalidated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondelegate/1407776-urlsession?language=objc
func (di *URLSessionDelegate) URLSessionDidBecomeInvalidWithError(session URLSession, error Error) {
	di._URLSessionDidBecomeInvalidWithError(session, error)
}

// ensure impl type implements protocol interface
var _ PURLSessionDelegate = (*URLSessionDelegateObject)(nil)

// A concrete type for the [PURLSessionDelegate] protocol.
type URLSessionDelegateObject struct {
	objc.Object
}

func (u_ URLSessionDelegateObject) HasURLSessionDidFinishEventsForBackgroundURLSession() bool {
	return u_.RespondsToSelector(objc.Sel("URLSessionDidFinishEventsForBackgroundURLSession:"))
}

// Tells the delegate that all messages enqueued for a session have been delivered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondelegate/1617185-urlsessiondidfinisheventsforback?language=objc
func (u_ URLSessionDelegateObject) URLSessionDidFinishEventsForBackgroundURLSession(session URLSession) {
	objc.Call[objc.Void](u_, objc.Sel("URLSessionDidFinishEventsForBackgroundURLSession:"), session)
}

func (u_ URLSessionDelegateObject) HasURLSessionDidBecomeInvalidWithError() bool {
	return u_.RespondsToSelector(objc.Sel("URLSession:didBecomeInvalidWithError:"))
}

// Tells the URL session that the session has been invalidated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiondelegate/1407776-urlsession?language=objc
func (u_ URLSessionDelegateObject) URLSessionDidBecomeInvalidWithError(session URLSession, error Error) {
	objc.Call[objc.Void](u_, objc.Sel("URLSession:didBecomeInvalidWithError:"), session, error)
}
