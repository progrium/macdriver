// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ToolbarItemGroup] class.
var ToolbarItemGroupClass = _ToolbarItemGroupClass{objc.GetClass("NSToolbarItemGroup")}

type _ToolbarItemGroupClass struct {
	objc.Class
}

// An interface definition for the [ToolbarItemGroup] class.
type IToolbarItemGroup interface {
	IToolbarItem
	SetSelectedAtIndex(selected bool, index int)
	IsSelectedAtIndex(index int) bool
	SelectionMode() ToolbarItemGroupSelectionMode
	SetSelectionMode(value ToolbarItemGroupSelectionMode)
	Subitems() []ToolbarItem
	SetSubitems(value []IToolbarItem)
	ControlRepresentation() ToolbarItemGroupControlRepresentation
	SetControlRepresentation(value ToolbarItemGroupControlRepresentation)
	SelectedIndex() int
	SetSelectedIndex(value int)
}

// A group of subitems in a toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup?language=objc
type ToolbarItemGroup struct {
	ToolbarItem
}

func ToolbarItemGroupFrom(ptr unsafe.Pointer) ToolbarItemGroup {
	return ToolbarItemGroup{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}

func (tc _ToolbarItemGroupClass) GroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction(itemIdentifier ToolbarItemIdentifier, titles []string, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.Selector) ToolbarItemGroup {
	rv := objc.Call[ToolbarItemGroup](tc, objc.Sel("groupWithItemIdentifier:titles:selectionMode:labels:target:action:"), itemIdentifier, titles, selectionMode, labels, target, action)
	return rv
}

// Creates a grouped toolbar item with labels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup/3242764-groupwithitemidentifier?language=objc
func ToolbarItemGroup_GroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction(itemIdentifier ToolbarItemIdentifier, titles []string, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.Selector) ToolbarItemGroup {
	return ToolbarItemGroupClass.GroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction(itemIdentifier, titles, selectionMode, labels, target, action)
}

func (tc _ToolbarItemGroupClass) Alloc() ToolbarItemGroup {
	rv := objc.Call[ToolbarItemGroup](tc, objc.Sel("alloc"))
	return rv
}

func ToolbarItemGroup_Alloc() ToolbarItemGroup {
	return ToolbarItemGroupClass.Alloc()
}

func (tc _ToolbarItemGroupClass) New() ToolbarItemGroup {
	rv := objc.Call[ToolbarItemGroup](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewToolbarItemGroup() ToolbarItemGroup {
	return ToolbarItemGroupClass.New()
}

func (t_ ToolbarItemGroup) Init() ToolbarItemGroup {
	rv := objc.Call[ToolbarItemGroup](t_, objc.Sel("init"))
	return rv
}

func (t_ ToolbarItemGroup) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) ToolbarItemGroup {
	rv := objc.Call[ToolbarItemGroup](t_, objc.Sel("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

// Creates a toolbar item with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/1534084-initwithitemidentifier?language=objc
func ToolbarItemGroup_InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) ToolbarItemGroup {
	return ToolbarItemGroupClass.Alloc().InitWithItemIdentifier(itemIdentifier)
}

// Sets the selected state of a subitem in a grouped toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup/3242768-setselected?language=objc
func (t_ ToolbarItemGroup) SetSelectedAtIndex(selected bool, index int) {
	objc.Call[objc.Void](t_, objc.Sel("setSelected:atIndex:"), selected, index)
}

// Indicates whether a specified index is currently selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup/3242765-isselectedatindex?language=objc
func (t_ ToolbarItemGroup) IsSelectedAtIndex(index int) bool {
	rv := objc.Call[bool](t_, objc.Sel("isSelectedAtIndex:"), index)
	return rv
}

// The selection mode of the grouped toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup/3242767-selectionmode?language=objc
func (t_ ToolbarItemGroup) SelectionMode() ToolbarItemGroupSelectionMode {
	rv := objc.Call[ToolbarItemGroupSelectionMode](t_, objc.Sel("selectionMode"))
	return rv
}

// The selection mode of the grouped toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup/3242767-selectionmode?language=objc
func (t_ ToolbarItemGroup) SetSelectionMode(value ToolbarItemGroupSelectionMode) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectionMode:"), value)
}

// The subitems of the grouped toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup/1529923-subitems?language=objc
func (t_ ToolbarItemGroup) Subitems() []ToolbarItem {
	rv := objc.Call[[]ToolbarItem](t_, objc.Sel("subitems"))
	return rv
}

// The subitems of the grouped toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup/1529923-subitems?language=objc
func (t_ ToolbarItemGroup) SetSubitems(value []IToolbarItem) {
	objc.Call[objc.Void](t_, objc.Sel("setSubitems:"), value)
}

// A value that represents how a toolbar displays a grouped toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup/3242762-controlrepresentation?language=objc
func (t_ ToolbarItemGroup) ControlRepresentation() ToolbarItemGroupControlRepresentation {
	rv := objc.Call[ToolbarItemGroupControlRepresentation](t_, objc.Sel("controlRepresentation"))
	return rv
}

// A value that represents how a toolbar displays a grouped toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup/3242762-controlrepresentation?language=objc
func (t_ ToolbarItemGroup) SetControlRepresentation(value ToolbarItemGroupControlRepresentation) {
	objc.Call[objc.Void](t_, objc.Sel("setControlRepresentation:"), value)
}

// The index value for the most recently selected subitem of a grouped toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup/3242766-selectedindex?language=objc
func (t_ ToolbarItemGroup) SelectedIndex() int {
	rv := objc.Call[int](t_, objc.Sel("selectedIndex"))
	return rv
}

// The index value for the most recently selected subitem of a grouped toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritemgroup/3242766-selectedindex?language=objc
func (t_ ToolbarItemGroup) SetSelectedIndex(value int) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectedIndex:"), value)
}
