// AUTO-GENERATED CODE, DO NOT MODIFY

#import "CoreGraphics/CoreGraphics.h"
int32_t PostKeyboardEvent(uint16_t keyChar, uint16_t virtualKey, NSInteger keyDown) {
	return CGPostKeyboardEvent(keyChar, virtualKey, keyDown);
}
void ContextStrokeEllipseInRect(void * c, CGRect rect) {
	return CGContextStrokeEllipseInRect(c, rect);
}
CGRect ContextConvertRectToDeviceSpace(void * c, CGRect rect) {
	return CGContextConvertRectToDeviceSpace(c, rect);
}
void EventSetSource(void * event, void * source) {
	return CGEventSetSource(event, source);
}
void ContextSetFontSize(void * c, float size) {
	return CGContextSetFontSize(c, size);
}
void ContextClipToRects(void * c, CGRect* rects, NSUInteger count) {
	return CGContextClipToRects(c, rects, count);
}
CGRect RectApplyAffineTransform(CGRect rect, CGAffineTransform t) {
	return CGRectApplyAffineTransform(rect, t);
}
int64_t EventSourceGetUserData(void * source) {
	return CGEventSourceGetUserData(source);
}
uint32_t ImageGetAlphaInfo(void * image) {
	return CGImageGetAlphaInfo(image);
}
NSInteger DisplayIsOnline(uint32_t display) {
	return CGDisplayIsOnline(display);
}
int32_t ReleaseDisplayFadeReservation(uint32_t token) {
	return CGReleaseDisplayFadeReservation(token);
}
BOOL ColorSpaceUsesExtendedRange(void * space) {
	return CGColorSpaceUsesExtendedRange(space);
}
NSUInteger BitmapContextGetBitsPerComponent(void * context) {
	return CGBitmapContextGetBitsPerComponent(context);
}
void * PDFContextCreate(void * consumer, CGRect* mediaBox, void * auxiliaryInfo) {
	return CGPDFContextCreate(consumer, mediaBox, auxiliaryInfo);
}
void * ColorSpaceGetName(void * space) {
	return CGColorSpaceGetName(space);
}
void ContextBeginTransparencyLayer(void * c, void * auxiliaryInfo) {
	return CGContextBeginTransparencyLayer(c, auxiliaryInfo);
}
void * ColorGetColorSpace(void * color) {
	return CGColorGetColorSpace(color);
}
int64_t EventGetIntegerValueField(void * event, uint32_t field) {
	return CGEventGetIntegerValueField(event, field);
}
void ContextSetTextPosition(void * c, float x, float y) {
	return CGContextSetTextPosition(c, x, y);
}
void * ColorSpaceCopyICCData(void * space) {
	return CGColorSpaceCopyICCData(space);
}
void ContextSetFont(void * c, void * font) {
	return CGContextSetFont(c, font);
}
void ContextDrawPDFDocument(void * c, CGRect rect, void * document, NSInteger page) {
	return CGContextDrawPDFDocument(c, rect, document, page);
}
CGRect ContextConvertRectToUserSpace(void * c, CGRect rect) {
	return CGContextConvertRectToUserSpace(c, rect);
}
void * ColorRetain(void * color) {
	return CGColorRetain(color);
}
void ContextBeginPath(void * c) {
	return CGContextBeginPath(c);
}
BOOL ColorSpaceIsHLGBased(void * s) {
	return CGColorSpaceIsHLGBased(s);
}
void * ColorCreateGenericGray(float gray, float alpha) {
	return CGColorCreateGenericGray(gray, alpha);
}
void DataProviderRelease(void * provider) {
	return CGDataProviderRelease(provider);
}
void PathAddArc(void * path, CGAffineTransform* m, float x, float y, float radius, float startAngle, float endAngle, BOOL clockwise) {
	return CGPathAddArc(path, m, x, y, radius, startAngle, endAngle, clockwise);
}
int32_t EventGetTypeID() {
	return CGEventGetTypeID();
}
void ContextDrawLinearGradient(void * c, void * gradient, CGPoint startPoint, CGPoint endPoint, uint32_t options) {
	return CGContextDrawLinearGradient(c, gradient, startPoint, endPoint, options);
}
CGRect PDFPageGetBoxRect(void * page, int32_t box) {
	return CGPDFPageGetBoxRect(page, box);
}
void ContextStrokeRectWithWidth(void * c, CGRect rect, float width) {
	return CGContextStrokeRectWithWidth(c, rect, width);
}
void * SessionCopyCurrentDictionary() {
	return CGSessionCopyCurrentDictionary();
}
int32_t ShadingGetTypeID() {
	return CGShadingGetTypeID();
}
void * ColorSpaceCreatePattern(void * baseSpace) {
	return CGColorSpaceCreatePattern(baseSpace);
}
CGRect DisplayBounds(uint32_t display) {
	return CGDisplayBounds(display);
}
void ContextSetTextDrawingMode(void * c, int32_t mode) {
	return CGContextSetTextDrawingMode(c, mode);
}
void * ColorSpaceCreateDeviceGray() {
	return CGColorSpaceCreateDeviceGray();
}
void * ColorSpaceCreateCalibratedGray(float* whitePoint, float* blackPoint, float gamma) {
	return CGColorSpaceCreateCalibratedGray(whitePoint, blackPoint, gamma);
}
void * PathCreateCopyByStrokingPath(void * path, CGAffineTransform* transform, float lineWidth, int32_t lineCap, int32_t lineJoin, float miterLimit) {
	return CGPathCreateCopyByStrokingPath(path, transform, lineWidth, lineCap, lineJoin, miterLimit);
}
BOOL ColorEqualToColor(void * color1, void * color2) {
	return CGColorEqualToColor(color1, color2);
}
BOOL PathContainsPoint(void * path, CGAffineTransform* m, CGPoint point, BOOL eoFill) {
	return CGPathContainsPoint(path, m, point, eoFill);
}
void * FontCopyVariations(void * font) {
	return CGFontCopyVariations(font);
}
void * FontRetain(void * font) {
	return CGFontRetain(font);
}
void * DisplayCopyColorSpace(uint32_t display) {
	return CGDisplayCopyColorSpace(display);
}
NSUInteger ImageGetWidth(void * image) {
	return CGImageGetWidth(image);
}
void * PDFDocumentCreateWithURL(void * url) {
	return CGPDFDocumentCreateWithURL(url);
}
void EventSetFlags(void * event, uint64_t flags) {
	return CGEventSetFlags(event, flags);
}
CGRect RectMake(float x, float y, float width, float height) {
	return CGRectMake(x, y, width, height);
}
NSUInteger ColorGetNumberOfComponents(void * color) {
	return CGColorGetNumberOfComponents(color);
}
void * PathCreateWithRect(CGRect rect, CGAffineTransform* transform) {
	return CGPathCreateWithRect(rect, transform);
}
CGPoint ContextGetPathCurrentPoint(void * c) {
	return CGContextGetPathCurrentPoint(c);
}
void PathCloseSubpath(void * path) {
	return CGPathCloseSubpath(path);
}
void PathAddPath(void * path1, CGAffineTransform* m, void * path2) {
	return CGPathAddPath(path1, m, path2);
}
BOOL RequestScreenCaptureAccess() {
	return CGRequestScreenCaptureAccess();
}
int32_t DisplayStreamGetTypeID() {
	return CGDisplayStreamGetTypeID();
}
BOOL SizeMakeWithDictionaryRepresentation(void * dict, CGSize* size) {
	return CGSizeMakeWithDictionaryRepresentation(dict, size);
}
BOOL ColorSpaceIsWideGamutRGB(void * arg0) {
	return CGColorSpaceIsWideGamutRGB(arg0);
}
void PDFOperatorTableRelease(void * table) {
	return CGPDFOperatorTableRelease(table);
}
NSInteger FontGetUnitsPerEm(void * font) {
	return CGFontGetUnitsPerEm(font);
}
void * PDFContentStreamRetain(void * cs) {
	return CGPDFContentStreamRetain(cs);
}
void ContextSetAlpha(void * c, float alpha) {
	return CGContextSetAlpha(c, alpha);
}
void ContextAddCurveToPoint(void * c, float cp1x, float cp1y, float cp2x, float cp2y, float x, float y) {
	return CGContextAddCurveToPoint(c, cp1x, cp1y, cp2x, cp2y, x, y);
}
void * ColorCreateGenericCMYK(float cyan, float magenta, float yellow, float black, float alpha) {
	return CGColorCreateGenericCMYK(cyan, magenta, yellow, black, alpha);
}
void * ColorSpaceCreateExtendedLinearized(void * space) {
	return CGColorSpaceCreateExtendedLinearized(space);
}
void * ColorSpaceCreateWithICCData(void* data) {
	return CGColorSpaceCreateWithICCData(data);
}
int32_t GetDisplaysWithRect(CGRect rect, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount) {
	return CGGetDisplaysWithRect(rect, maxDisplays, displays, matchingDisplayCount);
}
int32_t PostMouseEvent(CGPoint mouseCursorPosition, NSInteger updateMouseCursorPosition, uint32_t buttonCount, NSInteger mouseButtonDown) {
	return CGPostMouseEvent(mouseCursorPosition, updateMouseCursorPosition, buttonCount, mouseButtonDown);
}
NSUInteger DisplayModeGetWidth(void * mode) {
	return CGDisplayModeGetWidth(mode);
}
void ContextSaveGState(void * c) {
	return CGContextSaveGState(c);
}
NSInteger FontGetCapHeight(void * font) {
	return CGFontGetCapHeight(font);
}
NSUInteger PDFArrayGetCount(void * array) {
	return CGPDFArrayGetCount(array);
}
void FunctionRelease(void * function) {
	return CGFunctionRelease(function);
}
BOOL PDFDocumentUnlockWithPassword(void * document, uint8_t* password) {
	return CGPDFDocumentUnlockWithPassword(document, password);
}
double DisplayRotation(uint32_t display) {
	return CGDisplayRotation(display);
}
CGAffineTransform AffineTransformMakeScale(float sx, float sy) {
	return CGAffineTransformMakeScale(sx, sy);
}
void ContextDrawLayerInRect(void * context, CGRect rect, void * layer) {
	return CGContextDrawLayerInRect(context, rect, layer);
}
BOOL RectIsInfinite(CGRect rect) {
	return CGRectIsInfinite(rect);
}
void ContextShowGlyphsWithAdvances(void * c, CGGlyph* glyphs, CGSize* advances, NSUInteger count) {
	return CGContextShowGlyphsWithAdvances(c, glyphs, advances, count);
}
void * ImageCreateCopyWithColorSpace(void * image, void * space) {
	return CGImageCreateCopyWithColorSpace(image, space);
}
BOOL PreflightPostEventAccess() {
	return CGPreflightPostEventAccess();
}
void ContextRelease(void * c) {
	return CGContextRelease(c);
}
int32_t DataConsumerGetTypeID() {
	return CGDataConsumerGetTypeID();
}
int32_t CompleteDisplayConfiguration(void * config, uint32_t option) {
	return CGCompleteDisplayConfiguration(config, option);
}
void PDFContextEndTag(void * context) {
	return CGPDFContextEndTag(context);
}
void * ColorSpaceCopyICCProfile(void * space) {
	return CGColorSpaceCopyICCProfile(space);
}
void * ColorCreateGenericGrayGamma2_2(float gray, float alpha) {
	return CGColorCreateGenericGrayGamma2_2(gray, alpha);
}
CGRect PathGetBoundingBox(void * path) {
	return CGPathGetBoundingBox(path);
}
void PDFDocumentGetVersion(void * document, NSInteger* majorVersion, NSInteger* minorVersion) {
	return CGPDFDocumentGetVersion(document, majorVersion, minorVersion);
}
void PathAddRoundedRect(void * path, CGAffineTransform* transform, CGRect rect, float cornerWidth, float cornerHeight) {
	return CGPathAddRoundedRect(path, transform, rect, cornerWidth, cornerHeight);
}
void * DisplayAvailableModes(uint32_t dsp) {
	return CGDisplayAvailableModes(dsp);
}
int32_t WindowLevelForKey(int32_t key) {
	return CGWindowLevelForKey(key);
}
void ContextSetRenderingIntent(void * c, int32_t intent) {
	return CGContextSetRenderingIntent(c, intent);
}
void ContextSetShadowWithColor(void * c, CGSize offset, float blur, void * color) {
	return CGContextSetShadowWithColor(c, offset, blur, color);
}
int32_t PatternGetTypeID() {
	return CGPatternGetTypeID();
}
void * DisplayModeRetain(void * mode) {
	return CGDisplayModeRetain(mode);
}
uint64_t EventGetTimestamp(void * event) {
	return CGEventGetTimestamp(event);
}
void * PathCreateMutableCopy(void * path) {
	return CGPathCreateMutableCopy(path);
}
NSInteger PDFDocumentGetRotationAngle(void * document, NSInteger page) {
	return CGPDFDocumentGetRotationAngle(document, page);
}
CGAffineTransform ContextGetCTM(void * c) {
	return CGContextGetCTM(c);
}
void * EventCreateData(void * allocator, void * event) {
	return CGEventCreateData(allocator, event);
}
NSInteger DisplayIsStereo(uint32_t display) {
	return CGDisplayIsStereo(display);
}
void ContextConcatCTM(void * c, CGAffineTransform transform) {
	return CGContextConcatCTM(c, transform);
}
CGSize ContextConvertSizeToUserSpace(void * c, CGSize size) {
	return CGContextConvertSizeToUserSpace(c, size);
}
BOOL PDFScannerPopInteger(void * scanner, CGPDFInteger* value) {
	return CGPDFScannerPopInteger(scanner, value);
}
void PDFContextAddDocumentMetadata(void * context, void * metadata) {
	return CGPDFContextAddDocumentMetadata(context, metadata);
}
CGSize DisplayScreenSize(uint32_t display) {
	return CGDisplayScreenSize(display);
}
void ColorSpaceGetColorTable(void * space, uint8_t* table) {
	return CGColorSpaceGetColorTable(space, table);
}
void * PDFPageRetain(void * page) {
	return CGPDFPageRetain(page);
}
void ContextFillEllipseInRect(void * c, CGRect rect) {
	return CGContextFillEllipseInRect(c, rect);
}
void * DisplayCopyDisplayMode(uint32_t display) {
	return CGDisplayCopyDisplayMode(display);
}
void PDFContextSetDestinationForRect(void * context, void * name, CGRect rect) {
	return CGPDFContextSetDestinationForRect(context, name, rect);
}
int32_t GetDisplaysWithOpenGLDisplayMask(uint32_t mask, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount) {
	return CGGetDisplaysWithOpenGLDisplayMask(mask, maxDisplays, displays, matchingDisplayCount);
}
BOOL PDFDictionaryGetArray(void * dict, uint8_t* key, void * value) {
	return CGPDFDictionaryGetArray(dict, key, value);
}
void * ShadingCreateAxial(void * space, CGPoint start, CGPoint end, void * function, BOOL extendStart, BOOL extendEnd) {
	return CGShadingCreateAxial(space, start, end, function, extendStart, extendEnd);
}
void ContextSetShadow(void * c, CGSize offset, float blur) {
	return CGContextSetShadow(c, offset, blur);
}
BOOL PointMakeWithDictionaryRepresentation(void * dict, CGPoint* point) {
	return CGPointMakeWithDictionaryRepresentation(dict, point);
}
CGPoint ContextGetTextPosition(void * c) {
	return CGContextGetTextPosition(c);
}
NSUInteger BitmapContextGetBitsPerPixel(void * context) {
	return CGBitmapContextGetBitsPerPixel(context);
}
void ContextAddArc(void * c, float x, float y, float radius, float startAngle, float endAngle, NSInteger clockwise) {
	return CGContextAddArc(c, x, y, radius, startAngle, endAngle, clockwise);
}
void EventSetType(void * event, uint32_t type_) {
	return CGEventSetType(event, type_);
}
void * EventCreateScrollWheelEvent2(void * source, uint32_t units, uint32_t wheelCount, int32_t wheel1, int32_t wheel2, int32_t wheel3) {
	return CGEventCreateScrollWheelEvent2(source, units, wheelCount, wheel1, wheel2, wheel3);
}
int32_t ColorSpaceGetModel(void * space) {
	return CGColorSpaceGetModel(space);
}
void * ImageCreateWithPNGDataProvider(void * source, float* decode, BOOL shouldInterpolate, int32_t intent) {
	return CGImageCreateWithPNGDataProvider(source, decode, shouldInterpolate, intent);
}
void PathAddEllipseInRect(void * path, CGAffineTransform* m, CGRect rect) {
	return CGPathAddEllipseInRect(path, m, rect);
}
void * EventCreateFromData(void * allocator, void * data) {
	return CGEventCreateFromData(allocator, data);
}
CGRect PDFDocumentGetArtBox(void * document, NSInteger page) {
	return CGPDFDocumentGetArtBox(document, page);
}
void * ColorSpaceCreateDeviceCMYK() {
	return CGColorSpaceCreateDeviceCMYK();
}
int32_t ConfigureDisplayFadeEffect(void * config, float fadeOutSeconds, float fadeInSeconds, float fadeRed, float fadeGreen, float fadeBlue) {
	return CGConfigureDisplayFadeEffect(config, fadeOutSeconds, fadeInSeconds, fadeRed, fadeGreen, fadeBlue);
}
int32_t DisplayFade(uint32_t token, float duration, float startBlend, float endBlend, float redBlend, float greenBlend, float blueBlend, NSInteger synchronous) {
	return CGDisplayFade(token, duration, startBlend, endBlend, redBlend, greenBlend, blueBlend, synchronous);
}
BOOL SizeEqualToSize(CGSize size1, CGSize size2) {
	return CGSizeEqualToSize(size1, size2);
}
void ContextSetShouldSubpixelQuantizeFonts(void * c, BOOL shouldSubpixelQuantizeFonts) {
	return CGContextSetShouldSubpixelQuantizeFonts(c, shouldSubpixelQuantizeFonts);
}
CGPoint EventGetLocation(void * event) {
	return CGEventGetLocation(event);
}
BOOL PSConverterAbort(void * converter) {
	return CGPSConverterAbort(converter);
}
int32_t DisplayStreamUpdateGetTypeID() {
	return CGDisplayStreamUpdateGetTypeID();
}
void ContextScaleCTM(void * c, float sx, float sy) {
	return CGContextScaleCTM(c, sx, sy);
}
void * SizeCreateDictionaryRepresentation(CGSize size) {
	return CGSizeCreateDictionaryRepresentation(size);
}
BOOL EventSourceButtonState(int32_t stateID, uint32_t button) {
	return CGEventSourceButtonState(stateID, button);
}
void EventTapEnable(void * tap, BOOL enable) {
	return CGEventTapEnable(tap, enable);
}
CGAffineTransform AffineTransformMakeRotation(float angle) {
	return CGAffineTransformMakeRotation(angle);
}
void * ColorSpaceCreateICCBased(NSUInteger nComponents, float* range_, void * profile, void * alternate) {
	return CGColorSpaceCreateICCBased(nComponents, range_, profile, alternate);
}
int32_t AssociateMouseAndMouseCursorPosition(NSInteger connected) {
	return CGAssociateMouseAndMouseCursorPosition(connected);
}
void * PDFDocumentRetain(void * document) {
	return CGPDFDocumentRetain(document);
}
CGPoint PathGetCurrentPoint(void * path) {
	return CGPathGetCurrentPoint(path);
}
void ContextSetRGBStrokeColor(void * c, float red, float green, float blue, float alpha) {
	return CGContextSetRGBStrokeColor(c, red, green, blue, alpha);
}
BOOL PreflightScreenCaptureAccess() {
	return CGPreflightScreenCaptureAccess();
}
int32_t SetDisplayTransferByTable(uint32_t display, uint32_t tableSize, CGGammaValue* redTable, CGGammaValue* greenTable, CGGammaValue* blueTable) {
	return CGSetDisplayTransferByTable(display, tableSize, redTable, greenTable, blueTable);
}
void ContextSetFlatness(void * c, float flatness) {
	return CGContextSetFlatness(c, flatness);
}
NSInteger DisplayIsActive(uint32_t display) {
	return CGDisplayIsActive(display);
}
void ContextSetPatternPhase(void * c, CGSize phase) {
	return CGContextSetPatternPhase(c, phase);
}
void PDFScannerRelease(void * scanner) {
	return CGPDFScannerRelease(scanner);
}
void ContextSetLineJoin(void * c, int32_t join) {
	return CGContextSetLineJoin(c, join);
}
void * FontCopyTableTags(void * font) {
	return CGFontCopyTableTags(font);
}
int32_t GetActiveDisplayList(uint32_t maxDisplays, CGDirectDisplayID* activeDisplays, uint32_t* displayCount) {
	return CGGetActiveDisplayList(maxDisplays, activeDisplays, displayCount);
}
BOOL PathIsRect(void * path, CGRect* rect) {
	return CGPathIsRect(path, rect);
}
void * DataProviderCreateWithURL(void * url) {
	return CGDataProviderCreateWithURL(url);
}
void * EventSourceCreate(int32_t stateID) {
	return CGEventSourceCreate(stateID);
}
int32_t ContextGetInterpolationQuality(void * c) {
	return CGContextGetInterpolationQuality(c);
}
void * ColorConversionInfoCreate(void * src, void * dst) {
	return CGColorConversionInfoCreate(src, dst);
}
CGRect PDFDocumentGetTrimBox(void * document, NSInteger page) {
	return CGPDFDocumentGetTrimBox(document, page);
}
BOOL RectContainsRect(CGRect rect1, CGRect rect2) {
	return CGRectContainsRect(rect1, rect2);
}
void * FontCreatePostScriptEncoding(void * font, CGGlyph* encoding) {
	return CGFontCreatePostScriptEncoding(font, encoding);
}
void ContextDrawLayerAtPoint(void * context, CGPoint point, void * layer) {
	return CGContextDrawLayerAtPoint(context, point, layer);
}
BOOL PDFArrayGetStream(void * array, NSUInteger index, void * value) {
	return CGPDFArrayGetStream(array, index, value);
}
void ContextShowTextAtPoint(void * c, float x, float y, uint8_t* string_, NSUInteger length) {
	return CGContextShowTextAtPoint(c, x, y, string_, length);
}
void EventSourceSetPixelsPerLine(void * source, double pixelsPerLine) {
	return CGEventSourceSetPixelsPerLine(source, pixelsPerLine);
}
void * EventCreateKeyboardEvent(void * source, uint16_t virtualKey, BOOL keyDown) {
	return CGEventCreateKeyboardEvent(source, virtualKey, keyDown);
}
uint32_t EventSourceCounterForEventType(int32_t stateID, uint32_t eventType) {
	return CGEventSourceCounterForEventType(stateID, eventType);
}
void PathAddRects(void * path, CGAffineTransform* m, CGRect* rects, NSUInteger count) {
	return CGPathAddRects(path, m, rects, count);
}
void EventSetTimestamp(void * event, uint64_t timestamp) {
	return CGEventSetTimestamp(event, timestamp);
}
BOOL EventSourceKeyState(int32_t stateID, uint16_t key) {
	return CGEventSourceKeyState(stateID, key);
}
float RectGetHeight(CGRect rect) {
	return CGRectGetHeight(rect);
}
void ContextShowGlyphsAtPoint(void * c, float x, float y, CGGlyph* glyphs, NSUInteger count) {
	return CGContextShowGlyphsAtPoint(c, x, y, glyphs, count);
}
BOOL PDFDictionaryGetInteger(void * dict, uint8_t* key, CGPDFInteger* value) {
	return CGPDFDictionaryGetInteger(dict, key, value);
}
double EventSourceSecondsSinceLastEventType(int32_t stateID, uint32_t eventType) {
	return CGEventSourceSecondsSinceLastEventType(stateID, eventType);
}
CGRect RectIntegral(CGRect rect) {
	return CGRectIntegral(rect);
}
void ContextSetFillColorSpace(void * c, void * space) {
	return CGContextSetFillColorSpace(c, space);
}
BOOL PDFArrayGetNumber(void * array, NSUInteger index, CGPDFReal* value) {
	return CGPDFArrayGetNumber(array, index, value);
}
void * ColorCreateWithPattern(void * space, void * pattern, float* components) {
	return CGColorCreateWithPattern(space, pattern, components);
}
void ContextTranslateCTM(void * c, float tx, float ty) {
	return CGContextTranslateCTM(c, tx, ty);
}
NSInteger DisplayUsesOpenGLAcceleration(uint32_t display) {
	return CGDisplayUsesOpenGLAcceleration(display);
}
void ContextSetBlendMode(void * c, int32_t mode) {
	return CGContextSetBlendMode(c, mode);
}
void GetLastMouseDelta(int32_t* deltaX, int32_t* deltaY) {
	return CGGetLastMouseDelta(deltaX, deltaY);
}
BOOL FontGetGlyphBBoxes(void * font, CGGlyph* glyphs, NSUInteger count, CGRect* bboxes) {
	return CGFontGetGlyphBBoxes(font, glyphs, count, bboxes);
}
int32_t GetDisplaysWithPoint(CGPoint point, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount) {
	return CGGetDisplaysWithPoint(point, maxDisplays, displays, matchingDisplayCount);
}
void ImageRelease(void * image) {
	return CGImageRelease(image);
}
float RectGetMinX(CGRect rect) {
	return CGRectGetMinX(rect);
}
void ContextReplacePathWithStrokedPath(void * c) {
	return CGContextReplacePathWithStrokedPath(c);
}
NSUInteger BitmapContextGetBytesPerRow(void * context) {
	return CGBitmapContextGetBytesPerRow(context);
}
void * WindowListCreate(uint32_t option, uint32_t relativeToWindow) {
	return CGWindowListCreate(option, relativeToWindow);
}
CGSize SizeApplyAffineTransform(CGSize size, CGAffineTransform t) {
	return CGSizeApplyAffineTransform(size, t);
}
void EventTapPostEvent(void * proxy, void * event) {
	return CGEventTapPostEvent(proxy, event);
}
NSInteger DisplayIsInMirrorSet(uint32_t display) {
	return CGDisplayIsInMirrorSet(display);
}
BOOL ColorSpaceIsPQBased(void * s) {
	return CGColorSpaceIsPQBased(s);
}
void ContextClearRect(void * c, CGRect rect) {
	return CGContextClearRect(c, rect);
}
int32_t DisplayRelease(uint32_t display) {
	return CGDisplayRelease(display);
}
void * PDFPageGetDocument(void * page) {
	return CGPDFPageGetDocument(page);
}
int32_t GetDisplayTransferByTable(uint32_t display, uint32_t capacity, CGGammaValue* redTable, CGGammaValue* greenTable, CGGammaValue* blueTable, uint32_t* sampleCount) {
	return CGGetDisplayTransferByTable(display, capacity, redTable, greenTable, blueTable, sampleCount);
}
BOOL AffineTransformEqualToTransform(CGAffineTransform t1, CGAffineTransform t2) {
	return CGAffineTransformEqualToTransform(t1, t2);
}
void * PDFScannerGetContentStream(void * scanner) {
	return CGPDFScannerGetContentStream(scanner);
}
void PathAddLineToPoint(void * path, CGAffineTransform* m, float x, float y) {
	return CGPathAddLineToPoint(path, m, x, y);
}
BOOL PreflightListenEventAccess() {
	return CGPreflightListenEventAccess();
}
uint32_t EventGetType(void * event) {
	return CGEventGetType(event);
}
BOOL PDFScannerPopNumber(void * scanner, CGPDFReal* value) {
	return CGPDFScannerPopNumber(scanner, value);
}
void * ColorSpaceCreateLab(float* whitePoint, float* blackPoint, float* range_) {
	return CGColorSpaceCreateLab(whitePoint, blackPoint, range_);
}
void PDFContextEndPage(void * context) {
	return CGPDFContextEndPage(context);
}
CGRect ContextGetClipBoundingBox(void * c) {
	return CGContextGetClipBoundingBox(c);
}
void * PDFOperatorTableCreate() {
	return CGPDFOperatorTableCreate();
}
void PathAddRect(void * path, CGAffineTransform* m, CGRect rect) {
	return CGPathAddRect(path, m, rect);
}
CGPoint PointApplyAffineTransform(CGPoint point, CGAffineTransform t) {
	return CGPointApplyAffineTransform(point, t);
}
void * ImageCreateWithJPEGDataProvider(void * source, float* decode, BOOL shouldInterpolate, int32_t intent) {
	return CGImageCreateWithJPEGDataProvider(source, decode, shouldInterpolate, intent);
}
int32_t EventSourceGetTypeID() {
	return CGEventSourceGetTypeID();
}
void * ImageCreate(NSUInteger width, NSUInteger height, NSUInteger bitsPerComponent, NSUInteger bitsPerPixel, NSUInteger bytesPerRow, void * space, uint32_t bitmapInfo, void * provider, float* decode, BOOL shouldInterpolate, int32_t intent) {
	return CGImageCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, bitmapInfo, provider, decode, shouldInterpolate, intent);
}
BOOL RectIsNull(CGRect rect) {
	return CGRectIsNull(rect);
}
int32_t EventSourceGetSourceStateID(void * source) {
	return CGEventSourceGetSourceStateID(source);
}
void * GradientRetain(void * gradient) {
	return CGGradientRetain(gradient);
}
BOOL PDFArrayGetBoolean(void * array, NSUInteger index, CGPDFBoolean* value) {
	return CGPDFArrayGetBoolean(array, index, value);
}
void ContextSetFillColor(void * c, float* components) {
	return CGContextSetFillColor(c, components);
}
void * ColorGetPattern(void * color) {
	return CGColorGetPattern(color);
}
CGRect PathGetPathBoundingBox(void * path) {
	return CGPathGetPathBoundingBox(path);
}
int32_t CaptureAllDisplaysWithOptions(uint32_t options) {
	return CGCaptureAllDisplaysWithOptions(options);
}
void * WindowListCreateDescriptionFromArray(void * windowArray) {
	return CGWindowListCreateDescriptionFromArray(windowArray);
}
void * PDFContentStreamCreateWithStream(void * stream, void * streamResources, void * parent) {
	return CGPDFContentStreamCreateWithStream(stream, streamResources, parent);
}
void ContextDrawRadialGradient(void * c, void * gradient, CGPoint startCenter, float startRadius, CGPoint endCenter, float endRadius, uint32_t options) {
	return CGContextDrawRadialGradient(c, gradient, startCenter, startRadius, endCenter, endRadius, options);
}
void * PDFStringCopyDate(void * string_) {
	return CGPDFStringCopyDate(string_);
}
void * WindowServerCFMachPort() {
	return CGWindowServerCFMachPort();
}
int32_t GetOnlineDisplayList(uint32_t maxDisplays, CGDirectDisplayID* onlineDisplays, uint32_t* displayCount) {
	return CGGetOnlineDisplayList(maxDisplays, onlineDisplays, displayCount);
}
CGPoint ContextConvertPointToDeviceSpace(void * c, CGPoint point) {
	return CGContextConvertPointToDeviceSpace(c, point);
}
void * PathCreateCopy(void * path) {
	return CGPathCreateCopy(path);
}
void * BitmapContextCreateImage(void * context) {
	return CGBitmapContextCreateImage(context);
}
NSUInteger ColorSpaceGetNumberOfComponents(void * space) {
	return CGColorSpaceGetNumberOfComponents(space);
}
void * FunctionRetain(void * function) {
	return CGFunctionRetain(function);
}
int32_t GetDisplayTransferByFormula(uint32_t display, CGGammaValue* redMin, CGGammaValue* redMax, CGGammaValue* redGamma, CGGammaValue* greenMin, CGGammaValue* greenMax, CGGammaValue* greenGamma, CGGammaValue* blueMin, CGGammaValue* blueMax, CGGammaValue* blueGamma) {
	return CGGetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma);
}
BOOL PDFScannerPopStream(void * scanner, void * value) {
	return CGPDFScannerPopStream(scanner, value);
}
void * ColorSpaceRetain(void * space) {
	return CGColorSpaceRetain(space);
}
void * ColorCreateSRGB(float red, float green, float blue, float alpha) {
	return CGColorCreateSRGB(red, green, blue, alpha);
}
CGRect* DisplayStreamUpdateGetRects(void * updateRef, int32_t rectType, NSUInteger* rectCount) {
	return CGDisplayStreamUpdateGetRects(updateRef, rectType, rectCount);
}
void * ImageGetColorSpace(void * image) {
	return CGImageGetColorSpace(image);
}
int32_t SetLocalEventsFilterDuringSuppressionState(uint32_t filter, uint32_t state) {
	return CGSetLocalEventsFilterDuringSuppressionState(filter, state);
}
void * ShadingCreateRadial(void * space, CGPoint start, float startRadius, CGPoint end, float endRadius, void * function, BOOL extendStart, BOOL extendEnd) {
	return CGShadingCreateRadial(space, start, startRadius, end, endRadius, function, extendStart, extendEnd);
}
void RectDivide(CGRect rect, CGRect* slice, CGRect* remainder, float amount, uint32_t edge) {
	return CGRectDivide(rect, slice, remainder, amount, edge);
}
void ContextSetAllowsAntialiasing(void * c, BOOL allowsAntialiasing) {
	return CGContextSetAllowsAntialiasing(c, allowsAntialiasing);
}
NSUInteger BitmapContextGetHeight(void * context) {
	return CGBitmapContextGetHeight(context);
}
float ColorGetAlpha(void * color) {
	return CGColorGetAlpha(color);
}
NSUInteger ImageGetHeight(void * image) {
	return CGImageGetHeight(image);
}
void * ColorSpaceCreateWithName(void * name) {
	return CGColorSpaceCreateWithName(name);
}
void ContextAddQuadCurveToPoint(void * c, float cpx, float cpy, float x, float y) {
	return CGContextAddQuadCurveToPoint(c, cpx, cpy, x, y);
}
int32_t ConfigureDisplayWithDisplayMode(void * config, uint32_t display, void * mode, void * options) {
	return CGConfigureDisplayWithDisplayMode(config, display, mode, options);
}
void * DataProviderCopyData(void * provider) {
	return CGDataProviderCopyData(provider);
}
NSInteger DisplayIsInHWMirrorSet(uint32_t display) {
	return CGDisplayIsInHWMirrorSet(display);
}
uint32_t EventSourceGetLocalEventsFilterDuringSuppressionState(void * source, uint32_t state) {
	return CGEventSourceGetLocalEventsFilterDuringSuppressionState(source, state);
}
void * FontCopyTableForTag(void * font, uint32_t tag) {
	return CGFontCopyTableForTag(font, tag);
}
uint8_t* PDFStringGetBytePtr(void * string_) {
	return CGPDFStringGetBytePtr(string_);
}
CGRect ContextGetPathBoundingBox(void * c) {
	return CGContextGetPathBoundingBox(c);
}
void * LayerRetain(void * layer) {
	return CGLayerRetain(layer);
}
void ContextSetAllowsFontSubpixelPositioning(void * c, BOOL allowsFontSubpixelPositioning) {
	return CGContextSetAllowsFontSubpixelPositioning(c, allowsFontSubpixelPositioning);
}
int32_t ColorGetTypeID() {
	return CGColorGetTypeID();
}
uint32_t ImageGetByteOrderInfo(void * image) {
	return CGImageGetByteOrderInfo(image);
}
void * PDFDocumentCreateWithProvider(void * provider) {
	return CGPDFDocumentCreateWithProvider(provider);
}
BOOL FontGetGlyphAdvances(void * font, CGGlyph* glyphs, NSUInteger count, NSInteger* advances) {
	return CGFontGetGlyphAdvances(font, glyphs, count, advances);
}
uint32_t ImageGetPixelFormatInfo(void * image) {
	return CGImageGetPixelFormatInfo(image);
}
void * FontCopyVariationAxes(void * font) {
	return CGFontCopyVariationAxes(font);
}
CGRect PDFDocumentGetMediaBox(void * document, NSInteger page) {
	return CGPDFDocumentGetMediaBox(document, page);
}
void* ColorSpaceCopyPropertyList(void * space) {
	return CGColorSpaceCopyPropertyList(space);
}
void ContextSetStrokePattern(void * c, void * pattern, float* components) {
	return CGContextSetStrokePattern(c, pattern, components);
}
void * DisplayBestModeForParameters(uint32_t display, NSUInteger bitsPerPixel, NSUInteger width, NSUInteger height, boolean_t* exactMatch) {
	return CGDisplayBestModeForParameters(display, bitsPerPixel, width, height, exactMatch);
}
BOOL PDFDictionaryGetDictionary(void * dict, uint8_t* key, void * value) {
	return CGPDFDictionaryGetDictionary(dict, key, value);
}
void ContextClipToRect(void * c, CGRect rect) {
	return CGContextClipToRect(c, rect);
}
void ContextResetClip(void * c) {
	return CGContextResetClip(c);
}
BOOL PDFArrayGetDictionary(void * array, NSUInteger index, void * value) {
	return CGPDFArrayGetDictionary(array, index, value);
}
void PathAddQuadCurveToPoint(void * path, CGAffineTransform* m, float cpx, float cpy, float x, float y) {
	return CGPathAddQuadCurveToPoint(path, m, cpx, cpy, x, y);
}
int32_t WarpMouseCursorPosition(CGPoint newCursorPosition) {
	return CGWarpMouseCursorPosition(newCursorPosition);
}
int32_t DisplayMoveCursorToPoint(uint32_t display, CGPoint point) {
	return CGDisplayMoveCursorToPoint(display, point);
}
void * WindowListCopyWindowInfo(uint32_t option, uint32_t relativeToWindow) {
	return CGWindowListCopyWindowInfo(option, relativeToWindow);
}
uint32_t ShieldingWindowID(uint32_t display) {
	return CGShieldingWindowID(display);
}
void PDFContextSetOutline(void * context, void * outline) {
	return CGPDFContextSetOutline(context, outline);
}
void PathAddArcToPoint(void * path, CGAffineTransform* m, float x1, float y1, float x2, float y2, float radius) {
	return CGPathAddArcToPoint(path, m, x1, y1, x2, y2, radius);
}
NSUInteger ImageGetBitsPerComponent(void * image) {
	return CGImageGetBitsPerComponent(image);
}
int32_t CaptureAllDisplays() {
	return CGCaptureAllDisplays();
}
double EventSourceGetPixelsPerLine(void * source) {
	return CGEventSourceGetPixelsPerLine(source);
}
CGRect RectInset(CGRect rect, float dx, float dy) {
	return CGRectInset(rect, dx, dy);
}
int32_t BeginDisplayConfiguration(void * config) {
	return CGBeginDisplayConfiguration(config);
}
BOOL PDFArrayGetNull(void * array, NSUInteger index) {
	return CGPDFArrayGetNull(array, index);
}
float* ImageGetDecode(void * image) {
	return CGImageGetDecode(image);
}
void* DataProviderGetInfo(void * provider) {
	return CGDataProviderGetInfo(provider);
}
void * DisplayStreamUpdateCreateMergedUpdate(void * firstUpdate, void * secondUpdate) {
	return CGDisplayStreamUpdateCreateMergedUpdate(firstUpdate, secondUpdate);
}
void * PDFStreamCopyData(void * stream, CGPDFDataFormat* format) {
	return CGPDFStreamCopyData(stream, format);
}
void * PatternRetain(void * pattern) {
	return CGPatternRetain(pattern);
}
BOOL PDFDictionaryGetObject(void * dict, uint8_t* key, void * value) {
	return CGPDFDictionaryGetObject(dict, key, value);
}
void * BitmapContextGetColorSpace(void * context) {
	return CGBitmapContextGetColorSpace(context);
}
BOOL RectIsEmpty(CGRect rect) {
	return CGRectIsEmpty(rect);
}
float RectGetMinY(CGRect rect) {
	return CGRectGetMinY(rect);
}
void * WindowListCreateImage(CGRect screenBounds, uint32_t listOption, uint32_t windowID, uint32_t imageOption) {
	return CGWindowListCreateImage(screenBounds, listOption, windowID, imageOption);
}
int32_t PostScrollWheelEvent(uint32_t wheelCount, int32_t wheel1) {
	return CGPostScrollWheelEvent(wheelCount, wheel1);
}
int32_t ColorSpaceGetTypeID() {
	return CGColorSpaceGetTypeID();
}
void * ShadingRetain(void * shading) {
	return CGShadingRetain(shading);
}
void * ColorCreateCopy(void * color) {
	return CGColorCreateCopy(color);
}
void ContextBeginTransparencyLayerWithRect(void * c, CGRect rect, void * auxInfo) {
	return CGContextBeginTransparencyLayerWithRect(c, rect, auxInfo);
}
void PDFPageRelease(void * page) {
	return CGPDFPageRelease(page);
}
int32_t ImageGetTypeID() {
	return CGImageGetTypeID();
}
BOOL PDFScannerPopDictionary(void * scanner, void * value) {
	return CGPDFScannerPopDictionary(scanner, value);
}
void ContextRotateCTM(void * c, float angle) {
	return CGContextRotateCTM(c, angle);
}
void ContextEndTransparencyLayer(void * c) {
	return CGContextEndTransparencyLayer(c);
}
void EventSourceSetLocalEventsSuppressionInterval(void * source, double seconds) {
	return CGEventSourceSetLocalEventsSuppressionInterval(source, seconds);
}
CGRect RectIntersection(CGRect r1, CGRect r2) {
	return CGRectIntersection(r1, r2);
}
int32_t DisplayModeGetIODisplayModeID(void * mode) {
	return CGDisplayModeGetIODisplayModeID(mode);
}
uint32_t DisplayPrimaryDisplay(uint32_t display) {
	return CGDisplayPrimaryDisplay(display);
}
void PDFContextAddDestinationAtPoint(void * context, void * name, CGPoint point) {
	return CGPDFContextAddDestinationAtPoint(context, name, point);
}
NSInteger FontGetDescent(void * font) {
	return CGFontGetDescent(font);
}
void ContextSetStrokeColor(void * c, float* components) {
	return CGContextSetStrokeColor(c, components);
}
BOOL PDFDictionaryGetString(void * dict, uint8_t* key, void * value) {
	return CGPDFDictionaryGetString(dict, key, value);
}
void DisplayModeRelease(void * mode) {
	return CGDisplayModeRelease(mode);
}
void ContextAddRects(void * c, CGRect* rects, NSUInteger count) {
	return CGContextAddRects(c, rects, count);
}
void ContextSelectFont(void * c, uint8_t* name, float size, int32_t textEncoding) {
	return CGContextSelectFont(c, name, size, textEncoding);
}
void * PathCreateMutableCopyByTransformingPath(void * path, CGAffineTransform* transform) {
	return CGPathCreateMutableCopyByTransformingPath(path, transform);
}
CGAffineTransform AffineTransformMakeTranslation(float tx, float ty) {
	return CGAffineTransformMakeTranslation(tx, ty);
}
void ContextSetCharacterSpacing(void * c, float spacing) {
	return CGContextSetCharacterSpacing(c, spacing);
}
uint32_t DisplayModeGetIOFlags(void * mode) {
	return CGDisplayModeGetIOFlags(mode);
}
void PDFContextClose(void * context) {
	return CGPDFContextClose(context);
}
void * DataProviderRetain(void * provider) {
	return CGDataProviderRetain(provider);
}
BOOL ColorSpaceSupportsOutput(void * space) {
	return CGColorSpaceSupportsOutput(space);
}
void EventSetDoubleValueField(void * event, uint32_t field, double value) {
	return CGEventSetDoubleValueField(event, field, value);
}
void ContextDrawPDFPage(void * c, void * page) {
	return CGContextDrawPDFPage(c, page);
}
void * PDFContextCreateWithURL(void * url, CGRect* mediaBox, void * auxiliaryInfo) {
	return CGPDFContextCreateWithURL(url, mediaBox, auxiliaryInfo);
}
int32_t PDFPageGetTypeID() {
	return CGPDFPageGetTypeID();
}
CGRect PDFDocumentGetBleedBox(void * document, NSInteger page) {
	return CGPDFDocumentGetBleedBox(document, page);
}
void * PDFContentStreamCreateWithPage(void * page) {
	return CGPDFContentStreamCreateWithPage(page);
}
void * FontCreateWithDataProvider(void * provider) {
	return CGFontCreateWithDataProvider(provider);
}
BOOL PDFDictionaryGetNumber(void * dict, uint8_t* key, CGPDFReal* value) {
	return CGPDFDictionaryGetNumber(dict, key, value);
}
void ColorSpaceRelease(void * space) {
	return CGColorSpaceRelease(space);
}
void * ImageCreateCopy(void * image) {
	return CGImageCreateCopy(image);
}
void ContextEndPage(void * c) {
	return CGContextEndPage(c);
}
BOOL ColorSpaceIsHDR(void * arg0) {
	return CGColorSpaceIsHDR(arg0);
}
void * PDFDocumentGetID(void * document) {
	return CGPDFDocumentGetID(document);
}
void ContextSetTextMatrix(void * c, CGAffineTransform t) {
	return CGContextSetTextMatrix(c, t);
}
NSUInteger FontGetNumberOfGlyphs(void * font) {
	return CGFontGetNumberOfGlyphs(font);
}
void ContextClipToMask(void * c, CGRect rect, void * mask) {
	return CGContextClipToMask(c, rect, mask);
}
void * PDFPageGetDictionary(void * page) {
	return CGPDFPageGetDictionary(page);
}
uint32_t DisplayGammaTableCapacity(uint32_t display) {
	return CGDisplayGammaTableCapacity(display);
}
void EventSetIntegerValueField(void * event, uint32_t field, int64_t value) {
	return CGEventSetIntegerValueField(event, field, value);
}
void ContextAddPath(void * c, void * path) {
	return CGContextAddPath(c, path);
}
CGAffineTransform ContextGetTextMatrix(void * c) {
	return CGContextGetTextMatrix(c);
}
void ContextStrokePath(void * c) {
	return CGContextStrokePath(c);
}
NSUInteger DisplayPixelsWide(uint32_t display) {
	return CGDisplayPixelsWide(display);
}
void ContextSetShouldAntialias(void * c, BOOL shouldAntialias) {
	return CGContextSetShouldAntialias(c, shouldAntialias);
}
int32_t DisplaySwitchToMode(uint32_t display, void * mode) {
	return CGDisplaySwitchToMode(display, mode);
}
BOOL EventTapIsEnabled(void * tap) {
	return CGEventTapIsEnabled(tap);
}
BOOL PDFDictionaryGetStream(void * dict, uint8_t* key, void * value) {
	return CGPDFDictionaryGetStream(dict, key, value);
}
void * ColorSpaceCreateIndexed(void * baseSpace, NSUInteger lastIndex, uint8_t* colorTable) {
	return CGColorSpaceCreateIndexed(baseSpace, lastIndex, colorTable);
}
uint8_t* PDFTagTypeGetName(int32_t tagType) {
	return CGPDFTagTypeGetName(tagType);
}
void ContextSetGrayFillColor(void * c, float gray, float alpha) {
	return CGContextSetGrayFillColor(c, gray, alpha);
}
void * PDFContentStreamGetResource(void * cs, uint8_t* category, uint8_t* name) {
	return CGPDFContentStreamGetResource(cs, category, name);
}
BOOL PDFDictionaryGetBoolean(void * dict, uint8_t* key, CGPDFBoolean* value) {
	return CGPDFDictionaryGetBoolean(dict, key, value);
}
BOOL PSConverterConvert(void * converter, void * provider, void * consumer, void * options) {
	return CGPSConverterConvert(converter, provider, consumer, options);
}
void PDFContentStreamRelease(void * cs) {
	return CGPDFContentStreamRelease(cs);
}
CGRect PDFDocumentGetCropBox(void * document, NSInteger page) {
	return CGPDFDocumentGetCropBox(document, page);
}
uint32_t DisplayIDToOpenGLDisplayMask(uint32_t display) {
	return CGDisplayIDToOpenGLDisplayMask(display);
}
float RectGetMaxX(CGRect rect) {
	return CGRectGetMaxX(rect);
}
BOOL ContextIsPathEmpty(void * c) {
	return CGContextIsPathEmpty(c);
}
int32_t ColorConversionInfoGetTypeID() {
	return CGColorConversionInfoGetTypeID();
}
NSUInteger ImageGetBitsPerPixel(void * image) {
	return CGImageGetBitsPerPixel(image);
}
BOOL RectMakeWithDictionaryRepresentation(void * dict, CGRect* rect) {
	return CGRectMakeWithDictionaryRepresentation(dict, rect);
}
NSUInteger DisplayModeGetPixelHeight(void * mode) {
	return CGDisplayModeGetPixelHeight(mode);
}
void * ColorSpaceCreateDeviceRGB() {
	return CGColorSpaceCreateDeviceRGB();
}
void ContextSetStrokeColorWithColor(void * c, void * color) {
	return CGContextSetStrokeColorWithColor(c, color);
}
void * LayerGetContext(void * layer) {
	return CGLayerGetContext(layer);
}
BOOL PDFArrayGetInteger(void * array, NSUInteger index, CGPDFInteger* value) {
	return CGPDFArrayGetInteger(array, index, value);
}
float FontGetItalicAngle(void * font) {
	return CGFontGetItalicAngle(font);
}
int32_t ConfigureDisplayOrigin(void * config, uint32_t display, int32_t x, int32_t y) {
	return CGConfigureDisplayOrigin(config, display, x, y);
}
void * ColorConversionInfoCreateWithOptions(void * src, void * dst, void * options) {
	return CGColorConversionInfoCreateWithOptions(src, dst, options);
}
NSInteger DisplayFadeOperationInProgress() {
	return CGDisplayFadeOperationInProgress();
}
void RestorePermanentDisplayConfiguration() {
	return CGRestorePermanentDisplayConfiguration();
}
uint32_t BitmapContextGetAlphaInfo(void * context) {
	return CGBitmapContextGetAlphaInfo(context);
}
void ContextClosePath(void * c) {
	return CGContextClosePath(c);
}
BOOL PathEqualToPath(void * path1, void * path2) {
	return CGPathEqualToPath(path1, path2);
}
void * EventCreateMouseEvent(void * source, uint32_t mouseType, CGPoint mouseCursorPosition, uint32_t mouseButton) {
	return CGEventCreateMouseEvent(source, mouseType, mouseCursorPosition, mouseButton);
}
uint32_t MainDisplayID() {
	return CGMainDisplayID();
}
CGAffineTransform AffineTransformTranslate(CGAffineTransform t, float tx, float ty) {
	return CGAffineTransformTranslate(t, tx, ty);
}
void * PDFDocumentGetPage(void * document, NSUInteger pageNumber) {
	return CGPDFDocumentGetPage(document, pageNumber);
}
float RectGetMidX(CGRect rect) {
	return CGRectGetMidX(rect);
}
uint32_t EventSourceGetKeyboardType(void * source) {
	return CGEventSourceGetKeyboardType(source);
}
void PDFContextBeginPage(void * context, void * pageInfo) {
	return CGPDFContextBeginPage(context, pageInfo);
}
void * ColorCreateCopyWithAlpha(void * color, float alpha) {
	return CGColorCreateCopyWithAlpha(color, alpha);
}
uint32_t PDFDocumentGetAccessPermissions(void * document) {
	return CGPDFDocumentGetAccessPermissions(document);
}
CGAffineTransform AffineTransformRotate(CGAffineTransform t, float angle) {
	return CGAffineTransformRotate(t, angle);
}
void * EventCreateCopy(void * event) {
	return CGEventCreateCopy(event);
}
void ContextSynchronize(void * c) {
	return CGContextSynchronize(c);
}
CGAffineTransform PDFPageGetDrawingTransform(void * page, int32_t box, CGRect rect, NSInteger rotate, BOOL preserveAspectRatio) {
	return CGPDFPageGetDrawingTransform(page, box, rect, rotate, preserveAspectRatio);
}
int32_t ShieldingWindowLevel() {
	return CGShieldingWindowLevel();
}
BOOL DisplayModeIsUsableForDesktopGUI(void * mode) {
	return CGDisplayModeIsUsableForDesktopGUI(mode);
}
int32_t GradientGetTypeID() {
	return CGGradientGetTypeID();
}
void * PointCreateDictionaryRepresentation(CGPoint point) {
	return CGPointCreateDictionaryRepresentation(point);
}
CGRect RectOffset(CGRect rect, float dx, float dy) {
	return CGRectOffset(rect, dx, dy);
}
BOOL PDFArrayGetName(void * array, NSUInteger index, uint8_t* value) {
	return CGPDFArrayGetName(array, index, value);
}
void ContextSetCMYKStrokeColor(void * c, float cyan, float magenta, float yellow, float black, float alpha) {
	return CGContextSetCMYKStrokeColor(c, cyan, magenta, yellow, black, alpha);
}
void * ColorSpaceGetBaseColorSpace(void * space) {
	return CGColorSpaceGetBaseColorSpace(space);
}
void * LayerCreateWithContext(void * context, CGSize size, void * auxiliaryInfo) {
	return CGLayerCreateWithContext(context, size, auxiliaryInfo);
}
NSUInteger DisplayModeGetPixelWidth(void * mode) {
	return CGDisplayModeGetPixelWidth(mode);
}
void * ImageCreateWithMaskingColors(void * image, float* components) {
	return CGImageCreateWithMaskingColors(image, components);
}
void ContextFillRects(void * c, CGRect* rects, NSUInteger count) {
	return CGContextFillRects(c, rects, count);
}
void * FontCopyPostScriptName(void * font) {
	return CGFontCopyPostScriptName(font);
}
BOOL ImageIsMask(void * image) {
	return CGImageIsMask(image);
}
void ContextSetStrokeColorSpace(void * c, void * space) {
	return CGContextSetStrokeColorSpace(c, space);
}
CGRect RectStandardize(CGRect rect) {
	return CGRectStandardize(rect);
}
CGRect FontGetFontBBox(void * font) {
	return CGFontGetFontBBox(font);
}
BOOL PDFDictionaryGetName(void * dict, uint8_t* key, uint8_t* value) {
	return CGPDFDictionaryGetName(dict, key, value);
}
BOOL PDFScannerPopObject(void * scanner, void * value) {
	return CGPDFScannerPopObject(scanner, value);
}
void * DisplayStreamGetRunLoopSource(void * displayStream) {
	return CGDisplayStreamGetRunLoopSource(displayStream);
}
int32_t DisplayCapture(uint32_t display) {
	return CGDisplayCapture(display);
}
void ContextAddEllipseInRect(void * c, CGRect rect) {
	return CGContextAddEllipseInRect(c, rect);
}
void * GradientCreateWithColorComponents(void * space, float* components, float* locations, NSUInteger count) {
	return CGGradientCreateWithColorComponents(space, components, locations, count);
}
BOOL RectEqualToRect(CGRect rect1, CGRect rect2) {
	return CGRectEqualToRect(rect1, rect2);
}
void * ColorSpaceCreateCalibratedRGB(float* whitePoint, float* blackPoint, float* gamma, float* matrix) {
	return CGColorSpaceCreateCalibratedRGB(whitePoint, blackPoint, gamma, matrix);
}
void ContextDrawShading(void * c, void * shading) {
	return CGContextDrawShading(c, shading);
}
float RectGetMidY(CGRect rect) {
	return CGRectGetMidY(rect);
}
void * DisplayCreateImageForRect(uint32_t display, CGRect rect) {
	return CGDisplayCreateImageForRect(display, rect);
}
void * ColorCreate(void * space, float* components) {
	return CGColorCreate(space, components);
}
void PathMoveToPoint(void * path, CGAffineTransform* m, float x, float y) {
	return CGPathMoveToPoint(path, m, x, y);
}
void ColorRelease(void * color) {
	return CGColorRelease(color);
}
NSUInteger PDFDocumentGetNumberOfPages(void * document) {
	return CGPDFDocumentGetNumberOfPages(document);
}
void EventSourceSetKeyboardType(void * source, uint32_t keyboardType) {
	return CGEventSourceSetKeyboardType(source, keyboardType);
}
void * DisplayCopyAllDisplayModes(uint32_t display, void * options) {
	return CGDisplayCopyAllDisplayModes(display, options);
}
NSUInteger DisplayPixelsHigh(uint32_t display) {
	return CGDisplayPixelsHigh(display);
}
void DisplayStreamUpdateGetMovedRectsDelta(void * updateRef, float* dx, float* dy) {
	return CGDisplayStreamUpdateGetMovedRectsDelta(updateRef, dx, dy);
}
int32_t SetDisplayTransferByByteTable(uint32_t display, uint32_t tableSize, uint8_t* redTable, uint8_t* greenTable, uint8_t* blueTable) {
	return CGSetDisplayTransferByByteTable(display, tableSize, redTable, greenTable, blueTable);
}
void * ImageGetDataProvider(void * image) {
	return CGImageGetDataProvider(image);
}
NSInteger DisplayIsCaptured(uint32_t display) {
	return CGDisplayIsCaptured(display);
}
void ContextSetFillPattern(void * c, void * pattern, float* components) {
	return CGContextSetFillPattern(c, pattern, components);
}
int32_t WaitForScreenRefreshRects(CGRect* rects, uint32_t* count) {
	return CGWaitForScreenRefreshRects(rects, count);
}
int32_t DataProviderGetTypeID() {
	return CGDataProviderGetTypeID();
}
void * DataProviderCreateWithFilename(uint8_t* filename) {
	return CGDataProviderCreateWithFilename(filename);
}
BOOL PDFScannerPopString(void * scanner, void * value) {
	return CGPDFScannerPopString(scanner, value);
}
NSUInteger ImageGetBytesPerRow(void * image) {
	return CGImageGetBytesPerRow(image);
}
void * DataProviderCreateWithCFData(void * data) {
	return CGDataProviderCreateWithCFData(data);
}
BOOL PDFDocumentAllowsCopying(void * document) {
	return CGPDFDocumentAllowsCopying(document);
}
void * ColorCreateCopyByMatchingToColorSpace(void * arg0, int32_t intent, void * color, void * options) {
	return CGColorCreateCopyByMatchingToColorSpace(arg0, intent, color, options);
}
int32_t InhibitLocalEvents(NSInteger inhibit) {
	return CGInhibitLocalEvents(inhibit);
}
CGSize SizeMake(float width, float height) {
	return CGSizeMake(width, height);
}
float RectGetMaxY(CGRect rect) {
	return CGRectGetMaxY(rect);
}
BOOL PDFDocumentAllowsPrinting(void * document) {
	return CGPDFDocumentAllowsPrinting(document);
}
void ContextSetAllowsFontSmoothing(void * c, BOOL allowsFontSmoothing) {
	return CGContextSetAllowsFontSmoothing(c, allowsFontSmoothing);
}
void ContextDrawPath(void * c, int32_t mode) {
	return CGContextDrawPath(c, mode);
}
NSInteger FontGetGlyphWithGlyphName(void * font, void * name) {
	return CGFontGetGlyphWithGlyphName(font, name);
}
BOOL PDFScannerPopArray(void * scanner, void * value) {
	return CGPDFScannerPopArray(scanner, value);
}
void * RectCreateDictionaryRepresentation(CGRect arg0) {
	return CGRectCreateDictionaryRepresentation(arg0);
}
void* BitmapContextGetData(void * context) {
	return CGBitmapContextGetData(context);
}
NSInteger FontGetAscent(void * font) {
	return CGFontGetAscent(font);
}
void * DisplayModeCopyPixelEncoding(void * mode) {
	return CGDisplayModeCopyPixelEncoding(mode);
}
int32_t DisplaySetStereoOperation(uint32_t display, NSInteger stereo, NSInteger forceBlueLine, uint32_t option) {
	return CGDisplaySetStereoOperation(display, stereo, forceBlueLine, option);
}
BOOL ImageGetShouldInterpolate(void * image) {
	return CGImageGetShouldInterpolate(image);
}
void ContextFillPath(void * c) {
	return CGContextFillPath(c);
}
int32_t PDFDocumentGetTypeID() {
	return CGPDFDocumentGetTypeID();
}
BOOL PDFArrayGetString(void * array, NSUInteger index, void * value) {
	return CGPDFArrayGetString(array, index, value);
}
void ContextShowText(void * c, uint8_t* string_, NSUInteger length) {
	return CGContextShowText(c, string_, length);
}
int32_t ImageGetRenderingIntent(void * image) {
	return CGImageGetRenderingIntent(image);
}
void EventPost(uint32_t tap, void * event) {
	return CGEventPost(tap, event);
}
float FontGetStemV(void * font) {
	return CGFontGetStemV(font);
}
NSInteger CursorIsVisible() {
	return CGCursorIsVisible();
}
void ContextStrokeLineSegments(void * c, CGPoint* points, NSUInteger count) {
	return CGContextStrokeLineSegments(c, points, count);
}
void ContextEOFillPath(void * c) {
	return CGContextEOFillPath(c);
}
BOOL PDFArrayGetObject(void * array, NSUInteger index, void * value) {
	return CGPDFArrayGetObject(array, index, value);
}
int32_t DisplayCaptureWithOptions(uint32_t display, uint32_t options) {
	return CGDisplayCaptureWithOptions(display, options);
}
BOOL ContextPathContainsPoint(void * c, CGPoint point, int32_t mode) {
	return CGContextPathContainsPoint(c, point, mode);
}
BOOL ColorSpaceUsesITUR_2100TF(void * arg0) {
	return CGColorSpaceUsesITUR_2100TF(arg0);
}
void * PDFDocumentGetOutline(void * document) {
	return CGPDFDocumentGetOutline(document);
}
int32_t DisplayStreamStop(void * displayStream) {
	return CGDisplayStreamStop(displayStream);
}
void PathRelease(void * path) {
	return CGPathRelease(path);
}
int32_t PathGetTypeID() {
	return CGPathGetTypeID();
}
void PDFDocumentRelease(void * document) {
	return CGPDFDocumentRelease(document);
}
void * WindowServerCreateServerPort() {
	return CGWindowServerCreateServerPort();
}
CGAffineTransform ContextGetUserSpaceToDeviceSpaceTransform(void * c) {
	return CGContextGetUserSpaceToDeviceSpaceTransform(c);
}
void ContextDrawImage(void * c, CGRect rect, void * image) {
	return CGContextDrawImage(c, rect, image);
}
void * PathCreateWithRoundedRect(CGRect rect, float cornerWidth, float cornerHeight, CGAffineTransform* transform) {
	return CGPathCreateWithRoundedRect(rect, cornerWidth, cornerHeight, transform);
}
void PDFContextBeginTag(void * context, int32_t tagType, void * tagProperties) {
	return CGPDFContextBeginTag(context, tagType, tagProperties);
}
int32_t PDFObjectGetType(void * object) {
	return CGPDFObjectGetType(object);
}
void * ImageMaskCreate(NSUInteger width, NSUInteger height, NSUInteger bitsPerComponent, NSUInteger bitsPerPixel, NSUInteger bytesPerRow, void * provider, float* decode, BOOL shouldInterpolate) {
	return CGImageMaskCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, provider, decode, shouldInterpolate);
}
void * ColorCreateGenericRGB(float red, float green, float blue, float alpha) {
	return CGColorCreateGenericRGB(red, green, blue, alpha);
}
double EventGetDoubleValueField(void * event, uint32_t field) {
	return CGEventGetDoubleValueField(event, field);
}
void * PDFDocumentGetInfo(void * document) {
	return CGPDFDocumentGetInfo(document);
}
BOOL PSConverterIsConverting(void * converter) {
	return CGPSConverterIsConverting(converter);
}
void * FontCreateCopyWithVariations(void * font, void * variations) {
	return CGFontCreateCopyWithVariations(font, variations);
}
int32_t ConfigureDisplayMirrorOfDisplay(void * config, uint32_t display, uint32_t master) {
	return CGConfigureDisplayMirrorOfDisplay(config, display, master);
}
double DisplayModeGetRefreshRate(void * mode) {
	return CGDisplayModeGetRefreshRate(mode);
}
double EventSourceGetLocalEventsSuppressionInterval(void * source) {
	return CGEventSourceGetLocalEventsSuppressionInterval(source);
}
void ContextAddArcToPoint(void * c, float x1, float y1, float x2, float y2, float radius) {
	return CGContextAddArcToPoint(c, x1, y1, x2, y2, radius);
}
void * DataConsumerCreateWithURL(void * url) {
	return CGDataConsumerCreateWithURL(url);
}
CGSize LayerGetSize(void * layer) {
	return CGLayerGetSize(layer);
}
void * ContextRetain(void * c) {
	return CGContextRetain(c);
}
int32_t AcquireDisplayFadeReservation(float seconds, CGDisplayFadeReservationToken* token) {
	return CGAcquireDisplayFadeReservation(seconds, token);
}
uint64_t EventGetFlags(void * event) {
	return CGEventGetFlags(event);
}
void * FontCopyGlyphNameForGlyph(void * font, NSInteger glyph) {
	return CGFontCopyGlyphNameForGlyph(font, glyph);
}
void ContextSetRGBFillColor(void * c, float red, float green, float blue, float alpha) {
	return CGContextSetRGBFillColor(c, red, green, blue, alpha);
}
void ContextRestoreGState(void * c) {
	return CGContextRestoreGState(c);
}
void * PathCreateCopyByTransformingPath(void * path, CGAffineTransform* transform) {
	return CGPathCreateCopyByTransformingPath(path, transform);
}
void * FontCopyFullName(void * font) {
	return CGFontCopyFullName(font);
}
CGAffineTransform AffineTransformConcat(CGAffineTransform t1, CGAffineTransform t2) {
	return CGAffineTransformConcat(t1, t2);
}
NSUInteger PDFPageGetPageNumber(void * page) {
	return CGPDFPageGetPageNumber(page);
}
void EventSetLocation(void * event, CGPoint location) {
	return CGEventSetLocation(event, location);
}
void FontRelease(void * font) {
	return CGFontRelease(font);
}
void * PDFStringCopyTextString(void * string_) {
	return CGPDFStringCopyTextString(string_);
}
BOOL PDFDocumentIsEncrypted(void * document) {
	return CGPDFDocumentIsEncrypted(document);
}
void ContextStrokeRect(void * c, CGRect rect) {
	return CGContextStrokeRect(c, rect);
}
void ContextSetShouldSubpixelPositionFonts(void * c, BOOL shouldSubpixelPositionFonts) {
	return CGContextSetShouldSubpixelPositionFonts(c, shouldSubpixelPositionFonts);
}
void * ContextCopyPath(void * c) {
	return CGContextCopyPath(c);
}
void * ImageGetUTType(void * image) {
	return CGImageGetUTType(image);
}
BOOL FontCanCreatePostScriptSubset(void * font, int32_t format) {
	return CGFontCanCreatePostScriptSubset(font, format);
}
float* ColorGetComponents(void * color) {
	return CGColorGetComponents(color);
}
int32_t ContextGetTypeID() {
	return CGContextGetTypeID();
}
void * DataConsumerCreateWithCFData(void * data) {
	return CGDataConsumerCreateWithCFData(data);
}
void ContextAddRect(void * c, CGRect rect) {
	return CGContextAddRect(c, rect);
}
void PathAddRelativeArc(void * path, CGAffineTransform* matrix, float x, float y, float radius, float startAngle, float delta) {
	return CGPathAddRelativeArc(path, matrix, x, y, radius, startAngle, delta);
}
void DataConsumerRelease(void * consumer) {
	return CGDataConsumerRelease(consumer);
}
int32_t ConfigureDisplayStereoOperation(void * config, uint32_t display, NSInteger stereo, NSInteger forceBlueLine) {
	return CGConfigureDisplayStereoOperation(config, display, stereo, forceBlueLine);
}
CGSize ContextConvertSizeToDeviceSpace(void * c, CGSize size) {
	return CGContextConvertSizeToDeviceSpace(c, size);
}
void PathAddCurveToPoint(void * path, CGAffineTransform* m, float cp1x, float cp1y, float cp2x, float cp2y, float x, float y) {
	return CGPathAddCurveToPoint(path, m, cp1x, cp1y, cp2x, cp2y, x, y);
}
NSInteger DisplayIsBuiltin(uint32_t display) {
	return CGDisplayIsBuiltin(display);
}
int32_t DisplayHideCursor(uint32_t display) {
	return CGDisplayHideCursor(display);
}
uint32_t DisplayUnitNumber(uint32_t display) {
	return CGDisplayUnitNumber(display);
}
int32_t DisplaySetDisplayMode(uint32_t display, void * mode, void * options) {
	return CGDisplaySetDisplayMode(display, mode, options);
}
void ContextShowGlyphsAtPositions(void * c, CGGlyph* glyphs, CGPoint* Lpositions, NSUInteger count) {
	return CGContextShowGlyphsAtPositions(c, glyphs, Lpositions, count);
}
void ContextSetFillColorWithColor(void * c, void * color) {
	return CGContextSetFillColorWithColor(c, color);
}
NSInteger FontGetLeading(void * font) {
	return CGFontGetLeading(font);
}
int32_t PSConverterGetTypeID() {
	return CGPSConverterGetTypeID();
}
void * GradientCreateWithColors(void * space, void * colors, float* locations) {
	return CGGradientCreateWithColors(space, colors, locations);
}
NSInteger DisplayIsAlwaysInMirrorSet(uint32_t display) {
	return CGDisplayIsAlwaysInMirrorSet(display);
}
void * PathCreateCopyByDashingPath(void * path, CGAffineTransform* transform, float phase, float* lengths, NSUInteger count) {
	return CGPathCreateCopyByDashingPath(path, transform, phase, lengths, count);
}
void ContextSetCMYKFillColor(void * c, float cyan, float magenta, float yellow, float black, float alpha) {
	return CGContextSetCMYKFillColor(c, cyan, magenta, yellow, black, alpha);
}
int32_t CancelDisplayConfiguration(void * config) {
	return CGCancelDisplayConfiguration(config);
}
void PathAddLines(void * path, CGAffineTransform* m, CGPoint* points, NSUInteger count) {
	return CGPathAddLines(path, m, points, count);
}
int32_t SetDisplayTransferByFormula(uint32_t display, float redMin, float redMax, float redGamma, float greenMin, float greenMax, float greenGamma, float blueMin, float blueMax, float blueGamma) {
	return CGSetDisplayTransferByFormula(display, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma);
}
NSInteger DisplayIsAsleep(uint32_t display) {
	return CGDisplayIsAsleep(display);
}
uint32_t ImageGetBitmapInfo(void * image) {
	return CGImageGetBitmapInfo(image);
}
void * WindowListCreateImageFromArray(CGRect screenBounds, void * windowArray, uint32_t imageOption) {
	return CGWindowListCreateImageFromArray(screenBounds, windowArray, imageOption);
}
NSUInteger BitmapContextGetWidth(void * context) {
	return CGBitmapContextGetWidth(context);
}
CGAffineTransform AffineTransformMake(float a, float b, float c, float d, float tx, float ty) {
	return CGAffineTransformMake(a, b, c, d, tx, ty);
}
void * PathCreateWithEllipseInRect(CGRect rect, CGAffineTransform* transform) {
	return CGPathCreateWithEllipseInRect(rect, transform);
}
BOOL RectIntersectsRect(CGRect rect1, CGRect rect2) {
	return CGRectIntersectsRect(rect1, rect2);
}
void * DisplayCreateImage(uint32_t displayID) {
	return CGDisplayCreateImage(displayID);
}
void ContextSetLineWidth(void * c, float width) {
	return CGContextSetLineWidth(c, width);
}
int32_t EnableEventStateCombining(NSInteger combineState) {
	return CGEnableEventStateCombining(combineState);
}
void * ColorSpaceCreateWithICCProfile(void * data) {
	return CGColorSpaceCreateWithICCProfile(data);
}
void PatternRelease(void * pattern) {
	return CGPatternRelease(pattern);
}
void ContextEOClip(void * c) {
	return CGContextEOClip(c);
}
void * ColorSpaceCreateExtended(void * space) {
	return CGColorSpaceCreateExtended(space);
}
void ContextSetAllowsFontSubpixelQuantization(void * c, BOOL allowsFontSubpixelQuantization) {
	return CGContextSetAllowsFontSubpixelQuantization(c, allowsFontSubpixelQuantization);
}
CGPoint ContextConvertPointToUserSpace(void * c, CGPoint point) {
	return CGContextConvertPointToUserSpace(c, point);
}
void * DataConsumerRetain(void * consumer) {
	return CGDataConsumerRetain(consumer);
}
void EventSourceSetLocalEventsFilterDuringSuppressionState(void * source, uint32_t filter, uint32_t state) {
	return CGEventSourceSetLocalEventsFilterDuringSuppressionState(source, filter, state);
}
void * DisplayCurrentMode(uint32_t display) {
	return CGDisplayCurrentMode(display);
}
int32_t WaitForScreenUpdateRects(uint32_t requestedOperations, CGScreenUpdateOperation* currentOperation, CGRect* rects, NSUInteger* rectCount, CGScreenUpdateMoveDelta* delta) {
	return CGWaitForScreenUpdateRects(requestedOperations, currentOperation, rects, rectCount, delta);
}
void ContextSetGrayStrokeColor(void * c, float gray, float alpha) {
	return CGContextSetGrayStrokeColor(c, gray, alpha);
}
uint32_t DisplayVendorNumber(uint32_t display) {
	return CGDisplayVendorNumber(display);
}
BOOL RectContainsPoint(CGRect rect, CGPoint point) {
	return CGRectContainsPoint(rect, point);
}
int32_t FontGetTypeID() {
	return CGFontGetTypeID();
}
void * DisplayGetDrawingContext(uint32_t display) {
	return CGDisplayGetDrawingContext(display);
}
BOOL PDFScannerPopName(void * scanner, uint8_t* value) {
	return CGPDFScannerPopName(scanner, value);
}
void ContextSetShouldSmoothFonts(void * c, BOOL shouldSmoothFonts) {
	return CGContextSetShouldSmoothFonts(c, shouldSmoothFonts);
}
void * EventCreate(void * source) {
	return CGEventCreate(source);
}
uint32_t DisplayMirrorsDisplay(uint32_t display) {
	return CGDisplayMirrorsDisplay(display);
}
BOOL AffineTransformIsIdentity(CGAffineTransform t) {
	return CGAffineTransformIsIdentity(t);
}
void ContextBeginPage(void * c, CGRect* mediaBox) {
	return CGContextBeginPage(c, mediaBox);
}
float RectGetWidth(CGRect rect) {
	return CGRectGetWidth(rect);
}
CGPoint PointMake(float x, float y) {
	return CGPointMake(x, y);
}
void * EventCreateScrollWheelEvent(void * source, uint32_t units, uint32_t wheelCount, int32_t wheel1) {
	return CGEventCreateScrollWheelEvent(source, units, wheelCount, wheel1);
}
void * ColorGetConstantColor(void * colorName) {
	return CGColorGetConstantColor(colorName);
}
int32_t SetLocalEventsSuppressionInterval(double seconds) {
	return CGSetLocalEventsSuppressionInterval(seconds);
}
void * EventCreateSourceFromEvent(void * event) {
	return CGEventCreateSourceFromEvent(event);
}
NSInteger FontGetXHeight(void * font) {
	return CGFontGetXHeight(font);
}
int32_t DisplayStreamStart(void * displayStream) {
	return CGDisplayStreamStart(displayStream);
}
CGPoint EventGetUnflippedLocation(void * event) {
	return CGEventGetUnflippedLocation(event);
}
void * ImageCreateWithMask(void * image, void * mask) {
	return CGImageCreateWithMask(image, mask);
}
void PDFContextSetURLForRect(void * context, void * url, CGRect rect) {
	return CGPDFContextSetURLForRect(context, url, rect);
}
void * ColorSpaceCreateLinearized(void * space) {
	return CGColorSpaceCreateLinearized(space);
}
NSUInteger ColorSpaceGetColorTableCount(void * space) {
	return CGColorSpaceGetColorTableCount(space);
}
void * DisplayBestModeForParametersAndRefreshRate(uint32_t display, NSUInteger bitsPerPixel, NSUInteger width, NSUInteger height, double refreshRate, boolean_t* exactMatch) {
	return CGDisplayBestModeForParametersAndRefreshRate(display, bitsPerPixel, width, height, refreshRate, exactMatch);
}
uint32_t OpenGLDisplayMaskToDisplayID(uint32_t mask) {
	return CGOpenGLDisplayMaskToDisplayID(mask);
}
NSUInteger DisplayStreamUpdateGetDropCount(void * updateRef) {
	return CGDisplayStreamUpdateGetDropCount(updateRef);
}
BOOL PointEqualToPoint(CGPoint point1, CGPoint point2) {
	return CGPointEqualToPoint(point1, point2);
}
void ContextFlush(void * c) {
	return CGContextFlush(c);
}
CGVector VectorMake(float dx, float dy) {
	return CGVectorMake(dx, dy);
}
NSUInteger DisplayModeGetHeight(void * mode) {
	return CGDisplayModeGetHeight(mode);
}
CGRect RectUnion(CGRect r1, CGRect r2) {
	return CGRectUnion(r1, r2);
}
BOOL PDFScannerScan(void * scanner) {
	return CGPDFScannerScan(scanner);
}
void ContextAddLines(void * c, CGPoint* points, NSUInteger count) {
	return CGContextAddLines(c, points, count);
}
void ContextSetMiterLimit(void * c, float limit) {
	return CGContextSetMiterLimit(c, limit);
}
void * PDFStreamGetDictionary(void * stream) {
	return CGPDFStreamGetDictionary(stream);
}
NSUInteger PDFStringGetLength(void * string_) {
	return CGPDFStringGetLength(string_);
}
void ContextClip(void * c) {
	return CGContextClip(c);
}
int32_t ReleaseAllDisplays() {
	return CGReleaseAllDisplays();
}
NSInteger CursorIsDrawnInFramebuffer() {
	return CGCursorIsDrawnInFramebuffer();
}
uint32_t DisplaySerialNumber(uint32_t display) {
	return CGDisplaySerialNumber(display);
}
BOOL PDFArrayGetArray(void * array, NSUInteger index, void * value) {
	return CGPDFArrayGetArray(array, index, value);
}
void * ColorSpaceCopyName(void * space) {
	return CGColorSpaceCopyName(space);
}
void ContextSetLineCap(void * c, int32_t cap) {
	return CGContextSetLineCap(c, cap);
}
void ShadingRelease(void * shading) {
	return CGShadingRelease(shading);
}
void * PDFOperatorTableRetain(void * table) {
	return CGPDFOperatorTableRetain(table);
}
uint32_t DisplayModelNumber(uint32_t display) {
	return CGDisplayModelNumber(display);
}
void EventSourceSetUserData(void * source, int64_t userData) {
	return CGEventSourceSetUserData(source, userData);
}
void GradientRelease(void * gradient) {
	return CGGradientRelease(gradient);
}
BOOL RequestListenEventAccess() {
	return CGRequestListenEventAccess();
}
void ContextMoveToPoint(void * c, float x, float y) {
	return CGContextMoveToPoint(c, x, y);
}
CGAffineTransform AffineTransformInvert(CGAffineTransform t) {
	return CGAffineTransformInvert(t);
}
void * PDFScannerRetain(void * scanner) {
	return CGPDFScannerRetain(scanner);
}
BOOL PathIsEmpty(void * path) {
	return CGPathIsEmpty(path);
}
void ContextDrawTiledImage(void * c, CGRect rect, void * image) {
	return CGContextDrawTiledImage(c, rect, image);
}
void * ImageRetain(void * image) {
	return CGImageRetain(image);
}
void * ImageCreateWithImageInRect(void * image, CGRect rect) {
	return CGImageCreateWithImageInRect(image, rect);
}
NSInteger PDFPageGetRotationAngle(void * page) {
	return CGPDFPageGetRotationAngle(page);
}
int32_t ConfigureDisplayMode(void * config, uint32_t display, void * mode) {
	return CGConfigureDisplayMode(config, display, mode);
}
void ContextAddLineToPoint(void * c, float x, float y) {
	return CGContextAddLineToPoint(c, x, y);
}
BOOL PDFScannerPopBoolean(void * scanner, CGPDFBoolean* value) {
	return CGPDFScannerPopBoolean(scanner, value);
}
void * FontCreateWithFontName(void * name) {
	return CGFontCreateWithFontName(name);
}
void ContextSetInterpolationQuality(void * c, int32_t quality) {
	return CGContextSetInterpolationQuality(c, quality);
}
void * PDFDocumentGetCatalog(void * document) {
	return CGPDFDocumentGetCatalog(document);
}
void LayerRelease(void * layer) {
	return CGLayerRelease(layer);
}
int32_t LayerGetTypeID() {
	return CGLayerGetTypeID();
}
int32_t FunctionGetTypeID() {
	return CGFunctionGetTypeID();
}
void * PDFContentStreamGetStreams(void * cs) {
	return CGPDFContentStreamGetStreams(cs);
}
NSUInteger PDFDictionaryGetCount(void * dict) {
	return CGPDFDictionaryGetCount(dict);
}
int32_t DisplayShowCursor(uint32_t display) {
	return CGDisplayShowCursor(display);
}
uint32_t BitmapContextGetBitmapInfo(void * context) {
	return CGBitmapContextGetBitmapInfo(context);
}
CGAffineTransform AffineTransformScale(CGAffineTransform t, float sx, float sy) {
	return CGAffineTransformScale(t, sx, sy);
}
void * FontCreatePostScriptSubset(void * font, void * subsetName, int32_t format, CGGlyph* glyphs, NSUInteger count, CGGlyph* encoding) {
	return CGFontCreatePostScriptSubset(font, subsetName, format, glyphs, count, encoding);
}
NSInteger DisplayIsMain(uint32_t display) {
	return CGDisplayIsMain(display);
}
void * PathCreateMutable() {
	return CGPathCreateMutable();
}
void ReleaseScreenRefreshRects(CGRect* rects) {
	return CGReleaseScreenRefreshRects(rects);
}
void ContextShowGlyphs(void * c, CGGlyph* g, NSUInteger count) {
	return CGContextShowGlyphs(c, g, count);
}
void DisplayRestoreColorSyncSettings() {
	return CGDisplayRestoreColorSyncSettings();
}
void ContextSetLineDash(void * c, float phase, float* lengths, NSUInteger count) {
	return CGContextSetLineDash(c, phase, lengths, count);
}
int32_t DisplayModeGetTypeID() {
	return CGDisplayModeGetTypeID();
}
uint64_t EventSourceFlagsState(int32_t stateID) {
	return CGEventSourceFlagsState(stateID);
}
BOOL RequestPostEventAccess() {
	return CGRequestPostEventAccess();
}
void ContextFillRect(void * c, CGRect rect) {
	return CGContextFillRect(c, rect);
}
void * PathRetain(void * path) {
	return CGPathRetain(path);
}
BOOL PDFDocumentIsUnlocked(void * document) {
	return CGPDFDocumentIsUnlocked(document);
}
