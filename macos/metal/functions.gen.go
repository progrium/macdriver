// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import "Metal/Metal.h"
// void * CreateSystemDefaultDevice();
import "C"
import (
	"github.com/progrium/macdriver/objc"
)

// Returns the device instance Metal selects as the default. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/1433401-mtlcreatesystemdefaultdevice?language=objc
func CreateSystemDefaultDevice() DeviceWrapper {
	rv := C.CreateSystemDefaultDevice()
	// *typing.ProtocolType
	return DeviceWrapper{objc.ObjectFrom(rv)}
}
