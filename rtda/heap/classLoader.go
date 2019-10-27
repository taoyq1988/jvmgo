package heap

import (
	"fmt"
	"github.com/taoyq1988/jvmgo/classfile"
	"github.com/taoyq1988/jvmgo/classpath"
	"github.com/taoyq1988/jvmgo/vmerrors"
)

const (
	jlObjectClassName       = "java/lang/Object"
	jlClassClassName        = "java/lang/Class"
	jlStringClassName       = "java/lang/String"
	jlThreadClassName       = "java/lang/Thread"
	jlCloneableClassName    = "java/lang/Cloneable"
	ioSerializableClassName = "java/io/Serializable"
)

var (
	bootLoader           *ClassLoader
	_jlObjectClass       *Class
	_jlClassClass        *Class
	_jlStringClass       *Class
	_jlThreadClass       *Class
	_jlCloneableClass    *Class
	_ioSerializableClass *Class
)

type ClassLoader struct {
	classpath *classpath.ClassPath
	classMap  map[string]*Class
	verbose   bool // debug
}

func BootLoader() *ClassLoader {
	return bootLoader
}

func InitBootLoader(cp *classpath.ClassPath, verbose bool) {
	bootLoader = &ClassLoader{
		classpath: cp,
		classMap:  make(map[string]*Class),
		verbose:   verbose,
	}
	bootLoader._init()
}

func (loader *ClassLoader) ClassPath() *classpath.ClassPath {
	return loader.classpath
}

func (loader *ClassLoader) JLObjectClass() *Class {
	return _jlObjectClass
}
func (loader *ClassLoader) JLClassClass() *Class {
	return _jlClassClass
}
func (loader *ClassLoader) JLStringClass() *Class {
	return _jlStringClass
}
func (loader *ClassLoader) JLThreadClass() *Class {
	return _jlThreadClass
}

func (loader *ClassLoader) getClass(name string) *Class {
	if class, ok := loader.classMap[name]; ok {
		return class
	}
	panic("class not loaded " + name)
}

func (loader *ClassLoader) GetPrimitiveClass(name string) *Class {
	return loader.getClass(name)
}

func (loader *ClassLoader) _init() {
	_jlObjectClass = loader.LoadClass(jlObjectClassName)
	_jlClassClass = loader.LoadClass(jlClassClassName)
	for _, class := range loader.classMap {
		if class.JClass == nil {
			class.JClass = _jlClassClass.NewObj()
			class.JClass.Extra = class
		}
	}
	_jlCloneableClass = loader.LoadClass(jlCloneableClassName)
	_ioSerializableClass = loader.LoadClass(ioSerializableClassName)
	_jlThreadClass = loader.LoadClass(jlThreadClassName)
	_jlStringClass = loader.LoadClass(jlStringClassName)
	loader.loadPrimitiveClasses()
	loader.loadPrimitiveArrayClasses()
}

func (loader *ClassLoader) LoadClass(className string) *Class {
	if class, ok := loader.classMap[className]; ok {
		return class
	}
	if className[0] == '[' {
		return loader._getRefArrayClass(className)
	} else {
		return loader.reallyLoadClass(className)
	}
}

func (loader *ClassLoader) loadPrimitiveClasses() {
	for _, primitiveType := range PrimitiveTypes {
		loader.loadPrimitiveClass(primitiveType.Name)
	}
}
func (loader *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{Name: className}
	//class.classLoader = loader
	class.JClass = _jlClassClass.NewObj()
	class.JClass.Extra = class
	class.MarkFullyInitialized()
	loader.classMap[className] = class
}

func (loader *ClassLoader) loadPrimitiveArrayClasses() {
	for _, primitiveType := range PrimitiveTypes {
		loader.loadArrayClass(primitiveType.ArrayClassName)
	}
}
func (loader *ClassLoader) loadArrayClass(className string) *Class {
	class := &Class{Name: className}
	class.SuperClass = _jlObjectClass
	class.Interfaces = []*Class{_jlCloneableClass, _ioSerializableClass}
	class.JClass = _jlClassClass.NewObj()
	class.JClass.Extra = class
	createVTable(class)
	class.MarkFullyInitialized()
	loader.classMap[className] = class
	return class
}

func (loader *ClassLoader) getArrayClass(componentClass *Class) *Class {
	arrClassName := "[L" + componentClass.Name + ";"
	return loader._getRefArrayClass(arrClassName)
}

func (loader *ClassLoader) _getRefArrayClass(className string) *Class {
	if class, ok := loader.classMap[className]; ok {
		return class
	}
	return loader.loadArrayClass(className)
}

func (loader *ClassLoader) reallyLoadClass(className string) *Class {
	data, err := loader.classpath.ReadClass(className)
	if err != nil {
		panic(vmerrors.NewClassNotFoundError(SlashToDot(className)))
	}
	class := loader._loadClass(className, data)
	if loader.verbose {
		fmt.Printf("[Loaded class %s\n", className)
	}
	return class
}

func (loader *ClassLoader) _loadClass(name string, data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic(fmt.Sprintf("failed to parse class file %s, for error %s", name, err.Error()))
	}
	class := newClass(cf)
	//1. resolve super class
	if class.superClassName != "" {
		class.SuperClass = loader.LoadClass(class.superClassName)
	}
	//2. resolve interfaces
	if len(class.interfaceNames) > 0 {
		class.Interfaces = make([]*Class, len(class.interfaceNames))
		for i, interfaceName := range class.interfaceNames {
			class.Interfaces[i] = loader.LoadClass(interfaceName)
		}
	}
	calcStaticFieldSlotIDs(class)
	calcInstanceFieldSlotIDs(class)
	createVTable(class)
	prepare(class)
	loader.classMap[name] = class

	if _jlClassClass != nil {
		class.JClass = _jlClassClass.NewObj()
		class.JClass.Extra = class
	}
	return class
}

func calcStaticFieldSlotIDs(class *Class) {
	slotID := uint(0)
	for _, field := range class.Fields {
		if field.IsStatic() {
			field.SlotID = slotID
			slotID++
		}
	}
	class.staticFieldCount = slotID
}

func calcInstanceFieldSlotIDs(class *Class) {
	slotID := uint(0)
	if class.superClassName != "" {
		slotID = class.SuperClass.instanceFieldCount
	}
	for _, field := range class.Fields {
		if !field.IsStatic() {
			field.SlotID = slotID
			slotID++
		}
	}
	class.instanceFieldCount = slotID
}

func prepare(class *Class) {
	class.StaticFieldSlots = make([]Slot, class.staticFieldCount)
	for _, field := range class.Fields {
		if field.IsStatic() {
			class.StaticFieldSlots[field.SlotID] = EmptySlot
		}
	}
}
