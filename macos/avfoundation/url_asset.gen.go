// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLAsset] class.
var URLAssetClass = _URLAssetClass{objc.GetClass("AVURLAsset")}

type _URLAssetClass struct {
	objc.Class
}

// An interface definition for the [URLAsset] class.
type IURLAsset interface {
	IAsset
	FindCompatibleTrackForCompositionTrackCompletionHandler(compositionTrack ICompositionTrack, completionHandler func(arg0 AssetTrack, arg1 foundation.Error))
	ResourceLoader() AssetResourceLoader
	MayRequireContentKeysForMediaDataProcessing() bool
	URL() foundation.URL
	AssetCache() AssetCache
	Variants() []AssetVariant
}

// An asset that represents media at a local or remote URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset?language=objc
type URLAsset struct {
	Asset
}

func URLAssetFrom(ptr unsafe.Pointer) URLAsset {
	return URLAsset{
		Asset: AssetFrom(ptr),
	}
}

func (uc _URLAssetClass) URLAssetWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) URLAsset {
	rv := objc.Call[URLAsset](uc, objc.Sel("URLAssetWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Returns an asset that models the media resource found at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1508727-urlassetwithurl?language=objc
func URLAsset_URLAssetWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) URLAsset {
	return URLAssetClass.URLAssetWithURLOptions(URL, options)
}

func (u_ URLAsset) InitWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) URLAsset {
	rv := objc.Call[URLAsset](u_, objc.Sel("initWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Creates an asset that models the media resource at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1385698-initwithurl?language=objc
func NewURLAssetWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) URLAsset {
	instance := URLAssetClass.Alloc().InitWithURLOptions(URL, options)
	instance.Autorelease()
	return instance
}

func (uc _URLAssetClass) Alloc() URLAsset {
	rv := objc.Call[URLAsset](uc, objc.Sel("alloc"))
	return rv
}

func URLAsset_Alloc() URLAsset {
	return URLAssetClass.Alloc()
}

func (uc _URLAssetClass) New() URLAsset {
	rv := objc.Call[URLAsset](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLAsset() URLAsset {
	return URLAssetClass.New()
}

func (u_ URLAsset) Init() URLAsset {
	rv := objc.Call[URLAsset](u_, objc.Sel("init"))
	return rv
}

func (uc _URLAssetClass) AssetWithURL(URL foundation.IURL) URLAsset {
	rv := objc.Call[URLAsset](uc, objc.Sel("assetWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates an asset that models the media at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389943-assetwithurl?language=objc
func URLAsset_AssetWithURL(URL foundation.IURL) URLAsset {
	return URLAssetClass.AssetWithURL(URL)
}

// Returns an array of the file types the asset supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1386800-audiovisualtypes?language=objc
func (uc _URLAssetClass) AudiovisualTypes() []FileType {
	rv := objc.Call[[]FileType](uc, objc.Sel("audiovisualTypes"))
	return rv
}

// Returns an array of the file types the asset supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1386800-audiovisualtypes?language=objc
func URLAsset_AudiovisualTypes() []FileType {
	return URLAssetClass.AudiovisualTypes()
}

// Returns a Boolean value that indicates whether the asset is playable with the specified codecs and container type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1387142-isplayableextendedmimetype?language=objc
func (uc _URLAssetClass) IsPlayableExtendedMIMEType(extendedMIMEType string) bool {
	rv := objc.Call[bool](uc, objc.Sel("isPlayableExtendedMIMEType:"), extendedMIMEType)
	return rv
}

// Returns a Boolean value that indicates whether the asset is playable with the specified codecs and container type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1387142-isplayableextendedmimetype?language=objc
func URLAsset_IsPlayableExtendedMIMEType(extendedMIMEType string) bool {
	return URLAssetClass.IsPlayableExtendedMIMEType(extendedMIMEType)
}

// Returns an array of the MIME types the asset supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1390006-audiovisualmimetypes?language=objc
func (uc _URLAssetClass) AudiovisualMIMETypes() []string {
	rv := objc.Call[[]string](uc, objc.Sel("audiovisualMIMETypes"))
	return rv
}

// Returns an array of the MIME types the asset supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1390006-audiovisualmimetypes?language=objc
func URLAsset_AudiovisualMIMETypes() []string {
	return URLAssetClass.AudiovisualMIMETypes()
}

// Loads an asset track from which you can insert any time range into the composition track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/3746535-findcompatibletrackforcompositio?language=objc
func (u_ URLAsset) FindCompatibleTrackForCompositionTrackCompletionHandler(compositionTrack ICompositionTrack, completionHandler func(arg0 AssetTrack, arg1 foundation.Error)) {
	objc.Call[objc.Void](u_, objc.Sel("findCompatibleTrackForCompositionTrack:completionHandler:"), objc.Ptr(compositionTrack), completionHandler)
}

// The resource loader for the asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1389118-resourceloader?language=objc
func (u_ URLAsset) ResourceLoader() AssetResourceLoader {
	rv := objc.Call[AssetResourceLoader](u_, objc.Sel("resourceLoader"))
	return rv
}

// A Boolean value that indicates whether you can add this asset as a content key recipient to a content key session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/2806807-mayrequirecontentkeysformediadat?language=objc
func (u_ URLAsset) MayRequireContentKeysForMediaDataProcessing() bool {
	rv := objc.Call[bool](u_, objc.Sel("mayRequireContentKeysForMediaDataProcessing"))
	return rv
}

// A URL to the asset’s media. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1388127-url?language=objc
func (u_ URLAsset) URL() foundation.URL {
	rv := objc.Call[foundation.URL](u_, objc.Sel("URL"))
	return rv
}

// The asset’s associated asset cache, if it exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/1823714-assetcache?language=objc
func (u_ URLAsset) AssetCache() AssetCache {
	rv := objc.Call[AssetCache](u_, objc.Sel("assetCache"))
	return rv
}

// An array of variants that an asset contains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/3746536-variants?language=objc
func (u_ URLAsset) Variants() []AssetVariant {
	rv := objc.Call[[]AssetVariant](u_, objc.Sel("variants"))
	return rv
}
