package orm

import (
	"os"
	"path/filepath"
)

// Host are instances under a organisation.
type Host struct {
	hostname string
	ip       string
}

// Organisation contains one or more instances.
type Organisation struct {
	name  string
	tags  string
	id    int16
	hosts []Host
}

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func (o *Organisation) readConfig(hostfile string) {
	// hostconfig, err := os.Open(hostfile)
}
