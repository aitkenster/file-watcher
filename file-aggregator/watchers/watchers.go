package watchers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aitkenster/file-watcher/file-aggregator/lib"
)

type directoryResponse struct {
	Files []lib.FileMetadata `json:"files"`
}

type WatcherConfig struct {
	baseUrls []string
}

func NewConfig() *WatcherConfig {
	baseUrls := strings.Split(os.Getenv("WATCHER_ADDRESSES"), ",")
	if baseUrls[0] == "" {
		log.Fatal("[ERROR] no WATCHER_ADDRESSES env var set")
	}

	return &WatcherConfig{
		baseUrls: baseUrls,
	}
}

func (wc *WatcherConfig) GetDirectoryFiles() ([]lib.FileMetadata, error) {
	fl := []lib.FileMetadata{}
	for _, url := range wc.baseUrls {
		resp, err := http.Get(url + "/directory")
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
		fl = append(fl, unmarshalledBody.Files...)
	}
	return fl, nil
}
