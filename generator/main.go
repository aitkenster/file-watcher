package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type EnvConfig struct {
	WatchedFolders []string
}

func newConfig() *EnvConfig {
	folders := strings.Split(os.Getenv("WATCHED_FOLDERS"), ",")
	if folders[0] == "" {
		log.Fatal("[ERROR] no WATCHED_FOLDERS env var set")
	}
	return &EnvConfig{
		WatchedFolders: folders,
	}
}

func (cfg *EnvConfig) Transform() {
	if len(cfg.WatchedFolders) == 1 {
		cfg.WatchedFolders = strings.Split(cfg.WatchedFolders[0], ":")
	}
}

func main() {
	config := newConfig()
	config.Transform()

	// load template from file and add function to it
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}
	tmpl, err := template.New("docker-compose.tmpl").Funcs(funcMap).ParseFiles("./generator/docker-compose.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("./.docker-compose-gen.yml")
	if err != nil {
		log.Println("[INFO] create file: ", err)
		return
	}
	err = tmpl.Execute(f, config.WatchedFolders)
	if err != nil {
		log.Fatal(err)
	}
}
