package utils

import (
	cp "github.com/otiai10/copy"
)

func Copy(source string, destination string) error {
	err := cp.Copy(source, destination)
	return err
}
