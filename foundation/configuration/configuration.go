package configuration

import (
	"cmp"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	core "github.com/porsit-enterprise/platform/core"
	pkg_file "github.com/porsit-enterprise/platform/pkg/file"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

var instance Properties

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Load(basePath string) (Properties, error) {
	basePath = cmp.Or(basePath, ".")

	// Load file

	cfile := CONFIG_FILE

	if os.Getenv(core.ENVIRONMENT) == core.ENVIRONMENT_TEST {
		cfile = _CONFIG_FILE_TEST
	}

	configPath := filepath.Join(basePath, core.RESOURCES_DIRECTORY, cfile)

	slog.Info("load configuration", slog.String("path", configPath))

	file, err := pkg_file.Open(configPath)
	if file != nil {
		defer func() {
			err := file.Close()
			if err != nil {
				slog.Warn("error in closing config file", slog.Any("error", err))
			}
		}()
	}
	if err != nil && !errors.Is(err, io.EOF) {
		return Properties{}, err
	}

	// Set config

	prop := new(Properties)

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(prop); err != nil {
		return Properties{}, fmt.Errorf("error in decoding config file: %w", err)
	}

	// Validate config

	if prop.Version != VERSION {
		return Properties{}, fmt.Errorf("invalid config version: %s, expected: %s", prop.Version, VERSION)
	}

	slog.Debug("configuration", slog.Any("properties", *prop))

	instance = *prop

	return instance, nil
}

func Get() Properties {
	return instance
}
