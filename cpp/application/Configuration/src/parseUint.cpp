/**
 * @name Configuration/src/parseUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Configuration_ParseUint_H
#define Configuration_ParseUint_H

/**
 * @name parseUint
 * @brief Parse a given input (v) and return the unsigned integer representation.
 *        If an empty value "" is encountered, return the default value.
 *
 * @throws out_of_range exception if input (v) is < 0.
 *
 * @param v string
 * @param defaultValue uint
 * @return uint.
 */
int Configuration::parseUint(string v, uint defaultValue) {
    uint result = 0;
    if (v == "") {
        result = defaultValue;
    } else if (stoi(v) < 0) {
        throw out_of_range("signed integer (i<0) is out of bounds for uint.");
    } else {
        result = stoul(v);
    }
    return result;
}

#endif /*  Configuration_ParseUint_H */