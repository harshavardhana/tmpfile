//go:build !linux

package tmpfile

import (
	"os"
)

// TempFile this is wrapper around os.CreateTemp(dir, "") on non-Linux
func TempFile(dir string) (f *os.File, remove bool, err error) {
	f, err = os.CreateTemp(dir, "")
	return f, err == nil, err
}

// Link this is wrapper around os.Rename(old, new) on non-Linux.
func Link(f *os.File, newpath string) error {
	return os.Rename(f.Name(), newpath)
}
