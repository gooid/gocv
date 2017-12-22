package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type BFMatcher struct {
	*DescriptorMatcher
}

func NewBFMatcher(addr int64) (rcvr *BFMatcher) {
	rcvr = &BFMatcher{}
	rcvr.DescriptorMatcher = NewDescriptorMatcher(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewBFMatcher2(normType int, crossCheck bool) (rcvr *BFMatcher) {
	rcvr = &BFMatcher{}
	rcvr.DescriptorMatcher = NewDescriptorMatcher(BFMatcherNative_BFMatcher_0(normType, crossCheck))
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewBFMatcher3() (rcvr *BFMatcher) {
	rcvr = &BFMatcher{}
	rcvr.DescriptorMatcher = NewDescriptorMatcher(BFMatcherNative_BFMatcher_1())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func BFMatcherCreate(normType int, crossCheck bool) *BFMatcher {
	retVal := NewBFMatcher(BFMatcherNative_create_0(normType, crossCheck))
	return retVal
}
func BFMatcherCreate2() *BFMatcher {
	retVal := NewBFMatcher(BFMatcherNative_create_1())
	return retVal
}
func (rcvr *BFMatcher) finalize() {
	BFMatcherNative_delete(rcvr.GetNativeObjAddr())
}
