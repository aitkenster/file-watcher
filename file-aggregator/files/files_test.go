package files

import (
	"container/list"
	"reflect"
	"testing"

	"github.com/aitkenster/file-watcher/file-aggregator/lib"
)

func TestModifyFilesList(t *testing.T) {
	tests := []struct {
		scenario          string
		operation         lib.PatchOperation
		initialFileNames  []string
		expectedFileNames []string
	}{
		{
			scenario: "add file",
			operation: lib.PatchOperation{
				Op:   "add",
				Path: "/files",
				Value: lib.FileMetadata{
					Name: "newfile.txt",
				},
			},
			initialFileNames:  []string{"oldfile.txt"},
			expectedFileNames: []string{"newfile.txt", "oldfile.txt"},
		},
		{
			scenario: "remove file",
			operation: lib.PatchOperation{
				Op:   "remove",
				Path: "/files",
				Value: lib.FileMetadata{
					Name: "removeme.txt",
				},
			},
			initialFileNames: []string{
				"oldfile.txt",
				"removeme.txt",
				"somefile.txt",
			},
			expectedFileNames: []string{"oldfile.txt", "somefile.txt"},
		},
		{
			scenario: "unknown op",
			operation: lib.PatchOperation{
				Op:   "modify",
				Path: "/files",
				Value: lib.FileMetadata{
					Name: "name.txt",
				},
			},
			initialFileNames:  []string{"oldfile.txt", "somefile.txt"},
			expectedFileNames: []string{"oldfile.txt", "somefile.txt"},
		},
		{
			scenario: "remove unknown file",
			operation: lib.PatchOperation{
				Op:   "remove",
				Path: "/files",
				Value: lib.FileMetadata{
					Name: "name.txt",
				},
			},
			initialFileNames:  []string{"oldfile.txt", "somefile.txt"},
			expectedFileNames: []string{"oldfile.txt", "somefile.txt"},
		},
		{
			scenario: "remove unknown file",
			operation: lib.PatchOperation{
				Op:   "remove",
				Path: "/files",
				Value: lib.FileMetadata{
					Name: "name.txt",
				},
			},
			initialFileNames:  []string{"oldfile.txt", "somefile.txt"},
			expectedFileNames: []string{"oldfile.txt", "somefile.txt"},
		},
		{
			scenario: "add file to empty list",
			operation: lib.PatchOperation{
				Op:   "add",
				Path: "/files",
				Value: lib.FileMetadata{
					Name: "name.txt",
				},
			},
			initialFileNames:  []string{},
			expectedFileNames: []string{"name.txt"},
		},
	}

	for _, test := range tests {
		actualFiles := Filestore{
			list: createOrderedFileList(test.initialFileNames),
		}
		actualFiles.ModifyList(test.operation)
		actualFileNames := getOrderedFileNames(actualFiles)
		if !reflect.DeepEqual(actualFileNames, test.expectedFileNames) {
			t.Errorf(
				"%s: expected: %v, got: %v",
				test.scenario,
				test.expectedFileNames,
				actualFileNames,
			)
		}
	}
}

func createOrderedFileList(items []string) *list.List {
	list := list.New()
	for _, item := range items {
		list.PushBack(item)
	}
	return list
}

func getOrderedFileNames(files Filestore) []string {
	names := []string{}
	for e := files.list.Front(); e != nil; e = e.Next() {
		names = append(names, e.Value.(string))
	}
	return names
}
