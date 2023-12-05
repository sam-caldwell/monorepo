/**
 * @name Configuration/src/parseInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Configuration_ParseInt_H
#define Configuration_ParseInt_H

/**
 * @name parseInt
 * @brief Parse a given input (v) and return the integer representation.
 *        If an empty value "" is encountered, return the default value.
 *
 * @param v string
 * @param defaultValue int
 * @return int.
 */
int Configuration::parseInt(string v, int defaultValue) {
    int result = 0;
    if (v == "") {
        result = defaultValue;
    } else {
        result = stoi(v);
    }
    return result;
}

#endif /* Configuration_ParseInt_H */