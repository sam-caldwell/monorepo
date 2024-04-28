package threatQL

import (
	"github.com/sam-caldwell/monorepo/go/fs/file"
)

// QueryFile -  The Data structure used to represent a Query file in memory.
//
//	This is mostly intended to be used on the server side to store
//	queries for processing.
//
// ToDo: Query files should b encrypted for the target client using GPG
//
//	so that only the target can decrypt the file and execute the query.
type QueryFile struct {
	query    Query
	fileName *file.File
}
