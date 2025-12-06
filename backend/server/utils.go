package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ensure creates a dir if missing
func ensure(path string) error {
	return os.MkdirAll(path, 0o755)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func writeFileAtomic(path string, content []byte) error {
	dir := filepath.Dir(path)
	if err := ensure(dir); err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := ioutil.WriteFile(tmp, content, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func slugFromFilename(name string) string {
	n := strings.TrimSuffix(name, filepath.Ext(name))
	n = strings.ToLower(strings.ReplaceAll(n, " ", "-"))
	n = strings.ReplaceAll(n, "_", "-")
	return n
}

func joinTags(slice []string) string {
	return strings.Join(slice, ",")
}

func splitTags(s string) []string {
	if s == "" {
		return []string{}
	}
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func errf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
