package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aitkenster/file-watcher/watcher-node/filestore"
	"github.com/aitkenster/file-watcher/watcher-node/lib"
)

type listResponse struct {
	Files []lib.FileMetadata `json:"files"`
}

func DirectoryHandler(store *filestore.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !(r.Method == http.MethodGet) {
			log.Println("[ERROR] invalid request method :", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		filesMeta := []lib.FileMetadata{}
		for name, _ := range store.GetList() {
			filesMeta = append(filesMeta, lib.FileMetadata{
				Name: name,
			})
		}
		json.NewEncoder(w).Encode(listResponse{
			Files: filesMeta,
		})
	})
}
