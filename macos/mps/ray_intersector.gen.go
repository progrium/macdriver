// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RayIntersector] class.
var RayIntersectorClass = _RayIntersectorClass{objc.GetClass("MPSRayIntersector")}

type _RayIntersectorClass struct {
	objc.Class
}

// An interface definition for the [RayIntersector] class.
type IRayIntersector interface {
	IKernel
}

// A kernel that performs intersection tests between rays and geometry. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector?language=objc
type RayIntersector struct {
	Kernel
}

func RayIntersectorFrom(ptr unsafe.Pointer) RayIntersector {
	return RayIntersector{
		Kernel: KernelFrom(ptr),
	}
}

func (rc _RayIntersectorClass) Alloc() RayIntersector {
	rv := objc.Call[RayIntersector](rc, objc.Sel("alloc"))
	return rv
}

func RayIntersector_Alloc() RayIntersector {
	return RayIntersectorClass.Alloc()
}

func (rc _RayIntersectorClass) New() RayIntersector {
	rv := objc.Call[RayIntersector](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRayIntersector() RayIntersector {
	return RayIntersectorClass.New()
}

func (r_ RayIntersector) Init() RayIntersector {
	rv := objc.Call[RayIntersector](r_, objc.Sel("init"))
	return rv
}

func (r_ RayIntersector) InitWithDevice(device metal.PDevice) RayIntersector {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RayIntersector](r_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewRayIntersectorWithDevice(device metal.PDevice) RayIntersector {
	instance := RayIntersectorClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (r_ RayIntersector) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) RayIntersector {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RayIntersector](r_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func RayIntersector_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) RayIntersector {
	instance := RayIntersectorClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
