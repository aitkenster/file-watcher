package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aitkenster/file-watcher/file-aggregator/files"
	server "github.com/aitkenster/file-watcher/file-aggregator/server"
	"github.com/aitkenster/file-watcher/file-aggregator/watchers"
)

func main() {
	var store *files.Filestore

	http.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		if store != nil {
			server.FilesRequestHandler(w, r, store)
		}
	})

	watcherConfig := watchers.NewConfig()
	directoryFiles, err := watcherConfig.GetDirectoryFiles()
	if err != nil {
		log.Println("[ERROR]: ", err)
		// probably need a kill or retry here
	}

	store = files.New(directoryFiles)

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
