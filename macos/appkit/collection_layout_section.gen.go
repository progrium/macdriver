// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionLayoutSection] class.
var CollectionLayoutSectionClass = _CollectionLayoutSectionClass{objc.GetClass("NSCollectionLayoutSection")}

type _CollectionLayoutSectionClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutSection] class.
type ICollectionLayoutSection interface {
	objc.IObject
	SupplementariesFollowContentInsets() bool
	SetSupplementariesFollowContentInsets(value bool)
	OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior
	SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior)
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
	VisibleItemsInvalidationHandler() CollectionLayoutSectionVisibleItemsInvalidationHandler
	SetVisibleItemsInvalidationHandler(value CollectionLayoutSectionVisibleItemsInvalidationHandler)
	DecorationItems() []CollectionLayoutDecorationItem
	SetDecorationItems(value []ICollectionLayoutDecorationItem)
	InterGroupSpacing() float64
	SetInterGroupSpacing(value float64)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem)
}

// A container that combines a set of groups into distinct visual groupings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection?language=objc
type CollectionLayoutSection struct {
	objc.Object
}

func CollectionLayoutSectionFrom(ptr unsafe.Pointer) CollectionLayoutSection {
	return CollectionLayoutSection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionLayoutSectionClass) SectionWithGroup(group ICollectionLayoutGroup) CollectionLayoutSection {
	rv := objc.Call[CollectionLayoutSection](cc, objc.Sel("sectionWithGroup:"), objc.Ptr(group))
	return rv
}

// Creates a section containing the specified group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3213881-sectionwithgroup?language=objc
func CollectionLayoutSection_SectionWithGroup(group ICollectionLayoutGroup) CollectionLayoutSection {
	return CollectionLayoutSectionClass.SectionWithGroup(group)
}

func (cc _CollectionLayoutSectionClass) Alloc() CollectionLayoutSection {
	rv := objc.Call[CollectionLayoutSection](cc, objc.Sel("alloc"))
	return rv
}

func CollectionLayoutSection_Alloc() CollectionLayoutSection {
	return CollectionLayoutSectionClass.Alloc()
}

func (cc _CollectionLayoutSectionClass) New() CollectionLayoutSection {
	rv := objc.Call[CollectionLayoutSection](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSection() CollectionLayoutSection {
	return CollectionLayoutSectionClass.New()
}

func (c_ CollectionLayoutSection) Init() CollectionLayoutSection {
	rv := objc.Call[CollectionLayoutSection](c_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the section’s supplementary items follow the specified content insets for the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199095-supplementariesfollowcontentinse?language=objc
func (c_ CollectionLayoutSection) SupplementariesFollowContentInsets() bool {
	rv := objc.Call[bool](c_, objc.Sel("supplementariesFollowContentInsets"))
	return rv
}

// A Boolean value that indicates whether the section’s supplementary items follow the specified content insets for the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199095-supplementariesfollowcontentinse?language=objc
func (c_ CollectionLayoutSection) SetSupplementariesFollowContentInsets(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setSupplementariesFollowContentInsets:"), value)
}

// The section’s scrolling behavior in relation to the main layout axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199094-orthogonalscrollingbehavior?language=objc
func (c_ CollectionLayoutSection) OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior {
	rv := objc.Call[CollectionLayoutSectionOrthogonalScrollingBehavior](c_, objc.Sel("orthogonalScrollingBehavior"))
	return rv
}

// The section’s scrolling behavior in relation to the main layout axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199094-orthogonalscrollingbehavior?language=objc
func (c_ CollectionLayoutSection) SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior) {
	objc.Call[objc.Void](c_, objc.Sel("setOrthogonalScrollingBehavior:"), value)
}

// The amount of space between the content of the section and its boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199090-contentinsets?language=objc
func (c_ CollectionLayoutSection) ContentInsets() DirectionalEdgeInsets {
	rv := objc.Call[DirectionalEdgeInsets](c_, objc.Sel("contentInsets"))
	return rv
}

// The amount of space between the content of the section and its boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199090-contentinsets?language=objc
func (c_ CollectionLayoutSection) SetContentInsets(value DirectionalEdgeInsets) {
	objc.Call[objc.Void](c_, objc.Sel("setContentInsets:"), value)
}

// A closure called before each layout cycle to allow modification of the items in the section immediately before they’re displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199096-visibleitemsinvalidationhandler?language=objc
func (c_ CollectionLayoutSection) VisibleItemsInvalidationHandler() CollectionLayoutSectionVisibleItemsInvalidationHandler {
	rv := objc.Call[CollectionLayoutSectionVisibleItemsInvalidationHandler](c_, objc.Sel("visibleItemsInvalidationHandler"))
	return rv
}

// A closure called before each layout cycle to allow modification of the items in the section immediately before they’re displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199096-visibleitemsinvalidationhandler?language=objc
func (c_ CollectionLayoutSection) SetVisibleItemsInvalidationHandler(value CollectionLayoutSectionVisibleItemsInvalidationHandler) {
	objc.Call[objc.Void](c_, objc.Sel("setVisibleItemsInvalidationHandler:"), value)
}

// An array of the decoration items that are anchored to the section, such as background decoration views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199091-decorationitems?language=objc
func (c_ CollectionLayoutSection) DecorationItems() []CollectionLayoutDecorationItem {
	rv := objc.Call[[]CollectionLayoutDecorationItem](c_, objc.Sel("decorationItems"))
	return rv
}

// An array of the decoration items that are anchored to the section, such as background decoration views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199091-decorationitems?language=objc
func (c_ CollectionLayoutSection) SetDecorationItems(value []ICollectionLayoutDecorationItem) {
	objc.Call[objc.Void](c_, objc.Sel("setDecorationItems:"), value)
}

// The amount of space between the groups in the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199093-intergroupspacing?language=objc
func (c_ CollectionLayoutSection) InterGroupSpacing() float64 {
	rv := objc.Call[float64](c_, objc.Sel("interGroupSpacing"))
	return rv
}

// The amount of space between the groups in the section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199093-intergroupspacing?language=objc
func (c_ CollectionLayoutSection) SetInterGroupSpacing(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setInterGroupSpacing:"), value)
}

// An array of the supplementary items that are associated with the boundary edges of the section, such as headers and footers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199089-boundarysupplementaryitems?language=objc
func (c_ CollectionLayoutSection) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	rv := objc.Call[[]CollectionLayoutBoundarySupplementaryItem](c_, objc.Sel("boundarySupplementaryItems"))
	return rv
}

// An array of the supplementary items that are associated with the boundary edges of the section, such as headers and footers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutsection/3199089-boundarysupplementaryitems?language=objc
func (c_ CollectionLayoutSection) SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem) {
	objc.Call[objc.Void](c_, objc.Sel("setBoundarySupplementaryItems:"), value)
}
