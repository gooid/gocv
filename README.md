### Go bindings for OpenCV 3.3.1 ###

以 opencv-331.jar 为基础转换过来的（除highgui外），utils中的内容放到了core里。
在用到for java的jni库，比如在windows,amd64下运行需对应的 build/java/x64/opencv_java331.dll，在android下运行需opencv-3.3.1-android-sdk.zip中的OpenCV-android-sdk\sdk\native\libs下对应arch的libopencv_java3.so 。

有例子在:
    cmd/facedetect
    cmd/inception