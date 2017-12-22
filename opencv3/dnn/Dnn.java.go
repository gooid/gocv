package dnn

import (
	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const DNN_BACKEND_DEFAULT = 0
const DNN_BACKEND_HALIDE = 1
const DNN_TARGET_CPU = 0
const DNN_TARGET_OPENCL = 1

func BlobFromImage(image *Mat, scalefactor float64, size *Size, mean *Scalar, swapRB bool, crop bool) *Mat {
	retVal := NewMat(DnnNative_blobFromImage_0(image.GetNativeObjAddr(), scalefactor, size.Width, size.Height, mean.Val[0], mean.Val[1], mean.Val[2], mean.Val[3], swapRB, crop))
	return retVal
}
func BlobFromImage2(image *Mat) *Mat {
	retVal := NewMat(DnnNative_blobFromImage_1(image.GetNativeObjAddr()))
	return retVal
}
func BlobFromImages(images []*Mat, scalefactor float64, size *Size, mean *Scalar, swapRB bool, crop bool) *Mat {
	images_mat := ConvertersVectorMat(images)
	retVal := NewMat(DnnNative_blobFromImages_0(images_mat.GetNativeObjAddr(), scalefactor, size.Width, size.Height, mean.Val[0], mean.Val[1], mean.Val[2], mean.Val[3], swapRB, crop))
	return retVal
}
func BlobFromImages2(images []*Mat) *Mat {
	images_mat := ConvertersVectorMat(images)
	retVal := NewMat(DnnNative_blobFromImages_1(images_mat.GetNativeObjAddr()))
	return retVal
}
func CreateCaffeImporter(prototxt string, caffeModel string) *Importer {
	retVal := NewImporter(DnnNative_createCaffeImporter_0(prototxt, caffeModel))
	return retVal
}
func CreateCaffeImporter2(prototxt string) *Importer {
	retVal := NewImporter(DnnNative_createCaffeImporter_1(prototxt))
	return retVal
}
func CreateTensorflowImporter(model string) *Importer {
	retVal := NewImporter(DnnNative_createTensorflowImporter_0(model))
	return retVal
}
func CreateTorchImporter(filename string, isBinary bool) *Importer {
	retVal := NewImporter(DnnNative_createTorchImporter_0(filename, isBinary))
	return retVal
}
func CreateTorchImporter2(filename string) *Importer {
	retVal := NewImporter(DnnNative_createTorchImporter_1(filename))
	return retVal
}
func ReadNetFromCaffe(prototxt string, caffeModel string) *Net {
	retVal := NewNet(DnnNative_readNetFromCaffe_0(prototxt, caffeModel))
	return retVal
}
func ReadNetFromCaffe2(prototxt string) *Net {
	retVal := NewNet(DnnNative_readNetFromCaffe_1(prototxt))
	return retVal
}
func ReadNetFromDarknet(cfgFile string, darknetModel string) *Net {
	retVal := NewNet(DnnNative_readNetFromDarknet_0(cfgFile, darknetModel))
	return retVal
}
func ReadNetFromDarknet2(cfgFile string) *Net {
	retVal := NewNet(DnnNative_readNetFromDarknet_1(cfgFile))
	return retVal
}
func ReadNetFromTensorflow(model string, config string) *Net {
	retVal := NewNet(DnnNative_readNetFromTensorflow_0(model, config))
	return retVal
}
func ReadNetFromTensorflow2(model string) *Net {
	retVal := NewNet(DnnNative_readNetFromTensorflow_1(model))
	return retVal
}
func ReadNetFromTorch(model string, isBinary bool) *Net {
	retVal := NewNet(DnnNative_readNetFromTorch_0(model, isBinary))
	return retVal
}
func ReadNetFromTorch2(model string) *Net {
	retVal := NewNet(DnnNative_readNetFromTorch_1(model))
	return retVal
}
func ReadTorchBlob(filename string, isBinary bool) *Mat {
	retVal := NewMat(DnnNative_readTorchBlob_0(filename, isBinary))
	return retVal
}
func ReadTorchBlob2(filename string) *Mat {
	retVal := NewMat(DnnNative_readTorchBlob_1(filename))
	return retVal
}
func ShrinkCaffeModel(src string, dst string) {
	DnnNative_shrinkCaffeModel_0(src, dst)
	return
}
