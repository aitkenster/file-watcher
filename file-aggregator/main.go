package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// get all files from each node and store in list
	files := newFiles()

	// set up listeners to each node
	http.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		filesRequestHandler(w, r, files)
	})

	//port := os.Getenv("PORT")
	port := "9090"
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

	// insert new file into list
}

type filesPatchRequest []patchOperation

// Op should be enum
type patchOperation struct {
	Op    string       `json:"op"`
	Path  string       `json:"path"`
	Value fileMetadata `json:"value"`
}

type fileMetadata struct {
	Name string `json:"name"`
}

func filesRequestHandler(w http.ResponseWriter, r *http.Request, files *files) {
	if !(r.Method == http.MethodPatch) {
		log.Println("[ERROR] invalid request method :", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR]", err)
	}

	var req filesPatchRequest
	json.Unmarshal(body, &req)

	for _, op := range req {
		files.modifyList(op)
	}

	files.printList()

	log.Println(req)
}
