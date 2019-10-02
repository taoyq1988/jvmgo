package classpath

import (
	"os"
	"strings"
)

const (
	pathListSeparator = string(os.PathListSeparator)
	pathSeparator     = string(os.PathSeparator)
)

type Entry interface {
	readClass(name string) ([]byte, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip") {
		return newZipEntry(path)
	}
	return newDirEntry(pathListSeparator)
}
