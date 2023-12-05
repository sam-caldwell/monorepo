/**
 * @name ArbitraryKeyValue/src/ArbitraryKeyValue.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary key-value.
 * The data value is typed at runtime and stored in memory
 * with a string key (name).
 *
 */

/**
 * @name Default class constructor
 * @brief noop
 */
ArbitraryKeyValue::ArbitraryKeyValue() {};

/**
 * @name class constructor (provides object key)
 * @brief Class constructor (used to set the key name, no value).
 *
 * @param name string
 */
ArbitraryKeyValue::ArbitraryKeyValue(string name) {
    key(name);
}

/**
 * @class constructor (boolean)
 * @brief Class constructor (used to set key name and bool value).
 * @param name string
 * @param v bool
 */
ArbitraryKeyValue::ArbitraryKeyValue(string name, bool v) {
    key(name);
    setBool(v);
}

/**
 * Class constructor (used to set key name and double value).
 * @param name string
 * @param value double
 */
ArbitraryKeyValue::ArbitraryKeyValue(string name, double v) {
    key(name);
    setDouble(v);
}

/**
 * Class constructor (used to set key name and float value).
 * @param name string
 * @param value float
 */
ArbitraryKeyValue::ArbitraryKeyValue(string name, float value) {
    key(name);
    setFloat(value);
}

/**
 * Class constructor (used to set key name and int value).
 * @param name string
 * @param value int
 */
ArbitraryKeyValue::ArbitraryKeyValue(string name, int value) {
    key(name);
    setInt(value);
}

/**
 * Class constructor (used to set key name and string value).
 * @param name string
 * @param value string
 */
ArbitraryKeyValue::ArbitraryKeyValue(string name, string value) {
    key(name);
    setString(value);
}

/**
 * Class constructor (used to set key name and uint value).
 * @param name string
 * @param value uint
 */
ArbitraryKeyValue::ArbitraryKeyValue(string name, uint value) {
    key(name);
    setUint(value);
}

/**
* key property setter
 * @param string s
*/
void ArbitraryKeyValue::key(string s) {
    if (regex_match(s, regex("[a-zA-Z][a-zA-Z0-9._-]+[a-zA-Z0-9]{0,62}")))
        keyName = s;
    else
        throw InvalidKeyString();
}

/**
 * key property getter
 * @return string
 */
string ArbitraryKeyValue::key() {
    return keyName;
}
