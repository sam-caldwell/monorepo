package monorepo

import (
	"bufio"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os"
	"strings"
)

func LoadManifest(manifest *string) (projects map[string]string, err error) {
	projects = make(map[string]string)
	f, err := os.Open(*manifest)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "#") {
			continue
		}
		project, language := func() (p string, l string) {
			//
			//project:language #comment string is ignored
			//
			parts := strings.Split(text, "#")
			p = parts[0]
			parts = strings.Split(p, ":")
			p = parts[0]
			l = parts[1]
			return p, l
		}()
		ansi.Yellow().
			Println(project).
			Println(language).
			Reset()
	}
	return
}
