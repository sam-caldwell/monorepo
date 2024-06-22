/**
 * Parsers/yaml/src/Parsers.h/parse.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <string>
#include <vector>
#include "../../../formatter/formatter.h"
#include "TokenType.cpp"
#include "../../../Breadcrumbs/src/Breadcrumbs.h"
#include "extractTagOnly.cpp"

/**
 * @name parse
 * @brief stream the tokens to the given output stream.
 *
 * @throws runtime_error if header not found when expected.
 * @throws runtime_error if Syntax error occurs
 * @throws runtime_error if Unexpected tokens are encountered.
 *
 */
void Parser::parse() {
    /**
     * Loop through the YAML file.
     */
    while (!file.eof()) {
        fetchNextLine();
        switch (expectedToken) {
            case MultiLineTag:
                routeMultiLineTag(&line);
                break;
            case KeyValueTag:
                routeKeyValueTag();
                break;
            case SequenceKeyValue:
                tags.push("["+to_string(index++)+"]");
                routeSequenceKeyValue();
                break;
            case SequenceMap:
                routeSequenceMap();
                break;
            case SequenceItem:
                tags.push("["+to_string(index++)+"]");
                routeSequenceItem();
                tags.pop();
                break;
            case TagMap:
                routeTagMap();
                break;
            case Value:
                //ToDo: grab the value.
                break;
            case Header:
                routeHeader();
                break;
            default:
                routeAny();
                break;
        }/* switch (expectedToken) */
    }/* while (!file.eof()) */
}/* end of parse() */
