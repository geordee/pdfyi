package app

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/geordee/pdfyi/util"
	"golang.org/x/net/html"
)

// Generate API
func Generate(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.BadRequest(&w, "bad_request")
		return
	}
	defer r.Body.Close()

	// Ensure a valid HTML body
	reader := strings.NewReader(string(data))
	tokenizer := html.NewTokenizer(reader)
	for {
		token := tokenizer.Next()
		if token == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			util.BadRequest(&w, "bad_html")
			return
		}
	}

	pdfg, err := pdf.NewPDFGenerator()
	if err != nil {
		log.Println("Error in creating PDF generator")
		log.Println(err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	pdfg.AddPage(pdf.NewPageReader(strings.NewReader(string(data))))
	err = pdfg.Create()
	if err != nil {
		log.Println("Error in creating PDF")
		log.Println(err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(pdfg.Bytes())
}
