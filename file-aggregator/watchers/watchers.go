package watchers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aitkenster/file-watcher/file-aggregator/lib"
)

type directoryResponse struct {
	Files []lib.FileMetadata `json:"files"`
}

func GetDirectoryFiles() ([]lib.FileMetadata, error) {
	resp, err := http.Get("http://localhost:6060/directory")
	if err != nil {
		log.Println("[ERROR]", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var unmarshalledBody directoryResponse
	err = json.Unmarshal(body, &unmarshalledBody)
	if err != nil {
	}
	return unmarshalledBody.Files, nil
}
