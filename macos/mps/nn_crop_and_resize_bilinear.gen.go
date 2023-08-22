// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNCropAndResizeBilinear] class.
var NNCropAndResizeBilinearClass = _NNCropAndResizeBilinearClass{objc.GetClass("MPSNNCropAndResizeBilinear")}

type _NNCropAndResizeBilinearClass struct {
	objc.Class
}

// An interface definition for the [NNCropAndResizeBilinear] class.
type INNCropAndResizeBilinear interface {
	ICNNKernel
	ResizeHeight() uint
	ResizeWidth() uint
	Regions() *Region
	NumberOfRegions() uint
}

// A cropping and bilinear resizing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear?language=objc
type NNCropAndResizeBilinear struct {
	CNNKernel
}

func NNCropAndResizeBilinearFrom(ptr unsafe.Pointer) NNCropAndResizeBilinear {
	return NNCropAndResizeBilinear{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (n_ NNCropAndResizeBilinear) InitWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions(device metal.PDevice, resizeWidth uint, resizeHeight uint, numberOfRegions uint, regions *Region) NNCropAndResizeBilinear {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNCropAndResizeBilinear](n_, objc.Sel("initWithDevice:resizeWidth:resizeHeight:numberOfRegions:regions:"), po0, resizeWidth, resizeHeight, numberOfRegions, regions)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013789-initwithdevice?language=objc
func NewNNCropAndResizeBilinearWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions(device metal.PDevice, resizeWidth uint, resizeHeight uint, numberOfRegions uint, regions *Region) NNCropAndResizeBilinear {
	instance := NNCropAndResizeBilinearClass.Alloc().InitWithDeviceResizeWidthResizeHeightNumberOfRegionsRegions(device, resizeWidth, resizeHeight, numberOfRegions, regions)
	instance.Autorelease()
	return instance
}

func (nc _NNCropAndResizeBilinearClass) Alloc() NNCropAndResizeBilinear {
	rv := objc.Call[NNCropAndResizeBilinear](nc, objc.Sel("alloc"))
	return rv
}

func NNCropAndResizeBilinear_Alloc() NNCropAndResizeBilinear {
	return NNCropAndResizeBilinearClass.Alloc()
}

func (nc _NNCropAndResizeBilinearClass) New() NNCropAndResizeBilinear {
	rv := objc.Call[NNCropAndResizeBilinear](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNCropAndResizeBilinear() NNCropAndResizeBilinear {
	return NNCropAndResizeBilinearClass.New()
}

func (n_ NNCropAndResizeBilinear) Init() NNCropAndResizeBilinear {
	rv := objc.Call[NNCropAndResizeBilinear](n_, objc.Sel("init"))
	return rv
}

func (n_ NNCropAndResizeBilinear) InitWithDevice(device metal.PDevice) NNCropAndResizeBilinear {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNCropAndResizeBilinear](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewNNCropAndResizeBilinearWithDevice(device metal.PDevice) NNCropAndResizeBilinear {
	instance := NNCropAndResizeBilinearClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNCropAndResizeBilinear) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNCropAndResizeBilinear {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNCropAndResizeBilinear](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNCropAndResizeBilinear_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNCropAndResizeBilinear {
	instance := NNCropAndResizeBilinearClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013792-resizeheight?language=objc
func (n_ NNCropAndResizeBilinear) ResizeHeight() uint {
	rv := objc.Call[uint](n_, objc.Sel("resizeHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013793-resizewidth?language=objc
func (n_ NNCropAndResizeBilinear) ResizeWidth() uint {
	rv := objc.Call[uint](n_, objc.Sel("resizeWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013791-regions?language=objc
func (n_ NNCropAndResizeBilinear) Regions() *Region {
	rv := objc.Call[*Region](n_, objc.Sel("regions"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncropandresizebilinear/3013790-numberofregions?language=objc
func (n_ NNCropAndResizeBilinear) NumberOfRegions() uint {
	rv := objc.Call[uint](n_, objc.Sel("numberOfRegions"))
	return rv
}
