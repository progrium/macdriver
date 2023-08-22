// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetReaderOutputCaptionAdaptor] class.
var AssetReaderOutputCaptionAdaptorClass = _AssetReaderOutputCaptionAdaptorClass{objc.GetClass("AVAssetReaderOutputCaptionAdaptor")}

type _AssetReaderOutputCaptionAdaptorClass struct {
	objc.Class
}

// An interface definition for the [AssetReaderOutputCaptionAdaptor] class.
type IAssetReaderOutputCaptionAdaptor interface {
	objc.IObject
	CaptionsNotPresentInPreviousGroupsInCaptionGroup(captionGroup ICaptionGroup) []Caption
	NextCaptionGroup() CaptionGroup
	AssetReaderTrackOutput() AssetReaderTrackOutput
	ValidationDelegate() AssetReaderCaptionValidationHandlingWrapper
	SetValidationDelegate(value PAssetReaderCaptionValidationHandling)
	SetValidationDelegateObject(valueObject objc.IObject)
}

// An object that reads caption group objects from an asset track that contains timed text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputcaptionadaptor?language=objc
type AssetReaderOutputCaptionAdaptor struct {
	objc.Object
}

func AssetReaderOutputCaptionAdaptorFrom(ptr unsafe.Pointer) AssetReaderOutputCaptionAdaptor {
	return AssetReaderOutputCaptionAdaptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AssetReaderOutputCaptionAdaptor) InitWithAssetReaderTrackOutput(trackOutput IAssetReaderTrackOutput) AssetReaderOutputCaptionAdaptor {
	rv := objc.Call[AssetReaderOutputCaptionAdaptor](a_, objc.Sel("initWithAssetReaderTrackOutput:"), objc.Ptr(trackOutput))
	return rv
}

// Creates a caption adaptor that reads from a track output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputcaptionadaptor/3752797-initwithassetreadertrackoutput?language=objc
func NewAssetReaderOutputCaptionAdaptorWithAssetReaderTrackOutput(trackOutput IAssetReaderTrackOutput) AssetReaderOutputCaptionAdaptor {
	instance := AssetReaderOutputCaptionAdaptorClass.Alloc().InitWithAssetReaderTrackOutput(trackOutput)
	instance.Autorelease()
	return instance
}

func (ac _AssetReaderOutputCaptionAdaptorClass) AssetReaderOutputCaptionAdaptorWithAssetReaderTrackOutput(trackOutput IAssetReaderTrackOutput) AssetReaderOutputCaptionAdaptor {
	rv := objc.Call[AssetReaderOutputCaptionAdaptor](ac, objc.Sel("assetReaderOutputCaptionAdaptorWithAssetReaderTrackOutput:"), objc.Ptr(trackOutput))
	return rv
}

// A class method that creates a caption adaptor that reads from a track output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputcaptionadaptor/3752794-assetreaderoutputcaptionadaptorw?language=objc
func AssetReaderOutputCaptionAdaptor_AssetReaderOutputCaptionAdaptorWithAssetReaderTrackOutput(trackOutput IAssetReaderTrackOutput) AssetReaderOutputCaptionAdaptor {
	return AssetReaderOutputCaptionAdaptorClass.AssetReaderOutputCaptionAdaptorWithAssetReaderTrackOutput(trackOutput)
}

func (ac _AssetReaderOutputCaptionAdaptorClass) Alloc() AssetReaderOutputCaptionAdaptor {
	rv := objc.Call[AssetReaderOutputCaptionAdaptor](ac, objc.Sel("alloc"))
	return rv
}

func AssetReaderOutputCaptionAdaptor_Alloc() AssetReaderOutputCaptionAdaptor {
	return AssetReaderOutputCaptionAdaptorClass.Alloc()
}

func (ac _AssetReaderOutputCaptionAdaptorClass) New() AssetReaderOutputCaptionAdaptor {
	rv := objc.Call[AssetReaderOutputCaptionAdaptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetReaderOutputCaptionAdaptor() AssetReaderOutputCaptionAdaptor {
	return AssetReaderOutputCaptionAdaptorClass.New()
}

func (a_ AssetReaderOutputCaptionAdaptor) Init() AssetReaderOutputCaptionAdaptor {
	rv := objc.Call[AssetReaderOutputCaptionAdaptor](a_, objc.Sel("init"))
	return rv
}

// Returns the set of captions in the caption group that weren’t vended by the adaptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputcaptionadaptor/3752796-captionsnotpresentinpreviousgrou?language=objc
func (a_ AssetReaderOutputCaptionAdaptor) CaptionsNotPresentInPreviousGroupsInCaptionGroup(captionGroup ICaptionGroup) []Caption {
	rv := objc.Call[[]Caption](a_, objc.Sel("captionsNotPresentInPreviousGroupsInCaptionGroup:"), objc.Ptr(captionGroup))
	return rv
}

// Returns the next caption group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputcaptionadaptor/3752798-nextcaptiongroup?language=objc
func (a_ AssetReaderOutputCaptionAdaptor) NextCaptionGroup() CaptionGroup {
	rv := objc.Call[CaptionGroup](a_, objc.Sel("nextCaptionGroup"))
	return rv
}

// The associated asset reader track output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputcaptionadaptor/3752795-assetreadertrackoutput?language=objc
func (a_ AssetReaderOutputCaptionAdaptor) AssetReaderTrackOutput() AssetReaderTrackOutput {
	rv := objc.Call[AssetReaderTrackOutput](a_, objc.Sel("assetReaderTrackOutput"))
	return rv
}

// A delegate object that handles callbacks to the caption adaptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputcaptionadaptor/3752799-validationdelegate?language=objc
func (a_ AssetReaderOutputCaptionAdaptor) ValidationDelegate() AssetReaderCaptionValidationHandlingWrapper {
	rv := objc.Call[AssetReaderCaptionValidationHandlingWrapper](a_, objc.Sel("validationDelegate"))
	return rv
}

// A delegate object that handles callbacks to the caption adaptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputcaptionadaptor/3752799-validationdelegate?language=objc
func (a_ AssetReaderOutputCaptionAdaptor) SetValidationDelegate(value PAssetReaderCaptionValidationHandling) {
	po0 := objc.WrapAsProtocol("AVAssetReaderCaptionValidationHandling", value)
	objc.SetAssociatedObject(a_, objc.AssociationKey("setValidationDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](a_, objc.Sel("setValidationDelegate:"), po0)
}

// A delegate object that handles callbacks to the caption adaptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutputcaptionadaptor/3752799-validationdelegate?language=objc
func (a_ AssetReaderOutputCaptionAdaptor) SetValidationDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setValidationDelegate:"), objc.Ptr(valueObject))
}
