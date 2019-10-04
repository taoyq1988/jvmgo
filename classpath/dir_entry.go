package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (dir *DirEntry) readClass(className string) ([]byte, error) {
	fileName := filepath.Join(dir.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dir *DirEntry) open() error {
	return nil
}

func (dir *DirEntry) release() error {
	return nil
}

func (dir *DirEntry) String() string {
	return dir.absDir
}
