// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Movie] class.
var MovieClass = _MovieClass{objc.GetClass("AVMovie")}

type _MovieClass struct {
	objc.Class
}

// An interface definition for the [Movie] class.
type IMovie interface {
	IAsset
	WriteMovieHeaderToURLFileTypeOptionsError(URL foundation.IURL, fileType FileType, options MovieWritingOptions, outError foundation.IError) bool
	IsCompatibleWithFileType(fileType FileType) bool
	MovieHeaderWithFileTypeError(fileType FileType, outError foundation.IError) []byte
	DefaultMediaDataStorage() MediaDataStorage
	ContainsMovieFragments() bool
	Data() []byte
	URL() foundation.URL
	CanContainMovieFragments() bool
}

// An object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie?language=objc
type Movie struct {
	Asset
}

func MovieFrom(ptr unsafe.Pointer) Movie {
	return Movie{
		Asset: AssetFrom(ptr),
	}
}

func (m_ Movie) InitWithDataOptions(data []byte, options map[string]objc.IObject) Movie {
	rv := objc.Call[Movie](m_, objc.Sel("initWithData:options:"), data, options)
	return rv
}

// Creates a movie object from a movie file’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1388090-initwithdata?language=objc
func NewMovieWithDataOptions(data []byte, options map[string]objc.IObject) Movie {
	instance := MovieClass.Alloc().InitWithDataOptions(data, options)
	instance.Autorelease()
	return instance
}

func (mc _MovieClass) MovieWithDataOptions(data []byte, options map[string]objc.IObject) Movie {
	rv := objc.Call[Movie](mc, objc.Sel("movieWithData:options:"), data, options)
	return rv
}

// Returns a new movie object from a movie file’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1458261-moviewithdata?language=objc
func Movie_MovieWithDataOptions(data []byte, options map[string]objc.IObject) Movie {
	return MovieClass.MovieWithDataOptions(data, options)
}

func (m_ Movie) InitWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) Movie {
	rv := objc.Call[Movie](m_, objc.Sel("initWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Creates a movie object from a movie header stored in a QuickTime movie file of ISO base media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1387923-initwithurl?language=objc
func NewMovieWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) Movie {
	instance := MovieClass.Alloc().InitWithURLOptions(URL, options)
	instance.Autorelease()
	return instance
}

func (mc _MovieClass) MovieWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) Movie {
	rv := objc.Call[Movie](mc, objc.Sel("movieWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Returns a new movie object from a movie header stored in a QuickTime movie file of ISO base media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1458223-moviewithurl?language=objc
func Movie_MovieWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) Movie {
	return MovieClass.MovieWithURLOptions(URL, options)
}

func (mc _MovieClass) Alloc() Movie {
	rv := objc.Call[Movie](mc, objc.Sel("alloc"))
	return rv
}

func Movie_Alloc() Movie {
	return MovieClass.Alloc()
}

func (mc _MovieClass) New() Movie {
	rv := objc.Call[Movie](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMovie() Movie {
	return MovieClass.New()
}

func (m_ Movie) Init() Movie {
	rv := objc.Call[Movie](m_, objc.Sel("init"))
	return rv
}

func (mc _MovieClass) AssetWithURL(URL foundation.IURL) Movie {
	rv := objc.Call[Movie](mc, objc.Sel("assetWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates an asset that models the media at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389943-assetwithurl?language=objc
func Movie_AssetWithURL(URL foundation.IURL) Movie {
	return MovieClass.AssetWithURL(URL)
}

// Returns the file types that a movie supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1388690-movietypes?language=objc
func (mc _MovieClass) MovieTypes() []FileType {
	rv := objc.Call[[]FileType](mc, objc.Sel("movieTypes"))
	return rv
}

// Returns the file types that a movie supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1388690-movietypes?language=objc
func Movie_MovieTypes() []FileType {
	return MovieClass.MovieTypes()
}

// Writes the movie header to the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1386682-writemovieheadertourl?language=objc
func (m_ Movie) WriteMovieHeaderToURLFileTypeOptionsError(URL foundation.IURL, fileType FileType, options MovieWritingOptions, outError foundation.IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("writeMovieHeaderToURL:fileType:options:error:"), objc.Ptr(URL), fileType, options, objc.Ptr(outError))
	return rv
}

// Returns a Boolean value that indicates whether the system can create a movie header of the specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1385982-iscompatiblewithfiletype?language=objc
func (m_ Movie) IsCompatibleWithFileType(fileType FileType) bool {
	rv := objc.Call[bool](m_, objc.Sel("isCompatibleWithFileType:"), fileType)
	return rv
}

// Creates a header for a movie for the specified file type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1386686-movieheaderwithfiletype?language=objc
func (m_ Movie) MovieHeaderWithFileTypeError(fileType FileType, outError foundation.IError) []byte {
	rv := objc.Call[[]byte](m_, objc.Sel("movieHeaderWithFileType:error:"), fileType, objc.Ptr(outError))
	return rv
}

// The default storage container for media data added to a movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1388424-defaultmediadatastorage?language=objc
func (m_ Movie) DefaultMediaDataStorage() MediaDataStorage {
	rv := objc.Call[MediaDataStorage](m_, objc.Sel("defaultMediaDataStorage"))
	return rv
}

// A Boolean value that indicates whether at least one movie fragment extends the movie file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1388597-containsmoviefragments?language=objc
func (m_ Movie) ContainsMovieFragments() bool {
	rv := objc.Call[bool](m_, objc.Sel("containsMovieFragments"))
	return rv
}

// A data object that contains the movie file’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1388017-data?language=objc
func (m_ Movie) Data() []byte {
	rv := objc.Call[[]byte](m_, objc.Sel("data"))
	return rv
}

// A URL to a QuickTime or ISO base media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1386990-url?language=objc
func (m_ Movie) URL() foundation.URL {
	rv := objc.Call[foundation.URL](m_, objc.Sel("URL"))
	return rv
}

// A Boolean value that indicates whether fragments can extend the movie file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1387333-cancontainmoviefragments?language=objc
func (m_ Movie) CanContainMovieFragments() bool {
	rv := objc.Call[bool](m_, objc.Sel("canContainMovieFragments"))
	return rv
}
