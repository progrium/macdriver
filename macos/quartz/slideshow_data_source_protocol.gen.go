// Code generated by DarwinKit. DO NOT EDIT.

package quartz

import (
	"github.com/progrium/darwinkit/objc"
)

// The IKSlideshowDataSource protocol describes the methods that an IKSlideshow object uses to access the contents of its data source object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshowdatasource?language=objc
type PSlideshowDataSource interface {
	// optional
	CanExportSlideshowItemAtIndexToApplication(index uint, applicationBundleIdentifier string) bool
	HasCanExportSlideshowItemAtIndexToApplication() bool

	// optional
	NameOfSlideshowItemAtIndex(index uint) string
	HasNameOfSlideshowItemAtIndex() bool

	// optional
	NumberOfSlideshowItems() uint
	HasNumberOfSlideshowItems() bool

	// optional
	SlideshowWillStart()
	HasSlideshowWillStart() bool

	// optional
	SlideshowDidStop()
	HasSlideshowDidStop() bool

	// optional
	SlideshowItemAtIndex(index uint) objc.Object
	HasSlideshowItemAtIndex() bool

	// optional
	SlideshowDidChangeCurrentIndex(newIndex uint)
	HasSlideshowDidChangeCurrentIndex() bool
}

// ensure impl type implements protocol interface
var _ PSlideshowDataSource = (*SlideshowDataSourceObject)(nil)

// A concrete type for the [PSlideshowDataSource] protocol.
type SlideshowDataSourceObject struct {
	objc.Object
}

func (s_ SlideshowDataSourceObject) HasCanExportSlideshowItemAtIndexToApplication() bool {
	return s_.RespondsToSelector(objc.Sel("canExportSlideshowItemAtIndex:toApplication:"))
}

// Reports whether the export button should be enabled for a slideshow item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshowdatasource/1505226-canexportslideshowitematindex?language=objc
func (s_ SlideshowDataSourceObject) CanExportSlideshowItemAtIndexToApplication(index uint, applicationBundleIdentifier string) bool {
	rv := objc.Call[bool](s_, objc.Sel("canExportSlideshowItemAtIndex:toApplication:"), index, applicationBundleIdentifier)
	return rv
}

func (s_ SlideshowDataSourceObject) HasNameOfSlideshowItemAtIndex() bool {
	return s_.RespondsToSelector(objc.Sel("nameOfSlideshowItemAtIndex:"))
}

// Returns the display name for item at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshowdatasource/1503638-nameofslideshowitematindex?language=objc
func (s_ SlideshowDataSourceObject) NameOfSlideshowItemAtIndex(index uint) string {
	rv := objc.Call[string](s_, objc.Sel("nameOfSlideshowItemAtIndex:"), index)
	return rv
}

func (s_ SlideshowDataSourceObject) HasNumberOfSlideshowItems() bool {
	return s_.RespondsToSelector(objc.Sel("numberOfSlideshowItems"))
}

// Returns the number of items in a slideshow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshowdatasource/1503441-numberofslideshowitems?language=objc
func (s_ SlideshowDataSourceObject) NumberOfSlideshowItems() uint {
	rv := objc.Call[uint](s_, objc.Sel("numberOfSlideshowItems"))
	return rv
}

func (s_ SlideshowDataSourceObject) HasSlideshowWillStart() bool {
	return s_.RespondsToSelector(objc.Sel("slideshowWillStart"))
}

// Performs custom tasks when the slideshow is about to start. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshowdatasource/1504337-slideshowwillstart?language=objc
func (s_ SlideshowDataSourceObject) SlideshowWillStart() {
	objc.Call[objc.Void](s_, objc.Sel("slideshowWillStart"))
}

func (s_ SlideshowDataSourceObject) HasSlideshowDidStop() bool {
	return s_.RespondsToSelector(objc.Sel("slideshowDidStop"))
}

// Performs custom tasks when the slideshow stops. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshowdatasource/1504870-slideshowdidstop?language=objc
func (s_ SlideshowDataSourceObject) SlideshowDidStop() {
	objc.Call[objc.Void](s_, objc.Sel("slideshowDidStop"))
}

func (s_ SlideshowDataSourceObject) HasSlideshowItemAtIndex() bool {
	return s_.RespondsToSelector(objc.Sel("slideshowItemAtIndex:"))
}

// Returns the item for a given index [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshowdatasource/1503729-slideshowitematindex?language=objc
func (s_ SlideshowDataSourceObject) SlideshowItemAtIndex(index uint) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("slideshowItemAtIndex:"), index)
	return rv
}

func (s_ SlideshowDataSourceObject) HasSlideshowDidChangeCurrentIndex() bool {
	return s_.RespondsToSelector(objc.Sel("slideshowDidChangeCurrentIndex:"))
}

// Performs custom tasks when the slideshow changes to the item at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikslideshowdatasource/1504272-slideshowdidchangecurrentindex?language=objc
func (s_ SlideshowDataSourceObject) SlideshowDidChangeCurrentIndex(newIndex uint) {
	objc.Call[objc.Void](s_, objc.Sel("slideshowDidChangeCurrentIndex:"), newIndex)
}
