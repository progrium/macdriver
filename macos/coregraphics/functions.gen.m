// AUTO-GENERATED CODE, DO NOT MODIFY

#import "CoreGraphics/CoreGraphics.h"
bool PDFDictionaryGetArray(void * dict, char* key, void * value) {
	return (bool)CGPDFDictionaryGetArray(
		// *typing.RefType
		(CGPDFDictionaryRef*)dict,
		// *typing.CStringType
		(char*)key,
		// *typing.RefType
		(CGPDFArrayRef*)value
	);
}
