/**
 * @name formatter/src/trim.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Formatter_TestTrim_H
#define Formatter_TestTrim_H

#include <string>
#include "ltrim.cpp"
#include "rtrim.cpp"

namespace Strings {
    /**
     * @name trim
     * @brief trim from both ends (in place)
     *
     * @param s string reference
     */
    static inline void trim(std::string *s) {
        Strings::ltrim(s);
        Strings::rtrim(s);
    }

    /**
     * @name trim
     * @brief trim from both ends (in place)
     *
     * @param s string value
     * @return string (trimmed)
     */
    static inline string trim(std::string s) {
        Strings::trim(&s);
        return s;
    }
}/* Strings Namespace */
#endif /* Formatter_TestTrim_H */