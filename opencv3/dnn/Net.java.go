package dnn

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type Net struct {
	nativeObj int64
}

func NewNet(addr int64) (rcvr *Net) {
	rcvr = &Net{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewNet2() (rcvr *Net) {
	rcvr = &Net{}
	rcvr.nativeObj = NetNative_Net_0()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func (rcvr *Net) Connect(outPin string, inpPin string) {
	NetNative_connect_0(rcvr.nativeObj, outPin, inpPin)
	return
}
func (rcvr *Net) DeleteLayer(layer *DictValue) {
	NetNative_deleteLayer_0(rcvr.nativeObj, layer.nativeObj)
	return
}
func (rcvr *Net) Empty() bool {
	retVal := NetNative_empty_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Net) EnableFusion(fusion bool) {
	NetNative_enableFusion_0(rcvr.nativeObj, fusion)
	return
}
func (rcvr *Net) finalize() {
	NetNative_delete(rcvr.nativeObj)
}
func (rcvr *Net) Forward(outputName string) *Mat {
	retVal := NewMat(NetNative_forward_0(rcvr.nativeObj, outputName))
	return retVal
}
func (rcvr *Net) Forward2() *Mat {
	retVal := NewMat(NetNative_forward_1(rcvr.nativeObj))
	return retVal
}
func (rcvr *Net) Forward3(outputBlobs []*Mat, outputName string) {
	outputBlobs_mat := ConvertersVectorMat(outputBlobs)
	NetNative_forward_2(rcvr.nativeObj, outputBlobs_mat.GetNativeObjAddr(), outputName)
	return
}
func (rcvr *Net) Forward4(outputBlobs []*Mat) {
	outputBlobs_mat := ConvertersVectorMat(outputBlobs)
	NetNative_forward_3(rcvr.nativeObj, outputBlobs_mat.GetNativeObjAddr())
	return
}

//func (rcvr *Net) Forward5(outputBlobs []*Mat, outBlobNames []*Mat) {
//	outputBlobs_mat := Convertersector_Mat(outputBlobs)
//	NetNative_forward_4(rcvr.GetNativeObjAddr(), outputBlobs_mat.GetNativeObjAddr(), outBlobNames)
//	return
//}
func (rcvr *Net) GetFLOPS(netInputShape *MatOfInt) int64 {
	netInputShape_mat := netInputShape
	retVal := NetNative_getFLOPS_0(rcvr.nativeObj, netInputShape_mat.GetNativeObjAddr())
	return retVal
}
func (rcvr *Net) GetFLOPS2(layerId int, netInputShape *MatOfInt) int64 {
	netInputShape_mat := netInputShape
	retVal := NetNative_getFLOPS_1(rcvr.nativeObj, layerId, netInputShape_mat.GetNativeObjAddr())
	return retVal
}

//func (rcvr *Net) GetFLOPS3(layerId int, netInputShapes []*Mat) (int64) {
//	retVal := NetNative_getFLOPS_2(rcvr.GetNativeObjAddr(), layerId, netInputShapes)
//	return retVal
//}
//func (rcvr *Net) GetFLOPS4(netInputShapes []*Mat) (int64) {
//	retVal := NetNative_getFLOPS_3(rcvr.GetNativeObjAddr(), netInputShapes)
//	return retVal
//}
func (rcvr *Net) GetLayer(layerId *DictValue) *Layer {
	retVal := NewLayer(NetNative_getLayer_0(rcvr.nativeObj, layerId.nativeObj))
	return retVal
}
func (rcvr *Net) GetLayerId(layer string) int {
	retVal := NetNative_getLayerId_0(rcvr.nativeObj, layer)
	return retVal
}

//func (rcvr *Net) GetLayerNames() ([]*Mat) {
//	retVal := NetNative_getLayerNames_0(rcvr.GetNativeObjAddr())
//	return retVal
//}
//func (rcvr *Net) GetLayerTypes(layersTypes []*Mat) {
//	NetNative_getLayerTypes_0(rcvr.GetNativeObjAddr(), layersTypes)
//	return
//}
func (rcvr *Net) GetLayersCount(layerType string) int {
	retVal := NetNative_getLayersCount_0(rcvr.nativeObj, layerType)
	return retVal
}
func (rcvr *Net) GetMemoryConsumption(netInputShape *MatOfInt, weights []int64, blobs []int64) {
	netInputShape_mat := netInputShape
	weights_out := make([]float64, 1)
	blobs_out := make([]float64, 1)
	NetNative_getMemoryConsumption_0(rcvr.nativeObj, netInputShape_mat.GetNativeObjAddr(), weights_out, blobs_out)
	if weights != nil {
		weights[0] = int64(weights_out[0])
	}
	if blobs != nil {
		blobs[0] = int64(blobs_out[0])
	}
	return
}
func (rcvr *Net) GetMemoryConsumption2(layerId int, netInputShape *MatOfInt, weights []int64, blobs []int64) {
	netInputShape_mat := netInputShape
	weights_out := make([]float64, 1)
	blobs_out := make([]float64, 1)
	NetNative_getMemoryConsumption_1(rcvr.nativeObj, layerId, netInputShape_mat.GetNativeObjAddr(), weights_out, blobs_out)
	if weights != nil {
		weights[0] = int64(weights_out[0])
	}
	if blobs != nil {
		blobs[0] = int64(blobs_out[0])
	}
	return
}

//func (rcvr *Net) GetMemoryConsumption3(layerId int, netInputShapes []*Mat, weights []int64, blobs []int64) {
//	weights_out := make([]float64, 1)
//	blobs_out := make([]float64, 1)
//	NetNative_getMemoryConsumption_2(rcvr.GetNativeObjAddr(), layerId, netInputShapes, weights_out, blobs_out)
//	if weights != nil {
//		weights[0] = int64(weights_out[0])
//	}
//	if blobs != nil {
//		blobs[0] = int64(blobs_out[0])
//	}
//	return
//}
func (rcvr *Net) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *Net) GetParam(layer *DictValue, numParam int) *Mat {
	retVal := NewMat(NetNative_getParam_0(rcvr.nativeObj, layer.GetNativeObjAddr(), numParam))
	return retVal
}
func (rcvr *Net) GetParam2(layer *DictValue) *Mat {
	retVal := NewMat(NetNative_getParam_1(rcvr.nativeObj, layer.GetNativeObjAddr()))
	return retVal
}
func (rcvr *Net) GetPerfProfile(timings *MatOfDouble) int64 {
	timings_mat := timings
	retVal := NetNative_getPerfProfile_0(rcvr.nativeObj, timings_mat.GetNativeObjAddr())
	return retVal
}
func (rcvr *Net) GetUnconnectedOutLayers() *MatOfInt {
	retVal := MatOfIntFromNativeAddr(NetNative_getUnconnectedOutLayers_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *Net) SetHalideScheduler(scheduler string) {
	NetNative_setHalideScheduler_0(rcvr.nativeObj, scheduler)
	return
}
func (rcvr *Net) SetInput(blob *Mat, name string) {
	NetNative_setInput_0(rcvr.nativeObj, blob.GetNativeObjAddr(), name)
	return
}
func (rcvr *Net) SetInput2(blob *Mat) {
	NetNative_setInput_1(rcvr.nativeObj, blob.GetNativeObjAddr())
	return
}

//func (rcvr *Net) SetInputsNames(inputBlobNames []*Mat) {
//	NetNative_setInputsNames_0(rcvr.GetNativeObjAddr(), inputBlobNames)
//	return
//}
func (rcvr *Net) SetParam(layer *DictValue, numParam int, blob *Mat) {
	NetNative_setParam_0(rcvr.nativeObj, layer.GetNativeObjAddr(), numParam, blob.GetNativeObjAddr())
	return
}
func (rcvr *Net) SetPreferableBackend(backendId int) {
	NetNative_setPreferableBackend_0(rcvr.nativeObj, backendId)
	return
}
func (rcvr *Net) SetPreferableTarget(targetId int) {
	NetNative_setPreferableTarget_0(rcvr.nativeObj, targetId)
	return
}
