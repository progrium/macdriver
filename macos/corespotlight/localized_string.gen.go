// AUTO-GENERATED CODE, DO NOT MODIFY

package corespotlight

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LocalizedString] class.
var LocalizedStringClass = _LocalizedStringClass{objc.GetClass("CSLocalizedString")}

type _LocalizedStringClass struct {
	objc.Class
}

// An interface definition for the [LocalizedString] class.
type ILocalizedString interface {
	foundation.IString
	LocalizedString() string
}

// An object displaying localized text in search results related to your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cslocalizedstring?language=objc
type LocalizedString struct {
	foundation.String
}

func LocalizedStringFrom(ptr unsafe.Pointer) LocalizedString {
	return LocalizedString{
		String: foundation.StringFrom(ptr),
	}
}

func (l_ LocalizedString) InitWithLocalizedStrings(localizedStrings foundation.Dictionary) LocalizedString {
	rv := objc.Call[LocalizedString](l_, objc.Sel("initWithLocalizedStrings:"), localizedStrings)
	return rv
}

// Initializes a CSLocalizedString object with the specified dictionary of localized strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cslocalizedstring/1616403-initwithlocalizedstrings?language=objc
func LocalizedString_InitWithLocalizedStrings(localizedStrings foundation.Dictionary) LocalizedString {
	return LocalizedStringClass.Alloc().InitWithLocalizedStrings(localizedStrings)
}

func (lc _LocalizedStringClass) Alloc() LocalizedString {
	rv := objc.Call[LocalizedString](lc, objc.Sel("alloc"))
	return rv
}

func LocalizedString_Alloc() LocalizedString {
	return LocalizedStringClass.Alloc()
}

func (lc _LocalizedStringClass) New() LocalizedString {
	rv := objc.Call[LocalizedString](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLocalizedString() LocalizedString {
	return LocalizedStringClass.New()
}

func (l_ LocalizedString) Init() LocalizedString {
	rv := objc.Call[LocalizedString](l_, objc.Sel("init"))
	return rv
}

func (l_ LocalizedString) InitWithCharactersNoCopyLengthDeallocator(chars *foundation.Unichar, len uint, deallocator func(arg0 *foundation.Unichar, arg1 uint)) LocalizedString {
	rv := objc.Call[LocalizedString](l_, objc.Sel("initWithCharactersNoCopy:length:deallocator:"), chars, len, deallocator)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/3547180-initwithcharactersnocopy?language=objc
func LocalizedString_InitWithCharactersNoCopyLengthDeallocator(chars *foundation.Unichar, len uint, deallocator func(arg0 *foundation.Unichar, arg1 uint)) LocalizedString {
	return LocalizedStringClass.Alloc().InitWithCharactersNoCopyLengthDeallocator(chars, len, deallocator)
}

func (lc _LocalizedStringClass) StringWithString(string_ string) LocalizedString {
	rv := objc.Call[LocalizedString](lc, objc.Sel("stringWithString:"), string_)
	return rv
}

// Returns a string created by copying the characters from another given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497372-stringwithstring?language=objc
func LocalizedString_StringWithString(string_ string) LocalizedString {
	return LocalizedStringClass.StringWithString(string_)
}

func (l_ LocalizedString) InitWithFormatLocale(format string, locale objc.IObject, args ...any) LocalizedString {
	rv := objc.Call[LocalizedString](l_, objc.Sel("initWithFormat:locale:"), append([]any{format, locale}, args...)...)
	return rv
}

// Returns an NSString object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497317-initwithformat?language=objc
func LocalizedString_InitWithFormatLocale(format string, locale objc.IObject, args ...any) LocalizedString {
	return LocalizedStringClass.Alloc().InitWithFormatLocale(format, locale, args...)
}

func (l_ LocalizedString) InitWithDataEncoding(data []byte, encoding foundation.StringEncoding) LocalizedString {
	rv := objc.Call[LocalizedString](l_, objc.Sel("initWithData:encoding:"), data, encoding)
	return rv
}

// Returns an NSString object initialized by converting given data into UTF-16 code units using a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1416374-initwithdata?language=objc
func LocalizedString_InitWithDataEncoding(data []byte, encoding foundation.StringEncoding) LocalizedString {
	return LocalizedStringClass.Alloc().InitWithDataEncoding(data, encoding)
}

func (lc _LocalizedStringClass) StringWithFormat(format string, args ...any) LocalizedString {
	rv := objc.Call[LocalizedString](lc, objc.Sel("stringWithFormat:"), append([]any{format}, args...)...)
	return rv
}

// Returns a string created by using a given format string as a template into which the remaining argument values are substituted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497275-stringwithformat?language=objc
func LocalizedString_StringWithFormat(format string, args ...any) LocalizedString {
	return LocalizedStringClass.StringWithFormat(format, args...)
}

func (l_ LocalizedString) InitWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len uint, encoding foundation.StringEncoding, deallocator func(arg0 unsafe.Pointer, arg1 uint)) LocalizedString {
	rv := objc.Call[LocalizedString](l_, objc.Sel("initWithBytesNoCopy:length:encoding:deallocator:"), bytes, len, encoding, deallocator)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/3547179-initwithbytesnocopy?language=objc
func LocalizedString_InitWithBytesNoCopyLengthEncodingDeallocator(bytes unsafe.Pointer, len uint, encoding foundation.StringEncoding, deallocator func(arg0 unsafe.Pointer, arg1 uint)) LocalizedString {
	return LocalizedStringClass.Alloc().InitWithBytesNoCopyLengthEncodingDeallocator(bytes, len, encoding, deallocator)
}

func (l_ LocalizedString) InitWithString(aString string) LocalizedString {
	rv := objc.Call[LocalizedString](l_, objc.Sel("initWithString:"), aString)
	return rv
}

// Returns an NSString object initialized by copying the characters from another given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1411293-initwithstring?language=objc
func LocalizedString_InitWithString(aString string) LocalizedString {
	return LocalizedStringClass.Alloc().InitWithString(aString)
}

func (lc _LocalizedStringClass) String() LocalizedString {
	rv := objc.Call[LocalizedString](lc, objc.Sel("string"))
	return rv
}

// Returns an empty string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497312-string?language=objc
func LocalizedString_String() LocalizedString {
	return LocalizedStringClass.String()
}

func (l_ LocalizedString) InitWithBytesLengthEncoding(bytes unsafe.Pointer, len uint, encoding foundation.StringEncoding) LocalizedString {
	rv := objc.Call[LocalizedString](l_, objc.Sel("initWithBytes:length:encoding:"), bytes, len, encoding)
	return rv
}

// Returns an initialized NSString object containing a given number of bytes from a given buffer of bytes interpreted in a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1407339-initwithbytes?language=objc
func LocalizedString_InitWithBytesLengthEncoding(bytes unsafe.Pointer, len uint, encoding foundation.StringEncoding) LocalizedString {
	return LocalizedStringClass.Alloc().InitWithBytesLengthEncoding(bytes, len, encoding)
}

func (lc _LocalizedStringClass) StringWithCharactersLength(characters *foundation.Unichar, length uint) LocalizedString {
	rv := objc.Call[LocalizedString](lc, objc.Sel("stringWithCharacters:length:"), characters, length)
	return rv
}

// Returns a string containing a given number of characters taken from a given C array of UTF-16 code units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497248-stringwithcharacters?language=objc
func LocalizedString_StringWithCharactersLength(characters *foundation.Unichar, length uint) LocalizedString {
	return LocalizedStringClass.StringWithCharactersLength(characters, length)
}

func (l_ LocalizedString) InitWithCharactersLength(characters *foundation.Unichar, length uint) LocalizedString {
	rv := objc.Call[LocalizedString](l_, objc.Sel("initWithCharacters:length:"), characters, length)
	return rv
}

// Returns an initialized NSString object that contains a given number of characters from a given C array of UTF-16 code units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1410997-initwithcharacters?language=objc
func LocalizedString_InitWithCharactersLength(characters *foundation.Unichar, length uint) LocalizedString {
	return LocalizedStringClass.Alloc().InitWithCharactersLength(characters, length)
}

func (l_ LocalizedString) InitWithUTF8String(nullTerminatedCString *uint8) LocalizedString {
	rv := objc.Call[LocalizedString](l_, objc.Sel("initWithUTF8String:"), nullTerminatedCString)
	return rv
}

// Returns an NSString object initialized by copying the characters from a given C array of UTF8-encoded bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1412128-initwithutf8string?language=objc
func LocalizedString_InitWithUTF8String(nullTerminatedCString *uint8) LocalizedString {
	return LocalizedStringClass.Alloc().InitWithUTF8String(nullTerminatedCString)
}

func (lc _LocalizedStringClass) LocalizedStringWithFormat(format string, args ...any) LocalizedString {
	rv := objc.Call[LocalizedString](lc, objc.Sel("localizedStringWithFormat:"), append([]any{format}, args...)...)
	return rv
}

// Returns a string created by using a given format string as a template into which the remaining argument values are substituted according to the current locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497301-localizedstringwithformat?language=objc
func LocalizedString_LocalizedStringWithFormat(format string, args ...any) LocalizedString {
	return LocalizedStringClass.LocalizedStringWithFormat(format, args...)
}

func (lc _LocalizedStringClass) StringWithUTF8String(nullTerminatedCString *uint8) LocalizedString {
	rv := objc.Call[LocalizedString](lc, objc.Sel("stringWithUTF8String:"), nullTerminatedCString)
	return rv
}

// Returns a string created by copying the data from a given C array of UTF8-encoded bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497379-stringwithutf8string?language=objc
func LocalizedString_StringWithUTF8String(nullTerminatedCString *uint8) LocalizedString {
	return LocalizedStringClass.StringWithUTF8String(nullTerminatedCString)
}

func (lc _LocalizedStringClass) StringWithContentsOfURLEncodingError(url foundation.IURL, enc foundation.StringEncoding, error foundation.IError) LocalizedString {
	rv := objc.Call[LocalizedString](lc, objc.Sel("stringWithContentsOfURL:encoding:error:"), objc.Ptr(url), enc, objc.Ptr(error))
	return rv
}

// Returns a string created by reading data from a given URL interpreted using a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497360-stringwithcontentsofurl?language=objc
func LocalizedString_StringWithContentsOfURLEncodingError(url foundation.IURL, enc foundation.StringEncoding, error foundation.IError) LocalizedString {
	return LocalizedStringClass.StringWithContentsOfURLEncodingError(url, enc, error)
}

func (lc _LocalizedStringClass) StringWithContentsOfFileEncodingError(path string, enc foundation.StringEncoding, error foundation.IError) LocalizedString {
	rv := objc.Call[LocalizedString](lc, objc.Sel("stringWithContentsOfFile:encoding:error:"), path, enc, objc.Ptr(error))
	return rv
}

// Returns a string created by reading data from the file at a given path interpreted using a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstring/1497327-stringwithcontentsoffile?language=objc
func LocalizedString_StringWithContentsOfFileEncodingError(path string, enc foundation.StringEncoding, error foundation.IError) LocalizedString {
	return LocalizedStringClass.StringWithContentsOfFileEncodingError(path, enc, error)
}

// Returns the localized string for the current language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cslocalizedstring/1616401-localizedstring?language=objc
func (l_ LocalizedString) LocalizedString() string {
	rv := objc.Call[string](l_, objc.Sel("localizedString"))
	return rv
}
