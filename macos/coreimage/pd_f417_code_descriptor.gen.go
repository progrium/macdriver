// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PDF417CodeDescriptor] class.
var PDF417CodeDescriptorClass = _PDF417CodeDescriptorClass{objc.GetClass("CIPDF417CodeDescriptor")}

type _PDF417CodeDescriptorClass struct {
	objc.Class
}

// An interface definition for the [PDF417CodeDescriptor] class.
type IPDF417CodeDescriptor interface {
	IBarcodeDescriptor
	IsCompact() bool
	ColumnCount() int
	RowCount() int
	ErrorCorrectedPayload() []byte
}

// A concrete subclass of CIBarcodeDescriptor that represents a PDF 417 symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417codedescriptor?language=objc
type PDF417CodeDescriptor struct {
	BarcodeDescriptor
}

func PDF417CodeDescriptorFrom(ptr unsafe.Pointer) PDF417CodeDescriptor {
	return PDF417CodeDescriptor{
		BarcodeDescriptor: BarcodeDescriptorFrom(ptr),
	}
}

func (p_ PDF417CodeDescriptor) InitWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload []byte, isCompact bool, rowCount int, columnCount int) PDF417CodeDescriptor {
	rv := objc.Call[PDF417CodeDescriptor](p_, objc.Sel("initWithPayload:isCompact:rowCount:columnCount:"), errorCorrectedPayload, isCompact, rowCount, columnCount)
	return rv
}

// Initializes a descriptor that can be used as input to the CIBarcodeGenerator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417codedescriptor/2875182-initwithpayload?language=objc
func NewPDF417CodeDescriptorWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload []byte, isCompact bool, rowCount int, columnCount int) PDF417CodeDescriptor {
	instance := PDF417CodeDescriptorClass.Alloc().InitWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload, isCompact, rowCount, columnCount)
	instance.Autorelease()
	return instance
}

func (pc _PDF417CodeDescriptorClass) DescriptorWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload []byte, isCompact bool, rowCount int, columnCount int) PDF417CodeDescriptor {
	rv := objc.Call[PDF417CodeDescriptor](pc, objc.Sel("descriptorWithPayload:isCompact:rowCount:columnCount:"), errorCorrectedPayload, isCompact, rowCount, columnCount)
	return rv
}

// Creates a PDF417 code descriptor encoding the given payload and parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417codedescriptor/2875195-descriptorwithpayload?language=objc
func PDF417CodeDescriptor_DescriptorWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload []byte, isCompact bool, rowCount int, columnCount int) PDF417CodeDescriptor {
	return PDF417CodeDescriptorClass.DescriptorWithPayloadIsCompactRowCountColumnCount(errorCorrectedPayload, isCompact, rowCount, columnCount)
}

func (pc _PDF417CodeDescriptorClass) Alloc() PDF417CodeDescriptor {
	rv := objc.Call[PDF417CodeDescriptor](pc, objc.Sel("alloc"))
	return rv
}

func PDF417CodeDescriptor_Alloc() PDF417CodeDescriptor {
	return PDF417CodeDescriptorClass.Alloc()
}

func (pc _PDF417CodeDescriptorClass) New() PDF417CodeDescriptor {
	rv := objc.Call[PDF417CodeDescriptor](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPDF417CodeDescriptor() PDF417CodeDescriptor {
	return PDF417CodeDescriptorClass.New()
}

func (p_ PDF417CodeDescriptor) Init() PDF417CodeDescriptor {
	rv := objc.Call[PDF417CodeDescriptor](p_, objc.Sel("init"))
	return rv
}

// A boolean value telling if the PDF417 code is compact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417codedescriptor/2875194-iscompact?language=objc
func (p_ PDF417CodeDescriptor) IsCompact() bool {
	rv := objc.Call[bool](p_, objc.Sel("isCompact"))
	return rv
}

// The number of columns in the PDF417 code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417codedescriptor/2875171-columncount?language=objc
func (p_ PDF417CodeDescriptor) ColumnCount() int {
	rv := objc.Call[int](p_, objc.Sel("columnCount"))
	return rv
}

// The number of rows in the PDF417 code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417codedescriptor/2875199-rowcount?language=objc
func (p_ PDF417CodeDescriptor) RowCount() int {
	rv := objc.Call[int](p_, objc.Sel("rowCount"))
	return rv
}

// The error-corrected payload containing the data encoded in the PDF417 code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417codedescriptor/2875204-errorcorrectedpayload?language=objc
func (p_ PDF417CodeDescriptor) ErrorCorrectedPayload() []byte {
	rv := objc.Call[[]byte](p_, objc.Sel("errorCorrectedPayload"))
	return rv
}
