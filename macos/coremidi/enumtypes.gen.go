// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

// MIDI status types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicvstatus?language=objc
type CVStatus int

const (
	KCVStatusAssignableControl    CVStatus = 3
	KCVStatusAssignablePNC        CVStatus = 1
	KCVStatusChannelPressure      CVStatus = 13
	KCVStatusControlChange        CVStatus = 11
	KCVStatusNoteOff              CVStatus = 8
	KCVStatusNoteOn               CVStatus = 9
	KCVStatusPerNoteMgmt          CVStatus = 15
	KCVStatusPerNotePitchBend     CVStatus = 6
	KCVStatusPitchBend            CVStatus = 14
	KCVStatusPolyPressure         CVStatus = 10
	KCVStatusProgramChange        CVStatus = 12
	KCVStatusRegisteredControl    CVStatus = 2
	KCVStatusRegisteredPNC        CVStatus = 0
	KCVStatusRelAssignableControl CVStatus = 5
	KCVStatusRelRegisteredControl CVStatus = 4
)

// MIDI Channel, 0~15 (channels 1 through 16, respectively), or MIDIChannelsWholePort. Per the MIDI-CI specification, this is always a single byte. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midichannelnumber?language=objc
type ChannelNumber uint8

const (
	ChannelsWholePort ChannelNumber = 127
)

// An object that maintains per-client state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiclientref?language=objc
type ClientRef ObjectRef

// A list of MIDI devices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/mididevicelistref?language=objc
type DeviceListRef ObjectRef

// A MIDI device that contains entities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midideviceref?language=objc
type DeviceRef ObjectRef

// A MIDI source or destination an entity owns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiendpointref?language=objc
type EndpointRef ObjectRef

// An entity that a device owns and that contains endpoints. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midientityref?language=objc
type EntityRef ObjectRef

// Supported MIDI message types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midimessagetype?language=objc
type MessageType int

const (
	KMessageTypeChannelVoice1 MessageType = 2
	KMessageTypeChannelVoice2 MessageType = 4
	KMessageTypeData128       MessageType = 5
	KMessageTypeSysEx         MessageType = 3
	KMessageTypeSystem        MessageType = 1
	KMessageTypeUnknownF      MessageType = 15
	KMessageTypeUtility       MessageType = 0
)

// A 32-bit MIDI message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midimessage_32?language=objc
type Message_32 uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkconnectionpolicy?language=objc
type NetworkConnectionPolicy uint

const (
	NetworkConnectionPolicy_Anyone             NetworkConnectionPolicy = 2
	NetworkConnectionPolicy_HostsInContactList NetworkConnectionPolicy = 1
	NetworkConnectionPolicy_NoOne              NetworkConnectionPolicy = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinoteattribute?language=objc
type NoteAttribute uint8

const (
	KNoteAttributeManufacturerSpecific NoteAttribute = 1
	KNoteAttributeNone                 NoteAttribute = 0
	KNoteAttributePitch                NoteAttribute = 3
	KNoteAttributeProfileSpecific      NoteAttribute = 2
)

// The types of state changes the system supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinotificationmessageid?language=objc
type NotificationMessageID int32

const (
	KMsgIOError                NotificationMessageID = 7
	KMsgObjectAdded            NotificationMessageID = 2
	KMsgObjectRemoved          NotificationMessageID = 3
	KMsgPropertyChanged        NotificationMessageID = 4
	KMsgSerialPortOwnerChanged NotificationMessageID = 6
	KMsgSetupChanged           NotificationMessageID = 1
	KMsgThruConnectionsChanged NotificationMessageID = 5
)

// The common base class for many of the framework’s objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiobjectref?language=objc
type ObjectRef uint32

// The MIDI object types that the system supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiobjecttype?language=objc
type ObjectType int32

const (
	KObjectType_Destination         ObjectType = 3
	KObjectType_Device              ObjectType = 0
	KObjectType_Entity              ObjectType = 1
	KObjectType_ExternalDestination ObjectType = 19
	KObjectType_ExternalDevice      ObjectType = 16
	KObjectType_ExternalEntity      ObjectType = 17
	KObjectType_ExternalMask        ObjectType = 16
	KObjectType_ExternalSource      ObjectType = 18
	KObjectType_Other               ObjectType = -1
	KObjectType_Source              ObjectType = 2
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midipernotemanagementoptions?language=objc
type PerNoteManagementOptions uint8

const (
	KPerNoteManagementDetach PerNoteManagementOptions = 2
	KPerNoteManagementReset  PerNoteManagementOptions = 1
)

// A MIDI connection that a client maintains. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiportref?language=objc
type PortRef ObjectRef

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiprogramchangeoptions?language=objc
type ProgramChangeOptions uint8

const (
	KProgramChangeBankValid ProgramChangeOptions = 1
)

// Specifies a MIDI protocol variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiprotocolid?language=objc
type ProtocolID int32

const (
	KProtocol_1_0 ProtocolID = 1
	KProtocol_2_0 ProtocolID = 2
)

// A type that represents the global state of the MIDI system, that contains lists of the devices and serial port owners. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midisetupref?language=objc
type SetupRef ObjectRef

// MIDI System Exclusive (SysEx) types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midisysexstatus?language=objc
type SysExStatus int

const (
	KSysExStatusComplete            SysExStatus = 0
	KSysExStatusContinue            SysExStatus = 2
	KSysExStatusEnd                 SysExStatus = 3
	KSysExStatusMixedDataSetHeader  SysExStatus = 8
	KSysExStatusMixedDataSetPayload SysExStatus = 9
	KSysExStatusStart               SysExStatus = 1
)

// MIDI System status types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midisystemstatus?language=objc
type SystemStatus int

const (
	KStatusActiveSending    SystemStatus = 254
	KStatusActiveSensing    SystemStatus = 254
	KStatusContinue         SystemStatus = 251
	KStatusEndOfExclusive   SystemStatus = 247
	KStatusMTC              SystemStatus = 241
	KStatusSongPosPointer   SystemStatus = 242
	KStatusSongSelect       SystemStatus = 243
	KStatusStart            SystemStatus = 250
	KStatusStartOfExclusive SystemStatus = 240
	KStatusStop             SystemStatus = 252
	KStatusSystemReset      SystemStatus = 255
	KStatusTimingClock      SystemStatus = 248
	KStatusTuneRequest      SystemStatus = 246
)

// An opaque reference to a play-through connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midithruconnectionref?language=objc
type ThruConnectionRef ObjectRef

// The time on the host clock when the event occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/miditimestamp?language=objc
type TimeStamp uint64

// A set of values that indicate how to interpret control numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/miditransformcontroltype?language=objc
type TransformControlType uint8

const (
	KControlType_14Bit     TransformControlType = 1
	KControlType_14BitNRPN TransformControlType = 5
	KControlType_14BitRPN  TransformControlType = 3
	KControlType_7Bit      TransformControlType = 0
	KControlType_7BitNRPN  TransformControlType = 4
	KControlType_7BitRPN   TransformControlType = 2
)

// Values that specify the type of MIDI transformation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/miditransformtype?language=objc
type TransformType uint16

const (
	KTransform_Add        TransformType = 8
	KTransform_FilterOut  TransformType = 1
	KTransform_MapControl TransformType = 2
	KTransform_MapValue   TransformType = 12
	KTransform_MaxValue   TransformType = 11
	KTransform_MinValue   TransformType = 10
	KTransform_None       TransformType = 0
	KTransform_Scale      TransformType = 9
)

// A MIDI object’s unique identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiuniqueid?language=objc
type UniqueID int32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiutilitystatus?language=objc
type UtilityStatus int

const (
	KUtilityStatusJitterReductionClock     UtilityStatus = 1
	KUtilityStatusJitterReductionTimestamp UtilityStatus = 2
	KUtilityStatusNOOP                     UtilityStatus = 0
)
