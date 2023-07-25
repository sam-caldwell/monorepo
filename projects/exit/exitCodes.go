package exit

/*
 * projects/exit/exitCodes.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines our program exit codes.
 *
 * See README.md
 */

const (
	Success        = 0
	GeneralError   = 1
	InvalidCommand = 2
	InvalidInput   = 3

	MissingArg     = 10
	UnknownCommand = 11
	MissingColor   = 13

	NotFound = 20

	LockCreateFailed = 98
	FailedToFreeLock = 99
)
