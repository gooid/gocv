// +build ignore

package core

import "fmt"

const serialVersionUIDCvException = 1

type CvException struct {
	*RuntimeException
}

func NewCvException(msg string) (rcvr *CvException) {
	rcvr = &CvException{}
	rcvr.RuntimeException = NewRuntimeException(msg)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *CvException) String() string {
	return fmt.Sprintf("%v%v%v", "CvException [", super.toString(), "]")
}
