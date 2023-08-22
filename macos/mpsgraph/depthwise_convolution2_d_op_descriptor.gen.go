// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DepthwiseConvolution2DOpDescriptor] class.
var DepthwiseConvolution2DOpDescriptorClass = _DepthwiseConvolution2DOpDescriptorClass{objc.GetClass("MPSGraphDepthwiseConvolution2DOpDescriptor")}

type _DepthwiseConvolution2DOpDescriptorClass struct {
	objc.Class
}

// An interface definition for the [DepthwiseConvolution2DOpDescriptor] class.
type IDepthwiseConvolution2DOpDescriptor interface {
	objc.IObject
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)
	StrideInX() uint
	SetStrideInX(value uint)
	PaddingStyle() PaddingStyle
	SetPaddingStyle(value PaddingStyle)
	PaddingBottom() uint
	SetPaddingBottom(value uint)
	PaddingTop() uint
	SetPaddingTop(value uint)
	StrideInY() uint
	SetStrideInY(value uint)
	WeightsLayout() TensorNamedDataLayout
	SetWeightsLayout(value TensorNamedDataLayout)
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	DataLayout() TensorNamedDataLayout
	SetDataLayout(value TensorNamedDataLayout)
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	PaddingRight() uint
	SetPaddingRight(value uint)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor?language=objc
type DepthwiseConvolution2DOpDescriptor struct {
	objc.Object
}

func DepthwiseConvolution2DOpDescriptorFrom(ptr unsafe.Pointer) DepthwiseConvolution2DOpDescriptor {
	return DepthwiseConvolution2DOpDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DepthwiseConvolution2DOpDescriptorClass) DescriptorWithDataLayoutWeightsLayout(dataLayout TensorNamedDataLayout, weightsLayout TensorNamedDataLayout) DepthwiseConvolution2DOpDescriptor {
	rv := objc.Call[DepthwiseConvolution2DOpDescriptor](dc, objc.Sel("descriptorWithDataLayout:weightsLayout:"), dataLayout, weightsLayout)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667492-descriptorwithdatalayout?language=objc
func DepthwiseConvolution2DOpDescriptor_DescriptorWithDataLayoutWeightsLayout(dataLayout TensorNamedDataLayout, weightsLayout TensorNamedDataLayout) DepthwiseConvolution2DOpDescriptor {
	return DepthwiseConvolution2DOpDescriptorClass.DescriptorWithDataLayoutWeightsLayout(dataLayout, weightsLayout)
}

func (dc _DepthwiseConvolution2DOpDescriptorClass) DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle PaddingStyle, dataLayout TensorNamedDataLayout, weightsLayout TensorNamedDataLayout) DepthwiseConvolution2DOpDescriptor {
	rv := objc.Call[DepthwiseConvolution2DOpDescriptor](dc, objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667493-descriptorwithstrideinx?language=objc
func DepthwiseConvolution2DOpDescriptor_DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle PaddingStyle, dataLayout TensorNamedDataLayout, weightsLayout TensorNamedDataLayout) DepthwiseConvolution2DOpDescriptor {
	return DepthwiseConvolution2DOpDescriptorClass.DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX, strideInY, dilationRateInX, dilationRateInY, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
}

func (dc _DepthwiseConvolution2DOpDescriptorClass) Alloc() DepthwiseConvolution2DOpDescriptor {
	rv := objc.Call[DepthwiseConvolution2DOpDescriptor](dc, objc.Sel("alloc"))
	return rv
}

func DepthwiseConvolution2DOpDescriptor_Alloc() DepthwiseConvolution2DOpDescriptor {
	return DepthwiseConvolution2DOpDescriptorClass.Alloc()
}

func (dc _DepthwiseConvolution2DOpDescriptorClass) New() DepthwiseConvolution2DOpDescriptor {
	rv := objc.Call[DepthwiseConvolution2DOpDescriptor](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDepthwiseConvolution2DOpDescriptor() DepthwiseConvolution2DOpDescriptor {
	return DepthwiseConvolution2DOpDescriptorClass.New()
}

func (d_ DepthwiseConvolution2DOpDescriptor) Init() DepthwiseConvolution2DOpDescriptor {
	rv := objc.Call[DepthwiseConvolution2DOpDescriptor](d_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667501-setexplicitpaddingwithpaddinglef?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Call[objc.Void](d_, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667502-strideinx?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) StrideInX() uint {
	rv := objc.Call[uint](d_, objc.Sel("strideInX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667502-strideinx?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetStrideInX(value uint) {
	objc.Call[objc.Void](d_, objc.Sel("setStrideInX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667499-paddingstyle?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) PaddingStyle() PaddingStyle {
	rv := objc.Call[PaddingStyle](d_, objc.Sel("paddingStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667499-paddingstyle?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetPaddingStyle(value PaddingStyle) {
	objc.Call[objc.Void](d_, objc.Sel("setPaddingStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667496-paddingbottom?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) PaddingBottom() uint {
	rv := objc.Call[uint](d_, objc.Sel("paddingBottom"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667496-paddingbottom?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetPaddingBottom(value uint) {
	objc.Call[objc.Void](d_, objc.Sel("setPaddingBottom:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667500-paddingtop?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) PaddingTop() uint {
	rv := objc.Call[uint](d_, objc.Sel("paddingTop"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667500-paddingtop?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetPaddingTop(value uint) {
	objc.Call[objc.Void](d_, objc.Sel("setPaddingTop:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667503-strideiny?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) StrideInY() uint {
	rv := objc.Call[uint](d_, objc.Sel("strideInY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667503-strideiny?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetStrideInY(value uint) {
	objc.Call[objc.Void](d_, objc.Sel("setStrideInY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667504-weightslayout?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) WeightsLayout() TensorNamedDataLayout {
	rv := objc.Call[TensorNamedDataLayout](d_, objc.Sel("weightsLayout"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667504-weightslayout?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetWeightsLayout(value TensorNamedDataLayout) {
	objc.Call[objc.Void](d_, objc.Sel("setWeightsLayout:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667497-paddingleft?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Call[uint](d_, objc.Sel("paddingLeft"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667497-paddingleft?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Call[objc.Void](d_, objc.Sel("setPaddingLeft:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667494-dilationrateinx?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) DilationRateInX() uint {
	rv := objc.Call[uint](d_, objc.Sel("dilationRateInX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667494-dilationrateinx?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Call[objc.Void](d_, objc.Sel("setDilationRateInX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667491-datalayout?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) DataLayout() TensorNamedDataLayout {
	rv := objc.Call[TensorNamedDataLayout](d_, objc.Sel("dataLayout"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667491-datalayout?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetDataLayout(value TensorNamedDataLayout) {
	objc.Call[objc.Void](d_, objc.Sel("setDataLayout:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667495-dilationrateiny?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) DilationRateInY() uint {
	rv := objc.Call[uint](d_, objc.Sel("dilationRateInY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667495-dilationrateiny?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Call[objc.Void](d_, objc.Sel("setDilationRateInY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667498-paddingright?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) PaddingRight() uint {
	rv := objc.Call[uint](d_, objc.Sel("paddingRight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphdepthwiseconvolution2dopdescriptor/3667498-paddingright?language=objc
func (d_ DepthwiseConvolution2DOpDescriptor) SetPaddingRight(value uint) {
	objc.Call[objc.Void](d_, objc.Sel("setPaddingRight:"), value)
}
