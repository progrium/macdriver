// AUTO-GENERATED CODE, DO NOT MODIFY

package coregraphics

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import "CoreGraphics/CoreGraphics.h"
// bool PDFDictionaryGetArray(void * dict, char* key, void * value);
import "C"
import (
	"unsafe"
)

// Returns whether there is a PDF array associated with a specified key in a PDF dictionary and, if so, retrieves that array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430229-cgpdfdictionarygetarray?language=objc
func PDFDictionaryGetArray(dict unsafe.Pointer, key string, value unsafe.Pointer) bool {
	keyVal := C.CString(key)
	defer C.free(unsafe.Pointer(keyVal))
	rv := C.PDFDictionaryGetArray(
		// *typing.RefType
		unsafe.Pointer(dict),
		// *typing.CStringType
		keyVal,
		// *typing.RefType
		unsafe.Pointer(value),
	)
	// *typing.PrimitiveType
	return bool(rv)
}
