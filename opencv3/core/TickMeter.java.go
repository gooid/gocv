package core

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type TickMeter struct {
	nativeObj int64
}

func NewTickMeter(addr int64) (rcvr *TickMeter) {
	rcvr = &TickMeter{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewTickMeter2() (rcvr *TickMeter) {
	rcvr = &TickMeter{}
	rcvr.nativeObj = TickMeterNative_TickMeter_0()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func (rcvr *TickMeter) finalize() {
	TickMeterNative_delete(rcvr.nativeObj)
}
func (rcvr *TickMeter) GetCounter() int64 {
	retVal := TickMeterNative_getCounter_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *TickMeter) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *TickMeter) GetTimeMicro() float64 {
	retVal := TickMeterNative_getTimeMicro_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *TickMeter) GetTimeMilli() float64 {
	retVal := TickMeterNative_getTimeMilli_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *TickMeter) GetTimeSec() float64 {
	retVal := TickMeterNative_getTimeSec_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *TickMeter) GetTimeTicks() int64 {
	retVal := TickMeterNative_getTimeTicks_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *TickMeter) Reset() {
	TickMeterNative_reset_0(rcvr.nativeObj)
	return
}
func (rcvr *TickMeter) Start() {
	TickMeterNative_start_0(rcvr.nativeObj)
	return
}
func (rcvr *TickMeter) Stop() {
	TickMeterNative_stop_0(rcvr.nativeObj)
	return
}
