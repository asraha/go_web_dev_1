package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseGlob("templates/*")

	if err != nil {

		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)

	if err != nil {

		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "file-1.gohtml", nil)
	if err != nil {

		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "file-2.gohtml", nil)
	if err != nil {

		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "file-3.gohtml", nil)
	if err != nil {

		log.Fatalln(err)
	}

}
