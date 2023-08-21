// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ArrayController] class.
var ArrayControllerClass = _ArrayControllerClass{objc.GetClass("NSArrayController")}

type _ArrayControllerClass struct {
	objc.Class
}

// An interface definition for the [ArrayController] class.
type IArrayController interface {
	IObjectController
	RemoveObjectAtArrangedObjectIndex(index uint)
	SetSelectionIndex(index uint) bool
	RemoveSelectionIndexes(indexes foundation.IIndexSet) bool
	SetSelectedObjects(objects []objc.IObject) bool
	AddSelectionIndexes(indexes foundation.IIndexSet) bool
	InsertObjectAtArrangedObjectIndex(object objc.IObject, index uint)
	SetSelectionIndexes(indexes foundation.IIndexSet) bool
	RemoveObjectsAtArrangedObjectIndexes(indexes foundation.IIndexSet)
	DidChangeArrangementCriteria()
	RearrangeObjects()
	SelectPrevious(sender objc.IObject) objc.Object
	RemoveObjects(objects []objc.IObject)
	AddSelectedObjects(objects []objc.IObject) bool
	ArrangeObjects(objects []objc.IObject) []objc.Object
	AddObjects(objects []objc.IObject)
	RemoveSelectedObjects(objects []objc.IObject) bool
	SelectNext(sender objc.IObject) objc.Object
	Insert(sender objc.IObject) objc.Object
	InsertObjectsAtArrangedObjectIndexes(objects []objc.IObject, indexes foundation.IIndexSet)
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)
	FilterPredicate() foundation.Predicate
	SetFilterPredicate(value foundation.IPredicate)
	AutomaticRearrangementKeyPaths() []string
	CanSelectPrevious() bool
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.ISortDescriptor)
	ArrangedObjects() objc.Object
	SelectionIndexes() foundation.IndexSet
	CanInsert() bool
	CanSelectNext() bool
	SelectionIndex() uint
	AutomaticallyRearrangesObjects() bool
	SetAutomaticallyRearrangesObjects(value bool)
	ClearsFilterPredicateOnInsertion() bool
	SetClearsFilterPredicateOnInsertion(value bool)
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
}

// A bindings-compatible controller that manages a collection of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller?language=objc
type ArrayController struct {
	ObjectController
}

func ArrayControllerFrom(ptr unsafe.Pointer) ArrayController {
	return ArrayController{
		ObjectController: ObjectControllerFrom(ptr),
	}
}

func (ac _ArrayControllerClass) Alloc() ArrayController {
	rv := objc.Call[ArrayController](ac, objc.Sel("alloc"))
	return rv
}

func ArrayController_Alloc() ArrayController {
	return ArrayControllerClass.Alloc()
}

func (ac _ArrayControllerClass) New() ArrayController {
	rv := objc.Call[ArrayController](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewArrayController() ArrayController {
	return ArrayControllerClass.New()
}

func (a_ ArrayController) Init() ArrayController {
	rv := objc.Call[ArrayController](a_, objc.Sel("init"))
	return rv
}

func (a_ ArrayController) InitWithContent(content objc.IObject) ArrayController {
	rv := objc.Call[ArrayController](a_, objc.Sel("initWithContent:"), content)
	return rv
}

// Initializes and returns an NSObjectController object with the given content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1529422-initwithcontent?language=objc
func NewArrayControllerWithContent(content objc.IObject) ArrayController {
	instance := ArrayControllerClass.Alloc().InitWithContent(content)
	instance.Autorelease()
	return instance
}

// Removes the object at the specified index in the receiver’s arranged objects from the receiver’s content array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1534738-removeobjectatarrangedobjectinde?language=objc
func (a_ ArrayController) RemoveObjectAtArrangedObjectIndex(index uint) {
	objc.Call[objc.Void](a_, objc.Sel("removeObjectAtArrangedObjectIndex:"), index)
}

// Sets the receiver’s selection to the given index, and returns a Boolean value that indicates whether the selection was changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1526094-setselectionindex?language=objc
func (a_ ArrayController) SetSelectionIndex(index uint) bool {
	rv := objc.Call[bool](a_, objc.Sel("setSelectionIndex:"), index)
	return rv
}

// Removes the object as the specified indexes from the receiver’s current selection, returning YES if the selection was changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1525061-removeselectionindexes?language=objc
func (a_ ArrayController) RemoveSelectionIndexes(indexes foundation.IIndexSet) bool {
	rv := objc.Call[bool](a_, objc.Sel("removeSelectionIndexes:"), objc.Ptr(indexes))
	return rv
}

// Sets objects as the receiver’s current selection, returning YES if the selection was changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1533937-setselectedobjects?language=objc
func (a_ ArrayController) SetSelectedObjects(objects []objc.IObject) bool {
	rv := objc.Call[bool](a_, objc.Sel("setSelectedObjects:"), objects)
	return rv
}

// Adds the objects at the specified indexes in the receiver’s content array to the current selection, returning YES if the selection was changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1533997-addselectionindexes?language=objc
func (a_ ArrayController) AddSelectionIndexes(indexes foundation.IIndexSet) bool {
	rv := objc.Call[bool](a_, objc.Sel("addSelectionIndexes:"), objc.Ptr(indexes))
	return rv
}

// Inserts object into the receiver’s arranged objects array at the location specified by index, and adds it to the receiver’s content collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1532302-insertobject?language=objc
func (a_ ArrayController) InsertObjectAtArrangedObjectIndex(object objc.IObject, index uint) {
	objc.Call[objc.Void](a_, objc.Sel("insertObject:atArrangedObjectIndex:"), object, index)
}

// Sets the receiver’s selection indexes and returns a Boolean value that indicates whether the selection changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1525801-setselectionindexes?language=objc
func (a_ ArrayController) SetSelectionIndexes(indexes foundation.IIndexSet) bool {
	rv := objc.Call[bool](a_, objc.Sel("setSelectionIndexes:"), objc.Ptr(indexes))
	return rv
}

// Removes the objects at the specified indexes in the receiver’s arranged objects from the content array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1533508-removeobjectsatarrangedobjectind?language=objc
func (a_ ArrayController) RemoveObjectsAtArrangedObjectIndexes(indexes foundation.IIndexSet) {
	objc.Call[objc.Void](a_, objc.Sel("removeObjectsAtArrangedObjectIndexes:"), objc.Ptr(indexes))
}

// Invoked when any criteria for arranging objects change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1526171-didchangearrangementcriteria?language=objc
func (a_ ArrayController) DidChangeArrangementCriteria() {
	objc.Call[objc.Void](a_, objc.Sel("didChangeArrangementCriteria"))
}

// Triggers filtering of the receiver’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1534319-rearrangeobjects?language=objc
func (a_ ArrayController) RearrangeObjects() {
	objc.Call[objc.Void](a_, objc.Sel("rearrangeObjects"))
}

// Selects the previous object, relative to the current selection, in the receiver’s arranged content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1528448-selectprevious?language=objc
func (a_ ArrayController) SelectPrevious(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("selectPrevious:"), sender)
	return rv
}

// Removes objects from the receiver’s content collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1533426-removeobjects?language=objc
func (a_ ArrayController) RemoveObjects(objects []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("removeObjects:"), objects)
}

// Adds objects from the receiver’s content array to the current selection, returning YES if the selection was changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1533846-addselectedobjects?language=objc
func (a_ ArrayController) AddSelectedObjects(objects []objc.IObject) bool {
	rv := objc.Call[bool](a_, objc.Sel("addSelectedObjects:"), objects)
	return rv
}

// Returns a given array, appropriately sorted and filtered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1533881-arrangeobjects?language=objc
func (a_ ArrayController) ArrangeObjects(objects []objc.IObject) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("arrangeObjects:"), objects)
	return rv
}

// Adds objects to the receiver’s content collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1533977-addobjects?language=objc
func (a_ ArrayController) AddObjects(objects []objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("addObjects:"), objects)
}

// Removes objects from the receiver’s current selection, returning YES if the selection was changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1532684-removeselectedobjects?language=objc
func (a_ ArrayController) RemoveSelectedObjects(objects []objc.IObject) bool {
	rv := objc.Call[bool](a_, objc.Sel("removeSelectedObjects:"), objects)
	return rv
}

// Selects the next object, relative to the current selection, in the receiver’s arranged content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1525178-selectnext?language=objc
func (a_ ArrayController) SelectNext(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("selectNext:"), sender)
	return rv
}

// Creates a new object and inserts it into the receiver’s content array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1535089-insert?language=objc
func (a_ ArrayController) Insert(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("insert:"), sender)
	return rv
}

// Inserts objects into the receiver’s arranged objects array at the locations specified in indexes, and adds it to the receiver’s content collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1527973-insertobjects?language=objc
func (a_ ArrayController) InsertObjectsAtArrangedObjectIndexes(objects []objc.IObject, indexes foundation.IIndexSet) {
	objc.Call[objc.Void](a_, objc.Sel("insertObjects:atArrangedObjectIndexes:"), objects, objc.Ptr(indexes))
}

// A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1530628-preservesselection?language=objc
func (a_ ArrayController) PreservesSelection() bool {
	rv := objc.Call[bool](a_, objc.Sel("preservesSelection"))
	return rv
}

// A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1530628-preservesselection?language=objc
func (a_ ArrayController) SetPreservesSelection(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setPreservesSelection:"), value)
}

// A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1527129-alwaysusesmultiplevaluesmarker?language=objc
func (a_ ArrayController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.Call[bool](a_, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}

// A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1527129-alwaysusesmultiplevaluesmarker?language=objc
func (a_ ArrayController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAlwaysUsesMultipleValuesMarker:"), value)
}

// A predicate used by the receiver to filter the array controller contents [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1524815-filterpredicate?language=objc
func (a_ ArrayController) FilterPredicate() foundation.Predicate {
	rv := objc.Call[foundation.Predicate](a_, objc.Sel("filterPredicate"))
	return rv
}

// A predicate used by the receiver to filter the array controller contents [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1524815-filterpredicate?language=objc
func (a_ ArrayController) SetFilterPredicate(value foundation.IPredicate) {
	objc.Call[objc.Void](a_, objc.Sel("setFilterPredicate:"), objc.Ptr(value))
}

// An array of key paths that trigger automatic content sorting or filtering [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1526714-automaticrearrangementkeypaths?language=objc
func (a_ ArrayController) AutomaticRearrangementKeyPaths() []string {
	rv := objc.Call[[]string](a_, objc.Sel("automaticRearrangementKeyPaths"))
	return rv
}

// A Boolean value indicating whether the previous object, relative to the current selection, in the receiver’s content array can be selected [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1534274-canselectprevious?language=objc
func (a_ ArrayController) CanSelectPrevious() bool {
	rv := objc.Call[bool](a_, objc.Sel("canSelectPrevious"))
	return rv
}

// An array of NSSortDescriptor objects, used by the receiver to arrange its content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1525707-sortdescriptors?language=objc
func (a_ ArrayController) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.Call[[]foundation.SortDescriptor](a_, objc.Sel("sortDescriptors"))
	return rv
}

// An array of NSSortDescriptor objects, used by the receiver to arrange its content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1525707-sortdescriptors?language=objc
func (a_ ArrayController) SetSortDescriptors(value []foundation.ISortDescriptor) {
	objc.Call[objc.Void](a_, objc.Sel("setSortDescriptors:"), value)
}

// An array containing the receiver’s content objects arranged using arrangeObjects:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1534299-arrangedobjects?language=objc
func (a_ ArrayController) ArrangedObjects() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("arrangedObjects"))
	return rv
}

// An index set containing the indexes of the receiver’s currently selected objects in the content array [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1529908-selectionindexes?language=objc
func (a_ ArrayController) SelectionIndexes() foundation.IndexSet {
	rv := objc.Call[foundation.IndexSet](a_, objc.Sel("selectionIndexes"))
	return rv
}

// Returns a Boolean value that indicates whether an object can be inserted into the receiver’s content collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1532900-caninsert?language=objc
func (a_ ArrayController) CanInsert() bool {
	rv := objc.Call[bool](a_, objc.Sel("canInsert"))
	return rv
}

// A Boolean value indicating whether the next object, relative to the current selection, in the receiver’s content array can be selected [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1534863-canselectnext?language=objc
func (a_ ArrayController) CanSelectNext() bool {
	rv := objc.Call[bool](a_, objc.Sel("canSelectNext"))
	return rv
}

// The index of the first object in the receiver’s selection [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1535885-selectionindex?language=objc
func (a_ ArrayController) SelectionIndex() uint {
	rv := objc.Call[uint](a_, objc.Sel("selectionIndex"))
	return rv
}

// A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1524747-automaticallyrearrangesobjects?language=objc
func (a_ ArrayController) AutomaticallyRearrangesObjects() bool {
	rv := objc.Call[bool](a_, objc.Sel("automaticallyRearrangesObjects"))
	return rv
}

// A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1524747-automaticallyrearrangesobjects?language=objc
func (a_ ArrayController) SetAutomaticallyRearrangesObjects(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAutomaticallyRearrangesObjects:"), value)
}

// A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1534701-clearsfilterpredicateoninsertion?language=objc
func (a_ ArrayController) ClearsFilterPredicateOnInsertion() bool {
	rv := objc.Call[bool](a_, objc.Sel("clearsFilterPredicateOnInsertion"))
	return rv
}

// A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1534701-clearsfilterpredicateoninsertion?language=objc
func (a_ ArrayController) SetClearsFilterPredicateOnInsertion(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setClearsFilterPredicateOnInsertion:"), value)
}

// A Boolean value that indicates whether the receiver automatically selects inserted objects [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1527427-selectsinsertedobjects?language=objc
func (a_ ArrayController) SelectsInsertedObjects() bool {
	rv := objc.Call[bool](a_, objc.Sel("selectsInsertedObjects"))
	return rv
}

// A Boolean value that indicates whether the receiver automatically selects inserted objects [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1527427-selectsinsertedobjects?language=objc
func (a_ ArrayController) SetSelectsInsertedObjects(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setSelectsInsertedObjects:"), value)
}

// A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1532785-avoidsemptyselection?language=objc
func (a_ ArrayController) AvoidsEmptySelection() bool {
	rv := objc.Call[bool](a_, objc.Sel("avoidsEmptySelection"))
	return rv
}

// A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/1532785-avoidsemptyselection?language=objc
func (a_ ArrayController) SetAvoidsEmptySelection(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAvoidsEmptySelection:"), value)
}
