// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A set of methods that you use to manage the content in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement?language=objc
type PCollectionViewElement interface {
	// optional
	WillTransitionFromLayoutToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	HasWillTransitionFromLayoutToLayout() bool

	// optional
	ApplyLayoutAttributes(layoutAttributes CollectionViewLayoutAttributes)
	HasApplyLayoutAttributes() bool

	// optional
	DidTransitionFromLayoutToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	HasDidTransitionFromLayoutToLayout() bool

	// optional
	PrepareForReuse()
	HasPrepareForReuse() bool

	// optional
	PreferredLayoutAttributesFittingAttributes(layoutAttributes CollectionViewLayoutAttributes) CollectionViewLayoutAttributes
	HasPreferredLayoutAttributesFittingAttributes() bool
}

// ensure impl type implements protocol interface
var _ PCollectionViewElement = (*CollectionViewElementObject)(nil)

// A concrete type for the [PCollectionViewElement] protocol.
type CollectionViewElementObject struct {
	objc.Object
}

func (c_ CollectionViewElementObject) HasWillTransitionFromLayoutToLayout() bool {
	return c_.RespondsToSelector(objc.Sel("willTransitionFromLayout:toLayout:"))
}

// Tells the element that the layout object of the collection view is about to change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement/1528165-willtransitionfromlayout?language=objc
func (c_ CollectionViewElementObject) WillTransitionFromLayoutToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout) {
	objc.Call[objc.Void](c_, objc.Sel("willTransitionFromLayout:toLayout:"), oldLayout, newLayout)
}

func (c_ CollectionViewElementObject) HasApplyLayoutAttributes() bool {
	return c_.RespondsToSelector(objc.Sel("applyLayoutAttributes:"))
}

// Applies the specified layout attributes to the element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement/1528294-applylayoutattributes?language=objc
func (c_ CollectionViewElementObject) ApplyLayoutAttributes(layoutAttributes CollectionViewLayoutAttributes) {
	objc.Call[objc.Void](c_, objc.Sel("applyLayoutAttributes:"), layoutAttributes)
}

func (c_ CollectionViewElementObject) HasDidTransitionFromLayoutToLayout() bool {
	return c_.RespondsToSelector(objc.Sel("didTransitionFromLayout:toLayout:"))
}

// Tells the element that the layout object of the collection view changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement/1525851-didtransitionfromlayout?language=objc
func (c_ CollectionViewElementObject) DidTransitionFromLayoutToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout) {
	objc.Call[objc.Void](c_, objc.Sel("didTransitionFromLayout:toLayout:"), oldLayout, newLayout)
}

func (c_ CollectionViewElementObject) HasPrepareForReuse() bool {
	return c_.RespondsToSelector(objc.Sel("prepareForReuse"))
}

// Performs any necessary cleanup to prepare the element for use again. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement/1528248-prepareforreuse?language=objc
func (c_ CollectionViewElementObject) PrepareForReuse() {
	objc.Call[objc.Void](c_, objc.Sel("prepareForReuse"))
}

func (c_ CollectionViewElementObject) HasPreferredLayoutAttributesFittingAttributes() bool {
	return c_.RespondsToSelector(objc.Sel("preferredLayoutAttributesFittingAttributes:"))
}

// Asks your element if it wants to modify any layout attributes before they are applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewelement/1528259-preferredlayoutattributesfitting?language=objc
func (c_ CollectionViewElementObject) PreferredLayoutAttributesFittingAttributes(layoutAttributes CollectionViewLayoutAttributes) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("preferredLayoutAttributesFittingAttributes:"), layoutAttributes)
	return rv
}
