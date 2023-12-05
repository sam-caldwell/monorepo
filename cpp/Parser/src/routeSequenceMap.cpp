/**
 * @name Parser/src/routeSequenceMap.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name routeSequenceMap
 * @brief We have detected a sequence map ("- map:") and we expect a child node or sequence.
 *        so all we will do here is append the tag to our tags collection.  There is no value yet.
 *
 * @return bool
 */
inline bool Parser::routeSequenceMap() {
    tags.push(extractTagOnly(&line));

    expectedToken = Any;
    getNextLine = true;
    return true;
}
