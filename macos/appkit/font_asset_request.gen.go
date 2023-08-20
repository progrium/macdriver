// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FontAssetRequest] class.
var FontAssetRequestClass = _FontAssetRequestClass{objc.GetClass("NSFontAssetRequest")}

type _FontAssetRequestClass struct {
	objc.Class
}

// An interface definition for the [FontAssetRequest] class.
type IFontAssetRequest interface {
	objc.IObject
	DownloadFontAssetsWithCompletionHandler(completionHandler func(error foundation.Error) bool)
	DownloadedFontDescriptors() []FontDescriptor
	Progress() foundation.Progress
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontassetrequest?language=objc
type FontAssetRequest struct {
	objc.Object
}

func FontAssetRequestFrom(ptr unsafe.Pointer) FontAssetRequest {
	return FontAssetRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FontAssetRequest) InitWithFontDescriptorsOptions(fontDescriptors []IFontDescriptor, options FontAssetRequestOptions) FontAssetRequest {
	rv := objc.Call[FontAssetRequest](f_, objc.Sel("initWithFontDescriptors:options:"), fontDescriptors, options)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontassetrequest/2890807-initwithfontdescriptors?language=objc
func FontAssetRequest_InitWithFontDescriptorsOptions(fontDescriptors []IFontDescriptor, options FontAssetRequestOptions) FontAssetRequest {
	return FontAssetRequestClass.Alloc().InitWithFontDescriptorsOptions(fontDescriptors, options)
}

func (fc _FontAssetRequestClass) Alloc() FontAssetRequest {
	rv := objc.Call[FontAssetRequest](fc, objc.Sel("alloc"))
	return rv
}

func FontAssetRequest_Alloc() FontAssetRequest {
	return FontAssetRequestClass.Alloc()
}

func (fc _FontAssetRequestClass) New() FontAssetRequest {
	rv := objc.Call[FontAssetRequest](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFontAssetRequest() FontAssetRequest {
	return FontAssetRequestClass.New()
}

func (f_ FontAssetRequest) Init() FontAssetRequest {
	rv := objc.Call[FontAssetRequest](f_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontassetrequest/2890808-downloadfontassetswithcompletion?language=objc
func (f_ FontAssetRequest) DownloadFontAssetsWithCompletionHandler(completionHandler func(error foundation.Error) bool) {
	objc.Call[objc.Void](f_, objc.Sel("downloadFontAssetsWithCompletionHandler:"), completionHandler)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontassetrequest/2890802-downloadedfontdescriptors?language=objc
func (f_ FontAssetRequest) DownloadedFontDescriptors() []FontDescriptor {
	rv := objc.Call[[]FontDescriptor](f_, objc.Sel("downloadedFontDescriptors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontassetrequest/2890804-progress?language=objc
func (f_ FontAssetRequest) Progress() foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("progress"))
	return rv
}