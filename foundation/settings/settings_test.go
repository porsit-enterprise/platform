package settings_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/porsit-enterprise/platform/core"
	. "github.com/porsit-enterprise/platform/foundation/settings"
	_ "github.com/porsit-enterprise/platform/test"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func TestLoad(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "porsit_platform_test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpdir)
	}()

	resdir := filepath.Join(tmpdir, core.RESOURCES_DIRECTORY)
	if err := os.MkdirAll(resdir, 0755); err != nil {
		t.Fatal(err)
	}

	fpath := filepath.Join(resdir, SETTINGS_FILE)
	if err := os.WriteFile(fpath, []byte("key: value\n"), 0644); err != nil {
		t.Fatal(err)
	}

	s := &struct {
		Key string `yaml:"key"`
	}{}

	err = Load(tmpdir, s)
	if err != nil {
		t.Fatal(err)
	}

	if s.Key != "value" {
		t.Fatalf("expected key to be 'value', got '%s'", s.Key)
	}
}
