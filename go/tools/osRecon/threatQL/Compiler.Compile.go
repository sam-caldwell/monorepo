package threatQL

import (
	"bufio"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
	"path/filepath"
	"strings"
)

// Compile - Compile the human-readable input from stdin until EOF or CRLF into binary form.
func (compiler *Compiler) Compile() error {
	var content []byte

	fileName := filepath.Join(compiler.queryDir.String(), fmt.Sprintf("*.%s", queryFileExtension))

	// Read lines from stdin until EOF or error
	if query, err := func() (string, error) {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		return scanner.Text(), scanner.Err()
	}(); err != nil {
		return err
	} else {

		if q, err := compileQuery(query); err != nil {
			return err
		} else {
			if content, err = q.Bytes(); err != nil {
				return err
			} else {
				// Write the line to the query file
				if err = os.WriteFile(fileName, content, 0600); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// compileQuery - given a query, return a compiled binary object
func compileQuery(query string) (q *Query, err error) {
	/*
	   	   	 * Text query string examples
	   	   	 *
	   	   	 *      cpu.loadavg1
	   	   	 *      cpu.loadavg5
	   	   	 *      cpu.loadavg15
	   	   	 *      cpu.user            -> return %user cpu load
	   	   	 *      cpu.wait            -> return wait time
	   	   	 *      cpu.idle            -> return idle time
	   	   	 *      cpu.cores           -> return number cores
	   	   	 *      mem.total           -> total memory (KB)
	            *      mem.free            -> free memory (KB)
	            *      mem.swap            -> virtual memory (KB)
	            *      process.pid,user,command|pid==3312
	*/
	parts := strings.Split(query, words.Space)
	return &(Query{}), err
}
