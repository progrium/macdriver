// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionLayoutDecorationItem] class.
var CollectionLayoutDecorationItemClass = _CollectionLayoutDecorationItemClass{objc.GetClass("NSCollectionLayoutDecorationItem")}

type _CollectionLayoutDecorationItemClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutDecorationItem] class.
type ICollectionLayoutDecorationItem interface {
	ICollectionLayoutItem
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

// An object used to add a background to a section of a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdecorationitem?language=objc
type CollectionLayoutDecorationItem struct {
	CollectionLayoutItem
}

func CollectionLayoutDecorationItemFrom(ptr unsafe.Pointer) CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItem{
		CollectionLayoutItem: CollectionLayoutItemFrom(ptr),
	}
}

func (cc _CollectionLayoutDecorationItemClass) BackgroundDecorationItemWithElementKind(elementKind string) CollectionLayoutDecorationItem {
	rv := objc.Call[CollectionLayoutDecorationItem](cc, objc.Sel("backgroundDecorationItemWithElementKind:"), elementKind)
	return rv
}

// Creates a section background with a string to identify the element kind. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdecorationitem/3199051-backgrounddecorationitemwithelem?language=objc
func CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(elementKind string) CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.BackgroundDecorationItemWithElementKind(elementKind)
}

func (cc _CollectionLayoutDecorationItemClass) Alloc() CollectionLayoutDecorationItem {
	rv := objc.Call[CollectionLayoutDecorationItem](cc, objc.Sel("alloc"))
	return rv
}

func CollectionLayoutDecorationItem_Alloc() CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.Alloc()
}

func (cc _CollectionLayoutDecorationItemClass) New() CollectionLayoutDecorationItem {
	rv := objc.Call[CollectionLayoutDecorationItem](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutDecorationItem() CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.New()
}

func (c_ CollectionLayoutDecorationItem) Init() CollectionLayoutDecorationItem {
	rv := objc.Call[CollectionLayoutDecorationItem](c_, objc.Sel("init"))
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutDecorationItem {
	rv := objc.Call[CollectionLayoutDecorationItem](cc, objc.Sel("itemWithLayoutSize:"), objc.Ptr(layoutSize))
	return rv
}

// Creates an item of the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem/3213871-itemwithlayoutsize?language=objc
func CollectionLayoutDecorationItem_ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.ItemWithLayoutSize(layoutSize)
}

// A string that identifies the type of decoration item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdecorationitem/3213831-elementkind?language=objc
func (c_ CollectionLayoutDecorationItem) ElementKind() string {
	rv := objc.Call[string](c_, objc.Sel("elementKind"))
	return rv
}

// The vertical stacking order of the decoration item in relation to other items in the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdecorationitem/3199053-zindex?language=objc
func (c_ CollectionLayoutDecorationItem) ZIndex() int {
	rv := objc.Call[int](c_, objc.Sel("zIndex"))
	return rv
}

// The vertical stacking order of the decoration item in relation to other items in the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutdecorationitem/3199053-zindex?language=objc
func (c_ CollectionLayoutDecorationItem) SetZIndex(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setZIndex:"), value)
}