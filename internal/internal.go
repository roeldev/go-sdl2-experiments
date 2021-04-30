package internal

import (
	"path"
	"runtime"
)

func ExampleName() string {
	result := "sdlkit example"
	if _, file, _, ok := runtime.Caller(1); ok {
		name := path.Base(path.Dir(file))
		if name != "." && name != "/" {
			result += ": " + name
		}
	}

	return result
}
