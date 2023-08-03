// this file is marked tmp because its not generated, but should be
package coregraphics

// #import <CoreGraphics/CGGeometry.h>
import "C"
import "github.com/progrium/macdriver/objc"

// struct def should be sync with struct in <CoreGraphics/CGGeometry.h> <CoreGraphics/CGAffineTransform.h>

type Float = float64

type AffineTransform struct {
	A  Float
	B  Float
	C  Float
	D  Float
	TX Float
	TY Float
}

type Point struct {
	X Float
	Y Float
}

type Rect struct {
	Origin Point
	Size   Size
}

type Size struct {
	Width  Float
	Height Float
}

var RectNull = objc.ForceCast[C.CGRect, Rect](C.CGRectNull)
var RectInfinite = objc.ForceCast[C.CGRect, Rect](C.CGRectInfinite)
