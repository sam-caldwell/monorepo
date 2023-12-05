/**
 * @name formatter/src/formatter.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef formatter_h
#define formatter_h
/**
 * Select whether we use wide-char strings (e.g. en_US.utf8).
 */
#define USE_WIDE_CHAR_STRING true

/**
 * define our default locale
 */
#ifdef USE_WIDE_CHAR_STRING

#ifdef LOCALE
#undef LOCALE
#endif /* ifdef LOCALE*/

#define LOCALE "en_US.utf8"

#endif /* ifdef USE_WIDE_CHAR_STRING */

/**
 * Make sure CONVERSION_FUNC is undefined
 */
#ifdef CONVERSION_FUNC
#undef CONVERSION_FUNC
#endif
/**
 * Include the formatter project files.
 */
#include "src/intToString.cpp"
#include "src/toLower.cpp"
#include "src/toUpper.cpp"
#include "src/ltrim.cpp"
#include "src/rtrim.cpp"
#include "src/rtrimChar.cpp"
#include "src/trim.cpp"
#include "src/setToString.cpp"
#include "src/StringToSet.cpp"

#endif /*formatter_h*/