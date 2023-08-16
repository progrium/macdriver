// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Dictionary] class.
var DictionaryClass = _DictionaryClass{objc.GetClass("NSDictionary")}

type _DictionaryClass struct {
	objc.Class
}

// An interface definition for the [Dictionary] class.
type IDictionary interface {
	objc.IObject
	ObjectEnumerator() Enumerator
	ObjectForKey(aKey objc.IObject) objc.Object
	FileGroupOwnerAccountID() Number
	FileOwnerAccountID() Number
	ObjectsForKeysNotFoundMarker(keys []objc.IObject, marker objc.IObject) []objc.Object
	FileSystemNumber() int
	FileHFSCreatorCode() uint
	KeysOfEntriesPassingTest(predicate func(key objc.Object, obj objc.Object, stop *bool) bool) Set
	KeysSortedByValueWithOptionsUsingComparator(opts SortOptions, cmptr Comparator) []objc.Object
	ObjectForKeyedSubscript(key objc.IObject) objc.Object
	KeysSortedByValueUsingSelector(comparator objc.Selector) []objc.Object
	FileHFSTypeCode() uint
	GetObjectsAndKeysCount(objects objc.IObject, keys objc.IObject, count uint)
	CountByEnumeratingWithStateObjectsCount(state *FastEnumerationState, buffer objc.IObject, len uint) uint
	KeysSortedByValueUsingComparator(cmptr Comparator) []objc.Object
	FileCreationDate() Date
	FilePosixPermissions() uint
	FileExtensionHidden() bool
	FileType() string
	FileSystemFileNumber() uint
	FileIsAppendOnly() bool
	FileIsImmutable() bool
	FileModificationDate() Date
	FileOwnerAccountName() string
	FileSize() int64
	WriteToURLError(url IURL, error IError) bool
	KeyEnumerator() Enumerator
	DescriptionWithLocale(locale objc.IObject) string
	EnumerateKeysAndObjectsUsingBlock(block func(key objc.Object, obj objc.Object, stop *bool))
	IsEqualToDictionary(otherDictionary Dictionary) bool
	KeysOfEntriesWithOptionsPassingTest(opts EnumerationOptions, predicate func(key objc.Object, obj objc.Object, stop *bool) bool) Set
	FileGroupOwnerAccountName() string
	ValueForKey(key string) objc.Object
	EnumerateKeysAndObjectsWithOptionsUsingBlock(opts EnumerationOptions, block func(key objc.Object, obj objc.Object, stop *bool))
	AllKeysForObject(anObject objc.IObject) []objc.Object
	Description() string
	Count() uint
	AllValues() []objc.Object
	AllKeys() []objc.Object
	DescriptionInStringsFileFormat() string
}

// A static collection of objects associated with unique keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary?language=objc
type Dictionary struct {
	objc.Object
}

func DictionaryFrom(ptr unsafe.Pointer) Dictionary {
	return Dictionary{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DictionaryClass) DictionaryWithObjectsForKeys(objects []objc.IObject, keys []PCopying) Dictionary {
	rv := objc.Call[Dictionary](dc, objc.Sel("dictionaryWithObjects:forKeys:"), objects, keys)
	return rv
}

// Creates a dictionary containing entries constructed from the contents of an array of keys and an array of values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1574183-dictionarywithobjects?language=objc
func Dictionary_DictionaryWithObjectsForKeys(objects []objc.IObject, keys []PCopying) Dictionary {
	return DictionaryClass.DictionaryWithObjectsForKeys(objects, keys)
}

func (d_ Dictionary) InitWithObjectsAndKeys(firstObject objc.IObject) Dictionary {
	rv := objc.Call[Dictionary](d_, objc.Sel("initWithObjectsAndKeys:"), firstObject)
	return rv
}

// Initializes a newly allocated dictionary with entries constructed from the specified set of values and keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1574190-initwithobjectsandkeys?language=objc
func Dictionary_InitWithObjectsAndKeys(firstObject objc.IObject) Dictionary {
	return DictionaryClass.Alloc().InitWithObjectsAndKeys(firstObject)
}

func (d_ Dictionary) InitWithObjectsForKeys(objects []objc.IObject, keys []PCopying) Dictionary {
	rv := objc.Call[Dictionary](d_, objc.Sel("initWithObjects:forKeys:"), objects, keys)
	return rv
}

// Initializes a newly allocated dictionary with key-value pairs constructed from the provided arrays of keys and objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1410010-initwithobjects?language=objc
func Dictionary_InitWithObjectsForKeys(objects []objc.IObject, keys []PCopying) Dictionary {
	return DictionaryClass.Alloc().InitWithObjectsForKeys(objects, keys)
}

func (dc _DictionaryClass) DictionaryWithObjectForKey(object objc.IObject, key PCopying) Dictionary {
	po1 := objc.WrapAsProtocol("NSCopying", key)
	rv := objc.Call[Dictionary](dc, objc.Sel("dictionaryWithObject:forKey:"), objc.Ptr(object), po1)
	return rv
}

// Creates a dictionary containing a given key and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1414965-dictionarywithobject?language=objc
func Dictionary_DictionaryWithObjectForKey(object objc.IObject, key PCopying) Dictionary {
	return DictionaryClass.DictionaryWithObjectForKey(object, key)
}

func (dc _DictionaryClass) Dictionary() Dictionary {
	rv := objc.Call[Dictionary](dc, objc.Sel("dictionary"))
	return rv
}

// Creates an empty dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1574180-dictionary?language=objc
func Dictionary_Dictionary() Dictionary {
	return DictionaryClass.Dictionary()
}

func (dc _DictionaryClass) DictionaryWithObjectsAndKeys(firstObject objc.IObject) Dictionary {
	rv := objc.Call[Dictionary](dc, objc.Sel("dictionaryWithObjectsAndKeys:"), firstObject)
	return rv
}

// Creates a dictionary containing entries constructed from the specified set of values and keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1574181-dictionarywithobjectsandkeys?language=objc
func Dictionary_DictionaryWithObjectsAndKeys(firstObject objc.IObject) Dictionary {
	return DictionaryClass.DictionaryWithObjectsAndKeys(firstObject)
}

func (dc _DictionaryClass) DictionaryWithDictionary(dict Dictionary) Dictionary {
	rv := objc.Call[Dictionary](dc, objc.Sel("dictionaryWithDictionary:"), dict)
	return rv
}

// Creates a dictionary containing the keys and values from another given dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1574191-dictionarywithdictionary?language=objc
func Dictionary_DictionaryWithDictionary(dict Dictionary) Dictionary {
	return DictionaryClass.DictionaryWithDictionary(dict)
}

func (d_ Dictionary) InitWithDictionary(otherDictionary Dictionary) Dictionary {
	rv := objc.Call[Dictionary](d_, objc.Sel("initWithDictionary:"), otherDictionary)
	return rv
}

// Initializes a newly allocated dictionary by placing in it the keys and values contained in another given dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1418434-initwithdictionary?language=objc
func Dictionary_InitWithDictionary(otherDictionary Dictionary) Dictionary {
	return DictionaryClass.Alloc().InitWithDictionary(otherDictionary)
}

func (d_ Dictionary) Init() Dictionary {
	rv := objc.Call[Dictionary](d_, objc.Sel("init"))
	return rv
}

func (dc _DictionaryClass) Alloc() Dictionary {
	rv := objc.Call[Dictionary](dc, objc.Sel("alloc"))
	return rv
}

func Dictionary_Alloc() Dictionary {
	return DictionaryClass.Alloc()
}

func (dc _DictionaryClass) New() Dictionary {
	rv := objc.Call[Dictionary](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDictionary() Dictionary {
	return DictionaryClass.New()
}

// Returns an enumerator object that lets you access each value in the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1409600-objectenumerator?language=objc
func (d_ Dictionary) ObjectEnumerator() Enumerator {
	rv := objc.Call[Enumerator](d_, objc.Sel("objectEnumerator"))
	return rv
}

// Creates a shared key set object for the specified keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1408190-sharedkeysetforkeys?language=objc
func (dc _DictionaryClass) SharedKeySetForKeys(keys []PCopying) objc.Object {
	rv := objc.Call[objc.Object](dc, objc.Sel("sharedKeySetForKeys:"), keys)
	return rv
}

// Creates a shared key set object for the specified keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1408190-sharedkeysetforkeys?language=objc
func Dictionary_SharedKeySetForKeys(keys []PCopying) objc.Object {
	return DictionaryClass.SharedKeySetForKeys(keys)
}

// Returns the value associated with a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1414347-objectforkey?language=objc
func (d_ Dictionary) ObjectForKey(aKey objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("objectForKey:"), objc.Ptr(aKey))
	return rv
}

// Returns file’s group owner account ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1413626-filegroupowneraccountid?language=objc
func (d_ Dictionary) FileGroupOwnerAccountID() Number {
	rv := objc.Call[Number](d_, objc.Sel("fileGroupOwnerAccountID"))
	return rv
}

// Returns the file’s owner account ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1412281-fileowneraccountid?language=objc
func (d_ Dictionary) FileOwnerAccountID() Number {
	rv := objc.Call[Number](d_, objc.Sel("fileOwnerAccountID"))
	return rv
}

// Returns as a static array the set of objects from the dictionary that corresponds to the specified keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1408769-objectsforkeys?language=objc
func (d_ Dictionary) ObjectsForKeysNotFoundMarker(keys []objc.IObject, marker objc.IObject) []objc.Object {
	rv := objc.Call[[]objc.Object](d_, objc.Sel("objectsForKeys:notFoundMarker:"), keys, objc.Ptr(marker))
	return rv
}

// Returns the filesystem number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1415329-filesystemnumber?language=objc
func (d_ Dictionary) FileSystemNumber() int {
	rv := objc.Call[int](d_, objc.Sel("fileSystemNumber"))
	return rv
}

// Returns the file’s HFS creator code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1415065-filehfscreatorcode?language=objc
func (d_ Dictionary) FileHFSCreatorCode() uint {
	rv := objc.Call[uint](d_, objc.Sel("fileHFSCreatorCode"))
	return rv
}

// Returns the set of keys whose corresponding value satisfies a constraint described by a block object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1407186-keysofentriespassingtest?language=objc
func (d_ Dictionary) KeysOfEntriesPassingTest(predicate func(key objc.Object, obj objc.Object, stop *bool) bool) Set {
	rv := objc.Call[Set](d_, objc.Sel("keysOfEntriesPassingTest:"), predicate)
	return rv
}

// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block and a specified set of options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1415717-keyssortedbyvaluewithoptions?language=objc
func (d_ Dictionary) KeysSortedByValueWithOptionsUsingComparator(opts SortOptions, cmptr Comparator) []objc.Object {
	rv := objc.Call[[]objc.Object](d_, objc.Sel("keysSortedByValueWithOptions:usingComparator:"), opts, cmptr)
	return rv
}

// Returns the value associated with a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1415430-objectforkeyedsubscript?language=objc
func (d_ Dictionary) ObjectForKeyedSubscript(key objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("objectForKeyedSubscript:"), objc.Ptr(key))
	return rv
}

// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1412484-keyssortedbyvalueusingselector?language=objc
func (d_ Dictionary) KeysSortedByValueUsingSelector(comparator objc.Selector) []objc.Object {
	rv := objc.Call[[]objc.Object](d_, objc.Sel("keysSortedByValueUsingSelector:"), comparator)
	return rv
}

// Returns file’s HFS type code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1417215-filehfstypecode?language=objc
func (d_ Dictionary) FileHFSTypeCode() uint {
	rv := objc.Call[uint](d_, objc.Sel("fileHFSTypeCode"))
	return rv
}

// Returns by reference C arrays of the keys and values in the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1409973-getobjects?language=objc
func (d_ Dictionary) GetObjectsAndKeysCount(objects objc.IObject, keys objc.IObject, count uint) {
	objc.Call[objc.Void](d_, objc.Sel("getObjects:andKeys:count:"), objc.Ptr(objects), objc.Ptr(keys), count)
}

// Returns by reference a C array of objects over which the sender should iterate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/2865769-countbyenumeratingwithstate?language=objc
func (d_ Dictionary) CountByEnumeratingWithStateObjectsCount(state *FastEnumerationState, buffer objc.IObject, len uint) uint {
	rv := objc.Call[uint](d_, objc.Sel("countByEnumeratingWithState:objects:count:"), state, objc.Ptr(buffer), len)
	return rv
}

// Returns an array of the dictionary’s keys, in the order they would be in if the dictionary were sorted by its values using a given comparator block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1411105-keyssortedbyvalueusingcomparator?language=objc
func (d_ Dictionary) KeysSortedByValueUsingComparator(cmptr Comparator) []objc.Object {
	rv := objc.Call[[]objc.Object](d_, objc.Sel("keysSortedByValueUsingComparator:"), cmptr)
	return rv
}

// Returns the file’s creation date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1417071-filecreationdate?language=objc
func (d_ Dictionary) FileCreationDate() Date {
	rv := objc.Call[Date](d_, objc.Sel("fileCreationDate"))
	return rv
}

// Returns the file’s POSIX permissions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1409446-fileposixpermissions?language=objc
func (d_ Dictionary) FilePosixPermissions() uint {
	rv := objc.Call[uint](d_, objc.Sel("filePosixPermissions"))
	return rv
}

// Returns a Boolean value indicating whether the file hides its extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1413177-fileextensionhidden?language=objc
func (d_ Dictionary) FileExtensionHidden() bool {
	rv := objc.Call[bool](d_, objc.Sel("fileExtensionHidden"))
	return rv
}

// Returns the file type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1416809-filetype?language=objc
func (d_ Dictionary) FileType() string {
	rv := objc.Call[string](d_, objc.Sel("fileType"))
	return rv
}

// Returns the filesystem file number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1408396-filesystemfilenumber?language=objc
func (d_ Dictionary) FileSystemFileNumber() uint {
	rv := objc.Call[uint](d_, objc.Sel("fileSystemFileNumber"))
	return rv
}

// Returns a Boolean value indicating whether the file is append only. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1416083-fileisappendonly?language=objc
func (d_ Dictionary) FileIsAppendOnly() bool {
	rv := objc.Call[bool](d_, objc.Sel("fileIsAppendOnly"))
	return rv
}

// Returns a Boolean value indicating whether the file is immutable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1416500-fileisimmutable?language=objc
func (d_ Dictionary) FileIsImmutable() bool {
	rv := objc.Call[bool](d_, objc.Sel("fileIsImmutable"))
	return rv
}

// Returns file’s modification date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1408811-filemodificationdate?language=objc
func (d_ Dictionary) FileModificationDate() Date {
	rv := objc.Call[Date](d_, objc.Sel("fileModificationDate"))
	return rv
}

// Returns the file’s owner account name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1417533-fileowneraccountname?language=objc
func (d_ Dictionary) FileOwnerAccountName() string {
	rv := objc.Call[string](d_, objc.Sel("fileOwnerAccountName"))
	return rv
}

// Returns the file’s size, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1413465-filesize?language=objc
func (d_ Dictionary) FileSize() int64 {
	rv := objc.Call[int64](d_, objc.Sel("fileSize"))
	return rv
}

// Writes a property list representation of the contents of the dictionary to a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/2879139-writetourl?language=objc
func (d_ Dictionary) WriteToURLError(url IURL, error IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("writeToURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Provides an enumerator to access the keys in the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1411011-keyenumerator?language=objc
func (d_ Dictionary) KeyEnumerator() Enumerator {
	rv := objc.Call[Enumerator](d_, objc.Sel("keyEnumerator"))
	return rv
}

// Returns a string object that represents the contents of the dictionary, formatted as a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1417665-descriptionwithlocale?language=objc
func (d_ Dictionary) DescriptionWithLocale(locale objc.IObject) string {
	rv := objc.Call[string](d_, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

// Applies a given block object to the entries of the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1414570-enumeratekeysandobjectsusingbloc?language=objc
func (d_ Dictionary) EnumerateKeysAndObjectsUsingBlock(block func(key objc.Object, obj objc.Object, stop *bool)) {
	objc.Call[objc.Void](d_, objc.Sel("enumerateKeysAndObjectsUsingBlock:"), block)
}

// Returns a Boolean value that indicates whether the contents of the receiving dictionary are equal to the contents of another given dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1415445-isequaltodictionary?language=objc
func (d_ Dictionary) IsEqualToDictionary(otherDictionary Dictionary) bool {
	rv := objc.Call[bool](d_, objc.Sel("isEqualToDictionary:"), otherDictionary)
	return rv
}

// Returns the set of keys whose corresponding value satisfies a constraint described by a block object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1416706-keysofentrieswithoptions?language=objc
func (d_ Dictionary) KeysOfEntriesWithOptionsPassingTest(opts EnumerationOptions, predicate func(key objc.Object, obj objc.Object, stop *bool) bool) Set {
	rv := objc.Call[Set](d_, objc.Sel("keysOfEntriesWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns the file’s group owner account name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1416788-filegroupowneraccountname?language=objc
func (d_ Dictionary) FileGroupOwnerAccountName() string {
	rv := objc.Call[string](d_, objc.Sel("fileGroupOwnerAccountName"))
	return rv
}

// Returns the value associated with a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1410210-valueforkey?language=objc
func (d_ Dictionary) ValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("valueForKey:"), key)
	return rv
}

// Applies a given block object to the entries of the dictionary, with options specifying how the enumeration is performed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1409739-enumeratekeysandobjectswithoptio?language=objc
func (d_ Dictionary) EnumerateKeysAndObjectsWithOptionsUsingBlock(opts EnumerationOptions, block func(key objc.Object, obj objc.Object, stop *bool)) {
	objc.Call[objc.Void](d_, objc.Sel("enumerateKeysAndObjectsWithOptions:usingBlock:"), opts, block)
}

// Returns a new array containing the keys corresponding to all occurrences of a given object in the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1417147-allkeysforobject?language=objc
func (d_ Dictionary) AllKeysForObject(anObject objc.IObject) []objc.Object {
	rv := objc.Call[[]objc.Object](d_, objc.Sel("allKeysForObject:"), objc.Ptr(anObject))
	return rv
}

// A string that represents the contents of the dictionary, formatted as a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1410799-description?language=objc
func (d_ Dictionary) Description() string {
	rv := objc.Call[string](d_, objc.Sel("description"))
	return rv
}

// The number of entries in the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1409628-count?language=objc
func (d_ Dictionary) Count() uint {
	rv := objc.Call[uint](d_, objc.Sel("count"))
	return rv
}

// A new array containing the dictionary’s values, or an empty array if the dictionary has no entries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1408915-allvalues?language=objc
func (d_ Dictionary) AllValues() []objc.Object {
	rv := objc.Call[[]objc.Object](d_, objc.Sel("allValues"))
	return rv
}

// A new array containing the dictionary’s keys, or an empty array if the dictionary has no entries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1409150-allkeys?language=objc
func (d_ Dictionary) AllKeys() []objc.Object {
	rv := objc.Call[[]objc.Object](d_, objc.Sel("allKeys"))
	return rv
}

// A string that represents the contents of the dictionary, formatted in .strings file format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdictionary/1413282-descriptioninstringsfileformat?language=objc
func (d_ Dictionary) DescriptionInStringsFileFormat() string {
	rv := objc.Call[string](d_, objc.Sel("descriptionInStringsFileFormat"))
	return rv
}
