/**
 * @name formatter/src/rtrimChar.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Formatter_TestRTrimChar_H
#define Formatter_TestRTrimChar_H

#include <string>

namespace Strings {

    /**
     * @name rtrim
     * @brief trim from the right end (in place)
     *
     * @param s string reference
     * @param removeChar char reference
     *
     */
    inline void rtrimChar(string *s, unsigned char *removeChar) {
        s->erase(std::find_if(s->rbegin(), s->rend(), [removeChar](unsigned char ch) {
            return ch != *removeChar;
        }).base(), s->end());
    }

    /**
     * @name rtrim
     * @brief trim from both ends (in place)
     *
     * @param s string value
     * @param removeChar char
     * @return string (trimmed)
     */
    inline string rtrimChar(string s, unsigned char removeChar) {
        rtrimChar(&s, &removeChar);
        return s;
    }

}/* Strings Namespace */
#endif /* Formatter_TestRTrimChar_H */