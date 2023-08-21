// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitDispersion] class.
var UnitDispersionClass = _UnitDispersionClass{objc.GetClass("NSUnitDispersion")}

type _UnitDispersionClass struct {
	objc.Class
}

// An interface definition for the [UnitDispersion] class.
type IUnitDispersion interface {
	IDimension
}

// A unit of measure for specific quantities of dispersion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitdispersion?language=objc
type UnitDispersion struct {
	Dimension
}

func UnitDispersionFrom(ptr unsafe.Pointer) UnitDispersion {
	return UnitDispersion{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitDispersionClass) Alloc() UnitDispersion {
	rv := objc.Call[UnitDispersion](uc, objc.Sel("alloc"))
	return rv
}

func UnitDispersion_Alloc() UnitDispersion {
	return UnitDispersionClass.Alloc()
}

func (uc _UnitDispersionClass) New() UnitDispersion {
	rv := objc.Call[UnitDispersion](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitDispersion() UnitDispersion {
	return UnitDispersionClass.New()
}

func (u_ UnitDispersion) Init() UnitDispersion {
	rv := objc.Call[UnitDispersion](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitDispersionClass) BaseUnit() UnitDispersion {
	rv := objc.Call[UnitDispersion](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitDispersion_BaseUnit() UnitDispersion {
	return UnitDispersionClass.BaseUnit()
}

func (u_ UnitDispersion) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitDispersion {
	rv := objc.Call[UnitDispersion](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func NewUnitDispersionWithSymbolConverter(symbol string, converter IUnitConverter) UnitDispersion {
	instance := UnitDispersionClass.Alloc().InitWithSymbolConverter(symbol, converter)
	instance.Autorelease()
	return instance
}

func (u_ UnitDispersion) InitWithSymbol(symbol string) UnitDispersion {
	rv := objc.Call[UnitDispersion](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func NewUnitDispersionWithSymbol(symbol string) UnitDispersion {
	instance := UnitDispersionClass.Alloc().InitWithSymbol(symbol)
	instance.Autorelease()
	return instance
}

// The parts per million unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitdispersion/1690700-partspermillion?language=objc
func (uc _UnitDispersionClass) PartsPerMillion() UnitDispersion {
	rv := objc.Call[UnitDispersion](uc, objc.Sel("partsPerMillion"))
	return rv
}

// The parts per million unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitdispersion/1690700-partspermillion?language=objc
func UnitDispersion_PartsPerMillion() UnitDispersion {
	return UnitDispersionClass.PartsPerMillion()
}
