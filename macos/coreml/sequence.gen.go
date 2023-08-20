// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Sequence] class.
var SequenceClass = _SequenceClass{objc.GetClass("MLSequence")}

type _SequenceClass struct {
	objc.Class
}

// An interface definition for the [Sequence] class.
type ISequence interface {
	objc.IObject
	StringValues() []string
	Int64Values() []foundation.Number
	Type() FeatureType
}

// A machine learning collection type that stores a series of strings or integers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence?language=objc
type Sequence struct {
	objc.Object
}

func SequenceFrom(ptr unsafe.Pointer) Sequence {
	return Sequence{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SequenceClass) SequenceWithStringArray(stringValues []string) Sequence {
	rv := objc.Call[Sequence](sc, objc.Sel("sequenceWithStringArray:"), stringValues)
	return rv
}

// Creates a sequence of strings from a string array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/2962873-sequencewithstringarray?language=objc
func Sequence_SequenceWithStringArray(stringValues []string) Sequence {
	return SequenceClass.SequenceWithStringArray(stringValues)
}

func (sc _SequenceClass) SequenceWithInt64Array(int64Values []foundation.INumber) Sequence {
	rv := objc.Call[Sequence](sc, objc.Sel("sequenceWithInt64Array:"), int64Values)
	return rv
}

// Creates a sequence of integers from an array of numbers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/2962872-sequencewithint64array?language=objc
func Sequence_SequenceWithInt64Array(int64Values []foundation.INumber) Sequence {
	return SequenceClass.SequenceWithInt64Array(int64Values)
}

func (sc _SequenceClass) EmptySequenceWithType(type_ FeatureType) Sequence {
	rv := objc.Call[Sequence](sc, objc.Sel("emptySequenceWithType:"), type_)
	return rv
}

// Creates an empty sequence of strings or integers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/2962870-emptysequencewithtype?language=objc
func Sequence_EmptySequenceWithType(type_ FeatureType) Sequence {
	return SequenceClass.EmptySequenceWithType(type_)
}

func (sc _SequenceClass) Alloc() Sequence {
	rv := objc.Call[Sequence](sc, objc.Sel("alloc"))
	return rv
}

func Sequence_Alloc() Sequence {
	return SequenceClass.Alloc()
}

func (sc _SequenceClass) New() Sequence {
	rv := objc.Call[Sequence](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSequence() Sequence {
	return SequenceClass.New()
}

func (s_ Sequence) Init() Sequence {
	rv := objc.Call[Sequence](s_, objc.Sel("init"))
	return rv
}

// An array of strings in the sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/2962874-stringvalues?language=objc
func (s_ Sequence) StringValues() []string {
	rv := objc.Call[[]string](s_, objc.Sel("stringValues"))
	return rv
}

// An array of 64-bit integers in the sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/2962871-int64values?language=objc
func (s_ Sequence) Int64Values() []foundation.Number {
	rv := objc.Call[[]foundation.Number](s_, objc.Sel("int64Values"))
	return rv
}

// The underlying type of the sequence’s elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/2962875-type?language=objc
func (s_ Sequence) Type() FeatureType {
	rv := objc.Call[FeatureType](s_, objc.Sel("type"))
	return rv
}