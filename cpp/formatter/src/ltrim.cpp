/**
 * @name formatter/src/ltrim.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Formatter_TestLTrim_H
#define Formatter_TestLTrim_H

#include <string>

namespace Strings {

    /**
     * @name ltrim
     * @brief trim from the left end (in place)
     *
     * @param s string reference
     */
    static inline void ltrim(std::string *s) {
        s->erase(s->begin(), std::find_if(s->begin(), s->end(), [](unsigned char ch) {
            return !std::isspace(ch);
        }));
    }

    /**
     * @name ltrim
     * @brief trim from both ends (in place)
     *
     * @param s string value
     * @return string (trimmed)
     */
    static inline string ltrim(std::string s) {
        ltrim(&s);
        return s;
    }
}/* Strings Namespace */
#endif /* Formatter_TestLTrim_H */