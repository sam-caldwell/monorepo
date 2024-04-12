package fileWatcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sam-caldwell/monorepo/go/types"
)

// eventHandlerFunc - A handler function to fire when a fsnotify event occurs.
type eventHandlerFunc func(event fsnotify.Event, pathFile types.FilePath)
