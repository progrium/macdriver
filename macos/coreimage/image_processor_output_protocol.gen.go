// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/macos/corevideo"
	"github.com/progrium/darwinkit/macos/iosurface"
	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// A container for writing image data and information produced by a custom image processor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput?language=objc
type PImageProcessorOutput interface {
	// optional
	Surface() iosurface.Ref
	HasSurface() bool

	// optional
	MetalTexture() metal.TextureObject
	HasMetalTexture() bool

	// optional
	MetalCommandBuffer() metal.CommandBufferObject
	HasMetalCommandBuffer() bool

	// optional
	PixelBuffer() corevideo.PixelBufferRef
	HasPixelBuffer() bool

	// optional
	BytesPerRow() uint
	HasBytesPerRow() bool

	// optional
	BaseAddress() unsafe.Pointer
	HasBaseAddress() bool

	// optional
	Region() coregraphics.Rect
	HasRegion() bool

	// optional
	Format() Format
	HasFormat() bool
}

// ensure impl type implements protocol interface
var _ PImageProcessorOutput = (*ImageProcessorOutputObject)(nil)

// A concrete type for the [PImageProcessorOutput] protocol.
type ImageProcessorOutputObject struct {
	objc.Object
}

func (i_ ImageProcessorOutputObject) HasSurface() bool {
	return i_.RespondsToSelector(objc.Sel("surface"))
}

// An IOSurface object to which you can write output pixel data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639627-surface?language=objc
func (i_ ImageProcessorOutputObject) Surface() iosurface.Ref {
	rv := objc.Call[iosurface.Ref](i_, objc.Sel("surface"))
	return rv
}

func (i_ ImageProcessorOutputObject) HasMetalTexture() bool {
	return i_.RespondsToSelector(objc.Sel("metalTexture"))
}

// A Metal texture to which you can write output pixel data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639631-metaltexture?language=objc
func (i_ ImageProcessorOutputObject) MetalTexture() metal.TextureObject {
	rv := objc.Call[metal.TextureObject](i_, objc.Sel("metalTexture"))
	return rv
}

func (i_ ImageProcessorOutputObject) HasMetalCommandBuffer() bool {
	return i_.RespondsToSelector(objc.Sel("metalCommandBuffer"))
}

// A command buffer to use for image processing using Metal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639641-metalcommandbuffer?language=objc
func (i_ ImageProcessorOutputObject) MetalCommandBuffer() metal.CommandBufferObject {
	rv := objc.Call[metal.CommandBufferObject](i_, objc.Sel("metalCommandBuffer"))
	return rv
}

func (i_ ImageProcessorOutputObject) HasPixelBuffer() bool {
	return i_.RespondsToSelector(objc.Sel("pixelBuffer"))
}

// A CoreVideo pixel buffer to which you can write output pixel data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639647-pixelbuffer?language=objc
func (i_ ImageProcessorOutputObject) PixelBuffer() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](i_, objc.Sel("pixelBuffer"))
	return rv
}

func (i_ ImageProcessorOutputObject) HasBytesPerRow() bool {
	return i_.RespondsToSelector(objc.Sel("bytesPerRow"))
}

// The number of bytes per row of pixels for the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639635-bytesperrow?language=objc
func (i_ ImageProcessorOutputObject) BytesPerRow() uint {
	rv := objc.Call[uint](i_, objc.Sel("bytesPerRow"))
	return rv
}

func (i_ ImageProcessorOutputObject) HasBaseAddress() bool {
	return i_.RespondsToSelector(objc.Sel("baseAddress"))
}

// A pointer to CPU memory at which to write output pixel data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639626-baseaddress?language=objc
func (i_ ImageProcessorOutputObject) BaseAddress() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](i_, objc.Sel("baseAddress"))
	return rv
}

func (i_ ImageProcessorOutputObject) HasRegion() bool {
	return i_.RespondsToSelector(objc.Sel("region"))
}

// The rectangular region of the output image that your processor must provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639629-region?language=objc
func (i_ ImageProcessorOutputObject) Region() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](i_, objc.Sel("region"))
	return rv
}

func (i_ ImageProcessorOutputObject) HasFormat() bool {
	return i_.RespondsToSelector(objc.Sel("format"))
}

// The per-pixel data format expected of the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639628-format?language=objc
func (i_ ImageProcessorOutputObject) Format() Format {
	rv := objc.Call[Format](i_, objc.Sel("format"))
	return rv
}
