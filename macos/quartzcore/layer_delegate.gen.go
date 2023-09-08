// Code generated by DarwinKit. DO NOT EDIT.

package quartzcore

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// Methods your app can implement to respond to layer-related events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate?language=objc
type PLayerDelegate interface {
	// optional
	DrawLayerInContext(layer Layer, ctx coregraphics.ContextRef)
	HasDrawLayerInContext() bool

	// optional
	DisplayLayer(layer Layer)
	HasDisplayLayer() bool

	// optional
	LayoutSublayersOfLayer(layer Layer)
	HasLayoutSublayersOfLayer() bool

	// optional
	ActionForLayerForKey(layer Layer, event string) ActionObject
	HasActionForLayerForKey() bool

	// optional
	LayerWillDraw(layer Layer)
	HasLayerWillDraw() bool
}

// A delegate implementation builder for the [PLayerDelegate] protocol.
type LayerDelegate struct {
	_DrawLayerInContext     func(layer Layer, ctx coregraphics.ContextRef)
	_DisplayLayer           func(layer Layer)
	_LayoutSublayersOfLayer func(layer Layer)
	_ActionForLayerForKey   func(layer Layer, event string) ActionObject
	_LayerWillDraw          func(layer Layer)
}

func (di *LayerDelegate) HasDrawLayerInContext() bool {
	return di._DrawLayerInContext != nil
}

// Tells the delegate to implement the display process using the layer's CGContextRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097262-drawlayer?language=objc
func (di *LayerDelegate) SetDrawLayerInContext(f func(layer Layer, ctx coregraphics.ContextRef)) {
	di._DrawLayerInContext = f
}

// Tells the delegate to implement the display process using the layer's CGContextRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097262-drawlayer?language=objc
func (di *LayerDelegate) DrawLayerInContext(layer Layer, ctx coregraphics.ContextRef) {
	di._DrawLayerInContext(layer, ctx)
}
func (di *LayerDelegate) HasDisplayLayer() bool {
	return di._DisplayLayer != nil
}

// Tells the delegate to implement the display process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097261-displaylayer?language=objc
func (di *LayerDelegate) SetDisplayLayer(f func(layer Layer)) {
	di._DisplayLayer = f
}

// Tells the delegate to implement the display process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097261-displaylayer?language=objc
func (di *LayerDelegate) DisplayLayer(layer Layer) {
	di._DisplayLayer(layer)
}
func (di *LayerDelegate) HasLayoutSublayersOfLayer() bool {
	return di._LayoutSublayersOfLayer != nil
}

// Tells the delegate a layer's bounds have changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097257-layoutsublayersoflayer?language=objc
func (di *LayerDelegate) SetLayoutSublayersOfLayer(f func(layer Layer)) {
	di._LayoutSublayersOfLayer = f
}

// Tells the delegate a layer's bounds have changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097257-layoutsublayersoflayer?language=objc
func (di *LayerDelegate) LayoutSublayersOfLayer(layer Layer) {
	di._LayoutSublayersOfLayer(layer)
}
func (di *LayerDelegate) HasActionForLayerForKey() bool {
	return di._ActionForLayerForKey != nil
}

// Returns the default action of the [quartzcore/calayer/actionforkey] method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097264-actionforlayer?language=objc
func (di *LayerDelegate) SetActionForLayerForKey(f func(layer Layer, event string) ActionObject) {
	di._ActionForLayerForKey = f
}

// Returns the default action of the [quartzcore/calayer/actionforkey] method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097264-actionforlayer?language=objc
func (di *LayerDelegate) ActionForLayerForKey(layer Layer, event string) ActionObject {
	return di._ActionForLayerForKey(layer, event)
}
func (di *LayerDelegate) HasLayerWillDraw() bool {
	return di._LayerWillDraw != nil
}

// Notifies the delegate of an imminent draw. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097263-layerwilldraw?language=objc
func (di *LayerDelegate) SetLayerWillDraw(f func(layer Layer)) {
	di._LayerWillDraw = f
}

// Notifies the delegate of an imminent draw. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097263-layerwilldraw?language=objc
func (di *LayerDelegate) LayerWillDraw(layer Layer) {
	di._LayerWillDraw(layer)
}

// ensure impl type implements protocol interface
var _ PLayerDelegate = (*LayerDelegateObject)(nil)

// A concrete type for the [PLayerDelegate] protocol.
type LayerDelegateObject struct {
	objc.Object
}

func (l_ LayerDelegateObject) HasDrawLayerInContext() bool {
	return l_.RespondsToSelector(objc.Sel("drawLayer:inContext:"))
}

// Tells the delegate to implement the display process using the layer's CGContextRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097262-drawlayer?language=objc
func (l_ LayerDelegateObject) DrawLayerInContext(layer Layer, ctx coregraphics.ContextRef) {
	objc.Call[objc.Void](l_, objc.Sel("drawLayer:inContext:"), objc.Ptr(layer), ctx)
}

func (l_ LayerDelegateObject) HasDisplayLayer() bool {
	return l_.RespondsToSelector(objc.Sel("displayLayer:"))
}

// Tells the delegate to implement the display process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097261-displaylayer?language=objc
func (l_ LayerDelegateObject) DisplayLayer(layer Layer) {
	objc.Call[objc.Void](l_, objc.Sel("displayLayer:"), objc.Ptr(layer))
}

func (l_ LayerDelegateObject) HasLayoutSublayersOfLayer() bool {
	return l_.RespondsToSelector(objc.Sel("layoutSublayersOfLayer:"))
}

// Tells the delegate a layer's bounds have changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097257-layoutsublayersoflayer?language=objc
func (l_ LayerDelegateObject) LayoutSublayersOfLayer(layer Layer) {
	objc.Call[objc.Void](l_, objc.Sel("layoutSublayersOfLayer:"), objc.Ptr(layer))
}

func (l_ LayerDelegateObject) HasActionForLayerForKey() bool {
	return l_.RespondsToSelector(objc.Sel("actionForLayer:forKey:"))
}

// Returns the default action of the [quartzcore/calayer/actionforkey] method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097264-actionforlayer?language=objc
func (l_ LayerDelegateObject) ActionForLayerForKey(layer Layer, event string) ActionObject {
	rv := objc.Call[ActionObject](l_, objc.Sel("actionForLayer:forKey:"), objc.Ptr(layer), event)
	return rv
}

func (l_ LayerDelegateObject) HasLayerWillDraw() bool {
	return l_.RespondsToSelector(objc.Sel("layerWillDraw:"))
}

// Notifies the delegate of an imminent draw. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayerdelegate/2097263-layerwilldraw?language=objc
func (l_ LayerDelegateObject) LayerWillDraw(layer Layer) {
	objc.Call[objc.Void](l_, objc.Sel("layerWillDraw:"), objc.Ptr(layer))
}
