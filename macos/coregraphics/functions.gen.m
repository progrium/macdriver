// AUTO-GENERATED CODE, DO NOT MODIFY

#import "CoreGraphics/CoreGraphics.h"
void ContextStrokeEllipseInRect(void * c, CGRect rect) {
	return (void)CGContextStrokeEllipseInRect(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect
	);
}
CGRect ContextConvertRectToDeviceSpace(void * c, CGRect rect) {
	return (CGRect)CGContextConvertRectToDeviceSpace(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect
	);
}
void EventSetSource(void * event, void * source) {
	return (void)CGEventSetSource(
		// *typing.RefType
		(CGEventRef)event,
		// *typing.RefType
		(CGEventSourceRef)source
	);
}
void ContextSetFontSize(void * c, float size) {
	return (void)CGContextSetFontSize(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)size
	);
}
void ContextClipToRects(void * c, const CGRect* rects, uint count) {
	return (void)CGContextClipToRects(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)rects,
		// *typing.PrimitiveType
		(NSUInteger)count
	);
}
CGRect RectApplyAffineTransform(CGRect rect, CGAffineTransform t) {
	return (CGRect)CGRectApplyAffineTransform(
		// *typing.StructType
		(CGRect)rect,
		// *typing.StructType
		(CGAffineTransform)t
	);
}
int64_t EventSourceGetUserData(void * source) {
	return (int64_t)CGEventSourceGetUserData(
		// *typing.RefType
		(CGEventSourceRef)source
	);
}
CGImageAlphaInfo ImageGetAlphaInfo(void * image) {
	return (CGImageAlphaInfo)CGImageGetAlphaInfo(
		// *typing.RefType
		(CGImageRef)image
	);
}
boolean_t DisplayIsOnline(CGDirectDisplayID display) {
	return (boolean_t)CGDisplayIsOnline(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
CGError ReleaseDisplayFadeReservation(CGDisplayFadeReservationToken token) {
	return (CGError)CGReleaseDisplayFadeReservation(
		// *typing.AliasType
		(CGDisplayFadeReservationToken)token
	);
}
bool ColorSpaceUsesExtendedRange(void * space) {
	return (bool)CGColorSpaceUsesExtendedRange(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
uint BitmapContextGetBitsPerComponent(void * context) {
	return (uint)CGBitmapContextGetBitsPerComponent(
		// *typing.RefType
		(CGContextRef)context
	);
}
void * PDFContextCreate(void * consumer, const CGRect* mediaBox, void * auxiliaryInfo) {
	return (void *)CGPDFContextCreate(
		// *typing.RefType
		(CGDataConsumerRef)consumer,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)mediaBox,
		// *typing.RefType
		(CFDictionaryRef)auxiliaryInfo
	);
}
void * ColorSpaceGetName(void * space) {
	return (void *)CGColorSpaceGetName(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void ContextBeginTransparencyLayer(void * c, void * auxiliaryInfo) {
	return (void)CGContextBeginTransparencyLayer(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CFDictionaryRef)auxiliaryInfo
	);
}
void * ColorGetColorSpace(void * color) {
	return (void *)CGColorGetColorSpace(
		// *typing.RefType
		(CGColorRef)color
	);
}
int64_t EventGetIntegerValueField(void * event, CGEventField field) {
	return (int64_t)CGEventGetIntegerValueField(
		// *typing.RefType
		(CGEventRef)event,
		// *typing.AliasType
		(CGEventField)field
	);
}
void ContextSetTextPosition(void * c, float x, float y) {
	return (void)CGContextSetTextPosition(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y
	);
}
void * ColorSpaceCopyICCData(void * space) {
	return (void *)CGColorSpaceCopyICCData(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void ContextSetFont(void * c, void * font) {
	return (void)CGContextSetFont(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGFontRef)font
	);
}
CGRect ContextConvertRectToUserSpace(void * c, CGRect rect) {
	return (CGRect)CGContextConvertRectToUserSpace(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect
	);
}
void * ColorRetain(void * color) {
	return (void *)CGColorRetain(
		// *typing.RefType
		(CGColorRef)color
	);
}
void ContextBeginPath(void * c) {
	return (void)CGContextBeginPath(
		// *typing.RefType
		(CGContextRef)c
	);
}
bool ColorSpaceIsHLGBased(void * s) {
	return (bool)CGColorSpaceIsHLGBased(
		// *typing.RefType
		(CGColorSpaceRef)s
	);
}
void * ColorCreateGenericGray(float gray, float alpha) {
	return (void *)CGColorCreateGenericGray(
		// *typing.PrimitiveType
		(CGFloat)gray,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
void DataProviderRelease(void * provider) {
	return (void)CGDataProviderRelease(
		// *typing.RefType
		(CGDataProviderRef)provider
	);
}
void PathAddArc(void * path, const CGAffineTransform* m, float x, float y, float radius, float startAngle, float endAngle, bool clockwise) {
	return (void)CGPathAddArc(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y,
		// *typing.PrimitiveType
		(CGFloat)radius,
		// *typing.PrimitiveType
		(CGFloat)startAngle,
		// *typing.PrimitiveType
		(CGFloat)endAngle,
		// *typing.PrimitiveType
		(BOOL)clockwise
	);
}
CFTypeID EventGetTypeID() {
	return (CFTypeID)CGEventGetTypeID(
	);
}
void ContextDrawLinearGradient(void * c, void * gradient, CGPoint startPoint, CGPoint endPoint, CGGradientDrawingOptions options) {
	return (void)CGContextDrawLinearGradient(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGGradientRef)gradient,
		// *typing.StructType
		(CGPoint)startPoint,
		// *typing.StructType
		(CGPoint)endPoint,
		// *typing.AliasType
		(CGGradientDrawingOptions)options
	);
}
CGRect PDFPageGetBoxRect(void * page, CGPDFBox box) {
	return (CGRect)CGPDFPageGetBoxRect(
		// *typing.RefType
		(CGPDFPageRef)page,
		// *typing.AliasType
		(CGPDFBox)box
	);
}
void ContextStrokeRectWithWidth(void * c, CGRect rect, float width) {
	return (void)CGContextStrokeRectWithWidth(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect,
		// *typing.PrimitiveType
		(CGFloat)width
	);
}
void * SessionCopyCurrentDictionary() {
	return (void *)CGSessionCopyCurrentDictionary(
	);
}
CFTypeID ShadingGetTypeID() {
	return (CFTypeID)CGShadingGetTypeID(
	);
}
void * ColorSpaceCreatePattern(void * baseSpace) {
	return (void *)CGColorSpaceCreatePattern(
		// *typing.RefType
		(CGColorSpaceRef)baseSpace
	);
}
CGRect DisplayBounds(CGDirectDisplayID display) {
	return (CGRect)CGDisplayBounds(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void ContextSetTextDrawingMode(void * c, CGTextDrawingMode mode) {
	return (void)CGContextSetTextDrawingMode(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.AliasType
		(CGTextDrawingMode)mode
	);
}
void * ColorSpaceCreateDeviceGray() {
	return (void *)CGColorSpaceCreateDeviceGray(
	);
}
void * ColorSpaceCreateCalibratedGray(const float* whitePoint, const float* blackPoint, float gamma) {
	return (void *)CGColorSpaceCreateCalibratedGray(
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)whitePoint,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)blackPoint,
		// *typing.PrimitiveType
		(CGFloat)gamma
	);
}
void * PathCreateCopyByStrokingPath(void * path, const CGAffineTransform* transform, float lineWidth, CGLineCap lineCap, CGLineJoin lineJoin, float miterLimit) {
	return (void *)CGPathCreateCopyByStrokingPath(
		// *typing.RefType
		(CGPathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)transform,
		// *typing.PrimitiveType
		(CGFloat)lineWidth,
		// *typing.AliasType
		(CGLineCap)lineCap,
		// *typing.AliasType
		(CGLineJoin)lineJoin,
		// *typing.PrimitiveType
		(CGFloat)miterLimit
	);
}
bool ColorEqualToColor(void * color1, void * color2) {
	return (bool)CGColorEqualToColor(
		// *typing.RefType
		(CGColorRef)color1,
		// *typing.RefType
		(CGColorRef)color2
	);
}
bool PathContainsPoint(void * path, const CGAffineTransform* m, CGPoint point, bool eoFill) {
	return (bool)CGPathContainsPoint(
		// *typing.RefType
		(CGPathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.StructType
		(CGPoint)point,
		// *typing.PrimitiveType
		(BOOL)eoFill
	);
}
void * FontCopyVariations(void * font) {
	return (void *)CGFontCopyVariations(
		// *typing.RefType
		(CGFontRef)font
	);
}
void * FontRetain(void * font) {
	return (void *)CGFontRetain(
		// *typing.RefType
		(CGFontRef)font
	);
}
void * DisplayCopyColorSpace(CGDirectDisplayID display) {
	return (void *)CGDisplayCopyColorSpace(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
uint ImageGetWidth(void * image) {
	return (uint)CGImageGetWidth(
		// *typing.RefType
		(CGImageRef)image
	);
}
void * PDFDocumentCreateWithURL(void * url) {
	return (void *)CGPDFDocumentCreateWithURL(
		// *typing.RefType
		(CFURLRef)url
	);
}
void EventSetFlags(void * event, CGEventFlags flags) {
	return (void)CGEventSetFlags(
		// *typing.RefType
		(CGEventRef)event,
		// *typing.AliasType
		(CGEventFlags)flags
	);
}
CGRect RectMake(float x, float y, float width, float height) {
	return (CGRect)CGRectMake(
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y,
		// *typing.PrimitiveType
		(CGFloat)width,
		// *typing.PrimitiveType
		(CGFloat)height
	);
}
uint ColorGetNumberOfComponents(void * color) {
	return (uint)CGColorGetNumberOfComponents(
		// *typing.RefType
		(CGColorRef)color
	);
}
void * PathCreateWithRect(CGRect rect, const CGAffineTransform* transform) {
	return (void *)CGPathCreateWithRect(
		// *typing.StructType
		(CGRect)rect,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)transform
	);
}
CGPoint ContextGetPathCurrentPoint(void * c) {
	return (CGPoint)CGContextGetPathCurrentPoint(
		// *typing.RefType
		(CGContextRef)c
	);
}
void PathCloseSubpath(void * path) {
	return (void)CGPathCloseSubpath(
		// *typing.RefType
		(CGMutablePathRef)path
	);
}
void PathAddPath(void * path1, const CGAffineTransform* m, void * path2) {
	return (void)CGPathAddPath(
		// *typing.RefType
		(CGMutablePathRef)path1,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.RefType
		(CGPathRef)path2
	);
}
bool RequestScreenCaptureAccess() {
	return (bool)CGRequestScreenCaptureAccess(
	);
}
bool SizeMakeWithDictionaryRepresentation(void * dict, CGSize* size) {
	return (bool)CGSizeMakeWithDictionaryRepresentation(
		// *typing.RefType
		(CFDictionaryRef)dict,
		// *typing.PointerType
		// -> *typing.StructType
		(CGSize*)size
	);
}
bool ColorSpaceIsWideGamutRGB(void * arg0) {
	return (bool)CGColorSpaceIsWideGamutRGB(
		// *typing.RefType
		(CGColorSpaceRef)arg0
	);
}
void PDFOperatorTableRelease(void * table) {
	return (void)CGPDFOperatorTableRelease(
		// *typing.RefType
		(CGPDFOperatorTableRef)table
	);
}
int FontGetUnitsPerEm(void * font) {
	return (int)CGFontGetUnitsPerEm(
		// *typing.RefType
		(CGFontRef)font
	);
}
void * PDFContentStreamRetain(void * cs) {
	return (void *)CGPDFContentStreamRetain(
		// *typing.RefType
		(CGPDFContentStreamRef)cs
	);
}
void ContextSetAlpha(void * c, float alpha) {
	return (void)CGContextSetAlpha(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
void ContextAddCurveToPoint(void * c, float cp1x, float cp1y, float cp2x, float cp2y, float x, float y) {
	return (void)CGContextAddCurveToPoint(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)cp1x,
		// *typing.PrimitiveType
		(CGFloat)cp1y,
		// *typing.PrimitiveType
		(CGFloat)cp2x,
		// *typing.PrimitiveType
		(CGFloat)cp2y,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y
	);
}
void * ColorCreateGenericCMYK(float cyan, float magenta, float yellow, float black, float alpha) {
	return (void *)CGColorCreateGenericCMYK(
		// *typing.PrimitiveType
		(CGFloat)cyan,
		// *typing.PrimitiveType
		(CGFloat)magenta,
		// *typing.PrimitiveType
		(CGFloat)yellow,
		// *typing.PrimitiveType
		(CGFloat)black,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
void * ColorSpaceCreateExtendedLinearized(void * space) {
	return (void *)CGColorSpaceCreateExtendedLinearized(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void * ColorSpaceCreateWithICCData(CFTypeRef data) {
	return (void *)CGColorSpaceCreateWithICCData(
		// *typing.AliasType
		(CFTypeRef)data
	);
}
CGError GetDisplaysWithRect(CGRect rect, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount) {
	return (CGError)CGGetDisplaysWithRect(
		// *typing.StructType
		(CGRect)rect,
		// *typing.PrimitiveType
		(uint32_t)maxDisplays,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGDirectDisplayID*)displays,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint32_t*)matchingDisplayCount
	);
}
uint DisplayModeGetWidth(void * mode) {
	return (uint)CGDisplayModeGetWidth(
		// *typing.RefType
		(CGDisplayModeRef)mode
	);
}
void ContextSaveGState(void * c) {
	return (void)CGContextSaveGState(
		// *typing.RefType
		(CGContextRef)c
	);
}
int FontGetCapHeight(void * font) {
	return (int)CGFontGetCapHeight(
		// *typing.RefType
		(CGFontRef)font
	);
}
uint PDFArrayGetCount(void * array) {
	return (uint)CGPDFArrayGetCount(
		// *typing.RefType
		(CGPDFArrayRef)array
	);
}
void FunctionRelease(void * function) {
	return (void)CGFunctionRelease(
		// *typing.RefType
		(CGFunctionRef)function
	);
}
bool PDFDocumentUnlockWithPassword(void * document, char* password) {
	return (bool)CGPDFDocumentUnlockWithPassword(
		// *typing.RefType
		(CGPDFDocumentRef)document,
		// *typing.CStringType
		(char*)password
	);
}
double DisplayRotation(CGDirectDisplayID display) {
	return (double)CGDisplayRotation(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
CGAffineTransform AffineTransformMakeScale(float sx, float sy) {
	return (CGAffineTransform)CGAffineTransformMakeScale(
		// *typing.PrimitiveType
		(CGFloat)sx,
		// *typing.PrimitiveType
		(CGFloat)sy
	);
}
void ContextDrawLayerInRect(void * context, CGRect rect, void * layer) {
	return (void)CGContextDrawLayerInRect(
		// *typing.RefType
		(CGContextRef)context,
		// *typing.StructType
		(CGRect)rect,
		// *typing.RefType
		(CGLayerRef)layer
	);
}
bool RectIsInfinite(CGRect rect) {
	return (bool)CGRectIsInfinite(
		// *typing.StructType
		(CGRect)rect
	);
}
void * ImageCreateCopyWithColorSpace(void * image, void * space) {
	return (void *)CGImageCreateCopyWithColorSpace(
		// *typing.RefType
		(CGImageRef)image,
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
bool PreflightPostEventAccess() {
	return (bool)CGPreflightPostEventAccess(
	);
}
void ContextRelease(void * c) {
	return (void)CGContextRelease(
		// *typing.RefType
		(CGContextRef)c
	);
}
CFTypeID DataConsumerGetTypeID() {
	return (CFTypeID)CGDataConsumerGetTypeID(
	);
}
CGError CompleteDisplayConfiguration(void * config, CGConfigureOption option) {
	return (CGError)CGCompleteDisplayConfiguration(
		// *typing.RefType
		(CGDisplayConfigRef)config,
		// *typing.AliasType
		(CGConfigureOption)option
	);
}
void PDFContextEndTag(void * context) {
	return (void)CGPDFContextEndTag(
		// *typing.RefType
		(CGContextRef)context
	);
}
void * ColorCreateGenericGrayGamma2_2(float gray, float alpha) {
	return (void *)CGColorCreateGenericGrayGamma2_2(
		// *typing.PrimitiveType
		(CGFloat)gray,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
CGRect PathGetBoundingBox(void * path) {
	return (CGRect)CGPathGetBoundingBox(
		// *typing.RefType
		(CGPathRef)path
	);
}
void PDFDocumentGetVersion(void * document, int* majorVersion, int* minorVersion) {
	return (void)CGPDFDocumentGetVersion(
		// *typing.RefType
		(CGPDFDocumentRef)document,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(int*)majorVersion,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(int*)minorVersion
	);
}
void PathAddRoundedRect(void * path, const CGAffineTransform* transform, CGRect rect, float cornerWidth, float cornerHeight) {
	return (void)CGPathAddRoundedRect(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)transform,
		// *typing.StructType
		(CGRect)rect,
		// *typing.PrimitiveType
		(CGFloat)cornerWidth,
		// *typing.PrimitiveType
		(CGFloat)cornerHeight
	);
}
CGWindowLevel WindowLevelForKey(CGWindowLevelKey key) {
	return (CGWindowLevel)CGWindowLevelForKey(
		// *typing.AliasType
		(CGWindowLevelKey)key
	);
}
void ContextSetRenderingIntent(void * c, CGColorRenderingIntent intent) {
	return (void)CGContextSetRenderingIntent(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.AliasType
		(CGColorRenderingIntent)intent
	);
}
void ContextSetShadowWithColor(void * c, CGSize offset, float blur, void * color) {
	return (void)CGContextSetShadowWithColor(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGSize)offset,
		// *typing.PrimitiveType
		(CGFloat)blur,
		// *typing.RefType
		(CGColorRef)color
	);
}
CFTypeID PatternGetTypeID() {
	return (CFTypeID)CGPatternGetTypeID(
	);
}
void * DisplayModeRetain(void * mode) {
	return (void *)CGDisplayModeRetain(
		// *typing.RefType
		(CGDisplayModeRef)mode
	);
}
CGEventTimestamp EventGetTimestamp(void * event) {
	return (CGEventTimestamp)CGEventGetTimestamp(
		// *typing.RefType
		(CGEventRef)event
	);
}
void * PathCreateMutableCopy(void * path) {
	return (void *)CGPathCreateMutableCopy(
		// *typing.RefType
		(CGPathRef)path
	);
}
CGAffineTransform ContextGetCTM(void * c) {
	return (CGAffineTransform)CGContextGetCTM(
		// *typing.RefType
		(CGContextRef)c
	);
}
void * EventCreateData(void * allocator, void * event) {
	return (void *)CGEventCreateData(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CGEventRef)event
	);
}
boolean_t DisplayIsStereo(CGDirectDisplayID display) {
	return (boolean_t)CGDisplayIsStereo(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void ContextConcatCTM(void * c, CGAffineTransform transform) {
	return (void)CGContextConcatCTM(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGAffineTransform)transform
	);
}
CGSize ContextConvertSizeToUserSpace(void * c, CGSize size) {
	return (CGSize)CGContextConvertSizeToUserSpace(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGSize)size
	);
}
bool PDFScannerPopInteger(void * scanner, CGPDFInteger* value) {
	return (bool)CGPDFScannerPopInteger(
		// *typing.RefType
		(CGPDFScannerRef)scanner,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGPDFInteger*)value
	);
}
void PDFContextAddDocumentMetadata(void * context, void * metadata) {
	return (void)CGPDFContextAddDocumentMetadata(
		// *typing.RefType
		(CGContextRef)context,
		// *typing.RefType
		(CFDataRef)metadata
	);
}
CGSize DisplayScreenSize(CGDirectDisplayID display) {
	return (CGSize)CGDisplayScreenSize(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void ColorSpaceGetColorTable(void * space, uint8_t* table) {
	return (void)CGColorSpaceGetColorTable(
		// *typing.RefType
		(CGColorSpaceRef)space,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)table
	);
}
void * PDFPageRetain(void * page) {
	return (void *)CGPDFPageRetain(
		// *typing.RefType
		(CGPDFPageRef)page
	);
}
void ContextFillEllipseInRect(void * c, CGRect rect) {
	return (void)CGContextFillEllipseInRect(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect
	);
}
void * DisplayCopyDisplayMode(CGDirectDisplayID display) {
	return (void *)CGDisplayCopyDisplayMode(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void PDFContextSetDestinationForRect(void * context, void * name, CGRect rect) {
	return (void)CGPDFContextSetDestinationForRect(
		// *typing.RefType
		(CGContextRef)context,
		// *typing.RefType
		(CFStringRef)name,
		// *typing.StructType
		(CGRect)rect
	);
}
CGError GetDisplaysWithOpenGLDisplayMask(CGOpenGLDisplayMask mask, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount) {
	return (CGError)CGGetDisplaysWithOpenGLDisplayMask(
		// *typing.AliasType
		(CGOpenGLDisplayMask)mask,
		// *typing.PrimitiveType
		(uint32_t)maxDisplays,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGDirectDisplayID*)displays,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint32_t*)matchingDisplayCount
	);
}
void * ShadingCreateAxial(void * space, CGPoint start, CGPoint end, void * function, bool extendStart, bool extendEnd) {
	return (void *)CGShadingCreateAxial(
		// *typing.RefType
		(CGColorSpaceRef)space,
		// *typing.StructType
		(CGPoint)start,
		// *typing.StructType
		(CGPoint)end,
		// *typing.RefType
		(CGFunctionRef)function,
		// *typing.PrimitiveType
		(BOOL)extendStart,
		// *typing.PrimitiveType
		(BOOL)extendEnd
	);
}
void ContextSetShadow(void * c, CGSize offset, float blur) {
	return (void)CGContextSetShadow(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGSize)offset,
		// *typing.PrimitiveType
		(CGFloat)blur
	);
}
bool PointMakeWithDictionaryRepresentation(void * dict, CGPoint* point) {
	return (bool)CGPointMakeWithDictionaryRepresentation(
		// *typing.RefType
		(CFDictionaryRef)dict,
		// *typing.PointerType
		// -> *typing.StructType
		(CGPoint*)point
	);
}
CGPoint ContextGetTextPosition(void * c) {
	return (CGPoint)CGContextGetTextPosition(
		// *typing.RefType
		(CGContextRef)c
	);
}
uint BitmapContextGetBitsPerPixel(void * context) {
	return (uint)CGBitmapContextGetBitsPerPixel(
		// *typing.RefType
		(CGContextRef)context
	);
}
void ContextAddArc(void * c, float x, float y, float radius, float startAngle, float endAngle, int clockwise) {
	return (void)CGContextAddArc(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y,
		// *typing.PrimitiveType
		(CGFloat)radius,
		// *typing.PrimitiveType
		(CGFloat)startAngle,
		// *typing.PrimitiveType
		(CGFloat)endAngle,
		// *typing.PrimitiveType
		(int)clockwise
	);
}
void EventSetType(void * event, CGEventType type_) {
	return (void)CGEventSetType(
		// *typing.RefType
		(CGEventRef)event,
		// *typing.AliasType
		(CGEventType)type_
	);
}
void * EventCreateScrollWheelEvent2(void * source, CGScrollEventUnit units, uint32_t wheelCount, int32_t wheel1, int32_t wheel2, int32_t wheel3) {
	return (void *)CGEventCreateScrollWheelEvent2(
		// *typing.RefType
		(CGEventSourceRef)source,
		// *typing.AliasType
		(CGScrollEventUnit)units,
		// *typing.PrimitiveType
		(uint32_t)wheelCount,
		// *typing.PrimitiveType
		(int32_t)wheel1,
		// *typing.PrimitiveType
		(int32_t)wheel2,
		// *typing.PrimitiveType
		(int32_t)wheel3
	);
}
CGColorSpaceModel ColorSpaceGetModel(void * space) {
	return (CGColorSpaceModel)CGColorSpaceGetModel(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void * ImageCreateWithPNGDataProvider(void * source, const float* decode, bool shouldInterpolate, CGColorRenderingIntent intent) {
	return (void *)CGImageCreateWithPNGDataProvider(
		// *typing.RefType
		(CGDataProviderRef)source,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)decode,
		// *typing.PrimitiveType
		(BOOL)shouldInterpolate,
		// *typing.AliasType
		(CGColorRenderingIntent)intent
	);
}
void PathAddEllipseInRect(void * path, const CGAffineTransform* m, CGRect rect) {
	return (void)CGPathAddEllipseInRect(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.StructType
		(CGRect)rect
	);
}
void * EventCreateFromData(void * allocator, void * data) {
	return (void *)CGEventCreateFromData(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFDataRef)data
	);
}
void * ColorSpaceCreateDeviceCMYK() {
	return (void *)CGColorSpaceCreateDeviceCMYK(
	);
}
CGError ConfigureDisplayFadeEffect(void * config, CGDisplayFadeInterval fadeOutSeconds, CGDisplayFadeInterval fadeInSeconds, float fadeRed, float fadeGreen, float fadeBlue) {
	return (CGError)CGConfigureDisplayFadeEffect(
		// *typing.RefType
		(CGDisplayConfigRef)config,
		// *typing.AliasType
		(CGDisplayFadeInterval)fadeOutSeconds,
		// *typing.AliasType
		(CGDisplayFadeInterval)fadeInSeconds,
		// *typing.PrimitiveType
		(float)fadeRed,
		// *typing.PrimitiveType
		(float)fadeGreen,
		// *typing.PrimitiveType
		(float)fadeBlue
	);
}
CGError DisplayFade(CGDisplayFadeReservationToken token, CGDisplayFadeInterval duration, CGDisplayBlendFraction startBlend, CGDisplayBlendFraction endBlend, float redBlend, float greenBlend, float blueBlend, boolean_t synchronous) {
	return (CGError)CGDisplayFade(
		// *typing.AliasType
		(CGDisplayFadeReservationToken)token,
		// *typing.AliasType
		(CGDisplayFadeInterval)duration,
		// *typing.AliasType
		(CGDisplayBlendFraction)startBlend,
		// *typing.AliasType
		(CGDisplayBlendFraction)endBlend,
		// *typing.PrimitiveType
		(float)redBlend,
		// *typing.PrimitiveType
		(float)greenBlend,
		// *typing.PrimitiveType
		(float)blueBlend,
		// *typing.AliasType
		(boolean_t)synchronous
	);
}
void ContextSetShouldSubpixelQuantizeFonts(void * c, bool shouldSubpixelQuantizeFonts) {
	return (void)CGContextSetShouldSubpixelQuantizeFonts(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(BOOL)shouldSubpixelQuantizeFonts
	);
}
CGPoint EventGetLocation(void * event) {
	return (CGPoint)CGEventGetLocation(
		// *typing.RefType
		(CGEventRef)event
	);
}
bool PSConverterAbort(void * converter) {
	return (bool)CGPSConverterAbort(
		// *typing.RefType
		(CGPSConverterRef)converter
	);
}
void ContextScaleCTM(void * c, float sx, float sy) {
	return (void)CGContextScaleCTM(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)sx,
		// *typing.PrimitiveType
		(CGFloat)sy
	);
}
void * SizeCreateDictionaryRepresentation(CGSize size) {
	return (void *)CGSizeCreateDictionaryRepresentation(
		// *typing.StructType
		(CGSize)size
	);
}
bool EventSourceButtonState(CGEventSourceStateID stateID, CGMouseButton button) {
	return (bool)CGEventSourceButtonState(
		// *typing.AliasType
		(CGEventSourceStateID)stateID,
		// *typing.AliasType
		(CGMouseButton)button
	);
}
void EventTapEnable(void * tap, bool enable) {
	return (void)CGEventTapEnable(
		// *typing.RefType
		(CFMachPortRef)tap,
		// *typing.PrimitiveType
		(BOOL)enable
	);
}
CGAffineTransform AffineTransformMakeRotation(float angle) {
	return (CGAffineTransform)CGAffineTransformMakeRotation(
		// *typing.PrimitiveType
		(CGFloat)angle
	);
}
void * ColorSpaceCreateICCBased(uint nComponents, const float* range_, void * profile, void * alternate) {
	return (void *)CGColorSpaceCreateICCBased(
		// *typing.PrimitiveType
		(NSUInteger)nComponents,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)range_,
		// *typing.RefType
		(CGDataProviderRef)profile,
		// *typing.RefType
		(CGColorSpaceRef)alternate
	);
}
CGError AssociateMouseAndMouseCursorPosition(boolean_t connected) {
	return (CGError)CGAssociateMouseAndMouseCursorPosition(
		// *typing.AliasType
		(boolean_t)connected
	);
}
void * PDFDocumentRetain(void * document) {
	return (void *)CGPDFDocumentRetain(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
CGPoint PathGetCurrentPoint(void * path) {
	return (CGPoint)CGPathGetCurrentPoint(
		// *typing.RefType
		(CGPathRef)path
	);
}
void ContextSetRGBStrokeColor(void * c, float red, float green, float blue, float alpha) {
	return (void)CGContextSetRGBStrokeColor(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)red,
		// *typing.PrimitiveType
		(CGFloat)green,
		// *typing.PrimitiveType
		(CGFloat)blue,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
bool PreflightScreenCaptureAccess() {
	return (bool)CGPreflightScreenCaptureAccess(
	);
}
CGError SetDisplayTransferByTable(CGDirectDisplayID display, uint32_t tableSize, const CGGammaValue* redTable, const CGGammaValue* greenTable, const CGGammaValue* blueTable) {
	return (CGError)CGSetDisplayTransferByTable(
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.PrimitiveType
		(uint32_t)tableSize,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)redTable,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)greenTable,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)blueTable
	);
}
void ContextSetFlatness(void * c, float flatness) {
	return (void)CGContextSetFlatness(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)flatness
	);
}
boolean_t DisplayIsActive(CGDirectDisplayID display) {
	return (boolean_t)CGDisplayIsActive(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void ContextSetPatternPhase(void * c, CGSize phase) {
	return (void)CGContextSetPatternPhase(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGSize)phase
	);
}
void PDFScannerRelease(void * scanner) {
	return (void)CGPDFScannerRelease(
		// *typing.RefType
		(CGPDFScannerRef)scanner
	);
}
void ContextSetLineJoin(void * c, CGLineJoin join) {
	return (void)CGContextSetLineJoin(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.AliasType
		(CGLineJoin)join
	);
}
void * FontCopyTableTags(void * font) {
	return (void *)CGFontCopyTableTags(
		// *typing.RefType
		(CGFontRef)font
	);
}
CGError GetActiveDisplayList(uint32_t maxDisplays, CGDirectDisplayID* activeDisplays, uint32_t* displayCount) {
	return (CGError)CGGetActiveDisplayList(
		// *typing.PrimitiveType
		(uint32_t)maxDisplays,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGDirectDisplayID*)activeDisplays,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint32_t*)displayCount
	);
}
bool PathIsRect(void * path, CGRect* rect) {
	return (bool)CGPathIsRect(
		// *typing.RefType
		(CGPathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)rect
	);
}
void * DataProviderCreateWithURL(void * url) {
	return (void *)CGDataProviderCreateWithURL(
		// *typing.RefType
		(CFURLRef)url
	);
}
void * EventSourceCreate(CGEventSourceStateID stateID) {
	return (void *)CGEventSourceCreate(
		// *typing.AliasType
		(CGEventSourceStateID)stateID
	);
}
CGInterpolationQuality ContextGetInterpolationQuality(void * c) {
	return (CGInterpolationQuality)CGContextGetInterpolationQuality(
		// *typing.RefType
		(CGContextRef)c
	);
}
void * ColorConversionInfoCreate(void * src, void * dst) {
	return (void *)CGColorConversionInfoCreate(
		// *typing.RefType
		(CGColorSpaceRef)src,
		// *typing.RefType
		(CGColorSpaceRef)dst
	);
}
bool RectContainsRect(CGRect rect1, CGRect rect2) {
	return (bool)CGRectContainsRect(
		// *typing.StructType
		(CGRect)rect1,
		// *typing.StructType
		(CGRect)rect2
	);
}
void * FontCreatePostScriptEncoding(void * font, const CGGlyph* encoding) {
	return (void *)CGFontCreatePostScriptEncoding(
		// *typing.RefType
		(CGFontRef)font,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGlyph*)encoding
	);
}
void ContextDrawLayerAtPoint(void * context, CGPoint point, void * layer) {
	return (void)CGContextDrawLayerAtPoint(
		// *typing.RefType
		(CGContextRef)context,
		// *typing.StructType
		(CGPoint)point,
		// *typing.RefType
		(CGLayerRef)layer
	);
}
void EventSourceSetPixelsPerLine(void * source, double pixelsPerLine) {
	return (void)CGEventSourceSetPixelsPerLine(
		// *typing.RefType
		(CGEventSourceRef)source,
		// *typing.PrimitiveType
		(double)pixelsPerLine
	);
}
void * EventCreateKeyboardEvent(void * source, CGKeyCode virtualKey, bool keyDown) {
	return (void *)CGEventCreateKeyboardEvent(
		// *typing.RefType
		(CGEventSourceRef)source,
		// *typing.AliasType
		(CGKeyCode)virtualKey,
		// *typing.PrimitiveType
		(BOOL)keyDown
	);
}
uint32_t EventSourceCounterForEventType(CGEventSourceStateID stateID, CGEventType eventType) {
	return (uint32_t)CGEventSourceCounterForEventType(
		// *typing.AliasType
		(CGEventSourceStateID)stateID,
		// *typing.AliasType
		(CGEventType)eventType
	);
}
void PathAddRects(void * path, const CGAffineTransform* m, const CGRect* rects, uint count) {
	return (void)CGPathAddRects(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)rects,
		// *typing.PrimitiveType
		(NSUInteger)count
	);
}
void EventSetTimestamp(void * event, CGEventTimestamp timestamp) {
	return (void)CGEventSetTimestamp(
		// *typing.RefType
		(CGEventRef)event,
		// *typing.AliasType
		(CGEventTimestamp)timestamp
	);
}
bool EventSourceKeyState(CGEventSourceStateID stateID, CGKeyCode key) {
	return (bool)CGEventSourceKeyState(
		// *typing.AliasType
		(CGEventSourceStateID)stateID,
		// *typing.AliasType
		(CGKeyCode)key
	);
}
float RectGetHeight(CGRect rect) {
	return (float)CGRectGetHeight(
		// *typing.StructType
		(CGRect)rect
	);
}
bool PDFDictionaryGetInteger(void * dict, char* key, CGPDFInteger* value) {
	return (bool)CGPDFDictionaryGetInteger(
		// *typing.RefType
		(CGPDFDictionaryRef)dict,
		// *typing.CStringType
		(char*)key,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGPDFInteger*)value
	);
}
CFTimeInterval EventSourceSecondsSinceLastEventType(CGEventSourceStateID stateID, CGEventType eventType) {
	return (CFTimeInterval)CGEventSourceSecondsSinceLastEventType(
		// *typing.AliasType
		(CGEventSourceStateID)stateID,
		// *typing.AliasType
		(CGEventType)eventType
	);
}
CGRect RectIntegral(CGRect rect) {
	return (CGRect)CGRectIntegral(
		// *typing.StructType
		(CGRect)rect
	);
}
void ContextSetFillColorSpace(void * c, void * space) {
	return (void)CGContextSetFillColorSpace(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
bool PDFArrayGetNumber(void * array, uint index, CGPDFReal* value) {
	return (bool)CGPDFArrayGetNumber(
		// *typing.RefType
		(CGPDFArrayRef)array,
		// *typing.PrimitiveType
		(NSUInteger)index,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGPDFReal*)value
	);
}
void * ColorCreateWithPattern(void * space, void * pattern, const float* components) {
	return (void *)CGColorCreateWithPattern(
		// *typing.RefType
		(CGColorSpaceRef)space,
		// *typing.RefType
		(CGPatternRef)pattern,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)components
	);
}
void ContextTranslateCTM(void * c, float tx, float ty) {
	return (void)CGContextTranslateCTM(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)tx,
		// *typing.PrimitiveType
		(CGFloat)ty
	);
}
boolean_t DisplayUsesOpenGLAcceleration(CGDirectDisplayID display) {
	return (boolean_t)CGDisplayUsesOpenGLAcceleration(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void ContextSetBlendMode(void * c, CGBlendMode mode) {
	return (void)CGContextSetBlendMode(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.AliasType
		(CGBlendMode)mode
	);
}
void GetLastMouseDelta(int32_t* deltaX, int32_t* deltaY) {
	return (void)CGGetLastMouseDelta(
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(int32_t*)deltaX,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(int32_t*)deltaY
	);
}
bool FontGetGlyphBBoxes(void * font, const CGGlyph* glyphs, uint count, CGRect* bboxes) {
	return (bool)CGFontGetGlyphBBoxes(
		// *typing.RefType
		(CGFontRef)font,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGlyph*)glyphs,
		// *typing.PrimitiveType
		(NSUInteger)count,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)bboxes
	);
}
CGError GetDisplaysWithPoint(CGPoint point, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount) {
	return (CGError)CGGetDisplaysWithPoint(
		// *typing.StructType
		(CGPoint)point,
		// *typing.PrimitiveType
		(uint32_t)maxDisplays,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGDirectDisplayID*)displays,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint32_t*)matchingDisplayCount
	);
}
void ImageRelease(void * image) {
	return (void)CGImageRelease(
		// *typing.RefType
		(CGImageRef)image
	);
}
float RectGetMinX(CGRect rect) {
	return (float)CGRectGetMinX(
		// *typing.StructType
		(CGRect)rect
	);
}
void ContextReplacePathWithStrokedPath(void * c) {
	return (void)CGContextReplacePathWithStrokedPath(
		// *typing.RefType
		(CGContextRef)c
	);
}
uint BitmapContextGetBytesPerRow(void * context) {
	return (uint)CGBitmapContextGetBytesPerRow(
		// *typing.RefType
		(CGContextRef)context
	);
}
void * WindowListCreate(CGWindowListOption option, CGWindowID relativeToWindow) {
	return (void *)CGWindowListCreate(
		// *typing.AliasType
		(CGWindowListOption)option,
		// *typing.AliasType
		(CGWindowID)relativeToWindow
	);
}
CGSize SizeApplyAffineTransform(CGSize size, CGAffineTransform t) {
	return (CGSize)CGSizeApplyAffineTransform(
		// *typing.StructType
		(CGSize)size,
		// *typing.StructType
		(CGAffineTransform)t
	);
}
void EventTapPostEvent(CGEventTapProxy proxy, void * event) {
	return (void)CGEventTapPostEvent(
		// *typing.AliasType
		(CGEventTapProxy)proxy,
		// *typing.RefType
		(CGEventRef)event
	);
}
boolean_t DisplayIsInMirrorSet(CGDirectDisplayID display) {
	return (boolean_t)CGDisplayIsInMirrorSet(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
bool ColorSpaceIsPQBased(void * s) {
	return (bool)CGColorSpaceIsPQBased(
		// *typing.RefType
		(CGColorSpaceRef)s
	);
}
void ContextClearRect(void * c, CGRect rect) {
	return (void)CGContextClearRect(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect
	);
}
CGError DisplayRelease(CGDirectDisplayID display) {
	return (CGError)CGDisplayRelease(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void * PDFPageGetDocument(void * page) {
	return (void *)CGPDFPageGetDocument(
		// *typing.RefType
		(CGPDFPageRef)page
	);
}
CGError GetDisplayTransferByTable(CGDirectDisplayID display, uint32_t capacity, CGGammaValue* redTable, CGGammaValue* greenTable, CGGammaValue* blueTable, uint32_t* sampleCount) {
	return (CGError)CGGetDisplayTransferByTable(
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.PrimitiveType
		(uint32_t)capacity,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)redTable,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)greenTable,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)blueTable,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint32_t*)sampleCount
	);
}
bool AffineTransformEqualToTransform(CGAffineTransform t1, CGAffineTransform t2) {
	return (bool)CGAffineTransformEqualToTransform(
		// *typing.StructType
		(CGAffineTransform)t1,
		// *typing.StructType
		(CGAffineTransform)t2
	);
}
void * PDFScannerGetContentStream(void * scanner) {
	return (void *)CGPDFScannerGetContentStream(
		// *typing.RefType
		(CGPDFScannerRef)scanner
	);
}
void PathAddLineToPoint(void * path, const CGAffineTransform* m, float x, float y) {
	return (void)CGPathAddLineToPoint(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y
	);
}
bool PreflightListenEventAccess() {
	return (bool)CGPreflightListenEventAccess(
	);
}
CGEventType EventGetType(void * event) {
	return (CGEventType)CGEventGetType(
		// *typing.RefType
		(CGEventRef)event
	);
}
bool PDFScannerPopNumber(void * scanner, CGPDFReal* value) {
	return (bool)CGPDFScannerPopNumber(
		// *typing.RefType
		(CGPDFScannerRef)scanner,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGPDFReal*)value
	);
}
void * ColorSpaceCreateLab(const float* whitePoint, const float* blackPoint, const float* range_) {
	return (void *)CGColorSpaceCreateLab(
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)whitePoint,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)blackPoint,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)range_
	);
}
void PDFContextEndPage(void * context) {
	return (void)CGPDFContextEndPage(
		// *typing.RefType
		(CGContextRef)context
	);
}
CGRect ContextGetClipBoundingBox(void * c) {
	return (CGRect)CGContextGetClipBoundingBox(
		// *typing.RefType
		(CGContextRef)c
	);
}
void * PDFOperatorTableCreate() {
	return (void *)CGPDFOperatorTableCreate(
	);
}
void PathAddRect(void * path, const CGAffineTransform* m, CGRect rect) {
	return (void)CGPathAddRect(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.StructType
		(CGRect)rect
	);
}
CGPoint PointApplyAffineTransform(CGPoint point, CGAffineTransform t) {
	return (CGPoint)CGPointApplyAffineTransform(
		// *typing.StructType
		(CGPoint)point,
		// *typing.StructType
		(CGAffineTransform)t
	);
}
void * ImageCreateWithJPEGDataProvider(void * source, const float* decode, bool shouldInterpolate, CGColorRenderingIntent intent) {
	return (void *)CGImageCreateWithJPEGDataProvider(
		// *typing.RefType
		(CGDataProviderRef)source,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)decode,
		// *typing.PrimitiveType
		(BOOL)shouldInterpolate,
		// *typing.AliasType
		(CGColorRenderingIntent)intent
	);
}
CFTypeID EventSourceGetTypeID() {
	return (CFTypeID)CGEventSourceGetTypeID(
	);
}
void * ImageCreate(uint width, uint height, uint bitsPerComponent, uint bitsPerPixel, uint bytesPerRow, void * space, CGBitmapInfo bitmapInfo, void * provider, const float* decode, bool shouldInterpolate, CGColorRenderingIntent intent) {
	return (void *)CGImageCreate(
		// *typing.PrimitiveType
		(NSUInteger)width,
		// *typing.PrimitiveType
		(NSUInteger)height,
		// *typing.PrimitiveType
		(NSUInteger)bitsPerComponent,
		// *typing.PrimitiveType
		(NSUInteger)bitsPerPixel,
		// *typing.PrimitiveType
		(NSUInteger)bytesPerRow,
		// *typing.RefType
		(CGColorSpaceRef)space,
		// *typing.AliasType
		(CGBitmapInfo)bitmapInfo,
		// *typing.RefType
		(CGDataProviderRef)provider,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)decode,
		// *typing.PrimitiveType
		(BOOL)shouldInterpolate,
		// *typing.AliasType
		(CGColorRenderingIntent)intent
	);
}
bool RectIsNull(CGRect rect) {
	return (bool)CGRectIsNull(
		// *typing.StructType
		(CGRect)rect
	);
}
CGEventSourceStateID EventSourceGetSourceStateID(void * source) {
	return (CGEventSourceStateID)CGEventSourceGetSourceStateID(
		// *typing.RefType
		(CGEventSourceRef)source
	);
}
void * GradientRetain(void * gradient) {
	return (void *)CGGradientRetain(
		// *typing.RefType
		(CGGradientRef)gradient
	);
}
bool PDFArrayGetBoolean(void * array, uint index, CGPDFBoolean* value) {
	return (bool)CGPDFArrayGetBoolean(
		// *typing.RefType
		(CGPDFArrayRef)array,
		// *typing.PrimitiveType
		(NSUInteger)index,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGPDFBoolean*)value
	);
}
void ContextSetFillColor(void * c, const float* components) {
	return (void)CGContextSetFillColor(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)components
	);
}
void * ColorGetPattern(void * color) {
	return (void *)CGColorGetPattern(
		// *typing.RefType
		(CGColorRef)color
	);
}
CGRect PathGetPathBoundingBox(void * path) {
	return (CGRect)CGPathGetPathBoundingBox(
		// *typing.RefType
		(CGPathRef)path
	);
}
CGError CaptureAllDisplaysWithOptions(CGCaptureOptions options) {
	return (CGError)CGCaptureAllDisplaysWithOptions(
		// *typing.AliasType
		(CGCaptureOptions)options
	);
}
void * WindowListCreateDescriptionFromArray(void * windowArray) {
	return (void *)CGWindowListCreateDescriptionFromArray(
		// *typing.RefType
		(CFArrayRef)windowArray
	);
}
void * PDFContentStreamCreateWithStream(void * stream, void * streamResources, void * parent) {
	return (void *)CGPDFContentStreamCreateWithStream(
		// *typing.RefType
		(CGPDFStreamRef)stream,
		// *typing.RefType
		(CGPDFDictionaryRef)streamResources,
		// *typing.RefType
		(CGPDFContentStreamRef)parent
	);
}
void ContextDrawRadialGradient(void * c, void * gradient, CGPoint startCenter, float startRadius, CGPoint endCenter, float endRadius, CGGradientDrawingOptions options) {
	return (void)CGContextDrawRadialGradient(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGGradientRef)gradient,
		// *typing.StructType
		(CGPoint)startCenter,
		// *typing.PrimitiveType
		(CGFloat)startRadius,
		// *typing.StructType
		(CGPoint)endCenter,
		// *typing.PrimitiveType
		(CGFloat)endRadius,
		// *typing.AliasType
		(CGGradientDrawingOptions)options
	);
}
void * PDFStringCopyDate(void * string_) {
	return (void *)CGPDFStringCopyDate(
		// *typing.RefType
		(CGPDFStringRef)string_
	);
}
CGError GetOnlineDisplayList(uint32_t maxDisplays, CGDirectDisplayID* onlineDisplays, uint32_t* displayCount) {
	return (CGError)CGGetOnlineDisplayList(
		// *typing.PrimitiveType
		(uint32_t)maxDisplays,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGDirectDisplayID*)onlineDisplays,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint32_t*)displayCount
	);
}
CGPoint ContextConvertPointToDeviceSpace(void * c, CGPoint point) {
	return (CGPoint)CGContextConvertPointToDeviceSpace(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGPoint)point
	);
}
void * PathCreateCopy(void * path) {
	return (void *)CGPathCreateCopy(
		// *typing.RefType
		(CGPathRef)path
	);
}
void * BitmapContextCreateImage(void * context) {
	return (void *)CGBitmapContextCreateImage(
		// *typing.RefType
		(CGContextRef)context
	);
}
uint ColorSpaceGetNumberOfComponents(void * space) {
	return (uint)CGColorSpaceGetNumberOfComponents(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void * FunctionRetain(void * function) {
	return (void *)CGFunctionRetain(
		// *typing.RefType
		(CGFunctionRef)function
	);
}
CGError GetDisplayTransferByFormula(CGDirectDisplayID display, CGGammaValue* redMin, CGGammaValue* redMax, CGGammaValue* redGamma, CGGammaValue* greenMin, CGGammaValue* greenMax, CGGammaValue* greenGamma, CGGammaValue* blueMin, CGGammaValue* blueMax, CGGammaValue* blueGamma) {
	return (CGError)CGGetDisplayTransferByFormula(
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)redMin,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)redMax,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)redGamma,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)greenMin,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)greenMax,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)greenGamma,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)blueMin,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)blueMax,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGammaValue*)blueGamma
	);
}
void * ColorSpaceRetain(void * space) {
	return (void *)CGColorSpaceRetain(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void * ColorCreateSRGB(float red, float green, float blue, float alpha) {
	return (void *)CGColorCreateSRGB(
		// *typing.PrimitiveType
		(CGFloat)red,
		// *typing.PrimitiveType
		(CGFloat)green,
		// *typing.PrimitiveType
		(CGFloat)blue,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
void * ImageGetColorSpace(void * image) {
	return (void *)CGImageGetColorSpace(
		// *typing.RefType
		(CGImageRef)image
	);
}
void * ShadingCreateRadial(void * space, CGPoint start, float startRadius, CGPoint end, float endRadius, void * function, bool extendStart, bool extendEnd) {
	return (void *)CGShadingCreateRadial(
		// *typing.RefType
		(CGColorSpaceRef)space,
		// *typing.StructType
		(CGPoint)start,
		// *typing.PrimitiveType
		(CGFloat)startRadius,
		// *typing.StructType
		(CGPoint)end,
		// *typing.PrimitiveType
		(CGFloat)endRadius,
		// *typing.RefType
		(CGFunctionRef)function,
		// *typing.PrimitiveType
		(BOOL)extendStart,
		// *typing.PrimitiveType
		(BOOL)extendEnd
	);
}
void RectDivide(CGRect rect, CGRect* slice, CGRect* remainder, float amount, CGRectEdge edge) {
	return (void)CGRectDivide(
		// *typing.StructType
		(CGRect)rect,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)slice,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)remainder,
		// *typing.PrimitiveType
		(CGFloat)amount,
		// *typing.AliasType
		(CGRectEdge)edge
	);
}
void ContextSetAllowsAntialiasing(void * c, bool allowsAntialiasing) {
	return (void)CGContextSetAllowsAntialiasing(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(BOOL)allowsAntialiasing
	);
}
uint BitmapContextGetHeight(void * context) {
	return (uint)CGBitmapContextGetHeight(
		// *typing.RefType
		(CGContextRef)context
	);
}
float ColorGetAlpha(void * color) {
	return (float)CGColorGetAlpha(
		// *typing.RefType
		(CGColorRef)color
	);
}
uint ImageGetHeight(void * image) {
	return (uint)CGImageGetHeight(
		// *typing.RefType
		(CGImageRef)image
	);
}
void * ColorSpaceCreateWithName(void * name) {
	return (void *)CGColorSpaceCreateWithName(
		// *typing.RefType
		(CFStringRef)name
	);
}
void ContextAddQuadCurveToPoint(void * c, float cpx, float cpy, float x, float y) {
	return (void)CGContextAddQuadCurveToPoint(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)cpx,
		// *typing.PrimitiveType
		(CGFloat)cpy,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y
	);
}
CGError ConfigureDisplayWithDisplayMode(void * config, CGDirectDisplayID display, void * mode, void * options) {
	return (CGError)CGConfigureDisplayWithDisplayMode(
		// *typing.RefType
		(CGDisplayConfigRef)config,
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.RefType
		(CGDisplayModeRef)mode,
		// *typing.RefType
		(CFDictionaryRef)options
	);
}
void * DataProviderCopyData(void * provider) {
	return (void *)CGDataProviderCopyData(
		// *typing.RefType
		(CGDataProviderRef)provider
	);
}
boolean_t DisplayIsInHWMirrorSet(CGDirectDisplayID display) {
	return (boolean_t)CGDisplayIsInHWMirrorSet(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
CGEventFilterMask EventSourceGetLocalEventsFilterDuringSuppressionState(void * source, CGEventSuppressionState state) {
	return (CGEventFilterMask)CGEventSourceGetLocalEventsFilterDuringSuppressionState(
		// *typing.RefType
		(CGEventSourceRef)source,
		// *typing.AliasType
		(CGEventSuppressionState)state
	);
}
void * FontCopyTableForTag(void * font, uint32_t tag) {
	return (void *)CGFontCopyTableForTag(
		// *typing.RefType
		(CGFontRef)font,
		// *typing.PrimitiveType
		(uint32_t)tag
	);
}
char* PDFStringGetBytePtr(void * string_) {
	return (char*)CGPDFStringGetBytePtr(
		// *typing.RefType
		(CGPDFStringRef)string_
	);
}
CGRect ContextGetPathBoundingBox(void * c) {
	return (CGRect)CGContextGetPathBoundingBox(
		// *typing.RefType
		(CGContextRef)c
	);
}
void * LayerRetain(void * layer) {
	return (void *)CGLayerRetain(
		// *typing.RefType
		(CGLayerRef)layer
	);
}
void ContextSetAllowsFontSubpixelPositioning(void * c, bool allowsFontSubpixelPositioning) {
	return (void)CGContextSetAllowsFontSubpixelPositioning(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(BOOL)allowsFontSubpixelPositioning
	);
}
CFTypeID ColorGetTypeID() {
	return (CFTypeID)CGColorGetTypeID(
	);
}
CGImageByteOrderInfo ImageGetByteOrderInfo(void * image) {
	return (CGImageByteOrderInfo)CGImageGetByteOrderInfo(
		// *typing.RefType
		(CGImageRef)image
	);
}
void * PDFDocumentCreateWithProvider(void * provider) {
	return (void *)CGPDFDocumentCreateWithProvider(
		// *typing.RefType
		(CGDataProviderRef)provider
	);
}
bool FontGetGlyphAdvances(void * font, const CGGlyph* glyphs, uint count, int* advances) {
	return (bool)CGFontGetGlyphAdvances(
		// *typing.RefType
		(CGFontRef)font,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGlyph*)glyphs,
		// *typing.PrimitiveType
		(NSUInteger)count,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(int*)advances
	);
}
CGImagePixelFormatInfo ImageGetPixelFormatInfo(void * image) {
	return (CGImagePixelFormatInfo)CGImageGetPixelFormatInfo(
		// *typing.RefType
		(CGImageRef)image
	);
}
void * FontCopyVariationAxes(void * font) {
	return (void *)CGFontCopyVariationAxes(
		// *typing.RefType
		(CGFontRef)font
	);
}
CFPropertyListRef ColorSpaceCopyPropertyList(void * space) {
	return (CFPropertyListRef)CGColorSpaceCopyPropertyList(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void ContextSetStrokePattern(void * c, void * pattern, const float* components) {
	return (void)CGContextSetStrokePattern(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGPatternRef)pattern,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)components
	);
}
void ContextClipToRect(void * c, CGRect rect) {
	return (void)CGContextClipToRect(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect
	);
}
void ContextResetClip(void * c) {
	return (void)CGContextResetClip(
		// *typing.RefType
		(CGContextRef)c
	);
}
void PathAddQuadCurveToPoint(void * path, const CGAffineTransform* m, float cpx, float cpy, float x, float y) {
	return (void)CGPathAddQuadCurveToPoint(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.PrimitiveType
		(CGFloat)cpx,
		// *typing.PrimitiveType
		(CGFloat)cpy,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y
	);
}
CGError WarpMouseCursorPosition(CGPoint newCursorPosition) {
	return (CGError)CGWarpMouseCursorPosition(
		// *typing.StructType
		(CGPoint)newCursorPosition
	);
}
CGError DisplayMoveCursorToPoint(CGDirectDisplayID display, CGPoint point) {
	return (CGError)CGDisplayMoveCursorToPoint(
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.StructType
		(CGPoint)point
	);
}
void * WindowListCopyWindowInfo(CGWindowListOption option, CGWindowID relativeToWindow) {
	return (void *)CGWindowListCopyWindowInfo(
		// *typing.AliasType
		(CGWindowListOption)option,
		// *typing.AliasType
		(CGWindowID)relativeToWindow
	);
}
CGWindowID ShieldingWindowID(CGDirectDisplayID display) {
	return (CGWindowID)CGShieldingWindowID(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void PDFContextSetOutline(void * context, void * outline) {
	return (void)CGPDFContextSetOutline(
		// *typing.RefType
		(CGContextRef)context,
		// *typing.RefType
		(CFDictionaryRef)outline
	);
}
void PathAddArcToPoint(void * path, const CGAffineTransform* m, float x1, float y1, float x2, float y2, float radius) {
	return (void)CGPathAddArcToPoint(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.PrimitiveType
		(CGFloat)x1,
		// *typing.PrimitiveType
		(CGFloat)y1,
		// *typing.PrimitiveType
		(CGFloat)x2,
		// *typing.PrimitiveType
		(CGFloat)y2,
		// *typing.PrimitiveType
		(CGFloat)radius
	);
}
uint ImageGetBitsPerComponent(void * image) {
	return (uint)CGImageGetBitsPerComponent(
		// *typing.RefType
		(CGImageRef)image
	);
}
CGError CaptureAllDisplays() {
	return (CGError)CGCaptureAllDisplays(
	);
}
double EventSourceGetPixelsPerLine(void * source) {
	return (double)CGEventSourceGetPixelsPerLine(
		// *typing.RefType
		(CGEventSourceRef)source
	);
}
CGRect RectInset(CGRect rect, float dx, float dy) {
	return (CGRect)CGRectInset(
		// *typing.StructType
		(CGRect)rect,
		// *typing.PrimitiveType
		(CGFloat)dx,
		// *typing.PrimitiveType
		(CGFloat)dy
	);
}
bool PDFArrayGetNull(void * array, uint index) {
	return (bool)CGPDFArrayGetNull(
		// *typing.RefType
		(CGPDFArrayRef)array,
		// *typing.PrimitiveType
		(NSUInteger)index
	);
}
const float* ImageGetDecode(void * image) {
	return (const float*)CGImageGetDecode(
		// *typing.RefType
		(CGImageRef)image
	);
}
void * DataProviderGetInfo(void * provider) {
	return (void *)CGDataProviderGetInfo(
		// *typing.RefType
		(CGDataProviderRef)provider
	);
}
void * PDFStreamCopyData(void * stream, CGPDFDataFormat* format) {
	return (void *)CGPDFStreamCopyData(
		// *typing.RefType
		(CGPDFStreamRef)stream,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGPDFDataFormat*)format
	);
}
void * PatternRetain(void * pattern) {
	return (void *)CGPatternRetain(
		// *typing.RefType
		(CGPatternRef)pattern
	);
}
void * BitmapContextGetColorSpace(void * context) {
	return (void *)CGBitmapContextGetColorSpace(
		// *typing.RefType
		(CGContextRef)context
	);
}
bool RectIsEmpty(CGRect rect) {
	return (bool)CGRectIsEmpty(
		// *typing.StructType
		(CGRect)rect
	);
}
float RectGetMinY(CGRect rect) {
	return (float)CGRectGetMinY(
		// *typing.StructType
		(CGRect)rect
	);
}
CFTypeID ColorSpaceGetTypeID() {
	return (CFTypeID)CGColorSpaceGetTypeID(
	);
}
void * ShadingRetain(void * shading) {
	return (void *)CGShadingRetain(
		// *typing.RefType
		(CGShadingRef)shading
	);
}
void * ColorCreateCopy(void * color) {
	return (void *)CGColorCreateCopy(
		// *typing.RefType
		(CGColorRef)color
	);
}
void ContextBeginTransparencyLayerWithRect(void * c, CGRect rect, void * auxInfo) {
	return (void)CGContextBeginTransparencyLayerWithRect(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect,
		// *typing.RefType
		(CFDictionaryRef)auxInfo
	);
}
void PDFPageRelease(void * page) {
	return (void)CGPDFPageRelease(
		// *typing.RefType
		(CGPDFPageRef)page
	);
}
CFTypeID ImageGetTypeID() {
	return (CFTypeID)CGImageGetTypeID(
	);
}
void ContextRotateCTM(void * c, float angle) {
	return (void)CGContextRotateCTM(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)angle
	);
}
void ContextEndTransparencyLayer(void * c) {
	return (void)CGContextEndTransparencyLayer(
		// *typing.RefType
		(CGContextRef)c
	);
}
void EventSourceSetLocalEventsSuppressionInterval(void * source, CFTimeInterval seconds) {
	return (void)CGEventSourceSetLocalEventsSuppressionInterval(
		// *typing.RefType
		(CGEventSourceRef)source,
		// *typing.AliasType
		(CFTimeInterval)seconds
	);
}
CGRect RectIntersection(CGRect r1, CGRect r2) {
	return (CGRect)CGRectIntersection(
		// *typing.StructType
		(CGRect)r1,
		// *typing.StructType
		(CGRect)r2
	);
}
int32_t DisplayModeGetIODisplayModeID(void * mode) {
	return (int32_t)CGDisplayModeGetIODisplayModeID(
		// *typing.RefType
		(CGDisplayModeRef)mode
	);
}
CGDirectDisplayID DisplayPrimaryDisplay(CGDirectDisplayID display) {
	return (CGDirectDisplayID)CGDisplayPrimaryDisplay(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void PDFContextAddDestinationAtPoint(void * context, void * name, CGPoint point) {
	return (void)CGPDFContextAddDestinationAtPoint(
		// *typing.RefType
		(CGContextRef)context,
		// *typing.RefType
		(CFStringRef)name,
		// *typing.StructType
		(CGPoint)point
	);
}
int FontGetDescent(void * font) {
	return (int)CGFontGetDescent(
		// *typing.RefType
		(CGFontRef)font
	);
}
void ContextSetStrokeColor(void * c, const float* components) {
	return (void)CGContextSetStrokeColor(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)components
	);
}
void DisplayModeRelease(void * mode) {
	return (void)CGDisplayModeRelease(
		// *typing.RefType
		(CGDisplayModeRef)mode
	);
}
void ContextAddRects(void * c, const CGRect* rects, uint count) {
	return (void)CGContextAddRects(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)rects,
		// *typing.PrimitiveType
		(NSUInteger)count
	);
}
void * PathCreateMutableCopyByTransformingPath(void * path, const CGAffineTransform* transform) {
	return (void *)CGPathCreateMutableCopyByTransformingPath(
		// *typing.RefType
		(CGPathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)transform
	);
}
CGAffineTransform AffineTransformMakeTranslation(float tx, float ty) {
	return (CGAffineTransform)CGAffineTransformMakeTranslation(
		// *typing.PrimitiveType
		(CGFloat)tx,
		// *typing.PrimitiveType
		(CGFloat)ty
	);
}
void ContextSetCharacterSpacing(void * c, float spacing) {
	return (void)CGContextSetCharacterSpacing(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)spacing
	);
}
uint32_t DisplayModeGetIOFlags(void * mode) {
	return (uint32_t)CGDisplayModeGetIOFlags(
		// *typing.RefType
		(CGDisplayModeRef)mode
	);
}
void PDFContextClose(void * context) {
	return (void)CGPDFContextClose(
		// *typing.RefType
		(CGContextRef)context
	);
}
void * DataProviderRetain(void * provider) {
	return (void *)CGDataProviderRetain(
		// *typing.RefType
		(CGDataProviderRef)provider
	);
}
bool ColorSpaceSupportsOutput(void * space) {
	return (bool)CGColorSpaceSupportsOutput(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void EventSetDoubleValueField(void * event, CGEventField field, double value) {
	return (void)CGEventSetDoubleValueField(
		// *typing.RefType
		(CGEventRef)event,
		// *typing.AliasType
		(CGEventField)field,
		// *typing.PrimitiveType
		(double)value
	);
}
void ContextDrawPDFPage(void * c, void * page) {
	return (void)CGContextDrawPDFPage(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGPDFPageRef)page
	);
}
void * PDFContextCreateWithURL(void * url, const CGRect* mediaBox, void * auxiliaryInfo) {
	return (void *)CGPDFContextCreateWithURL(
		// *typing.RefType
		(CFURLRef)url,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)mediaBox,
		// *typing.RefType
		(CFDictionaryRef)auxiliaryInfo
	);
}
CFTypeID PDFPageGetTypeID() {
	return (CFTypeID)CGPDFPageGetTypeID(
	);
}
void * PDFContentStreamCreateWithPage(void * page) {
	return (void *)CGPDFContentStreamCreateWithPage(
		// *typing.RefType
		(CGPDFPageRef)page
	);
}
void * FontCreateWithDataProvider(void * provider) {
	return (void *)CGFontCreateWithDataProvider(
		// *typing.RefType
		(CGDataProviderRef)provider
	);
}
bool PDFDictionaryGetNumber(void * dict, char* key, CGPDFReal* value) {
	return (bool)CGPDFDictionaryGetNumber(
		// *typing.RefType
		(CGPDFDictionaryRef)dict,
		// *typing.CStringType
		(char*)key,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGPDFReal*)value
	);
}
void ColorSpaceRelease(void * space) {
	return (void)CGColorSpaceRelease(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void * ImageCreateCopy(void * image) {
	return (void *)CGImageCreateCopy(
		// *typing.RefType
		(CGImageRef)image
	);
}
void ContextEndPage(void * c) {
	return (void)CGContextEndPage(
		// *typing.RefType
		(CGContextRef)c
	);
}
void * PDFDocumentGetID(void * document) {
	return (void *)CGPDFDocumentGetID(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
void ContextSetTextMatrix(void * c, CGAffineTransform t) {
	return (void)CGContextSetTextMatrix(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGAffineTransform)t
	);
}
uint FontGetNumberOfGlyphs(void * font) {
	return (uint)CGFontGetNumberOfGlyphs(
		// *typing.RefType
		(CGFontRef)font
	);
}
void ContextClipToMask(void * c, CGRect rect, void * mask) {
	return (void)CGContextClipToMask(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect,
		// *typing.RefType
		(CGImageRef)mask
	);
}
void * PDFPageGetDictionary(void * page) {
	return (void *)CGPDFPageGetDictionary(
		// *typing.RefType
		(CGPDFPageRef)page
	);
}
uint32_t DisplayGammaTableCapacity(CGDirectDisplayID display) {
	return (uint32_t)CGDisplayGammaTableCapacity(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void EventSetIntegerValueField(void * event, CGEventField field, int64_t value) {
	return (void)CGEventSetIntegerValueField(
		// *typing.RefType
		(CGEventRef)event,
		// *typing.AliasType
		(CGEventField)field,
		// *typing.PrimitiveType
		(int64_t)value
	);
}
void ContextAddPath(void * c, void * path) {
	return (void)CGContextAddPath(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGPathRef)path
	);
}
CGAffineTransform ContextGetTextMatrix(void * c) {
	return (CGAffineTransform)CGContextGetTextMatrix(
		// *typing.RefType
		(CGContextRef)c
	);
}
void ContextStrokePath(void * c) {
	return (void)CGContextStrokePath(
		// *typing.RefType
		(CGContextRef)c
	);
}
uint DisplayPixelsWide(CGDirectDisplayID display) {
	return (uint)CGDisplayPixelsWide(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void ContextSetShouldAntialias(void * c, bool shouldAntialias) {
	return (void)CGContextSetShouldAntialias(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(BOOL)shouldAntialias
	);
}
bool EventTapIsEnabled(void * tap) {
	return (bool)CGEventTapIsEnabled(
		// *typing.RefType
		(CFMachPortRef)tap
	);
}
char* PDFTagTypeGetName(CGPDFTagType tagType) {
	return (char*)CGPDFTagTypeGetName(
		// *typing.AliasType
		(CGPDFTagType)tagType
	);
}
void ContextSetGrayFillColor(void * c, float gray, float alpha) {
	return (void)CGContextSetGrayFillColor(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)gray,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
void * PDFContentStreamGetResource(void * cs, char* category, char* name) {
	return (void *)CGPDFContentStreamGetResource(
		// *typing.RefType
		(CGPDFContentStreamRef)cs,
		// *typing.CStringType
		(char*)category,
		// *typing.CStringType
		(char*)name
	);
}
bool PDFDictionaryGetBoolean(void * dict, char* key, CGPDFBoolean* value) {
	return (bool)CGPDFDictionaryGetBoolean(
		// *typing.RefType
		(CGPDFDictionaryRef)dict,
		// *typing.CStringType
		(char*)key,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGPDFBoolean*)value
	);
}
bool PSConverterConvert(void * converter, void * provider, void * consumer, void * options) {
	return (bool)CGPSConverterConvert(
		// *typing.RefType
		(CGPSConverterRef)converter,
		// *typing.RefType
		(CGDataProviderRef)provider,
		// *typing.RefType
		(CGDataConsumerRef)consumer,
		// *typing.RefType
		(CFDictionaryRef)options
	);
}
void PDFContentStreamRelease(void * cs) {
	return (void)CGPDFContentStreamRelease(
		// *typing.RefType
		(CGPDFContentStreamRef)cs
	);
}
CGOpenGLDisplayMask DisplayIDToOpenGLDisplayMask(CGDirectDisplayID display) {
	return (CGOpenGLDisplayMask)CGDisplayIDToOpenGLDisplayMask(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
float RectGetMaxX(CGRect rect) {
	return (float)CGRectGetMaxX(
		// *typing.StructType
		(CGRect)rect
	);
}
bool ContextIsPathEmpty(void * c) {
	return (bool)CGContextIsPathEmpty(
		// *typing.RefType
		(CGContextRef)c
	);
}
CFTypeID ColorConversionInfoGetTypeID() {
	return (CFTypeID)CGColorConversionInfoGetTypeID(
	);
}
uint ImageGetBitsPerPixel(void * image) {
	return (uint)CGImageGetBitsPerPixel(
		// *typing.RefType
		(CGImageRef)image
	);
}
bool RectMakeWithDictionaryRepresentation(void * dict, CGRect* rect) {
	return (bool)CGRectMakeWithDictionaryRepresentation(
		// *typing.RefType
		(CFDictionaryRef)dict,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)rect
	);
}
uint DisplayModeGetPixelHeight(void * mode) {
	return (uint)CGDisplayModeGetPixelHeight(
		// *typing.RefType
		(CGDisplayModeRef)mode
	);
}
void * ColorSpaceCreateDeviceRGB() {
	return (void *)CGColorSpaceCreateDeviceRGB(
	);
}
void ContextSetStrokeColorWithColor(void * c, void * color) {
	return (void)CGContextSetStrokeColorWithColor(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGColorRef)color
	);
}
void * LayerGetContext(void * layer) {
	return (void *)CGLayerGetContext(
		// *typing.RefType
		(CGLayerRef)layer
	);
}
bool PDFArrayGetInteger(void * array, uint index, CGPDFInteger* value) {
	return (bool)CGPDFArrayGetInteger(
		// *typing.RefType
		(CGPDFArrayRef)array,
		// *typing.PrimitiveType
		(NSUInteger)index,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGPDFInteger*)value
	);
}
float FontGetItalicAngle(void * font) {
	return (float)CGFontGetItalicAngle(
		// *typing.RefType
		(CGFontRef)font
	);
}
CGError ConfigureDisplayOrigin(void * config, CGDirectDisplayID display, int32_t x, int32_t y) {
	return (CGError)CGConfigureDisplayOrigin(
		// *typing.RefType
		(CGDisplayConfigRef)config,
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.PrimitiveType
		(int32_t)x,
		// *typing.PrimitiveType
		(int32_t)y
	);
}
void * ColorConversionInfoCreateWithOptions(void * src, void * dst, void * options) {
	return (void *)CGColorConversionInfoCreateWithOptions(
		// *typing.RefType
		(CGColorSpaceRef)src,
		// *typing.RefType
		(CGColorSpaceRef)dst,
		// *typing.RefType
		(CFDictionaryRef)options
	);
}
void RestorePermanentDisplayConfiguration() {
	return (void)CGRestorePermanentDisplayConfiguration(
	);
}
CGImageAlphaInfo BitmapContextGetAlphaInfo(void * context) {
	return (CGImageAlphaInfo)CGBitmapContextGetAlphaInfo(
		// *typing.RefType
		(CGContextRef)context
	);
}
void ContextClosePath(void * c) {
	return (void)CGContextClosePath(
		// *typing.RefType
		(CGContextRef)c
	);
}
bool PathEqualToPath(void * path1, void * path2) {
	return (bool)CGPathEqualToPath(
		// *typing.RefType
		(CGPathRef)path1,
		// *typing.RefType
		(CGPathRef)path2
	);
}
void * EventCreateMouseEvent(void * source, CGEventType mouseType, CGPoint mouseCursorPosition, CGMouseButton mouseButton) {
	return (void *)CGEventCreateMouseEvent(
		// *typing.RefType
		(CGEventSourceRef)source,
		// *typing.AliasType
		(CGEventType)mouseType,
		// *typing.StructType
		(CGPoint)mouseCursorPosition,
		// *typing.AliasType
		(CGMouseButton)mouseButton
	);
}
CGDirectDisplayID MainDisplayID() {
	return (CGDirectDisplayID)CGMainDisplayID(
	);
}
CGAffineTransform AffineTransformTranslate(CGAffineTransform t, float tx, float ty) {
	return (CGAffineTransform)CGAffineTransformTranslate(
		// *typing.StructType
		(CGAffineTransform)t,
		// *typing.PrimitiveType
		(CGFloat)tx,
		// *typing.PrimitiveType
		(CGFloat)ty
	);
}
void * PDFDocumentGetPage(void * document, uint pageNumber) {
	return (void *)CGPDFDocumentGetPage(
		// *typing.RefType
		(CGPDFDocumentRef)document,
		// *typing.PrimitiveType
		(NSUInteger)pageNumber
	);
}
float RectGetMidX(CGRect rect) {
	return (float)CGRectGetMidX(
		// *typing.StructType
		(CGRect)rect
	);
}
CGEventSourceKeyboardType EventSourceGetKeyboardType(void * source) {
	return (CGEventSourceKeyboardType)CGEventSourceGetKeyboardType(
		// *typing.RefType
		(CGEventSourceRef)source
	);
}
void PDFContextBeginPage(void * context, void * pageInfo) {
	return (void)CGPDFContextBeginPage(
		// *typing.RefType
		(CGContextRef)context,
		// *typing.RefType
		(CFDictionaryRef)pageInfo
	);
}
void * ColorCreateCopyWithAlpha(void * color, float alpha) {
	return (void *)CGColorCreateCopyWithAlpha(
		// *typing.RefType
		(CGColorRef)color,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
CGPDFAccessPermissions PDFDocumentGetAccessPermissions(void * document) {
	return (CGPDFAccessPermissions)CGPDFDocumentGetAccessPermissions(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
CGAffineTransform AffineTransformRotate(CGAffineTransform t, float angle) {
	return (CGAffineTransform)CGAffineTransformRotate(
		// *typing.StructType
		(CGAffineTransform)t,
		// *typing.PrimitiveType
		(CGFloat)angle
	);
}
void * EventCreateCopy(void * event) {
	return (void *)CGEventCreateCopy(
		// *typing.RefType
		(CGEventRef)event
	);
}
void ContextSynchronize(void * c) {
	return (void)CGContextSynchronize(
		// *typing.RefType
		(CGContextRef)c
	);
}
CGAffineTransform PDFPageGetDrawingTransform(void * page, CGPDFBox box, CGRect rect, int rotate, bool preserveAspectRatio) {
	return (CGAffineTransform)CGPDFPageGetDrawingTransform(
		// *typing.RefType
		(CGPDFPageRef)page,
		// *typing.AliasType
		(CGPDFBox)box,
		// *typing.StructType
		(CGRect)rect,
		// *typing.PrimitiveType
		(int)rotate,
		// *typing.PrimitiveType
		(BOOL)preserveAspectRatio
	);
}
CGWindowLevel ShieldingWindowLevel() {
	return (CGWindowLevel)CGShieldingWindowLevel(
	);
}
bool DisplayModeIsUsableForDesktopGUI(void * mode) {
	return (bool)CGDisplayModeIsUsableForDesktopGUI(
		// *typing.RefType
		(CGDisplayModeRef)mode
	);
}
CFTypeID GradientGetTypeID() {
	return (CFTypeID)CGGradientGetTypeID(
	);
}
void * PointCreateDictionaryRepresentation(CGPoint point) {
	return (void *)CGPointCreateDictionaryRepresentation(
		// *typing.StructType
		(CGPoint)point
	);
}
CGRect RectOffset(CGRect rect, float dx, float dy) {
	return (CGRect)CGRectOffset(
		// *typing.StructType
		(CGRect)rect,
		// *typing.PrimitiveType
		(CGFloat)dx,
		// *typing.PrimitiveType
		(CGFloat)dy
	);
}
void ContextSetCMYKStrokeColor(void * c, float cyan, float magenta, float yellow, float black, float alpha) {
	return (void)CGContextSetCMYKStrokeColor(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)cyan,
		// *typing.PrimitiveType
		(CGFloat)magenta,
		// *typing.PrimitiveType
		(CGFloat)yellow,
		// *typing.PrimitiveType
		(CGFloat)black,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
void * ColorSpaceGetBaseColorSpace(void * space) {
	return (void *)CGColorSpaceGetBaseColorSpace(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void * LayerCreateWithContext(void * context, CGSize size, void * auxiliaryInfo) {
	return (void *)CGLayerCreateWithContext(
		// *typing.RefType
		(CGContextRef)context,
		// *typing.StructType
		(CGSize)size,
		// *typing.RefType
		(CFDictionaryRef)auxiliaryInfo
	);
}
uint DisplayModeGetPixelWidth(void * mode) {
	return (uint)CGDisplayModeGetPixelWidth(
		// *typing.RefType
		(CGDisplayModeRef)mode
	);
}
void * ImageCreateWithMaskingColors(void * image, const float* components) {
	return (void *)CGImageCreateWithMaskingColors(
		// *typing.RefType
		(CGImageRef)image,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)components
	);
}
void ContextFillRects(void * c, const CGRect* rects, uint count) {
	return (void)CGContextFillRects(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)rects,
		// *typing.PrimitiveType
		(NSUInteger)count
	);
}
void * FontCopyPostScriptName(void * font) {
	return (void *)CGFontCopyPostScriptName(
		// *typing.RefType
		(CGFontRef)font
	);
}
bool ImageIsMask(void * image) {
	return (bool)CGImageIsMask(
		// *typing.RefType
		(CGImageRef)image
	);
}
void ContextSetStrokeColorSpace(void * c, void * space) {
	return (void)CGContextSetStrokeColorSpace(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
CGRect RectStandardize(CGRect rect) {
	return (CGRect)CGRectStandardize(
		// *typing.StructType
		(CGRect)rect
	);
}
CGRect FontGetFontBBox(void * font) {
	return (CGRect)CGFontGetFontBBox(
		// *typing.RefType
		(CGFontRef)font
	);
}
CGError DisplayCapture(CGDirectDisplayID display) {
	return (CGError)CGDisplayCapture(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void ContextAddEllipseInRect(void * c, CGRect rect) {
	return (void)CGContextAddEllipseInRect(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect
	);
}
void * GradientCreateWithColorComponents(void * space, const float* components, const float* locations, uint count) {
	return (void *)CGGradientCreateWithColorComponents(
		// *typing.RefType
		(CGColorSpaceRef)space,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)components,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)locations,
		// *typing.PrimitiveType
		(NSUInteger)count
	);
}
void * ColorSpaceCreateCalibratedRGB(const float* whitePoint, const float* blackPoint, const float* gamma, const float* matrix) {
	return (void *)CGColorSpaceCreateCalibratedRGB(
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)whitePoint,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)blackPoint,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)gamma,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)matrix
	);
}
void ContextDrawShading(void * c, void * shading) {
	return (void)CGContextDrawShading(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGShadingRef)shading
	);
}
float RectGetMidY(CGRect rect) {
	return (float)CGRectGetMidY(
		// *typing.StructType
		(CGRect)rect
	);
}
void * DisplayCreateImageForRect(CGDirectDisplayID display, CGRect rect) {
	return (void *)CGDisplayCreateImageForRect(
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.StructType
		(CGRect)rect
	);
}
void * ColorCreate(void * space, const float* components) {
	return (void *)CGColorCreate(
		// *typing.RefType
		(CGColorSpaceRef)space,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)components
	);
}
void PathMoveToPoint(void * path, const CGAffineTransform* m, float x, float y) {
	return (void)CGPathMoveToPoint(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y
	);
}
void ColorRelease(void * color) {
	return (void)CGColorRelease(
		// *typing.RefType
		(CGColorRef)color
	);
}
uint PDFDocumentGetNumberOfPages(void * document) {
	return (uint)CGPDFDocumentGetNumberOfPages(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
void EventSourceSetKeyboardType(void * source, CGEventSourceKeyboardType keyboardType) {
	return (void)CGEventSourceSetKeyboardType(
		// *typing.RefType
		(CGEventSourceRef)source,
		// *typing.AliasType
		(CGEventSourceKeyboardType)keyboardType
	);
}
void * DisplayCopyAllDisplayModes(CGDirectDisplayID display, void * options) {
	return (void *)CGDisplayCopyAllDisplayModes(
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.RefType
		(CFDictionaryRef)options
	);
}
uint DisplayPixelsHigh(CGDirectDisplayID display) {
	return (uint)CGDisplayPixelsHigh(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
CGError SetDisplayTransferByByteTable(CGDirectDisplayID display, uint32_t tableSize, const uint8_t* redTable, const uint8_t* greenTable, const uint8_t* blueTable) {
	return (CGError)CGSetDisplayTransferByByteTable(
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.PrimitiveType
		(uint32_t)tableSize,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)redTable,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)greenTable,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)blueTable
	);
}
void * ImageGetDataProvider(void * image) {
	return (void *)CGImageGetDataProvider(
		// *typing.RefType
		(CGImageRef)image
	);
}
void ContextSetFillPattern(void * c, void * pattern, const float* components) {
	return (void)CGContextSetFillPattern(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGPatternRef)pattern,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)components
	);
}
CFTypeID DataProviderGetTypeID() {
	return (CFTypeID)CGDataProviderGetTypeID(
	);
}
void * DataProviderCreateWithFilename(char* filename) {
	return (void *)CGDataProviderCreateWithFilename(
		// *typing.CStringType
		(char*)filename
	);
}
uint ImageGetBytesPerRow(void * image) {
	return (uint)CGImageGetBytesPerRow(
		// *typing.RefType
		(CGImageRef)image
	);
}
void * DataProviderCreateWithCFData(void * data) {
	return (void *)CGDataProviderCreateWithCFData(
		// *typing.RefType
		(CFDataRef)data
	);
}
bool PDFDocumentAllowsCopying(void * document) {
	return (bool)CGPDFDocumentAllowsCopying(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
void * ColorCreateCopyByMatchingToColorSpace(void * arg0, CGColorRenderingIntent intent, void * color, void * options) {
	return (void *)CGColorCreateCopyByMatchingToColorSpace(
		// *typing.RefType
		(CGColorSpaceRef)arg0,
		// *typing.AliasType
		(CGColorRenderingIntent)intent,
		// *typing.RefType
		(CGColorRef)color,
		// *typing.RefType
		(CFDictionaryRef)options
	);
}
CGSize SizeMake(float width, float height) {
	return (CGSize)CGSizeMake(
		// *typing.PrimitiveType
		(CGFloat)width,
		// *typing.PrimitiveType
		(CGFloat)height
	);
}
float RectGetMaxY(CGRect rect) {
	return (float)CGRectGetMaxY(
		// *typing.StructType
		(CGRect)rect
	);
}
bool PDFDocumentAllowsPrinting(void * document) {
	return (bool)CGPDFDocumentAllowsPrinting(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
void ContextSetAllowsFontSmoothing(void * c, bool allowsFontSmoothing) {
	return (void)CGContextSetAllowsFontSmoothing(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(BOOL)allowsFontSmoothing
	);
}
void ContextDrawPath(void * c, CGPathDrawingMode mode) {
	return (void)CGContextDrawPath(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.AliasType
		(CGPathDrawingMode)mode
	);
}
CGGlyph FontGetGlyphWithGlyphName(void * font, void * name) {
	return (CGGlyph)CGFontGetGlyphWithGlyphName(
		// *typing.RefType
		(CGFontRef)font,
		// *typing.RefType
		(CFStringRef)name
	);
}
void * RectCreateDictionaryRepresentation(CGRect arg0) {
	return (void *)CGRectCreateDictionaryRepresentation(
		// *typing.StructType
		(CGRect)arg0
	);
}
void * BitmapContextGetData(void * context) {
	return (void *)CGBitmapContextGetData(
		// *typing.RefType
		(CGContextRef)context
	);
}
int FontGetAscent(void * font) {
	return (int)CGFontGetAscent(
		// *typing.RefType
		(CGFontRef)font
	);
}
CGError DisplaySetStereoOperation(CGDirectDisplayID display, boolean_t stereo, boolean_t forceBlueLine, CGConfigureOption option) {
	return (CGError)CGDisplaySetStereoOperation(
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.AliasType
		(boolean_t)stereo,
		// *typing.AliasType
		(boolean_t)forceBlueLine,
		// *typing.AliasType
		(CGConfigureOption)option
	);
}
bool ImageGetShouldInterpolate(void * image) {
	return (bool)CGImageGetShouldInterpolate(
		// *typing.RefType
		(CGImageRef)image
	);
}
void ContextFillPath(void * c) {
	return (void)CGContextFillPath(
		// *typing.RefType
		(CGContextRef)c
	);
}
CFTypeID PDFDocumentGetTypeID() {
	return (CFTypeID)CGPDFDocumentGetTypeID(
	);
}
CGColorRenderingIntent ImageGetRenderingIntent(void * image) {
	return (CGColorRenderingIntent)CGImageGetRenderingIntent(
		// *typing.RefType
		(CGImageRef)image
	);
}
void EventPost(CGEventTapLocation tap, void * event) {
	return (void)CGEventPost(
		// *typing.AliasType
		(CGEventTapLocation)tap,
		// *typing.RefType
		(CGEventRef)event
	);
}
float FontGetStemV(void * font) {
	return (float)CGFontGetStemV(
		// *typing.RefType
		(CGFontRef)font
	);
}
void ContextStrokeLineSegments(void * c, const CGPoint* points, uint count) {
	return (void)CGContextStrokeLineSegments(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PointerType
		// -> *typing.StructType
		(CGPoint*)points,
		// *typing.PrimitiveType
		(NSUInteger)count
	);
}
void ContextEOFillPath(void * c) {
	return (void)CGContextEOFillPath(
		// *typing.RefType
		(CGContextRef)c
	);
}
CGError DisplayCaptureWithOptions(CGDirectDisplayID display, CGCaptureOptions options) {
	return (CGError)CGDisplayCaptureWithOptions(
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.AliasType
		(CGCaptureOptions)options
	);
}
bool ContextPathContainsPoint(void * c, CGPoint point, CGPathDrawingMode mode) {
	return (bool)CGContextPathContainsPoint(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGPoint)point,
		// *typing.AliasType
		(CGPathDrawingMode)mode
	);
}
bool ColorSpaceUsesITUR_2100TF(void * arg0) {
	return (bool)CGColorSpaceUsesITUR_2100TF(
		// *typing.RefType
		(CGColorSpaceRef)arg0
	);
}
void * PDFDocumentGetOutline(void * document) {
	return (void *)CGPDFDocumentGetOutline(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
void PathRelease(void * path) {
	return (void)CGPathRelease(
		// *typing.RefType
		(CGPathRef)path
	);
}
CFTypeID PathGetTypeID() {
	return (CFTypeID)CGPathGetTypeID(
	);
}
void PDFDocumentRelease(void * document) {
	return (void)CGPDFDocumentRelease(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
void * WindowServerCreateServerPort() {
	return (void *)CGWindowServerCreateServerPort(
	);
}
CGAffineTransform ContextGetUserSpaceToDeviceSpaceTransform(void * c) {
	return (CGAffineTransform)CGContextGetUserSpaceToDeviceSpaceTransform(
		// *typing.RefType
		(CGContextRef)c
	);
}
void ContextDrawImage(void * c, CGRect rect, void * image) {
	return (void)CGContextDrawImage(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect,
		// *typing.RefType
		(CGImageRef)image
	);
}
void * PathCreateWithRoundedRect(CGRect rect, float cornerWidth, float cornerHeight, const CGAffineTransform* transform) {
	return (void *)CGPathCreateWithRoundedRect(
		// *typing.StructType
		(CGRect)rect,
		// *typing.PrimitiveType
		(CGFloat)cornerWidth,
		// *typing.PrimitiveType
		(CGFloat)cornerHeight,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)transform
	);
}
void PDFContextBeginTag(void * context, CGPDFTagType tagType, void * tagProperties) {
	return (void)CGPDFContextBeginTag(
		// *typing.RefType
		(CGContextRef)context,
		// *typing.AliasType
		(CGPDFTagType)tagType,
		// *typing.RefType
		(CFDictionaryRef)tagProperties
	);
}
CGPDFObjectType PDFObjectGetType(void * object) {
	return (CGPDFObjectType)CGPDFObjectGetType(
		// *typing.RefType
		(CGPDFObjectRef)object
	);
}
void * ImageMaskCreate(uint width, uint height, uint bitsPerComponent, uint bitsPerPixel, uint bytesPerRow, void * provider, const float* decode, bool shouldInterpolate) {
	return (void *)CGImageMaskCreate(
		// *typing.PrimitiveType
		(NSUInteger)width,
		// *typing.PrimitiveType
		(NSUInteger)height,
		// *typing.PrimitiveType
		(NSUInteger)bitsPerComponent,
		// *typing.PrimitiveType
		(NSUInteger)bitsPerPixel,
		// *typing.PrimitiveType
		(NSUInteger)bytesPerRow,
		// *typing.RefType
		(CGDataProviderRef)provider,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)decode,
		// *typing.PrimitiveType
		(BOOL)shouldInterpolate
	);
}
void * ColorCreateGenericRGB(float red, float green, float blue, float alpha) {
	return (void *)CGColorCreateGenericRGB(
		// *typing.PrimitiveType
		(CGFloat)red,
		// *typing.PrimitiveType
		(CGFloat)green,
		// *typing.PrimitiveType
		(CGFloat)blue,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
double EventGetDoubleValueField(void * event, CGEventField field) {
	return (double)CGEventGetDoubleValueField(
		// *typing.RefType
		(CGEventRef)event,
		// *typing.AliasType
		(CGEventField)field
	);
}
void * PDFDocumentGetInfo(void * document) {
	return (void *)CGPDFDocumentGetInfo(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
bool PSConverterIsConverting(void * converter) {
	return (bool)CGPSConverterIsConverting(
		// *typing.RefType
		(CGPSConverterRef)converter
	);
}
void * FontCreateCopyWithVariations(void * font, void * variations) {
	return (void *)CGFontCreateCopyWithVariations(
		// *typing.RefType
		(CGFontRef)font,
		// *typing.RefType
		(CFDictionaryRef)variations
	);
}
CGError ConfigureDisplayMirrorOfDisplay(void * config, CGDirectDisplayID display, CGDirectDisplayID master) {
	return (CGError)CGConfigureDisplayMirrorOfDisplay(
		// *typing.RefType
		(CGDisplayConfigRef)config,
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.AliasType
		(CGDirectDisplayID)master
	);
}
double DisplayModeGetRefreshRate(void * mode) {
	return (double)CGDisplayModeGetRefreshRate(
		// *typing.RefType
		(CGDisplayModeRef)mode
	);
}
CFTimeInterval EventSourceGetLocalEventsSuppressionInterval(void * source) {
	return (CFTimeInterval)CGEventSourceGetLocalEventsSuppressionInterval(
		// *typing.RefType
		(CGEventSourceRef)source
	);
}
void ContextAddArcToPoint(void * c, float x1, float y1, float x2, float y2, float radius) {
	return (void)CGContextAddArcToPoint(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)x1,
		// *typing.PrimitiveType
		(CGFloat)y1,
		// *typing.PrimitiveType
		(CGFloat)x2,
		// *typing.PrimitiveType
		(CGFloat)y2,
		// *typing.PrimitiveType
		(CGFloat)radius
	);
}
void * DataConsumerCreateWithURL(void * url) {
	return (void *)CGDataConsumerCreateWithURL(
		// *typing.RefType
		(CFURLRef)url
	);
}
CGSize LayerGetSize(void * layer) {
	return (CGSize)CGLayerGetSize(
		// *typing.RefType
		(CGLayerRef)layer
	);
}
void * ContextRetain(void * c) {
	return (void *)CGContextRetain(
		// *typing.RefType
		(CGContextRef)c
	);
}
CGError AcquireDisplayFadeReservation(CGDisplayReservationInterval seconds, CGDisplayFadeReservationToken* token) {
	return (CGError)CGAcquireDisplayFadeReservation(
		// *typing.AliasType
		(CGDisplayReservationInterval)seconds,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGDisplayFadeReservationToken*)token
	);
}
CGEventFlags EventGetFlags(void * event) {
	return (CGEventFlags)CGEventGetFlags(
		// *typing.RefType
		(CGEventRef)event
	);
}
void * FontCopyGlyphNameForGlyph(void * font, CGGlyph glyph) {
	return (void *)CGFontCopyGlyphNameForGlyph(
		// *typing.RefType
		(CGFontRef)font,
		// *typing.AliasType
		(CGGlyph)glyph
	);
}
void ContextSetRGBFillColor(void * c, float red, float green, float blue, float alpha) {
	return (void)CGContextSetRGBFillColor(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)red,
		// *typing.PrimitiveType
		(CGFloat)green,
		// *typing.PrimitiveType
		(CGFloat)blue,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
void ContextRestoreGState(void * c) {
	return (void)CGContextRestoreGState(
		// *typing.RefType
		(CGContextRef)c
	);
}
void * PathCreateCopyByTransformingPath(void * path, const CGAffineTransform* transform) {
	return (void *)CGPathCreateCopyByTransformingPath(
		// *typing.RefType
		(CGPathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)transform
	);
}
void * FontCopyFullName(void * font) {
	return (void *)CGFontCopyFullName(
		// *typing.RefType
		(CGFontRef)font
	);
}
CGAffineTransform AffineTransformConcat(CGAffineTransform t1, CGAffineTransform t2) {
	return (CGAffineTransform)CGAffineTransformConcat(
		// *typing.StructType
		(CGAffineTransform)t1,
		// *typing.StructType
		(CGAffineTransform)t2
	);
}
uint PDFPageGetPageNumber(void * page) {
	return (uint)CGPDFPageGetPageNumber(
		// *typing.RefType
		(CGPDFPageRef)page
	);
}
void EventSetLocation(void * event, CGPoint location) {
	return (void)CGEventSetLocation(
		// *typing.RefType
		(CGEventRef)event,
		// *typing.StructType
		(CGPoint)location
	);
}
void FontRelease(void * font) {
	return (void)CGFontRelease(
		// *typing.RefType
		(CGFontRef)font
	);
}
void * PDFStringCopyTextString(void * string_) {
	return (void *)CGPDFStringCopyTextString(
		// *typing.RefType
		(CGPDFStringRef)string_
	);
}
bool PDFDocumentIsEncrypted(void * document) {
	return (bool)CGPDFDocumentIsEncrypted(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
void ContextStrokeRect(void * c, CGRect rect) {
	return (void)CGContextStrokeRect(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect
	);
}
void ContextSetShouldSubpixelPositionFonts(void * c, bool shouldSubpixelPositionFonts) {
	return (void)CGContextSetShouldSubpixelPositionFonts(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(BOOL)shouldSubpixelPositionFonts
	);
}
void * ContextCopyPath(void * c) {
	return (void *)CGContextCopyPath(
		// *typing.RefType
		(CGContextRef)c
	);
}
void * ImageGetUTType(void * image) {
	return (void *)CGImageGetUTType(
		// *typing.RefType
		(CGImageRef)image
	);
}
bool FontCanCreatePostScriptSubset(void * font, CGFontPostScriptFormat format) {
	return (bool)CGFontCanCreatePostScriptSubset(
		// *typing.RefType
		(CGFontRef)font,
		// *typing.AliasType
		(CGFontPostScriptFormat)format
	);
}
const float* ColorGetComponents(void * color) {
	return (const float*)CGColorGetComponents(
		// *typing.RefType
		(CGColorRef)color
	);
}
CFTypeID ContextGetTypeID() {
	return (CFTypeID)CGContextGetTypeID(
	);
}
void * DataConsumerCreateWithCFData(void * data) {
	return (void *)CGDataConsumerCreateWithCFData(
		// *typing.RefType
		(CFMutableDataRef)data
	);
}
void ContextAddRect(void * c, CGRect rect) {
	return (void)CGContextAddRect(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect
	);
}
void PathAddRelativeArc(void * path, const CGAffineTransform* matrix, float x, float y, float radius, float startAngle, float delta) {
	return (void)CGPathAddRelativeArc(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)matrix,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y,
		// *typing.PrimitiveType
		(CGFloat)radius,
		// *typing.PrimitiveType
		(CGFloat)startAngle,
		// *typing.PrimitiveType
		(CGFloat)delta
	);
}
void DataConsumerRelease(void * consumer) {
	return (void)CGDataConsumerRelease(
		// *typing.RefType
		(CGDataConsumerRef)consumer
	);
}
CGError ConfigureDisplayStereoOperation(void * config, CGDirectDisplayID display, boolean_t stereo, boolean_t forceBlueLine) {
	return (CGError)CGConfigureDisplayStereoOperation(
		// *typing.RefType
		(CGDisplayConfigRef)config,
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.AliasType
		(boolean_t)stereo,
		// *typing.AliasType
		(boolean_t)forceBlueLine
	);
}
CGSize ContextConvertSizeToDeviceSpace(void * c, CGSize size) {
	return (CGSize)CGContextConvertSizeToDeviceSpace(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGSize)size
	);
}
void PathAddCurveToPoint(void * path, const CGAffineTransform* m, float cp1x, float cp1y, float cp2x, float cp2y, float x, float y) {
	return (void)CGPathAddCurveToPoint(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.PrimitiveType
		(CGFloat)cp1x,
		// *typing.PrimitiveType
		(CGFloat)cp1y,
		// *typing.PrimitiveType
		(CGFloat)cp2x,
		// *typing.PrimitiveType
		(CGFloat)cp2y,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y
	);
}
boolean_t DisplayIsBuiltin(CGDirectDisplayID display) {
	return (boolean_t)CGDisplayIsBuiltin(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
CGError DisplayHideCursor(CGDirectDisplayID display) {
	return (CGError)CGDisplayHideCursor(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
uint32_t DisplayUnitNumber(CGDirectDisplayID display) {
	return (uint32_t)CGDisplayUnitNumber(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
CGError DisplaySetDisplayMode(CGDirectDisplayID display, void * mode, void * options) {
	return (CGError)CGDisplaySetDisplayMode(
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.RefType
		(CGDisplayModeRef)mode,
		// *typing.RefType
		(CFDictionaryRef)options
	);
}
void ContextShowGlyphsAtPositions(void * c, const CGGlyph* glyphs, const CGPoint* Lpositions, uint count) {
	return (void)CGContextShowGlyphsAtPositions(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGlyph*)glyphs,
		// *typing.PointerType
		// -> *typing.StructType
		(CGPoint*)Lpositions,
		// *typing.PrimitiveType
		(NSUInteger)count
	);
}
void ContextSetFillColorWithColor(void * c, void * color) {
	return (void)CGContextSetFillColorWithColor(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.RefType
		(CGColorRef)color
	);
}
int FontGetLeading(void * font) {
	return (int)CGFontGetLeading(
		// *typing.RefType
		(CGFontRef)font
	);
}
CFTypeID PSConverterGetTypeID() {
	return (CFTypeID)CGPSConverterGetTypeID(
	);
}
void * GradientCreateWithColors(void * space, void * colors, const float* locations) {
	return (void *)CGGradientCreateWithColors(
		// *typing.RefType
		(CGColorSpaceRef)space,
		// *typing.RefType
		(CFArrayRef)colors,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)locations
	);
}
boolean_t DisplayIsAlwaysInMirrorSet(CGDirectDisplayID display) {
	return (boolean_t)CGDisplayIsAlwaysInMirrorSet(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void * PathCreateCopyByDashingPath(void * path, const CGAffineTransform* transform, float phase, const float* lengths, uint count) {
	return (void *)CGPathCreateCopyByDashingPath(
		// *typing.RefType
		(CGPathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)transform,
		// *typing.PrimitiveType
		(CGFloat)phase,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)lengths,
		// *typing.PrimitiveType
		(NSUInteger)count
	);
}
void ContextSetCMYKFillColor(void * c, float cyan, float magenta, float yellow, float black, float alpha) {
	return (void)CGContextSetCMYKFillColor(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)cyan,
		// *typing.PrimitiveType
		(CGFloat)magenta,
		// *typing.PrimitiveType
		(CGFloat)yellow,
		// *typing.PrimitiveType
		(CGFloat)black,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
CGError CancelDisplayConfiguration(void * config) {
	return (CGError)CGCancelDisplayConfiguration(
		// *typing.RefType
		(CGDisplayConfigRef)config
	);
}
void PathAddLines(void * path, const CGAffineTransform* m, const CGPoint* points, uint count) {
	return (void)CGPathAddLines(
		// *typing.RefType
		(CGMutablePathRef)path,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)m,
		// *typing.PointerType
		// -> *typing.StructType
		(CGPoint*)points,
		// *typing.PrimitiveType
		(NSUInteger)count
	);
}
CGError SetDisplayTransferByFormula(CGDirectDisplayID display, CGGammaValue redMin, CGGammaValue redMax, CGGammaValue redGamma, CGGammaValue greenMin, CGGammaValue greenMax, CGGammaValue greenGamma, CGGammaValue blueMin, CGGammaValue blueMax, CGGammaValue blueGamma) {
	return (CGError)CGSetDisplayTransferByFormula(
		// *typing.AliasType
		(CGDirectDisplayID)display,
		// *typing.AliasType
		(CGGammaValue)redMin,
		// *typing.AliasType
		(CGGammaValue)redMax,
		// *typing.AliasType
		(CGGammaValue)redGamma,
		// *typing.AliasType
		(CGGammaValue)greenMin,
		// *typing.AliasType
		(CGGammaValue)greenMax,
		// *typing.AliasType
		(CGGammaValue)greenGamma,
		// *typing.AliasType
		(CGGammaValue)blueMin,
		// *typing.AliasType
		(CGGammaValue)blueMax,
		// *typing.AliasType
		(CGGammaValue)blueGamma
	);
}
boolean_t DisplayIsAsleep(CGDirectDisplayID display) {
	return (boolean_t)CGDisplayIsAsleep(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
CGBitmapInfo ImageGetBitmapInfo(void * image) {
	return (CGBitmapInfo)CGImageGetBitmapInfo(
		// *typing.RefType
		(CGImageRef)image
	);
}
uint BitmapContextGetWidth(void * context) {
	return (uint)CGBitmapContextGetWidth(
		// *typing.RefType
		(CGContextRef)context
	);
}
CGAffineTransform AffineTransformMake(float a, float b, float c, float d, float tx, float ty) {
	return (CGAffineTransform)CGAffineTransformMake(
		// *typing.PrimitiveType
		(CGFloat)a,
		// *typing.PrimitiveType
		(CGFloat)b,
		// *typing.PrimitiveType
		(CGFloat)c,
		// *typing.PrimitiveType
		(CGFloat)d,
		// *typing.PrimitiveType
		(CGFloat)tx,
		// *typing.PrimitiveType
		(CGFloat)ty
	);
}
void * PathCreateWithEllipseInRect(CGRect rect, const CGAffineTransform* transform) {
	return (void *)CGPathCreateWithEllipseInRect(
		// *typing.StructType
		(CGRect)rect,
		// *typing.PointerType
		// -> *typing.StructType
		(CGAffineTransform*)transform
	);
}
bool RectIntersectsRect(CGRect rect1, CGRect rect2) {
	return (bool)CGRectIntersectsRect(
		// *typing.StructType
		(CGRect)rect1,
		// *typing.StructType
		(CGRect)rect2
	);
}
void * DisplayCreateImage(CGDirectDisplayID displayID) {
	return (void *)CGDisplayCreateImage(
		// *typing.AliasType
		(CGDirectDisplayID)displayID
	);
}
void ContextSetLineWidth(void * c, float width) {
	return (void)CGContextSetLineWidth(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)width
	);
}
void PatternRelease(void * pattern) {
	return (void)CGPatternRelease(
		// *typing.RefType
		(CGPatternRef)pattern
	);
}
void ContextEOClip(void * c) {
	return (void)CGContextEOClip(
		// *typing.RefType
		(CGContextRef)c
	);
}
void * ColorSpaceCreateExtended(void * space) {
	return (void *)CGColorSpaceCreateExtended(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void ContextSetAllowsFontSubpixelQuantization(void * c, bool allowsFontSubpixelQuantization) {
	return (void)CGContextSetAllowsFontSubpixelQuantization(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(BOOL)allowsFontSubpixelQuantization
	);
}
CGPoint ContextConvertPointToUserSpace(void * c, CGPoint point) {
	return (CGPoint)CGContextConvertPointToUserSpace(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGPoint)point
	);
}
void * DataConsumerRetain(void * consumer) {
	return (void *)CGDataConsumerRetain(
		// *typing.RefType
		(CGDataConsumerRef)consumer
	);
}
void EventSourceSetLocalEventsFilterDuringSuppressionState(void * source, CGEventFilterMask filter, CGEventSuppressionState state) {
	return (void)CGEventSourceSetLocalEventsFilterDuringSuppressionState(
		// *typing.RefType
		(CGEventSourceRef)source,
		// *typing.AliasType
		(CGEventFilterMask)filter,
		// *typing.AliasType
		(CGEventSuppressionState)state
	);
}
void ContextSetGrayStrokeColor(void * c, float gray, float alpha) {
	return (void)CGContextSetGrayStrokeColor(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)gray,
		// *typing.PrimitiveType
		(CGFloat)alpha
	);
}
uint32_t DisplayVendorNumber(CGDirectDisplayID display) {
	return (uint32_t)CGDisplayVendorNumber(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
bool RectContainsPoint(CGRect rect, CGPoint point) {
	return (bool)CGRectContainsPoint(
		// *typing.StructType
		(CGRect)rect,
		// *typing.StructType
		(CGPoint)point
	);
}
CFTypeID FontGetTypeID() {
	return (CFTypeID)CGFontGetTypeID(
	);
}
void * DisplayGetDrawingContext(CGDirectDisplayID display) {
	return (void *)CGDisplayGetDrawingContext(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void ContextSetShouldSmoothFonts(void * c, bool shouldSmoothFonts) {
	return (void)CGContextSetShouldSmoothFonts(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(BOOL)shouldSmoothFonts
	);
}
void * EventCreate(void * source) {
	return (void *)CGEventCreate(
		// *typing.RefType
		(CGEventSourceRef)source
	);
}
CGDirectDisplayID DisplayMirrorsDisplay(CGDirectDisplayID display) {
	return (CGDirectDisplayID)CGDisplayMirrorsDisplay(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
bool AffineTransformIsIdentity(CGAffineTransform t) {
	return (bool)CGAffineTransformIsIdentity(
		// *typing.StructType
		(CGAffineTransform)t
	);
}
void ContextBeginPage(void * c, const CGRect* mediaBox) {
	return (void)CGContextBeginPage(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PointerType
		// -> *typing.StructType
		(CGRect*)mediaBox
	);
}
float RectGetWidth(CGRect rect) {
	return (float)CGRectGetWidth(
		// *typing.StructType
		(CGRect)rect
	);
}
CGPoint PointMake(float x, float y) {
	return (CGPoint)CGPointMake(
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y
	);
}
void * EventCreateScrollWheelEvent(void * source, CGScrollEventUnit units, uint32_t wheelCount, int32_t wheel1) {
	return (void *)CGEventCreateScrollWheelEvent(
		// *typing.RefType
		(CGEventSourceRef)source,
		// *typing.AliasType
		(CGScrollEventUnit)units,
		// *typing.PrimitiveType
		(uint32_t)wheelCount,
		// *typing.PrimitiveType
		(int32_t)wheel1
	);
}
void * ColorGetConstantColor(void * colorName) {
	return (void *)CGColorGetConstantColor(
		// *typing.RefType
		(CFStringRef)colorName
	);
}
void * EventCreateSourceFromEvent(void * event) {
	return (void *)CGEventCreateSourceFromEvent(
		// *typing.RefType
		(CGEventRef)event
	);
}
int FontGetXHeight(void * font) {
	return (int)CGFontGetXHeight(
		// *typing.RefType
		(CGFontRef)font
	);
}
CGPoint EventGetUnflippedLocation(void * event) {
	return (CGPoint)CGEventGetUnflippedLocation(
		// *typing.RefType
		(CGEventRef)event
	);
}
void * ImageCreateWithMask(void * image, void * mask) {
	return (void *)CGImageCreateWithMask(
		// *typing.RefType
		(CGImageRef)image,
		// *typing.RefType
		(CGImageRef)mask
	);
}
void PDFContextSetURLForRect(void * context, void * url, CGRect rect) {
	return (void)CGPDFContextSetURLForRect(
		// *typing.RefType
		(CGContextRef)context,
		// *typing.RefType
		(CFURLRef)url,
		// *typing.StructType
		(CGRect)rect
	);
}
void * ColorSpaceCreateLinearized(void * space) {
	return (void *)CGColorSpaceCreateLinearized(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
uint ColorSpaceGetColorTableCount(void * space) {
	return (uint)CGColorSpaceGetColorTableCount(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
CGDirectDisplayID OpenGLDisplayMaskToDisplayID(CGOpenGLDisplayMask mask) {
	return (CGDirectDisplayID)CGOpenGLDisplayMaskToDisplayID(
		// *typing.AliasType
		(CGOpenGLDisplayMask)mask
	);
}
void ContextFlush(void * c) {
	return (void)CGContextFlush(
		// *typing.RefType
		(CGContextRef)c
	);
}
CGVector VectorMake(float dx, float dy) {
	return (CGVector)CGVectorMake(
		// *typing.PrimitiveType
		(CGFloat)dx,
		// *typing.PrimitiveType
		(CGFloat)dy
	);
}
uint DisplayModeGetHeight(void * mode) {
	return (uint)CGDisplayModeGetHeight(
		// *typing.RefType
		(CGDisplayModeRef)mode
	);
}
CGRect RectUnion(CGRect r1, CGRect r2) {
	return (CGRect)CGRectUnion(
		// *typing.StructType
		(CGRect)r1,
		// *typing.StructType
		(CGRect)r2
	);
}
bool PDFScannerScan(void * scanner) {
	return (bool)CGPDFScannerScan(
		// *typing.RefType
		(CGPDFScannerRef)scanner
	);
}
void ContextAddLines(void * c, const CGPoint* points, uint count) {
	return (void)CGContextAddLines(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PointerType
		// -> *typing.StructType
		(CGPoint*)points,
		// *typing.PrimitiveType
		(NSUInteger)count
	);
}
void ContextSetMiterLimit(void * c, float limit) {
	return (void)CGContextSetMiterLimit(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)limit
	);
}
void * PDFStreamGetDictionary(void * stream) {
	return (void *)CGPDFStreamGetDictionary(
		// *typing.RefType
		(CGPDFStreamRef)stream
	);
}
uint PDFStringGetLength(void * string_) {
	return (uint)CGPDFStringGetLength(
		// *typing.RefType
		(CGPDFStringRef)string_
	);
}
void ContextClip(void * c) {
	return (void)CGContextClip(
		// *typing.RefType
		(CGContextRef)c
	);
}
CGError ReleaseAllDisplays() {
	return (CGError)CGReleaseAllDisplays(
	);
}
uint32_t DisplaySerialNumber(CGDirectDisplayID display) {
	return (uint32_t)CGDisplaySerialNumber(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void * ColorSpaceCopyName(void * space) {
	return (void *)CGColorSpaceCopyName(
		// *typing.RefType
		(CGColorSpaceRef)space
	);
}
void ContextSetLineCap(void * c, CGLineCap cap) {
	return (void)CGContextSetLineCap(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.AliasType
		(CGLineCap)cap
	);
}
void ShadingRelease(void * shading) {
	return (void)CGShadingRelease(
		// *typing.RefType
		(CGShadingRef)shading
	);
}
void * PDFOperatorTableRetain(void * table) {
	return (void *)CGPDFOperatorTableRetain(
		// *typing.RefType
		(CGPDFOperatorTableRef)table
	);
}
uint32_t DisplayModelNumber(CGDirectDisplayID display) {
	return (uint32_t)CGDisplayModelNumber(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void EventSourceSetUserData(void * source, int64_t userData) {
	return (void)CGEventSourceSetUserData(
		// *typing.RefType
		(CGEventSourceRef)source,
		// *typing.PrimitiveType
		(int64_t)userData
	);
}
void GradientRelease(void * gradient) {
	return (void)CGGradientRelease(
		// *typing.RefType
		(CGGradientRef)gradient
	);
}
bool RequestListenEventAccess() {
	return (bool)CGRequestListenEventAccess(
	);
}
void ContextMoveToPoint(void * c, float x, float y) {
	return (void)CGContextMoveToPoint(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y
	);
}
CGAffineTransform AffineTransformInvert(CGAffineTransform t) {
	return (CGAffineTransform)CGAffineTransformInvert(
		// *typing.StructType
		(CGAffineTransform)t
	);
}
void * PDFScannerRetain(void * scanner) {
	return (void *)CGPDFScannerRetain(
		// *typing.RefType
		(CGPDFScannerRef)scanner
	);
}
bool PathIsEmpty(void * path) {
	return (bool)CGPathIsEmpty(
		// *typing.RefType
		(CGPathRef)path
	);
}
void ContextDrawTiledImage(void * c, CGRect rect, void * image) {
	return (void)CGContextDrawTiledImage(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect,
		// *typing.RefType
		(CGImageRef)image
	);
}
void * ImageRetain(void * image) {
	return (void *)CGImageRetain(
		// *typing.RefType
		(CGImageRef)image
	);
}
void * ImageCreateWithImageInRect(void * image, CGRect rect) {
	return (void *)CGImageCreateWithImageInRect(
		// *typing.RefType
		(CGImageRef)image,
		// *typing.StructType
		(CGRect)rect
	);
}
int PDFPageGetRotationAngle(void * page) {
	return (int)CGPDFPageGetRotationAngle(
		// *typing.RefType
		(CGPDFPageRef)page
	);
}
void ContextAddLineToPoint(void * c, float x, float y) {
	return (void)CGContextAddLineToPoint(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)x,
		// *typing.PrimitiveType
		(CGFloat)y
	);
}
bool PDFScannerPopBoolean(void * scanner, CGPDFBoolean* value) {
	return (bool)CGPDFScannerPopBoolean(
		// *typing.RefType
		(CGPDFScannerRef)scanner,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGPDFBoolean*)value
	);
}
void * FontCreateWithFontName(void * name) {
	return (void *)CGFontCreateWithFontName(
		// *typing.RefType
		(CFStringRef)name
	);
}
void ContextSetInterpolationQuality(void * c, CGInterpolationQuality quality) {
	return (void)CGContextSetInterpolationQuality(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.AliasType
		(CGInterpolationQuality)quality
	);
}
void * PDFDocumentGetCatalog(void * document) {
	return (void *)CGPDFDocumentGetCatalog(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
void LayerRelease(void * layer) {
	return (void)CGLayerRelease(
		// *typing.RefType
		(CGLayerRef)layer
	);
}
CFTypeID LayerGetTypeID() {
	return (CFTypeID)CGLayerGetTypeID(
	);
}
CFTypeID FunctionGetTypeID() {
	return (CFTypeID)CGFunctionGetTypeID(
	);
}
void * PDFContentStreamGetStreams(void * cs) {
	return (void *)CGPDFContentStreamGetStreams(
		// *typing.RefType
		(CGPDFContentStreamRef)cs
	);
}
uint PDFDictionaryGetCount(void * dict) {
	return (uint)CGPDFDictionaryGetCount(
		// *typing.RefType
		(CGPDFDictionaryRef)dict
	);
}
CGError DisplayShowCursor(CGDirectDisplayID display) {
	return (CGError)CGDisplayShowCursor(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
CGBitmapInfo BitmapContextGetBitmapInfo(void * context) {
	return (CGBitmapInfo)CGBitmapContextGetBitmapInfo(
		// *typing.RefType
		(CGContextRef)context
	);
}
CGAffineTransform AffineTransformScale(CGAffineTransform t, float sx, float sy) {
	return (CGAffineTransform)CGAffineTransformScale(
		// *typing.StructType
		(CGAffineTransform)t,
		// *typing.PrimitiveType
		(CGFloat)sx,
		// *typing.PrimitiveType
		(CGFloat)sy
	);
}
void * FontCreatePostScriptSubset(void * font, void * subsetName, CGFontPostScriptFormat format, const CGGlyph* glyphs, uint count, const CGGlyph* encoding) {
	return (void *)CGFontCreatePostScriptSubset(
		// *typing.RefType
		(CGFontRef)font,
		// *typing.RefType
		(CFStringRef)subsetName,
		// *typing.AliasType
		(CGFontPostScriptFormat)format,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGlyph*)glyphs,
		// *typing.PrimitiveType
		(NSUInteger)count,
		// *typing.PointerType
		// -> *typing.AliasType
		(CGGlyph*)encoding
	);
}
boolean_t DisplayIsMain(CGDirectDisplayID display) {
	return (boolean_t)CGDisplayIsMain(
		// *typing.AliasType
		(CGDirectDisplayID)display
	);
}
void * PathCreateMutable() {
	return (void *)CGPathCreateMutable(
	);
}
void DisplayRestoreColorSyncSettings() {
	return (void)CGDisplayRestoreColorSyncSettings(
	);
}
void ContextSetLineDash(void * c, float phase, const float* lengths, uint count) {
	return (void)CGContextSetLineDash(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.PrimitiveType
		(CGFloat)phase,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(CGFloat*)lengths,
		// *typing.PrimitiveType
		(NSUInteger)count
	);
}
CFTypeID DisplayModeGetTypeID() {
	return (CFTypeID)CGDisplayModeGetTypeID(
	);
}
CGEventFlags EventSourceFlagsState(CGEventSourceStateID stateID) {
	return (CGEventFlags)CGEventSourceFlagsState(
		// *typing.AliasType
		(CGEventSourceStateID)stateID
	);
}
bool RequestPostEventAccess() {
	return (bool)CGRequestPostEventAccess(
	);
}
void ContextFillRect(void * c, CGRect rect) {
	return (void)CGContextFillRect(
		// *typing.RefType
		(CGContextRef)c,
		// *typing.StructType
		(CGRect)rect
	);
}
void * PathRetain(void * path) {
	return (void *)CGPathRetain(
		// *typing.RefType
		(CGPathRef)path
	);
}
bool PDFDocumentIsUnlocked(void * document) {
	return (bool)CGPDFDocumentIsUnlocked(
		// *typing.RefType
		(CGPDFDocumentRef)document
	);
}
