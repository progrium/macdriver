// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptionConversionValidator] class.
var CaptionConversionValidatorClass = _CaptionConversionValidatorClass{objc.GetClass("AVCaptionConversionValidator")}

type _CaptionConversionValidatorClass struct {
	objc.Class
}

// An interface definition for the [CaptionConversionValidator] class.
type ICaptionConversionValidator interface {
	objc.IObject
	StopValidating()
	ValidateCaptionConversionWithWarningHandler(handler func(warning CaptionConversionWarning))
	Captions() []Caption
	TimeRange() coremedia.TimeRange
	Status() CaptionConversionValidatorStatus
	Warnings() []CaptionConversionWarning
}

// An object that validates captions for a conversion operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator?language=objc
type CaptionConversionValidator struct {
	objc.Object
}

func CaptionConversionValidatorFrom(ptr unsafe.Pointer) CaptionConversionValidator {
	return CaptionConversionValidator{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptionConversionValidatorClass) CaptionConversionValidatorWithCaptionsTimeRangeConversionSettings(captions []ICaption, timeRange coremedia.TimeRange, conversionSettings map[CaptionSettingsKey]objc.IObject) CaptionConversionValidator {
	rv := objc.Call[CaptionConversionValidator](cc, objc.Sel("captionConversionValidatorWithCaptions:timeRange:conversionSettings:"), captions, timeRange, conversionSettings)
	return rv
}

// A convenience initializer to create an object that validates captions for a conversion operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/3752935-captionconversionvalidatorwithca?language=objc
func CaptionConversionValidator_CaptionConversionValidatorWithCaptionsTimeRangeConversionSettings(captions []ICaption, timeRange coremedia.TimeRange, conversionSettings map[CaptionSettingsKey]objc.IObject) CaptionConversionValidator {
	return CaptionConversionValidatorClass.CaptionConversionValidatorWithCaptionsTimeRangeConversionSettings(captions, timeRange, conversionSettings)
}

func (c_ CaptionConversionValidator) InitWithCaptionsTimeRangeConversionSettings(captions []ICaption, timeRange coremedia.TimeRange, conversionSettings map[CaptionSettingsKey]objc.IObject) CaptionConversionValidator {
	rv := objc.Call[CaptionConversionValidator](c_, objc.Sel("initWithCaptions:timeRange:conversionSettings:"), captions, timeRange, conversionSettings)
	return rv
}

// Creates an object that validates captions for a conversion operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/3752937-initwithcaptions?language=objc
func NewCaptionConversionValidatorWithCaptionsTimeRangeConversionSettings(captions []ICaption, timeRange coremedia.TimeRange, conversionSettings map[CaptionSettingsKey]objc.IObject) CaptionConversionValidator {
	instance := CaptionConversionValidatorClass.Alloc().InitWithCaptionsTimeRangeConversionSettings(captions, timeRange, conversionSettings)
	instance.Autorelease()
	return instance
}

func (cc _CaptionConversionValidatorClass) Alloc() CaptionConversionValidator {
	rv := objc.Call[CaptionConversionValidator](cc, objc.Sel("alloc"))
	return rv
}

func CaptionConversionValidator_Alloc() CaptionConversionValidator {
	return CaptionConversionValidatorClass.Alloc()
}

func (cc _CaptionConversionValidatorClass) New() CaptionConversionValidator {
	rv := objc.Call[CaptionConversionValidator](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptionConversionValidator() CaptionConversionValidator {
	return CaptionConversionValidatorClass.New()
}

func (c_ CaptionConversionValidator) Init() CaptionConversionValidator {
	rv := objc.Call[CaptionConversionValidator](c_, objc.Sel("init"))
	return rv
}

// Stops the active validation operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/3752939-stopvalidating?language=objc
func (c_ CaptionConversionValidator) StopValidating() {
	objc.Call[objc.Void](c_, objc.Sel("stopValidating"))
}

// Validates the object’s captions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/3752941-validatecaptionconversionwithwar?language=objc
func (c_ CaptionConversionValidator) ValidateCaptionConversionWithWarningHandler(handler func(warning CaptionConversionWarning)) {
	objc.Call[objc.Void](c_, objc.Sel("validateCaptionConversionWithWarningHandler:"), handler)
}

// The array of captions that the system validates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/3752936-captions?language=objc
func (c_ CaptionConversionValidator) Captions() []Caption {
	rv := objc.Call[[]Caption](c_, objc.Sel("captions"))
	return rv
}

// The time range of the media timeline in which the captions must exist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/3752940-timerange?language=objc
func (c_ CaptionConversionValidator) TimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](c_, objc.Sel("timeRange"))
	return rv
}

// A value that indicates the status of validation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/3752938-status?language=objc
func (c_ CaptionConversionValidator) Status() CaptionConversionValidatorStatus {
	rv := objc.Call[CaptionConversionValidatorStatus](c_, objc.Sel("status"))
	return rv
}

// The collection of warnings the validator encountered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/3752942-warnings?language=objc
func (c_ CaptionConversionValidator) Warnings() []CaptionConversionWarning {
	rv := objc.Call[[]CaptionConversionWarning](c_, objc.Sel("warnings"))
	return rv
}
