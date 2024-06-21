// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a palette centroid filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid?language=objc
type PPaletteCentroid interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetPerceptual(value bool)
	HasSetPerceptual() bool

	// optional
	Perceptual() bool
	HasPerceptual() bool

	// optional
	SetPaletteImage(value Image)
	HasSetPaletteImage() bool

	// optional
	PaletteImage() Image
	HasPaletteImage() bool
}

// ensure impl type implements protocol interface
var _ PPaletteCentroid = (*PaletteCentroidObject)(nil)

// A concrete type for the [PPaletteCentroid] protocol.
type PaletteCentroidObject struct {
	objc.Object
}

func (p_ PaletteCentroidObject) HasSetInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228632-inputimage?language=objc
func (p_ PaletteCentroidObject) SetInputImage(value Image) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), value)
}

func (p_ PaletteCentroidObject) HasInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228632-inputimage?language=objc
func (p_ PaletteCentroidObject) InputImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("inputImage"))
	return rv
}

func (p_ PaletteCentroidObject) HasSetPerceptual() bool {
	return p_.RespondsToSelector(objc.Sel("setPerceptual:"))
}

// A Boolean value that specifies whether the filter applies the color palette in a perceptual color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228634-perceptual?language=objc
func (p_ PaletteCentroidObject) SetPerceptual(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setPerceptual:"), value)
}

func (p_ PaletteCentroidObject) HasPerceptual() bool {
	return p_.RespondsToSelector(objc.Sel("perceptual"))
}

// A Boolean value that specifies whether the filter applies the color palette in a perceptual color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228634-perceptual?language=objc
func (p_ PaletteCentroidObject) Perceptual() bool {
	rv := objc.Call[bool](p_, objc.Sel("perceptual"))
	return rv
}

func (p_ PaletteCentroidObject) HasSetPaletteImage() bool {
	return p_.RespondsToSelector(objc.Sel("setPaletteImage:"))
}

// The input color palette, obtained by using a k-means clustering filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228633-paletteimage?language=objc
func (p_ PaletteCentroidObject) SetPaletteImage(value Image) {
	objc.Call[objc.Void](p_, objc.Sel("setPaletteImage:"), value)
}

func (p_ PaletteCentroidObject) HasPaletteImage() bool {
	return p_.RespondsToSelector(objc.Sel("paletteImage"))
}

// The input color palette, obtained by using a k-means clustering filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228633-paletteimage?language=objc
func (p_ PaletteCentroidObject) PaletteImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("paletteImage"))
	return rv
}
