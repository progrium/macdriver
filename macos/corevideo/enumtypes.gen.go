// AUTO-GENERATED CODE, DO NOT MODIFY

package corevideo

// The propagation modes of a Core Video buffer attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvattachmentmode?language=objc
type AttachmentMode uint32

const (
	KAttachmentMode_ShouldNotPropagate AttachmentMode = 0
	KAttachmentMode_ShouldPropagate    AttachmentMode = 1
)

// The flags to be used for the display link output callback function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvoptionflags?language=objc
type OptionFlags uint64

// The flags to pass to CVPixelBufferLockBaseAddress and CVPixelBufferUnlockBaseAddress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvpixelbufferlockflags?language=objc
type PixelBufferLockFlags OptionFlags

const (
	KPixelBufferLock_ReadOnly PixelBufferLockFlags = 1
)

// The flags to pass to flush the pool. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvpixelbufferpoolflushflags?language=objc
type PixelBufferPoolFlushFlags OptionFlags

const (
	KPixelBufferPoolFlushExcessBuffers PixelBufferPoolFlushFlags = 1
)

// A Core Video error type return value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvreturn?language=objc
type Return int32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvsmptetimeflags?language=objc
type SMPTETimeFlags uint32

const (
	KSMPTETimeRunning SMPTETimeFlags = 2
	KSMPTETimeValid   SMPTETimeFlags = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvsmptetimetype?language=objc
type SMPTETimeType uint32

const (
	KSMPTETimeType24       SMPTETimeType = 0
	KSMPTETimeType25       SMPTETimeType = 1
	KSMPTETimeType2997     SMPTETimeType = 4
	KSMPTETimeType2997Drop SMPTETimeType = 5
	KSMPTETimeType30       SMPTETimeType = 3
	KSMPTETimeType30Drop   SMPTETimeType = 2
	KSMPTETimeType5994     SMPTETimeType = 7
	KSMPTETimeType60       SMPTETimeType = 6
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvtimeflags?language=objc
type TimeFlags int32

const (
	KTimeIsIndefinite TimeFlags = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corevideo/cvtimestampflags?language=objc
type TimeStampFlags uint64

const (
	KTimeStampBottomField             TimeStampFlags = 131072
	KTimeStampHostTimeValid           TimeStampFlags = 2
	KTimeStampIsInterlaced            TimeStampFlags = 196608
	KTimeStampRateScalarValid         TimeStampFlags = 16
	KTimeStampSMPTETimeValid          TimeStampFlags = 4
	KTimeStampTopField                TimeStampFlags = 65536
	KTimeStampVideoHostTimeValid      TimeStampFlags = 3
	KTimeStampVideoRefreshPeriodValid TimeStampFlags = 8
	KTimeStampVideoTimeValid          TimeStampFlags = 1
)
