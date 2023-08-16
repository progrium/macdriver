// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [InflectionRule] class.
var InflectionRuleClass = _InflectionRuleClass{objc.GetClass("NSInflectionRule")}

type _InflectionRuleClass struct {
	objc.Class
}

// An interface definition for the [InflectionRule] class.
type IInflectionRule interface {
	objc.IObject
}

// A rule that affects how an attributed string performs automatic grammatical agreement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinflectionrule?language=objc
type InflectionRule struct {
	objc.Object
}

func InflectionRuleFrom(ptr unsafe.Pointer) InflectionRule {
	return InflectionRule{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _InflectionRuleClass) Alloc() InflectionRule {
	rv := objc.Call[InflectionRule](ic, objc.Sel("alloc"))
	return rv
}

func InflectionRule_Alloc() InflectionRule {
	return InflectionRuleClass.Alloc()
}

func (ic _InflectionRuleClass) New() InflectionRule {
	rv := objc.Call[InflectionRule](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewInflectionRule() InflectionRule {
	return InflectionRuleClass.New()
}

func (i_ InflectionRule) Init() InflectionRule {
	rv := objc.Call[InflectionRule](i_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value that indicates whether the rule can inflect a given language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinflectionrule/3746911-caninflectlanguage?language=objc
func (ic _InflectionRuleClass) CanInflectLanguage(language string) bool {
	rv := objc.Call[bool](ic, objc.Sel("canInflectLanguage:"), language)
	return rv
}

// Returns a Boolean value that indicates whether the rule can inflect a given language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinflectionrule/3746911-caninflectlanguage?language=objc
func InflectionRule_CanInflectLanguage(language string) bool {
	return InflectionRuleClass.CanInflectLanguage(language)
}

// An inflection rule that performs automatic grammar agreement with default transformations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinflectionrule/3746910-automaticrule?language=objc
func (ic _InflectionRuleClass) AutomaticRule() InflectionRule {
	rv := objc.Call[InflectionRule](ic, objc.Sel("automaticRule"))
	return rv
}

// An inflection rule that performs automatic grammar agreement with default transformations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinflectionrule/3746910-automaticrule?language=objc
func InflectionRule_AutomaticRule() InflectionRule {
	return InflectionRuleClass.AutomaticRule()
}

// A Boolean value that indicates whether the rule can inflect the user’s current preferred localization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinflectionrule/3746912-caninflectpreferredlocalization?language=objc
func (ic _InflectionRuleClass) CanInflectPreferredLocalization() bool {
	rv := objc.Call[bool](ic, objc.Sel("canInflectPreferredLocalization"))
	return rv
}

// A Boolean value that indicates whether the rule can inflect the user’s current preferred localization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinflectionrule/3746912-caninflectpreferredlocalization?language=objc
func InflectionRule_CanInflectPreferredLocalization() bool {
	return InflectionRuleClass.CanInflectPreferredLocalization()
}
