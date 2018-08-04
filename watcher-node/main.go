package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"

	"github.com/aitkenster/file-watcher/watcher-node/aggregator"
	"github.com/aitkenster/file-watcher/watcher-node/filestore"
	"github.com/aitkenster/file-watcher/watcher-node/server"
)

const (
	mountedDir = "/host/watched-folder"
	add        = "add"
	remove     = "remove"
)

func main() {
	var directory = flag.String("dir", mountedDir, "the path of the directory to watch")
	flag.Parse()

	aggregatorClient := aggregator.New(&http.Client{})

	store, err := initializeStoreForDirectory(*directory)
	if err != nil {
		log.Fatal("[ERROR]", err)
	}

	http.HandleFunc("/directory", server.DirectoryHandler(store))

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("[ERROR]", err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				handleEvent(event, store, aggregatorClient)
			case err := <-watcher.Errors:
				log.Println("[ERROR]", err)
			}
		}
	}()

	if err := watcher.Add(*directory); err != nil {
		log.Println("[ERROR]", err)
	}

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func initializeStoreForDirectory(directory string) (*filestore.Store, error) {
	store := filestore.New()

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	store.AddFiles(files)
	return store, nil
}

func handleEvent(
	event fsnotify.Event,
	store *filestore.Store,
	aggregator *aggregator.Aggregator,
) {
	op := getOp(event)
	if op == "" {
		return
	}
	filename := filepath.Base(event.Name)

	store.Update(op, filename)

	err := aggregator.NotifyUpdate(op, filename)
	if err != nil {
		log.Println("[ERROR]: ", err)
	}
}

func getOp(event fsnotify.Event) string {
	switch event.Op {
	case fsnotify.Create:
		return add
	case fsnotify.Remove, fsnotify.Rename:
		return remove
	}
	return ""
}
