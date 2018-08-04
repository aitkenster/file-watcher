package lib

type FileMetadata struct {
	Name string `json:"name"`
}

type PatchOperation struct {
	Op    string       `json:"op"`
	Path  string       `json:"path"`
	Value FileMetadata `json:"value"`
}
