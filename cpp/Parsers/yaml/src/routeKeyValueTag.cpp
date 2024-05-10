/**
 * Parsers/yaml/src/Parsers.h/routeKeyValueTag.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/common/Parser/src/extractKvValue.cpp"

/**
 * @name routeKeyValueTag
 * @brief We have detected a simple key:value pattern ("identifier: value")
 *        In this method, we capture the tag (key) then expect the value in
 *        the next pass.
 *
 * @return bool
 */
inline bool Parser::routeKeyValueTag() {

    //Step back one tag if we detect indent has stepped back.
    if (indent.current() <= indent.prev()) {
        tags.pop();
        indent.pop();
    }

    tags.push(extractTagOnly(&line));
    string value = extractKvValue(&line);

    tokens.insert(pair<string, string>(tags.toString(), value));
    /**
     * Next we need to clean up our state for the next pass...
     * - pop our Multi-line tag from the end of our prefix tags since we
     *   will have no child tags.
     */
    tags.pop();
    indent.pop();
    getNextLine = true;
    expectedToken = Any;
    return true;
}
