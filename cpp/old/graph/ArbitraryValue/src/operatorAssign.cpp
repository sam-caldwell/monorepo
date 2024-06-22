/**
 * @name ArbitraryValue/src/operatorAssign.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */

/**
 * @name Assignment operator (=) boolean
 * @param v bool
 * @return ArbitraryValue
 */
ArbitraryValue ArbitraryValue::operator=(bool v) {
    setBool(v);
    return *this;
}

/**
 * @name Assignment operator (=) double-precision floating-point number
 * @param v double
 * @return ArbitraryValue
 */
ArbitraryValue ArbitraryValue::operator=(double v) {
    setDouble(v);
    return *this;
}

/**
 * @name Assignment operator (=) single-precision floating-point number
 * @param v float
 * @return ArbitraryValue
 */
ArbitraryValue ArbitraryValue::operator=(float v) {
    setFloat(v);
    return *this;
}

/**
 * @name Assignment operator(=) integer
 * @param v int
 * @return ArbitraryValue
 */
ArbitraryValue ArbitraryValue::operator=(int v) {
    setInt(v);
    return *this;
}

/**
 * @name Assignment operator (=) string
 * @param v string
 * @return ArbitraryValue
 */
ArbitraryValue ArbitraryValue::operator=(string v) {
    setString(v);
    return *this;
}

/**
 * @name Assignment operator (=) unsigned integer
 * @param v uint
 * @return ArbitraryValue
 */

ArbitraryValue ArbitraryValue::operator=(uint v) {
    setUint(v);
    return *this;
}

/**
 * @name Assignment operator (=) ArbitraryValue
 * @param v ArbitraryValue
 * @return ArbitraryValue
 */
ArbitraryValue ArbitraryValue::operator=(ArbitraryValue v) {
    lock->wait()->on();
    type = v.type;
    value = v.value;
    sz = v.sz;
    lock->off();
    return *this;
}

/**
 * @name Assignment operator (=) ArbitraryValue
 * @param *v ArbitraryValue
 * @return ArbitraryValue pointer
 */
ArbitraryValue *ArbitraryValue::operator=(ArbitraryValue *v) {
    lock->wait()->on();
    type = v->type;
    value = v->value;
    sz = v->sz;
    lock->off();
    return this;
}
