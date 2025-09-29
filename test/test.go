package test_test

import (
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/porsit-enterprise/platform/core"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

const (
	TestFAIL string = "%s(%q) == %q, expected %q"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

var ConfigPath string = filepath.Join("..", "..")

//──────────────────────────────────────────────────────────────────────────────────────────────────

func init() {
	if err := os.Setenv(core.ENVIRONMENT, core.ENVIRONMENT_TEST); err != nil {
		log.Fatal("can't set test envrionment variable")
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))
}
