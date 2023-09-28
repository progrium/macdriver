package metal

import "unsafe"

type Coordinate2D = SamplePosition

// todo
type Region struct{}
type Origin struct{}
type TextureSwizzleChannels struct{}
type Size struct {
	Width  uint
	Height uint
	Depth  uint
}
type Viewport struct{}
type ScissorRect struct{}
type VertexAmplificationViewMapping struct{}
type SizeAndAlign struct{}
type SamplePosition struct {
	X float32
	Y float32
}
type ClearColor struct{}
type AccelerationStructureSizes struct{}

type AutoreleasedComputePipelineReflection unsafe.Pointer
type AutoreleasedRenderPipelineReflection unsafe.Pointer
