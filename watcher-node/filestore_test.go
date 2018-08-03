package main

import (
	"reflect"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestProcessingEvents(t *testing.T) {
	tests := []struct {
		scenario          string
		event             fsnotify.Event
		actualFilestore   filestore
		expectedFilestore filestore
	}{
		{
			scenario: "add file",
			event: fsnotify.Event{
				Op:   fsnotify.Create,
				Name: "/my/test/file.txt",
			},
			actualFilestore: filestore{},
			expectedFilestore: filestore{
				"file.txt": {},
			},
		},
		{
			scenario: "delete file",
			event: fsnotify.Event{
				Op:   fsnotify.Remove,
				Name: "/my/test/file.txt",
			},
			actualFilestore: filestore{
				"file.txt": {},
			},
			expectedFilestore: filestore{},
		},
		{
			scenario: "rename file",
			event: fsnotify.Event{
				Op:   fsnotify.Rename,
				Name: "/my/test/file.txt",
			},
			actualFilestore: filestore{
				"file.txt": {},
			},
			expectedFilestore: filestore{},
		},
		{
			scenario: "unknown op",
			event: fsnotify.Event{
				Op:   fsnotify.Chmod,
				Name: "/my/test/file.txt",
			},
			actualFilestore: filestore{
				"file.txt": {},
			},
			expectedFilestore: filestore{
				"file.txt": {},
			},
		},
	}

	for _, test := range tests {
		test.actualFilestore.processEvent(test.event)
		if !reflect.DeepEqual(test.actualFilestore, test.expectedFilestore) {
			t.Errorf("%s, expected: %v, got: %v", test.scenario, test.expectedFilestore, test.actualFilestore)
		}
	}
}
