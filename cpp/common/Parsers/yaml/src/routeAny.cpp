/**
 * Parsers/yaml/src/Parsers.h/routeAny.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name routeAny
 * @brief top-level parse-router will evaluate a line when there is no more-specific expectation
 *        to determine what kind of token is on our current line and route execution accordingly.
 *        Lines are preserved for the next, directed pass.
 *
 * @return bool
 */
inline bool Parser::routeAny() {

    getNextLine = false; //Preserve this line for the next pass.

    if (std::regex_match(line, std::regex(PATTERN_TAG_MULTILINE))) {
        // We have detected a multi-line tag ("identifier: |") pattern, which
        // we will consume for both key and expect a child token (value text block).
        expectedToken = MultiLineTag;

    } else if (std::regex_match(line, std::regex(PATTERN_TAG_MAP))) {
        // We have detected a tag ("identifier:") pattern, which we will
        // consume and then evaluate any (expected) child tokens.
        expectedToken = TagMap;

    } else if (std::regex_match(line, std::regex(PATTERN_SEQUENCE_KV_ITEM))) {
        // We have detected a Sequence item of type Key-Value ("- identifier: value").
        expectedToken = SequenceKeyValue;

    } else if (std::regex_match(line, std::regex(PATTERN_SEQUENCE_MAP_ITEM))) {
        // We have detected a Sequence item of type MapKey: ("- identifier:").
        expectedToken = SequenceMap;

    } else if (std::regex_match(line, std::regex(PATTERN_SEQUENCE_ELEMENTS))) {
        // We have detected a Sequence pattern that does not match the above,
        // so now know this is just a simple sequence item (value)
        expectedToken = SequenceItem;

    } else if (std::regex_match(line, std::regex(PATTERN_TAG_KEYVALUE))) {
        // We have detected a key-value tag ("identifier: value") pattern,
        // which we will consume for both key and value and expect no children.
        expectedToken = KeyValueTag;

    } else if (line != "") {
        throw runtime_error("Syntax Error: Unknown or unsupported pattern: '" + line + "'");
        // We have not identified the content of this line which means we are
        // in a syntax error.
    }
    return true;
}