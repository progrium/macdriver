// AUTO-GENERATED CODE, DO NOT MODIFY

#import "CoreFoundation/CoreFoundation.h"
void * DictionaryCreateMutableCopy(void * allocator, CFIndex capacity, void * theDict) {
	return (void *)CFDictionaryCreateMutableCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFIndex)capacity,
		// *typing.RefType
		(CFDictionaryRef)theDict
	);
}
CFTypeID NullGetTypeID() {
	return (CFTypeID)CFNullGetTypeID(
	);
}
void * WriteStreamCopyError(void * stream) {
	return (void *)CFWriteStreamCopyError(
		// *typing.RefType
		(CFWriteStreamRef)stream
	);
}
void * BundleCopyResourceURLForLocalization(void * bundle, void * resourceName, void * resourceType, void * subDirName, void * localizationName) {
	return (void *)CFBundleCopyResourceURLForLocalization(
		// *typing.RefType
		(CFBundleRef)bundle,
		// *typing.RefType
		(CFStringRef)resourceName,
		// *typing.RefType
		(CFStringRef)resourceType,
		// *typing.RefType
		(CFStringRef)subDirName,
		// *typing.RefType
		(CFStringRef)localizationName
	);
}
void * StringCreateCopy(void * alloc, void * theString) {
	return (void *)CFStringCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFStringRef)theString
	);
}
void CalendarSetFirstWeekday(void * calendar, CFIndex wkdy) {
	return (void)CFCalendarSetFirstWeekday(
		// *typing.RefType
		(CFCalendarRef)calendar,
		// *typing.AliasType
		(CFIndex)wkdy
	);
}
CFTypeID NumberFormatterGetTypeID() {
	return (CFTypeID)CFNumberFormatterGetTypeID(
	);
}
bool BundleIsArchitectureLoadable(cpu_type_t arch) {
	return (bool)CFBundleIsArchitectureLoadable(
		// *typing.AliasType
		(cpu_type_t)arch
	);
}
void * BundleCopyInfoDictionaryInDirectory(void * bundleURL) {
	return (void *)CFBundleCopyInfoDictionaryInDirectory(
		// *typing.RefType
		(CFURLRef)bundleURL
	);
}
uint32_t SwapInt32(uint32_t arg) {
	return (uint32_t)CFSwapInt32(
		// *typing.PrimitiveType
		(uint32_t)arg
	);
}
void * URLCopyStrictPath(void * anURL, bool* isAbsolute) {
	return (void *)CFURLCopyStrictPath(
		// *typing.RefType
		(CFURLRef)anURL,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(Boolean*)isAbsolute
	);
}
void * TimeZoneCopySystem() {
	return (void *)CFTimeZoneCopySystem(
	);
}
void * FileSecurityCreateCopy(void * allocator, void * fileSec) {
	return (void *)CFFileSecurityCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFFileSecurityRef)fileSec
	);
}
void * URLCreateWithFileSystemPath(void * allocator, void * filePath, CFURLPathStyle pathStyle, bool isDirectory) {
	return (void *)CFURLCreateWithFileSystemPath(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFStringRef)filePath,
		// *typing.AliasType
		(CFURLPathStyle)pathStyle,
		// *typing.PrimitiveType
		(Boolean)isDirectory
	);
}
CFIndex BagGetCount(void * theBag) {
	return (CFIndex)CFBagGetCount(
		// *typing.RefType
		(CFBagRef)theBag
	);
}
void * BundleCopyAuxiliaryExecutableURL(void * bundle, void * executableName) {
	return (void *)CFBundleCopyAuxiliaryExecutableURL(
		// *typing.RefType
		(CFBundleRef)bundle,
		// *typing.RefType
		(CFStringRef)executableName
	);
}
void * BagCreateCopy(void * allocator, void * theBag) {
	return (void *)CFBagCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFBagRef)theBag
	);
}
void WriteStreamClose(void * stream) {
	return (void)CFWriteStreamClose(
		// *typing.RefType
		(CFWriteStreamRef)stream
	);
}
bool PreferencesGetAppBooleanValue(void * key, void * applicationID, bool* keyExistsAndHasValidFormat) {
	return (bool)CFPreferencesGetAppBooleanValue(
		// *typing.RefType
		(CFStringRef)key,
		// *typing.RefType
		(CFStringRef)applicationID,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(Boolean*)keyExistsAndHasValidFormat
	);
}
void * SetCreateCopy(void * allocator, void * theSet) {
	return (void *)CFSetCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFSetRef)theSet
	);
}
void DateFormatterSetFormat(void * formatter, void * formatString) {
	return (void)CFDateFormatterSetFormat(
		// *typing.RefType
		(CFDateFormatterRef)formatter,
		// *typing.RefType
		(CFStringRef)formatString
	);
}
void TreeRemoveAllChildren(void * tree) {
	return (void)CFTreeRemoveAllChildren(
		// *typing.RefType
		(CFTreeRef)tree
	);
}
void * CharacterSetCreateWithBitmapRepresentation(void * alloc, void * theData) {
	return (void *)CFCharacterSetCreateWithBitmapRepresentation(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFDataRef)theData
	);
}
CFBit BitVectorGetBitAtIndex(void * bv, CFIndex idx) {
	return (CFBit)CFBitVectorGetBitAtIndex(
		// *typing.RefType
		(CFBitVectorRef)bv,
		// *typing.AliasType
		(CFIndex)idx
	);
}
void * GetAllocator(CFTypeRef cf) {
	return (void *)CFGetAllocator(
		// *typing.AliasType
		(CFTypeRef)cf
	);
}
void * BundleCopyBuiltInPlugInsURL(void * bundle) {
	return (void *)CFBundleCopyBuiltInPlugInsURL(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
CFIndex ErrorGetCode(void * err) {
	return (CFIndex)CFErrorGetCode(
		// *typing.RefType
		(CFErrorRef)err
	);
}
CFIndex URLGetBytes(void * url, uint8_t* buffer, CFIndex bufferLength) {
	return (CFIndex)CFURLGetBytes(
		// *typing.RefType
		(CFURLRef)url,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)buffer,
		// *typing.AliasType
		(CFIndex)bufferLength
	);
}
void BitVectorSetBitAtIndex(void * bv, CFIndex idx, CFBit value) {
	return (void)CFBitVectorSetBitAtIndex(
		// *typing.RefType
		(CFMutableBitVectorRef)bv,
		// *typing.AliasType
		(CFIndex)idx,
		// *typing.AliasType
		(CFBit)value
	);
}
void PlugInAddInstanceForFactory(void * factoryID) {
	return (void)CFPlugInAddInstanceForFactory(
		// *typing.RefType
		(CFUUIDRef)factoryID
	);
}
CFIndex CalendarGetMinimumDaysInFirstWeek(void * calendar) {
	return (CFIndex)CFCalendarGetMinimumDaysInFirstWeek(
		// *typing.RefType
		(CFCalendarRef)calendar
	);
}
void DataIncreaseLength(void * theData, CFIndex extraLength) {
	return (void)CFDataIncreaseLength(
		// *typing.RefType
		(CFMutableDataRef)theData,
		// *typing.AliasType
		(CFIndex)extraLength
	);
}
bool RunLoopTimerIsValid(void * timer) {
	return (bool)CFRunLoopTimerIsValid(
		// *typing.RefType
		(CFRunLoopTimerRef)timer
	);
}
void * BundleCopyResourceURL(void * bundle, void * resourceName, void * resourceType, void * subDirName) {
	return (void *)CFBundleCopyResourceURL(
		// *typing.RefType
		(CFBundleRef)bundle,
		// *typing.RefType
		(CFStringRef)resourceName,
		// *typing.RefType
		(CFStringRef)resourceType,
		// *typing.RefType
		(CFStringRef)subDirName
	);
}
void * StringCreateWithCharacters(void * alloc, const UniChar* chars, CFIndex numChars) {
	return (void *)CFStringCreateWithCharacters(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.PointerType
		// -> *typing.AliasType
		(UniChar*)chars,
		// *typing.AliasType
		(CFIndex)numChars
	);
}
CFIndex ReadStreamRead(void * stream, uint8_t* buffer, CFIndex bufferLength) {
	return (CFIndex)CFReadStreamRead(
		// *typing.RefType
		(CFReadStreamRef)stream,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)buffer,
		// *typing.AliasType
		(CFIndex)bufferLength
	);
}
CFTypeID RunLoopObserverGetTypeID() {
	return (CFTypeID)CFRunLoopObserverGetTypeID(
	);
}
void * NumberFormatterCreate(void * allocator, void * locale, CFNumberFormatterStyle style) {
	return (void *)CFNumberFormatterCreate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFLocaleRef)locale,
		// *typing.AliasType
		(CFNumberFormatterStyle)style
	);
}
void * NumberFormatterGetLocale(void * formatter) {
	return (void *)CFNumberFormatterGetLocale(
		// *typing.RefType
		(CFNumberFormatterRef)formatter
	);
}
void * MessagePortCreateRemote(void * allocator, void * name) {
	return (void *)CFMessagePortCreateRemote(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFStringRef)name
	);
}
CFOptionFlags RunLoopObserverGetActivities(void * observer) {
	return (CFOptionFlags)CFRunLoopObserverGetActivities(
		// *typing.RefType
		(CFRunLoopObserverRef)observer
	);
}
bool PropertyListIsValid(CFPropertyListRef plist, CFPropertyListFormat format) {
	return (bool)CFPropertyListIsValid(
		// *typing.AliasType
		(CFPropertyListRef)plist,
		// *typing.AliasType
		(CFPropertyListFormat)format
	);
}
CFTimeInterval RunLoopTimerGetInterval(void * timer) {
	return (CFTimeInterval)CFRunLoopTimerGetInterval(
		// *typing.RefType
		(CFRunLoopTimerRef)timer
	);
}
double StringGetDoubleValue(void * str) {
	return (double)CFStringGetDoubleValue(
		// *typing.RefType
		(CFStringRef)str
	);
}
void * TimeZoneCopyAbbreviationDictionary() {
	return (void *)CFTimeZoneCopyAbbreviationDictionary(
	);
}
void FileDescriptorEnableCallBacks(void * f, CFOptionFlags callBackTypes) {
	return (void)CFFileDescriptorEnableCallBacks(
		// *typing.RefType
		(CFFileDescriptorRef)f,
		// *typing.AliasType
		(CFOptionFlags)callBackTypes
	);
}
void * BundleGetInfoDictionary(void * bundle) {
	return (void *)CFBundleGetInfoDictionary(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
void * AttributedStringCreate(void * alloc, void * str, void * attributes) {
	return (void *)CFAttributedStringCreate(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFStringRef)str,
		// *typing.RefType
		(CFDictionaryRef)attributes
	);
}
uint64_t SwapInt64HostToLittle(uint64_t arg) {
	return (uint64_t)CFSwapInt64HostToLittle(
		// *typing.PrimitiveType
		(uint64_t)arg
	);
}
void FileDescriptorInvalidate(void * f) {
	return (void)CFFileDescriptorInvalidate(
		// *typing.RefType
		(CFFileDescriptorRef)f
	);
}
void * DataCreate(void * allocator, const uint8_t* bytes, CFIndex length) {
	return (void *)CFDataCreate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)bytes,
		// *typing.AliasType
		(CFIndex)length
	);
}
void * BundleCopyInfoDictionaryForURL(void * url) {
	return (void *)CFBundleCopyInfoDictionaryForURL(
		// *typing.RefType
		(CFURLRef)url
	);
}
void * StringConvertEncodingToIANACharSetName(CFStringEncoding encoding) {
	return (void *)CFStringConvertEncodingToIANACharSetName(
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
void * URLGetString(void * anURL) {
	return (void *)CFURLGetString(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
void DictionaryRemoveAllValues(void * theDict) {
	return (void)CFDictionaryRemoveAllValues(
		// *typing.RefType
		(CFMutableDictionaryRef)theDict
	);
}
void * DataCreateCopy(void * allocator, void * theData) {
	return (void *)CFDataCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFDataRef)theData
	);
}
bool BundleIsExecutableLoaded(void * bundle) {
	return (bool)CFBundleIsExecutableLoaded(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
CFStreamStatus ReadStreamGetStatus(void * stream) {
	return (CFStreamStatus)CFReadStreamGetStatus(
		// *typing.RefType
		(CFReadStreamRef)stream
	);
}
void * URLEnumeratorCreateForDirectoryURL(void * alloc, void * directoryURL, CFURLEnumeratorOptions option, void * propertyKeys) {
	return (void *)CFURLEnumeratorCreateForDirectoryURL(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFURLRef)directoryURL,
		// *typing.AliasType
		(CFURLEnumeratorOptions)option,
		// *typing.RefType
		(CFArrayRef)propertyKeys
	);
}
void * LocaleCopyCurrent() {
	return (void *)CFLocaleCopyCurrent(
	);
}
void * TimeZoneCopyAbbreviation(void * tz, CFAbsoluteTime at) {
	return (void *)CFTimeZoneCopyAbbreviation(
		// *typing.RefType
		(CFTimeZoneRef)tz,
		// *typing.AliasType
		(CFAbsoluteTime)at
	);
}
uint16_t SwapInt16BigToHost(uint16_t arg) {
	return (uint16_t)CFSwapInt16BigToHost(
		// *typing.PrimitiveType
		(uint16_t)arg
	);
}
CFTypeID TimeZoneGetTypeID() {
	return (CFTypeID)CFTimeZoneGetTypeID(
	);
}
void * PreferencesCopyMultiple(void * keysToFetch, void * applicationID, void * userName, void * hostName) {
	return (void *)CFPreferencesCopyMultiple(
		// *typing.RefType
		(CFArrayRef)keysToFetch,
		// *typing.RefType
		(CFStringRef)applicationID,
		// *typing.RefType
		(CFStringRef)userName,
		// *typing.RefType
		(CFStringRef)hostName
	);
}
CFTypeRef URLCreateResourcePropertyForKeyFromBookmarkData(void * allocator, void * resourcePropertyKey, void * bookmark) {
	return (CFTypeRef)CFURLCreateResourcePropertyForKeyFromBookmarkData(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFStringRef)resourcePropertyKey,
		// *typing.RefType
		(CFDataRef)bookmark
	);
}
bool MessagePortSetName(void * ms, void * newName) {
	return (bool)CFMessagePortSetName(
		// *typing.RefType
		(CFMessagePortRef)ms,
		// *typing.RefType
		(CFStringRef)newName
	);
}
void * TreeFindRoot(void * tree) {
	return (void *)CFTreeFindRoot(
		// *typing.RefType
		(CFTreeRef)tree
	);
}
void * CharacterSetCreateInvertedSet(void * alloc, void * theSet) {
	return (void *)CFCharacterSetCreateInvertedSet(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFCharacterSetRef)theSet
	);
}
int32_t UserNotificationUpdate(void * userNotification, CFTimeInterval timeout, CFOptionFlags flags, void * dictionary) {
	return (int32_t)CFUserNotificationUpdate(
		// *typing.RefType
		(CFUserNotificationRef)userNotification,
		// *typing.AliasType
		(CFTimeInterval)timeout,
		// *typing.AliasType
		(CFOptionFlags)flags,
		// *typing.RefType
		(CFDictionaryRef)dictionary
	);
}
void * ErrorCopyRecoverySuggestion(void * err) {
	return (void *)CFErrorCopyRecoverySuggestion(
		// *typing.RefType
		(CFErrorRef)err
	);
}
void * URLEnumeratorCreateForMountedVolumes(void * alloc, CFURLEnumeratorOptions option, void * propertyKeys) {
	return (void *)CFURLEnumeratorCreateForMountedVolumes(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.AliasType
		(CFURLEnumeratorOptions)option,
		// *typing.RefType
		(CFArrayRef)propertyKeys
	);
}
CFTimeInterval RunLoopTimerGetTolerance(void * timer) {
	return (CFTimeInterval)CFRunLoopTimerGetTolerance(
		// *typing.RefType
		(CFRunLoopTimerRef)timer
	);
}
void * BundleCopyLocalizedString(void * bundle, void * key, void * value, void * tableName) {
	return (void *)CFBundleCopyLocalizedString(
		// *typing.RefType
		(CFBundleRef)bundle,
		// *typing.RefType
		(CFStringRef)key,
		// *typing.RefType
		(CFStringRef)value,
		// *typing.RefType
		(CFStringRef)tableName
	);
}
void * RunLoopGetCurrent() {
	return (void *)CFRunLoopGetCurrent(
	);
}
void PreferencesAddSuitePreferencesToApp(void * applicationID, void * suiteID) {
	return (void)CFPreferencesAddSuitePreferencesToApp(
		// *typing.RefType
		(CFStringRef)applicationID,
		// *typing.RefType
		(CFStringRef)suiteID
	);
}
void * CharacterSetCreateCopy(void * alloc, void * theSet) {
	return (void *)CFCharacterSetCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFCharacterSetRef)theSet
	);
}
void Release(CFTypeRef cf) {
	return (void)CFRelease(
		// *typing.AliasType
		(CFTypeRef)cf
	);
}
bool FileDescriptorIsValid(void * f) {
	return (bool)CFFileDescriptorIsValid(
		// *typing.RefType
		(CFFileDescriptorRef)f
	);
}
void * CopyDescription(CFTypeRef cf) {
	return (void *)CFCopyDescription(
		// *typing.AliasType
		(CFTypeRef)cf
	);
}
CFTypeID TreeGetTypeID() {
	return (CFTypeID)CFTreeGetTypeID(
	);
}
void TreeInsertSibling(void * tree, void * newSibling) {
	return (void)CFTreeInsertSibling(
		// *typing.RefType
		(CFTreeRef)tree,
		// *typing.RefType
		(CFTreeRef)newSibling
	);
}
void * WriteStreamCreateWithAllocatedBuffers(void * alloc, void * bufferAllocator) {
	return (void *)CFWriteStreamCreateWithAllocatedBuffers(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFAllocatorRef)bufferAllocator
	);
}
void CalendarSetTimeZone(void * calendar, void * tz) {
	return (void)CFCalendarSetTimeZone(
		// *typing.RefType
		(CFCalendarRef)calendar,
		// *typing.RefType
		(CFTimeZoneRef)tz
	);
}
bool FileSecuritySetGroupUUID(void * fileSec, void * groupUUID) {
	return (bool)CFFileSecuritySetGroupUUID(
		// *typing.RefType
		(CFFileSecurityRef)fileSec,
		// *typing.RefType
		(CFUUIDRef)groupUUID
	);
}
void * CharacterSetCreateMutable(void * alloc) {
	return (void *)CFCharacterSetCreateMutable(
		// *typing.RefType
		(CFAllocatorRef)alloc
	);
}
bool FileSecuritySetGroup(void * fileSec, gid_t group) {
	return (bool)CFFileSecuritySetGroup(
		// *typing.RefType
		(CFFileSecurityRef)fileSec,
		// *typing.AliasType
		(gid_t)group
	);
}
CFStreamStatus WriteStreamGetStatus(void * stream) {
	return (CFStreamStatus)CFWriteStreamGetStatus(
		// *typing.RefType
		(CFWriteStreamRef)stream
	);
}
void * AttributedStringCreateMutableCopy(void * alloc, CFIndex maxLength, void * aStr) {
	return (void *)CFAttributedStringCreateMutableCopy(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.AliasType
		(CFIndex)maxLength,
		// *typing.RefType
		(CFAttributedStringRef)aStr
	);
}
void * TimeZoneGetName(void * tz) {
	return (void *)CFTimeZoneGetName(
		// *typing.RefType
		(CFTimeZoneRef)tz
	);
}
bool BundleGetPackageInfoInDirectory(void * url, uint32_t* packageType, uint32_t* packageCreator) {
	return (bool)CFBundleGetPackageInfoInDirectory(
		// *typing.RefType
		(CFURLRef)url,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint32_t*)packageType,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint32_t*)packageCreator
	);
}
void ArrayExchangeValuesAtIndices(void * theArray, CFIndex idx1, CFIndex idx2) {
	return (void)CFArrayExchangeValuesAtIndices(
		// *typing.RefType
		(CFMutableArrayRef)theArray,
		// *typing.AliasType
		(CFIndex)idx1,
		// *typing.AliasType
		(CFIndex)idx2
	);
}
CFTypeRef MakeCollectable(CFTypeRef cf) {
	return (CFTypeRef)CFMakeCollectable(
		// *typing.AliasType
		(CFTypeRef)cf
	);
}
void CalendarSetMinimumDaysInFirstWeek(void * calendar, CFIndex mwd) {
	return (void)CFCalendarSetMinimumDaysInFirstWeek(
		// *typing.RefType
		(CFCalendarRef)calendar,
		// *typing.AliasType
		(CFIndex)mwd
	);
}
void CharacterSetIntersect(void * theSet, void * theOtherSet) {
	return (void)CFCharacterSetIntersect(
		// *typing.RefType
		(CFMutableCharacterSetRef)theSet,
		// *typing.RefType
		(CFCharacterSetRef)theOtherSet
	);
}
void * XMLCreateStringByUnescapingEntities(void * allocator, void * string_, void * entitiesDictionary) {
	return (void *)CFXMLCreateStringByUnescapingEntities(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFStringRef)string_,
		// *typing.RefType
		(CFDictionaryRef)entitiesDictionary
	);
}
void * URLCopyPath(void * anURL) {
	return (void *)CFURLCopyPath(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
void SocketEnableCallBacks(void * s, CFOptionFlags callBackTypes) {
	return (void)CFSocketEnableCallBacks(
		// *typing.RefType
		(CFSocketRef)s,
		// *typing.AliasType
		(CFOptionFlags)callBackTypes
	);
}
void ReadStreamClose(void * stream) {
	return (void)CFReadStreamClose(
		// *typing.RefType
		(CFReadStreamRef)stream
	);
}
CFTypeID CalendarGetTypeID() {
	return (CFTypeID)CFCalendarGetTypeID(
	);
}
void * BundleGetIdentifier(void * bundle) {
	return (void *)CFBundleGetIdentifier(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
void * URLCopyPassword(void * anURL) {
	return (void *)CFURLCopyPassword(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
int32_t StringGetIntValue(void * str) {
	return (int32_t)CFStringGetIntValue(
		// *typing.RefType
		(CFStringRef)str
	);
}
bool FileSecurityGetMode(void * fileSec, mode_t* mode) {
	return (bool)CFFileSecurityGetMode(
		// *typing.RefType
		(CFFileSecurityRef)fileSec,
		// *typing.PointerType
		// -> *typing.AliasType
		(mode_t*)mode
	);
}
CFTypeID RunLoopGetTypeID() {
	return (CFTypeID)CFRunLoopGetTypeID(
	);
}
void URLEnumeratorSkipDescendents(void * enumerator) {
	return (void)CFURLEnumeratorSkipDescendents(
		// *typing.RefType
		(CFURLEnumeratorRef)enumerator
	);
}
CFTypeID RunLoopTimerGetTypeID() {
	return (CFTypeID)CFRunLoopTimerGetTypeID(
	);
}
void * URLCopyFileSystemPath(void * anURL, CFURLPathStyle pathStyle) {
	return (void *)CFURLCopyFileSystemPath(
		// *typing.RefType
		(CFURLRef)anURL,
		// *typing.AliasType
		(CFURLPathStyle)pathStyle
	);
}
void ArrayRemoveAllValues(void * theArray) {
	return (void)CFArrayRemoveAllValues(
		// *typing.RefType
		(CFMutableArrayRef)theArray
	);
}
void * MachPortCreateRunLoopSource(void * allocator, void * port, CFIndex order) {
	return (void *)CFMachPortCreateRunLoopSource(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFMachPortRef)port,
		// *typing.AliasType
		(CFIndex)order
	);
}
void * StringGetNameOfEncoding(CFStringEncoding encoding) {
	return (void *)CFStringGetNameOfEncoding(
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
void PlugInRemoveInstanceForFactory(void * factoryID) {
	return (void)CFPlugInRemoveInstanceForFactory(
		// *typing.RefType
		(CFUUIDRef)factoryID
	);
}
CFHashCode Hash(CFTypeRef cf) {
	return (CFHashCode)CFHash(
		// *typing.AliasType
		(CFTypeRef)cf
	);
}
bool URLGetFileSystemRepresentation(void * url, bool resolveAgainstBase, uint8_t* buffer, CFIndex maxBufLen) {
	return (bool)CFURLGetFileSystemRepresentation(
		// *typing.RefType
		(CFURLRef)url,
		// *typing.PrimitiveType
		(Boolean)resolveAgainstBase,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)buffer,
		// *typing.AliasType
		(CFIndex)maxBufLen
	);
}
void * ReadStreamCopyError(void * stream) {
	return (void *)CFReadStreamCopyError(
		// *typing.RefType
		(CFReadStreamRef)stream
	);
}
void TreePrependChild(void * tree, void * newChild) {
	return (void)CFTreePrependChild(
		// *typing.RefType
		(CFTreeRef)tree,
		// *typing.RefType
		(CFTreeRef)newChild
	);
}
void * LocaleCopyISOCurrencyCodes() {
	return (void *)CFLocaleCopyISOCurrencyCodes(
	);
}
void * AllocatorAllocate(void * allocator, CFIndex size, CFOptionFlags hint) {
	return (void *)CFAllocatorAllocate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFIndex)size,
		// *typing.AliasType
		(CFOptionFlags)hint
	);
}
CFIndex DataGetLength(void * theData) {
	return (CFIndex)CFDataGetLength(
		// *typing.RefType
		(CFDataRef)theData
	);
}
bool PreferencesAppValueIsForced(void * key, void * applicationID) {
	return (bool)CFPreferencesAppValueIsForced(
		// *typing.RefType
		(CFStringRef)key,
		// *typing.RefType
		(CFStringRef)applicationID
	);
}
CFStringEncoding StringGetFastestEncoding(void * theString) {
	return (CFStringEncoding)CFStringGetFastestEncoding(
		// *typing.RefType
		(CFStringRef)theString
	);
}
void * StringCreateWithBytes(void * alloc, const uint8_t* bytes, CFIndex numBytes, CFStringEncoding encoding, bool isExternalRepresentation) {
	return (void *)CFStringCreateWithBytes(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)bytes,
		// *typing.AliasType
		(CFIndex)numBytes,
		// *typing.AliasType
		(CFStringEncoding)encoding,
		// *typing.PrimitiveType
		(Boolean)isExternalRepresentation
	);
}
bool StringGetCString(void * theString, char* buffer, CFIndex bufferSize, CFStringEncoding encoding) {
	return (bool)CFStringGetCString(
		// *typing.RefType
		(CFStringRef)theString,
		// *typing.CStringType
		(char*)buffer,
		// *typing.AliasType
		(CFIndex)bufferSize,
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
bool PlugInUnregisterFactory(void * factoryUUID) {
	return (bool)CFPlugInUnregisterFactory(
		// *typing.RefType
		(CFUUIDRef)factoryUUID
	);
}
int32_t UserNotificationCancel(void * userNotification) {
	return (int32_t)CFUserNotificationCancel(
		// *typing.RefType
		(CFUserNotificationRef)userNotification
	);
}
bool StringIsEncodingAvailable(CFStringEncoding encoding) {
	return (bool)CFStringIsEncodingAvailable(
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
bool StringIsSurrogateHighCharacter(UniChar character) {
	return (bool)CFStringIsSurrogateHighCharacter(
		// *typing.AliasType
		(UniChar)character
	);
}
void * ErrorCopyFailureReason(void * err) {
	return (void *)CFErrorCopyFailureReason(
		// *typing.RefType
		(CFErrorRef)err
	);
}
void * DictionaryCreateCopy(void * allocator, void * theDict) {
	return (void *)CFDictionaryCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFDictionaryRef)theDict
	);
}
void * URLCopyScheme(void * anURL) {
	return (void *)CFURLCopyScheme(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
CFTypeID NumberGetTypeID() {
	return (CFTypeID)CFNumberGetTypeID(
	);
}
bool CalendarGetTimeRangeOfUnit(void * calendar, CFCalendarUnit unit, CFAbsoluteTime at, CFAbsoluteTime* startp, CFTimeInterval* tip) {
	return (bool)CFCalendarGetTimeRangeOfUnit(
		// *typing.RefType
		(CFCalendarRef)calendar,
		// *typing.AliasType
		(CFCalendarUnit)unit,
		// *typing.AliasType
		(CFAbsoluteTime)at,
		// *typing.PointerType
		// -> *typing.AliasType
		(CFAbsoluteTime*)startp,
		// *typing.PointerType
		// -> *typing.AliasType
		(CFTimeInterval*)tip
	);
}
void PreferencesSetAppValue(void * key, CFPropertyListRef value, void * applicationID) {
	return (void)CFPreferencesSetAppValue(
		// *typing.RefType
		(CFStringRef)key,
		// *typing.AliasType
		(CFPropertyListRef)value,
		// *typing.RefType
		(CFStringRef)applicationID
	);
}
void ArrayRemoveValueAtIndex(void * theArray, CFIndex idx) {
	return (void)CFArrayRemoveValueAtIndex(
		// *typing.RefType
		(CFMutableArrayRef)theArray,
		// *typing.AliasType
		(CFIndex)idx
	);
}
uint64_t SwapInt64(uint64_t arg) {
	return (uint64_t)CFSwapInt64(
		// *typing.PrimitiveType
		(uint64_t)arg
	);
}
CFIndex CalendarGetFirstWeekday(void * calendar) {
	return (CFIndex)CFCalendarGetFirstWeekday(
		// *typing.RefType
		(CFCalendarRef)calendar
	);
}
void * BundleGetAllBundles() {
	return (void *)CFBundleGetAllBundles(
	);
}
bool URLHasDirectoryPath(void * anURL) {
	return (bool)CFURLHasDirectoryPath(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
void * BitVectorCreate(void * allocator, const uint8_t* bytes, CFIndex numBits) {
	return (void *)CFBitVectorCreate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)bytes,
		// *typing.AliasType
		(CFIndex)numBits
	);
}
void * NotificationCenterGetDistributedCenter() {
	return (void *)CFNotificationCenterGetDistributedCenter(
	);
}
void StringAppendCharacters(void * theString, const UniChar* chars, CFIndex numChars) {
	return (void)CFStringAppendCharacters(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.PointerType
		// -> *typing.AliasType
		(UniChar*)chars,
		// *typing.AliasType
		(CFIndex)numChars
	);
}
bool PlugInRegisterFactoryFunctionByName(void * factoryUUID, void * plugIn, void * functionName) {
	return (bool)CFPlugInRegisterFactoryFunctionByName(
		// *typing.RefType
		(CFUUIDRef)factoryUUID,
		// *typing.RefType
		(CFPlugInRef)plugIn,
		// *typing.RefType
		(CFStringRef)functionName
	);
}
bool CalendarAddComponents(void * calendar, CFAbsoluteTime* at, CFOptionFlags options, char* componentDesc) {
	return (bool)CFCalendarAddComponents(
		// *typing.RefType
		(CFCalendarRef)calendar,
		// *typing.PointerType
		// -> *typing.AliasType
		(CFAbsoluteTime*)at,
		// *typing.AliasType
		(CFOptionFlags)options,
		// *typing.CStringType
		(char*)componentDesc
	);
}
CFNumberType NumberGetType(void * number) {
	return (CFNumberType)CFNumberGetType(
		// *typing.RefType
		(CFNumberRef)number
	);
}
void * DateFormatterGetFormat(void * formatter) {
	return (void *)CFDateFormatterGetFormat(
		// *typing.RefType
		(CFDateFormatterRef)formatter
	);
}
void * BundleCopyBundleLocalizations(void * bundle) {
	return (void *)CFBundleCopyBundleLocalizations(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
void * BundleGetMainBundle() {
	return (void *)CFBundleGetMainBundle(
	);
}
bool RunLoopObserverIsValid(void * observer) {
	return (bool)CFRunLoopObserverIsValid(
		// *typing.RefType
		(CFRunLoopObserverRef)observer
	);
}
void AttributedStringEndEditing(void * aStr) {
	return (void)CFAttributedStringEndEditing(
		// *typing.RefType
		(CFMutableAttributedStringRef)aStr
	);
}
bool CharacterSetIsCharacterMember(void * theSet, UniChar theChar) {
	return (bool)CFCharacterSetIsCharacterMember(
		// *typing.RefType
		(CFCharacterSetRef)theSet,
		// *typing.AliasType
		(UniChar)theChar
	);
}
void * URLCopyAbsoluteURL(void * relativeURL) {
	return (void *)CFURLCopyAbsoluteURL(
		// *typing.RefType
		(CFURLRef)relativeURL
	);
}
void * BundleCopySharedSupportURL(void * bundle) {
	return (void *)CFBundleCopySharedSupportURL(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
void StringTrim(void * theString, void * trimString) {
	return (void)CFStringTrim(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.RefType
		(CFStringRef)trimString
	);
}
CFOptionFlags UserNotificationSecureTextField(CFIndex i) {
	return (CFOptionFlags)CFUserNotificationSecureTextField(
		// *typing.AliasType
		(CFIndex)i
	);
}
CFOptionFlags SocketGetSocketFlags(void * s) {
	return (CFOptionFlags)CFSocketGetSocketFlags(
		// *typing.RefType
		(CFSocketRef)s
	);
}
CFTypeID StringGetTypeID() {
	return (CFTypeID)CFStringGetTypeID(
	);
}
uint32_t StringConvertEncodingToWindowsCodepage(CFStringEncoding encoding) {
	return (uint32_t)CFStringConvertEncodingToWindowsCodepage(
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
void * LocaleCopyCommonISOCurrencyCodes() {
	return (void *)CFLocaleCopyCommonISOCurrencyCodes(
	);
}
void * URLCreateWithBytes(void * allocator, const uint8_t* URLBytes, CFIndex length, CFStringEncoding encoding, void * baseURL) {
	return (void *)CFURLCreateWithBytes(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)URLBytes,
		// *typing.AliasType
		(CFIndex)length,
		// *typing.AliasType
		(CFStringEncoding)encoding,
		// *typing.RefType
		(CFURLRef)baseURL
	);
}
CFTypeID URLEnumeratorGetTypeID() {
	return (CFTypeID)CFURLEnumeratorGetTypeID(
	);
}
CFStringTokenizerTokenType StringTokenizerAdvanceToNextToken(void * tokenizer) {
	return (CFStringTokenizerTokenType)CFStringTokenizerAdvanceToNextToken(
		// *typing.RefType
		(CFStringTokenizerRef)tokenizer
	);
}
void * StringCreateMutableCopy(void * alloc, CFIndex maxLength, void * theString) {
	return (void *)CFStringCreateMutableCopy(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.AliasType
		(CFIndex)maxLength,
		// *typing.RefType
		(CFStringRef)theString
	);
}
void RunLoopTimerSetNextFireDate(void * timer, CFAbsoluteTime fireDate) {
	return (void)CFRunLoopTimerSetNextFireDate(
		// *typing.RefType
		(CFRunLoopTimerRef)timer,
		// *typing.AliasType
		(CFAbsoluteTime)fireDate
	);
}
void * TimeZoneCreateWithName(void * allocator, void * name, bool tryAbbrev) {
	return (void *)CFTimeZoneCreateWithName(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFStringRef)name,
		// *typing.PrimitiveType
		(Boolean)tryAbbrev
	);
}
void MessagePortInvalidate(void * ms) {
	return (void)CFMessagePortInvalidate(
		// *typing.RefType
		(CFMessagePortRef)ms
	);
}
bool CalendarDecomposeAbsoluteTime(void * calendar, CFAbsoluteTime at, char* componentDesc) {
	return (bool)CFCalendarDecomposeAbsoluteTime(
		// *typing.RefType
		(CFCalendarRef)calendar,
		// *typing.AliasType
		(CFAbsoluteTime)at,
		// *typing.CStringType
		(char*)componentDesc
	);
}
CFTypeRef Autorelease(CFTypeRef arg) {
	return (CFTypeRef)CFAutorelease(
		// *typing.AliasType
		(CFTypeRef)arg
	);
}
void * BundleCopyResourceURLsOfTypeInDirectory(void * bundleURL, void * resourceType, void * subDirName) {
	return (void *)CFBundleCopyResourceURLsOfTypeInDirectory(
		// *typing.RefType
		(CFURLRef)bundleURL,
		// *typing.RefType
		(CFStringRef)resourceType,
		// *typing.RefType
		(CFStringRef)subDirName
	);
}
void RunLoopTimerSetTolerance(void * timer, CFTimeInterval tolerance) {
	return (void)CFRunLoopTimerSetTolerance(
		// *typing.RefType
		(CFRunLoopTimerRef)timer,
		// *typing.AliasType
		(CFTimeInterval)tolerance
	);
}
void StringSetExternalCharactersNoCopy(void * theString, UniChar* chars, CFIndex length, CFIndex capacity) {
	return (void)CFStringSetExternalCharactersNoCopy(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.PointerType
		// -> *typing.AliasType
		(UniChar*)chars,
		// *typing.AliasType
		(CFIndex)length,
		// *typing.AliasType
		(CFIndex)capacity
	);
}
void * BundleGetBundleWithIdentifier(void * bundleID) {
	return (void *)CFBundleGetBundleWithIdentifier(
		// *typing.RefType
		(CFStringRef)bundleID
	);
}
void * StringCreateMutable(void * alloc, CFIndex maxLength) {
	return (void *)CFStringCreateMutable(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.AliasType
		(CFIndex)maxLength
	);
}
void * URLCreateWithFileSystemPathRelativeToBase(void * allocator, void * filePath, CFURLPathStyle pathStyle, bool isDirectory, void * baseURL) {
	return (void *)CFURLCreateWithFileSystemPathRelativeToBase(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFStringRef)filePath,
		// *typing.AliasType
		(CFURLPathStyle)pathStyle,
		// *typing.PrimitiveType
		(Boolean)isDirectory,
		// *typing.RefType
		(CFURLRef)baseURL
	);
}
void * URLCreateWithString(void * allocator, void * URLString, void * baseURL) {
	return (void *)CFURLCreateWithString(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFStringRef)URLString,
		// *typing.RefType
		(CFURLRef)baseURL
	);
}
void * BagCreateMutableCopy(void * allocator, CFIndex capacity, void * theBag) {
	return (void *)CFBagCreateMutableCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFIndex)capacity,
		// *typing.RefType
		(CFBagRef)theBag
	);
}
CFTypeRef BundleGetValueForInfoDictionaryKey(void * bundle, void * key) {
	return (CFTypeRef)CFBundleGetValueForInfoDictionaryKey(
		// *typing.RefType
		(CFBundleRef)bundle,
		// *typing.RefType
		(CFStringRef)key
	);
}
void BitVectorSetCount(void * bv, CFIndex count) {
	return (void)CFBitVectorSetCount(
		// *typing.RefType
		(CFMutableBitVectorRef)bv,
		// *typing.AliasType
		(CFIndex)count
	);
}
void * URLGetBaseURL(void * anURL) {
	return (void *)CFURLGetBaseURL(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
CFTypeID BagGetTypeID() {
	return (CFTypeID)CFBagGetTypeID(
	);
}
bool BooleanGetValue(void * boolean) {
	return (bool)CFBooleanGetValue(
		// *typing.RefType
		(CFBooleanRef)boolean
	);
}
void * ErrorCopyDescription(void * err) {
	return (void *)CFErrorCopyDescription(
		// *typing.RefType
		(CFErrorRef)err
	);
}
void * PlugInInstanceGetInstanceData(void * instance) {
	return (void *)CFPlugInInstanceGetInstanceData(
		// *typing.RefType
		(CFPlugInInstanceRef)instance
	);
}
CFNumberFormatterStyle NumberFormatterGetStyle(void * formatter) {
	return (CFNumberFormatterStyle)CFNumberFormatterGetStyle(
		// *typing.RefType
		(CFNumberFormatterRef)formatter
	);
}
void StringCapitalize(void * theString, void * locale) {
	return (void)CFStringCapitalize(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.RefType
		(CFLocaleRef)locale
	);
}
const uint8_t* DataGetBytePtr(void * theData) {
	return (const uint8_t*)CFDataGetBytePtr(
		// *typing.RefType
		(CFDataRef)theData
	);
}
bool FileSecuritySetMode(void * fileSec, mode_t mode) {
	return (bool)CFFileSecuritySetMode(
		// *typing.RefType
		(CFFileSecurityRef)fileSec,
		// *typing.AliasType
		(mode_t)mode
	);
}
CFOptionFlags UserNotificationCheckBoxChecked(CFIndex i) {
	return (CFOptionFlags)CFUserNotificationCheckBoxChecked(
		// *typing.AliasType
		(CFIndex)i
	);
}
void * URLCopyLastPathComponent(void * url) {
	return (void *)CFURLCopyLastPathComponent(
		// *typing.RefType
		(CFURLRef)url
	);
}
CFIndex StringGetMaximumSizeOfFileSystemRepresentation(void * string_) {
	return (CFIndex)CFStringGetMaximumSizeOfFileSystemRepresentation(
		// *typing.RefType
		(CFStringRef)string_
	);
}
void * BundleCopyResourcesDirectoryURL(void * bundle) {
	return (void *)CFBundleCopyResourcesDirectoryURL(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
CFTypeID BooleanGetTypeID() {
	return (CFTypeID)CFBooleanGetTypeID(
	);
}
void * AttributedStringGetMutableString(void * aStr) {
	return (void *)CFAttributedStringGetMutableString(
		// *typing.RefType
		(CFMutableAttributedStringRef)aStr
	);
}
void URLClearResourcePropertyCacheForKey(void * url, void * key) {
	return (void)CFURLClearResourcePropertyCacheForKey(
		// *typing.RefType
		(CFURLRef)url,
		// *typing.RefType
		(CFStringRef)key
	);
}
void * URLCopyFragment(void * anURL, void * charactersToLeaveEscaped) {
	return (void *)CFURLCopyFragment(
		// *typing.RefType
		(CFURLRef)anURL,
		// *typing.RefType
		(CFStringRef)charactersToLeaveEscaped
	);
}
void SocketInvalidate(void * s) {
	return (void)CFSocketInvalidate(
		// *typing.RefType
		(CFSocketRef)s
	);
}
void * UUIDCreateWithBytes(void * alloc, uint8_t byte0, uint8_t byte1, uint8_t byte2, uint8_t byte3, uint8_t byte4, uint8_t byte5, uint8_t byte6, uint8_t byte7, uint8_t byte8, uint8_t byte9, uint8_t byte10, uint8_t byte11, uint8_t byte12, uint8_t byte13, uint8_t byte14, uint8_t byte15) {
	return (void *)CFUUIDCreateWithBytes(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.PrimitiveType
		(uint8_t)byte0,
		// *typing.PrimitiveType
		(uint8_t)byte1,
		// *typing.PrimitiveType
		(uint8_t)byte2,
		// *typing.PrimitiveType
		(uint8_t)byte3,
		// *typing.PrimitiveType
		(uint8_t)byte4,
		// *typing.PrimitiveType
		(uint8_t)byte5,
		// *typing.PrimitiveType
		(uint8_t)byte6,
		// *typing.PrimitiveType
		(uint8_t)byte7,
		// *typing.PrimitiveType
		(uint8_t)byte8,
		// *typing.PrimitiveType
		(uint8_t)byte9,
		// *typing.PrimitiveType
		(uint8_t)byte10,
		// *typing.PrimitiveType
		(uint8_t)byte11,
		// *typing.PrimitiveType
		(uint8_t)byte12,
		// *typing.PrimitiveType
		(uint8_t)byte13,
		// *typing.PrimitiveType
		(uint8_t)byte14,
		// *typing.PrimitiveType
		(uint8_t)byte15
	);
}
void * UserNotificationCreate(void * allocator, CFTimeInterval timeout, CFOptionFlags flags, int32_t* error, void * dictionary) {
	return (void *)CFUserNotificationCreate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFTimeInterval)timeout,
		// *typing.AliasType
		(CFOptionFlags)flags,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(int32_t*)error,
		// *typing.RefType
		(CFDictionaryRef)dictionary
	);
}
void ShowStr(void * str) {
	return (void)CFShowStr(
		// *typing.RefType
		(CFStringRef)str
	);
}
void * WriteStreamCreateWithBuffer(void * alloc, uint8_t* buffer, CFIndex bufferCapacity) {
	return (void *)CFWriteStreamCreateWithBuffer(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)buffer,
		// *typing.AliasType
		(CFIndex)bufferCapacity
	);
}
void * SetCreateMutableCopy(void * allocator, CFIndex capacity, void * theSet) {
	return (void *)CFSetCreateMutableCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFIndex)capacity,
		// *typing.RefType
		(CFSetRef)theSet
	);
}
void AllocatorSetDefault(void * allocator) {
	return (void)CFAllocatorSetDefault(
		// *typing.RefType
		(CFAllocatorRef)allocator
	);
}
bool WriteStreamCanAcceptBytes(void * stream) {
	return (bool)CFWriteStreamCanAcceptBytes(
		// *typing.RefType
		(CFWriteStreamRef)stream
	);
}
CFTypeID SetGetTypeID() {
	return (CFTypeID)CFSetGetTypeID(
	);
}
CFTypeID StringTokenizerGetTypeID() {
	return (CFTypeID)CFStringTokenizerGetTypeID(
	);
}
void * TimeZoneCopyLocalizedName(void * tz, CFTimeZoneNameStyle style, void * locale) {
	return (void *)CFTimeZoneCopyLocalizedName(
		// *typing.RefType
		(CFTimeZoneRef)tz,
		// *typing.AliasType
		(CFTimeZoneNameStyle)style,
		// *typing.RefType
		(CFLocaleRef)locale
	);
}
CFPropertyListRef PreferencesCopyAppValue(void * key, void * applicationID) {
	return (CFPropertyListRef)CFPreferencesCopyAppValue(
		// *typing.RefType
		(CFStringRef)key,
		// *typing.RefType
		(CFStringRef)applicationID
	);
}
CFTypeID AttributedStringGetTypeID() {
	return (CFTypeID)CFAttributedStringGetTypeID(
	);
}
void * URLCopyNetLocation(void * anURL) {
	return (void *)CFURLCopyNetLocation(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
void PreferencesSetValue(void * key, CFPropertyListRef value, void * applicationID, void * userName, void * hostName) {
	return (void)CFPreferencesSetValue(
		// *typing.RefType
		(CFStringRef)key,
		// *typing.AliasType
		(CFPropertyListRef)value,
		// *typing.RefType
		(CFStringRef)applicationID,
		// *typing.RefType
		(CFStringRef)userName,
		// *typing.RefType
		(CFStringRef)hostName
	);
}
void FileDescriptorDisableCallBacks(void * f, CFOptionFlags callBackTypes) {
	return (void)CFFileDescriptorDisableCallBacks(
		// *typing.RefType
		(CFFileDescriptorRef)f,
		// *typing.AliasType
		(CFOptionFlags)callBackTypes
	);
}
void * DateFormatterCreateStringWithAbsoluteTime(void * allocator, void * formatter, CFAbsoluteTime at) {
	return (void *)CFDateFormatterCreateStringWithAbsoluteTime(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFDateFormatterRef)formatter,
		// *typing.AliasType
		(CFAbsoluteTime)at
	);
}
bool PlugInIsLoadOnDemand(void * plugIn) {
	return (bool)CFPlugInIsLoadOnDemand(
		// *typing.RefType
		(CFPlugInRef)plugIn
	);
}
CFIndex RunLoopTimerGetOrder(void * timer) {
	return (CFIndex)CFRunLoopTimerGetOrder(
		// *typing.RefType
		(CFRunLoopTimerRef)timer
	);
}
void StringReplaceAll(void * theString, void * replacement) {
	return (void)CFStringReplaceAll(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.RefType
		(CFStringRef)replacement
	);
}
void * TimeZoneCopyDefault() {
	return (void *)CFTimeZoneCopyDefault(
	);
}
uint32_t SwapInt32BigToHost(uint32_t arg) {
	return (uint32_t)CFSwapInt32BigToHost(
		// *typing.PrimitiveType
		(uint32_t)arg
	);
}
CFTimeInterval TimeZoneGetSecondsFromGMT(void * tz, CFAbsoluteTime at) {
	return (CFTimeInterval)CFTimeZoneGetSecondsFromGMT(
		// *typing.RefType
		(CFTimeZoneRef)tz,
		// *typing.AliasType
		(CFAbsoluteTime)at
	);
}
void * ReadStreamCreateWithFile(void * alloc, void * fileURL) {
	return (void *)CFReadStreamCreateWithFile(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFURLRef)fileURL
	);
}
void DataAppendBytes(void * theData, const uint8_t* bytes, CFIndex length) {
	return (void)CFDataAppendBytes(
		// *typing.RefType
		(CFMutableDataRef)theData,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)bytes,
		// *typing.AliasType
		(CFIndex)length
	);
}
void * LocaleGetSystem() {
	return (void *)CFLocaleGetSystem(
	);
}
void * StringCreateWithCharactersNoCopy(void * alloc, const UniChar* chars, CFIndex numChars, void * contentsDeallocator) {
	return (void *)CFStringCreateWithCharactersNoCopy(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.PointerType
		// -> *typing.AliasType
		(UniChar*)chars,
		// *typing.AliasType
		(CFIndex)numChars,
		// *typing.RefType
		(CFAllocatorRef)contentsDeallocator
	);
}
uint64_t SwapInt64LittleToHost(uint64_t arg) {
	return (uint64_t)CFSwapInt64LittleToHost(
		// *typing.PrimitiveType
		(uint64_t)arg
	);
}
CFIndex ArrayGetCount(void * theArray) {
	return (CFIndex)CFArrayGetCount(
		// *typing.RefType
		(CFArrayRef)theArray
	);
}
void * TreeGetParent(void * tree) {
	return (void *)CFTreeGetParent(
		// *typing.RefType
		(CFTreeRef)tree
	);
}
uint64_t SwapInt64HostToBig(uint64_t arg) {
	return (uint64_t)CFSwapInt64HostToBig(
		// *typing.PrimitiveType
		(uint64_t)arg
	);
}
void RunLoopRun() {
	return (void)CFRunLoopRun(
	);
}
void * RunLoopGetMain() {
	return (void *)CFRunLoopGetMain(
	);
}
void * NumberFormatterCreateStringWithNumber(void * allocator, void * formatter, void * number) {
	return (void *)CFNumberFormatterCreateStringWithNumber(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFNumberFormatterRef)formatter,
		// *typing.RefType
		(CFNumberRef)number
	);
}
CFIndex URLEnumeratorGetDescendentLevel(void * enumerator) {
	return (CFIndex)CFURLEnumeratorGetDescendentLevel(
		// *typing.RefType
		(CFURLEnumeratorRef)enumerator
	);
}
CFTypeID SocketGetTypeID() {
	return (CFTypeID)CFSocketGetTypeID(
	);
}
bool MessagePortIsRemote(void * ms) {
	return (bool)CFMessagePortIsRemote(
		// *typing.RefType
		(CFMessagePortRef)ms
	);
}
void MachPortInvalidate(void * port) {
	return (void)CFMachPortInvalidate(
		// *typing.RefType
		(CFMachPortRef)port
	);
}
void CharacterSetAddCharactersInString(void * theSet, void * theString) {
	return (void)CFCharacterSetAddCharactersInString(
		// *typing.RefType
		(CFMutableCharacterSetRef)theSet,
		// *typing.RefType
		(CFStringRef)theString
	);
}
void * LocaleCopyISOLanguageCodes() {
	return (void *)CFLocaleCopyISOLanguageCodes(
	);
}
CFSocketError SocketSendData(void * s, void * address, void * data, CFTimeInterval timeout) {
	return (CFSocketError)CFSocketSendData(
		// *typing.RefType
		(CFSocketRef)s,
		// *typing.RefType
		(CFDataRef)address,
		// *typing.RefType
		(CFDataRef)data,
		// *typing.AliasType
		(CFTimeInterval)timeout
	);
}
CFSocketNativeHandle SocketGetNative(void * s) {
	return (CFSocketNativeHandle)CFSocketGetNative(
		// *typing.RefType
		(CFSocketRef)s
	);
}
void SocketSetDefaultNameRegistryPortNumber(uint16_t port) {
	return (void)CFSocketSetDefaultNameRegistryPortNumber(
		// *typing.PrimitiveType
		(uint16_t)port
	);
}
CFStringEncoding StringConvertIANACharSetNameToEncoding(void * theString) {
	return (CFStringEncoding)CFStringConvertIANACharSetNameToEncoding(
		// *typing.RefType
		(CFStringRef)theString
	);
}
void * StringCreateFromExternalRepresentation(void * alloc, void * data, CFStringEncoding encoding) {
	return (void *)CFStringCreateFromExternalRepresentation(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFDataRef)data,
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
bool StringHasSuffix(void * theString, void * suffix) {
	return (bool)CFStringHasSuffix(
		// *typing.RefType
		(CFStringRef)theString,
		// *typing.RefType
		(CFStringRef)suffix
	);
}
bool FileSecurityGetOwner(void * fileSec, uid_t* owner) {
	return (bool)CFFileSecurityGetOwner(
		// *typing.RefType
		(CFFileSecurityRef)fileSec,
		// *typing.PointerType
		// -> *typing.AliasType
		(uid_t*)owner
	);
}
void * PlugInInstanceGetFactoryName(void * instance) {
	return (void *)CFPlugInInstanceGetFactoryName(
		// *typing.RefType
		(CFPlugInInstanceRef)instance
	);
}
void * BundleGetDevelopmentRegion(void * bundle) {
	return (void *)CFBundleGetDevelopmentRegion(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
void DataSetLength(void * theData, CFIndex length) {
	return (void)CFDataSetLength(
		// *typing.RefType
		(CFMutableDataRef)theData,
		// *typing.AliasType
		(CFIndex)length
	);
}
void * BundleCopyLocalizationsForPreferences(void * locArray, void * prefArray) {
	return (void *)CFBundleCopyLocalizationsForPreferences(
		// *typing.RefType
		(CFArrayRef)locArray,
		// *typing.RefType
		(CFArrayRef)prefArray
	);
}
void * URLCreateData(void * allocator, void * url, CFStringEncoding encoding, bool escapeWhitespace) {
	return (void *)CFURLCreateData(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFURLRef)url,
		// *typing.AliasType
		(CFStringEncoding)encoding,
		// *typing.PrimitiveType
		(Boolean)escapeWhitespace
	);
}
void * LocaleCreateCopy(void * allocator, void * locale) {
	return (void *)CFLocaleCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFLocaleRef)locale
	);
}
CFLocaleLanguageDirection LocaleGetLanguageLineDirection(void * isoLangCode) {
	return (CFLocaleLanguageDirection)CFLocaleGetLanguageLineDirection(
		// *typing.RefType
		(CFStringRef)isoLangCode
	);
}
void * DateFormatterCreate(void * allocator, void * locale, CFDateFormatterStyle dateStyle, CFDateFormatterStyle timeStyle) {
	return (void *)CFDateFormatterCreate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFLocaleRef)locale,
		// *typing.AliasType
		(CFDateFormatterStyle)dateStyle,
		// *typing.AliasType
		(CFDateFormatterStyle)timeStyle
	);
}
void * PlugInFindFactoriesForPlugInType(void * typeUUID) {
	return (void *)CFPlugInFindFactoriesForPlugInType(
		// *typing.RefType
		(CFUUIDRef)typeUUID
	);
}
void * PlugInCreate(void * allocator, void * plugInURL) {
	return (void *)CFPlugInCreate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFURLRef)plugInURL
	);
}
CFAbsoluteTime RunLoopTimerGetNextFireDate(void * timer) {
	return (CFAbsoluteTime)CFRunLoopTimerGetNextFireDate(
		// *typing.RefType
		(CFRunLoopTimerRef)timer
	);
}
bool NumberFormatterGetDecimalInfoForCurrencyCode(void * currencyCode, int32_t* defaultFractionDigits, double* roundingIncrement) {
	return (bool)CFNumberFormatterGetDecimalInfoForCurrencyCode(
		// *typing.RefType
		(CFStringRef)currencyCode,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(int32_t*)defaultFractionDigits,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(double*)roundingIncrement
	);
}
void * CalendarCopyLocale(void * calendar) {
	return (void *)CFCalendarCopyLocale(
		// *typing.RefType
		(CFCalendarRef)calendar
	);
}
CFIndex WriteStreamWrite(void * stream, const uint8_t* buffer, CFIndex bufferLength) {
	return (CFIndex)CFWriteStreamWrite(
		// *typing.RefType
		(CFWriteStreamRef)stream,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)buffer,
		// *typing.AliasType
		(CFIndex)bufferLength
	);
}
void Show(CFTypeRef obj) {
	return (void)CFShow(
		// *typing.AliasType
		(CFTypeRef)obj
	);
}
bool StringIsHyphenationAvailableForLocale(void * locale) {
	return (bool)CFStringIsHyphenationAvailableForLocale(
		// *typing.RefType
		(CFLocaleRef)locale
	);
}
CFTypeRef StringTokenizerCopyCurrentTokenAttribute(void * tokenizer, CFOptionFlags attribute) {
	return (CFTypeRef)CFStringTokenizerCopyCurrentTokenAttribute(
		// *typing.RefType
		(CFStringTokenizerRef)tokenizer,
		// *typing.AliasType
		(CFOptionFlags)attribute
	);
}
void * UUIDCreateFromString(void * alloc, void * uuidStr) {
	return (void *)CFUUIDCreateFromString(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFStringRef)uuidStr
	);
}
CFTypeID ArrayGetTypeID() {
	return (CFTypeID)CFArrayGetTypeID(
	);
}
void RunLoopSourceSignal(void * source) {
	return (void)CFRunLoopSourceSignal(
		// *typing.RefType
		(CFRunLoopSourceRef)source
	);
}
CFSocketError SocketSetAddress(void * s, void * address) {
	return (CFSocketError)CFSocketSetAddress(
		// *typing.RefType
		(CFSocketRef)s,
		// *typing.RefType
		(CFDataRef)address
	);
}
void * StringCreateWithCString(void * alloc, char* cStr, CFStringEncoding encoding) {
	return (void *)CFStringCreateWithCString(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.CStringType
		(char*)cStr,
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
void * DataCreateMutableCopy(void * allocator, CFIndex capacity, void * theData) {
	return (void *)CFDataCreateMutableCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFIndex)capacity,
		// *typing.RefType
		(CFDataRef)theData
	);
}
void * AttributedStringCreateMutable(void * alloc, CFIndex maxLength) {
	return (void *)CFAttributedStringCreateMutable(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.AliasType
		(CFIndex)maxLength
	);
}
void * MessagePortGetName(void * ms) {
	return (void *)CFMessagePortGetName(
		// *typing.RefType
		(CFMessagePortRef)ms
	);
}
bool FileSecuritySetOwnerUUID(void * fileSec, void * ownerUUID) {
	return (bool)CFFileSecuritySetOwnerUUID(
		// *typing.RefType
		(CFFileSecurityRef)fileSec,
		// *typing.RefType
		(CFUUIDRef)ownerUUID
	);
}
CFAbsoluteTime DateGetAbsoluteTime(void * theDate) {
	return (CFAbsoluteTime)CFDateGetAbsoluteTime(
		// *typing.RefType
		(CFDateRef)theDate
	);
}
void CalendarSetLocale(void * calendar, void * locale) {
	return (void)CFCalendarSetLocale(
		// *typing.RefType
		(CFCalendarRef)calendar,
		// *typing.RefType
		(CFLocaleRef)locale
	);
}
void * BundleCopyExecutableArchitectures(void * bundle) {
	return (void *)CFBundleCopyExecutableArchitectures(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
void RunLoopObserverInvalidate(void * observer) {
	return (void)CFRunLoopObserverInvalidate(
		// *typing.RefType
		(CFRunLoopObserverRef)observer
	);
}
void PreferencesSetMultiple(void * keysToSet, void * keysToRemove, void * applicationID, void * userName, void * hostName) {
	return (void)CFPreferencesSetMultiple(
		// *typing.RefType
		(CFDictionaryRef)keysToSet,
		// *typing.RefType
		(CFArrayRef)keysToRemove,
		// *typing.RefType
		(CFStringRef)applicationID,
		// *typing.RefType
		(CFStringRef)userName,
		// *typing.RefType
		(CFStringRef)hostName
	);
}
CFIndex StringGetMaximumSizeForEncoding(CFIndex length, CFStringEncoding encoding) {
	return (CFIndex)CFStringGetMaximumSizeForEncoding(
		// *typing.AliasType
		(CFIndex)length,
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
bool RunLoopTimerDoesRepeat(void * timer) {
	return (bool)CFRunLoopTimerDoesRepeat(
		// *typing.RefType
		(CFRunLoopTimerRef)timer
	);
}
void * LocaleCopyAvailableLocaleIdentifiers() {
	return (void *)CFLocaleCopyAvailableLocaleIdentifiers(
	);
}
void * BundleCopySharedFrameworksURL(void * bundle) {
	return (void *)CFBundleCopySharedFrameworksURL(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
uint16_t SwapInt16(uint16_t arg) {
	return (uint16_t)CFSwapInt16(
		// *typing.PrimitiveType
		(uint16_t)arg
	);
}
void * BundleCreate(void * allocator, void * bundleURL) {
	return (void *)CFBundleCreate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFURLRef)bundleURL
	);
}
void * URLCreateStringByReplacingPercentEscapes(void * allocator, void * originalString, void * charactersToLeaveEscaped) {
	return (void *)CFURLCreateStringByReplacingPercentEscapes(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFStringRef)originalString,
		// *typing.RefType
		(CFStringRef)charactersToLeaveEscaped
	);
}
CFTimeInterval TimeZoneGetDaylightSavingTimeOffset(void * tz, CFAbsoluteTime at) {
	return (CFTimeInterval)CFTimeZoneGetDaylightSavingTimeOffset(
		// *typing.RefType
		(CFTimeZoneRef)tz,
		// *typing.AliasType
		(CFAbsoluteTime)at
	);
}
void BitVectorSetAllBits(void * bv, CFBit value) {
	return (void)CFBitVectorSetAllBits(
		// *typing.RefType
		(CFMutableBitVectorRef)bv,
		// *typing.AliasType
		(CFBit)value
	);
}
void * XMLCreateStringByEscapingEntities(void * allocator, void * string_, void * entitiesDictionary) {
	return (void *)CFXMLCreateStringByEscapingEntities(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFStringRef)string_,
		// *typing.RefType
		(CFDictionaryRef)entitiesDictionary
	);
}
CFTypeID GetTypeID(CFTypeRef cf) {
	return (CFTypeID)CFGetTypeID(
		// *typing.AliasType
		(CFTypeRef)cf
	);
}
CFIndex PreferencesGetAppIntegerValue(void * key, void * applicationID, bool* keyExistsAndHasValidFormat) {
	return (CFIndex)CFPreferencesGetAppIntegerValue(
		// *typing.RefType
		(CFStringRef)key,
		// *typing.RefType
		(CFStringRef)applicationID,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(Boolean*)keyExistsAndHasValidFormat
	);
}
void * StringCreateWithCStringNoCopy(void * alloc, char* cStr, CFStringEncoding encoding, void * contentsDeallocator) {
	return (void *)CFStringCreateWithCStringNoCopy(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.CStringType
		(char*)cStr,
		// *typing.AliasType
		(CFStringEncoding)encoding,
		// *typing.RefType
		(CFAllocatorRef)contentsDeallocator
	);
}
CFTypeID ReadStreamGetTypeID() {
	return (CFTypeID)CFReadStreamGetTypeID(
	);
}
void * DateCreate(void * allocator, CFAbsoluteTime at) {
	return (void *)CFDateCreate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFAbsoluteTime)at
	);
}
void StringLowercase(void * theString, void * locale) {
	return (void)CFStringLowercase(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.RefType
		(CFLocaleRef)locale
	);
}
CFStringEncoding StringConvertNSStringEncodingToEncoding(int32_t encoding) {
	return (CFStringEncoding)CFStringConvertNSStringEncodingToEncoding(
		// *typing.PrimitiveType
		(int32_t)encoding
	);
}
CFStringEncoding StringGetSmallestEncoding(void * theString) {
	return (CFStringEncoding)CFStringGetSmallestEncoding(
		// *typing.RefType
		(CFStringRef)theString
	);
}
bool StringHasPrefix(void * theString, void * prefix) {
	return (bool)CFStringHasPrefix(
		// *typing.RefType
		(CFStringRef)theString,
		// *typing.RefType
		(CFStringRef)prefix
	);
}
void * DateFormatterCreateDateFormatFromTemplate(void * allocator, void * tmplate, CFOptionFlags options, void * locale) {
	return (void *)CFDateFormatterCreateDateFormatFromTemplate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFStringRef)tmplate,
		// *typing.AliasType
		(CFOptionFlags)options,
		// *typing.RefType
		(CFLocaleRef)locale
	);
}
bool FileSecurityClearProperties(void * fileSec, CFFileSecurityClearOptions clearPropertyMask) {
	return (bool)CFFileSecurityClearProperties(
		// *typing.RefType
		(CFFileSecurityRef)fileSec,
		// *typing.AliasType
		(CFFileSecurityClearOptions)clearPropertyMask
	);
}
uint8_t* DataGetMutableBytePtr(void * theData) {
	return (uint8_t*)CFDataGetMutableBytePtr(
		// *typing.RefType
		(CFMutableDataRef)theData
	);
}
CFIndex BitVectorGetCount(void * bv) {
	return (CFIndex)CFBitVectorGetCount(
		// *typing.RefType
		(CFBitVectorRef)bv
	);
}
void RunLoopStop(void * rl) {
	return (void)CFRunLoopStop(
		// *typing.RefType
		(CFRunLoopRef)rl
	);
}
bool CharacterSetHasMemberInPlane(void * theSet, CFIndex thePlane) {
	return (bool)CFCharacterSetHasMemberInPlane(
		// *typing.RefType
		(CFCharacterSetRef)theSet,
		// *typing.AliasType
		(CFIndex)thePlane
	);
}
void RunLoopWakeUp(void * rl) {
	return (void)CFRunLoopWakeUp(
		// *typing.RefType
		(CFRunLoopRef)rl
	);
}
bool StringGetFileSystemRepresentation(void * string_, char* buffer, CFIndex maxBufLen) {
	return (bool)CFStringGetFileSystemRepresentation(
		// *typing.RefType
		(CFStringRef)string_,
		// *typing.CStringType
		(char*)buffer,
		// *typing.AliasType
		(CFIndex)maxBufLen
	);
}
bool URLCanBeDecomposed(void * anURL) {
	return (bool)CFURLCanBeDecomposed(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
void * UserNotificationGetResponseValue(void * userNotification, void * key, CFIndex idx) {
	return (void *)CFUserNotificationGetResponseValue(
		// *typing.RefType
		(CFUserNotificationRef)userNotification,
		// *typing.RefType
		(CFStringRef)key,
		// *typing.AliasType
		(CFIndex)idx
	);
}
void * ArrayCreateMutableCopy(void * allocator, CFIndex capacity, void * theArray) {
	return (void *)CFArrayCreateMutableCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFIndex)capacity,
		// *typing.RefType
		(CFArrayRef)theArray
	);
}
void * ReadStreamCreateWithBytesNoCopy(void * alloc, const uint8_t* bytes, CFIndex length, void * bytesDeallocator) {
	return (void *)CFReadStreamCreateWithBytesNoCopy(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)bytes,
		// *typing.AliasType
		(CFIndex)length,
		// *typing.RefType
		(CFAllocatorRef)bytesDeallocator
	);
}
int32_t UserNotificationDisplayAlert(CFTimeInterval timeout, CFOptionFlags flags, void * iconURL, void * soundURL, void * localizationURL, void * alertHeader, void * alertMessage, void * defaultButtonTitle, void * alternateButtonTitle, void * otherButtonTitle, CFOptionFlags* responseFlags) {
	return (int32_t)CFUserNotificationDisplayAlert(
		// *typing.AliasType
		(CFTimeInterval)timeout,
		// *typing.AliasType
		(CFOptionFlags)flags,
		// *typing.RefType
		(CFURLRef)iconURL,
		// *typing.RefType
		(CFURLRef)soundURL,
		// *typing.RefType
		(CFURLRef)localizationURL,
		// *typing.RefType
		(CFStringRef)alertHeader,
		// *typing.RefType
		(CFStringRef)alertMessage,
		// *typing.RefType
		(CFStringRef)defaultButtonTitle,
		// *typing.RefType
		(CFStringRef)alternateButtonTitle,
		// *typing.RefType
		(CFStringRef)otherButtonTitle,
		// *typing.PointerType
		// -> *typing.AliasType
		(CFOptionFlags*)responseFlags
	);
}
void * URLCopyHostName(void * anURL) {
	return (void *)CFURLCopyHostName(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
void * ArrayGetValueAtIndex(void * theArray, CFIndex idx) {
	return (void *)CFArrayGetValueAtIndex(
		// *typing.RefType
		(CFArrayRef)theArray,
		// *typing.AliasType
		(CFIndex)idx
	);
}
void PlugInSetLoadOnDemand(void * plugIn, bool flag) {
	return (void)CFPlugInSetLoadOnDemand(
		// *typing.RefType
		(CFPlugInRef)plugIn,
		// *typing.PrimitiveType
		(Boolean)flag
	);
}
void * TimeZoneCreateWithTimeIntervalFromGMT(void * allocator, CFTimeInterval ti) {
	return (void *)CFTimeZoneCreateWithTimeIntervalFromGMT(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFTimeInterval)ti
	);
}
uint32_t BundleGetVersionNumber(void * bundle) {
	return (uint32_t)CFBundleGetVersionNumber(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
void DateFormatterSetProperty(void * formatter, void * key, CFTypeRef value) {
	return (void)CFDateFormatterSetProperty(
		// *typing.RefType
		(CFDateFormatterRef)formatter,
		// *typing.RefType
		(CFStringRef)key,
		// *typing.AliasType
		(CFTypeRef)value
	);
}
CFTypeID MachPortGetTypeID() {
	return (CFTypeID)CFMachPortGetTypeID(
	);
}
void * URLCreateFromFileSystemRepresentation(void * allocator, const uint8_t* buffer, CFIndex bufLen, bool isDirectory) {
	return (void *)CFURLCreateFromFileSystemRepresentation(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)buffer,
		// *typing.AliasType
		(CFIndex)bufLen,
		// *typing.PrimitiveType
		(Boolean)isDirectory
	);
}
void * UUIDCreateString(void * alloc, void * uuid) {
	return (void *)CFUUIDCreateString(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFUUIDRef)uuid
	);
}
CFStringEncoding StringConvertWindowsCodepageToEncoding(uint32_t codepage) {
	return (CFStringEncoding)CFStringConvertWindowsCodepageToEncoding(
		// *typing.PrimitiveType
		(uint32_t)codepage
	);
}
void StringTrimWhitespace(void * theString) {
	return (void)CFStringTrimWhitespace(
		// *typing.RefType
		(CFMutableStringRef)theString
	);
}
bool MachPortIsValid(void * port) {
	return (bool)CFMachPortIsValid(
		// *typing.RefType
		(CFMachPortRef)port
	);
}
CFTypeID BinaryHeapGetTypeID() {
	return (CFTypeID)CFBinaryHeapGetTypeID(
	);
}
void AttributedStringBeginEditing(void * aStr) {
	return (void)CFAttributedStringBeginEditing(
		// *typing.RefType
		(CFMutableAttributedStringRef)aStr
	);
}
bool RunLoopSourceIsValid(void * source) {
	return (bool)CFRunLoopSourceIsValid(
		// *typing.RefType
		(CFRunLoopSourceRef)source
	);
}
void * URLCreateCopyDeletingPathExtension(void * allocator, void * url) {
	return (void *)CFURLCreateCopyDeletingPathExtension(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFURLRef)url
	);
}
CFIndex GetRetainCount(CFTypeRef cf) {
	return (CFIndex)CFGetRetainCount(
		// *typing.AliasType
		(CFTypeRef)cf
	);
}
void * TimeZoneCreate(void * allocator, void * name, void * data) {
	return (void *)CFTimeZoneCreate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFStringRef)name,
		// *typing.RefType
		(CFDataRef)data
	);
}
void * StringCreateWithFileSystemRepresentation(void * alloc, char* buffer) {
	return (void *)CFStringCreateWithFileSystemRepresentation(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.CStringType
		(char*)buffer
	);
}
void * DateFormatterCreateISO8601Formatter(void * allocator, CFISO8601DateFormatOptions formatOptions) {
	return (void *)CFDateFormatterCreateISO8601Formatter(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFISO8601DateFormatOptions)formatOptions
	);
}
CFStringEncoding StringGetSystemEncoding() {
	return (CFStringEncoding)CFStringGetSystemEncoding(
	);
}
char* StringGetCStringPtr(void * theString, CFStringEncoding encoding) {
	return (char*)CFStringGetCStringPtr(
		// *typing.RefType
		(CFStringRef)theString,
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
uint16_t SocketGetDefaultNameRegistryPortNumber() {
	return (uint16_t)CFSocketGetDefaultNameRegistryPortNumber(
	);
}
void * CharacterSetCreateBitmapRepresentation(void * alloc, void * theSet) {
	return (void *)CFCharacterSetCreateBitmapRepresentation(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFCharacterSetRef)theSet
	);
}
CFAbsoluteTime AbsoluteTimeGetCurrent() {
	return (CFAbsoluteTime)CFAbsoluteTimeGetCurrent(
	);
}
bool StringIsSurrogateLowCharacter(UniChar character) {
	return (bool)CFStringIsSurrogateLowCharacter(
		// *typing.AliasType
		(UniChar)character
	);
}
bool BundleIsExecutableLoadableForURL(void * url) {
	return (bool)CFBundleIsExecutableLoadableForURL(
		// *typing.RefType
		(CFURLRef)url
	);
}
void * BundleCreateBundlesFromDirectory(void * allocator, void * directoryURL, void * bundleType) {
	return (void *)CFBundleCreateBundlesFromDirectory(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFURLRef)directoryURL,
		// *typing.RefType
		(CFStringRef)bundleType
	);
}
void SocketDisableCallBacks(void * s, CFOptionFlags callBackTypes) {
	return (void)CFSocketDisableCallBacks(
		// *typing.RefType
		(CFSocketRef)s,
		// *typing.AliasType
		(CFOptionFlags)callBackTypes
	);
}
void * NumberFormatterGetFormat(void * formatter) {
	return (void *)CFNumberFormatterGetFormat(
		// *typing.RefType
		(CFNumberFormatterRef)formatter
	);
}
void * URLCopyResourceSpecifier(void * anURL) {
	return (void *)CFURLCopyResourceSpecifier(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
void * BitVectorCreateMutableCopy(void * allocator, CFIndex capacity, void * bv) {
	return (void *)CFBitVectorCreateMutableCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFIndex)capacity,
		// *typing.RefType
		(CFBitVectorRef)bv
	);
}
void * BundleCopyExecutableURL(void * bundle) {
	return (void *)CFBundleCopyExecutableURL(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
void TimeZoneSetDefault(void * tz) {
	return (void)CFTimeZoneSetDefault(
		// *typing.RefType
		(CFTimeZoneRef)tz
	);
}
void * CharacterSetCreateMutableCopy(void * alloc, void * theSet) {
	return (void *)CFCharacterSetCreateMutableCopy(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFCharacterSetRef)theSet
	);
}
void BundleGetPackageInfo(void * bundle, uint32_t* packageType, uint32_t* packageCreator) {
	return (void)CFBundleGetPackageInfo(
		// *typing.RefType
		(CFBundleRef)bundle,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint32_t*)packageType,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint32_t*)packageCreator
	);
}
CFTypeID FileDescriptorGetTypeID() {
	return (CFTypeID)CFFileDescriptorGetTypeID(
	);
}
void * StringCreateMutableWithExternalCharactersNoCopy(void * alloc, UniChar* chars, CFIndex numChars, CFIndex capacity, void * externalCharactersAllocator) {
	return (void *)CFStringCreateMutableWithExternalCharactersNoCopy(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.PointerType
		// -> *typing.AliasType
		(UniChar*)chars,
		// *typing.AliasType
		(CFIndex)numChars,
		// *typing.AliasType
		(CFIndex)capacity,
		// *typing.RefType
		(CFAllocatorRef)externalCharactersAllocator
	);
}
CFIndex StringGetLength(void * theString) {
	return (CFIndex)CFStringGetLength(
		// *typing.RefType
		(CFStringRef)theString
	);
}
uint32_t SwapInt32HostToLittle(uint32_t arg) {
	return (uint32_t)CFSwapInt32HostToLittle(
		// *typing.PrimitiveType
		(uint32_t)arg
	);
}
void * URLCreateResourcePropertiesForKeysFromBookmarkData(void * allocator, void * resourcePropertiesToReturn, void * bookmark) {
	return (void *)CFURLCreateResourcePropertiesForKeysFromBookmarkData(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFArrayRef)resourcePropertiesToReturn,
		// *typing.RefType
		(CFDataRef)bookmark
	);
}
void * ArrayCreateCopy(void * allocator, void * theArray) {
	return (void *)CFArrayCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFArrayRef)theArray
	);
}
void StringPad(void * theString, void * padString, CFIndex length, CFIndex indexIntoPad) {
	return (void)CFStringPad(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.RefType
		(CFStringRef)padString,
		// *typing.AliasType
		(CFIndex)length,
		// *typing.AliasType
		(CFIndex)indexIntoPad
	);
}
void * CalendarCopyCurrent() {
	return (void *)CFCalendarCopyCurrent(
	);
}
void * URLCopyQueryString(void * anURL, void * charactersToLeaveEscaped) {
	return (void *)CFURLCopyQueryString(
		// *typing.RefType
		(CFURLRef)anURL,
		// *typing.RefType
		(CFStringRef)charactersToLeaveEscaped
	);
}
void * LocaleCopyPreferredLanguages() {
	return (void *)CFLocaleCopyPreferredLanguages(
	);
}
void * NotificationCenterGetDarwinNotifyCenter() {
	return (void *)CFNotificationCenterGetDarwinNotifyCenter(
	);
}
void * DataCreateWithBytesNoCopy(void * allocator, const uint8_t* bytes, CFIndex length, void * bytesDeallocator) {
	return (void *)CFDataCreateWithBytesNoCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)bytes,
		// *typing.AliasType
		(CFIndex)length,
		// *typing.RefType
		(CFAllocatorRef)bytesDeallocator
	);
}
void * BundleCopyBundleURL(void * bundle) {
	return (void *)CFBundleCopyBundleURL(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
void * AttributedStringGetString(void * aStr) {
	return (void *)CFAttributedStringGetString(
		// *typing.RefType
		(CFAttributedStringRef)aStr
	);
}
bool PreferencesSynchronize(void * applicationID, void * userName, void * hostName) {
	return (bool)CFPreferencesSynchronize(
		// *typing.RefType
		(CFStringRef)applicationID,
		// *typing.RefType
		(CFStringRef)userName,
		// *typing.RefType
		(CFStringRef)hostName
	);
}
CFIndex RunLoopSourceGetOrder(void * source) {
	return (CFIndex)CFRunLoopSourceGetOrder(
		// *typing.RefType
		(CFRunLoopSourceRef)source
	);
}
void * TreeGetFirstChild(void * tree) {
	return (void *)CFTreeGetFirstChild(
		// *typing.RefType
		(CFTreeRef)tree
	);
}
void PreferencesRemoveSuitePreferencesFromApp(void * applicationID, void * suiteID) {
	return (void)CFPreferencesRemoveSuitePreferencesFromApp(
		// *typing.RefType
		(CFStringRef)applicationID,
		// *typing.RefType
		(CFStringRef)suiteID
	);
}
CFDateFormatterStyle DateFormatterGetTimeStyle(void * formatter) {
	return (CFDateFormatterStyle)CFDateFormatterGetTimeStyle(
		// *typing.RefType
		(CFDateFormatterRef)formatter
	);
}
bool Equal(CFTypeRef cf1, CFTypeRef cf2) {
	return (bool)CFEqual(
		// *typing.AliasType
		(CFTypeRef)cf1,
		// *typing.AliasType
		(CFTypeRef)cf2
	);
}
uint16_t SwapInt16HostToBig(uint16_t arg) {
	return (uint16_t)CFSwapInt16HostToBig(
		// *typing.PrimitiveType
		(uint16_t)arg
	);
}
uint16_t SwapInt16LittleToHost(uint16_t arg) {
	return (uint16_t)CFSwapInt16LittleToHost(
		// *typing.PrimitiveType
		(uint16_t)arg
	);
}
void * AllocatorGetDefault() {
	return (void *)CFAllocatorGetDefault(
	);
}
bool FileSecuritySetOwner(void * fileSec, uid_t owner) {
	return (bool)CFFileSecuritySetOwner(
		// *typing.RefType
		(CFFileSecurityRef)fileSec,
		// *typing.AliasType
		(uid_t)owner
	);
}
void * BundleGetPlugIn(void * bundle) {
	return (void *)CFBundleGetPlugIn(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
bool NumberIsFloatType(void * number) {
	return (bool)CFNumberIsFloatType(
		// *typing.RefType
		(CFNumberRef)number
	);
}
int32_t UserNotificationDisplayNotice(CFTimeInterval timeout, CFOptionFlags flags, void * iconURL, void * soundURL, void * localizationURL, void * alertHeader, void * alertMessage, void * defaultButtonTitle) {
	return (int32_t)CFUserNotificationDisplayNotice(
		// *typing.AliasType
		(CFTimeInterval)timeout,
		// *typing.AliasType
		(CFOptionFlags)flags,
		// *typing.RefType
		(CFURLRef)iconURL,
		// *typing.RefType
		(CFURLRef)soundURL,
		// *typing.RefType
		(CFURLRef)localizationURL,
		// *typing.RefType
		(CFStringRef)alertHeader,
		// *typing.RefType
		(CFStringRef)alertMessage,
		// *typing.RefType
		(CFStringRef)defaultButtonTitle
	);
}
CFTypeID BundleGetTypeID() {
	return (CFTypeID)CFBundleGetTypeID(
	);
}
CFIndex AllocatorGetPreferredSizeForSize(void * allocator, CFIndex size, CFOptionFlags hint) {
	return (CFIndex)CFAllocatorGetPreferredSizeForSize(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFIndex)size,
		// *typing.AliasType
		(CFOptionFlags)hint
	);
}
void StringNormalize(void * theString, CFStringNormalizationForm theForm) {
	return (void)CFStringNormalize(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.AliasType
		(CFStringNormalizationForm)theForm
	);
}
void SetRemoveAllValues(void * theSet) {
	return (void)CFSetRemoveAllValues(
		// *typing.RefType
		(CFMutableSetRef)theSet
	);
}
CFTypeRef Retain(CFTypeRef cf) {
	return (CFTypeRef)CFRetain(
		// *typing.AliasType
		(CFTypeRef)cf
	);
}
void * BitVectorCreateCopy(void * allocator, void * bv) {
	return (void *)CFBitVectorCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFBitVectorRef)bv
	);
}
void TreeAppendChild(void * tree, void * newChild) {
	return (void)CFTreeAppendChild(
		// *typing.RefType
		(CFTreeRef)tree,
		// *typing.RefType
		(CFTreeRef)newChild
	);
}
void TimeZoneResetSystem() {
	return (void)CFTimeZoneResetSystem(
	);
}
CFTypeID WriteStreamGetTypeID() {
	return (CFTypeID)CFWriteStreamGetTypeID(
	);
}
CFIndex RunLoopObserverGetOrder(void * observer) {
	return (CFIndex)CFRunLoopObserverGetOrder(
		// *typing.RefType
		(CFRunLoopObserverRef)observer
	);
}
void BundleUnloadExecutable(void * bundle) {
	return (void)CFBundleUnloadExecutable(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
bool URLIsFileReferenceURL(void * url) {
	return (bool)CFURLIsFileReferenceURL(
		// *typing.RefType
		(CFURLRef)url
	);
}
void URLStopAccessingSecurityScopedResource(void * url) {
	return (void)CFURLStopAccessingSecurityScopedResource(
		// *typing.RefType
		(CFURLRef)url
	);
}
CFTimeInterval DateGetTimeIntervalSinceDate(void * theDate, void * otherDate) {
	return (CFTimeInterval)CFDateGetTimeIntervalSinceDate(
		// *typing.RefType
		(CFDateRef)theDate,
		// *typing.RefType
		(CFDateRef)otherDate
	);
}
CFFileDescriptorNativeDescriptor FileDescriptorGetNativeDescriptor(void * f) {
	return (CFFileDescriptorNativeDescriptor)CFFileDescriptorGetNativeDescriptor(
		// *typing.RefType
		(CFFileDescriptorRef)f
	);
}
void * FileSecurityCreate(void * allocator) {
	return (void *)CFFileSecurityCreate(
		// *typing.RefType
		(CFAllocatorRef)allocator
	);
}
void * CopyTypeIDDescription(CFTypeID type_id) {
	return (void *)CFCopyTypeIDDescription(
		// *typing.AliasType
		(CFTypeID)type_id
	);
}
void * FileDescriptorCreateRunLoopSource(void * allocator, void * f, CFIndex order) {
	return (void *)CFFileDescriptorCreateRunLoopSource(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFFileDescriptorRef)f,
		// *typing.AliasType
		(CFIndex)order
	);
}
CFStringTokenizerTokenType StringTokenizerGoToTokenAtIndex(void * tokenizer, CFIndex index) {
	return (CFStringTokenizerTokenType)CFStringTokenizerGoToTokenAtIndex(
		// *typing.RefType
		(CFStringTokenizerRef)tokenizer,
		// *typing.AliasType
		(CFIndex)index
	);
}
CFStringEncoding StringGetMostCompatibleMacStringEncoding(CFStringEncoding encoding) {
	return (CFStringEncoding)CFStringGetMostCompatibleMacStringEncoding(
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
void * StringCreateByCombiningStrings(void * alloc, void * theArray, void * separatorString) {
	return (void *)CFStringCreateByCombiningStrings(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFArrayRef)theArray,
		// *typing.RefType
		(CFStringRef)separatorString
	);
}
void * DataCreateMutable(void * allocator, CFIndex capacity) {
	return (void *)CFDataCreateMutable(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFIndex)capacity
	);
}
void * PreferencesCopyKeyList(void * applicationID, void * userName, void * hostName) {
	return (void *)CFPreferencesCopyKeyList(
		// *typing.RefType
		(CFStringRef)applicationID,
		// *typing.RefType
		(CFStringRef)userName,
		// *typing.RefType
		(CFStringRef)hostName
	);
}
void * BundleCopySupportFilesDirectoryURL(void * bundle) {
	return (void *)CFBundleCopySupportFilesDirectoryURL(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
CFTypeID MessagePortGetTypeID() {
	return (CFTypeID)CFMessagePortGetTypeID(
	);
}
void * BundleCopyPrivateFrameworksURL(void * bundle) {
	return (void *)CFBundleCopyPrivateFrameworksURL(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
void * SocketCopyPeerAddress(void * s) {
	return (void *)CFSocketCopyPeerAddress(
		// *typing.RefType
		(CFSocketRef)s
	);
}
bool CharacterSetIsSupersetOfSet(void * theSet, void * theOtherset) {
	return (bool)CFCharacterSetIsSupersetOfSet(
		// *typing.RefType
		(CFCharacterSetRef)theSet,
		// *typing.RefType
		(CFCharacterSetRef)theOtherset
	);
}
void BinaryHeapRemoveMinimumValue(void * heap) {
	return (void)CFBinaryHeapRemoveMinimumValue(
		// *typing.RefType
		(CFBinaryHeapRef)heap
	);
}
void StringAppend(void * theString, void * appendedString) {
	return (void)CFStringAppend(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.RefType
		(CFStringRef)appendedString
	);
}
void * UUIDCreate(void * alloc) {
	return (void *)CFUUIDCreate(
		// *typing.RefType
		(CFAllocatorRef)alloc
	);
}
CFIndex AttributedStringGetLength(void * aStr) {
	return (CFIndex)CFAttributedStringGetLength(
		// *typing.RefType
		(CFAttributedStringRef)aStr
	);
}
void * BundleGetLocalInfoDictionary(void * bundle) {
	return (void *)CFBundleGetLocalInfoDictionary(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
void * ErrorCopyUserInfo(void * err) {
	return (void *)CFErrorCopyUserInfo(
		// *typing.RefType
		(CFErrorRef)err
	);
}
void * TreeGetNextSibling(void * tree) {
	return (void *)CFTreeGetNextSibling(
		// *typing.RefType
		(CFTreeRef)tree
	);
}
void * URLCopyUserName(void * anURL) {
	return (void *)CFURLCopyUserName(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
void * URLCreateAbsoluteURLWithBytes(void * alloc, const uint8_t* relativeURLBytes, CFIndex length, CFStringEncoding encoding, void * baseURL, bool useCompatibilityMode) {
	return (void *)CFURLCreateAbsoluteURLWithBytes(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)relativeURLBytes,
		// *typing.AliasType
		(CFIndex)length,
		// *typing.AliasType
		(CFStringEncoding)encoding,
		// *typing.RefType
		(CFURLRef)baseURL,
		// *typing.PrimitiveType
		(Boolean)useCompatibilityMode
	);
}
void * DateFormatterCreateStringWithDate(void * allocator, void * formatter, void * date) {
	return (void *)CFDateFormatterCreateStringWithDate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFDateFormatterRef)formatter,
		// *typing.RefType
		(CFDateRef)date
	);
}
bool CalendarGetComponentDifference(void * calendar, CFAbsoluteTime startingAT, CFAbsoluteTime resultAT, CFOptionFlags options, char* componentDesc) {
	return (bool)CFCalendarGetComponentDifference(
		// *typing.RefType
		(CFCalendarRef)calendar,
		// *typing.AliasType
		(CFAbsoluteTime)startingAT,
		// *typing.AliasType
		(CFAbsoluteTime)resultAT,
		// *typing.AliasType
		(CFOptionFlags)options,
		// *typing.CStringType
		(char*)componentDesc
	);
}
void * TreeGetChildAtIndex(void * tree, CFIndex idx) {
	return (void *)CFTreeGetChildAtIndex(
		// *typing.RefType
		(CFTreeRef)tree,
		// *typing.AliasType
		(CFIndex)idx
	);
}
CFTypeID DictionaryGetTypeID() {
	return (CFTypeID)CFDictionaryGetTypeID(
	);
}
void * PlugInGetBundle(void * plugIn) {
	return (void *)CFPlugInGetBundle(
		// *typing.RefType
		(CFPlugInRef)plugIn
	);
}
void * BundleCopyResourceURLsOfTypeForLocalization(void * bundle, void * resourceType, void * subDirName, void * localizationName) {
	return (void *)CFBundleCopyResourceURLsOfTypeForLocalization(
		// *typing.RefType
		(CFBundleRef)bundle,
		// *typing.RefType
		(CFStringRef)resourceType,
		// *typing.RefType
		(CFStringRef)subDirName,
		// *typing.RefType
		(CFStringRef)localizationName
	);
}
void * UUIDGetConstantUUIDWithBytes(void * alloc, uint8_t byte0, uint8_t byte1, uint8_t byte2, uint8_t byte3, uint8_t byte4, uint8_t byte5, uint8_t byte6, uint8_t byte7, uint8_t byte8, uint8_t byte9, uint8_t byte10, uint8_t byte11, uint8_t byte12, uint8_t byte13, uint8_t byte14, uint8_t byte15) {
	return (void *)CFUUIDGetConstantUUIDWithBytes(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.PrimitiveType
		(uint8_t)byte0,
		// *typing.PrimitiveType
		(uint8_t)byte1,
		// *typing.PrimitiveType
		(uint8_t)byte2,
		// *typing.PrimitiveType
		(uint8_t)byte3,
		// *typing.PrimitiveType
		(uint8_t)byte4,
		// *typing.PrimitiveType
		(uint8_t)byte5,
		// *typing.PrimitiveType
		(uint8_t)byte6,
		// *typing.PrimitiveType
		(uint8_t)byte7,
		// *typing.PrimitiveType
		(uint8_t)byte8,
		// *typing.PrimitiveType
		(uint8_t)byte9,
		// *typing.PrimitiveType
		(uint8_t)byte10,
		// *typing.PrimitiveType
		(uint8_t)byte11,
		// *typing.PrimitiveType
		(uint8_t)byte12,
		// *typing.PrimitiveType
		(uint8_t)byte13,
		// *typing.PrimitiveType
		(uint8_t)byte14,
		// *typing.PrimitiveType
		(uint8_t)byte15
	);
}
bool TimeZoneIsDaylightSavingTime(void * tz, CFAbsoluteTime at) {
	return (bool)CFTimeZoneIsDaylightSavingTime(
		// *typing.RefType
		(CFTimeZoneRef)tz,
		// *typing.AliasType
		(CFAbsoluteTime)at
	);
}
void URLClearResourcePropertyCache(void * url) {
	return (void)CFURLClearResourcePropertyCache(
		// *typing.RefType
		(CFURLRef)url
	);
}
bool SocketIsValid(void * s) {
	return (bool)CFSocketIsValid(
		// *typing.RefType
		(CFSocketRef)s
	);
}
void * BundleGetDataPointerForName(void * bundle, void * symbolName) {
	return (void *)CFBundleGetDataPointerForName(
		// *typing.RefType
		(CFBundleRef)bundle,
		// *typing.RefType
		(CFStringRef)symbolName
	);
}
int32_t StringConvertEncodingToNSStringEncoding(CFStringEncoding encoding) {
	return (int32_t)CFStringConvertEncodingToNSStringEncoding(
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
void * SocketCopyAddress(void * s) {
	return (void *)CFSocketCopyAddress(
		// *typing.RefType
		(CFSocketRef)s
	);
}
void * BitVectorCreateMutable(void * allocator, CFIndex capacity) {
	return (void *)CFBitVectorCreateMutable(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFIndex)capacity
	);
}
const uint8_t* ReadStreamGetBuffer(void * stream, CFIndex maxBytesToRead, CFIndex* numBytesRead) {
	return (const uint8_t*)CFReadStreamGetBuffer(
		// *typing.RefType
		(CFReadStreamRef)stream,
		// *typing.AliasType
		(CFIndex)maxBytesToRead,
		// *typing.PointerType
		// -> *typing.AliasType
		(CFIndex*)numBytesRead
	);
}
bool PreferencesAppSynchronize(void * applicationID) {
	return (bool)CFPreferencesAppSynchronize(
		// *typing.RefType
		(CFStringRef)applicationID
	);
}
bool RunLoopIsWaiting(void * rl) {
	return (bool)CFRunLoopIsWaiting(
		// *typing.RefType
		(CFRunLoopRef)rl
	);
}
void * URLCreateCopyDeletingLastPathComponent(void * allocator, void * url) {
	return (void *)CFURLCreateCopyDeletingLastPathComponent(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFURLRef)url
	);
}
const CFStringEncoding* StringGetListOfAvailableEncodings() {
	return (const CFStringEncoding*)CFStringGetListOfAvailableEncodings(
	);
}
void * CharacterSetCreateWithCharactersInString(void * alloc, void * theString) {
	return (void *)CFCharacterSetCreateWithCharactersInString(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFStringRef)theString
	);
}
void TimeZoneSetAbbreviationDictionary(void * dict) {
	return (void)CFTimeZoneSetAbbreviationDictionary(
		// *typing.RefType
		(CFDictionaryRef)dict
	);
}
void * URLCreateCopyAppendingPathExtension(void * allocator, void * url, void * extension) {
	return (void *)CFURLCreateCopyAppendingPathExtension(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFURLRef)url,
		// *typing.RefType
		(CFStringRef)extension
	);
}
void StringUppercase(void * theString, void * locale) {
	return (void)CFStringUppercase(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.RefType
		(CFLocaleRef)locale
	);
}
CFPropertyListRef PreferencesCopyValue(void * key, void * applicationID, void * userName, void * hostName) {
	return (CFPropertyListRef)CFPreferencesCopyValue(
		// *typing.RefType
		(CFStringRef)key,
		// *typing.RefType
		(CFStringRef)applicationID,
		// *typing.RefType
		(CFStringRef)userName,
		// *typing.RefType
		(CFStringRef)hostName
	);
}
CFTypeID PlugInGetTypeID() {
	return (CFTypeID)CFPlugInGetTypeID(
	);
}
void * StringCreateArrayBySeparatingStrings(void * alloc, void * theString, void * separatorString) {
	return (void *)CFStringCreateArrayBySeparatingStrings(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFStringRef)theString,
		// *typing.RefType
		(CFStringRef)separatorString
	);
}
void * PlugInInstanceCreate(void * allocator, void * factoryUUID, void * typeUUID) {
	return (void *)CFPlugInInstanceCreate(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFUUIDRef)factoryUUID,
		// *typing.RefType
		(CFUUIDRef)typeUUID
	);
}
void * DateFormatterGetLocale(void * formatter) {
	return (void *)CFDateFormatterGetLocale(
		// *typing.RefType
		(CFDateFormatterRef)formatter
	);
}
uint32_t SwapInt32LittleToHost(uint32_t arg) {
	return (uint32_t)CFSwapInt32LittleToHost(
		// *typing.PrimitiveType
		(uint32_t)arg
	);
}
void * CharacterSetGetPredefined(CFCharacterSetPredefinedSet theSetIdentifier) {
	return (void *)CFCharacterSetGetPredefined(
		// *typing.AliasType
		(CFCharacterSetPredefinedSet)theSetIdentifier
	);
}
CFTypeID DataGetTypeID() {
	return (CFTypeID)CFDataGetTypeID(
	);
}
int32_t UserNotificationReceiveResponse(void * userNotification, CFTimeInterval timeout, CFOptionFlags* responseFlags) {
	return (int32_t)CFUserNotificationReceiveResponse(
		// *typing.RefType
		(CFUserNotificationRef)userNotification,
		// *typing.AliasType
		(CFTimeInterval)timeout,
		// *typing.PointerType
		// -> *typing.AliasType
		(CFOptionFlags*)responseFlags
	);
}
CFTypeID CharacterSetGetTypeID() {
	return (CFTypeID)CFCharacterSetGetTypeID(
	);
}
void * NotificationCenterGetLocalCenter() {
	return (void *)CFNotificationCenterGetLocalCenter(
	);
}
void RunLoopSourceInvalidate(void * source) {
	return (void)CFRunLoopSourceInvalidate(
		// *typing.RefType
		(CFRunLoopSourceRef)source
	);
}
uint64_t SwapInt64BigToHost(uint64_t arg) {
	return (uint64_t)CFSwapInt64BigToHost(
		// *typing.PrimitiveType
		(uint64_t)arg
	);
}
CFIndex DictionaryGetCount(void * theDict) {
	return (CFIndex)CFDictionaryGetCount(
		// *typing.RefType
		(CFDictionaryRef)theDict
	);
}
void * URLCopyPathExtension(void * url) {
	return (void *)CFURLCopyPathExtension(
		// *typing.RefType
		(CFURLRef)url
	);
}
bool MessagePortIsValid(void * ms) {
	return (bool)CFMessagePortIsValid(
		// *typing.RefType
		(CFMessagePortRef)ms
	);
}
uint32_t SwapInt32HostToBig(uint32_t arg) {
	return (uint32_t)CFSwapInt32HostToBig(
		// *typing.PrimitiveType
		(uint32_t)arg
	);
}
CFIndex NumberGetByteSize(void * number) {
	return (CFIndex)CFNumberGetByteSize(
		// *typing.RefType
		(CFNumberRef)number
	);
}
void * URLCreateCopyAppendingPathComponent(void * allocator, void * url, void * pathComponent, bool isDirectory) {
	return (void *)CFURLCreateCopyAppendingPathComponent(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFURLRef)url,
		// *typing.RefType
		(CFStringRef)pathComponent,
		// *typing.PrimitiveType
		(Boolean)isDirectory
	);
}
CFComparisonResult StringCompare(void * theString1, void * theString2, CFStringCompareFlags compareOptions) {
	return (CFComparisonResult)CFStringCompare(
		// *typing.RefType
		(CFStringRef)theString1,
		// *typing.RefType
		(CFStringRef)theString2,
		// *typing.AliasType
		(CFStringCompareFlags)compareOptions
	);
}
CFOptionFlags UserNotificationPopUpSelection(CFIndex n) {
	return (CFOptionFlags)CFUserNotificationPopUpSelection(
		// *typing.AliasType
		(CFIndex)n
	);
}
CFByteOrder ByteOrderGetCurrent() {
	return (CFByteOrder)CFByteOrderGetCurrent(
	);
}
bool ReadStreamOpen(void * stream) {
	return (bool)CFReadStreamOpen(
		// *typing.RefType
		(CFReadStreamRef)stream
	);
}
CFTypeID DateGetTypeID() {
	return (CFTypeID)CFDateGetTypeID(
	);
}
uint16_t SwapInt16HostToLittle(uint16_t arg) {
	return (uint16_t)CFSwapInt16HostToLittle(
		// *typing.PrimitiveType
		(uint16_t)arg
	);
}
void * AttributedStringCreateCopy(void * alloc, void * aStr) {
	return (void *)CFAttributedStringCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFAttributedStringRef)aStr
	);
}
void NumberFormatterSetFormat(void * formatter, void * formatString) {
	return (void)CFNumberFormatterSetFormat(
		// *typing.RefType
		(CFNumberFormatterRef)formatter,
		// *typing.RefType
		(CFStringRef)formatString
	);
}
void * TimeZoneCopyKnownNames() {
	return (void *)CFTimeZoneCopyKnownNames(
	);
}
void * SocketCreateRunLoopSource(void * allocator, void * s, CFIndex order) {
	return (void *)CFSocketCreateRunLoopSource(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFSocketRef)s,
		// *typing.AliasType
		(CFIndex)order
	);
}
CFTypeID BitVectorGetTypeID() {
	return (CFTypeID)CFBitVectorGetTypeID(
	);
}
void CharacterSetUnion(void * theSet, void * theOtherSet) {
	return (void)CFCharacterSetUnion(
		// *typing.RefType
		(CFMutableCharacterSetRef)theSet,
		// *typing.RefType
		(CFCharacterSetRef)theOtherSet
	);
}
void * RunLoopCopyAllModes(void * rl) {
	return (void *)CFRunLoopCopyAllModes(
		// *typing.RefType
		(CFRunLoopRef)rl
	);
}
CFIndex SetGetCount(void * theSet) {
	return (CFIndex)CFSetGetCount(
		// *typing.RefType
		(CFSetRef)theSet
	);
}
bool WriteStreamOpen(void * stream) {
	return (bool)CFWriteStreamOpen(
		// *typing.RefType
		(CFWriteStreamRef)stream
	);
}
void SocketSetSocketFlags(void * s, CFOptionFlags flags) {
	return (void)CFSocketSetSocketFlags(
		// *typing.RefType
		(CFSocketRef)s,
		// *typing.AliasType
		(CFOptionFlags)flags
	);
}
CFTypeID DateFormatterGetTypeID() {
	return (CFTypeID)CFDateFormatterGetTypeID(
	);
}
CFTypeID URLGetTypeID() {
	return (CFTypeID)CFURLGetTypeID(
	);
}
void BagRemoveAllValues(void * theBag) {
	return (void)CFBagRemoveAllValues(
		// *typing.RefType
		(CFMutableBagRef)theBag
	);
}
void BinaryHeapRemoveAllValues(void * heap) {
	return (void)CFBinaryHeapRemoveAllValues(
		// *typing.RefType
		(CFBinaryHeapRef)heap
	);
}
void BitVectorFlipBitAtIndex(void * bv, CFIndex idx) {
	return (void)CFBitVectorFlipBitAtIndex(
		// *typing.RefType
		(CFMutableBitVectorRef)bv,
		// *typing.AliasType
		(CFIndex)idx
	);
}
CFTypeID NotificationCenterGetTypeID() {
	return (CFTypeID)CFNotificationCenterGetTypeID(
	);
}
CFIndex TreeGetChildCount(void * tree) {
	return (CFIndex)CFTreeGetChildCount(
		// *typing.RefType
		(CFTreeRef)tree
	);
}
CFTypeID FileSecurityGetTypeID() {
	return (CFTypeID)CFFileSecurityGetTypeID(
	);
}
void * BundleCopyResourceURLInDirectory(void * bundleURL, void * resourceName, void * resourceType, void * subDirName) {
	return (void *)CFBundleCopyResourceURLInDirectory(
		// *typing.RefType
		(CFURLRef)bundleURL,
		// *typing.RefType
		(CFStringRef)resourceName,
		// *typing.RefType
		(CFStringRef)resourceType,
		// *typing.RefType
		(CFStringRef)subDirName
	);
}
bool BundleIsExecutableLoadable(void * bundle) {
	return (bool)CFBundleIsExecutableLoadable(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
CFTypeID AllocatorGetTypeID() {
	return (CFTypeID)CFAllocatorGetTypeID(
	);
}
void URLSetTemporaryResourcePropertyForKey(void * url, void * key, CFTypeRef propertyValue) {
	return (void)CFURLSetTemporaryResourcePropertyForKey(
		// *typing.RefType
		(CFURLRef)url,
		// *typing.RefType
		(CFStringRef)key,
		// *typing.AliasType
		(CFTypeRef)propertyValue
	);
}
CFTypeID PlugInInstanceGetTypeID() {
	return (CFTypeID)CFPlugInInstanceGetTypeID(
	);
}
bool PlugInRegisterPlugInType(void * factoryUUID, void * typeUUID) {
	return (bool)CFPlugInRegisterPlugInType(
		// *typing.RefType
		(CFUUIDRef)factoryUUID,
		// *typing.RefType
		(CFUUIDRef)typeUUID
	);
}
CFTypeID UUIDGetTypeID() {
	return (CFTypeID)CFUUIDGetTypeID(
	);
}
void * TimeZoneGetData(void * tz) {
	return (void *)CFTimeZoneGetData(
		// *typing.RefType
		(CFTimeZoneRef)tz
	);
}
CFLocaleLanguageDirection LocaleGetLanguageCharacterDirection(void * isoLangCode) {
	return (CFLocaleLanguageDirection)CFLocaleGetLanguageCharacterDirection(
		// *typing.RefType
		(CFStringRef)isoLangCode
	);
}
void * StringCreateWithBytesNoCopy(void * alloc, const uint8_t* bytes, CFIndex numBytes, CFStringEncoding encoding, bool isExternalRepresentation, void * contentsDeallocator) {
	return (void *)CFStringCreateWithBytesNoCopy(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.PointerType
		// -> *typing.PrimitiveType
		(uint8_t*)bytes,
		// *typing.AliasType
		(CFIndex)numBytes,
		// *typing.AliasType
		(CFStringEncoding)encoding,
		// *typing.PrimitiveType
		(Boolean)isExternalRepresentation,
		// *typing.RefType
		(CFAllocatorRef)contentsDeallocator
	);
}
void TreeRemove(void * tree) {
	return (void)CFTreeRemove(
		// *typing.RefType
		(CFTreeRef)tree
	);
}
bool BundleLoadExecutable(void * bundle) {
	return (bool)CFBundleLoadExecutable(
		// *typing.RefType
		(CFBundleRef)bundle
	);
}
bool RunLoopObserverDoesRepeat(void * observer) {
	return (bool)CFRunLoopObserverDoesRepeat(
		// *typing.RefType
		(CFRunLoopObserverRef)observer
	);
}
void RunLoopTimerInvalidate(void * timer) {
	return (void)CFRunLoopTimerInvalidate(
		// *typing.RefType
		(CFRunLoopTimerRef)timer
	);
}
CFAbsoluteTime TimeZoneGetNextDaylightSavingTimeTransition(void * tz, CFAbsoluteTime at) {
	return (CFAbsoluteTime)CFTimeZoneGetNextDaylightSavingTimeTransition(
		// *typing.RefType
		(CFTimeZoneRef)tz,
		// *typing.AliasType
		(CFAbsoluteTime)at
	);
}
void * BinaryHeapGetMinimum(void * heap) {
	return (void *)CFBinaryHeapGetMinimum(
		// *typing.RefType
		(CFBinaryHeapRef)heap
	);
}
void * BundleCopyPreferredLocalizationsFromArray(void * locArray) {
	return (void *)CFBundleCopyPreferredLocalizationsFromArray(
		// *typing.RefType
		(CFArrayRef)locArray
	);
}
void * MessagePortCreateRunLoopSource(void * allocator, void * local, CFIndex order) {
	return (void *)CFMessagePortCreateRunLoopSource(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.RefType
		(CFMessagePortRef)local,
		// *typing.AliasType
		(CFIndex)order
	);
}
CFTypeID LocaleGetTypeID() {
	return (CFTypeID)CFLocaleGetTypeID(
	);
}
bool CalendarComposeAbsoluteTime(void * calendar, CFAbsoluteTime* at, char* componentDesc) {
	return (bool)CFCalendarComposeAbsoluteTime(
		// *typing.RefType
		(CFCalendarRef)calendar,
		// *typing.PointerType
		// -> *typing.AliasType
		(CFAbsoluteTime*)at,
		// *typing.CStringType
		(char*)componentDesc
	);
}
bool PlugInUnregisterPlugInType(void * factoryUUID, void * typeUUID) {
	return (bool)CFPlugInUnregisterPlugInType(
		// *typing.RefType
		(CFUUIDRef)factoryUUID,
		// *typing.RefType
		(CFUUIDRef)typeUUID
	);
}
bool ReadStreamHasBytesAvailable(void * stream) {
	return (bool)CFReadStreamHasBytesAvailable(
		// *typing.RefType
		(CFReadStreamRef)stream
	);
}
UniChar StringGetCharacterAtIndex(void * theString, CFIndex idx) {
	return (UniChar)CFStringGetCharacterAtIndex(
		// *typing.RefType
		(CFStringRef)theString,
		// *typing.AliasType
		(CFIndex)idx
	);
}
CFIndex CalendarGetOrdinalityOfUnit(void * calendar, CFCalendarUnit smallerUnit, CFCalendarUnit biggerUnit, CFAbsoluteTime at) {
	return (CFIndex)CFCalendarGetOrdinalityOfUnit(
		// *typing.RefType
		(CFCalendarRef)calendar,
		// *typing.AliasType
		(CFCalendarUnit)smallerUnit,
		// *typing.AliasType
		(CFCalendarUnit)biggerUnit,
		// *typing.AliasType
		(CFAbsoluteTime)at
	);
}
CFPropertyListRef PropertyListCreateDeepCopy(void * allocator, CFPropertyListRef propertyList, CFOptionFlags mutabilityOption) {
	return (CFPropertyListRef)CFPropertyListCreateDeepCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFPropertyListRef)propertyList,
		// *typing.AliasType
		(CFOptionFlags)mutabilityOption
	);
}
void * CalendarCopyTimeZone(void * calendar) {
	return (void *)CFCalendarCopyTimeZone(
		// *typing.RefType
		(CFCalendarRef)calendar
	);
}
int32_t URLGetPortNumber(void * anURL) {
	return (int32_t)CFURLGetPortNumber(
		// *typing.RefType
		(CFURLRef)anURL
	);
}
void * LocaleCopyISOCountryCodes() {
	return (void *)CFLocaleCopyISOCountryCodes(
	);
}
void * BinaryHeapCreateCopy(void * allocator, CFIndex capacity, void * heap) {
	return (void *)CFBinaryHeapCreateCopy(
		// *typing.RefType
		(CFAllocatorRef)allocator,
		// *typing.AliasType
		(CFIndex)capacity,
		// *typing.RefType
		(CFBinaryHeapRef)heap
	);
}
bool FileSecurityGetGroup(void * fileSec, gid_t* group) {
	return (bool)CFFileSecurityGetGroup(
		// *typing.RefType
		(CFFileSecurityRef)fileSec,
		// *typing.PointerType
		// -> *typing.AliasType
		(gid_t*)group
	);
}
void * WriteStreamCreateWithFile(void * alloc, void * fileURL) {
	return (void *)CFWriteStreamCreateWithFile(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFURLRef)fileURL
	);
}
void CharacterSetInvert(void * theSet) {
	return (void)CFCharacterSetInvert(
		// *typing.RefType
		(CFMutableCharacterSetRef)theSet
	);
}
CFTypeID ErrorGetTypeID() {
	return (CFTypeID)CFErrorGetTypeID(
	);
}
void * BundleGetFunctionPointerForName(void * bundle, void * functionName) {
	return (void *)CFBundleGetFunctionPointerForName(
		// *typing.RefType
		(CFBundleRef)bundle,
		// *typing.RefType
		(CFStringRef)functionName
	);
}
void StringFold(void * theString, CFStringCompareFlags theFlags, void * theLocale) {
	return (void)CFStringFold(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.AliasType
		(CFStringCompareFlags)theFlags,
		// *typing.RefType
		(CFLocaleRef)theLocale
	);
}
bool URLStartAccessingSecurityScopedResource(void * url) {
	return (bool)CFURLStartAccessingSecurityScopedResource(
		// *typing.RefType
		(CFURLRef)url
	);
}
void * StringCreateExternalRepresentation(void * alloc, void * theString, CFStringEncoding encoding, uint8_t lossByte) {
	return (void *)CFStringCreateExternalRepresentation(
		// *typing.RefType
		(CFAllocatorRef)alloc,
		// *typing.RefType
		(CFStringRef)theString,
		// *typing.AliasType
		(CFStringEncoding)encoding,
		// *typing.PrimitiveType
		(uint8_t)lossByte
	);
}
CFSocketError SocketConnectToAddress(void * s, void * address, CFTimeInterval timeout) {
	return (CFSocketError)CFSocketConnectToAddress(
		// *typing.RefType
		(CFSocketRef)s,
		// *typing.RefType
		(CFDataRef)address,
		// *typing.AliasType
		(CFTimeInterval)timeout
	);
}
void StringAppendCString(void * theString, char* cStr, CFStringEncoding encoding) {
	return (void)CFStringAppendCString(
		// *typing.RefType
		(CFMutableStringRef)theString,
		// *typing.CStringType
		(char*)cStr,
		// *typing.AliasType
		(CFStringEncoding)encoding
	);
}
CFDateFormatterStyle DateFormatterGetDateStyle(void * formatter) {
	return (CFDateFormatterStyle)CFDateFormatterGetDateStyle(
		// *typing.RefType
		(CFDateFormatterRef)formatter
	);
}
void * BundleCopyLocalizationsForURL(void * url) {
	return (void *)CFBundleCopyLocalizationsForURL(
		// *typing.RefType
		(CFURLRef)url
	);
}
void * UserNotificationGetResponseDictionary(void * userNotification) {
	return (void *)CFUserNotificationGetResponseDictionary(
		// *typing.RefType
		(CFUserNotificationRef)userNotification
	);
}
CFIndex BinaryHeapGetCount(void * heap) {
	return (CFIndex)CFBinaryHeapGetCount(
		// *typing.RefType
		(CFBinaryHeapRef)heap
	);
}
CFTypeID UserNotificationGetTypeID() {
	return (CFTypeID)CFUserNotificationGetTypeID(
	);
}
void * BundleCopyResourceURLsOfType(void * bundle, void * resourceType, void * subDirName) {
	return (void *)CFBundleCopyResourceURLsOfType(
		// *typing.RefType
		(CFBundleRef)bundle,
		// *typing.RefType
		(CFStringRef)resourceType,
		// *typing.RefType
		(CFStringRef)subDirName
	);
}
CFTypeID RunLoopSourceGetTypeID() {
	return (CFTypeID)CFRunLoopSourceGetTypeID(
	);
}
void StringInsert(void * str, CFIndex idx, void * insertedStr) {
	return (void)CFStringInsert(
		// *typing.RefType
		(CFMutableStringRef)str,
		// *typing.AliasType
		(CFIndex)idx,
		// *typing.RefType
		(CFStringRef)insertedStr
	);
}
