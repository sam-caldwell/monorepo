package ansi

//	 Ansi Color code constants
//
//		(c) 2023 Sam Caldwell.  MIT License
const (
	LineFeed = "\n"
	Space    = " "
	Tab      = "\t"

	//CodeReset - This ansi code resets the system to the default state
	CodeReset = "\033[0m"
	//CodeClear - This ansi code clears a screen
	CodeClear = "\033[2J"

	//CodeBold - print text in bold
	CodeBold          = "\033[1m" // Set CodeBold text
	CodeDim           = "\033[2m" // Set CodeDim (decreased intensity) text
	CodeUnderline     = "\033[4m" // Set underlined text
	CodeBlink         = "\033[5m" // Set blinking text (not widely supported)
	CodeReverse       = "\033[7m" // Set CodeReverse video (swap foreground and background color)
	CodeHiddenText    = "\033[8m" // Set hidden text (invisible)
	CodeStrikeThrough = "\033[9m" // Set CodeStrikeThrough text (not widely supported)

	//CodeBgBlack - Black background
	CodeBgBlack   = "\033[40m"
	CodeBgRed     = "\033[41m"
	CodeBgGreen   = "\033[42m"
	CodeBgYellow  = "\033[43m"
	CodeBgBlue    = "\033[44m"
	CodeBgMagenta = "\033[45m"
	CodeBgCyan    = "\033[46m"
	CodeBgWhite   = "\033[47m"

	//CodeFgBlack - Black Foreground
	CodeFgBlack   = "\033[30m"
	CodeFgRed     = "\033[31m"
	CodeFgGreen   = "\033[32m"
	CodeFgYellow  = "\033[33m"
	CodeFgBlue    = "\033[34m"
	CodeFgMagenta = "\033[35m"
	CodeFgCyan    = "\033[36m"
	CodeFgWhite   = "\033[37m"

	CodeSetTopLeft = "\033[H"
	CodeMoveUp     = "\033[%dA"
	CodeMoveDown   = "\033[%dB"
	CodeMoveRight  = "\033[%dC"
	CodeMoveLeft   = "\033[%dD"
)
