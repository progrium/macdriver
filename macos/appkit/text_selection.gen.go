// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextSelection] class.
var TextSelectionClass = _TextSelectionClass{objc.GetClass("NSTextSelection")}

type _TextSelectionClass struct {
	objc.Class
}

// An interface definition for the [TextSelection] class.
type ITextSelection interface {
	objc.IObject
	TextSelectionWithTextRanges(textRanges []ITextRange) TextSelection
	IsLogical() bool
	SetLogical(value bool)
	IsTransient() bool
	Affinity() TextSelectionAffinity
	AnchorPositionOffset() float64
	SetAnchorPositionOffset(value float64)
	SecondarySelectionLocation() TextLocationWrapper
	SetSecondarySelectionLocation(value PTextLocation)
	SetSecondarySelectionLocationObject(valueObject objc.IObject)
	TextRanges() []TextRange
	TypingAttributes() map[foundation.AttributedStringKey]objc.Object
	SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	Granularity() TextSelectionGranularity
}

// A class that represents a single logical selection context that corresponds to an insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection?language=objc
type TextSelection struct {
	objc.Object
}

func TextSelectionFrom(ptr unsafe.Pointer) TextSelection {
	return TextSelection{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextSelection) InitWithRangeAffinityGranularity(range_ ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	rv := objc.Call[TextSelection](t_, objc.Sel("initWithRange:affinity:granularity:"), objc.Ptr(range_), affinity, granularity)
	return rv
}

// Creates a new text selection with the range, selection affinity, and granularity you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801821-initwithrange?language=objc
func TextSelection_InitWithRangeAffinityGranularity(range_ ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	return TextSelectionClass.Alloc().InitWithRangeAffinityGranularity(range_, affinity, granularity)
}

func (t_ TextSelection) InitWithLocationAffinity(location PTextLocation, affinity TextSelectionAffinity) TextSelection {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextSelection](t_, objc.Sel("initWithLocation:affinity:"), po0, affinity)
	return rv
}

// Creates a new text selection with the location and selection affinity you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801820-initwithlocation?language=objc
func TextSelection_InitWithLocationAffinity(location PTextLocation, affinity TextSelectionAffinity) TextSelection {
	return TextSelectionClass.Alloc().InitWithLocationAffinity(location, affinity)
}

func (t_ TextSelection) InitWithRangesAffinityGranularity(textRanges []ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	rv := objc.Call[TextSelection](t_, objc.Sel("initWithRanges:affinity:granularity:"), textRanges, affinity, granularity)
	return rv
}

// Creates a new text selection with the ranges, selection affinity, and granularity you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801822-initwithranges?language=objc
func TextSelection_InitWithRangesAffinityGranularity(textRanges []ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	return TextSelectionClass.Alloc().InitWithRangesAffinityGranularity(textRanges, affinity, granularity)
}

func (tc _TextSelectionClass) Alloc() TextSelection {
	rv := objc.Call[TextSelection](tc, objc.Sel("alloc"))
	return rv
}

func TextSelection_Alloc() TextSelection {
	return TextSelectionClass.Alloc()
}

func (tc _TextSelectionClass) New() TextSelection {
	rv := objc.Call[TextSelection](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextSelection() TextSelection {
	return TextSelectionClass.New()
}

func (t_ TextSelection) Init() TextSelection {
	rv := objc.Call[TextSelection](t_, objc.Sel("init"))
	return rv
}

// Creates a subselection of the current text selection with the ranges you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801826-textselectionwithtextranges?language=objc
func (t_ TextSelection) TextSelectionWithTextRanges(textRanges []ITextRange) TextSelection {
	rv := objc.Call[TextSelection](t_, objc.Sel("textSelectionWithTextRanges:"), textRanges)
	return rv
}

// A Boolean value that indicates whether the framework interprets the selection as logical or visual. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801823-logical?language=objc
func (t_ TextSelection) IsLogical() bool {
	rv := objc.Call[bool](t_, objc.Sel("isLogical"))
	return rv
}

// A Boolean value that indicates whether the framework interprets the selection as logical or visual. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801823-logical?language=objc
func (t_ TextSelection) SetLogical(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setLogical:"), value)
}

// A Boolean value that indicates transient text selection during drag handling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801827-transient?language=objc
func (t_ TextSelection) IsTransient() bool {
	rv := objc.Call[bool](t_, objc.Sel("isTransient"))
	return rv
}

// Returns the selection affinity of the text selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801816-affinity?language=objc
func (t_ TextSelection) Affinity() TextSelectionAffinity {
	rv := objc.Call[TextSelectionAffinity](t_, objc.Sel("affinity"))
	return rv
}

// Represents the anchor position offset from the beginning of a line fragment in the visual order for the initial tap or click location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801817-anchorpositionoffset?language=objc
func (t_ TextSelection) AnchorPositionOffset() float64 {
	rv := objc.Call[float64](t_, objc.Sel("anchorPositionOffset"))
	return rv
}

// Represents the anchor position offset from the beginning of a line fragment in the visual order for the initial tap or click location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801817-anchorpositionoffset?language=objc
func (t_ TextSelection) SetAnchorPositionOffset(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setAnchorPositionOffset:"), value)
}

// Specifies the secondary character location when user taps or clicks at a directional boundary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801824-secondaryselectionlocation?language=objc
func (t_ TextSelection) SecondarySelectionLocation() TextLocationWrapper {
	rv := objc.Call[TextLocationWrapper](t_, objc.Sel("secondarySelectionLocation"))
	return rv
}

// Specifies the secondary character location when user taps or clicks at a directional boundary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801824-secondaryselectionlocation?language=objc
func (t_ TextSelection) SetSecondarySelectionLocation(value PTextLocation) {
	po0 := objc.WrapAsProtocol("NSTextLocation", value)
	objc.Call[objc.Void](t_, objc.Sel("setSecondarySelectionLocation:"), po0)
}

// Specifies the secondary character location when user taps or clicks at a directional boundary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801824-secondaryselectionlocation?language=objc
func (t_ TextSelection) SetSecondarySelectionLocationObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setSecondarySelectionLocation:"), objc.Ptr(valueObject))
}

// Represents an array of noncontiguous logical ranges in the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801825-textranges?language=objc
func (t_ TextSelection) TextRanges() []TextRange {
	rv := objc.Call[[]TextRange](t_, objc.Sel("textRanges"))
	return rv
}

// The template attributes the framework uses for characters that replace the contents of this selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801828-typingattributes?language=objc
func (t_ TextSelection) TypingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.Call[map[foundation.AttributedStringKey]objc.Object](t_, objc.Sel("typingAttributes"))
	return rv
}

// The template attributes the framework uses for characters that replace the contents of this selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801828-typingattributes?language=objc
func (t_ TextSelection) SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setTypingAttributes:"), value)
}

// The granularity of the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextselection/3801818-granularity?language=objc
func (t_ TextSelection) Granularity() TextSelectionGranularity {
	rv := objc.Call[TextSelectionGranularity](t_, objc.Sel("granularity"))
	return rv
}
