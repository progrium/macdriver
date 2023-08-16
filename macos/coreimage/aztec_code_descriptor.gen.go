// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AztecCodeDescriptor] class.
var AztecCodeDescriptorClass = _AztecCodeDescriptorClass{objc.GetClass("CIAztecCodeDescriptor")}

type _AztecCodeDescriptorClass struct {
	objc.Class
}

// An interface definition for the [AztecCodeDescriptor] class.
type IAztecCodeDescriptor interface {
	IBarcodeDescriptor
	IsCompact() bool
	DataCodewordCount() int
	ErrorCorrectedPayload() []byte
	LayerCount() int
}

// A concrete subclass of CIBarcodeDescriptor that represents an Aztec code symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodedescriptor?language=objc
type AztecCodeDescriptor struct {
	BarcodeDescriptor
}

func AztecCodeDescriptorFrom(ptr unsafe.Pointer) AztecCodeDescriptor {
	return AztecCodeDescriptor{
		BarcodeDescriptor: BarcodeDescriptorFrom(ptr),
	}
}

func (a_ AztecCodeDescriptor) InitWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload []byte, isCompact bool, layerCount int, dataCodewordCount int) AztecCodeDescriptor {
	rv := objc.Call[AztecCodeDescriptor](a_, objc.Sel("initWithPayload:isCompact:layerCount:dataCodewordCount:"), errorCorrectedPayload, isCompact, layerCount, dataCodewordCount)
	return rv
}

// Initializes a descriptor that can be used as input to the CIBarcodeGenerator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodedescriptor/2875188-initwithpayload?language=objc
func AztecCodeDescriptor_InitWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload []byte, isCompact bool, layerCount int, dataCodewordCount int) AztecCodeDescriptor {
	return AztecCodeDescriptorClass.Alloc().InitWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload, isCompact, layerCount, dataCodewordCount)
}

func (ac _AztecCodeDescriptorClass) DescriptorWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload []byte, isCompact bool, layerCount int, dataCodewordCount int) AztecCodeDescriptor {
	rv := objc.Call[AztecCodeDescriptor](ac, objc.Sel("descriptorWithPayload:isCompact:layerCount:dataCodewordCount:"), errorCorrectedPayload, isCompact, layerCount, dataCodewordCount)
	return rv
}

// Creates an Aztec code descriptor encoding the given payload and parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodedescriptor/2875192-descriptorwithpayload?language=objc
func AztecCodeDescriptor_DescriptorWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload []byte, isCompact bool, layerCount int, dataCodewordCount int) AztecCodeDescriptor {
	return AztecCodeDescriptorClass.DescriptorWithPayloadIsCompactLayerCountDataCodewordCount(errorCorrectedPayload, isCompact, layerCount, dataCodewordCount)
}

func (ac _AztecCodeDescriptorClass) Alloc() AztecCodeDescriptor {
	rv := objc.Call[AztecCodeDescriptor](ac, objc.Sel("alloc"))
	return rv
}

func AztecCodeDescriptor_Alloc() AztecCodeDescriptor {
	return AztecCodeDescriptorClass.Alloc()
}

func (ac _AztecCodeDescriptorClass) New() AztecCodeDescriptor {
	rv := objc.Call[AztecCodeDescriptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAztecCodeDescriptor() AztecCodeDescriptor {
	return AztecCodeDescriptorClass.New()
}

func (a_ AztecCodeDescriptor) Init() AztecCodeDescriptor {
	rv := objc.Call[AztecCodeDescriptor](a_, objc.Sel("init"))
	return rv
}

// A Boolean value telling if the Aztec code is compact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodedescriptor/2875203-iscompact?language=objc
func (a_ AztecCodeDescriptor) IsCompact() bool {
	rv := objc.Call[bool](a_, objc.Sel("isCompact"))
	return rv
}

// The number of data codewords in the Aztec code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodedescriptor/2875166-datacodewordcount?language=objc
func (a_ AztecCodeDescriptor) DataCodewordCount() int {
	rv := objc.Call[int](a_, objc.Sel("dataCodewordCount"))
	return rv
}

// The error-corrected payload containing the data encoded in the Aztec code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodedescriptor/2875187-errorcorrectedpayload?language=objc
func (a_ AztecCodeDescriptor) ErrorCorrectedPayload() []byte {
	rv := objc.Call[[]byte](a_, objc.Sel("errorCorrectedPayload"))
	return rv
}

// The number of layers embedded in the Aztec code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciazteccodedescriptor/2875174-layercount?language=objc
func (a_ AztecCodeDescriptor) LayerCount() int {
	rv := objc.Call[int](a_, objc.Sel("layerCount"))
	return rv
}
