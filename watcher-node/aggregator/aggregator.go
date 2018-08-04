package aggregator

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/aitkenster/file-watcher/watcher-node/lib"
)

type Aggregator struct {
	baseUrl string
	client  *http.Client
}

type patchOperation struct {
	Op    string           `json:"op"`
	Path  string           `json:"path"`
	Value lib.FileMetadata `json:"value"`
}

func New(baseUrl string, client *http.Client) *Aggregator {
	return &Aggregator{
		baseUrl: baseUrl,
		client:  client,
	}
}

func (ag *Aggregator) NotifyUpdate(op string, filename string) error {
	body := []patchOperation{
		patchOperation{
			Op:   op,
			Path: "/files",
			Value: lib.FileMetadata{
				Name: filename,
			},
		},
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPatch,
		ag.baseUrl+"/files",
		bytes.NewReader(payload),
	)

	if err != nil {
		return err
	}

	_, err = ag.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
