// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserInterfaceCompressionOptions] class.
var UserInterfaceCompressionOptionsClass = _UserInterfaceCompressionOptionsClass{objc.GetClass("NSUserInterfaceCompressionOptions")}

type _UserInterfaceCompressionOptionsClass struct {
	objc.Class
}

// An interface definition for the [UserInterfaceCompressionOptions] class.
type IUserInterfaceCompressionOptions interface {
	objc.IObject
	OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions
	ContainsOptions(options IUserInterfaceCompressionOptions) bool
	OptionsByAddingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions
	IntersectsOptions(options IUserInterfaceCompressionOptions) bool
	IsEmpty() bool
}

// An object that specifies how user interface elements resize themselves when space is constrained. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions?language=objc
type UserInterfaceCompressionOptions struct {
	objc.Object
}

func UserInterfaceCompressionOptionsFrom(ptr unsafe.Pointer) UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptions{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ UserInterfaceCompressionOptions) InitWithCompressionOptions(options foundation.ISet) UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](u_, objc.Sel("initWithCompressionOptions:"), objc.Ptr(options))
	return rv
}

// Creates an option object that represents the union of the supplied options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909987-initwithcompressionoptions?language=objc
func NewUserInterfaceCompressionOptionsWithCompressionOptions(options foundation.ISet) UserInterfaceCompressionOptions {
	instance := UserInterfaceCompressionOptionsClass.Alloc().InitWithCompressionOptions(options)
	instance.Autorelease()
	return instance
}

func (u_ UserInterfaceCompressionOptions) InitWithIdentifier(identifier string) UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](u_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Creates an option object with the given identifier string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909979-initwithidentifier?language=objc
func NewUserInterfaceCompressionOptionsWithIdentifier(identifier string) UserInterfaceCompressionOptions {
	instance := UserInterfaceCompressionOptionsClass.Alloc().InitWithIdentifier(identifier)
	instance.Autorelease()
	return instance
}

func (u_ UserInterfaceCompressionOptions) Init() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](u_, objc.Sel("init"))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) Alloc() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](uc, objc.Sel("alloc"))
	return rv
}

func UserInterfaceCompressionOptions_Alloc() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.Alloc()
}

func (uc _UserInterfaceCompressionOptionsClass) New() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserInterfaceCompressionOptions() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.New()
}

// Creates a new compression options object with the supplied options removed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909969-optionsbyremovingoptions?language=objc
func (u_ UserInterfaceCompressionOptions) OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](u_, objc.Sel("optionsByRemovingOptions:"), objc.Ptr(options))
	return rv
}

// Determines whether the supplied compression options are all present in the current instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909983-containsoptions?language=objc
func (u_ UserInterfaceCompressionOptions) ContainsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.Call[bool](u_, objc.Sel("containsOptions:"), objc.Ptr(options))
	return rv
}

// Creates a new compression options object representing the union with the provided options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909965-optionsbyaddingoptions?language=objc
func (u_ UserInterfaceCompressionOptions) OptionsByAddingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](u_, objc.Sel("optionsByAddingOptions:"), objc.Ptr(options))
	return rv
}

// Determines whether the supplied compression options intersect with the current instance's options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909977-intersectsoptions?language=objc
func (u_ UserInterfaceCompressionOptions) IntersectsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.Call[bool](u_, objc.Sel("intersectsOptions:"), objc.Ptr(options))
	return rv
}

// An option specifying that views should reduce their internal metrics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909976-reducemetricsoption?language=objc
func (uc _UserInterfaceCompressionOptionsClass) ReduceMetricsOption() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](uc, objc.Sel("reduceMetricsOption"))
	return rv
}

// An option specifying that views should reduce their internal metrics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909976-reducemetricsoption?language=objc
func UserInterfaceCompressionOptions_ReduceMetricsOption() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.ReduceMetricsOption()
}

// A Boolean value that denotes whether the option is empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909986-empty?language=objc
func (u_ UserInterfaceCompressionOptions) IsEmpty() bool {
	rv := objc.Call[bool](u_, objc.Sel("isEmpty"))
	return rv
}

// An option specifying that views should no longer maintain equal width constraints. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909966-breakequalwidthsoption?language=objc
func (uc _UserInterfaceCompressionOptionsClass) BreakEqualWidthsOption() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](uc, objc.Sel("breakEqualWidthsOption"))
	return rv
}

// An option specifying that views should no longer maintain equal width constraints. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909966-breakequalwidthsoption?language=objc
func UserInterfaceCompressionOptions_BreakEqualWidthsOption() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.BreakEqualWidthsOption()
}

// An option that represents the union of all standard compression options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909974-standardoptions?language=objc
func (uc _UserInterfaceCompressionOptionsClass) StandardOptions() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](uc, objc.Sel("standardOptions"))
	return rv
}

// An option that represents the union of all standard compression options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909974-standardoptions?language=objc
func UserInterfaceCompressionOptions_StandardOptions() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.StandardOptions()
}

// An option specifying that views should hide their text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909971-hidetextoption?language=objc
func (uc _UserInterfaceCompressionOptionsClass) HideTextOption() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](uc, objc.Sel("hideTextOption"))
	return rv
}

// An option specifying that views should hide their text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909971-hidetextoption?language=objc
func UserInterfaceCompressionOptions_HideTextOption() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.HideTextOption()
}

// An option specifying that views should hide their images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909980-hideimagesoption?language=objc
func (uc _UserInterfaceCompressionOptionsClass) HideImagesOption() UserInterfaceCompressionOptions {
	rv := objc.Call[UserInterfaceCompressionOptions](uc, objc.Sel("hideImagesOption"))
	return rv
}

// An option specifying that views should hide their images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserinterfacecompressionoptions/2909980-hideimagesoption?language=objc
func UserInterfaceCompressionOptions_HideImagesOption() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.HideImagesOption()
}
