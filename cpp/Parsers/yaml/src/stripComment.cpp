/**
 * Parsers/yaml/src/Parsers.h/stripComment.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Parser_stripComment_H
#define Parser_stripComment_H

#include <string>
#include "../../../formatter/formatter.h"

/**
 * @name stripComment
 * @brief strip the comments from the current line.
 *
 * @param *line string (by ref)
 * @return bool
 */
inline bool stripComment(string *line) {
    std::size_t found = line->find("#");
    if (found != std::string::npos) {
        line->erase(found, line->length());
        Strings::trim(line);
        return true;
    }
    return false;
}

#endif /* Parser_stripComment_H */