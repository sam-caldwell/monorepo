package file

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"path/filepath"
	"strings"
)

// GetExtension - return the current file's extension
//
//		 The goal is to create an abstraction layer so changes to golang will not require
//		 significant effort to update all projects.
//
//	     - returns empty string if not initialized
//
//		 (c) 2023 Sam Caldwell.  MIT License
func (fp *File) GetExtension() string {
	if fp.handle == nil {
		return words.EmptyString
	}
	return strings.ToLower(filepath.Ext(fp.handle.Name()))
}
