// AUTO-GENERATED CODE, DO NOT MODIFY

package coregraphics

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import "CoreGraphics/CoreGraphics.h"
// uint32_t MainDisplayID();
// void * DisplayCreateImage(uint32_t displayID);
import "C"

// Returns the display ID of the main display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455620-cgmaindisplayid?language=objc
func MainDisplayID() DirectDisplayID {
	rv := C.MainDisplayID()
	return DirectDisplayID(rv)
}

// Returns an image containing the contents of the specified display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455691-cgdisplaycreateimage?language=objc
func DisplayCreateImage(displayID DirectDisplayID) ImageRef {
	rv := C.DisplayCreateImage(C.uint32_t(displayID))
	return ImageRef(rv)
}
