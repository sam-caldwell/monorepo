package Postgres

import (
	"regexp"
)

func SanitizeDBName(dbName *string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{1,32}$`)
	return pattern.MatchString(*dbName)
}
