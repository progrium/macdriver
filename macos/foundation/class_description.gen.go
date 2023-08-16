// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ClassDescription] class.
var ClassDescriptionClass = _ClassDescriptionClass{objc.GetClass("NSClassDescription")}

type _ClassDescriptionClass struct {
	objc.Class
}

// An interface definition for the [ClassDescription] class.
type IClassDescription interface {
	objc.IObject
	InverseForRelationshipKey(relationshipKey string) string
	AttributeKeys() []string
	ToManyRelationshipKeys() []string
	ToOneRelationshipKeys() []string
}

// An abstract class that provides the interface for querying the relationships and properties of a class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclassdescription?language=objc
type ClassDescription struct {
	objc.Object
}

func ClassDescriptionFrom(ptr unsafe.Pointer) ClassDescription {
	return ClassDescription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ClassDescriptionClass) Alloc() ClassDescription {
	rv := objc.Call[ClassDescription](cc, objc.Sel("alloc"))
	return rv
}

func ClassDescription_Alloc() ClassDescription {
	return ClassDescriptionClass.Alloc()
}

func (cc _ClassDescriptionClass) New() ClassDescription {
	rv := objc.Call[ClassDescription](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewClassDescription() ClassDescription {
	return ClassDescriptionClass.New()
}

func (c_ ClassDescription) Init() ClassDescription {
	rv := objc.Call[ClassDescription](c_, objc.Sel("init"))
	return rv
}

// Registers an NSClassDescription object for a given class in the NSClassDescription cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclassdescription/1407444-registerclassdescription?language=objc
func (cc _ClassDescriptionClass) RegisterClassDescriptionForClass(description IClassDescription, aClass objc.IClass) {
	objc.Call[objc.Void](cc, objc.Sel("registerClassDescription:forClass:"), objc.Ptr(description), objc.Ptr(aClass))
}

// Registers an NSClassDescription object for a given class in the NSClassDescription cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclassdescription/1407444-registerclassdescription?language=objc
func ClassDescription_RegisterClassDescriptionForClass(description IClassDescription, aClass objc.IClass) {
	ClassDescriptionClass.RegisterClassDescriptionForClass(description, aClass)
}

// Removes all NSClassDescription objects from the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclassdescription/1413040-invalidateclassdescriptioncache?language=objc
func (cc _ClassDescriptionClass) InvalidateClassDescriptionCache() {
	objc.Call[objc.Void](cc, objc.Sel("invalidateClassDescriptionCache"))
}

// Removes all NSClassDescription objects from the cache. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclassdescription/1413040-invalidateclassdescriptioncache?language=objc
func ClassDescription_InvalidateClassDescriptionCache() {
	ClassDescriptionClass.InvalidateClassDescriptionCache()
}

// Overridden by subclasses to return the name of the inverse relationship from a relationship specified by a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclassdescription/1410945-inverseforrelationshipkey?language=objc
func (c_ ClassDescription) InverseForRelationshipKey(relationshipKey string) string {
	rv := objc.Call[string](c_, objc.Sel("inverseForRelationshipKey:"), relationshipKey)
	return rv
}

// Returns the class description for a given class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclassdescription/1413026-classdescriptionforclass?language=objc
func (cc _ClassDescriptionClass) ClassDescriptionForClass(aClass objc.IClass) ClassDescription {
	rv := objc.Call[ClassDescription](cc, objc.Sel("classDescriptionForClass:"), objc.Ptr(aClass))
	return rv
}

// Returns the class description for a given class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclassdescription/1413026-classdescriptionforclass?language=objc
func ClassDescription_ClassDescriptionForClass(aClass objc.IClass) ClassDescription {
	return ClassDescriptionClass.ClassDescriptionForClass(aClass)
}

// Overridden by subclasses to return the names of attributes of instances of the described class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclassdescription/1415001-attributekeys?language=objc
func (c_ ClassDescription) AttributeKeys() []string {
	rv := objc.Call[[]string](c_, objc.Sel("attributeKeys"))
	return rv
}

// Overridden by subclasses to return the keys for the to-many relationship properties of instances of the described class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclassdescription/1408530-tomanyrelationshipkeys?language=objc
func (c_ ClassDescription) ToManyRelationshipKeys() []string {
	rv := objc.Call[[]string](c_, objc.Sel("toManyRelationshipKeys"))
	return rv
}

// Overridden by subclasses to return the keys for the to-one relationship properties of instances of the described class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclassdescription/1411937-toonerelationshipkeys?language=objc
func (c_ ClassDescription) ToOneRelationshipKeys() []string {
	rv := objc.Call[[]string](c_, objc.Sel("toOneRelationshipKeys"))
	return rv
}
