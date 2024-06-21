// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A set of methods an app can implement to provide Spotlight for Help for its own custom help data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemsearching?language=objc
type PUserInterfaceItemSearching interface {
	// optional
	LocalizedTitlesForItem(item objc.Object) []string
	HasLocalizedTitlesForItem() bool

	// optional
	PerformActionForItem(item objc.Object)
	HasPerformActionForItem() bool

	// optional
	SearchForItemsWithSearchStringResultLimitMatchedItemHandler(searchString string, resultLimit int, handleMatchedItems func(items []objc.Object))
	HasSearchForItemsWithSearchStringResultLimitMatchedItemHandler() bool

	// optional
	ShowAllHelpTopicsForSearchString(searchString string)
	HasShowAllHelpTopicsForSearchString() bool
}

// ensure impl type implements protocol interface
var _ PUserInterfaceItemSearching = (*UserInterfaceItemSearchingObject)(nil)

// A concrete type for the [PUserInterfaceItemSearching] protocol.
type UserInterfaceItemSearchingObject struct {
	objc.Object
}

func (u_ UserInterfaceItemSearchingObject) HasLocalizedTitlesForItem() bool {
	return u_.RespondsToSelector(objc.Sel("localizedTitlesForItem:"))
}

// Returns an array of localized strings that will form the help menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemsearching/1420814-localizedtitlesforitem?language=objc
func (u_ UserInterfaceItemSearchingObject) LocalizedTitlesForItem(item objc.Object) []string {
	rv := objc.Call[[]string](u_, objc.Sel("localizedTitlesForItem:"), item)
	return rv
}

func (u_ UserInterfaceItemSearchingObject) HasPerformActionForItem() bool {
	return u_.RespondsToSelector(objc.Sel("performActionForItem:"))
}

// Invoked when the user selects a search result in Help menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemsearching/1420812-performactionforitem?language=objc
func (u_ UserInterfaceItemSearchingObject) PerformActionForItem(item objc.Object) {
	objc.Call[objc.Void](u_, objc.Sel("performActionForItem:"), item)
}

func (u_ UserInterfaceItemSearchingObject) HasSearchForItemsWithSearchStringResultLimitMatchedItemHandler() bool {
	return u_.RespondsToSelector(objc.Sel("searchForItemsWithSearchString:resultLimit:matchedItemHandler:"))
}

// Search for the specified items, with the result limit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemsearching/1420816-searchforitemswithsearchstring?language=objc
func (u_ UserInterfaceItemSearchingObject) SearchForItemsWithSearchStringResultLimitMatchedItemHandler(searchString string, resultLimit int, handleMatchedItems func(items []objc.Object)) {
	objc.Call[objc.Void](u_, objc.Sel("searchForItemsWithSearchString:resultLimit:matchedItemHandler:"), searchString, resultLimit, handleMatchedItems)
}

func (u_ UserInterfaceItemSearchingObject) HasShowAllHelpTopicsForSearchString() bool {
	return u_.RespondsToSelector(objc.Sel("showAllHelpTopicsForSearchString:"))
}

// If this method is implemented, a "Show All Help Topics" item will appear in the menu and this method is called when the user selects it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfaceitemsearching/1420806-showallhelptopicsforsearchstring?language=objc
func (u_ UserInterfaceItemSearchingObject) ShowAllHelpTopicsForSearchString(searchString string) {
	objc.Call[objc.Void](u_, objc.Sel("showAllHelpTopicsForSearchString:"), searchString)
}
