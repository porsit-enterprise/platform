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

func Open(path string) (*os.File, error) {
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, fs.ErrNotExist
		}
		return nil, fmt.Errorf("error in accessing file: %w", err)
	}

	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("error in opening file: %w", err)
	}

	fileStat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("error in file: %w", err)
	}

	if fileStat.Size() == 0 {
		return nil, io.EOF
	}

	return file, nil
}
