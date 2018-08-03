package main

import (
	"container/list"
	"fmt"
	"log"
)

type files struct {
	list *list.List
}

func newFiles() *files {
	return &files{
		list: list.New(),
	}
}

func (f *files) modifyList(op patchOperation) {
	switch op.Op {
	case "add":
		f.addFile(op.Value.Name)
	case "remove":
		f.removeFile(op.Value.Name)
	}
}

func (f *files) printList() {
	for e := f.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func (f *files) addFile(name string) {
	for e := f.list.Front(); e != nil; e = e.Next() {
		if e.Value.(string) > name {
			f.list.InsertBefore(name, e)
			return
		}
	}
	f.list.PushBack(name)
}

func (f *files) removeFile(name string) {
	for e := f.list.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == name {
			f.list.Remove(e)
			return
		}
	}
	log.Printf("[Error] file not found: %s", name)
}
