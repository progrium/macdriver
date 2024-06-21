// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"github.com/progrium/darwinkit/objc"
)

// A protocol that defines methods that URL session instances call on their delegates to handle task-level events specific to WebSocket tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketdelegate?language=objc
type PURLSessionWebSocketDelegate interface {
	// optional
	URLSessionWebSocketTaskDidOpenWithProtocol(session URLSession, webSocketTask URLSessionWebSocketTask, protocol string)
	HasURLSessionWebSocketTaskDidOpenWithProtocol() bool
}

// A delegate implementation builder for the [PURLSessionWebSocketDelegate] protocol.
type URLSessionWebSocketDelegate struct {
	_URLSessionWebSocketTaskDidOpenWithProtocol func(session URLSession, webSocketTask URLSessionWebSocketTask, protocol string)
}

func (di *URLSessionWebSocketDelegate) HasURLSessionWebSocketTaskDidOpenWithProtocol() bool {
	return di._URLSessionWebSocketTaskDidOpenWithProtocol != nil
}

// Tells the delegate that the WebSocket task successfully negotiated the handshake with the endpoint, indicating the negotiated protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketdelegate/3181189-urlsession?language=objc
func (di *URLSessionWebSocketDelegate) SetURLSessionWebSocketTaskDidOpenWithProtocol(f func(session URLSession, webSocketTask URLSessionWebSocketTask, protocol string)) {
	di._URLSessionWebSocketTaskDidOpenWithProtocol = f
}

// Tells the delegate that the WebSocket task successfully negotiated the handshake with the endpoint, indicating the negotiated protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketdelegate/3181189-urlsession?language=objc
func (di *URLSessionWebSocketDelegate) URLSessionWebSocketTaskDidOpenWithProtocol(session URLSession, webSocketTask URLSessionWebSocketTask, protocol string) {
	di._URLSessionWebSocketTaskDidOpenWithProtocol(session, webSocketTask, protocol)
}

// ensure impl type implements protocol interface
var _ PURLSessionWebSocketDelegate = (*URLSessionWebSocketDelegateObject)(nil)

// A concrete type for the [PURLSessionWebSocketDelegate] protocol.
type URLSessionWebSocketDelegateObject struct {
	objc.Object
}

func (u_ URLSessionWebSocketDelegateObject) HasURLSessionWebSocketTaskDidOpenWithProtocol() bool {
	return u_.RespondsToSelector(objc.Sel("URLSession:webSocketTask:didOpenWithProtocol:"))
}

// Tells the delegate that the WebSocket task successfully negotiated the handshake with the endpoint, indicating the negotiated protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketdelegate/3181189-urlsession?language=objc
func (u_ URLSessionWebSocketDelegateObject) URLSessionWebSocketTaskDidOpenWithProtocol(session URLSession, webSocketTask URLSessionWebSocketTask, protocol string) {
	objc.Call[objc.Void](u_, objc.Sel("URLSession:webSocketTask:didOpenWithProtocol:"), session, webSocketTask, protocol)
}
