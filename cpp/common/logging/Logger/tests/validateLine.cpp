/**
 * @name Logger/tests/logLineRegex.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logger_Test_ValidateLine_cpp
#define Logger_Test_ValidateLine_cpp

#include <regex>
#include "../src/Logger.h"

/**
 * @name validateLine
 * @brief Test and expect that the line is properly formatted, etc.
 *
 * @params sv - log line severity
 * @params t - tag string
 * @params l - log message
 * @params m - isolated message content
 * @return bool
 */
bool validateLine(Logging::Severity sv, string *t, string *l, string *m) {
    string s = Logging::ReverseSeverity.at(sv);
    std::regex pattern(
            "^\\{\"n\":[0-9]{10},\"c\":[0-9]{1,3},\"s\":\"" + s + "\",\"m\":\"" + *m + "\",\"t\":\\[" + *t +"\\]\\}$");
    return std::regex_match(*l, pattern);
}

#endif /* Logger_Test_ValidateLine_cpp */