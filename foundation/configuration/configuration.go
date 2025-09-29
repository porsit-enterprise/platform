package configuration

import (
	"cmp"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	core "github.com/porsit-enterprise/platform/core"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Load(basePath string) (Properties, error) {
	basePath = cmp.Or(basePath, ".")

	// Load file

	configPath := filepath.Join(basePath, core.RESOURCES_DIRECTORY, CONFIG_FILE)

	slog.Info("load configuration", slog.String("path", configPath))

	fileBase, err := open(configPath)
	if fileBase != nil {
		defer fileBase.Close()
	}
	if err != nil && !errors.Is(err, io.EOF) {
		return Properties{}, err
	}

	// Set config

	prop := new(Properties)

	decoder := yaml.NewDecoder(fileBase)
	if err := decoder.Decode(prop); err != nil {
		return Properties{}, fmt.Errorf("error in decoding config base file: %w", err)
	}

	slog.Debug("configuration", slog.Any("properties", *prop))

	return *prop, nil
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func open(configPath string) (*os.File, error) {
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
