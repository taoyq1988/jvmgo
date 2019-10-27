package io

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
	"os"
)

const (
	//private native void writeBytes(byte b[], int off, int len, boolean append) throws IOException;
	fileOutPutStream = "java/io/FileOutputStream"
)

func init() {
	_fos(writeBytes, "writeBytes", "([BIIZ)V")
}

func _fos(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod(fileOutPutStream, name, desc, method)
}

func writeBytes(frame *rtda.Frame) {
	fosObj := frame.GetThis()
	byteArray := frame.GetRefVar(1)
	offset := frame.GetIntVar(2)
	length := frame.GetIntVar(3)

	fbObj := fosObj.GetFieldValue("fd", "Ljava/io/FileDescriptor;").Ref
	fd := fbObj.GetFieldValue("fd", "I").IntValue()
	if fd == 1 {
		printBytes := byteArray.GoBytes()[offset : offset+length]
		_, _ = os.Stdout.Write(printBytes)
	}
}
