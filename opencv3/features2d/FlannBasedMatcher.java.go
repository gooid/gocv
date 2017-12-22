package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type FlannBasedMatcher struct {
	*DescriptorMatcher
}

func NewFlannBasedMatcher(addr int64) (rcvr *FlannBasedMatcher) {
	rcvr = &FlannBasedMatcher{}
	rcvr.DescriptorMatcher = NewDescriptorMatcher(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewFlannBasedMatcher2() (rcvr *FlannBasedMatcher) {
	rcvr = &FlannBasedMatcher{}
	rcvr.DescriptorMatcher = NewDescriptorMatcher(FlannBasedMatcherNative_FlannBasedMatcher_0())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func FlannBasedMatcherCreate() *FlannBasedMatcher {
	retVal := NewFlannBasedMatcher(FlannBasedMatcherNative_create_0())
	return retVal
}
func (rcvr *FlannBasedMatcher) finalize() {
	FlannBasedMatcherNative_delete(rcvr.GetNativeObjAddr())
}
