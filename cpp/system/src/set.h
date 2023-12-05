/**
 * @name system/src/set.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/common/formatter/formatter.h"

namespace Environment {
    /**
     * @name set
     * @brief set an Environment variable with a given string value.
     *
     * @param varname string
     * @param value string
     * @return bool (operation outcome)
     */
    inline bool set(string varname, string value) {
        return setenv(varname.c_str(), value.c_str(),1) == 0;
    }
};