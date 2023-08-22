// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetReaderAudioMixOutput] class.
var AssetReaderAudioMixOutputClass = _AssetReaderAudioMixOutputClass{objc.GetClass("AVAssetReaderAudioMixOutput")}

type _AssetReaderAudioMixOutputClass struct {
	objc.Class
}

// An interface definition for the [AssetReaderAudioMixOutput] class.
type IAssetReaderAudioMixOutput interface {
	IAssetReaderOutput
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm)
	AudioSettings() map[string]objc.Object
	AudioMix() AudioMix
	SetAudioMix(value IAudioMix)
	AudioTracks() []AssetTrack
}

// An object that reads audio samples that result from mixing audio from one or more tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderaudiomixoutput?language=objc
type AssetReaderAudioMixOutput struct {
	AssetReaderOutput
}

func AssetReaderAudioMixOutputFrom(ptr unsafe.Pointer) AssetReaderAudioMixOutput {
	return AssetReaderAudioMixOutput{
		AssetReaderOutput: AssetReaderOutputFrom(ptr),
	}
}

func (ac _AssetReaderAudioMixOutputClass) AssetReaderAudioMixOutputWithAudioTracksAudioSettings(audioTracks []IAssetTrack, audioSettings map[string]objc.IObject) AssetReaderAudioMixOutput {
	rv := objc.Call[AssetReaderAudioMixOutput](ac, objc.Sel("assetReaderAudioMixOutputWithAudioTracks:audioSettings:"), audioTracks, audioSettings)
	return rv
}

// Creates an object that reads mixed audio from the specified audio tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderaudiomixoutput/1490328-assetreaderaudiomixoutputwithaud?language=objc
func AssetReaderAudioMixOutput_AssetReaderAudioMixOutputWithAudioTracksAudioSettings(audioTracks []IAssetTrack, audioSettings map[string]objc.IObject) AssetReaderAudioMixOutput {
	return AssetReaderAudioMixOutputClass.AssetReaderAudioMixOutputWithAudioTracksAudioSettings(audioTracks, audioSettings)
}

func (a_ AssetReaderAudioMixOutput) InitWithAudioTracksAudioSettings(audioTracks []IAssetTrack, audioSettings map[string]objc.IObject) AssetReaderAudioMixOutput {
	rv := objc.Call[AssetReaderAudioMixOutput](a_, objc.Sel("initWithAudioTracks:audioSettings:"), audioTracks, audioSettings)
	return rv
}

// Creates an object that reads mixed audio from the specified audio tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderaudiomixoutput/1388883-initwithaudiotracks?language=objc
func NewAssetReaderAudioMixOutputWithAudioTracksAudioSettings(audioTracks []IAssetTrack, audioSettings map[string]objc.IObject) AssetReaderAudioMixOutput {
	instance := AssetReaderAudioMixOutputClass.Alloc().InitWithAudioTracksAudioSettings(audioTracks, audioSettings)
	instance.Autorelease()
	return instance
}

func (ac _AssetReaderAudioMixOutputClass) Alloc() AssetReaderAudioMixOutput {
	rv := objc.Call[AssetReaderAudioMixOutput](ac, objc.Sel("alloc"))
	return rv
}

func AssetReaderAudioMixOutput_Alloc() AssetReaderAudioMixOutput {
	return AssetReaderAudioMixOutputClass.Alloc()
}

func (ac _AssetReaderAudioMixOutputClass) New() AssetReaderAudioMixOutput {
	rv := objc.Call[AssetReaderAudioMixOutput](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetReaderAudioMixOutput() AssetReaderAudioMixOutput {
	return AssetReaderAudioMixOutputClass.New()
}

func (a_ AssetReaderAudioMixOutput) Init() AssetReaderAudioMixOutput {
	rv := objc.Call[AssetReaderAudioMixOutput](a_, objc.Sel("init"))
	return rv
}

// The processing algorithm to use for scaled audio edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderaudiomixoutput/1388713-audiotimepitchalgorithm?language=objc
func (a_ AssetReaderAudioMixOutput) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Call[AudioTimePitchAlgorithm](a_, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}

// The processing algorithm to use for scaled audio edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderaudiomixoutput/1388713-audiotimepitchalgorithm?language=objc
func (a_ AssetReaderAudioMixOutput) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm) {
	objc.Call[objc.Void](a_, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}

// The audio settings that the output uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderaudiomixoutput/1388860-audiosettings?language=objc
func (a_ AssetReaderAudioMixOutput) AudioSettings() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](a_, objc.Sel("audioSettings"))
	return rv
}

// The audio mix to use with this output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderaudiomixoutput/1387074-audiomix?language=objc
func (a_ AssetReaderAudioMixOutput) AudioMix() AudioMix {
	rv := objc.Call[AudioMix](a_, objc.Sel("audioMix"))
	return rv
}

// The audio mix to use with this output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderaudiomixoutput/1387074-audiomix?language=objc
func (a_ AssetReaderAudioMixOutput) SetAudioMix(value IAudioMix) {
	objc.Call[objc.Void](a_, objc.Sel("setAudioMix:"), objc.Ptr(value))
}

// The tracks from which the output reads audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderaudiomixoutput/1385635-audiotracks?language=objc
func (a_ AssetReaderAudioMixOutput) AudioTracks() []AssetTrack {
	rv := objc.Call[[]AssetTrack](a_, objc.Sel("audioTracks"))
	return rv
}
