package configuration_test

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/porsit-enterprise/platform/core"
	"github.com/stretchr/testify/assert"

	. "github.com/porsit-enterprise/platform/foundation/configuration"
	test_testing "github.com/porsit-enterprise/platform/test"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func TestLoad(t *testing.T) {
	config := loadConfigurationForTesting(t)
	assert.NotZero(t, config)
}

func TestLoadPathNotExist(t *testing.T) {
	_, err := Load("/tmp/dummy_path")
	if assert.Error(t, err) {
		assert.Equal(t, fs.ErrNotExist, err)
	}
}

func TestVersion(t *testing.T) {
	config := loadConfigurationForTesting(t)
	t.Logf("version: %s", config.Version)
	assert.NotEmpty(t, config.Version)
}

func TestLoadEmptyFile(t *testing.T) {
	err := os.Mkdir(filepath.Join(os.TempDir(), core.RESOURCES_DIRECTORY), fs.ModePerm)
	if err != nil && !errors.Is(err, os.ErrExist) {
		t.Errorf("error in creating temp directory: %v", err)
		return
	}
	defer func() {
		err := os.RemoveAll(filepath.Join(os.TempDir(), core.RESOURCES_DIRECTORY))
		if err != nil {
			t.Errorf("error in removing temp directory: %v", err)
		}
	}()
	fileBase, err := os.Create(filepath.Join(os.TempDir(), core.RESOURCES_DIRECTORY, CONFIG_FILE))
	if err != nil {
		t.Errorf("error in creating temp config global file: %v", err)
		return
	}
	_ = fileBase.Close()
	file, err := os.Create(filepath.Join(os.TempDir(), core.RESOURCES_DIRECTORY, CONFIG_FILE))
	if err != nil {
		t.Errorf("error in creating temp config file: %v", err)
		return
	}
	_ = file.Close()
	_, err = Load(os.TempDir())
	assert.Error(t, err)
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func loadConfigurationForTesting(t *testing.T) Properties {
	t.Helper()
	config, err := Load(test_testing.ConfigPath)
	if err != nil {
		t.Fatalf("error in loading configuration: %v", err)
	}
	return config
}
