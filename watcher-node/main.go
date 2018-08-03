package main

import (
	"flag"
	"io/ioutil"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
)

const mountedDir = "/host/watched-folder"

func main() {
	var directory = flag.String("dir", mountedDir, "the path of the directory to watch")
	flag.Parse()

	log.Println("[INFO] Retrieving file list")
	files, err := ioutil.ReadDir(*directory)
	if err != nil {
		log.Fatal("[ERROR]", err)
	}

	var mutex sync.Mutex
	filestore := filestore{}

	filestore.addFiles(files)

	log.Println("[INFO]", filestore)

	log.Println("[INFO] Starting up file watcher")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("[ERROR]", err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				mutex.Lock()
				filestore.processEvent(event)
				mutex.Unlock()
				log.Println("[INFO]", filestore)

			case err := <-watcher.Errors:
				log.Println("[ERROR]", err)
			}
		}
	}()

	if err := watcher.Add(*directory); err != nil {
		log.Println("[ERROR]", err)
	}

	<-done
}
