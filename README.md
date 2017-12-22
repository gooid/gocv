### Go bindings for OpenCV 3.3.1 ###

以 opencv-331.jar 为基础转换过来的（除highgui外），utils中的内容放到了core里。
需用到opencv对应SDK中for java的jni库.
比如在windows,amd64下运行需要用到 opencv-3.3.1-vc14.exe 解出来的 build/java/x64/opencv_java331.dll，
在android下运行需opencv-3.3.1-android-sdk.zip中的OpenCV-android-sdk\sdk\native\libs下对应arch的libopencv_java3.so 。

有例子在:
```
    cmd/facedetect
    cmd/inception
```

关于Converters要特别说明一下，在原始opencv-331.jar中Converters是在utils目录下的，现在把它移到了core中。并且对它的函数名进行了重命名，基本规则是：

java | go | 
---------|----------|
 vector_Point_to_Mat | ConvertersVectorPoint | 
 Mat_to_vector_Point | ConvertersToVectorPoint |


如果你想试试opencv的功能，欢迎您使用它，所需要的一些图片和资源可以在opencv的sdk中找到。
