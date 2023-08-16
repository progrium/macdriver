// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersonNameComponents] class.
var PersonNameComponentsClass = _PersonNameComponentsClass{objc.GetClass("NSPersonNameComponents")}

type _PersonNameComponentsClass struct {
	objc.Class
}

// An interface definition for the [PersonNameComponents] class.
type IPersonNameComponents interface {
	objc.IObject
	PhoneticRepresentation() PersonNameComponents
	SetPhoneticRepresentation(value IPersonNameComponents)
	NamePrefix() string
	SetNamePrefix(value string)
	MiddleName() string
	SetMiddleName(value string)
	GivenName() string
	SetGivenName(value string)
	NameSuffix() string
	SetNameSuffix(value string)
	FamilyName() string
	SetFamilyName(value string)
	Nickname() string
	SetNickname(value string)
}

// An object that manages the separate parts of a person's name to allow locale-aware formatting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents?language=objc
type PersonNameComponents struct {
	objc.Object
}

func PersonNameComponentsFrom(ptr unsafe.Pointer) PersonNameComponents {
	return PersonNameComponents{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PersonNameComponentsClass) Alloc() PersonNameComponents {
	rv := objc.Call[PersonNameComponents](pc, objc.Sel("alloc"))
	return rv
}

func PersonNameComponents_Alloc() PersonNameComponents {
	return PersonNameComponentsClass.Alloc()
}

func (pc _PersonNameComponentsClass) New() PersonNameComponents {
	rv := objc.Call[PersonNameComponents](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersonNameComponents() PersonNameComponents {
	return PersonNameComponentsClass.New()
}

func (p_ PersonNameComponents) Init() PersonNameComponents {
	rv := objc.Call[PersonNameComponents](p_, objc.Sel("init"))
	return rv
}

// The phonetic representation name components of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1412193-phoneticrepresentation?language=objc
func (p_ PersonNameComponents) PhoneticRepresentation() PersonNameComponents {
	rv := objc.Call[PersonNameComponents](p_, objc.Sel("phoneticRepresentation"))
	return rv
}

// The phonetic representation name components of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1412193-phoneticrepresentation?language=objc
func (p_ PersonNameComponents) SetPhoneticRepresentation(value IPersonNameComponents) {
	objc.Call[objc.Void](p_, objc.Sel("setPhoneticRepresentation:"), objc.Ptr(value))
}

// The portion of a name’s full form of address that precedes the name itself (for example, “Dr.,” “Mr.,” “Ms.”). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1410275-nameprefix?language=objc
func (p_ PersonNameComponents) NamePrefix() string {
	rv := objc.Call[string](p_, objc.Sel("namePrefix"))
	return rv
}

// The portion of a name’s full form of address that precedes the name itself (for example, “Dr.,” “Mr.,” “Ms.”). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1410275-nameprefix?language=objc
func (p_ PersonNameComponents) SetNamePrefix(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setNamePrefix:"), value)
}

// Secondary name bestowed upon an individual to differentiate them from others that have the same given name (for example, “Maple”). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1418183-middlename?language=objc
func (p_ PersonNameComponents) MiddleName() string {
	rv := objc.Call[string](p_, objc.Sel("middleName"))
	return rv
}

// Secondary name bestowed upon an individual to differentiate them from others that have the same given name (for example, “Maple”). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1418183-middlename?language=objc
func (p_ PersonNameComponents) SetMiddleName(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setMiddleName:"), value)
}

// Name bestowed upon an individual to differentiate them from other members of a group that share a family name (for example, “Johnathan”). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1407259-givenname?language=objc
func (p_ PersonNameComponents) GivenName() string {
	rv := objc.Call[string](p_, objc.Sel("givenName"))
	return rv
}

// Name bestowed upon an individual to differentiate them from other members of a group that share a family name (for example, “Johnathan”). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1407259-givenname?language=objc
func (p_ PersonNameComponents) SetGivenName(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setGivenName:"), value)
}

// The portion of a name’s full form of address that follows the name itself (for example, “Esq.,” “Jr.,” “Ph.D.”). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1410776-namesuffix?language=objc
func (p_ PersonNameComponents) NameSuffix() string {
	rv := objc.Call[string](p_, objc.Sel("nameSuffix"))
	return rv
}

// The portion of a name’s full form of address that follows the name itself (for example, “Esq.,” “Jr.,” “Ph.D.”). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1410776-namesuffix?language=objc
func (p_ PersonNameComponents) SetNameSuffix(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setNameSuffix:"), value)
}

// Name bestowed upon an individual to denote membership in a group or family. (for example, “Appleseed”). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1413354-familyname?language=objc
func (p_ PersonNameComponents) FamilyName() string {
	rv := objc.Call[string](p_, objc.Sel("familyName"))
	return rv
}

// Name bestowed upon an individual to denote membership in a group or family. (for example, “Appleseed”). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1413354-familyname?language=objc
func (p_ PersonNameComponents) SetFamilyName(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setFamilyName:"), value)
}

// Name substituted for the purposes of familiarity (for example, "Johnny"). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1414892-nickname?language=objc
func (p_ PersonNameComponents) Nickname() string {
	rv := objc.Call[string](p_, objc.Sel("nickname"))
	return rv
}

// Name substituted for the purposes of familiarity (for example, "Johnny"). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspersonnamecomponents/1414892-nickname?language=objc
func (p_ PersonNameComponents) SetNickname(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setNickname:"), value)
}
