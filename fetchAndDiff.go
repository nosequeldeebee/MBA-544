package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ledongthuc/pdf"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {
	// URL of the PDF file
	pdfURL := "example.pdf"

	// Fetch the PDF
	resp, err := http.Get(pdfURL)
	if err != nil {
		fmt.Println("Error fetching PDF:", err)
		return
	}
	defer resp.Body.Close()

	// Create a temporary file to store the PDF
	tmpfile, err := os.CreateTemp("", "temp-*.pdf")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer os.Remove(tmpfile.Name())

	// Copy the PDF content to the temporary file
	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		fmt.Println("Error writing to temp file:", err)
		return
	}

	// Close the temp file before reading it
	tmpfile.Close()

	// Read the PDF file
	content, err := readPdf(tmpfile.Name())
	if err != nil {
		fmt.Println("Error reading PDF:", err)
		return
	}

	// URL of the PDF file
	pdfURL2 := "example2.pdf"

	// Fetch the PDF
	resp, err = http.Get(pdfURL2)
	if err != nil {
		fmt.Println("Error fetching PDF:", err)
		return
	}
	defer resp.Body.Close()

	// Create a temporary file to store the PDF
	tmpfile, err = os.CreateTemp("", "temp-*.pdf")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer os.Remove(tmpfile.Name())

	// Copy the PDF content to the temporary file
	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		fmt.Println("Error writing to temp file:", err)
		return
	}

	// Close the temp file before reading it
	tmpfile.Close()

	// Read the PDF file
	content2, err := readPdf(tmpfile.Name())
	if err != nil {
		fmt.Println("Error reading PDF:", err)
		return
	}

	// Create a new diff-match-patch object
	dmp := diffmatchpatch.New()

	// Calculate the diff
	diffs := dmp.DiffMain(content, content2, false)

	// Print the diff in a human-readable format
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffEqual:
			fmt.Print(diff.Text)
		case diffmatchpatch.DiffInsert:
			fmt.Printf("\033[32m%s\033[0m", diff.Text) // Green for insertions
		case diffmatchpatch.DiffDelete:
			fmt.Printf("\033[31m%s\033[0m", diff.Text) // Red for deletions
		}
	}
	fmt.Println()

}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var content string
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		text, err := p.GetPlainText(nil)
		if err != nil {
			return "", err
		}
		content += text
	}

	return content, nil
}
