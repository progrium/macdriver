// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FragmentedMovie] class.
var FragmentedMovieClass = _FragmentedMovieClass{objc.GetClass("AVFragmentedMovie")}

type _FragmentedMovieClass struct {
	objc.Class
}

// An interface definition for the [FragmentedMovie] class.
type IFragmentedMovie interface {
	IMovie
}

// An object that represents a fragmented movie file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovie?language=objc
type FragmentedMovie struct {
	Movie
}

func FragmentedMovieFrom(ptr unsafe.Pointer) FragmentedMovie {
	return FragmentedMovie{
		Movie: MovieFrom(ptr),
	}
}

func (fc _FragmentedMovieClass) Alloc() FragmentedMovie {
	rv := objc.Call[FragmentedMovie](fc, objc.Sel("alloc"))
	return rv
}

func FragmentedMovie_Alloc() FragmentedMovie {
	return FragmentedMovieClass.Alloc()
}

func (fc _FragmentedMovieClass) New() FragmentedMovie {
	rv := objc.Call[FragmentedMovie](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFragmentedMovie() FragmentedMovie {
	return FragmentedMovieClass.New()
}

func (f_ FragmentedMovie) Init() FragmentedMovie {
	rv := objc.Call[FragmentedMovie](f_, objc.Sel("init"))
	return rv
}

func (f_ FragmentedMovie) InitWithDataOptions(data []byte, options map[string]objc.IObject) FragmentedMovie {
	rv := objc.Call[FragmentedMovie](f_, objc.Sel("initWithData:options:"), data, options)
	return rv
}

// Creates a movie object from a movie file’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1388090-initwithdata?language=objc
func NewFragmentedMovieWithDataOptions(data []byte, options map[string]objc.IObject) FragmentedMovie {
	instance := FragmentedMovieClass.Alloc().InitWithDataOptions(data, options)
	instance.Autorelease()
	return instance
}

func (fc _FragmentedMovieClass) MovieWithDataOptions(data []byte, options map[string]objc.IObject) FragmentedMovie {
	rv := objc.Call[FragmentedMovie](fc, objc.Sel("movieWithData:options:"), data, options)
	return rv
}

// Returns a new movie object from a movie file’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1458261-moviewithdata?language=objc
func FragmentedMovie_MovieWithDataOptions(data []byte, options map[string]objc.IObject) FragmentedMovie {
	return FragmentedMovieClass.MovieWithDataOptions(data, options)
}

func (f_ FragmentedMovie) InitWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) FragmentedMovie {
	rv := objc.Call[FragmentedMovie](f_, objc.Sel("initWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Creates a movie object from a movie header stored in a QuickTime movie file of ISO base media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1387923-initwithurl?language=objc
func NewFragmentedMovieWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) FragmentedMovie {
	instance := FragmentedMovieClass.Alloc().InitWithURLOptions(URL, options)
	instance.Autorelease()
	return instance
}

func (fc _FragmentedMovieClass) MovieWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) FragmentedMovie {
	rv := objc.Call[FragmentedMovie](fc, objc.Sel("movieWithURL:options:"), objc.Ptr(URL), options)
	return rv
}

// Returns a new movie object from a movie header stored in a QuickTime movie file of ISO base media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmovie/1458223-moviewithurl?language=objc
func FragmentedMovie_MovieWithURLOptions(URL foundation.IURL, options map[string]objc.IObject) FragmentedMovie {
	return FragmentedMovieClass.MovieWithURLOptions(URL, options)
}

func (fc _FragmentedMovieClass) AssetWithURL(URL foundation.IURL) FragmentedMovie {
	rv := objc.Call[FragmentedMovie](fc, objc.Sel("assetWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates an asset that models the media at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389943-assetwithurl?language=objc
func FragmentedMovie_AssetWithURL(URL foundation.IURL) FragmentedMovie {
	return FragmentedMovieClass.AssetWithURL(URL)
}
