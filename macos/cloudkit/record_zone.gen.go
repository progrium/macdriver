// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecordZone] class.
var RecordZoneClass = _RecordZoneClass{objc.GetClass("CKRecordZone")}

type _RecordZoneClass struct {
	objc.Class
}

// An interface definition for the [RecordZone] class.
type IRecordZone interface {
	objc.IObject
	Share() Reference
	Capabilities() RecordZoneCapabilities
	ZoneID() RecordZoneID
}

// A database partition that contains related records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone?language=objc
type RecordZone struct {
	objc.Object
}

func RecordZoneFrom(ptr unsafe.Pointer) RecordZone {
	return RecordZone{
		Object: objc.ObjectFrom(ptr),
	}
}

func (r_ RecordZone) InitWithZoneName(zoneName string) RecordZone {
	rv := objc.Call[RecordZone](r_, objc.Sel("initWithZoneName:"), zoneName)
	return rv
}

// Creates a record zone object with the specified zone name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone/1515102-initwithzonename?language=objc
func RecordZone_InitWithZoneName(zoneName string) RecordZone {
	return RecordZoneClass.Alloc().InitWithZoneName(zoneName)
}

func (r_ RecordZone) InitWithZoneID(zoneID IRecordZoneID) RecordZone {
	rv := objc.Call[RecordZone](r_, objc.Sel("initWithZoneID:"), objc.Ptr(zoneID))
	return rv
}

// Creates a record zone object with the specified zone ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone/1515207-initwithzoneid?language=objc
func RecordZone_InitWithZoneID(zoneID IRecordZoneID) RecordZone {
	return RecordZoneClass.Alloc().InitWithZoneID(zoneID)
}

func (rc _RecordZoneClass) Alloc() RecordZone {
	rv := objc.Call[RecordZone](rc, objc.Sel("alloc"))
	return rv
}

func RecordZone_Alloc() RecordZone {
	return RecordZoneClass.Alloc()
}

func (rc _RecordZoneClass) New() RecordZone {
	rv := objc.Call[RecordZone](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecordZone() RecordZone {
	return RecordZoneClass.New()
}

func (r_ RecordZone) Init() RecordZone {
	rv := objc.Call[RecordZone](r_, objc.Sel("init"))
	return rv
}

// Returns the default record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone/1514919-defaultrecordzone?language=objc
func (rc _RecordZoneClass) DefaultRecordZone() RecordZone {
	rv := objc.Call[RecordZone](rc, objc.Sel("defaultRecordZone"))
	return rv
}

// Returns the default record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone/1514919-defaultrecordzone?language=objc
func RecordZone_DefaultRecordZone() RecordZone {
	return RecordZoneClass.DefaultRecordZone()
}

// A reference to the record zone’s share record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone/3746822-share?language=objc
func (r_ RecordZone) Share() Reference {
	rv := objc.Call[Reference](r_, objc.Sel("share"))
	return rv
}

// The capabilities that the zone supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone/1515194-capabilities?language=objc
func (r_ RecordZone) Capabilities() RecordZoneCapabilities {
	rv := objc.Call[RecordZoneCapabilities](r_, objc.Sel("capabilities"))
	return rv
}

// The unique ID of the zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzone/1514917-zoneid?language=objc
func (r_ RecordZone) ZoneID() RecordZoneID {
	rv := objc.Call[RecordZoneID](r_, objc.Sel("zoneID"))
	return rv
}
