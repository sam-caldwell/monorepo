package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/convert"
	"github.com/sam-caldwell/monorepo/projects/exit"
	"github.com/sam-caldwell/monorepo/projects/fs/file"
	"github.com/sam-caldwell/monorepo/projects/misc/words"
	"github.com/sam-caldwell/monorepo/projects/version"
	"strings"
)

func main() {
	const (
		pdfExtension = ".pdf"
		mdExtension  = ".md"
	)
	v := flag.Bool("version", false, "show version")
	pdfFile := flag.String("in", "", "Pdf path/file")
	mdFile := flag.String("out", "", "output file (markdown)")
	flag.Parse()
	if *v {
		version.Show()
	}
	if strings.TrimSpace(*pdfFile) == words.EmptyString {
		ansi.Red().Println("-in is required (cannot be empty)").Fatal(exit.InvalidInput).Reset()
		return
	}
	if !file.Existsp(pdfFile) {
		ansi.Red().Printf("File not found (%s)", *pdfFile).LF().Reset()
		return
	}
	if strings.TrimSpace(*mdFile) == words.EmptyString {
		ansi.Blue().
			Printf("-out not specified.  Appending %s to %s as the output file\n", mdExtension, *pdfFile).
			Reset()
		if strings.HasSuffix(*pdfFile, pdfExtension) {
			*mdFile = strings.TrimSuffix(*pdfFile, pdfExtension)
		}
		*pdfFile += mdExtension
	}

	if err := convert.PdfToMarkdown(*pdfFile, *mdFile); err != nil {
		ansi.Red().Printf("Error converting PDF to Markdown: %v\n", err).Reset()
		return
	}
	ansi.Green().Println("Conversion completed successfully.").Reset()
}
