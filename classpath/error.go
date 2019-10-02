package classpath

import "errors"

var (
	NotSupportPathError = errors.New("not support path, pelease usr jar or zip")
	NoSuchClassError    = errors.New("no such class error")
	CanNotFildClassFile = errors.New("can not find class")
)
