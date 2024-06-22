/**
 * Parsers/yaml/src/Parsers.h/routeSequenceItem.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "extractKvValue.cpp"

/**
 * @name routeSequenceItem
 * @brief we detected a simple sequence item ("- identifier") which we will store as a value.
 *
 * @return bool
 */
inline bool Parser::routeSequenceItem() {
    if (std::regex_match(line, std::regex(PATTERN_SEQUENCE_ELEMENTS))) {
        string value = [](string *line) {
            std::size_t startPos = line->find_first_of("-");
            if (startPos != string::npos)
                return line->substr(startPos + 2, string::npos);
            return string("");
        }(&line);

        tokens.insert(pair<string,string>(tags.toString(), value));
        /**
         * After a sequence item, we expect anything could occur in our file.
         */
        expectedToken = SequenceItem;
        getNextLine = true;
        return true;
    } else {
        tags.pop();
        tags.pop();

        expectedToken = Any;
        getNextLine = false;

        return false;
    }
}
