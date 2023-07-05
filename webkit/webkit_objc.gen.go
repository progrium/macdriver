package webkit

import (
	cocoa "github.com/progrium/macdriver/cocoa"
	core "github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework WebKit
#define __OBJC2__ 1
#include <objc/message.h>
#include <stdlib.h>

#include <WebKit/WebKit.h>

bool webkit_convertObjCBool(BOOL b) {
	if (b) { return true; }
	return false;
}


void* WKNavigation_type_alloc() {
	return [WKNavigation
		alloc];
}
void* WKUserScript_type_alloc() {
	return [WKUserScript
		alloc];
}
void* WKWebView_type_alloc() {
	return [WKWebView
		alloc];
}
BOOL WKWebView_type_handlesURLScheme(void* urlScheme) {
	return [WKWebView
		handlesURLScheme: urlScheme];
}
void* WKWebViewConfiguration_type_alloc() {
	return [WKWebViewConfiguration
		alloc];
}
void* WKPreferences_type_alloc() {
	return [WKPreferences
		alloc];
}


void* WKNavigation_inst_init(void *id) {
	return [(WKNavigation*)id
		init];
}

void* WKUserScript_inst_init(void *id) {
	return [(WKUserScript*)id
		init];
}

void* WKUserScript_inst_source(void *id) {
	return [(WKUserScript*)id
		source];
}

BOOL WKUserScript_inst_isForMainFrameOnly(void *id) {
	return [(WKUserScript*)id
		isForMainFrameOnly];
}

void* WKWebView_inst_goBack(void *id) {
	return [(WKWebView*)id
		goBack];
}

void WKWebView_inst_goBack_(void *id, void* sender) {
	[(WKWebView*)id
		goBack: sender];
}

void* WKWebView_inst_goForward(void *id) {
	return [(WKWebView*)id
		goForward];
}

void WKWebView_inst_goForward_(void *id, void* sender) {
	[(WKWebView*)id
		goForward: sender];
}

void* WKWebView_inst_initWithFrame_configuration(void *id, NSRect frame, void* configuration) {
	return [(WKWebView*)id
		initWithFrame: frame
		configuration: configuration];
}

void* WKWebView_inst_loadData_MIMEType_characterEncodingName_baseURL(void *id, void* data, void* MIMEType, void* characterEncodingName, void* baseURL) {
	return [(WKWebView*)id
		loadData: data
		MIMEType: MIMEType
		characterEncodingName: characterEncodingName
		baseURL: baseURL];
}

void* WKWebView_inst_loadFileRequest_allowingReadAccessToURL(void *id, void* request, void* readAccessURL) {
	return [(WKWebView*)id
		loadFileRequest: request
		allowingReadAccessToURL: readAccessURL];
}

void* WKWebView_inst_loadFileURL_allowingReadAccessToURL(void *id, void* URL, void* readAccessURL) {
	return [(WKWebView*)id
		loadFileURL: URL
		allowingReadAccessToURL: readAccessURL];
}

void* WKWebView_inst_loadHTMLString_baseURL(void *id, void* string, void* baseURL) {
	return [(WKWebView*)id
		loadHTMLString: string
		baseURL: baseURL];
}

void* WKWebView_inst_loadRequest(void *id, void* request) {
	return [(WKWebView*)id
		loadRequest: request];
}

void* WKWebView_inst_loadSimulatedRequest_responseHTMLString(void *id, void* request, void* string) {
	return [(WKWebView*)id
		loadSimulatedRequest: request
		responseHTMLString: string];
}

void* WKWebView_inst_reload(void *id) {
	return [(WKWebView*)id
		reload];
}

void WKWebView_inst_reload_(void *id, void* sender) {
	[(WKWebView*)id
		reload: sender];
}

void* WKWebView_inst_reloadFromOrigin(void *id) {
	return [(WKWebView*)id
		reloadFromOrigin];
}

void WKWebView_inst_reloadFromOrigin_(void *id, void* sender) {
	[(WKWebView*)id
		reloadFromOrigin: sender];
}

void WKWebView_inst_stopLoading(void *id) {
	[(WKWebView*)id
		stopLoading];
}

void WKWebView_inst_stopLoading_(void *id, void* sender) {
	[(WKWebView*)id
		stopLoading: sender];
}

void* WKWebView_inst_init(void *id) {
	return [(WKWebView*)id
		init];
}

void* WKWebView_inst_configuration(void *id) {
	return [(WKWebView*)id
		configuration];
}

void* WKWebView_inst_UIDelegate(void *id) {
	return [(WKWebView*)id
		UIDelegate];
}

void WKWebView_inst_setUIDelegate(void *id, void* value) {
	[(WKWebView*)id
		setUIDelegate: value];
}

void* WKWebView_inst_navigationDelegate(void *id) {
	return [(WKWebView*)id
		navigationDelegate];
}

void WKWebView_inst_setNavigationDelegate(void *id, void* value) {
	[(WKWebView*)id
		setNavigationDelegate: value];
}

BOOL WKWebView_inst_isLoading(void *id) {
	return [(WKWebView*)id
		isLoading];
}

void* WKWebView_inst_title(void *id) {
	return [(WKWebView*)id
		title];
}

void* WKWebView_inst_URL(void *id) {
	return [(WKWebView*)id
		URL];
}

void* WKWebView_inst_mediaType(void *id) {
	return [(WKWebView*)id
		mediaType];
}

void WKWebView_inst_setMediaType(void *id, void* value) {
	[(WKWebView*)id
		setMediaType: value];
}

void* WKWebView_inst_customUserAgent(void *id) {
	return [(WKWebView*)id
		customUserAgent];
}

void WKWebView_inst_setCustomUserAgent(void *id, void* value) {
	[(WKWebView*)id
		setCustomUserAgent: value];
}

BOOL WKWebView_inst_hasOnlySecureContent(void *id) {
	return [(WKWebView*)id
		hasOnlySecureContent];
}

BOOL WKWebView_inst_allowsMagnification(void *id) {
	return [(WKWebView*)id
		allowsMagnification];
}

void WKWebView_inst_setAllowsMagnification(void *id, BOOL value) {
	[(WKWebView*)id
		setAllowsMagnification: value];
}

double WKWebView_inst_magnification(void *id) {
	return [(WKWebView*)id
		magnification];
}

void WKWebView_inst_setMagnification(void *id, double value) {
	[(WKWebView*)id
		setMagnification: value];
}

BOOL WKWebView_inst_allowsBackForwardNavigationGestures(void *id) {
	return [(WKWebView*)id
		allowsBackForwardNavigationGestures];
}

void WKWebView_inst_setAllowsBackForwardNavigationGestures(void *id, BOOL value) {
	[(WKWebView*)id
		setAllowsBackForwardNavigationGestures: value];
}

BOOL WKWebView_inst_canGoBack(void *id) {
	return [(WKWebView*)id
		canGoBack];
}

BOOL WKWebView_inst_canGoForward(void *id) {
	return [(WKWebView*)id
		canGoForward];
}

BOOL WKWebView_inst_allowsLinkPreview(void *id) {
	return [(WKWebView*)id
		allowsLinkPreview];
}

void WKWebView_inst_setAllowsLinkPreview(void *id, BOOL value) {
	[(WKWebView*)id
		setAllowsLinkPreview: value];
}

void* WKWebView_inst_interactionState(void *id) {
	return [(WKWebView*)id
		interactionState];
}

void WKWebView_inst_setInteractionState(void *id, void* value) {
	[(WKWebView*)id
		setInteractionState: value];
}

void WKWebViewConfiguration_inst_setURLSchemeHandler_forURLScheme(void *id, void* urlSchemeHandler, void* urlScheme) {
	[(WKWebViewConfiguration*)id
		setURLSchemeHandler: urlSchemeHandler
		forURLScheme: urlScheme];
}

void* WKWebViewConfiguration_inst_urlSchemeHandlerForURLScheme(void *id, void* urlScheme) {
	return [(WKWebViewConfiguration*)id
		urlSchemeHandlerForURLScheme: urlScheme];
}

void* WKWebViewConfiguration_inst_init(void *id) {
	return [(WKWebViewConfiguration*)id
		init];
}

void* WKWebViewConfiguration_inst_applicationNameForUserAgent(void *id) {
	return [(WKWebViewConfiguration*)id
		applicationNameForUserAgent];
}

void WKWebViewConfiguration_inst_setApplicationNameForUserAgent(void *id, void* value) {
	[(WKWebViewConfiguration*)id
		setApplicationNameForUserAgent: value];
}

BOOL WKWebViewConfiguration_inst_limitsNavigationsToAppBoundDomains(void *id) {
	return [(WKWebViewConfiguration*)id
		limitsNavigationsToAppBoundDomains];
}

void WKWebViewConfiguration_inst_setLimitsNavigationsToAppBoundDomains(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setLimitsNavigationsToAppBoundDomains: value];
}

void* WKWebViewConfiguration_inst_preferences(void *id) {
	return [(WKWebViewConfiguration*)id
		preferences];
}

void WKWebViewConfiguration_inst_setPreferences(void *id, void* value) {
	[(WKWebViewConfiguration*)id
		setPreferences: value];
}

BOOL WKWebViewConfiguration_inst_ignoresViewportScaleLimits(void *id) {
	return [(WKWebViewConfiguration*)id
		ignoresViewportScaleLimits];
}

void WKWebViewConfiguration_inst_setIgnoresViewportScaleLimits(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setIgnoresViewportScaleLimits: value];
}

BOOL WKWebViewConfiguration_inst_suppressesIncrementalRendering(void *id) {
	return [(WKWebViewConfiguration*)id
		suppressesIncrementalRendering];
}

void WKWebViewConfiguration_inst_setSuppressesIncrementalRendering(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setSuppressesIncrementalRendering: value];
}

BOOL WKWebViewConfiguration_inst_allowsInlineMediaPlayback(void *id) {
	return [(WKWebViewConfiguration*)id
		allowsInlineMediaPlayback];
}

void WKWebViewConfiguration_inst_setAllowsInlineMediaPlayback(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setAllowsInlineMediaPlayback: value];
}

BOOL WKWebViewConfiguration_inst_allowsAirPlayForMediaPlayback(void *id) {
	return [(WKWebViewConfiguration*)id
		allowsAirPlayForMediaPlayback];
}

void WKWebViewConfiguration_inst_setAllowsAirPlayForMediaPlayback(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setAllowsAirPlayForMediaPlayback: value];
}

BOOL WKWebViewConfiguration_inst_allowsPictureInPictureMediaPlayback(void *id) {
	return [(WKWebViewConfiguration*)id
		allowsPictureInPictureMediaPlayback];
}

void WKWebViewConfiguration_inst_setAllowsPictureInPictureMediaPlayback(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setAllowsPictureInPictureMediaPlayback: value];
}

BOOL WKWebViewConfiguration_inst_upgradeKnownHostsToHTTPS(void *id) {
	return [(WKWebViewConfiguration*)id
		upgradeKnownHostsToHTTPS];
}

void WKWebViewConfiguration_inst_setUpgradeKnownHostsToHTTPS(void *id, BOOL value) {
	[(WKWebViewConfiguration*)id
		setUpgradeKnownHostsToHTTPS: value];
}

void WKPreferences_inst_setValue_forKey(void *id, void* value, void* key) {
	[(WKPreferences*)id
		setValue: value
		forKey: key];
}

void* WKPreferences_inst_init(void *id) {
	return [(WKPreferences*)id
		init];
}

double WKPreferences_inst_minimumFontSize(void *id) {
	return [(WKPreferences*)id
		minimumFontSize];
}

void WKPreferences_inst_setMinimumFontSize(void *id, double value) {
	[(WKPreferences*)id
		setMinimumFontSize: value];
}

BOOL WKPreferences_inst_tabFocusesLinks(void *id) {
	return [(WKPreferences*)id
		tabFocusesLinks];
}

void WKPreferences_inst_setTabFocusesLinks(void *id, BOOL value) {
	[(WKPreferences*)id
		setTabFocusesLinks: value];
}

BOOL WKPreferences_inst_javaScriptCanOpenWindowsAutomatically(void *id) {
	return [(WKPreferences*)id
		javaScriptCanOpenWindowsAutomatically];
}

void WKPreferences_inst_setJavaScriptCanOpenWindowsAutomatically(void *id, BOOL value) {
	[(WKPreferences*)id
		setJavaScriptCanOpenWindowsAutomatically: value];
}

BOOL WKPreferences_inst_isFraudulentWebsiteWarningEnabled(void *id) {
	return [(WKPreferences*)id
		isFraudulentWebsiteWarningEnabled];
}

void WKPreferences_inst_setFraudulentWebsiteWarningEnabled(void *id, BOOL value) {
	[(WKPreferences*)id
		setFraudulentWebsiteWarningEnabled: value];
}

BOOL WKPreferences_inst_isTextInteractionEnabled(void *id) {
	return [(WKPreferences*)id
		isTextInteractionEnabled];
}

void WKPreferences_inst_setTextInteractionEnabled(void *id, BOOL value) {
	[(WKPreferences*)id
		setTextInteractionEnabled: value];
}


BOOL webkit_objc_bool_true = YES;
BOOL webkit_objc_bool_false = NO;

*/
import "C"

func convertObjCBoolToGo(b C.BOOL) bool {
	// NOTE: the prefix here is used to namespace these since the linker will
	// otherwise report a "duplicate symbol" because the C functions have the
	// same name.
	return bool(C.webkit_convertObjCBool(b))
}

func convertToObjCBool(b bool) C.BOOL {
	if b {
		return C.webkit_objc_bool_true
	}
	return C.webkit_objc_bool_false
}

// WKNavigation_alloc
//
// See  for details.
func WKNavigation_alloc() WKNavigation {
	ret := C.WKNavigation_type_alloc()

	return WKNavigation_fromPointer(ret)

}

// WKUserScript_alloc
//
// See  for details.
func WKUserScript_alloc() WKUserScript {
	ret := C.WKUserScript_type_alloc()

	return WKUserScript_fromPointer(ret)

}

// WKWebView_alloc
//
// See  for details.
func WKWebView_alloc() WKWebView {
	ret := C.WKWebView_type_alloc()

	return WKWebView_fromPointer(ret)

}

// WKWebView_handlesURLScheme returns a boolean value that indicates whether webkit natively supports resources with the specified url scheme.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/2875370-handlesurlscheme?language=objc for details.
func WKWebView_handlesURLScheme(urlScheme core.NSStringRef) bool {
	ret := C.WKWebView_type_handlesURLScheme(
		objc.RefPointer(urlScheme),
	)

	return convertObjCBoolToGo(ret)

}

// WKWebViewConfiguration_alloc
//
// See  for details.
func WKWebViewConfiguration_alloc() WKWebViewConfiguration {
	ret := C.WKWebViewConfiguration_type_alloc()

	return WKWebViewConfiguration_fromPointer(ret)

}

// WKPreferences_alloc
//
// See  for details.
func WKPreferences_alloc() WKPreferences {
	ret := C.WKPreferences_type_alloc()

	return WKPreferences_fromPointer(ret)

}

type WKNavigationRef interface {
	Pointer() uintptr
	Init_asWKNavigation() WKNavigation
}

type gen_WKNavigation struct {
	objc.Object
}

func WKNavigation_fromPointer(ptr unsafe.Pointer) WKNavigation {
	return WKNavigation{gen_WKNavigation{
		objc.Object_fromPointer(ptr),
	}}
}

func WKNavigation_fromRef(ref objc.Ref) WKNavigation {
	return WKNavigation_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init
//
// See  for details.
func (x gen_WKNavigation) Init_asWKNavigation() WKNavigation {
	ret := C.WKNavigation_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKNavigation_fromPointer(ret)

}

type WKUserScriptRef interface {
	Pointer() uintptr
	Init_asWKUserScript() WKUserScript
}

type gen_WKUserScript struct {
	objc.Object
}

func WKUserScript_fromPointer(ptr unsafe.Pointer) WKUserScript {
	return WKUserScript{gen_WKUserScript{
		objc.Object_fromPointer(ptr),
	}}
}

func WKUserScript_fromRef(ref objc.Ref) WKUserScript {
	return WKUserScript_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// Init
//
// See  for details.
func (x gen_WKUserScript) Init_asWKUserScript() WKUserScript {
	ret := C.WKUserScript_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKUserScript_fromPointer(ret)

}

// Source returns the script’s source code.
//
// See https://developer.apple.com/documentation/webkit/wkuserscript/1537787-source?language=objc for details.
func (x gen_WKUserScript) Source() core.NSString {
	ret := C.WKUserScript_inst_source(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// IsForMainFrameOnly returns a boolean value that indicates whether to inject the script into the main frame or all frames.
//
// See https://developer.apple.com/documentation/webkit/wkuserscript/1537856-formainframeonly?language=objc for details.
func (x gen_WKUserScript) IsForMainFrameOnly() bool {
	ret := C.WKUserScript_inst_isForMainFrameOnly(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

type WKWebViewRef interface {
	Pointer() uintptr
	Init_asWKWebView() WKWebView
}

type gen_WKWebView struct {
	cocoa.NSView
}

func WKWebView_fromPointer(ptr unsafe.Pointer) WKWebView {
	return WKWebView{gen_WKWebView{
		cocoa.NSView_fromPointer(ptr),
	}}
}

func WKWebView_fromRef(ref objc.Ref) WKWebView {
	return WKWebView_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// GoBack navigates to the back item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414952-goback?language=objc for details.
func (x gen_WKWebView) GoBack() WKNavigation {
	ret := C.WKWebView_inst_goBack(
		unsafe.Pointer(x.Pointer()),
	)

	return WKNavigation_fromPointer(ret)

}

// GoBack_ navigates to the back item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414975-goback?language=objc for details.
func (x gen_WKWebView) GoBack_(
	sender objc.Ref,
) {
	C.WKWebView_inst_goBack_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// GoForward navigates to the forward item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414993-goforward?language=objc for details.
func (x gen_WKWebView) GoForward() WKNavigation {
	ret := C.WKWebView_inst_goForward(
		unsafe.Pointer(x.Pointer()),
	)

	return WKNavigation_fromPointer(ret)

}

// GoForward_ navigates to the forward item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414960-goforward?language=objc for details.
func (x gen_WKWebView) GoForward_(
	sender objc.Ref,
) {
	C.WKWebView_inst_goForward_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// InitWithFrame_configuration creates a web view and initializes it with the specified frame and configuration data.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414998-initwithframe?language=objc for details.
func (x gen_WKWebView) InitWithFrame_configuration_asWKWebView(
	frame core.NSRect,
	configuration WKWebViewConfigurationRef,
) WKWebView {
	ret := C.WKWebView_inst_initWithFrame_configuration(
		unsafe.Pointer(x.Pointer()),
		*(*C.NSRect)(unsafe.Pointer(&frame)),
		objc.RefPointer(configuration),
	)

	return WKWebView_fromPointer(ret)

}

// LoadData_MIMEType_characterEncodingName_baseURL loads the content of the specified data object and navigates to it.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415011-loaddata?language=objc for details.
func (x gen_WKWebView) LoadData_MIMEType_characterEncodingName_baseURL(
	data core.NSDataRef,
	MIMEType core.NSStringRef,
	characterEncodingName core.NSStringRef,
	baseURL core.NSURLRef,
) WKNavigation {
	ret := C.WKWebView_inst_loadData_MIMEType_characterEncodingName_baseURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(data),
		objc.RefPointer(MIMEType),
		objc.RefPointer(characterEncodingName),
		objc.RefPointer(baseURL),
	)

	return WKNavigation_fromPointer(ret)

}

// LoadFileRequest_allowingReadAccessToURL
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3752237-loadfilerequest?language=objc for details.
func (x gen_WKWebView) LoadFileRequest_allowingReadAccessToURL(
	request core.NSURLRequestRef,
	readAccessURL core.NSURLRef,
) WKNavigation {
	ret := C.WKWebView_inst_loadFileRequest_allowingReadAccessToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(request),
		objc.RefPointer(readAccessURL),
	)

	return WKNavigation_fromPointer(ret)

}

// LoadFileURL_allowingReadAccessToURL loads the web content from the specified file and navigates to it.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414973-loadfileurl?language=objc for details.
func (x gen_WKWebView) LoadFileURL_allowingReadAccessToURL(
	URL core.NSURLRef,
	readAccessURL core.NSURLRef,
) WKNavigation {
	ret := C.WKWebView_inst_loadFileURL_allowingReadAccessToURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(URL),
		objc.RefPointer(readAccessURL),
	)

	return WKNavigation_fromPointer(ret)

}

// LoadHTMLString_baseURL loads the contents of the specified html string and navigates to it.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415004-loadhtmlstring?language=objc for details.
func (x gen_WKWebView) LoadHTMLString_baseURL(
	string core.NSStringRef,
	baseURL core.NSURLRef,
) WKNavigation {
	ret := C.WKWebView_inst_loadHTMLString_baseURL(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(string),
		objc.RefPointer(baseURL),
	)

	return WKNavigation_fromPointer(ret)

}

// LoadRequest loads the web content referenced by the specified url request object and navigates to it.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414954-loadrequest?language=objc for details.
func (x gen_WKWebView) LoadRequest(
	request core.NSURLRequestRef,
) WKNavigation {
	ret := C.WKWebView_inst_loadRequest(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(request),
	)

	return WKNavigation_fromPointer(ret)

}

// LoadSimulatedRequest_responseHTMLString
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3763095-loadsimulatedrequest?language=objc for details.
func (x gen_WKWebView) LoadSimulatedRequest_responseHTMLString(
	request core.NSURLRequestRef,
	string core.NSStringRef,
) WKNavigation {
	ret := C.WKWebView_inst_loadSimulatedRequest_responseHTMLString(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(request),
		objc.RefPointer(string),
	)

	return WKNavigation_fromPointer(ret)

}

// Reload reloads the current webpage.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414969-reload?language=objc for details.
func (x gen_WKWebView) Reload() WKNavigation {
	ret := C.WKWebView_inst_reload(
		unsafe.Pointer(x.Pointer()),
	)

	return WKNavigation_fromPointer(ret)

}

// Reload_ reloads the current webpage.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414987-reload?language=objc for details.
func (x gen_WKWebView) Reload_(
	sender objc.Ref,
) {
	C.WKWebView_inst_reload_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// ReloadFromOrigin reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414956-reloadfromorigin?language=objc for details.
func (x gen_WKWebView) ReloadFromOrigin() WKNavigation {
	ret := C.WKWebView_inst_reloadFromOrigin(
		unsafe.Pointer(x.Pointer()),
	)

	return WKNavigation_fromPointer(ret)

}

// ReloadFromOrigin_ reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414989-reloadfromorigin?language=objc for details.
func (x gen_WKWebView) ReloadFromOrigin_(
	sender objc.Ref,
) {
	C.WKWebView_inst_reloadFromOrigin_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// StopLoading stops loading all resources on the current page.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414981-stoploading?language=objc for details.
func (x gen_WKWebView) StopLoading() {
	C.WKWebView_inst_stopLoading(
		unsafe.Pointer(x.Pointer()),
	)

	return

}

// StopLoading_ stops loading all resources on the current page.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415013-stoploading?language=objc for details.
func (x gen_WKWebView) StopLoading_(
	sender objc.Ref,
) {
	C.WKWebView_inst_stopLoading_(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(sender),
	)

	return

}

// Init
//
// See  for details.
func (x gen_WKWebView) Init_asWKWebView() WKWebView {
	ret := C.WKWebView_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKWebView_fromPointer(ret)

}

// Configuration returns the object that contains the configuration details for the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414979-configuration?language=objc for details.
func (x gen_WKWebView) Configuration() WKWebViewConfiguration {
	ret := C.WKWebView_inst_configuration(
		unsafe.Pointer(x.Pointer()),
	)

	return WKWebViewConfiguration_fromPointer(ret)

}

// UIDelegate returns the object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415009-uidelegate?language=objc for details.
func (x gen_WKWebView) UIDelegate() objc.Object {
	ret := C.WKWebView_inst_UIDelegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetUIDelegate returns the object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415009-uidelegate?language=objc for details.
func (x gen_WKWebView) SetUIDelegate(
	value objc.Ref,
) {
	C.WKWebView_inst_setUIDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// NavigationDelegate returns the object you use to manage navigation behavior for the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414971-navigationdelegate?language=objc for details.
func (x gen_WKWebView) NavigationDelegate() objc.Object {
	ret := C.WKWebView_inst_navigationDelegate(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetNavigationDelegate returns the object you use to manage navigation behavior for the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414971-navigationdelegate?language=objc for details.
func (x gen_WKWebView) SetNavigationDelegate(
	value objc.Ref,
) {
	C.WKWebView_inst_setNavigationDelegate(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IsLoading returns a boolean value that indicates whether the view is currently loading content.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414964-loading?language=objc for details.
func (x gen_WKWebView) IsLoading() bool {
	ret := C.WKWebView_inst_isLoading(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// Title returns the page title.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415015-title?language=objc for details.
func (x gen_WKWebView) Title() core.NSString {
	ret := C.WKWebView_inst_title(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// URL returns the url for the current webpage.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415005-url?language=objc for details.
func (x gen_WKWebView) URL() core.NSURL {
	ret := C.WKWebView_inst_URL(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSURL_fromPointer(ret)

}

// MediaType returns the media type for the contents of the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3516410-mediatype?language=objc for details.
func (x gen_WKWebView) MediaType() core.NSString {
	ret := C.WKWebView_inst_mediaType(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetMediaType returns the media type for the contents of the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3516410-mediatype?language=objc for details.
func (x gen_WKWebView) SetMediaType(
	value core.NSStringRef,
) {
	C.WKWebView_inst_setMediaType(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// CustomUserAgent returns the custom user agent string.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414950-customuseragent?language=objc for details.
func (x gen_WKWebView) CustomUserAgent() core.NSString {
	ret := C.WKWebView_inst_customUserAgent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetCustomUserAgent returns the custom user agent string.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414950-customuseragent?language=objc for details.
func (x gen_WKWebView) SetCustomUserAgent(
	value core.NSStringRef,
) {
	C.WKWebView_inst_setCustomUserAgent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// HasOnlySecureContent returns a boolean value that indicates whether the web view loaded all resources on the page through securely encrypted connections.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415002-hasonlysecurecontent?language=objc for details.
func (x gen_WKWebView) HasOnlySecureContent() bool {
	ret := C.WKWebView_inst_hasOnlySecureContent(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AllowsMagnification returns a boolean value that indicates whether magnify gestures change the web view’s magnification.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414983-allowsmagnification?language=objc for details.
func (x gen_WKWebView) AllowsMagnification() bool {
	ret := C.WKWebView_inst_allowsMagnification(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsMagnification returns a boolean value that indicates whether magnify gestures change the web view’s magnification.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414983-allowsmagnification?language=objc for details.
func (x gen_WKWebView) SetAllowsMagnification(
	value bool,
) {
	C.WKWebView_inst_setAllowsMagnification(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Magnification returns the factor by which the page content is currently scaled.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414985-magnification?language=objc for details.
func (x gen_WKWebView) Magnification() core.CGFloat {
	ret := C.WKWebView_inst_magnification(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SetMagnification returns the factor by which the page content is currently scaled.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414985-magnification?language=objc for details.
func (x gen_WKWebView) SetMagnification(
	value core.CGFloat,
) {
	C.WKWebView_inst_setMagnification(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// AllowsBackForwardNavigationGestures returns a boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414995-allowsbackforwardnavigationgestu?language=objc for details.
func (x gen_WKWebView) AllowsBackForwardNavigationGestures() bool {
	ret := C.WKWebView_inst_allowsBackForwardNavigationGestures(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsBackForwardNavigationGestures returns a boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414995-allowsbackforwardnavigationgestu?language=objc for details.
func (x gen_WKWebView) SetAllowsBackForwardNavigationGestures(
	value bool,
) {
	C.WKWebView_inst_setAllowsBackForwardNavigationGestures(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// CanGoBack returns a boolean value that indicates whether there is a valid back item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414966-cangoback?language=objc for details.
func (x gen_WKWebView) CanGoBack() bool {
	ret := C.WKWebView_inst_canGoBack(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// CanGoForward returns a boolean value that indicates whether there is a valid forward item in the back-forward list.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1414962-cangoforward?language=objc for details.
func (x gen_WKWebView) CanGoForward() bool {
	ret := C.WKWebView_inst_canGoForward(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// AllowsLinkPreview returns a boolean value that determines whether pressing a link displays a preview of the destination for the link.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415000-allowslinkpreview?language=objc for details.
func (x gen_WKWebView) AllowsLinkPreview() bool {
	ret := C.WKWebView_inst_allowsLinkPreview(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsLinkPreview returns a boolean value that determines whether pressing a link displays a preview of the destination for the link.
//
// See https://developer.apple.com/documentation/webkit/wkwebview/1415000-allowslinkpreview?language=objc for details.
func (x gen_WKWebView) SetAllowsLinkPreview(
	value bool,
) {
	C.WKWebView_inst_setAllowsLinkPreview(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// InteractionState
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3752236-interactionstate?language=objc for details.
func (x gen_WKWebView) InteractionState() objc.Object {
	ret := C.WKWebView_inst_interactionState(
		unsafe.Pointer(x.Pointer()),
	)

	return objc.Object_fromPointer(ret)

}

// SetInteractionState
//
// See https://developer.apple.com/documentation/webkit/wkwebview/3752236-interactionstate?language=objc for details.
func (x gen_WKWebView) SetInteractionState(
	value objc.Ref,
) {
	C.WKWebView_inst_setInteractionState(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

type WKWebViewConfigurationRef interface {
	Pointer() uintptr
	Init_asWKWebViewConfiguration() WKWebViewConfiguration
}

type gen_WKWebViewConfiguration struct {
	objc.Object
}

func WKWebViewConfiguration_fromPointer(ptr unsafe.Pointer) WKWebViewConfiguration {
	return WKWebViewConfiguration{gen_WKWebViewConfiguration{
		objc.Object_fromPointer(ptr),
	}}
}

func WKWebViewConfiguration_fromRef(ref objc.Ref) WKWebViewConfiguration {
	return WKWebViewConfiguration_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// SetURLSchemeHandler_forURLScheme registers an object to load resources associated with the specified url scheme.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/2875766-seturlschemehandler?language=objc for details.
func (x gen_WKWebViewConfiguration) SetURLSchemeHandler_forURLScheme(
	urlSchemeHandler objc.Ref,
	urlScheme core.NSStringRef,
) {
	C.WKWebViewConfiguration_inst_setURLSchemeHandler_forURLScheme(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(urlSchemeHandler),
		objc.RefPointer(urlScheme),
	)

	return

}

// UrlSchemeHandlerForURLScheme returns the currently registered handler object for the specified url scheme.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/2875767-urlschemehandlerforurlscheme?language=objc for details.
func (x gen_WKWebViewConfiguration) UrlSchemeHandlerForURLScheme(
	urlScheme core.NSStringRef,
) objc.Object {
	ret := C.WKWebViewConfiguration_inst_urlSchemeHandlerForURLScheme(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(urlScheme),
	)

	return objc.Object_fromPointer(ret)

}

// Init
//
// See  for details.
func (x gen_WKWebViewConfiguration) Init_asWKWebViewConfiguration() WKWebViewConfiguration {
	ret := C.WKWebViewConfiguration_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKWebViewConfiguration_fromPointer(ret)

}

// ApplicationNameForUserAgent returns the app name that appears in the user agent string.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395665-applicationnameforuseragent?language=objc for details.
func (x gen_WKWebViewConfiguration) ApplicationNameForUserAgent() core.NSString {
	ret := C.WKWebViewConfiguration_inst_applicationNameForUserAgent(
		unsafe.Pointer(x.Pointer()),
	)

	return core.NSString_fromPointer(ret)

}

// SetApplicationNameForUserAgent returns the app name that appears in the user agent string.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395665-applicationnameforuseragent?language=objc for details.
func (x gen_WKWebViewConfiguration) SetApplicationNameForUserAgent(
	value core.NSStringRef,
) {
	C.WKWebViewConfiguration_inst_setApplicationNameForUserAgent(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// LimitsNavigationsToAppBoundDomains returns a boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3585117-limitsnavigationstoappbounddomai?language=objc for details.
func (x gen_WKWebViewConfiguration) LimitsNavigationsToAppBoundDomains() bool {
	ret := C.WKWebViewConfiguration_inst_limitsNavigationsToAppBoundDomains(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetLimitsNavigationsToAppBoundDomains returns a boolean value that indicates whether the web view limits navigation to pages within the app’s domain.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3585117-limitsnavigationstoappbounddomai?language=objc for details.
func (x gen_WKWebViewConfiguration) SetLimitsNavigationsToAppBoundDomains(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setLimitsNavigationsToAppBoundDomains(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// Preferences returns the object that manages the preference-related settings for the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395666-preferences?language=objc for details.
func (x gen_WKWebViewConfiguration) Preferences() WKPreferences {
	ret := C.WKWebViewConfiguration_inst_preferences(
		unsafe.Pointer(x.Pointer()),
	)

	return WKPreferences_fromPointer(ret)

}

// SetPreferences returns the object that manages the preference-related settings for the web view.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395666-preferences?language=objc for details.
func (x gen_WKWebViewConfiguration) SetPreferences(
	value WKPreferencesRef,
) {
	C.WKWebViewConfiguration_inst_setPreferences(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
	)

	return

}

// IgnoresViewportScaleLimits returns a boolean value that determines whether a web view allows scaling of the webpage.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/2274633-ignoresviewportscalelimits?language=objc for details.
func (x gen_WKWebViewConfiguration) IgnoresViewportScaleLimits() bool {
	ret := C.WKWebViewConfiguration_inst_ignoresViewportScaleLimits(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetIgnoresViewportScaleLimits returns a boolean value that determines whether a web view allows scaling of the webpage.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/2274633-ignoresviewportscalelimits?language=objc for details.
func (x gen_WKWebViewConfiguration) SetIgnoresViewportScaleLimits(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setIgnoresViewportScaleLimits(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// SuppressesIncrementalRendering returns a boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395663-suppressesincrementalrendering?language=objc for details.
func (x gen_WKWebViewConfiguration) SuppressesIncrementalRendering() bool {
	ret := C.WKWebViewConfiguration_inst_suppressesIncrementalRendering(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetSuppressesIncrementalRendering returns a boolean value that indicates whether the web view suppresses content rendering until the content is fully loaded into memory.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395663-suppressesincrementalrendering?language=objc for details.
func (x gen_WKWebViewConfiguration) SetSuppressesIncrementalRendering(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setSuppressesIncrementalRendering(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsInlineMediaPlayback returns a boolean value that indicates whether html5 videos play inline or use the native full-screen controller.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1614793-allowsinlinemediaplayback?language=objc for details.
func (x gen_WKWebViewConfiguration) AllowsInlineMediaPlayback() bool {
	ret := C.WKWebViewConfiguration_inst_allowsInlineMediaPlayback(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsInlineMediaPlayback returns a boolean value that indicates whether html5 videos play inline or use the native full-screen controller.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1614793-allowsinlinemediaplayback?language=objc for details.
func (x gen_WKWebViewConfiguration) SetAllowsInlineMediaPlayback(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setAllowsInlineMediaPlayback(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsAirPlayForMediaPlayback returns a boolean value that indicates whether the web view allows media playback over airplay.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395673-allowsairplayformediaplayback?language=objc for details.
func (x gen_WKWebViewConfiguration) AllowsAirPlayForMediaPlayback() bool {
	ret := C.WKWebViewConfiguration_inst_allowsAirPlayForMediaPlayback(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsAirPlayForMediaPlayback returns a boolean value that indicates whether the web view allows media playback over airplay.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1395673-allowsairplayformediaplayback?language=objc for details.
func (x gen_WKWebViewConfiguration) SetAllowsAirPlayForMediaPlayback(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setAllowsAirPlayForMediaPlayback(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// AllowsPictureInPictureMediaPlayback returns a boolean value that indicates whether html5 videos can play picture in picture.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1614792-allowspictureinpicturemediaplayb?language=objc for details.
func (x gen_WKWebViewConfiguration) AllowsPictureInPictureMediaPlayback() bool {
	ret := C.WKWebViewConfiguration_inst_allowsPictureInPictureMediaPlayback(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetAllowsPictureInPictureMediaPlayback returns a boolean value that indicates whether html5 videos can play picture in picture.
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/1614792-allowspictureinpicturemediaplayb?language=objc for details.
func (x gen_WKWebViewConfiguration) SetAllowsPictureInPictureMediaPlayback(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setAllowsPictureInPictureMediaPlayback(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// UpgradeKnownHostsToHTTPS
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3752243-upgradeknownhoststohttps?language=objc for details.
func (x gen_WKWebViewConfiguration) UpgradeKnownHostsToHTTPS() bool {
	ret := C.WKWebViewConfiguration_inst_upgradeKnownHostsToHTTPS(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetUpgradeKnownHostsToHTTPS
//
// See https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/3752243-upgradeknownhoststohttps?language=objc for details.
func (x gen_WKWebViewConfiguration) SetUpgradeKnownHostsToHTTPS(
	value bool,
) {
	C.WKWebViewConfiguration_inst_setUpgradeKnownHostsToHTTPS(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

type WKPreferencesRef interface {
	Pointer() uintptr
	Init_asWKPreferences() WKPreferences
}

type gen_WKPreferences struct {
	objc.Object
}

func WKPreferences_fromPointer(ptr unsafe.Pointer) WKPreferences {
	return WKPreferences{gen_WKPreferences{
		objc.Object_fromPointer(ptr),
	}}
}

func WKPreferences_fromRef(ref objc.Ref) WKPreferences {
	return WKPreferences_fromPointer(unsafe.Pointer(ref.Pointer()))
}

// SetValue_forKey
//
// See  for details.
func (x gen_WKPreferences) SetValue_forKey(
	value objc.Ref,
	key core.NSStringRef,
) {
	C.WKPreferences_inst_setValue_forKey(
		unsafe.Pointer(x.Pointer()),
		objc.RefPointer(value),
		objc.RefPointer(key),
	)

	return

}

// Init
//
// See  for details.
func (x gen_WKPreferences) Init_asWKPreferences() WKPreferences {
	ret := C.WKPreferences_inst_init(
		unsafe.Pointer(x.Pointer()),
	)

	return WKPreferences_fromPointer(ret)

}

// MinimumFontSize returns the minimum font size, in points.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/1537155-minimumfontsize?language=objc for details.
func (x gen_WKPreferences) MinimumFontSize() core.CGFloat {
	ret := C.WKPreferences_inst_minimumFontSize(
		unsafe.Pointer(x.Pointer()),
	)

	return core.CGFloat(ret)

}

// SetMinimumFontSize returns the minimum font size, in points.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/1537155-minimumfontsize?language=objc for details.
func (x gen_WKPreferences) SetMinimumFontSize(
	value core.CGFloat,
) {
	C.WKPreferences_inst_setMinimumFontSize(
		unsafe.Pointer(x.Pointer()),
		C.double(value),
	)

	return

}

// TabFocusesLinks returns a boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/2818595-tabfocuseslinks?language=objc for details.
func (x gen_WKPreferences) TabFocusesLinks() bool {
	ret := C.WKPreferences_inst_tabFocusesLinks(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetTabFocusesLinks returns a boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/2818595-tabfocuseslinks?language=objc for details.
func (x gen_WKPreferences) SetTabFocusesLinks(
	value bool,
) {
	C.WKPreferences_inst_setTabFocusesLinks(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// JavaScriptCanOpenWindowsAutomatically returns a boolean value that indicates whether javascript can open windows without user interaction.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/1536573-javascriptcanopenwindowsautomati?language=objc for details.
func (x gen_WKPreferences) JavaScriptCanOpenWindowsAutomatically() bool {
	ret := C.WKPreferences_inst_javaScriptCanOpenWindowsAutomatically(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetJavaScriptCanOpenWindowsAutomatically returns a boolean value that indicates whether javascript can open windows without user interaction.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/1536573-javascriptcanopenwindowsautomati?language=objc for details.
func (x gen_WKPreferences) SetJavaScriptCanOpenWindowsAutomatically(
	value bool,
) {
	C.WKPreferences_inst_setJavaScriptCanOpenWindowsAutomatically(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsFraudulentWebsiteWarningEnabled returns a boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/3335219-fraudulentwebsitewarningenabled?language=objc for details.
func (x gen_WKPreferences) IsFraudulentWebsiteWarningEnabled() bool {
	ret := C.WKPreferences_inst_isFraudulentWebsiteWarningEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetFraudulentWebsiteWarningEnabled returns a boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/3335219-fraudulentwebsitewarningenabled?language=objc for details.
func (x gen_WKPreferences) SetFraudulentWebsiteWarningEnabled(
	value bool,
) {
	C.WKPreferences_inst_setFraudulentWebsiteWarningEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}

// IsTextInteractionEnabled
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/3727362-textinteractionenabled?language=objc for details.
func (x gen_WKPreferences) IsTextInteractionEnabled() bool {
	ret := C.WKPreferences_inst_isTextInteractionEnabled(
		unsafe.Pointer(x.Pointer()),
	)

	return convertObjCBoolToGo(ret)

}

// SetTextInteractionEnabled
//
// See https://developer.apple.com/documentation/webkit/wkpreferences/3727362-textinteractionenabled?language=objc for details.
func (x gen_WKPreferences) SetTextInteractionEnabled(
	value bool,
) {
	C.WKPreferences_inst_setTextInteractionEnabled(
		unsafe.Pointer(x.Pointer()),
		convertToObjCBool(value),
	)

	return

}
