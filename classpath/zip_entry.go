package classpath

import (
	zzip "archive/zip"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
	reader  *zzip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absZip, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	zip := &ZipEntry{absZip, nil}
	_ = zip.open()
	return zip
}

func (zip *ZipEntry) readClass(class string) ([]byte, error) {
	f, err := zip.findClass(class)
	if err != nil {
		return nil, err
	}
	return readClass(f)
}

func (zip *ZipEntry) findClass(class string) (*zzip.File, error) {
	for _, f := range zip.reader.File {
		if f.Name == class {
			return f, nil
		}
	}
	return nil, CanNotFildClassFile
}

func readClass(file *zzip.File) ([]byte, error) {
	rc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (zip *ZipEntry) open() error {
	r, err := zzip.OpenReader(zip.absPath)
	if err != nil {
		return err
	}
	zip.reader = r
	return nil
}

func (zip *ZipEntry) release() error {
	return zip.reader.Close()
}

func (zip *ZipEntry) String() string {
	return zip.absPath
}
