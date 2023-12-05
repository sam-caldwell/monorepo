/**
 * @name AnsiColors/AnsiColors.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @note see https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
 */
#ifndef ANSICOLORS_H
#define ANSICOLORS_H

#define LINE_FEED  "\n"
#define SPACE      " "
#define	TAB        "\t"

//General
#define	RESET  "\033[0m"
#define	CLEAR  "\033[2J"

// Formatting
#define	BOLD           "\033[1m" // Set bold text
#define	DIM            "\033[2m" // Set dim (decreased intensity) text
#define	UNDERLINE      "\033[4m" // Set underlined text
#define	BLINK          "\033[5m" // Set blinking text (not widely supported)
#define	REVERSE        "\033[7m" // Set reverse video (swap foreground and background color)
#define	HIDDEN_TEXT    "\033[8m" // Set hidden text (invisible)
#define	STRIKETHROUGH  "\033[9m" // Set strikethrough text (not widely supported)

//Background Colors
#define	BG_BLACK    "\033[40m"
#define	BG_RED      "\033[41m"
#define	BG_GREEN    "\033[42m"
#define	BG_YELLOW   "\033[43m"
#define	BG_BLUE     "\033[44m"
#define	BG_MAGENTA  "\033[45m"
#define	BG_CYAN     "\033[46m"
#define	BG_WHITE    "\033[47m"

//Foreground Colors
#define	FG_BLACK    "\033[30m"
#define	FG_RED      "\033[31m"
#define	FG_GREEN    "\033[32m"
#define	FG_YELLOW   "\033[33m"
#define	FG_BLUE     "\033[34m"
#define	FG_MAGENTA  "\033[35m"
#define	FG_CYAN     "\033[36m"
#define	FG_WHITE    "\033[37m"

//Navigation
// Move the cursor up by 2 rows
#define	SET_TOP_LEFT  "\033[H"
#define	MOVE_UP       "\033[%dA"
#define	MOVE_DOWN     "\033[%dB"
#define	MOVE_RIGHT    "\033[%dC"
#define	MOVE_LEFT     "\033[%dD"

#endif /*ANSICOLORS_H*/
