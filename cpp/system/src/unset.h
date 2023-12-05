/**
 * @name system/src/unset.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace Environment {

    /**
     * @name unset
     * @brief unset an Environment variable (s).
     *
     * @param s string
     * @return bool
     */
    bool unset(string s) {
        return unsetenv(s.c_str()) == 0;
    }

};