//«Chill» © Nikita Nikonov, 2021
//License: Have it in all the holes, for free

//About: Chilling static articles generator/compiler, written in a few hours
//Performance: it does 100 articles (1500 words / 10000 symb. without MD markup) in ~50ms
//Thank's to Igor Adamenko (https://igoradamenko.github.io/awsm.css/) & community that created "github.com/gomarkdown/markdown" package

//How to use: put some markdown files into input/, run program and get index.html with all articles in one website.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gomarkdown/markdown"
)

func main() {

	//future result/stat values
	start := time.Now()
	var processedNumber int

	//create file in which will be written content
	f, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//some styles
	f.WriteString(`
		<link rel="stylesheet" href="https://igoradamenko.github.io/awsm.css/css/awsm_theme_mischka.min.css">
	`)

	//walk through current directory & check each file
	err = filepath.Walk("input/", func(path string, info os.FileInfo, err error) error {

		//check if file extension is .md
		if matched, err := filepath.Match("*.md", filepath.Base(path)); err != nil {
			return err
		} else if matched {

			//read file
			md, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal("Fucked up reading file: ", err)
			}

			//translate content to html
			output := markdown.ToHTML([]byte(md), nil, nil)

			//write html into file
			_, err = f.WriteString("<article>")
			_, err = f.Write(output)
			_, err = f.WriteString("</article>")

			if err != nil {
				log.Fatal("Can't write.")
			}

			//increace processed files counter
			processedNumber++

		}
		return nil

	})
	if err != nil {
		log.Fatal("Shit happens.")
	}

	//beautiful results
	fmt.Print("--- \n Success. \n Processed " + fmt.Sprint(processedNumber) + " files in " + fmt.Sprint(time.Since(start).Milliseconds()) + "ms. \n---")

	//waitin' for user to quit
	var input string
	fmt.Scanf("%v", &input)

}
