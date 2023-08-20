// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitAngle] class.
var UnitAngleClass = _UnitAngleClass{objc.GetClass("NSUnitAngle")}

type _UnitAngleClass struct {
	objc.Class
}

// An interface definition for the [UnitAngle] class.
type IUnitAngle interface {
	IDimension
}

// A unit of measure for planar angle and rotation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle?language=objc
type UnitAngle struct {
	Dimension
}

func UnitAngleFrom(ptr unsafe.Pointer) UnitAngle {
	return UnitAngle{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitAngleClass) Alloc() UnitAngle {
	rv := objc.Call[UnitAngle](uc, objc.Sel("alloc"))
	return rv
}

func UnitAngle_Alloc() UnitAngle {
	return UnitAngleClass.Alloc()
}

func (uc _UnitAngleClass) New() UnitAngle {
	rv := objc.Call[UnitAngle](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitAngle() UnitAngle {
	return UnitAngleClass.New()
}

func (u_ UnitAngle) Init() UnitAngle {
	rv := objc.Call[UnitAngle](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitAngleClass) BaseUnit() UnitAngle {
	rv := objc.Call[UnitAngle](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitAngle_BaseUnit() UnitAngle {
	return UnitAngleClass.BaseUnit()
}

func (u_ UnitAngle) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitAngle {
	rv := objc.Call[UnitAngle](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func UnitAngle_InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitAngle {
	return UnitAngleClass.Alloc().InitWithSymbolConverter(symbol, converter)
}

func (u_ UnitAngle) InitWithSymbol(symbol string) UnitAngle {
	rv := objc.Call[UnitAngle](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func UnitAngle_InitWithSymbol(symbol string) UnitAngle {
	return UnitAngleClass.Alloc().InitWithSymbol(symbol)
}

// The revolutions unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1855992-revolutions?language=objc
func (uc _UnitAngleClass) Revolutions() UnitAngle {
	rv := objc.Call[UnitAngle](uc, objc.Sel("revolutions"))
	return rv
}

// The revolutions unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1855992-revolutions?language=objc
func UnitAngle_Revolutions() UnitAngle {
	return UnitAngleClass.Revolutions()
}

// The arc seconds unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1856114-arcseconds?language=objc
func (uc _UnitAngleClass) ArcSeconds() UnitAngle {
	rv := objc.Call[UnitAngle](uc, objc.Sel("arcSeconds"))
	return rv
}

// The arc seconds unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1856114-arcseconds?language=objc
func UnitAngle_ArcSeconds() UnitAngle {
	return UnitAngleClass.ArcSeconds()
}

// The arc minutes unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1856069-arcminutes?language=objc
func (uc _UnitAngleClass) ArcMinutes() UnitAngle {
	rv := objc.Call[UnitAngle](uc, objc.Sel("arcMinutes"))
	return rv
}

// The arc minutes unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1856069-arcminutes?language=objc
func UnitAngle_ArcMinutes() UnitAngle {
	return UnitAngleClass.ArcMinutes()
}

// The gradians unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1855988-gradians?language=objc
func (uc _UnitAngleClass) Gradians() UnitAngle {
	rv := objc.Call[UnitAngle](uc, objc.Sel("gradians"))
	return rv
}

// The gradians unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1855988-gradians?language=objc
func UnitAngle_Gradians() UnitAngle {
	return UnitAngleClass.Gradians()
}

// The radians unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1856062-radians?language=objc
func (uc _UnitAngleClass) Radians() UnitAngle {
	rv := objc.Call[UnitAngle](uc, objc.Sel("radians"))
	return rv
}

// The radians unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1856062-radians?language=objc
func UnitAngle_Radians() UnitAngle {
	return UnitAngleClass.Radians()
}

// The degrees unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1856083-degrees?language=objc
func (uc _UnitAngleClass) Degrees() UnitAngle {
	rv := objc.Call[UnitAngle](uc, objc.Sel("degrees"))
	return rv
}

// The degrees unit of angle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitangle/1856083-degrees?language=objc
func UnitAngle_Degrees() UnitAngle {
	return UnitAngleClass.Degrees()
}