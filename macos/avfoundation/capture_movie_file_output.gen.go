// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureMovieFileOutput] class.
var CaptureMovieFileOutputClass = _CaptureMovieFileOutputClass{objc.GetClass("AVCaptureMovieFileOutput")}

type _CaptureMovieFileOutputClass struct {
	objc.Class
}

// An interface definition for the [CaptureMovieFileOutput] class.
type ICaptureMovieFileOutput interface {
	ICaptureFileOutput
	OutputSettingsForConnection(connection ICaptureConnection) map[string]objc.Object
	SetOutputSettingsForConnection(outputSettings map[string]objc.IObject, connection ICaptureConnection)
	SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingRestrictedSwitchingBehaviorConditions(switchingBehavior CapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions)
	PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
	MovieFragmentInterval() coremedia.Time
	SetMovieFragmentInterval(value coremedia.Time)
	Metadata() []MetadataItem
	SetMetadata(value []IMetadataItem)
	IsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled() bool
	SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled(value bool)
}

// A capture output that records video and audio to a QuickTime movie file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput?language=objc
type CaptureMovieFileOutput struct {
	CaptureFileOutput
}

func CaptureMovieFileOutputFrom(ptr unsafe.Pointer) CaptureMovieFileOutput {
	return CaptureMovieFileOutput{
		CaptureFileOutput: CaptureFileOutputFrom(ptr),
	}
}

func (cc _CaptureMovieFileOutputClass) New() CaptureMovieFileOutput {
	rv := objc.Call[CaptureMovieFileOutput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureMovieFileOutput() CaptureMovieFileOutput {
	return CaptureMovieFileOutputClass.New()
}

func (c_ CaptureMovieFileOutput) Init() CaptureMovieFileOutput {
	rv := objc.Call[CaptureMovieFileOutput](c_, objc.Sel("init"))
	return rv
}

func (cc _CaptureMovieFileOutputClass) Alloc() CaptureMovieFileOutput {
	rv := objc.Call[CaptureMovieFileOutput](cc, objc.Sel("alloc"))
	return rv
}

func CaptureMovieFileOutput_Alloc() CaptureMovieFileOutput {
	return CaptureMovieFileOutputClass.Alloc()
}

// Returns the options used to reencode media from a given connection as it's being recorded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/1386479-outputsettingsforconnection?language=objc
func (c_ CaptureMovieFileOutput) OutputSettingsForConnection(connection ICaptureConnection) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("outputSettingsForConnection:"), objc.Ptr(connection))
	return rv
}

// Sets the options dictionary used to reencode media from the given connection as it's being recorded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/1388448-setoutputsettings?language=objc
func (c_ CaptureMovieFileOutput) SetOutputSettingsForConnection(outputSettings map[string]objc.IObject, connection ICaptureConnection) {
	objc.Call[objc.Void](c_, objc.Sel("setOutputSettings:forConnection:"), outputSettings, objc.Ptr(connection))
}

// Sets the camera switching behavior to use during recording. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/3875327-setprimaryconstituentdeviceswitc?language=objc
func (c_ CaptureMovieFileOutput) SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingRestrictedSwitchingBehaviorConditions(switchingBehavior CapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimaryConstituentDeviceSwitchingBehaviorForRecording:restrictedSwitchingBehaviorConditions:"), switchingBehavior, restrictedSwitchingBehaviorConditions)
}

// The conditions during which camera switching may occur while recording. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/3875324-primaryconstituentdevicerestrict?language=objc
func (c_ CaptureMovieFileOutput) PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions {
	rv := objc.Call[CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions](c_, objc.Sel("primaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording"))
	return rv
}

// The number of seconds of output that are written per fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/1387146-moviefragmentinterval?language=objc
func (c_ CaptureMovieFileOutput) MovieFragmentInterval() coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("movieFragmentInterval"))
	return rv
}

// The number of seconds of output that are written per fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/1387146-moviefragmentinterval?language=objc
func (c_ CaptureMovieFileOutput) SetMovieFragmentInterval(value coremedia.Time) {
	objc.Call[objc.Void](c_, objc.Sel("setMovieFragmentInterval:"), value)
}

// The metadata for the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/1387808-metadata?language=objc
func (c_ CaptureMovieFileOutput) Metadata() []MetadataItem {
	rv := objc.Call[[]MetadataItem](c_, objc.Sel("metadata"))
	return rv
}

// The metadata for the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/1387808-metadata?language=objc
func (c_ CaptureMovieFileOutput) SetMetadata(value []IMetadataItem) {
	objc.Call[objc.Void](c_, objc.Sel("setMetadata:"), value)
}

// A Boolean value that indicates whether to restrict constituent device switching behavior during recording. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/3875326-primaryconstituentdeviceswitchin?language=objc
func (c_ CaptureMovieFileOutput) IsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled() bool {
	rv := objc.Call[bool](c_, objc.Sel("isPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled"))
	return rv
}

// A Boolean value that indicates whether to restrict constituent device switching behavior during recording. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/3875326-primaryconstituentdeviceswitchin?language=objc
func (c_ CaptureMovieFileOutput) SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled:"), value)
}
