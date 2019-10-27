package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
	"runtime"
	"sort"
	"time"
	"unsafe"
)

const (
	system = "java/lang/System"
)

func init() {
	_system(arraycopy, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V")
	_system(currentTimeMillis, "currentTimeMillis", "()J")
	_system(identityHashCode, "identityHashCode", "(Ljava/lang/Object;)I")
	_system(initProperties, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;")
	_system(nanoTime, "nanoTime", "()J")
	_system(setErr0, "setErr0", "(Ljava/io/PrintStream;)V")
	_system(setIn0, "setIn0", "(Ljava/io/InputStream;)V")
	_system(setOut0, "setOut0", "(Ljava/io/PrintStream;)V")

	_system(mapLibraryName, "mapLibraryName", "(Ljava/lang/String;)Ljava/lang/String;")
}

func mapLibraryName(frame *rtda.Frame) {
	frame.PushRef(frame.GetRefVar(0))
}

func _system(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod(system, name, desc, method)
}

func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class
	destClass := dest.Class

	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	if srcClass.IsPrimitiveArray() || destClass.IsPrimitiveArray() {
		return srcClass == destClass
	}
	return true
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
func arraycopy(frame *rtda.Frame) {
	src := frame.GetRefVar(0)
	srcPos := frame.GetIntVar(1)
	dest := frame.GetRefVar(2)
	destPos := frame.GetIntVar(3)
	length := frame.GetIntVar(4)

	// NullPointerException
	if src == nil || dest == nil {
		panic("NPE") // todo
	}
	// ArrayStoreException
	if !checkArrayCopy(src, dest) {
		panic("ArrayStoreException")
	}
	// IndexOutOfBoundsException
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > heap.ArrayLength(src) ||
		destPos+length > heap.ArrayLength(dest) {

		panic("IndexOutOfBoundsException") // todo
	}

	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

// public static native long currentTimeMillis();
// ()J
func currentTimeMillis(frame *rtda.Frame) {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	frame.PushLong(millis)
}

// public static native int identityHashCode(Object x);
// (Ljava/lang/Object;)I
func identityHashCode(frame *rtda.Frame) {
	ref := frame.GetRefVar(0)

	// todo
	hashCode := int32(uintptr(unsafe.Pointer(ref)))
	frame.PushInt(hashCode)
}

func _getSysProps(absJavaHome string) map[string]string {
	return map[string]string{
		"java.version":         "1.8.0",
		"java.home":            absJavaHome,
		"java.class.version":   "52.0",
		"java.class.path":      heap.BootLoader().ClassPath().String(),
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              runtime.GOOS,   // todo
		"os.arch":              runtime.GOARCH, // todo
		"os.version":           "",             // todo
		"file.separator":       "/",            // todo os.PathSeparator
		"path.separator":       ":",            // todo os.PathListSeparator
		"line.separator":       "\n",           // todo
		"user.name":            "",             // todo
		"user.home":            "",             // todo
		"user.dir":             ".",            // todo
		"user.country":         "CN",           // todo
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
	}
}

func _getSysPropKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// private static native Properties initProperties(Properties props);
// (Ljava/util/Properties;)Ljava/util/Properties;
func initProperties(frame *rtda.Frame) {
	props := frame.GetRefVar(0)

	frame.PushRef(props)
	// public synchronized Object setProperty(String key, String value)
	setPropMethod := props.Class.GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	thread := frame.Thread

	sysPropMap := _getSysProps("/path/to/jre")
	sysPropKeys := _getSysPropKeys(sysPropMap)

	for _, key := range sysPropKeys {
		val := sysPropMap[key]
		jKey := heap.JSFromGoStr(key)
		jVal := heap.JSFromGoStr(val)
		args := []heap.Slot{heap.NewRefSlot(props), heap.NewRefSlot(jKey), heap.NewRefSlot(jVal)}
		thread.InvokeMethodWithShim(setPropMethod, args)
	}
}

// public static native long nanoTime();
// ()J
func nanoTime(frame *rtda.Frame) {
	nanoTime := time.Now().UnixNano()
	frame.PushLong(nanoTime)
}

// private static native void setErr0(PrintStream err);
// (Ljava/io/PrintStream;)V
func setErr0(frame *rtda.Frame) {
	err := frame.GetRefVar(0) // TODO

	sysClass := frame.GetClass()
	sysClass.SetStaticValue("err", "Ljava/io/PrintStream;", heap.NewRefSlot(err))
}

// private static native void setIn0(InputStream in);
// (Ljava/io/InputStream;)V
func setIn0(frame *rtda.Frame) {
	in := frame.GetRefVar(0) // TODO

	sysClass := frame.GetClass()
	sysClass.SetStaticValue("in", "Ljava/io/InputStream;", heap.NewRefSlot(in))
}

// private static native void setOut0(PrintStream out);
// (Ljava/io/PrintStream;)V
func setOut0(frame *rtda.Frame) {
	out := frame.GetRefVar(0) // TODO

	sysClass := frame.GetClass()
	sysClass.SetStaticValue("out", "Ljava/io/PrintStream;", heap.NewRefSlot(out))
}
