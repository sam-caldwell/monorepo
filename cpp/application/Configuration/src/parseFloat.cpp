/**
 * @name Configuration/src/parseFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#ifndef Configuration_ParseFloat_H
#define Configuration_ParseFloat_H

/**
 * @name parseFloat
 * @brief Parse a given input (v) and return the single-precision floating point
 *        representation.  If an empty value "" is encountered, return the default
 *        value.
 *
 * @param v string
 * @param defaultValue float
 * @return double.
 */
float Configuration::parseFloat(string v, float defaultValue) {
    float result = 0.0;
    if (v == "") {
        result = defaultValue;
    } else {
        result = stof(v);
    }
    return result;
}

#endif /* Configuration_ParseFloat_H */