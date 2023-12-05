/**
 * @name formatter/src/rtrim.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Formatter_TestRTrim_H
#define Formatter_TestRTrim_H

#include <string>

namespace Strings {

    /**
     * @name rtrim
     * @brief trim from the right end (in place)
     *
     * @param s string reference
     */
    void rtrim(string *s) {
        s->erase(std::find_if(s->rbegin(), s->rend(), [](unsigned char ch) {
            return !std::isspace(ch);
        }).base(), s->end());
    }

    /**
     * @name rtrim
     * @brief trim from both ends (in place)
     *
     * @param s string value
     * @return string (trimmed)
     */
    inline string rtrim(string s) {
        rtrim(&s);
        return s;
    }
}/* Strings Namespace */
#endif /* Formatter_TestRTrim_H */