package main

import (
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type filestore map[string]struct{}

func (fs filestore) addFiles(files []os.FileInfo) {
	for _, file := range files {
		fs[file.Name()] = struct{}{}
	}
}

func (fs filestore) processEvent(event fsnotify.Event) {
	filename := filepath.Base(event.Name)
	switch event.Op {
	case fsnotify.Create:
		fs[filename] = struct{}{}
	case fsnotify.Remove, fsnotify.Rename:
		delete(fs, filename)
	}
}
