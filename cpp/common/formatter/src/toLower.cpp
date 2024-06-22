/**
 * @name formatter/src/toLower.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef toLower_H
#define toLower_H
/**
 * If USE_WIDE_CHAR_STRING is defined, configure <cwctype>.
 */
#ifdef USE_WIDE_CHAR_STRING

#include <cwctype>
#undef CONVERSION_FUNC
#define CONVERSION_FUNC(c) std::towlower(c);

#else
/**
 * Else If USE_WIDE_CHAR_STRING not defined, then configure <cctype>
 */
#ifdef LOCALE
#undef LOCALE
#endif /*ifdef LOCALE*/

#include <cctype>
#undef CONVERSION_FUNC
#define CONVERSION_FUNC(c) std::tolower(c);

#endif /*if USE_WIDE_CHAR_STRING...*/

#include <string>

using namespace std;

/**
 * @name toLower
 * @brief convert a string to upper case (uses wide char)
 *
 * @param s string
 * @return string (lowercase version of s)
 */
string toLower(string s) {
#ifdef LOCALE
    setlocale(LC_ALL, LOCALE);
#endif /*LOCALE*/
    transform(s.begin(), s.end(), s.begin(), [](unsigned char c) {
        return CONVERSION_FUNC(c)
    });
    return s;
}

#endif /* toLower_H */