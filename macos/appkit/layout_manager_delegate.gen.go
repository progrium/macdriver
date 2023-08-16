// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods that delegates of layout manager objects implement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanagerdelegate?language=objc
type PLayoutManagerDelegate interface {
	// optional
	LayoutManagerDidInvalidateLayout(sender LayoutManager)
	HasLayoutManagerDidInvalidateLayout() bool

	// optional
	LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange foundation.RangePointer) map[foundation.AttributedStringKey]objc.IObject
	HasLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange() bool
}

// A delegate implementation builder for the [PLayoutManagerDelegate] protocol.
type LayoutManagerDelegate struct {
	_LayoutManagerDidInvalidateLayout                                                          func(sender LayoutManager)
	_LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange func(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange foundation.RangePointer) map[foundation.AttributedStringKey]objc.IObject
}

func (di *LayoutManagerDelegate) HasLayoutManagerDidInvalidateLayout() bool {
	return di._LayoutManagerDidInvalidateLayout != nil
}

// Informs the delegate when the specified layout manager invalidates layout information (not glyph information). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanagerdelegate/1402993-layoutmanagerdidinvalidatelayout?language=objc
func (di *LayoutManagerDelegate) SetLayoutManagerDidInvalidateLayout(f func(sender LayoutManager)) {
	di._LayoutManagerDidInvalidateLayout = f
}

// Informs the delegate when the specified layout manager invalidates layout information (not glyph information). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanagerdelegate/1402993-layoutmanagerdidinvalidatelayout?language=objc
func (di *LayoutManagerDelegate) LayoutManagerDidInvalidateLayout(sender LayoutManager) {
	di._LayoutManagerDidInvalidateLayout(sender)
}
func (di *LayoutManagerDelegate) HasLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange() bool {
	return di._LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange != nil
}

// Asks the delegate whether to use temporary attributes when drawing the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanagerdelegate/1403085-layoutmanager?language=objc
func (di *LayoutManagerDelegate) SetLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(f func(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange foundation.RangePointer) map[foundation.AttributedStringKey]objc.IObject) {
	di._LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange = f
}

// Asks the delegate whether to use temporary attributes when drawing the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanagerdelegate/1403085-layoutmanager?language=objc
func (di *LayoutManagerDelegate) LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange foundation.RangePointer) map[foundation.AttributedStringKey]objc.IObject {
	return di._LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager, attrs, toScreen, charIndex, effectiveCharRange)
}

// A concrete type wrapper for the [PLayoutManagerDelegate] protocol.
type LayoutManagerDelegateWrapper struct {
	objc.Object
}

func (l_ LayoutManagerDelegateWrapper) HasLayoutManagerDidInvalidateLayout() bool {
	return l_.RespondsToSelector(objc.Sel("layoutManagerDidInvalidateLayout:"))
}

// Informs the delegate when the specified layout manager invalidates layout information (not glyph information). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanagerdelegate/1402993-layoutmanagerdidinvalidatelayout?language=objc
func (l_ LayoutManagerDelegateWrapper) LayoutManagerDidInvalidateLayout(sender ILayoutManager) {
	objc.Call[objc.Void](l_, objc.Sel("layoutManagerDidInvalidateLayout:"), objc.Ptr(sender))
}

func (l_ LayoutManagerDelegateWrapper) HasLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange() bool {
	return l_.RespondsToSelector(objc.Sel("layoutManager:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:"))
}

// Asks the delegate whether to use temporary attributes when drawing the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanagerdelegate/1403085-layoutmanager?language=objc
func (l_ LayoutManagerDelegateWrapper) LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager ILayoutManager, attrs map[foundation.AttributedStringKey]objc.IObject, toScreen bool, charIndex uint, effectiveCharRange foundation.RangePointer) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.Call[map[foundation.AttributedStringKey]objc.Object](l_, objc.Sel("layoutManager:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:"), objc.Ptr(layoutManager), attrs, toScreen, charIndex, effectiveCharRange)
	return rv
}
