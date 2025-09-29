package file

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Open(configPath string) (*os.File, error) {
	if _, err := os.Stat(configPath); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, fs.ErrNotExist
		}
		return nil, fmt.Errorf("error in accessing config file: %w", err)
	}

	file, err := os.Open(filepath.Clean(configPath))
	if err != nil {
		return nil, fmt.Errorf("error in opening config file: %w", err)
	}

	fileStat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("error in config file: %w", err)
	}

	if fileStat.Size() == 0 {
		return nil, io.EOF
	}

	return file, nil
}
