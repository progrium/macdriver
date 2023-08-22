// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VideoCompositionCoreAnimationTool] class.
var VideoCompositionCoreAnimationToolClass = _VideoCompositionCoreAnimationToolClass{objc.GetClass("AVVideoCompositionCoreAnimationTool")}

type _VideoCompositionCoreAnimationToolClass struct {
	objc.Class
}

// An interface definition for the [VideoCompositionCoreAnimationTool] class.
type IVideoCompositionCoreAnimationTool interface {
	objc.IObject
}

// An object used to incorporate Core Animation into a video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositioncoreanimationtool?language=objc
type VideoCompositionCoreAnimationTool struct {
	objc.Object
}

func VideoCompositionCoreAnimationToolFrom(ptr unsafe.Pointer) VideoCompositionCoreAnimationTool {
	return VideoCompositionCoreAnimationTool{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VideoCompositionCoreAnimationToolClass) VideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayerInLayer(videoLayer quartzcore.ILayer, animationLayer quartzcore.ILayer) VideoCompositionCoreAnimationTool {
	rv := objc.Call[VideoCompositionCoreAnimationTool](vc, objc.Sel("videoCompositionCoreAnimationToolWithPostProcessingAsVideoLayer:inLayer:"), objc.Ptr(videoLayer), objc.Ptr(animationLayer))
	return rv
}

// Composes the composited video frame with a Core Animation layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositioncoreanimationtool/1389594-videocompositioncoreanimationtoo?language=objc
func VideoCompositionCoreAnimationTool_VideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayerInLayer(videoLayer quartzcore.ILayer, animationLayer quartzcore.ILayer) VideoCompositionCoreAnimationTool {
	return VideoCompositionCoreAnimationToolClass.VideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayerInLayer(videoLayer, animationLayer)
}

func (vc _VideoCompositionCoreAnimationToolClass) Alloc() VideoCompositionCoreAnimationTool {
	rv := objc.Call[VideoCompositionCoreAnimationTool](vc, objc.Sel("alloc"))
	return rv
}

func VideoCompositionCoreAnimationTool_Alloc() VideoCompositionCoreAnimationTool {
	return VideoCompositionCoreAnimationToolClass.Alloc()
}

func (vc _VideoCompositionCoreAnimationToolClass) New() VideoCompositionCoreAnimationTool {
	rv := objc.Call[VideoCompositionCoreAnimationTool](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVideoCompositionCoreAnimationTool() VideoCompositionCoreAnimationTool {
	return VideoCompositionCoreAnimationToolClass.New()
}

func (v_ VideoCompositionCoreAnimationTool) Init() VideoCompositionCoreAnimationTool {
	rv := objc.Call[VideoCompositionCoreAnimationTool](v_, objc.Sel("init"))
	return rv
}
