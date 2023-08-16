// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [IndexPath] class.
var IndexPathClass = _IndexPathClass{objc.GetClass("NSIndexPath")}

type _IndexPathClass struct {
	objc.Class
}

// An interface definition for the [IndexPath] class.
type IIndexPath interface {
	objc.IObject
	IndexPathByAddingIndex(index uint) IndexPath
	IndexAtPosition(position uint) uint
	IndexPathByRemovingLastIndex() IndexPath
	Compare(otherObject IIndexPath) ComparisonResult
	Item() int
	Section() int
	Length() uint
}

// A list of indexes that together represent the path to a specific location in a tree of nested arrays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath?language=objc
type IndexPath struct {
	objc.Object
}

func IndexPathFrom(ptr unsafe.Pointer) IndexPath {
	return IndexPath{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _IndexPathClass) IndexPathWithIndexesLength(indexes *uint, length uint) IndexPath {
	rv := objc.Call[IndexPath](ic, objc.Sel("indexPathWithIndexes:length:"), indexes, length)
	return rv
}

// Creates an index path with one or more nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1521015-indexpathwithindexes?language=objc
func IndexPath_IndexPathWithIndexesLength(indexes *uint, length uint) IndexPath {
	return IndexPathClass.IndexPathWithIndexesLength(indexes, length)
}

func (i_ IndexPath) InitWithIndexesLength(indexes *uint, length uint) IndexPath {
	rv := objc.Call[IndexPath](i_, objc.Sel("initWithIndexes:length:"), indexes, length)
	return rv
}

// Initializes an index path with the given nodes and length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1416906-initwithindexes?language=objc
func IndexPath_InitWithIndexesLength(indexes *uint, length uint) IndexPath {
	return IndexPathClass.Alloc().InitWithIndexesLength(indexes, length)
}

func (ic _IndexPathClass) IndexPathForItemInSection(item int, section int) IndexPath {
	rv := objc.Call[IndexPath](ic, objc.Sel("indexPathForItem:inSection:"), item, section)
	return rv
}

// Initializes an index path with the indexes of a specific item and section in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1526053-indexpathforitem?language=objc
func IndexPath_IndexPathForItemInSection(item int, section int) IndexPath {
	return IndexPathClass.IndexPathForItemInSection(item, section)
}

func (ic _IndexPathClass) IndexPathWithIndex(index uint) IndexPath {
	rv := objc.Call[IndexPath](ic, objc.Sel("indexPathWithIndex:"), index)
	return rv
}

// Creates a one-node index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1521019-indexpathwithindex?language=objc
func IndexPath_IndexPathWithIndex(index uint) IndexPath {
	return IndexPathClass.IndexPathWithIndex(index)
}

func (i_ IndexPath) InitWithIndex(index uint) IndexPath {
	rv := objc.Call[IndexPath](i_, objc.Sel("initWithIndex:"), index)
	return rv
}

// Initializes an index path with a single node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1416855-initwithindex?language=objc
func IndexPath_InitWithIndex(index uint) IndexPath {
	return IndexPathClass.Alloc().InitWithIndex(index)
}

func (ic _IndexPathClass) Alloc() IndexPath {
	rv := objc.Call[IndexPath](ic, objc.Sel("alloc"))
	return rv
}

func IndexPath_Alloc() IndexPath {
	return IndexPathClass.Alloc()
}

func (ic _IndexPathClass) New() IndexPath {
	rv := objc.Call[IndexPath](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewIndexPath() IndexPath {
	return IndexPathClass.New()
}

func (i_ IndexPath) Init() IndexPath {
	rv := objc.Call[IndexPath](i_, objc.Sel("init"))
	return rv
}

// Returns an index path containing the nodes in the receiving index path plus another given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1417553-indexpathbyaddingindex?language=objc
func (i_ IndexPath) IndexPathByAddingIndex(index uint) IndexPath {
	rv := objc.Call[IndexPath](i_, objc.Sel("indexPathByAddingIndex:"), index)
	return rv
}

// Provides the value at a particular node in the index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1414179-indexatposition?language=objc
func (i_ IndexPath) IndexAtPosition(position uint) uint {
	rv := objc.Call[uint](i_, objc.Sel("indexAtPosition:"), position)
	return rv
}

// Returns an index path with the nodes in the receiving index path, excluding the last one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1415624-indexpathbyremovinglastindex?language=objc
func (i_ IndexPath) IndexPathByRemovingLastIndex() IndexPath {
	rv := objc.Call[IndexPath](i_, objc.Sel("indexPathByRemovingLastIndex"))
	return rv
}

// Indicates the depth-first traversal order of the receiving index path and another index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1407552-compare?language=objc
func (i_ IndexPath) Compare(otherObject IIndexPath) ComparisonResult {
	rv := objc.Call[ComparisonResult](i_, objc.Sel("compare:"), objc.Ptr(otherObject))
	return rv
}

// An index number identifying an item in a section of a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1528293-item?language=objc
func (i_ IndexPath) Item() int {
	rv := objc.Call[int](i_, objc.Sel("item"))
	return rv
}

// An index number identifying a section in a table view or collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1528298-section?language=objc
func (i_ IndexPath) Section() int {
	rv := objc.Call[int](i_, objc.Sel("section"))
	return rv
}

// The number of nodes in the index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexpath/1412001-length?language=objc
func (i_ IndexPath) Length() uint {
	rv := objc.Call[uint](i_, objc.Sel("length"))
	return rv
}
