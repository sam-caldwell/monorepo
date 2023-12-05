/**
 * @name Parser/src/routeSequenceKeyValue.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/common/Parser/src/extractKvValue.cpp"

/**
 * @name routeSequenceKeyValue
 * @brief We have detected a sequence key-value pair.  We know that our fully-qualified tag
 *        name at this point is the parent identity of the overall sequence, but our current
 *        key and its value must be indexed as part of that sequence ("- key:value").
 *
 * @return bool
 */
inline bool Parser::routeSequenceKeyValue() {
    if (std::regex_match(line, std::regex(PATTERN_SEQUENCE_KV_ITEM))) {
        tags.push(extractTagOnly(&line));
        string value = extractKvValue(&line);

        tokens.insert(pair<string, string>(tags.toString(), value));
        /**
         * Next we need to clean up our state for the next pass...
         * - pop our Multi-line tag from the end of our prefix tags since we
         *   will have no child tags.
         */
        tags.pop(); //pop current tag
        tags.pop(); //pop index
        getNextLine = true;
        expectedToken = SequenceKeyValue;
        return true;
    }else{
        tags.pop();
        indent.pop();
        expectedToken = Any;
        getNextLine = false;
        return false;
    }
}