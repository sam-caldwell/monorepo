/**
 * @name Validators/src/directoryTraversal.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef Logger_directoryTraversal_h
#define Logger_directoryTraversal_h

#include <string>
#include <regex>

/**
 * @name directoryTraversal
 * @brief identify if a directory traversal risk exists in a path.
 *
 * @param s - string reference
 * @return bool
 */
bool directoryTraversal(string *s) {
    std::regex p1(".*[.]{2}.*");
    std::regex p2(".*%2e%2e.*");

    return std::regex_match(*s, p1) ||
           std::regex_match(*s, p2);
}


#endif /* Logger_directoryTraversal_h */