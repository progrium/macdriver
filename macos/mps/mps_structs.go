package mps

// TODO:

// https://developer.apple.com/documentation/metalperformanceshaders/mpsregion?language=objc
type Region struct{}

// https://developer.apple.com/documentation/metalperformanceshaders/mpsoffset?language=objc
type Offset struct{}

// https://developer.apple.com/documentation/metalperformanceshaders/mpsstatetextureinfo?language=objc
type StateTextureInfo struct{}

// https://developer.apple.com/documentation/metalperformanceshaders/mpsimagekeypointrangeinfo?language=objc
type ImageKeypointRangeInfo struct{}

// https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistograminfo?language=objc
type ImageHistogramInfo struct{}

// https://developer.apple.com/documentation/metalperformanceshaders/mpsscaletransform?language=objc
type ScaleTransform struct{}

// https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopyoffsets?language=objc
type MatrixCopyOffsets struct{}

// https://developer.apple.com/documentation/metalperformanceshaders/mpsdimensionslice?language=objc
type DimensionSlice struct{}

// https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecoordinate?language=objc
type ImageCoordinate struct{}

// https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayoffsets?language=objc
type NDArrayOffsets struct{}

// https://developer.apple.com/documentation/metalperformanceshaders/mpscustomkernelargumentcount?language=objc
type CustomKernelArgumentCount struct {
	broadcastTextureCount   uint64
	destinationTextureCount uint64
	sourceTextureCount      uint64
}

// https://developer.apple.com/documentation/metalperformanceshaders/mpsintegerdivisionparams?language=objc
type IntegerDivisionParams struct {
	addend  uint16
	divisor uint16
	recip   uint16
	shift   uint16
}
