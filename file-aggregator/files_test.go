package main

import (
	"container/list"
	"reflect"
	"testing"
)

func TestModifyFilesList(t *testing.T) {
	tests := []struct {
		scenario          string
		operation         patchOperation
		initialFileNames  []string
		expectedFileNames []string
	}{
		{
			scenario: "add file",
			operation: patchOperation{
				Op:   "add",
				Path: "/files",
				Value: fileMetadata{
					Name: "newfile.txt",
				},
			},
			initialFileNames:  []string{"oldfile.txt"},
			expectedFileNames: []string{"newfile.txt", "oldfile.txt"},
		},
		{
			scenario: "remove file",
			operation: patchOperation{
				Op:   "remove",
				Path: "/files",
				Value: fileMetadata{
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
			operation: patchOperation{
				Op:   "modify",
				Path: "/files",
				Value: fileMetadata{
					Name: "name.txt",
				},
			},
			initialFileNames:  []string{"oldfile.txt", "somefile.txt"},
			expectedFileNames: []string{"oldfile.txt", "somefile.txt"},
		},
		{
			scenario: "remove unknown file",
			operation: patchOperation{
				Op:   "remove",
				Path: "/files",
				Value: fileMetadata{
					Name: "name.txt",
				},
			},
			initialFileNames:  []string{"oldfile.txt", "somefile.txt"},
			expectedFileNames: []string{"oldfile.txt", "somefile.txt"},
		},
		{
			scenario: "remove unknown file",
			operation: patchOperation{
				Op:   "remove",
				Path: "/files",
				Value: fileMetadata{
					Name: "name.txt",
				},
			},
			initialFileNames:  []string{"oldfile.txt", "somefile.txt"},
			expectedFileNames: []string{"oldfile.txt", "somefile.txt"},
		},
		{
			scenario: "add file to empty list",
			operation: patchOperation{
				Op:   "add",
				Path: "/files",
				Value: fileMetadata{
					Name: "name.txt",
				},
			},
			initialFileNames:  []string{},
			expectedFileNames: []string{"name.txt"},
		},
	}

	for _, test := range tests {
		actualFiles := files{
			list: createOrderedFileList(test.initialFileNames),
		}
		actualFiles.modifyList(test.operation)
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

func getOrderedFileNames(files files) []string {
	names := []string{}
	for e := files.list.Front(); e != nil; e = e.Next() {
		names = append(names, e.Value.(string))
	}
	return names
}
