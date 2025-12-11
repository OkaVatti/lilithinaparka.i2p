package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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
