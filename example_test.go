package tmpfile

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ExampleTempFile() {
	f, remove, err := TempFile(".")
	if err != nil {
		panic(err)
	}
	if remove {
		defer os.Remove(f.Name()) // clean up
	}
	defer f.Close()

	if _, err := io.WriteString(f, "example"); err != nil {
		log.Fatal(err)
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
	// Output: example
}

func ExampleLink() {
	f, _, err := TempFile(".")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := io.WriteString(f, "example"); err != nil {
		log.Fatal(err)
	}

	dir, err := os.MkdirTemp(".", "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(dir)

	path := filepath.Join(dir, "link-example")
	defer os.Remove(path)

	if err := Link(f, path); err != nil {
		log.Fatal(err)
	}

	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
	// Output: example
}
