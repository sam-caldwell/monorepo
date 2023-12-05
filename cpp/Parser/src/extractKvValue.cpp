/**
 * @name Parser/src/extractKvValue.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Parser_extractKvValue_H
#define Parser_extractKvValue_H

#include <string>
#include <vector>
#include "projects/common/formatter/formatter.h"

/**
 * @name extractKvValue
 * @brief Given an input in the format "identifier: value", this function will
 *        strip everything to the left of ": " and return what is to the right
 *        as a value.
 *
 * @param *line string pointer
 * @return string (stripped value)
 */
inline string extractKvValue(string *line) {
    std::size_t startPos = line->find_first_of(":");
    if (startPos != string::npos)
        return line->substr(startPos + 2, string::npos);
    return "";
}

#endif /* Parser_extractKvValue_H */