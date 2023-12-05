/**
 * @name system/src/get.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
namespace Environment {

    /**
     * @name get
     * @brief get the value of a given Environment variable.
     *
     * @param varname string
     * @return string
     */
    inline string get(string varname) {
        const char *s = getenv(varname.c_str());
        if (s == NULL)
            throw runtime_error("missing Environment variable:" + varname);
        else
            return (string)(s);
    }

};