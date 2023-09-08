// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Bundle] class.
var BundleClass = _BundleClass{objc.GetClass("NSBundle")}

type _BundleClass struct {
	objc.Class
}

// An interface definition for the [Bundle] class.
type IBundle interface {
	objc.IObject
	LoadNibNamedOwnerTopLevelObjects(nibName objc.IObject, owner objc.IObject, topLevelObjects []objc.IObject) bool
	PathForResourceOfType(name string, ext string) string
	URLForResourceWithExtensionSubdirectoryLocalization(name string, ext string, subpath string, localizationName string) URL
	LoadAndReturnError(error IError) bool
	LocalizedStringForKeyValueTable(key string, value string, tableName string) string
	PathForResourceOfTypeInDirectoryForLocalization(name string, ext string, subpath string, localizationName string) string
	Unload() bool
	PreflightAndReturnError(error IError) bool
	Load() bool
	URLForImageResource(name objc.IObject) URL
	URLForAuxiliaryExecutable(executableName string) URL
	ObjectForInfoDictionaryKey(key string) objc.Object
	PathForImageResource(name objc.IObject) string
	PathsForResourcesOfTypeInDirectoryForLocalization(ext string, subpath string, localizationName string) []string
	LocalizedAttributedStringForKeyValueTable(key string, value string, tableName string) AttributedString
	ImageForResource(name objc.IObject) objc.Object
	ClassNamed(className string) objc.Class
	URLForResourceWithExtension(name string, ext string) URL
	URLsForResourcesWithExtensionSubdirectory(ext string, subpath string) []URL
	PathForAuxiliaryExecutable(executableName string) string
	PathForSoundResource(name objc.IObject) string
	URLsForResourcesWithExtensionSubdirectoryLocalization(ext string, subpath string, localizationName string) []URL
	URLForResourceWithExtensionSubdirectory(name string, ext string, subpath string) URL
	ContextHelpForKey(key objc.IObject) AttributedString
	DevelopmentLocalization() string
	BuiltInPlugInsPath() string
	SharedSupportPath() string
	BundleIdentifier() string
	PrivateFrameworksPath() string
	SharedSupportURL() URL
	PreferredLocalizations() []string
	AppStoreReceiptURL() URL
	ResourceURL() URL
	PrincipalClass() objc.Class
	ExecutableURL() URL
	LocalizedInfoDictionary() map[string]objc.Object
	BundleURL() URL
	ExecutableArchitectures() []Number
	InfoDictionary() map[string]objc.Object
	SharedFrameworksPath() string
	ExecutablePath() string
	IsLoaded() bool
	SharedFrameworksURL() URL
	BundlePath() string
	BuiltInPlugInsURL() URL
	ResourcePath() string
	Localizations() []string
	PrivateFrameworksURL() URL
}

// A representation of the code and resources stored in a bundle directory on disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle?language=objc
type Bundle struct {
	objc.Object
}

func BundleFrom(ptr unsafe.Pointer) Bundle {
	return Bundle{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BundleClass) BundleWithPath(path string) Bundle {
	rv := objc.Call[Bundle](bc, objc.Sel("bundleWithPath:"), path)
	return rv
}

// Returns an NSBundle object that corresponds to the specified directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1495012-bundlewithpath?language=objc
func Bundle_BundleWithPath(path string) Bundle {
	return BundleClass.BundleWithPath(path)
}

func (b_ Bundle) InitWithURL(url IURL) Bundle {
	rv := objc.Call[Bundle](b_, objc.Sel("initWithURL:"), objc.Ptr(url))
	return rv
}

// Returns an NSBundle object initialized to correspond to the specified file URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1409352-initwithurl?language=objc
func NewBundleWithURL(url IURL) Bundle {
	instance := BundleClass.Alloc().InitWithURL(url)
	instance.Autorelease()
	return instance
}

func (bc _BundleClass) BundleWithURL(url IURL) Bundle {
	rv := objc.Call[Bundle](bc, objc.Sel("bundleWithURL:"), objc.Ptr(url))
	return rv
}

// Returns an NSBundle object that corresponds to the specified file URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1494992-bundlewithurl?language=objc
func Bundle_BundleWithURL(url IURL) Bundle {
	return BundleClass.BundleWithURL(url)
}

func (b_ Bundle) InitWithPath(path string) Bundle {
	rv := objc.Call[Bundle](b_, objc.Sel("initWithPath:"), path)
	return rv
}

// Returns an NSBundle object initialized to correspond to the specified directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1412741-initwithpath?language=objc
func NewBundleWithPath(path string) Bundle {
	instance := BundleClass.Alloc().InitWithPath(path)
	instance.Autorelease()
	return instance
}

func (bc _BundleClass) Alloc() Bundle {
	rv := objc.Call[Bundle](bc, objc.Sel("alloc"))
	return rv
}

func (bc _BundleClass) New() Bundle {
	rv := objc.Call[Bundle](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBundle() Bundle {
	return BundleClass.New()
}

func (b_ Bundle) Init() Bundle {
	rv := objc.Call[Bundle](b_, objc.Sel("init"))
	return rv
}

// Loads a nib from the bundle with the specified file name and owner. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1402909-loadnibnamed?language=objc
func (b_ Bundle) LoadNibNamedOwnerTopLevelObjects(nibName objc.IObject, owner objc.IObject, topLevelObjects []objc.IObject) bool {
	rv := objc.Call[bool](b_, objc.Sel("loadNibNamed:owner:topLevelObjects:"), objc.Ptr(nibName), owner, topLevelObjects)
	return rv
}

// Returns the full pathname for the resource identified by the specified name and file extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1410989-pathforresource?language=objc
func (b_ Bundle) PathForResourceOfType(name string, ext string) string {
	rv := objc.Call[string](b_, objc.Sel("pathForResource:ofType:"), name, ext)
	return rv
}

// Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, within the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1409807-urlsforresourceswithextension?language=objc
func (bc _BundleClass) URLsForResourcesWithExtensionSubdirectoryInBundleWithURL(ext string, subpath string, bundleURL IURL) []URL {
	rv := objc.Call[[]URL](bc, objc.Sel("URLsForResourcesWithExtension:subdirectory:inBundleWithURL:"), ext, subpath, objc.Ptr(bundleURL))
	return rv
}

// Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, within the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1409807-urlsforresourceswithextension?language=objc
func Bundle_URLsForResourcesWithExtensionSubdirectoryInBundleWithURL(ext string, subpath string, bundleURL IURL) []URL {
	return BundleClass.URLsForResourcesWithExtensionSubdirectoryInBundleWithURL(ext, subpath, bundleURL)
}

// Returns locale identifiers for which a bundle would provide localized content, given a specified list of candidates for a user's language preferences. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1409418-preferredlocalizationsfromarray?language=objc
func (bc _BundleClass) PreferredLocalizationsFromArrayForPreferences(localizationsArray []string, preferencesArray []string) []string {
	rv := objc.Call[[]string](bc, objc.Sel("preferredLocalizationsFromArray:forPreferences:"), localizationsArray, preferencesArray)
	return rv
}

// Returns locale identifiers for which a bundle would provide localized content, given a specified list of candidates for a user's language preferences. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1409418-preferredlocalizationsfromarray?language=objc
func Bundle_PreferredLocalizationsFromArrayForPreferences(localizationsArray []string, preferencesArray []string) []string {
	return BundleClass.PreferredLocalizationsFromArrayForPreferences(localizationsArray, preferencesArray)
}

// Returns the NSBundle object with which the specified class is associated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1417717-bundleforclass?language=objc
func (bc _BundleClass) BundleForClass(aClass objc.IClass) Bundle {
	rv := objc.Call[Bundle](bc, objc.Sel("bundleForClass:"), objc.Ptr(aClass))
	return rv
}

// Returns the NSBundle object with which the specified class is associated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1417717-bundleforclass?language=objc
func Bundle_BundleForClass(aClass objc.IClass) Bundle {
	return BundleClass.BundleForClass(aClass)
}

// Returns an array containing the pathnames for all bundle resources having the specified extension and residing in the bundle directory at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1415876-pathsforresourcesoftype?language=objc
func (bc _BundleClass) PathsForResourcesOfTypeInDirectory_(ext string, bundlePath string) []string {
	rv := objc.Call[[]string](bc, objc.Sel("pathsForResourcesOfType:inDirectory:"), ext, bundlePath)
	return rv
}

// Returns an array containing the pathnames for all bundle resources having the specified extension and residing in the bundle directory at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1415876-pathsforresourcesoftype?language=objc
func Bundle_PathsForResourcesOfTypeInDirectory_(ext string, bundlePath string) []string {
	return BundleClass.PathsForResourcesOfTypeInDirectory_(ext, bundlePath)
}

// Returns the file URL for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1417378-urlforresource?language=objc
func (b_ Bundle) URLForResourceWithExtensionSubdirectoryLocalization(name string, ext string, subpath string, localizationName string) URL {
	rv := objc.Call[URL](b_, objc.Sel("URLForResource:withExtension:subdirectory:localization:"), name, ext, subpath, localizationName)
	return rv
}

// Loads the bundle’s executable code and returns any errors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1411819-loadandreturnerror?language=objc
func (b_ Bundle) LoadAndReturnError(error IError) bool {
	rv := objc.Call[bool](b_, objc.Sel("loadAndReturnError:"), objc.Ptr(error))
	return rv
}

// Returns a localized version of the string designated by the specified key and residing in the specified table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1417694-localizedstringforkey?language=objc
func (b_ Bundle) LocalizedStringForKeyValueTable(key string, value string, tableName string) string {
	rv := objc.Call[string](b_, objc.Sel("localizedStringForKey:value:table:"), key, value, tableName)
	return rv
}

// Returns the full pathname for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1413471-pathforresource?language=objc
func (b_ Bundle) PathForResourceOfTypeInDirectoryForLocalization(name string, ext string, subpath string, localizationName string) string {
	rv := objc.Call[string](b_, objc.Sel("pathForResource:ofType:inDirectory:forLocalization:"), name, ext, subpath, localizationName)
	return rv
}

// Unloads the code associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1412388-unload?language=objc
func (b_ Bundle) Unload() bool {
	rv := objc.Call[bool](b_, objc.Sel("unload"))
	return rv
}

// Returns the NSBundle instance that has the specified bundle identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1411929-bundlewithidentifier?language=objc
func (bc _BundleClass) BundleWithIdentifier(identifier string) Bundle {
	rv := objc.Call[Bundle](bc, objc.Sel("bundleWithIdentifier:"), identifier)
	return rv
}

// Returns the NSBundle instance that has the specified bundle identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1411929-bundlewithidentifier?language=objc
func Bundle_BundleWithIdentifier(identifier string) Bundle {
	return BundleClass.BundleWithIdentifier(identifier)
}

// Returns a Boolean value indicating whether the bundle’s executable code could be loaded successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1415083-preflightandreturnerror?language=objc
func (b_ Bundle) PreflightAndReturnError(error IError) bool {
	rv := objc.Call[bool](b_, objc.Sel("preflightAndReturnError:"), objc.Ptr(error))
	return rv
}

// Dynamically loads the bundle’s executable code into a running program, if the code has not already been loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1415927-load?language=objc
func (b_ Bundle) Load() bool {
	rv := objc.Call[bool](b_, objc.Sel("load"))
	return rv
}

// Returns the location of the specified image resource as an NSURL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1519886-urlforimageresource?language=objc
func (b_ Bundle) URLForImageResource(name objc.IObject) URL {
	rv := objc.Call[URL](b_, objc.Sel("URLForImageResource:"), objc.Ptr(name))
	return rv
}

// Creates and returns a file URL for the resource with the specified name and extension in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1416361-urlforresource?language=objc
func (bc _BundleClass) URLForResourceWithExtensionSubdirectoryInBundleWithURL(name string, ext string, subpath string, bundleURL IURL) URL {
	rv := objc.Call[URL](bc, objc.Sel("URLForResource:withExtension:subdirectory:inBundleWithURL:"), name, ext, subpath, objc.Ptr(bundleURL))
	return rv
}

// Creates and returns a file URL for the resource with the specified name and extension in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1416361-urlforresource?language=objc
func Bundle_URLForResourceWithExtensionSubdirectoryInBundleWithURL(name string, ext string, subpath string, bundleURL IURL) URL {
	return BundleClass.URLForResourceWithExtensionSubdirectoryInBundleWithURL(name, ext, subpath, bundleURL)
}

// Returns the file URL of the executable with the specified name in the receiver’s bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1411412-urlforauxiliaryexecutable?language=objc
func (b_ Bundle) URLForAuxiliaryExecutable(executableName string) URL {
	rv := objc.Call[URL](b_, objc.Sel("URLForAuxiliaryExecutable:"), executableName)
	return rv
}

// Returns the value associated with the specified key in the receiver's information property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1408696-objectforinfodictionarykey?language=objc
func (b_ Bundle) ObjectForInfoDictionaryKey(key string) objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("objectForInfoDictionaryKey:"), key)
	return rv
}

// Returns the location of the specified image resource file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1519854-pathforimageresource?language=objc
func (b_ Bundle) PathForImageResource(name objc.IObject) string {
	rv := objc.Call[string](b_, objc.Sel("pathForImageResource:"), objc.Ptr(name))
	return rv
}

// Returns an array containing the file for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1416940-pathsforresourcesoftype?language=objc
func (b_ Bundle) PathsForResourcesOfTypeInDirectoryForLocalization(ext string, subpath string, localizationName string) []string {
	rv := objc.Call[[]string](b_, objc.Sel("pathsForResourcesOfType:inDirectory:forLocalization:"), ext, subpath, localizationName)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/3746904-localizedattributedstringforkey?language=objc
func (b_ Bundle) LocalizedAttributedStringForKeyValueTable(key string, value string, tableName string) AttributedString {
	rv := objc.Call[AttributedString](b_, objc.Sel("localizedAttributedStringForKey:value:table:"), key, value, tableName)
	return rv
}

// Returns an NSImage instance associated with the specified name, which can be backed by multiple files representing different resolution versions of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1519901-imageforresource?language=objc
func (b_ Bundle) ImageForResource(name objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("imageForResource:"), objc.Ptr(name))
	return rv
}

// Returns the Class object for the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1407299-classnamed?language=objc
func (b_ Bundle) ClassNamed(className string) objc.Class {
	rv := objc.Call[objc.Class](b_, objc.Sel("classNamed:"), className)
	return rv
}

// Returns the file URL for the resource identified by the specified name and file extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1411540-urlforresource?language=objc
func (b_ Bundle) URLForResourceWithExtension(name string, ext string) URL {
	rv := objc.Call[URL](b_, objc.Sel("URLForResource:withExtension:"), name, ext)
	return rv
}

// Returns an array of file URLs for all resources identified by the specified file extension and located in the specified bundle subdirectory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1407424-urlsforresourceswithextension?language=objc
func (b_ Bundle) URLsForResourcesWithExtensionSubdirectory(ext string, subpath string) []URL {
	rv := objc.Call[[]URL](b_, objc.Sel("URLsForResourcesWithExtension:subdirectory:"), ext, subpath)
	return rv
}

// Returns one or more localizations from the specified list that a bundle object would use to locate resources for the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1417249-preferredlocalizationsfromarray?language=objc
func (bc _BundleClass) PreferredLocalizationsFromArray(localizationsArray []string) []string {
	rv := objc.Call[[]string](bc, objc.Sel("preferredLocalizationsFromArray:"), localizationsArray)
	return rv
}

// Returns one or more localizations from the specified list that a bundle object would use to locate resources for the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1417249-preferredlocalizationsfromarray?language=objc
func Bundle_PreferredLocalizationsFromArray(localizationsArray []string) []string {
	return BundleClass.PreferredLocalizationsFromArray(localizationsArray)
}

// Returns the full pathname of the executable with the specified name in the receiver’s bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1415214-pathforauxiliaryexecutable?language=objc
func (b_ Bundle) PathForAuxiliaryExecutable(executableName string) string {
	rv := objc.Call[string](b_, objc.Sel("pathForAuxiliaryExecutable:"), executableName)
	return rv
}

// Returns the location of the specified sound resource file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1477280-pathforsoundresource?language=objc
func (b_ Bundle) PathForSoundResource(name objc.IObject) string {
	rv := objc.Call[string](b_, objc.Sel("pathForSoundResource:"), objc.Ptr(name))
	return rv
}

// Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1414688-urlsforresourceswithextension?language=objc
func (b_ Bundle) URLsForResourcesWithExtensionSubdirectoryLocalization(ext string, subpath string, localizationName string) []URL {
	rv := objc.Call[[]URL](b_, objc.Sel("URLsForResourcesWithExtension:subdirectory:localization:"), ext, subpath, localizationName)
	return rv
}

// Returns the file URL for the resource file identified by the specified name and extension and residing in a given bundle directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1416712-urlforresource?language=objc
func (b_ Bundle) URLForResourceWithExtensionSubdirectory(name string, ext string, subpath string) URL {
	rv := objc.Call[URL](b_, objc.Sel("URLForResource:withExtension:subdirectory:"), name, ext, subpath)
	return rv
}

// Returns the full pathname for the resource file identified by the specified name and extension and residing in a given bundle directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1409523-pathforresource?language=objc
func (bc _BundleClass) PathForResourceOfTypeInDirectory_(name string, ext string, bundlePath string) string {
	rv := objc.Call[string](bc, objc.Sel("pathForResource:ofType:inDirectory:"), name, ext, bundlePath)
	return rv
}

// Returns the full pathname for the resource file identified by the specified name and extension and residing in a given bundle directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1409523-pathforresource?language=objc
func Bundle_PathForResourceOfTypeInDirectory_(name string, ext string, bundlePath string) string {
	return BundleClass.PathForResourceOfTypeInDirectory_(name, ext, bundlePath)
}

// Returns the context-sensitive help for the specified key from the bundle's help file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1500918-contexthelpforkey?language=objc
func (b_ Bundle) ContextHelpForKey(key objc.IObject) AttributedString {
	rv := objc.Call[AttributedString](b_, objc.Sel("contextHelpForKey:"), objc.Ptr(key))
	return rv
}

// The localization for the development language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1417526-developmentlocalization?language=objc
func (b_ Bundle) DevelopmentLocalization() string {
	rv := objc.Call[string](b_, objc.Sel("developmentLocalization"))
	return rv
}

// Returns an array of all of the application’s bundles that represent frameworks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1408056-allframeworks?language=objc
func (bc _BundleClass) AllFrameworks() []Bundle {
	rv := objc.Call[[]Bundle](bc, objc.Sel("allFrameworks"))
	return rv
}

// Returns an array of all of the application’s bundles that represent frameworks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1408056-allframeworks?language=objc
func Bundle_AllFrameworks() []Bundle {
	return BundleClass.AllFrameworks()
}

// The full pathname of the receiver's subdirectory containing plug-ins. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1408900-builtinpluginspath?language=objc
func (b_ Bundle) BuiltInPlugInsPath() string {
	rv := objc.Call[string](b_, objc.Sel("builtInPlugInsPath"))
	return rv
}

// The full pathname of the bundle’s subdirectory containing shared support files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1411609-sharedsupportpath?language=objc
func (b_ Bundle) SharedSupportPath() string {
	rv := objc.Call[string](b_, objc.Sel("sharedSupportPath"))
	return rv
}

// The receiver’s bundle identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1418023-bundleidentifier?language=objc
func (b_ Bundle) BundleIdentifier() string {
	rv := objc.Call[string](b_, objc.Sel("bundleIdentifier"))
	return rv
}

// The full pathname of the bundle’s subdirectory containing private frameworks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1415562-privateframeworkspath?language=objc
func (b_ Bundle) PrivateFrameworksPath() string {
	rv := objc.Call[string](b_, objc.Sel("privateFrameworksPath"))
	return rv
}

// The file URL of the bundle’s subdirectory containing shared support files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1416823-sharedsupporturl?language=objc
func (b_ Bundle) SharedSupportURL() URL {
	rv := objc.Call[URL](b_, objc.Sel("sharedSupportURL"))
	return rv
}

// An ordered list of preferred localizations contained in the bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1413220-preferredlocalizations?language=objc
func (b_ Bundle) PreferredLocalizations() []string {
	rv := objc.Call[[]string](b_, objc.Sel("preferredLocalizations"))
	return rv
}

// The file URL for the bundle’s App Store receipt. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1407276-appstorereceipturl?language=objc
func (b_ Bundle) AppStoreReceiptURL() URL {
	rv := objc.Call[URL](b_, objc.Sel("appStoreReceiptURL"))
	return rv
}

// The file URL of the bundle’s subdirectory containing resource files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1414821-resourceurl?language=objc
func (b_ Bundle) ResourceURL() URL {
	rv := objc.Call[URL](b_, objc.Sel("resourceURL"))
	return rv
}

// The bundle’s principal class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1409048-principalclass?language=objc
func (b_ Bundle) PrincipalClass() objc.Class {
	rv := objc.Call[objc.Class](b_, objc.Sel("principalClass"))
	return rv
}

// The file URL of the receiver's executable file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1410470-executableurl?language=objc
func (b_ Bundle) ExecutableURL() URL {
	rv := objc.Call[URL](b_, objc.Sel("executableURL"))
	return rv
}

// A dictionary with the keys from the bundle’s localized property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1407645-localizedinfodictionary?language=objc
func (b_ Bundle) LocalizedInfoDictionary() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](b_, objc.Sel("localizedInfoDictionary"))
	return rv
}

// The full URL of the receiver’s bundle directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1415654-bundleurl?language=objc
func (b_ Bundle) BundleURL() URL {
	rv := objc.Call[URL](b_, objc.Sel("bundleURL"))
	return rv
}

// An array of numbers indicating the architecture types supported by the bundle’s executable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1415499-executablearchitectures?language=objc
func (b_ Bundle) ExecutableArchitectures() []Number {
	rv := objc.Call[[]Number](b_, objc.Sel("executableArchitectures"))
	return rv
}

// Returns an array of all the application’s non-framework bundles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1413705-allbundles?language=objc
func (bc _BundleClass) AllBundles() []Bundle {
	rv := objc.Call[[]Bundle](bc, objc.Sel("allBundles"))
	return rv
}

// Returns an array of all the application’s non-framework bundles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1413705-allbundles?language=objc
func Bundle_AllBundles() []Bundle {
	return BundleClass.AllBundles()
}

// A dictionary, constructed from the bundle’s Info.plist file, that contains information about the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1413477-infodictionary?language=objc
func (b_ Bundle) InfoDictionary() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](b_, objc.Sel("infoDictionary"))
	return rv
}

// The full pathname of the bundle’s subdirectory containing shared frameworks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1417226-sharedframeworkspath?language=objc
func (b_ Bundle) SharedFrameworksPath() string {
	rv := objc.Call[string](b_, objc.Sel("sharedFrameworksPath"))
	return rv
}

// Returns the bundle object that contains the current executable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1410786-mainbundle?language=objc
func (bc _BundleClass) MainBundle() Bundle {
	rv := objc.Call[Bundle](bc, objc.Sel("mainBundle"))
	return rv
}

// Returns the bundle object that contains the current executable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1410786-mainbundle?language=objc
func Bundle_MainBundle() Bundle {
	return BundleClass.MainBundle()
}

// The full pathname of the receiver's executable file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1409078-executablepath?language=objc
func (b_ Bundle) ExecutablePath() string {
	rv := objc.Call[string](b_, objc.Sel("executablePath"))
	return rv
}

// The load status of a bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1413594-loaded?language=objc
func (b_ Bundle) IsLoaded() bool {
	rv := objc.Call[bool](b_, objc.Sel("isLoaded"))
	return rv
}

// The file URL of the receiver's subdirectory containing shared frameworks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1411774-sharedframeworksurl?language=objc
func (b_ Bundle) SharedFrameworksURL() URL {
	rv := objc.Call[URL](b_, objc.Sel("sharedFrameworksURL"))
	return rv
}

// The full pathname of the receiver’s bundle directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1407973-bundlepath?language=objc
func (b_ Bundle) BundlePath() string {
	rv := objc.Call[string](b_, objc.Sel("bundlePath"))
	return rv
}

// The file URL of the receiver's subdirectory containing plug-ins. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1409603-builtinpluginsurl?language=objc
func (b_ Bundle) BuiltInPlugInsURL() URL {
	rv := objc.Call[URL](b_, objc.Sel("builtInPlugInsURL"))
	return rv
}

// The full pathname of the bundle’s subdirectory containing resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1417723-resourcepath?language=objc
func (b_ Bundle) ResourcePath() string {
	rv := objc.Call[string](b_, objc.Sel("resourcePath"))
	return rv
}

// A list of all the localizations contained in the bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1417415-localizations?language=objc
func (b_ Bundle) Localizations() []string {
	rv := objc.Call[[]string](b_, objc.Sel("localizations"))
	return rv
}

// The file URL of the bundle’s subdirectory containing private frameworks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbundle/1417617-privateframeworksurl?language=objc
func (b_ Bundle) PrivateFrameworksURL() URL {
	rv := objc.Call[URL](b_, objc.Sel("privateFrameworksURL"))
	return rv
}
