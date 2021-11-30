package fs

import (
	"os"
	"path/filepath"
)

func CreateNotExistsDir(dir string, perm os.FileMode) error {
	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err = os.MkdirAll(dir, perm); err != nil {
			return err
		}
	}

	return nil
}

func CreateNotExistsFilepath(path string) error {
	err := CreateNotExistsDir(filepath.Dir(path), 0775)
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if f, err := os.Create(path); err != nil {
			return err
		} else {
			f.Close()
		}
	}

	return nil
}
