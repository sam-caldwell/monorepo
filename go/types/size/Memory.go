package size

type Memory uint

const (
	/*
	 * Memory size coefficients
	 */
	KB Memory = 1024
	MB Memory = 1024 * 1024
	GB Memory = 1024 * 1024 * 1024
	TB Memory = 1024 * 1024 * 1024 * 1024
	PB Memory = 1024 * 1024 * 1024 * 1024 * 1024
)
