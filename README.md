# Go bindings for OpenCV 3.3.1 （full functions, easy to use）

是以 **opencv-331.jar** 为基础转换过来的（除highgui外），所以是个全功能的opencv版本，并且模块划分及接口都是原汁原味的。

## 安装（使用）Installation
需用到 opencv SDK中的jni库，仅仅需要运行库，而不需要java环境的支持。（这里用了一个简单的技巧来免除对java的依赖）

比如在windos和android下对应的运行库：

platform  | runtime library  |
---------|----------|
windows,amd64|opencv-3.3.1-vc14.exe/opencv/build/java/x64/opencv_java331.dll
android|opencv-3.3.1-android-sdk.zip/sdk/native/libs/{arch}/libopencv_java3.so

如在windows,amd64下只需把上面对应的opencv_java331.dll拷贝到环境变量PATH所包含的目录下，确定 go 和 gcc 都已准备好，然后
```
go get github.com/gooid/gocv/opencv3
cd github.com/gooid/gocv/opencv3/cmd/facedetect
go run facedetect.go ..\testdata\lena.jpg ..\testdata\haarcascade_frontalface_alt2.xml ..\testdata\haarcascade_eye_tree_eyeglasses.xml
```


##### facedetect.go
```
import (
	...
	"github.com/gooid/gocv/opencv3/core"
	"github.com/gooid/gocv/opencv3/imgcodecs"
	"github.com/gooid/gocv/opencv3/imgproc"
	"github.com/gooid/gocv/opencv3/objdetect"
) 

...
    org := imgcodecs.Imread("lena.jpg", imgcodecs.IMREAD_ANYCOLOR)
    defer org.Release() 
    
    gray := core.NewMat2()
    imgproc.CvtColor2(org, gray, imgproc.COLOR_RGB2GRAY)
    
    cascade := objdetect.NewCascadeClassifier2("haarcascade_frontalface_alt2.xml") 

    rv := core.NewMatOfRect()    
    cascade.DetectMultiScale2(gray, rv)
    // faces
    for _, r := range rv.ToArray() {
        fmt.Println("\t", r.X, r.Y, r.Width, r.Height) 
    }
...

```

有例子在(sample):
```
    opencv3/cmd/facedetect
    opencv3/cmd/inception
```

## 关于Converters的特别说明：
在原始opencv-331.jar中Converters是在utils目录下的，现在把它移到了core中。并且对它的函数名进行了重命名，基本规则是：

java | go | 
---------|----------|
 vector_Point_to_Mat | ConvertersVectorPoint | 
 Mat_to_vector_Point | ConvertersToVectorPoint |

还有就是ConvertersToVectorPoint对cols限制为1，但在使用中发现经常会得到rows为1但要用ConvertersToVector*转换的情况，所以顺手改了一下，现在兼容rows为1的情况，不再需要先把mat转90度了。

## 已知问题
#### 文件名和路径不能有非ASCII码（opencv的问题，用opencv-331.jar一样有问题）
#### 有几个函数没有实现（看了对应的opencv JNI函数就会知道原因了）：
* func (rcvr *TrainData) GetNames(names []*Mat)
* func (rcvr *Net) Forward5(outputBlobs []*Mat, outBlobNames []*Mat)
* func (rcvr *Net) GetFLOPS3(layerId int, netInputShapes []*Mat) (int64)
* func (rcvr *Net) GetFLOPS4(netInputShapes []*Mat) (int64)
* func (rcvr *Net) GetLayerNames() ([]*Mat) 
* func (rcvr *Net) GetLayerTypes(layersTypes []*Mat) 
* func (rcvr *Net) GetMemoryConsumption3(layerId int, netInputShapes []*Mat, weights []int64, blobs []int64) 
* func (rcvr *Net) SetInputsNames(inputBlobNames []*Mat)

## 其它
如果您想试试opencv的功能，欢迎您使用它。我也是想在android下跑opencv，但又不想用java。找了一圈，也没合适我用的，才决定自己动手。

所需要的一些图片和资源可以在opencv的sdk中找到。

## License

MIT

## Author

gooid

