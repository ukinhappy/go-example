package main

import (
	"bytes"
	"fmt"
	"github.com/ledongthuc/pdf"
)

func main() {
	filepath := ""
	content, err := readPdf(filepath) // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	return
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
