/**
 * @name ArbitraryValue/src/constructors.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */


/**
 * @name default constructor
 * @brief Class constructor (default)
 */
ArbitraryValue::ArbitraryValue() {
    lock = new SimpleLock(ARBITRARY_VALUE_LOCK_TIMEOUT);
    type = Null;
    value = NULL;
    sz = 0;
}

/**
 * @name Class Constructor (bool)
 * @brief initialize boolean object
 *
 * @param v bool
 */
ArbitraryValue::ArbitraryValue(bool v) {
    lock = new SimpleLock(ARBITRARY_VALUE_LOCK_TIMEOUT);
    type = Null;
    value = NULL;
    sz = 0;
    setBool(v);
}

/**
 * @name Class Constructor (double)
 * @brief initialize double object
 *
 * @param v double
 */
ArbitraryValue::ArbitraryValue(double v) {
    lock = new SimpleLock(ARBITRARY_VALUE_LOCK_TIMEOUT);
    type = Null;
    value = NULL;
    sz = 0;
    setDouble(v);
}

/**
 * @name Class Constructor (float)
 * @brief initialize float object
 *
 * @param v float
 */
ArbitraryValue::ArbitraryValue(float v) {
    lock = new SimpleLock(ARBITRARY_VALUE_LOCK_TIMEOUT);
    type = Null;
    value = NULL;
    sz = 0;
    setFloat(v);
}

/**
 * @name Class Constructor (int)
 * @brief initialize int object
 *
 * @param v int
 */
ArbitraryValue::ArbitraryValue(int v) {
    lock = new SimpleLock(ARBITRARY_VALUE_LOCK_TIMEOUT);
    type = Null;
    value = NULL;
    sz = 0;
    setInt(v);
}

/**
 * @name constructor(string)
 * @brief initialize string object
 *
 * @param v string
 */
ArbitraryValue::ArbitraryValue(string v) {
    lock = new SimpleLock(ARBITRARY_VALUE_LOCK_TIMEOUT);
    type = Null;
    value = NULL;
    sz = 0;
    setString(v);
}

/**
 * @name Class Constructor (uint)
 * @brief initialize uint object
 *
 * @param v uint
 */
ArbitraryValue::ArbitraryValue(uint v) {
    lock = new SimpleLock(ARBITRARY_VALUE_LOCK_TIMEOUT);
    type = Null;
    value = NULL;
    sz = 0;
    setUint(v);
}
