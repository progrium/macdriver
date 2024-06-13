// Code generated by DarwinKit. DO NOT EDIT.

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// The methods to adopt in an object that monitors changes to a webpage’s cookies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestoreobserver?language=objc
type PHTTPCookieStoreObserver interface {
	// optional
	CookiesDidChangeInCookieStore(cookieStore HTTPCookieStore)
	HasCookiesDidChangeInCookieStore() bool
}

// ensure impl type implements protocol interface
var _ PHTTPCookieStoreObserver = (*HTTPCookieStoreObserverObject)(nil)

// A concrete type for the [PHTTPCookieStoreObserver] protocol.
type HTTPCookieStoreObserverObject struct {
	objc.Object
}

func (h_ HTTPCookieStoreObserverObject) HasCookiesDidChangeInCookieStore() bool {
	return h_.RespondsToSelector(objc.Sel("cookiesDidChangeInCookieStore:"))
}

// Tells the delegate that the cookies in the specified cookie store changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkhttpcookiestoreobserver/2882008-cookiesdidchangeincookiestore?language=objc
func (h_ HTTPCookieStoreObserverObject) CookiesDidChangeInCookieStore(cookieStore HTTPCookieStore) {
	objc.Call[objc.Void](h_, objc.Sel("cookiesDidChangeInCookieStore:"), objc.Ptr(cookieStore))
}