package classpath

import "strings"

type CompositeEntry struct{
	entries []Entry
}

func newCompositeEntry(pathList string) *CompositeEntry {
	compoundEntry := &CompositeEntry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compoundEntry.addEntry(entry)
	}
	return compoundEntry
}

func (composite *CompositeEntry) addEntry(subEntry Entry) {
	composite.entries = append(composite.entries, subEntry)
}

func(composite *CompositeEntry) readClass(className string) (data []byte, err error) {
	for _, e := range composite.entries {
		data, err = e.readClass(className)
		if err == nil {
			return
		}
	}
	return
}

func(composite *CompositeEntry) open() error {
	//for _, e := range composite.entries {
	//	err := e.open()
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}

func(composite *CompositeEntry) release() error {
	//for _, e := range composite.entries {
	//	err := e.release()
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}

func(composite *CompositeEntry) String() string {
	paths := make([]string, 0)
	for _, e := range composite.entries {
		paths = append(paths, e.String())
	}
	return strings.Join(paths, pathListSeparator)
}

