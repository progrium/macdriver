package metal

import "testing"

func TestMetalValid(t *testing.T) {

}

func TestCreateSystemDefaultDevice(t *testing.T) {
	d := CreateSystemDefaultDevice()
	if d.HasSupportsFamily() {
		t.Log(d.SupportsFamily(GPUFamily(5001)))
		t.Log("MaxThreadgroupMemoryLength", d.MaxThreadgroupMemoryLength())
		t.Log("MaxBufferLength", d.MaxBufferLength())
		t.Log("MaxThreadsPerThreadgroup", d.MaxThreadsPerThreadgroup())
	}
}
