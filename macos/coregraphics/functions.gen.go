// AUTO-GENERATED CODE, DO NOT MODIFY

package coregraphics

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import "CoreGraphics/CoreGraphics.h"
// int32_t PostKeyboardEvent(uint16_t keyChar, uint16_t virtualKey, NSInteger keyDown);
// void ContextStrokeEllipseInRect(void * c, CGRect rect);
// CGRect ContextConvertRectToDeviceSpace(void * c, CGRect rect);
// void EventSetSource(void * event, void * source);
// void ContextSetFontSize(void * c, float size);
// void ContextClipToRects(void * c, CGRect* rects, NSUInteger count);
// CGRect RectApplyAffineTransform(CGRect rect, CGAffineTransform t);
// int64_t EventSourceGetUserData(void * source);
// uint32_t ImageGetAlphaInfo(void * image);
// int DisplayIsOnline(uint32_t display);
// int32_t ReleaseDisplayFadeReservation(uint32_t token);
// bool ColorSpaceUsesExtendedRange(void * space);
// uint BitmapContextGetBitsPerComponent(void * context);
// void * PDFContextCreate(void * consumer, CGRect* mediaBox, void * auxiliaryInfo);
// void * ColorSpaceGetName(void * space);
// void ContextBeginTransparencyLayer(void * c, void * auxiliaryInfo);
// void * ColorGetColorSpace(void * color);
// int64_t EventGetIntegerValueField(void * event, uint32_t field);
// void ContextSetTextPosition(void * c, float x, float y);
// void * ColorSpaceCopyICCData(void * space);
// void ContextSetFont(void * c, void * font);
// void ContextDrawPDFDocument(void * c, CGRect rect, void * document, NSInteger page);
// CGRect ContextConvertRectToUserSpace(void * c, CGRect rect);
// void * ColorSpaceCreateWithPropertyList(void* plist);
// void * ColorRetain(void * color);
// void ContextBeginPath(void * c);
// bool ColorSpaceIsHLGBased(void * s);
// void * ColorCreateGenericGray(float gray, float alpha);
// void DataProviderRelease(void * provider);
// void PathAddArc(void * path, CGAffineTransform* m, float x, float y, float radius, float startAngle, float endAngle, BOOL clockwise);
// int32_t EventGetTypeID();
// void ContextDrawLinearGradient(void * c, void * gradient, CGPoint startPoint, CGPoint endPoint, uint32_t options);
// CGRect PDFPageGetBoxRect(void * page, int32_t box);
// void ContextStrokeRectWithWidth(void * c, CGRect rect, float width);
// void * SessionCopyCurrentDictionary();
// int32_t ShadingGetTypeID();
// void * ColorSpaceCreatePattern(void * baseSpace);
// CGRect DisplayBounds(uint32_t display);
// void ContextSetTextDrawingMode(void * c, int32_t mode);
// void * ColorSpaceCreateDeviceGray();
// void * ColorSpaceCreateCalibratedGray(float* whitePoint, float* blackPoint, float gamma);
// void * PathCreateCopyByStrokingPath(void * path, CGAffineTransform* transform, float lineWidth, int32_t lineCap, int32_t lineJoin, float miterLimit);
// bool ColorEqualToColor(void * color1, void * color2);
// bool PathContainsPoint(void * path, CGAffineTransform* m, CGPoint point, BOOL eoFill);
// void * FontCopyVariations(void * font);
// void * FontRetain(void * font);
// void * DisplayCopyColorSpace(uint32_t display);
// uint ImageGetWidth(void * image);
// void * PDFDocumentCreateWithURL(void * url);
// void EventSetFlags(void * event, uint64_t flags);
// CGRect RectMake(float x, float y, float width, float height);
// uint ColorGetNumberOfComponents(void * color);
// void * PathCreateWithRect(CGRect rect, CGAffineTransform* transform);
// CGPoint ContextGetPathCurrentPoint(void * c);
// void PathCloseSubpath(void * path);
// void PathAddPath(void * path1, CGAffineTransform* m, void * path2);
// bool RequestScreenCaptureAccess();
// int32_t DisplayStreamGetTypeID();
// bool SizeMakeWithDictionaryRepresentation(void * dict, CGSize* size);
// bool ColorSpaceIsWideGamutRGB(void * );
// void PDFOperatorTableRelease(void * table);
// int FontGetUnitsPerEm(void * font);
// void * PDFContentStreamRetain(void * cs);
// void ContextSetAlpha(void * c, float alpha);
// void ContextAddCurveToPoint(void * c, float cp1x, float cp1y, float cp2x, float cp2y, float x, float y);
// void * ColorCreateGenericCMYK(float cyan, float magenta, float yellow, float black, float alpha);
// void * ColorSpaceCreateExtendedLinearized(void * space);
// void * DisplayIOServicePort(uint32_t display);
// void * ColorSpaceCreateWithICCData(void* data);
// int32_t GetDisplaysWithRect(CGRect rect, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount);
// int32_t PostMouseEvent(CGPoint mouseCursorPosition, NSInteger updateMouseCursorPosition, uint32_t buttonCount, NSInteger mouseButtonDown);
// uint DisplayModeGetWidth(void * mode);
// void ContextSaveGState(void * c);
// int FontGetCapHeight(void * font);
// uint PDFArrayGetCount(void * array);
// void FunctionRelease(void * function);
// bool PDFDocumentUnlockWithPassword(void * document, uint8_t* password);
// double DisplayRotation(uint32_t display);
// CGAffineTransform AffineTransformMakeScale(float sx, float sy);
// void ContextDrawLayerInRect(void * context, CGRect rect, void * layer);
// bool RectIsInfinite(CGRect rect);
// void ContextShowGlyphsWithAdvances(void * c, CGGlyph* glyphs, CGSize* advances, NSUInteger count);
// void * ImageCreateCopyWithColorSpace(void * image, void * space);
// bool PreflightPostEventAccess();
// void ContextRelease(void * c);
// int32_t DataConsumerGetTypeID();
// int32_t CompleteDisplayConfiguration(void * config, uint32_t option);
// void PDFContextEndTag(void * context);
// void * ColorSpaceCopyICCProfile(void * space);
// void * ColorCreateGenericGrayGamma2_2(float gray, float alpha);
// CGRect PathGetBoundingBox(void * path);
// void PDFDocumentGetVersion(void * document, NSInteger* majorVersion, NSInteger* minorVersion);
// void PathAddRoundedRect(void * path, CGAffineTransform* transform, CGRect rect, float cornerWidth, float cornerHeight);
// void * DisplayAvailableModes(uint32_t dsp);
// int32_t WindowLevelForKey(int32_t key);
// void ContextSetRenderingIntent(void * c, int32_t intent);
// void ContextSetShadowWithColor(void * c, CGSize offset, float blur, void * color);
// int32_t PatternGetTypeID();
// void * DisplayModeRetain(void * mode);
// uint64_t EventGetTimestamp(void * event);
// void * PathCreateMutableCopy(void * path);
// int PDFDocumentGetRotationAngle(void * document, NSInteger page);
// CGAffineTransform ContextGetCTM(void * c);
// void * EventCreateData(void * allocator, void * event);
// int DisplayIsStereo(uint32_t display);
// void ContextConcatCTM(void * c, CGAffineTransform transform);
// CGSize ContextConvertSizeToUserSpace(void * c, CGSize size);
// bool PDFScannerPopInteger(void * scanner, CGPDFInteger* value);
// void PDFContextAddDocumentMetadata(void * context, void * metadata);
// CGSize DisplayScreenSize(uint32_t display);
// void ColorSpaceGetColorTable(void * space, uint8_t* table);
// void * PDFPageRetain(void * page);
// void ContextFillEllipseInRect(void * c, CGRect rect);
// void * DisplayCopyDisplayMode(uint32_t display);
// void PDFContextSetDestinationForRect(void * context, void * name, CGRect rect);
// int32_t GetDisplaysWithOpenGLDisplayMask(uint32_t mask, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount);
// bool PDFDictionaryGetArray(void * dict, uint8_t* key, void * value);
// void * ShadingCreateAxial(void * space, CGPoint start, CGPoint end, void * function, BOOL extendStart, BOOL extendEnd);
// void ContextSetShadow(void * c, CGSize offset, float blur);
// bool PointMakeWithDictionaryRepresentation(void * dict, CGPoint* point);
// void EventPostToPid(NSObject* pid, void * event);
// CGPoint ContextGetTextPosition(void * c);
// uint BitmapContextGetBitsPerPixel(void * context);
// void ContextAddArc(void * c, float x, float y, float radius, float startAngle, float endAngle, NSInteger clockwise);
// void EventSetType(void * event, uint32_t type);
// void * EventCreateScrollWheelEvent2(void * source, uint32_t units, uint32_t wheelCount, int32_t wheel1, int32_t wheel2, int32_t wheel3);
// int32_t ColorSpaceGetModel(void * space);
// void * ImageCreateWithPNGDataProvider(void * source, float* decode, BOOL shouldInterpolate, int32_t intent);
// void PathAddEllipseInRect(void * path, CGAffineTransform* m, CGRect rect);
// void * EventCreateFromData(void * allocator, void * data);
// CGRect PDFDocumentGetArtBox(void * document, NSInteger page);
// void * ColorSpaceCreateDeviceCMYK();
// int32_t ConfigureDisplayFadeEffect(void * config, float fadeOutSeconds, float fadeInSeconds, float fadeRed, float fadeGreen, float fadeBlue);
// int32_t DisplayFade(uint32_t token, float duration, float startBlend, float endBlend, float redBlend, float greenBlend, float blueBlend, NSInteger synchronous);
// bool SizeEqualToSize(CGSize size1, CGSize size2);
// void ContextSetShouldSubpixelQuantizeFonts(void * c, BOOL shouldSubpixelQuantizeFonts);
// CGPoint EventGetLocation(void * event);
// bool PSConverterAbort(void * converter);
// int32_t DisplayStreamUpdateGetTypeID();
// void ContextScaleCTM(void * c, float sx, float sy);
// void * SizeCreateDictionaryRepresentation(CGSize size);
// bool EventSourceButtonState(int32_t stateID, uint32_t button);
// void EventTapEnable(void * tap, BOOL enable);
// CGAffineTransform AffineTransformMakeRotation(float angle);
// void * ColorSpaceCreateICCBased(NSUInteger nComponents, float* range, void * profile, void * alternate);
// int32_t AssociateMouseAndMouseCursorPosition(NSInteger connected);
// void * PDFDocumentRetain(void * document);
// CGPoint PathGetCurrentPoint(void * path);
// void ContextSetRGBStrokeColor(void * c, float red, float green, float blue, float alpha);
// bool PreflightScreenCaptureAccess();
// int32_t SetDisplayTransferByTable(uint32_t display, uint32_t tableSize, CGGammaValue* redTable, CGGammaValue* greenTable, CGGammaValue* blueTable);
// void ContextSetFlatness(void * c, float flatness);
// int DisplayIsActive(uint32_t display);
// void ContextSetPatternPhase(void * c, CGSize phase);
// void PDFScannerRelease(void * scanner);
// void ContextSetLineJoin(void * c, int32_t join);
// void * FontCopyTableTags(void * font);
// int32_t GetActiveDisplayList(uint32_t maxDisplays, CGDirectDisplayID* activeDisplays, uint32_t* displayCount);
// bool PathIsRect(void * path, CGRect* rect);
// void * DataProviderCreateWithURL(void * url);
// void * EventSourceCreate(int32_t stateID);
// int32_t ContextGetInterpolationQuality(void * c);
// void * ColorConversionInfoCreate(void * src, void * dst);
// CGRect PDFDocumentGetTrimBox(void * document, NSInteger page);
// bool RectContainsRect(CGRect rect1, CGRect rect2);
// void * FontCreatePostScriptEncoding(void * font, CGGlyph* encoding);
// void ContextDrawLayerAtPoint(void * context, CGPoint point, void * layer);
// bool PDFArrayGetStream(void * array, NSUInteger index, void * value);
// void ContextShowTextAtPoint(void * c, float x, float y, uint8_t* string, NSUInteger length);
// void EventSourceSetPixelsPerLine(void * source, double pixelsPerLine);
// void * EventCreateKeyboardEvent(void * source, uint16_t virtualKey, BOOL keyDown);
// uint32_t EventSourceCounterForEventType(int32_t stateID, uint32_t eventType);
// void PathAddRects(void * path, CGAffineTransform* m, CGRect* rects, NSUInteger count);
// void EventSetTimestamp(void * event, uint64_t timestamp);
// bool EventSourceKeyState(int32_t stateID, uint16_t key);
// float RectGetHeight(CGRect rect);
// void ContextShowGlyphsAtPoint(void * c, float x, float y, CGGlyph* glyphs, NSUInteger count);
// bool PDFDictionaryGetInteger(void * dict, uint8_t* key, CGPDFInteger* value);
// double EventSourceSecondsSinceLastEventType(int32_t stateID, uint32_t eventType);
// CGRect RectIntegral(CGRect rect);
// void ContextSetFillColorSpace(void * c, void * space);
// bool PDFArrayGetNumber(void * array, NSUInteger index, CGPDFReal* value);
// void * ColorCreateWithPattern(void * space, void * pattern, float* components);
// void ContextTranslateCTM(void * c, float tx, float ty);
// int DisplayUsesOpenGLAcceleration(uint32_t display);
// void ContextSetBlendMode(void * c, int32_t mode);
// void GetLastMouseDelta(int32_t* deltaX, int32_t* deltaY);
// bool FontGetGlyphBBoxes(void * font, CGGlyph* glyphs, NSUInteger count, CGRect* bboxes);
// int32_t GetDisplaysWithPoint(CGPoint point, uint32_t maxDisplays, CGDirectDisplayID* displays, uint32_t* matchingDisplayCount);
// void ImageRelease(void * image);
// float RectGetMinX(CGRect rect);
// void ContextReplacePathWithStrokedPath(void * c);
// uint BitmapContextGetBytesPerRow(void * context);
// void * WindowListCreate(uint32_t option, uint32_t relativeToWindow);
// CGSize SizeApplyAffineTransform(CGSize size, CGAffineTransform t);
// void EventTapPostEvent(void * proxy, void * event);
// int DisplayIsInMirrorSet(uint32_t display);
// bool ColorSpaceIsPQBased(void * s);
// void ContextClearRect(void * c, CGRect rect);
// void EventKeyboardSetUnicodeString(void * event, NSObject* stringLength, UniChar* unicodeString);
// int32_t DisplayRelease(uint32_t display);
// void * PDFPageGetDocument(void * page);
// int32_t GetDisplayTransferByTable(uint32_t display, uint32_t capacity, CGGammaValue* redTable, CGGammaValue* greenTable, CGGammaValue* blueTable, uint32_t* sampleCount);
// bool AffineTransformEqualToTransform(CGAffineTransform t1, CGAffineTransform t2);
// void * PDFScannerGetContentStream(void * scanner);
// void PathAddLineToPoint(void * path, CGAffineTransform* m, float x, float y);
// bool PreflightListenEventAccess();
// uint32_t EventGetType(void * event);
// bool PDFScannerPopNumber(void * scanner, CGPDFReal* value);
// void * ColorSpaceCreateLab(float* whitePoint, float* blackPoint, float* range);
// void PDFContextEndPage(void * context);
// CGRect ContextGetClipBoundingBox(void * c);
// void * PDFOperatorTableCreate();
// void PathAddRect(void * path, CGAffineTransform* m, CGRect rect);
// CGPoint PointApplyAffineTransform(CGPoint point, CGAffineTransform t);
// void * ImageCreateWithJPEGDataProvider(void * source, float* decode, BOOL shouldInterpolate, int32_t intent);
// int32_t EventSourceGetTypeID();
// void * ImageCreate(NSUInteger width, NSUInteger height, NSUInteger bitsPerComponent, NSUInteger bitsPerPixel, NSUInteger bytesPerRow, void * space, uint32_t bitmapInfo, void * provider, float* decode, BOOL shouldInterpolate, int32_t intent);
// bool RectIsNull(CGRect rect);
// int32_t EventSourceGetSourceStateID(void * source);
// void * GradientRetain(void * gradient);
// bool PDFArrayGetBoolean(void * array, NSUInteger index, CGPDFBoolean* value);
// void ContextSetFillColor(void * c, float* components);
// void * ColorGetPattern(void * color);
// CGRect PathGetPathBoundingBox(void * path);
// int32_t CaptureAllDisplaysWithOptions(uint32_t options);
// void * WindowListCreateDescriptionFromArray(void * windowArray);
// void * PDFContentStreamCreateWithStream(void * stream, void * streamResources, void * parent);
// void ContextDrawRadialGradient(void * c, void * gradient, CGPoint startCenter, float startRadius, CGPoint endCenter, float endRadius, uint32_t options);
// void * PDFStringCopyDate(void * string);
// void * WindowServerCFMachPort();
// int32_t GetOnlineDisplayList(uint32_t maxDisplays, CGDirectDisplayID* onlineDisplays, uint32_t* displayCount);
// CGPoint ContextConvertPointToDeviceSpace(void * c, CGPoint point);
// void * PathCreateCopy(void * path);
// void * BitmapContextCreateImage(void * context);
// uint ColorSpaceGetNumberOfComponents(void * space);
// void * FunctionRetain(void * function);
// int32_t GetDisplayTransferByFormula(uint32_t display, CGGammaValue* redMin, CGGammaValue* redMax, CGGammaValue* redGamma, CGGammaValue* greenMin, CGGammaValue* greenMax, CGGammaValue* greenGamma, CGGammaValue* blueMin, CGGammaValue* blueMax, CGGammaValue* blueGamma);
// bool PDFScannerPopStream(void * scanner, void * value);
// void * ColorSpaceRetain(void * space);
// void * ColorCreateSRGB(float red, float green, float blue, float alpha);
// CGRect* DisplayStreamUpdateGetRects(void * updateRef, int32_t rectType, NSUInteger* rectCount);
// void * ImageGetColorSpace(void * image);
// int32_t SetLocalEventsFilterDuringSuppressionState(uint32_t filter, uint32_t state);
// void * ShadingCreateRadial(void * space, CGPoint start, float startRadius, CGPoint end, float endRadius, void * function, BOOL extendStart, BOOL extendEnd);
// void RectDivide(CGRect rect, CGRect* slice, CGRect* remainder, float amount, uint32_t edge);
// void ContextSetAllowsAntialiasing(void * c, BOOL allowsAntialiasing);
// uint BitmapContextGetHeight(void * context);
// float ColorGetAlpha(void * color);
// uint ImageGetHeight(void * image);
// void * ColorSpaceCreateWithName(void * name);
// void ContextAddQuadCurveToPoint(void * c, float cpx, float cpy, float x, float y);
// int32_t ConfigureDisplayWithDisplayMode(void * config, uint32_t display, void * mode, void * options);
// void * DataProviderCopyData(void * provider);
// int DisplayIsInHWMirrorSet(uint32_t display);
// uint32_t EventSourceGetLocalEventsFilterDuringSuppressionState(void * source, uint32_t state);
// void * FontCopyTableForTag(void * font, uint32_t tag);
// uint8_t* PDFStringGetBytePtr(void * string);
// CGRect ContextGetPathBoundingBox(void * c);
// void EventKeyboardGetUnicodeString(void * event, NSObject* maxStringLength, NSObject* actualStringLength, UniChar* unicodeString);
// void * LayerRetain(void * layer);
// void ContextSetAllowsFontSubpixelPositioning(void * c, BOOL allowsFontSubpixelPositioning);
// int32_t ColorGetTypeID();
// uint32_t ImageGetByteOrderInfo(void * image);
// void * PDFDocumentCreateWithProvider(void * provider);
// bool FontGetGlyphAdvances(void * font, CGGlyph* glyphs, NSUInteger count, NSInteger* advances);
// uint32_t ImageGetPixelFormatInfo(void * image);
// void * FontCopyVariationAxes(void * font);
// CGRect PDFDocumentGetMediaBox(void * document, NSInteger page);
// void* ColorSpaceCopyPropertyList(void * space);
// void ContextSetStrokePattern(void * c, void * pattern, float* components);
// void * DisplayBestModeForParameters(uint32_t display, NSUInteger bitsPerPixel, NSUInteger width, NSUInteger height, boolean_t* exactMatch);
// bool PDFDictionaryGetDictionary(void * dict, uint8_t* key, void * value);
// void ContextClipToRect(void * c, CGRect rect);
// void ContextResetClip(void * c);
// bool PDFArrayGetDictionary(void * array, NSUInteger index, void * value);
// void PathAddQuadCurveToPoint(void * path, CGAffineTransform* m, float cpx, float cpy, float x, float y);
// int32_t WarpMouseCursorPosition(CGPoint newCursorPosition);
// int32_t DisplayMoveCursorToPoint(uint32_t display, CGPoint point);
// void * WindowListCopyWindowInfo(uint32_t option, uint32_t relativeToWindow);
// uint32_t ShieldingWindowID(uint32_t display);
// void PDFContextSetOutline(void * context, void * outline);
// void PathAddArcToPoint(void * path, CGAffineTransform* m, float x1, float y1, float x2, float y2, float radius);
// uint ImageGetBitsPerComponent(void * image);
// int32_t CaptureAllDisplays();
// double EventSourceGetPixelsPerLine(void * source);
// CGRect RectInset(CGRect rect, float dx, float dy);
// int32_t BeginDisplayConfiguration(void * config);
// bool PDFArrayGetNull(void * array, NSUInteger index);
// float* ImageGetDecode(void * image);
// void* DataProviderGetInfo(void * provider);
// void * DisplayStreamUpdateCreateMergedUpdate(void * firstUpdate, void * secondUpdate);
// void * PDFStreamCopyData(void * stream, CGPDFDataFormat* format);
// void * PatternRetain(void * pattern);
// bool PDFDictionaryGetObject(void * dict, uint8_t* key, void * value);
// void * BitmapContextGetColorSpace(void * context);
// bool RectIsEmpty(CGRect rect);
// float RectGetMinY(CGRect rect);
// void * WindowListCreateImage(CGRect screenBounds, uint32_t listOption, uint32_t windowID, uint32_t imageOption);
// int32_t PostScrollWheelEvent(uint32_t wheelCount, int32_t wheel1);
// int32_t ColorSpaceGetTypeID();
// void * ShadingRetain(void * shading);
// void * ColorCreateCopy(void * color);
// void ContextBeginTransparencyLayerWithRect(void * c, CGRect rect, void * auxInfo);
// void PDFPageRelease(void * page);
// int32_t ImageGetTypeID();
// bool PDFScannerPopDictionary(void * scanner, void * value);
// void ContextRotateCTM(void * c, float angle);
// void ContextEndTransparencyLayer(void * c);
// void EventSourceSetLocalEventsSuppressionInterval(void * source, double seconds);
// CGRect RectIntersection(CGRect r1, CGRect r2);
// int32_t DisplayModeGetIODisplayModeID(void * mode);
// uint32_t DisplayPrimaryDisplay(uint32_t display);
// void PDFContextAddDestinationAtPoint(void * context, void * name, CGPoint point);
// int FontGetDescent(void * font);
// void ContextSetStrokeColor(void * c, float* components);
// bool PDFDictionaryGetString(void * dict, uint8_t* key, void * value);
// void DisplayModeRelease(void * mode);
// void ContextAddRects(void * c, CGRect* rects, NSUInteger count);
// void ContextSelectFont(void * c, uint8_t* name, float size, int32_t textEncoding);
// void * PathCreateMutableCopyByTransformingPath(void * path, CGAffineTransform* transform);
// CGAffineTransform AffineTransformMakeTranslation(float tx, float ty);
// void ContextSetCharacterSpacing(void * c, float spacing);
// uint32_t DisplayModeGetIOFlags(void * mode);
// void PDFContextClose(void * context);
// void * DataProviderRetain(void * provider);
// bool ColorSpaceSupportsOutput(void * space);
// void EventSetDoubleValueField(void * event, uint32_t field, double value);
// void ContextDrawPDFPage(void * c, void * page);
// void * PDFContextCreateWithURL(void * url, CGRect* mediaBox, void * auxiliaryInfo);
// int32_t PDFPageGetTypeID();
// CGRect PDFDocumentGetBleedBox(void * document, NSInteger page);
// void * PDFContentStreamCreateWithPage(void * page);
// void * FontCreateWithDataProvider(void * provider);
// bool PDFDictionaryGetNumber(void * dict, uint8_t* key, CGPDFReal* value);
// void ColorSpaceRelease(void * space);
// void * ImageCreateCopy(void * image);
// void ContextEndPage(void * c);
// bool ColorSpaceIsHDR(void * );
// void * PDFDocumentGetID(void * document);
// void ContextSetTextMatrix(void * c, CGAffineTransform t);
// uint FontGetNumberOfGlyphs(void * font);
// void ContextClipToMask(void * c, CGRect rect, void * mask);
// void * PDFPageGetDictionary(void * page);
// uint32_t DisplayGammaTableCapacity(uint32_t display);
// void EventSetIntegerValueField(void * event, uint32_t field, int64_t value);
// void ContextAddPath(void * c, void * path);
// CGAffineTransform ContextGetTextMatrix(void * c);
// void ContextStrokePath(void * c);
// uint DisplayPixelsWide(uint32_t display);
// void ContextSetShouldAntialias(void * c, BOOL shouldAntialias);
// int32_t DisplaySwitchToMode(uint32_t display, void * mode);
// int32_t GetEventTapList(uint32_t maxNumberOfTaps, CGEventTapInformation* tapList, uint32_t* eventTapCount);
// bool EventTapIsEnabled(void * tap);
// bool PDFDictionaryGetStream(void * dict, uint8_t* key, void * value);
// void * ColorSpaceCreateIndexed(void * baseSpace, NSUInteger lastIndex, uint8_t* colorTable);
// uint8_t* PDFTagTypeGetName(int32_t tagType);
// void ContextSetGrayFillColor(void * c, float gray, float alpha);
// void * PDFContentStreamGetResource(void * cs, uint8_t* category, uint8_t* name);
// bool PDFDictionaryGetBoolean(void * dict, uint8_t* key, CGPDFBoolean* value);
// bool PSConverterConvert(void * converter, void * provider, void * consumer, void * options);
// void PDFContentStreamRelease(void * cs);
// CGRect PDFDocumentGetCropBox(void * document, NSInteger page);
// uint32_t DisplayIDToOpenGLDisplayMask(uint32_t display);
// float RectGetMaxX(CGRect rect);
// bool ContextIsPathEmpty(void * c);
// int32_t ColorConversionInfoGetTypeID();
// uint ImageGetBitsPerPixel(void * image);
// bool RectMakeWithDictionaryRepresentation(void * dict, CGRect* rect);
// uint DisplayModeGetPixelHeight(void * mode);
// void * ColorSpaceCreateDeviceRGB();
// void ContextSetStrokeColorWithColor(void * c, void * color);
// void * LayerGetContext(void * layer);
// bool PDFArrayGetInteger(void * array, NSUInteger index, CGPDFInteger* value);
// float FontGetItalicAngle(void * font);
// int32_t ConfigureDisplayOrigin(void * config, uint32_t display, int32_t x, int32_t y);
// void * ColorConversionInfoCreateWithOptions(void * src, void * dst, void * options);
// int DisplayFadeOperationInProgress();
// void RestorePermanentDisplayConfiguration();
// uint32_t BitmapContextGetAlphaInfo(void * context);
// void ContextClosePath(void * c);
// bool PathEqualToPath(void * path1, void * path2);
// void * EventCreateMouseEvent(void * source, uint32_t mouseType, CGPoint mouseCursorPosition, uint32_t mouseButton);
// uint32_t MainDisplayID();
// CGAffineTransform AffineTransformTranslate(CGAffineTransform t, float tx, float ty);
// void * PDFDocumentGetPage(void * document, NSUInteger pageNumber);
// float RectGetMidX(CGRect rect);
// uint32_t EventSourceGetKeyboardType(void * source);
// void PDFContextBeginPage(void * context, void * pageInfo);
// void * ColorCreateCopyWithAlpha(void * color, float alpha);
// uint32_t PDFDocumentGetAccessPermissions(void * document);
// CGAffineTransform AffineTransformRotate(CGAffineTransform t, float angle);
// void * EventCreateCopy(void * event);
// void ContextSynchronize(void * c);
// CGAffineTransform PDFPageGetDrawingTransform(void * page, int32_t box, CGRect rect, NSInteger rotate, BOOL preserveAspectRatio);
// int32_t ShieldingWindowLevel();
// bool DisplayModeIsUsableForDesktopGUI(void * mode);
// int32_t GradientGetTypeID();
// void * PointCreateDictionaryRepresentation(CGPoint point);
// CGRect RectOffset(CGRect rect, float dx, float dy);
// bool PDFArrayGetName(void * array, NSUInteger index, uint8_t* value);
// void ContextSetCMYKStrokeColor(void * c, float cyan, float magenta, float yellow, float black, float alpha);
// void * ColorSpaceGetBaseColorSpace(void * space);
// void * LayerCreateWithContext(void * context, CGSize size, void * auxiliaryInfo);
// uint DisplayModeGetPixelWidth(void * mode);
// void * ImageCreateWithMaskingColors(void * image, float* components);
// void ContextFillRects(void * c, CGRect* rects, NSUInteger count);
// void * FontCopyPostScriptName(void * font);
// bool ImageIsMask(void * image);
// void ContextSetStrokeColorSpace(void * c, void * space);
// CGRect RectStandardize(CGRect rect);
// CGRect FontGetFontBBox(void * font);
// bool PDFDictionaryGetName(void * dict, uint8_t* key, uint8_t* value);
// bool PDFScannerPopObject(void * scanner, void * value);
// void * DisplayStreamGetRunLoopSource(void * displayStream);
// int32_t DisplayCapture(uint32_t display);
// void ContextAddEllipseInRect(void * c, CGRect rect);
// void * GradientCreateWithColorComponents(void * space, float* components, float* locations, NSUInteger count);
// bool RectEqualToRect(CGRect rect1, CGRect rect2);
// void * ColorSpaceCreateCalibratedRGB(float* whitePoint, float* blackPoint, float* gamma, float* matrix);
// void ContextDrawShading(void * c, void * shading);
// float RectGetMidY(CGRect rect);
// void * DisplayCreateImageForRect(uint32_t display, CGRect rect);
// void * ColorCreate(void * space, float* components);
// void PathMoveToPoint(void * path, CGAffineTransform* m, float x, float y);
// void ColorRelease(void * color);
// uint PDFDocumentGetNumberOfPages(void * document);
// void EventSourceSetKeyboardType(void * source, uint32_t keyboardType);
// void * DisplayCopyAllDisplayModes(uint32_t display, void * options);
// uint DisplayPixelsHigh(uint32_t display);
// void DisplayStreamUpdateGetMovedRectsDelta(void * updateRef, float* dx, float* dy);
// int32_t SetDisplayTransferByByteTable(uint32_t display, uint32_t tableSize, uint8_t* redTable, uint8_t* greenTable, uint8_t* blueTable);
// void * ImageGetDataProvider(void * image);
// int DisplayIsCaptured(uint32_t display);
// void ContextSetFillPattern(void * c, void * pattern, float* components);
// int32_t WaitForScreenRefreshRects(CGRect* rects, uint32_t* count);
// int32_t DataProviderGetTypeID();
// void * DataProviderCreateWithFilename(uint8_t* filename);
// bool PDFScannerPopString(void * scanner, void * value);
// uint ImageGetBytesPerRow(void * image);
// void * DataProviderCreateWithCFData(void * data);
// bool PDFDocumentAllowsCopying(void * document);
// void * ColorCreateCopyByMatchingToColorSpace(void * , int32_t intent, void * color, void * options);
// int32_t InhibitLocalEvents(NSInteger inhibit);
// CGSize SizeMake(float width, float height);
// float RectGetMaxY(CGRect rect);
// bool PDFDocumentAllowsPrinting(void * document);
// void ContextSetAllowsFontSmoothing(void * c, BOOL allowsFontSmoothing);
// void ContextDrawPath(void * c, int32_t mode);
// int FontGetGlyphWithGlyphName(void * font, void * name);
// bool PDFScannerPopArray(void * scanner, void * value);
// void * RectCreateDictionaryRepresentation(CGRect );
// void* BitmapContextGetData(void * context);
// int FontGetAscent(void * font);
// void * DisplayModeCopyPixelEncoding(void * mode);
// int32_t DisplaySetStereoOperation(uint32_t display, NSInteger stereo, NSInteger forceBlueLine, uint32_t option);
// bool ImageGetShouldInterpolate(void * image);
// void ContextFillPath(void * c);
// int32_t PDFDocumentGetTypeID();
// bool PDFArrayGetString(void * array, NSUInteger index, void * value);
// void ContextShowText(void * c, uint8_t* string, NSUInteger length);
// int32_t ImageGetRenderingIntent(void * image);
// void EventPost(uint32_t tap, void * event);
// float FontGetStemV(void * font);
// int CursorIsVisible();
// void ContextStrokeLineSegments(void * c, CGPoint* points, NSUInteger count);
// void ContextEOFillPath(void * c);
// bool PDFArrayGetObject(void * array, NSUInteger index, void * value);
// int32_t DisplayCaptureWithOptions(uint32_t display, uint32_t options);
// bool ContextPathContainsPoint(void * c, CGPoint point, int32_t mode);
// bool ColorSpaceUsesITUR_2100TF(void * );
// void * PDFDocumentGetOutline(void * document);
// int32_t DisplayStreamStop(void * displayStream);
// void PathRelease(void * path);
// int32_t PathGetTypeID();
// void PDFDocumentRelease(void * document);
// void * WindowServerCreateServerPort();
// CGAffineTransform ContextGetUserSpaceToDeviceSpaceTransform(void * c);
// void ContextDrawImage(void * c, CGRect rect, void * image);
// void * PathCreateWithRoundedRect(CGRect rect, float cornerWidth, float cornerHeight, CGAffineTransform* transform);
// void PDFContextBeginTag(void * context, int32_t tagType, void * tagProperties);
// int32_t PDFObjectGetType(void * object);
// void * ImageMaskCreate(NSUInteger width, NSUInteger height, NSUInteger bitsPerComponent, NSUInteger bitsPerPixel, NSUInteger bytesPerRow, void * provider, float* decode, BOOL shouldInterpolate);
// void * ColorCreateGenericRGB(float red, float green, float blue, float alpha);
// double EventGetDoubleValueField(void * event, uint32_t field);
// void * PDFDocumentGetInfo(void * document);
// bool PSConverterIsConverting(void * converter);
// void * FontCreateCopyWithVariations(void * font, void * variations);
// int32_t ConfigureDisplayMirrorOfDisplay(void * config, uint32_t display, uint32_t master);
// double DisplayModeGetRefreshRate(void * mode);
// double EventSourceGetLocalEventsSuppressionInterval(void * source);
// void ContextAddArcToPoint(void * c, float x1, float y1, float x2, float y2, float radius);
// void * DataConsumerCreateWithURL(void * url);
// CGSize LayerGetSize(void * layer);
// void * ContextRetain(void * c);
// int32_t AcquireDisplayFadeReservation(float seconds, CGDisplayFadeReservationToken* token);
// uint64_t EventGetFlags(void * event);
// void * FontCopyGlyphNameForGlyph(void * font, NSInteger glyph);
// void ContextSetRGBFillColor(void * c, float red, float green, float blue, float alpha);
// void ContextRestoreGState(void * c);
// void * PathCreateCopyByTransformingPath(void * path, CGAffineTransform* transform);
// void * FontCopyFullName(void * font);
// CGAffineTransform AffineTransformConcat(CGAffineTransform t1, CGAffineTransform t2);
// uint PDFPageGetPageNumber(void * page);
// void EventSetLocation(void * event, CGPoint location);
// void FontRelease(void * font);
// void * PDFStringCopyTextString(void * string);
// bool PDFDocumentIsEncrypted(void * document);
// void ContextStrokeRect(void * c, CGRect rect);
// void ContextSetShouldSubpixelPositionFonts(void * c, BOOL shouldSubpixelPositionFonts);
// void * ContextCopyPath(void * c);
// void * ImageGetUTType(void * image);
// bool FontCanCreatePostScriptSubset(void * font, int32_t format);
// float* ColorGetComponents(void * color);
// int32_t ContextGetTypeID();
// void * DataConsumerCreateWithCFData(void * data);
// void ContextAddRect(void * c, CGRect rect);
// void PathAddRelativeArc(void * path, CGAffineTransform* matrix, float x, float y, float radius, float startAngle, float delta);
// void DataConsumerRelease(void * consumer);
// int32_t ConfigureDisplayStereoOperation(void * config, uint32_t display, NSInteger stereo, NSInteger forceBlueLine);
// CGSize ContextConvertSizeToDeviceSpace(void * c, CGSize size);
// void PathAddCurveToPoint(void * path, CGAffineTransform* m, float cp1x, float cp1y, float cp2x, float cp2y, float x, float y);
// int DisplayIsBuiltin(uint32_t display);
// int32_t DisplayHideCursor(uint32_t display);
// uint32_t DisplayUnitNumber(uint32_t display);
// int32_t DisplaySetDisplayMode(uint32_t display, void * mode, void * options);
// void ContextShowGlyphsAtPositions(void * c, CGGlyph* glyphs, CGPoint* Lpositions, NSUInteger count);
// void ContextSetFillColorWithColor(void * c, void * color);
// int FontGetLeading(void * font);
// int32_t PSConverterGetTypeID();
// void * GradientCreateWithColors(void * space, void * colors, float* locations);
// int DisplayIsAlwaysInMirrorSet(uint32_t display);
// void * PathCreateCopyByDashingPath(void * path, CGAffineTransform* transform, float phase, float* lengths, NSUInteger count);
// void ContextSetCMYKFillColor(void * c, float cyan, float magenta, float yellow, float black, float alpha);
// int32_t CancelDisplayConfiguration(void * config);
// void PathAddLines(void * path, CGAffineTransform* m, CGPoint* points, NSUInteger count);
// int32_t SetDisplayTransferByFormula(uint32_t display, float redMin, float redMax, float redGamma, float greenMin, float greenMax, float greenGamma, float blueMin, float blueMax, float blueGamma);
// int DisplayIsAsleep(uint32_t display);
// uint32_t ImageGetBitmapInfo(void * image);
// void * WindowListCreateImageFromArray(CGRect screenBounds, void * windowArray, uint32_t imageOption);
// uint BitmapContextGetWidth(void * context);
// CGAffineTransform AffineTransformMake(float a, float b, float c, float d, float tx, float ty);
// void * PathCreateWithEllipseInRect(CGRect rect, CGAffineTransform* transform);
// bool RectIntersectsRect(CGRect rect1, CGRect rect2);
// void * DisplayCreateImage(uint32_t displayID);
// void ContextSetLineWidth(void * c, float width);
// int32_t EnableEventStateCombining(NSInteger combineState);
// void * ColorSpaceCreateWithICCProfile(void * data);
// void PatternRelease(void * pattern);
// void ContextEOClip(void * c);
// void * ColorSpaceCreateExtended(void * space);
// void ContextSetAllowsFontSubpixelQuantization(void * c, BOOL allowsFontSubpixelQuantization);
// CGPoint ContextConvertPointToUserSpace(void * c, CGPoint point);
// void * DataConsumerRetain(void * consumer);
// void EventSourceSetLocalEventsFilterDuringSuppressionState(void * source, uint32_t filter, uint32_t state);
// void * DisplayCurrentMode(uint32_t display);
// int32_t WaitForScreenUpdateRects(uint32_t requestedOperations, CGScreenUpdateOperation* currentOperation, CGRect* rects, NSUInteger* rectCount, CGScreenUpdateMoveDelta* delta);
// void * ColorConversionInfoCreateFromListWithArguments(void * options, void * , uint32_t , int32_t , NSObject* );
// void ContextSetGrayStrokeColor(void * c, float gray, float alpha);
// uint32_t DisplayVendorNumber(uint32_t display);
// bool RectContainsPoint(CGRect rect, CGPoint point);
// int32_t FontGetTypeID();
// void * DisplayGetDrawingContext(uint32_t display);
// bool PDFScannerPopName(void * scanner, uint8_t* value);
// void ContextSetShouldSmoothFonts(void * c, BOOL shouldSmoothFonts);
// void * EventCreate(void * source);
// uint32_t DisplayMirrorsDisplay(uint32_t display);
// bool AffineTransformIsIdentity(CGAffineTransform t);
// void ContextBeginPage(void * c, CGRect* mediaBox);
// float RectGetWidth(CGRect rect);
// CGPoint PointMake(float x, float y);
// void * EventCreateScrollWheelEvent(void * source, uint32_t units, uint32_t wheelCount, int32_t wheel1);
// void * ColorGetConstantColor(void * colorName);
// int32_t SetLocalEventsSuppressionInterval(double seconds);
// void * EventCreateSourceFromEvent(void * event);
// int FontGetXHeight(void * font);
// int32_t DisplayStreamStart(void * displayStream);
// CGPoint EventGetUnflippedLocation(void * event);
// void * ImageCreateWithMask(void * image, void * mask);
// void PDFContextSetURLForRect(void * context, void * url, CGRect rect);
// void * ColorSpaceCreateLinearized(void * space);
// uint ColorSpaceGetColorTableCount(void * space);
// void * DisplayBestModeForParametersAndRefreshRate(uint32_t display, NSUInteger bitsPerPixel, NSUInteger width, NSUInteger height, double refreshRate, boolean_t* exactMatch);
// uint32_t OpenGLDisplayMaskToDisplayID(uint32_t mask);
// uint DisplayStreamUpdateGetDropCount(void * updateRef);
// bool PointEqualToPoint(CGPoint point1, CGPoint point2);
// void ContextFlush(void * c);
// CGVector VectorMake(float dx, float dy);
// uint DisplayModeGetHeight(void * mode);
// CGRect RectUnion(CGRect r1, CGRect r2);
// bool PDFScannerScan(void * scanner);
// void ContextAddLines(void * c, CGPoint* points, NSUInteger count);
// void ContextSetMiterLimit(void * c, float limit);
// void * PDFStreamGetDictionary(void * stream);
// uint PDFStringGetLength(void * string);
// void ContextClip(void * c);
// int32_t ReleaseAllDisplays();
// int CursorIsDrawnInFramebuffer();
// uint32_t DisplaySerialNumber(uint32_t display);
// bool PDFArrayGetArray(void * array, NSUInteger index, void * value);
// void * ColorSpaceCopyName(void * space);
// void ContextSetLineCap(void * c, int32_t cap);
// void ShadingRelease(void * shading);
// void * PDFOperatorTableRetain(void * table);
// uint32_t DisplayModelNumber(uint32_t display);
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
// int32_t ConfigureDisplayMode(void * config, uint32_t display, void * mode);
// void ContextAddLineToPoint(void * c, float x, float y);
// bool PDFScannerPopBoolean(void * scanner, CGPDFBoolean* value);
// void * FontCreateWithFontName(void * name);
// void ContextSetInterpolationQuality(void * c, int32_t quality);
// void * PDFDocumentGetCatalog(void * document);
// void LayerRelease(void * layer);
// int32_t LayerGetTypeID();
// int32_t FunctionGetTypeID();
// void * PDFContentStreamGetStreams(void * cs);
// uint PDFDictionaryGetCount(void * dict);
// int32_t DisplayShowCursor(uint32_t display);
// uint32_t BitmapContextGetBitmapInfo(void * context);
// CGAffineTransform AffineTransformScale(CGAffineTransform t, float sx, float sy);
// void * FontCreatePostScriptSubset(void * font, void * subsetName, int32_t format, CGGlyph* glyphs, NSUInteger count, CGGlyph* encoding);
// int DisplayIsMain(uint32_t display);
// void * PathCreateMutable();
// void ReleaseScreenRefreshRects(CGRect* rects);
// void ContextShowGlyphs(void * c, CGGlyph* g, NSUInteger count);
// void DisplayRestoreColorSyncSettings();
// void ContextSetLineDash(void * c, float phase, float* lengths, NSUInteger count);
// int32_t DisplayModeGetTypeID();
// uint64_t EventSourceFlagsState(int32_t stateID);
// bool RequestPostEventAccess();
// void ContextFillRect(void * c, CGRect rect);
// void * PathRetain(void * path);
// bool PDFDocumentIsUnlocked(void * document);
import "C"
import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/objc"
)

// Synthesizes a low-level keyboard event on the local machine. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541792-cgpostkeyboardevent?language=objc
func PostKeyboardEvent(keyChar CharCode, virtualKey KeyCode, keyDown kernel.Boolean_t) Error {
	rv := C.PostKeyboardEvent(C.uint16_t(keyChar), C.uint16_t(virtualKey), C.NSInteger(keyDown))
	return Error(rv)
}

// Strokes an ellipse that fits inside the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455774-cgcontextstrokeellipseinrect?language=objc
func ContextStrokeEllipseInRect(c ContextRef, rect Rect) {
	C.ContextStrokeEllipseInRect(c, rect)
}

// Returns a rectangle that is transformed from user space coordinate to device space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456017-cgcontextconvertrecttodevicespac?language=objc
func ContextConvertRectToDeviceSpace(c ContextRef, rect Rect) Rect {
	rv := C.ContextConvertRectToDeviceSpace(c, rect)
	return Rect(rv)
}

// Sets the event source of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455500-cgeventsetsource?language=objc
func EventSetSource(event EventRef, source EventSourceRef) {
	C.EventSetSource(event, source)
}

// Sets the current font size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456426-cgcontextsetfontsize?language=objc
func ContextSetFontSize(c ContextRef, size float64) {
	C.ContextSetFontSize(c, size)
}

// Sets the clipping path to the intersection of the current clipping path with the region defined by an array of rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454626-cgcontextcliptorects?language=objc
func ContextClipToRects(c ContextRef, rects *Rect, count uint) {
	C.ContextClipToRects(c, rects, count)
}

// Applies an affine transform to a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455875-cgrectapplyaffinetransform?language=objc
func RectApplyAffineTransform(rect Rect, t AffineTransform) Rect {
	rv := C.RectApplyAffineTransform(rect, t)
	return Rect(rv)
}

// Returns the 64-bit user-specified data for a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408777-cgeventsourcegetuserdata?language=objc
func EventSourceGetUserData(source EventSourceRef) int64 {
	rv := C.EventSourceGetUserData(source)
	return int64(rv)
}

// Returns the alpha channel information for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455401-cgimagegetalphainfo?language=objc
func ImageGetAlphaInfo(image ImageRef) ImageAlphaInfo {
	rv := C.ImageGetAlphaInfo(image)
	return ImageAlphaInfo(rv)
}

// Returns a Boolean value indicating whether a display is connected or online. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454476-cgdisplayisonline?language=objc
func DisplayIsOnline(display DirectDisplayID) kernel.Boolean_t {
	rv := C.DisplayIsOnline(C.uint32_t(display))
	return kernel.Boolean_t(rv)
}

// Releases a display fade reservation, and unfades the display if needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455230-cgreleasedisplayfadereservation?language=objc
func ReleaseDisplayFadeReservation(token DisplayFadeReservationToken) Error {
	rv := C.ReleaseDisplayFadeReservation(C.uint32_t(token))
	return Error(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3579522-cgcolorspaceusesextendedrange?language=objc
func ColorSpaceUsesExtendedRange(space ColorSpaceRef) bool {
	rv := C.ColorSpaceUsesExtendedRange(space)
	return bool(rv)
}

// Returns the bits per component of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455383-cgbitmapcontextgetbitspercompone?language=objc
func BitmapContextGetBitsPerComponent(context ContextRef) uint {
	rv := C.BitmapContextGetBitsPerComponent(context)
	return uint(rv)
}

// Creates a PDF graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454204-cgpdfcontextcreate?language=objc
func PDFContextCreate(consumer DataConsumerRef, mediaBox *Rect, auxiliaryInfo corefoundation.DictionaryRef) ContextRef {
	rv := C.PDFContextCreate(consumer, mediaBox, auxiliaryInfo)
	return ContextRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2923324-cgcolorspacegetname?language=objc
func ColorSpaceGetName(space ColorSpaceRef) corefoundation.StringRef {
	rv := C.ColorSpaceGetName(space)
	return corefoundation.StringRef(rv)
}

// Begins a transparency layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456011-cgcontextbegintransparencylayer?language=objc
func ContextBeginTransparencyLayer(c ContextRef, auxiliaryInfo corefoundation.DictionaryRef) {
	C.ContextBeginTransparencyLayer(c, auxiliaryInfo)
}

// Returns the color space associated with a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455744-cgcolorgetcolorspace?language=objc
func ColorGetColorSpace(color ColorRef) ColorSpaceRef {
	rv := C.ColorGetColorSpace(color)
	return ColorSpaceRef(rv)
}

// Returns the integer value of a field in a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455885-cgeventgetintegervaluefield?language=objc
func EventGetIntegerValueField(event EventRef, field EventField) int64 {
	rv := C.EventGetIntegerValueField(event, C.uint32_t(field))
	return int64(rv)
}

// Sets the location at which text is drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456069-cgcontextsettextposition?language=objc
func ContextSetTextPosition(c ContextRef, x float64, y float64) {
	C.ContextSetTextPosition(c, x, y)
}

// Returns a copy of the ICC profile data of the provided color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1644732-cgcolorspacecopyiccdata?language=objc
func ColorSpaceCopyICCData(space ColorSpaceRef) corefoundation.DataRef {
	rv := C.ColorSpaceCopyICCData(space)
	return corefoundation.DataRef(rv)
}

// Sets the platform font in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454950-cgcontextsetfont?language=objc
func ContextSetFont(c ContextRef, font FontRef) {
	C.ContextSetFont(c, font)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586508-cgcontextdrawpdfdocument?language=objc
func ContextDrawPDFDocument(c ContextRef, rect Rect, document PDFDocumentRef, page int) {
	C.ContextDrawPDFDocument(c, rect, document, page)
}

// Returns a rectangle that is transformed from device space coordinate to user space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454165-cgcontextconvertrecttouserspace?language=objc
func ContextConvertRectToUserSpace(c ContextRef, rect Rect) Rect {
	rv := C.ContextConvertRectToUserSpace(c, rect)
	return Rect(rv)
}

// Creates a color space from a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2962829-cgcolorspacecreatewithpropertyli?language=objc
func ColorSpaceCreateWithPropertyList(plist corefoundation.PropertyListRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateWithPropertyList(C.void * (plist))
	return ColorSpaceRef(rv)
}

// Increments the retain count of a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586339-cgcolorretain?language=objc
func ColorRetain(color ColorRef) ColorRef {
	rv := C.ColorRetain(color)
	return ColorRef(rv)
}

// Creates a new empty path in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456635-cgcontextbeginpath?language=objc
func ContextBeginPath(c ContextRef) {
	C.ContextBeginPath(c)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3861799-cgcolorspaceishlgbased?language=objc
func ColorSpaceIsHLGBased(s ColorSpaceRef) bool {
	rv := C.ColorSpaceIsHLGBased(s)
	return bool(rv)
}

// Creates a color in the Generic gray color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456453-cgcolorcreategenericgray?language=objc
func ColorCreateGenericGray(gray float64, alpha float64) ColorRef {
	rv := C.ColorCreateGenericGray(gray, alpha)
	return ColorRef(rv)
}

// Decrements the retain count of a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408304-cgdataproviderrelease?language=objc
func DataProviderRelease(provider DataProviderRef) {
	C.DataProviderRelease(provider)
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411147-cgpathaddarc?language=objc
func PathAddArc(path MutablePathRef, m *AffineTransform, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise bool) {
	C.PathAddArc(path, m, x, y, radius, startAngle, endAngle, clockwise)
}

// Returns the type identifier for the opaque type CGEventRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454922-cgeventgettypeid?language=objc
func EventGetTypeID() corefoundation.TypeID {
	rv := C.EventGetTypeID()
	return corefoundation.TypeID(rv)
}

// Paints a gradient fill that varies along the line defined by the provided starting and ending points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454782-cgcontextdrawlineargradient?language=objc
func ContextDrawLinearGradient(c ContextRef, gradient GradientRef, startPoint Point, endPoint Point, options GradientDrawingOptions) {
	C.ContextDrawLinearGradient(c, gradient, startPoint, endPoint, C.uint32_t(options))
}

// Returns the rectangle that represents a type of box for a content region or page dimensions of a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456114-cgpdfpagegetboxrect?language=objc
func PDFPageGetBoxRect(page PDFPageRef, box PDFBox) Rect {
	rv := C.PDFPageGetBoxRect(page, C.int32_t(box))
	return Rect(rv)
}

// Paints a rectangular path, using the specified line width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454679-cgcontextstrokerectwithwidth?language=objc
func ContextStrokeRectWithWidth(c ContextRef, rect Rect, width float64) {
	C.ContextStrokeRectWithWidth(c, rect, width)
}

// Returns information about the caller’s window server session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454780-cgsessioncopycurrentdictionary?language=objc
func SessionCopyCurrentDictionary() corefoundation.DictionaryRef {
	rv := C.SessionCopyCurrentDictionary()
	return corefoundation.DictionaryRef(rv)
}

// Returns the Core Foundation type identifier for Core Graphics shading objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454302-cgshadinggettypeid?language=objc
func ShadingGetTypeID() corefoundation.TypeID {
	rv := C.ShadingGetTypeID()
	return corefoundation.TypeID(rv)
}

// Creates a pattern color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408869-cgcolorspacecreatepattern?language=objc
func ColorSpaceCreatePattern(baseSpace ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceCreatePattern(baseSpace)
	return ColorSpaceRef(rv)
}

// Returns the bounds of a display in the global display coordinate space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456395-cgdisplaybounds?language=objc
func DisplayBounds(display DirectDisplayID) Rect {
	rv := C.DisplayBounds(C.uint32_t(display))
	return Rect(rv)
}

// Sets the current text drawing mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454253-cgcontextsettextdrawingmode?language=objc
func ContextSetTextDrawingMode(c ContextRef, mode TextDrawingMode) {
	C.ContextSetTextDrawingMode(c, C.int32_t(mode))
}

// Creates a device-dependent grayscale color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408908-cgcolorspacecreatedevicegray?language=objc
func ColorSpaceCreateDeviceGray() ColorSpaceRef {
	rv := C.ColorSpaceCreateDeviceGray()
	return ColorSpaceRef(rv)
}

// Creates a calibrated grayscale color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408887-cgcolorspacecreatecalibratedgray?language=objc
func ColorSpaceCreateCalibratedGray(whitePoint *float64, blackPoint *float64, gamma float64) ColorSpaceRef {
	rv := C.ColorSpaceCreateCalibratedGray(whitePoint, blackPoint, gamma)
	return ColorSpaceRef(rv)
}

// Creates a stroked copy of another path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411128-cgpathcreatecopybystrokingpath?language=objc
func PathCreateCopyByStrokingPath(path unsafe.Pointer, transform *AffineTransform, lineWidth float64, lineCap LineCap, lineJoin LineJoin, miterLimit float64) unsafe.Pointer {
	rv := C.PathCreateCopyByStrokingPath(path, transform, lineWidth, C.int32_t(lineCap), C.int32_t(lineJoin), miterLimit)
	return unsafe.Pointer(rv)
}

// Indicates whether two colors are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455217-cgcolorequaltocolor?language=objc
func ColorEqualToColor(color1 ColorRef, color2 ColorRef) bool {
	rv := C.ColorEqualToColor(color1, color2)
	return bool(rv)
}

// Checks whether a point is contained in a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411175-cgpathcontainspoint?language=objc
func PathContainsPoint(path unsafe.Pointer, m *AffineTransform, point Point, eoFill bool) bool {
	rv := C.PathContainsPoint(path, m, point, eoFill)
	return bool(rv)
}

// Returns the variation specification dictionary for a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396355-cgfontcopyvariations?language=objc
func FontCopyVariations(font FontRef) corefoundation.DictionaryRef {
	rv := C.FontCopyVariations(font)
	return corefoundation.DictionaryRef(rv)
}

// Increments the retain count of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396384-cgfontretain?language=objc
func FontRetain(font FontRef) FontRef {
	rv := C.FontRetain(font)
	return FontRef(rv)
}

// Returns the color space for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454190-cgdisplaycopycolorspace?language=objc
func DisplayCopyColorSpace(display DirectDisplayID) ColorSpaceRef {
	rv := C.DisplayCopyColorSpace(C.uint32_t(display))
	return ColorSpaceRef(rv)
}

// Returns the width of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456148-cgimagegetwidth?language=objc
func ImageGetWidth(image ImageRef) uint {
	rv := C.ImageGetWidth(image)
	return uint(rv)
}

// Creates a Core Graphics PDF document using data specified by a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402585-cgpdfdocumentcreatewithurl?language=objc
func PDFDocumentCreateWithURL(url corefoundation.URLRef) PDFDocumentRef {
	rv := C.PDFDocumentCreateWithURL(url)
	return PDFDocumentRef(rv)
}

// Sets the event flags of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455044-cgeventsetflags?language=objc
func EventSetFlags(event EventRef, flags EventFlags) {
	C.EventSetFlags(event, C.uint64_t(flags))
}

// Returns a rectangle with the specified coordinate and size values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455245-cgrectmake?language=objc
func RectMake(x float64, y float64, width float64, height float64) Rect {
	rv := C.RectMake(x, y, width, height)
	return Rect(rv)
}

// Returns the number of color components (including alpha) associated with a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454130-cgcolorgetnumberofcomponents?language=objc
func ColorGetNumberOfComponents(color ColorRef) uint {
	rv := C.ColorGetNumberOfComponents(color)
	return uint(rv)
}

// Create an immutable path of a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411155-cgpathcreatewithrect?language=objc
func PathCreateWithRect(rect Rect, transform *AffineTransform) unsafe.Pointer {
	rv := C.PathCreateWithRect(rect, transform)
	return unsafe.Pointer(rv)
}

// Returns the current point in a non-empty path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454788-cgcontextgetpathcurrentpoint?language=objc
func ContextGetPathCurrentPoint(c ContextRef) Point {
	rv := C.ContextGetPathCurrentPoint(c)
	return Point(rv)
}

// Closes and completes a subpath in a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411188-cgpathclosesubpath?language=objc
func PathCloseSubpath(path MutablePathRef) {
	C.PathCloseSubpath(path)
}

// Appends a path to onto a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411201-cgpathaddpath?language=objc
func PathAddPath(path1 MutablePathRef, m *AffineTransform, path2 unsafe.Pointer) {
	C.PathAddPath(path1, m, path2)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656524-cgrequestscreencaptureaccess?language=objc
func RequestScreenCaptureAccess() bool {
	rv := C.RequestScreenCaptureAccess()
	return bool(rv)
}

// Returns the type identifier of a Quartz display stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454784-cgdisplaystreamgettypeid?language=objc
func DisplayStreamGetTypeID() corefoundation.TypeID {
	rv := C.DisplayStreamGetTypeID()
	return corefoundation.TypeID(rv)
}

// Fills in a size using the contents of the specified dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454318-cgsizemakewithdictionaryrepresen?language=objc
func SizeMakeWithDictionaryRepresentation(dict corefoundation.DictionaryRef, size *Size) bool {
	rv := C.SizeMakeWithDictionaryRepresentation(dict, size)
	return bool(rv)
}

// Returns whether the RGB color space covers a significant portion of the NTSC color gamut. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1644737-cgcolorspaceiswidegamutrgb?language=objc
func ColorSpaceIsWideGamutRGB(arg0 ColorSpaceRef) bool {
	rv := C.ColorSpaceIsWideGamutRGB(arg0)
	return bool(rv)
}

// Decrements the retain count of a CGPDFOperatorTable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455277-cgpdfoperatortablerelease?language=objc
func PDFOperatorTableRelease(table unsafe.Pointer) {
	C.PDFOperatorTableRelease(table)
}

// Returns the number of glyph space units per em for the provided font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396344-cgfontgetunitsperem?language=objc
func FontGetUnitsPerEm(font FontRef) int {
	rv := C.FontGetUnitsPerEm(font)
	return int(rv)
}

// Increments the retain count of a PDF content stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410049-cgpdfcontentstreamretain?language=objc
func PDFContentStreamRetain(cs unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFContentStreamRetain(cs)
	return unsafe.Pointer(rv)
}

// Sets the opacity level for objects drawn in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456404-cgcontextsetalpha?language=objc
func ContextSetAlpha(c ContextRef, alpha float64) {
	C.ContextSetAlpha(c, alpha)
}

// Appends a cubic Bézier curve from the current point, using the provided control points and end point . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456393-cgcontextaddcurvetopoint?language=objc
func ContextAddCurveToPoint(c ContextRef, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {
	C.ContextAddCurveToPoint(c, cp1x, cp1y, cp2x, cp2y, x, y)
}

// Creates a color in the Generic CMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454222-cgcolorcreategenericcmyk?language=objc
func ColorCreateGenericCMYK(cyan float64, magenta float64, yellow float64, black float64, alpha float64) ColorRef {
	rv := C.ColorCreateGenericCMYK(cyan, magenta, yellow, black, alpha)
	return ColorRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3684558-cgcolorspacecreateextendedlinear?language=objc
func ColorSpaceCreateExtendedLinearized(space ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateExtendedLinearized(space)
	return ColorSpaceRef(rv)
}

// Returns the I/O Kit service port of the specified display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1543516-cgdisplayioserviceport?language=objc
func DisplayIOServicePort(display DirectDisplayID) kernel.Io_service_t {
	rv := C.DisplayIOServicePort(C.uint32_t(display))
	return kernel.Io_service_t(rv)
}

// Creates an ICC-based color space using the ICC profile contained in the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2866135-cgcolorspacecreatewithiccdata?language=objc
func ColorSpaceCreateWithICCData(data corefoundation.TypeRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateWithICCData(C.void * (data))
	return ColorSpaceRef(rv)
}

// Gets a list of online displays with bounds that intersect the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456071-cggetdisplayswithrect?language=objc
func GetDisplaysWithRect(rect Rect, maxDisplays uint32, displays *DirectDisplayID, matchingDisplayCount *uint32) Error {
	rv := C.GetDisplaysWithRect(rect, maxDisplays, displays, matchingDisplayCount)
	return Error(rv)
}

// Synthesizes a low-level mouse-button event on the local machine. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541785-cgpostmouseevent?language=objc
func PostMouseEvent(mouseCursorPosition Point, updateMouseCursorPosition kernel.Boolean_t, buttonCount ButtonCount, mouseButtonDown kernel.Boolean_t) Error {
	rv := C.PostMouseEvent(mouseCursorPosition, C.NSInteger(updateMouseCursorPosition), C.uint32_t(buttonCount), C.NSInteger(mouseButtonDown))
	return Error(rv)
}

// Returns the width of the specified display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454442-cgdisplaymodegetwidth?language=objc
func DisplayModeGetWidth(mode DisplayModeRef) uint {
	rv := C.DisplayModeGetWidth(mode)
	return uint(rv)
}

// Pushes a copy of the current graphics state onto the graphics state stack for the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456156-cgcontextsavegstate?language=objc
func ContextSaveGState(c ContextRef) {
	C.ContextSaveGState(c)
}

// Returns the cap height of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396338-cgfontgetcapheight?language=objc
func FontGetCapHeight(font FontRef) int {
	rv := C.FontGetCapHeight(font)
	return int(rv)
}

// Returns the number of items in a PDF array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455207-cgpdfarraygetcount?language=objc
func PDFArrayGetCount(array unsafe.Pointer) uint {
	rv := C.PDFArrayGetCount(array)
	return uint(rv)
}

// Decrements the retain count of a function object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1390864-cgfunctionrelease?language=objc
func FunctionRelease(function FunctionRef) {
	C.FunctionRelease(function)
}

// Unlocks an encrypted PDF document when a valid password is supplied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402599-cgpdfdocumentunlockwithpassword?language=objc
func PDFDocumentUnlockWithPassword(document PDFDocumentRef, password *uint8) bool {
	rv := C.PDFDocumentUnlockWithPassword(document, password)
	return bool(rv)
}

// Returns the rotation angle of a display in degrees. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455299-cgdisplayrotation?language=objc
func DisplayRotation(display DirectDisplayID) float64 {
	rv := C.DisplayRotation(C.uint32_t(display))
	return float64(rv)
}

// Returns an affine transformation matrix constructed from scaling values you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455016-cgaffinetransformmakescale?language=objc
func AffineTransformMakeScale(sx float64, sy float64) AffineTransform {
	rv := C.AffineTransformMakeScale(sx, sy)
	return AffineTransform(rv)
}

// Draws the contents of a layer object into the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450896-cgcontextdrawlayerinrect?language=objc
func ContextDrawLayerInRect(context ContextRef, rect Rect, layer LayerRef) {
	C.ContextDrawLayerInRect(context, rect, layer)
}

// Returns whether a rectangle is infinite. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455008-cgrectisinfinite?language=objc
func RectIsInfinite(rect Rect) bool {
	rv := C.RectIsInfinite(rect)
	return bool(rv)
}

// Draws an array of glyphs with varying offsets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586503-cgcontextshowglyphswithadvances?language=objc
func ContextShowGlyphsWithAdvances(c ContextRef, glyphs *Glyph, advances *Size, count uint) {
	C.ContextShowGlyphsWithAdvances(c, glyphs, advances, count)
}

// Creates a copy of a bitmap image, replacing its colorspace. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455355-cgimagecreatecopywithcolorspace?language=objc
func ImageCreateCopyWithColorSpace(image ImageRef, space ColorSpaceRef) ImageRef {
	rv := C.ImageCreateCopyWithColorSpace(image, space)
	return ImageRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656520-cgpreflightposteventaccess?language=objc
func PreflightPostEventAccess() bool {
	rv := C.PreflightPostEventAccess()
	return bool(rv)
}

// Decrements the retain count of a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586509-cgcontextrelease?language=objc
func ContextRelease(c ContextRef) {
	C.ContextRelease(c)
}

// Returns the Core Foundation type identifier for Core Graphics data consumers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455226-cgdataconsumergettypeid?language=objc
func DataConsumerGetTypeID() corefoundation.TypeID {
	rv := C.DataConsumerGetTypeID()
	return corefoundation.TypeID(rv)
}

// Completes a set of display configuration changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454488-cgcompletedisplayconfiguration?language=objc
func CompleteDisplayConfiguration(config unsafe.Pointer, option ConfigureOption) Error {
	rv := C.CompleteDisplayConfiguration(config, C.uint32_t(option))
	return Error(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3042357-cgpdfcontextendtag?language=objc
func PDFContextEndTag(context ContextRef) {
	C.PDFContextEndTag(context)
}

// Returns a copy of the ICC profile of the provided color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408889-cgcolorspacecopyiccprofile?language=objc
func ColorSpaceCopyICCProfile(space ColorSpaceRef) corefoundation.DataRef {
	rv := C.ColorSpaceCopyICCProfile(space)
	return corefoundation.DataRef(rv)
}

// Creates a color in the Generic gray color space with a gamma ramp of 2.2. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3042354-cgcolorcreategenericgraygamma2_2?language=objc
func ColorCreateGenericGrayGamma2_2(gray float64, alpha float64) ColorRef {
	rv := C.ColorCreateGenericGrayGamma2_2(gray, alpha)
	return ColorRef(rv)
}

// Returns the bounding box containing all points in a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411165-cgpathgetboundingbox?language=objc
func PathGetBoundingBox(path unsafe.Pointer) Rect {
	rv := C.PathGetBoundingBox(path)
	return Rect(rv)
}

// Returns the major and minor version numbers of a Core Graphics PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402604-cgpdfdocumentgetversion?language=objc
func PDFDocumentGetVersion(document PDFDocumentRef, majorVersion *int, minorVersion *int) {
	C.PDFDocumentGetVersion(document, majorVersion, minorVersion)
}

// Appends a rounded rectangle to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411124-cgpathaddroundedrect?language=objc
func PathAddRoundedRect(path MutablePathRef, transform *AffineTransform, rect Rect, cornerWidth float64, cornerHeight float64) {
	C.PathAddRoundedRect(path, transform, rect, cornerWidth, cornerHeight)
}

// Returns information about the currently available display modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1562068-cgdisplayavailablemodes?language=objc
func DisplayAvailableModes(dsp DirectDisplayID) corefoundation.ArrayRef {
	rv := C.DisplayAvailableModes(C.uint32_t(dsp))
	return corefoundation.ArrayRef(rv)
}

// Returns the window level that corresponds to one of the standard window types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454084-cgwindowlevelforkey?language=objc
func WindowLevelForKey(key WindowLevelKey) WindowLevel {
	rv := C.WindowLevelForKey(C.int32_t(key))
	return WindowLevel(rv)
}

// Sets the rendering intent in the current graphics state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455544-cgcontextsetrenderingintent?language=objc
func ContextSetRenderingIntent(c ContextRef, intent ColorRenderingIntent) {
	C.ContextSetRenderingIntent(c, C.int32_t(intent))
}

// Enables shadowing with color a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455205-cgcontextsetshadowwithcolor?language=objc
func ContextSetShadowWithColor(c ContextRef, offset Size, blur float64, color ColorRef) {
	C.ContextSetShadowWithColor(c, offset, blur, color)
}

// Returns the type identifier for Core Graphics patterns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456445-cgpatterngettypeid?language=objc
func PatternGetTypeID() corefoundation.TypeID {
	rv := C.PatternGetTypeID()
	return corefoundation.TypeID(rv)
}

// Retains a Core Graphics display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1562059-cgdisplaymoderetain?language=objc
func DisplayModeRetain(mode DisplayModeRef) DisplayModeRef {
	rv := C.DisplayModeRetain(mode)
	return DisplayModeRef(rv)
}

// Returns the timestamp of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455481-cgeventgettimestamp?language=objc
func EventGetTimestamp(event EventRef) EventTimestamp {
	rv := C.EventGetTimestamp(event)
	return EventTimestamp(rv)
}

// Creates a mutable copy of an existing graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411196-cgpathcreatemutablecopy?language=objc
func PathCreateMutableCopy(path unsafe.Pointer) MutablePathRef {
	rv := C.PathCreateMutableCopy(path)
	return MutablePathRef(rv)
}

// Returns the rotation angle of a page in a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402602-cgpdfdocumentgetrotationangle?language=objc
func PDFDocumentGetRotationAngle(document PDFDocumentRef, page int) int {
	rv := C.PDFDocumentGetRotationAngle(document, page)
	return int(rv)
}

// Returns the current transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454691-cgcontextgetctm?language=objc
func ContextGetCTM(c ContextRef) AffineTransform {
	rv := C.ContextGetCTM(c)
	return AffineTransform(rv)
}

// Returns a flattened data representation of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454381-cgeventcreatedata?language=objc
func EventCreateData(allocator corefoundation.AllocatorRef, event EventRef) corefoundation.DataRef {
	rv := C.EventCreateData(allocator, event)
	return corefoundation.DataRef(rv)
}

// Returns a Boolean value indicating whether a display is running in a stereo graphics mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455025-cgdisplayisstereo?language=objc
func DisplayIsStereo(display DirectDisplayID) kernel.Boolean_t {
	rv := C.DisplayIsStereo(C.uint32_t(display))
	return kernel.Boolean_t(rv)
}

// Transforms the user coordinate system in a context using a specified matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454897-cgcontextconcatctm?language=objc
func ContextConcatCTM(c ContextRef, transform AffineTransform) {
	C.ContextConcatCTM(c, transform)
}

// Returns a size that is transformed from device space coordinates to user space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456510-cgcontextconvertsizetouserspace?language=objc
func ContextConvertSizeToUserSpace(c ContextRef, size Size) Size {
	rv := C.ContextConvertSizeToUserSpace(c, size)
	return Size(rv)
}

// Retrieves an integer object from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454399-cgpdfscannerpopinteger?language=objc
func PDFScannerPopInteger(scanner unsafe.Pointer, value *PDFInteger) bool {
	rv := C.PDFScannerPopInteger(scanner, value)
	return bool(rv)
}

// Associates custom metadata with the PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456026-cgpdfcontextadddocumentmetadata?language=objc
func PDFContextAddDocumentMetadata(context ContextRef, metadata corefoundation.DataRef) {
	C.PDFContextAddDocumentMetadata(context, metadata)
}

// Returns the width and height of a display in millimeters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456599-cgdisplayscreensize?language=objc
func DisplayScreenSize(display DirectDisplayID) Size {
	rv := C.DisplayScreenSize(C.uint32_t(display))
	return Size(rv)
}

// Copies the entries in the color table of an indexed color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408853-cgcolorspacegetcolortable?language=objc
func ColorSpaceGetColorTable(space ColorSpaceRef, table *uint8) {
	C.ColorSpaceGetColorTable(space, table)
}

// Increments the retain count of a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1571723-cgpdfpageretain?language=objc
func PDFPageRetain(page PDFPageRef) PDFPageRef {
	rv := C.PDFPageRetain(page)
	return PDFPageRef(rv)
}

// Paints the area of the ellipse that fits inside the provided rectangle, using the fill color in the current graphics state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454371-cgcontextfillellipseinrect?language=objc
func ContextFillEllipseInRect(c ContextRef, rect Rect) {
	C.ContextFillEllipseInRect(c, rect)
}

// Returns information about a display’s current configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454099-cgdisplaycopydisplaymode?language=objc
func DisplayCopyDisplayMode(display DirectDisplayID) DisplayModeRef {
	rv := C.DisplayCopyDisplayMode(C.uint32_t(display))
	return DisplayModeRef(rv)
}

// Sets a destination to jump to when a rectangle in the current PDF page is clicked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456459-cgpdfcontextsetdestinationforrec?language=objc
func PDFContextSetDestinationForRect(context ContextRef, name corefoundation.StringRef, rect Rect) {
	C.PDFContextSetDestinationForRect(context, name, rect)
}

// Provides a list of displays that corresponds to the bits set in an OpenGL display mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454234-cggetdisplayswithopengldisplayma?language=objc
func GetDisplaysWithOpenGLDisplayMask(mask OpenGLDisplayMask, maxDisplays uint32, displays *DirectDisplayID, matchingDisplayCount *uint32) Error {
	rv := C.GetDisplaysWithOpenGLDisplayMask(C.uint32_t(mask), maxDisplays, displays, matchingDisplayCount)
	return Error(rv)
}

// Returns whether there is a PDF array associated with a specified key in a PDF dictionary and, if so, retrieves that array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430229-cgpdfdictionarygetarray?language=objc
func PDFDictionaryGetArray(dict unsafe.Pointer, key *uint8, value unsafe.Pointer) bool {
	rv := C.PDFDictionaryGetArray(dict, key, value)
	return bool(rv)
}

// Creates a shading object to use for axial shading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455224-cgshadingcreateaxial?language=objc
func ShadingCreateAxial(space ColorSpaceRef, start Point, end Point, function FunctionRef, extendStart bool, extendEnd bool) ShadingRef {
	rv := C.ShadingCreateAxial(space, start, end, function, extendStart, extendEnd)
	return ShadingRef(rv)
}

// Enables shadowing in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456082-cgcontextsetshadow?language=objc
func ContextSetShadow(c ContextRef, offset Size, blur float64) {
	C.ContextSetShadow(c, offset, blur)
}

// Fills in a point using the contents of the specified dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455338-cgpointmakewithdictionaryreprese?language=objc
func PointMakeWithDictionaryRepresentation(dict corefoundation.DictionaryRef, point *Point) bool {
	rv := C.PointMakeWithDictionaryRepresentation(dict, point)
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454804-cgeventposttopid?language=objc
func EventPostToPid(pid kernel.Pid_t, event EventRef) {
	C.EventPostToPid(C.NSObject*(pid), event)
}

// Returns the location at which text is drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454687-cgcontextgettextposition?language=objc
func ContextGetTextPosition(c ContextRef) Point {
	rv := C.ContextGetTextPosition(c)
	return Point(rv)
}

// Returns the bits per pixel of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455946-cgbitmapcontextgetbitsperpixel?language=objc
func BitmapContextGetBitsPerPixel(context ContextRef) uint {
	rv := C.BitmapContextGetBitsPerPixel(context)
	return uint(rv)
}

// Adds an arc of a circle to the current path, possibly preceded by a straight line segment [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455756-cgcontextaddarc?language=objc
func ContextAddArc(c ContextRef, x float64, y float64, radius float64, startAngle float64, endAngle float64, clockwise int) {
	C.ContextAddArc(c, x, y, radius, startAngle, endAngle, clockwise)
}

// Sets the event type of a Quartz event (left mouse down, for example). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454300-cgeventsettype?language=objc
func EventSetType(event EventRef, type_ EventType) {
	C.EventSetType(event, C.uint32_t(type_))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2919739-cgeventcreatescrollwheelevent2?language=objc
func EventCreateScrollWheelEvent2(source EventSourceRef, units ScrollEventUnit, wheelCount uint32, wheel1 int32, wheel2 int32, wheel3 int32) EventRef {
	rv := C.EventCreateScrollWheelEvent2(source, C.uint32_t(units), wheelCount, wheel1, wheel2, wheel3)
	return EventRef(rv)
}

// Returns the color space model of the provided color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408854-cgcolorspacegetmodel?language=objc
func ColorSpaceGetModel(space ColorSpaceRef) ColorSpaceModel {
	rv := C.ColorSpaceGetModel(space)
	return ColorSpaceModel(rv)
}

// Creates a bitmap image using PNG-encoded data supplied by a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454993-cgimagecreatewithpngdataprovider?language=objc
func ImageCreateWithPNGDataProvider(source DataProviderRef, decode *float64, shouldInterpolate bool, intent ColorRenderingIntent) ImageRef {
	rv := C.ImageCreateWithPNGDataProvider(source, decode, shouldInterpolate, C.int32_t(intent))
	return ImageRef(rv)
}

// Adds to a path an ellipse that fits inside a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411222-cgpathaddellipseinrect?language=objc
func PathAddEllipseInRect(path MutablePathRef, m *AffineTransform, rect Rect) {
	C.PathAddEllipseInRect(path, m, rect)
}

// Returns a Quartz event created from a flattened data representation of the event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454249-cgeventcreatefromdata?language=objc
func EventCreateFromData(allocator corefoundation.AllocatorRef, data corefoundation.DataRef) EventRef {
	rv := C.EventCreateFromData(allocator, data)
	return EventRef(rv)
}

// Returns the art box of a page in a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402601-cgpdfdocumentgetartbox?language=objc
func PDFDocumentGetArtBox(document PDFDocumentRef, page int) Rect {
	rv := C.PDFDocumentGetArtBox(document, page)
	return Rect(rv)
}

// Creates a device-dependent CMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408897-cgcolorspacecreatedevicecmyk?language=objc
func ColorSpaceCreateDeviceCMYK() ColorSpaceRef {
	rv := C.ColorSpaceCreateDeviceCMYK()
	return ColorSpaceRef(rv)
}

// Modifies the settings of the built-in fade effect that occurs during a display configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454103-cgconfiguredisplayfadeeffect?language=objc
func ConfigureDisplayFadeEffect(config unsafe.Pointer, fadeOutSeconds DisplayFadeInterval, fadeInSeconds DisplayFadeInterval, fadeRed float64, fadeGreen float64, fadeBlue float64) Error {
	rv := C.ConfigureDisplayFadeEffect(config, C.float(fadeOutSeconds), C.float(fadeInSeconds), fadeRed, fadeGreen, fadeBlue)
	return Error(rv)
}

// Performs a single fade operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456189-cgdisplayfade?language=objc
func DisplayFade(token DisplayFadeReservationToken, duration DisplayFadeInterval, startBlend DisplayBlendFraction, endBlend DisplayBlendFraction, redBlend float64, greenBlend float64, blueBlend float64, synchronous kernel.Boolean_t) Error {
	rv := C.DisplayFade(C.uint32_t(token), C.float(duration), C.float(startBlend), C.float(endBlend), redBlend, greenBlend, blueBlend, C.NSInteger(synchronous))
	return Error(rv)
}

// Returns whether two sizes are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455176-cgsizeequaltosize?language=objc
func SizeEqualToSize(size1 Size, size2 Size) bool {
	rv := C.SizeEqualToSize(size1, size2)
	return bool(rv)
}

// Enables or disables subpixel quantization in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455766-cgcontextsetshouldsubpixelquanti?language=objc
func ContextSetShouldSubpixelQuantizeFonts(c ContextRef, shouldSubpixelQuantizeFonts bool) {
	C.ContextSetShouldSubpixelQuantizeFonts(c, shouldSubpixelQuantizeFonts)
}

// Returns the location of a Quartz mouse event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455788-cgeventgetlocation?language=objc
func EventGetLocation(event EventRef) Point {
	rv := C.EventGetLocation(event)
	return Point(rv)
}

// Tells a PostScript converter to abort a conversion at the next available opportunity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455046-cgpsconverterabort?language=objc
func PSConverterAbort(converter PSConverterRef) bool {
	rv := C.PSConverterAbort(converter)
	return bool(rv)
}

// Returns the type identifier of a Quartz display stream update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454989-cgdisplaystreamupdategettypeid?language=objc
func DisplayStreamUpdateGetTypeID() corefoundation.TypeID {
	rv := C.DisplayStreamUpdateGetTypeID()
	return corefoundation.TypeID(rv)
}

// Changes the scale of the user coordinate system in a context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454659-cgcontextscalectm?language=objc
func ContextScaleCTM(c ContextRef, sx float64, sy float64) {
	C.ContextScaleCTM(c, sx, sy)
}

// Returns a dictionary representation of the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455274-cgsizecreatedictionaryrepresenta?language=objc
func SizeCreateDictionaryRepresentation(size Size) corefoundation.DictionaryRef {
	rv := C.SizeCreateDictionaryRepresentation(size)
	return corefoundation.DictionaryRef(rv)
}

// Returns a Boolean value indicating the current button state of a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408781-cgeventsourcebuttonstate?language=objc
func EventSourceButtonState(stateID EventSourceStateID, button MouseButton) bool {
	rv := C.EventSourceButtonState(C.int32_t(stateID), C.uint32_t(button))
	return bool(rv)
}

// Enables or disables an event tap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455445-cgeventtapenable?language=objc
func EventTapEnable(tap corefoundation.MachPortRef, enable bool) {
	C.EventTapEnable(tap, enable)
}

// Returns an affine transformation matrix constructed from a rotation value you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455666-cgaffinetransformmakerotation?language=objc
func AffineTransformMakeRotation(angle float64) AffineTransform {
	rv := C.AffineTransformMakeRotation(angle)
	return AffineTransform(rv)
}

// Creates a device-independent color space that is defined according to the ICC color profile specification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408881-cgcolorspacecreateiccbased?language=objc
func ColorSpaceCreateICCBased(nComponents uint, range_ *float64, profile DataProviderRef, alternate ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateICCBased(nComponents, range_, profile, alternate)
	return ColorSpaceRef(rv)
}

// Connects or disconnects the mouse and cursor while an application is in the foreground. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454486-cgassociatemouseandmousecursorpo?language=objc
func AssociateMouseAndMouseCursorPosition(connected kernel.Boolean_t) Error {
	rv := C.AssociateMouseAndMouseCursorPosition(C.NSInteger(connected))
	return Error(rv)
}

// Increments the retain count of a Core Graphics PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402587-cgpdfdocumentretain?language=objc
func PDFDocumentRetain(document PDFDocumentRef) PDFDocumentRef {
	rv := C.PDFDocumentRetain(document)
	return PDFDocumentRef(rv)
}

// Returns the current point in a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411132-cgpathgetcurrentpoint?language=objc
func PathGetCurrentPoint(path unsafe.Pointer) Point {
	rv := C.PathGetCurrentPoint(path)
	return Point(rv)
}

// Sets the current stroke color to a value in the DeviceRGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456378-cgcontextsetrgbstrokecolor?language=objc
func ContextSetRGBStrokeColor(c ContextRef, red float64, green float64, blue float64, alpha float64) {
	C.ContextSetRGBStrokeColor(c, red, green, blue, alpha)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656523-cgpreflightscreencaptureaccess?language=objc
func PreflightScreenCaptureAccess() bool {
	rv := C.PreflightScreenCaptureAccess()
	return bool(rv)
}

// Sets the color gamma function for a display by specifying the values in the RGB gamma tables. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456604-cgsetdisplaytransferbytable?language=objc
func SetDisplayTransferByTable(display DirectDisplayID, tableSize uint32, redTable *GammaValue, greenTable *GammaValue, blueTable *GammaValue) Error {
	rv := C.SetDisplayTransferByTable(C.uint32_t(display), tableSize, redTable, greenTable, blueTable)
	return Error(rv)
}

// Sets the accuracy of curved paths in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455798-cgcontextsetflatness?language=objc
func ContextSetFlatness(c ContextRef, flatness float64) {
	C.ContextSetFlatness(c, flatness)
}

// Returns a Boolean value indicating whether a display is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455222-cgdisplayisactive?language=objc
func DisplayIsActive(display DirectDisplayID) kernel.Boolean_t {
	rv := C.DisplayIsActive(C.uint32_t(display))
	return kernel.Boolean_t(rv)
}

// Sets the pattern phase of a context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455334-cgcontextsetpatternphase?language=objc
func ContextSetPatternPhase(c ContextRef, phase Size) {
	C.ContextSetPatternPhase(c, phase)
}

// Decrements the retain count of a scanner object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454962-cgpdfscannerrelease?language=objc
func PDFScannerRelease(scanner unsafe.Pointer) {
	C.PDFScannerRelease(scanner)
}

// Sets the style for the joins of connected lines in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455973-cgcontextsetlinejoin?language=objc
func ContextSetLineJoin(c ContextRef, join LineJoin) {
	C.ContextSetLineJoin(c, C.int32_t(join))
}

// Returns an array of tags that correspond to the font tables for a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396392-cgfontcopytabletags?language=objc
func FontCopyTableTags(font FontRef) corefoundation.ArrayRef {
	rv := C.FontCopyTableTags(font)
	return corefoundation.ArrayRef(rv)
}

// Provides a list of displays that are active for drawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454603-cggetactivedisplaylist?language=objc
func GetActiveDisplayList(maxDisplays uint32, activeDisplays *DirectDisplayID, displayCount *uint32) Error {
	rv := C.GetActiveDisplayList(maxDisplays, activeDisplays, displayCount)
	return Error(rv)
}

// Indicates whether or not a graphics path represents a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411163-cgpathisrect?language=objc
func PathIsRect(path unsafe.Pointer, rect *Rect) bool {
	rv := C.PathIsRect(path, rect)
	return bool(rv)
}

// Creates a direct-access data provider that uses a URL to supply data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408327-cgdataprovidercreatewithurl?language=objc
func DataProviderCreateWithURL(url corefoundation.URLRef) DataProviderRef {
	rv := C.DataProviderCreateWithURL(url)
	return DataProviderRef(rv)
}

// Returns a Quartz event source created with a specified source state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408776-cgeventsourcecreate?language=objc
func EventSourceCreate(stateID EventSourceStateID) EventSourceRef {
	rv := C.EventSourceCreate(C.int32_t(stateID))
	return EventSourceRef(rv)
}

// Returns the current level of interpolation quality for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454940-cgcontextgetinterpolationquality?language=objc
func ContextGetInterpolationQuality(c ContextRef) InterpolationQuality {
	rv := C.ContextGetInterpolationQuality(c)
	return InterpolationQuality(rv)
}

// Creates a conversion between two specified color spaces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2113677-cgcolorconversioninfocreate?language=objc
func ColorConversionInfoCreate(src ColorSpaceRef, dst ColorSpaceRef) ColorConversionInfoRef {
	rv := C.ColorConversionInfoCreate(src, dst)
	return ColorConversionInfoRef(rv)
}

// Returns the trim box of a page in a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402590-cgpdfdocumentgettrimbox?language=objc
func PDFDocumentGetTrimBox(document PDFDocumentRef, page int) Rect {
	rv := C.PDFDocumentGetTrimBox(document, page)
	return Rect(rv)
}

// Returns whether the first rectangle contains the second rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454186-cgrectcontainsrect?language=objc
func RectContainsRect(rect1 Rect, rect2 Rect) bool {
	rv := C.RectContainsRect(rect1, rect2)
	return bool(rv)
}

// Creates a PostScript encoding of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396348-cgfontcreatepostscriptencoding?language=objc
func FontCreatePostScriptEncoding(font FontRef, encoding *Glyph) corefoundation.DataRef {
	rv := C.FontCreatePostScriptEncoding(font, encoding)
	return corefoundation.DataRef(rv)
}

// Draws the contents of a CGLayer object at the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450894-cgcontextdrawlayeratpoint?language=objc
func ContextDrawLayerAtPoint(context ContextRef, point Point, layer LayerRef) {
	C.ContextDrawLayerAtPoint(context, point, layer)
}

// Returns whether an object at a given index in a PDF array is a PDF stream and, if so, retrieves that stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454424-cgpdfarraygetstream?language=objc
func PDFArrayGetStream(array unsafe.Pointer, index uint, value unsafe.Pointer) bool {
	rv := C.PDFArrayGetStream(array, index, value)
	return bool(rv)
}

// Displays a character string at a position you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586505-cgcontextshowtextatpoint?language=objc
func ContextShowTextAtPoint(c ContextRef, x float64, y float64, string *uint8, length uint) {
	C.ContextShowTextAtPoint(c, x, y, string_, length)
}

// Sets the scale of pixels per line in a scrolling event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408766-cgeventsourcesetpixelsperline?language=objc
func EventSourceSetPixelsPerLine(source EventSourceRef, pixelsPerLine float64) {
	C.EventSourceSetPixelsPerLine(source, pixelsPerLine)
}

// Returns a new Quartz keyboard event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456564-cgeventcreatekeyboardevent?language=objc
func EventCreateKeyboardEvent(source EventSourceRef, virtualKey KeyCode, keyDown bool) EventRef {
	rv := C.EventCreateKeyboardEvent(source, C.uint16_t(virtualKey), keyDown)
	return EventRef(rv)
}

// Returns a count of events of a given type seen since the window server started. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408794-cgeventsourcecounterforeventtype?language=objc
func EventSourceCounterForEventType(stateID EventSourceStateID, eventType EventType) uint32 {
	rv := C.EventSourceCounterForEventType(C.int32_t(stateID), C.uint32_t(eventType))
	return uint32(rv)
}

// Appends an array of rectangles to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411153-cgpathaddrects?language=objc
func PathAddRects(path MutablePathRef, m *AffineTransform, rects *Rect, count uint) {
	C.PathAddRects(path, m, rects, count)
}

// Sets the timestamp of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456611-cgeventsettimestamp?language=objc
func EventSetTimestamp(event EventRef, timestamp EventTimestamp) {
	C.EventSetTimestamp(event, C.uint64_t(timestamp))
}

// Returns a Boolean value indicating the current keyboard state of a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408768-cgeventsourcekeystate?language=objc
func EventSourceKeyState(stateID EventSourceStateID, key KeyCode) bool {
	rv := C.EventSourceKeyState(C.int32_t(stateID), C.uint16_t(key))
	return bool(rv)
}

// Returns the height of a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455645-cgrectgetheight?language=objc
func RectGetHeight(rect Rect) float64 {
	rv := C.RectGetHeight(rect)
	return float64(rv)
}

// Displays an array of glyphs at a position you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586502-cgcontextshowglyphsatpoint?language=objc
func ContextShowGlyphsAtPoint(c ContextRef, x float64, y float64, glyphs *Glyph, count uint) {
	C.ContextShowGlyphsAtPoint(c, x, y, glyphs, count)
}

// Returns whether there is a PDF integer associated with a specified key in a PDF dictionary and, if so, retrieves that integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430231-cgpdfdictionarygetinteger?language=objc
func PDFDictionaryGetInteger(dict unsafe.Pointer, key *uint8, value *PDFInteger) bool {
	rv := C.PDFDictionaryGetInteger(dict, key, value)
	return bool(rv)
}

// Returns the elapsed time since the last event for a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408790-cgeventsourcesecondssincelasteve?language=objc
func EventSourceSecondsSinceLastEventType(stateID EventSourceStateID, eventType EventType) corefoundation.TimeInterval {
	rv := C.EventSourceSecondsSinceLastEventType(C.int32_t(stateID), C.uint32_t(eventType))
	return corefoundation.TimeInterval(rv)
}

// Returns the smallest rectangle that results from converting the source rectangle values to integers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456348-cgrectintegral?language=objc
func RectIntegral(rect Rect) Rect {
	rv := C.RectIntegral(rect)
	return Rect(rv)
}

// Sets the fill color space in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455151-cgcontextsetfillcolorspace?language=objc
func ContextSetFillColorSpace(c ContextRef, space ColorSpaceRef) {
	C.ContextSetFillColorSpace(c, space)
}

// Returns whether an object at a given index in a PDF array is a PDF number and, if so, retrieves that object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455374-cgpdfarraygetnumber?language=objc
func PDFArrayGetNumber(array unsafe.Pointer, index uint, value *PDFReal) bool {
	rv := C.PDFArrayGetNumber(array, index, value)
	return bool(rv)
}

// Creates a color using a list of intensity values (including alpha), a pattern color space, and a pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455687-cgcolorcreatewithpattern?language=objc
func ColorCreateWithPattern(space ColorSpaceRef, pattern PatternRef, components *float64) ColorRef {
	rv := C.ColorCreateWithPattern(space, pattern, components)
	return ColorRef(rv)
}

// Changes the origin of the user coordinate system in a context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455286-cgcontexttranslatectm?language=objc
func ContextTranslateCTM(c ContextRef, tx float64, ty float64) {
	C.ContextTranslateCTM(c, tx, ty)
}

// Returns a Boolean value indicating whether Quartz is using OpenGL-based window acceleration (Quartz Extreme) to render in a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455721-cgdisplayusesopenglacceleration?language=objc
func DisplayUsesOpenGLAcceleration(display DirectDisplayID) kernel.Boolean_t {
	rv := C.DisplayUsesOpenGLAcceleration(C.uint32_t(display))
	return kernel.Boolean_t(rv)
}

// Sets how sample values are composited by a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455994-cgcontextsetblendmode?language=objc
func ContextSetBlendMode(c ContextRef, mode BlendMode) {
	C.ContextSetBlendMode(c, C.int32_t(mode))
}

// Reports the change in mouse position since the last mouse movement event received by the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456484-cggetlastmousedelta?language=objc
func GetLastMouseDelta(deltaX *int32, deltaY *int32) {
	C.GetLastMouseDelta(deltaX, deltaY)
}

// Get the bounding box of each glyph in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396342-cgfontgetglyphbboxes?language=objc
func FontGetGlyphBBoxes(font FontRef, glyphs *Glyph, count uint, bboxes *Rect) bool {
	rv := C.FontGetGlyphBBoxes(font, glyphs, count, bboxes)
	return bool(rv)
}

// Provides a list of online displays with bounds that include the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454385-cggetdisplayswithpoint?language=objc
func GetDisplaysWithPoint(point Point, maxDisplays uint32, displays *DirectDisplayID, matchingDisplayCount *uint32) Error {
	rv := C.GetDisplaysWithPoint(point, maxDisplays, displays, matchingDisplayCount)
	return Error(rv)
}

// Decrements the retain count of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1556742-cgimagerelease?language=objc
func ImageRelease(image ImageRef) {
	C.ImageRelease(image)
}

// Returns the smallest value for the x-coordinate of the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455948-cgrectgetminx?language=objc
func RectGetMinX(rect Rect) float64 {
	rv := C.RectGetMinX(rect)
	return float64(rv)
}

// Replaces the path in the graphics context with the stroked version of the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454517-cgcontextreplacepathwithstrokedp?language=objc
func ContextReplacePathWithStrokedPath(c ContextRef) {
	C.ContextReplacePathWithStrokedPath(c)
}

// Returns the bytes per row of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456129-cgbitmapcontextgetbytesperrow?language=objc
func BitmapContextGetBytesPerRow(context ContextRef) uint {
	rv := C.BitmapContextGetBytesPerRow(context)
	return uint(rv)
}

// Returns the list of window IDs associated with the specified windows in the current user session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1552209-cgwindowlistcreate?language=objc
func WindowListCreate(option WindowListOption, relativeToWindow WindowID) corefoundation.ArrayRef {
	rv := C.WindowListCreate(C.uint32_t(option), C.uint32_t(relativeToWindow))
	return corefoundation.ArrayRef(rv)
}

// Returns the height and width resulting from a transformation of an existing height and width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454806-cgsizeapplyaffinetransform?language=objc
func SizeApplyAffineTransform(size Size, t AffineTransform) Size {
	rv := C.SizeApplyAffineTransform(size, t)
	return Size(rv)
}

// Posts a Quartz event from an event tap into the event stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455172-cgeventtappostevent?language=objc
func EventTapPostEvent(proxy unsafe.Pointer, event EventRef) {
	C.EventTapPostEvent(proxy, event)
}

// Returns a Boolean value indicating whether a display is in a mirroring set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455558-cgdisplayisinmirrorset?language=objc
func DisplayIsInMirrorSet(display DirectDisplayID) kernel.Boolean_t {
	rv := C.DisplayIsInMirrorSet(C.uint32_t(display))
	return kernel.Boolean_t(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3861800-cgcolorspaceispqbased?language=objc
func ColorSpaceIsPQBased(s ColorSpaceRef) bool {
	rv := C.ColorSpaceIsPQBased(s)
	return bool(rv)
}

// Paints a transparent rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456457-cgcontextclearrect?language=objc
func ContextClearRect(c ContextRef, rect Rect) {
	C.ContextClearRect(c, rect)
}

// Sets the Unicode string associated with a Quartz keyboard event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456028-cgeventkeyboardsetunicodestring?language=objc
func EventKeyboardSetUnicodeString(event EventRef, stringLength objc.Object, unicodeString *kernel.UniChar) {
	C.EventKeyboardSetUnicodeString(event, stringLength, unicodeString)
}

// Releases a captured display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455685-cgdisplayrelease?language=objc
func DisplayRelease(display DirectDisplayID) Error {
	rv := C.DisplayRelease(C.uint32_t(display))
	return Error(rv)
}

// Returns the document for a page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456166-cgpdfpagegetdocument?language=objc
func PDFPageGetDocument(page PDFPageRef) PDFDocumentRef {
	rv := C.PDFPageGetDocument(page)
	return PDFDocumentRef(rv)
}

// Gets the values in the RGB gamma tables for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454974-cggetdisplaytransferbytable?language=objc
func GetDisplayTransferByTable(display DirectDisplayID, capacity uint32, redTable *GammaValue, greenTable *GammaValue, blueTable *GammaValue, sampleCount *uint32) Error {
	rv := C.GetDisplayTransferByTable(C.uint32_t(display), capacity, redTable, greenTable, blueTable, sampleCount)
	return Error(rv)
}

// Checks whether two affine transforms are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455732-cgaffinetransformequaltotransfor?language=objc
func AffineTransformEqualToTransform(t1 AffineTransform, t2 AffineTransform) bool {
	rv := C.AffineTransformEqualToTransform(t1, t2)
	return bool(rv)
}

// Returns the content stream associated with a PDF scanner object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454724-cgpdfscannergetcontentstream?language=objc
func PDFScannerGetContentStream(scanner unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFScannerGetContentStream(scanner)
	return unsafe.Pointer(rv)
}

// Appends a line segment to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411138-cgpathaddlinetopoint?language=objc
func PathAddLineToPoint(path MutablePathRef, m *AffineTransform, x float64, y float64) {
	C.PathAddLineToPoint(path, m, x, y)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656519-cgpreflightlisteneventaccess?language=objc
func PreflightListenEventAccess() bool {
	rv := C.PreflightListenEventAccess()
	return bool(rv)
}

// Returns the event type of a Quartz event (left mouse down, for example). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455634-cgeventgettype?language=objc
func EventGetType(event EventRef) EventType {
	rv := C.EventGetType(event)
	return EventType(rv)
}

// Retrieves a real value object from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456297-cgpdfscannerpopnumber?language=objc
func PDFScannerPopNumber(scanner unsafe.Pointer, value *PDFReal) bool {
	rv := C.PDFScannerPopNumber(scanner, value)
	return bool(rv)
}

// Creates a device-independent color space that is relative to human color perception, according to the CIE L*a*b* standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408879-cgcolorspacecreatelab?language=objc
func ColorSpaceCreateLab(whitePoint *float64, blackPoint *float64, range_ *float64) ColorSpaceRef {
	rv := C.ColorSpaceCreateLab(whitePoint, blackPoint, range_)
	return ColorSpaceRef(rv)
}

// Ends the current page in the PDF graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456122-cgpdfcontextendpage?language=objc
func PDFContextEndPage(context ContextRef) {
	C.PDFContextEndPage(context)
}

// Returns the bounding box of a clipping path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455387-cgcontextgetclipboundingbox?language=objc
func ContextGetClipBoundingBox(c ContextRef) Rect {
	rv := C.ContextGetClipBoundingBox(c)
	return Rect(rv)
}

// Creates an empty PDF operator table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455932-cgpdfoperatortablecreate?language=objc
func PDFOperatorTableCreate() unsafe.Pointer {
	rv := C.PDFOperatorTableCreate()
	return unsafe.Pointer(rv)
}

// Appends a rectangle to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411144-cgpathaddrect?language=objc
func PathAddRect(path MutablePathRef, m *AffineTransform, rect Rect) {
	C.PathAddRect(path, m, rect)
}

// Returns the point resulting from an affine transformation of an existing point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454251-cgpointapplyaffinetransform?language=objc
func PointApplyAffineTransform(point Point, t AffineTransform) Point {
	rv := C.PointApplyAffineTransform(point, t)
	return Point(rv)
}

// Creates a bitmap image using JPEG-encoded data supplied by a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454920-cgimagecreatewithjpegdataprovide?language=objc
func ImageCreateWithJPEGDataProvider(source DataProviderRef, decode *float64, shouldInterpolate bool, intent ColorRenderingIntent) ImageRef {
	rv := C.ImageCreateWithJPEGDataProvider(source, decode, shouldInterpolate, C.int32_t(intent))
	return ImageRef(rv)
}

// Returns the type identifier for the opaque type CGEventSourceRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408789-cgeventsourcegettypeid?language=objc
func EventSourceGetTypeID() corefoundation.TypeID {
	rv := C.EventSourceGetTypeID()
	return corefoundation.TypeID(rv)
}

// Creates a bitmap image from data supplied by a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455149-cgimagecreate?language=objc
func ImageCreate(width uint, height uint, bitsPerComponent uint, bitsPerPixel uint, bytesPerRow uint, space ColorSpaceRef, bitmapInfo BitmapInfo, provider DataProviderRef, decode *float64, shouldInterpolate bool, intent ColorRenderingIntent) ImageRef {
	rv := C.ImageCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, space, C.uint32_t(bitmapInfo), provider, decode, shouldInterpolate, C.int32_t(intent))
	return ImageRef(rv)
}

// Returns whether the rectangle is equal to the null rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455471-cgrectisnull?language=objc
func RectIsNull(rect Rect) bool {
	rv := C.RectIsNull(rect)
	return bool(rv)
}

// Returns the source state associated with a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408772-cgeventsourcegetsourcestateid?language=objc
func EventSourceGetSourceStateID(source EventSourceRef) EventSourceStateID {
	rv := C.EventSourceGetSourceStateID(source)
	return EventSourceStateID(rv)
}

// Increments the retain count of a CGGradient object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1398456-cggradientretain?language=objc
func GradientRetain(gradient GradientRef) GradientRef {
	rv := C.GradientRetain(gradient)
	return GradientRef(rv)
}

// Returns whether an object at a given index in a PDF array is a PDF Boolean and, if so, retrieves that Boolean. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454504-cgpdfarraygetboolean?language=objc
func PDFArrayGetBoolean(array unsafe.Pointer, index uint, value *PDFBoolean) bool {
	rv := C.PDFArrayGetBoolean(array, index, value)
	return bool(rv)
}

// Sets the current fill color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455296-cgcontextsetfillcolor?language=objc
func ContextSetFillColor(c ContextRef, components *float64) {
	C.ContextSetFillColor(c, components)
}

// Returns the pattern associated with a color in a pattern color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455937-cgcolorgetpattern?language=objc
func ColorGetPattern(color ColorRef) PatternRef {
	rv := C.ColorGetPattern(color)
	return PatternRef(rv)
}

// Returns the bounding box of a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411200-cgpathgetpathboundingbox?language=objc
func PathGetPathBoundingBox(path unsafe.Pointer) Rect {
	rv := C.PathGetPathBoundingBox(path)
	return Rect(rv)
}

// Captures all attached displays, using the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456514-cgcapturealldisplayswithoptions?language=objc
func CaptureAllDisplaysWithOptions(options CaptureOptions) Error {
	rv := C.CaptureAllDisplaysWithOptions(C.uint32_t(options))
	return Error(rv)
}

// Generates and returns information about windows with the specified window IDs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455215-cgwindowlistcreatedescriptionfro?language=objc
func WindowListCreateDescriptionFromArray(windowArray corefoundation.ArrayRef) corefoundation.ArrayRef {
	rv := C.WindowListCreateDescriptionFromArray(windowArray)
	return corefoundation.ArrayRef(rv)
}

// Creates a PDF content stream object from an existing PDF content stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410041-cgpdfcontentstreamcreatewithstre?language=objc
func PDFContentStreamCreateWithStream(stream unsafe.Pointer, streamResources unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFContentStreamCreateWithStream(stream, streamResources, parent)
	return unsafe.Pointer(rv)
}

// Paints a gradient fill that varies along the area defined by the provided starting and ending circles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455923-cgcontextdrawradialgradient?language=objc
func ContextDrawRadialGradient(c ContextRef, gradient GradientRef, startCenter Point, startRadius float64, endCenter Point, endRadius float64, options GradientDrawingOptions) {
	C.ContextDrawRadialGradient(c, gradient, startCenter, startRadius, endCenter, endRadius, C.uint32_t(options))
}

// Converts a string to a date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454295-cgpdfstringcopydate?language=objc
func PDFStringCopyDate(string unsafe.Pointer) corefoundation.DateRef {
	rv := C.PDFStringCopyDate(string_)
	return corefoundation.DateRef(rv)
}

// Returns a Core Foundation Mach port (CFMachPort) that corresponds to the macOS window server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541813-cgwindowservercfmachport?language=objc
func WindowServerCFMachPort() corefoundation.MachPortRef {
	rv := C.WindowServerCFMachPort()
	return corefoundation.MachPortRef(rv)
}

// Provides a list of displays that are online (active, mirrored, or sleeping). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454964-cggetonlinedisplaylist?language=objc
func GetOnlineDisplayList(maxDisplays uint32, onlineDisplays *DirectDisplayID, displayCount *uint32) Error {
	rv := C.GetOnlineDisplayList(maxDisplays, onlineDisplays, displayCount)
	return Error(rv)
}

// Returns a point that is transformed from user space coordinates to device space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455916-cgcontextconvertpointtodevicespa?language=objc
func ContextConvertPointToDeviceSpace(c ContextRef, point Point) Point {
	rv := C.ContextConvertPointToDeviceSpace(c, point)
	return Point(rv)
}

// Creates an immutable copy of a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411211-cgpathcreatecopy?language=objc
func PathCreateCopy(path unsafe.Pointer) unsafe.Pointer {
	rv := C.PathCreateCopy(path)
	return unsafe.Pointer(rv)
}

// Creates and returns a CGImage from the pixel data in a bitmap graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454225-cgbitmapcontextcreateimage?language=objc
func BitmapContextCreateImage(context ContextRef) ImageRef {
	rv := C.BitmapContextCreateImage(context)
	return ImageRef(rv)
}

// Returns the number of color components in a color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408848-cgcolorspacegetnumberofcomponent?language=objc
func ColorSpaceGetNumberOfComponents(space ColorSpaceRef) uint {
	rv := C.ColorSpaceGetNumberOfComponents(space)
	return uint(rv)
}

// Increments the retain count of a function object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1390869-cgfunctionretain?language=objc
func FunctionRetain(function FunctionRef) FunctionRef {
	rv := C.FunctionRetain(function)
	return FunctionRef(rv)
}

// Gets the coefficients of the gamma transfer formula for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456330-cggetdisplaytransferbyformula?language=objc
func GetDisplayTransferByFormula(display DirectDisplayID, redMin *GammaValue, redMax *GammaValue, redGamma *GammaValue, greenMin *GammaValue, greenMax *GammaValue, greenGamma *GammaValue, blueMin *GammaValue, blueMax *GammaValue, blueGamma *GammaValue) Error {
	rv := C.GetDisplayTransferByFormula(C.uint32_t(display), redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
	return Error(rv)
}

// Retrieves a PDF stream object from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454561-cgpdfscannerpopstream?language=objc
func PDFScannerPopStream(scanner unsafe.Pointer, value unsafe.Pointer) bool {
	rv := C.PDFScannerPopStream(scanner, value)
	return bool(rv)
}

// Increments the retain count of a color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408885-cgcolorspaceretain?language=objc
func ColorSpaceRetain(space ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceRetain(space)
	return ColorSpaceRef(rv)
}

// Creates a color in the sRGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3042355-cgcolorcreatesrgb?language=objc
func ColorCreateSRGB(red float64, green float64, blue float64, alpha float64) ColorRef {
	rv := C.ColorCreateSRGB(red, green, blue, alpha)
	return ColorRef(rv)
}

// Returns an array of rectangles that describe where the frame has changed since the previous frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455530-cgdisplaystreamupdategetrects?language=objc
func DisplayStreamUpdateGetRects(updateRef DisplayStreamUpdateRef, rectType DisplayStreamUpdateRectType, rectCount *uint) *Rect {
	rv := C.DisplayStreamUpdateGetRects(updateRef, C.int32_t(rectType), rectCount)
	return *Rect(rv)
}

// Return the color space for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454858-cgimagegetcolorspace?language=objc
func ImageGetColorSpace(image ImageRef) ColorSpaceRef {
	rv := C.ImageGetColorSpace(image)
	return ColorSpaceRef(rv)
}

// Filters local hardware events from the keyboard and mouse during the short interval after a synthetic event is posted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541776-cgsetlocaleventsfilterduringsupp?language=objc
func SetLocalEventsFilterDuringSuppressionState(filter EventFilterMask, state EventSuppressionState) Error {
	rv := C.SetLocalEventsFilterDuringSuppressionState(C.uint32_t(filter), C.uint32_t(state))
	return Error(rv)
}

// Creates a shading object to use for radial shading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456399-cgshadingcreateradial?language=objc
func ShadingCreateRadial(space ColorSpaceRef, start Point, startRadius float64, end Point, endRadius float64, function FunctionRef, extendStart bool, extendEnd bool) ShadingRef {
	rv := C.ShadingCreateRadial(space, start, startRadius, end, endRadius, function, extendStart, extendEnd)
	return ShadingRef(rv)
}

// Divides a source rectangle into two component rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455925-cgrectdivide?language=objc
func RectDivide(rect Rect, slice *Rect, remainder *Rect, amount float64, edge RectEdge) {
	C.RectDivide(rect, slice, remainder, amount, C.uint32_t(edge))
}

// Sets whether or not to allow antialiasing for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456310-cgcontextsetallowsantialiasing?language=objc
func ContextSetAllowsAntialiasing(c ContextRef, allowsAntialiasing bool) {
	C.ContextSetAllowsAntialiasing(c, allowsAntialiasing)
}

// Returns the height in pixels of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454681-cgbitmapcontextgetheight?language=objc
func BitmapContextGetHeight(context ContextRef) uint {
	rv := C.BitmapContextGetHeight(context)
	return uint(rv)
}

// Returns the value of the alpha component associated with a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456637-cgcolorgetalpha?language=objc
func ColorGetAlpha(color ColorRef) float64 {
	rv := C.ColorGetAlpha(color)
	return float64(rv)
}

// Returns the height of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455829-cgimagegetheight?language=objc
func ImageGetHeight(image ImageRef) uint {
	rv := C.ImageGetHeight(image)
	return uint(rv)
}

// Creates a specified type of Quartz color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408921-cgcolorspacecreatewithname?language=objc
func ColorSpaceCreateWithName(name corefoundation.StringRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateWithName(name)
	return ColorSpaceRef(rv)
}

// Appends a quadratic Bézier curve from the current point, using a control point and an end point you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454268-cgcontextaddquadcurvetopoint?language=objc
func ContextAddQuadCurveToPoint(c ContextRef, cpx float64, cpy float64, x float64, y float64) {
	C.ContextAddQuadCurveToPoint(c, cpx, cpy, x, y)
}

// Configures the display mode of a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454273-cgconfiguredisplaywithdisplaymod?language=objc
func ConfigureDisplayWithDisplayMode(config unsafe.Pointer, display DirectDisplayID, mode DisplayModeRef, options corefoundation.DictionaryRef) Error {
	rv := C.ConfigureDisplayWithDisplayMode(config, C.uint32_t(display), mode, options)
	return Error(rv)
}

// Returns a copy of the provider’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408309-cgdataprovidercopydata?language=objc
func DataProviderCopyData(provider DataProviderRef) corefoundation.DataRef {
	rv := C.DataProviderCopyData(provider)
	return corefoundation.DataRef(rv)
}

// Returns a Boolean value indicating whether a display is in a hardware mirroring set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454506-cgdisplayisinhwmirrorset?language=objc
func DisplayIsInHWMirrorSet(display DirectDisplayID) kernel.Boolean_t {
	rv := C.DisplayIsInHWMirrorSet(C.uint32_t(display))
	return kernel.Boolean_t(rv)
}

// Returns the mask that indicates which classes of local hardware events are enabled during event suppression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408785-cgeventsourcegetlocaleventsfilte?language=objc
func EventSourceGetLocalEventsFilterDuringSuppressionState(source EventSourceRef, state EventSuppressionState) EventFilterMask {
	rv := C.EventSourceGetLocalEventsFilterDuringSuppressionState(source, C.uint32_t(state))
	return EventFilterMask(rv)
}

// Returns the font table that corresponds to the provided tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396402-cgfontcopytablefortag?language=objc
func FontCopyTableForTag(font FontRef, tag uint32) corefoundation.DataRef {
	rv := C.FontCopyTableForTag(font, tag)
	return corefoundation.DataRef(rv)
}

// Returns a pointer to the bytes of a PDF string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455978-cgpdfstringgetbyteptr?language=objc
func PDFStringGetBytePtr(string unsafe.Pointer) *uint8 {
	rv := C.PDFStringGetBytePtr(string_)
	return *uint8(rv)
}

// Returns the smallest rectangle that contains the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454577-cgcontextgetpathboundingbox?language=objc
func ContextGetPathBoundingBox(c ContextRef) Rect {
	rv := C.ContextGetPathBoundingBox(c)
	return Rect(rv)
}

// Returns the Unicode string associated with a Quartz keyboard event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456120-cgeventkeyboardgetunicodestring?language=objc
func EventKeyboardGetUnicodeString(event EventRef, maxStringLength objc.Object, actualStringLength objc.Object, unicodeString *kernel.UniChar) {
	C.EventKeyboardGetUnicodeString(event, maxStringLength, actualStringLength, unicodeString)
}

// Increments the retain count of a layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450900-cglayerretain?language=objc
func LayerRetain(layer LayerRef) LayerRef {
	rv := C.LayerRetain(layer)
	return LayerRef(rv)
}

// Sets whether or not to allow subpixel positioning for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454942-cgcontextsetallowsfontsubpixelpo?language=objc
func ContextSetAllowsFontSubpixelPositioning(c ContextRef, allowsFontSubpixelPositioning bool) {
	C.ContextSetAllowsFontSubpixelPositioning(c, allowsFontSubpixelPositioning)
}

// Returns the Core Foundation type identifier for a color data type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455568-cgcolorgettypeid?language=objc
func ColorGetTypeID() corefoundation.TypeID {
	rv := C.ColorGetTypeID()
	return corefoundation.TypeID(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2962830-cgimagegetbyteorderinfo?language=objc
func ImageGetByteOrderInfo(image ImageRef) ImageByteOrderInfo {
	rv := C.ImageGetByteOrderInfo(image)
	return ImageByteOrderInfo(rv)
}

// Creates a Core Graphics PDF document using a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402603-cgpdfdocumentcreatewithprovider?language=objc
func PDFDocumentCreateWithProvider(provider DataProviderRef) PDFDocumentRef {
	rv := C.PDFDocumentCreateWithProvider(provider)
	return PDFDocumentRef(rv)
}

// Gets the advance width of each glyph in the provided array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396332-cgfontgetglyphadvances?language=objc
func FontGetGlyphAdvances(font FontRef, glyphs *Glyph, count uint, advances *int) bool {
	rv := C.FontGetGlyphAdvances(font, glyphs, count, advances)
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2962831-cgimagegetpixelformatinfo?language=objc
func ImageGetPixelFormatInfo(image ImageRef) ImagePixelFormatInfo {
	rv := C.ImageGetPixelFormatInfo(image)
	return ImagePixelFormatInfo(rv)
}

// Returns an array of the variation axis dictionaries for a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396376-cgfontcopyvariationaxes?language=objc
func FontCopyVariationAxes(font FontRef) corefoundation.ArrayRef {
	rv := C.FontCopyVariationAxes(font)
	return corefoundation.ArrayRef(rv)
}

// Returns the media box of a page in a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402592-cgpdfdocumentgetmediabox?language=objc
func PDFDocumentGetMediaBox(document PDFDocumentRef, page int) Rect {
	rv := C.PDFDocumentGetMediaBox(document, page)
	return Rect(rv)
}

// Returns a copy of the color space's properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2962828-cgcolorspacecopypropertylist?language=objc
func ColorSpaceCopyPropertyList(space ColorSpaceRef) corefoundation.PropertyListRef {
	rv := C.ColorSpaceCopyPropertyList(space)
	return corefoundation.PropertyListRef(rv)
}

// Sets the stroke pattern in the specified graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454796-cgcontextsetstrokepattern?language=objc
func ContextSetStrokePattern(c ContextRef, pattern PatternRef, components *float64) {
	C.ContextSetStrokePattern(c, pattern, components)
}

// Returns information about the display mode closest to a specified depth and screen size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1562060-cgdisplaybestmodeforparameters?language=objc
func DisplayBestModeForParameters(display DirectDisplayID, bitsPerPixel uint, width uint, height uint, exactMatch *kernel.Boolean_t) corefoundation.DictionaryRef {
	rv := C.DisplayBestModeForParameters(C.uint32_t(display), bitsPerPixel, width, height, exactMatch)
	return corefoundation.DictionaryRef(rv)
}

// Returns whether there is another PDF dictionary associated with a specified key in a PDF dictionary and, if so, retrieves that dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430220-cgpdfdictionarygetdictionary?language=objc
func PDFDictionaryGetDictionary(dict unsafe.Pointer, key *uint8, value unsafe.Pointer) bool {
	rv := C.PDFDictionaryGetDictionary(dict, key, value)
	return bool(rv)
}

// Sets the clipping path to the intersection of the current clipping path with the area defined by the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454716-cgcontextcliptorect?language=objc
func ContextClipToRect(c ContextRef, rect Rect) {
	C.ContextClipToRect(c, rect)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2881687-cgcontextresetclip?language=objc
func ContextResetClip(c ContextRef) {
	C.ContextResetClip(c)
}

// Returns whether an object at a given index in a PDF array is a PDF dictionary and, if so, retrieves that dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454139-cgpdfarraygetdictionary?language=objc
func PDFArrayGetDictionary(array unsafe.Pointer, index uint, value unsafe.Pointer) bool {
	rv := C.PDFArrayGetDictionary(array, index, value)
	return bool(rv)
}

// Appends a quadratic Bézier curve to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411157-cgpathaddquadcurvetopoint?language=objc
func PathAddQuadCurveToPoint(path MutablePathRef, m *AffineTransform, cpx float64, cpy float64, x float64, y float64) {
	C.PathAddQuadCurveToPoint(path, m, cpx, cpy, x, y)
}

// Moves the mouse cursor without generating events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456387-cgwarpmousecursorposition?language=objc
func WarpMouseCursorPosition(newCursorPosition Point) Error {
	rv := C.WarpMouseCursorPosition(newCursorPosition)
	return Error(rv)
}

// Moves the mouse cursor to a specified point relative to the upper-left corner of the display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455258-cgdisplaymovecursortopoint?language=objc
func DisplayMoveCursorToPoint(display DirectDisplayID, point Point) Error {
	rv := C.DisplayMoveCursorToPoint(C.uint32_t(display), point)
	return Error(rv)
}

// Generates and returns information about the selected windows in the current user session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455137-cgwindowlistcopywindowinfo?language=objc
func WindowListCopyWindowInfo(option WindowListOption, relativeToWindow WindowID) corefoundation.ArrayRef {
	rv := C.WindowListCopyWindowInfo(C.uint32_t(option), C.uint32_t(relativeToWindow))
	return corefoundation.ArrayRef(rv)
}

// Returns the window ID of the shield window for a captured display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454279-cgshieldingwindowid?language=objc
func ShieldingWindowID(display DirectDisplayID) WindowID {
	rv := C.ShieldingWindowID(C.uint32_t(display))
	return WindowID(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2875542-cgpdfcontextsetoutline?language=objc
func PDFContextSetOutline(context ContextRef, outline corefoundation.DictionaryRef) {
	C.PDFContextSetOutline(context, outline)
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411173-cgpathaddarctopoint?language=objc
func PathAddArcToPoint(path MutablePathRef, m *AffineTransform, x1 float64, y1 float64, x2 float64, y2 float64, radius float64) {
	C.PathAddArcToPoint(path, m, x1, y1, x2, y2, radius)
}

// Returns the number of bits allocated for a single color component of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454980-cgimagegetbitspercomponent?language=objc
func ImageGetBitsPerComponent(image ImageRef) uint {
	rv := C.ImageGetBitsPerComponent(image)
	return uint(rv)
}

// Obtains exclusive use of all active displays, preventing other applications and system services from using the display or changing its configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455221-cgcapturealldisplays?language=objc
func CaptureAllDisplays() Error {
	rv := C.CaptureAllDisplays()
	return Error(rv)
}

// Gets the scale of pixels per line in a scrolling event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408775-cgeventsourcegetpixelsperline?language=objc
func EventSourceGetPixelsPerLine(source EventSourceRef) float64 {
	rv := C.EventSourceGetPixelsPerLine(source)
	return float64(rv)
}

// Returns a rectangle that is smaller or larger than the source rectangle, with the same center point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454218-cgrectinset?language=objc
func RectInset(rect Rect, dx float64, dy float64) Rect {
	rv := C.RectInset(rect, dx, dy)
	return Rect(rv)
}

// Begins a new set of display configuration changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455235-cgbegindisplayconfiguration?language=objc
func BeginDisplayConfiguration(config unsafe.Pointer) Error {
	rv := C.BeginDisplayConfiguration(config)
	return Error(rv)
}

// Returns whether an object at a given index in a Quartz PDF array is a PDF null. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456173-cgpdfarraygetnull?language=objc
func PDFArrayGetNull(array unsafe.Pointer, index uint) bool {
	rv := C.PDFArrayGetNull(array, index)
	return bool(rv)
}

// Returns the decode array for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454575-cgimagegetdecode?language=objc
func ImageGetDecode(image ImageRef) *float64 {
	rv := C.ImageGetDecode(image)
	return *float64(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2866188-cgdataprovidergetinfo?language=objc
func DataProviderGetInfo(provider DataProviderRef) unsafe.Pointer {
	rv := C.DataProviderGetInfo(provider)
	return unsafe.Pointer(rv)
}

// Combines two updates into a new update that includes the metadata for both source updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454341-cgdisplaystreamupdatecreatemerge?language=objc
func DisplayStreamUpdateCreateMergedUpdate(firstUpdate DisplayStreamUpdateRef, secondUpdate DisplayStreamUpdateRef) DisplayStreamUpdateRef {
	rv := C.DisplayStreamUpdateCreateMergedUpdate(firstUpdate, secondUpdate)
	return DisplayStreamUpdateRef(rv)
}

// Returns the data associated with a PDF stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454657-cgpdfstreamcopydata?language=objc
func PDFStreamCopyData(stream unsafe.Pointer, format *PDFDataFormat) corefoundation.DataRef {
	rv := C.PDFStreamCopyData(stream, format)
	return corefoundation.DataRef(rv)
}

// Increments the retain count of a Core Graphics pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1552265-cgpatternretain?language=objc
func PatternRetain(pattern PatternRef) PatternRef {
	rv := C.PatternRetain(pattern)
	return PatternRef(rv)
}

// Returns whether there is a PDF object associated with a specified key in a PDF dictionary and, if so, retrieves that object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430214-cgpdfdictionarygetobject?language=objc
func PDFDictionaryGetObject(dict unsafe.Pointer, key *uint8, value unsafe.Pointer) bool {
	rv := C.PDFDictionaryGetObject(dict, key, value)
	return bool(rv)
}

// Returns the color space of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454058-cgbitmapcontextgetcolorspace?language=objc
func BitmapContextGetColorSpace(context ContextRef) ColorSpaceRef {
	rv := C.BitmapContextGetColorSpace(context)
	return ColorSpaceRef(rv)
}

// Returns whether a rectangle has zero width or height, or is a null rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454917-cgrectisempty?language=objc
func RectIsEmpty(rect Rect) bool {
	rv := C.RectIsEmpty(rect)
	return bool(rv)
}

// Returns the smallest value for the y-coordinate of the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454832-cgrectgetminy?language=objc
func RectGetMinY(rect Rect) float64 {
	rv := C.RectGetMinY(rect)
	return float64(rv)
}

// Returns a composite image based on a dynamically generated list of windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454852-cgwindowlistcreateimage?language=objc
func WindowListCreateImage(screenBounds Rect, listOption WindowListOption, windowID WindowID, imageOption WindowImageOption) ImageRef {
	rv := C.WindowListCreateImage(screenBounds, C.uint32_t(listOption), C.uint32_t(windowID), C.uint32_t(imageOption))
	return ImageRef(rv)
}

// Synthesizes a low-level scrolling event on the local machine. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541788-cgpostscrollwheelevent?language=objc
func PostScrollWheelEvent(wheelCount WheelCount, wheel1 int32) Error {
	rv := C.PostScrollWheelEvent(C.uint32_t(wheelCount), wheel1)
	return Error(rv)
}

// Returns the Core Foundation type identifier for Quartz color spaces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408926-cgcolorspacegettypeid?language=objc
func ColorSpaceGetTypeID() corefoundation.TypeID {
	rv := C.ColorSpaceGetTypeID()
	return corefoundation.TypeID(rv)
}

// Increments the retain count of a shading object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1573767-cgshadingretain?language=objc
func ShadingRetain(shading ShadingRef) ShadingRef {
	rv := C.ShadingRetain(shading)
	return ShadingRef(rv)
}

// Creates a copy of an existing color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456134-cgcolorcreatecopy?language=objc
func ColorCreateCopy(color ColorRef) ColorRef {
	rv := C.ColorCreateCopy(color)
	return ColorRef(rv)
}

// Begins a transparency layer whose contents are bounded by the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454368-cgcontextbegintransparencylayerw?language=objc
func ContextBeginTransparencyLayerWithRect(c ContextRef, rect Rect, auxInfo corefoundation.DictionaryRef) {
	C.ContextBeginTransparencyLayerWithRect(c, rect, auxInfo)
}

// Decrements the retain count of a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1571724-cgpdfpagerelease?language=objc
func PDFPageRelease(page PDFPageRef) {
	C.PDFPageRelease(page)
}

// Returns the type identifier for CGImage objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455014-cgimagegettypeid?language=objc
func ImageGetTypeID() corefoundation.TypeID {
	rv := C.ImageGetTypeID()
	return corefoundation.TypeID(rv)
}

// Retrieves a PDF dictionary object from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456538-cgpdfscannerpopdictionary?language=objc
func PDFScannerPopDictionary(scanner unsafe.Pointer, value unsafe.Pointer) bool {
	rv := C.PDFScannerPopDictionary(scanner, value)
	return bool(rv)
}

// Rotates the user coordinate system in a context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456228-cgcontextrotatectm?language=objc
func ContextRotateCTM(c ContextRef, angle float64) {
	C.ContextRotateCTM(c, angle)
}

// Ends a transparency layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456554-cgcontextendtransparencylayer?language=objc
func ContextEndTransparencyLayer(c ContextRef) {
	C.ContextEndTransparencyLayer(c)
}

// Sets the interval that local hardware events may be suppressed following the posting of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408783-cgeventsourcesetlocaleventssuppr?language=objc
func EventSourceSetLocalEventsSuppressionInterval(source EventSourceRef, seconds corefoundation.TimeInterval) {
	C.EventSourceSetLocalEventsSuppressionInterval(source, C.double(seconds))
}

// Returns the intersection of two rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455346-cgrectintersection?language=objc
func RectIntersection(r1 Rect, r2 Rect) Rect {
	rv := C.RectIntersection(r1, r2)
	return Rect(rv)
}

// Returns the I/O Kit display mode ID of the specified display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454728-cgdisplaymodegetiodisplaymodeid?language=objc
func DisplayModeGetIODisplayModeID(mode DisplayModeRef) int32 {
	rv := C.DisplayModeGetIODisplayModeID(mode)
	return int32(rv)
}

// Returns the primary display in a hardware mirroring set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454403-cgdisplayprimarydisplay?language=objc
func DisplayPrimaryDisplay(display DirectDisplayID) DirectDisplayID {
	rv := C.DisplayPrimaryDisplay(C.uint32_t(display))
	return DirectDisplayID(rv)
}

// Sets a destination to jump to when a point in the current page of a PDF graphics context is clicked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455424-cgpdfcontextadddestinationatpoin?language=objc
func PDFContextAddDestinationAtPoint(context ContextRef, name corefoundation.StringRef, point Point) {
	C.PDFContextAddDestinationAtPoint(context, name, point)
}

// Returns the descent of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396351-cgfontgetdescent?language=objc
func FontGetDescent(font FontRef) int {
	rv := C.FontGetDescent(font)
	return int(rv)
}

// Sets the current stroke color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456283-cgcontextsetstrokecolor?language=objc
func ContextSetStrokeColor(c ContextRef, components *float64) {
	C.ContextSetStrokeColor(c, components)
}

// Returns whether there is a PDF string associated with a specified key in a PDF dictionary and, if so, retrieves that string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430224-cgpdfdictionarygetstring?language=objc
func PDFDictionaryGetString(dict unsafe.Pointer, key *uint8, value unsafe.Pointer) bool {
	rv := C.PDFDictionaryGetString(dict, key, value)
	return bool(rv)
}

// Releases a Core Graphics display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1562069-cgdisplaymoderelease?language=objc
func DisplayModeRelease(mode DisplayModeRef) {
	C.DisplayModeRelease(mode)
}

// Adds a set of rectangular paths to the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454734-cgcontextaddrects?language=objc
func ContextAddRects(c ContextRef, rects *Rect, count uint) {
	C.ContextAddRects(c, rects, count)
}

// Sets the font and font size in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586511-cgcontextselectfont?language=objc
func ContextSelectFont(c ContextRef, name *uint8, size float64, textEncoding TextEncoding) {
	C.ContextSelectFont(c, name, size, C.int32_t(textEncoding))
}

// Creates a mutable copy of a graphics path transformed by a transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411150-cgpathcreatemutablecopybytransfo?language=objc
func PathCreateMutableCopyByTransformingPath(path unsafe.Pointer, transform *AffineTransform) MutablePathRef {
	rv := C.PathCreateMutableCopyByTransformingPath(path, transform)
	return MutablePathRef(rv)
}

// Returns an affine transformation matrix constructed from translation values you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454909-cgaffinetransformmaketranslation?language=objc
func AffineTransformMakeTranslation(tx float64, ty float64) AffineTransform {
	rv := C.AffineTransformMakeTranslation(tx, ty)
	return AffineTransform(rv)
}

// Sets the current character spacing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454786-cgcontextsetcharacterspacing?language=objc
func ContextSetCharacterSpacing(c ContextRef, spacing float64) {
	C.ContextSetCharacterSpacing(c, spacing)
}

// Returns the I/O Kit flags of the specified display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454092-cgdisplaymodegetioflags?language=objc
func DisplayModeGetIOFlags(mode DisplayModeRef) uint32 {
	rv := C.DisplayModeGetIOFlags(mode)
	return uint32(rv)
}

// Closes a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454306-cgpdfcontextclose?language=objc
func PDFContextClose(context ContextRef) {
	C.PDFContextClose(context)
}

// Increments the retain count of a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408276-cgdataproviderretain?language=objc
func DataProviderRetain(provider DataProviderRef) DataProviderRef {
	rv := C.DataProviderRetain(provider)
	return DataProviderRef(rv)
}

// Returns a Boolean indicating whether the color space can be used as a destination color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1690958-cgcolorspacesupportsoutput?language=objc
func ColorSpaceSupportsOutput(space ColorSpaceRef) bool {
	rv := C.ColorSpaceSupportsOutput(space)
	return bool(rv)
}

// Sets the floating-point value of a field in a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455526-cgeventsetdoublevaluefield?language=objc
func EventSetDoubleValueField(event EventRef, field EventField, value float64) {
	C.EventSetDoubleValueField(event, C.uint32_t(field), value)
}

// Draws the content of a PDF page into the current graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456255-cgcontextdrawpdfpage?language=objc
func ContextDrawPDFPage(c ContextRef, page PDFPageRef) {
	C.ContextDrawPDFPage(c, page)
}

// Creates a URL-based PDF graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456290-cgpdfcontextcreatewithurl?language=objc
func PDFContextCreateWithURL(url corefoundation.URLRef, mediaBox *Rect, auxiliaryInfo corefoundation.DictionaryRef) ContextRef {
	rv := C.PDFContextCreateWithURL(url, mediaBox, auxiliaryInfo)
	return ContextRef(rv)
}

// Returns the CFType ID for PDF page objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454478-cgpdfpagegettypeid?language=objc
func PDFPageGetTypeID() corefoundation.TypeID {
	rv := C.PDFPageGetTypeID()
	return corefoundation.TypeID(rv)
}

// Returns the bleed box of a page in a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402596-cgpdfdocumentgetbleedbox?language=objc
func PDFDocumentGetBleedBox(document PDFDocumentRef, page int) Rect {
	rv := C.PDFDocumentGetBleedBox(document, page)
	return Rect(rv)
}

// Creates a content stream object from a PDF page object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410047-cgpdfcontentstreamcreatewithpage?language=objc
func PDFContentStreamCreateWithPage(page PDFPageRef) unsafe.Pointer {
	rv := C.PDFContentStreamCreateWithPage(page)
	return unsafe.Pointer(rv)
}

// Creates a font object from data supplied from a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396367-cgfontcreatewithdataprovider?language=objc
func FontCreateWithDataProvider(provider DataProviderRef) FontRef {
	rv := C.FontCreateWithDataProvider(provider)
	return FontRef(rv)
}

// Returns whether there is a PDF number associated with a specified key in a PDF dictionary and, if so, retrieves that number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430228-cgpdfdictionarygetnumber?language=objc
func PDFDictionaryGetNumber(dict unsafe.Pointer, key *uint8, value *PDFReal) bool {
	rv := C.PDFDictionaryGetNumber(dict, key, value)
	return bool(rv)
}

// Decrements the retain count of a color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408855-cgcolorspacerelease?language=objc
func ColorSpaceRelease(space ColorSpaceRef) {
	C.ColorSpaceRelease(space)
}

// Creates a copy of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455615-cgimagecreatecopy?language=objc
func ImageCreateCopy(image ImageRef) ImageRef {
	rv := C.ImageCreateCopy(image)
	return ImageRef(rv)
}

// Ends the current page in a page-based graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455027-cgcontextendpage?language=objc
func ContextEndPage(c ContextRef) {
	C.ContextEndPage(c)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3325503-cgcolorspaceishdr?language=objc
func ColorSpaceIsHDR(arg0 ColorSpaceRef) bool {
	rv := C.ColorSpaceIsHDR(arg0)
	return bool(rv)
}

// Gets the file identifier for a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402600-cgpdfdocumentgetid?language=objc
func PDFDocumentGetID(document PDFDocumentRef) unsafe.Pointer {
	rv := C.PDFDocumentGetID(document)
	return unsafe.Pointer(rv)
}

// Sets the current text matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455611-cgcontextsettextmatrix?language=objc
func ContextSetTextMatrix(c ContextRef, t AffineTransform) {
	C.ContextSetTextMatrix(c, t)
}

// Returns the number of glyphs in a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396371-cgfontgetnumberofglyphs?language=objc
func FontGetNumberOfGlyphs(font FontRef) uint {
	rv := C.FontGetNumberOfGlyphs(font)
	return uint(rv)
}

// Maps a mask into the specified rectangle and intersects it with the current clipping area of the graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456497-cgcontextcliptomask?language=objc
func ContextClipToMask(c ContextRef, rect Rect, mask ImageRef) {
	C.ContextClipToMask(c, rect, mask)
}

// Returns the dictionary of a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455125-cgpdfpagegetdictionary?language=objc
func PDFPageGetDictionary(page PDFPageRef) unsafe.Pointer {
	rv := C.PDFPageGetDictionary(page)
	return unsafe.Pointer(rv)
}

// Returns the capacity, or number of entries, in the gamma table for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454310-cgdisplaygammatablecapacity?language=objc
func DisplayGammaTableCapacity(display DirectDisplayID) uint32 {
	rv := C.DisplayGammaTableCapacity(C.uint32_t(display))
	return uint32(rv)
}

// Sets the integer value of a field in a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455556-cgeventsetintegervaluefield?language=objc
func EventSetIntegerValueField(event EventRef, field EventField, value int64) {
	C.EventSetIntegerValueField(event, C.uint32_t(field), value)
}

// Adds a previously created path object to the current path in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456628-cgcontextaddpath?language=objc
func ContextAddPath(c ContextRef, path unsafe.Pointer) {
	C.ContextAddPath(c, path)
}

// Returns the current text matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456154-cgcontextgettextmatrix?language=objc
func ContextGetTextMatrix(c ContextRef) AffineTransform {
	rv := C.ContextGetTextMatrix(c)
	return AffineTransform(rv)
}

// Paints a line along the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454490-cgcontextstrokepath?language=objc
func ContextStrokePath(c ContextRef) {
	C.ContextStrokePath(c)
}

// Returns the display width in pixel units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456361-cgdisplaypixelswide?language=objc
func DisplayPixelsWide(display DirectDisplayID) uint {
	rv := C.DisplayPixelsWide(C.uint32_t(display))
	return uint(rv)
}

// Sets antialiasing on or off for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455178-cgcontextsetshouldantialias?language=objc
func ContextSetShouldAntialias(c ContextRef, shouldAntialias bool) {
	C.ContextSetShouldAntialias(c, shouldAntialias)
}

// Switches a display to a different mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1562065-cgdisplayswitchtomode?language=objc
func DisplaySwitchToMode(display DirectDisplayID, mode corefoundation.DictionaryRef) Error {
	rv := C.DisplaySwitchToMode(C.uint32_t(display), mode)
	return Error(rv)
}

// Gets a list of currently installed event taps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455395-cggeteventtaplist?language=objc
func GetEventTapList(maxNumberOfTaps uint32, tapList *EventTapInformation, eventTapCount *uint32) Error {
	rv := C.GetEventTapList(maxNumberOfTaps, tapList, eventTapCount)
	return Error(rv)
}

// Returns a Boolean value indicating whether an event tap is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456102-cgeventtapisenabled?language=objc
func EventTapIsEnabled(tap corefoundation.MachPortRef) bool {
	rv := C.EventTapIsEnabled(tap)
	return bool(rv)
}

// Returns whether there is a PDF stream associated with a specified key in a PDF dictionary and, if so, retrieves that stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430213-cgpdfdictionarygetstream?language=objc
func PDFDictionaryGetStream(dict unsafe.Pointer, key *uint8, value unsafe.Pointer) bool {
	rv := C.PDFDictionaryGetStream(dict, key, value)
	return bool(rv)
}

// Creates an indexed color space, consisting of colors specified by a color lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408899-cgcolorspacecreateindexed?language=objc
func ColorSpaceCreateIndexed(baseSpace ColorSpaceRef, lastIndex uint, colorTable *uint8) ColorSpaceRef {
	rv := C.ColorSpaceCreateIndexed(baseSpace, lastIndex, colorTable)
	return ColorSpaceRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3042409-cgpdftagtypegetname?language=objc
func PDFTagTypeGetName(tagType PDFTagType) *uint8 {
	rv := C.PDFTagTypeGetName(C.int32_t(tagType))
	return *uint8(rv)
}

// Sets the current fill color to a value in the DeviceGray color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454255-cgcontextsetgrayfillcolor?language=objc
func ContextSetGrayFillColor(c ContextRef, gray float64, alpha float64) {
	C.ContextSetGrayFillColor(c, gray, alpha)
}

// Gets the specified resource from a PDF content stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410044-cgpdfcontentstreamgetresource?language=objc
func PDFContentStreamGetResource(cs unsafe.Pointer, category *uint8, name *uint8) unsafe.Pointer {
	rv := C.PDFContentStreamGetResource(cs, category, name)
	return unsafe.Pointer(rv)
}

// Returns whether there is a PDF Boolean value associated with a specified key in a PDF dictionary and, if so, retrieves the Boolean value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430226-cgpdfdictionarygetboolean?language=objc
func PDFDictionaryGetBoolean(dict unsafe.Pointer, key *uint8, value *PDFBoolean) bool {
	rv := C.PDFDictionaryGetBoolean(dict, key, value)
	return bool(rv)
}

// Uses a PostScript converter to convert PostScript data to PDF data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455368-cgpsconverterconvert?language=objc
func PSConverterConvert(converter PSConverterRef, provider DataProviderRef, consumer DataConsumerRef, options corefoundation.DictionaryRef) bool {
	rv := C.PSConverterConvert(converter, provider, consumer, options)
	return bool(rv)
}

// Decrements the retain count of a PDF content stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410046-cgpdfcontentstreamrelease?language=objc
func PDFContentStreamRelease(cs unsafe.Pointer) {
	C.PDFContentStreamRelease(cs)
}

// Returns the crop box of a page in a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402598-cgpdfdocumentgetcropbox?language=objc
func PDFDocumentGetCropBox(document PDFDocumentRef, page int) Rect {
	rv := C.PDFDocumentGetCropBox(document, page)
	return Rect(rv)
}

// Maps a display ID to an OpenGL display mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455554-cgdisplayidtoopengldisplaymask?language=objc
func DisplayIDToOpenGLDisplayMask(display DirectDisplayID) OpenGLDisplayMask {
	rv := C.DisplayIDToOpenGLDisplayMask(C.uint32_t(display))
	return OpenGLDisplayMask(rv)
}

// Returns the largest value of the x-coordinate for the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454334-cgrectgetmaxx?language=objc
func RectGetMaxX(rect Rect) float64 {
	rv := C.RectGetMaxX(rect)
	return float64(rv)
}

// Indicates whether the current path contains any subpaths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455772-cgcontextispathempty?language=objc
func ContextIsPathEmpty(c ContextRef) bool {
	rv := C.ContextIsPathEmpty(c)
	return bool(rv)
}

// Returns the Core Foundation type identifier for a color conversion info data type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2113681-cgcolorconversioninfogettypeid?language=objc
func ColorConversionInfoGetTypeID() corefoundation.TypeID {
	rv := C.ColorConversionInfoGetTypeID()
	return corefoundation.TypeID(rv)
}

// Returns the number of bits allocated for a single pixel in a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454599-cgimagegetbitsperpixel?language=objc
func ImageGetBitsPerPixel(image ImageRef) uint {
	rv := C.ImageGetBitsPerPixel(image)
	return uint(rv)
}

// Fills in a rectangle using the contents of the specified dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456558-cgrectmakewithdictionaryrepresen?language=objc
func RectMakeWithDictionaryRepresentation(dict corefoundation.DictionaryRef, rect *Rect) bool {
	rv := C.RectMakeWithDictionaryRepresentation(dict, rect)
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456406-cgdisplaymodegetpixelheight?language=objc
func DisplayModeGetPixelHeight(mode DisplayModeRef) uint {
	rv := C.DisplayModeGetPixelHeight(mode)
	return uint(rv)
}

// Creates a device-dependent RGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408837-cgcolorspacecreatedevicergb?language=objc
func ColorSpaceCreateDeviceRGB() ColorSpaceRef {
	rv := C.ColorSpaceCreateDeviceRGB()
	return ColorSpaceRef(rv)
}

// Sets the current stroke color in a context, using a CGColor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456196-cgcontextsetstrokecolorwithcolor?language=objc
func ContextSetStrokeColorWithColor(c ContextRef, color ColorRef) {
	C.ContextSetStrokeColorWithColor(c, color)
}

// Returns the graphics context associated with a layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450902-cglayergetcontext?language=objc
func LayerGetContext(layer LayerRef) ContextRef {
	rv := C.LayerGetContext(layer)
	return ContextRef(rv)
}

// Returns whether an object at a given index in a PDF array is a PDF integer and, if so, retrieves that object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456053-cgpdfarraygetinteger?language=objc
func PDFArrayGetInteger(array unsafe.Pointer, index uint, value *PDFInteger) bool {
	rv := C.PDFArrayGetInteger(array, index, value)
	return bool(rv)
}

// Returns the italic angle of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396404-cgfontgetitalicangle?language=objc
func FontGetItalicAngle(font FontRef) float64 {
	rv := C.FontGetItalicAngle(font)
	return float64(rv)
}

// Configures the origin of a display relative to the global display coordinate space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454090-cgconfiguredisplayorigin?language=objc
func ConfigureDisplayOrigin(config unsafe.Pointer, display DirectDisplayID, x int32, y int32) Error {
	rv := C.ConfigureDisplayOrigin(config, C.uint32_t(display), x, y)
	return Error(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3227891-cgcolorconversioninfocreatewitho?language=objc
func ColorConversionInfoCreateWithOptions(src ColorSpaceRef, dst ColorSpaceRef, options corefoundation.DictionaryRef) ColorConversionInfoRef {
	rv := C.ColorConversionInfoCreateWithOptions(src, dst, options)
	return ColorConversionInfoRef(rv)
}

// Returns a Boolean value indicating whether a fade operation is currently in progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1571962-cgdisplayfadeoperationinprogress?language=objc
func DisplayFadeOperationInProgress() kernel.Boolean_t {
	rv := C.DisplayFadeOperationInProgress()
	return kernel.Boolean_t(rv)
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
	rv := C.BitmapContextGetAlphaInfo(context)
	return ImageAlphaInfo(rv)
}

// Closes and terminates the current path’s subpath. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454508-cgcontextclosepath?language=objc
func ContextClosePath(c ContextRef) {
	C.ContextClosePath(c)
}

// Indicates whether two graphics paths are equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411167-cgpathequaltopath?language=objc
func PathEqualToPath(path1 unsafe.Pointer, path2 unsafe.Pointer) bool {
	rv := C.PathEqualToPath(path1, path2)
	return bool(rv)
}

// Returns a new Quartz mouse event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454356-cgeventcreatemouseevent?language=objc
func EventCreateMouseEvent(source EventSourceRef, mouseType EventType, mouseCursorPosition Point, mouseButton MouseButton) EventRef {
	rv := C.EventCreateMouseEvent(source, C.uint32_t(mouseType), mouseCursorPosition, C.uint32_t(mouseButton))
	return EventRef(rv)
}

// Returns the display ID of the main display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455620-cgmaindisplayid?language=objc
func MainDisplayID() DirectDisplayID {
	rv := C.MainDisplayID()
	return DirectDisplayID(rv)
}

// Returns an affine transformation matrix constructed by translating an existing affine transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455822-cgaffinetransformtranslate?language=objc
func AffineTransformTranslate(t AffineTransform, tx float64, ty float64) AffineTransform {
	rv := C.AffineTransformTranslate(t, tx, ty)
	return AffineTransform(rv)
}

// Returns a page from a Core Graphics PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402586-cgpdfdocumentgetpage?language=objc
func PDFDocumentGetPage(document PDFDocumentRef, pageNumber uint) PDFPageRef {
	rv := C.PDFDocumentGetPage(document, pageNumber)
	return PDFPageRef(rv)
}

// Returns the x- coordinate that establishes the center of a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456175-cgrectgetmidx?language=objc
func RectGetMidX(rect Rect) float64 {
	rv := C.RectGetMidX(rect)
	return float64(rv)
}

// Returns the keyboard type to be used with a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408787-cgeventsourcegetkeyboardtype?language=objc
func EventSourceGetKeyboardType(source EventSourceRef) EventSourceKeyboardType {
	rv := C.EventSourceGetKeyboardType(source)
	return EventSourceKeyboardType(rv)
}

// Begins a new page in a PDF graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456578-cgpdfcontextbeginpage?language=objc
func PDFContextBeginPage(context ContextRef, pageInfo corefoundation.DictionaryRef) {
	C.PDFContextBeginPage(context, pageInfo)
}

// Creates a copy of an existing color, substituting a new alpha value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455986-cgcolorcreatecopywithalpha?language=objc
func ColorCreateCopyWithAlpha(color ColorRef, alpha float64) ColorRef {
	rv := C.ColorCreateCopyWithAlpha(color, alpha)
	return ColorRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2875137-cgpdfdocumentgetaccesspermission?language=objc
func PDFDocumentGetAccessPermissions(document PDFDocumentRef) PDFAccessPermissions {
	rv := C.PDFDocumentGetAccessPermissions(document)
	return PDFAccessPermissions(rv)
}

// Returns an affine transformation matrix constructed by rotating an existing affine transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455962-cgaffinetransformrotate?language=objc
func AffineTransformRotate(t AffineTransform, angle float64) AffineTransform {
	rv := C.AffineTransformRotate(t, angle)
	return AffineTransform(rv)
}

// Returns a copy of an existing Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454071-cgeventcreatecopy?language=objc
func EventCreateCopy(event EventRef) EventRef {
	rv := C.EventCreateCopy(event)
	return EventRef(rv)
}

// Marks a window context for update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455450-cgcontextsynchronize?language=objc
func ContextSynchronize(c ContextRef) {
	C.ContextSynchronize(c)
}

// Returns the affine transform that maps a box to a given rectangle on a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454893-cgpdfpagegetdrawingtransform?language=objc
func PDFPageGetDrawingTransform(page PDFPageRef, box PDFBox, rect Rect, rotate int, preserveAspectRatio bool) AffineTransform {
	rv := C.PDFPageGetDrawingTransform(page, C.int32_t(box), rect, rotate, preserveAspectRatio)
	return AffineTransform(rv)
}

// Returns the window level of the shield window for a captured display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456352-cgshieldingwindowlevel?language=objc
func ShieldingWindowLevel() WindowLevel {
	rv := C.ShieldingWindowLevel()
	return WindowLevel(rv)
}

// Returns a Boolean value indicating whether the specified display mode is usable for a desktop graphical user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454928-cgdisplaymodeisusablefordesktopg?language=objc
func DisplayModeIsUsableForDesktopGUI(mode DisplayModeRef) bool {
	rv := C.DisplayModeIsUsableForDesktopGUI(mode)
	return bool(rv)
}

// Returns the Core Foundation type identifier for CGGradient objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1398453-cggradientgettypeid?language=objc
func GradientGetTypeID() corefoundation.TypeID {
	rv := C.GradientGetTypeID()
	return corefoundation.TypeID(rv)
}

// Returns a dictionary representation of the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455382-cgpointcreatedictionaryrepresent?language=objc
func PointCreateDictionaryRepresentation(point Point) corefoundation.DictionaryRef {
	rv := C.PointCreateDictionaryRepresentation(point)
	return corefoundation.DictionaryRef(rv)
}

// Returns a rectangle with an origin that is offset from that of the source rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454841-cgrectoffset?language=objc
func RectOffset(rect Rect, dx float64, dy float64) Rect {
	rv := C.RectOffset(rect, dx, dy)
	return Rect(rv)
}

// Returns whether an object at a given index in a PDF array is a PDF name reference (represented as a constant C string) and, if so, retrieves that name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455034-cgpdfarraygetname?language=objc
func PDFArrayGetName(array unsafe.Pointer, index uint, value *uint8) bool {
	rv := C.PDFArrayGetName(array, index, value)
	return bool(rv)
}

// Sets the current stroke color to a value in the DeviceCMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455358-cgcontextsetcmykstrokecolor?language=objc
func ContextSetCMYKStrokeColor(c ContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64) {
	C.ContextSetCMYKStrokeColor(c, cyan, magenta, yellow, black, alpha)
}

// Returns the base color space of a pattern or indexed color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408839-cgcolorspacegetbasecolorspace?language=objc
func ColorSpaceGetBaseColorSpace(space ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceGetBaseColorSpace(space)
	return ColorSpaceRef(rv)
}

// Creates a layer object that is associated with a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450892-cglayercreatewithcontext?language=objc
func LayerCreateWithContext(context ContextRef, size Size, auxiliaryInfo corefoundation.DictionaryRef) LayerRef {
	rv := C.LayerCreateWithContext(context, size, auxiliaryInfo)
	return LayerRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454756-cgdisplaymodegetpixelwidth?language=objc
func DisplayModeGetPixelWidth(mode DisplayModeRef) uint {
	rv := C.DisplayModeGetPixelWidth(mode)
	return uint(rv)
}

// Creates a bitmap image by masking an existing bitmap image with the provided color values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454358-cgimagecreatewithmaskingcolors?language=objc
func ImageCreateWithMaskingColors(image ImageRef, components *float64) ImageRef {
	rv := C.ImageCreateWithMaskingColors(image, components)
	return ImageRef(rv)
}

// Paints the areas contained within the provided rectangles, using the fill color in the current graphics state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454132-cgcontextfillrects?language=objc
func ContextFillRects(c ContextRef, rects *Rect, count uint) {
	C.ContextFillRects(c, rects, count)
}

// Obtains the PostScript name of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396346-cgfontcopypostscriptname?language=objc
func FontCopyPostScriptName(font FontRef) corefoundation.StringRef {
	rv := C.FontCopyPostScriptName(font)
	return corefoundation.StringRef(rv)
}

// Returns whether a bitmap image is an image mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454229-cgimageismask?language=objc
func ImageIsMask(image ImageRef) bool {
	rv := C.ImageIsMask(image)
	return bool(rv)
}

// Sets the stroke color space in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454396-cgcontextsetstrokecolorspace?language=objc
func ContextSetStrokeColorSpace(c ContextRef, space ColorSpaceRef) {
	C.ContextSetStrokeColorSpace(c, space)
}

// Returns a rectangle with a positive width and height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456432-cgrectstandardize?language=objc
func RectStandardize(rect Rect) Rect {
	rv := C.RectStandardize(rect)
	return Rect(rv)
}

// Returns the bounding box of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396353-cgfontgetfontbbox?language=objc
func FontGetFontBBox(font FontRef) Rect {
	rv := C.FontGetFontBBox(font)
	return Rect(rv)
}

// Returns whether an object with a specified key in a PDF dictionary is a PDF name reference (represented as a constant C string) and, if so, retrieves that name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430230-cgpdfdictionarygetname?language=objc
func PDFDictionaryGetName(dict unsafe.Pointer, key *uint8, value *uint8) bool {
	rv := C.PDFDictionaryGetName(dict, key, value)
	return bool(rv)
}

// Retrieves an object from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455971-cgpdfscannerpopobject?language=objc
func PDFScannerPopObject(scanner unsafe.Pointer, value unsafe.Pointer) bool {
	rv := C.PDFScannerPopObject(scanner, value)
	return bool(rv)
}

// Gets the run loop source for a display stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455403-cgdisplaystreamgetrunloopsource?language=objc
func DisplayStreamGetRunLoopSource(displayStream DisplayStreamRef) corefoundation.RunLoopSourceRef {
	rv := C.DisplayStreamGetRunLoopSource(displayStream)
	return corefoundation.RunLoopSourceRef(rv)
}

// Obtains exclusive use of a display, preventing other applications and system services from using the display or changing its configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456259-cgdisplaycapture?language=objc
func DisplayCapture(display DirectDisplayID) Error {
	rv := C.DisplayCapture(C.uint32_t(display))
	return Error(rv)
}

// Adds an ellipse that fits inside the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456420-cgcontextaddellipseinrect?language=objc
func ContextAddEllipseInRect(c ContextRef, rect Rect) {
	C.ContextAddEllipseInRect(c, rect)
}

// Creates a CGGradient object from a color space and the provided color components and locations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1398454-cggradientcreatewithcolorcompone?language=objc
func GradientCreateWithColorComponents(space ColorSpaceRef, components *float64, locations *float64, count uint) GradientRef {
	rv := C.GradientCreateWithColorComponents(space, components, locations, count)
	return GradientRef(rv)
}

// Returns whether two rectangles are equal in size and position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456516-cgrectequaltorect?language=objc
func RectEqualToRect(rect1 Rect, rect2 Rect) bool {
	rv := C.RectEqualToRect(rect1, rect2)
	return bool(rv)
}

// Creates a calibrated RGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408861-cgcolorspacecreatecalibratedrgb?language=objc
func ColorSpaceCreateCalibratedRGB(whitePoint *float64, blackPoint *float64, gamma *float64, matrix *float64) ColorSpaceRef {
	rv := C.ColorSpaceCreateCalibratedRGB(whitePoint, blackPoint, gamma, matrix)
	return ColorSpaceRef(rv)
}

// Fills the clipping path of a context with the specified shading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456643-cgcontextdrawshading?language=objc
func ContextDrawShading(c ContextRef, shading ShadingRef) {
	C.ContextDrawShading(c, shading)
}

// Returns the y-coordinate that establishes the center of the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456550-cgrectgetmidy?language=objc
func RectGetMidY(rect Rect) float64 {
	rv := C.RectGetMidY(rect)
	return float64(rv)
}

// Returns an image containing the contents of a portion of the specified display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454595-cgdisplaycreateimageforrect?language=objc
func DisplayCreateImageForRect(display DirectDisplayID, rect Rect) ImageRef {
	rv := C.DisplayCreateImageForRect(C.uint32_t(display), rect)
	return ImageRef(rv)
}

// Creates a color using a list of intensity values (including alpha) and an associated color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455927-cgcolorcreate?language=objc
func ColorCreate(space ColorSpaceRef, components *float64) ColorRef {
	rv := C.ColorCreate(space, components)
	return ColorRef(rv)
}

// Starts a new subpath at a specified location in a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411146-cgpathmovetopoint?language=objc
func PathMoveToPoint(path MutablePathRef, m *AffineTransform, x float64, y float64) {
	C.PathMoveToPoint(path, m, x, y)
}

// Decrements the retain count of a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586340-cgcolorrelease?language=objc
func ColorRelease(color ColorRef) {
	C.ColorRelease(color)
}

// Returns the number of pages in a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402595-cgpdfdocumentgetnumberofpages?language=objc
func PDFDocumentGetNumberOfPages(document PDFDocumentRef) uint {
	rv := C.PDFDocumentGetNumberOfPages(document)
	return uint(rv)
}

// Sets the keyboard type to be used with a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408795-cgeventsourcesetkeyboardtype?language=objc
func EventSourceSetKeyboardType(source EventSourceRef, keyboardType EventSourceKeyboardType) {
	C.EventSourceSetKeyboardType(source, C.uint32_t(keyboardType))
}

// Returns information about the currently available display modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455537-cgdisplaycopyalldisplaymodes?language=objc
func DisplayCopyAllDisplayModes(display DirectDisplayID, options corefoundation.DictionaryRef) corefoundation.ArrayRef {
	rv := C.DisplayCopyAllDisplayModes(C.uint32_t(display), options)
	return corefoundation.ArrayRef(rv)
}

// Returns the display height in pixel units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454247-cgdisplaypixelshigh?language=objc
func DisplayPixelsHigh(display DirectDisplayID) uint {
	rv := C.DisplayPixelsHigh(C.uint32_t(display))
	return uint(rv)
}

// Return the movement delta values for a single update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455469-cgdisplaystreamupdategetmovedrec?language=objc
func DisplayStreamUpdateGetMovedRectsDelta(updateRef DisplayStreamUpdateRef, dx *float64, dy *float64) {
	C.DisplayStreamUpdateGetMovedRectsDelta(updateRef, dx, dy)
}

// Sets the byte values in the 8-bit RGB gamma tables for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455896-cgsetdisplaytransferbybytetable?language=objc
func SetDisplayTransferByByteTable(display DirectDisplayID, tableSize uint32, redTable *uint8, greenTable *uint8, blueTable *uint8) Error {
	rv := C.SetDisplayTransferByByteTable(C.uint32_t(display), tableSize, redTable, greenTable, blueTable)
	return Error(rv)
}

// Returns the data provider for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455260-cgimagegetdataprovider?language=objc
func ImageGetDataProvider(image ImageRef) DataProviderRef {
	rv := C.ImageGetDataProvider(image)
	return DataProviderRef(rv)
}

// Returns a Boolean value indicating whether a display is captured. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1562061-cgdisplayiscaptured?language=objc
func DisplayIsCaptured(display DirectDisplayID) kernel.Boolean_t {
	rv := C.DisplayIsCaptured(C.uint32_t(display))
	return kernel.Boolean_t(rv)
}

// Sets the fill pattern in the specified graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456334-cgcontextsetfillpattern?language=objc
func ContextSetFillPattern(c ContextRef, pattern PatternRef, components *float64) {
	C.ContextSetFillPattern(c, pattern, components)
}

// Waits for screen refresh operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541809-cgwaitforscreenrefreshrects?language=objc
func WaitForScreenRefreshRects(rects *Rect, count *uint32) Error {
	rv := C.WaitForScreenRefreshRects(rects, count)
	return Error(rv)
}

// Returns the Core Foundation type identifier for data providers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408290-cgdataprovidergettypeid?language=objc
func DataProviderGetTypeID() corefoundation.TypeID {
	rv := C.DataProviderGetTypeID()
	return corefoundation.TypeID(rv)
}

// Creates a direct-access data provider that uses a file to supply data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408294-cgdataprovidercreatewithfilename?language=objc
func DataProviderCreateWithFilename(filename *uint8) DataProviderRef {
	rv := C.DataProviderCreateWithFilename(filename)
	return DataProviderRef(rv)
}

// Retrieves a string object from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455018-cgpdfscannerpopstring?language=objc
func PDFScannerPopString(scanner unsafe.Pointer, value unsafe.Pointer) bool {
	rv := C.PDFScannerPopString(scanner, value)
	return bool(rv)
}

// Returns the number of bytes allocated for a single row of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455425-cgimagegetbytesperrow?language=objc
func ImageGetBytesPerRow(image ImageRef) uint {
	rv := C.ImageGetBytesPerRow(image)
	return uint(rv)
}

// Creates a data provider that reads from a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408284-cgdataprovidercreatewithcfdata?language=objc
func DataProviderCreateWithCFData(data corefoundation.DataRef) DataProviderRef {
	rv := C.DataProviderCreateWithCFData(data)
	return DataProviderRef(rv)
}

// Returns whether the specified PDF document allows copying. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402588-cgpdfdocumentallowscopying?language=objc
func PDFDocumentAllowsCopying(document PDFDocumentRef) bool {
	rv := C.PDFDocumentAllowsCopying(document)
	return bool(rv)
}

// Creates a new color in a different color space that matches the provided color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455493-cgcolorcreatecopybymatchingtocol?language=objc
func ColorCreateCopyByMatchingToColorSpace(arg0 ColorSpaceRef, intent ColorRenderingIntent, color ColorRef, options corefoundation.DictionaryRef) ColorRef {
	rv := C.ColorCreateCopyByMatchingToColorSpace(arg0, C.int32_t(intent), color, options)
	return ColorRef(rv)
}

// Turns off local hardware events in the current session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541773-cginhibitlocalevents?language=objc
func InhibitLocalEvents(inhibit kernel.Boolean_t) Error {
	rv := C.InhibitLocalEvents(C.NSInteger(inhibit))
	return Error(rv)
}

// Returns a size with the specified dimension values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455082-cgsizemake?language=objc
func SizeMake(width float64, height float64) Size {
	rv := C.SizeMake(width, height)
	return Size(rv)
}

// Returns the largest value for the y-coordinate of the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454060-cgrectgetmaxy?language=objc
func RectGetMaxY(rect Rect) float64 {
	rv := C.RectGetMaxY(rect)
	return float64(rv)
}

// Returns whether a PDF document allows printing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402594-cgpdfdocumentallowsprinting?language=objc
func PDFDocumentAllowsPrinting(document PDFDocumentRef) bool {
	rv := C.PDFDocumentAllowsPrinting(document)
	return bool(rv)
}

// Sets whether or not to allow font smoothing for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454767-cgcontextsetallowsfontsmoothing?language=objc
func ContextSetAllowsFontSmoothing(c ContextRef, allowsFontSmoothing bool) {
	C.ContextSetAllowsFontSmoothing(c, allowsFontSmoothing)
}

// Draws the current path using the provided drawing mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455195-cgcontextdrawpath?language=objc
func ContextDrawPath(c ContextRef, mode PathDrawingMode) {
	C.ContextDrawPath(c, C.int32_t(mode))
}

// Returns the glyph for the glyph name associated with the specified font object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396340-cgfontgetglyphwithglyphname?language=objc
func FontGetGlyphWithGlyphName(font FontRef, name corefoundation.StringRef) Glyph {
	rv := C.FontGetGlyphWithGlyphName(font, name)
	return Glyph(rv)
}

// Retrieves an array object from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454360-cgpdfscannerpoparray?language=objc
func PDFScannerPopArray(scanner unsafe.Pointer, value unsafe.Pointer) bool {
	rv := C.PDFScannerPopArray(scanner, value)
	return bool(rv)
}

// Returns a dictionary representation of the provided rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455760-cgrectcreatedictionaryrepresenta?language=objc
func RectCreateDictionaryRepresentation(arg0 Rect) corefoundation.DictionaryRef {
	rv := C.RectCreateDictionaryRepresentation(arg0)
	return corefoundation.DictionaryRef(rv)
}

// Returns a pointer to the image data associated with a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455517-cgbitmapcontextgetdata?language=objc
func BitmapContextGetData(context ContextRef) unsafe.Pointer {
	rv := C.BitmapContextGetData(context)
	return unsafe.Pointer(rv)
}

// Returns the ascent of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396359-cgfontgetascent?language=objc
func FontGetAscent(font FontRef) int {
	rv := C.FontGetAscent(font)
	return int(rv)
}

// Returns the pixel encoding of the specified display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455067-cgdisplaymodecopypixelencoding?language=objc
func DisplayModeCopyPixelEncoding(mode DisplayModeRef) corefoundation.StringRef {
	rv := C.DisplayModeCopyPixelEncoding(mode)
	return corefoundation.StringRef(rv)
}

// Immediately enables or disables stereo operation for a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454316-cgdisplaysetstereooperation?language=objc
func DisplaySetStereoOperation(display DirectDisplayID, stereo kernel.Boolean_t, forceBlueLine kernel.Boolean_t, option ConfigureOption) Error {
	rv := C.DisplaySetStereoOperation(C.uint32_t(display), C.NSInteger(stereo), C.NSInteger(forceBlueLine), C.uint32_t(option))
	return Error(rv)
}

// Returns the interpolation setting for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455363-cgimagegetshouldinterpolate?language=objc
func ImageGetShouldInterpolate(image ImageRef) bool {
	rv := C.ImageGetShouldInterpolate(image)
	return bool(rv)
}

// Paints the area within the current path, using the nonzero winding number rule. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456306-cgcontextfillpath?language=objc
func ContextFillPath(c ContextRef) {
	C.ContextFillPath(c)
}

// Returns the type identifier for Core Graphics PDF documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402597-cgpdfdocumentgettypeid?language=objc
func PDFDocumentGetTypeID() corefoundation.TypeID {
	rv := C.PDFDocumentGetTypeID()
	return corefoundation.TypeID(rv)
}

// Returns whether an object at a given index in a PDF array is a PDF string and, if so, retrieves that string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456104-cgpdfarraygetstring?language=objc
func PDFArrayGetString(array unsafe.Pointer, index uint, value unsafe.Pointer) bool {
	rv := C.PDFArrayGetString(array, index, value)
	return bool(rv)
}

// Displays a character array at the current text position, a point specified by the current text matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586507-cgcontextshowtext?language=objc
func ContextShowText(c ContextRef, string *uint8, length uint) {
	C.ContextShowText(c, string_, length)
}

// Returns the rendering intent setting for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456350-cgimagegetrenderingintent?language=objc
func ImageGetRenderingIntent(image ImageRef) ColorRenderingIntent {
	rv := C.ImageGetRenderingIntent(image)
	return ColorRenderingIntent(rv)
}

// Posts a Quartz event into the event stream at a specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456527-cgeventpost?language=objc
func EventPost(tap EventTapLocation, event EventRef) {
	C.EventPost(C.uint32_t(tap), event)
}

// Returns the thickness of the dominant vertical stems of glyphs in a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396380-cgfontgetstemv?language=objc
func FontGetStemV(font FontRef) float64 {
	rv := C.FontGetStemV(font)
	return float64(rv)
}

// Returns a Boolean value indicating whether the mouse cursor is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541812-cgcursorisvisible?language=objc
func CursorIsVisible() kernel.Boolean_t {
	rv := C.CursorIsVisible()
	return kernel.Boolean_t(rv)
}

// Strokes a sequence of line segments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454389-cgcontextstrokelinesegments?language=objc
func ContextStrokeLineSegments(c ContextRef, points *Point, count uint) {
	C.ContextStrokeLineSegments(c, points, count)
}

// Paints the area within the current path, using the even-odd fill rule. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454865-cgcontexteofillpath?language=objc
func ContextEOFillPath(c ContextRef) {
	C.ContextEOFillPath(c)
}

// Returns whether an object at a given index in a PDF array is a PDF object and, if so, retrieves that object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456631-cgpdfarraygetobject?language=objc
func PDFArrayGetObject(array unsafe.Pointer, index uint, value unsafe.Pointer) bool {
	rv := C.PDFArrayGetObject(array, index, value)
	return bool(rv)
}

// Obtains exclusive use of a display for an application using the options you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454934-cgdisplaycapturewithoptions?language=objc
func DisplayCaptureWithOptions(display DirectDisplayID, options CaptureOptions) Error {
	rv := C.DisplayCaptureWithOptions(C.uint32_t(display), C.uint32_t(options))
	return Error(rv)
}

// Checks to see whether the specified point is contained in the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454778-cgcontextpathcontainspoint?language=objc
func ContextPathContainsPoint(c ContextRef, point Point, mode PathDrawingMode) bool {
	rv := C.ContextPathContainsPoint(c, point, C.int32_t(mode))
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3684559-cgcolorspaceusesitur_2100tf?language=objc
func ColorSpaceUsesITUR_2100TF(arg0 ColorSpaceRef) bool {
	rv := C.ColorSpaceUsesITUR_2100TF(arg0)
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2875135-cgpdfdocumentgetoutline?language=objc
func PDFDocumentGetOutline(document PDFDocumentRef) corefoundation.DictionaryRef {
	rv := C.PDFDocumentGetOutline(document)
	return corefoundation.DictionaryRef(rv)
}

// Tells a stream to stop sending updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455658-cgdisplaystreamstop?language=objc
func DisplayStreamStop(displayStream DisplayStreamRef) Error {
	rv := C.DisplayStreamStop(displayStream)
	return Error(rv)
}

// Decrements the retain count of a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411148-cgpathrelease?language=objc
func PathRelease(path unsafe.Pointer) {
	C.PathRelease(path)
}

// Returns the Core Foundation type identifier for Core Graphics paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411192-cgpathgettypeid?language=objc
func PathGetTypeID() corefoundation.TypeID {
	rv := C.PathGetTypeID()
	return corefoundation.TypeID(rv)
}

// Decrements the retain count of a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402593-cgpdfdocumentrelease?language=objc
func PDFDocumentRelease(document PDFDocumentRef) {
	C.PDFDocumentRelease(document)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454440-cgwindowservercreateserverport?language=objc
func WindowServerCreateServerPort() corefoundation.MachPortRef {
	rv := C.WindowServerCreateServerPort()
	return corefoundation.MachPortRef(rv)
}

// Returns an affine transform that maps user space coordinates to device space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455677-cgcontextgetuserspacetodevicespa?language=objc
func ContextGetUserSpaceToDeviceSpaceTransform(c ContextRef) AffineTransform {
	rv := C.ContextGetUserSpaceToDeviceSpaceTransform(c)
	return AffineTransform(rv)
}

// Draws an image into a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454845-cgcontextdrawimage?language=objc
func ContextDrawImage(c ContextRef, rect Rect, image ImageRef) {
	C.ContextDrawImage(c, rect, image)
}

// Create an immutable path of a rounded rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411218-cgpathcreatewithroundedrect?language=objc
func PathCreateWithRoundedRect(rect Rect, cornerWidth float64, cornerHeight float64, transform *AffineTransform) unsafe.Pointer {
	rv := C.PathCreateWithRoundedRect(rect, cornerWidth, cornerHeight, transform)
	return unsafe.Pointer(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3042356-cgpdfcontextbegintag?language=objc
func PDFContextBeginTag(context ContextRef, tagType PDFTagType, tagProperties corefoundation.DictionaryRef) {
	C.PDFContextBeginTag(context, C.int32_t(tagType), tagProperties)
}

// Returns the PDF type identifier of an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455189-cgpdfobjectgettype?language=objc
func PDFObjectGetType(object unsafe.Pointer) PDFObjectType {
	rv := C.PDFObjectGetType(object)
	return PDFObjectType(rv)
}

// Creates a bitmap image mask from data supplied by a data provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455089-cgimagemaskcreate?language=objc
func ImageMaskCreate(width uint, height uint, bitsPerComponent uint, bitsPerPixel uint, bytesPerRow uint, provider DataProviderRef, decode *float64, shouldInterpolate bool) ImageRef {
	rv := C.ImageMaskCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow, provider, decode, shouldInterpolate)
	return ImageRef(rv)
}

// Creates a color in the Generic RGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455631-cgcolorcreategenericrgb?language=objc
func ColorCreateGenericRGB(red float64, green float64, blue float64, alpha float64) ColorRef {
	rv := C.ColorCreateGenericRGB(red, green, blue, alpha)
	return ColorRef(rv)
}

// Returns the floating-point value of a field in a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455506-cgeventgetdoublevaluefield?language=objc
func EventGetDoubleValueField(event EventRef, field EventField) float64 {
	rv := C.EventGetDoubleValueField(event, C.uint32_t(field))
	return float64(rv)
}

// Gets the information dictionary for a PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402589-cgpdfdocumentgetinfo?language=objc
func PDFDocumentGetInfo(document PDFDocumentRef) unsafe.Pointer {
	rv := C.PDFDocumentGetInfo(document)
	return unsafe.Pointer(rv)
}

// Checks whether the converter is currently converting data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454582-cgpsconverterisconverting?language=objc
func PSConverterIsConverting(converter PSConverterRef) bool {
	rv := C.PSConverterIsConverting(converter)
	return bool(rv)
}

// Creates a copy of a font using a variation specification dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396373-cgfontcreatecopywithvariations?language=objc
func FontCreateCopyWithVariations(font FontRef, variations corefoundation.DictionaryRef) FontRef {
	rv := C.FontCreateCopyWithVariations(font, variations)
	return FontRef(rv)
}

// Changes the configuration of a mirroring set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454531-cgconfiguredisplaymirrorofdispla?language=objc
func ConfigureDisplayMirrorOfDisplay(config unsafe.Pointer, display DirectDisplayID, master DirectDisplayID) Error {
	rv := C.ConfigureDisplayMirrorOfDisplay(config, C.uint32_t(display), C.uint32_t(master))
	return Error(rv)
}

// Returns the refresh rate of the specified display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454661-cgdisplaymodegetrefreshrate?language=objc
func DisplayModeGetRefreshRate(mode DisplayModeRef) float64 {
	rv := C.DisplayModeGetRefreshRate(mode)
	return float64(rv)
}

// Returns the interval that local hardware events may be suppressed following the posting of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408774-cgeventsourcegetlocaleventssuppr?language=objc
func EventSourceGetLocalEventsSuppressionInterval(source EventSourceRef) corefoundation.TimeInterval {
	rv := C.EventSourceGetLocalEventsSuppressionInterval(source)
	return corefoundation.TimeInterval(rv)
}

// Adds an arc of a circle to the current path, using a radius and tangent points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456238-cgcontextaddarctopoint?language=objc
func ContextAddArcToPoint(c ContextRef, x1 float64, y1 float64, x2 float64, y2 float64, radius float64) {
	C.ContextAddArcToPoint(c, x1, y1, x2, y2, radius)
}

// Creates a data consumer that writes data to a location specified by a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454474-cgdataconsumercreatewithurl?language=objc
func DataConsumerCreateWithURL(url corefoundation.URLRef) DataConsumerRef {
	rv := C.DataConsumerCreateWithURL(url)
	return DataConsumerRef(rv)
}

// Returns the width and height of a layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450890-cglayergetsize?language=objc
func LayerGetSize(layer LayerRef) Size {
	rv := C.LayerGetSize(layer)
	return Size(rv)
}

// Increments the retain count of a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586506-cgcontextretain?language=objc
func ContextRetain(c ContextRef) ContextRef {
	rv := C.ContextRetain(c)
	return ContextRef(rv)
}

// Reserves the fade hardware for a specified time interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456391-cgacquiredisplayfadereservation?language=objc
func AcquireDisplayFadeReservation(seconds DisplayReservationInterval, token *DisplayFadeReservationToken) Error {
	rv := C.AcquireDisplayFadeReservation(C.float(seconds), token)
	return Error(rv)
}

// Returns the event flags of a Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455642-cgeventgetflags?language=objc
func EventGetFlags(event EventRef) EventFlags {
	rv := C.EventGetFlags(event)
	return EventFlags(rv)
}

// Returns the glyph name of the specified glyph in the specified font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396349-cgfontcopyglyphnameforglyph?language=objc
func FontCopyGlyphNameForGlyph(font FontRef, glyph Glyph) corefoundation.StringRef {
	rv := C.FontCopyGlyphNameForGlyph(font, C.NSInteger(glyph))
	return corefoundation.StringRef(rv)
}

// Sets the current fill color to a value in the DeviceRGB color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455624-cgcontextsetrgbfillcolor?language=objc
func ContextSetRGBFillColor(c ContextRef, red float64, green float64, blue float64, alpha float64) {
	C.ContextSetRGBFillColor(c, red, green, blue, alpha)
}

// Sets the current graphics state to the state most recently saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455391-cgcontextrestoregstate?language=objc
func ContextRestoreGState(c ContextRef) {
	C.ContextRestoreGState(c)
}

// Creates an immutable copy of a graphics path transformed by a transformation matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411161-cgpathcreatecopybytransformingpa?language=objc
func PathCreateCopyByTransformingPath(path unsafe.Pointer, transform *AffineTransform) unsafe.Pointer {
	rv := C.PathCreateCopyByTransformingPath(path, transform)
	return unsafe.Pointer(rv)
}

// Returns the full name associated with a font object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396357-cgfontcopyfullname?language=objc
func FontCopyFullName(font FontRef) corefoundation.StringRef {
	rv := C.FontCopyFullName(font)
	return corefoundation.StringRef(rv)
}

// Returns an affine transformation matrix constructed by combining two existing affine transforms. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455996-cgaffinetransformconcat?language=objc
func AffineTransformConcat(t1 AffineTransform, t2 AffineTransform) AffineTransform {
	rv := C.AffineTransformConcat(t1, t2)
	return AffineTransform(rv)
}

// Returns the page number of the specified PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454587-cgpdfpagegetpagenumber?language=objc
func PDFPageGetPageNumber(page PDFPageRef) uint {
	rv := C.PDFPageGetPageNumber(page)
	return uint(rv)
}

// Sets the location of a Quartz mouse event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456389-cgeventsetlocation?language=objc
func EventSetLocation(event EventRef, location Point) {
	C.EventSetLocation(event, location)
}

// Decrements the retain count of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396363-cgfontrelease?language=objc
func FontRelease(font FontRef) {
	C.FontRelease(font)
}

// Returns a CFString object that represents a PDF string as a text string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456234-cgpdfstringcopytextstring?language=objc
func PDFStringCopyTextString(string unsafe.Pointer) corefoundation.StringRef {
	rv := C.PDFStringCopyTextString(string_)
	return corefoundation.StringRef(rv)
}

// Returns whether the specified PDF file is encrypted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402591-cgpdfdocumentisencrypted?language=objc
func PDFDocumentIsEncrypted(document PDFDocumentRef) bool {
	rv := C.PDFDocumentIsEncrypted(document)
	return bool(rv)
}

// Paints a rectangular path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454675-cgcontextstrokerect?language=objc
func ContextStrokeRect(c ContextRef, rect Rect) {
	C.ContextStrokeRect(c, rect)
}

// Enables or disables subpixel positioning in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455671-cgcontextsetshouldsubpixelpositi?language=objc
func ContextSetShouldSubpixelPositionFonts(c ContextRef, shouldSubpixelPositionFonts bool) {
	C.ContextSetShouldSubpixelPositionFonts(c, shouldSubpixelPositionFonts)
}

// Returns a path object built from the current path information in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455397-cgcontextcopypath?language=objc
func ContextCopyPath(c ContextRef) unsafe.Pointer {
	rv := C.ContextCopyPath(c)
	return unsafe.Pointer(rv)
}

// The Universal Type Identifier for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456067-cgimagegetuttype?language=objc
func ImageGetUTType(image ImageRef) corefoundation.StringRef {
	rv := C.ImageGetUTType(image)
	return corefoundation.StringRef(rv)
}

// Determines whether Core Graphics can create a subset of the font in PostScript format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396365-cgfontcancreatepostscriptsubset?language=objc
func FontCanCreatePostScriptSubset(font FontRef, format FontPostScriptFormat) bool {
	rv := C.FontCanCreatePostScriptSubset(font, C.int32_t(format))
	return bool(rv)
}

// Returns the values of the color components (including alpha) associated with a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455930-cgcolorgetcomponents?language=objc
func ColorGetComponents(color ColorRef) *float64 {
	rv := C.ColorGetComponents(color)
	return *float64(rv)
}

// Returns the type identifier for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455504-cgcontextgettypeid?language=objc
func ContextGetTypeID() corefoundation.TypeID {
	rv := C.ContextGetTypeID()
	return corefoundation.TypeID(rv)
}

// Creates a data consumer that writes to a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456292-cgdataconsumercreatewithcfdata?language=objc
func DataConsumerCreateWithCFData(data unsafe.Pointer) DataConsumerRef {
	rv := C.DataConsumerCreateWithCFData(data)
	return DataConsumerRef(rv)
}

// Adds a rectangular path to the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456617-cgcontextaddrect?language=objc
func ContextAddRect(c ContextRef, rect Rect) {
	C.ContextAddRect(c, rect)
}

// Appends an arc to a mutable graphics path, possibly preceded by a straight line segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411136-cgpathaddrelativearc?language=objc
func PathAddRelativeArc(path MutablePathRef, matrix *AffineTransform, x float64, y float64, radius float64, startAngle float64, delta float64) {
	C.PathAddRelativeArc(path, matrix, x, y, radius, startAngle, delta)
}

// Decrements the retain count of a data consumer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1508424-cgdataconsumerrelease?language=objc
func DataConsumerRelease(consumer DataConsumerRef) {
	C.DataConsumerRelease(consumer)
}

// Enables or disables stereo operation for a display, as part of a display configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456308-cgconfiguredisplaystereooperatio?language=objc
func ConfigureDisplayStereoOperation(config unsafe.Pointer, display DirectDisplayID, stereo kernel.Boolean_t, forceBlueLine kernel.Boolean_t) Error {
	rv := C.ConfigureDisplayStereoOperation(config, C.uint32_t(display), C.NSInteger(stereo), C.NSInteger(forceBlueLine))
	return Error(rv)
}

// Returns a size that is transformed from user space coordinates to device space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456619-cgcontextconvertsizetodevicespac?language=objc
func ContextConvertSizeToDeviceSpace(c ContextRef, size Size) Size {
	rv := C.ContextConvertSizeToDeviceSpace(c, size)
	return Size(rv)
}

// Appends a cubic Bézier curve to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411212-cgpathaddcurvetopoint?language=objc
func PathAddCurveToPoint(path MutablePathRef, m *AffineTransform, cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {
	C.PathAddCurveToPoint(path, m, cp1x, cp1y, cp2x, cp2y, x, y)
}

// Returns a Boolean value indicating whether a display is built-in, such as the internal display in portable systems. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454566-cgdisplayisbuiltin?language=objc
func DisplayIsBuiltin(display DirectDisplayID) kernel.Boolean_t {
	rv := C.DisplayIsBuiltin(C.uint32_t(display))
	return kernel.Boolean_t(rv)
}

// Hides the mouse cursor, and increments the hide cursor count. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456534-cgdisplayhidecursor?language=objc
func DisplayHideCursor(display DirectDisplayID) Error {
	rv := C.DisplayHideCursor(C.uint32_t(display))
	return Error(rv)
}

// Returns the logical unit number of a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456489-cgdisplayunitnumber?language=objc
func DisplayUnitNumber(display DirectDisplayID) uint32 {
	rv := C.DisplayUnitNumber(C.uint32_t(display))
	return uint32(rv)
}

// Switches a display to a different mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454760-cgdisplaysetdisplaymode?language=objc
func DisplaySetDisplayMode(display DirectDisplayID, mode DisplayModeRef, options corefoundation.DictionaryRef) Error {
	rv := C.DisplaySetDisplayMode(C.uint32_t(display), mode, options)
	return Error(rv)
}

// Draws glyphs at the provided position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456200-cgcontextshowglyphsatpositions?language=objc
func ContextShowGlyphsAtPositions(c ContextRef, glyphs *Glyph, Lpositions *Point, count uint) {
	C.ContextShowGlyphsAtPositions(c, glyphs, Lpositions, count)
}

// Sets the current fill color in a graphics context, using a CGColor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454079-cgcontextsetfillcolorwithcolor?language=objc
func ContextSetFillColorWithColor(c ContextRef, color ColorRef) {
	C.ContextSetFillColorWithColor(c, color)
}

// Returns the leading of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396390-cgfontgetleading?language=objc
func FontGetLeading(font FontRef) int {
	rv := C.FontGetLeading(font)
	return int(rv)
}

// Returns the Core Foundation type identifier for PostScript converters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454188-cgpsconvertergettypeid?language=objc
func PSConverterGetTypeID() corefoundation.TypeID {
	rv := C.PSConverterGetTypeID()
	return corefoundation.TypeID(rv)
}

// Creates a gradient object from a color space and the provided color objects and locations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1398458-cggradientcreatewithcolors?language=objc
func GradientCreateWithColors(space ColorSpaceRef, colors corefoundation.ArrayRef, locations *float64) GradientRef {
	rv := C.GradientCreateWithColors(space, colors, locations)
	return GradientRef(rv)
}

// Returns a Boolean value indicating whether a display is always in a mirroring set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456110-cgdisplayisalwaysinmirrorset?language=objc
func DisplayIsAlwaysInMirrorSet(display DirectDisplayID) kernel.Boolean_t {
	rv := C.DisplayIsAlwaysInMirrorSet(C.uint32_t(display))
	return kernel.Boolean_t(rv)
}

// Creates a dashed copy of another path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411134-cgpathcreatecopybydashingpath?language=objc
func PathCreateCopyByDashingPath(path unsafe.Pointer, transform *AffineTransform, phase float64, lengths *float64, count uint) unsafe.Pointer {
	rv := C.PathCreateCopyByDashingPath(path, transform, phase, lengths, count)
	return unsafe.Pointer(rv)
}

// Sets the current fill color to a value in the DeviceCMYK color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454214-cgcontextsetcmykfillcolor?language=objc
func ContextSetCMYKFillColor(c ContextRef, cyan float64, magenta float64, yellow float64, black float64, alpha float64) {
	C.ContextSetCMYKFillColor(c, cyan, magenta, yellow, black, alpha)
}

// Cancels a set of display configuration changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455522-cgcanceldisplayconfiguration?language=objc
func CancelDisplayConfiguration(config unsafe.Pointer) Error {
	rv := C.CancelDisplayConfiguration(config)
	return Error(rv)
}

// Appends an array of new line segments to a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411171-cgpathaddlines?language=objc
func PathAddLines(path MutablePathRef, m *AffineTransform, points *Point, count uint) {
	C.PathAddLines(path, m, points, count)
}

// Sets the gamma function for a display by specifying the coefficients of the gamma transfer formula. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454126-cgsetdisplaytransferbyformula?language=objc
func SetDisplayTransferByFormula(display DirectDisplayID, redMin GammaValue, redMax GammaValue, redGamma GammaValue, greenMin GammaValue, greenMax GammaValue, greenGamma GammaValue, blueMin GammaValue, blueMax GammaValue, blueGamma GammaValue) Error {
	rv := C.SetDisplayTransferByFormula(C.uint32_t(display), C.float(redMin), C.float(redMax), C.float(redGamma), C.float(greenMin), C.float(greenMax), C.float(greenGamma), C.float(blueMin), C.float(blueMax), C.float(blueGamma))
	return Error(rv)
}

// Returns a Boolean value indicating whether a display is sleeping (and is therefore not drawable). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454260-cgdisplayisasleep?language=objc
func DisplayIsAsleep(display DirectDisplayID) kernel.Boolean_t {
	rv := C.DisplayIsAsleep(C.uint32_t(display))
	return kernel.Boolean_t(rv)
}

// Returns the bitmap information for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454200-cgimagegetbitmapinfo?language=objc
func ImageGetBitmapInfo(image ImageRef) BitmapInfo {
	rv := C.ImageGetBitmapInfo(image)
	return BitmapInfo(rv)
}

// Returns a composite image of the specified windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455730-cgwindowlistcreateimagefromarray?language=objc
func WindowListCreateImageFromArray(screenBounds Rect, windowArray corefoundation.ArrayRef, imageOption WindowImageOption) ImageRef {
	rv := C.WindowListCreateImageFromArray(screenBounds, windowArray, C.uint32_t(imageOption))
	return ImageRef(rv)
}

// Returns the width in pixels of a bitmap context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455607-cgbitmapcontextgetwidth?language=objc
func BitmapContextGetWidth(context ContextRef) uint {
	rv := C.BitmapContextGetWidth(context)
	return uint(rv)
}

// Returns an affine transformation matrix constructed from values you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455865-cgaffinetransformmake?language=objc
func AffineTransformMake(a float64, b float64, c float64, d float64, tx float64, ty float64) AffineTransform {
	rv := C.AffineTransformMake(a, b, c, d, tx, ty)
	return AffineTransform(rv)
}

// Create an immutable path of an ellipse. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411177-cgpathcreatewithellipseinrect?language=objc
func PathCreateWithEllipseInRect(rect Rect, transform *AffineTransform) unsafe.Pointer {
	rv := C.PathCreateWithEllipseInRect(rect, transform)
	return unsafe.Pointer(rv)
}

// Returns whether two rectangles intersect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454747-cgrectintersectsrect?language=objc
func RectIntersectsRect(rect1 Rect, rect2 Rect) bool {
	rv := C.RectIntersectsRect(rect1, rect2)
	return bool(rv)
}

// Returns an image containing the contents of the specified display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455691-cgdisplaycreateimage?language=objc
func DisplayCreateImage(displayID DirectDisplayID) ImageRef {
	rv := C.DisplayCreateImage(C.uint32_t(displayID))
	return ImageRef(rv)
}

// Sets the line width for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455270-cgcontextsetlinewidth?language=objc
func ContextSetLineWidth(c ContextRef, width float64) {
	C.ContextSetLineWidth(c, width)
}

// Enables or disables the merging of actual key and mouse state with the application-specified state in a synthetic event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541787-cgenableeventstatecombining?language=objc
func EnableEventStateCombining(combineState kernel.Boolean_t) Error {
	rv := C.EnableEventStateCombining(C.NSInteger(combineState))
	return Error(rv)
}

// Creates an ICC-based color space using the ICC profile contained in the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408895-cgcolorspacecreatewithiccprofile?language=objc
func ColorSpaceCreateWithICCProfile(data corefoundation.DataRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateWithICCProfile(data)
	return ColorSpaceRef(rv)
}

// Decrements the retain count of a Core Graphics pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1552266-cgpatternrelease?language=objc
func PatternRelease(pattern PatternRef) {
	C.PatternRelease(pattern)
}

// Modifies the current clipping path, using the even-odd rule. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455944-cgcontexteoclip?language=objc
func ContextEOClip(c ContextRef) {
	C.ContextEOClip(c)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3684557-cgcolorspacecreateextended?language=objc
func ColorSpaceCreateExtended(space ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateExtended(space)
	return ColorSpaceRef(rv)
}

// Sets whether or not to allow subpixel quantization for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456263-cgcontextsetallowsfontsubpixelqu?language=objc
func ContextSetAllowsFontSubpixelQuantization(c ContextRef, allowsFontSubpixelQuantization bool) {
	C.ContextSetAllowsFontSubpixelQuantization(c, allowsFontSubpixelQuantization)
}

// Returns a point that is transformed from device space coordinates to user space coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456451-cgcontextconvertpointtouserspace?language=objc
func ContextConvertPointToUserSpace(c ContextRef, point Point) Point {
	rv := C.ContextConvertPointToUserSpace(c, point)
	return Point(rv)
}

// Increments the retain count of a data consumer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1508422-cgdataconsumerretain?language=objc
func DataConsumerRetain(consumer DataConsumerRef) DataConsumerRef {
	rv := C.DataConsumerRetain(consumer)
	return DataConsumerRef(rv)
}

// Sets the mask that indicates which classes of local hardware events are enabled during event suppression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408770-cgeventsourcesetlocaleventsfilte?language=objc
func EventSourceSetLocalEventsFilterDuringSuppressionState(source EventSourceRef, filter EventFilterMask, state EventSuppressionState) {
	C.EventSourceSetLocalEventsFilterDuringSuppressionState(source, C.uint32_t(filter), C.uint32_t(state))
}

// Returns information about the current display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1562062-cgdisplaycurrentmode?language=objc
func DisplayCurrentMode(display DirectDisplayID) corefoundation.DictionaryRef {
	rv := C.DisplayCurrentMode(C.uint32_t(display))
	return corefoundation.DictionaryRef(rv)
}

// Waits for screen update operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541790-cgwaitforscreenupdaterects?language=objc
func WaitForScreenUpdateRects(requestedOperations ScreenUpdateOperation, currentOperation *ScreenUpdateOperation, rects *Rect, rectCount *uint, delta *ScreenUpdateMoveDelta) Error {
	rv := C.WaitForScreenUpdateRects(C.uint32_t(requestedOperations), currentOperation, rects, rectCount, delta)
	return Error(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2866016-cgcolorconversioninfocreatefroml?language=objc
func ColorConversionInfoCreateFromListWithArguments(options corefoundation.DictionaryRef, arg0 ColorSpaceRef, arg1 ColorConversionInfoTransformType, arg2 ColorRenderingIntent, arg3 kernel.Va_list) ColorConversionInfoRef {
	rv := C.ColorConversionInfoCreateFromListWithArguments(options, arg0, C.uint32_t(arg1), C.int32_t(arg2), C.NSObject*(arg3))
	return ColorConversionInfoRef(rv)
}

// Sets the current stroke color to a value in the DeviceGray color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455209-cgcontextsetgraystrokecolor?language=objc
func ContextSetGrayStrokeColor(c ContextRef, gray float64, alpha float64) {
	C.ContextSetGrayStrokeColor(c, gray, alpha)
}

// Returns the vendor number of the specified display’s monitor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455233-cgdisplayvendornumber?language=objc
func DisplayVendorNumber(display DirectDisplayID) uint32 {
	rv := C.DisplayVendorNumber(C.uint32_t(display))
	return uint32(rv)
}

// Returns whether a rectangle contains a specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456316-cgrectcontainspoint?language=objc
func RectContainsPoint(rect Rect, point Point) bool {
	rv := C.RectContainsPoint(rect, point)
	return bool(rv)
}

// Returns the Core Foundation type identifier for Core Graphics fonts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396369-cgfontgettypeid?language=objc
func FontGetTypeID() corefoundation.TypeID {
	rv := C.FontGetTypeID()
	return corefoundation.TypeID(rv)
}

// Returns a graphics context suitable for drawing to a captured display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456576-cgdisplaygetdrawingcontext?language=objc
func DisplayGetDrawingContext(display DirectDisplayID) ContextRef {
	rv := C.DisplayGetDrawingContext(C.uint32_t(display))
	return ContextRef(rv)
}

// Retrieves a character string from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454584-cgpdfscannerpopname?language=objc
func PDFScannerPopName(scanner unsafe.Pointer, value *uint8) bool {
	rv := C.PDFScannerPopName(scanner, value)
	return bool(rv)
}

// Enables or disables font smoothing in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455816-cgcontextsetshouldsmoothfonts?language=objc
func ContextSetShouldSmoothFonts(c ContextRef, shouldSmoothFonts bool) {
	C.ContextSetShouldSmoothFonts(c, shouldSmoothFonts)
}

// Returns a new Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454913-cgeventcreate?language=objc
func EventCreate(source EventSourceRef) EventRef {
	rv := C.EventCreate(source)
	return EventRef(rv)
}

// For a secondary display in a mirroring set, returns the primary display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454556-cgdisplaymirrorsdisplay?language=objc
func DisplayMirrorsDisplay(display DirectDisplayID) DirectDisplayID {
	rv := C.DisplayMirrorsDisplay(C.uint32_t(display))
	return DirectDisplayID(rv)
}

// Checks whether an affine transform is the identity transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455754-cgaffinetransformisidentity?language=objc
func AffineTransformIsIdentity(t AffineTransform) bool {
	rv := C.AffineTransformIsIdentity(t)
	return bool(rv)
}

// Starts a new page in a page-based graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454794-cgcontextbeginpage?language=objc
func ContextBeginPage(c ContextRef, mediaBox *Rect) {
	C.ContextBeginPage(c, mediaBox)
}

// Returns the width of a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454758-cgrectgetwidth?language=objc
func RectGetWidth(rect Rect) float64 {
	rv := C.RectGetWidth(rect)
	return float64(rv)
}

// Returns a point with the specified coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455746-cgpointmake?language=objc
func PointMake(x float64, y float64) Point {
	rv := C.PointMake(x, y)
	return Point(rv)
}

// Returns a new Quartz scrolling event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541327-cgeventcreatescrollwheelevent?language=objc
func EventCreateScrollWheelEvent(source EventSourceRef, units ScrollEventUnit, wheelCount uint32, wheel1 int32) EventRef {
	rv := C.EventCreateScrollWheelEvent(source, C.uint32_t(units), wheelCount, wheel1)
	return EventRef(rv)
}

// Returns a color object that represents a constant color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454283-cgcolorgetconstantcolor?language=objc
func ColorGetConstantColor(colorName corefoundation.StringRef) ColorRef {
	rv := C.ColorGetConstantColor(colorName)
	return ColorRef(rv)
}

// Sets the time interval in seconds that local hardware events are suppressed after posting a synthetic event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541800-cgsetlocaleventssuppressioninter?language=objc
func SetLocalEventsSuppressionInterval(seconds corefoundation.TimeInterval) Error {
	rv := C.SetLocalEventsSuppressionInterval(C.double(seconds))
	return Error(rv)
}

// Returns a Quartz event source created from an existing Quartz event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455393-cgeventcreatesourcefromevent?language=objc
func EventCreateSourceFromEvent(event EventRef) EventSourceRef {
	rv := C.EventCreateSourceFromEvent(event)
	return EventSourceRef(rv)
}

// Returns the x-height of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396410-cgfontgetxheight?language=objc
func FontGetXHeight(font FontRef) int {
	rv := C.FontGetXHeight(font)
	return int(rv)
}

// Tells a stream to start sending updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454870-cgdisplaystreamstart?language=objc
func DisplayStreamStart(displayStream DisplayStreamRef) Error {
	rv := C.DisplayStreamStart(displayStream)
	return Error(rv)
}

// Returns the location of a Quartz mouse event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455589-cgeventgetunflippedlocation?language=objc
func EventGetUnflippedLocation(event EventRef) Point {
	rv := C.EventGetUnflippedLocation(event)
	return Point(rv)
}

// Creates a bitmap image from an existing image and an image mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456337-cgimagecreatewithmask?language=objc
func ImageCreateWithMask(image ImageRef, mask ImageRef) ImageRef {
	rv := C.ImageCreateWithMask(image, mask)
	return ImageRef(rv)
}

// Sets the URL associated with a rectangle in a PDF graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455622-cgpdfcontextseturlforrect?language=objc
func PDFContextSetURLForRect(context ContextRef, url corefoundation.URLRef, rect Rect) {
	C.PDFContextSetURLForRect(context, url, rect)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/2091882-cgcolorspacecreatelinearized?language=objc
func ColorSpaceCreateLinearized(space ColorSpaceRef) ColorSpaceRef {
	rv := C.ColorSpaceCreateLinearized(space)
	return ColorSpaceRef(rv)
}

// Returns the number of entries in the color table of an indexed color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408883-cgcolorspacegetcolortablecount?language=objc
func ColorSpaceGetColorTableCount(space ColorSpaceRef) uint {
	rv := C.ColorSpaceGetColorTableCount(space)
	return uint(rv)
}

// Returns information about the display mode closest to a specified depth, screen size, and refresh rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1562066-cgdisplaybestmodeforparametersan?language=objc
func DisplayBestModeForParametersAndRefreshRate(display DirectDisplayID, bitsPerPixel uint, width uint, height uint, refreshRate RefreshRate, exactMatch *kernel.Boolean_t) corefoundation.DictionaryRef {
	rv := C.DisplayBestModeForParametersAndRefreshRate(C.uint32_t(display), bitsPerPixel, width, height, C.double(refreshRate), exactMatch)
	return corefoundation.DictionaryRef(rv)
}

// Maps an OpenGL display mask to a display ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455386-cgopengldisplaymasktodisplayid?language=objc
func OpenGLDisplayMaskToDisplayID(mask OpenGLDisplayMask) DirectDisplayID {
	rv := C.OpenGLDisplayMaskToDisplayID(C.uint32_t(mask))
	return DirectDisplayID(rv)
}

// Returns the number of frames that have been dropped since the last call to your update handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456212-cgdisplaystreamupdategetdropcoun?language=objc
func DisplayStreamUpdateGetDropCount(updateRef DisplayStreamUpdateRef) uint {
	rv := C.DisplayStreamUpdateGetDropCount(updateRef)
	return uint(rv)
}

// Returns whether two points are equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456179-cgpointequaltopoint?language=objc
func PointEqualToPoint(point1 Point, point2 Point) bool {
	rv := C.PointEqualToPoint(point1, point2)
	return bool(rv)
}

// Forces all pending drawing operations in a window context to be rendered immediately to the destination device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454895-cgcontextflush?language=objc
func ContextFlush(c ContextRef) {
	C.ContextFlush(c)
}

// Returns a vector with the specified dimension values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454811-cgvectormake?language=objc
func VectorMake(dx float64, dy float64) Vector {
	rv := C.VectorMake(dx, dy)
	return Vector(rv)
}

// Returns the height of the specified display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455380-cgdisplaymodegetheight?language=objc
func DisplayModeGetHeight(mode DisplayModeRef) uint {
	rv := C.DisplayModeGetHeight(mode)
	return uint(rv)
}

// Returns the smallest rectangle that contains the two source rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455837-cgrectunion?language=objc
func RectUnion(r1 Rect, r2 Rect) Rect {
	rv := C.RectUnion(r1, r2)
	return Rect(rv)
}

// Parses the content stream of a PDF scanner object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454698-cgpdfscannerscan?language=objc
func PDFScannerScan(scanner unsafe.Pointer) bool {
	rv := C.PDFScannerScan(scanner)
	return bool(rv)
}

// Adds a sequence of connected straight-line segments to the current path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455461-cgcontextaddlines?language=objc
func ContextAddLines(c ContextRef, points *Point, count uint) {
	C.ContextAddLines(c, points, count)
}

// Sets the miter limit for the joins of connected lines in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456499-cgcontextsetmiterlimit?language=objc
func ContextSetMiterLimit(c ContextRef, limit float64) {
	C.ContextSetMiterLimit(c, limit)
}

// Returns the dictionary associated with a PDF stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456118-cgpdfstreamgetdictionary?language=objc
func PDFStreamGetDictionary(stream unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFStreamGetDictionary(stream)
	return unsafe.Pointer(rv)
}

// Returns the number of bytes in a PDF string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454095-cgpdfstringgetlength?language=objc
func PDFStringGetLength(string unsafe.Pointer) uint {
	rv := C.PDFStringGetLength(string_)
	return uint(rv)
}

// Modifies the current clipping path, using the nonzero winding number rule. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455262-cgcontextclip?language=objc
func ContextClip(c ContextRef) {
	C.ContextClip(c)
}

// Releases all captured displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454901-cgreleasealldisplays?language=objc
func ReleaseAllDisplays() Error {
	rv := C.ReleaseAllDisplays()
	return Error(rv)
}

// Returns a Boolean value indicating whether the mouse cursor is drawn in framebuffer memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541804-cgcursorisdrawninframebuffer?language=objc
func CursorIsDrawnInFramebuffer() kernel.Boolean_t {
	rv := C.CursorIsDrawnInFramebuffer()
	return kernel.Boolean_t(rv)
}

// Returns the serial number of a display monitor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455409-cgdisplayserialnumber?language=objc
func DisplaySerialNumber(display DirectDisplayID) uint32 {
	rv := C.DisplaySerialNumber(C.uint32_t(display))
	return uint32(rv)
}

// Returns whether an object at a given index in a PDF array is another PDF array and, if so, retrieves that array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454834-cgpdfarraygetarray?language=objc
func PDFArrayGetArray(array unsafe.Pointer, index uint, value unsafe.Pointer) bool {
	rv := C.PDFArrayGetArray(array, index, value)
	return bool(rv)
}

// Returns the name used to create the specified color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408903-cgcolorspacecopyname?language=objc
func ColorSpaceCopyName(space ColorSpaceRef) corefoundation.StringRef {
	rv := C.ColorSpaceCopyName(space)
	return corefoundation.StringRef(rv)
}

// Sets the style for the endpoints of lines drawn in a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454326-cgcontextsetlinecap?language=objc
func ContextSetLineCap(c ContextRef, cap LineCap) {
	C.ContextSetLineCap(c, C.int32_t(cap))
}

// Decrements the retain count of a shading object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1573766-cgshadingrelease?language=objc
func ShadingRelease(shading ShadingRef) {
	C.ShadingRelease(shading)
}

// Increments the retain count of a CGPDFOperatorTable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454547-cgpdfoperatortableretain?language=objc
func PDFOperatorTableRetain(table unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFOperatorTableRetain(table)
	return unsafe.Pointer(rv)
}

// Returns the model number of a display monitor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454149-cgdisplaymodelnumber?language=objc
func DisplayModelNumber(display DirectDisplayID) uint32 {
	rv := C.DisplayModelNumber(C.uint32_t(display))
	return uint32(rv)
}

// Sets the 64-bit user-specified data for a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408779-cgeventsourcesetuserdata?language=objc
func EventSourceSetUserData(source EventSourceRef, userData int64) {
	C.EventSourceSetUserData(source, userData)
}

// Decrements the retain count of a CGGradient object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1398460-cggradientrelease?language=objc
func GradientRelease(gradient GradientRef) {
	C.GradientRelease(gradient)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656521-cgrequestlisteneventaccess?language=objc
func RequestListenEventAccess() bool {
	rv := C.RequestListenEventAccess()
	return bool(rv)
}

// Begins a new subpath at the point you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454738-cgcontextmovetopoint?language=objc
func ContextMoveToPoint(c ContextRef, x float64, y float64) {
	C.ContextMoveToPoint(c, x, y)
}

// Returns an affine transformation matrix constructed by inverting an existing affine transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455264-cgaffinetransforminvert?language=objc
func AffineTransformInvert(t AffineTransform) AffineTransform {
	rv := C.AffineTransformInvert(t)
	return AffineTransform(rv)
}

// Increments the retain count of a scanner object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455810-cgpdfscannerretain?language=objc
func PDFScannerRetain(scanner unsafe.Pointer) unsafe.Pointer {
	rv := C.PDFScannerRetain(scanner)
	return unsafe.Pointer(rv)
}

// Indicates whether or not a graphics path is empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411149-cgpathisempty?language=objc
func PathIsEmpty(path unsafe.Pointer) bool {
	rv := C.PathIsEmpty(path)
	return bool(rv)
}

// Repeatedly draws an image, scaled to the provided rectangle, to fill the current clip region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456240-cgcontextdrawtiledimage?language=objc
func ContextDrawTiledImage(c ContextRef, rect Rect, image ImageRef) {
	C.ContextDrawTiledImage(c, rect, image)
}

// Increments the retain count of a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1556741-cgimageretain?language=objc
func ImageRetain(image ImageRef) ImageRef {
	rv := C.ImageRetain(image)
	return ImageRef(rv)
}

// Creates a bitmap image using the data contained within a subregion of an existing bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454683-cgimagecreatewithimageinrect?language=objc
func ImageCreateWithImageInRect(image ImageRef, rect Rect) ImageRef {
	rv := C.ImageCreateWithImageInRect(image, rect)
	return ImageRef(rv)
}

// Returns the rotation angle of a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455550-cgpdfpagegetrotationangle?language=objc
func PDFPageGetRotationAngle(page PDFPageRef) int {
	rv := C.PDFPageGetRotationAngle(page)
	return int(rv)
}

// Configures the display mode of a display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1543535-cgconfiguredisplaymode?language=objc
func ConfigureDisplayMode(config unsafe.Pointer, display DirectDisplayID, mode corefoundation.DictionaryRef) Error {
	rv := C.ConfigureDisplayMode(config, C.uint32_t(display), mode)
	return Error(rv)
}

// Appends a straight line segment from the current point to the provided point . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455213-cgcontextaddlinetopoint?language=objc
func ContextAddLineToPoint(c ContextRef, x float64, y float64) {
	C.ContextAddLineToPoint(c, x, y)
}

// Retrieves a Boolean object from the scanner stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454663-cgpdfscannerpopboolean?language=objc
func PDFScannerPopBoolean(scanner unsafe.Pointer, value *PDFBoolean) bool {
	rv := C.PDFScannerPopBoolean(scanner, value)
	return bool(rv)
}

// Creates a font object corresponding to the font specified by a PostScript or full name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396330-cgfontcreatewithfontname?language=objc
func FontCreateWithFontName(name corefoundation.StringRef) FontRef {
	rv := C.FontCreateWithFontName(name)
	return FontRef(rv)
}

// Sets the level of interpolation quality for a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455656-cgcontextsetinterpolationquality?language=objc
func ContextSetInterpolationQuality(c ContextRef, quality InterpolationQuality) {
	C.ContextSetInterpolationQuality(c, C.int32_t(quality))
}

// Returns the document catalog of a Core Graphics PDF document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402606-cgpdfdocumentgetcatalog?language=objc
func PDFDocumentGetCatalog(document PDFDocumentRef) unsafe.Pointer {
	rv := C.PDFDocumentGetCatalog(document)
	return unsafe.Pointer(rv)
}

// Decrements the retain count of a layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450898-cglayerrelease?language=objc
func LayerRelease(layer LayerRef) {
	C.LayerRelease(layer)
}

// Returns the unique type identifier used for CGLayerRef objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1450888-cglayergettypeid?language=objc
func LayerGetTypeID() corefoundation.TypeID {
	rv := C.LayerGetTypeID()
	return corefoundation.TypeID(rv)
}

// Returns the type identifier for Core Graphics function objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1390879-cgfunctiongettypeid?language=objc
func FunctionGetTypeID() corefoundation.TypeID {
	rv := C.FunctionGetTypeID()
	return corefoundation.TypeID(rv)
}

// Gets the array of PDF content streams contained in a PDF content stream object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1410048-cgpdfcontentstreamgetstreams?language=objc
func PDFContentStreamGetStreams(cs unsafe.Pointer) corefoundation.ArrayRef {
	rv := C.PDFContentStreamGetStreams(cs)
	return corefoundation.ArrayRef(rv)
}

// Returns the number of entries in a PDF dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1430218-cgpdfdictionarygetcount?language=objc
func PDFDictionaryGetCount(dict unsafe.Pointer) uint {
	rv := C.PDFDictionaryGetCount(dict)
	return uint(rv)
}

// Decrements the hide cursor count, and shows the mouse cursor if the count is 0. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455867-cgdisplayshowcursor?language=objc
func DisplayShowCursor(display DirectDisplayID) Error {
	rv := C.DisplayShowCursor(C.uint32_t(display))
	return Error(rv)
}

// Obtains the bitmap information associated with a bitmap graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455839-cgbitmapcontextgetbitmapinfo?language=objc
func BitmapContextGetBitmapInfo(context ContextRef) BitmapInfo {
	rv := C.BitmapContextGetBitmapInfo(context)
	return BitmapInfo(rv)
}

// Returns an affine transformation matrix constructed by scaling an existing affine transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455882-cgaffinetransformscale?language=objc
func AffineTransformScale(t AffineTransform, sx float64, sy float64) AffineTransform {
	rv := C.AffineTransformScale(t, sx, sy)
	return AffineTransform(rv)
}

// Creates a subset of the font in the specified PostScript format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1396324-cgfontcreatepostscriptsubset?language=objc
func FontCreatePostScriptSubset(font FontRef, subsetName corefoundation.StringRef, format FontPostScriptFormat, glyphs *Glyph, count uint, encoding *Glyph) corefoundation.DataRef {
	rv := C.FontCreatePostScriptSubset(font, subsetName, C.int32_t(format), glyphs, count, encoding)
	return corefoundation.DataRef(rv)
}

// Returns a Boolean value indicating whether a display is the main display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1455448-cgdisplayismain?language=objc
func DisplayIsMain(display DirectDisplayID) kernel.Boolean_t {
	rv := C.DisplayIsMain(C.uint32_t(display))
	return kernel.Boolean_t(rv)
}

// Creates a mutable graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411209-cgpathcreatemutable?language=objc
func PathCreateMutable() MutablePathRef {
	rv := C.PathCreateMutable()
	return MutablePathRef(rv)
}

// Deallocates a list of rectangles that represent changed areas on local displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1541780-cgreleasescreenrefreshrects?language=objc
func ReleaseScreenRefreshRects(rects *Rect) {
	C.ReleaseScreenRefreshRects(rects)
}

// Displays an array of glyphs at the current text position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1586500-cgcontextshowglyphs?language=objc
func ContextShowGlyphs(c ContextRef, g *Glyph, count uint) {
	C.ContextShowGlyphs(c, g, count)
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
	C.ContextSetLineDash(c, phase, lengths, count)
}

// Returns the type identifier of Quartz display modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1456191-cgdisplaymodegettypeid?language=objc
func DisplayModeGetTypeID() corefoundation.TypeID {
	rv := C.DisplayModeGetTypeID()
	return corefoundation.TypeID(rv)
}

// Returns the current flags of a Quartz event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1408792-cgeventsourceflagsstate?language=objc
func EventSourceFlagsState(stateID EventSourceStateID) EventFlags {
	rv := C.EventSourceFlagsState(C.int32_t(stateID))
	return EventFlags(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/3656522-cgrequestposteventaccess?language=objc
func RequestPostEventAccess() bool {
	rv := C.RequestPostEventAccess()
	return bool(rv)
}

// Paints the area contained within the provided rectangle, using the fill color in the current graphics state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1454700-cgcontextfillrect?language=objc
func ContextFillRect(c ContextRef, rect Rect) {
	C.ContextFillRect(c, rect)
}

// Increments the retain count of a graphics path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1411181-cgpathretain?language=objc
func PathRetain(path unsafe.Pointer) unsafe.Pointer {
	rv := C.PathRetain(path)
	return unsafe.Pointer(rv)
}

// Returns whether the specified PDF document is currently unlocked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/1402607-cgpdfdocumentisunlocked?language=objc
func PDFDocumentIsUnlocked(document PDFDocumentRef) bool {
	rv := C.PDFDocumentIsUnlocked(document)
	return bool(rv)
}
