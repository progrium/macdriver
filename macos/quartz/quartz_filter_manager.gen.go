// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QuartzFilterManager] class.
var QuartzFilterManagerClass = _QuartzFilterManagerClass{objc.GetClass("QuartzFilterManager")}

type _QuartzFilterManagerClass struct {
	objc.Class
}

// An interface definition for the [QuartzFilterManager] class.
type IQuartzFilterManager interface {
	objc.IObject
	SetDelegate(aDelegate objc.IObject)
	SelectFilter(filter IQuartzFilter) bool
	FilterPanel() appkit.Panel
	Delegate() objc.Object
	FilterView() QuartzFilterView
	SelectedFilter() QuartzFilter
	ImportFilter(filterProperties foundation.Dictionary) QuartzFilter
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager?language=objc
type QuartzFilterManager struct {
	objc.Object
}

func QuartzFilterManagerFrom(ptr unsafe.Pointer) QuartzFilterManager {
	return QuartzFilterManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (qc _QuartzFilterManagerClass) Alloc() QuartzFilterManager {
	rv := objc.Call[QuartzFilterManager](qc, objc.Sel("alloc"))
	return rv
}

func QuartzFilterManager_Alloc() QuartzFilterManager {
	return QuartzFilterManagerClass.Alloc()
}

func (qc _QuartzFilterManagerClass) New() QuartzFilterManager {
	rv := objc.Call[QuartzFilterManager](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQuartzFilterManager() QuartzFilterManager {
	return QuartzFilterManagerClass.New()
}

func (q_ QuartzFilterManager) Init() QuartzFilterManager {
	rv := objc.Call[QuartzFilterManager](q_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager/1505098-setdelegate?language=objc
func (q_ QuartzFilterManager) SetDelegate(aDelegate objc.IObject) {
	objc.Call[objc.Void](q_, objc.Sel("setDelegate:"), aDelegate)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager/1503913-selectfilter?language=objc
func (q_ QuartzFilterManager) SelectFilter(filter IQuartzFilter) bool {
	rv := objc.Call[bool](q_, objc.Sel("selectFilter:"), objc.Ptr(filter))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager/1505297-filterpanel?language=objc
func (q_ QuartzFilterManager) FilterPanel() appkit.Panel {
	rv := objc.Call[appkit.Panel](q_, objc.Sel("filterPanel"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager/1503409-delegate?language=objc
func (q_ QuartzFilterManager) Delegate() objc.Object {
	rv := objc.Call[objc.Object](q_, objc.Sel("delegate"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager/1505290-filterview?language=objc
func (q_ QuartzFilterManager) FilterView() QuartzFilterView {
	rv := objc.Call[QuartzFilterView](q_, objc.Sel("filterView"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager/1518336-filtermanager?language=objc
func (qc _QuartzFilterManagerClass) FilterManager() QuartzFilterManager {
	rv := objc.Call[QuartzFilterManager](qc, objc.Sel("filterManager"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager/1518336-filtermanager?language=objc
func QuartzFilterManager_FilterManager() QuartzFilterManager {
	return QuartzFilterManagerClass.FilterManager()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager/1503432-filtersindomains?language=objc
func (qc _QuartzFilterManagerClass) FiltersInDomains(domains []objc.IObject) []objc.Object {
	rv := objc.Call[[]objc.Object](qc, objc.Sel("filtersInDomains:"), domains)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager/1503432-filtersindomains?language=objc
func QuartzFilterManager_FiltersInDomains(domains []objc.IObject) []objc.Object {
	return QuartzFilterManagerClass.FiltersInDomains(domains)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager/1504816-selectedfilter?language=objc
func (q_ QuartzFilterManager) SelectedFilter() QuartzFilter {
	rv := objc.Call[QuartzFilter](q_, objc.Sel("selectedFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/quartzfiltermanager/1503784-importfilter?language=objc
func (q_ QuartzFilterManager) ImportFilter(filterProperties foundation.Dictionary) QuartzFilter {
	rv := objc.Call[QuartzFilter](q_, objc.Sel("importFilter:"), filterProperties)
	return rv
}
