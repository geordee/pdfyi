package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	app "github.com/geordee/pdfyi/app"
	config "github.com/geordee/pdfyi/config"
)

func main() {
	config.InitializeApp()

	if config.App.AllowInsecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	http.HandleFunc("/pdfs", app.Generate)
	http.HandleFunc("/", app.Index)

	url := fmt.Sprintf(":%d", config.App.ListenPort)
	log.Println("Starting server at " + url)
	log.Fatalln(http.ListenAndServe(url, nil))
}
