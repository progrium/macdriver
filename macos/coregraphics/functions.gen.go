// AUTO-GENERATED CODE, DO NOT MODIFY

package coregraphics

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import "CoreGraphics/CoreGraphics.h"
// void ContextStrokeEllipseInRect(void * c, CGRect rect);
// CGRect ContextConvertRectToDeviceSpace(void * c, CGRect rect);
// void EventSetSource(void * event, void * source);
// void ContextSetFontSize(void * c, float size);
// void ContextClipToRects(void * c, const CGRect* rects, uint count);
// CGRect RectApplyAffineTransform(CGRect rect, CGAffineTransform t);
// int64_t EventSourceGetUserData(void * source);
// CGImageAlphaInfo ImageGetAlphaInfo(void * image);
// boolean_t DisplayIsOnline(CGDirectDisplayID display);
// CGError ReleaseDisplayFadeReservation(CGDisplayFadeReservationToken token);
// bool ColorSpaceUsesExtendedRange(void * space);
// uint BitmapContextGetBitsPerComponent(void * context);
// void * PDFContextCreate(void * consumer, const CGRect* mediaBox, void * auxiliaryInfo);
// void * ColorSpaceGetName(void * space);
// void ContextBeginTransparencyLayer(void * c, void * auxiliaryInfo);
// void * ColorGetColorSpace(void * color);
// int64_t EventGetIntegerValueField(void * event, CGEventField field);
// void ContextSetTextPosition(void * c, float x, float y);
// void * ColorSpaceCopyICCData(void * space);
// void ContextSetFont(void * c, void * font);
// CGRect ContextConvertRectToUserSpace(void * c, CGRect rect);
// void * ColorRetain(void * color);
// void ContextBeginPath(void * c);
// bool ColorSpaceIsHLGBased(void * s);
// void * ColorCreateGenericGray(float gray, float alpha);
// void DataProviderRelease(void * provider);
// void PathAddArc(void * path, const CGAffineTransform* m, float x, float y, float radius, float startAngle, float endAngle, bool clockwise);
// CFTypeID EventGetTypeID();
// void ContextDrawLinearGradient(void * c, void * gradient, CGPoint startPoint, CGPoint endPoint, CGGradientDrawingOptions options);
// CGRect PDFPageGetBoxRect(void * page, CGPDFBox box);
// void ContextStrokeRectWithWidth(void * c, CGRect rect, float width);
// void * SessionCopyCurrentDictionary();
// CFTypeID ShadingGetTypeID();
// void * ColorSpaceCreatePattern(void * baseSpace);
// CGRect DisplayBounds(CGDirectDisplayID display);
// void ContextSetTextDrawingMode(void * c, CGTextDrawingMode mode);
// void * ColorSpaceCreateDeviceGray();
// void * ColorSpaceCreateCalibratedGray(const float* whitePoint, const float* blackPoint, float gamma);
// void * PathCreateCopyByStrokingPath(void * path, const CGAffineTransform* transform, float lineWidth, CGLineCap lineCap, CGLineJoin lineJoin, float miterLimit);
// bool ColorEqualToColor(void * color1, void * color2);
// bool PathContainsPoint(void * path, const CGAffineTransform* m, CGPoint point, bool eoFill);
// void * FontCopyVariations(void * font);
// void * FontRetain(void * font);
// void * DisplayCopyColorSpace(CGDirectDisplayID display);
// uint ImageGetWidth(void * image);
// void * PDFDocumentCreateWithURL(void * url);
// void EventSetFlags(void * event, CGEventFlags flags);
// CGRect RectMake(float x, float y, float width, float height);
// uint ColorGetNumberOfComponents(void * color);
// void * PathCreateWithRect(CGRect rect, const CGAffineTransform* transform);
// CGPoint ContextGetPathCurrentPoint(void * c);
// void PathCloseSubpath(void * path);
// void PathAddPath(void * path1, const CGAffineTransform* m, void * path2);
// bool RequestScreenCaptureAccess();
// bool SizeMakeWithDictionaryRepresentation(void * dict, CGSize* size);
// bool ColorSpaceIsWideGamutRGB(void * );
// void PDFOperatorTableRelease(void * table);
// int FontGetUnitsPerEm(void * font);
// void * PDFContentStreamRetain(void * cs);
// void ContextSetAlpha(void * c, float alpha);
// void ContextAddCurveToPoint(void * c, float cp1x, float cp1y, float cp2x, float cp2y, float x, float y);
// void * ColorCreateGenericCMYK(float cyan, float magenta, float yellow, float black, float alpha);
// void * ColorSpaceCreateExtendedLinearized(void * space);
// void * ColorSpaceCreateWithICCData(CFTypeRef data);
// CGError GetDisplaysWithRect(CGRect rect, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount);
// uint DisplayModeGetWidth(void * mode);
// void ContextSaveGState(void * c);
// int FontGetCapHeight(void * font);
// uint PDFArrayGetCount(void * array);
// void FunctionRelease(void * function);
// bool PDFDocumentUnlockWithPassword(void * document, char* password);
// double DisplayRotation(CGDirectDisplayID display);
// CGAffineTransform AffineTransformMakeScale(float sx, float sy);
// void ContextDrawLayerInRect(void * context, CGRect rect, void * layer);
// bool RectIsInfinite(CGRect rect);
// void * ImageCreateCopyWithColorSpace(void * image, void * space);
// bool PreflightPostEventAccess();
// void ContextRelease(void * c);
// CFTypeID DataConsumerGetTypeID();
// CGError CompleteDisplayConfiguration(void * config, CGConfigureOption option);
// void PDFContextEndTag(void * context);
// void * ColorCreateGenericGrayGamma2_2(float gray, float alpha);
// CGRect PathGetBoundingBox(void * path);
// void PDFDocumentGetVersion(void * document, int* majorVersion, int* minorVersion);
// void PathAddRoundedRect(void * path, const CGAffineTransform* transform, CGRect rect, float cornerWidth, float cornerHeight);
// CGWindowLevel WindowLevelForKey(CGWindowLevelKey key);
// void ContextSetRenderingIntent(void * c, CGColorRenderingIntent intent);
// void ContextSetShadowWithColor(void * c, CGSize offset, float blur, void * color);
// CFTypeID PatternGetTypeID();
// void * DisplayModeRetain(void * mode);
// CGEventTimestamp EventGetTimestamp(void * event);
// void * PathCreateMutableCopy(void * path);
// CGAffineTransform ContextGetCTM(void * c);
// void * EventCreateData(void * allocator, void * event);
// boolean_t DisplayIsStereo(CGDirectDisplayID display);
// void ContextConcatCTM(void * c, CGAffineTransform transform);
// CGSize ContextConvertSizeToUserSpace(void * c, CGSize size);
// bool PDFScannerPopInteger(void * scanner, CGPDFInteger* value);
// void PDFContextAddDocumentMetadata(void * context, void * metadata);
// CGSize DisplayScreenSize(CGDirectDisplayID display);
// void ColorSpaceGetColorTable(void * space, uint8_t* table);
// void * PDFPageRetain(void * page);
// void ContextFillEllipseInRect(void * c, CGRect rect);
// void * DisplayCopyDisplayMode(CGDirectDisplayID display);
// void PDFContextSetDestinationForRect(void * context, void * name, CGRect rect);
// CGError GetDisplaysWithOpenGLDisplayMask(CGOpenGLDisplayMask mask, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount);
// void * ShadingCreateAxial(void * space, CGPoint start, CGPoint end, void * function, bool extendStart, bool extendEnd);
// void ContextSetShadow(void * c, CGSize offset, float blur);
// bool PointMakeWithDictionaryRepresentation(void * dict, CGPoint* point);
// CGPoint ContextGetTextPosition(void * c);
// uint BitmapContextGetBitsPerPixel(void * context);
// void ContextAddArc(void * c, float x, float y, float radius, float startAngle, float endAngle, int clockwise);
// void EventSetType(void * event, CGEventType type_);
// void * EventCreateScrollWheelEvent2(void * source, CGScrollEventUnit units, uint32_t wheelCount, int32_t wheel1, int32_t wheel2, int32_t wheel3);
// CGColorSpaceModel ColorSpaceGetModel(void * space);
// void * ImageCreateWithPNGDataProvider(void * source, const float* decode, bool shouldInterpolate, CGColorRenderingIntent intent);
// void PathAddEllipseInRect(void * path, const CGAffineTransform* m, CGRect rect);
// void * EventCreateFromData(void * allocator, void * data);
// void * ColorSpaceCreateDeviceCMYK();
// CGError ConfigureDisplayFadeEffect(void * config, CGDisplayFadeInterval fadeOutSeconds, CGDisplayFadeInterval fadeInSeconds, float fadeRed, float fadeGreen, float fadeBlue);
// CGError DisplayFade(CGDisplayFadeReservationToken token, CGDisplayFadeInterval duration, CGDisplayBlendFraction startBlend, CGDisplayBlendFraction endBlend, float redBlend, float greenBlend, float blueBlend, boolean_t synchronous);
// void ContextSetShouldSubpixelQuantizeFonts(void * c, bool shouldSubpixelQuantizeFonts);
// CGPoint EventGetLocation(void * event);
// bool PSConverterAbort(void * converter);
// void ContextScaleCTM(void * c, float sx, float sy);
// void * SizeCreateDictionaryRepresentation(CGSize size);
// bool EventSourceButtonState(CGEventSourceStateID stateID, CGMouseButton button);
// void EventTapEnable(void * tap, bool enable);
// CGAffineTransform AffineTransformMakeRotation(float angle);
// void * ColorSpaceCreateICCBased(uint nComponents, const float* range_, void * profile, void * alternate);
// CGError AssociateMouseAndMouseCursorPosition(boolean_t connected);
// void * PDFDocumentRetain(void * document);
// CGPoint PathGetCurrentPoint(void * path);
// void ContextSetRGBStrokeColor(void * c, float red, float green, float blue, float alpha);
// bool PreflightScreenCaptureAccess();
// CGError SetDisplayTransferByTable(CGDirectDisplayID display, uint32_t tableSize, const CGGammaValue* redTable, const CGGammaValue* greenTable, const CGGammaValue* blueTable);
// void ContextSetFlatness(void * c, float flatness);
// boolean_t DisplayIsActive(CGDirectDisplayID display);
// void ContextSetPatternPhase(void * c, CGSize phase);
// void PDFScannerRelease(void * scanner);
// void ContextSetLineJoin(void * c, CGLineJoin join);
// void * FontCopyTableTags(void * font);
// CGError GetActiveDisplayList(uint32_t maxDisplays, CGDirectDisplayID* activeDisplays, uint32_t* displayCount);
// bool PathIsRect(void * path, CGRect* rect);
// void * DataProviderCreateWithURL(void * url);
// void * EventSourceCreate(CGEventSourceStateID stateID);
// CGInterpolationQuality ContextGetInterpolationQuality(void * c);
// void * ColorConversionInfoCreate(void * src, void * dst);
// bool RectContainsRect(CGRect rect1, CGRect rect2);
// void * FontCreatePostScriptEncoding(void * font, const CGGlyph* encoding);
// void ContextDrawLayerAtPoint(void * context, CGPoint point, void * layer);
// void EventSourceSetPixelsPerLine(void * source, double pixelsPerLine);
// void * EventCreateKeyboardEvent(void * source, CGKeyCode virtualKey, bool keyDown);
// uint32_t EventSourceCounterForEventType(CGEventSourceStateID stateID, CGEventType eventType);
// void PathAddRects(void * path, const CGAffineTransform* m, const CGRect* rects, uint count);
// void EventSetTimestamp(void * event, CGEventTimestamp timestamp);
// bool EventSourceKeyState(CGEventSourceStateID stateID, CGKeyCode key);
// float RectGetHeight(CGRect rect);
// bool PDFDictionaryGetInteger(void * dict, char* key, CGPDFInteger* value);
// CFTimeInterval EventSourceSecondsSinceLastEventType(CGEventSourceStateID stateID, CGEventType eventType);
// CGRect RectIntegral(CGRect rect);
// void ContextSetFillColorSpace(void * c, void * space);
// bool PDFArrayGetNumber(void * array, uint index, CGPDFReal* value);
// void * ColorCreateWithPattern(void * space, void * pattern, const float* components);
// void ContextTranslateCTM(void * c, float tx, float ty);
// boolean_t DisplayUsesOpenGLAcceleration(CGDirectDisplayID display);
// void ContextSetBlendMode(void * c, CGBlendMode mode);
// void GetLastMouseDelta(int32_t* deltaX, int32_t* deltaY);
// bool FontGetGlyphBBoxes(void * font, const CGGlyph* glyphs, uint count, CGRect* bboxes);
// CGError GetDisplaysWithPoint(CGPoint point, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount);
// void ImageRelease(void * image);
// float RectGetMinX(CGRect rect);
// void ContextReplacePathWithStrokedPath(void * c);
// uint BitmapContextGetBytesPerRow(void * context);
// void * WindowListCreate(CGWindowListOption option, CGWindowID relativeToWindow);
// CGSize SizeApplyAffineTransform(CGSize size, CGAffineTransform t);
// void EventTapPostEvent(CGEventTapProxy proxy, void * event);
// boolean_t DisplayIsInMirrorSet(CGDirectDisplayID display);
// bool ColorSpaceIsPQBased(void * s);
// void ContextClearRect(void * c, CGRect rect);
// CGError DisplayRelease(CGDirectDisplayID display);
// void * PDFPageGetDocument(void * page);
// CGError GetDisplayTransferByTable(CGDirectDisplayID display, uint32_t capacity, CGGammaValue* redTable, CGGammaValue* greenTable, CGGammaValue* blueTable, uint32_t* sampleCount);
// bool AffineTransformEqualToTransform(CGAffineTransform t1, CGAffineTransform t2);
// void * PDFScannerGetContentStream(void * scanner);
// void PathAddLineToPoint(void * path, const CGAffineTransform* m, float x, float y);
// bool PreflightListenEventAccess();
// CGEventType EventGetType(void * event);
// bool PDFScannerPopNumber(void * scanner, CGPDFReal* value);
// void * ColorSpaceCreateLab(const float* whitePoint, const float* blackPoint, const float* range_);
// void PDFContextEndPage(void * context);
// CGRect ContextGetClipBoundingBox(void * c);
// void * PDFOperatorTableCreate();
// void PathAddRect(void * path, const CGAffineTransform* m, CGRect rect);
// CGPoint PointApplyAffineTransform(CGPoint point, CGAffineTransform t);
// void * ImageCreateWithJPEGDataProvider(void * source, const float* decode, bool shouldInterpolate, CGColorRenderingIntent intent);
// CFTypeID EventSourceGetTypeID();
// void * ImageCreate(uint width, uint height, uint bitsPerComponent, uint bitsPerPixel, uint bytesPerRow, void * space, CGBitmapInfo bitmapInfo, void * provider, const float* decode, bool shouldInterpolate, CGColorRenderingIntent intent);
// bool RectIsNull(CGRect rect);
// CGEventSourceStateID EventSourceGetSourceStateID(void * source);
// void * GradientRetain(void * gradient);
// bool PDFArrayGetBoolean(void * array, uint index, CGPDFBoolean* value);
// void ContextSetFillColor(void * c, const float* components);
// void * ColorGetPattern(void * color);
// CGRect PathGetPathBoundingBox(void * path);
// CGError CaptureAllDisplaysWithOptions(CGCaptureOptions options);
// void * WindowListCreateDescriptionFromArray(void * windowArray);
// void * PDFContentStreamCreateWithStream(void * stream, void * streamResources, void * parent);
// void ContextDrawRadialGradient(void * c, void * gradient, CGPoint startCenter, float startRadius, CGPoint endCenter, float endRadius, CGGradientDrawingOptions options);
// void * PDFStringCopyDate(void * string_);
// CGError GetOnlineDisplayList(uint32_t maxDisplays, CGDirectDisplayID* onlineDisplays, uint32_t* displayCount);
// CGPoint ContextConvertPointToDeviceSpace(void * c, CGPoint point);
// void * PathCreateCopy(void * path);
// void * BitmapContextCreateImage(void * context);
// uint ColorSpaceGetNumberOfComponents(void * space);
// void * FunctionRetain(void * function);
// CGError GetDisplayTransferByFormula(CGDirectDisplayID display, CGGammaValue* redMin, CGGammaValue* redMax, CGGammaValue* redGamma, CGGammaValue* greenMin, CGGammaValue* greenMax, CGGammaValue* greenGamma, CGGammaValue* blueMin, CGGammaValue* blueMax, CGGammaValue* blueGamma);
// void * ColorSpaceRetain(void * space);
// void * ColorCreateSRGB(float red, float green, float blue, float alpha);
// void * ImageGetColorSpace(void * image);
// void * ShadingCreateRadial(void * space, CGPoint start, float startRadius, CGPoint end, float endRadius, void * function, bool extendStart, bool extendEnd);
// void RectDivide(CGRect rect, CGRect* slice, CGRect* remainder, float amount, CGRectEdge edge);
// void ContextSetAllowsAntialiasing(void * c, bool allowsAntialiasing);
// uint BitmapContextGetHeight(void * context);
// float ColorGetAlpha(void * color);
// uint ImageGetHeight(void * image);
// void * ColorSpaceCreateWithName(void * name);
// void ContextAddQuadCurveToPoint(void * c, float cpx, float cpy, float x, float y);
// CGError ConfigureDisplayWithDisplayMode(void * config, CGDirectDisplayID display, void * mode, void * options);
// void * DataProviderCopyData(void * provider);
// boolean_t DisplayIsInHWMirrorSet(CGDirectDisplayID display);
// CGEventFilterMask EventSourceGetLocalEventsFilterDuringSuppressionState(void * source, CGEventSuppressionState state);
// void * FontCopyTableForTag(void * font, uint32_t tag);
// char* PDFStringGetBytePtr(void * string_);
// CGRect ContextGetPathBoundingBox(void * c);
// void * LayerRetain(void * layer);
// void ContextSetAllowsFontSubpixelPositioning(void * c, bool allowsFontSubpixelPositioning);
// CFTypeID ColorGetTypeID();
// CGImageByteOrderInfo ImageGetByteOrderInfo(void * image);
// void * PDFDocumentCreateWithProvider(void * provider);
// bool FontGetGlyphAdvances(void * font, const CGGlyph* glyphs, uint count, int* advances);
// CGImagePixelFormatInfo ImageGetPixelFormatInfo(void * image);
// void * FontCopyVariationAxes(void * font);
// CFPropertyListRef ColorSpaceCopyPropertyList(void * space);
// void ContextSetStrokePattern(void * c, void * pattern, const float* components);
// void ContextClipToRect(void * c, CGRect rect);
// void ContextResetClip(void * c);
// void PathAddQuadCurveToPoint(void * path, const CGAffineTransform* m, float cpx, float cpy, float x, float y);
// CGError WarpMouseCursorPosition(CGPoint newCursorPosition);
// CGError DisplayMoveCursorToPoint(CGDirectDisplayID display, CGPoint point);
// void * WindowListCopyWindowInfo(CGWindowListOption option, CGWindowID relativeToWindow);
// CGWindowID ShieldingWindowID(CGDirectDisplayID display);
// void PDFContextSetOutline(void * context, void * outline);
// void PathAddArcToPoint(void * path, const CGAffineTransform* m, float x1, float y1, float x2, float y2, float radius);
// uint ImageGetBitsPerComponent(void * image);
// CGError CaptureAllDisplays();
// double EventSourceGetPixelsPerLine(void * source);
// CGRect RectInset(CGRect rect, float dx, float dy);
// bool PDFArrayGetNull(void * array, uint index);
// const float* ImageGetDecode(void * image);
// void * DataProviderGetInfo(void * provider);
// void * PDFStreamCopyData(void * stream, CGPDFDataFormat* format);
// void * PatternRetain(void * pattern);
// void * BitmapContextGetColorSpace(void * context);
// bool RectIsEmpty(CGRect rect);
// float RectGetMinY(CGRect rect);
// CFTypeID ColorSpaceGetTypeID();
// void * ShadingRetain(void * shading);
// void * ColorCreateCopy(void * color);
// void ContextBeginTransparencyLayerWithRect(void * c, CGRect rect, void * auxInfo);
// void PDFPageRelease(void * page);
// CFTypeID ImageGetTypeID();
// void ContextRotateCTM(void * c, float angle);
// void ContextEndTransparencyLayer(void * c);
// void EventSourceSetLocalEventsSuppressionInterval(void * source, CFTimeInterval seconds);
// CGRect RectIntersection(CGRect r1, CGRect r2);
// int32_t DisplayModeGetIODisplayModeID(void * mode);
// CGDirectDisplayID DisplayPrimaryDisplay(CGDirectDisplayID display);
// void PDFContextAddDestinationAtPoint(void * context, void * name, CGPoint point);
// int FontGetDescent(void * font);
// void ContextSetStrokeColor(void * c, const float* components);
// void DisplayModeRelease(void * mode);
// void ContextAddRects(void * c, const CGRect* rects, uint count);
// void * PathCreateMutableCopyByTransformingPath(void * path, const CGAffineTransform* transform);
// CGAffineTransform AffineTransformMakeTranslation(float tx, float ty);
// void ContextSetCharacterSpacing(void * c, float spacing);
// uint32_t DisplayModeGetIOFlags(void * mode);
// void PDFContextClose(void * context);
// void * DataProviderRetain(void * provider);
// bool ColorSpaceSupportsOutput(void * space);
// void EventSetDoubleValueField(void * event, CGEventField field, double value);
// void ContextDrawPDFPage(void * c, void * page);
// void * PDFContextCreateWithURL(void * url, const CGRect* mediaBox, void * auxiliaryInfo);
// CFTypeID PDFPageGetTypeID();
// void * PDFContentStreamCreateWithPage(void * page);
// void * FontCreateWithDataProvider(void * provider);
// bool PDFDictionaryGetNumber(void * dict, char* key, CGPDFReal* value);
// void ColorSpaceRelease(void * space);
// void * ImageCreateCopy(void * image);
// void ContextEndPage(void * c);
// void * PDFDocumentGetID(void * document);
// void ContextSetTextMatrix(void * c, CGAffineTransform t);
// uint FontGetNumberOfGlyphs(void * font);
// void ContextClipToMask(void * c, CGRect rect, void * mask);
// void * PDFPageGetDictionary(void * page);
// uint32_t DisplayGammaTableCapacity(CGDirectDisplayID display);
// void EventSetIntegerValueField(void * event, CGEventField field, int64_t value);
// void ContextAddPath(void * c, void * path);
// CGAffineTransform ContextGetTextMatrix(void * c);
// void ContextStrokePath(void * c);
// uint DisplayPixelsWide(CGDirectDisplayID display);
// void ContextSetShouldAntialias(void * c, bool shouldAntialias);
// bool EventTapIsEnabled(void * tap);
// char* PDFTagTypeGetName(CGPDFTagType tagType);
// void ContextSetGrayFillColor(void * c, float gray, float alpha);
// void * PDFContentStreamGetResource(void * cs, char* category, char* name);
// bool PDFDictionaryGetBoolean(void * dict, char* key, CGPDFBoolean* value);
// bool PSConverterConvert(void * converter, void * provider, void * consumer, void * options);
// void PDFContentStreamRelease(void * cs);
// CGOpenGLDisplayMask DisplayIDToOpenGLDisplayMask(CGDirectDisplayID display);
// float RectGetMaxX(CGRect rect);
// bool ContextIsPathEmpty(void * c);
// CFTypeID ColorConversionInfoGetTypeID();
// uint ImageGetBitsPerPixel(void * image);
// bool RectMakeWithDictionaryRepresentation(void * dict, CGRect* rect);
// uint DisplayModeGetPixelHeight(void * mode);
// void * ColorSpaceCreateDeviceRGB();
// void ContextSetStrokeColorWithColor(void * c, void * color);
// void * LayerGetContext(void * layer);
// bool PDFArrayGetInteger(void * array, uint index, CGPDFInteger* value);
// float FontGetItalicAngle(void * font);
// CGError ConfigureDisplayOrigin(void * config, CGDirectDisplayID display, int32_t x, int32_t y);
// void * ColorConversionInfoCreateWithOptions(void * src, void * dst, void * options);
// void RestorePermanentDisplayConfiguration();
// CGImageAlphaInfo BitmapContextGetAlphaInfo(void * context);
// void ContextClosePath(void * c);
// bool PathEqualToPath(void * path1, void * path2);
// void * EventCreateMouseEvent(void * source, CGEventType mouseType, CGPoint mouseCursorPosition, CGMouseButton mouseButton);
// CGDirectDisplayID MainDisplayID();
// CGAffineTransform AffineTransformTranslate(CGAffineTransform t, float tx, float ty);
// void * PDFDocumentGetPage(void * document, uint pageNumber);
// float RectGetMidX(CGRect rect);
// CGEventSourceKeyboardType EventSourceGetKeyboardType(void * source);
// void PDFContextBeginPage(void * context, void * pageInfo);
// void * ColorCreateCopyWithAlpha(void * color, float alpha);
// CGPDFAccessPermissions PDFDocumentGetAccessPermissions(void * document);
// CGAffineTransform AffineTransformRotate(CGAffineTransform t, float angle);
// void * EventCreateCopy(void * event);
// void ContextSynchronize(void * c);
// CGAffineTransform PDFPageGetDrawingTransform(void * page, CGPDFBox box, CGRect rect, int rotate, bool preserveAspectRatio);
// CGWindowLevel ShieldingWindowLevel();
// bool DisplayModeIsUsableForDesktopGUI(void * mode);
// CFTypeID GradientGetTypeID();
// void * PointCreateDictionaryRepresentation(CGPoint point);
// CGRect RectOffset(CGRect rect, float dx, float dy);
// void ContextSetCMYKStrokeColor(void * c, float cyan, float magenta, float yellow, float black, float alpha);
// void * ColorSpaceGetBaseColorSpace(void * space);
// void * LayerCreateWithContext(void * context, CGSize size, void * auxiliaryInfo);
// uint DisplayModeGetPixelWidth(void * mode);
// void * ImageCreateWithMaskingColors(void * image, const float* components);
// void ContextFillRects(void * c, const CGRect* rects, uint count);
// void * FontCopyPostScriptName(void * font);
// bool ImageIsMask(void * image);
// void ContextSetStrokeColorSpace(void * c, void * space);
// CGRect RectStandardize(CGRect rect);
// CGRect FontGetFontBBox(void * font);
// CGError DisplayCapture(CGDirectDisplayID display);
// void ContextAddEllipseInRect(void * c, CGRect rect);
// void * GradientCreateWithColorComponents(void * space, const float* components, const float* locations, uint count);
// void * ColorSpaceCreateCalibratedRGB(const float* whitePoint, const float* blackPoint, const float* gamma, const float* matrix);
// void ContextDrawShading(void * c, void * shading);
// float RectGetMidY(CGRect rect);
// void * DisplayCreateImageForRect(CGDirectDisplayID display, CGRect rect);
// void * ColorCreate(void * space, const float* components);
// void PathMoveToPoint(void * path, const CGAffineTransform* m, float x, float y);
// void ColorRelease(void * color);
// uint PDFDocumentGetNumberOfPages(void * document);
// void EventSourceSetKeyboardType(void * source, CGEventSourceKeyboardType keyboardType);
// void * DisplayCopyAllDisplayModes(CGDirectDisplayID display, void * options);
// uint DisplayPixelsHigh(CGDirectDisplayID display);
// CGError SetDisplayTransferByByteTable(CGDirectDisplayID display, uint32_t tableSize, const uint8_t* redTable, const uint8_t* greenTable, const uint8_t* blueTable);
// void * ImageGetDataProvider(void * image);
// void ContextSetFillPattern(void * c, void * pattern, const float* components);
// CFTypeID DataProviderGetTypeID();
// void * DataProviderCreateWithFilename(char* filename);
// uint ImageGetBytesPerRow(void * image);
// void * DataProviderCreateWithCFData(void * data);
// bool PDFDocumentAllowsCopying(void * document);
// void * ColorCreateCopyByMatchingToColorSpace(void * , CGColorRenderingIntent intent, void * color, void * options);
// CGSize SizeMake(float width, float height);
// float RectGetMaxY(CGRect rect);
// bool PDFDocumentAllowsPrinting(void * document);
// void ContextSetAllowsFontSmoothing(void * c, bool allowsFontSmoothing);
// void ContextDrawPath(void * c, CGPathDrawingMode mode);
// CGGlyph FontGetGlyphWithGlyphName(void * font, void * name);
// void * RectCreateDictionaryRepresentation(CGRect );
// void * BitmapContextGetData(void * context);
// int FontGetAscent(void * font);
// CGError DisplaySetStereoOperation(CGDirectDisplayID display, boolean_t stereo, boolean_t forceBlueLine, CGConfigureOption option);
// bool ImageGetShouldInterpolate(void * image);
// void ContextFillPath(void * c);
// CFTypeID PDFDocumentGetTypeID();
// CGColorRenderingIntent ImageGetRenderingIntent(void * image);
// void EventPost(CGEventTapLocation tap, void * event);
// float FontGetStemV(void * font);
// void ContextStrokeLineSegments(void * c, const CGPoint* points, uint count);
// void ContextEOFillPath(void * c);
// CGError DisplayCaptureWithOptions(CGDirectDisplayID display, CGCaptureOptions options);
// bool ContextPathContainsPoint(void * c, CGPoint point, CGPathDrawingMode mode);
// bool ColorSpaceUsesITUR_2100TF(void * );
// void * PDFDocumentGetOutline(void * document);
// void PathRelease(void * path);
// CFTypeID PathGetTypeID();
// void PDFDocumentRelease(void * document);
// void * WindowServerCreateServerPort();
// CGAffineTransform ContextGetUserSpaceToDeviceSpaceTransform(void * c);
// void ContextDrawImage(void * c, CGRect rect, void * image);
// void * PathCreateWithRoundedRect(CGRect rect, float cornerWidth, float cornerHeight, const CGAffineTransform* transform);
// void PDFContextBeginTag(void * context, CGPDFTagType tagType, void * tagProperties);
// CGPDFObjectType PDFObjectGetType(void * object);
// void * ImageMaskCreate(uint width, uint height, uint bitsPerComponent, uint bitsPerPixel, uint bytesPerRow, void * provider, const float* decode, bool shouldInterpolate);
// void * ColorCreateGenericRGB(float red, float green, float blue, float alpha);
// double EventGetDoubleValueField(void * event, CGEventField field);
// void * PDFDocumentGetInfo(void * document);
// bool PSConverterIsConverting(void * converter);
// void * FontCreateCopyWithVariations(void * font, void * variations);
// CGError ConfigureDisplayMirrorOfDisplay(void * config, CGDirectDisplayID display, CGDirectDisplayID master);
// double DisplayModeGetRefreshRate(void * mode);
// CFTimeInterval EventSourceGetLocalEventsSuppressionInterval(void * source);
// void ContextAddArcToPoint(void * c, float x1, float y1, float x2, float y2, float radius);
// void * DataConsumerCreateWithURL(void * url);
// CGSize LayerGetSize(void * layer);
// void * ContextRetain(void * c);
// CGError AcquireDisplayFadeReservation(CGDisplayReservationInterval seconds, CGDisplayFadeReservationToken* token);
// CGEventFlags EventGetFlags(void * event);
// void * FontCopyGlyphNameForGlyph(void * font, CGGlyph glyph);
// void ContextSetRGBFillColor(void * c, float red, float green, float blue, float alpha);
// void ContextRestoreGState(void * c);
// void * PathCreateCopyByTransformingPath(void * path, const CGAffineTransform* transform);
// void * FontCopyFullName(void * font);
// CGAffineTransform AffineTransformConcat(CGAffineTransform t1, CGAffineTransform t2);
// uint PDFPageGetPageNumber(void * page);
// void EventSetLocation(void * event, CGPoint location);
// void FontRelease(void * font);
// void * PDFStringCopyTextString(void * string_);
// bool PDFDocumentIsEncrypted(void * document);
// void ContextStrokeRect(void * c, CGRect rect);
// void ContextSetShouldSubpixelPositionFonts(void * c, bool shouldSubpixelPositionFonts);
// void * ContextCopyPath(void * c);
// void * ImageGetUTType(void * image);
// bool FontCanCreatePostScriptSubset(void * font, CGFontPostScriptFormat format);
// const float* ColorGetComponents(void * color);
// CFTypeID ContextGetTypeID();
// void * DataConsumerCreateWithCFData(void * data);
// void ContextAddRect(void * c, CGRect rect);
// void PathAddRelativeArc(void * path, const CGAffineTransform* matrix, float x, float y, float radius, float startAngle, float delta);
// void DataConsumerRelease(void * consumer);
// CGError ConfigureDisplayStereoOperation(void * config, CGDirectDisplayID display, boolean_t stereo, boolean_t forceBlueLine);
// CGSize ContextConvertSizeToDeviceSpace(void * c, CGSize size);
// void PathAddCurveToPoint(void * path, const CGAffineTransform* m, float cp1x, float cp1y, float cp2x, float cp2y, float x, float y);
// boolean_t DisplayIsBuiltin(CGDirectDisplayID display);
// CGError DisplayHideCursor(CGDirectDisplayID display);
// uint32_t DisplayUnitNumber(CGDirectDisplayID display);
// CGError DisplaySetDisplayMode(CGDirectDisplayID display, void * mode, void * options);
// void ContextShowGlyphsAtPositions(void * c, const CGGlyph* glyphs, const CGPoint* Lpositions, uint count);
// void ContextSetFillColorWithColor(void * c, void * color);
// int FontGetLeading(void * font);
// CFTypeID PSConverterGetTypeID();
// void * GradientCreateWithColors(void * space, void * colors, const float* locations);
// boolean_t DisplayIsAlwaysInMirrorSet(CGDirectDisplayID display);
// void * PathCreateCopyByDashingPath(void * path, const CGAffineTransform* transform, float phase, const float* lengths, uint count);
// void ContextSetCMYKFillColor(void * c, float cyan, float magenta, float yellow, float black, float alpha);
// CGError CancelDisplayConfiguration(void * config);
// void PathAddLines(void * path, const CGAffineTransform* m, const CGPoint* points, uint count);
// CGError SetDisplayTransferByFormula(CGDirectDisplayID display, CGGammaValue redMin, CGGammaValue redMax, CGGammaValue redGamma, CGGammaValue greenMin, CGGammaValue greenMax, CGGammaValue greenGamma, CGGammaValue blueMin, CGGammaValue blueMax, CGGammaValue blueGamma);
// boolean_t DisplayIsAsleep(CGDirectDisplayID display);
// CGBitmapInfo ImageGetBitmapInfo(void * image);
// uint BitmapContextGetWidth(void * context);
// CGAffineTransform AffineTransformMake(float a, float b, float c, float d, float tx, float ty);
// void * PathCreateWithEllipseInRect(CGRect rect, const CGAffineTransform* transform);
// bool RectIntersectsRect(CGRect rect1, CGRect rect2);
// void * DisplayCreateImage(CGDirectDisplayID displayID);
// void ContextSetLineWidth(void * c, float width);
// void PatternRelease(void * pattern);
// void ContextEOClip(void * c);
// void * ColorSpaceCreateExtended(void * space);
// void ContextSetAllowsFontSubpixelQuantization(void * c, bool allowsFontSubpixelQuantization);
// CGPoint ContextConvertPointToUserSpace(void * c, CGPoint point);
// void * DataConsumerRetain(void * consumer);
// void EventSourceSetLocalEventsFilterDuringSuppressionState(void * source, CGEventFilterMask filter, CGEventSuppressionState state);
// void ContextSetGrayStrokeColor(void * c, float gray, float alpha);
// uint32_t DisplayVendorNumber(CGDirectDisplayID display);
// bool RectContainsPoint(CGRect rect, CGPoint point);
// CFTypeID FontGetTypeID();
// void * DisplayGetDrawingContext(CGDirectDisplayID display);
// void ContextSetShouldSmoothFonts(void * c, bool shouldSmoothFonts);
// void * EventCreate(void * source);
// CGDirectDisplayID DisplayMirrorsDisplay(CGDirectDisplayID display);
// bool AffineTransformIsIdentity(CGAffineTransform t);
// void ContextBeginPage(void * c, const CGRect* mediaBox);
// float RectGetWidth(CGRect rect);
// CGPoint PointMake(float x, float y);
// void * EventCreateScrollWheelEvent(void * source, CGScrollEventUnit units, uint32_t wheelCount, int32_t wheel1);
// void * ColorGetConstantColor(void * colorName);
// void * EventCreateSourceFromEvent(void * event);
// int FontGetXHeight(void * font);
// CGPoint EventGetUnflippedLocation(void * event);
// void * ImageCreateWithMask(void * image, void * mask);
// void PDFContextSetURLForRect(void * context, void * url, CGRect rect);
// void * ColorSpaceCreateLinearized(void * space);
// uint ColorSpaceGetColorTableCount(void * space);
// CGDirectDisplayID OpenGLDisplayMaskToDisplayID(CGOpenGLDisplayMask mask);
// void ContextFlush(void * c);
// CGVector VectorMake(float dx, float dy);
// uint DisplayModeGetHeight(void * mode);
// CGRect RectUnion(CGRect r1, CGRect r2);
// bool PDFScannerScan(void * scanner);
// void ContextAddLines(void * c, const CGPoint* points, uint count);
// void ContextSetMiterLimit(void * c, float limit);
// void * PDFStreamGetDictionary(void * stream);
// uint PDFStringGetLength(void * string_);
// void ContextClip(void * c);
// CGError ReleaseAllDisplays();
// uint32_t DisplaySerialNumber(CGDirectDisplayID display);
// void * ColorSpaceCopyName(void * space);
// void ContextSetLineCap(void * c, CGLineCap cap);
// void ShadingRelease(void * shading);
// void * PDFOperatorTableRetain(void * table);
// uint32_t DisplayModelNumber(CGDirectDisplayID display);
// void EventSourceSetUserData(void * source, int64_t userData);
// void GradientRelease(void * gradient);
// bool RequestListenEventAccess();
// void ContextMoveToPoint(void * c, float x, float y);
// CGAffineTransform AffineTransformInvert(CGAffineTransform t);
// void * PDFScannerRetain(void * scanner);
// bool PathIsEmpty(void * path);
// void ContextDrawTiledImage(void * c, CGRect rect, void * image);
// void * ImageRetain(void * image);
// void * ImageCreateWithImageInRect(void * image, CGRect rect);
// int PDFPageGetRotationAngle(void * page);
// void ContextAddLineToPoint(void * c, float x, float y);
// bool PDFScannerPopBoolean(void * scanner, CGPDFBoolean* value);
// void * FontCreateWithFontName(void * name);
// void ContextSetInterpolationQuality(void * c, CGInterpolationQuality quality);
// void * PDFDocumentGetCatalog(void * document);
// void LayerRelease(void * layer);
// CFTypeID LayerGetTypeID();
// CFTypeID FunctionGetTypeID();
// void * PDFContentStreamGetStreams(void * cs);
// uint PDFDictionaryGetCount(void * dict);
// CGError DisplayShowCursor(CGDirectDisplayID display);
// CGBitmapInfo BitmapContextGetBitmapInfo(void * context);
// CGAffineTransform AffineTransformScale(CGAffineTransform t, float sx, float sy);
// void * FontCreatePostScriptSubset(void * font, void * subsetName, CGFontPostScriptFormat format, const CGGlyph* glyphs, uint count, const CGGlyph* encoding);
// boolean_t DisplayIsMain(CGDirectDisplayID display);
// void * PathCreateMutable();
// void DisplayRestoreColorSyncSettings();
// void ContextSetLineDash(void * c, float phase, const float* lengths, uint count);
// CFTypeID DisplayModeGetTypeID();
// CGEventFlags EventSourceFlagsState(CGEventSourceStateID stateID);
// bool RequestPostEventAccess();
// void ContextFillRect(void * c, CGRect rect);
// void * PathRetain(void * path);
// bool PDFDocumentIsUnlocked(void * document);
import "C"
import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
)

// Strokes an ellipse that fits inside the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455774-cgcontextstrokeellipseinrect?language=objc
func ContextStrokeEllipseInRect(c ContextRef, rect Rect) {
	C.ContextStrokeEllipseInRect(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

// Returns a rectangle that is transformed from user space coordinate to device space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456017-cgcontextconvertrecttodevicespac?language=objc
func ContextConvertRectToDeviceSpace(c ContextRef, rect Rect) Rect {
	rv := C.ContextConvertRectToDeviceSpace(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Sets the event source of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455500-cgeventsetsource?language=objc
func EventSetSource(event EventRef, source EventSourceRef) {
	C.EventSetSource(
		// *typing.RefType
		unsafe.Pointer(event),
		// *typing.RefType
		unsafe.Pointer(source),
	)
}

// Sets the current font size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456426-cgcontextsetfontsize?language=objc
func ContextSetFontSize(c ContextRef, size float64) {
	C.ContextSetFontSize(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(size),
	)
}

// Sets the clipping path to the intersection of the current clipping path with the region defined by an array of rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454626-cgcontextcliptorects?language=objc
func ContextClipToRects(c ContextRef, rects *Rect, count uint) {
	C.ContextClipToRects(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&rects)),
		// *typing.PrimitiveType
		C.uint(count),
	)
}

// Applies an affine transform to a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455875-cgrectapplyaffinetransform?language=objc
func RectApplyAffineTransform(rect Rect, t AffineTransform) Rect {
	rv := C.RectApplyAffineTransform(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t)),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Returns the 64-bit user-specified data for a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408777-cgeventsourcegetuserdata?language=objc
func EventSourceGetUserData(source EventSourceRef) int64 {
	rv := C.EventSourceGetUserData(
		// *typing.RefType
		unsafe.Pointer(source),
	)
	// *typing.PrimitiveType
	return int64(rv)
}

// Returns the alpha channel information for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455401-cgimagegetalphainfo?language=objc
func ImageGetAlphaInfo(image ImageRef) ImageAlphaInfo {
	rv := C.ImageGetAlphaInfo(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.AliasType
	return ImageAlphaInfo(rv)
}

// Returns a Boolean value indicating whether a display is connected or online. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454476-cgdisplayisonline?language=objc
func DisplayIsOnline(display DirectDisplayID) int {
	rv := C.DisplayIsOnline(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return int(rv)
}

// Releases a display fade reservation, and unfades the display if needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455230-cgreleasedisplayfadereservation?language=objc
func ReleaseDisplayFadeReservation(token DisplayFadeReservationToken) Error {
	rv := C.ReleaseDisplayFadeReservation(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDisplayFadeReservationToken)(token),
	)
	// *typing.AliasType
	return Error(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3579522-cgcolorspaceusesextendedrange?language=objc
func ColorSpaceUsesExtendedRange(space ColorSpaceRef) bool {
	rv := C.ColorSpaceUsesExtendedRange(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the bits per component of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455383-cgbitmapcontextgetbitspercompone?language=objc
func BitmapContextGetBitsPerComponent(context ContextRef) uint {
	rv := C.BitmapContextGetBitsPerComponent(
		// *typing.RefType
		unsafe.Pointer(context),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Creates a PDF graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454204-cgpdfcontextcreate?language=objc
func PDFContextCreate(consumer DataConsumerRef, mediaBox *Rect, auxiliaryInfo corefoundation.DictionaryRef) ContextRef {
	rv := C.PDFContextCreate(
		// *typing.RefType
		unsafe.Pointer(consumer),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&mediaBox)),
		// *typing.RefType
		unsafe.Pointer(auxiliaryInfo),
	)
	// *typing.RefType
	return ContextRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2923324-cgcolorspacegetname?language=objc
func ColorSpaceGetName(space ColorSpaceRef) corefoundation.StringRef {
	rv := C.ColorSpaceGetName(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.RefType
	return corefoundation.StringRef(rv)
}

// Begins a transparency layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456011-cgcontextbegintransparencylayer?language=objc
func ContextBeginTransparencyLayer(c ContextRef, auxiliaryInfo corefoundation.DictionaryRef) {
	C.ContextBeginTransparencyLayer(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(auxiliaryInfo),
	)
}

// Returns the color space associated with a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455744-cgcolorgetcolorspace?language=objc
func ColorGetColorSpace(color ColorRef) ColorSpaceRef {
	rv := C.ColorGetColorSpace(
		// *typing.RefType
		unsafe.Pointer(color),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Returns the integer value of a field in a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455885-cgeventgetintegervaluefield?language=objc
func EventGetIntegerValueField(event EventRef, field EventField) int64 {
	rv := C.EventGetIntegerValueField(
		// *typing.RefType
		unsafe.Pointer(event),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventField)(field),
	)
	// *typing.PrimitiveType
	return int64(rv)
}

// Sets the location at which text is drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456069-cgcontextsettextposition?language=objc
func ContextSetTextPosition(c ContextRef, x float64, y float64) {
	C.ContextSetTextPosition(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
}

// Returns a copy of the ICC profile data of the provided color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1644732-cgcolorspacecopyiccdata?language=objc
func ColorSpaceCopyICCData(space ColorSpaceRef) corefoundation.DataRef {
	rv := C.ColorSpaceCopyICCData(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.RefType
	return corefoundation.DataRef(rv)
}

// Sets the platform font in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454950-cgcontextsetfont?language=objc
func ContextSetFont(c ContextRef, font FontRef) {
	C.ContextSetFont(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(font),
	)
}

// Returns a rectangle that is transformed from device space coordinate to user space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454165-cgcontextconvertrecttouserspace?language=objc
func ContextConvertRectToUserSpace(c ContextRef, rect Rect) Rect {
	rv := C.ContextConvertRectToUserSpace(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Increments the retain count of a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586339-cgcolorretain?language=objc
func ColorRetain(color ColorRef) ColorRef {
	rv := C.ColorRetain(
		// *typing.RefType
		unsafe.Pointer(color),
	)
	// *typing.RefType
	return ColorRef(rv)
}

// Creates a new empty path in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456635-cgcontextbeginpath?language=objc
func ContextBeginPath(c ContextRef) {
	C.ContextBeginPath(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3861799-cgcolorspaceishlgbased?language=objc
func ColorSpaceIsHLGBased(s ColorSpaceRef) bool {
	rv := C.ColorSpaceIsHLGBased(
		// *typing.RefType
		unsafe.Pointer(s),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a color in the Generic gray color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456453-cgcolorcreategenericgray?language=objc
func ColorCreateGenericGray(gray float64, alpha float64) ColorRef {
	rv := C.ColorCreateGenericGray(
		// *typing.PrimitiveType
		C.float(gray),
		// *typing.PrimitiveType
		C.float(alpha),
	)
	// *typing.RefType
	return ColorRef(rv)
}

// Decrements the retain count of a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408304-cgdataproviderrelease?language=objc
func DataProviderRelease(provider DataProviderRef) {
	C.DataProviderRelease(
		// *typing.RefType
		unsafe.Pointer(provider),
	)
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411147-cgpathaddarc?language=objc
func PathAddArc(path MutablePathRef, m *AffineTransform, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise bool) {
	C.PathAddArc(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
		// *typing.PrimitiveType
		C.float(radius),
		// *typing.PrimitiveType
		C.float(startAngle),
		// *typing.PrimitiveType
		C.float(endAngle),
		// *typing.PrimitiveType
		C.bool(clockwise),
	)
}

// Returns the type identifier for the opaque type CGEventRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454922-cgeventgettypeid?language=objc
func EventGetTypeID() corefoundation.TypeID {
	rv := C.EventGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Paints a gradient fill that varies along the line defined by the provided starting and ending points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454782-cgcontextdrawlineargradient?language=objc
func ContextDrawLinearGradient(c ContextRef, gradient GradientRef, startPoint Point, endPoint Point, options GradientDrawingOptions) {
	C.ContextDrawLinearGradient(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(gradient),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&startPoint)),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&endPoint)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGGradientDrawingOptions)(options),
	)
}

// Returns the rectangle that represents a type of box for a content region or page dimensions of a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456114-cgpdfpagegetboxrect?language=objc
func PDFPageGetBoxRect(page PDFPageRef, box PDFBox) Rect {
	rv := C.PDFPageGetBoxRect(
		// *typing.RefType
		unsafe.Pointer(page),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGPDFBox)(box),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Paints a rectangular path, using the specified line width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454679-cgcontextstrokerectwithwidth?language=objc
func ContextStrokeRectWithWidth(c ContextRef, rect Rect, width float64) {
	C.ContextStrokeRectWithWidth(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.PrimitiveType
		C.float(width),
	)
}

// Returns information about the caller’s window server session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454780-cgsessioncopycurrentdictionary?language=objc
func SessionCopyCurrentDictionary() corefoundation.DictionaryRef {
	rv := C.SessionCopyCurrentDictionary()
	// *typing.RefType
	return corefoundation.DictionaryRef(rv)
}

// Returns the Core Foundation type identifier for Core Graphics shading objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454302-cgshadinggettypeid?language=objc
func ShadingGetTypeID() corefoundation.TypeID {
	rv := C.ShadingGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Creates a pattern color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408869-cgcolorspacecreatepattern?language=objc
func ColorSpaceCreatePattern(baseSpace ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceCreatePattern(
		// *typing.RefType
		unsafe.Pointer(baseSpace),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Returns the bounds of a display in the global display coordinate space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456395-cgdisplaybounds?language=objc
func DisplayBounds(display DirectDisplayID) Rect {
	rv := C.DisplayBounds(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Sets the current text drawing mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454253-cgcontextsettextdrawingmode?language=objc
func ContextSetTextDrawingMode(c ContextRef, mode TextDrawingMode) {
	C.ContextSetTextDrawingMode(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGTextDrawingMode)(mode),
	)
}

// Creates a device-dependent grayscale color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408908-cgcolorspacecreatedevicegray?language=objc
func ColorSpaceCreateDeviceGray() ColorSpaceRef {
	rv := C.ColorSpaceCreateDeviceGray()
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Creates a calibrated grayscale color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408887-cgcolorspacecreatecalibratedgray?language=objc
func ColorSpaceCreateCalibratedGray(whitePoint *float64, blackPoint *float64, gamma float64) ColorSpaceRef {
	rv := C.ColorSpaceCreateCalibratedGray(
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&whitePoint)),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&blackPoint)),
		// *typing.PrimitiveType
		C.float(gamma),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Creates a stroked copy of another path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411128-cgpathcreatecopybystrokingpath?language=objc
func PathCreateCopyByStrokingPath(path unsafe.Pointer, transform *AffineTransform, lineWidth float64, lineCap LineCap, lineJoin LineJoin, miterLimit float64) unsafe.Pointer {
	rv := C.PathCreateCopyByStrokingPath(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&transform)),
		// *typing.PrimitiveType
		C.float(lineWidth),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGLineCap)(lineCap),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGLineJoin)(lineJoin),
		// *typing.PrimitiveType
		C.float(miterLimit),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Indicates whether two colors are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455217-cgcolorequaltocolor?language=objc
func ColorEqualToColor(color1 ColorRef, color2 ColorRef) bool {
	rv := C.ColorEqualToColor(
		// *typing.RefType
		unsafe.Pointer(color1),
		// *typing.RefType
		unsafe.Pointer(color2),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Checks whether a point is contained in a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411175-cgpathcontainspoint?language=objc
func PathContainsPoint(path unsafe.Pointer, m *AffineTransform, point Point, eoFill bool) bool {
	rv := C.PathContainsPoint(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&point)),
		// *typing.PrimitiveType
		C.bool(eoFill),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the variation specification dictionary for a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396355-cgfontcopyvariations?language=objc
func FontCopyVariations(font FontRef) corefoundation.DictionaryRef {
	rv := C.FontCopyVariations(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.RefType
	return corefoundation.DictionaryRef(rv)
}

// Increments the retain count of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396384-cgfontretain?language=objc
func FontRetain(font FontRef) FontRef {
	rv := C.FontRetain(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.RefType
	return FontRef(rv)
}

// Returns the color space for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454190-cgdisplaycopycolorspace?language=objc
func DisplayCopyColorSpace(display DirectDisplayID) ColorSpaceRef {
	rv := C.DisplayCopyColorSpace(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Returns the width of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456148-cgimagegetwidth?language=objc
func ImageGetWidth(image ImageRef) uint {
	rv := C.ImageGetWidth(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Creates a Core Graphics PDF document using data specified by a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402585-cgpdfdocumentcreatewithurl?language=objc
func PDFDocumentCreateWithURL(url corefoundation.URLRef) PDFDocumentRef {
	rv := C.PDFDocumentCreateWithURL(
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.RefType
	return PDFDocumentRef(rv)
}

// Sets the event flags of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455044-cgeventsetflags?language=objc
func EventSetFlags(event EventRef, flags EventFlags) {
	C.EventSetFlags(
		// *typing.RefType
		unsafe.Pointer(event),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventFlags)(flags),
	)
}

// Returns a rectangle with the specified coordinate and size values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455245-cgrectmake?language=objc
func RectMake(x float64, y float64, width float64, height float64) Rect {
	rv := C.RectMake(
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
		// *typing.PrimitiveType
		C.float(width),
		// *typing.PrimitiveType
		C.float(height),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Returns the number of color components (including alpha) associated with a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454130-cgcolorgetnumberofcomponents?language=objc
func ColorGetNumberOfComponents(color ColorRef) uint {
	rv := C.ColorGetNumberOfComponents(
		// *typing.RefType
		unsafe.Pointer(color),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Create an immutable path of a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411155-cgpathcreatewithrect?language=objc
func PathCreateWithRect(rect Rect, transform *AffineTransform) unsafe.Pointer {
	rv := C.PathCreateWithRect(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&transform)),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns the current point in a non-empty path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454788-cgcontextgetpathcurrentpoint?language=objc
func ContextGetPathCurrentPoint(c ContextRef) Point {
	rv := C.ContextGetPathCurrentPoint(
		// *typing.RefType
		unsafe.Pointer(c),
	)
	// *typing.StructType
	return *(*Point)(unsafe.Pointer(&rv))
}

// Closes and completes a subpath in a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411188-cgpathclosesubpath?language=objc
func PathCloseSubpath(path MutablePathRef) {
	C.PathCloseSubpath(
		// *typing.RefType
		unsafe.Pointer(path),
	)
}

// Appends a path to onto a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411201-cgpathaddpath?language=objc
func PathAddPath(path1 MutablePathRef, m *AffineTransform, path2 unsafe.Pointer) {
	C.PathAddPath(
		// *typing.RefType
		unsafe.Pointer(path1),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.RefType
		unsafe.Pointer(path2),
	)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656524-cgrequestscreencaptureaccess?language=objc
func RequestScreenCaptureAccess() bool {
	rv := C.RequestScreenCaptureAccess()
	// *typing.PrimitiveType
	return bool(rv)
}

// Fills in a size using the contents of the specified dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454318-cgsizemakewithdictionaryrepresen?language=objc
func SizeMakeWithDictionaryRepresentation(dict corefoundation.DictionaryRef, size *Size) bool {
	rv := C.SizeMakeWithDictionaryRepresentation(
		// *typing.RefType
		unsafe.Pointer(dict),
		// *typing.PointerType
		(*C.CGSize)(unsafe.Pointer(&size)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns whether the RGB color space covers a significant portion of the NTSC color gamut. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1644737-cgcolorspaceiswidegamutrgb?language=objc
func ColorSpaceIsWideGamutRGB(arg0 ColorSpaceRef) bool {
	rv := C.ColorSpaceIsWideGamutRGB(
		// *typing.RefType
		unsafe.Pointer(arg0),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Decrements the retain count of a CGPDFOperatorTable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455277-cgpdfoperatortablerelease?language=objc
func PDFOperatorTableRelease(table unsafe.Pointer) {
	C.PDFOperatorTableRelease(
		// *typing.RefType
		unsafe.Pointer(table),
	)
}

// Returns the number of glyph space units per em for the provided font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396344-cgfontgetunitsperem?language=objc
func FontGetUnitsPerEm(font FontRef) int {
	rv := C.FontGetUnitsPerEm(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.PrimitiveType
	return int(rv)
}

// Increments the retain count of a PDF content stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410049-cgpdfcontentstreamretain?language=objc
func PDFContentStreamRetain(cs unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFContentStreamRetain(
		// *typing.RefType
		unsafe.Pointer(cs),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Sets the opacity level for objects drawn in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456404-cgcontextsetalpha?language=objc
func ContextSetAlpha(c ContextRef, alpha float64) {
	C.ContextSetAlpha(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(alpha),
	)
}

// Appends a cubic Bézier curve from the current point, using the provided control points and end point . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456393-cgcontextaddcurvetopoint?language=objc
func ContextAddCurveToPoint(c ContextRef, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {
	C.ContextAddCurveToPoint(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(cp1x),
		// *typing.PrimitiveType
		C.float(cp1y),
		// *typing.PrimitiveType
		C.float(cp2x),
		// *typing.PrimitiveType
		C.float(cp2y),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
}

// Creates a color in the Generic CMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454222-cgcolorcreategenericcmyk?language=objc
func ColorCreateGenericCMYK(cyan float64, magenta float64, yellow float64, black float64, alpha float64) ColorRef {
	rv := C.ColorCreateGenericCMYK(
		// *typing.PrimitiveType
		C.float(cyan),
		// *typing.PrimitiveType
		C.float(magenta),
		// *typing.PrimitiveType
		C.float(yellow),
		// *typing.PrimitiveType
		C.float(black),
		// *typing.PrimitiveType
		C.float(alpha),
	)
	// *typing.RefType
	return ColorRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3684558-cgcolorspacecreateextendedlinear?language=objc
func ColorSpaceCreateExtendedLinearized(space ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateExtendedLinearized(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Creates an ICC-based color space using the ICC profile contained in the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2866135-cgcolorspacecreatewithiccdata?language=objc
func ColorSpaceCreateWithICCData(data corefoundation.TypeRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateWithICCData(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(data),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Gets a list of online displays with bounds that intersect the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456071-cggetdisplayswithrect?language=objc
func GetDisplaysWithRect(rect Rect, maxDisplays uint32, displays *DirectDisplayID, matchingDisplayCount *uint32) Error {
	rv := C.GetDisplaysWithRect(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.PrimitiveType
		C.uint32_t(maxDisplays),
		// *typing.PointerType
		(*C.CGDirectDisplayID)(unsafe.Pointer(&displays)),
		// *typing.PointerType
		(*C.uint32_t)(unsafe.Pointer(&matchingDisplayCount)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Returns the width of the specified display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454442-cgdisplaymodegetwidth?language=objc
func DisplayModeGetWidth(mode DisplayModeRef) uint {
	rv := C.DisplayModeGetWidth(
		// *typing.RefType
		unsafe.Pointer(mode),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Pushes a copy of the current graphics state onto the graphics state stack for the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456156-cgcontextsavegstate?language=objc
func ContextSaveGState(c ContextRef) {
	C.ContextSaveGState(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Returns the cap height of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396338-cgfontgetcapheight?language=objc
func FontGetCapHeight(font FontRef) int {
	rv := C.FontGetCapHeight(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.PrimitiveType
	return int(rv)
}

// Returns the number of items in a PDF array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455207-cgpdfarraygetcount?language=objc
func PDFArrayGetCount(array unsafe.Pointer) uint {
	rv := C.PDFArrayGetCount(
		// *typing.RefType
		unsafe.Pointer(array),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Decrements the retain count of a function object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1390864-cgfunctionrelease?language=objc
func FunctionRelease(function FunctionRef) {
	C.FunctionRelease(
		// *typing.RefType
		unsafe.Pointer(function),
	)
}

// Unlocks an encrypted PDF document when a valid password is supplied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402599-cgpdfdocumentunlockwithpassword?language=objc
func PDFDocumentUnlockWithPassword(document PDFDocumentRef, password string) bool {
	passwordVal := C.CString(password)
	defer C.free(unsafe.Pointer(passwordVal))
	rv := C.PDFDocumentUnlockWithPassword(
		// *typing.RefType
		unsafe.Pointer(document),
		// *typing.CStringType
		passwordVal,
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the rotation angle of a display in degrees. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455299-cgdisplayrotation?language=objc
func DisplayRotation(display DirectDisplayID) float64 {
	rv := C.DisplayRotation(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Returns an affine transformation matrix constructed from scaling values you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455016-cgaffinetransformmakescale?language=objc
func AffineTransformMakeScale(sx float64, sy float64) AffineTransform {
	rv := C.AffineTransformMakeScale(
		// *typing.PrimitiveType
		C.float(sx),
		// *typing.PrimitiveType
		C.float(sy),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Draws the contents of a layer object into the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450896-cgcontextdrawlayerinrect?language=objc
func ContextDrawLayerInRect(context ContextRef, rect Rect, layer LayerRef) {
	C.ContextDrawLayerInRect(
		// *typing.RefType
		unsafe.Pointer(context),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.RefType
		unsafe.Pointer(layer),
	)
}

// Returns whether a rectangle is infinite. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455008-cgrectisinfinite?language=objc
func RectIsInfinite(rect Rect) bool {
	rv := C.RectIsInfinite(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a copy of a bitmap image, replacing its colorspace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455355-cgimagecreatecopywithcolorspace?language=objc
func ImageCreateCopyWithColorSpace(image ImageRef, space ColorSpaceRef) ImageRef {
	rv := C.ImageCreateCopyWithColorSpace(
		// *typing.RefType
		unsafe.Pointer(image),
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.RefType
	return ImageRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656520-cgpreflightposteventaccess?language=objc
func PreflightPostEventAccess() bool {
	rv := C.PreflightPostEventAccess()
	// *typing.PrimitiveType
	return bool(rv)
}

// Decrements the retain count of a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586509-cgcontextrelease?language=objc
func ContextRelease(c ContextRef) {
	C.ContextRelease(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Returns the Core Foundation type identifier for Core Graphics data consumers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455226-cgdataconsumergettypeid?language=objc
func DataConsumerGetTypeID() corefoundation.TypeID {
	rv := C.DataConsumerGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Completes a set of display configuration changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454488-cgcompletedisplayconfiguration?language=objc
func CompleteDisplayConfiguration(config unsafe.Pointer, option ConfigureOption) Error {
	rv := C.CompleteDisplayConfiguration(
		// *typing.RefType
		unsafe.Pointer(config),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGConfigureOption)(option),
	)
	// *typing.AliasType
	return Error(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3042357-cgpdfcontextendtag?language=objc
func PDFContextEndTag(context ContextRef) {
	C.PDFContextEndTag(
		// *typing.RefType
		unsafe.Pointer(context),
	)
}

// Creates a color in the Generic gray color space with a gamma ramp of 2.2. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3042354-cgcolorcreategenericgraygamma2_2?language=objc
func ColorCreateGenericGrayGamma2_2(gray float64, alpha float64) ColorRef {
	rv := C.ColorCreateGenericGrayGamma2_2(
		// *typing.PrimitiveType
		C.float(gray),
		// *typing.PrimitiveType
		C.float(alpha),
	)
	// *typing.RefType
	return ColorRef(rv)
}

// Returns the bounding box containing all points in a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411165-cgpathgetboundingbox?language=objc
func PathGetBoundingBox(path unsafe.Pointer) Rect {
	rv := C.PathGetBoundingBox(
		// *typing.RefType
		unsafe.Pointer(path),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Returns the major and minor version numbers of a Core Graphics PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402604-cgpdfdocumentgetversion?language=objc
func PDFDocumentGetVersion(document PDFDocumentRef, majorVersion *int, minorVersion *int) {
	C.PDFDocumentGetVersion(
		// *typing.RefType
		unsafe.Pointer(document),
		// *typing.PointerType
		(*C.int)(unsafe.Pointer(&majorVersion)),
		// *typing.PointerType
		(*C.int)(unsafe.Pointer(&minorVersion)),
	)
}

// Appends a rounded rectangle to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411124-cgpathaddroundedrect?language=objc
func PathAddRoundedRect(path MutablePathRef, transform *AffineTransform, rect Rect, cornerWidth float64, cornerHeight float64) {
	C.PathAddRoundedRect(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&transform)),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.PrimitiveType
		C.float(cornerWidth),
		// *typing.PrimitiveType
		C.float(cornerHeight),
	)
}

// Returns the window level that corresponds to one of the standard window types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454084-cgwindowlevelforkey?language=objc
func WindowLevelForKey(key WindowLevelKey) WindowLevel {
	rv := C.WindowLevelForKey(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGWindowLevelKey)(key),
	)
	// *typing.AliasType
	return WindowLevel(rv)
}

// Sets the rendering intent in the current graphics state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455544-cgcontextsetrenderingintent?language=objc
func ContextSetRenderingIntent(c ContextRef, intent ColorRenderingIntent) {
	C.ContextSetRenderingIntent(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGColorRenderingIntent)(intent),
	)
}

// Enables shadowing with color a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455205-cgcontextsetshadowwithcolor?language=objc
func ContextSetShadowWithColor(c ContextRef, offset Size, blur float64, color ColorRef) {
	C.ContextSetShadowWithColor(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGSize)(unsafe.Pointer(&offset)),
		// *typing.PrimitiveType
		C.float(blur),
		// *typing.RefType
		unsafe.Pointer(color),
	)
}

// Returns the type identifier for Core Graphics patterns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456445-cgpatterngettypeid?language=objc
func PatternGetTypeID() corefoundation.TypeID {
	rv := C.PatternGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Retains a Core Graphics display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1562059-cgdisplaymoderetain?language=objc
func DisplayModeRetain(mode DisplayModeRef) DisplayModeRef {
	rv := C.DisplayModeRetain(
		// *typing.RefType
		unsafe.Pointer(mode),
	)
	// *typing.RefType
	return DisplayModeRef(rv)
}

// Returns the timestamp of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455481-cgeventgettimestamp?language=objc
func EventGetTimestamp(event EventRef) EventTimestamp {
	rv := C.EventGetTimestamp(
		// *typing.RefType
		unsafe.Pointer(event),
	)
	// *typing.AliasType
	return EventTimestamp(rv)
}

// Creates a mutable copy of an existing graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411196-cgpathcreatemutablecopy?language=objc
func PathCreateMutableCopy(path unsafe.Pointer) MutablePathRef {
	rv := C.PathCreateMutableCopy(
		// *typing.RefType
		unsafe.Pointer(path),
	)
	// *typing.RefType
	return MutablePathRef(rv)
}

// Returns the current transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454691-cgcontextgetctm?language=objc
func ContextGetCTM(c ContextRef) AffineTransform {
	rv := C.ContextGetCTM(
		// *typing.RefType
		unsafe.Pointer(c),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Returns a flattened data representation of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454381-cgeventcreatedata?language=objc
func EventCreateData(allocator corefoundation.AllocatorRef, event EventRef) corefoundation.DataRef {
	rv := C.EventCreateData(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(event),
	)
	// *typing.RefType
	return corefoundation.DataRef(rv)
}

// Returns a Boolean value indicating whether a display is running in a stereo graphics mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455025-cgdisplayisstereo?language=objc
func DisplayIsStereo(display DirectDisplayID) int {
	rv := C.DisplayIsStereo(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return int(rv)
}

// Transforms the user coordinate system in a context using a specified matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454897-cgcontextconcatctm?language=objc
func ContextConcatCTM(c ContextRef, transform AffineTransform) {
	C.ContextConcatCTM(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&transform)),
	)
}

// Returns a size that is transformed from device space coordinates to user space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456510-cgcontextconvertsizetouserspace?language=objc
func ContextConvertSizeToUserSpace(c ContextRef, size Size) Size {
	rv := C.ContextConvertSizeToUserSpace(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGSize)(unsafe.Pointer(&size)),
	)
	// *typing.StructType
	return *(*Size)(unsafe.Pointer(&rv))
}

// Retrieves an integer object from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454399-cgpdfscannerpopinteger?language=objc
func PDFScannerPopInteger(scanner unsafe.Pointer, value *PDFInteger) bool {
	rv := C.PDFScannerPopInteger(
		// *typing.RefType
		unsafe.Pointer(scanner),
		// *typing.PointerType
		(*C.CGPDFInteger)(unsafe.Pointer(&value)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Associates custom metadata with the PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456026-cgpdfcontextadddocumentmetadata?language=objc
func PDFContextAddDocumentMetadata(context ContextRef, metadata corefoundation.DataRef) {
	C.PDFContextAddDocumentMetadata(
		// *typing.RefType
		unsafe.Pointer(context),
		// *typing.RefType
		unsafe.Pointer(metadata),
	)
}

// Returns the width and height of a display in millimeters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456599-cgdisplayscreensize?language=objc
func DisplayScreenSize(display DirectDisplayID) Size {
	rv := C.DisplayScreenSize(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.StructType
	return *(*Size)(unsafe.Pointer(&rv))
}

// Copies the entries in the color table of an indexed color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408853-cgcolorspacegetcolortable?language=objc
func ColorSpaceGetColorTable(space ColorSpaceRef, table *uint8) {
	C.ColorSpaceGetColorTable(
		// *typing.RefType
		unsafe.Pointer(space),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&table)),
	)
}

// Increments the retain count of a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1571723-cgpdfpageretain?language=objc
func PDFPageRetain(page PDFPageRef) PDFPageRef {
	rv := C.PDFPageRetain(
		// *typing.RefType
		unsafe.Pointer(page),
	)
	// *typing.RefType
	return PDFPageRef(rv)
}

// Paints the area of the ellipse that fits inside the provided rectangle, using the fill color in the current graphics state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454371-cgcontextfillellipseinrect?language=objc
func ContextFillEllipseInRect(c ContextRef, rect Rect) {
	C.ContextFillEllipseInRect(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

// Returns information about a display’s current configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454099-cgdisplaycopydisplaymode?language=objc
func DisplayCopyDisplayMode(display DirectDisplayID) DisplayModeRef {
	rv := C.DisplayCopyDisplayMode(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.RefType
	return DisplayModeRef(rv)
}

// Sets a destination to jump to when a rectangle in the current PDF page is clicked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456459-cgpdfcontextsetdestinationforrec?language=objc
func PDFContextSetDestinationForRect(context ContextRef, name corefoundation.StringRef, rect Rect) {
	C.PDFContextSetDestinationForRect(
		// *typing.RefType
		unsafe.Pointer(context),
		// *typing.RefType
		unsafe.Pointer(name),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

// Provides a list of displays that corresponds to the bits set in an OpenGL display mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454234-cggetdisplayswithopengldisplayma?language=objc
func GetDisplaysWithOpenGLDisplayMask(mask OpenGLDisplayMask, maxDisplays uint32, displays *DirectDisplayID, matchingDisplayCount *uint32) Error {
	rv := C.GetDisplaysWithOpenGLDisplayMask(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGOpenGLDisplayMask)(mask),
		// *typing.PrimitiveType
		C.uint32_t(maxDisplays),
		// *typing.PointerType
		(*C.CGDirectDisplayID)(unsafe.Pointer(&displays)),
		// *typing.PointerType
		(*C.uint32_t)(unsafe.Pointer(&matchingDisplayCount)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Creates a shading object to use for axial shading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455224-cgshadingcreateaxial?language=objc
func ShadingCreateAxial(space ColorSpaceRef, start Point, end Point, function FunctionRef, extendStart bool, extendEnd bool) ShadingRef {
	rv := C.ShadingCreateAxial(
		// *typing.RefType
		unsafe.Pointer(space),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&start)),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&end)),
		// *typing.RefType
		unsafe.Pointer(function),
		// *typing.PrimitiveType
		C.bool(extendStart),
		// *typing.PrimitiveType
		C.bool(extendEnd),
	)
	// *typing.RefType
	return ShadingRef(rv)
}

// Enables shadowing in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456082-cgcontextsetshadow?language=objc
func ContextSetShadow(c ContextRef, offset Size, blur float64) {
	C.ContextSetShadow(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGSize)(unsafe.Pointer(&offset)),
		// *typing.PrimitiveType
		C.float(blur),
	)
}

// Fills in a point using the contents of the specified dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455338-cgpointmakewithdictionaryreprese?language=objc
func PointMakeWithDictionaryRepresentation(dict corefoundation.DictionaryRef, point *Point) bool {
	rv := C.PointMakeWithDictionaryRepresentation(
		// *typing.RefType
		unsafe.Pointer(dict),
		// *typing.PointerType
		(*C.CGPoint)(unsafe.Pointer(&point)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the location at which text is drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454687-cgcontextgettextposition?language=objc
func ContextGetTextPosition(c ContextRef) Point {
	rv := C.ContextGetTextPosition(
		// *typing.RefType
		unsafe.Pointer(c),
	)
	// *typing.StructType
	return *(*Point)(unsafe.Pointer(&rv))
}

// Returns the bits per pixel of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455946-cgbitmapcontextgetbitsperpixel?language=objc
func BitmapContextGetBitsPerPixel(context ContextRef) uint {
	rv := C.BitmapContextGetBitsPerPixel(
		// *typing.RefType
		unsafe.Pointer(context),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Adds an arc of a circle to the current path, possibly preceded by a straight line segment [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455756-cgcontextaddarc?language=objc
func ContextAddArc(c ContextRef, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise int) {
	C.ContextAddArc(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
		// *typing.PrimitiveType
		C.float(radius),
		// *typing.PrimitiveType
		C.float(startAngle),
		// *typing.PrimitiveType
		C.float(endAngle),
		// *typing.PrimitiveType
		C.int(clockwise),
	)
}

// Sets the event type of a Quartz event (left mouse down, for example). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454300-cgeventsettype?language=objc
func EventSetType(event EventRef, type_ EventType) {
	C.EventSetType(
		// *typing.RefType
		unsafe.Pointer(event),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventType)(type_),
	)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2919739-cgeventcreatescrollwheelevent2?language=objc
func EventCreateScrollWheelEvent2(source EventSourceRef, units ScrollEventUnit, wheelCount uint32, wheel1 int32, wheel2 int32, wheel3 int32) EventRef {
	rv := C.EventCreateScrollWheelEvent2(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGScrollEventUnit)(units),
		// *typing.PrimitiveType
		C.uint32_t(wheelCount),
		// *typing.PrimitiveType
		C.int32_t(wheel1),
		// *typing.PrimitiveType
		C.int32_t(wheel2),
		// *typing.PrimitiveType
		C.int32_t(wheel3),
	)
	// *typing.RefType
	return EventRef(rv)
}

// Returns the color space model of the provided color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408854-cgcolorspacegetmodel?language=objc
func ColorSpaceGetModel(space ColorSpaceRef) ColorSpaceModel {
	rv := C.ColorSpaceGetModel(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.AliasType
	return ColorSpaceModel(rv)
}

// Creates a bitmap image using PNG-encoded data supplied by a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454993-cgimagecreatewithpngdataprovider?language=objc
func ImageCreateWithPNGDataProvider(source DataProviderRef, decode *float64, shouldInterpolate bool, intent ColorRenderingIntent) ImageRef {
	rv := C.ImageCreateWithPNGDataProvider(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&decode)),
		// *typing.PrimitiveType
		C.bool(shouldInterpolate),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGColorRenderingIntent)(intent),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Adds to a path an ellipse that fits inside a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411222-cgpathaddellipseinrect?language=objc
func PathAddEllipseInRect(path MutablePathRef, m *AffineTransform, rect Rect) {
	C.PathAddEllipseInRect(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

// Returns a Quartz event created from a flattened data representation of the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454249-cgeventcreatefromdata?language=objc
func EventCreateFromData(allocator corefoundation.AllocatorRef, data corefoundation.DataRef) EventRef {
	rv := C.EventCreateFromData(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(data),
	)
	// *typing.RefType
	return EventRef(rv)
}

// Creates a device-dependent CMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408897-cgcolorspacecreatedevicecmyk?language=objc
func ColorSpaceCreateDeviceCMYK() ColorSpaceRef {
	rv := C.ColorSpaceCreateDeviceCMYK()
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Modifies the settings of the built-in fade effect that occurs during a display configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454103-cgconfiguredisplayfadeeffect?language=objc
func ConfigureDisplayFadeEffect(config unsafe.Pointer, fadeOutSeconds DisplayFadeInterval, fadeInSeconds DisplayFadeInterval, fadeRed float64, fadeGreen float64, fadeBlue float64) Error {
	rv := C.ConfigureDisplayFadeEffect(
		// *typing.RefType
		unsafe.Pointer(config),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDisplayFadeInterval)(fadeOutSeconds),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDisplayFadeInterval)(fadeInSeconds),
		// *typing.PrimitiveType
		C.float(fadeRed),
		// *typing.PrimitiveType
		C.float(fadeGreen),
		// *typing.PrimitiveType
		C.float(fadeBlue),
	)
	// *typing.AliasType
	return Error(rv)
}

// Performs a single fade operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456189-cgdisplayfade?language=objc
func DisplayFade(token DisplayFadeReservationToken, duration DisplayFadeInterval, startBlend DisplayBlendFraction, endBlend DisplayBlendFraction, redBlend float64, greenBlend float64, blueBlend float64, synchronous int) Error {
	rv := C.DisplayFade(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDisplayFadeReservationToken)(token),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDisplayFadeInterval)(duration),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDisplayBlendFraction)(startBlend),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDisplayBlendFraction)(endBlend),
		// *typing.PrimitiveType
		C.float(redBlend),
		// *typing.PrimitiveType
		C.float(greenBlend),
		// *typing.PrimitiveType
		C.float(blueBlend),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.boolean_t)(synchronous),
	)
	// *typing.AliasType
	return Error(rv)
}

// Enables or disables subpixel quantization in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455766-cgcontextsetshouldsubpixelquanti?language=objc
func ContextSetShouldSubpixelQuantizeFonts(c ContextRef, shouldSubpixelQuantizeFonts bool) {
	C.ContextSetShouldSubpixelQuantizeFonts(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.bool(shouldSubpixelQuantizeFonts),
	)
}

// Returns the location of a Quartz mouse event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455788-cgeventgetlocation?language=objc
func EventGetLocation(event EventRef) Point {
	rv := C.EventGetLocation(
		// *typing.RefType
		unsafe.Pointer(event),
	)
	// *typing.StructType
	return *(*Point)(unsafe.Pointer(&rv))
}

// Tells a PostScript converter to abort a conversion at the next available opportunity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455046-cgpsconverterabort?language=objc
func PSConverterAbort(converter PSConverterRef) bool {
	rv := C.PSConverterAbort(
		// *typing.RefType
		unsafe.Pointer(converter),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Changes the scale of the user coordinate system in a context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454659-cgcontextscalectm?language=objc
func ContextScaleCTM(c ContextRef, sx float64, sy float64) {
	C.ContextScaleCTM(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(sx),
		// *typing.PrimitiveType
		C.float(sy),
	)
}

// Returns a dictionary representation of the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455274-cgsizecreatedictionaryrepresenta?language=objc
func SizeCreateDictionaryRepresentation(size Size) corefoundation.DictionaryRef {
	rv := C.SizeCreateDictionaryRepresentation(
		// *typing.StructType
		*(*C.CGSize)(unsafe.Pointer(&size)),
	)
	// *typing.RefType
	return corefoundation.DictionaryRef(rv)
}

// Returns a Boolean value indicating the current button state of a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408781-cgeventsourcebuttonstate?language=objc
func EventSourceButtonState(stateID EventSourceStateID, button MouseButton) bool {
	rv := C.EventSourceButtonState(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventSourceStateID)(stateID),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGMouseButton)(button),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Enables or disables an event tap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455445-cgeventtapenable?language=objc
func EventTapEnable(tap corefoundation.MachPortRef, enable bool) {
	C.EventTapEnable(
		// *typing.RefType
		unsafe.Pointer(tap),
		// *typing.PrimitiveType
		C.bool(enable),
	)
}

// Returns an affine transformation matrix constructed from a rotation value you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455666-cgaffinetransformmakerotation?language=objc
func AffineTransformMakeRotation(angle float64) AffineTransform {
	rv := C.AffineTransformMakeRotation(
		// *typing.PrimitiveType
		C.float(angle),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Creates a device-independent color space that is defined according to the ICC color profile specification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408881-cgcolorspacecreateiccbased?language=objc
func ColorSpaceCreateICCBased(nComponents uint, range_ *float64, profile DataProviderRef, alternate ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateICCBased(
		// *typing.PrimitiveType
		C.uint(nComponents),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&range_)),
		// *typing.RefType
		unsafe.Pointer(profile),
		// *typing.RefType
		unsafe.Pointer(alternate),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Connects or disconnects the mouse and cursor while an application is in the foreground. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454486-cgassociatemouseandmousecursorpo?language=objc
func AssociateMouseAndMouseCursorPosition(connected int) Error {
	rv := C.AssociateMouseAndMouseCursorPosition(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.boolean_t)(connected),
	)
	// *typing.AliasType
	return Error(rv)
}

// Increments the retain count of a Core Graphics PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402587-cgpdfdocumentretain?language=objc
func PDFDocumentRetain(document PDFDocumentRef) PDFDocumentRef {
	rv := C.PDFDocumentRetain(
		// *typing.RefType
		unsafe.Pointer(document),
	)
	// *typing.RefType
	return PDFDocumentRef(rv)
}

// Returns the current point in a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411132-cgpathgetcurrentpoint?language=objc
func PathGetCurrentPoint(path unsafe.Pointer) Point {
	rv := C.PathGetCurrentPoint(
		// *typing.RefType
		unsafe.Pointer(path),
	)
	// *typing.StructType
	return *(*Point)(unsafe.Pointer(&rv))
}

// Sets the current stroke color to a value in the DeviceRGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456378-cgcontextsetrgbstrokecolor?language=objc
func ContextSetRGBStrokeColor(c ContextRef, red float64, green float64, blue float64, alpha float64) {
	C.ContextSetRGBStrokeColor(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(red),
		// *typing.PrimitiveType
		C.float(green),
		// *typing.PrimitiveType
		C.float(blue),
		// *typing.PrimitiveType
		C.float(alpha),
	)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656523-cgpreflightscreencaptureaccess?language=objc
func PreflightScreenCaptureAccess() bool {
	rv := C.PreflightScreenCaptureAccess()
	// *typing.PrimitiveType
	return bool(rv)
}

// Sets the color gamma function for a display by specifying the values in the RGB gamma tables. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456604-cgsetdisplaytransferbytable?language=objc
func SetDisplayTransferByTable(display DirectDisplayID, tableSize uint32, redTable *GammaValue, greenTable *GammaValue, blueTable *GammaValue) Error {
	rv := C.SetDisplayTransferByTable(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.PrimitiveType
		C.uint32_t(tableSize),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&redTable)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&greenTable)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&blueTable)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Sets the accuracy of curved paths in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455798-cgcontextsetflatness?language=objc
func ContextSetFlatness(c ContextRef, flatness float64) {
	C.ContextSetFlatness(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(flatness),
	)
}

// Returns a Boolean value indicating whether a display is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455222-cgdisplayisactive?language=objc
func DisplayIsActive(display DirectDisplayID) int {
	rv := C.DisplayIsActive(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return int(rv)
}

// Sets the pattern phase of a context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455334-cgcontextsetpatternphase?language=objc
func ContextSetPatternPhase(c ContextRef, phase Size) {
	C.ContextSetPatternPhase(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGSize)(unsafe.Pointer(&phase)),
	)
}

// Decrements the retain count of a scanner object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454962-cgpdfscannerrelease?language=objc
func PDFScannerRelease(scanner unsafe.Pointer) {
	C.PDFScannerRelease(
		// *typing.RefType
		unsafe.Pointer(scanner),
	)
}

// Sets the style for the joins of connected lines in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455973-cgcontextsetlinejoin?language=objc
func ContextSetLineJoin(c ContextRef, join LineJoin) {
	C.ContextSetLineJoin(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGLineJoin)(join),
	)
}

// Returns an array of tags that correspond to the font tables for a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396392-cgfontcopytabletags?language=objc
func FontCopyTableTags(font FontRef) corefoundation.ArrayRef {
	rv := C.FontCopyTableTags(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.RefType
	return corefoundation.ArrayRef(rv)
}

// Provides a list of displays that are active for drawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454603-cggetactivedisplaylist?language=objc
func GetActiveDisplayList(maxDisplays uint32, activeDisplays *DirectDisplayID, displayCount *uint32) Error {
	rv := C.GetActiveDisplayList(
		// *typing.PrimitiveType
		C.uint32_t(maxDisplays),
		// *typing.PointerType
		(*C.CGDirectDisplayID)(unsafe.Pointer(&activeDisplays)),
		// *typing.PointerType
		(*C.uint32_t)(unsafe.Pointer(&displayCount)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Indicates whether or not a graphics path represents a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411163-cgpathisrect?language=objc
func PathIsRect(path unsafe.Pointer, rect *Rect) bool {
	rv := C.PathIsRect(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a direct-access data provider that uses a URL to supply data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408327-cgdataprovidercreatewithurl?language=objc
func DataProviderCreateWithURL(url corefoundation.URLRef) DataProviderRef {
	rv := C.DataProviderCreateWithURL(
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.RefType
	return DataProviderRef(rv)
}

// Returns a Quartz event source created with a specified source state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408776-cgeventsourcecreate?language=objc
func EventSourceCreate(stateID EventSourceStateID) EventSourceRef {
	rv := C.EventSourceCreate(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventSourceStateID)(stateID),
	)
	// *typing.RefType
	return EventSourceRef(rv)
}

// Returns the current level of interpolation quality for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454940-cgcontextgetinterpolationquality?language=objc
func ContextGetInterpolationQuality(c ContextRef) InterpolationQuality {
	rv := C.ContextGetInterpolationQuality(
		// *typing.RefType
		unsafe.Pointer(c),
	)
	// *typing.AliasType
	return InterpolationQuality(rv)
}

// Creates a conversion between two specified color spaces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2113677-cgcolorconversioninfocreate?language=objc
func ColorConversionInfoCreate(src ColorSpaceRef, dst ColorSpaceRef) ColorConversionInfoRef {
	rv := C.ColorConversionInfoCreate(
		// *typing.RefType
		unsafe.Pointer(src),
		// *typing.RefType
		unsafe.Pointer(dst),
	)
	// *typing.RefType
	return ColorConversionInfoRef(rv)
}

// Returns whether the first rectangle contains the second rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454186-cgrectcontainsrect?language=objc
func RectContainsRect(rect1 Rect, rect2 Rect) bool {
	rv := C.RectContainsRect(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect1)),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect2)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a PostScript encoding of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396348-cgfontcreatepostscriptencoding?language=objc
func FontCreatePostScriptEncoding(font FontRef, encoding *Glyph) corefoundation.DataRef {
	rv := C.FontCreatePostScriptEncoding(
		// *typing.RefType
		unsafe.Pointer(font),
		// *typing.PointerType
		(*C.CGGlyph)(unsafe.Pointer(&encoding)),
	)
	// *typing.RefType
	return corefoundation.DataRef(rv)
}

// Draws the contents of a CGLayer object at the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450894-cgcontextdrawlayeratpoint?language=objc
func ContextDrawLayerAtPoint(context ContextRef, point Point, layer LayerRef) {
	C.ContextDrawLayerAtPoint(
		// *typing.RefType
		unsafe.Pointer(context),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&point)),
		// *typing.RefType
		unsafe.Pointer(layer),
	)
}

// Sets the scale of pixels per line in a scrolling event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408766-cgeventsourcesetpixelsperline?language=objc
func EventSourceSetPixelsPerLine(source EventSourceRef, pixelsPerLine float64) {
	C.EventSourceSetPixelsPerLine(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.PrimitiveType
		C.double(pixelsPerLine),
	)
}

// Returns a new Quartz keyboard event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456564-cgeventcreatekeyboardevent?language=objc
func EventCreateKeyboardEvent(source EventSourceRef, virtualKey KeyCode, keyDown bool) EventRef {
	rv := C.EventCreateKeyboardEvent(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGKeyCode)(virtualKey),
		// *typing.PrimitiveType
		C.bool(keyDown),
	)
	// *typing.RefType
	return EventRef(rv)
}

// Returns a count of events of a given type seen since the window server started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408794-cgeventsourcecounterforeventtype?language=objc
func EventSourceCounterForEventType(stateID EventSourceStateID, eventType EventType) uint32 {
	rv := C.EventSourceCounterForEventType(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventSourceStateID)(stateID),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventType)(eventType),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Appends an array of rectangles to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411153-cgpathaddrects?language=objc
func PathAddRects(path MutablePathRef, m *AffineTransform, rects *Rect, count uint) {
	C.PathAddRects(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&rects)),
		// *typing.PrimitiveType
		C.uint(count),
	)
}

// Sets the timestamp of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456611-cgeventsettimestamp?language=objc
func EventSetTimestamp(event EventRef, timestamp EventTimestamp) {
	C.EventSetTimestamp(
		// *typing.RefType
		unsafe.Pointer(event),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventTimestamp)(timestamp),
	)
}

// Returns a Boolean value indicating the current keyboard state of a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408768-cgeventsourcekeystate?language=objc
func EventSourceKeyState(stateID EventSourceStateID, key KeyCode) bool {
	rv := C.EventSourceKeyState(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventSourceStateID)(stateID),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGKeyCode)(key),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the height of a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455645-cgrectgetheight?language=objc
func RectGetHeight(rect Rect) float64 {
	rv := C.RectGetHeight(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Returns whether there is a PDF integer associated with a specified key in a PDF dictionary and, if so, retrieves that integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430231-cgpdfdictionarygetinteger?language=objc
func PDFDictionaryGetInteger(dict unsafe.Pointer, key string, value *PDFInteger) bool {
	keyVal := C.CString(key)
	defer C.free(unsafe.Pointer(keyVal))
	rv := C.PDFDictionaryGetInteger(
		// *typing.RefType
		unsafe.Pointer(dict),
		// *typing.CStringType
		keyVal,
		// *typing.PointerType
		(*C.CGPDFInteger)(unsafe.Pointer(&value)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the elapsed time since the last event for a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408790-cgeventsourcesecondssincelasteve?language=objc
func EventSourceSecondsSinceLastEventType(stateID EventSourceStateID, eventType EventType) corefoundation.TimeInterval {
	rv := C.EventSourceSecondsSinceLastEventType(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventSourceStateID)(stateID),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventType)(eventType),
	)
	// *typing.AliasType
	return corefoundation.TimeInterval(rv)
}

// Returns the smallest rectangle that results from converting the source rectangle values to integers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456348-cgrectintegral?language=objc
func RectIntegral(rect Rect) Rect {
	rv := C.RectIntegral(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Sets the fill color space in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455151-cgcontextsetfillcolorspace?language=objc
func ContextSetFillColorSpace(c ContextRef, space ColorSpaceRef) {
	C.ContextSetFillColorSpace(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(space),
	)
}

// Returns whether an object at a given index in a PDF array is a PDF number and, if so, retrieves that object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455374-cgpdfarraygetnumber?language=objc
func PDFArrayGetNumber(array unsafe.Pointer, index uint, value *PDFReal) bool {
	rv := C.PDFArrayGetNumber(
		// *typing.RefType
		unsafe.Pointer(array),
		// *typing.PrimitiveType
		C.uint(index),
		// *typing.PointerType
		(*C.CGPDFReal)(unsafe.Pointer(&value)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a color using a list of intensity values (including alpha), a pattern color space, and a pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455687-cgcolorcreatewithpattern?language=objc
func ColorCreateWithPattern(space ColorSpaceRef, pattern PatternRef, components *float64) ColorRef {
	rv := C.ColorCreateWithPattern(
		// *typing.RefType
		unsafe.Pointer(space),
		// *typing.RefType
		unsafe.Pointer(pattern),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&components)),
	)
	// *typing.RefType
	return ColorRef(rv)
}

// Changes the origin of the user coordinate system in a context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455286-cgcontexttranslatectm?language=objc
func ContextTranslateCTM(c ContextRef, tx float64, ty float64) {
	C.ContextTranslateCTM(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(tx),
		// *typing.PrimitiveType
		C.float(ty),
	)
}

// Returns a Boolean value indicating whether Quartz is using OpenGL-based window acceleration (Quartz Extreme) to render in a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455721-cgdisplayusesopenglacceleration?language=objc
func DisplayUsesOpenGLAcceleration(display DirectDisplayID) int {
	rv := C.DisplayUsesOpenGLAcceleration(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return int(rv)
}

// Sets how sample values are composited by a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455994-cgcontextsetblendmode?language=objc
func ContextSetBlendMode(c ContextRef, mode BlendMode) {
	C.ContextSetBlendMode(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGBlendMode)(mode),
	)
}

// Reports the change in mouse position since the last mouse movement event received by the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456484-cggetlastmousedelta?language=objc
func GetLastMouseDelta(deltaX *int32, deltaY *int32) {
	C.GetLastMouseDelta(
		// *typing.PointerType
		(*C.int32_t)(unsafe.Pointer(&deltaX)),
		// *typing.PointerType
		(*C.int32_t)(unsafe.Pointer(&deltaY)),
	)
}

// Get the bounding box of each glyph in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396342-cgfontgetglyphbboxes?language=objc
func FontGetGlyphBBoxes(font FontRef, glyphs *Glyph, count uint, bboxes *Rect) bool {
	rv := C.FontGetGlyphBBoxes(
		// *typing.RefType
		unsafe.Pointer(font),
		// *typing.PointerType
		(*C.CGGlyph)(unsafe.Pointer(&glyphs)),
		// *typing.PrimitiveType
		C.uint(count),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&bboxes)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Provides a list of online displays with bounds that include the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454385-cggetdisplayswithpoint?language=objc
func GetDisplaysWithPoint(point Point, maxDisplays uint32, displays *DirectDisplayID, matchingDisplayCount *uint32) Error {
	rv := C.GetDisplaysWithPoint(
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&point)),
		// *typing.PrimitiveType
		C.uint32_t(maxDisplays),
		// *typing.PointerType
		(*C.CGDirectDisplayID)(unsafe.Pointer(&displays)),
		// *typing.PointerType
		(*C.uint32_t)(unsafe.Pointer(&matchingDisplayCount)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Decrements the retain count of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1556742-cgimagerelease?language=objc
func ImageRelease(image ImageRef) {
	C.ImageRelease(
		// *typing.RefType
		unsafe.Pointer(image),
	)
}

// Returns the smallest value for the x-coordinate of the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455948-cgrectgetminx?language=objc
func RectGetMinX(rect Rect) float64 {
	rv := C.RectGetMinX(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Replaces the path in the graphics context with the stroked version of the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454517-cgcontextreplacepathwithstrokedp?language=objc
func ContextReplacePathWithStrokedPath(c ContextRef) {
	C.ContextReplacePathWithStrokedPath(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Returns the bytes per row of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456129-cgbitmapcontextgetbytesperrow?language=objc
func BitmapContextGetBytesPerRow(context ContextRef) uint {
	rv := C.BitmapContextGetBytesPerRow(
		// *typing.RefType
		unsafe.Pointer(context),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Returns the list of window IDs associated with the specified windows in the current user session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1552209-cgwindowlistcreate?language=objc
func WindowListCreate(option WindowListOption, relativeToWindow WindowID) corefoundation.ArrayRef {
	rv := C.WindowListCreate(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGWindowListOption)(option),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGWindowID)(relativeToWindow),
	)
	// *typing.RefType
	return corefoundation.ArrayRef(rv)
}

// Returns the height and width resulting from a transformation of an existing height and width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454806-cgsizeapplyaffinetransform?language=objc
func SizeApplyAffineTransform(size Size, t AffineTransform) Size {
	rv := C.SizeApplyAffineTransform(
		// *typing.StructType
		*(*C.CGSize)(unsafe.Pointer(&size)),
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t)),
	)
	// *typing.StructType
	return *(*Size)(unsafe.Pointer(&rv))
}

// Posts a Quartz event from an event tap into the event stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455172-cgeventtappostevent?language=objc
func EventTapPostEvent(proxy EventTapProxy, event EventRef) {
	C.EventTapPostEvent(
		// *typing.AliasType
		// *typing.PointerType
		(C.CGEventTapProxy)(proxy),
		// *typing.RefType
		unsafe.Pointer(event),
	)
}

// Returns a Boolean value indicating whether a display is in a mirroring set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455558-cgdisplayisinmirrorset?language=objc
func DisplayIsInMirrorSet(display DirectDisplayID) int {
	rv := C.DisplayIsInMirrorSet(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return int(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3861800-cgcolorspaceispqbased?language=objc
func ColorSpaceIsPQBased(s ColorSpaceRef) bool {
	rv := C.ColorSpaceIsPQBased(
		// *typing.RefType
		unsafe.Pointer(s),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Paints a transparent rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456457-cgcontextclearrect?language=objc
func ContextClearRect(c ContextRef, rect Rect) {
	C.ContextClearRect(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

// Releases a captured display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455685-cgdisplayrelease?language=objc
func DisplayRelease(display DirectDisplayID) Error {
	rv := C.DisplayRelease(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return Error(rv)
}

// Returns the document for a page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456166-cgpdfpagegetdocument?language=objc
func PDFPageGetDocument(page PDFPageRef) PDFDocumentRef {
	rv := C.PDFPageGetDocument(
		// *typing.RefType
		unsafe.Pointer(page),
	)
	// *typing.RefType
	return PDFDocumentRef(rv)
}

// Gets the values in the RGB gamma tables for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454974-cggetdisplaytransferbytable?language=objc
func GetDisplayTransferByTable(display DirectDisplayID, capacity uint32, redTable *GammaValue, greenTable *GammaValue, blueTable *GammaValue, sampleCount *uint32) Error {
	rv := C.GetDisplayTransferByTable(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.PrimitiveType
		C.uint32_t(capacity),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&redTable)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&greenTable)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&blueTable)),
		// *typing.PointerType
		(*C.uint32_t)(unsafe.Pointer(&sampleCount)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Checks whether two affine transforms are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455732-cgaffinetransformequaltotransfor?language=objc
func AffineTransformEqualToTransform(t1 AffineTransform, t2 AffineTransform) bool {
	rv := C.AffineTransformEqualToTransform(
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t1)),
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t2)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the content stream associated with a PDF scanner object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454724-cgpdfscannergetcontentstream?language=objc
func PDFScannerGetContentStream(scanner unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFScannerGetContentStream(
		// *typing.RefType
		unsafe.Pointer(scanner),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Appends a line segment to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411138-cgpathaddlinetopoint?language=objc
func PathAddLineToPoint(path MutablePathRef, m *AffineTransform, x float64, y float64) {
	C.PathAddLineToPoint(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656519-cgpreflightlisteneventaccess?language=objc
func PreflightListenEventAccess() bool {
	rv := C.PreflightListenEventAccess()
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the event type of a Quartz event (left mouse down, for example). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455634-cgeventgettype?language=objc
func EventGetType(event EventRef) EventType {
	rv := C.EventGetType(
		// *typing.RefType
		unsafe.Pointer(event),
	)
	// *typing.AliasType
	return EventType(rv)
}

// Retrieves a real value object from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456297-cgpdfscannerpopnumber?language=objc
func PDFScannerPopNumber(scanner unsafe.Pointer, value *PDFReal) bool {
	rv := C.PDFScannerPopNumber(
		// *typing.RefType
		unsafe.Pointer(scanner),
		// *typing.PointerType
		(*C.CGPDFReal)(unsafe.Pointer(&value)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a device-independent color space that is relative to human color perception, according to the CIE L*a*b* standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408879-cgcolorspacecreatelab?language=objc
func ColorSpaceCreateLab(whitePoint *float64, blackPoint *float64, range_ *float64) ColorSpaceRef {
	rv := C.ColorSpaceCreateLab(
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&whitePoint)),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&blackPoint)),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&range_)),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Ends the current page in the PDF graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456122-cgpdfcontextendpage?language=objc
func PDFContextEndPage(context ContextRef) {
	C.PDFContextEndPage(
		// *typing.RefType
		unsafe.Pointer(context),
	)
}

// Returns the bounding box of a clipping path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455387-cgcontextgetclipboundingbox?language=objc
func ContextGetClipBoundingBox(c ContextRef) Rect {
	rv := C.ContextGetClipBoundingBox(
		// *typing.RefType
		unsafe.Pointer(c),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Creates an empty PDF operator table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455932-cgpdfoperatortablecreate?language=objc
func PDFOperatorTableCreate() unsafe.Pointer {
	rv := C.PDFOperatorTableCreate()
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Appends a rectangle to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411144-cgpathaddrect?language=objc
func PathAddRect(path MutablePathRef, m *AffineTransform, rect Rect) {
	C.PathAddRect(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

// Returns the point resulting from an affine transformation of an existing point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454251-cgpointapplyaffinetransform?language=objc
func PointApplyAffineTransform(point Point, t AffineTransform) Point {
	rv := C.PointApplyAffineTransform(
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&point)),
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t)),
	)
	// *typing.StructType
	return *(*Point)(unsafe.Pointer(&rv))
}

// Creates a bitmap image using JPEG-encoded data supplied by a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454920-cgimagecreatewithjpegdataprovide?language=objc
func ImageCreateWithJPEGDataProvider(source DataProviderRef, decode *float64, shouldInterpolate bool, intent ColorRenderingIntent) ImageRef {
	rv := C.ImageCreateWithJPEGDataProvider(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&decode)),
		// *typing.PrimitiveType
		C.bool(shouldInterpolate),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGColorRenderingIntent)(intent),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Returns the type identifier for the opaque type CGEventSourceRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408789-cgeventsourcegettypeid?language=objc
func EventSourceGetTypeID() corefoundation.TypeID {
	rv := C.EventSourceGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Creates a bitmap image from data supplied by a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455149-cgimagecreate?language=objc
func ImageCreate(width uint, height uint, bitsPerComponent uint, bitsPerPixel uint, bytesPerRow uint, space ColorSpaceRef, bitmapInfo BitmapInfo, provider DataProviderRef, decode *float64, shouldInterpolate bool, intent ColorRenderingIntent) ImageRef {
	rv := C.ImageCreate(
		// *typing.PrimitiveType
		C.uint(width),
		// *typing.PrimitiveType
		C.uint(height),
		// *typing.PrimitiveType
		C.uint(bitsPerComponent),
		// *typing.PrimitiveType
		C.uint(bitsPerPixel),
		// *typing.PrimitiveType
		C.uint(bytesPerRow),
		// *typing.RefType
		unsafe.Pointer(space),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGBitmapInfo)(bitmapInfo),
		// *typing.RefType
		unsafe.Pointer(provider),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&decode)),
		// *typing.PrimitiveType
		C.bool(shouldInterpolate),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGColorRenderingIntent)(intent),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Returns whether the rectangle is equal to the null rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455471-cgrectisnull?language=objc
func RectIsNull(rect Rect) bool {
	rv := C.RectIsNull(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the source state associated with a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408772-cgeventsourcegetsourcestateid?language=objc
func EventSourceGetSourceStateID(source EventSourceRef) EventSourceStateID {
	rv := C.EventSourceGetSourceStateID(
		// *typing.RefType
		unsafe.Pointer(source),
	)
	// *typing.AliasType
	return EventSourceStateID(rv)
}

// Increments the retain count of a CGGradient object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1398456-cggradientretain?language=objc
func GradientRetain(gradient GradientRef) GradientRef {
	rv := C.GradientRetain(
		// *typing.RefType
		unsafe.Pointer(gradient),
	)
	// *typing.RefType
	return GradientRef(rv)
}

// Returns whether an object at a given index in a PDF array is a PDF Boolean and, if so, retrieves that Boolean. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454504-cgpdfarraygetboolean?language=objc
func PDFArrayGetBoolean(array unsafe.Pointer, index uint, value *PDFBoolean) bool {
	rv := C.PDFArrayGetBoolean(
		// *typing.RefType
		unsafe.Pointer(array),
		// *typing.PrimitiveType
		C.uint(index),
		// *typing.PointerType
		(*C.CGPDFBoolean)(unsafe.Pointer(&value)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Sets the current fill color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455296-cgcontextsetfillcolor?language=objc
func ContextSetFillColor(c ContextRef, components *float64) {
	C.ContextSetFillColor(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&components)),
	)
}

// Returns the pattern associated with a color in a pattern color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455937-cgcolorgetpattern?language=objc
func ColorGetPattern(color ColorRef) PatternRef {
	rv := C.ColorGetPattern(
		// *typing.RefType
		unsafe.Pointer(color),
	)
	// *typing.RefType
	return PatternRef(rv)
}

// Returns the bounding box of a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411200-cgpathgetpathboundingbox?language=objc
func PathGetPathBoundingBox(path unsafe.Pointer) Rect {
	rv := C.PathGetPathBoundingBox(
		// *typing.RefType
		unsafe.Pointer(path),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Captures all attached displays, using the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456514-cgcapturealldisplayswithoptions?language=objc
func CaptureAllDisplaysWithOptions(options CaptureOptions) Error {
	rv := C.CaptureAllDisplaysWithOptions(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGCaptureOptions)(options),
	)
	// *typing.AliasType
	return Error(rv)
}

// Generates and returns information about windows with the specified window IDs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455215-cgwindowlistcreatedescriptionfro?language=objc
func WindowListCreateDescriptionFromArray(windowArray corefoundation.ArrayRef) corefoundation.ArrayRef {
	rv := C.WindowListCreateDescriptionFromArray(
		// *typing.RefType
		unsafe.Pointer(windowArray),
	)
	// *typing.RefType
	return corefoundation.ArrayRef(rv)
}

// Creates a PDF content stream object from an existing PDF content stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410041-cgpdfcontentstreamcreatewithstre?language=objc
func PDFContentStreamCreateWithStream(stream unsafe.Pointer, streamResources unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFContentStreamCreateWithStream(
		// *typing.RefType
		unsafe.Pointer(stream),
		// *typing.RefType
		unsafe.Pointer(streamResources),
		// *typing.RefType
		unsafe.Pointer(parent),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Paints a gradient fill that varies along the area defined by the provided starting and ending circles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455923-cgcontextdrawradialgradient?language=objc
func ContextDrawRadialGradient(c ContextRef, gradient GradientRef, startCenter Point, startRadius float64, endCenter Point, endRadius float64, options GradientDrawingOptions) {
	C.ContextDrawRadialGradient(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(gradient),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&startCenter)),
		// *typing.PrimitiveType
		C.float(startRadius),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&endCenter)),
		// *typing.PrimitiveType
		C.float(endRadius),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGGradientDrawingOptions)(options),
	)
}

// Converts a string to a date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454295-cgpdfstringcopydate?language=objc
func PDFStringCopyDate(string_ unsafe.Pointer) corefoundation.DateRef {
	rv := C.PDFStringCopyDate(
		// *typing.RefType
		unsafe.Pointer(string_),
	)
	// *typing.RefType
	return corefoundation.DateRef(rv)
}

// Provides a list of displays that are online (active, mirrored, or sleeping). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454964-cggetonlinedisplaylist?language=objc
func GetOnlineDisplayList(maxDisplays uint32, onlineDisplays *DirectDisplayID, displayCount *uint32) Error {
	rv := C.GetOnlineDisplayList(
		// *typing.PrimitiveType
		C.uint32_t(maxDisplays),
		// *typing.PointerType
		(*C.CGDirectDisplayID)(unsafe.Pointer(&onlineDisplays)),
		// *typing.PointerType
		(*C.uint32_t)(unsafe.Pointer(&displayCount)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Returns a point that is transformed from user space coordinates to device space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455916-cgcontextconvertpointtodevicespa?language=objc
func ContextConvertPointToDeviceSpace(c ContextRef, point Point) Point {
	rv := C.ContextConvertPointToDeviceSpace(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&point)),
	)
	// *typing.StructType
	return *(*Point)(unsafe.Pointer(&rv))
}

// Creates an immutable copy of a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411211-cgpathcreatecopy?language=objc
func PathCreateCopy(path unsafe.Pointer) unsafe.Pointer {
	rv := C.PathCreateCopy(
		// *typing.RefType
		unsafe.Pointer(path),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Creates and returns a CGImage from the pixel data in a bitmap graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454225-cgbitmapcontextcreateimage?language=objc
func BitmapContextCreateImage(context ContextRef) ImageRef {
	rv := C.BitmapContextCreateImage(
		// *typing.RefType
		unsafe.Pointer(context),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Returns the number of color components in a color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408848-cgcolorspacegetnumberofcomponent?language=objc
func ColorSpaceGetNumberOfComponents(space ColorSpaceRef) uint {
	rv := C.ColorSpaceGetNumberOfComponents(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Increments the retain count of a function object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1390869-cgfunctionretain?language=objc
func FunctionRetain(function FunctionRef) FunctionRef {
	rv := C.FunctionRetain(
		// *typing.RefType
		unsafe.Pointer(function),
	)
	// *typing.RefType
	return FunctionRef(rv)
}

// Gets the coefficients of the gamma transfer formula for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456330-cggetdisplaytransferbyformula?language=objc
func GetDisplayTransferByFormula(display DirectDisplayID, redMin *GammaValue, redMax *GammaValue, redGamma *GammaValue, greenMin *GammaValue, greenMax *GammaValue, greenGamma *GammaValue, blueMin *GammaValue, blueMax *GammaValue, blueGamma *GammaValue) Error {
	rv := C.GetDisplayTransferByFormula(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&redMin)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&redMax)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&redGamma)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&greenMin)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&greenMax)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&greenGamma)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&blueMin)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&blueMax)),
		// *typing.PointerType
		(*C.CGGammaValue)(unsafe.Pointer(&blueGamma)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Increments the retain count of a color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408885-cgcolorspaceretain?language=objc
func ColorSpaceRetain(space ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceRetain(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Creates a color in the sRGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3042355-cgcolorcreatesrgb?language=objc
func ColorCreateSRGB(red float64, green float64, blue float64, alpha float64) ColorRef {
	rv := C.ColorCreateSRGB(
		// *typing.PrimitiveType
		C.float(red),
		// *typing.PrimitiveType
		C.float(green),
		// *typing.PrimitiveType
		C.float(blue),
		// *typing.PrimitiveType
		C.float(alpha),
	)
	// *typing.RefType
	return ColorRef(rv)
}

// Return the color space for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454858-cgimagegetcolorspace?language=objc
func ImageGetColorSpace(image ImageRef) ColorSpaceRef {
	rv := C.ImageGetColorSpace(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Creates a shading object to use for radial shading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456399-cgshadingcreateradial?language=objc
func ShadingCreateRadial(space ColorSpaceRef, start Point, startRadius float64, end Point, endRadius float64, function FunctionRef, extendStart bool, extendEnd bool) ShadingRef {
	rv := C.ShadingCreateRadial(
		// *typing.RefType
		unsafe.Pointer(space),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&start)),
		// *typing.PrimitiveType
		C.float(startRadius),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&end)),
		// *typing.PrimitiveType
		C.float(endRadius),
		// *typing.RefType
		unsafe.Pointer(function),
		// *typing.PrimitiveType
		C.bool(extendStart),
		// *typing.PrimitiveType
		C.bool(extendEnd),
	)
	// *typing.RefType
	return ShadingRef(rv)
}

// Divides a source rectangle into two component rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455925-cgrectdivide?language=objc
func RectDivide(rect Rect, slice *Rect, remainder *Rect, amount float64, edge RectEdge) {
	C.RectDivide(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&slice)),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&remainder)),
		// *typing.PrimitiveType
		C.float(amount),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGRectEdge)(edge),
	)
}

// Sets whether or not to allow antialiasing for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456310-cgcontextsetallowsantialiasing?language=objc
func ContextSetAllowsAntialiasing(c ContextRef, allowsAntialiasing bool) {
	C.ContextSetAllowsAntialiasing(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.bool(allowsAntialiasing),
	)
}

// Returns the height in pixels of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454681-cgbitmapcontextgetheight?language=objc
func BitmapContextGetHeight(context ContextRef) uint {
	rv := C.BitmapContextGetHeight(
		// *typing.RefType
		unsafe.Pointer(context),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Returns the value of the alpha component associated with a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456637-cgcolorgetalpha?language=objc
func ColorGetAlpha(color ColorRef) float64 {
	rv := C.ColorGetAlpha(
		// *typing.RefType
		unsafe.Pointer(color),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Returns the height of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455829-cgimagegetheight?language=objc
func ImageGetHeight(image ImageRef) uint {
	rv := C.ImageGetHeight(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Creates a specified type of Quartz color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408921-cgcolorspacecreatewithname?language=objc
func ColorSpaceCreateWithName(name corefoundation.StringRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateWithName(
		// *typing.RefType
		unsafe.Pointer(name),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Appends a quadratic Bézier curve from the current point, using a control point and an end point you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454268-cgcontextaddquadcurvetopoint?language=objc
func ContextAddQuadCurveToPoint(c ContextRef, cpx float64, cpy float64, x float64, y float64) {
	C.ContextAddQuadCurveToPoint(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(cpx),
		// *typing.PrimitiveType
		C.float(cpy),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
}

// Configures the display mode of a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454273-cgconfiguredisplaywithdisplaymod?language=objc
func ConfigureDisplayWithDisplayMode(config unsafe.Pointer, display DirectDisplayID, mode DisplayModeRef, options corefoundation.DictionaryRef) Error {
	rv := C.ConfigureDisplayWithDisplayMode(
		// *typing.RefType
		unsafe.Pointer(config),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.RefType
		unsafe.Pointer(mode),
		// *typing.RefType
		unsafe.Pointer(options),
	)
	// *typing.AliasType
	return Error(rv)
}

// Returns a copy of the provider’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408309-cgdataprovidercopydata?language=objc
func DataProviderCopyData(provider DataProviderRef) corefoundation.DataRef {
	rv := C.DataProviderCopyData(
		// *typing.RefType
		unsafe.Pointer(provider),
	)
	// *typing.RefType
	return corefoundation.DataRef(rv)
}

// Returns a Boolean value indicating whether a display is in a hardware mirroring set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454506-cgdisplayisinhwmirrorset?language=objc
func DisplayIsInHWMirrorSet(display DirectDisplayID) int {
	rv := C.DisplayIsInHWMirrorSet(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return int(rv)
}

// Returns the mask that indicates which classes of local hardware events are enabled during event suppression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408785-cgeventsourcegetlocaleventsfilte?language=objc
func EventSourceGetLocalEventsFilterDuringSuppressionState(source EventSourceRef, state EventSuppressionState) EventFilterMask {
	rv := C.EventSourceGetLocalEventsFilterDuringSuppressionState(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventSuppressionState)(state),
	)
	// *typing.AliasType
	return EventFilterMask(rv)
}

// Returns the font table that corresponds to the provided tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396402-cgfontcopytablefortag?language=objc
func FontCopyTableForTag(font FontRef, tag uint32) corefoundation.DataRef {
	rv := C.FontCopyTableForTag(
		// *typing.RefType
		unsafe.Pointer(font),
		// *typing.PrimitiveType
		C.uint32_t(tag),
	)
	// *typing.RefType
	return corefoundation.DataRef(rv)
}

// Returns a pointer to the bytes of a PDF string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455978-cgpdfstringgetbyteptr?language=objc
func PDFStringGetBytePtr(string_ unsafe.Pointer) string {
	rv := C.PDFStringGetBytePtr(
		// *typing.RefType
		unsafe.Pointer(string_),
	)
	// *typing.CStringType
	return C.GoString(rv)
}

// Returns the smallest rectangle that contains the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454577-cgcontextgetpathboundingbox?language=objc
func ContextGetPathBoundingBox(c ContextRef) Rect {
	rv := C.ContextGetPathBoundingBox(
		// *typing.RefType
		unsafe.Pointer(c),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Increments the retain count of a layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450900-cglayerretain?language=objc
func LayerRetain(layer LayerRef) LayerRef {
	rv := C.LayerRetain(
		// *typing.RefType
		unsafe.Pointer(layer),
	)
	// *typing.RefType
	return LayerRef(rv)
}

// Sets whether or not to allow subpixel positioning for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454942-cgcontextsetallowsfontsubpixelpo?language=objc
func ContextSetAllowsFontSubpixelPositioning(c ContextRef, allowsFontSubpixelPositioning bool) {
	C.ContextSetAllowsFontSubpixelPositioning(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.bool(allowsFontSubpixelPositioning),
	)
}

// Returns the Core Foundation type identifier for a color data type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455568-cgcolorgettypeid?language=objc
func ColorGetTypeID() corefoundation.TypeID {
	rv := C.ColorGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2962830-cgimagegetbyteorderinfo?language=objc
func ImageGetByteOrderInfo(image ImageRef) ImageByteOrderInfo {
	rv := C.ImageGetByteOrderInfo(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.AliasType
	return ImageByteOrderInfo(rv)
}

// Creates a Core Graphics PDF document using a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402603-cgpdfdocumentcreatewithprovider?language=objc
func PDFDocumentCreateWithProvider(provider DataProviderRef) PDFDocumentRef {
	rv := C.PDFDocumentCreateWithProvider(
		// *typing.RefType
		unsafe.Pointer(provider),
	)
	// *typing.RefType
	return PDFDocumentRef(rv)
}

// Gets the advance width of each glyph in the provided array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396332-cgfontgetglyphadvances?language=objc
func FontGetGlyphAdvances(font FontRef, glyphs *Glyph, count uint, advances *int) bool {
	rv := C.FontGetGlyphAdvances(
		// *typing.RefType
		unsafe.Pointer(font),
		// *typing.PointerType
		(*C.CGGlyph)(unsafe.Pointer(&glyphs)),
		// *typing.PrimitiveType
		C.uint(count),
		// *typing.PointerType
		(*C.int)(unsafe.Pointer(&advances)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2962831-cgimagegetpixelformatinfo?language=objc
func ImageGetPixelFormatInfo(image ImageRef) ImagePixelFormatInfo {
	rv := C.ImageGetPixelFormatInfo(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.AliasType
	return ImagePixelFormatInfo(rv)
}

// Returns an array of the variation axis dictionaries for a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396376-cgfontcopyvariationaxes?language=objc
func FontCopyVariationAxes(font FontRef) corefoundation.ArrayRef {
	rv := C.FontCopyVariationAxes(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.RefType
	return corefoundation.ArrayRef(rv)
}

// Returns a copy of the color space's properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2962828-cgcolorspacecopypropertylist?language=objc
func ColorSpaceCopyPropertyList(space ColorSpaceRef) corefoundation.PropertyListRef {
	rv := C.ColorSpaceCopyPropertyList(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.AliasType
	return corefoundation.PropertyListRef(rv)
}

// Sets the stroke pattern in the specified graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454796-cgcontextsetstrokepattern?language=objc
func ContextSetStrokePattern(c ContextRef, pattern PatternRef, components *float64) {
	C.ContextSetStrokePattern(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(pattern),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&components)),
	)
}

// Sets the clipping path to the intersection of the current clipping path with the area defined by the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454716-cgcontextcliptorect?language=objc
func ContextClipToRect(c ContextRef, rect Rect) {
	C.ContextClipToRect(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2881687-cgcontextresetclip?language=objc
func ContextResetClip(c ContextRef) {
	C.ContextResetClip(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Appends a quadratic Bézier curve to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411157-cgpathaddquadcurvetopoint?language=objc
func PathAddQuadCurveToPoint(path MutablePathRef, m *AffineTransform, cpx float64, cpy float64, x float64, y float64) {
	C.PathAddQuadCurveToPoint(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.PrimitiveType
		C.float(cpx),
		// *typing.PrimitiveType
		C.float(cpy),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
}

// Moves the mouse cursor without generating events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456387-cgwarpmousecursorposition?language=objc
func WarpMouseCursorPosition(newCursorPosition Point) Error {
	rv := C.WarpMouseCursorPosition(
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&newCursorPosition)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Moves the mouse cursor to a specified point relative to the upper-left corner of the display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455258-cgdisplaymovecursortopoint?language=objc
func DisplayMoveCursorToPoint(display DirectDisplayID, point Point) Error {
	rv := C.DisplayMoveCursorToPoint(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&point)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Generates and returns information about the selected windows in the current user session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455137-cgwindowlistcopywindowinfo?language=objc
func WindowListCopyWindowInfo(option WindowListOption, relativeToWindow WindowID) corefoundation.ArrayRef {
	rv := C.WindowListCopyWindowInfo(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGWindowListOption)(option),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGWindowID)(relativeToWindow),
	)
	// *typing.RefType
	return corefoundation.ArrayRef(rv)
}

// Returns the window ID of the shield window for a captured display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454279-cgshieldingwindowid?language=objc
func ShieldingWindowID(display DirectDisplayID) WindowID {
	rv := C.ShieldingWindowID(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return WindowID(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2875542-cgpdfcontextsetoutline?language=objc
func PDFContextSetOutline(context ContextRef, outline corefoundation.DictionaryRef) {
	C.PDFContextSetOutline(
		// *typing.RefType
		unsafe.Pointer(context),
		// *typing.RefType
		unsafe.Pointer(outline),
	)
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411173-cgpathaddarctopoint?language=objc
func PathAddArcToPoint(path MutablePathRef, m *AffineTransform, x1 float64, y1 float64, x2 float64, y2 float64, radius float64) {
	C.PathAddArcToPoint(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.PrimitiveType
		C.float(x1),
		// *typing.PrimitiveType
		C.float(y1),
		// *typing.PrimitiveType
		C.float(x2),
		// *typing.PrimitiveType
		C.float(y2),
		// *typing.PrimitiveType
		C.float(radius),
	)
}

// Returns the number of bits allocated for a single color component of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454980-cgimagegetbitspercomponent?language=objc
func ImageGetBitsPerComponent(image ImageRef) uint {
	rv := C.ImageGetBitsPerComponent(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Obtains exclusive use of all active displays, preventing other applications and system services from using the display or changing its configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455221-cgcapturealldisplays?language=objc
func CaptureAllDisplays() Error {
	rv := C.CaptureAllDisplays()
	// *typing.AliasType
	return Error(rv)
}

// Gets the scale of pixels per line in a scrolling event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408775-cgeventsourcegetpixelsperline?language=objc
func EventSourceGetPixelsPerLine(source EventSourceRef) float64 {
	rv := C.EventSourceGetPixelsPerLine(
		// *typing.RefType
		unsafe.Pointer(source),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Returns a rectangle that is smaller or larger than the source rectangle, with the same center point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454218-cgrectinset?language=objc
func RectInset(rect Rect, dx float64, dy float64) Rect {
	rv := C.RectInset(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.PrimitiveType
		C.float(dx),
		// *typing.PrimitiveType
		C.float(dy),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Returns whether an object at a given index in a Quartz PDF array is a PDF null. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456173-cgpdfarraygetnull?language=objc
func PDFArrayGetNull(array unsafe.Pointer, index uint) bool {
	rv := C.PDFArrayGetNull(
		// *typing.RefType
		unsafe.Pointer(array),
		// *typing.PrimitiveType
		C.uint(index),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the decode array for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454575-cgimagegetdecode?language=objc
func ImageGetDecode(image ImageRef) *float64 {
	rv := C.ImageGetDecode(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.PointerType
	return *(**float64)(unsafe.Pointer(&rv))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2866188-cgdataprovidergetinfo?language=objc
func DataProviderGetInfo(provider DataProviderRef) unsafe.Pointer {
	rv := C.DataProviderGetInfo(
		// *typing.RefType
		unsafe.Pointer(provider),
	)
	// *typing.VoidPointerType
	return unsafe.Pointer(rv)
}

// Returns the data associated with a PDF stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454657-cgpdfstreamcopydata?language=objc
func PDFStreamCopyData(stream unsafe.Pointer, format *PDFDataFormat) corefoundation.DataRef {
	rv := C.PDFStreamCopyData(
		// *typing.RefType
		unsafe.Pointer(stream),
		// *typing.PointerType
		(*C.CGPDFDataFormat)(unsafe.Pointer(&format)),
	)
	// *typing.RefType
	return corefoundation.DataRef(rv)
}

// Increments the retain count of a Core Graphics pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1552265-cgpatternretain?language=objc
func PatternRetain(pattern PatternRef) PatternRef {
	rv := C.PatternRetain(
		// *typing.RefType
		unsafe.Pointer(pattern),
	)
	// *typing.RefType
	return PatternRef(rv)
}

// Returns the color space of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454058-cgbitmapcontextgetcolorspace?language=objc
func BitmapContextGetColorSpace(context ContextRef) ColorSpaceRef {
	rv := C.BitmapContextGetColorSpace(
		// *typing.RefType
		unsafe.Pointer(context),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Returns whether a rectangle has zero width or height, or is a null rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454917-cgrectisempty?language=objc
func RectIsEmpty(rect Rect) bool {
	rv := C.RectIsEmpty(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the smallest value for the y-coordinate of the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454832-cgrectgetminy?language=objc
func RectGetMinY(rect Rect) float64 {
	rv := C.RectGetMinY(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Returns the Core Foundation type identifier for Quartz color spaces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408926-cgcolorspacegettypeid?language=objc
func ColorSpaceGetTypeID() corefoundation.TypeID {
	rv := C.ColorSpaceGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Increments the retain count of a shading object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1573767-cgshadingretain?language=objc
func ShadingRetain(shading ShadingRef) ShadingRef {
	rv := C.ShadingRetain(
		// *typing.RefType
		unsafe.Pointer(shading),
	)
	// *typing.RefType
	return ShadingRef(rv)
}

// Creates a copy of an existing color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456134-cgcolorcreatecopy?language=objc
func ColorCreateCopy(color ColorRef) ColorRef {
	rv := C.ColorCreateCopy(
		// *typing.RefType
		unsafe.Pointer(color),
	)
	// *typing.RefType
	return ColorRef(rv)
}

// Begins a transparency layer whose contents are bounded by the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454368-cgcontextbegintransparencylayerw?language=objc
func ContextBeginTransparencyLayerWithRect(c ContextRef, rect Rect, auxInfo corefoundation.DictionaryRef) {
	C.ContextBeginTransparencyLayerWithRect(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.RefType
		unsafe.Pointer(auxInfo),
	)
}

// Decrements the retain count of a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1571724-cgpdfpagerelease?language=objc
func PDFPageRelease(page PDFPageRef) {
	C.PDFPageRelease(
		// *typing.RefType
		unsafe.Pointer(page),
	)
}

// Returns the type identifier for CGImage objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455014-cgimagegettypeid?language=objc
func ImageGetTypeID() corefoundation.TypeID {
	rv := C.ImageGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Rotates the user coordinate system in a context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456228-cgcontextrotatectm?language=objc
func ContextRotateCTM(c ContextRef, angle float64) {
	C.ContextRotateCTM(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(angle),
	)
}

// Ends a transparency layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456554-cgcontextendtransparencylayer?language=objc
func ContextEndTransparencyLayer(c ContextRef) {
	C.ContextEndTransparencyLayer(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Sets the interval that local hardware events may be suppressed following the posting of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408783-cgeventsourcesetlocaleventssuppr?language=objc
func EventSourceSetLocalEventsSuppressionInterval(source EventSourceRef, seconds corefoundation.TimeInterval) {
	C.EventSourceSetLocalEventsSuppressionInterval(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFTimeInterval)(seconds),
	)
}

// Returns the intersection of two rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455346-cgrectintersection?language=objc
func RectIntersection(r1 Rect, r2 Rect) Rect {
	rv := C.RectIntersection(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&r1)),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&r2)),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Returns the I/O Kit display mode ID of the specified display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454728-cgdisplaymodegetiodisplaymodeid?language=objc
func DisplayModeGetIODisplayModeID(mode DisplayModeRef) int32 {
	rv := C.DisplayModeGetIODisplayModeID(
		// *typing.RefType
		unsafe.Pointer(mode),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Returns the primary display in a hardware mirroring set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454403-cgdisplayprimarydisplay?language=objc
func DisplayPrimaryDisplay(display DirectDisplayID) DirectDisplayID {
	rv := C.DisplayPrimaryDisplay(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return DirectDisplayID(rv)
}

// Sets a destination to jump to when a point in the current page of a PDF graphics context is clicked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455424-cgpdfcontextadddestinationatpoin?language=objc
func PDFContextAddDestinationAtPoint(context ContextRef, name corefoundation.StringRef, point Point) {
	C.PDFContextAddDestinationAtPoint(
		// *typing.RefType
		unsafe.Pointer(context),
		// *typing.RefType
		unsafe.Pointer(name),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&point)),
	)
}

// Returns the descent of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396351-cgfontgetdescent?language=objc
func FontGetDescent(font FontRef) int {
	rv := C.FontGetDescent(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.PrimitiveType
	return int(rv)
}

// Sets the current stroke color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456283-cgcontextsetstrokecolor?language=objc
func ContextSetStrokeColor(c ContextRef, components *float64) {
	C.ContextSetStrokeColor(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&components)),
	)
}

// Releases a Core Graphics display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1562069-cgdisplaymoderelease?language=objc
func DisplayModeRelease(mode DisplayModeRef) {
	C.DisplayModeRelease(
		// *typing.RefType
		unsafe.Pointer(mode),
	)
}

// Adds a set of rectangular paths to the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454734-cgcontextaddrects?language=objc
func ContextAddRects(c ContextRef, rects *Rect, count uint) {
	C.ContextAddRects(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&rects)),
		// *typing.PrimitiveType
		C.uint(count),
	)
}

// Creates a mutable copy of a graphics path transformed by a transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411150-cgpathcreatemutablecopybytransfo?language=objc
func PathCreateMutableCopyByTransformingPath(path unsafe.Pointer, transform *AffineTransform) MutablePathRef {
	rv := C.PathCreateMutableCopyByTransformingPath(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&transform)),
	)
	// *typing.RefType
	return MutablePathRef(rv)
}

// Returns an affine transformation matrix constructed from translation values you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454909-cgaffinetransformmaketranslation?language=objc
func AffineTransformMakeTranslation(tx float64, ty float64) AffineTransform {
	rv := C.AffineTransformMakeTranslation(
		// *typing.PrimitiveType
		C.float(tx),
		// *typing.PrimitiveType
		C.float(ty),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Sets the current character spacing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454786-cgcontextsetcharacterspacing?language=objc
func ContextSetCharacterSpacing(c ContextRef, spacing float64) {
	C.ContextSetCharacterSpacing(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(spacing),
	)
}

// Returns the I/O Kit flags of the specified display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454092-cgdisplaymodegetioflags?language=objc
func DisplayModeGetIOFlags(mode DisplayModeRef) uint32 {
	rv := C.DisplayModeGetIOFlags(
		// *typing.RefType
		unsafe.Pointer(mode),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Closes a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454306-cgpdfcontextclose?language=objc
func PDFContextClose(context ContextRef) {
	C.PDFContextClose(
		// *typing.RefType
		unsafe.Pointer(context),
	)
}

// Increments the retain count of a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408276-cgdataproviderretain?language=objc
func DataProviderRetain(provider DataProviderRef) DataProviderRef {
	rv := C.DataProviderRetain(
		// *typing.RefType
		unsafe.Pointer(provider),
	)
	// *typing.RefType
	return DataProviderRef(rv)
}

// Returns a Boolean indicating whether the color space can be used as a destination color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1690958-cgcolorspacesupportsoutput?language=objc
func ColorSpaceSupportsOutput(space ColorSpaceRef) bool {
	rv := C.ColorSpaceSupportsOutput(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Sets the floating-point value of a field in a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455526-cgeventsetdoublevaluefield?language=objc
func EventSetDoubleValueField(event EventRef, field EventField, value float64) {
	C.EventSetDoubleValueField(
		// *typing.RefType
		unsafe.Pointer(event),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventField)(field),
		// *typing.PrimitiveType
		C.double(value),
	)
}

// Draws the content of a PDF page into the current graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456255-cgcontextdrawpdfpage?language=objc
func ContextDrawPDFPage(c ContextRef, page PDFPageRef) {
	C.ContextDrawPDFPage(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(page),
	)
}

// Creates a URL-based PDF graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456290-cgpdfcontextcreatewithurl?language=objc
func PDFContextCreateWithURL(url corefoundation.URLRef, mediaBox *Rect, auxiliaryInfo corefoundation.DictionaryRef) ContextRef {
	rv := C.PDFContextCreateWithURL(
		// *typing.RefType
		unsafe.Pointer(url),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&mediaBox)),
		// *typing.RefType
		unsafe.Pointer(auxiliaryInfo),
	)
	// *typing.RefType
	return ContextRef(rv)
}

// Returns the CFType ID for PDF page objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454478-cgpdfpagegettypeid?language=objc
func PDFPageGetTypeID() corefoundation.TypeID {
	rv := C.PDFPageGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Creates a content stream object from a PDF page object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410047-cgpdfcontentstreamcreatewithpage?language=objc
func PDFContentStreamCreateWithPage(page PDFPageRef) unsafe.Pointer {
	rv := C.PDFContentStreamCreateWithPage(
		// *typing.RefType
		unsafe.Pointer(page),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Creates a font object from data supplied from a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396367-cgfontcreatewithdataprovider?language=objc
func FontCreateWithDataProvider(provider DataProviderRef) FontRef {
	rv := C.FontCreateWithDataProvider(
		// *typing.RefType
		unsafe.Pointer(provider),
	)
	// *typing.RefType
	return FontRef(rv)
}

// Returns whether there is a PDF number associated with a specified key in a PDF dictionary and, if so, retrieves that number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430228-cgpdfdictionarygetnumber?language=objc
func PDFDictionaryGetNumber(dict unsafe.Pointer, key string, value *PDFReal) bool {
	keyVal := C.CString(key)
	defer C.free(unsafe.Pointer(keyVal))
	rv := C.PDFDictionaryGetNumber(
		// *typing.RefType
		unsafe.Pointer(dict),
		// *typing.CStringType
		keyVal,
		// *typing.PointerType
		(*C.CGPDFReal)(unsafe.Pointer(&value)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Decrements the retain count of a color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408855-cgcolorspacerelease?language=objc
func ColorSpaceRelease(space ColorSpaceRef) {
	C.ColorSpaceRelease(
		// *typing.RefType
		unsafe.Pointer(space),
	)
}

// Creates a copy of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455615-cgimagecreatecopy?language=objc
func ImageCreateCopy(image ImageRef) ImageRef {
	rv := C.ImageCreateCopy(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Ends the current page in a page-based graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455027-cgcontextendpage?language=objc
func ContextEndPage(c ContextRef) {
	C.ContextEndPage(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Gets the file identifier for a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402600-cgpdfdocumentgetid?language=objc
func PDFDocumentGetID(document PDFDocumentRef) unsafe.Pointer {
	rv := C.PDFDocumentGetID(
		// *typing.RefType
		unsafe.Pointer(document),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Sets the current text matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455611-cgcontextsettextmatrix?language=objc
func ContextSetTextMatrix(c ContextRef, t AffineTransform) {
	C.ContextSetTextMatrix(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t)),
	)
}

// Returns the number of glyphs in a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396371-cgfontgetnumberofglyphs?language=objc
func FontGetNumberOfGlyphs(font FontRef) uint {
	rv := C.FontGetNumberOfGlyphs(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Maps a mask into the specified rectangle and intersects it with the current clipping area of the graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456497-cgcontextcliptomask?language=objc
func ContextClipToMask(c ContextRef, rect Rect, mask ImageRef) {
	C.ContextClipToMask(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.RefType
		unsafe.Pointer(mask),
	)
}

// Returns the dictionary of a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455125-cgpdfpagegetdictionary?language=objc
func PDFPageGetDictionary(page PDFPageRef) unsafe.Pointer {
	rv := C.PDFPageGetDictionary(
		// *typing.RefType
		unsafe.Pointer(page),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns the capacity, or number of entries, in the gamma table for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454310-cgdisplaygammatablecapacity?language=objc
func DisplayGammaTableCapacity(display DirectDisplayID) uint32 {
	rv := C.DisplayGammaTableCapacity(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Sets the integer value of a field in a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455556-cgeventsetintegervaluefield?language=objc
func EventSetIntegerValueField(event EventRef, field EventField, value int64) {
	C.EventSetIntegerValueField(
		// *typing.RefType
		unsafe.Pointer(event),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventField)(field),
		// *typing.PrimitiveType
		C.int64_t(value),
	)
}

// Adds a previously created path object to the current path in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456628-cgcontextaddpath?language=objc
func ContextAddPath(c ContextRef, path unsafe.Pointer) {
	C.ContextAddPath(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(path),
	)
}

// Returns the current text matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456154-cgcontextgettextmatrix?language=objc
func ContextGetTextMatrix(c ContextRef) AffineTransform {
	rv := C.ContextGetTextMatrix(
		// *typing.RefType
		unsafe.Pointer(c),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Paints a line along the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454490-cgcontextstrokepath?language=objc
func ContextStrokePath(c ContextRef) {
	C.ContextStrokePath(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Returns the display width in pixel units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456361-cgdisplaypixelswide?language=objc
func DisplayPixelsWide(display DirectDisplayID) uint {
	rv := C.DisplayPixelsWide(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Sets antialiasing on or off for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455178-cgcontextsetshouldantialias?language=objc
func ContextSetShouldAntialias(c ContextRef, shouldAntialias bool) {
	C.ContextSetShouldAntialias(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.bool(shouldAntialias),
	)
}

// Returns a Boolean value indicating whether an event tap is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456102-cgeventtapisenabled?language=objc
func EventTapIsEnabled(tap corefoundation.MachPortRef) bool {
	rv := C.EventTapIsEnabled(
		// *typing.RefType
		unsafe.Pointer(tap),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3042409-cgpdftagtypegetname?language=objc
func PDFTagTypeGetName(tagType PDFTagType) string {
	rv := C.PDFTagTypeGetName(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGPDFTagType)(tagType),
	)
	// *typing.CStringType
	return C.GoString(rv)
}

// Sets the current fill color to a value in the DeviceGray color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454255-cgcontextsetgrayfillcolor?language=objc
func ContextSetGrayFillColor(c ContextRef, gray float64, alpha float64) {
	C.ContextSetGrayFillColor(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(gray),
		// *typing.PrimitiveType
		C.float(alpha),
	)
}

// Gets the specified resource from a PDF content stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410044-cgpdfcontentstreamgetresource?language=objc
func PDFContentStreamGetResource(cs unsafe.Pointer, category string, name string) unsafe.Pointer {
	categoryVal := C.CString(category)
	defer C.free(unsafe.Pointer(categoryVal))
	nameVal := C.CString(name)
	defer C.free(unsafe.Pointer(nameVal))
	rv := C.PDFContentStreamGetResource(
		// *typing.RefType
		unsafe.Pointer(cs),
		// *typing.CStringType
		categoryVal,
		// *typing.CStringType
		nameVal,
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns whether there is a PDF Boolean value associated with a specified key in a PDF dictionary and, if so, retrieves the Boolean value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430226-cgpdfdictionarygetboolean?language=objc
func PDFDictionaryGetBoolean(dict unsafe.Pointer, key string, value *PDFBoolean) bool {
	keyVal := C.CString(key)
	defer C.free(unsafe.Pointer(keyVal))
	rv := C.PDFDictionaryGetBoolean(
		// *typing.RefType
		unsafe.Pointer(dict),
		// *typing.CStringType
		keyVal,
		// *typing.PointerType
		(*C.CGPDFBoolean)(unsafe.Pointer(&value)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Uses a PostScript converter to convert PostScript data to PDF data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455368-cgpsconverterconvert?language=objc
func PSConverterConvert(converter PSConverterRef, provider DataProviderRef, consumer DataConsumerRef, options corefoundation.DictionaryRef) bool {
	rv := C.PSConverterConvert(
		// *typing.RefType
		unsafe.Pointer(converter),
		// *typing.RefType
		unsafe.Pointer(provider),
		// *typing.RefType
		unsafe.Pointer(consumer),
		// *typing.RefType
		unsafe.Pointer(options),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Decrements the retain count of a PDF content stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410046-cgpdfcontentstreamrelease?language=objc
func PDFContentStreamRelease(cs unsafe.Pointer) {
	C.PDFContentStreamRelease(
		// *typing.RefType
		unsafe.Pointer(cs),
	)
}

// Maps a display ID to an OpenGL display mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455554-cgdisplayidtoopengldisplaymask?language=objc
func DisplayIDToOpenGLDisplayMask(display DirectDisplayID) OpenGLDisplayMask {
	rv := C.DisplayIDToOpenGLDisplayMask(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return OpenGLDisplayMask(rv)
}

// Returns the largest value of the x-coordinate for the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454334-cgrectgetmaxx?language=objc
func RectGetMaxX(rect Rect) float64 {
	rv := C.RectGetMaxX(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Indicates whether the current path contains any subpaths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455772-cgcontextispathempty?language=objc
func ContextIsPathEmpty(c ContextRef) bool {
	rv := C.ContextIsPathEmpty(
		// *typing.RefType
		unsafe.Pointer(c),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the Core Foundation type identifier for a color conversion info data type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2113681-cgcolorconversioninfogettypeid?language=objc
func ColorConversionInfoGetTypeID() corefoundation.TypeID {
	rv := C.ColorConversionInfoGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Returns the number of bits allocated for a single pixel in a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454599-cgimagegetbitsperpixel?language=objc
func ImageGetBitsPerPixel(image ImageRef) uint {
	rv := C.ImageGetBitsPerPixel(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Fills in a rectangle using the contents of the specified dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456558-cgrectmakewithdictionaryrepresen?language=objc
func RectMakeWithDictionaryRepresentation(dict corefoundation.DictionaryRef, rect *Rect) bool {
	rv := C.RectMakeWithDictionaryRepresentation(
		// *typing.RefType
		unsafe.Pointer(dict),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456406-cgdisplaymodegetpixelheight?language=objc
func DisplayModeGetPixelHeight(mode DisplayModeRef) uint {
	rv := C.DisplayModeGetPixelHeight(
		// *typing.RefType
		unsafe.Pointer(mode),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Creates a device-dependent RGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408837-cgcolorspacecreatedevicergb?language=objc
func ColorSpaceCreateDeviceRGB() ColorSpaceRef {
	rv := C.ColorSpaceCreateDeviceRGB()
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Sets the current stroke color in a context, using a CGColor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456196-cgcontextsetstrokecolorwithcolor?language=objc
func ContextSetStrokeColorWithColor(c ContextRef, color ColorRef) {
	C.ContextSetStrokeColorWithColor(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(color),
	)
}

// Returns the graphics context associated with a layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450902-cglayergetcontext?language=objc
func LayerGetContext(layer LayerRef) ContextRef {
	rv := C.LayerGetContext(
		// *typing.RefType
		unsafe.Pointer(layer),
	)
	// *typing.RefType
	return ContextRef(rv)
}

// Returns whether an object at a given index in a PDF array is a PDF integer and, if so, retrieves that object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456053-cgpdfarraygetinteger?language=objc
func PDFArrayGetInteger(array unsafe.Pointer, index uint, value *PDFInteger) bool {
	rv := C.PDFArrayGetInteger(
		// *typing.RefType
		unsafe.Pointer(array),
		// *typing.PrimitiveType
		C.uint(index),
		// *typing.PointerType
		(*C.CGPDFInteger)(unsafe.Pointer(&value)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the italic angle of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396404-cgfontgetitalicangle?language=objc
func FontGetItalicAngle(font FontRef) float64 {
	rv := C.FontGetItalicAngle(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Configures the origin of a display relative to the global display coordinate space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454090-cgconfiguredisplayorigin?language=objc
func ConfigureDisplayOrigin(config unsafe.Pointer, display DirectDisplayID, x int32, y int32) Error {
	rv := C.ConfigureDisplayOrigin(
		// *typing.RefType
		unsafe.Pointer(config),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.PrimitiveType
		C.int32_t(x),
		// *typing.PrimitiveType
		C.int32_t(y),
	)
	// *typing.AliasType
	return Error(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3227891-cgcolorconversioninfocreatewitho?language=objc
func ColorConversionInfoCreateWithOptions(src ColorSpaceRef, dst ColorSpaceRef, options corefoundation.DictionaryRef) ColorConversionInfoRef {
	rv := C.ColorConversionInfoCreateWithOptions(
		// *typing.RefType
		unsafe.Pointer(src),
		// *typing.RefType
		unsafe.Pointer(dst),
		// *typing.RefType
		unsafe.Pointer(options),
	)
	// *typing.RefType
	return ColorConversionInfoRef(rv)
}

// Restores the permanent display configuration settings for the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456402-cgrestorepermanentdisplayconfigu?language=objc
func RestorePermanentDisplayConfiguration() {
	C.RestorePermanentDisplayConfiguration()
}

// Returns the alpha information associated with the context, which indicates how a bitmap context handles the alpha component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454960-cgbitmapcontextgetalphainfo?language=objc
func BitmapContextGetAlphaInfo(context ContextRef) ImageAlphaInfo {
	rv := C.BitmapContextGetAlphaInfo(
		// *typing.RefType
		unsafe.Pointer(context),
	)
	// *typing.AliasType
	return ImageAlphaInfo(rv)
}

// Closes and terminates the current path’s subpath. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454508-cgcontextclosepath?language=objc
func ContextClosePath(c ContextRef) {
	C.ContextClosePath(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Indicates whether two graphics paths are equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411167-cgpathequaltopath?language=objc
func PathEqualToPath(path1 unsafe.Pointer, path2 unsafe.Pointer) bool {
	rv := C.PathEqualToPath(
		// *typing.RefType
		unsafe.Pointer(path1),
		// *typing.RefType
		unsafe.Pointer(path2),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a new Quartz mouse event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454356-cgeventcreatemouseevent?language=objc
func EventCreateMouseEvent(source EventSourceRef, mouseType EventType, mouseCursorPosition Point, mouseButton MouseButton) EventRef {
	rv := C.EventCreateMouseEvent(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventType)(mouseType),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&mouseCursorPosition)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGMouseButton)(mouseButton),
	)
	// *typing.RefType
	return EventRef(rv)
}

// Returns the display ID of the main display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455620-cgmaindisplayid?language=objc
func MainDisplayID() DirectDisplayID {
	rv := C.MainDisplayID()
	// *typing.AliasType
	return DirectDisplayID(rv)
}

// Returns an affine transformation matrix constructed by translating an existing affine transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455822-cgaffinetransformtranslate?language=objc
func AffineTransformTranslate(t AffineTransform, tx float64, ty float64) AffineTransform {
	rv := C.AffineTransformTranslate(
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t)),
		// *typing.PrimitiveType
		C.float(tx),
		// *typing.PrimitiveType
		C.float(ty),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Returns a page from a Core Graphics PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402586-cgpdfdocumentgetpage?language=objc
func PDFDocumentGetPage(document PDFDocumentRef, pageNumber uint) PDFPageRef {
	rv := C.PDFDocumentGetPage(
		// *typing.RefType
		unsafe.Pointer(document),
		// *typing.PrimitiveType
		C.uint(pageNumber),
	)
	// *typing.RefType
	return PDFPageRef(rv)
}

// Returns the x- coordinate that establishes the center of a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456175-cgrectgetmidx?language=objc
func RectGetMidX(rect Rect) float64 {
	rv := C.RectGetMidX(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Returns the keyboard type to be used with a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408787-cgeventsourcegetkeyboardtype?language=objc
func EventSourceGetKeyboardType(source EventSourceRef) EventSourceKeyboardType {
	rv := C.EventSourceGetKeyboardType(
		// *typing.RefType
		unsafe.Pointer(source),
	)
	// *typing.AliasType
	return EventSourceKeyboardType(rv)
}

// Begins a new page in a PDF graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456578-cgpdfcontextbeginpage?language=objc
func PDFContextBeginPage(context ContextRef, pageInfo corefoundation.DictionaryRef) {
	C.PDFContextBeginPage(
		// *typing.RefType
		unsafe.Pointer(context),
		// *typing.RefType
		unsafe.Pointer(pageInfo),
	)
}

// Creates a copy of an existing color, substituting a new alpha value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455986-cgcolorcreatecopywithalpha?language=objc
func ColorCreateCopyWithAlpha(color ColorRef, alpha float64) ColorRef {
	rv := C.ColorCreateCopyWithAlpha(
		// *typing.RefType
		unsafe.Pointer(color),
		// *typing.PrimitiveType
		C.float(alpha),
	)
	// *typing.RefType
	return ColorRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2875137-cgpdfdocumentgetaccesspermission?language=objc
func PDFDocumentGetAccessPermissions(document PDFDocumentRef) PDFAccessPermissions {
	rv := C.PDFDocumentGetAccessPermissions(
		// *typing.RefType
		unsafe.Pointer(document),
	)
	// *typing.AliasType
	return PDFAccessPermissions(rv)
}

// Returns an affine transformation matrix constructed by rotating an existing affine transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455962-cgaffinetransformrotate?language=objc
func AffineTransformRotate(t AffineTransform, angle float64) AffineTransform {
	rv := C.AffineTransformRotate(
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t)),
		// *typing.PrimitiveType
		C.float(angle),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Returns a copy of an existing Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454071-cgeventcreatecopy?language=objc
func EventCreateCopy(event EventRef) EventRef {
	rv := C.EventCreateCopy(
		// *typing.RefType
		unsafe.Pointer(event),
	)
	// *typing.RefType
	return EventRef(rv)
}

// Marks a window context for update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455450-cgcontextsynchronize?language=objc
func ContextSynchronize(c ContextRef) {
	C.ContextSynchronize(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Returns the affine transform that maps a box to a given rectangle on a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454893-cgpdfpagegetdrawingtransform?language=objc
func PDFPageGetDrawingTransform(page PDFPageRef, box PDFBox, rect Rect, rotate int, preserveAspectRatio bool) AffineTransform {
	rv := C.PDFPageGetDrawingTransform(
		// *typing.RefType
		unsafe.Pointer(page),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGPDFBox)(box),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.PrimitiveType
		C.int(rotate),
		// *typing.PrimitiveType
		C.bool(preserveAspectRatio),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Returns the window level of the shield window for a captured display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456352-cgshieldingwindowlevel?language=objc
func ShieldingWindowLevel() WindowLevel {
	rv := C.ShieldingWindowLevel()
	// *typing.AliasType
	return WindowLevel(rv)
}

// Returns a Boolean value indicating whether the specified display mode is usable for a desktop graphical user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454928-cgdisplaymodeisusablefordesktopg?language=objc
func DisplayModeIsUsableForDesktopGUI(mode DisplayModeRef) bool {
	rv := C.DisplayModeIsUsableForDesktopGUI(
		// *typing.RefType
		unsafe.Pointer(mode),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the Core Foundation type identifier for CGGradient objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1398453-cggradientgettypeid?language=objc
func GradientGetTypeID() corefoundation.TypeID {
	rv := C.GradientGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Returns a dictionary representation of the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455382-cgpointcreatedictionaryrepresent?language=objc
func PointCreateDictionaryRepresentation(point Point) corefoundation.DictionaryRef {
	rv := C.PointCreateDictionaryRepresentation(
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&point)),
	)
	// *typing.RefType
	return corefoundation.DictionaryRef(rv)
}

// Returns a rectangle with an origin that is offset from that of the source rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454841-cgrectoffset?language=objc
func RectOffset(rect Rect, dx float64, dy float64) Rect {
	rv := C.RectOffset(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.PrimitiveType
		C.float(dx),
		// *typing.PrimitiveType
		C.float(dy),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Sets the current stroke color to a value in the DeviceCMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455358-cgcontextsetcmykstrokecolor?language=objc
func ContextSetCMYKStrokeColor(c ContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64) {
	C.ContextSetCMYKStrokeColor(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(cyan),
		// *typing.PrimitiveType
		C.float(magenta),
		// *typing.PrimitiveType
		C.float(yellow),
		// *typing.PrimitiveType
		C.float(black),
		// *typing.PrimitiveType
		C.float(alpha),
	)
}

// Returns the base color space of a pattern or indexed color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408839-cgcolorspacegetbasecolorspace?language=objc
func ColorSpaceGetBaseColorSpace(space ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceGetBaseColorSpace(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Creates a layer object that is associated with a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450892-cglayercreatewithcontext?language=objc
func LayerCreateWithContext(context ContextRef, size Size, auxiliaryInfo corefoundation.DictionaryRef) LayerRef {
	rv := C.LayerCreateWithContext(
		// *typing.RefType
		unsafe.Pointer(context),
		// *typing.StructType
		*(*C.CGSize)(unsafe.Pointer(&size)),
		// *typing.RefType
		unsafe.Pointer(auxiliaryInfo),
	)
	// *typing.RefType
	return LayerRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454756-cgdisplaymodegetpixelwidth?language=objc
func DisplayModeGetPixelWidth(mode DisplayModeRef) uint {
	rv := C.DisplayModeGetPixelWidth(
		// *typing.RefType
		unsafe.Pointer(mode),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Creates a bitmap image by masking an existing bitmap image with the provided color values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454358-cgimagecreatewithmaskingcolors?language=objc
func ImageCreateWithMaskingColors(image ImageRef, components *float64) ImageRef {
	rv := C.ImageCreateWithMaskingColors(
		// *typing.RefType
		unsafe.Pointer(image),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&components)),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Paints the areas contained within the provided rectangles, using the fill color in the current graphics state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454132-cgcontextfillrects?language=objc
func ContextFillRects(c ContextRef, rects *Rect, count uint) {
	C.ContextFillRects(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&rects)),
		// *typing.PrimitiveType
		C.uint(count),
	)
}

// Obtains the PostScript name of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396346-cgfontcopypostscriptname?language=objc
func FontCopyPostScriptName(font FontRef) corefoundation.StringRef {
	rv := C.FontCopyPostScriptName(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.RefType
	return corefoundation.StringRef(rv)
}

// Returns whether a bitmap image is an image mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454229-cgimageismask?language=objc
func ImageIsMask(image ImageRef) bool {
	rv := C.ImageIsMask(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Sets the stroke color space in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454396-cgcontextsetstrokecolorspace?language=objc
func ContextSetStrokeColorSpace(c ContextRef, space ColorSpaceRef) {
	C.ContextSetStrokeColorSpace(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(space),
	)
}

// Returns a rectangle with a positive width and height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456432-cgrectstandardize?language=objc
func RectStandardize(rect Rect) Rect {
	rv := C.RectStandardize(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Returns the bounding box of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396353-cgfontgetfontbbox?language=objc
func FontGetFontBBox(font FontRef) Rect {
	rv := C.FontGetFontBBox(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Obtains exclusive use of a display, preventing other applications and system services from using the display or changing its configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456259-cgdisplaycapture?language=objc
func DisplayCapture(display DirectDisplayID) Error {
	rv := C.DisplayCapture(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return Error(rv)
}

// Adds an ellipse that fits inside the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456420-cgcontextaddellipseinrect?language=objc
func ContextAddEllipseInRect(c ContextRef, rect Rect) {
	C.ContextAddEllipseInRect(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

// Creates a CGGradient object from a color space and the provided color components and locations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1398454-cggradientcreatewithcolorcompone?language=objc
func GradientCreateWithColorComponents(space ColorSpaceRef, components *float64, locations *float64, count uint) GradientRef {
	rv := C.GradientCreateWithColorComponents(
		// *typing.RefType
		unsafe.Pointer(space),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&components)),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&locations)),
		// *typing.PrimitiveType
		C.uint(count),
	)
	// *typing.RefType
	return GradientRef(rv)
}

// Creates a calibrated RGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408861-cgcolorspacecreatecalibratedrgb?language=objc
func ColorSpaceCreateCalibratedRGB(whitePoint *float64, blackPoint *float64, gamma *float64, matrix *float64) ColorSpaceRef {
	rv := C.ColorSpaceCreateCalibratedRGB(
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&whitePoint)),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&blackPoint)),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&gamma)),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&matrix)),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Fills the clipping path of a context with the specified shading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456643-cgcontextdrawshading?language=objc
func ContextDrawShading(c ContextRef, shading ShadingRef) {
	C.ContextDrawShading(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(shading),
	)
}

// Returns the y-coordinate that establishes the center of the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456550-cgrectgetmidy?language=objc
func RectGetMidY(rect Rect) float64 {
	rv := C.RectGetMidY(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Returns an image containing the contents of a portion of the specified display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454595-cgdisplaycreateimageforrect?language=objc
func DisplayCreateImageForRect(display DirectDisplayID, rect Rect) ImageRef {
	rv := C.DisplayCreateImageForRect(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Creates a color using a list of intensity values (including alpha) and an associated color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455927-cgcolorcreate?language=objc
func ColorCreate(space ColorSpaceRef, components *float64) ColorRef {
	rv := C.ColorCreate(
		// *typing.RefType
		unsafe.Pointer(space),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&components)),
	)
	// *typing.RefType
	return ColorRef(rv)
}

// Starts a new subpath at a specified location in a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411146-cgpathmovetopoint?language=objc
func PathMoveToPoint(path MutablePathRef, m *AffineTransform, x float64, y float64) {
	C.PathMoveToPoint(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
}

// Decrements the retain count of a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586340-cgcolorrelease?language=objc
func ColorRelease(color ColorRef) {
	C.ColorRelease(
		// *typing.RefType
		unsafe.Pointer(color),
	)
}

// Returns the number of pages in a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402595-cgpdfdocumentgetnumberofpages?language=objc
func PDFDocumentGetNumberOfPages(document PDFDocumentRef) uint {
	rv := C.PDFDocumentGetNumberOfPages(
		// *typing.RefType
		unsafe.Pointer(document),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Sets the keyboard type to be used with a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408795-cgeventsourcesetkeyboardtype?language=objc
func EventSourceSetKeyboardType(source EventSourceRef, keyboardType EventSourceKeyboardType) {
	C.EventSourceSetKeyboardType(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventSourceKeyboardType)(keyboardType),
	)
}

// Returns information about the currently available display modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455537-cgdisplaycopyalldisplaymodes?language=objc
func DisplayCopyAllDisplayModes(display DirectDisplayID, options corefoundation.DictionaryRef) corefoundation.ArrayRef {
	rv := C.DisplayCopyAllDisplayModes(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.RefType
		unsafe.Pointer(options),
	)
	// *typing.RefType
	return corefoundation.ArrayRef(rv)
}

// Returns the display height in pixel units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454247-cgdisplaypixelshigh?language=objc
func DisplayPixelsHigh(display DirectDisplayID) uint {
	rv := C.DisplayPixelsHigh(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Sets the byte values in the 8-bit RGB gamma tables for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455896-cgsetdisplaytransferbybytetable?language=objc
func SetDisplayTransferByByteTable(display DirectDisplayID, tableSize uint32, redTable *uint8, greenTable *uint8, blueTable *uint8) Error {
	rv := C.SetDisplayTransferByByteTable(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.PrimitiveType
		C.uint32_t(tableSize),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&redTable)),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&greenTable)),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&blueTable)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Returns the data provider for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455260-cgimagegetdataprovider?language=objc
func ImageGetDataProvider(image ImageRef) DataProviderRef {
	rv := C.ImageGetDataProvider(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.RefType
	return DataProviderRef(rv)
}

// Sets the fill pattern in the specified graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456334-cgcontextsetfillpattern?language=objc
func ContextSetFillPattern(c ContextRef, pattern PatternRef, components *float64) {
	C.ContextSetFillPattern(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(pattern),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&components)),
	)
}

// Returns the Core Foundation type identifier for data providers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408290-cgdataprovidergettypeid?language=objc
func DataProviderGetTypeID() corefoundation.TypeID {
	rv := C.DataProviderGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Creates a direct-access data provider that uses a file to supply data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408294-cgdataprovidercreatewithfilename?language=objc
func DataProviderCreateWithFilename(filename string) DataProviderRef {
	filenameVal := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameVal))
	rv := C.DataProviderCreateWithFilename(
		// *typing.CStringType
		filenameVal,
	)
	// *typing.RefType
	return DataProviderRef(rv)
}

// Returns the number of bytes allocated for a single row of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455425-cgimagegetbytesperrow?language=objc
func ImageGetBytesPerRow(image ImageRef) uint {
	rv := C.ImageGetBytesPerRow(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Creates a data provider that reads from a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408284-cgdataprovidercreatewithcfdata?language=objc
func DataProviderCreateWithCFData(data corefoundation.DataRef) DataProviderRef {
	rv := C.DataProviderCreateWithCFData(
		// *typing.RefType
		unsafe.Pointer(data),
	)
	// *typing.RefType
	return DataProviderRef(rv)
}

// Returns whether the specified PDF document allows copying. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402588-cgpdfdocumentallowscopying?language=objc
func PDFDocumentAllowsCopying(document PDFDocumentRef) bool {
	rv := C.PDFDocumentAllowsCopying(
		// *typing.RefType
		unsafe.Pointer(document),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a new color in a different color space that matches the provided color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455493-cgcolorcreatecopybymatchingtocol?language=objc
func ColorCreateCopyByMatchingToColorSpace(arg0 ColorSpaceRef, intent ColorRenderingIntent, color ColorRef, options corefoundation.DictionaryRef) ColorRef {
	rv := C.ColorCreateCopyByMatchingToColorSpace(
		// *typing.RefType
		unsafe.Pointer(arg0),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGColorRenderingIntent)(intent),
		// *typing.RefType
		unsafe.Pointer(color),
		// *typing.RefType
		unsafe.Pointer(options),
	)
	// *typing.RefType
	return ColorRef(rv)
}

// Returns a size with the specified dimension values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455082-cgsizemake?language=objc
func SizeMake(width float64, height float64) Size {
	rv := C.SizeMake(
		// *typing.PrimitiveType
		C.float(width),
		// *typing.PrimitiveType
		C.float(height),
	)
	// *typing.StructType
	return *(*Size)(unsafe.Pointer(&rv))
}

// Returns the largest value for the y-coordinate of the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454060-cgrectgetmaxy?language=objc
func RectGetMaxY(rect Rect) float64 {
	rv := C.RectGetMaxY(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Returns whether a PDF document allows printing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402594-cgpdfdocumentallowsprinting?language=objc
func PDFDocumentAllowsPrinting(document PDFDocumentRef) bool {
	rv := C.PDFDocumentAllowsPrinting(
		// *typing.RefType
		unsafe.Pointer(document),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Sets whether or not to allow font smoothing for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454767-cgcontextsetallowsfontsmoothing?language=objc
func ContextSetAllowsFontSmoothing(c ContextRef, allowsFontSmoothing bool) {
	C.ContextSetAllowsFontSmoothing(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.bool(allowsFontSmoothing),
	)
}

// Draws the current path using the provided drawing mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455195-cgcontextdrawpath?language=objc
func ContextDrawPath(c ContextRef, mode PathDrawingMode) {
	C.ContextDrawPath(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGPathDrawingMode)(mode),
	)
}

// Returns the glyph for the glyph name associated with the specified font object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396340-cgfontgetglyphwithglyphname?language=objc
func FontGetGlyphWithGlyphName(font FontRef, name corefoundation.StringRef) Glyph {
	rv := C.FontGetGlyphWithGlyphName(
		// *typing.RefType
		unsafe.Pointer(font),
		// *typing.RefType
		unsafe.Pointer(name),
	)
	// *typing.AliasType
	return Glyph(rv)
}

// Returns a dictionary representation of the provided rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455760-cgrectcreatedictionaryrepresenta?language=objc
func RectCreateDictionaryRepresentation(arg0 Rect) corefoundation.DictionaryRef {
	rv := C.RectCreateDictionaryRepresentation(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&arg0)),
	)
	// *typing.RefType
	return corefoundation.DictionaryRef(rv)
}

// Returns a pointer to the image data associated with a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455517-cgbitmapcontextgetdata?language=objc
func BitmapContextGetData(context ContextRef) unsafe.Pointer {
	rv := C.BitmapContextGetData(
		// *typing.RefType
		unsafe.Pointer(context),
	)
	// *typing.VoidPointerType
	return unsafe.Pointer(rv)
}

// Returns the ascent of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396359-cgfontgetascent?language=objc
func FontGetAscent(font FontRef) int {
	rv := C.FontGetAscent(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.PrimitiveType
	return int(rv)
}

// Immediately enables or disables stereo operation for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454316-cgdisplaysetstereooperation?language=objc
func DisplaySetStereoOperation(display DirectDisplayID, stereo int, forceBlueLine int, option ConfigureOption) Error {
	rv := C.DisplaySetStereoOperation(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.boolean_t)(stereo),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.boolean_t)(forceBlueLine),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGConfigureOption)(option),
	)
	// *typing.AliasType
	return Error(rv)
}

// Returns the interpolation setting for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455363-cgimagegetshouldinterpolate?language=objc
func ImageGetShouldInterpolate(image ImageRef) bool {
	rv := C.ImageGetShouldInterpolate(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Paints the area within the current path, using the nonzero winding number rule. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456306-cgcontextfillpath?language=objc
func ContextFillPath(c ContextRef) {
	C.ContextFillPath(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Returns the type identifier for Core Graphics PDF documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402597-cgpdfdocumentgettypeid?language=objc
func PDFDocumentGetTypeID() corefoundation.TypeID {
	rv := C.PDFDocumentGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Returns the rendering intent setting for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456350-cgimagegetrenderingintent?language=objc
func ImageGetRenderingIntent(image ImageRef) ColorRenderingIntent {
	rv := C.ImageGetRenderingIntent(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.AliasType
	return ColorRenderingIntent(rv)
}

// Posts a Quartz event into the event stream at a specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456527-cgeventpost?language=objc
func EventPost(tap EventTapLocation, event EventRef) {
	C.EventPost(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventTapLocation)(tap),
		// *typing.RefType
		unsafe.Pointer(event),
	)
}

// Returns the thickness of the dominant vertical stems of glyphs in a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396380-cgfontgetstemv?language=objc
func FontGetStemV(font FontRef) float64 {
	rv := C.FontGetStemV(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Strokes a sequence of line segments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454389-cgcontextstrokelinesegments?language=objc
func ContextStrokeLineSegments(c ContextRef, points *Point, count uint) {
	C.ContextStrokeLineSegments(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PointerType
		(*C.CGPoint)(unsafe.Pointer(&points)),
		// *typing.PrimitiveType
		C.uint(count),
	)
}

// Paints the area within the current path, using the even-odd fill rule. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454865-cgcontexteofillpath?language=objc
func ContextEOFillPath(c ContextRef) {
	C.ContextEOFillPath(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Obtains exclusive use of a display for an application using the options you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454934-cgdisplaycapturewithoptions?language=objc
func DisplayCaptureWithOptions(display DirectDisplayID, options CaptureOptions) Error {
	rv := C.DisplayCaptureWithOptions(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGCaptureOptions)(options),
	)
	// *typing.AliasType
	return Error(rv)
}

// Checks to see whether the specified point is contained in the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454778-cgcontextpathcontainspoint?language=objc
func ContextPathContainsPoint(c ContextRef, point Point, mode PathDrawingMode) bool {
	rv := C.ContextPathContainsPoint(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&point)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGPathDrawingMode)(mode),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3684559-cgcolorspaceusesitur_2100tf?language=objc
func ColorSpaceUsesITUR_2100TF(arg0 ColorSpaceRef) bool {
	rv := C.ColorSpaceUsesITUR_2100TF(
		// *typing.RefType
		unsafe.Pointer(arg0),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2875135-cgpdfdocumentgetoutline?language=objc
func PDFDocumentGetOutline(document PDFDocumentRef) corefoundation.DictionaryRef {
	rv := C.PDFDocumentGetOutline(
		// *typing.RefType
		unsafe.Pointer(document),
	)
	// *typing.RefType
	return corefoundation.DictionaryRef(rv)
}

// Decrements the retain count of a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411148-cgpathrelease?language=objc
func PathRelease(path unsafe.Pointer) {
	C.PathRelease(
		// *typing.RefType
		unsafe.Pointer(path),
	)
}

// Returns the Core Foundation type identifier for Core Graphics paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411192-cgpathgettypeid?language=objc
func PathGetTypeID() corefoundation.TypeID {
	rv := C.PathGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Decrements the retain count of a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402593-cgpdfdocumentrelease?language=objc
func PDFDocumentRelease(document PDFDocumentRef) {
	C.PDFDocumentRelease(
		// *typing.RefType
		unsafe.Pointer(document),
	)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454440-cgwindowservercreateserverport?language=objc
func WindowServerCreateServerPort() corefoundation.MachPortRef {
	rv := C.WindowServerCreateServerPort()
	// *typing.RefType
	return corefoundation.MachPortRef(rv)
}

// Returns an affine transform that maps user space coordinates to device space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455677-cgcontextgetuserspacetodevicespa?language=objc
func ContextGetUserSpaceToDeviceSpaceTransform(c ContextRef) AffineTransform {
	rv := C.ContextGetUserSpaceToDeviceSpaceTransform(
		// *typing.RefType
		unsafe.Pointer(c),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Draws an image into a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454845-cgcontextdrawimage?language=objc
func ContextDrawImage(c ContextRef, rect Rect, image ImageRef) {
	C.ContextDrawImage(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.RefType
		unsafe.Pointer(image),
	)
}

// Create an immutable path of a rounded rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411218-cgpathcreatewithroundedrect?language=objc
func PathCreateWithRoundedRect(rect Rect, cornerWidth float64, cornerHeight float64, transform *AffineTransform) unsafe.Pointer {
	rv := C.PathCreateWithRoundedRect(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.PrimitiveType
		C.float(cornerWidth),
		// *typing.PrimitiveType
		C.float(cornerHeight),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&transform)),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3042356-cgpdfcontextbegintag?language=objc
func PDFContextBeginTag(context ContextRef, tagType PDFTagType, tagProperties corefoundation.DictionaryRef) {
	C.PDFContextBeginTag(
		// *typing.RefType
		unsafe.Pointer(context),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGPDFTagType)(tagType),
		// *typing.RefType
		unsafe.Pointer(tagProperties),
	)
}

// Returns the PDF type identifier of an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455189-cgpdfobjectgettype?language=objc
func PDFObjectGetType(object unsafe.Pointer) PDFObjectType {
	rv := C.PDFObjectGetType(
		// *typing.RefType
		unsafe.Pointer(object),
	)
	// *typing.AliasType
	return PDFObjectType(rv)
}

// Creates a bitmap image mask from data supplied by a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455089-cgimagemaskcreate?language=objc
func ImageMaskCreate(width uint, height uint, bitsPerComponent uint, bitsPerPixel uint, bytesPerRow uint, provider DataProviderRef, decode *float64, shouldInterpolate bool) ImageRef {
	rv := C.ImageMaskCreate(
		// *typing.PrimitiveType
		C.uint(width),
		// *typing.PrimitiveType
		C.uint(height),
		// *typing.PrimitiveType
		C.uint(bitsPerComponent),
		// *typing.PrimitiveType
		C.uint(bitsPerPixel),
		// *typing.PrimitiveType
		C.uint(bytesPerRow),
		// *typing.RefType
		unsafe.Pointer(provider),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&decode)),
		// *typing.PrimitiveType
		C.bool(shouldInterpolate),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Creates a color in the Generic RGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455631-cgcolorcreategenericrgb?language=objc
func ColorCreateGenericRGB(red float64, green float64, blue float64, alpha float64) ColorRef {
	rv := C.ColorCreateGenericRGB(
		// *typing.PrimitiveType
		C.float(red),
		// *typing.PrimitiveType
		C.float(green),
		// *typing.PrimitiveType
		C.float(blue),
		// *typing.PrimitiveType
		C.float(alpha),
	)
	// *typing.RefType
	return ColorRef(rv)
}

// Returns the floating-point value of a field in a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455506-cgeventgetdoublevaluefield?language=objc
func EventGetDoubleValueField(event EventRef, field EventField) float64 {
	rv := C.EventGetDoubleValueField(
		// *typing.RefType
		unsafe.Pointer(event),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventField)(field),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Gets the information dictionary for a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402589-cgpdfdocumentgetinfo?language=objc
func PDFDocumentGetInfo(document PDFDocumentRef) unsafe.Pointer {
	rv := C.PDFDocumentGetInfo(
		// *typing.RefType
		unsafe.Pointer(document),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Checks whether the converter is currently converting data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454582-cgpsconverterisconverting?language=objc
func PSConverterIsConverting(converter PSConverterRef) bool {
	rv := C.PSConverterIsConverting(
		// *typing.RefType
		unsafe.Pointer(converter),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a copy of a font using a variation specification dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396373-cgfontcreatecopywithvariations?language=objc
func FontCreateCopyWithVariations(font FontRef, variations corefoundation.DictionaryRef) FontRef {
	rv := C.FontCreateCopyWithVariations(
		// *typing.RefType
		unsafe.Pointer(font),
		// *typing.RefType
		unsafe.Pointer(variations),
	)
	// *typing.RefType
	return FontRef(rv)
}

// Changes the configuration of a mirroring set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454531-cgconfiguredisplaymirrorofdispla?language=objc
func ConfigureDisplayMirrorOfDisplay(config unsafe.Pointer, display DirectDisplayID, master DirectDisplayID) Error {
	rv := C.ConfigureDisplayMirrorOfDisplay(
		// *typing.RefType
		unsafe.Pointer(config),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(master),
	)
	// *typing.AliasType
	return Error(rv)
}

// Returns the refresh rate of the specified display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454661-cgdisplaymodegetrefreshrate?language=objc
func DisplayModeGetRefreshRate(mode DisplayModeRef) float64 {
	rv := C.DisplayModeGetRefreshRate(
		// *typing.RefType
		unsafe.Pointer(mode),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Returns the interval that local hardware events may be suppressed following the posting of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408774-cgeventsourcegetlocaleventssuppr?language=objc
func EventSourceGetLocalEventsSuppressionInterval(source EventSourceRef) corefoundation.TimeInterval {
	rv := C.EventSourceGetLocalEventsSuppressionInterval(
		// *typing.RefType
		unsafe.Pointer(source),
	)
	// *typing.AliasType
	return corefoundation.TimeInterval(rv)
}

// Adds an arc of a circle to the current path, using a radius and tangent points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456238-cgcontextaddarctopoint?language=objc
func ContextAddArcToPoint(c ContextRef, x1 float64, y1 float64, x2 float64, y2 float64, radius float64) {
	C.ContextAddArcToPoint(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(x1),
		// *typing.PrimitiveType
		C.float(y1),
		// *typing.PrimitiveType
		C.float(x2),
		// *typing.PrimitiveType
		C.float(y2),
		// *typing.PrimitiveType
		C.float(radius),
	)
}

// Creates a data consumer that writes data to a location specified by a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454474-cgdataconsumercreatewithurl?language=objc
func DataConsumerCreateWithURL(url corefoundation.URLRef) DataConsumerRef {
	rv := C.DataConsumerCreateWithURL(
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.RefType
	return DataConsumerRef(rv)
}

// Returns the width and height of a layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450890-cglayergetsize?language=objc
func LayerGetSize(layer LayerRef) Size {
	rv := C.LayerGetSize(
		// *typing.RefType
		unsafe.Pointer(layer),
	)
	// *typing.StructType
	return *(*Size)(unsafe.Pointer(&rv))
}

// Increments the retain count of a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586506-cgcontextretain?language=objc
func ContextRetain(c ContextRef) ContextRef {
	rv := C.ContextRetain(
		// *typing.RefType
		unsafe.Pointer(c),
	)
	// *typing.RefType
	return ContextRef(rv)
}

// Reserves the fade hardware for a specified time interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456391-cgacquiredisplayfadereservation?language=objc
func AcquireDisplayFadeReservation(seconds DisplayReservationInterval, token *DisplayFadeReservationToken) Error {
	rv := C.AcquireDisplayFadeReservation(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDisplayReservationInterval)(seconds),
		// *typing.PointerType
		(*C.CGDisplayFadeReservationToken)(unsafe.Pointer(&token)),
	)
	// *typing.AliasType
	return Error(rv)
}

// Returns the event flags of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455642-cgeventgetflags?language=objc
func EventGetFlags(event EventRef) EventFlags {
	rv := C.EventGetFlags(
		// *typing.RefType
		unsafe.Pointer(event),
	)
	// *typing.AliasType
	return EventFlags(rv)
}

// Returns the glyph name of the specified glyph in the specified font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396349-cgfontcopyglyphnameforglyph?language=objc
func FontCopyGlyphNameForGlyph(font FontRef, glyph Glyph) corefoundation.StringRef {
	rv := C.FontCopyGlyphNameForGlyph(
		// *typing.RefType
		unsafe.Pointer(font),
		// *typing.AliasType
		// *typing.AliasType
		(C.CGGlyph)(glyph),
	)
	// *typing.RefType
	return corefoundation.StringRef(rv)
}

// Sets the current fill color to a value in the DeviceRGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455624-cgcontextsetrgbfillcolor?language=objc
func ContextSetRGBFillColor(c ContextRef, red float64, green float64, blue float64, alpha float64) {
	C.ContextSetRGBFillColor(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(red),
		// *typing.PrimitiveType
		C.float(green),
		// *typing.PrimitiveType
		C.float(blue),
		// *typing.PrimitiveType
		C.float(alpha),
	)
}

// Sets the current graphics state to the state most recently saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455391-cgcontextrestoregstate?language=objc
func ContextRestoreGState(c ContextRef) {
	C.ContextRestoreGState(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Creates an immutable copy of a graphics path transformed by a transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411161-cgpathcreatecopybytransformingpa?language=objc
func PathCreateCopyByTransformingPath(path unsafe.Pointer, transform *AffineTransform) unsafe.Pointer {
	rv := C.PathCreateCopyByTransformingPath(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&transform)),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns the full name associated with a font object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396357-cgfontcopyfullname?language=objc
func FontCopyFullName(font FontRef) corefoundation.StringRef {
	rv := C.FontCopyFullName(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.RefType
	return corefoundation.StringRef(rv)
}

// Returns an affine transformation matrix constructed by combining two existing affine transforms. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455996-cgaffinetransformconcat?language=objc
func AffineTransformConcat(t1 AffineTransform, t2 AffineTransform) AffineTransform {
	rv := C.AffineTransformConcat(
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t1)),
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t2)),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Returns the page number of the specified PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454587-cgpdfpagegetpagenumber?language=objc
func PDFPageGetPageNumber(page PDFPageRef) uint {
	rv := C.PDFPageGetPageNumber(
		// *typing.RefType
		unsafe.Pointer(page),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Sets the location of a Quartz mouse event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456389-cgeventsetlocation?language=objc
func EventSetLocation(event EventRef, location Point) {
	C.EventSetLocation(
		// *typing.RefType
		unsafe.Pointer(event),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&location)),
	)
}

// Decrements the retain count of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396363-cgfontrelease?language=objc
func FontRelease(font FontRef) {
	C.FontRelease(
		// *typing.RefType
		unsafe.Pointer(font),
	)
}

// Returns a CFString object that represents a PDF string as a text string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456234-cgpdfstringcopytextstring?language=objc
func PDFStringCopyTextString(string_ unsafe.Pointer) corefoundation.StringRef {
	rv := C.PDFStringCopyTextString(
		// *typing.RefType
		unsafe.Pointer(string_),
	)
	// *typing.RefType
	return corefoundation.StringRef(rv)
}

// Returns whether the specified PDF file is encrypted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402591-cgpdfdocumentisencrypted?language=objc
func PDFDocumentIsEncrypted(document PDFDocumentRef) bool {
	rv := C.PDFDocumentIsEncrypted(
		// *typing.RefType
		unsafe.Pointer(document),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Paints a rectangular path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454675-cgcontextstrokerect?language=objc
func ContextStrokeRect(c ContextRef, rect Rect) {
	C.ContextStrokeRect(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

// Enables or disables subpixel positioning in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455671-cgcontextsetshouldsubpixelpositi?language=objc
func ContextSetShouldSubpixelPositionFonts(c ContextRef, shouldSubpixelPositionFonts bool) {
	C.ContextSetShouldSubpixelPositionFonts(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.bool(shouldSubpixelPositionFonts),
	)
}

// Returns a path object built from the current path information in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455397-cgcontextcopypath?language=objc
func ContextCopyPath(c ContextRef) unsafe.Pointer {
	rv := C.ContextCopyPath(
		// *typing.RefType
		unsafe.Pointer(c),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// The Universal Type Identifier for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456067-cgimagegetuttype?language=objc
func ImageGetUTType(image ImageRef) corefoundation.StringRef {
	rv := C.ImageGetUTType(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.RefType
	return corefoundation.StringRef(rv)
}

// Determines whether Core Graphics can create a subset of the font in PostScript format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396365-cgfontcancreatepostscriptsubset?language=objc
func FontCanCreatePostScriptSubset(font FontRef, format FontPostScriptFormat) bool {
	rv := C.FontCanCreatePostScriptSubset(
		// *typing.RefType
		unsafe.Pointer(font),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGFontPostScriptFormat)(format),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the values of the color components (including alpha) associated with a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455930-cgcolorgetcomponents?language=objc
func ColorGetComponents(color ColorRef) *float64 {
	rv := C.ColorGetComponents(
		// *typing.RefType
		unsafe.Pointer(color),
	)
	// *typing.PointerType
	return *(**float64)(unsafe.Pointer(&rv))
}

// Returns the type identifier for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455504-cgcontextgettypeid?language=objc
func ContextGetTypeID() corefoundation.TypeID {
	rv := C.ContextGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Creates a data consumer that writes to a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456292-cgdataconsumercreatewithcfdata?language=objc
func DataConsumerCreateWithCFData(data unsafe.Pointer) DataConsumerRef {
	rv := C.DataConsumerCreateWithCFData(
		// *typing.RefType
		unsafe.Pointer(data),
	)
	// *typing.RefType
	return DataConsumerRef(rv)
}

// Adds a rectangular path to the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456617-cgcontextaddrect?language=objc
func ContextAddRect(c ContextRef, rect Rect) {
	C.ContextAddRect(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411136-cgpathaddrelativearc?language=objc
func PathAddRelativeArc(path MutablePathRef, matrix *AffineTransform, x float64, y float64, radius float64, startAngle float64, delta float64) {
	C.PathAddRelativeArc(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&matrix)),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
		// *typing.PrimitiveType
		C.float(radius),
		// *typing.PrimitiveType
		C.float(startAngle),
		// *typing.PrimitiveType
		C.float(delta),
	)
}

// Decrements the retain count of a data consumer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1508424-cgdataconsumerrelease?language=objc
func DataConsumerRelease(consumer DataConsumerRef) {
	C.DataConsumerRelease(
		// *typing.RefType
		unsafe.Pointer(consumer),
	)
}

// Enables or disables stereo operation for a display, as part of a display configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456308-cgconfiguredisplaystereooperatio?language=objc
func ConfigureDisplayStereoOperation(config unsafe.Pointer, display DirectDisplayID, stereo int, forceBlueLine int) Error {
	rv := C.ConfigureDisplayStereoOperation(
		// *typing.RefType
		unsafe.Pointer(config),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.boolean_t)(stereo),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.boolean_t)(forceBlueLine),
	)
	// *typing.AliasType
	return Error(rv)
}

// Returns a size that is transformed from user space coordinates to device space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456619-cgcontextconvertsizetodevicespac?language=objc
func ContextConvertSizeToDeviceSpace(c ContextRef, size Size) Size {
	rv := C.ContextConvertSizeToDeviceSpace(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGSize)(unsafe.Pointer(&size)),
	)
	// *typing.StructType
	return *(*Size)(unsafe.Pointer(&rv))
}

// Appends a cubic Bézier curve to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411212-cgpathaddcurvetopoint?language=objc
func PathAddCurveToPoint(path MutablePathRef, m *AffineTransform, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {
	C.PathAddCurveToPoint(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.PrimitiveType
		C.float(cp1x),
		// *typing.PrimitiveType
		C.float(cp1y),
		// *typing.PrimitiveType
		C.float(cp2x),
		// *typing.PrimitiveType
		C.float(cp2y),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
}

// Returns a Boolean value indicating whether a display is built-in, such as the internal display in portable systems. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454566-cgdisplayisbuiltin?language=objc
func DisplayIsBuiltin(display DirectDisplayID) int {
	rv := C.DisplayIsBuiltin(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return int(rv)
}

// Hides the mouse cursor, and increments the hide cursor count. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456534-cgdisplayhidecursor?language=objc
func DisplayHideCursor(display DirectDisplayID) Error {
	rv := C.DisplayHideCursor(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return Error(rv)
}

// Returns the logical unit number of a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456489-cgdisplayunitnumber?language=objc
func DisplayUnitNumber(display DirectDisplayID) uint32 {
	rv := C.DisplayUnitNumber(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Switches a display to a different mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454760-cgdisplaysetdisplaymode?language=objc
func DisplaySetDisplayMode(display DirectDisplayID, mode DisplayModeRef, options corefoundation.DictionaryRef) Error {
	rv := C.DisplaySetDisplayMode(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.RefType
		unsafe.Pointer(mode),
		// *typing.RefType
		unsafe.Pointer(options),
	)
	// *typing.AliasType
	return Error(rv)
}

// Draws glyphs at the provided position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456200-cgcontextshowglyphsatpositions?language=objc
func ContextShowGlyphsAtPositions(c ContextRef, glyphs *Glyph, Lpositions *Point, count uint) {
	C.ContextShowGlyphsAtPositions(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PointerType
		(*C.CGGlyph)(unsafe.Pointer(&glyphs)),
		// *typing.PointerType
		(*C.CGPoint)(unsafe.Pointer(&Lpositions)),
		// *typing.PrimitiveType
		C.uint(count),
	)
}

// Sets the current fill color in a graphics context, using a CGColor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454079-cgcontextsetfillcolorwithcolor?language=objc
func ContextSetFillColorWithColor(c ContextRef, color ColorRef) {
	C.ContextSetFillColorWithColor(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.RefType
		unsafe.Pointer(color),
	)
}

// Returns the leading of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396390-cgfontgetleading?language=objc
func FontGetLeading(font FontRef) int {
	rv := C.FontGetLeading(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.PrimitiveType
	return int(rv)
}

// Returns the Core Foundation type identifier for PostScript converters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454188-cgpsconvertergettypeid?language=objc
func PSConverterGetTypeID() corefoundation.TypeID {
	rv := C.PSConverterGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Creates a gradient object from a color space and the provided color objects and locations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1398458-cggradientcreatewithcolors?language=objc
func GradientCreateWithColors(space ColorSpaceRef, colors corefoundation.ArrayRef, locations *float64) GradientRef {
	rv := C.GradientCreateWithColors(
		// *typing.RefType
		unsafe.Pointer(space),
		// *typing.RefType
		unsafe.Pointer(colors),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&locations)),
	)
	// *typing.RefType
	return GradientRef(rv)
}

// Returns a Boolean value indicating whether a display is always in a mirroring set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456110-cgdisplayisalwaysinmirrorset?language=objc
func DisplayIsAlwaysInMirrorSet(display DirectDisplayID) int {
	rv := C.DisplayIsAlwaysInMirrorSet(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return int(rv)
}

// Creates a dashed copy of another path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411134-cgpathcreatecopybydashingpath?language=objc
func PathCreateCopyByDashingPath(path unsafe.Pointer, transform *AffineTransform, phase float64, lengths *float64, count uint) unsafe.Pointer {
	rv := C.PathCreateCopyByDashingPath(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&transform)),
		// *typing.PrimitiveType
		C.float(phase),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&lengths)),
		// *typing.PrimitiveType
		C.uint(count),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Sets the current fill color to a value in the DeviceCMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454214-cgcontextsetcmykfillcolor?language=objc
func ContextSetCMYKFillColor(c ContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64) {
	C.ContextSetCMYKFillColor(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(cyan),
		// *typing.PrimitiveType
		C.float(magenta),
		// *typing.PrimitiveType
		C.float(yellow),
		// *typing.PrimitiveType
		C.float(black),
		// *typing.PrimitiveType
		C.float(alpha),
	)
}

// Cancels a set of display configuration changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455522-cgcanceldisplayconfiguration?language=objc
func CancelDisplayConfiguration(config unsafe.Pointer) Error {
	rv := C.CancelDisplayConfiguration(
		// *typing.RefType
		unsafe.Pointer(config),
	)
	// *typing.AliasType
	return Error(rv)
}

// Appends an array of new line segments to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411171-cgpathaddlines?language=objc
func PathAddLines(path MutablePathRef, m *AffineTransform, points *Point, count uint) {
	C.PathAddLines(
		// *typing.RefType
		unsafe.Pointer(path),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&m)),
		// *typing.PointerType
		(*C.CGPoint)(unsafe.Pointer(&points)),
		// *typing.PrimitiveType
		C.uint(count),
	)
}

// Sets the gamma function for a display by specifying the coefficients of the gamma transfer formula. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454126-cgsetdisplaytransferbyformula?language=objc
func SetDisplayTransferByFormula(display DirectDisplayID, redMin GammaValue, redMax GammaValue, redGamma GammaValue, greenMin GammaValue, greenMax GammaValue, greenGamma GammaValue, blueMin GammaValue, blueMax GammaValue, blueGamma GammaValue) Error {
	rv := C.SetDisplayTransferByFormula(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGGammaValue)(redMin),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGGammaValue)(redMax),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGGammaValue)(redGamma),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGGammaValue)(greenMin),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGGammaValue)(greenMax),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGGammaValue)(greenGamma),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGGammaValue)(blueMin),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGGammaValue)(blueMax),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGGammaValue)(blueGamma),
	)
	// *typing.AliasType
	return Error(rv)
}

// Returns a Boolean value indicating whether a display is sleeping (and is therefore not drawable). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454260-cgdisplayisasleep?language=objc
func DisplayIsAsleep(display DirectDisplayID) int {
	rv := C.DisplayIsAsleep(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return int(rv)
}

// Returns the bitmap information for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454200-cgimagegetbitmapinfo?language=objc
func ImageGetBitmapInfo(image ImageRef) BitmapInfo {
	rv := C.ImageGetBitmapInfo(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.AliasType
	return BitmapInfo(rv)
}

// Returns the width in pixels of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455607-cgbitmapcontextgetwidth?language=objc
func BitmapContextGetWidth(context ContextRef) uint {
	rv := C.BitmapContextGetWidth(
		// *typing.RefType
		unsafe.Pointer(context),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Returns an affine transformation matrix constructed from values you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455865-cgaffinetransformmake?language=objc
func AffineTransformMake(a float64, b float64, c float64, d float64, tx float64, ty float64) AffineTransform {
	rv := C.AffineTransformMake(
		// *typing.PrimitiveType
		C.float(a),
		// *typing.PrimitiveType
		C.float(b),
		// *typing.PrimitiveType
		C.float(c),
		// *typing.PrimitiveType
		C.float(d),
		// *typing.PrimitiveType
		C.float(tx),
		// *typing.PrimitiveType
		C.float(ty),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Create an immutable path of an ellipse. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411177-cgpathcreatewithellipseinrect?language=objc
func PathCreateWithEllipseInRect(rect Rect, transform *AffineTransform) unsafe.Pointer {
	rv := C.PathCreateWithEllipseInRect(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.PointerType
		(*C.CGAffineTransform)(unsafe.Pointer(&transform)),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns whether two rectangles intersect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454747-cgrectintersectsrect?language=objc
func RectIntersectsRect(rect1 Rect, rect2 Rect) bool {
	rv := C.RectIntersectsRect(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect1)),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect2)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns an image containing the contents of the specified display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455691-cgdisplaycreateimage?language=objc
func DisplayCreateImage(displayID DirectDisplayID) ImageRef {
	rv := C.DisplayCreateImage(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(displayID),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Sets the line width for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455270-cgcontextsetlinewidth?language=objc
func ContextSetLineWidth(c ContextRef, width float64) {
	C.ContextSetLineWidth(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(width),
	)
}

// Decrements the retain count of a Core Graphics pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1552266-cgpatternrelease?language=objc
func PatternRelease(pattern PatternRef) {
	C.PatternRelease(
		// *typing.RefType
		unsafe.Pointer(pattern),
	)
}

// Modifies the current clipping path, using the even-odd rule. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455944-cgcontexteoclip?language=objc
func ContextEOClip(c ContextRef) {
	C.ContextEOClip(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3684557-cgcolorspacecreateextended?language=objc
func ColorSpaceCreateExtended(space ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateExtended(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Sets whether or not to allow subpixel quantization for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456263-cgcontextsetallowsfontsubpixelqu?language=objc
func ContextSetAllowsFontSubpixelQuantization(c ContextRef, allowsFontSubpixelQuantization bool) {
	C.ContextSetAllowsFontSubpixelQuantization(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.bool(allowsFontSubpixelQuantization),
	)
}

// Returns a point that is transformed from device space coordinates to user space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456451-cgcontextconvertpointtouserspace?language=objc
func ContextConvertPointToUserSpace(c ContextRef, point Point) Point {
	rv := C.ContextConvertPointToUserSpace(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&point)),
	)
	// *typing.StructType
	return *(*Point)(unsafe.Pointer(&rv))
}

// Increments the retain count of a data consumer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1508422-cgdataconsumerretain?language=objc
func DataConsumerRetain(consumer DataConsumerRef) DataConsumerRef {
	rv := C.DataConsumerRetain(
		// *typing.RefType
		unsafe.Pointer(consumer),
	)
	// *typing.RefType
	return DataConsumerRef(rv)
}

// Sets the mask that indicates which classes of local hardware events are enabled during event suppression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408770-cgeventsourcesetlocaleventsfilte?language=objc
func EventSourceSetLocalEventsFilterDuringSuppressionState(source EventSourceRef, filter EventFilterMask, state EventSuppressionState) {
	C.EventSourceSetLocalEventsFilterDuringSuppressionState(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventFilterMask)(filter),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventSuppressionState)(state),
	)
}

// Sets the current stroke color to a value in the DeviceGray color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455209-cgcontextsetgraystrokecolor?language=objc
func ContextSetGrayStrokeColor(c ContextRef, gray float64, alpha float64) {
	C.ContextSetGrayStrokeColor(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(gray),
		// *typing.PrimitiveType
		C.float(alpha),
	)
}

// Returns the vendor number of the specified display’s monitor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455233-cgdisplayvendornumber?language=objc
func DisplayVendorNumber(display DirectDisplayID) uint32 {
	rv := C.DisplayVendorNumber(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Returns whether a rectangle contains a specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456316-cgrectcontainspoint?language=objc
func RectContainsPoint(rect Rect, point Point) bool {
	rv := C.RectContainsPoint(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.StructType
		*(*C.CGPoint)(unsafe.Pointer(&point)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the Core Foundation type identifier for Core Graphics fonts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396369-cgfontgettypeid?language=objc
func FontGetTypeID() corefoundation.TypeID {
	rv := C.FontGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Returns a graphics context suitable for drawing to a captured display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456576-cgdisplaygetdrawingcontext?language=objc
func DisplayGetDrawingContext(display DirectDisplayID) ContextRef {
	rv := C.DisplayGetDrawingContext(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.RefType
	return ContextRef(rv)
}

// Enables or disables font smoothing in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455816-cgcontextsetshouldsmoothfonts?language=objc
func ContextSetShouldSmoothFonts(c ContextRef, shouldSmoothFonts bool) {
	C.ContextSetShouldSmoothFonts(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.bool(shouldSmoothFonts),
	)
}

// Returns a new Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454913-cgeventcreate?language=objc
func EventCreate(source EventSourceRef) EventRef {
	rv := C.EventCreate(
		// *typing.RefType
		unsafe.Pointer(source),
	)
	// *typing.RefType
	return EventRef(rv)
}

// For a secondary display in a mirroring set, returns the primary display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454556-cgdisplaymirrorsdisplay?language=objc
func DisplayMirrorsDisplay(display DirectDisplayID) DirectDisplayID {
	rv := C.DisplayMirrorsDisplay(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return DirectDisplayID(rv)
}

// Checks whether an affine transform is the identity transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455754-cgaffinetransformisidentity?language=objc
func AffineTransformIsIdentity(t AffineTransform) bool {
	rv := C.AffineTransformIsIdentity(
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Starts a new page in a page-based graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454794-cgcontextbeginpage?language=objc
func ContextBeginPage(c ContextRef, mediaBox *Rect) {
	C.ContextBeginPage(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PointerType
		(*C.CGRect)(unsafe.Pointer(&mediaBox)),
	)
}

// Returns the width of a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454758-cgrectgetwidth?language=objc
func RectGetWidth(rect Rect) float64 {
	rv := C.RectGetWidth(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Returns a point with the specified coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455746-cgpointmake?language=objc
func PointMake(x float64, y float64) Point {
	rv := C.PointMake(
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
	// *typing.StructType
	return *(*Point)(unsafe.Pointer(&rv))
}

// Returns a new Quartz scrolling event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541327-cgeventcreatescrollwheelevent?language=objc
func EventCreateScrollWheelEvent(source EventSourceRef, units ScrollEventUnit, wheelCount uint32, wheel1 int32) EventRef {
	rv := C.EventCreateScrollWheelEvent(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGScrollEventUnit)(units),
		// *typing.PrimitiveType
		C.uint32_t(wheelCount),
		// *typing.PrimitiveType
		C.int32_t(wheel1),
	)
	// *typing.RefType
	return EventRef(rv)
}

// Returns a color object that represents a constant color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454283-cgcolorgetconstantcolor?language=objc
func ColorGetConstantColor(colorName corefoundation.StringRef) ColorRef {
	rv := C.ColorGetConstantColor(
		// *typing.RefType
		unsafe.Pointer(colorName),
	)
	// *typing.RefType
	return ColorRef(rv)
}

// Returns a Quartz event source created from an existing Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455393-cgeventcreatesourcefromevent?language=objc
func EventCreateSourceFromEvent(event EventRef) EventSourceRef {
	rv := C.EventCreateSourceFromEvent(
		// *typing.RefType
		unsafe.Pointer(event),
	)
	// *typing.RefType
	return EventSourceRef(rv)
}

// Returns the x-height of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396410-cgfontgetxheight?language=objc
func FontGetXHeight(font FontRef) int {
	rv := C.FontGetXHeight(
		// *typing.RefType
		unsafe.Pointer(font),
	)
	// *typing.PrimitiveType
	return int(rv)
}

// Returns the location of a Quartz mouse event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455589-cgeventgetunflippedlocation?language=objc
func EventGetUnflippedLocation(event EventRef) Point {
	rv := C.EventGetUnflippedLocation(
		// *typing.RefType
		unsafe.Pointer(event),
	)
	// *typing.StructType
	return *(*Point)(unsafe.Pointer(&rv))
}

// Creates a bitmap image from an existing image and an image mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456337-cgimagecreatewithmask?language=objc
func ImageCreateWithMask(image ImageRef, mask ImageRef) ImageRef {
	rv := C.ImageCreateWithMask(
		// *typing.RefType
		unsafe.Pointer(image),
		// *typing.RefType
		unsafe.Pointer(mask),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Sets the URL associated with a rectangle in a PDF graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455622-cgpdfcontextseturlforrect?language=objc
func PDFContextSetURLForRect(context ContextRef, url corefoundation.URLRef, rect Rect) {
	C.PDFContextSetURLForRect(
		// *typing.RefType
		unsafe.Pointer(context),
		// *typing.RefType
		unsafe.Pointer(url),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2091882-cgcolorspacecreatelinearized?language=objc
func ColorSpaceCreateLinearized(space ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateLinearized(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.RefType
	return ColorSpaceRef(rv)
}

// Returns the number of entries in the color table of an indexed color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408883-cgcolorspacegetcolortablecount?language=objc
func ColorSpaceGetColorTableCount(space ColorSpaceRef) uint {
	rv := C.ColorSpaceGetColorTableCount(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Maps an OpenGL display mask to a display ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455386-cgopengldisplaymasktodisplayid?language=objc
func OpenGLDisplayMaskToDisplayID(mask OpenGLDisplayMask) DirectDisplayID {
	rv := C.OpenGLDisplayMaskToDisplayID(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGOpenGLDisplayMask)(mask),
	)
	// *typing.AliasType
	return DirectDisplayID(rv)
}

// Forces all pending drawing operations in a window context to be rendered immediately to the destination device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454895-cgcontextflush?language=objc
func ContextFlush(c ContextRef) {
	C.ContextFlush(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Returns a vector with the specified dimension values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454811-cgvectormake?language=objc
func VectorMake(dx float64, dy float64) Vector {
	rv := C.VectorMake(
		// *typing.PrimitiveType
		C.float(dx),
		// *typing.PrimitiveType
		C.float(dy),
	)
	// *typing.StructType
	return *(*Vector)(unsafe.Pointer(&rv))
}

// Returns the height of the specified display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455380-cgdisplaymodegetheight?language=objc
func DisplayModeGetHeight(mode DisplayModeRef) uint {
	rv := C.DisplayModeGetHeight(
		// *typing.RefType
		unsafe.Pointer(mode),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Returns the smallest rectangle that contains the two source rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455837-cgrectunion?language=objc
func RectUnion(r1 Rect, r2 Rect) Rect {
	rv := C.RectUnion(
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&r1)),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&r2)),
	)
	// *typing.StructType
	return *(*Rect)(unsafe.Pointer(&rv))
}

// Parses the content stream of a PDF scanner object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454698-cgpdfscannerscan?language=objc
func PDFScannerScan(scanner unsafe.Pointer) bool {
	rv := C.PDFScannerScan(
		// *typing.RefType
		unsafe.Pointer(scanner),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Adds a sequence of connected straight-line segments to the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455461-cgcontextaddlines?language=objc
func ContextAddLines(c ContextRef, points *Point, count uint) {
	C.ContextAddLines(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PointerType
		(*C.CGPoint)(unsafe.Pointer(&points)),
		// *typing.PrimitiveType
		C.uint(count),
	)
}

// Sets the miter limit for the joins of connected lines in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456499-cgcontextsetmiterlimit?language=objc
func ContextSetMiterLimit(c ContextRef, limit float64) {
	C.ContextSetMiterLimit(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(limit),
	)
}

// Returns the dictionary associated with a PDF stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456118-cgpdfstreamgetdictionary?language=objc
func PDFStreamGetDictionary(stream unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFStreamGetDictionary(
		// *typing.RefType
		unsafe.Pointer(stream),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns the number of bytes in a PDF string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454095-cgpdfstringgetlength?language=objc
func PDFStringGetLength(string_ unsafe.Pointer) uint {
	rv := C.PDFStringGetLength(
		// *typing.RefType
		unsafe.Pointer(string_),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Modifies the current clipping path, using the nonzero winding number rule. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455262-cgcontextclip?language=objc
func ContextClip(c ContextRef) {
	C.ContextClip(
		// *typing.RefType
		unsafe.Pointer(c),
	)
}

// Releases all captured displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454901-cgreleasealldisplays?language=objc
func ReleaseAllDisplays() Error {
	rv := C.ReleaseAllDisplays()
	// *typing.AliasType
	return Error(rv)
}

// Returns the serial number of a display monitor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455409-cgdisplayserialnumber?language=objc
func DisplaySerialNumber(display DirectDisplayID) uint32 {
	rv := C.DisplaySerialNumber(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Returns the name used to create the specified color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408903-cgcolorspacecopyname?language=objc
func ColorSpaceCopyName(space ColorSpaceRef) corefoundation.StringRef {
	rv := C.ColorSpaceCopyName(
		// *typing.RefType
		unsafe.Pointer(space),
	)
	// *typing.RefType
	return corefoundation.StringRef(rv)
}

// Sets the style for the endpoints of lines drawn in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454326-cgcontextsetlinecap?language=objc
func ContextSetLineCap(c ContextRef, cap LineCap) {
	C.ContextSetLineCap(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGLineCap)(cap),
	)
}

// Decrements the retain count of a shading object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1573766-cgshadingrelease?language=objc
func ShadingRelease(shading ShadingRef) {
	C.ShadingRelease(
		// *typing.RefType
		unsafe.Pointer(shading),
	)
}

// Increments the retain count of a CGPDFOperatorTable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454547-cgpdfoperatortableretain?language=objc
func PDFOperatorTableRetain(table unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFOperatorTableRetain(
		// *typing.RefType
		unsafe.Pointer(table),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns the model number of a display monitor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454149-cgdisplaymodelnumber?language=objc
func DisplayModelNumber(display DirectDisplayID) uint32 {
	rv := C.DisplayModelNumber(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Sets the 64-bit user-specified data for a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408779-cgeventsourcesetuserdata?language=objc
func EventSourceSetUserData(source EventSourceRef, userData int64) {
	C.EventSourceSetUserData(
		// *typing.RefType
		unsafe.Pointer(source),
		// *typing.PrimitiveType
		C.int64_t(userData),
	)
}

// Decrements the retain count of a CGGradient object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1398460-cggradientrelease?language=objc
func GradientRelease(gradient GradientRef) {
	C.GradientRelease(
		// *typing.RefType
		unsafe.Pointer(gradient),
	)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656521-cgrequestlisteneventaccess?language=objc
func RequestListenEventAccess() bool {
	rv := C.RequestListenEventAccess()
	// *typing.PrimitiveType
	return bool(rv)
}

// Begins a new subpath at the point you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454738-cgcontextmovetopoint?language=objc
func ContextMoveToPoint(c ContextRef, x float64, y float64) {
	C.ContextMoveToPoint(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
}

// Returns an affine transformation matrix constructed by inverting an existing affine transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455264-cgaffinetransforminvert?language=objc
func AffineTransformInvert(t AffineTransform) AffineTransform {
	rv := C.AffineTransformInvert(
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t)),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Increments the retain count of a scanner object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455810-cgpdfscannerretain?language=objc
func PDFScannerRetain(scanner unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFScannerRetain(
		// *typing.RefType
		unsafe.Pointer(scanner),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Indicates whether or not a graphics path is empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411149-cgpathisempty?language=objc
func PathIsEmpty(path unsafe.Pointer) bool {
	rv := C.PathIsEmpty(
		// *typing.RefType
		unsafe.Pointer(path),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Repeatedly draws an image, scaled to the provided rectangle, to fill the current clip region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456240-cgcontextdrawtiledimage?language=objc
func ContextDrawTiledImage(c ContextRef, rect Rect, image ImageRef) {
	C.ContextDrawTiledImage(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
		// *typing.RefType
		unsafe.Pointer(image),
	)
}

// Increments the retain count of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1556741-cgimageretain?language=objc
func ImageRetain(image ImageRef) ImageRef {
	rv := C.ImageRetain(
		// *typing.RefType
		unsafe.Pointer(image),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Creates a bitmap image using the data contained within a subregion of an existing bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454683-cgimagecreatewithimageinrect?language=objc
func ImageCreateWithImageInRect(image ImageRef, rect Rect) ImageRef {
	rv := C.ImageCreateWithImageInRect(
		// *typing.RefType
		unsafe.Pointer(image),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
	// *typing.RefType
	return ImageRef(rv)
}

// Returns the rotation angle of a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455550-cgpdfpagegetrotationangle?language=objc
func PDFPageGetRotationAngle(page PDFPageRef) int {
	rv := C.PDFPageGetRotationAngle(
		// *typing.RefType
		unsafe.Pointer(page),
	)
	// *typing.PrimitiveType
	return int(rv)
}

// Appends a straight line segment from the current point to the provided point . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455213-cgcontextaddlinetopoint?language=objc
func ContextAddLineToPoint(c ContextRef, x float64, y float64) {
	C.ContextAddLineToPoint(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(x),
		// *typing.PrimitiveType
		C.float(y),
	)
}

// Retrieves a Boolean object from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454663-cgpdfscannerpopboolean?language=objc
func PDFScannerPopBoolean(scanner unsafe.Pointer, value *PDFBoolean) bool {
	rv := C.PDFScannerPopBoolean(
		// *typing.RefType
		unsafe.Pointer(scanner),
		// *typing.PointerType
		(*C.CGPDFBoolean)(unsafe.Pointer(&value)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a font object corresponding to the font specified by a PostScript or full name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396330-cgfontcreatewithfontname?language=objc
func FontCreateWithFontName(name corefoundation.StringRef) FontRef {
	rv := C.FontCreateWithFontName(
		// *typing.RefType
		unsafe.Pointer(name),
	)
	// *typing.RefType
	return FontRef(rv)
}

// Sets the level of interpolation quality for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455656-cgcontextsetinterpolationquality?language=objc
func ContextSetInterpolationQuality(c ContextRef, quality InterpolationQuality) {
	C.ContextSetInterpolationQuality(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGInterpolationQuality)(quality),
	)
}

// Returns the document catalog of a Core Graphics PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402606-cgpdfdocumentgetcatalog?language=objc
func PDFDocumentGetCatalog(document PDFDocumentRef) unsafe.Pointer {
	rv := C.PDFDocumentGetCatalog(
		// *typing.RefType
		unsafe.Pointer(document),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Decrements the retain count of a layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450898-cglayerrelease?language=objc
func LayerRelease(layer LayerRef) {
	C.LayerRelease(
		// *typing.RefType
		unsafe.Pointer(layer),
	)
}

// Returns the unique type identifier used for CGLayerRef objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450888-cglayergettypeid?language=objc
func LayerGetTypeID() corefoundation.TypeID {
	rv := C.LayerGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Returns the type identifier for Core Graphics function objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1390879-cgfunctiongettypeid?language=objc
func FunctionGetTypeID() corefoundation.TypeID {
	rv := C.FunctionGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Gets the array of PDF content streams contained in a PDF content stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410048-cgpdfcontentstreamgetstreams?language=objc
func PDFContentStreamGetStreams(cs unsafe.Pointer) corefoundation.ArrayRef {
	rv := C.PDFContentStreamGetStreams(
		// *typing.RefType
		unsafe.Pointer(cs),
	)
	// *typing.RefType
	return corefoundation.ArrayRef(rv)
}

// Returns the number of entries in a PDF dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430218-cgpdfdictionarygetcount?language=objc
func PDFDictionaryGetCount(dict unsafe.Pointer) uint {
	rv := C.PDFDictionaryGetCount(
		// *typing.RefType
		unsafe.Pointer(dict),
	)
	// *typing.PrimitiveType
	return uint(rv)
}

// Decrements the hide cursor count, and shows the mouse cursor if the count is 0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455867-cgdisplayshowcursor?language=objc
func DisplayShowCursor(display DirectDisplayID) Error {
	rv := C.DisplayShowCursor(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return Error(rv)
}

// Obtains the bitmap information associated with a bitmap graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455839-cgbitmapcontextgetbitmapinfo?language=objc
func BitmapContextGetBitmapInfo(context ContextRef) BitmapInfo {
	rv := C.BitmapContextGetBitmapInfo(
		// *typing.RefType
		unsafe.Pointer(context),
	)
	// *typing.AliasType
	return BitmapInfo(rv)
}

// Returns an affine transformation matrix constructed by scaling an existing affine transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455882-cgaffinetransformscale?language=objc
func AffineTransformScale(t AffineTransform, sx float64, sy float64) AffineTransform {
	rv := C.AffineTransformScale(
		// *typing.StructType
		*(*C.CGAffineTransform)(unsafe.Pointer(&t)),
		// *typing.PrimitiveType
		C.float(sx),
		// *typing.PrimitiveType
		C.float(sy),
	)
	// *typing.StructType
	return *(*AffineTransform)(unsafe.Pointer(&rv))
}

// Creates a subset of the font in the specified PostScript format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396324-cgfontcreatepostscriptsubset?language=objc
func FontCreatePostScriptSubset(font FontRef, subsetName corefoundation.StringRef, format FontPostScriptFormat, glyphs *Glyph, count uint, encoding *Glyph) corefoundation.DataRef {
	rv := C.FontCreatePostScriptSubset(
		// *typing.RefType
		unsafe.Pointer(font),
		// *typing.RefType
		unsafe.Pointer(subsetName),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGFontPostScriptFormat)(format),
		// *typing.PointerType
		(*C.CGGlyph)(unsafe.Pointer(&glyphs)),
		// *typing.PrimitiveType
		C.uint(count),
		// *typing.PointerType
		(*C.CGGlyph)(unsafe.Pointer(&encoding)),
	)
	// *typing.RefType
	return corefoundation.DataRef(rv)
}

// Returns a Boolean value indicating whether a display is the main display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455448-cgdisplayismain?language=objc
func DisplayIsMain(display DirectDisplayID) int {
	rv := C.DisplayIsMain(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGDirectDisplayID)(display),
	)
	// *typing.AliasType
	return int(rv)
}

// Creates a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411209-cgpathcreatemutable?language=objc
func PathCreateMutable() MutablePathRef {
	rv := C.PathCreateMutable()
	// *typing.RefType
	return MutablePathRef(rv)
}

// Restores the gamma tables to the values in the user’s ColorSync display profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454702-cgdisplayrestorecolorsyncsetting?language=objc
func DisplayRestoreColorSyncSettings() {
	C.DisplayRestoreColorSyncSettings()
}

// Sets the pattern for dashed lines in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455911-cgcontextsetlinedash?language=objc
func ContextSetLineDash(c ContextRef, phase float64, lengths *float64, count uint) {
	C.ContextSetLineDash(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.PrimitiveType
		C.float(phase),
		// *typing.PointerType
		(*C.float)(unsafe.Pointer(&lengths)),
		// *typing.PrimitiveType
		C.uint(count),
	)
}

// Returns the type identifier of Quartz display modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456191-cgdisplaymodegettypeid?language=objc
func DisplayModeGetTypeID() corefoundation.TypeID {
	rv := C.DisplayModeGetTypeID()
	// *typing.AliasType
	return corefoundation.TypeID(rv)
}

// Returns the current flags of a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408792-cgeventsourceflagsstate?language=objc
func EventSourceFlagsState(stateID EventSourceStateID) EventFlags {
	rv := C.EventSourceFlagsState(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CGEventSourceStateID)(stateID),
	)
	// *typing.AliasType
	return EventFlags(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656522-cgrequestposteventaccess?language=objc
func RequestPostEventAccess() bool {
	rv := C.RequestPostEventAccess()
	// *typing.PrimitiveType
	return bool(rv)
}

// Paints the area contained within the provided rectangle, using the fill color in the current graphics state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454700-cgcontextfillrect?language=objc
func ContextFillRect(c ContextRef, rect Rect) {
	C.ContextFillRect(
		// *typing.RefType
		unsafe.Pointer(c),
		// *typing.StructType
		*(*C.CGRect)(unsafe.Pointer(&rect)),
	)
}

// Increments the retain count of a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411181-cgpathretain?language=objc
func PathRetain(path unsafe.Pointer) unsafe.Pointer {
	rv := C.PathRetain(
		// *typing.RefType
		unsafe.Pointer(path),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns whether the specified PDF document is currently unlocked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402607-cgpdfdocumentisunlocked?language=objc
func PDFDocumentIsUnlocked(document PDFDocumentRef) bool {
	rv := C.PDFDocumentIsUnlocked(
		// *typing.RefType
		unsafe.Pointer(document),
	)
	// *typing.PrimitiveType
	return bool(rv)
}
