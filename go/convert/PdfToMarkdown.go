package convert

import (
	"fmt"
	"github.com/dslipak/pdf"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
)

// PdfToMarkdown - reads a PDF file and converts its text to Markdown format
//
//	(c) 2023 Sam Caldwell.  MIT License
func PdfToMarkdown(pdfPath string, mdPath string) (err error) {
	var r *os.File
	var pdfReader *pdf.Reader
	var mdContent string
	// Open the PDF file
	if r, err = os.Open(pdfPath); err != nil {
		return fmt.Errorf("failed to open PDF file: %v", err)
	}
	defer func() { _ = r.Close() }()

	// Read the PDF content
	if pdfReader, err = pdf.NewReader(r, file.GetFileSize(r)); err != nil {
		return fmt.Errorf("failed to create PDF reader: %v", err)
	}
	// Iterate over the pages
	for pageIndex := 1; pageIndex <= pdfReader.NumPage(); pageIndex++ {
		var page pdf.Page
		var content string
		if page = pdfReader.Page(pageIndex); page.V.IsNull() {
			continue
		}
		// Extract text from the page
		if content, err = page.GetPlainText(nil); err != nil {
			return fmt.Errorf("failed to extract text from page %d: %v", pageIndex, err)
		}
		mdContent += content + "\n\n"
	}
	// Write the Markdown content to a file
	if err = os.WriteFile(mdPath, []byte(mdContent), 0644); err != nil {
		return fmt.Errorf("failed to write Markdown file: %v", err)
	}
	return err
}
