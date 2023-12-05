/**
 * @name Configuration/src/parseBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Configuration_ParseBool_H
#define Configuration_ParseBool_H

/**
 * @name parseBool
 * @brief Parse a given input (v) and return the boolean equivalent.
 *        Where an empty value "" is encountered, return the default value.
 *
 * @throws runtime_error on non-sense, non-boolean data.
 *
 * @param v string
 * @param defaultValue bool
 * @return bool
 */
bool Configuration::parseBool(string v, bool defaultValue) {
    bool result = false;
    string s = toLower(v);
    if ((s == "true") || (s == "1"))
        result = true;
    else if ((s == "false") || (s == "0"))
        result = false;
    else if (v == "")
        result = defaultValue;
    else
        throw runtime_error("Invalid boolean value.  Expected true or false");
    return result;
}

#endif /* Configuration_ParseBool_H */