package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aitkenster/file-watcher/file-aggregator/files"
	"github.com/aitkenster/file-watcher/file-aggregator/lib"
)

type filesPatchRequest []lib.PatchOperation

type listFilesResponse struct {
	Items []lib.FileMetadata `json:"items"`
}

func FilesRequestHandler(w http.ResponseWriter, r *http.Request, files *files.Filestore) {
	switch r.Method {
	case http.MethodPatch:
		handlePatchFiles(w, r, files)
	case http.MethodGet:
		handleGetFiles(w, r, files)
	default:
		log.Println("[ERROR] invalid request method :", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func handlePatchFiles(w http.ResponseWriter, r *http.Request, files *files.Filestore) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR]", err)
	}

	var req filesPatchRequest
	json.Unmarshal(body, &req)

	for _, op := range req {
		files.ModifyList(op)
	}

	log.Println(req)
}

func handleGetFiles(w http.ResponseWriter, r *http.Request, files *files.Filestore) {
	response := listFilesResponse{
		Items: files.RetrieveAll(),
	}
	json.NewEncoder(w).Encode(response)
}
