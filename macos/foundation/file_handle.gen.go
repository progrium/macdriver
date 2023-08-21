// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileHandle] class.
var FileHandleClass = _FileHandleClass{objc.GetClass("NSFileHandle")}

type _FileHandleClass struct {
	objc.Class
}

// An interface definition for the [FileHandle] class.
type IFileHandle interface {
	objc.IObject
	GetOffsetError(offsetInFile *int64, error IError) bool
	ReadInBackgroundAndNotify()
	TruncateAtOffsetError(offset int64, error IError) bool
	ReadDataToEndOfFileAndReturnError(error IError) []byte
	WriteDataError(data []byte, error IError) bool
	SeekToEndReturningOffsetError(offsetInFile *int64, error IError) bool
	WaitForDataInBackgroundAndNotifyForModes(modes []RunLoopMode)
	AcceptConnectionInBackgroundAndNotifyForModes(modes []RunLoopMode)
	ReadToEndOfFileInBackgroundAndNotify()
	ReadDataUpToLengthError(length uint, error IError) []byte
	SynchronizeAndReturnError(error IError) bool
	CloseAndReturnError(error IError) bool
	ReadInBackgroundAndNotifyForModes(modes []RunLoopMode)
	SeekToOffsetError(offset int64, error IError) bool
	AvailableData() []byte
	ReadabilityHandler() func(arg0 FileHandle)
	SetReadabilityHandler(value func(arg0 FileHandle))
	WriteabilityHandler() func(arg0 FileHandle)
	SetWriteabilityHandler(value func(arg0 FileHandle))
	FileDescriptor() int
}

// An object-oriented wrapper for a file descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle?language=objc
type FileHandle struct {
	objc.Object
}

func FileHandleFrom(ptr unsafe.Pointer) FileHandle {
	return FileHandle{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FileHandleClass) FileHandleForUpdatingURLError(url IURL, error IError) FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("fileHandleForUpdatingURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1417026-filehandleforupdatingurl?language=objc
func FileHandle_FileHandleForUpdatingURLError(url IURL, error IError) FileHandle {
	return FileHandleClass.FileHandleForUpdatingURLError(url, error)
}

func (fc _FileHandleClass) FileHandleForUpdatingAtPath(path string) FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("fileHandleForUpdatingAtPath:"), path)
	return rv
}

// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1411131-filehandleforupdatingatpath?language=objc
func FileHandle_FileHandleForUpdatingAtPath(path string) FileHandle {
	return FileHandleClass.FileHandleForUpdatingAtPath(path)
}

func (fc _FileHandleClass) FileHandleForWritingToURLError(url IURL, error IError) FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("fileHandleForWritingToURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Returns a file handle initialized for writing to the file, device, or named socket at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1416892-filehandleforwritingtourl?language=objc
func FileHandle_FileHandleForWritingToURLError(url IURL, error IError) FileHandle {
	return FileHandleClass.FileHandleForWritingToURLError(url, error)
}

func (f_ FileHandle) InitWithFileDescriptor(fd int) FileHandle {
	rv := objc.Call[FileHandle](f_, objc.Sel("initWithFileDescriptor:"), fd)
	return rv
}

// Creates and returns a file handle object associated with the specified file descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1409825-initwithfiledescriptor?language=objc
func NewFileHandleWithFileDescriptor(fd int) FileHandle {
	instance := FileHandleClass.Alloc().InitWithFileDescriptor(fd)
	instance.Autorelease()
	return instance
}

func (fc _FileHandleClass) FileHandleForWritingAtPath(path string) FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("fileHandleForWritingAtPath:"), path)
	return rv
}

// Returns a file handle initialized for writing to the file, device, or named socket at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1414405-filehandleforwritingatpath?language=objc
func FileHandle_FileHandleForWritingAtPath(path string) FileHandle {
	return FileHandleClass.FileHandleForWritingAtPath(path)
}

func (fc _FileHandleClass) FileHandleForReadingAtPath(path string) FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("fileHandleForReadingAtPath:"), path)
	return rv
}

// Returns a file handle initialized for reading the file, device, or named socket at the specified path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1411250-filehandleforreadingatpath?language=objc
func FileHandle_FileHandleForReadingAtPath(path string) FileHandle {
	return FileHandleClass.FileHandleForReadingAtPath(path)
}

func (fc _FileHandleClass) FileHandleForReadingFromURLError(url IURL, error IError) FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("fileHandleForReadingFromURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Returns a file handle initialized for reading the file, device, or named socket at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1408422-filehandleforreadingfromurl?language=objc
func FileHandle_FileHandleForReadingFromURLError(url IURL, error IError) FileHandle {
	return FileHandleClass.FileHandleForReadingFromURLError(url, error)
}

func (fc _FileHandleClass) Alloc() FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("alloc"))
	return rv
}

func FileHandle_Alloc() FileHandle {
	return FileHandleClass.Alloc()
}

func (fc _FileHandleClass) New() FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileHandle() FileHandle {
	return FileHandleClass.New()
}

func (f_ FileHandle) Init() FileHandle {
	rv := objc.Call[FileHandle](f_, objc.Sel("init"))
	return rv
}

// Get the current position of the file pointer within the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/3172526-getoffset?language=objc
func (f_ FileHandle) GetOffsetError(offsetInFile *int64, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("getOffset:error:"), offsetInFile, objc.Ptr(error))
	return rv
}

// Reads from the file or communications channel in the background and posts a notification when finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1417635-readinbackgroundandnotify?language=objc
func (f_ FileHandle) ReadInBackgroundAndNotify() {
	objc.Call[objc.Void](f_, objc.Sel("readInBackgroundAndNotify"))
}

// Truncates or extends the file represented by the file handle to a specified offset within the file and puts the file pointer at that position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/3172532-truncateatoffset?language=objc
func (f_ FileHandle) TruncateAtOffsetError(offset int64, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("truncateAtOffset:error:"), offset, objc.Ptr(error))
	return rv
}

// Reads the available data synchronously up to the end of file or maximum number of bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/3172527-readdatatoendoffileandreturnerro?language=objc
func (f_ FileHandle) ReadDataToEndOfFileAndReturnError(error IError) []byte {
	rv := objc.Call[[]byte](f_, objc.Sel("readDataToEndOfFileAndReturnError:"), objc.Ptr(error))
	return rv
}

// Writes the specified data synchronously to the file handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/3172533-writedata?language=objc
func (f_ FileHandle) WriteDataError(data []byte, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("writeData:error:"), data, objc.Ptr(error))
	return rv
}

// Places the file pointer at the end of the file referenced by the file handle and returns the new file offset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/3172529-seektoendreturningoffset?language=objc
func (f_ FileHandle) SeekToEndReturningOffsetError(offsetInFile *int64, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("seekToEndReturningOffset:error:"), offsetInFile, objc.Ptr(error))
	return rv
}

// Asynchronously checks to see if data is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1414643-waitfordatainbackgroundandnotify?language=objc
func (f_ FileHandle) WaitForDataInBackgroundAndNotifyForModes(modes []RunLoopMode) {
	objc.Call[objc.Void](f_, objc.Sel("waitForDataInBackgroundAndNotifyForModes:"), modes)
}

// Accepts a socket connection (for stream-type sockets only) in the background and creates a file handle for the “near” (client) end of the communications channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1412997-acceptconnectioninbackgroundandn?language=objc
func (f_ FileHandle) AcceptConnectionInBackgroundAndNotifyForModes(modes []RunLoopMode) {
	objc.Call[objc.Void](f_, objc.Sel("acceptConnectionInBackgroundAndNotifyForModes:"), modes)
}

// Reads to the end of file from the file or communications channel in the background and posts a notification when finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1415313-readtoendoffileinbackgroundandno?language=objc
func (f_ FileHandle) ReadToEndOfFileInBackgroundAndNotify() {
	objc.Call[objc.Void](f_, objc.Sel("readToEndOfFileInBackgroundAndNotify"))
}

// Reads data synchronously up to the specified number of bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/3172528-readdatauptolength?language=objc
func (f_ FileHandle) ReadDataUpToLengthError(length uint, error IError) []byte {
	rv := objc.Call[[]byte](f_, objc.Sel("readDataUpToLength:error:"), length, objc.Ptr(error))
	return rv
}

// Causes all in-memory data and attributes of the file represented by the file handle to write to permanent storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/3172531-synchronizeandreturnerror?language=objc
func (f_ FileHandle) SynchronizeAndReturnError(error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("synchronizeAndReturnError:"), objc.Ptr(error))
	return rv
}

// Disallows further access to the represented file or communications channel and signals end of file on communications channels that permit writing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/3172525-closeandreturnerror?language=objc
func (f_ FileHandle) CloseAndReturnError(error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("closeAndReturnError:"), objc.Ptr(error))
	return rv
}

// Reads from the file or communications channel in the background and posts a notification when finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1416294-readinbackgroundandnotifyformode?language=objc
func (f_ FileHandle) ReadInBackgroundAndNotifyForModes(modes []RunLoopMode) {
	objc.Call[objc.Void](f_, objc.Sel("readInBackgroundAndNotifyForModes:"), modes)
}

// Moves the file pointer to the specified offset within the file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/3172530-seektooffset?language=objc
func (f_ FileHandle) SeekToOffsetError(offset int64, error IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("seekToOffset:error:"), offset, objc.Ptr(error))
	return rv
}

// The file handle associated with a null device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1413881-filehandlewithnulldevice?language=objc
func (fc _FileHandleClass) FileHandleWithNullDevice() FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("fileHandleWithNullDevice"))
	return rv
}

// The file handle associated with a null device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1413881-filehandlewithnulldevice?language=objc
func FileHandle_FileHandleWithNullDevice() FileHandle {
	return FileHandleClass.FileHandleWithNullDevice()
}

// The file handle associated with the standard error file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1411001-filehandlewithstandarderror?language=objc
func (fc _FileHandleClass) FileHandleWithStandardError() FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("fileHandleWithStandardError"))
	return rv
}

// The file handle associated with the standard error file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1411001-filehandlewithstandarderror?language=objc
func FileHandle_FileHandleWithStandardError() FileHandle {
	return FileHandleClass.FileHandleWithStandardError()
}

// The file handle associated with the standard input file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1413686-filehandlewithstandardinput?language=objc
func (fc _FileHandleClass) FileHandleWithStandardInput() FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("fileHandleWithStandardInput"))
	return rv
}

// The file handle associated with the standard input file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1413686-filehandlewithstandardinput?language=objc
func FileHandle_FileHandleWithStandardInput() FileHandle {
	return FileHandleClass.FileHandleWithStandardInput()
}

// The data currently available in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1411463-availabledata?language=objc
func (f_ FileHandle) AvailableData() []byte {
	rv := objc.Call[[]byte](f_, objc.Sel("availableData"))
	return rv
}

// The block to use for reading the contents of the file handle asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1412413-readabilityhandler?language=objc
func (f_ FileHandle) ReadabilityHandler() func(arg0 FileHandle) {
	rv := objc.Call[func(arg0 FileHandle)](f_, objc.Sel("readabilityHandler"))
	return rv
}

// The block to use for reading the contents of the file handle asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1412413-readabilityhandler?language=objc
func (f_ FileHandle) SetReadabilityHandler(value func(arg0 FileHandle)) {
	objc.Call[objc.Void](f_, objc.Sel("setReadabilityHandler:"), value)
}

// The file handle associated with the standard output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1416965-filehandlewithstandardoutput?language=objc
func (fc _FileHandleClass) FileHandleWithStandardOutput() FileHandle {
	rv := objc.Call[FileHandle](fc, objc.Sel("fileHandleWithStandardOutput"))
	return rv
}

// The file handle associated with the standard output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1416965-filehandlewithstandardoutput?language=objc
func FileHandle_FileHandleWithStandardOutput() FileHandle {
	return FileHandleClass.FileHandleWithStandardOutput()
}

// The block to use for writing the contents of the file handle asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1415367-writeabilityhandler?language=objc
func (f_ FileHandle) WriteabilityHandler() func(arg0 FileHandle) {
	rv := objc.Call[func(arg0 FileHandle)](f_, objc.Sel("writeabilityHandler"))
	return rv
}

// The block to use for writing the contents of the file handle asynchronously. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1415367-writeabilityhandler?language=objc
func (f_ FileHandle) SetWriteabilityHandler(value func(arg0 FileHandle)) {
	objc.Call[objc.Void](f_, objc.Sel("setWriteabilityHandler:"), value)
}

// The POSIX file descriptor associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandle/1410326-filedescriptor?language=objc
func (f_ FileHandle) FileDescriptor() int {
	rv := objc.Call[int](f_, objc.Sel("fileDescriptor"))
	return rv
}
