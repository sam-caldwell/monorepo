/**
 * @name Configuration/src/parseDouble.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#ifndef Configuration_ParseDouble_H
#define Configuration_ParseDouble_H

/**
 * @name parseDouble
 * @brief Parse a given input (v) and return the double-precision floating point
 *        representation.  If an empty value "" is encountered, return the default
 *        value.
 *
 * @param v string
 * @param defaultValue double
 * @return double.
 */
double Configuration::parseDouble(string v, double defaultValue) {
    double result = 0.0;
    if (v.compare("") == 0) {
        result = defaultValue;
    } else {
        result = stod(v);
    }
    return result;
}

#endif /* Configuration_ParseDouble_H */