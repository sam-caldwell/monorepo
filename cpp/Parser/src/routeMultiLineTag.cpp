/**
 * @name Parser/src/routeMultiLineTag.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <sstream>

/**
 * @name routeMultiLineTag
 * @brief We have identified a multi-line tag ("identifier: |") which we will consume and add to our
 *        tags collection.  Note that we do not store a Node to our tokens collection here because we
 *        only know the tag, and we do not know the value.
 *
 * @param *line string pointer
 * @return bool
 */
inline bool Parser::routeMultiLineTag(string *line) {
    tagIndent.push(indent.current());
    tags.push(extractTagOnly(line));

    unsigned i;
    unsigned lastIndent = 0;
    stringstream lineBuffer;

    while (!file.eof()) {
        getline(file, *line);
        i = getIndentation(line);
        Strings::trim(line);
        if (i < lastIndent) {
            break;
        } else {
            lineBuffer << *line << " ";
        }
        lastIndent = i;
    }
    string temp = lineBuffer.str();
    Strings::trim(&temp);

    tokens.insert(pair<string, string>(tags.toString(), temp));

    tagIndent.pop();
    indent.pop();
    tags.pop();
    getNextLine = false;
    expectedToken = Any;
    return true;
}