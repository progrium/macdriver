// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ConditionLock] class.
var ConditionLockClass = _ConditionLockClass{objc.GetClass("NSConditionLock")}

type _ConditionLockClass struct {
	objc.Class
}

// An interface definition for the [ConditionLock] class.
type IConditionLock interface {
	objc.IObject
	LockWhenCondition(condition int)
	UnlockWithCondition(condition int)
	TryLockWhenCondition(condition int) bool
	TryLock() bool
	LockBeforeDate(limit IDate) bool
	Name() string
	SetName(value string)
	Condition() int
}

// A lock that can be associated with specific, user-defined conditions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock?language=objc
type ConditionLock struct {
	objc.Object
}

func ConditionLockFrom(ptr unsafe.Pointer) ConditionLock {
	return ConditionLock{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ ConditionLock) InitWithCondition(condition int) ConditionLock {
	rv := objc.Call[ConditionLock](c_, objc.Sel("initWithCondition:"), condition)
	return rv
}

// Initializes a newly allocated NSConditionLock object and sets its condition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/1413018-initwithcondition?language=objc
func ConditionLock_InitWithCondition(condition int) ConditionLock {
	return ConditionLockClass.Alloc().InitWithCondition(condition)
}

func (cc _ConditionLockClass) Alloc() ConditionLock {
	rv := objc.Call[ConditionLock](cc, objc.Sel("alloc"))
	return rv
}

func ConditionLock_Alloc() ConditionLock {
	return ConditionLockClass.Alloc()
}

func (cc _ConditionLockClass) New() ConditionLock {
	rv := objc.Call[ConditionLock](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewConditionLock() ConditionLock {
	return ConditionLockClass.New()
}

func (c_ ConditionLock) Init() ConditionLock {
	rv := objc.Call[ConditionLock](c_, objc.Sel("init"))
	return rv
}

// Attempts to acquire a lock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/1417832-lockwhencondition?language=objc
func (c_ ConditionLock) LockWhenCondition(condition int) {
	objc.Call[objc.Void](c_, objc.Sel("lockWhenCondition:"), condition)
}

// Relinquishes the lock and sets the receiver’s condition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/1412052-unlockwithcondition?language=objc
func (c_ ConditionLock) UnlockWithCondition(condition int) {
	objc.Call[objc.Void](c_, objc.Sel("unlockWithCondition:"), condition)
}

// Attempts to acquire a lock if the receiver’s condition is equal to the specified condition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/1413548-trylockwhencondition?language=objc
func (c_ ConditionLock) TryLockWhenCondition(condition int) bool {
	rv := objc.Call[bool](c_, objc.Sel("tryLockWhenCondition:"), condition)
	return rv
}

// Attempts to acquire a lock without regard to the receiver’s condition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/1409458-trylock?language=objc
func (c_ ConditionLock) TryLock() bool {
	rv := objc.Call[bool](c_, objc.Sel("tryLock"))
	return rv
}

// Attempts to acquire a lock before a specified moment in time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/1418253-lockbeforedate?language=objc
func (c_ ConditionLock) LockBeforeDate(limit IDate) bool {
	rv := objc.Call[bool](c_, objc.Sel("lockBeforeDate:"), objc.Ptr(limit))
	return rv
}

// The name associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/1411294-name?language=objc
func (c_ ConditionLock) Name() string {
	rv := objc.Call[string](c_, objc.Sel("name"))
	return rv
}

// The name associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/1411294-name?language=objc
func (c_ ConditionLock) SetName(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setName:"), value)
}

// The condition associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsconditionlock/1408807-condition?language=objc
func (c_ ConditionLock) Condition() int {
	rv := objc.Call[int](c_, objc.Sel("condition"))
	return rv
}
