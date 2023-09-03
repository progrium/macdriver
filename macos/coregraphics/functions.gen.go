// AUTO-GENERATED CODE, DO NOT MODIFY

package coregraphics

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import "CoreGraphics/CoreGraphics.h"
// void * ColorSpaceCreateCalibratedGray(float* whitePoint, float* blackPoint, float gamma);
// bool PDFDocumentUnlockWithPassword(void * document, uint8_t* password);
import "C"
import (
	"unsafe"
)

// Creates a calibrated grayscale color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408887-cgcolorspacecreatecalibratedgray?language=objc
func ColorSpaceCreateCalibratedGray(whitePoint *float64, blackPoint *float64, gamma float64) ColorSpaceRef {
	rv := C.ColorSpaceCreateCalibratedGray(
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(whitePoint)),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(blackPoint)),
		// *typing.PrimitiveType
		C.float(gamma),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Unlocks an encrypted PDF document when a valid password is supplied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402599-cgpdfdocumentunlockwithpassword?language=objc
func PDFDocumentUnlockWithPassword(document PDFDocumentRef, password *uint8) bool {
	rv := C.PDFDocumentUnlockWithPassword(
		// *typing.RefType
		unsafe.Pointer(document),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(password)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}
