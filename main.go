package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
)

func main() {

	type Event struct {
		Date     string
		Header   string
		Markdown string
		HTML     template.HTML
	}

	// Load data.json
	type Config struct {
		Avatar  string
		Name    string
		About   string
		Contact string
		GitHub  string
		Social  []map[string]string
		Events  []Event
	}

	// Load config
	rawConfig, _ := ioutil.ReadFile("config.json")
	cfg := Config{}
	err := json.Unmarshal([]byte(rawConfig), &cfg)
	if err != nil {
		fmt.Print("Please, fill config.json. More info in README.md: ", err.Error())
		os.Exit(0)
	}

	// Process MD
	for i := 0; i < len(cfg.Events); i++ {
		f, err := ioutil.ReadFile("markdown/" + cfg.Events[i].Markdown)
		cfg.Events[i].HTML = template.HTML(string(markdown.ToHTML(f, nil, nil)))
		if err != nil {
			cfg.Events[i].HTML = template.HTML(string(markdown.ToHTML([]byte(cfg.Events[i].Markdown), nil, nil)))
		}
	}

	// Render layout (index)
	ts, err := template.ParseFiles("layout/index.html")
	if err != nil {
		fmt.Print("Your index.html can't be parsed: ", err.Error())
		os.Exit(0)
	}
	f, err := os.Create("public/index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = ts.Execute(f, cfg)
	if err != nil {
		log.Println(err.Error())
	}

	// Copy static files to public and add rendered index.html
	filepath.Walk("layout/", func(wPath string, info os.FileInfo, err error) error {
		os.Link("layout/"+wPath, "public/"+wPath)
		return nil
	})

	// Push public dir to gh-pages branch

}

// TODO: (you can help developing)

// Write err handlers
// Write tests
// Add configuration of input/output dirs in data.json
