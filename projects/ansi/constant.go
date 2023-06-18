package ansi

const (
	LineFeed = "\n"
	Space    = " "
	Tab      = "\t"

	//General
	reset = "\033[0m"
	clear = "\033[2J"

	// Formatting
	bold          = "\033[1m" // Set bold text
	dim           = "\033[2m" // Set dim (decreased intensity) text
	underline     = "\033[4m" // Set underlined text
	blink         = "\033[5m" // Set blinking text (not widely supported)
	reverse       = "\033[7m" // Set reverse video (swap foreground and background colors)
	hiddenText    = "\033[8m" // Set hidden text (invisible)
	strikethrough = "\033[9m" // Set strikethrough text (not widely supported)

	//Background Colors
	bgBlack   = "\033[40m"
	bgRed     = "\033[41m"
	bgGreen   = "\033[42m"
	bgYellow  = "\033[43m"
	bgBlue    = "\033[44m"
	bgMagenta = "\033[45m"
	bgCyan    = "\033[46m"
	bgWhite   = "\033[47m"

	//Foreground Colors
	fgBlack   = "\033[30m"
	fgRed     = "\033[31m"
	fgGreen   = "\033[32m"
	fgYellow  = "\033[33m"
	fgBlue    = "\033[34m"
	fgMagenta = "\033[35m"
	fgCyan    = "\033[36m"
	fgWhite   = "\033[37m"

	//Navigation
	// Move the cursor up by 2 rows
	setTopLeft = "\033[H"
	moveUp     = "\033[%dA"
	moveDown   = "\033[%dB"
	moveRight  = "\033[%dC"
	moveLeft   = "\033[%dD"
)
