/**
 * @name Parser/src/routeTagMap.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name routeTagMap
 * @brief We have detected a map ("map:") and we expect a child node or sequence.
 *        so all we will do here is append the tag to our tags collection.
 *        There is no value yet.
 *
 * @return bool
 */
inline bool Parser::routeTagMap() {
    unsigned i = getIndentation(&line);
    if (i == 0) {
        tags.clear();
        indent.clear();
        tagIndent.clear();
    } else if (indent.current() <= tagIndent.current()) {
        tags.pop();
        indent.pop();
    }
    tags.push(extractTagOnly(&line));
    expectedToken = Any;
    getNextLine = true;

    tagIndent.push(indent.current());
    return true;
}
