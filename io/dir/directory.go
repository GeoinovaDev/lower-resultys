package dir

import (
	"os"
	"path/filepath"
	"runtime"
)

// BasePath ...
func BasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}

// Cwd ..
func Cwd() string {
	d, _ := os.Getwd()

	return d
}

// Remove remove diretorio e todo conteudo
func Remove(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	os.Remove(dir)
	return nil
}
