/**
 * @name system/src/exists.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#ifndef System_File_Exists_H
#define System_File_Exists_H

#include <string>
#include <sys/stat.h>

namespace File {
    /**
     * @name exists
     * @brief Determine if a given file exists.
     *
     * @param &name string (file name)
     * @return bool
     */
    inline bool exists(const std::string &fileName) {
        struct stat buffer;
        return (stat(fileName.c_str(), &buffer) == 0);
    }
}/* end of namespace File */

#endif /* System_File_Exists_H */