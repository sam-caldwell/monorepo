package environment

import (
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"strconv"
	"strings"
)

// GetFloatp - Return the floating-point value of a given environment variable
func GetFloatp(name *string) (r float64, e error) {
	v := os.Getenv(*name)
	if strings.TrimSpace(v) == "" {
		return defaultFloatValue, nil
	}
	return strconv.ParseFloat(v, floatVarSize)
}
