package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"path/filepath"
	"strings"
)

// filterDirectory - filter out undesired directory names
func filterDirectory(directoryList []string) []string {
	var result []string

	BlockedDirectories := map[string]int{
		"docs":               0,
		"external":           0,
		".git":               0,
		".github":            0,
		".idea":              0,
		"Makefile.d":         0,
		"database/sqldbtest": 0,
	}

	for _, dirPath := range directoryList {
		dirName := filepath.Base(dirPath)
		if _, ok := BlockedDirectories[dirName]; !ok {
			result = append(result, dirPath)
		}
	}
	return result
}

func errIfEmpty(filter *string) (err error) {
	if strings.TrimSpace(*filter) == "" {
		err = fmt.Errorf("missing argument.  use --help for usage")
	}
	return err
}

func sanitizeFilter(filter *string) (err error) {
	if strings.TrimSpace(strings.ToLower(*filter)) == "all" {
		*filter = "*"
	}
	err = errIfEmpty(filter)
	return err
}

func getDirectoryList(rootDirectory, filter *string) (list []string, err error) {
	var rawList []string
	if rawList, err = filepath.Glob(filepath.Join(*rootDirectory, *filter)); err != nil {
		return nil, err
	}
	for _, item := range rawList {
		if directory.Existsp(&item) && file.Exists(filepath.Join(item, "Manifest.yml")) {
			list = append(list, item)
		}
	}
	return list, err
}

func GetProjectList(classFilter, projectFilter *string) (list map[string][]string, err error) {
	/*
	   <root>/class/<project><code files>

	       class - string
	       project - path of string
	*/
	var classList []string
	var projectList []string
	rootDirectory := directory.GetCurrent()

	if err = sanitizeFilter(classFilter); err != nil {
		return list, err
	}
	if err = sanitizeFilter(projectFilter); err != nil {
		return list, err
	}
	list = make(map[string][]string)
	if classList, err = getDirectoryList(&rootDirectory, classFilter); err != nil {
		return nil, err
	}
	for _, classRoot := range filterDirectory(classList) {
		filter := *projectFilter
		for i := 0; i < 2; i++ {

			if projectList, err = getDirectoryList(&classRoot, &filter); err != nil {
				return nil, err
			}
			for _, item := range filterDirectory(projectList) {
				list[classRoot] = append(list[classRoot], item)
			}
			filter += "/*"
		}
	}
	return list, err
}
