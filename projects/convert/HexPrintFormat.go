package convert

import "fmt"

func HexPrintFormat(input []byte, width int) (out string) {
	for pos, b := range input {
		if pos%width == 0 {
			out += "\n"
		}
		out += fmt.Sprintf("%02x ", b)
	}
	return out
}
