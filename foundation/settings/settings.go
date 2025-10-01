package settings

import (
	"cmp"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v3"

	core "github.com/porsit-enterprise/platform/core"
	pkg_file "github.com/porsit-enterprise/platform/pkg/file"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func Load(basePath string, properties any) error {
	basePath = cmp.Or(basePath, ".")

	// Load file

	settingsPath := filepath.Join(basePath, core.RESOURCES_DIRECTORY, SETTINGS_FILE)

	slog.Info("load settings", slog.String("path", settingsPath))

	file, err := pkg_file.Open(settingsPath)
	if file != nil {
		defer func() {
			err := file.Close()
			if err != nil {
				slog.Warn("error in closing config file", slog.Any("error", err))
			}
		}()
	}
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	// Set settings

	if properties == nil {
		return fmt.Errorf("properties argument must not be nil")
	}
	if reflect.ValueOf(properties).Kind() != reflect.Pointer {
		return fmt.Errorf("properties argument must be a pointer")
	}

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(properties); err != nil {
		return fmt.Errorf("error in decoding settings file: %w", err)
	}

	slog.Debug("settings", slog.Any("properties", properties))

	return nil
}
