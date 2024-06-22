/**
 * @name Validators/src/validateTags.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logger_validateTags_h
#define Logger_validateTags_h

#include <regex>

/**
 * @name validateTags
 * @brief parse the quoted comma-delimited string of tags and throw an exception if not properly
 *        formatted.
 *
 * @throws out_of_range exception if validation fails.
 *
 * @param *tags - string (by reference)
 */
void validateTags(string *tags) {
    std::regex pattern("^\"[a-zA-Z.\\-_0-9]{1,63}\"$|^\"[a-zA-Z.\\-_0-9]{1,63}:[a-zA-Z.\\-_0-9]{1,63}"\
                       "\"$|^(\"[a-zA-Z.\\-_0-9]{1,63}\",){1,254}\"[a-zA-Z.\\-_0-9]{1,63}\"$|"\
                       "^(\"[a-zA-Z.\\-_0-9]{1,63}:[a-zA-Z.\\-_0-9]+\",)+\"[a-zA-Z.\\-_0-9]{1,63}"\
                       ":[a-zA-Z.\\-_0-9]+\"$");
    if (tags) {
        if (tags->length() == 0) {
            tags = nullptr;
        } else {
            if (!std::regex_match(*tags, pattern))
                throw out_of_range("Logger cannot validate given tag set.");
        }
    }
}

#endif /* Logger_validateTags_h */

