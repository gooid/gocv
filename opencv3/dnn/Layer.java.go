package dnn

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type Layer struct {
	*Algorithm
}

func NewLayer(addr int64) (rcvr *Layer) {
	rcvr = &Layer{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *Layer) Finalize(inputs []*Mat) []*Mat {
	inputs_mat := ConvertersVectorMat(inputs)
	retValMat := NewMat(LayerNative_finalize_0(rcvr.GetNativeObjAddr(), inputs_mat.GetNativeObjAddr()))
	return ConvertersToVectorMat(retValMat)
}
func (rcvr *Layer) Finalize2(inputs []*Mat) (outputs []*Mat) {
	inputs_mat := ConvertersVectorMat(inputs)
	outputs_mat := NewMat2()
	LayerNative_finalize_1(rcvr.GetNativeObjAddr(), inputs_mat.GetNativeObjAddr(), outputs_mat.GetNativeObjAddr())
	outputs = ConvertersToVectorMat(outputs_mat)
	outputs_mat.Release()
	return
}
func (rcvr *Layer) finalize() {
	LayerNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *Layer) Forward(inputs []*Mat) (outputs []*Mat, internals []*Mat) {
	inputs_mat := ConvertersVectorMat(inputs)
	outputs_mat := ConvertersVectorMat(outputs)
	internals_mat := ConvertersVectorMat(internals)
	LayerNative_forward_0(rcvr.GetNativeObjAddr(), inputs_mat.GetNativeObjAddr(), outputs_mat.GetNativeObjAddr(), internals_mat.GetNativeObjAddr())
	outputs = ConvertersToVectorMat(outputs_mat)
	outputs_mat.Release()
	internals = ConvertersToVectorMat(internals_mat)
	internals_mat.Release()
	return
}
func (rcvr *Layer) Get_blobs() []*Mat {
	retValMat := NewMat(LayerNative_get_blobs_0(rcvr.GetNativeObjAddr()))
	return ConvertersToVectorMat(retValMat)
}
func (rcvr *Layer) Get_name() string {
	retVal := LayerNative_get_name_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *Layer) Get_preferableTarget() int {
	retVal := LayerNative_get_preferableTarget_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *Layer) Get_type() string {
	retVal := LayerNative_get_type_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *Layer) Run(inputs []*Mat) (outputs []*Mat, internals []*Mat) {
	inputs_mat := ConvertersVectorMat(inputs)
	outputs_mat := NewMat2()
	internals_mat := ConvertersVectorMat(internals)
	LayerNative_run_0(rcvr.GetNativeObjAddr(), inputs_mat.GetNativeObjAddr(), outputs_mat.GetNativeObjAddr(), internals_mat.GetNativeObjAddr())
	outputs = ConvertersToVectorMat(outputs_mat)
	outputs_mat.Release()
	internals = ConvertersToVectorMat(internals_mat)
	internals_mat.Release()
	return
}
func (rcvr *Layer) Set_blobs(blobs []*Mat) {
	blobs_mat := ConvertersVectorMat(blobs)
	LayerNative_set_blobs_0(rcvr.GetNativeObjAddr(), blobs_mat.GetNativeObjAddr())
	return
}
