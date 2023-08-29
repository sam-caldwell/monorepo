#ifndef MISC_TRIM_H
#define MISC_TRIM_H

#include <string>


namespace misc {

    const std::string whitespace = " \t\n\r";

    std::string trim(const std::string &str) {
        size_t firstNonSpace = str.find_first_not_of(whitespace);
        size_t lastNonSpace = str.find_last_not_of(whitespace);

        if (firstNonSpace == std::string::npos || lastNonSpace == std::string::npos) {
            return ""; // String is all whitespace
        }

        return str.substr(firstNonSpace, lastNonSpace - firstNonSpace + 1);
    }

};

#endif