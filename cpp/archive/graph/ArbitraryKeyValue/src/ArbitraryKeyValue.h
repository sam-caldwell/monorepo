/**
 * ArbitraryKeyValue/src/ArbitraryKeyValue.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * This is a class library defining an arbitrary key-value.
 * The data value is typed at runtime and stored in memory
 * with a string key (name).
 *
 */
#ifndef ArbitraryKeyValue_H
#define ArbitraryKeyValue_H

#include <regex>
#include "../../../../exceptions/exceptions.h"
#include "../../../../types/unsigned_int.h"
#include "../../ArbitraryValue/src/ArbitraryValue.h"

using namespace std;

class ArbitraryKeyValue : public ArbitraryValue {
public:
    /**
     * @name Default Class constructor
     * @brief initialize class state
     */
    ArbitraryKeyValue();

    /**
     * @name Class constructor
     * @brief initialize class state and set key name with no value.
     *
     * @param name string
     */
    ArbitraryKeyValue(string name);

    /**
     * @name Class constructor
     * @brief Class constructor (used to set key name and bool value).
     *
     * @param name string
     * @param v bool
     */
    ArbitraryKeyValue(string name, bool value);

    /**
     * @name Class Constructor
     * @brief Class constructor (used to set key name and double value).
     * @param name string
     * @param value double
     */
    ArbitraryKeyValue(string name, double value);

    /**
     * Class constructor (used to set key name and float value).
     * @param name string
     * @param value float
     */
    ArbitraryKeyValue(string name, float value);

    /**
     * Class constructor (used to set key name and int value).
     * @param name string
     * @param value int
     */
    ArbitraryKeyValue(string name, int value);

    /**
     * Class constructor (used to set key name and string value).
     * @param name string
     * @param value string
     */
    ArbitraryKeyValue(string name, string value);

    /**
     * Class constructor (used to set key name and uint value).
     * @param name string
     * @param value uint
     */
    ArbitraryKeyValue(string name, uint value);

    /**
    * key property setter
     * @param string s
    */
    string key();

    /**
     * key property getter
     * @return string
     */
    void key(string s);

protected:
    string keyName;
};

#include "ArbitraryKeyValue.cpp"

#endif /* ArbitraryKeyValue_H */