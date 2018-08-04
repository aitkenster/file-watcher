package filestore

import (
	"os"
	"sync"
)

type Store struct {
	list  map[string]struct{}
	mutex sync.Mutex
}

type fileList map[string]struct{}

func New() *Store {
	return &Store{
		list:  fileList{},
		mutex: sync.Mutex{},
	}
}

func (s *Store) AddFiles(files []os.FileInfo) {
	for _, file := range files {
		s.list[file.Name()] = struct{}{}
	}
}

func (s *Store) Update(op string, filename string) {
	s.lock()
	switch op {
	case "add":
		s.list[filename] = struct{}{}
	case "remove":
		delete(s.list, filename)
	}
	s.unlock()
}

func (s *Store) GetList() fileList {
	return s.list
}

func (s *Store) lock() {
	s.mutex.Lock()
}

func (s *Store) unlock() {
	s.mutex.Unlock()
}
