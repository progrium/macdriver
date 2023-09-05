// AUTO-GENERATED CODE, DO NOT MODIFY

#import "Metal/Metal.h"
MTLRegion RegionMake2D(uint x, uint y, uint width, uint height) {
	return (MTLRegion)MTLRegionMake2D(
		// *typing.PrimitiveType
		(NSUInteger)x,
		// *typing.PrimitiveType
		(NSUInteger)y,
		// *typing.PrimitiveType
		(NSUInteger)width,
		// *typing.PrimitiveType
		(NSUInteger)height
	);
}
void RemoveDeviceObserver(void * observer) {
	return (void)MTLRemoveDeviceObserver(
		// *typing.IDType
		(id)observer
	);
}
MTLRegion RegionMake1D(uint x, uint width) {
	return (MTLRegion)MTLRegionMake1D(
		// *typing.PrimitiveType
		(NSUInteger)x,
		// *typing.PrimitiveType
		(NSUInteger)width
	);
}
MTLCoordinate2D Coordinate2DMake(float x, float y) {
	return (MTLCoordinate2D)MTLCoordinate2DMake(
		// *typing.PrimitiveType
		(float)x,
		// *typing.PrimitiveType
		(float)y
	);
}
MTLRegion RegionMake3D(uint x, uint y, uint z, uint width, uint height, uint depth) {
	return (MTLRegion)MTLRegionMake3D(
		// *typing.PrimitiveType
		(NSUInteger)x,
		// *typing.PrimitiveType
		(NSUInteger)y,
		// *typing.PrimitiveType
		(NSUInteger)z,
		// *typing.PrimitiveType
		(NSUInteger)width,
		// *typing.PrimitiveType
		(NSUInteger)height,
		// *typing.PrimitiveType
		(NSUInteger)depth
	);
}
MTLSamplePosition SamplePositionMake(float x, float y) {
	return (MTLSamplePosition)MTLSamplePositionMake(
		// *typing.PrimitiveType
		(float)x,
		// *typing.PrimitiveType
		(float)y
	);
}
MTLClearColor ClearColorMake(double red, double green, double blue, double alpha) {
	return (MTLClearColor)MTLClearColorMake(
		// *typing.PrimitiveType
		(double)red,
		// *typing.PrimitiveType
		(double)green,
		// *typing.PrimitiveType
		(double)blue,
		// *typing.PrimitiveType
		(double)alpha
	);
}
MTLTextureSwizzleChannels TextureSwizzleChannelsMake(MTLTextureSwizzle r, MTLTextureSwizzle g, MTLTextureSwizzle b, MTLTextureSwizzle a) {
	return (MTLTextureSwizzleChannels)MTLTextureSwizzleChannelsMake(
		// *typing.AliasType
		(MTLTextureSwizzle)r,
		// *typing.AliasType
		(MTLTextureSwizzle)g,
		// *typing.AliasType
		(MTLTextureSwizzle)b,
		// *typing.AliasType
		(MTLTextureSwizzle)a
	);
}
MTLOrigin OriginMake(uint x, uint y, uint z) {
	return (MTLOrigin)MTLOriginMake(
		// *typing.PrimitiveType
		(NSUInteger)x,
		// *typing.PrimitiveType
		(NSUInteger)y,
		// *typing.PrimitiveType
		(NSUInteger)z
	);
}
void * CreateSystemDefaultDevice() {
	return (void *)MTLCreateSystemDefaultDevice(
	);
}
