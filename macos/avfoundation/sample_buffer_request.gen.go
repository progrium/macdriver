// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SampleBufferRequest] class.
var SampleBufferRequestClass = _SampleBufferRequestClass{objc.GetClass("AVSampleBufferRequest")}

type _SampleBufferRequestClass struct {
	objc.Class
}

// An interface definition for the [SampleBufferRequest] class.
type ISampleBufferRequest interface {
	objc.IObject
	Direction() SampleBufferRequestDirection
	SetDirection(value SampleBufferRequestDirection)
	LimitCursor() SampleCursor
	SetLimitCursor(value ISampleCursor)
	PreferredMinSampleCount() int
	SetPreferredMinSampleCount(value int)
	MaxSampleCount() int
	SetMaxSampleCount(value int)
	OverrideTime() coremedia.Time
	SetOverrideTime(value coremedia.Time)
	Mode() SampleBufferRequestMode
	SetMode(value SampleBufferRequestMode)
	StartCursor() SampleCursor
}

// An object that describes a sample buffer creation request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest?language=objc
type SampleBufferRequest struct {
	objc.Object
}

func SampleBufferRequestFrom(ptr unsafe.Pointer) SampleBufferRequest {
	return SampleBufferRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SampleBufferRequest) InitWithStartCursor(startCursor ISampleCursor) SampleBufferRequest {
	rv := objc.Call[SampleBufferRequest](s_, objc.Sel("initWithStartCursor:"), objc.Ptr(startCursor))
	return rv
}

// Creates a newly allocated sample buffer request with the specified sample cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1387449-initwithstartcursor?language=objc
func NewSampleBufferRequestWithStartCursor(startCursor ISampleCursor) SampleBufferRequest {
	instance := SampleBufferRequestClass.Alloc().InitWithStartCursor(startCursor)
	instance.Autorelease()
	return instance
}

func (sc _SampleBufferRequestClass) Alloc() SampleBufferRequest {
	rv := objc.Call[SampleBufferRequest](sc, objc.Sel("alloc"))
	return rv
}

func SampleBufferRequest_Alloc() SampleBufferRequest {
	return SampleBufferRequestClass.Alloc()
}

func (sc _SampleBufferRequestClass) New() SampleBufferRequest {
	rv := objc.Call[SampleBufferRequest](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSampleBufferRequest() SampleBufferRequest {
	return SampleBufferRequestClass.New()
}

func (s_ SampleBufferRequest) Init() SampleBufferRequest {
	rv := objc.Call[SampleBufferRequest](s_, objc.Sel("init"))
	return rv
}

// The buffer sample direction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1386442-direction?language=objc
func (s_ SampleBufferRequest) Direction() SampleBufferRequestDirection {
	rv := objc.Call[SampleBufferRequestDirection](s_, objc.Sel("direction"))
	return rv
}

// The buffer sample direction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1386442-direction?language=objc
func (s_ SampleBufferRequest) SetDirection(value SampleBufferRequestDirection) {
	objc.Call[objc.Void](s_, objc.Sel("setDirection:"), value)
}

// The limiting position for sample loading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1387466-limitcursor?language=objc
func (s_ SampleBufferRequest) LimitCursor() SampleCursor {
	rv := objc.Call[SampleCursor](s_, objc.Sel("limitCursor"))
	return rv
}

// The limiting position for sample loading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1387466-limitcursor?language=objc
func (s_ SampleBufferRequest) SetLimitCursor(value ISampleCursor) {
	objc.Call[objc.Void](s_, objc.Sel("setLimitCursor:"), objc.Ptr(value))
}

// The preferred minimum number of samples to load. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1386251-preferredminsamplecount?language=objc
func (s_ SampleBufferRequest) PreferredMinSampleCount() int {
	rv := objc.Call[int](s_, objc.Sel("preferredMinSampleCount"))
	return rv
}

// The preferred minimum number of samples to load. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1386251-preferredminsamplecount?language=objc
func (s_ SampleBufferRequest) SetPreferredMinSampleCount(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setPreferredMinSampleCount:"), value)
}

// The maximum number of samples to load. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1385645-maxsamplecount?language=objc
func (s_ SampleBufferRequest) MaxSampleCount() int {
	rv := objc.Call[int](s_, objc.Sel("maxSampleCount"))
	return rv
}

// The maximum number of samples to load. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1385645-maxsamplecount?language=objc
func (s_ SampleBufferRequest) SetMaxSampleCount(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setMaxSampleCount:"), value)
}

// The deadline for sample data and output PTS for the sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1385790-overridetime?language=objc
func (s_ SampleBufferRequest) OverrideTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](s_, objc.Sel("overrideTime"))
	return rv
}

// The deadline for sample data and output PTS for the sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1385790-overridetime?language=objc
func (s_ SampleBufferRequest) SetOverrideTime(value coremedia.Time) {
	objc.Call[objc.Void](s_, objc.Sel("setOverrideTime:"), value)
}

// The sample buffer request mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1387463-mode?language=objc
func (s_ SampleBufferRequest) Mode() SampleBufferRequestMode {
	rv := objc.Call[SampleBufferRequestMode](s_, objc.Sel("mode"))
	return rv
}

// The sample buffer request mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1387463-mode?language=objc
func (s_ SampleBufferRequest) SetMode(value SampleBufferRequestMode) {
	objc.Call[objc.Void](s_, objc.Sel("setMode:"), value)
}

// The starting cursor position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrequest/1387398-startcursor?language=objc
func (s_ SampleBufferRequest) StartCursor() SampleCursor {
	rv := objc.Call[SampleCursor](s_, objc.Sel("startCursor"))
	return rv
}
