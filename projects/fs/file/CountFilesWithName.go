package file

import (
	"os"
	"path/filepath"
)

func CountManifestFilesWalk(path string, fileName string) (int, error) {
	count := 0
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == fileName {
			count++
		}
		return nil
	})
	return count, err
}
