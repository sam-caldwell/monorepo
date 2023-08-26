package main

/*
 * locks command
 * (c) 2023 Sam Caldwell.  All Rights Reserved.
 *
 * This is a cross-platform lock tool for our Makefiles or anything really.
 * When executed, this tool will create a context-specific lock file.
 *
 * Note: By design this is a very trusting locking mechanism.  Do not use this
 *       for security purposes.
 *
 * lock create
 * 					<- creates a lock with the context and print the context string to stdout
 * 					<- if 'foo' exists, the program will fail (exit 1)
 * 				    <- if 'foo' does not exist, the lock is created (exit 0)
 *
 * lock check <contextId>
 *  				<- check to see if the lock exists:
 *								yes: exit 1 - locked,
 * 								 no: exit 0 - unlocked
 *
 * lock free <contextId>
 * 					<- delete the given context if it exists (exit 0)
 */
import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"github.com/sam-caldwell/go/v2/projects/go/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/go/fs/file"
	"github.com/sam-caldwell/go/v2/projects/go/lock"
	"github.com/sam-caldwell/go/v2/projects/go/misc/words"
	"os"
	"regexp"
	"strings"
)

const (
	cmdUsage = "\n\n" +
		"Usage:\n" +
		"\tlock %s            - creates a new lock file and returns its contextId\n" +
		"\tlock %s <contextId> - indicates if lock exists (exit code 0: no, exit code 1: yes)\n" +
		"\tlock %s <contextId>  - removes a given lock\n" +
		"\n"

	contextIdRegex = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	osArgCommand   = 1
	osArgContextId = 2

	lockExists = 1
	lockFree   = 0
)

func main() {
	exit.IfVersionRequested()
	exit.OnCondition(
		len(os.Args) < osArgCommand+1,
		exit.MissingArg,
		fmt.Sprintf(errors.MissingArgWithDetail, words.Command),
		fmt.Sprintf(cmdUsage, words.Create, words.Check, words.Free))

	getCheckContext := func() (string, error) {
		exit.OnCondition(
			len(os.Args) < osArgContextId+1,
			exit.MissingArg,
			errors.MissingContextId,
			fmt.Sprintf(cmdUsage, words.Create, words.Check, words.Free))
		contextId := strings.ToLower(strings.TrimSpace(os.Args[osArgContextId]))
		// Sanitize our contextId
		if regexp.MustCompile(contextIdRegex).MatchString(contextId) {
			return contextId, nil
		}
		return contextId, fmt.Errorf(errors.InvalidContextIdWithDetail, contextId)
	}

	var lock lock.File

	switch command := strings.ToLower(strings.TrimSpace(os.Args[osArgCommand])); command {
	case words.Create:
		contextId, err := lock.Create()
		if err != nil {
			exit.OnError(err, exit.LockCreateFailed, words.EmptyString)
		}
		fmt.Println(contextId)
	case words.Check:
		contextId, err := getCheckContext()
		if err != nil {
			exit.OnError(err, exit.InvalidInput, cmdUsage)
		}
		if err := lock.Check(contextId); err != nil {
			if err.Error() == errors.LockCheckFailed {
				exit.OnError(err, exit.GeneralError, cmdUsage)
			}
			os.Exit(lockExists)
		}
		os.Exit(lockFree)
	case words.Free:
		contextId, err := getCheckContext()
		if err != nil {
			exit.OnError(err, exit.InvalidInput, cmdUsage)
		}
		if err := lock.Free(contextId); err != nil {
			exit.OnError(err, exit.FailedToFreeLock, words.EmptyString)
		}

	default:
		exit.OnError(fmt.Errorf(errors.InvalidCommandWithDetail, command), exit.InvalidCommand, cmdUsage)
	}

	tempFilePath, err := file.CreateTempFile()
	if err != nil {
		fmt.Printf("Error creating temporary file: %s\n", err)
		return
	}

	fmt.Printf("Temporary file created: %s\n", tempFilePath.Name())
	defer func() {
		// Clean up the temporary file when done
		_ = os.Remove(tempFilePath.Name())
	}()
}
