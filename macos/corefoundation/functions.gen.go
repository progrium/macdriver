// AUTO-GENERATED CODE, DO NOT MODIFY

package corefoundation

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import "CoreFoundation/CoreFoundation.h"
// void * DictionaryCreateMutableCopy(void * allocator, CFIndex capacity, void * theDict);
// CFTypeID NullGetTypeID();
// void * WriteStreamCopyError(void * stream);
// void * BundleCopyResourceURLForLocalization(void * bundle, void * resourceName, void * resourceType, void * subDirName, void * localizationName);
// void * StringCreateCopy(void * alloc, void * theString);
// void CalendarSetFirstWeekday(void * calendar, CFIndex wkdy);
// CFTypeID NumberFormatterGetTypeID();
// bool BundleIsArchitectureLoadable(cpu_type_t arch);
// void * BundleCopyInfoDictionaryInDirectory(void * bundleURL);
// uint32_t SwapInt32(uint32_t arg);
// void * URLCopyStrictPath(void * anURL, bool* isAbsolute);
// void * TimeZoneCopySystem();
// void * FileSecurityCreateCopy(void * allocator, void * fileSec);
// void * URLCreateWithFileSystemPath(void * allocator, void * filePath, CFURLPathStyle pathStyle, bool isDirectory);
// CFIndex BagGetCount(void * theBag);
// void * BundleCopyAuxiliaryExecutableURL(void * bundle, void * executableName);
// void * BagCreateCopy(void * allocator, void * theBag);
// void WriteStreamClose(void * stream);
// bool PreferencesGetAppBooleanValue(void * key, void * applicationID, bool* keyExistsAndHasValidFormat);
// void * SetCreateCopy(void * allocator, void * theSet);
// void DateFormatterSetFormat(void * formatter, void * formatString);
// void TreeRemoveAllChildren(void * tree);
// void * CharacterSetCreateWithBitmapRepresentation(void * alloc, void * theData);
// CFBit BitVectorGetBitAtIndex(void * bv, CFIndex idx);
// void * GetAllocator(CFTypeRef cf);
// void * BundleCopyBuiltInPlugInsURL(void * bundle);
// CFIndex ErrorGetCode(void * err);
// CFIndex URLGetBytes(void * url, uint8_t* buffer, CFIndex bufferLength);
// void BitVectorSetBitAtIndex(void * bv, CFIndex idx, CFBit value);
// void PlugInAddInstanceForFactory(void * factoryID);
// CFIndex CalendarGetMinimumDaysInFirstWeek(void * calendar);
// void DataIncreaseLength(void * theData, CFIndex extraLength);
// bool RunLoopTimerIsValid(void * timer);
// void * BundleCopyResourceURL(void * bundle, void * resourceName, void * resourceType, void * subDirName);
// void * StringCreateWithCharacters(void * alloc, const UniChar* chars, CFIndex numChars);
// CFIndex ReadStreamRead(void * stream, uint8_t* buffer, CFIndex bufferLength);
// CFTypeID RunLoopObserverGetTypeID();
// void * NumberFormatterCreate(void * allocator, void * locale, CFNumberFormatterStyle style);
// void * NumberFormatterGetLocale(void * formatter);
// void * MessagePortCreateRemote(void * allocator, void * name);
// CFOptionFlags RunLoopObserverGetActivities(void * observer);
// bool PropertyListIsValid(CFPropertyListRef plist, CFPropertyListFormat format);
// CFTimeInterval RunLoopTimerGetInterval(void * timer);
// double StringGetDoubleValue(void * str);
// void * TimeZoneCopyAbbreviationDictionary();
// void FileDescriptorEnableCallBacks(void * f, CFOptionFlags callBackTypes);
// void * BundleGetInfoDictionary(void * bundle);
// void * AttributedStringCreate(void * alloc, void * str, void * attributes);
// uint64_t SwapInt64HostToLittle(uint64_t arg);
// void FileDescriptorInvalidate(void * f);
// void * DataCreate(void * allocator, const uint8_t* bytes, CFIndex length);
// void * BundleCopyInfoDictionaryForURL(void * url);
// void * StringConvertEncodingToIANACharSetName(CFStringEncoding encoding);
// void * URLGetString(void * anURL);
// void DictionaryRemoveAllValues(void * theDict);
// void * DataCreateCopy(void * allocator, void * theData);
// bool BundleIsExecutableLoaded(void * bundle);
// CFStreamStatus ReadStreamGetStatus(void * stream);
// void * URLEnumeratorCreateForDirectoryURL(void * alloc, void * directoryURL, CFURLEnumeratorOptions option, void * propertyKeys);
// void * LocaleCopyCurrent();
// void * TimeZoneCopyAbbreviation(void * tz, CFAbsoluteTime at);
// uint16_t SwapInt16BigToHost(uint16_t arg);
// CFTypeID TimeZoneGetTypeID();
// void * PreferencesCopyMultiple(void * keysToFetch, void * applicationID, void * userName, void * hostName);
// CFTypeRef URLCreateResourcePropertyForKeyFromBookmarkData(void * allocator, void * resourcePropertyKey, void * bookmark);
// bool MessagePortSetName(void * ms, void * newName);
// void * TreeFindRoot(void * tree);
// void * CharacterSetCreateInvertedSet(void * alloc, void * theSet);
// int32_t UserNotificationUpdate(void * userNotification, CFTimeInterval timeout, CFOptionFlags flags, void * dictionary);
// void * ErrorCopyRecoverySuggestion(void * err);
// void * URLEnumeratorCreateForMountedVolumes(void * alloc, CFURLEnumeratorOptions option, void * propertyKeys);
// CFTimeInterval RunLoopTimerGetTolerance(void * timer);
// void * BundleCopyLocalizedString(void * bundle, void * key, void * value, void * tableName);
// void * RunLoopGetCurrent();
// void PreferencesAddSuitePreferencesToApp(void * applicationID, void * suiteID);
// void * CharacterSetCreateCopy(void * alloc, void * theSet);
// void Release(CFTypeRef cf);
// bool FileDescriptorIsValid(void * f);
// void * CopyDescription(CFTypeRef cf);
// CFTypeID TreeGetTypeID();
// void TreeInsertSibling(void * tree, void * newSibling);
// void * WriteStreamCreateWithAllocatedBuffers(void * alloc, void * bufferAllocator);
// void CalendarSetTimeZone(void * calendar, void * tz);
// bool FileSecuritySetGroupUUID(void * fileSec, void * groupUUID);
// void * CharacterSetCreateMutable(void * alloc);
// bool FileSecuritySetGroup(void * fileSec, gid_t group);
// CFStreamStatus WriteStreamGetStatus(void * stream);
// void * AttributedStringCreateMutableCopy(void * alloc, CFIndex maxLength, void * aStr);
// void * TimeZoneGetName(void * tz);
// bool BundleGetPackageInfoInDirectory(void * url, uint32_t* packageType, uint32_t* packageCreator);
// void ArrayExchangeValuesAtIndices(void * theArray, CFIndex idx1, CFIndex idx2);
// CFTypeRef MakeCollectable(CFTypeRef cf);
// void CalendarSetMinimumDaysInFirstWeek(void * calendar, CFIndex mwd);
// void CharacterSetIntersect(void * theSet, void * theOtherSet);
// void * XMLCreateStringByUnescapingEntities(void * allocator, void * string_, void * entitiesDictionary);
// void * URLCopyPath(void * anURL);
// void SocketEnableCallBacks(void * s, CFOptionFlags callBackTypes);
// void ReadStreamClose(void * stream);
// CFTypeID CalendarGetTypeID();
// void * BundleGetIdentifier(void * bundle);
// void * URLCopyPassword(void * anURL);
// int32_t StringGetIntValue(void * str);
// bool FileSecurityGetMode(void * fileSec, mode_t* mode);
// CFTypeID RunLoopGetTypeID();
// void URLEnumeratorSkipDescendents(void * enumerator);
// CFTypeID RunLoopTimerGetTypeID();
// void * URLCopyFileSystemPath(void * anURL, CFURLPathStyle pathStyle);
// void ArrayRemoveAllValues(void * theArray);
// void * MachPortCreateRunLoopSource(void * allocator, void * port, CFIndex order);
// void * StringGetNameOfEncoding(CFStringEncoding encoding);
// void PlugInRemoveInstanceForFactory(void * factoryID);
// CFHashCode Hash(CFTypeRef cf);
// bool URLGetFileSystemRepresentation(void * url, bool resolveAgainstBase, uint8_t* buffer, CFIndex maxBufLen);
// void * ReadStreamCopyError(void * stream);
// void TreePrependChild(void * tree, void * newChild);
// void * LocaleCopyISOCurrencyCodes();
// void * AllocatorAllocate(void * allocator, CFIndex size, CFOptionFlags hint);
// CFIndex DataGetLength(void * theData);
// bool PreferencesAppValueIsForced(void * key, void * applicationID);
// CFStringEncoding StringGetFastestEncoding(void * theString);
// void * StringCreateWithBytes(void * alloc, const uint8_t* bytes, CFIndex numBytes, CFStringEncoding encoding, bool isExternalRepresentation);
// bool StringGetCString(void * theString, char* buffer, CFIndex bufferSize, CFStringEncoding encoding);
// bool PlugInUnregisterFactory(void * factoryUUID);
// int32_t UserNotificationCancel(void * userNotification);
// bool StringIsEncodingAvailable(CFStringEncoding encoding);
// bool StringIsSurrogateHighCharacter(UniChar character);
// void * ErrorCopyFailureReason(void * err);
// void * DictionaryCreateCopy(void * allocator, void * theDict);
// void * URLCopyScheme(void * anURL);
// CFTypeID NumberGetTypeID();
// bool CalendarGetTimeRangeOfUnit(void * calendar, CFCalendarUnit unit, CFAbsoluteTime at, CFAbsoluteTime* startp, CFTimeInterval* tip);
// void PreferencesSetAppValue(void * key, CFPropertyListRef value, void * applicationID);
// void ArrayRemoveValueAtIndex(void * theArray, CFIndex idx);
// uint64_t SwapInt64(uint64_t arg);
// CFIndex CalendarGetFirstWeekday(void * calendar);
// void * BundleGetAllBundles();
// bool URLHasDirectoryPath(void * anURL);
// void * BitVectorCreate(void * allocator, const uint8_t* bytes, CFIndex numBits);
// void * NotificationCenterGetDistributedCenter();
// void StringAppendCharacters(void * theString, const UniChar* chars, CFIndex numChars);
// bool PlugInRegisterFactoryFunctionByName(void * factoryUUID, void * plugIn, void * functionName);
// bool CalendarAddComponents(void * calendar, CFAbsoluteTime* at, CFOptionFlags options, char* componentDesc);
// CFNumberType NumberGetType(void * number);
// void * DateFormatterGetFormat(void * formatter);
// void * BundleCopyBundleLocalizations(void * bundle);
// void * BundleGetMainBundle();
// bool RunLoopObserverIsValid(void * observer);
// void AttributedStringEndEditing(void * aStr);
// bool CharacterSetIsCharacterMember(void * theSet, UniChar theChar);
// void * URLCopyAbsoluteURL(void * relativeURL);
// void * BundleCopySharedSupportURL(void * bundle);
// void StringTrim(void * theString, void * trimString);
// CFOptionFlags UserNotificationSecureTextField(CFIndex i);
// CFOptionFlags SocketGetSocketFlags(void * s);
// CFTypeID StringGetTypeID();
// uint32_t StringConvertEncodingToWindowsCodepage(CFStringEncoding encoding);
// void * LocaleCopyCommonISOCurrencyCodes();
// void * URLCreateWithBytes(void * allocator, const uint8_t* URLBytes, CFIndex length, CFStringEncoding encoding, void * baseURL);
// CFTypeID URLEnumeratorGetTypeID();
// CFStringTokenizerTokenType StringTokenizerAdvanceToNextToken(void * tokenizer);
// void * StringCreateMutableCopy(void * alloc, CFIndex maxLength, void * theString);
// void RunLoopTimerSetNextFireDate(void * timer, CFAbsoluteTime fireDate);
// void * TimeZoneCreateWithName(void * allocator, void * name, bool tryAbbrev);
// void MessagePortInvalidate(void * ms);
// bool CalendarDecomposeAbsoluteTime(void * calendar, CFAbsoluteTime at, char* componentDesc);
// CFTypeRef Autorelease(CFTypeRef arg);
// void * BundleCopyResourceURLsOfTypeInDirectory(void * bundleURL, void * resourceType, void * subDirName);
// void RunLoopTimerSetTolerance(void * timer, CFTimeInterval tolerance);
// void StringSetExternalCharactersNoCopy(void * theString, UniChar* chars, CFIndex length, CFIndex capacity);
// void * BundleGetBundleWithIdentifier(void * bundleID);
// void * StringCreateMutable(void * alloc, CFIndex maxLength);
// void * URLCreateWithFileSystemPathRelativeToBase(void * allocator, void * filePath, CFURLPathStyle pathStyle, bool isDirectory, void * baseURL);
// void * URLCreateWithString(void * allocator, void * URLString, void * baseURL);
// void * BagCreateMutableCopy(void * allocator, CFIndex capacity, void * theBag);
// CFTypeRef BundleGetValueForInfoDictionaryKey(void * bundle, void * key);
// void BitVectorSetCount(void * bv, CFIndex count);
// void * URLGetBaseURL(void * anURL);
// CFTypeID BagGetTypeID();
// bool BooleanGetValue(void * boolean);
// void * ErrorCopyDescription(void * err);
// void * PlugInInstanceGetInstanceData(void * instance);
// CFNumberFormatterStyle NumberFormatterGetStyle(void * formatter);
// void StringCapitalize(void * theString, void * locale);
// const uint8_t* DataGetBytePtr(void * theData);
// bool FileSecuritySetMode(void * fileSec, mode_t mode);
// CFOptionFlags UserNotificationCheckBoxChecked(CFIndex i);
// void * URLCopyLastPathComponent(void * url);
// CFIndex StringGetMaximumSizeOfFileSystemRepresentation(void * string_);
// void * BundleCopyResourcesDirectoryURL(void * bundle);
// CFTypeID BooleanGetTypeID();
// void * AttributedStringGetMutableString(void * aStr);
// void URLClearResourcePropertyCacheForKey(void * url, void * key);
// void * URLCopyFragment(void * anURL, void * charactersToLeaveEscaped);
// void SocketInvalidate(void * s);
// void * UUIDCreateWithBytes(void * alloc, uint8_t byte0, uint8_t byte1, uint8_t byte2, uint8_t byte3, uint8_t byte4, uint8_t byte5, uint8_t byte6, uint8_t byte7, uint8_t byte8, uint8_t byte9, uint8_t byte10, uint8_t byte11, uint8_t byte12, uint8_t byte13, uint8_t byte14, uint8_t byte15);
// void * UserNotificationCreate(void * allocator, CFTimeInterval timeout, CFOptionFlags flags, int32_t* error, void * dictionary);
// void ShowStr(void * str);
// void * WriteStreamCreateWithBuffer(void * alloc, uint8_t* buffer, CFIndex bufferCapacity);
// void * SetCreateMutableCopy(void * allocator, CFIndex capacity, void * theSet);
// void AllocatorSetDefault(void * allocator);
// bool WriteStreamCanAcceptBytes(void * stream);
// CFTypeID SetGetTypeID();
// CFTypeID StringTokenizerGetTypeID();
// void * TimeZoneCopyLocalizedName(void * tz, CFTimeZoneNameStyle style, void * locale);
// CFPropertyListRef PreferencesCopyAppValue(void * key, void * applicationID);
// CFTypeID AttributedStringGetTypeID();
// void * URLCopyNetLocation(void * anURL);
// void PreferencesSetValue(void * key, CFPropertyListRef value, void * applicationID, void * userName, void * hostName);
// void FileDescriptorDisableCallBacks(void * f, CFOptionFlags callBackTypes);
// void * DateFormatterCreateStringWithAbsoluteTime(void * allocator, void * formatter, CFAbsoluteTime at);
// bool PlugInIsLoadOnDemand(void * plugIn);
// CFIndex RunLoopTimerGetOrder(void * timer);
// void StringReplaceAll(void * theString, void * replacement);
// void * TimeZoneCopyDefault();
// uint32_t SwapInt32BigToHost(uint32_t arg);
// CFTimeInterval TimeZoneGetSecondsFromGMT(void * tz, CFAbsoluteTime at);
// void * ReadStreamCreateWithFile(void * alloc, void * fileURL);
// void DataAppendBytes(void * theData, const uint8_t* bytes, CFIndex length);
// void * LocaleGetSystem();
// void * StringCreateWithCharactersNoCopy(void * alloc, const UniChar* chars, CFIndex numChars, void * contentsDeallocator);
// uint64_t SwapInt64LittleToHost(uint64_t arg);
// CFIndex ArrayGetCount(void * theArray);
// void * TreeGetParent(void * tree);
// uint64_t SwapInt64HostToBig(uint64_t arg);
// void RunLoopRun();
// void * RunLoopGetMain();
// void * NumberFormatterCreateStringWithNumber(void * allocator, void * formatter, void * number);
// CFIndex URLEnumeratorGetDescendentLevel(void * enumerator);
// CFTypeID SocketGetTypeID();
// bool MessagePortIsRemote(void * ms);
// void MachPortInvalidate(void * port);
// void CharacterSetAddCharactersInString(void * theSet, void * theString);
// void * LocaleCopyISOLanguageCodes();
// CFSocketError SocketSendData(void * s, void * address, void * data, CFTimeInterval timeout);
// CFSocketNativeHandle SocketGetNative(void * s);
// void SocketSetDefaultNameRegistryPortNumber(uint16_t port);
// CFStringEncoding StringConvertIANACharSetNameToEncoding(void * theString);
// void * StringCreateFromExternalRepresentation(void * alloc, void * data, CFStringEncoding encoding);
// bool StringHasSuffix(void * theString, void * suffix);
// bool FileSecurityGetOwner(void * fileSec, uid_t* owner);
// void * PlugInInstanceGetFactoryName(void * instance);
// void * BundleGetDevelopmentRegion(void * bundle);
// void DataSetLength(void * theData, CFIndex length);
// void * BundleCopyLocalizationsForPreferences(void * locArray, void * prefArray);
// void * URLCreateData(void * allocator, void * url, CFStringEncoding encoding, bool escapeWhitespace);
// void * LocaleCreateCopy(void * allocator, void * locale);
// CFLocaleLanguageDirection LocaleGetLanguageLineDirection(void * isoLangCode);
// void * DateFormatterCreate(void * allocator, void * locale, CFDateFormatterStyle dateStyle, CFDateFormatterStyle timeStyle);
// void * PlugInFindFactoriesForPlugInType(void * typeUUID);
// void * PlugInCreate(void * allocator, void * plugInURL);
// CFAbsoluteTime RunLoopTimerGetNextFireDate(void * timer);
// bool NumberFormatterGetDecimalInfoForCurrencyCode(void * currencyCode, int32_t* defaultFractionDigits, double* roundingIncrement);
// void * CalendarCopyLocale(void * calendar);
// CFIndex WriteStreamWrite(void * stream, const uint8_t* buffer, CFIndex bufferLength);
// void Show(CFTypeRef obj);
// bool StringIsHyphenationAvailableForLocale(void * locale);
// CFTypeRef StringTokenizerCopyCurrentTokenAttribute(void * tokenizer, CFOptionFlags attribute);
// void * UUIDCreateFromString(void * alloc, void * uuidStr);
// CFTypeID ArrayGetTypeID();
// void RunLoopSourceSignal(void * source);
// CFSocketError SocketSetAddress(void * s, void * address);
// void * StringCreateWithCString(void * alloc, char* cStr, CFStringEncoding encoding);
// void * DataCreateMutableCopy(void * allocator, CFIndex capacity, void * theData);
// void * AttributedStringCreateMutable(void * alloc, CFIndex maxLength);
// void * MessagePortGetName(void * ms);
// bool FileSecuritySetOwnerUUID(void * fileSec, void * ownerUUID);
// CFAbsoluteTime DateGetAbsoluteTime(void * theDate);
// void CalendarSetLocale(void * calendar, void * locale);
// void * BundleCopyExecutableArchitectures(void * bundle);
// void RunLoopObserverInvalidate(void * observer);
// void PreferencesSetMultiple(void * keysToSet, void * keysToRemove, void * applicationID, void * userName, void * hostName);
// CFIndex StringGetMaximumSizeForEncoding(CFIndex length, CFStringEncoding encoding);
// bool RunLoopTimerDoesRepeat(void * timer);
// void * LocaleCopyAvailableLocaleIdentifiers();
// void * BundleCopySharedFrameworksURL(void * bundle);
// uint16_t SwapInt16(uint16_t arg);
// void * BundleCreate(void * allocator, void * bundleURL);
// void * URLCreateStringByReplacingPercentEscapes(void * allocator, void * originalString, void * charactersToLeaveEscaped);
// CFTimeInterval TimeZoneGetDaylightSavingTimeOffset(void * tz, CFAbsoluteTime at);
// void BitVectorSetAllBits(void * bv, CFBit value);
// void * XMLCreateStringByEscapingEntities(void * allocator, void * string_, void * entitiesDictionary);
// CFTypeID GetTypeID(CFTypeRef cf);
// CFIndex PreferencesGetAppIntegerValue(void * key, void * applicationID, bool* keyExistsAndHasValidFormat);
// void * StringCreateWithCStringNoCopy(void * alloc, char* cStr, CFStringEncoding encoding, void * contentsDeallocator);
// CFTypeID ReadStreamGetTypeID();
// void * DateCreate(void * allocator, CFAbsoluteTime at);
// void StringLowercase(void * theString, void * locale);
// CFStringEncoding StringConvertNSStringEncodingToEncoding(int32_t encoding);
// CFStringEncoding StringGetSmallestEncoding(void * theString);
// bool StringHasPrefix(void * theString, void * prefix);
// void * DateFormatterCreateDateFormatFromTemplate(void * allocator, void * tmplate, CFOptionFlags options, void * locale);
// bool FileSecurityClearProperties(void * fileSec, CFFileSecurityClearOptions clearPropertyMask);
// uint8_t* DataGetMutableBytePtr(void * theData);
// CFIndex BitVectorGetCount(void * bv);
// void RunLoopStop(void * rl);
// bool CharacterSetHasMemberInPlane(void * theSet, CFIndex thePlane);
// void RunLoopWakeUp(void * rl);
// bool StringGetFileSystemRepresentation(void * string_, char* buffer, CFIndex maxBufLen);
// bool URLCanBeDecomposed(void * anURL);
// void * UserNotificationGetResponseValue(void * userNotification, void * key, CFIndex idx);
// void * ArrayCreateMutableCopy(void * allocator, CFIndex capacity, void * theArray);
// void * ReadStreamCreateWithBytesNoCopy(void * alloc, const uint8_t* bytes, CFIndex length, void * bytesDeallocator);
// int32_t UserNotificationDisplayAlert(CFTimeInterval timeout, CFOptionFlags flags, void * iconURL, void * soundURL, void * localizationURL, void * alertHeader, void * alertMessage, void * defaultButtonTitle, void * alternateButtonTitle, void * otherButtonTitle, CFOptionFlags* responseFlags);
// void * URLCopyHostName(void * anURL);
// void * ArrayGetValueAtIndex(void * theArray, CFIndex idx);
// void PlugInSetLoadOnDemand(void * plugIn, bool flag);
// void * TimeZoneCreateWithTimeIntervalFromGMT(void * allocator, CFTimeInterval ti);
// uint32_t BundleGetVersionNumber(void * bundle);
// void DateFormatterSetProperty(void * formatter, void * key, CFTypeRef value);
// CFTypeID MachPortGetTypeID();
// void * URLCreateFromFileSystemRepresentation(void * allocator, const uint8_t* buffer, CFIndex bufLen, bool isDirectory);
// void * UUIDCreateString(void * alloc, void * uuid);
// CFStringEncoding StringConvertWindowsCodepageToEncoding(uint32_t codepage);
// void StringTrimWhitespace(void * theString);
// bool MachPortIsValid(void * port);
// CFTypeID BinaryHeapGetTypeID();
// void AttributedStringBeginEditing(void * aStr);
// bool RunLoopSourceIsValid(void * source);
// void * URLCreateCopyDeletingPathExtension(void * allocator, void * url);
// CFIndex GetRetainCount(CFTypeRef cf);
// void * TimeZoneCreate(void * allocator, void * name, void * data);
// void * StringCreateWithFileSystemRepresentation(void * alloc, char* buffer);
// void * DateFormatterCreateISO8601Formatter(void * allocator, CFISO8601DateFormatOptions formatOptions);
// CFStringEncoding StringGetSystemEncoding();
// char* StringGetCStringPtr(void * theString, CFStringEncoding encoding);
// uint16_t SocketGetDefaultNameRegistryPortNumber();
// void * CharacterSetCreateBitmapRepresentation(void * alloc, void * theSet);
// CFAbsoluteTime AbsoluteTimeGetCurrent();
// bool StringIsSurrogateLowCharacter(UniChar character);
// bool BundleIsExecutableLoadableForURL(void * url);
// void * BundleCreateBundlesFromDirectory(void * allocator, void * directoryURL, void * bundleType);
// void SocketDisableCallBacks(void * s, CFOptionFlags callBackTypes);
// void * NumberFormatterGetFormat(void * formatter);
// void * URLCopyResourceSpecifier(void * anURL);
// void * BitVectorCreateMutableCopy(void * allocator, CFIndex capacity, void * bv);
// void * BundleCopyExecutableURL(void * bundle);
// void TimeZoneSetDefault(void * tz);
// void * CharacterSetCreateMutableCopy(void * alloc, void * theSet);
// void BundleGetPackageInfo(void * bundle, uint32_t* packageType, uint32_t* packageCreator);
// CFTypeID FileDescriptorGetTypeID();
// void * StringCreateMutableWithExternalCharactersNoCopy(void * alloc, UniChar* chars, CFIndex numChars, CFIndex capacity, void * externalCharactersAllocator);
// CFIndex StringGetLength(void * theString);
// uint32_t SwapInt32HostToLittle(uint32_t arg);
// void * URLCreateResourcePropertiesForKeysFromBookmarkData(void * allocator, void * resourcePropertiesToReturn, void * bookmark);
// void * ArrayCreateCopy(void * allocator, void * theArray);
// void StringPad(void * theString, void * padString, CFIndex length, CFIndex indexIntoPad);
// void * CalendarCopyCurrent();
// void * URLCopyQueryString(void * anURL, void * charactersToLeaveEscaped);
// void * LocaleCopyPreferredLanguages();
// void * NotificationCenterGetDarwinNotifyCenter();
// void * DataCreateWithBytesNoCopy(void * allocator, const uint8_t* bytes, CFIndex length, void * bytesDeallocator);
// void * BundleCopyBundleURL(void * bundle);
// void * AttributedStringGetString(void * aStr);
// bool PreferencesSynchronize(void * applicationID, void * userName, void * hostName);
// CFIndex RunLoopSourceGetOrder(void * source);
// void * TreeGetFirstChild(void * tree);
// void PreferencesRemoveSuitePreferencesFromApp(void * applicationID, void * suiteID);
// CFDateFormatterStyle DateFormatterGetTimeStyle(void * formatter);
// bool Equal(CFTypeRef cf1, CFTypeRef cf2);
// uint16_t SwapInt16HostToBig(uint16_t arg);
// uint16_t SwapInt16LittleToHost(uint16_t arg);
// void * AllocatorGetDefault();
// bool FileSecuritySetOwner(void * fileSec, uid_t owner);
// void * BundleGetPlugIn(void * bundle);
// bool NumberIsFloatType(void * number);
// int32_t UserNotificationDisplayNotice(CFTimeInterval timeout, CFOptionFlags flags, void * iconURL, void * soundURL, void * localizationURL, void * alertHeader, void * alertMessage, void * defaultButtonTitle);
// CFTypeID BundleGetTypeID();
// CFIndex AllocatorGetPreferredSizeForSize(void * allocator, CFIndex size, CFOptionFlags hint);
// void StringNormalize(void * theString, CFStringNormalizationForm theForm);
// void SetRemoveAllValues(void * theSet);
// CFTypeRef Retain(CFTypeRef cf);
// void * BitVectorCreateCopy(void * allocator, void * bv);
// void TreeAppendChild(void * tree, void * newChild);
// void TimeZoneResetSystem();
// CFTypeID WriteStreamGetTypeID();
// CFIndex RunLoopObserverGetOrder(void * observer);
// void BundleUnloadExecutable(void * bundle);
// bool URLIsFileReferenceURL(void * url);
// void URLStopAccessingSecurityScopedResource(void * url);
// CFTimeInterval DateGetTimeIntervalSinceDate(void * theDate, void * otherDate);
// CFFileDescriptorNativeDescriptor FileDescriptorGetNativeDescriptor(void * f);
// void * FileSecurityCreate(void * allocator);
// void * CopyTypeIDDescription(CFTypeID type_id);
// void * FileDescriptorCreateRunLoopSource(void * allocator, void * f, CFIndex order);
// CFStringTokenizerTokenType StringTokenizerGoToTokenAtIndex(void * tokenizer, CFIndex index);
// CFStringEncoding StringGetMostCompatibleMacStringEncoding(CFStringEncoding encoding);
// void * StringCreateByCombiningStrings(void * alloc, void * theArray, void * separatorString);
// void * DataCreateMutable(void * allocator, CFIndex capacity);
// void * PreferencesCopyKeyList(void * applicationID, void * userName, void * hostName);
// void * BundleCopySupportFilesDirectoryURL(void * bundle);
// CFTypeID MessagePortGetTypeID();
// void * BundleCopyPrivateFrameworksURL(void * bundle);
// void * SocketCopyPeerAddress(void * s);
// bool CharacterSetIsSupersetOfSet(void * theSet, void * theOtherset);
// void BinaryHeapRemoveMinimumValue(void * heap);
// void StringAppend(void * theString, void * appendedString);
// void * UUIDCreate(void * alloc);
// CFIndex AttributedStringGetLength(void * aStr);
// void * BundleGetLocalInfoDictionary(void * bundle);
// void * ErrorCopyUserInfo(void * err);
// void * TreeGetNextSibling(void * tree);
// void * URLCopyUserName(void * anURL);
// void * URLCreateAbsoluteURLWithBytes(void * alloc, const uint8_t* relativeURLBytes, CFIndex length, CFStringEncoding encoding, void * baseURL, bool useCompatibilityMode);
// void * DateFormatterCreateStringWithDate(void * allocator, void * formatter, void * date);
// bool CalendarGetComponentDifference(void * calendar, CFAbsoluteTime startingAT, CFAbsoluteTime resultAT, CFOptionFlags options, char* componentDesc);
// void * TreeGetChildAtIndex(void * tree, CFIndex idx);
// CFTypeID DictionaryGetTypeID();
// void * PlugInGetBundle(void * plugIn);
// void * BundleCopyResourceURLsOfTypeForLocalization(void * bundle, void * resourceType, void * subDirName, void * localizationName);
// void * UUIDGetConstantUUIDWithBytes(void * alloc, uint8_t byte0, uint8_t byte1, uint8_t byte2, uint8_t byte3, uint8_t byte4, uint8_t byte5, uint8_t byte6, uint8_t byte7, uint8_t byte8, uint8_t byte9, uint8_t byte10, uint8_t byte11, uint8_t byte12, uint8_t byte13, uint8_t byte14, uint8_t byte15);
// bool TimeZoneIsDaylightSavingTime(void * tz, CFAbsoluteTime at);
// void URLClearResourcePropertyCache(void * url);
// bool SocketIsValid(void * s);
// void * BundleGetDataPointerForName(void * bundle, void * symbolName);
// int32_t StringConvertEncodingToNSStringEncoding(CFStringEncoding encoding);
// void * SocketCopyAddress(void * s);
// void * BitVectorCreateMutable(void * allocator, CFIndex capacity);
// const uint8_t* ReadStreamGetBuffer(void * stream, CFIndex maxBytesToRead, CFIndex* numBytesRead);
// bool PreferencesAppSynchronize(void * applicationID);
// bool RunLoopIsWaiting(void * rl);
// void * URLCreateCopyDeletingLastPathComponent(void * allocator, void * url);
// const CFStringEncoding* StringGetListOfAvailableEncodings();
// void * CharacterSetCreateWithCharactersInString(void * alloc, void * theString);
// void TimeZoneSetAbbreviationDictionary(void * dict);
// void * URLCreateCopyAppendingPathExtension(void * allocator, void * url, void * extension);
// void StringUppercase(void * theString, void * locale);
// CFPropertyListRef PreferencesCopyValue(void * key, void * applicationID, void * userName, void * hostName);
// CFTypeID PlugInGetTypeID();
// void * StringCreateArrayBySeparatingStrings(void * alloc, void * theString, void * separatorString);
// void * PlugInInstanceCreate(void * allocator, void * factoryUUID, void * typeUUID);
// void * DateFormatterGetLocale(void * formatter);
// uint32_t SwapInt32LittleToHost(uint32_t arg);
// void * CharacterSetGetPredefined(CFCharacterSetPredefinedSet theSetIdentifier);
// CFTypeID DataGetTypeID();
// int32_t UserNotificationReceiveResponse(void * userNotification, CFTimeInterval timeout, CFOptionFlags* responseFlags);
// CFTypeID CharacterSetGetTypeID();
// void * NotificationCenterGetLocalCenter();
// void RunLoopSourceInvalidate(void * source);
// uint64_t SwapInt64BigToHost(uint64_t arg);
// CFIndex DictionaryGetCount(void * theDict);
// void * URLCopyPathExtension(void * url);
// bool MessagePortIsValid(void * ms);
// uint32_t SwapInt32HostToBig(uint32_t arg);
// CFIndex NumberGetByteSize(void * number);
// void * URLCreateCopyAppendingPathComponent(void * allocator, void * url, void * pathComponent, bool isDirectory);
// CFComparisonResult StringCompare(void * theString1, void * theString2, CFStringCompareFlags compareOptions);
// CFOptionFlags UserNotificationPopUpSelection(CFIndex n);
// CFByteOrder ByteOrderGetCurrent();
// bool ReadStreamOpen(void * stream);
// CFTypeID DateGetTypeID();
// uint16_t SwapInt16HostToLittle(uint16_t arg);
// void * AttributedStringCreateCopy(void * alloc, void * aStr);
// void NumberFormatterSetFormat(void * formatter, void * formatString);
// void * TimeZoneCopyKnownNames();
// void * SocketCreateRunLoopSource(void * allocator, void * s, CFIndex order);
// CFTypeID BitVectorGetTypeID();
// void CharacterSetUnion(void * theSet, void * theOtherSet);
// void * RunLoopCopyAllModes(void * rl);
// CFIndex SetGetCount(void * theSet);
// bool WriteStreamOpen(void * stream);
// void SocketSetSocketFlags(void * s, CFOptionFlags flags);
// CFTypeID DateFormatterGetTypeID();
// CFTypeID URLGetTypeID();
// void BagRemoveAllValues(void * theBag);
// void BinaryHeapRemoveAllValues(void * heap);
// void BitVectorFlipBitAtIndex(void * bv, CFIndex idx);
// CFTypeID NotificationCenterGetTypeID();
// CFIndex TreeGetChildCount(void * tree);
// CFTypeID FileSecurityGetTypeID();
// void * BundleCopyResourceURLInDirectory(void * bundleURL, void * resourceName, void * resourceType, void * subDirName);
// bool BundleIsExecutableLoadable(void * bundle);
// CFTypeID AllocatorGetTypeID();
// void URLSetTemporaryResourcePropertyForKey(void * url, void * key, CFTypeRef propertyValue);
// CFTypeID PlugInInstanceGetTypeID();
// bool PlugInRegisterPlugInType(void * factoryUUID, void * typeUUID);
// CFTypeID UUIDGetTypeID();
// void * TimeZoneGetData(void * tz);
// CFLocaleLanguageDirection LocaleGetLanguageCharacterDirection(void * isoLangCode);
// void * StringCreateWithBytesNoCopy(void * alloc, const uint8_t* bytes, CFIndex numBytes, CFStringEncoding encoding, bool isExternalRepresentation, void * contentsDeallocator);
// void TreeRemove(void * tree);
// bool BundleLoadExecutable(void * bundle);
// bool RunLoopObserverDoesRepeat(void * observer);
// void RunLoopTimerInvalidate(void * timer);
// CFAbsoluteTime TimeZoneGetNextDaylightSavingTimeTransition(void * tz, CFAbsoluteTime at);
// void * BinaryHeapGetMinimum(void * heap);
// void * BundleCopyPreferredLocalizationsFromArray(void * locArray);
// void * MessagePortCreateRunLoopSource(void * allocator, void * local, CFIndex order);
// CFTypeID LocaleGetTypeID();
// bool CalendarComposeAbsoluteTime(void * calendar, CFAbsoluteTime* at, char* componentDesc);
// bool PlugInUnregisterPlugInType(void * factoryUUID, void * typeUUID);
// bool ReadStreamHasBytesAvailable(void * stream);
// UniChar StringGetCharacterAtIndex(void * theString, CFIndex idx);
// CFIndex CalendarGetOrdinalityOfUnit(void * calendar, CFCalendarUnit smallerUnit, CFCalendarUnit biggerUnit, CFAbsoluteTime at);
// CFPropertyListRef PropertyListCreateDeepCopy(void * allocator, CFPropertyListRef propertyList, CFOptionFlags mutabilityOption);
// void * CalendarCopyTimeZone(void * calendar);
// int32_t URLGetPortNumber(void * anURL);
// void * LocaleCopyISOCountryCodes();
// void * BinaryHeapCreateCopy(void * allocator, CFIndex capacity, void * heap);
// bool FileSecurityGetGroup(void * fileSec, gid_t* group);
// void * WriteStreamCreateWithFile(void * alloc, void * fileURL);
// void CharacterSetInvert(void * theSet);
// CFTypeID ErrorGetTypeID();
// void * BundleGetFunctionPointerForName(void * bundle, void * functionName);
// void StringFold(void * theString, CFStringCompareFlags theFlags, void * theLocale);
// bool URLStartAccessingSecurityScopedResource(void * url);
// void * StringCreateExternalRepresentation(void * alloc, void * theString, CFStringEncoding encoding, uint8_t lossByte);
// CFSocketError SocketConnectToAddress(void * s, void * address, CFTimeInterval timeout);
// void StringAppendCString(void * theString, char* cStr, CFStringEncoding encoding);
// CFDateFormatterStyle DateFormatterGetDateStyle(void * formatter);
// void * BundleCopyLocalizationsForURL(void * url);
// void * UserNotificationGetResponseDictionary(void * userNotification);
// CFIndex BinaryHeapGetCount(void * heap);
// CFTypeID UserNotificationGetTypeID();
// void * BundleCopyResourceURLsOfType(void * bundle, void * resourceType, void * subDirName);
// CFTypeID RunLoopSourceGetTypeID();
// void StringInsert(void * str, CFIndex idx, void * insertedStr);
import "C"
import (
	"unsafe"
)

// Creates a new mutable dictionary with the key-value pairs from another dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1516786-cfdictionarycreatemutablecopy?language=objc
func DictionaryCreateMutableCopy(allocator AllocatorRef, capacity Index, theDict DictionaryRef) unsafe.Pointer {
	rv := C.DictionaryCreateMutableCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(capacity),
		// *typing.RefType
		unsafe.Pointer(theDict),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns the type identifier for the CFNull opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521289-cfnullgettypeid?language=objc
func NullGetTypeID() TypeID {
	rv := C.NullGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the error associated with a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539634-cfwritestreamcopyerror?language=objc
func WriteStreamCopyError(stream WriteStreamRef) ErrorRef {
	rv := C.WriteStreamCopyError(
		// *typing.RefType
		unsafe.Pointer(stream),
	)
	// *typing.RefType
	return ErrorRef(rv)
}

// Returns the location of a localized resource in a bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537130-cfbundlecopyresourceurlforlocali?language=objc
func BundleCopyResourceURLForLocalization(bundle BundleRef, resourceName StringRef, resourceType StringRef, subDirName StringRef, localizationName StringRef) URLRef {
	rv := C.BundleCopyResourceURLForLocalization(
		// *typing.RefType
		unsafe.Pointer(bundle),
		// *typing.RefType
		unsafe.Pointer(resourceName),
		// *typing.RefType
		unsafe.Pointer(resourceType),
		// *typing.RefType
		unsafe.Pointer(subDirName),
		// *typing.RefType
		unsafe.Pointer(localizationName),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Creates an immutable copy of a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542087-cfstringcreatecopy?language=objc
func StringCreateCopy(alloc AllocatorRef, theString StringRef) StringRef {
	rv := C.StringCreateCopy(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(theString),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Sets the first weekday for a calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533524-cfcalendarsetfirstweekday?language=objc
func CalendarSetFirstWeekday(calendar CalendarRef, wkdy Index) {
	C.CalendarSetFirstWeekday(
		// *typing.RefType
		unsafe.Pointer(calendar),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(wkdy),
	)
}

// Returns the type identifier for the CFNumberFormatter opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1390736-cfnumberformattergettypeid?language=objc
func NumberFormatterGetTypeID() TypeID {
	rv := C.NumberFormatterGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/3684868-cfbundleisarchitectureloadable?language=objc
func BundleIsArchitectureLoadable(arch int) bool {
	rv := C.BundleIsArchitectureLoadable(
		// *typing.AliasType
		// *typing.AliasType
		(C.cpu_type_t)(arch),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a bundle’s information dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537135-cfbundlecopyinfodictionaryindire?language=objc
func BundleCopyInfoDictionaryInDirectory(bundleURL URLRef) DictionaryRef {
	rv := C.BundleCopyInfoDictionaryInDirectory(
		// *typing.RefType
		unsafe.Pointer(bundleURL),
	)
	// *typing.RefType
	return DictionaryRef(rv)
}

// Swaps the bytes of a 32-bit integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425262-cfswapint32?language=objc
func SwapInt32(arg uint32) uint32 {
	rv := C.SwapInt32(
		// *typing.PrimitiveType
		C.uint32_t(arg),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Returns the path portion of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542952-cfurlcopystrictpath?language=objc
func URLCopyStrictPath(anURL URLRef, isAbsolute *bool) StringRef {
	rv := C.URLCopyStrictPath(
		// *typing.RefType
		unsafe.Pointer(anURL),
		// *typing.PointerType
		(*C.bool)(unsafe.Pointer(&isAbsolute)),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the time zone currently used by the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543660-cftimezonecopysystem?language=objc
func TimeZoneCopySystem() TimeZoneRef {
	rv := C.TimeZoneCopySystem()
	// *typing.RefType
	return TimeZoneRef(rv)
}

// Creates a copy of a CFFileSecurityRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426498-cffilesecuritycreatecopy?language=objc
func FileSecurityCreateCopy(allocator AllocatorRef, fileSec FileSecurityRef) FileSecurityRef {
	rv := C.FileSecurityCreateCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(fileSec),
	)
	// *typing.RefType
	return FileSecurityRef(rv)
}

// Creates a CFURL object using a local file system path string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543250-cfurlcreatewithfilesystempath?language=objc
func URLCreateWithFileSystemPath(allocator AllocatorRef, filePath StringRef, pathStyle URLPathStyle, isDirectory bool) URLRef {
	rv := C.URLCreateWithFileSystemPath(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(filePath),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFURLPathStyle)(pathStyle),
		// *typing.PrimitiveType
		C.bool(isDirectory),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns the number of values currently in a bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1469310-cfbaggetcount?language=objc
func BagGetCount(theBag BagRef) Index {
	rv := C.BagGetCount(
		// *typing.RefType
		unsafe.Pointer(theBag),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns the location of a bundle’s auxiliary executable code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537098-cfbundlecopyauxiliaryexecutableu?language=objc
func BundleCopyAuxiliaryExecutableURL(bundle BundleRef, executableName StringRef) URLRef {
	rv := C.BundleCopyAuxiliaryExecutableURL(
		// *typing.RefType
		unsafe.Pointer(bundle),
		// *typing.RefType
		unsafe.Pointer(executableName),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Creates an immutable bag with the values of another bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1469264-cfbagcreatecopy?language=objc
func BagCreateCopy(allocator AllocatorRef, theBag BagRef) BagRef {
	rv := C.BagCreateCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(theBag),
	)
	// *typing.RefType
	return BagRef(rv)
}

// Closes a writable stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539729-cfwritestreamclose?language=objc
func WriteStreamClose(stream WriteStreamRef) {
	C.WriteStreamClose(
		// *typing.RefType
		unsafe.Pointer(stream),
	)
}

// Convenience function that directly obtains a Boolean preference value for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515514-cfpreferencesgetappbooleanvalue?language=objc
func PreferencesGetAppBooleanValue(key StringRef, applicationID StringRef, keyExistsAndHasValidFormat *bool) bool {
	rv := C.PreferencesGetAppBooleanValue(
		// *typing.RefType
		unsafe.Pointer(key),
		// *typing.RefType
		unsafe.Pointer(applicationID),
		// *typing.PointerType
		(*C.bool)(unsafe.Pointer(&keyExistsAndHasValidFormat)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates an immutable set containing the values of an existing set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1520420-cfsetcreatecopy?language=objc
func SetCreateCopy(allocator AllocatorRef, theSet SetRef) SetRef {
	rv := C.SetCreateCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(theSet),
	)
	// *typing.RefType
	return SetRef(rv)
}

// Sets the format string of the given date formatter to the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1396276-cfdateformattersetformat?language=objc
func DateFormatterSetFormat(formatter DateFormatterRef, formatString StringRef) {
	C.DateFormatterSetFormat(
		// *typing.RefType
		unsafe.Pointer(formatter),
		// *typing.RefType
		unsafe.Pointer(formatString),
	)
}

// Removes all the children of a tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401769-cftreeremoveallchildren?language=objc
func TreeRemoveAllChildren(tree TreeRef) {
	C.TreeRemoveAllChildren(
		// *typing.RefType
		unsafe.Pointer(tree),
	)
}

// Creates a new immutable character set with the bitmap representation specified by given data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542964-cfcharactersetcreatewithbitmapre?language=objc
func CharacterSetCreateWithBitmapRepresentation(alloc AllocatorRef, theData DataRef) CharacterSetRef {
	rv := C.CharacterSetCreateWithBitmapRepresentation(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(theData),
	)
	// *typing.RefType
	return CharacterSetRef(rv)
}

// Returns the bit value at a given index in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542874-cfbitvectorgetbitatindex?language=objc
func BitVectorGetBitAtIndex(bv BitVectorRef, idx Index) Bit {
	rv := C.BitVectorGetBitAtIndex(
		// *typing.RefType
		unsafe.Pointer(bv),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(idx),
	)
	// *typing.AliasType
	return Bit(rv)
}

// Returns the allocator used to allocate a Core Foundation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521280-cfgetallocator?language=objc
func GetAllocator(cf TypeRef) AllocatorRef {
	rv := C.GetAllocator(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(cf),
	)
	// *typing.RefType
	return AllocatorRef(rv)
}

// Returns the location of a bundle’s built in plug-in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537106-cfbundlecopybuiltinpluginsurl?language=objc
func BundleCopyBuiltInPlugInsURL(bundle BundleRef) URLRef {
	rv := C.BundleCopyBuiltInPlugInsURL(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns the error code for a given CFError. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1494656-cferrorgetcode?language=objc
func ErrorGetCode(err ErrorRef) Index {
	rv := C.ErrorGetCode(
		// *typing.RefType
		unsafe.Pointer(err),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns by reference the byte representation of a URL object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541551-cfurlgetbytes?language=objc
func URLGetBytes(url URLRef, buffer *uint8, bufferLength Index) Index {
	rv := C.URLGetBytes(
		// *typing.RefType
		unsafe.Pointer(url),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&buffer)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(bufferLength),
	)
	// *typing.AliasType
	return Index(rv)
}

// Sets the value of a particular bit in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542852-cfbitvectorsetbitatindex?language=objc
func BitVectorSetBitAtIndex(bv unsafe.Pointer, idx Index, value Bit) {
	C.BitVectorSetBitAtIndex(
		// *typing.RefType
		unsafe.Pointer(bv),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(idx),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFBit)(value),
	)
}

// Registers a new instance of a type with CFPlugIn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493891-cfpluginaddinstanceforfactory?language=objc
func PlugInAddInstanceForFactory(factoryID UUIDRef) {
	C.PlugInAddInstanceForFactory(
		// *typing.RefType
		unsafe.Pointer(factoryID),
	)
}

// Returns the minimum number of days in the first week of a specified calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533477-cfcalendargetminimumdaysinfirstw?language=objc
func CalendarGetMinimumDaysInFirstWeek(calendar CalendarRef) Index {
	rv := C.CalendarGetMinimumDaysInFirstWeek(
		// *typing.RefType
		unsafe.Pointer(calendar),
	)
	// *typing.AliasType
	return Index(rv)
}

// Increases the length of a CFMutableData object's internal byte buffer, zero-filling the extension to the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542081-cfdataincreaselength?language=objc
func DataIncreaseLength(theData unsafe.Pointer, extraLength Index) {
	C.DataIncreaseLength(
		// *typing.RefType
		unsafe.Pointer(theData),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(extraLength),
	)
}

// Returns a Boolean value that indicates whether a CFRunLoopTimer object is valid and able to fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543110-cfrunlooptimerisvalid?language=objc
func RunLoopTimerIsValid(timer RunLoopTimerRef) bool {
	rv := C.RunLoopTimerIsValid(
		// *typing.RefType
		unsafe.Pointer(timer),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the location of a resource contained in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537117-cfbundlecopyresourceurl?language=objc
func BundleCopyResourceURL(bundle BundleRef, resourceName StringRef, resourceType StringRef, subDirName StringRef) URLRef {
	rv := C.BundleCopyResourceURL(
		// *typing.RefType
		unsafe.Pointer(bundle),
		// *typing.RefType
		unsafe.Pointer(resourceName),
		// *typing.RefType
		unsafe.Pointer(resourceType),
		// *typing.RefType
		unsafe.Pointer(subDirName),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Creates a string from a buffer of Unicode characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541779-cfstringcreatewithcharacters?language=objc
func StringCreateWithCharacters(alloc AllocatorRef, chars *uint16, numChars Index) StringRef {
	rv := C.StringCreateWithCharacters(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.PointerType
		(*C.UniChar)(unsafe.Pointer(&chars)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(numChars),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Reads data from a readable stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539700-cfreadstreamread?language=objc
func ReadStreamRead(stream ReadStreamRef, buffer *uint8, bufferLength Index) Index {
	rv := C.ReadStreamRead(
		// *typing.RefType
		unsafe.Pointer(stream),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&buffer)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(bufferLength),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns the type identifier for the CFRunLoopObserver opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543288-cfrunloopobservergettypeid?language=objc
func RunLoopObserverGetTypeID() TypeID {
	rv := C.RunLoopObserverGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Creates a new CFNumberFormatter object, localized to the given locale, which will format numbers to the given style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1390708-cfnumberformattercreate?language=objc
func NumberFormatterCreate(allocator AllocatorRef, locale LocaleRef, style NumberFormatterStyle) NumberFormatterRef {
	rv := C.NumberFormatterCreate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(locale),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFNumberFormatterStyle)(style),
	)
	// *typing.RefType
	return NumberFormatterRef(rv)
}

// Returns the locale object used to create the given number formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1390690-cfnumberformattergetlocale?language=objc
func NumberFormatterGetLocale(formatter NumberFormatterRef) LocaleRef {
	rv := C.NumberFormatterGetLocale(
		// *typing.RefType
		unsafe.Pointer(formatter),
	)
	// *typing.RefType
	return LocaleRef(rv)
}

// Returns a CFMessagePort object connected to a remote port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542648-cfmessageportcreateremote?language=objc
func MessagePortCreateRemote(allocator AllocatorRef, name StringRef) MessagePortRef {
	rv := C.MessagePortCreateRemote(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(name),
	)
	// *typing.RefType
	return MessagePortRef(rv)
}

// Returns the run loop stages during which an observer runs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542898-cfrunloopobservergetactivities?language=objc
func RunLoopObserverGetActivities(observer RunLoopObserverRef) OptionFlags {
	rv := C.RunLoopObserverGetActivities(
		// *typing.RefType
		unsafe.Pointer(observer),
	)
	// *typing.AliasType
	return OptionFlags(rv)
}

// Determines if a property list is valid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1430019-cfpropertylistisvalid?language=objc
func PropertyListIsValid(plist PropertyListRef, format PropertyListFormat) bool {
	rv := C.PropertyListIsValid(
		// *typing.AliasType
		// *typing.AliasType
		(C.CFPropertyListRef)(plist),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFPropertyListFormat)(format),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the firing interval of a repeating CFRunLoopTimer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543291-cfrunlooptimergetinterval?language=objc
func RunLoopTimerGetInterval(timer RunLoopTimerRef) TimeInterval {
	rv := C.RunLoopTimerGetInterval(
		// *typing.RefType
		unsafe.Pointer(timer),
	)
	// *typing.AliasType
	return TimeInterval(rv)
}

// Returns the primary double value represented by a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543293-cfstringgetdoublevalue?language=objc
func StringGetDoubleValue(str StringRef) float64 {
	rv := C.StringGetDoubleValue(
		// *typing.RefType
		unsafe.Pointer(str),
	)
	// *typing.PrimitiveType
	return float64(rv)
}

// Returns a dictionary holding the mappings of time zone abbreviations to time zone names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541799-cftimezonecopyabbreviationdictio?language=objc
func TimeZoneCopyAbbreviationDictionary() DictionaryRef {
	rv := C.TimeZoneCopyAbbreviationDictionary()
	// *typing.RefType
	return DictionaryRef(rv)
}

// Enables callbacks for a given CFFileDescriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1477610-cffiledescriptorenablecallbacks?language=objc
func FileDescriptorEnableCallBacks(f FileDescriptorRef, callBackTypes OptionFlags) {
	C.FileDescriptorEnableCallBacks(
		// *typing.RefType
		unsafe.Pointer(f),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(callBackTypes),
	)
}

// Returns a bundle’s information dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537165-cfbundlegetinfodictionary?language=objc
func BundleGetInfoDictionary(bundle BundleRef) DictionaryRef {
	rv := C.BundleGetInfoDictionary(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return DictionaryRef(rv)
}

// Creates an attributed string with specified string and attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542924-cfattributedstringcreate?language=objc
func AttributedStringCreate(alloc AllocatorRef, str StringRef, attributes DictionaryRef) AttributedStringRef {
	rv := C.AttributedStringCreate(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(str),
		// *typing.RefType
		unsafe.Pointer(attributes),
	)
	// *typing.RefType
	return AttributedStringRef(rv)
}

// Converts a 64-bit integer from the host’s native byte order to little-endian format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425297-cfswapint64hosttolittle?language=objc
func SwapInt64HostToLittle(arg uint64) uint64 {
	rv := C.SwapInt64HostToLittle(
		// *typing.PrimitiveType
		C.uint64_t(arg),
	)
	// *typing.PrimitiveType
	return uint64(rv)
}

// Invalidates a CFFileDescriptor object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1477587-cffiledescriptorinvalidate?language=objc
func FileDescriptorInvalidate(f FileDescriptorRef) {
	C.FileDescriptorInvalidate(
		// *typing.RefType
		unsafe.Pointer(f),
	)
}

// Creates an immutable CFData object using data copied from a specified byte buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542359-cfdatacreate?language=objc
func DataCreate(allocator AllocatorRef, bytes *uint8, length Index) DataRef {
	rv := C.DataCreate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&bytes)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(length),
	)
	// *typing.RefType
	return DataRef(rv)
}

// Returns the information dictionary for a given URL location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537134-cfbundlecopyinfodictionaryforurl?language=objc
func BundleCopyInfoDictionaryForURL(url URLRef) DictionaryRef {
	rv := C.BundleCopyInfoDictionaryForURL(
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.RefType
	return DictionaryRef(rv)
}

// Returns the name of the IANA registry “charset” that is the closest mapping to a specified string encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542710-cfstringconvertencodingtoianacha?language=objc
func StringConvertEncodingToIANACharSetName(encoding StringEncoding) StringRef {
	rv := C.StringConvertEncodingToIANACharSetName(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the URL as a CFString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542826-cfurlgetstring?language=objc
func URLGetString(anURL URLRef) StringRef {
	rv := C.URLGetString(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Removes all the key-value pairs from a dictionary, making it empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1516800-cfdictionaryremoveallvalues?language=objc
func DictionaryRemoveAllValues(theDict unsafe.Pointer) {
	C.DictionaryRemoveAllValues(
		// *typing.RefType
		unsafe.Pointer(theDict),
	)
}

// Creates an immutable copy of a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541991-cfdatacreatecopy?language=objc
func DataCreateCopy(allocator AllocatorRef, theData DataRef) DataRef {
	rv := C.DataCreateCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(theData),
	)
	// *typing.RefType
	return DataRef(rv)
}

// Obtains information about the load status for a bundle’s main executable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537151-cfbundleisexecutableloaded?language=objc
func BundleIsExecutableLoaded(bundle BundleRef) bool {
	rv := C.BundleIsExecutableLoaded(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the current state of a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539623-cfreadstreamgetstatus?language=objc
func ReadStreamGetStatus(stream ReadStreamRef) StreamStatus {
	rv := C.ReadStreamGetStatus(
		// *typing.RefType
		unsafe.Pointer(stream),
	)
	// *typing.AliasType
	return StreamStatus(rv)
}

// Creates and returns a directory enumerator with provided enumerator behavior options and properties to be prefetched. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543646-cfurlenumeratorcreatefordirector?language=objc
func URLEnumeratorCreateForDirectoryURL(alloc AllocatorRef, directoryURL URLRef, option URLEnumeratorOptions, propertyKeys ArrayRef) URLEnumeratorRef {
	rv := C.URLEnumeratorCreateForDirectoryURL(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(directoryURL),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFURLEnumeratorOptions)(option),
		// *typing.RefType
		unsafe.Pointer(propertyKeys),
	)
	// *typing.RefType
	return URLEnumeratorRef(rv)
}

// Returns a copy of the logical locale for the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542508-cflocalecopycurrent?language=objc
func LocaleCopyCurrent() LocaleRef {
	rv := C.LocaleCopyCurrent()
	// *typing.RefType
	return LocaleRef(rv)
}

// Returns the abbreviation of a time zone at a specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543329-cftimezonecopyabbreviation?language=objc
func TimeZoneCopyAbbreviation(tz TimeZoneRef, at AbsoluteTime) StringRef {
	rv := C.TimeZoneCopyAbbreviation(
		// *typing.RefType
		unsafe.Pointer(tz),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(at),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Converts a 16-bit integer from big-endian format to the host’s native byte order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425282-cfswapint16bigtohost?language=objc
func SwapInt16BigToHost(arg uint16) uint16 {
	rv := C.SwapInt16BigToHost(
		// *typing.PrimitiveType
		C.uint16_t(arg),
	)
	// *typing.PrimitiveType
	return uint16(rv)
}

// Returns the type identifier for the CFTimeZone opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542256-cftimezonegettypeid?language=objc
func TimeZoneGetTypeID() TypeID {
	rv := C.TimeZoneGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns a dictionary containing preference values for multiple keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515498-cfpreferencescopymultiple?language=objc
func PreferencesCopyMultiple(keysToFetch ArrayRef, applicationID StringRef, userName StringRef, hostName StringRef) DictionaryRef {
	rv := C.PreferencesCopyMultiple(
		// *typing.RefType
		unsafe.Pointer(keysToFetch),
		// *typing.RefType
		unsafe.Pointer(applicationID),
		// *typing.RefType
		unsafe.Pointer(userName),
		// *typing.RefType
		unsafe.Pointer(hostName),
	)
	// *typing.RefType
	return DictionaryRef(rv)
}

// Returns the value of a resource property from specified bookmark data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543031-cfurlcreateresourcepropertyforke?language=objc
func URLCreateResourcePropertyForKeyFromBookmarkData(allocator AllocatorRef, resourcePropertyKey StringRef, bookmark DataRef) TypeRef {
	rv := C.URLCreateResourcePropertyForKeyFromBookmarkData(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(resourcePropertyKey),
		// *typing.RefType
		unsafe.Pointer(bookmark),
	)
	// *typing.AliasType
	return TypeRef(rv)
}

// Sets the name of a local CFMessagePort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543116-cfmessageportsetname?language=objc
func MessagePortSetName(ms MessagePortRef, newName StringRef) bool {
	rv := C.MessagePortSetName(
		// *typing.RefType
		unsafe.Pointer(ms),
		// *typing.RefType
		unsafe.Pointer(newName),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the root tree of a given tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401763-cftreefindroot?language=objc
func TreeFindRoot(tree TreeRef) TreeRef {
	rv := C.TreeFindRoot(
		// *typing.RefType
		unsafe.Pointer(tree),
	)
	// *typing.RefType
	return TreeRef(rv)
}

// Creates a new immutable character set that is the invert of the specified character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542432-cfcharactersetcreateinvertedset?language=objc
func CharacterSetCreateInvertedSet(alloc AllocatorRef, theSet CharacterSetRef) CharacterSetRef {
	rv := C.CharacterSetCreateInvertedSet(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(theSet),
	)
	// *typing.RefType
	return CharacterSetRef(rv)
}

// Updates a displayed user notification dialog with new user interface information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534529-cfusernotificationupdate?language=objc
func UserNotificationUpdate(userNotification UserNotificationRef, timeout TimeInterval, flags OptionFlags, dictionary DictionaryRef) int32 {
	rv := C.UserNotificationUpdate(
		// *typing.RefType
		unsafe.Pointer(userNotification),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFTimeInterval)(timeout),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(flags),
		// *typing.RefType
		unsafe.Pointer(dictionary),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Returns a human presentable recovery suggestion for a given error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1494650-cferrorcopyrecoverysuggestion?language=objc
func ErrorCopyRecoverySuggestion(err ErrorRef) StringRef {
	rv := C.ErrorCopyRecoverySuggestion(
		// *typing.RefType
		unsafe.Pointer(err),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Creates and returns a volume enumerator with provided enumerator behavior options and properties to be prefetched. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542110-cfurlenumeratorcreateformountedv?language=objc
func URLEnumeratorCreateForMountedVolumes(alloc AllocatorRef, option URLEnumeratorOptions, propertyKeys ArrayRef) URLEnumeratorRef {
	rv := C.URLEnumeratorCreateForMountedVolumes(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFURLEnumeratorOptions)(option),
		// *typing.RefType
		unsafe.Pointer(propertyKeys),
	)
	// *typing.RefType
	return URLEnumeratorRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543275-cfrunlooptimergettolerance?language=objc
func RunLoopTimerGetTolerance(timer RunLoopTimerRef) TimeInterval {
	rv := C.RunLoopTimerGetTolerance(
		// *typing.RefType
		unsafe.Pointer(timer),
	)
	// *typing.AliasType
	return TimeInterval(rv)
}

// Returns a localized string from a bundle’s strings file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537103-cfbundlecopylocalizedstring?language=objc
func BundleCopyLocalizedString(bundle BundleRef, key StringRef, value StringRef, tableName StringRef) StringRef {
	rv := C.BundleCopyLocalizedString(
		// *typing.RefType
		unsafe.Pointer(bundle),
		// *typing.RefType
		unsafe.Pointer(key),
		// *typing.RefType
		unsafe.Pointer(value),
		// *typing.RefType
		unsafe.Pointer(tableName),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the CFRunLoop object for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542428-cfrunloopgetcurrent?language=objc
func RunLoopGetCurrent() RunLoopRef {
	rv := C.RunLoopGetCurrent()
	// *typing.RefType
	return RunLoopRef(rv)
}

// Adds suite preferences to an application’s preference search chain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515518-cfpreferencesaddsuitepreferences?language=objc
func PreferencesAddSuitePreferencesToApp(applicationID StringRef, suiteID StringRef) {
	C.PreferencesAddSuitePreferencesToApp(
		// *typing.RefType
		unsafe.Pointer(applicationID),
		// *typing.RefType
		unsafe.Pointer(suiteID),
	)
}

// Creates a new character set with the values from a given character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543462-cfcharactersetcreatecopy?language=objc
func CharacterSetCreateCopy(alloc AllocatorRef, theSet CharacterSetRef) CharacterSetRef {
	rv := C.CharacterSetCreateCopy(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(theSet),
	)
	// *typing.RefType
	return CharacterSetRef(rv)
}

// Releases a Core Foundation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521153-cfrelease?language=objc
func Release(cf TypeRef) {
	C.Release(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(cf),
	)
}

// Returns a Boolean value that indicates whether the native file descriptor for a given CFFileDescriptor is valid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1477575-cffiledescriptorisvalid?language=objc
func FileDescriptorIsValid(f FileDescriptorRef) bool {
	rv := C.FileDescriptorIsValid(
		// *typing.RefType
		unsafe.Pointer(f),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a textual description of a Core Foundation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521252-cfcopydescription?language=objc
func CopyDescription(cf TypeRef) StringRef {
	rv := C.CopyDescription(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(cf),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the type identifier of the CFTree opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401798-cftreegettypeid?language=objc
func TreeGetTypeID() TypeID {
	rv := C.TreeGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Inserts a new sibling after a given tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401777-cftreeinsertsibling?language=objc
func TreeInsertSibling(tree TreeRef, newSibling TreeRef) {
	C.TreeInsertSibling(
		// *typing.RefType
		unsafe.Pointer(tree),
		// *typing.RefType
		unsafe.Pointer(newSibling),
	)
}

// Creates a writable stream for a growable block of memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539694-cfwritestreamcreatewithallocated?language=objc
func WriteStreamCreateWithAllocatedBuffers(alloc AllocatorRef, bufferAllocator AllocatorRef) WriteStreamRef {
	rv := C.WriteStreamCreateWithAllocatedBuffers(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(bufferAllocator),
	)
	// *typing.RefType
	return WriteStreamRef(rv)
}

// Sets the time zone for a calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533517-cfcalendarsettimezone?language=objc
func CalendarSetTimeZone(calendar CalendarRef, tz TimeZoneRef) {
	C.CalendarSetTimeZone(
		// *typing.RefType
		unsafe.Pointer(calendar),
		// *typing.RefType
		unsafe.Pointer(tz),
	)
}

// Sets the group UUID associated with a CFFileSecurityRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426492-cffilesecuritysetgroupuuid?language=objc
func FileSecuritySetGroupUUID(fileSec FileSecurityRef, groupUUID UUIDRef) bool {
	rv := C.FileSecuritySetGroupUUID(
		// *typing.RefType
		unsafe.Pointer(fileSec),
		// *typing.RefType
		unsafe.Pointer(groupUUID),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a new empty mutable character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543715-cfcharactersetcreatemutable?language=objc
func CharacterSetCreateMutable(alloc AllocatorRef) unsafe.Pointer {
	rv := C.CharacterSetCreateMutable(
		// *typing.RefType
		unsafe.Pointer(alloc),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Sets the group ID associated with a CFFileSecurityRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426524-cffilesecuritysetgroup?language=objc
func FileSecuritySetGroup(fileSec FileSecurityRef, group int) bool {
	rv := C.FileSecuritySetGroup(
		// *typing.RefType
		unsafe.Pointer(fileSec),
		// *typing.AliasType
		// *typing.StructType
		(C.gid_t)(group),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the current state of a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539731-cfwritestreamgetstatus?language=objc
func WriteStreamGetStatus(stream WriteStreamRef) StreamStatus {
	rv := C.WriteStreamGetStatus(
		// *typing.RefType
		unsafe.Pointer(stream),
	)
	// *typing.AliasType
	return StreamStatus(rv)
}

// Creates a mutable copy of an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543654-cfattributedstringcreatemutablec?language=objc
func AttributedStringCreateMutableCopy(alloc AllocatorRef, maxLength Index, aStr AttributedStringRef) unsafe.Pointer {
	rv := C.AttributedStringCreateMutableCopy(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(maxLength),
		// *typing.RefType
		unsafe.Pointer(aStr),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns the geopolitical region name that identifies a given time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543519-cftimezonegetname?language=objc
func TimeZoneGetName(tz TimeZoneRef) StringRef {
	rv := C.TimeZoneGetName(
		// *typing.RefType
		unsafe.Pointer(tz),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns a bundle’s package type and creator without having to create a CFBundle object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537156-cfbundlegetpackageinfoindirector?language=objc
func BundleGetPackageInfoInDirectory(url URLRef, packageType *uint32, packageCreator *uint32) bool {
	rv := C.BundleGetPackageInfoInDirectory(
		// *typing.RefType
		unsafe.Pointer(url),
		// *typing.PointerType
		(*C.uint32_t)(unsafe.Pointer(&packageType)),
		// *typing.PointerType
		(*C.uint32_t)(unsafe.Pointer(&packageCreator)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Exchanges the values at two indices of an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1388739-cfarrayexchangevaluesatindices?language=objc
func ArrayExchangeValuesAtIndices(theArray unsafe.Pointer, idx1 Index, idx2 Index) {
	C.ArrayExchangeValuesAtIndices(
		// *typing.RefType
		unsafe.Pointer(theArray),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(idx1),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(idx2),
	)
}

// Makes a newly-allocated Core Foundation object eligible for garbage collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521163-cfmakecollectable?language=objc
func MakeCollectable(cf TypeRef) TypeRef {
	rv := C.MakeCollectable(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(cf),
	)
	// *typing.AliasType
	return TypeRef(rv)
}

// Sets the minimum number of days in the first week of a specified calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533479-cfcalendarsetminimumdaysinfirstw?language=objc
func CalendarSetMinimumDaysInFirstWeek(calendar CalendarRef, mwd Index) {
	C.CalendarSetMinimumDaysInFirstWeek(
		// *typing.RefType
		unsafe.Pointer(calendar),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(mwd),
	)
}

// Forms an intersection of two character sets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543260-cfcharactersetintersect?language=objc
func CharacterSetIntersect(theSet unsafe.Pointer, theOtherSet CharacterSetRef) {
	C.CharacterSetIntersect(
		// *typing.RefType
		unsafe.Pointer(theSet),
		// *typing.RefType
		unsafe.Pointer(theOtherSet),
	)
}

// Given a CFString object containing XML source with escaped entities, returns a string with specified XML entities unescaped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541542-cfxmlcreatestringbyunescapingent?language=objc
func XMLCreateStringByUnescapingEntities(allocator AllocatorRef, string_ StringRef, entitiesDictionary DictionaryRef) StringRef {
	rv := C.XMLCreateStringByUnescapingEntities(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(string_),
		// *typing.RefType
		unsafe.Pointer(entitiesDictionary),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the path portion of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541982-cfurlcopypath?language=objc
func URLCopyPath(anURL URLRef) StringRef {
	rv := C.URLCopyPath(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Enables the callback function of a CFSocket object for certain types of socket activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543029-cfsocketenablecallbacks?language=objc
func SocketEnableCallBacks(s SocketRef, callBackTypes OptionFlags) {
	C.SocketEnableCallBacks(
		// *typing.RefType
		unsafe.Pointer(s),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(callBackTypes),
	)
}

// Closes a readable stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539712-cfreadstreamclose?language=objc
func ReadStreamClose(stream ReadStreamRef) {
	C.ReadStreamClose(
		// *typing.RefType
		unsafe.Pointer(stream),
	)
}

// Returns the type identifier for the CFCalendar opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533522-cfcalendargettypeid?language=objc
func CalendarGetTypeID() TypeID {
	rv := C.CalendarGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the bundle identifier from a bundle’s information property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537105-cfbundlegetidentifier?language=objc
func BundleGetIdentifier(bundle BundleRef) StringRef {
	rv := C.BundleGetIdentifier(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the password of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543510-cfurlcopypassword?language=objc
func URLCopyPassword(anURL URLRef) StringRef {
	rv := C.URLCopyPassword(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the integer value represented by a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541939-cfstringgetintvalue?language=objc
func StringGetIntValue(str StringRef) int32 {
	rv := C.StringGetIntValue(
		// *typing.RefType
		unsafe.Pointer(str),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Gets the file mode associated with a CFFileSecurityRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426517-cffilesecuritygetmode?language=objc
func FileSecurityGetMode(fileSec FileSecurityRef, mode *int) bool {
	rv := C.FileSecurityGetMode(
		// *typing.RefType
		unsafe.Pointer(fileSec),
		// *typing.PointerType
		(*C.mode_t)(unsafe.Pointer(&mode)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the type identifier for the CFRunLoop opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543626-cfrunloopgettypeid?language=objc
func RunLoopGetTypeID() TypeID {
	rv := C.RunLoopGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Tells a recursive enumerator not to descend into the directory at the URL that was returned by the most recent call to the CFURLEnumeratorGetNextURL function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541451-cfurlenumeratorskipdescendents?language=objc
func URLEnumeratorSkipDescendents(enumerator URLEnumeratorRef) {
	C.URLEnumeratorSkipDescendents(
		// *typing.RefType
		unsafe.Pointer(enumerator),
	)
}

// Returns the type identifier of the CFRunLoopTimer opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542912-cfrunlooptimergettypeid?language=objc
func RunLoopTimerGetTypeID() TypeID {
	rv := C.RunLoopTimerGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the path portion of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541581-cfurlcopyfilesystempath?language=objc
func URLCopyFileSystemPath(anURL URLRef, pathStyle URLPathStyle) StringRef {
	rv := C.URLCopyFileSystemPath(
		// *typing.RefType
		unsafe.Pointer(anURL),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFURLPathStyle)(pathStyle),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Removes all the values from an array, making it empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1388798-cfarrayremoveallvalues?language=objc
func ArrayRemoveAllValues(theArray unsafe.Pointer) {
	C.ArrayRemoveAllValues(
		// *typing.RefType
		unsafe.Pointer(theArray),
	)
}

// Creates a CFRunLoopSource object for a CFMachPort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1400928-cfmachportcreaterunloopsource?language=objc
func MachPortCreateRunLoopSource(allocator AllocatorRef, port MachPortRef, order Index) RunLoopSourceRef {
	rv := C.MachPortCreateRunLoopSource(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(port),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(order),
	)
	// *typing.RefType
	return RunLoopSourceRef(rv)
}

// Returns the canonical name of a specified string encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543585-cfstringgetnameofencoding?language=objc
func StringGetNameOfEncoding(encoding StringEncoding) StringRef {
	rv := C.StringGetNameOfEncoding(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Unregisters an instance of a type with CFPlugIn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493875-cfpluginremoveinstanceforfactory?language=objc
func PlugInRemoveInstanceForFactory(factoryID UUIDRef) {
	C.PlugInRemoveInstanceForFactory(
		// *typing.RefType
		unsafe.Pointer(factoryID),
	)
}

// Returns a code that can be used to identify an object in a hashing structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521137-cfhash?language=objc
func Hash(cf TypeRef) HashCode {
	rv := C.Hash(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(cf),
	)
	// *typing.AliasType
	return HashCode(rv)
}

// Fills a buffer with the file system's native string representation of a given URL's path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541515-cfurlgetfilesystemrepresentation?language=objc
func URLGetFileSystemRepresentation(url URLRef, resolveAgainstBase bool, buffer *uint8, maxBufLen Index) bool {
	rv := C.URLGetFileSystemRepresentation(
		// *typing.RefType
		unsafe.Pointer(url),
		// *typing.PrimitiveType
		C.bool(resolveAgainstBase),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&buffer)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(maxBufLen),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the error associated with a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539682-cfreadstreamcopyerror?language=objc
func ReadStreamCopyError(stream ReadStreamRef) ErrorRef {
	rv := C.ReadStreamCopyError(
		// *typing.RefType
		unsafe.Pointer(stream),
	)
	// *typing.RefType
	return ErrorRef(rv)
}

// Adds a new child to the specified tree as the first in its list of children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401771-cftreeprependchild?language=objc
func TreePrependChild(tree TreeRef, newChild TreeRef) {
	C.TreePrependChild(
		// *typing.RefType
		unsafe.Pointer(tree),
		// *typing.RefType
		unsafe.Pointer(newChild),
	)
}

// Returns an array of CFString objects that represents all known legal ISO currency codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542005-cflocalecopyisocurrencycodes?language=objc
func LocaleCopyISOCurrencyCodes() ArrayRef {
	rv := C.LocaleCopyISOCurrencyCodes()
	// *typing.RefType
	return ArrayRef(rv)
}

// Allocates memory using the specified allocator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521250-cfallocatorallocate?language=objc
func AllocatorAllocate(allocator AllocatorRef, size Index, hint OptionFlags) unsafe.Pointer {
	rv := C.AllocatorAllocate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(size),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(hint),
	)
	// *typing.VoidPointerType
	return unsafe.Pointer(rv)
}

// Returns the number of bytes contained by a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541728-cfdatagetlength?language=objc
func DataGetLength(theData DataRef) Index {
	rv := C.DataGetLength(
		// *typing.RefType
		unsafe.Pointer(theData),
	)
	// *typing.AliasType
	return Index(rv)
}

// Determines whether or not a given key has been imposed on the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515521-cfpreferencesappvalueisforced?language=objc
func PreferencesAppValueIsForced(key StringRef, applicationID StringRef) bool {
	rv := C.PreferencesAppValueIsForced(
		// *typing.RefType
		unsafe.Pointer(key),
		// *typing.RefType
		unsafe.Pointer(applicationID),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns for a CFString object the character encoding that requires the least conversion time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543163-cfstringgetfastestencoding?language=objc
func StringGetFastestEncoding(theString StringRef) StringEncoding {
	rv := C.StringGetFastestEncoding(
		// *typing.RefType
		unsafe.Pointer(theString),
	)
	// *typing.AliasType
	return StringEncoding(rv)
}

// Creates a string from a buffer containing characters in a specified encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543419-cfstringcreatewithbytes?language=objc
func StringCreateWithBytes(alloc AllocatorRef, bytes *uint8, numBytes Index, encoding StringEncoding, isExternalRepresentation bool) StringRef {
	rv := C.StringCreateWithBytes(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&bytes)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(numBytes),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
		// *typing.PrimitiveType
		C.bool(isExternalRepresentation),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Copies the character contents of a string to a local C string buffer after converting the characters to a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542721-cfstringgetcstring?language=objc
func StringGetCString(theString StringRef, buffer string, bufferSize Index, encoding StringEncoding) bool {
	bufferVal := C.CString(buffer)
	defer C.free(unsafe.Pointer(bufferVal))
	rv := C.StringGetCString(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.CStringType
		bufferVal,
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(bufferSize),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Removes the given function from a plug-in’s list of registered factory functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493856-cfpluginunregisterfactory?language=objc
func PlugInUnregisterFactory(factoryUUID UUIDRef) bool {
	rv := C.PlugInUnregisterFactory(
		// *typing.RefType
		unsafe.Pointer(factoryUUID),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Cancels a user notification dialog. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534499-cfusernotificationcancel?language=objc
func UserNotificationCancel(userNotification UserNotificationRef) int32 {
	rv := C.UserNotificationCancel(
		// *typing.RefType
		unsafe.Pointer(userNotification),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Determines whether a given Core Foundation string encoding is available on the current system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542860-cfstringisencodingavailable?language=objc
func StringIsEncodingAvailable(encoding StringEncoding) bool {
	rv := C.StringIsEncodingAvailable(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a Boolean value that indicates whether a given character is a high character in a surrogate pair. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543284-cfstringissurrogatehighcharacter?language=objc
func StringIsSurrogateHighCharacter(character uint16) bool {
	rv := C.StringIsSurrogateHighCharacter(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.UniChar)(character),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a human-presentable failure reason for a given error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1494651-cferrorcopyfailurereason?language=objc
func ErrorCopyFailureReason(err ErrorRef) StringRef {
	rv := C.ErrorCopyFailureReason(
		// *typing.RefType
		unsafe.Pointer(err),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Creates and returns a new immutable dictionary with the key-value pairs of another dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1516796-cfdictionarycreatecopy?language=objc
func DictionaryCreateCopy(allocator AllocatorRef, theDict DictionaryRef) DictionaryRef {
	rv := C.DictionaryCreateCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(theDict),
	)
	// *typing.RefType
	return DictionaryRef(rv)
}

// Returns the scheme portion of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542465-cfurlcopyscheme?language=objc
func URLCopyScheme(anURL URLRef) StringRef {
	rv := C.URLCopyScheme(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the type identifier for the CFNumber opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541730-cfnumbergettypeid?language=objc
func NumberGetTypeID() TypeID {
	rv := C.NumberGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns by reference the start time and duration of a given calendar unit that contains a given absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533501-cfcalendargettimerangeofunit?language=objc
func CalendarGetTimeRangeOfUnit(calendar CalendarRef, unit CalendarUnit, at AbsoluteTime, startp *AbsoluteTime, tip *TimeInterval) bool {
	rv := C.CalendarGetTimeRangeOfUnit(
		// *typing.RefType
		unsafe.Pointer(calendar),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFCalendarUnit)(unit),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(at),
		// *typing.PointerType
		(*C.CFAbsoluteTime)(unsafe.Pointer(&startp)),
		// *typing.PointerType
		(*C.CFTimeInterval)(unsafe.Pointer(&tip)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Adds, modifies, or removes a preference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515528-cfpreferencessetappvalue?language=objc
func PreferencesSetAppValue(key StringRef, value PropertyListRef, applicationID StringRef) {
	C.PreferencesSetAppValue(
		// *typing.RefType
		unsafe.Pointer(key),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFPropertyListRef)(value),
		// *typing.RefType
		unsafe.Pointer(applicationID),
	)
}

// Removes the value at a given index from an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1388765-cfarrayremovevalueatindex?language=objc
func ArrayRemoveValueAtIndex(theArray unsafe.Pointer, idx Index) {
	C.ArrayRemoveValueAtIndex(
		// *typing.RefType
		unsafe.Pointer(theArray),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(idx),
	)
}

// Swaps the bytes of a 64-bit integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425278-cfswapint64?language=objc
func SwapInt64(arg uint64) uint64 {
	rv := C.SwapInt64(
		// *typing.PrimitiveType
		C.uint64_t(arg),
	)
	// *typing.PrimitiveType
	return uint64(rv)
}

// Returns the index of first weekday for a specified calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533471-cfcalendargetfirstweekday?language=objc
func CalendarGetFirstWeekday(calendar CalendarRef) Index {
	rv := C.CalendarGetFirstWeekday(
		// *typing.RefType
		unsafe.Pointer(calendar),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns an array containing all of the bundles currently open in the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537082-cfbundlegetallbundles?language=objc
func BundleGetAllBundles() ArrayRef {
	rv := C.BundleGetAllBundles()
	// *typing.RefType
	return ArrayRef(rv)
}

// Determines if a given URL's path represents a directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542229-cfurlhasdirectorypath?language=objc
func URLHasDirectoryPath(anURL URLRef) bool {
	rv := C.URLHasDirectoryPath(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates an immutable bit vector from a block of memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543718-cfbitvectorcreate?language=objc
func BitVectorCreate(allocator AllocatorRef, bytes *uint8, numBits Index) BitVectorRef {
	rv := C.BitVectorCreate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&bytes)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(numBits),
	)
	// *typing.RefType
	return BitVectorRef(rv)
}

// Returns the application’s distributed notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543492-cfnotificationcentergetdistribut?language=objc
func NotificationCenterGetDistributedCenter() NotificationCenterRef {
	rv := C.NotificationCenterGetDistributedCenter()
	// *typing.RefType
	return NotificationCenterRef(rv)
}

// Appends a buffer of Unicode characters to the character contents of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543667-cfstringappendcharacters?language=objc
func StringAppendCharacters(theString unsafe.Pointer, chars *uint16, numChars Index) {
	C.StringAppendCharacters(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.PointerType
		(*C.UniChar)(unsafe.Pointer(&chars)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(numChars),
	)
}

// Registers a factory function with a CFPlugIn object using the function's name instead of its UUID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493893-cfpluginregisterfactoryfunctionb?language=objc
func PlugInRegisterFactoryFunctionByName(factoryUUID UUIDRef, plugIn unsafe.Pointer, functionName StringRef) bool {
	rv := C.PlugInRegisterFactoryFunctionByName(
		// *typing.RefType
		unsafe.Pointer(factoryUUID),
		// *typing.RefType
		unsafe.Pointer(plugIn),
		// *typing.RefType
		unsafe.Pointer(functionName),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Computes the absolute time when specified components are added to a given absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533487-cfcalendaraddcomponents?language=objc
func CalendarAddComponents(calendar CalendarRef, at *AbsoluteTime, options OptionFlags, componentDesc string) bool {
	componentDescVal := C.CString(componentDesc)
	defer C.free(unsafe.Pointer(componentDescVal))
	rv := C.CalendarAddComponents(
		// *typing.RefType
		unsafe.Pointer(calendar),
		// *typing.PointerType
		(*C.CFAbsoluteTime)(unsafe.Pointer(&at)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(options),
		// *typing.CStringType
		componentDescVal,
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the type used by a CFNumber object to store its value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543388-cfnumbergettype?language=objc
func NumberGetType(number NumberRef) NumberType {
	rv := C.NumberGetType(
		// *typing.RefType
		unsafe.Pointer(number),
	)
	// *typing.AliasType
	return NumberType(rv)
}

// Returns a format string for the given date formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1396261-cfdateformattergetformat?language=objc
func DateFormatterGetFormat(formatter DateFormatterRef) StringRef {
	rv := C.DateFormatterGetFormat(
		// *typing.RefType
		unsafe.Pointer(formatter),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns an array containing a bundle’s localizations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537157-cfbundlecopybundlelocalizations?language=objc
func BundleCopyBundleLocalizations(bundle BundleRef) ArrayRef {
	rv := C.BundleCopyBundleLocalizations(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Returns an application’s main bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537085-cfbundlegetmainbundle?language=objc
func BundleGetMainBundle() BundleRef {
	rv := C.BundleGetMainBundle()
	// *typing.RefType
	return BundleRef(rv)
}

// Returns a Boolean value that indicates whether a CFRunLoopObserver object is valid and able to fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543205-cfrunloopobserverisvalid?language=objc
func RunLoopObserverIsValid(observer RunLoopObserverRef) bool {
	rv := C.RunLoopObserverIsValid(
		// *typing.RefType
		unsafe.Pointer(observer),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Re-enables internal consistency-checking and coalescing for a mutable attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543264-cfattributedstringendediting?language=objc
func AttributedStringEndEditing(aStr unsafe.Pointer) {
	C.AttributedStringEndEditing(
		// *typing.RefType
		unsafe.Pointer(aStr),
	)
}

// Reports whether or not a given Unicode character is in a character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542024-cfcharactersetischaractermember?language=objc
func CharacterSetIsCharacterMember(theSet CharacterSetRef, theChar uint16) bool {
	rv := C.CharacterSetIsCharacterMember(
		// *typing.RefType
		unsafe.Pointer(theSet),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.UniChar)(theChar),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a new CFURL object by resolving the relative portion of a URL against its base. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543590-cfurlcopyabsoluteurl?language=objc
func URLCopyAbsoluteURL(relativeURL URLRef) URLRef {
	rv := C.URLCopyAbsoluteURL(
		// *typing.RefType
		unsafe.Pointer(relativeURL),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns the location of a bundle’s shared support files directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537131-cfbundlecopysharedsupporturl?language=objc
func BundleCopySharedSupportURL(bundle BundleRef) URLRef {
	rv := C.BundleCopySharedSupportURL(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Trims a specified substring from the beginning and end of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542426-cfstringtrim?language=objc
func StringTrim(theString unsafe.Pointer, trimString StringRef) {
	C.StringTrim(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.RefType
		unsafe.Pointer(trimString),
	)
}

// Returns a flag used to set the secure state of a text field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534532-cfusernotificationsecuretextfiel?language=objc
func UserNotificationSecureTextField(i Index) OptionFlags {
	rv := C.UserNotificationSecureTextField(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(i),
	)
	// *typing.AliasType
	return OptionFlags(rv)
}

// Returns flags that control certain behaviors of a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543754-cfsocketgetsocketflags?language=objc
func SocketGetSocketFlags(s SocketRef) OptionFlags {
	rv := C.SocketGetSocketFlags(
		// *typing.RefType
		unsafe.Pointer(s),
	)
	// *typing.AliasType
	return OptionFlags(rv)
}

// Returns the type identifier for the CFString opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542570-cfstringgettypeid?language=objc
func StringGetTypeID() TypeID {
	rv := C.StringGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the Windows codepage identifier that maps most closely to a given Core Foundation encoding constant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543125-cfstringconvertencodingtowindows?language=objc
func StringConvertEncodingToWindowsCodepage(encoding StringEncoding) uint32 {
	rv := C.StringConvertEncodingToWindowsCodepage(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Returns an array of strings that represents ISO currency codes for currencies in common use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543366-cflocalecopycommonisocurrencycod?language=objc
func LocaleCopyCommonISOCurrencyCodes() ArrayRef {
	rv := C.LocaleCopyCommonISOCurrencyCodes()
	// *typing.RefType
	return ArrayRef(rv)
}

// Creates a CFURL object using a given character bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542075-cfurlcreatewithbytes?language=objc
func URLCreateWithBytes(allocator AllocatorRef, URLBytes *uint8, length Index, encoding StringEncoding, baseURL URLRef) URLRef {
	rv := C.URLCreateWithBytes(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&URLBytes)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(length),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
		// *typing.RefType
		unsafe.Pointer(baseURL),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns the opaque type identifier for the CFURLEnumerator opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542687-cfurlenumeratorgettypeid?language=objc
func URLEnumeratorGetTypeID() TypeID {
	rv := C.URLEnumeratorGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Advances the tokenizer to the next token and sets that as the current token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542065-cfstringtokenizeradvancetonextto?language=objc
func StringTokenizerAdvanceToNextToken(tokenizer StringTokenizerRef) StringTokenizerTokenType {
	rv := C.StringTokenizerAdvanceToNextToken(
		// *typing.RefType
		unsafe.Pointer(tokenizer),
	)
	// *typing.AliasType
	return StringTokenizerTokenType(rv)
}

// Creates a mutable copy of a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541480-cfstringcreatemutablecopy?language=objc
func StringCreateMutableCopy(alloc AllocatorRef, maxLength Index, theString StringRef) unsafe.Pointer {
	rv := C.StringCreateMutableCopy(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(maxLength),
		// *typing.RefType
		unsafe.Pointer(theString),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Sets the next firing date for a CFRunLoopTimer object . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542501-cfrunlooptimersetnextfiredate?language=objc
func RunLoopTimerSetNextFireDate(timer RunLoopTimerRef, fireDate AbsoluteTime) {
	C.RunLoopTimerSetNextFireDate(
		// *typing.RefType
		unsafe.Pointer(timer),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(fireDate),
	)
}

// Returns the time zone object identified by a given name or abbreviation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541923-cftimezonecreatewithname?language=objc
func TimeZoneCreateWithName(allocator AllocatorRef, name StringRef, tryAbbrev bool) TimeZoneRef {
	rv := C.TimeZoneCreateWithName(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(name),
		// *typing.PrimitiveType
		C.bool(tryAbbrev),
	)
	// *typing.RefType
	return TimeZoneRef(rv)
}

// Invalidates a CFMessagePort object, stopping it from receiving or sending any more messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542430-cfmessageportinvalidate?language=objc
func MessagePortInvalidate(ms MessagePortRef) {
	C.MessagePortInvalidate(
		// *typing.RefType
		unsafe.Pointer(ms),
	)
}

// Computes the components which are indicated by the componentDesc description string for the given absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533506-cfcalendardecomposeabsolutetime?language=objc
func CalendarDecomposeAbsoluteTime(calendar CalendarRef, at AbsoluteTime, componentDesc string) bool {
	componentDescVal := C.CString(componentDesc)
	defer C.free(unsafe.Pointer(componentDescVal))
	rv := C.CalendarDecomposeAbsoluteTime(
		// *typing.RefType
		unsafe.Pointer(calendar),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(at),
		// *typing.CStringType
		componentDescVal,
	)
	// *typing.PrimitiveType
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521271-cfautorelease?language=objc
func Autorelease(arg TypeRef) TypeRef {
	rv := C.Autorelease(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(arg),
	)
	// *typing.AliasType
	return TypeRef(rv)
}

// Returns an array of CFURL objects describing the locations of all resources in a bundle of the specified type without needing to create a CFBundle object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537140-cfbundlecopyresourceurlsoftypein?language=objc
func BundleCopyResourceURLsOfTypeInDirectory(bundleURL URLRef, resourceType StringRef, subDirName StringRef) ArrayRef {
	rv := C.BundleCopyResourceURLsOfTypeInDirectory(
		// *typing.RefType
		unsafe.Pointer(bundleURL),
		// *typing.RefType
		unsafe.Pointer(resourceType),
		// *typing.RefType
		unsafe.Pointer(subDirName),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542980-cfrunlooptimersettolerance?language=objc
func RunLoopTimerSetTolerance(timer RunLoopTimerRef, tolerance TimeInterval) {
	C.RunLoopTimerSetTolerance(
		// *typing.RefType
		unsafe.Pointer(timer),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFTimeInterval)(tolerance),
	)
}

// Notifies a CFMutableString object that its external backing store of Unicode characters has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542035-cfstringsetexternalcharactersnoc?language=objc
func StringSetExternalCharactersNoCopy(theString unsafe.Pointer, chars *uint16, length Index, capacity Index) {
	C.StringSetExternalCharactersNoCopy(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.PointerType
		(*C.UniChar)(unsafe.Pointer(&chars)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(length),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(capacity),
	)
}

// Locate a bundle given its program-defined identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537139-cfbundlegetbundlewithidentifier?language=objc
func BundleGetBundleWithIdentifier(bundleID StringRef) BundleRef {
	rv := C.BundleGetBundleWithIdentifier(
		// *typing.RefType
		unsafe.Pointer(bundleID),
	)
	// *typing.RefType
	return BundleRef(rv)
}

// Creates an empty CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542987-cfstringcreatemutable?language=objc
func StringCreateMutable(alloc AllocatorRef, maxLength Index) unsafe.Pointer {
	rv := C.StringCreateMutable(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(maxLength),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Creates a CFURL object using a local file system path string relative to a base URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543552-cfurlcreatewithfilesystempathrel?language=objc
func URLCreateWithFileSystemPathRelativeToBase(allocator AllocatorRef, filePath StringRef, pathStyle URLPathStyle, isDirectory bool, baseURL URLRef) URLRef {
	rv := C.URLCreateWithFileSystemPathRelativeToBase(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(filePath),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFURLPathStyle)(pathStyle),
		// *typing.PrimitiveType
		C.bool(isDirectory),
		// *typing.RefType
		unsafe.Pointer(baseURL),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Creates a CFURL object using a given CFString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542055-cfurlcreatewithstring?language=objc
func URLCreateWithString(allocator AllocatorRef, URLString StringRef, baseURL URLRef) URLRef {
	rv := C.URLCreateWithString(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(URLString),
		// *typing.RefType
		unsafe.Pointer(baseURL),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Creates a new mutable bag with the values from another bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1469303-cfbagcreatemutablecopy?language=objc
func BagCreateMutableCopy(allocator AllocatorRef, capacity Index, theBag BagRef) unsafe.Pointer {
	rv := C.BagCreateMutableCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(capacity),
		// *typing.RefType
		unsafe.Pointer(theBag),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns a value (localized if possible) from a bundle’s information dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537102-cfbundlegetvalueforinfodictionar?language=objc
func BundleGetValueForInfoDictionaryKey(bundle BundleRef, key StringRef) TypeRef {
	rv := C.BundleGetValueForInfoDictionaryKey(
		// *typing.RefType
		unsafe.Pointer(bundle),
		// *typing.RefType
		unsafe.Pointer(key),
	)
	// *typing.AliasType
	return TypeRef(rv)
}

// Changes the size of a mutable bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541956-cfbitvectorsetcount?language=objc
func BitVectorSetCount(bv unsafe.Pointer, count Index) {
	C.BitVectorSetCount(
		// *typing.RefType
		unsafe.Pointer(bv),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(count),
	)
}

// Returns the base URL of a given URL if it exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542154-cfurlgetbaseurl?language=objc
func URLGetBaseURL(anURL URLRef) URLRef {
	rv := C.URLGetBaseURL(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns the type identifier for the CFBag opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1469299-cfbaggettypeid?language=objc
func BagGetTypeID() TypeID {
	rv := C.BagGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the value of a CFBoolean object as a standard C type Boolean. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541447-cfbooleangetvalue?language=objc
func BooleanGetValue(boolean BooleanRef) bool {
	rv := C.BooleanGetValue(
		// *typing.RefType
		unsafe.Pointer(boolean),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a human-presentable description for a given error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1494673-cferrorcopydescription?language=objc
func ErrorCopyDescription(err ErrorRef) StringRef {
	rv := C.ErrorCopyDescription(
		// *typing.RefType
		unsafe.Pointer(err),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Not recommended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493848-cfplugininstancegetinstancedata?language=objc
func PlugInInstanceGetInstanceData(instance PlugInInstanceRef) unsafe.Pointer {
	rv := C.PlugInInstanceGetInstanceData(
		// *typing.RefType
		unsafe.Pointer(instance),
	)
	// *typing.VoidPointerType
	return unsafe.Pointer(rv)
}

// Returns the number style used to create the given number formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1390732-cfnumberformattergetstyle?language=objc
func NumberFormatterGetStyle(formatter NumberFormatterRef) NumberFormatterStyle {
	rv := C.NumberFormatterGetStyle(
		// *typing.RefType
		unsafe.Pointer(formatter),
	)
	// *typing.AliasType
	return NumberFormatterStyle(rv)
}

// Changes the first character in each word of a string to uppercase (if it is a lowercase alphabetical character). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542102-cfstringcapitalize?language=objc
func StringCapitalize(theString unsafe.Pointer, locale LocaleRef) {
	C.StringCapitalize(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.RefType
		unsafe.Pointer(locale),
	)
}

// Returns a read-only pointer to the bytes of a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543330-cfdatagetbyteptr?language=objc
func DataGetBytePtr(theData DataRef) *uint8 {
	rv := C.DataGetBytePtr(
		// *typing.RefType
		unsafe.Pointer(theData),
	)
	// *typing.PointerType
	return *(**uint8)(unsafe.Pointer(&rv))
}

// Sets the file mode associated with a CFFileSecurityRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426496-cffilesecuritysetmode?language=objc
func FileSecuritySetMode(fileSec FileSecurityRef, mode int) bool {
	rv := C.FileSecuritySetMode(
		// *typing.RefType
		unsafe.Pointer(fileSec),
		// *typing.AliasType
		// *typing.StructType
		(C.mode_t)(mode),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a flag used to set or test a checkbox’s state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534480-cfusernotificationcheckboxchecke?language=objc
func UserNotificationCheckBoxChecked(i Index) OptionFlags {
	rv := C.UserNotificationCheckBoxChecked(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(i),
	)
	// *typing.AliasType
	return OptionFlags(rv)
}

// Returns the last path component of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542469-cfurlcopylastpathcomponent?language=objc
func URLCopyLastPathComponent(url URLRef) StringRef {
	rv := C.URLCopyLastPathComponent(
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Determines the upper bound on the number of bytes required to hold the file system representation of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543126-cfstringgetmaximumsizeoffilesyst?language=objc
func StringGetMaximumSizeOfFileSystemRepresentation(string_ StringRef) Index {
	rv := C.StringGetMaximumSizeOfFileSystemRepresentation(
		// *typing.RefType
		unsafe.Pointer(string_),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns the location of a bundle’s Resources directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537113-cfbundlecopyresourcesdirectoryur?language=objc
func BundleCopyResourcesDirectoryURL(bundle BundleRef) URLRef {
	rv := C.BundleCopyResourcesDirectoryURL(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns the Core Foundation type identifier for the CFBoolean opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541762-cfbooleangettypeid?language=objc
func BooleanGetTypeID() TypeID {
	rv := C.BooleanGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Gets as a mutable string the string for an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541996-cfattributedstringgetmutablestri?language=objc
func AttributedStringGetMutableString(aStr unsafe.Pointer) unsafe.Pointer {
	rv := C.AttributedStringGetMutableString(
		// *typing.RefType
		unsafe.Pointer(aStr),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Removes the cached resource value identified by a given key from the URL object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542054-cfurlclearresourcepropertycachef?language=objc
func URLClearResourcePropertyCacheForKey(url URLRef, key StringRef) {
	C.URLClearResourcePropertyCacheForKey(
		// *typing.RefType
		unsafe.Pointer(url),
		// *typing.RefType
		unsafe.Pointer(key),
	)
}

// Returns the fragment from a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543642-cfurlcopyfragment?language=objc
func URLCopyFragment(anURL URLRef, charactersToLeaveEscaped StringRef) StringRef {
	rv := C.URLCopyFragment(
		// *typing.RefType
		unsafe.Pointer(anURL),
		// *typing.RefType
		unsafe.Pointer(charactersToLeaveEscaped),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Invalidates a CFSocket object, stopping it from sending or receiving any more messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542010-cfsocketinvalidate?language=objc
func SocketInvalidate(s SocketRef) {
	C.SocketInvalidate(
		// *typing.RefType
		unsafe.Pointer(s),
	)
}

// Creates a CFUUID object from raw UUID bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542951-cfuuidcreatewithbytes?language=objc
func UUIDCreateWithBytes(alloc AllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) UUIDRef {
	rv := C.UUIDCreateWithBytes(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.PrimitiveType
		C.uint8_t(byte0),
		// *typing.PrimitiveType
		C.uint8_t(byte1),
		// *typing.PrimitiveType
		C.uint8_t(byte2),
		// *typing.PrimitiveType
		C.uint8_t(byte3),
		// *typing.PrimitiveType
		C.uint8_t(byte4),
		// *typing.PrimitiveType
		C.uint8_t(byte5),
		// *typing.PrimitiveType
		C.uint8_t(byte6),
		// *typing.PrimitiveType
		C.uint8_t(byte7),
		// *typing.PrimitiveType
		C.uint8_t(byte8),
		// *typing.PrimitiveType
		C.uint8_t(byte9),
		// *typing.PrimitiveType
		C.uint8_t(byte10),
		// *typing.PrimitiveType
		C.uint8_t(byte11),
		// *typing.PrimitiveType
		C.uint8_t(byte12),
		// *typing.PrimitiveType
		C.uint8_t(byte13),
		// *typing.PrimitiveType
		C.uint8_t(byte14),
		// *typing.PrimitiveType
		C.uint8_t(byte15),
	)
	// *typing.RefType
	return UUIDRef(rv)
}

// Creates a CFUserNotification object and displays its notification dialog on screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534528-cfusernotificationcreate?language=objc
func UserNotificationCreate(allocator AllocatorRef, timeout TimeInterval, flags OptionFlags, error *int32, dictionary DictionaryRef) UserNotificationRef {
	rv := C.UserNotificationCreate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFTimeInterval)(timeout),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(flags),
		// *typing.PointerType
		(*C.int32_t)(unsafe.Pointer(&error)),
		// *typing.RefType
		unsafe.Pointer(dictionary),
	)
	// *typing.RefType
	return UserNotificationRef(rv)
}

// Prints the attributes of a string during debugging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542067-cfshowstr?language=objc
func ShowStr(str StringRef) {
	C.ShowStr(
		// *typing.RefType
		unsafe.Pointer(str),
	)
}

// Creates a writable stream for a fixed-size block of memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539650-cfwritestreamcreatewithbuffer?language=objc
func WriteStreamCreateWithBuffer(alloc AllocatorRef, buffer *uint8, bufferCapacity Index) WriteStreamRef {
	rv := C.WriteStreamCreateWithBuffer(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&buffer)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(bufferCapacity),
	)
	// *typing.RefType
	return WriteStreamRef(rv)
}

// Creates a new mutable set with the values from another set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1520429-cfsetcreatemutablecopy?language=objc
func SetCreateMutableCopy(allocator AllocatorRef, capacity Index, theSet SetRef) unsafe.Pointer {
	rv := C.SetCreateMutableCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(capacity),
		// *typing.RefType
		unsafe.Pointer(theSet),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Sets the given allocator as the default for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521134-cfallocatorsetdefault?language=objc
func AllocatorSetDefault(allocator AllocatorRef) {
	C.AllocatorSetDefault(
		// *typing.RefType
		unsafe.Pointer(allocator),
	)
}

// Returns whether a writable stream can accept new data without blocking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539684-cfwritestreamcanacceptbytes?language=objc
func WriteStreamCanAcceptBytes(stream WriteStreamRef) bool {
	rv := C.WriteStreamCanAcceptBytes(
		// *typing.RefType
		unsafe.Pointer(stream),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the type identifier for the CFSet type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1520427-cfsetgettypeid?language=objc
func SetGetTypeID() TypeID {
	rv := C.SetGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the type ID for CFStringTokenizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541794-cfstringtokenizergettypeid?language=objc
func StringTokenizerGetTypeID() TypeID {
	rv := C.StringTokenizerGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the localized name of a given time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542227-cftimezonecopylocalizedname?language=objc
func TimeZoneCopyLocalizedName(tz TimeZoneRef, style TimeZoneNameStyle, locale LocaleRef) StringRef {
	rv := C.TimeZoneCopyLocalizedName(
		// *typing.RefType
		unsafe.Pointer(tz),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFTimeZoneNameStyle)(style),
		// *typing.RefType
		unsafe.Pointer(locale),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Obtains a preference value for the specified key and application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515497-cfpreferencescopyappvalue?language=objc
func PreferencesCopyAppValue(key StringRef, applicationID StringRef) PropertyListRef {
	rv := C.PreferencesCopyAppValue(
		// *typing.RefType
		unsafe.Pointer(key),
		// *typing.RefType
		unsafe.Pointer(applicationID),
	)
	// *typing.AliasType
	return PropertyListRef(rv)
}

// Returns the type identifier for the CFAttributedString opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541715-cfattributedstringgettypeid?language=objc
func AttributedStringGetTypeID() TypeID {
	rv := C.AttributedStringGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the net location portion of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543732-cfurlcopynetlocation?language=objc
func URLCopyNetLocation(anURL URLRef) StringRef {
	rv := C.URLCopyNetLocation(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Adds, modifies, or removes a preference value for the specified domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515525-cfpreferencessetvalue?language=objc
func PreferencesSetValue(key StringRef, value PropertyListRef, applicationID StringRef, userName StringRef, hostName StringRef) {
	C.PreferencesSetValue(
		// *typing.RefType
		unsafe.Pointer(key),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFPropertyListRef)(value),
		// *typing.RefType
		unsafe.Pointer(applicationID),
		// *typing.RefType
		unsafe.Pointer(userName),
		// *typing.RefType
		unsafe.Pointer(hostName),
	)
}

// Disables callbacks for a given CFFileDescriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1477597-cffiledescriptordisablecallbacks?language=objc
func FileDescriptorDisableCallBacks(f FileDescriptorRef, callBackTypes OptionFlags) {
	C.FileDescriptorDisableCallBacks(
		// *typing.RefType
		unsafe.Pointer(f),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(callBackTypes),
	)
}

// Returns a string representation of the given absolute time using the specified date formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1396242-cfdateformattercreatestringwitha?language=objc
func DateFormatterCreateStringWithAbsoluteTime(allocator AllocatorRef, formatter DateFormatterRef, at AbsoluteTime) StringRef {
	rv := C.DateFormatterCreateStringWithAbsoluteTime(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(formatter),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(at),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Determines whether or not a plug-in is loaded on demand. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493872-cfpluginisloadondemand?language=objc
func PlugInIsLoadOnDemand(plugIn unsafe.Pointer) bool {
	rv := C.PlugInIsLoadOnDemand(
		// *typing.RefType
		unsafe.Pointer(plugIn),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the ordering parameter for a CFRunLoopTimer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541633-cfrunlooptimergetorder?language=objc
func RunLoopTimerGetOrder(timer RunLoopTimerRef) Index {
	rv := C.RunLoopTimerGetOrder(
		// *typing.RefType
		unsafe.Pointer(timer),
	)
	// *typing.AliasType
	return Index(rv)
}

// Replaces all characters of a CFMutableString object with other characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543319-cfstringreplaceall?language=objc
func StringReplaceAll(theString unsafe.Pointer, replacement StringRef) {
	C.StringReplaceAll(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.RefType
		unsafe.Pointer(replacement),
	)
}

// Returns the default time zone set for your application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542117-cftimezonecopydefault?language=objc
func TimeZoneCopyDefault() TimeZoneRef {
	rv := C.TimeZoneCopyDefault()
	// *typing.RefType
	return TimeZoneRef(rv)
}

// Converts a 32-bit integer from big-endian format to the host’s native byte order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425266-cfswapint32bigtohost?language=objc
func SwapInt32BigToHost(arg uint32) uint32 {
	rv := C.SwapInt32BigToHost(
		// *typing.PrimitiveType
		C.uint32_t(arg),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Returns the difference in seconds between the receiver and Greenwich Mean Time (GMT) at the specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541927-cftimezonegetsecondsfromgmt?language=objc
func TimeZoneGetSecondsFromGMT(tz TimeZoneRef, at AbsoluteTime) TimeInterval {
	rv := C.TimeZoneGetSecondsFromGMT(
		// *typing.RefType
		unsafe.Pointer(tz),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(at),
	)
	// *typing.AliasType
	return TimeInterval(rv)
}

// Creates a readable stream for a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539727-cfreadstreamcreatewithfile?language=objc
func ReadStreamCreateWithFile(alloc AllocatorRef, fileURL URLRef) ReadStreamRef {
	rv := C.ReadStreamCreateWithFile(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(fileURL),
	)
	// *typing.RefType
	return ReadStreamRef(rv)
}

// Appends the bytes from a byte buffer to the contents of a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541769-cfdataappendbytes?language=objc
func DataAppendBytes(theData unsafe.Pointer, bytes *uint8, length Index) {
	C.DataAppendBytes(
		// *typing.RefType
		unsafe.Pointer(theData),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&bytes)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(length),
	)
}

// Returns the root, canonical locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543257-cflocalegetsystem?language=objc
func LocaleGetSystem() LocaleRef {
	rv := C.LocaleGetSystem()
	// *typing.RefType
	return LocaleRef(rv)
}

// Creates a string from a buffer of Unicode characters that might serve as the backing store for the object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542856-cfstringcreatewithcharactersnoco?language=objc
func StringCreateWithCharactersNoCopy(alloc AllocatorRef, chars *uint16, numChars Index, contentsDeallocator AllocatorRef) StringRef {
	rv := C.StringCreateWithCharactersNoCopy(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.PointerType
		(*C.UniChar)(unsafe.Pointer(&chars)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(numChars),
		// *typing.RefType
		unsafe.Pointer(contentsDeallocator),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Converts a 64-bit integer from little-endian format to the host’s native byte order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425274-cfswapint64littletohost?language=objc
func SwapInt64LittleToHost(arg uint64) uint64 {
	rv := C.SwapInt64LittleToHost(
		// *typing.PrimitiveType
		C.uint64_t(arg),
	)
	// *typing.PrimitiveType
	return uint64(rv)
}

// Returns the number of values currently in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1388772-cfarraygetcount?language=objc
func ArrayGetCount(theArray ArrayRef) Index {
	rv := C.ArrayGetCount(
		// *typing.RefType
		unsafe.Pointer(theArray),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns the parent of a given tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401785-cftreegetparent?language=objc
func TreeGetParent(tree TreeRef) TreeRef {
	rv := C.TreeGetParent(
		// *typing.RefType
		unsafe.Pointer(tree),
	)
	// *typing.RefType
	return TreeRef(rv)
}

// Converts a 64-bit integer from the host’s native byte order to big-endian format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425303-cfswapint64hosttobig?language=objc
func SwapInt64HostToBig(arg uint64) uint64 {
	rv := C.SwapInt64HostToBig(
		// *typing.PrimitiveType
		C.uint64_t(arg),
	)
	// *typing.PrimitiveType
	return uint64(rv)
}

// Runs the current thread’s CFRunLoop object in its default mode indefinitely. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542011-cfrunlooprun?language=objc
func RunLoopRun() {
	C.RunLoopRun()
}

// Returns the main CFRunLoop object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542890-cfrunloopgetmain?language=objc
func RunLoopGetMain() RunLoopRef {
	rv := C.RunLoopGetMain()
	// *typing.RefType
	return RunLoopRef(rv)
}

// Returns a string representation of the given number using the specified number formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1390814-cfnumberformattercreatestringwit?language=objc
func NumberFormatterCreateStringWithNumber(allocator AllocatorRef, formatter NumberFormatterRef, number NumberRef) StringRef {
	rv := C.NumberFormatterCreateStringWithNumber(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(formatter),
		// *typing.RefType
		unsafe.Pointer(number),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the number of levels a recursive directory enumerator has descended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542086-cfurlenumeratorgetdescendentleve?language=objc
func URLEnumeratorGetDescendentLevel(enumerator URLEnumeratorRef) Index {
	rv := C.URLEnumeratorGetDescendentLevel(
		// *typing.RefType
		unsafe.Pointer(enumerator),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns the type identifier for the CFSocket opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542549-cfsocketgettypeid?language=objc
func SocketGetTypeID() TypeID {
	rv := C.SocketGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns a Boolean value that indicates whether a CFMessagePort object represents a remote port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543277-cfmessageportisremote?language=objc
func MessagePortIsRemote(ms MessagePortRef) bool {
	rv := C.MessagePortIsRemote(
		// *typing.RefType
		unsafe.Pointer(ms),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Invalidates a CFMachPort object, stopping it from receiving any more messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1400922-cfmachportinvalidate?language=objc
func MachPortInvalidate(port MachPortRef) {
	C.MachPortInvalidate(
		// *typing.RefType
		unsafe.Pointer(port),
	)
}

// Adds the characters in a given string to a character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543710-cfcharactersetaddcharactersinstr?language=objc
func CharacterSetAddCharactersInString(theSet unsafe.Pointer, theString StringRef) {
	C.CharacterSetAddCharactersInString(
		// *typing.RefType
		unsafe.Pointer(theSet),
		// *typing.RefType
		unsafe.Pointer(theString),
	)
}

// Returns an array of CFString objects that represents all known legal ISO language codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542179-cflocalecopyisolanguagecodes?language=objc
func LocaleCopyISOLanguageCodes() ArrayRef {
	rv := C.LocaleCopyISOLanguageCodes()
	// *typing.RefType
	return ArrayRef(rv)
}

// Sends data over a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543414-cfsocketsenddata?language=objc
func SocketSendData(s SocketRef, address DataRef, data DataRef, timeout TimeInterval) SocketError {
	rv := C.SocketSendData(
		// *typing.RefType
		unsafe.Pointer(s),
		// *typing.RefType
		unsafe.Pointer(address),
		// *typing.RefType
		unsafe.Pointer(data),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFTimeInterval)(timeout),
	)
	// *typing.AliasType
	return SocketError(rv)
}

// Returns the native socket associated with a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543617-cfsocketgetnative?language=objc
func SocketGetNative(s SocketRef) SocketNativeHandle {
	rv := C.SocketGetNative(
		// *typing.RefType
		unsafe.Pointer(s),
	)
	// *typing.AliasType
	return SocketNativeHandle(rv)
}

// Sets the default port number with which to connect to a CFSocket name server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541710-cfsocketsetdefaultnameregistrypo?language=objc
func SocketSetDefaultNameRegistryPortNumber(port uint16) {
	C.SocketSetDefaultNameRegistryPortNumber(
		// *typing.PrimitiveType
		C.uint16_t(port),
	)
}

// Returns the Core Foundation encoding constant that is the closest mapping to a given IANA registry “charset” name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542975-cfstringconvertianacharsetnameto?language=objc
func StringConvertIANACharSetNameToEncoding(theString StringRef) StringEncoding {
	rv := C.StringConvertIANACharSetNameToEncoding(
		// *typing.RefType
		unsafe.Pointer(theString),
	)
	// *typing.AliasType
	return StringEncoding(rv)
}

// Creates a string from its “external representation.” [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542194-cfstringcreatefromexternalrepres?language=objc
func StringCreateFromExternalRepresentation(alloc AllocatorRef, data DataRef, encoding StringEncoding) StringRef {
	rv := C.StringCreateFromExternalRepresentation(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(data),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Determines if a string ends with a specified sequence of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542768-cfstringhassuffix?language=objc
func StringHasSuffix(theString StringRef, suffix StringRef) bool {
	rv := C.StringHasSuffix(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.RefType
		unsafe.Pointer(suffix),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Gets the owner ID associated with a CFFileSecurityRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426516-cffilesecuritygetowner?language=objc
func FileSecurityGetOwner(fileSec FileSecurityRef, owner *int) bool {
	rv := C.FileSecurityGetOwner(
		// *typing.RefType
		unsafe.Pointer(fileSec),
		// *typing.PointerType
		(*C.uid_t)(unsafe.Pointer(&owner)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Not recommended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493887-cfplugininstancegetfactoryname?language=objc
func PlugInInstanceGetFactoryName(instance PlugInInstanceRef) StringRef {
	rv := C.PlugInInstanceGetFactoryName(
		// *typing.RefType
		unsafe.Pointer(instance),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the bundle’s development region from the bundle’s information property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537145-cfbundlegetdevelopmentregion?language=objc
func BundleGetDevelopmentRegion(bundle BundleRef) StringRef {
	rv := C.BundleGetDevelopmentRegion(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Resets the length of a CFMutableData object's internal byte buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542375-cfdatasetlength?language=objc
func DataSetLength(theData unsafe.Pointer, length Index) {
	C.DataSetLength(
		// *typing.RefType
		unsafe.Pointer(theData),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(length),
	)
}

// Given an array of possible localizations and preferred locations, returns the one or more of them that CFBundle would use, without reference to the current application context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537153-cfbundlecopylocalizationsforpref?language=objc
func BundleCopyLocalizationsForPreferences(locArray ArrayRef, prefArray ArrayRef) ArrayRef {
	rv := C.BundleCopyLocalizationsForPreferences(
		// *typing.RefType
		unsafe.Pointer(locArray),
		// *typing.RefType
		unsafe.Pointer(prefArray),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Creates a CFData object containing the content of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543276-cfurlcreatedata?language=objc
func URLCreateData(allocator AllocatorRef, url URLRef, encoding StringEncoding, escapeWhitespace bool) DataRef {
	rv := C.URLCreateData(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(url),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
		// *typing.PrimitiveType
		C.bool(escapeWhitespace),
	)
	// *typing.RefType
	return DataRef(rv)
}

// Returns a copy of a locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542267-cflocalecreatecopy?language=objc
func LocaleCreateCopy(allocator AllocatorRef, locale LocaleRef) LocaleRef {
	rv := C.LocaleCreateCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(locale),
	)
	// *typing.RefType
	return LocaleRef(rv)
}

// Returns the line direction for the specified ISO language code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542965-cflocalegetlanguagelinedirection?language=objc
func LocaleGetLanguageLineDirection(isoLangCode StringRef) LocaleLanguageDirection {
	rv := C.LocaleGetLanguageLineDirection(
		// *typing.RefType
		unsafe.Pointer(isoLangCode),
	)
	// *typing.AliasType
	return LocaleLanguageDirection(rv)
}

// Creates a new CFDateFormatter object, localized to the given locale, which will format dates to the given date and time styles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1396269-cfdateformattercreate?language=objc
func DateFormatterCreate(allocator AllocatorRef, locale LocaleRef, dateStyle DateFormatterStyle, timeStyle DateFormatterStyle) DateFormatterRef {
	rv := C.DateFormatterCreate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(locale),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFDateFormatterStyle)(dateStyle),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFDateFormatterStyle)(timeStyle),
	)
	// *typing.RefType
	return DateFormatterRef(rv)
}

// Searches all registered plug-ins for factory functions capable of creating an instance of the given type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493884-cfpluginfindfactoriesforpluginty?language=objc
func PlugInFindFactoriesForPlugInType(typeUUID UUIDRef) ArrayRef {
	rv := C.PlugInFindFactoriesForPlugInType(
		// *typing.RefType
		unsafe.Pointer(typeUUID),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Creates a CFPlugIn given its URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493873-cfplugincreate?language=objc
func PlugInCreate(allocator AllocatorRef, plugInURL URLRef) unsafe.Pointer {
	rv := C.PlugInCreate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(plugInURL),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns the next firing time for a CFRunLoopTimer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543643-cfrunlooptimergetnextfiredate?language=objc
func RunLoopTimerGetNextFireDate(timer RunLoopTimerRef) AbsoluteTime {
	rv := C.RunLoopTimerGetNextFireDate(
		// *typing.RefType
		unsafe.Pointer(timer),
	)
	// *typing.AliasType
	return AbsoluteTime(rv)
}

// Returns the number of fraction digits that should be displayed, and the rounding increment, for a given currency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1390757-cfnumberformattergetdecimalinfof?language=objc
func NumberFormatterGetDecimalInfoForCurrencyCode(currencyCode StringRef, defaultFractionDigits *int32, roundingIncrement *float64) bool {
	rv := C.NumberFormatterGetDecimalInfoForCurrencyCode(
		// *typing.RefType
		unsafe.Pointer(currencyCode),
		// *typing.PointerType
		(*C.int32_t)(unsafe.Pointer(&defaultFractionDigits)),
		// *typing.PointerType
		(*C.double)(unsafe.Pointer(&roundingIncrement)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a locale object for a specified calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533515-cfcalendarcopylocale?language=objc
func CalendarCopyLocale(calendar CalendarRef) LocaleRef {
	rv := C.CalendarCopyLocale(
		// *typing.RefType
		unsafe.Pointer(calendar),
	)
	// *typing.RefType
	return LocaleRef(rv)
}

// Writes data to a writable stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539680-cfwritestreamwrite?language=objc
func WriteStreamWrite(stream WriteStreamRef, buffer *uint8, bufferLength Index) Index {
	rv := C.WriteStreamWrite(
		// *typing.RefType
		unsafe.Pointer(stream),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&buffer)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(bufferLength),
	)
	// *typing.AliasType
	return Index(rv)
}

// Prints a description of a Core Foundation object to stderr. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541433-cfshow?language=objc
func Show(obj TypeRef) {
	C.Show(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(obj),
	)
}

// Returns a Boolean value that indicates whether hyphenation data is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543237-cfstringishyphenationavailablefo?language=objc
func StringIsHyphenationAvailableForLocale(locale LocaleRef) bool {
	rv := C.StringIsHyphenationAvailableForLocale(
		// *typing.RefType
		unsafe.Pointer(locale),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a given attribute of the current token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542839-cfstringtokenizercopycurrenttoke?language=objc
func StringTokenizerCopyCurrentTokenAttribute(tokenizer StringTokenizerRef, attribute OptionFlags) TypeRef {
	rv := C.StringTokenizerCopyCurrentTokenAttribute(
		// *typing.RefType
		unsafe.Pointer(tokenizer),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(attribute),
	)
	// *typing.AliasType
	return TypeRef(rv)
}

// Creates a CFUUID object for a specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543502-cfuuidcreatefromstring?language=objc
func UUIDCreateFromString(alloc AllocatorRef, uuidStr StringRef) UUIDRef {
	rv := C.UUIDCreateFromString(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(uuidStr),
	)
	// *typing.RefType
	return UUIDRef(rv)
}

// Returns the type identifier for the CFArray opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1388776-cfarraygettypeid?language=objc
func ArrayGetTypeID() TypeID {
	rv := C.ArrayGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Signals a CFRunLoopSource object, marking it as ready to fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543700-cfrunloopsourcesignal?language=objc
func RunLoopSourceSignal(source RunLoopSourceRef) {
	C.RunLoopSourceSignal(
		// *typing.RefType
		unsafe.Pointer(source),
	)
}

// Binds a local address to a CFSocket object and configures it for listening. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542729-cfsocketsetaddress?language=objc
func SocketSetAddress(s SocketRef, address DataRef) SocketError {
	rv := C.SocketSetAddress(
		// *typing.RefType
		unsafe.Pointer(s),
		// *typing.RefType
		unsafe.Pointer(address),
	)
	// *typing.AliasType
	return SocketError(rv)
}

// Creates an immutable string from a C string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542942-cfstringcreatewithcstring?language=objc
func StringCreateWithCString(alloc AllocatorRef, cStr string, encoding StringEncoding) StringRef {
	cStrVal := C.CString(cStr)
	defer C.free(unsafe.Pointer(cStrVal))
	rv := C.StringCreateWithCString(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.CStringType
		cStrVal,
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Creates a CFMutableData object by copying another CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542459-cfdatacreatemutablecopy?language=objc
func DataCreateMutableCopy(allocator AllocatorRef, capacity Index, theData DataRef) unsafe.Pointer {
	rv := C.DataCreateMutableCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(capacity),
		// *typing.RefType
		unsafe.Pointer(theData),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Creates a mutable attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543247-cfattributedstringcreatemutable?language=objc
func AttributedStringCreateMutable(alloc AllocatorRef, maxLength Index) unsafe.Pointer {
	rv := C.AttributedStringCreateMutable(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(maxLength),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns the name with which a CFMessagePort object is registered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542032-cfmessageportgetname?language=objc
func MessagePortGetName(ms MessagePortRef) StringRef {
	rv := C.MessagePortGetName(
		// *typing.RefType
		unsafe.Pointer(ms),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Sets the owner UUID associated with a CFFileSecurityRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426494-cffilesecuritysetowneruuid?language=objc
func FileSecuritySetOwnerUUID(fileSec FileSecurityRef, ownerUUID UUIDRef) bool {
	rv := C.FileSecuritySetOwnerUUID(
		// *typing.RefType
		unsafe.Pointer(fileSec),
		// *typing.RefType
		unsafe.Pointer(ownerUUID),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a CFDate object’s absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542812-cfdategetabsolutetime?language=objc
func DateGetAbsoluteTime(theDate DateRef) AbsoluteTime {
	rv := C.DateGetAbsoluteTime(
		// *typing.RefType
		unsafe.Pointer(theDate),
	)
	// *typing.AliasType
	return AbsoluteTime(rv)
}

// Sets the locale for a calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533472-cfcalendarsetlocale?language=objc
func CalendarSetLocale(calendar CalendarRef, locale LocaleRef) {
	C.CalendarSetLocale(
		// *typing.RefType
		unsafe.Pointer(calendar),
		// *typing.RefType
		unsafe.Pointer(locale),
	)
}

// Returns an array of CFNumbers representing the architectures a given bundle provides. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537144-cfbundlecopyexecutablearchitectu?language=objc
func BundleCopyExecutableArchitectures(bundle BundleRef) ArrayRef {
	rv := C.BundleCopyExecutableArchitectures(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Invalidates a CFRunLoopObserver object, stopping it from ever firing again. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543197-cfrunloopobserverinvalidate?language=objc
func RunLoopObserverInvalidate(observer RunLoopObserverRef) {
	C.RunLoopObserverInvalidate(
		// *typing.RefType
		unsafe.Pointer(observer),
	)
}

// Convenience function that allows you to set and remove multiple preference values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515513-cfpreferencessetmultiple?language=objc
func PreferencesSetMultiple(keysToSet DictionaryRef, keysToRemove ArrayRef, applicationID StringRef, userName StringRef, hostName StringRef) {
	C.PreferencesSetMultiple(
		// *typing.RefType
		unsafe.Pointer(keysToSet),
		// *typing.RefType
		unsafe.Pointer(keysToRemove),
		// *typing.RefType
		unsafe.Pointer(applicationID),
		// *typing.RefType
		unsafe.Pointer(userName),
		// *typing.RefType
		unsafe.Pointer(hostName),
	)
}

// Returns the maximum number of bytes a string of a specified length (in Unicode characters) will take up if encoded in a specified encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542143-cfstringgetmaximumsizeforencodin?language=objc
func StringGetMaximumSizeForEncoding(length Index, encoding StringEncoding) Index {
	rv := C.StringGetMaximumSizeForEncoding(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(length),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns a Boolean value that indicates whether a CFRunLoopTimer object repeats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543545-cfrunlooptimerdoesrepeat?language=objc
func RunLoopTimerDoesRepeat(timer RunLoopTimerRef) bool {
	rv := C.RunLoopTimerDoesRepeat(
		// *typing.RefType
		unsafe.Pointer(timer),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns an array of CFString objects that represents all locales for which locale data is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541997-cflocalecopyavailablelocaleident?language=objc
func LocaleCopyAvailableLocaleIdentifiers() ArrayRef {
	rv := C.LocaleCopyAvailableLocaleIdentifiers()
	// *typing.RefType
	return ArrayRef(rv)
}

// Returns the location of a bundle’s shared frameworks directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537159-cfbundlecopysharedframeworksurl?language=objc
func BundleCopySharedFrameworksURL(bundle BundleRef) URLRef {
	rv := C.BundleCopySharedFrameworksURL(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Swaps the bytes of a 16-bit integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425304-cfswapint16?language=objc
func SwapInt16(arg uint16) uint16 {
	rv := C.SwapInt16(
		// *typing.PrimitiveType
		C.uint16_t(arg),
	)
	// *typing.PrimitiveType
	return uint16(rv)
}

// Creates a CFBundle object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537154-cfbundlecreate?language=objc
func BundleCreate(allocator AllocatorRef, bundleURL URLRef) BundleRef {
	rv := C.BundleCreate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(bundleURL),
	)
	// *typing.RefType
	return BundleRef(rv)
}

// Creates a new string by replacing any percent escape sequences with their character equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542938-cfurlcreatestringbyreplacingperc?language=objc
func URLCreateStringByReplacingPercentEscapes(allocator AllocatorRef, originalString StringRef, charactersToLeaveEscaped StringRef) StringRef {
	rv := C.URLCreateStringByReplacingPercentEscapes(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(originalString),
		// *typing.RefType
		unsafe.Pointer(charactersToLeaveEscaped),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the daylight saving time offset for a time zone at a given time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543630-cftimezonegetdaylightsavingtimeo?language=objc
func TimeZoneGetDaylightSavingTimeOffset(tz TimeZoneRef, at AbsoluteTime) TimeInterval {
	rv := C.TimeZoneGetDaylightSavingTimeOffset(
		// *typing.RefType
		unsafe.Pointer(tz),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(at),
	)
	// *typing.AliasType
	return TimeInterval(rv)
}

// Sets all bits in a bit vector to a particular value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543127-cfbitvectorsetallbits?language=objc
func BitVectorSetAllBits(bv unsafe.Pointer, value Bit) {
	C.BitVectorSetAllBits(
		// *typing.RefType
		unsafe.Pointer(bv),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFBit)(value),
	)
}

// Given a CFString object containing XML source with unescaped entities, returns a string with specified XML entities escaped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542736-cfxmlcreatestringbyescapingentit?language=objc
func XMLCreateStringByEscapingEntities(allocator AllocatorRef, string_ StringRef, entitiesDictionary DictionaryRef) StringRef {
	rv := C.XMLCreateStringByEscapingEntities(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(string_),
		// *typing.RefType
		unsafe.Pointer(entitiesDictionary),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the unique identifier of an opaque type to which a Core Foundation object belongs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521218-cfgettypeid?language=objc
func GetTypeID(cf TypeRef) TypeID {
	rv := C.GetTypeID(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(cf),
	)
	// *typing.AliasType
	return TypeID(rv)
}

// Convenience function that directly obtains an integer preference value for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515526-cfpreferencesgetappintegervalue?language=objc
func PreferencesGetAppIntegerValue(key StringRef, applicationID StringRef, keyExistsAndHasValidFormat *bool) Index {
	rv := C.PreferencesGetAppIntegerValue(
		// *typing.RefType
		unsafe.Pointer(key),
		// *typing.RefType
		unsafe.Pointer(applicationID),
		// *typing.PointerType
		(*C.bool)(unsafe.Pointer(&keyExistsAndHasValidFormat)),
	)
	// *typing.AliasType
	return Index(rv)
}

// Creates a CFString object from an external C string buffer that might serve as the backing store for the object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542134-cfstringcreatewithcstringnocopy?language=objc
func StringCreateWithCStringNoCopy(alloc AllocatorRef, cStr string, encoding StringEncoding, contentsDeallocator AllocatorRef) StringRef {
	cStrVal := C.CString(cStr)
	defer C.free(unsafe.Pointer(cStrVal))
	rv := C.StringCreateWithCStringNoCopy(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.CStringType
		cStrVal,
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
		// *typing.RefType
		unsafe.Pointer(contentsDeallocator),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the type identifier the CFReadStream opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539751-cfreadstreamgettypeid?language=objc
func ReadStreamGetTypeID() TypeID {
	rv := C.ReadStreamGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Creates a CFDate object given an absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541708-cfdatecreate?language=objc
func DateCreate(allocator AllocatorRef, at AbsoluteTime) DateRef {
	rv := C.DateCreate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(at),
	)
	// *typing.RefType
	return DateRef(rv)
}

// Changes all uppercase alphabetical characters in a CFMutableString to lowercase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543576-cfstringlowercase?language=objc
func StringLowercase(theString unsafe.Pointer, locale LocaleRef) {
	C.StringLowercase(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.RefType
		unsafe.Pointer(locale),
	)
}

// Returns the Core Foundation encoding constant that is the closest mapping to a given Cocoa encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542787-cfstringconvertnsstringencodingt?language=objc
func StringConvertNSStringEncodingToEncoding(encoding int32) StringEncoding {
	rv := C.StringConvertNSStringEncodingToEncoding(
		// *typing.PrimitiveType
		C.int32_t(encoding),
	)
	// *typing.AliasType
	return StringEncoding(rv)
}

// Returns the smallest encoding on the current system for the character contents of a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541970-cfstringgetsmallestencoding?language=objc
func StringGetSmallestEncoding(theString StringRef) StringEncoding {
	rv := C.StringGetSmallestEncoding(
		// *typing.RefType
		unsafe.Pointer(theString),
	)
	// *typing.AliasType
	return StringEncoding(rv)
}

// Determines if the character data of a string begin with a specified sequence of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542014-cfstringhasprefix?language=objc
func StringHasPrefix(theString StringRef, prefix StringRef) bool {
	rv := C.StringHasPrefix(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.RefType
		unsafe.Pointer(prefix),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a localized date format string representing the given date format components arranged appropriately for the specified locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1396320-cfdateformattercreatedateformatf?language=objc
func DateFormatterCreateDateFormatFromTemplate(allocator AllocatorRef, tmplate StringRef, options OptionFlags, locale LocaleRef) StringRef {
	rv := C.DateFormatterCreateDateFormatFromTemplate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(tmplate),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(options),
		// *typing.RefType
		unsafe.Pointer(locale),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Clears properties from a CFFileSecurityRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426500-cffilesecurityclearproperties?language=objc
func FileSecurityClearProperties(fileSec FileSecurityRef, clearPropertyMask FileSecurityClearOptions) bool {
	rv := C.FileSecurityClearProperties(
		// *typing.RefType
		unsafe.Pointer(fileSec),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFFileSecurityClearOptions)(clearPropertyMask),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a pointer to a mutable byte buffer of a CFMutableData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542870-cfdatagetmutablebyteptr?language=objc
func DataGetMutableBytePtr(theData unsafe.Pointer) *uint8 {
	rv := C.DataGetMutableBytePtr(
		// *typing.RefType
		unsafe.Pointer(theData),
	)
	// *typing.PointerType
	return *(**uint8)(unsafe.Pointer(&rv))
}

// Returns the number of bit values in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543392-cfbitvectorgetcount?language=objc
func BitVectorGetCount(bv BitVectorRef) Index {
	rv := C.BitVectorGetCount(
		// *typing.RefType
		unsafe.Pointer(bv),
	)
	// *typing.AliasType
	return Index(rv)
}

// Forces a CFRunLoop object to stop running. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541796-cfrunloopstop?language=objc
func RunLoopStop(rl RunLoopRef) {
	C.RunLoopStop(
		// *typing.RefType
		unsafe.Pointer(rl),
	)
}

// Reports whether or not a character set contains at least one member character in the specified plane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542152-cfcharactersethasmemberinplane?language=objc
func CharacterSetHasMemberInPlane(theSet CharacterSetRef, thePlane Index) bool {
	rv := C.CharacterSetHasMemberInPlane(
		// *typing.RefType
		unsafe.Pointer(theSet),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(thePlane),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Wakes a waiting CFRunLoop object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541622-cfrunloopwakeup?language=objc
func RunLoopWakeUp(rl RunLoopRef) {
	C.RunLoopWakeUp(
		// *typing.RefType
		unsafe.Pointer(rl),
	)
}

// Extracts the contents of a string as a NULL-terminated 8-bit string appropriate for passing to POSIX APIs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542019-cfstringgetfilesystemrepresentat?language=objc
func StringGetFileSystemRepresentation(string_ StringRef, buffer string, maxBufLen Index) bool {
	bufferVal := C.CString(buffer)
	defer C.free(unsafe.Pointer(bufferVal))
	rv := C.StringGetFileSystemRepresentation(
		// *typing.RefType
		unsafe.Pointer(string_),
		// *typing.CStringType
		bufferVal,
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(maxBufLen),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Determines if the given URL conforms to RFC 1808 and therefore can be decomposed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542131-cfurlcanbedecomposed?language=objc
func URLCanBeDecomposed(anURL URLRef) bool {
	rv := C.URLCanBeDecomposed(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Extracts the values of the text fields from a dismissed notification dialog. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534517-cfusernotificationgetresponseval?language=objc
func UserNotificationGetResponseValue(userNotification UserNotificationRef, key StringRef, idx Index) StringRef {
	rv := C.UserNotificationGetResponseValue(
		// *typing.RefType
		unsafe.Pointer(userNotification),
		// *typing.RefType
		unsafe.Pointer(key),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(idx),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Creates a new mutable array with the values from another array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1388800-cfarraycreatemutablecopy?language=objc
func ArrayCreateMutableCopy(allocator AllocatorRef, capacity Index, theArray ArrayRef) unsafe.Pointer {
	rv := C.ArrayCreateMutableCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(capacity),
		// *typing.RefType
		unsafe.Pointer(theArray),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Creates a readable stream for a block of memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539646-cfreadstreamcreatewithbytesnocop?language=objc
func ReadStreamCreateWithBytesNoCopy(alloc AllocatorRef, bytes *uint8, length Index, bytesDeallocator AllocatorRef) ReadStreamRef {
	rv := C.ReadStreamCreateWithBytesNoCopy(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&bytes)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(length),
		// *typing.RefType
		unsafe.Pointer(bytesDeallocator),
	)
	// *typing.RefType
	return ReadStreamRef(rv)
}

// Displays a user notification dialog and waits for a user response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534496-cfusernotificationdisplayalert?language=objc
func UserNotificationDisplayAlert(timeout TimeInterval, flags OptionFlags, iconURL URLRef, soundURL URLRef, localizationURL URLRef, alertHeader StringRef, alertMessage StringRef, defaultButtonTitle StringRef, alternateButtonTitle StringRef, otherButtonTitle StringRef, responseFlags *OptionFlags) int32 {
	rv := C.UserNotificationDisplayAlert(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFTimeInterval)(timeout),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(flags),
		// *typing.RefType
		unsafe.Pointer(iconURL),
		// *typing.RefType
		unsafe.Pointer(soundURL),
		// *typing.RefType
		unsafe.Pointer(localizationURL),
		// *typing.RefType
		unsafe.Pointer(alertHeader),
		// *typing.RefType
		unsafe.Pointer(alertMessage),
		// *typing.RefType
		unsafe.Pointer(defaultButtonTitle),
		// *typing.RefType
		unsafe.Pointer(alternateButtonTitle),
		// *typing.RefType
		unsafe.Pointer(otherButtonTitle),
		// *typing.PointerType
		(*C.CFOptionFlags)(unsafe.Pointer(&responseFlags)),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Returns the host name of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543550-cfurlcopyhostname?language=objc
func URLCopyHostName(anURL URLRef) StringRef {
	rv := C.URLCopyHostName(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Retrieves a value at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1388767-cfarraygetvalueatindex?language=objc
func ArrayGetValueAtIndex(theArray ArrayRef, idx Index) unsafe.Pointer {
	rv := C.ArrayGetValueAtIndex(
		// *typing.RefType
		unsafe.Pointer(theArray),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(idx),
	)
	// *typing.VoidPointerType
	return unsafe.Pointer(rv)
}

// Enables or disables load on demand for plug-ins that do dynamic registration (only when a client requests an instance of a supported type). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493898-cfpluginsetloadondemand?language=objc
func PlugInSetLoadOnDemand(plugIn unsafe.Pointer, flag bool) {
	C.PlugInSetLoadOnDemand(
		// *typing.RefType
		unsafe.Pointer(plugIn),
		// *typing.PrimitiveType
		C.bool(flag),
	)
}

// Returns a time zone object for the specified time interval offset from Greenwich Mean Time (GMT). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542743-cftimezonecreatewithtimeinterval?language=objc
func TimeZoneCreateWithTimeIntervalFromGMT(allocator AllocatorRef, ti TimeInterval) TimeZoneRef {
	rv := C.TimeZoneCreateWithTimeIntervalFromGMT(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFTimeInterval)(ti),
	)
	// *typing.RefType
	return TimeZoneRef(rv)
}

// Returns a bundle’s version number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537160-cfbundlegetversionnumber?language=objc
func BundleGetVersionNumber(bundle BundleRef) uint32 {
	rv := C.BundleGetVersionNumber(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Sets a date formatter property using a key-value pair. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1396265-cfdateformattersetproperty?language=objc
func DateFormatterSetProperty(formatter DateFormatterRef, key StringRef, value TypeRef) {
	C.DateFormatterSetProperty(
		// *typing.RefType
		unsafe.Pointer(formatter),
		// *typing.RefType
		unsafe.Pointer(key),
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(value),
	)
}

// Returns the type identifier for the CFMachPort opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1400918-cfmachportgettypeid?language=objc
func MachPortGetTypeID() TypeID {
	rv := C.MachPortGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Creates a new CFURL object for a file system entity using the native representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543314-cfurlcreatefromfilesystemreprese?language=objc
func URLCreateFromFileSystemRepresentation(allocator AllocatorRef, buffer *uint8, bufLen Index, isDirectory bool) URLRef {
	rv := C.URLCreateFromFileSystemRepresentation(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&buffer)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(bufLen),
		// *typing.PrimitiveType
		C.bool(isDirectory),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns the string representation of a specified CFUUID object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543625-cfuuidcreatestring?language=objc
func UUIDCreateString(alloc AllocatorRef, uuid UUIDRef) StringRef {
	rv := C.UUIDCreateString(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(uuid),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the Core Foundation encoding constant that is the closest mapping to a given Windows codepage identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542113-cfstringconvertwindowscodepageto?language=objc
func StringConvertWindowsCodepageToEncoding(codepage uint32) StringEncoding {
	rv := C.StringConvertWindowsCodepageToEncoding(
		// *typing.PrimitiveType
		C.uint32_t(codepage),
	)
	// *typing.AliasType
	return StringEncoding(rv)
}

// Trims whitespace from the beginning and end of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542093-cfstringtrimwhitespace?language=objc
func StringTrimWhitespace(theString unsafe.Pointer) {
	C.StringTrimWhitespace(
		// *typing.RefType
		unsafe.Pointer(theString),
	)
}

// Returns a Boolean value that indicates whether a CFMachPort object is valid and able to receive messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1400936-cfmachportisvalid?language=objc
func MachPortIsValid(port MachPortRef) bool {
	rv := C.MachPortIsValid(
		// *typing.RefType
		unsafe.Pointer(port),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the type identifier of the CFBinaryHeap opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1509309-cfbinaryheapgettypeid?language=objc
func BinaryHeapGetTypeID() TypeID {
	rv := C.BinaryHeapGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Defers internal consistency-checking and coalescing for a mutable attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542876-cfattributedstringbeginediting?language=objc
func AttributedStringBeginEditing(aStr unsafe.Pointer) {
	C.AttributedStringBeginEditing(
		// *typing.RefType
		unsafe.Pointer(aStr),
	)
}

// Returns a Boolean value that indicates whether a CFRunLoopSource object is valid and able to fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541509-cfrunloopsourceisvalid?language=objc
func RunLoopSourceIsValid(source RunLoopSourceRef) bool {
	rv := C.RunLoopSourceIsValid(
		// *typing.RefType
		unsafe.Pointer(source),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a copy of a given URL with its last path extension removed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542836-cfurlcreatecopydeletingpathexten?language=objc
func URLCreateCopyDeletingPathExtension(allocator AllocatorRef, url URLRef) URLRef {
	rv := C.URLCreateCopyDeletingPathExtension(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns the reference count of a Core Foundation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521288-cfgetretaincount?language=objc
func GetRetainCount(cf TypeRef) Index {
	rv := C.GetRetainCount(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(cf),
	)
	// *typing.AliasType
	return Index(rv)
}

// Creates a time zone with a given name and data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542130-cftimezonecreate?language=objc
func TimeZoneCreate(allocator AllocatorRef, name StringRef, data DataRef) TimeZoneRef {
	rv := C.TimeZoneCreate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(name),
		// *typing.RefType
		unsafe.Pointer(data),
	)
	// *typing.RefType
	return TimeZoneRef(rv)
}

// Creates a CFString from a zero-terminated POSIX file system representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543636-cfstringcreatewithfilesystemrepr?language=objc
func StringCreateWithFileSystemRepresentation(alloc AllocatorRef, buffer string) StringRef {
	bufferVal := C.CString(buffer)
	defer C.free(unsafe.Pointer(bufferVal))
	rv := C.StringCreateWithFileSystemRepresentation(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.CStringType
		bufferVal,
	)
	// *typing.RefType
	return StringRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1643402-cfdateformattercreateiso8601form?language=objc
func DateFormatterCreateISO8601Formatter(allocator AllocatorRef, formatOptions ISO8601DateFormatOptions) DateFormatterRef {
	rv := C.DateFormatterCreateISO8601Formatter(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFISO8601DateFormatOptions)(formatOptions),
	)
	// *typing.RefType
	return DateFormatterRef(rv)
}

// Returns the default encoding used by the operating system when it creates strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541720-cfstringgetsystemencoding?language=objc
func StringGetSystemEncoding() StringEncoding {
	rv := C.StringGetSystemEncoding()
	// *typing.AliasType
	return StringEncoding(rv)
}

// Quickly obtains a pointer to a C-string buffer containing the characters of a string in a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542133-cfstringgetcstringptr?language=objc
func StringGetCStringPtr(theString StringRef, encoding StringEncoding) string {
	rv := C.StringGetCStringPtr(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
	// *typing.CStringType
	return C.GoString(rv)
}

// Returns the default port number with which to connect to a CFSocket name server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541492-cfsocketgetdefaultnameregistrypo?language=objc
func SocketGetDefaultNameRegistryPortNumber() uint16 {
	rv := C.SocketGetDefaultNameRegistryPortNumber()
	// *typing.PrimitiveType
	return uint16(rv)
}

// Creates a new immutable data with the bitmap representation from the given character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542820-cfcharactersetcreatebitmaprepres?language=objc
func CharacterSetCreateBitmapRepresentation(alloc AllocatorRef, theSet CharacterSetRef) DataRef {
	rv := C.CharacterSetCreateBitmapRepresentation(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(theSet),
	)
	// *typing.RefType
	return DataRef(rv)
}

// Returns the current system absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543542-cfabsolutetimegetcurrent?language=objc
func AbsoluteTimeGetCurrent() AbsoluteTime {
	rv := C.AbsoluteTimeGetCurrent()
	// *typing.AliasType
	return AbsoluteTime(rv)
}

// Returns a Boolean value that indicates whether a given character is a low character in a surrogate pair. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541963-cfstringissurrogatelowcharacter?language=objc
func StringIsSurrogateLowCharacter(character uint16) bool {
	rv := C.StringIsSurrogateLowCharacter(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.UniChar)(character),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/3684870-cfbundleisexecutableloadableforu?language=objc
func BundleIsExecutableLoadableForURL(url URLRef) bool {
	rv := C.BundleIsExecutableLoadableForURL(
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Searches a directory and constructs an array of CFBundle objects from all valid bundles in the specified directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537125-cfbundlecreatebundlesfromdirecto?language=objc
func BundleCreateBundlesFromDirectory(allocator AllocatorRef, directoryURL URLRef, bundleType StringRef) ArrayRef {
	rv := C.BundleCreateBundlesFromDirectory(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(directoryURL),
		// *typing.RefType
		unsafe.Pointer(bundleType),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Disables the callback function of a CFSocket object for certain types of socket activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542983-cfsocketdisablecallbacks?language=objc
func SocketDisableCallBacks(s SocketRef, callBackTypes OptionFlags) {
	C.SocketDisableCallBacks(
		// *typing.RefType
		unsafe.Pointer(s),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(callBackTypes),
	)
}

// Returns a format string for the given number formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1390726-cfnumberformattergetformat?language=objc
func NumberFormatterGetFormat(formatter NumberFormatterRef) StringRef {
	rv := C.NumberFormatterGetFormat(
		// *typing.RefType
		unsafe.Pointer(formatter),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns any additional resource specifiers after the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542114-cfurlcopyresourcespecifier?language=objc
func URLCopyResourceSpecifier(anURL URLRef) StringRef {
	rv := C.URLCopyResourceSpecifier(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Creates a new mutable bit vector from a pre-existing bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543572-cfbitvectorcreatemutablecopy?language=objc
func BitVectorCreateMutableCopy(allocator AllocatorRef, capacity Index, bv BitVectorRef) unsafe.Pointer {
	rv := C.BitVectorCreateMutableCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(capacity),
		// *typing.RefType
		unsafe.Pointer(bv),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns the location of a bundle’s main executable code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537086-cfbundlecopyexecutableurl?language=objc
func BundleCopyExecutableURL(bundle BundleRef) URLRef {
	rv := C.BundleCopyExecutableURL(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Sets the default time zone for your application the given time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543481-cftimezonesetdefault?language=objc
func TimeZoneSetDefault(tz TimeZoneRef) {
	C.TimeZoneSetDefault(
		// *typing.RefType
		unsafe.Pointer(tz),
	)
}

// Creates a new mutable character set with the values from another character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543531-cfcharactersetcreatemutablecopy?language=objc
func CharacterSetCreateMutableCopy(alloc AllocatorRef, theSet CharacterSetRef) unsafe.Pointer {
	rv := C.CharacterSetCreateMutableCopy(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(theSet),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns a bundle’s package type and creator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537089-cfbundlegetpackageinfo?language=objc
func BundleGetPackageInfo(bundle BundleRef, packageType *uint32, packageCreator *uint32) {
	C.BundleGetPackageInfo(
		// *typing.RefType
		unsafe.Pointer(bundle),
		// *typing.PointerType
		(*C.uint32_t)(unsafe.Pointer(&packageType)),
		// *typing.PointerType
		(*C.uint32_t)(unsafe.Pointer(&packageCreator)),
	)
}

// Returns the type identifier for the CFFileDescriptor opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1477583-cffiledescriptorgettypeid?language=objc
func FileDescriptorGetTypeID() TypeID {
	rv := C.FileDescriptorGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Creates a CFMutableString object whose Unicode character buffer is controlled externally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542308-cfstringcreatemutablewithexterna?language=objc
func StringCreateMutableWithExternalCharactersNoCopy(alloc AllocatorRef, chars *uint16, numChars Index, capacity Index, externalCharactersAllocator AllocatorRef) unsafe.Pointer {
	rv := C.StringCreateMutableWithExternalCharactersNoCopy(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.PointerType
		(*C.UniChar)(unsafe.Pointer(&chars)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(numChars),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(capacity),
		// *typing.RefType
		unsafe.Pointer(externalCharactersAllocator),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns the number (in terms of UTF-16 code pairs) of Unicode characters in a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542853-cfstringgetlength?language=objc
func StringGetLength(theString StringRef) Index {
	rv := C.StringGetLength(
		// *typing.RefType
		unsafe.Pointer(theString),
	)
	// *typing.AliasType
	return Index(rv)
}

// Converts a 32-bit integer from the host’s native byte order to little-endian format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425248-cfswapint32hosttolittle?language=objc
func SwapInt32HostToLittle(arg uint32) uint32 {
	rv := C.SwapInt32HostToLittle(
		// *typing.PrimitiveType
		C.uint32_t(arg),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543621-cfurlcreateresourcepropertiesfor?language=objc
func URLCreateResourcePropertiesForKeysFromBookmarkData(allocator AllocatorRef, resourcePropertiesToReturn ArrayRef, bookmark DataRef) DictionaryRef {
	rv := C.URLCreateResourcePropertiesForKeysFromBookmarkData(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(resourcePropertiesToReturn),
		// *typing.RefType
		unsafe.Pointer(bookmark),
	)
	// *typing.RefType
	return DictionaryRef(rv)
}

// Creates a new immutable array with the values from another array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1388792-cfarraycreatecopy?language=objc
func ArrayCreateCopy(allocator AllocatorRef, theArray ArrayRef) ArrayRef {
	rv := C.ArrayCreateCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(theArray),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Enlarges a string, padding it with specified characters, or truncates the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543484-cfstringpad?language=objc
func StringPad(theString unsafe.Pointer, padString StringRef, length Index, indexIntoPad Index) {
	C.StringPad(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.RefType
		unsafe.Pointer(padString),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(length),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(indexIntoPad),
	)
}

// Returns a copy of the logical calendar for the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533497-cfcalendarcopycurrent?language=objc
func CalendarCopyCurrent() CalendarRef {
	rv := C.CalendarCopyCurrent()
	// *typing.RefType
	return CalendarRef(rv)
}

// Returns the query string of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542792-cfurlcopyquerystring?language=objc
func URLCopyQueryString(anURL URLRef, charactersToLeaveEscaped StringRef) StringRef {
	rv := C.URLCopyQueryString(
		// *typing.RefType
		unsafe.Pointer(anURL),
		// *typing.RefType
		unsafe.Pointer(charactersToLeaveEscaped),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns the array of canonicalized language IDs that the user prefers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542887-cflocalecopypreferredlanguages?language=objc
func LocaleCopyPreferredLanguages() ArrayRef {
	rv := C.LocaleCopyPreferredLanguages()
	// *typing.RefType
	return ArrayRef(rv)
}

// Returns the application’s Darwin notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542572-cfnotificationcentergetdarwinnot?language=objc
func NotificationCenterGetDarwinNotifyCenter() NotificationCenterRef {
	rv := C.NotificationCenterGetDarwinNotifyCenter()
	// *typing.RefType
	return NotificationCenterRef(rv)
}

// Creates an immutable CFData object from an external (client-owned) byte buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541971-cfdatacreatewithbytesnocopy?language=objc
func DataCreateWithBytesNoCopy(allocator AllocatorRef, bytes *uint8, length Index, bytesDeallocator AllocatorRef) DataRef {
	rv := C.DataCreateWithBytesNoCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&bytes)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(length),
		// *typing.RefType
		unsafe.Pointer(bytesDeallocator),
	)
	// *typing.RefType
	return DataRef(rv)
}

// Returns the location of a bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537142-cfbundlecopybundleurl?language=objc
func BundleCopyBundleURL(bundle BundleRef) URLRef {
	rv := C.BundleCopyBundleURL(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns the string for an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542271-cfattributedstringgetstring?language=objc
func AttributedStringGetString(aStr AttributedStringRef) StringRef {
	rv := C.AttributedStringGetString(
		// *typing.RefType
		unsafe.Pointer(aStr),
	)
	// *typing.RefType
	return StringRef(rv)
}

// For the specified domain, writes all pending changes to preference data to permanent storage, and reads latest preference data from permanent storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515504-cfpreferencessynchronize?language=objc
func PreferencesSynchronize(applicationID StringRef, userName StringRef, hostName StringRef) bool {
	rv := C.PreferencesSynchronize(
		// *typing.RefType
		unsafe.Pointer(applicationID),
		// *typing.RefType
		unsafe.Pointer(userName),
		// *typing.RefType
		unsafe.Pointer(hostName),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the ordering parameter for a CFRunLoopSource object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541533-cfrunloopsourcegetorder?language=objc
func RunLoopSourceGetOrder(source RunLoopSourceRef) Index {
	rv := C.RunLoopSourceGetOrder(
		// *typing.RefType
		unsafe.Pointer(source),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns the first child of a tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401781-cftreegetfirstchild?language=objc
func TreeGetFirstChild(tree TreeRef) TreeRef {
	rv := C.TreeGetFirstChild(
		// *typing.RefType
		unsafe.Pointer(tree),
	)
	// *typing.RefType
	return TreeRef(rv)
}

// Removes suite preferences from an application’s search chain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515501-cfpreferencesremovesuitepreferen?language=objc
func PreferencesRemoveSuitePreferencesFromApp(applicationID StringRef, suiteID StringRef) {
	C.PreferencesRemoveSuitePreferencesFromApp(
		// *typing.RefType
		unsafe.Pointer(applicationID),
		// *typing.RefType
		unsafe.Pointer(suiteID),
	)
}

// Returns the time style used to create the given date formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1396282-cfdateformattergettimestyle?language=objc
func DateFormatterGetTimeStyle(formatter DateFormatterRef) DateFormatterStyle {
	rv := C.DateFormatterGetTimeStyle(
		// *typing.RefType
		unsafe.Pointer(formatter),
	)
	// *typing.AliasType
	return DateFormatterStyle(rv)
}

// Determines whether two Core Foundation objects are considered equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521287-cfequal?language=objc
func Equal(cf1 TypeRef, cf2 TypeRef) bool {
	rv := C.Equal(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(cf1),
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(cf2),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Converts a 16-bit integer from the host’s native byte order to big-endian format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425301-cfswapint16hosttobig?language=objc
func SwapInt16HostToBig(arg uint16) uint16 {
	rv := C.SwapInt16HostToBig(
		// *typing.PrimitiveType
		C.uint16_t(arg),
	)
	// *typing.PrimitiveType
	return uint16(rv)
}

// Converts a 16-bit integer from little-endian format to the host’s native byte order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425272-cfswapint16littletohost?language=objc
func SwapInt16LittleToHost(arg uint16) uint16 {
	rv := C.SwapInt16LittleToHost(
		// *typing.PrimitiveType
		C.uint16_t(arg),
	)
	// *typing.PrimitiveType
	return uint16(rv)
}

// Gets the default allocator object for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521325-cfallocatorgetdefault?language=objc
func AllocatorGetDefault() AllocatorRef {
	rv := C.AllocatorGetDefault()
	// *typing.RefType
	return AllocatorRef(rv)
}

// Sets the owner ID associated with a CFFileSecurityRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426528-cffilesecuritysetowner?language=objc
func FileSecuritySetOwner(fileSec FileSecurityRef, owner int) bool {
	rv := C.FileSecuritySetOwner(
		// *typing.RefType
		unsafe.Pointer(fileSec),
		// *typing.AliasType
		// *typing.StructType
		(C.uid_t)(owner),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a bundle’s plug-in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537100-cfbundlegetplugin?language=objc
func BundleGetPlugIn(bundle BundleRef) unsafe.Pointer {
	rv := C.BundleGetPlugIn(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Determines whether a CFNumber object contains a value stored as one of the defined floating point types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543131-cfnumberisfloattype?language=objc
func NumberIsFloatType(number NumberRef) bool {
	rv := C.NumberIsFloatType(
		// *typing.RefType
		unsafe.Pointer(number),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Displays a user notification dialog that does not need a user response. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534478-cfusernotificationdisplaynotice?language=objc
func UserNotificationDisplayNotice(timeout TimeInterval, flags OptionFlags, iconURL URLRef, soundURL URLRef, localizationURL URLRef, alertHeader StringRef, alertMessage StringRef, defaultButtonTitle StringRef) int32 {
	rv := C.UserNotificationDisplayNotice(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFTimeInterval)(timeout),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(flags),
		// *typing.RefType
		unsafe.Pointer(iconURL),
		// *typing.RefType
		unsafe.Pointer(soundURL),
		// *typing.RefType
		unsafe.Pointer(localizationURL),
		// *typing.RefType
		unsafe.Pointer(alertHeader),
		// *typing.RefType
		unsafe.Pointer(alertMessage),
		// *typing.RefType
		unsafe.Pointer(defaultButtonTitle),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Returns the type identifier for the CFBundle opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537126-cfbundlegettypeid?language=objc
func BundleGetTypeID() TypeID {
	rv := C.BundleGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Obtains the number of bytes likely to be allocated upon a specific request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521138-cfallocatorgetpreferredsizeforsi?language=objc
func AllocatorGetPreferredSizeForSize(allocator AllocatorRef, size Index, hint OptionFlags) Index {
	rv := C.AllocatorGetPreferredSizeForSize(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(size),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(hint),
	)
	// *typing.AliasType
	return Index(rv)
}

// Normalizes the string into the specified form as described in Unicode Technical Report #15. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542778-cfstringnormalize?language=objc
func StringNormalize(theString unsafe.Pointer, theForm StringNormalizationForm) {
	C.StringNormalize(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFStringNormalizationForm)(theForm),
	)
}

// Removes all values from a CFMutableSet object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1520440-cfsetremoveallvalues?language=objc
func SetRemoveAllValues(theSet unsafe.Pointer) {
	C.SetRemoveAllValues(
		// *typing.RefType
		unsafe.Pointer(theSet),
	)
}

// Retains a Core Foundation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521269-cfretain?language=objc
func Retain(cf TypeRef) TypeRef {
	rv := C.Retain(
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(cf),
	)
	// *typing.AliasType
	return TypeRef(rv)
}

// Creates an immutable bit vector that is a copy of another bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542453-cfbitvectorcreatecopy?language=objc
func BitVectorCreateCopy(allocator AllocatorRef, bv BitVectorRef) BitVectorRef {
	rv := C.BitVectorCreateCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(bv),
	)
	// *typing.RefType
	return BitVectorRef(rv)
}

// Adds a new child to a tree as the last in its list of children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401794-cftreeappendchild?language=objc
func TreeAppendChild(tree TreeRef, newChild TreeRef) {
	C.TreeAppendChild(
		// *typing.RefType
		unsafe.Pointer(tree),
		// *typing.RefType
		unsafe.Pointer(newChild),
	)
}

// Clears the previously determined system time zone, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541629-cftimezoneresetsystem?language=objc
func TimeZoneResetSystem() {
	C.TimeZoneResetSystem()
}

// Returns the type identifier of all CFWriteStream objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539668-cfwritestreamgettypeid?language=objc
func WriteStreamGetTypeID() TypeID {
	rv := C.WriteStreamGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the ordering parameter for a CFRunLoopObserver object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543674-cfrunloopobservergetorder?language=objc
func RunLoopObserverGetOrder(observer RunLoopObserverRef) Index {
	rv := C.RunLoopObserverGetOrder(
		// *typing.RefType
		unsafe.Pointer(observer),
	)
	// *typing.AliasType
	return Index(rv)
}

// Unloads the main executable for the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537149-cfbundleunloadexecutable?language=objc
func BundleUnloadExecutable(bundle BundleRef) {
	C.BundleUnloadExecutable(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543161-cfurlisfilereferenceurl?language=objc
func URLIsFileReferenceURL(url URLRef) bool {
	rv := C.URLIsFileReferenceURL(
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542381-cfurlstopaccessingsecurityscoped?language=objc
func URLStopAccessingSecurityScopedResource(url URLRef) {
	C.URLStopAccessingSecurityScopedResource(
		// *typing.RefType
		unsafe.Pointer(url),
	)
}

// Returns the number of elapsed seconds between the given CFDate objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542497-cfdategettimeintervalsincedate?language=objc
func DateGetTimeIntervalSinceDate(theDate DateRef, otherDate DateRef) TimeInterval {
	rv := C.DateGetTimeIntervalSinceDate(
		// *typing.RefType
		unsafe.Pointer(theDate),
		// *typing.RefType
		unsafe.Pointer(otherDate),
	)
	// *typing.AliasType
	return TimeInterval(rv)
}

// Returns the native file descriptor for a given CFFileDescriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1477611-cffiledescriptorgetnativedescrip?language=objc
func FileDescriptorGetNativeDescriptor(f FileDescriptorRef) FileDescriptorNativeDescriptor {
	rv := C.FileDescriptorGetNativeDescriptor(
		// *typing.RefType
		unsafe.Pointer(f),
	)
	// *typing.AliasType
	return FileDescriptorNativeDescriptor(rv)
}

// Creates a CFFileSecurityRef object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426509-cffilesecuritycreate?language=objc
func FileSecurityCreate(allocator AllocatorRef) FileSecurityRef {
	rv := C.FileSecurityCreate(
		// *typing.RefType
		unsafe.Pointer(allocator),
	)
	// *typing.RefType
	return FileSecurityRef(rv)
}

// Returns a textual description of a Core Foundation type, as identified by its type ID, which can be used when debugging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521220-cfcopytypeiddescription?language=objc
func CopyTypeIDDescription(type_id TypeID) StringRef {
	rv := C.CopyTypeIDDescription(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFTypeID)(type_id),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Creates a new runloop source for a given CFFileDescriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1477593-cffiledescriptorcreaterunloopsou?language=objc
func FileDescriptorCreateRunLoopSource(allocator AllocatorRef, f FileDescriptorRef, order Index) RunLoopSourceRef {
	rv := C.FileDescriptorCreateRunLoopSource(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(f),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(order),
	)
	// *typing.RefType
	return RunLoopSourceRef(rv)
}

// Finds a token that includes the character at a given index, and set it as the current token. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542487-cfstringtokenizergototokenatinde?language=objc
func StringTokenizerGoToTokenAtIndex(tokenizer StringTokenizerRef, index Index) StringTokenizerTokenType {
	rv := C.StringTokenizerGoToTokenAtIndex(
		// *typing.RefType
		unsafe.Pointer(tokenizer),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(index),
	)
	// *typing.AliasType
	return StringTokenizerTokenType(rv)
}

// Returns the most compatible Mac OS script value for the given input encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541474-cfstringgetmostcompatiblemacstri?language=objc
func StringGetMostCompatibleMacStringEncoding(encoding StringEncoding) StringEncoding {
	rv := C.StringGetMostCompatibleMacStringEncoding(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
	// *typing.AliasType
	return StringEncoding(rv)
}

// Creates a single string from the individual CFString objects that comprise the elements of an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542718-cfstringcreatebycombiningstrings?language=objc
func StringCreateByCombiningStrings(alloc AllocatorRef, theArray ArrayRef, separatorString StringRef) StringRef {
	rv := C.StringCreateByCombiningStrings(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(theArray),
		// *typing.RefType
		unsafe.Pointer(separatorString),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Creates an empty CFMutableData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542051-cfdatacreatemutable?language=objc
func DataCreateMutable(allocator AllocatorRef, capacity Index) unsafe.Pointer {
	rv := C.DataCreateMutable(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(capacity),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Constructs and returns the list of all keys set in the specified domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515512-cfpreferencescopykeylist?language=objc
func PreferencesCopyKeyList(applicationID StringRef, userName StringRef, hostName StringRef) ArrayRef {
	rv := C.PreferencesCopyKeyList(
		// *typing.RefType
		unsafe.Pointer(applicationID),
		// *typing.RefType
		unsafe.Pointer(userName),
		// *typing.RefType
		unsafe.Pointer(hostName),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Returns the location of the bundle’s support files directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537118-cfbundlecopysupportfilesdirector?language=objc
func BundleCopySupportFilesDirectoryURL(bundle BundleRef) URLRef {
	rv := C.BundleCopySupportFilesDirectoryURL(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns the type identifier for the CFMessagePort opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541943-cfmessageportgettypeid?language=objc
func MessagePortGetTypeID() TypeID {
	rv := C.MessagePortGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the location of a bundle’s private Frameworks directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537092-cfbundlecopyprivateframeworksurl?language=objc
func BundleCopyPrivateFrameworksURL(bundle BundleRef) URLRef {
	rv := C.BundleCopyPrivateFrameworksURL(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns the remote address to which a CFSocket object is connected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542819-cfsocketcopypeeraddress?language=objc
func SocketCopyPeerAddress(s SocketRef) DataRef {
	rv := C.SocketCopyPeerAddress(
		// *typing.RefType
		unsafe.Pointer(s),
	)
	// *typing.RefType
	return DataRef(rv)
}

// Reports whether or not a character set is a superset of another set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542915-cfcharactersetissupersetofset?language=objc
func CharacterSetIsSupersetOfSet(theSet CharacterSetRef, theOtherset CharacterSetRef) bool {
	rv := C.CharacterSetIsSupersetOfSet(
		// *typing.RefType
		unsafe.Pointer(theSet),
		// *typing.RefType
		unsafe.Pointer(theOtherset),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Removes the minimum value from a binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1509298-cfbinaryheapremoveminimumvalue?language=objc
func BinaryHeapRemoveMinimumValue(heap BinaryHeapRef) {
	C.BinaryHeapRemoveMinimumValue(
		// *typing.RefType
		unsafe.Pointer(heap),
	)
}

// Appends the characters of a string to those of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542205-cfstringappend?language=objc
func StringAppend(theString unsafe.Pointer, appendedString StringRef) {
	C.StringAppend(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.RefType
		unsafe.Pointer(appendedString),
	)
}

// Creates a Universally Unique Identifier (UUID) object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542906-cfuuidcreate?language=objc
func UUIDCreate(alloc AllocatorRef) UUIDRef {
	rv := C.UUIDCreate(
		// *typing.RefType
		unsafe.Pointer(alloc),
	)
	// *typing.RefType
	return UUIDRef(rv)
}

// Returns the length of the attributed string in characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543267-cfattributedstringgetlength?language=objc
func AttributedStringGetLength(aStr AttributedStringRef) Index {
	rv := C.AttributedStringGetLength(
		// *typing.RefType
		unsafe.Pointer(aStr),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns a bundle’s localized information dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537148-cfbundlegetlocalinfodictionary?language=objc
func BundleGetLocalInfoDictionary(bundle BundleRef) DictionaryRef {
	rv := C.BundleGetLocalInfoDictionary(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.RefType
	return DictionaryRef(rv)
}

// Returns the user info dictionary for a given CFError. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1494648-cferrorcopyuserinfo?language=objc
func ErrorCopyUserInfo(err ErrorRef) DictionaryRef {
	rv := C.ErrorCopyUserInfo(
		// *typing.RefType
		unsafe.Pointer(err),
	)
	// *typing.RefType
	return DictionaryRef(rv)
}

// Returns the next sibling, adjacent to a given tree, in the parent's children list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401765-cftreegetnextsibling?language=objc
func TreeGetNextSibling(tree TreeRef) TreeRef {
	rv := C.TreeGetNextSibling(
		// *typing.RefType
		unsafe.Pointer(tree),
	)
	// *typing.RefType
	return TreeRef(rv)
}

// Returns the user name from a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542692-cfurlcopyusername?language=objc
func URLCopyUserName(anURL URLRef) StringRef {
	rv := C.URLCopyUserName(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Creates a new CFURL object by resolving the relative portion of a URL, specified as bytes, against its given base URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542795-cfurlcreateabsoluteurlwithbytes?language=objc
func URLCreateAbsoluteURLWithBytes(alloc AllocatorRef, relativeURLBytes *uint8, length Index, encoding StringEncoding, baseURL URLRef, useCompatibilityMode bool) URLRef {
	rv := C.URLCreateAbsoluteURLWithBytes(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&relativeURLBytes)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(length),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
		// *typing.RefType
		unsafe.Pointer(baseURL),
		// *typing.PrimitiveType
		C.bool(useCompatibilityMode),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns a string representation of the given date using the specified date formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1396312-cfdateformattercreatestringwithd?language=objc
func DateFormatterCreateStringWithDate(allocator AllocatorRef, formatter DateFormatterRef, date DateRef) StringRef {
	rv := C.DateFormatterCreateStringWithDate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(formatter),
		// *typing.RefType
		unsafe.Pointer(date),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Computes the difference between the two absolute times, in terms of specified calendrical components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533470-cfcalendargetcomponentdifference?language=objc
func CalendarGetComponentDifference(calendar CalendarRef, startingAT AbsoluteTime, resultAT AbsoluteTime, options OptionFlags, componentDesc string) bool {
	componentDescVal := C.CString(componentDesc)
	defer C.free(unsafe.Pointer(componentDescVal))
	rv := C.CalendarGetComponentDifference(
		// *typing.RefType
		unsafe.Pointer(calendar),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(startingAT),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(resultAT),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(options),
		// *typing.CStringType
		componentDescVal,
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the child of a tree at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401812-cftreegetchildatindex?language=objc
func TreeGetChildAtIndex(tree TreeRef, idx Index) TreeRef {
	rv := C.TreeGetChildAtIndex(
		// *typing.RefType
		unsafe.Pointer(tree),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(idx),
	)
	// *typing.RefType
	return TreeRef(rv)
}

// Returns the type identifier for the CFDictionary opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1516806-cfdictionarygettypeid?language=objc
func DictionaryGetTypeID() TypeID {
	rv := C.DictionaryGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns a plug-in's bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493851-cfplugingetbundle?language=objc
func PlugInGetBundle(plugIn unsafe.Pointer) BundleRef {
	rv := C.PlugInGetBundle(
		// *typing.RefType
		unsafe.Pointer(plugIn),
	)
	// *typing.RefType
	return BundleRef(rv)
}

// Returns an array containing copies of the URL locations for a specified bundle, resource, and localization name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537162-cfbundlecopyresourceurlsoftypefo?language=objc
func BundleCopyResourceURLsOfTypeForLocalization(bundle BundleRef, resourceType StringRef, subDirName StringRef, localizationName StringRef) ArrayRef {
	rv := C.BundleCopyResourceURLsOfTypeForLocalization(
		// *typing.RefType
		unsafe.Pointer(bundle),
		// *typing.RefType
		unsafe.Pointer(resourceType),
		// *typing.RefType
		unsafe.Pointer(subDirName),
		// *typing.RefType
		unsafe.Pointer(localizationName),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Returns a CFUUID object from raw UUID bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542189-cfuuidgetconstantuuidwithbytes?language=objc
func UUIDGetConstantUUIDWithBytes(alloc AllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) UUIDRef {
	rv := C.UUIDGetConstantUUIDWithBytes(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.PrimitiveType
		C.uint8_t(byte0),
		// *typing.PrimitiveType
		C.uint8_t(byte1),
		// *typing.PrimitiveType
		C.uint8_t(byte2),
		// *typing.PrimitiveType
		C.uint8_t(byte3),
		// *typing.PrimitiveType
		C.uint8_t(byte4),
		// *typing.PrimitiveType
		C.uint8_t(byte5),
		// *typing.PrimitiveType
		C.uint8_t(byte6),
		// *typing.PrimitiveType
		C.uint8_t(byte7),
		// *typing.PrimitiveType
		C.uint8_t(byte8),
		// *typing.PrimitiveType
		C.uint8_t(byte9),
		// *typing.PrimitiveType
		C.uint8_t(byte10),
		// *typing.PrimitiveType
		C.uint8_t(byte11),
		// *typing.PrimitiveType
		C.uint8_t(byte12),
		// *typing.PrimitiveType
		C.uint8_t(byte13),
		// *typing.PrimitiveType
		C.uint8_t(byte14),
		// *typing.PrimitiveType
		C.uint8_t(byte15),
	)
	// *typing.RefType
	return UUIDRef(rv)
}

// Returns whether or not a time zone is in daylight savings time at a specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541793-cftimezoneisdaylightsavingtime?language=objc
func TimeZoneIsDaylightSavingTime(tz TimeZoneRef, at AbsoluteTime) bool {
	rv := C.TimeZoneIsDaylightSavingTime(
		// *typing.RefType
		unsafe.Pointer(tz),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(at),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Removes all cached resource values and temporary resource values from the URL object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541959-cfurlclearresourcepropertycache?language=objc
func URLClearResourcePropertyCache(url URLRef) {
	C.URLClearResourcePropertyCache(
		// *typing.RefType
		unsafe.Pointer(url),
	)
}

// Returns a Boolean value that indicates whether a CFSocket object is valid and able to send or receive messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543588-cfsocketisvalid?language=objc
func SocketIsValid(s SocketRef) bool {
	rv := C.SocketIsValid(
		// *typing.RefType
		unsafe.Pointer(s),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a data pointer to a symbol of the given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537120-cfbundlegetdatapointerforname?language=objc
func BundleGetDataPointerForName(bundle BundleRef, symbolName StringRef) unsafe.Pointer {
	rv := C.BundleGetDataPointerForName(
		// *typing.RefType
		unsafe.Pointer(bundle),
		// *typing.RefType
		unsafe.Pointer(symbolName),
	)
	// *typing.VoidPointerType
	return unsafe.Pointer(rv)
}

// Returns the Cocoa encoding constant that maps most closely to a given Core Foundation encoding constant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543046-cfstringconvertencodingtonsstrin?language=objc
func StringConvertEncodingToNSStringEncoding(encoding StringEncoding) int32 {
	rv := C.StringConvertEncodingToNSStringEncoding(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Returns the local address of a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543303-cfsocketcopyaddress?language=objc
func SocketCopyAddress(s SocketRef) DataRef {
	rv := C.SocketCopyAddress(
		// *typing.RefType
		unsafe.Pointer(s),
	)
	// *typing.RefType
	return DataRef(rv)
}

// Creates a mutable bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543376-cfbitvectorcreatemutable?language=objc
func BitVectorCreateMutable(allocator AllocatorRef, capacity Index) unsafe.Pointer {
	rv := C.BitVectorCreateMutable(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(capacity),
	)
	// *typing.RefType
	return unsafe.Pointer(rv)
}

// Returns a pointer to a stream’s internal buffer of unread data, if possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539692-cfreadstreamgetbuffer?language=objc
func ReadStreamGetBuffer(stream ReadStreamRef, maxBytesToRead Index, numBytesRead *Index) *uint8 {
	rv := C.ReadStreamGetBuffer(
		// *typing.RefType
		unsafe.Pointer(stream),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(maxBytesToRead),
		// *typing.PointerType
		(*C.CFIndex)(unsafe.Pointer(&numBytesRead)),
	)
	// *typing.PointerType
	return *(**uint8)(unsafe.Pointer(&rv))
}

// Writes to permanent storage all pending changes to the preference data for the application, and reads the latest preference data from permanent storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515510-cfpreferencesappsynchronize?language=objc
func PreferencesAppSynchronize(applicationID StringRef) bool {
	rv := C.PreferencesAppSynchronize(
		// *typing.RefType
		unsafe.Pointer(applicationID),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a Boolean value that indicates whether the run loop is waiting for an event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542956-cfrunloopiswaiting?language=objc
func RunLoopIsWaiting(rl RunLoopRef) bool {
	rv := C.RunLoopIsWaiting(
		// *typing.RefType
		unsafe.Pointer(rl),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a copy of a given URL with the last path component deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543199-cfurlcreatecopydeletinglastpathc?language=objc
func URLCreateCopyDeletingLastPathComponent(allocator AllocatorRef, url URLRef) URLRef {
	rv := C.URLCreateCopyDeletingLastPathComponent(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Returns a pointer to a list of string encodings supported by the current system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542932-cfstringgetlistofavailableencodi?language=objc
func StringGetListOfAvailableEncodings() *StringEncoding {
	rv := C.StringGetListOfAvailableEncodings()
	// *typing.PointerType
	return *(**StringEncoding)(unsafe.Pointer(&rv))
}

// Creates a new character set with the values in the given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542761-cfcharactersetcreatewithcharacte?language=objc
func CharacterSetCreateWithCharactersInString(alloc AllocatorRef, theString StringRef) CharacterSetRef {
	rv := C.CharacterSetCreateWithCharactersInString(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(theString),
	)
	// *typing.RefType
	return CharacterSetRef(rv)
}

// Sets the abbreviation dictionary to a given dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542698-cftimezonesetabbreviationdiction?language=objc
func TimeZoneSetAbbreviationDictionary(dict DictionaryRef) {
	C.TimeZoneSetAbbreviationDictionary(
		// *typing.RefType
		unsafe.Pointer(dict),
	)
}

// Creates a copy of a given URL and appends a path extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541648-cfurlcreatecopyappendingpathexte?language=objc
func URLCreateCopyAppendingPathExtension(allocator AllocatorRef, url URLRef, extension StringRef) URLRef {
	rv := C.URLCreateCopyAppendingPathExtension(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(url),
		// *typing.RefType
		unsafe.Pointer(extension),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Changes all lowercase alphabetical characters in a CFMutableString object to uppercase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542348-cfstringuppercase?language=objc
func StringUppercase(theString unsafe.Pointer, locale LocaleRef) {
	C.StringUppercase(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.RefType
		unsafe.Pointer(locale),
	)
}

// Returns a preference value for a given domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1515500-cfpreferencescopyvalue?language=objc
func PreferencesCopyValue(key StringRef, applicationID StringRef, userName StringRef, hostName StringRef) PropertyListRef {
	rv := C.PreferencesCopyValue(
		// *typing.RefType
		unsafe.Pointer(key),
		// *typing.RefType
		unsafe.Pointer(applicationID),
		// *typing.RefType
		unsafe.Pointer(userName),
		// *typing.RefType
		unsafe.Pointer(hostName),
	)
	// *typing.AliasType
	return PropertyListRef(rv)
}

// Returns the type identifier for the CFPlugIn opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493861-cfplugingettypeid?language=objc
func PlugInGetTypeID() TypeID {
	rv := C.PlugInGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Creates an array of CFString objects from a single CFString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541864-cfstringcreatearraybyseparatings?language=objc
func StringCreateArrayBySeparatingStrings(alloc AllocatorRef, theString StringRef, separatorString StringRef) ArrayRef {
	rv := C.StringCreateArrayBySeparatingStrings(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.RefType
		unsafe.Pointer(separatorString),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Creates a CFPlugIn instance of a given type using a given factory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493857-cfplugininstancecreate?language=objc
func PlugInInstanceCreate(allocator AllocatorRef, factoryUUID UUIDRef, typeUUID UUIDRef) unsafe.Pointer {
	rv := C.PlugInInstanceCreate(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(factoryUUID),
		// *typing.RefType
		unsafe.Pointer(typeUUID),
	)
	// *typing.VoidPointerType
	return unsafe.Pointer(rv)
}

// Returns the locale object used to create the given date formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1396272-cfdateformattergetlocale?language=objc
func DateFormatterGetLocale(formatter DateFormatterRef) LocaleRef {
	rv := C.DateFormatterGetLocale(
		// *typing.RefType
		unsafe.Pointer(formatter),
	)
	// *typing.RefType
	return LocaleRef(rv)
}

// Converts a 32-bit integer from little-endian format to the host’s native byte order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425287-cfswapint32littletohost?language=objc
func SwapInt32LittleToHost(arg uint32) uint32 {
	rv := C.SwapInt32LittleToHost(
		// *typing.PrimitiveType
		C.uint32_t(arg),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Returns a predefined character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541624-cfcharactersetgetpredefined?language=objc
func CharacterSetGetPredefined(theSetIdentifier CharacterSetPredefinedSet) CharacterSetRef {
	rv := C.CharacterSetGetPredefined(
		// *typing.AliasType
		// *typing.AliasType
		(C.CFCharacterSetPredefinedSet)(theSetIdentifier),
	)
	// *typing.RefType
	return CharacterSetRef(rv)
}

// Returns the type identifier for the CFData opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541611-cfdatagettypeid?language=objc
func DataGetTypeID() TypeID {
	rv := C.DataGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Waits for the user to respond to a notification or for the notification to time out. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534477-cfusernotificationreceiverespons?language=objc
func UserNotificationReceiveResponse(userNotification UserNotificationRef, timeout TimeInterval, responseFlags *OptionFlags) int32 {
	rv := C.UserNotificationReceiveResponse(
		// *typing.RefType
		unsafe.Pointer(userNotification),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFTimeInterval)(timeout),
		// *typing.PointerType
		(*C.CFOptionFlags)(unsafe.Pointer(&responseFlags)),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Returns the type identifier of the CFCharacterSet opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543628-cfcharactersetgettypeid?language=objc
func CharacterSetGetTypeID() TypeID {
	rv := C.CharacterSetGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the application’s local notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542139-cfnotificationcentergetlocalcent?language=objc
func NotificationCenterGetLocalCenter() NotificationCenterRef {
	rv := C.NotificationCenterGetLocalCenter()
	// *typing.RefType
	return NotificationCenterRef(rv)
}

// Invalidates a CFRunLoopSource object, stopping it from ever firing again. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543336-cfrunloopsourceinvalidate?language=objc
func RunLoopSourceInvalidate(source RunLoopSourceRef) {
	C.RunLoopSourceInvalidate(
		// *typing.RefType
		unsafe.Pointer(source),
	)
}

// Converts a 64-bit integer from big-endian format to the host’s native byte order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425280-cfswapint64bigtohost?language=objc
func SwapInt64BigToHost(arg uint64) uint64 {
	rv := C.SwapInt64BigToHost(
		// *typing.PrimitiveType
		C.uint64_t(arg),
	)
	// *typing.PrimitiveType
	return uint64(rv)
}

// Returns the number of key-value pairs in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1516741-cfdictionarygetcount?language=objc
func DictionaryGetCount(theDict DictionaryRef) Index {
	rv := C.DictionaryGetCount(
		// *typing.RefType
		unsafe.Pointer(theDict),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns the path extension of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543269-cfurlcopypathextension?language=objc
func URLCopyPathExtension(url URLRef) StringRef {
	rv := C.URLCopyPathExtension(
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Returns a Boolean value that indicates whether a CFMessagePort object is valid and able to send or receive messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541942-cfmessageportisvalid?language=objc
func MessagePortIsValid(ms MessagePortRef) bool {
	rv := C.MessagePortIsValid(
		// *typing.RefType
		unsafe.Pointer(ms),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Converts a 32-bit integer from the host’s native byte order to big-endian format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425254-cfswapint32hosttobig?language=objc
func SwapInt32HostToBig(arg uint32) uint32 {
	rv := C.SwapInt32HostToBig(
		// *typing.PrimitiveType
		C.uint32_t(arg),
	)
	// *typing.PrimitiveType
	return uint32(rv)
}

// Returns the number of bytes used by a CFNumber object to store its value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542080-cfnumbergetbytesize?language=objc
func NumberGetByteSize(number NumberRef) Index {
	rv := C.NumberGetByteSize(
		// *typing.RefType
		unsafe.Pointer(number),
	)
	// *typing.AliasType
	return Index(rv)
}

// Creates a copy of a given URL and appends a path component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542127-cfurlcreatecopyappendingpathcomp?language=objc
func URLCreateCopyAppendingPathComponent(allocator AllocatorRef, url URLRef, pathComponent StringRef, isDirectory bool) URLRef {
	rv := C.URLCreateCopyAppendingPathComponent(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(url),
		// *typing.RefType
		unsafe.Pointer(pathComponent),
		// *typing.PrimitiveType
		C.bool(isDirectory),
	)
	// *typing.RefType
	return URLRef(rv)
}

// Compares one string with another string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542911-cfstringcompare?language=objc
func StringCompare(theString1 StringRef, theString2 StringRef, compareOptions StringCompareFlags) ComparisonResult {
	rv := C.StringCompare(
		// *typing.RefType
		unsafe.Pointer(theString1),
		// *typing.RefType
		unsafe.Pointer(theString2),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFStringCompareFlags)(compareOptions),
	)
	// *typing.AliasType
	return ComparisonResult(rv)
}

// Returns a flag used to set the selected element of a pop-up menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534520-cfusernotificationpopupselection?language=objc
func UserNotificationPopUpSelection(n Index) OptionFlags {
	rv := C.UserNotificationPopUpSelection(
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(n),
	)
	// *typing.AliasType
	return OptionFlags(rv)
}

// Returns the byte order of the current computer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425298-cfbyteordergetcurrent?language=objc
func ByteOrderGetCurrent() ByteOrder {
	rv := C.ByteOrderGetCurrent()
	// *typing.AliasType
	return ByteOrder(rv)
}

// Opens a stream for reading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539743-cfreadstreamopen?language=objc
func ReadStreamOpen(stream ReadStreamRef) bool {
	rv := C.ReadStreamOpen(
		// *typing.RefType
		unsafe.Pointer(stream),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the type identifier for the CFDate opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542050-cfdategettypeid?language=objc
func DateGetTypeID() TypeID {
	rv := C.DateGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Converts a 16-bit integer from the host’s native byte order to little-endian format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1425285-cfswapint16hosttolittle?language=objc
func SwapInt16HostToLittle(arg uint16) uint16 {
	rv := C.SwapInt16HostToLittle(
		// *typing.PrimitiveType
		C.uint16_t(arg),
	)
	// *typing.PrimitiveType
	return uint16(rv)
}

// Creates an immutable copy of an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542122-cfattributedstringcreatecopy?language=objc
func AttributedStringCreateCopy(alloc AllocatorRef, aStr AttributedStringRef) AttributedStringRef {
	rv := C.AttributedStringCreateCopy(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(aStr),
	)
	// *typing.RefType
	return AttributedStringRef(rv)
}

// Sets the format string of a number formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1390694-cfnumberformattersetformat?language=objc
func NumberFormatterSetFormat(formatter NumberFormatterRef, formatString StringRef) {
	C.NumberFormatterSetFormat(
		// *typing.RefType
		unsafe.Pointer(formatter),
		// *typing.RefType
		unsafe.Pointer(formatString),
	)
}

// Returns an array of strings containing the names of all the time zones known to the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542883-cftimezonecopyknownnames?language=objc
func TimeZoneCopyKnownNames() ArrayRef {
	rv := C.TimeZoneCopyKnownNames()
	// *typing.RefType
	return ArrayRef(rv)
}

// Creates a CFRunLoopSource object for a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542740-cfsocketcreaterunloopsource?language=objc
func SocketCreateRunLoopSource(allocator AllocatorRef, s SocketRef, order Index) RunLoopSourceRef {
	rv := C.SocketCreateRunLoopSource(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(s),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(order),
	)
	// *typing.RefType
	return RunLoopSourceRef(rv)
}

// Returns the type identifier for the CFBitVector opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542151-cfbitvectorgettypeid?language=objc
func BitVectorGetTypeID() TypeID {
	rv := C.BitVectorGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Forms the union of two character sets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542265-cfcharactersetunion?language=objc
func CharacterSetUnion(theSet unsafe.Pointer, theOtherSet CharacterSetRef) {
	C.CharacterSetUnion(
		// *typing.RefType
		unsafe.Pointer(theSet),
		// *typing.RefType
		unsafe.Pointer(theOtherSet),
	)
}

// Returns an array that contains all the defined modes for a CFRunLoop object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542184-cfrunloopcopyallmodes?language=objc
func RunLoopCopyAllModes(rl RunLoopRef) ArrayRef {
	rv := C.RunLoopCopyAllModes(
		// *typing.RefType
		unsafe.Pointer(rl),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Returns the number of values currently in a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1520419-cfsetgetcount?language=objc
func SetGetCount(theSet SetRef) Index {
	rv := C.SetGetCount(
		// *typing.RefType
		unsafe.Pointer(theSet),
	)
	// *typing.AliasType
	return Index(rv)
}

// Opens a stream for writing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539753-cfwritestreamopen?language=objc
func WriteStreamOpen(stream WriteStreamRef) bool {
	rv := C.WriteStreamOpen(
		// *typing.RefType
		unsafe.Pointer(stream),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Sets flags that control certain behaviors of a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543135-cfsocketsetsocketflags?language=objc
func SocketSetSocketFlags(s SocketRef, flags OptionFlags) {
	C.SocketSetSocketFlags(
		// *typing.RefType
		unsafe.Pointer(s),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(flags),
	)
}

// Returns the type identifier for CFDateFormatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1396230-cfdateformattergettypeid?language=objc
func DateFormatterGetTypeID() TypeID {
	rv := C.DateFormatterGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the type identifier for the CFURL opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543482-cfurlgettypeid?language=objc
func URLGetTypeID() TypeID {
	rv := C.URLGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Removes all values from a mutable bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1469287-cfbagremoveallvalues?language=objc
func BagRemoveAllValues(theBag unsafe.Pointer) {
	C.BagRemoveAllValues(
		// *typing.RefType
		unsafe.Pointer(theBag),
	)
}

// Removes all values from a binary heap, making it empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1509321-cfbinaryheapremoveallvalues?language=objc
func BinaryHeapRemoveAllValues(heap BinaryHeapRef) {
	C.BinaryHeapRemoveAllValues(
		// *typing.RefType
		unsafe.Pointer(heap),
	)
}

// Flips a bit value in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542355-cfbitvectorflipbitatindex?language=objc
func BitVectorFlipBitAtIndex(bv unsafe.Pointer, idx Index) {
	C.BitVectorFlipBitAtIndex(
		// *typing.RefType
		unsafe.Pointer(bv),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(idx),
	)
}

// Returns the type identifier for the CFNotificationCenter opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542734-cfnotificationcentergettypeid?language=objc
func NotificationCenterGetTypeID() TypeID {
	rv := C.NotificationCenterGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the number of children in a tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401810-cftreegetchildcount?language=objc
func TreeGetChildCount(tree TreeRef) Index {
	rv := C.TreeGetChildCount(
		// *typing.RefType
		unsafe.Pointer(tree),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns the type identifier for the CFFileSecurityRef opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426530-cffilesecuritygettypeid?language=objc
func FileSecurityGetTypeID() TypeID {
	rv := C.FileSecurityGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the location of a resource contained in the specified bundle directory without requiring the creation of a CFBundle object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537138-cfbundlecopyresourceurlindirecto?language=objc
func BundleCopyResourceURLInDirectory(bundleURL URLRef, resourceName StringRef, resourceType StringRef, subDirName StringRef) URLRef {
	rv := C.BundleCopyResourceURLInDirectory(
		// *typing.RefType
		unsafe.Pointer(bundleURL),
		// *typing.RefType
		unsafe.Pointer(resourceName),
		// *typing.RefType
		unsafe.Pointer(resourceType),
		// *typing.RefType
		unsafe.Pointer(subDirName),
	)
	// *typing.RefType
	return URLRef(rv)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/3684869-cfbundleisexecutableloadable?language=objc
func BundleIsExecutableLoadable(bundle BundleRef) bool {
	rv := C.BundleIsExecutableLoadable(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the type identifier for the CFAllocator opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521328-cfallocatorgettypeid?language=objc
func AllocatorGetTypeID() TypeID {
	rv := C.AllocatorGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Sets a temporary resource value on the URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542384-cfurlsettemporaryresourcepropert?language=objc
func URLSetTemporaryResourcePropertyForKey(url URLRef, key StringRef, propertyValue TypeRef) {
	C.URLSetTemporaryResourcePropertyForKey(
		// *typing.RefType
		unsafe.Pointer(url),
		// *typing.RefType
		unsafe.Pointer(key),
		// *typing.AliasType
		// *typing.VoidPointerType
		(C.CFTypeRef)(propertyValue),
	)
}

// Not recommended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493885-cfplugininstancegettypeid?language=objc
func PlugInInstanceGetTypeID() TypeID {
	rv := C.PlugInInstanceGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Registers a type and its corresponding factory function with a CFPlugIn object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493853-cfpluginregisterplugintype?language=objc
func PlugInRegisterPlugInType(factoryUUID UUIDRef, typeUUID UUIDRef) bool {
	rv := C.PlugInRegisterPlugInType(
		// *typing.RefType
		unsafe.Pointer(factoryUUID),
		// *typing.RefType
		unsafe.Pointer(typeUUID),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the type identifier for all CFUUID objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543685-cfuuidgettypeid?language=objc
func UUIDGetTypeID() TypeID {
	rv := C.UUIDGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns the data that stores the information used by a time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543470-cftimezonegetdata?language=objc
func TimeZoneGetData(tz TimeZoneRef) DataRef {
	rv := C.TimeZoneGetData(
		// *typing.RefType
		unsafe.Pointer(tz),
	)
	// *typing.RefType
	return DataRef(rv)
}

// Returns the character direction for the specified ISO language code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542786-cflocalegetlanguagecharacterdire?language=objc
func LocaleGetLanguageCharacterDirection(isoLangCode StringRef) LocaleLanguageDirection {
	rv := C.LocaleGetLanguageCharacterDirection(
		// *typing.RefType
		unsafe.Pointer(isoLangCode),
	)
	// *typing.AliasType
	return LocaleLanguageDirection(rv)
}

// Creates a string from a buffer, containing characters in a specified encoding, that might serve as the backing store for the new string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543597-cfstringcreatewithbytesnocopy?language=objc
func StringCreateWithBytesNoCopy(alloc AllocatorRef, bytes *uint8, numBytes Index, encoding StringEncoding, isExternalRepresentation bool, contentsDeallocator AllocatorRef) StringRef {
	rv := C.StringCreateWithBytesNoCopy(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.PointerType
		(*C.uint8_t)(unsafe.Pointer(&bytes)),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(numBytes),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
		// *typing.PrimitiveType
		C.bool(isExternalRepresentation),
		// *typing.RefType
		unsafe.Pointer(contentsDeallocator),
	)
	// *typing.RefType
	return StringRef(rv)
}

// Removes a tree from its parent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1401808-cftreeremove?language=objc
func TreeRemove(tree TreeRef) {
	C.TreeRemove(
		// *typing.RefType
		unsafe.Pointer(tree),
	)
}

// Loads a bundle’s main executable code into memory and dynamically links it into the running application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537116-cfbundleloadexecutable?language=objc
func BundleLoadExecutable(bundle BundleRef) bool {
	rv := C.BundleLoadExecutable(
		// *typing.RefType
		unsafe.Pointer(bundle),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a Boolean value that indicates whether a CFRunLoopObserver repeats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542616-cfrunloopobserverdoesrepeat?language=objc
func RunLoopObserverDoesRepeat(observer RunLoopObserverRef) bool {
	rv := C.RunLoopObserverDoesRepeat(
		// *typing.RefType
		unsafe.Pointer(observer),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Invalidates a CFRunLoopTimer object, stopping it from ever firing again. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542231-cfrunlooptimerinvalidate?language=objc
func RunLoopTimerInvalidate(timer RunLoopTimerRef) {
	C.RunLoopTimerInvalidate(
		// *typing.RefType
		unsafe.Pointer(timer),
	)
}

// Returns the time in a given time zone of the next daylight saving time transition after a given time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541854-cftimezonegetnextdaylightsavingt?language=objc
func TimeZoneGetNextDaylightSavingTimeTransition(tz TimeZoneRef, at AbsoluteTime) AbsoluteTime {
	rv := C.TimeZoneGetNextDaylightSavingTimeTransition(
		// *typing.RefType
		unsafe.Pointer(tz),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(at),
	)
	// *typing.AliasType
	return AbsoluteTime(rv)
}

// Returns the minimum value in a binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1509325-cfbinaryheapgetminimum?language=objc
func BinaryHeapGetMinimum(heap BinaryHeapRef) unsafe.Pointer {
	rv := C.BinaryHeapGetMinimum(
		// *typing.RefType
		unsafe.Pointer(heap),
	)
	// *typing.VoidPointerType
	return unsafe.Pointer(rv)
}

// Given an array of possible localizations, returns the one or more of them that CFBundle would use in the current application context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537164-cfbundlecopypreferredlocalizatio?language=objc
func BundleCopyPreferredLocalizationsFromArray(locArray ArrayRef) ArrayRef {
	rv := C.BundleCopyPreferredLocalizationsFromArray(
		// *typing.RefType
		unsafe.Pointer(locArray),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Creates a CFRunLoopSource object for a CFMessagePort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542908-cfmessageportcreaterunloopsource?language=objc
func MessagePortCreateRunLoopSource(allocator AllocatorRef, local MessagePortRef, order Index) RunLoopSourceRef {
	rv := C.MessagePortCreateRunLoopSource(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.RefType
		unsafe.Pointer(local),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(order),
	)
	// *typing.RefType
	return RunLoopSourceRef(rv)
}

// Returns the type identifier for the CFLocale opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541930-cflocalegettypeid?language=objc
func LocaleGetTypeID() TypeID {
	rv := C.LocaleGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Computes the absolute time from components in a description string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533489-cfcalendarcomposeabsolutetime?language=objc
func CalendarComposeAbsoluteTime(calendar CalendarRef, at *AbsoluteTime, componentDesc string) bool {
	componentDescVal := C.CString(componentDesc)
	defer C.free(unsafe.Pointer(componentDescVal))
	rv := C.CalendarComposeAbsoluteTime(
		// *typing.RefType
		unsafe.Pointer(calendar),
		// *typing.PointerType
		(*C.CFAbsoluteTime)(unsafe.Pointer(&at)),
		// *typing.CStringType
		componentDescVal,
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Removes the given type from a plug-in’s list of registered types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1493874-cfpluginunregisterplugintype?language=objc
func PlugInUnregisterPlugInType(factoryUUID UUIDRef, typeUUID UUIDRef) bool {
	rv := C.PlugInUnregisterPlugInType(
		// *typing.RefType
		unsafe.Pointer(factoryUUID),
		// *typing.RefType
		unsafe.Pointer(typeUUID),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns a Boolean value that indicates whether a readable stream has data that can be read without blocking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539638-cfreadstreamhasbytesavailable?language=objc
func ReadStreamHasBytesAvailable(stream ReadStreamRef) bool {
	rv := C.ReadStreamHasBytesAvailable(
		// *typing.RefType
		unsafe.Pointer(stream),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Returns the Unicode character at a specified location in a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542730-cfstringgetcharacteratindex?language=objc
func StringGetCharacterAtIndex(theString StringRef, idx Index) uint16 {
	rv := C.StringGetCharacterAtIndex(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(idx),
	)
	// *typing.AliasType
	return uint16(rv)
}

// Returns the ordinal number of a calendrical unit within a larger unit at a specified absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533491-cfcalendargetordinalityofunit?language=objc
func CalendarGetOrdinalityOfUnit(calendar CalendarRef, smallerUnit CalendarUnit, biggerUnit CalendarUnit, at AbsoluteTime) Index {
	rv := C.CalendarGetOrdinalityOfUnit(
		// *typing.RefType
		unsafe.Pointer(calendar),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFCalendarUnit)(smallerUnit),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFCalendarUnit)(biggerUnit),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFAbsoluteTime)(at),
	)
	// *typing.AliasType
	return Index(rv)
}

// Recursively creates a copy of a given property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1429997-cfpropertylistcreatedeepcopy?language=objc
func PropertyListCreateDeepCopy(allocator AllocatorRef, propertyList PropertyListRef, mutabilityOption OptionFlags) PropertyListRef {
	rv := C.PropertyListCreateDeepCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFPropertyListRef)(propertyList),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFOptionFlags)(mutabilityOption),
	)
	// *typing.AliasType
	return PropertyListRef(rv)
}

// Returns a time zone object for a specified calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1533499-cfcalendarcopytimezone?language=objc
func CalendarCopyTimeZone(calendar CalendarRef) TimeZoneRef {
	rv := C.CalendarCopyTimeZone(
		// *typing.RefType
		unsafe.Pointer(calendar),
	)
	// *typing.RefType
	return TimeZoneRef(rv)
}

// Returns the port number from a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542357-cfurlgetportnumber?language=objc
func URLGetPortNumber(anURL URLRef) int32 {
	rv := C.URLGetPortNumber(
		// *typing.RefType
		unsafe.Pointer(anURL),
	)
	// *typing.PrimitiveType
	return int32(rv)
}

// Returns an array of CFString objects that represents all known legal ISO country codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543372-cflocalecopyisocountrycodes?language=objc
func LocaleCopyISOCountryCodes() ArrayRef {
	rv := C.LocaleCopyISOCountryCodes()
	// *typing.RefType
	return ArrayRef(rv)
}

// Creates a new mutable or fixed-mutable binary heap with the values from a pre-existing binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1509295-cfbinaryheapcreatecopy?language=objc
func BinaryHeapCreateCopy(allocator AllocatorRef, capacity Index, heap BinaryHeapRef) BinaryHeapRef {
	rv := C.BinaryHeapCreateCopy(
		// *typing.RefType
		unsafe.Pointer(allocator),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(capacity),
		// *typing.RefType
		unsafe.Pointer(heap),
	)
	// *typing.RefType
	return BinaryHeapRef(rv)
}

// Gets the group ID associated with a CFFileSecurityRef object [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1426526-cffilesecuritygetgroup?language=objc
func FileSecurityGetGroup(fileSec FileSecurityRef, group *int) bool {
	rv := C.FileSecurityGetGroup(
		// *typing.RefType
		unsafe.Pointer(fileSec),
		// *typing.PointerType
		(*C.gid_t)(unsafe.Pointer(&group)),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates a writable stream for a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1539652-cfwritestreamcreatewithfile?language=objc
func WriteStreamCreateWithFile(alloc AllocatorRef, fileURL URLRef) WriteStreamRef {
	rv := C.WriteStreamCreateWithFile(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(fileURL),
	)
	// *typing.RefType
	return WriteStreamRef(rv)
}

// Inverts the content of a given character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543744-cfcharactersetinvert?language=objc
func CharacterSetInvert(theSet unsafe.Pointer) {
	C.CharacterSetInvert(
		// *typing.RefType
		unsafe.Pointer(theSet),
	)
}

// Returns the type identifier for the CFError opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1494645-cferrorgettypeid?language=objc
func ErrorGetTypeID() TypeID {
	rv := C.ErrorGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Returns a pointer to a function in a bundle’s executable code using the function name as the search key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537143-cfbundlegetfunctionpointerfornam?language=objc
func BundleGetFunctionPointerForName(bundle BundleRef, functionName StringRef) unsafe.Pointer {
	rv := C.BundleGetFunctionPointerForName(
		// *typing.RefType
		unsafe.Pointer(bundle),
		// *typing.RefType
		unsafe.Pointer(functionName),
	)
	// *typing.VoidPointerType
	return unsafe.Pointer(rv)
}

// Folds a given string into the form specified by optional flags. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542031-cfstringfold?language=objc
func StringFold(theString unsafe.Pointer, theFlags StringCompareFlags, theLocale LocaleRef) {
	C.StringFold(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.AliasType
		// *typing.AliasType
		(C.CFStringCompareFlags)(theFlags),
		// *typing.RefType
		unsafe.Pointer(theLocale),
	)
}

// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543318-cfurlstartaccessingsecurityscope?language=objc
func URLStartAccessingSecurityScopedResource(url URLRef) bool {
	rv := C.URLStartAccessingSecurityScopedResource(
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.PrimitiveType
	return bool(rv)
}

// Creates an “external representation” of a CFString object, that is, a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543169-cfstringcreateexternalrepresenta?language=objc
func StringCreateExternalRepresentation(alloc AllocatorRef, theString StringRef, encoding StringEncoding, lossByte uint8) DataRef {
	rv := C.StringCreateExternalRepresentation(
		// *typing.RefType
		unsafe.Pointer(alloc),
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
		// *typing.PrimitiveType
		C.uint8_t(lossByte),
	)
	// *typing.RefType
	return DataRef(rv)
}

// Opens a connection to a remote socket. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542331-cfsocketconnecttoaddress?language=objc
func SocketConnectToAddress(s SocketRef, address DataRef, timeout TimeInterval) SocketError {
	rv := C.SocketConnectToAddress(
		// *typing.RefType
		unsafe.Pointer(s),
		// *typing.RefType
		unsafe.Pointer(address),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFTimeInterval)(timeout),
	)
	// *typing.AliasType
	return SocketError(rv)
}

// Appends a C string to the character contents of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1543219-cfstringappendcstring?language=objc
func StringAppendCString(theString unsafe.Pointer, cStr string, encoding StringEncoding) {
	cStrVal := C.CString(cStr)
	defer C.free(unsafe.Pointer(cStrVal))
	C.StringAppendCString(
		// *typing.RefType
		unsafe.Pointer(theString),
		// *typing.CStringType
		cStrVal,
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFStringEncoding)(encoding),
	)
}

// Returns the date style used to create the given date formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1396280-cfdateformattergetdatestyle?language=objc
func DateFormatterGetDateStyle(formatter DateFormatterRef) DateFormatterStyle {
	rv := C.DateFormatterGetDateStyle(
		// *typing.RefType
		unsafe.Pointer(formatter),
	)
	// *typing.AliasType
	return DateFormatterStyle(rv)
}

// Returns an array containing the localizations for a bundle or executable at a particular location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537084-cfbundlecopylocalizationsforurl?language=objc
func BundleCopyLocalizationsForURL(url URLRef) ArrayRef {
	rv := C.BundleCopyLocalizationsForURL(
		// *typing.RefType
		unsafe.Pointer(url),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Returns the dictionary containing all the text field values from a dismissed notification dialog. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534500-cfusernotificationgetresponsedic?language=objc
func UserNotificationGetResponseDictionary(userNotification UserNotificationRef) DictionaryRef {
	rv := C.UserNotificationGetResponseDictionary(
		// *typing.RefType
		unsafe.Pointer(userNotification),
	)
	// *typing.RefType
	return DictionaryRef(rv)
}

// Returns the number of values currently in a binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1509300-cfbinaryheapgetcount?language=objc
func BinaryHeapGetCount(heap BinaryHeapRef) Index {
	rv := C.BinaryHeapGetCount(
		// *typing.RefType
		unsafe.Pointer(heap),
	)
	// *typing.AliasType
	return Index(rv)
}

// Returns the type identifier for the CFUserNotification opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1534522-cfusernotificationgettypeid?language=objc
func UserNotificationGetTypeID() TypeID {
	rv := C.UserNotificationGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Assembles an array of URLs specifying all of the resources of the specified type found in a bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1537121-cfbundlecopyresourceurlsoftype?language=objc
func BundleCopyResourceURLsOfType(bundle BundleRef, resourceType StringRef, subDirName StringRef) ArrayRef {
	rv := C.BundleCopyResourceURLsOfType(
		// *typing.RefType
		unsafe.Pointer(bundle),
		// *typing.RefType
		unsafe.Pointer(resourceType),
		// *typing.RefType
		unsafe.Pointer(subDirName),
	)
	// *typing.RefType
	return ArrayRef(rv)
}

// Returns the type identifier of the CFRunLoopSource opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1541979-cfrunloopsourcegettypeid?language=objc
func RunLoopSourceGetTypeID() TypeID {
	rv := C.RunLoopSourceGetTypeID()
	// *typing.AliasType
	return TypeID(rv)
}

// Inserts a string at a specified location in the character buffer of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1542829-cfstringinsert?language=objc
func StringInsert(str unsafe.Pointer, idx Index, insertedStr StringRef) {
	C.StringInsert(
		// *typing.RefType
		unsafe.Pointer(str),
		// *typing.AliasType
		// *typing.PrimitiveType
		(C.CFIndex)(idx),
		// *typing.RefType
		unsafe.Pointer(insertedStr),
	)
}
