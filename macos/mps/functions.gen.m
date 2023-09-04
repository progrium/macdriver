// AUTO-GENERATED CODE, DO NOT MODIFY

#import "MetalPerformanceShaders/MetalPerformanceShaders.h"
uint StateBatchResourceSize(MPSStateBatch* batch) {
	return (uint)MPSStateBatchResourceSize(
		// *typing.PointerType
		// -> *typing.AliasType
		(MPSStateBatch*)batch
	);
}
void HintTemporaryMemoryHighWaterMark(id<MTLCommandBuffer> cmdBuf, uint bytes) {
	return (void)MPSHintTemporaryMemoryHighWaterMark(
		// *typing.ProtocolType
		(id<MTLCommandBuffer>)cmdBuf,
		// *typing.PrimitiveType
		(NSUInteger)bytes
	);
}
uint ImageBatchResourceSize(MPSImageBatch* batch) {
	return (uint)MPSImageBatchResourceSize(
		// *typing.PointerType
		// -> *typing.AliasType
		(MPSImageBatch*)batch
	);
}
void SetHeapCacheDuration(id<MTLCommandBuffer> cmdBuf, double seconds) {
	return (void)MPSSetHeapCacheDuration(
		// *typing.ProtocolType
		(id<MTLCommandBuffer>)cmdBuf,
		// *typing.PrimitiveType
		(double)seconds
	);
}
uint StateBatchIncrementReadCount(MPSStateBatch* batch, long amount) {
	return (uint)MPSStateBatchIncrementReadCount(
		// *typing.PointerType
		// -> *typing.AliasType
		(MPSStateBatch*)batch,
		// *typing.PrimitiveType
		(NSInteger)amount
	);
}
void StateBatchSynchronize(MPSStateBatch* batch, id<MTLCommandBuffer> cmdBuf) {
	return (void)MPSStateBatchSynchronize(
		// *typing.PointerType
		// -> *typing.AliasType
		(MPSStateBatch*)batch,
		// *typing.ProtocolType
		(id<MTLCommandBuffer>)cmdBuf
	);
}
int32_t GetCustomKernelBatchedDestinationIndex(MPSCustomKernelArgumentCount c) {
	return (int32_t)MPSGetCustomKernelBatchedDestinationIndex(
		// *typing.StructType
		(MPSCustomKernelArgumentCount)c
	);
}
int32_t GetCustomKernelMaxBatchSize(MPSCustomKernelArgumentCount c, int32_t MPSMaxTextures) {
	return (int32_t)MPSGetCustomKernelMaxBatchSize(
		// *typing.StructType
		(MPSCustomKernelArgumentCount)c,
		// *typing.PrimitiveType
		(int32_t)MPSMaxTextures
	);
}
uint ImageBatchIncrementReadCount(MPSImageBatch* batch, long amount) {
	return (uint)MPSImageBatchIncrementReadCount(
		// *typing.PointerType
		// -> *typing.AliasType
		(MPSImageBatch*)batch,
		// *typing.PrimitiveType
		(NSInteger)amount
	);
}
id<MTLDevice> GetPreferredDevice(MPSDeviceOptions options) {
	return (id<MTLDevice>)MPSGetPreferredDevice(
		// *typing.AliasType
		(MPSDeviceOptions)options
	);
}
int32_t GetCustomKernelBroadcastSourceIndex(MPSCustomKernelArgumentCount c, int32_t sourceIndex, int32_t MPSMaxTextures) {
	return (int32_t)MPSGetCustomKernelBroadcastSourceIndex(
		// *typing.StructType
		(MPSCustomKernelArgumentCount)c,
		// *typing.PrimitiveType
		(int32_t)sourceIndex,
		// *typing.PrimitiveType
		(int32_t)MPSMaxTextures
	);
}
MPSImageType GetImageType(void * image) {
	return (MPSImageType)MPSGetImageType(
		// *typing.ClassType
		(MPSImage*)image
	);
}
MPSIntegerDivisionParams FindIntegerDivisionParams(uint16_t divisor) {
	return (MPSIntegerDivisionParams)MPSFindIntegerDivisionParams(
		// *typing.PrimitiveType
		(uint16_t)divisor
	);
}
int32_t GetCustomKernelBatchedSourceIndex(MPSCustomKernelArgumentCount c, int32_t sourceIndex, int32_t MPSMaxTextures) {
	return (int32_t)MPSGetCustomKernelBatchedSourceIndex(
		// *typing.StructType
		(MPSCustomKernelArgumentCount)c,
		// *typing.PrimitiveType
		(int32_t)sourceIndex,
		// *typing.PrimitiveType
		(int32_t)MPSMaxTextures
	);
}
void ImageBatchSynchronize(MPSImageBatch* batch, id<MTLCommandBuffer> cmdBuf) {
	return (void)MPSImageBatchSynchronize(
		// *typing.PointerType
		// -> *typing.AliasType
		(MPSImageBatch*)batch,
		// *typing.ProtocolType
		(id<MTLCommandBuffer>)cmdBuf
	);
}
bool SupportsMTLDevice(id<MTLDevice> device) {
	return (bool)MPSSupportsMTLDevice(
		// *typing.ProtocolType
		(id<MTLDevice>)device
	);
}
