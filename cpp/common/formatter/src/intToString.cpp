/**
 * @name formatter/src/intToString.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef IntToString_H
#define IntToString_H

#include <sstream>
#include <iomanip>

using namespace std;

/**
 * @name intToString
 * @brief Return a formatted string with a given width.
 *
 * @param i integer to format
 * @param width width (padded with leading zeros)
 * @return string
 */
string intToString(int i, int width=1){

    stringstream ss;

    ss << setw(width) << setfill('0') << i;

    return ss.str();
}

#endif /** IntToString_H */