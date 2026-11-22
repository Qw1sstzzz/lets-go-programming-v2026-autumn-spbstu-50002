package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o644
)

func WriteJSON(path string, data any) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPermissions); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, filePermissions)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}
