package yaml

import (
	"io"
	"os"

	"cyberpull.com/gotk/v2"
)

func Read[T any](reader io.Reader) (value T, err error) {
	var b []byte

	if b, err = io.ReadAll(reader); err != nil {
		return
	}

	value, err = Decode[T](b)

	return
}

func ReadFile[T any](name string) (value T, err error) {
	var file *os.File

	if file, err = os.Open(name); err != nil {
		return
	}

	defer file.Close()

	value, err = Read[T](file)

	return
}

func GetConfigFile[T any](paths ...string) (value T, err error) {
	var name string

	if name, err = gotk.PathFromExecutable(paths...); err != nil {
		return
	}

	value, err = ReadFile[T](name)

	return
}
