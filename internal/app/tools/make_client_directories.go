package tools

import (
	"errors"
	"os"
)

func MakeBaseDirectories() error {
	if err := os.MkdirAll("/tmp/keeper/db", 0777); err != nil {
		return err
	}

	if err := os.MkdirAll("/tmp/keeper/files", 0777); err != nil {
		return err
	}

	if _, err := os.Stat("/tmp/keeper/db/keeper.db"); errors.Is(err, os.ErrNotExist) {
		if _, err1 := os.Create("/tmp/keeper/db/keeper.db"); err != nil {
			return err1
		}
	}

	return nil
}

func MakeFilesDirectory(login string) error {
	if err := os.MkdirAll("/tmp/keeper/files/"+login, 0777); err != nil {
		return err
	}
	return nil
}
