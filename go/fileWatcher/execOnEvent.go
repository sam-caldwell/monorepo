package fileWatcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sam-caldwell/monorepo/go/types"
)

// ExecOnEvent - On a given file operation against pathFile, execute the provided handler function.
func ExecOnEvent(pathFile types.FilePath, watchEvent fsnotify.Op, handler eventHandlerFunc) (err error) {
	var watcher *fsnotify.Watcher
	if watcher, err = fsnotify.NewWatcher(); err != nil {
		return // exit go routine
	}
	if err = watcher.Add(pathFile.String()); err != nil {
		return // exit go routine
	}
	defer func() {
		_ = watcher.Close()
	}()

	// This go routine runs indefinitely
	go func() {
		// Start listening for events.
		for {
			var ok bool
			var event fsnotify.Event
			select {
			case event, ok = <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(watchEvent) {
					handler(event, pathFile)
				}
			case err, ok = <-watcher.Errors:
				if !ok {
					break
				}
			}
		}
	}()
	return err
}
