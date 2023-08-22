// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Kernel] class.
var KernelClass = _KernelClass{objc.GetClass("MPSKernel")}

type _KernelClass struct {
	objc.Class
}

// An interface definition for the [Kernel] class.
type IKernel interface {
	objc.IObject
	Options() KernelOptions
	SetOptions(value KernelOptions)
	Device() metal.DeviceWrapper
	Label() string
	SetLabel(value string)
}

// A standard interface for Metal Performance Shaders kernels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel?language=objc
type Kernel struct {
	objc.Object
}

func KernelFrom(ptr unsafe.Pointer) Kernel {
	return Kernel{
		Object: objc.ObjectFrom(ptr),
	}
}

func (k_ Kernel) InitWithDevice(device metal.PDevice) Kernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[Kernel](k_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewKernelWithDevice(device metal.PDevice) Kernel {
	instance := KernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (k_ Kernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) Kernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[Kernel](k_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func Kernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) Kernel {
	instance := KernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (kc _KernelClass) Alloc() Kernel {
	rv := objc.Call[Kernel](kc, objc.Sel("alloc"))
	return rv
}

func Kernel_Alloc() Kernel {
	return KernelClass.Alloc()
}

func (kc _KernelClass) New() Kernel {
	rv := objc.Call[Kernel](kc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewKernel() Kernel {
	return KernelClass.New()
}

func (k_ Kernel) Init() Kernel {
	rv := objc.Call[Kernel](k_, objc.Sel("init"))
	return rv
}

// The set of options used to run the kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618889-options?language=objc
func (k_ Kernel) Options() KernelOptions {
	rv := objc.Call[KernelOptions](k_, objc.Sel("options"))
	return rv
}

// The set of options used to run the kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618889-options?language=objc
func (k_ Kernel) SetOptions(value KernelOptions) {
	objc.Call[objc.Void](k_, objc.Sel("setOptions:"), value)
}

// The device on which the kernel will be used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618824-device?language=objc
func (k_ Kernel) Device() metal.DeviceWrapper {
	rv := objc.Call[metal.DeviceWrapper](k_, objc.Sel("device"))
	return rv
}

// The string that identifies the kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618803-label?language=objc
func (k_ Kernel) Label() string {
	rv := objc.Call[string](k_, objc.Sel("label"))
	return rv
}

// The string that identifies the kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618803-label?language=objc
func (k_ Kernel) SetLabel(value string) {
	objc.Call[objc.Void](k_, objc.Sel("setLabel:"), value)
}
