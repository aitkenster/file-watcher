package files

import (
	"container/list"
	"fmt"
	"log"
	"sync"

	"github.com/aitkenster/file-watcher/file-aggregator/lib"
)

type Filestore struct {
	list  *list.List
	mutex sync.Mutex
}

func New(fileBatch []lib.FileMetadata) *Filestore {
	f := Filestore{
		list: list.New(),
	}
	for _, file := range fileBatch {
		f.addFile(file.Name)
	}
	return &f
}

func (f *Filestore) ModifyList(op lib.PatchOperation) {
	switch op.Op {
	case "add":
		f.addFile(op.Value.Name)
	case "remove":
		f.removeFile(op.Value.Name)
	}
}

func (f *Filestore) RetrieveAll() []lib.FileMetadata {
	fileSlice := []lib.FileMetadata{}
	for e := f.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		fileSlice = append(fileSlice, lib.FileMetadata{
			Name: e.Value.(string),
		})
	}
	return fileSlice
}

func (f *Filestore) addFile(name string) {
	f.lock()
	defer f.unlock()
	for e := f.list.Front(); e != nil; e = e.Next() {
		if e.Value.(string) > name {
			f.list.InsertBefore(name, e)
			return
		}
	}
	f.list.PushBack(name)
}

func (f *Filestore) removeFile(name string) {
	f.lock()
	defer f.unlock()
	for e := f.list.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == name {
			f.list.Remove(e)
			return
		}
	}
	log.Printf("[Error] file not found: %s", name)
}

func (f *Filestore) lock() {
	f.mutex.Lock()
}

func (f *Filestore) unlock() {
	f.mutex.Unlock()
}
