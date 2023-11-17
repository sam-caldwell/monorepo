package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"os"
	"strings"
)

const svgTemplate = `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="20">
<linearGradient id="b" x2="0" y2="100%%">
	<stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
	<stop offset="1" stop-opacity=".1"/>
</linearGradient>
<mask id="a"><rect width="%d" height="20" rx="3" fill="#fff"/></mask>
	<g mask="url(#a)">
		<path fill="#555" d="M0 0h%dv20H0z"/>
		<path fill="%s" d="M%d 0h%dv20H%dz"/>
		<path fill="url(#b)" d="M0 0h%dv20H0z"/>
	</g>
	<g fill="#fff" text-anchor="middle"
		font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
		<text x="%f" y="15" fill="#010101" fill-opacity=".3">%s</text>
		<text x="%f" y="14">%s</text>
		<text x="%f" y="15" fill="#010101" fill-opacity=".3">%s</text>
		<text x="%f" y="14">%s</text>
	</g>
</svg>`

func main() {
	var name, status, color string
	exit.IfVersionRequested()
	flag.StringVar(&name, "name", "", "Name of the badge")
	flag.StringVar(&status, "status", "", "Status of the badge")
	flag.StringVar(&color, "color", "red", "Color of the badge")
	flag.Parse()

	if name == "" || status == "" {
		fmt.Println("Both 'name' and 'status' flags are required.")
		os.Exit(1)
	}

	fileName := fmt.Sprintf("badges/%s.svg", strings.ReplaceAll(name, " ", "_"))

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() { _ = f.Close() }()

	// calculate lengths and coordinates
	nameLength := len(name) * 8 // approximate text width unit
	statusLength := len(status) * 8
	width := nameLength + statusLength
	nameX := float64(nameLength) / 2.0
	statusX := float64(nameLength) + float64(statusLength)/2.0 + 2.0

	if color == "green" {
		color = "#4c1"
	}

	badge := fmt.Sprintf(svgTemplate,
		width, width, nameLength, color,
		nameLength, statusLength, nameLength,
		width, nameX, name, nameX, name,
		statusX, status, statusX, status)

	_, err = f.WriteString(badge)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("SVG badge created at %s\n", fileName)
}
