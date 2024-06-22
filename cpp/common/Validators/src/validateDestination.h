/**
 * @name Validators/src/validateDestination.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logger_validateDestination_h
#define Logger_validateDestination_h

#include <regex>
#include "../validators.h"

/**
 * @name validateDestination
 * @brief ensure the destination follows the expected patterns.
 *
 * @throws out_of_range exception if validation fails.
 *
 * @param *s - string reference
 * @return string reference
 */
void validateDestination(string *s) {
    std::regex ok("(^(syslog:|https:)\\/\\/[a-zA-Z0-9\\._\\-\\/]+:[0-9]{1,5}$|"\
                  "^(file:\\/|[.]{0,1})\\/[a-zA-Z0-9\\._\\-\\/]+$)");

    if (s) {
        if (s->length() == 0) {
            s = nullptr;
        } else {
            if (!std::regex_match(*s, ok))
                throw out_of_range("Logger destination path failed validation.");

            if (directoryTraversal(s))
                throw out_of_range("Logger destination path cannot include directory traversal (e.g. '..')");
        }
    }
}


#endif /* Logger_validateDestination_h */
