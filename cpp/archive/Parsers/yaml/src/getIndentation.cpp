/**
 * Parsers/yaml/src/Parsers.h/getIndentation.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Parser_GetIndentation_H
#define Parser_GetIndentation_H

/**
 * @name getIndentation
 * @brief count number of spaces in the current line
 *        from column 0 to the first non-whitespace char,
 *        then update the internal class state.
 *
 * @param *text string pointer
 * @return int (indentation character count)
 */
inline int getIndentation(string *text) {
    int i = 0;
    while ((*text)[i] == SPACE) i++;
    return i;
}

#endif /* Parser_GetIndentation_H */
