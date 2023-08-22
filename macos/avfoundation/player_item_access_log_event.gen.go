// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItemAccessLogEvent] class.
var PlayerItemAccessLogEventClass = _PlayerItemAccessLogEventClass{objc.GetClass("AVPlayerItemAccessLogEvent")}

type _PlayerItemAccessLogEventClass struct {
	objc.Class
}

// An interface definition for the [PlayerItemAccessLogEvent] class.
type IPlayerItemAccessLogEvent interface {
	objc.IObject
	AverageAudioBitrate() float64
	NumberOfMediaRequests() int
	TransferDuration() foundation.TimeInterval
	URI() string
	NumberOfDroppedVideoFrames() int
	NumberOfServerAddressChanges() int
	ObservedBitrate() float64
	ServerAddress() string
	AverageVideoBitrate() float64
	NumberOfBytesTransferred() int64
	PlaybackType() string
	IndicatedBitrate() float64
	DownloadOverdue() int
	IndicatedAverageBitrate() float64
	SwitchBitrate() float64
	PlaybackStartDate() foundation.Date
	ObservedBitrateStandardDeviation() float64
	DurationWatched() foundation.TimeInterval
	StartupTime() foundation.TimeInterval
	PlaybackStartOffset() foundation.TimeInterval
	MediaRequestsWWAN() int
	NumberOfStalls() int
	SegmentsDownloadedDuration() foundation.TimeInterval
	PlaybackSessionID() string
}

// A single entry in a player item's access log. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent?language=objc
type PlayerItemAccessLogEvent struct {
	objc.Object
}

func PlayerItemAccessLogEventFrom(ptr unsafe.Pointer) PlayerItemAccessLogEvent {
	return PlayerItemAccessLogEvent{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlayerItemAccessLogEventClass) Alloc() PlayerItemAccessLogEvent {
	rv := objc.Call[PlayerItemAccessLogEvent](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItemAccessLogEvent_Alloc() PlayerItemAccessLogEvent {
	return PlayerItemAccessLogEventClass.Alloc()
}

func (pc _PlayerItemAccessLogEventClass) New() PlayerItemAccessLogEvent {
	rv := objc.Call[PlayerItemAccessLogEvent](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItemAccessLogEvent() PlayerItemAccessLogEvent {
	return PlayerItemAccessLogEventClass.New()
}

func (p_ PlayerItemAccessLogEvent) Init() PlayerItemAccessLogEvent {
	rv := objc.Call[PlayerItemAccessLogEvent](p_, objc.Sel("init"))
	return rv
}

// The audio track’s average bit rate, in bits per second. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1643590-averageaudiobitrate?language=objc
func (p_ PlayerItemAccessLogEvent) AverageAudioBitrate() float64 {
	rv := objc.Call[float64](p_, objc.Sel("averageAudioBitrate"))
	return rv
}

// The number of media read requests from the server to this client. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1388357-numberofmediarequests?language=objc
func (p_ PlayerItemAccessLogEvent) NumberOfMediaRequests() int {
	rv := objc.Call[int](p_, objc.Sel("numberOfMediaRequests"))
	return rv
}

// The accumulated duration, in seconds, of active network transfer of bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1387370-transferduration?language=objc
func (p_ PlayerItemAccessLogEvent) TransferDuration() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](p_, objc.Sel("transferDuration"))
	return rv
}

// The URI of the playback item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1388643-uri?language=objc
func (p_ PlayerItemAccessLogEvent) URI() string {
	rv := objc.Call[string](p_, objc.Sel("URI"))
	return rv
}

// The total number of dropped video frames [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1388647-numberofdroppedvideoframes?language=objc
func (p_ PlayerItemAccessLogEvent) NumberOfDroppedVideoFrames() int {
	rv := objc.Call[int](p_, objc.Sel("numberOfDroppedVideoFrames"))
	return rv
}

// A count of changes to the server address over the last uninterrupted period of playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1388076-numberofserveraddresschanges?language=objc
func (p_ PlayerItemAccessLogEvent) NumberOfServerAddressChanges() int {
	rv := objc.Call[int](p_, objc.Sel("numberOfServerAddressChanges"))
	return rv
}

// The empirical throughput, in bits per second, across all media downloaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1390804-observedbitrate?language=objc
func (p_ PlayerItemAccessLogEvent) ObservedBitrate() float64 {
	rv := objc.Call[float64](p_, objc.Sel("observedBitrate"))
	return rv
}

// The IP address of the server that was the source of the last delivered media segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1390315-serveraddress?language=objc
func (p_ PlayerItemAccessLogEvent) ServerAddress() string {
	rv := objc.Call[string](p_, objc.Sel("serverAddress"))
	return rv
}

// The video track’s average bit rate, in bits per second. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1643592-averagevideobitrate?language=objc
func (p_ PlayerItemAccessLogEvent) AverageVideoBitrate() float64 {
	rv := objc.Call[float64](p_, objc.Sel("averageVideoBitrate"))
	return rv
}

// The accumulated number of bytes transferred by the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1387305-numberofbytestransferred?language=objc
func (p_ PlayerItemAccessLogEvent) NumberOfBytesTransferred() int64 {
	rv := objc.Call[int64](p_, objc.Sel("numberOfBytesTransferred"))
	return rv
}

// The playback type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1387218-playbacktype?language=objc
func (p_ PlayerItemAccessLogEvent) PlaybackType() string {
	rv := objc.Call[string](p_, objc.Sel("playbackType"))
	return rv
}

// The throughput, in bits per second, required to play the stream, as advertised by the server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1388123-indicatedbitrate?language=objc
func (p_ PlayerItemAccessLogEvent) IndicatedBitrate() float64 {
	rv := objc.Call[float64](p_, objc.Sel("indicatedBitrate"))
	return rv
}

// The total number of times that downloading the segments took too long. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1389213-downloadoverdue?language=objc
func (p_ PlayerItemAccessLogEvent) DownloadOverdue() int {
	rv := objc.Call[int](p_, objc.Sel("downloadOverdue"))
	return rv
}

// The average throughput, in bits per second, required to play the stream, as advertised by the server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1872546-indicatedaveragebitrate?language=objc
func (p_ PlayerItemAccessLogEvent) IndicatedAverageBitrate() float64 {
	rv := objc.Call[float64](p_, objc.Sel("indicatedAverageBitrate"))
	return rv
}

// The bandwidth value that causes a switch, up or down, in the item's quality being played. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1390645-switchbitrate?language=objc
func (p_ PlayerItemAccessLogEvent) SwitchBitrate() float64 {
	rv := objc.Call[float64](p_, objc.Sel("switchBitrate"))
	return rv
}

// The date and time at which playback began for this event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1390502-playbackstartdate?language=objc
func (p_ PlayerItemAccessLogEvent) PlaybackStartDate() foundation.Date {
	rv := objc.Call[foundation.Date](p_, objc.Sel("playbackStartDate"))
	return rv
}

// The standard deviation of the observed segment download bit rates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1386094-observedbitratestandarddeviation?language=objc
func (p_ PlayerItemAccessLogEvent) ObservedBitrateStandardDeviation() float64 {
	rv := objc.Call[float64](p_, objc.Sel("observedBitrateStandardDeviation"))
	return rv
}

// The accumulated duration, in seconds, of the media played. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1388200-durationwatched?language=objc
func (p_ PlayerItemAccessLogEvent) DurationWatched() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](p_, objc.Sel("durationWatched"))
	return rv
}

// The accumulated duration, in seconds, until the player item is ready to play. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1389138-startuptime?language=objc
func (p_ PlayerItemAccessLogEvent) StartupTime() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](p_, objc.Sel("startupTime"))
	return rv
}

// The offset, in seconds, in the playlist where the last uninterrupted period of playback began. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1385922-playbackstartoffset?language=objc
func (p_ PlayerItemAccessLogEvent) PlaybackStartOffset() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](p_, objc.Sel("playbackStartOffset"))
	return rv
}

// The number of network read requests over a WWAN. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1388549-mediarequestswwan?language=objc
func (p_ PlayerItemAccessLogEvent) MediaRequestsWWAN() int {
	rv := objc.Call[int](p_, objc.Sel("mediaRequestsWWAN"))
	return rv
}

// The total number of playback stalls encountered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1387712-numberofstalls?language=objc
func (p_ PlayerItemAccessLogEvent) NumberOfStalls() int {
	rv := objc.Call[int](p_, objc.Sel("numberOfStalls"))
	return rv
}

// The accumulated duration, in seconds, of the media segments downloaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1388147-segmentsdownloadedduration?language=objc
func (p_ PlayerItemAccessLogEvent) SegmentsDownloadedDuration() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](p_, objc.Sel("segmentsDownloadedDuration"))
	return rv
}

// A GUID that identifies the playback session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemaccesslogevent/1388462-playbacksessionid?language=objc
func (p_ PlayerItemAccessLogEvent) PlaybackSessionID() string {
	rv := objc.Call[string](p_, objc.Sel("playbackSessionID"))
	return rv
}
