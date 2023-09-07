package mps

import (
	"fmt"
	"testing"

	"github.com/progrium/macdriver/macos/metal"
)

func TestMPSValid(t *testing.T) {}

func TestX(t *testing.T) {
	d := metal.CreateSystemDefaultDevice()
	cq := d.NewCommandQueue()
	cb := cq.CommandBuffer()

	mm := NewMatrixMultiplication()
	m1 := NewMatrix()
	m2 := NewMatrix()
	m3 := NewMatrix()

	mm.EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix(cb, m1, m2, m3)

	//ce.SetComputePipelineState(
	fmt.Println(cb.Status())
	cb.Commit()
	fmt.Println(cb.Status())
	fmt.Println(cb.Status())
	fmt.Println(cb.Status())
	fmt.Println(cb.Status())
}
