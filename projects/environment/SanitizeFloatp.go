package environment

import "fmt"

// SanitizeFloatp - get the float env var and bounds check the value
func SanitizeFloatp(name *string, min float64, max float64) (float64, error) {
	value, err := GetFloatp(name)
	if err != nil {
		return 0, err
	}
	if (value < min) || (value > max) {
		return 0, fmt.Errorf(errEnvBoundCheck)
	}
	return value, err
}
