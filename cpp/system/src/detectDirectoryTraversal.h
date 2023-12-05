/**
 * @name system/src/detectDirectoryTraversal.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <regex>
#include <string>

namespace File {
    /**
     * @name detectDirectoryTraversal
     * @brief detect when a path/filename contains directory traversal.
     *
     * @param &path string
     * @return bool
     */
    inline bool detectDirectoryTraversal(string path) {
        const string pattern = ".*(\\.\\./|/\\.\\.|%2e%2e/)+.*";
        return (std::regex_match(path, std::regex(pattern)));
    }
}