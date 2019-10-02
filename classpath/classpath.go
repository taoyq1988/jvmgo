package classpath

import (
	"strings"
)

const (
	classSuffix = ".class"
)

type ClassPath struct {
	CompositeEntry
}

func NewClassPath(bootClassPath, applicationClassPath string) *ClassPath {
	cp := &ClassPath{}
	cp.addEntry(newZipEntry(bootClassPath))
	cp.addEntry(newDirEntry(applicationClassPath))
	return cp
}

func(cp *ClassPath) ReadClass(className string) ([]byte, error) {
	className = strings.Replace(className, ".", pathSeparator, -1)
	className += classSuffix
	return cp.readClass(className)
}

func (cp *ClassPath) String() string {
	userClassPath := cp.CompositeEntry.entries[1]
	return userClassPath.String()
}
