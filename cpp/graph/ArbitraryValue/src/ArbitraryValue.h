/**
 * ArbitraryValue/src/ArbitraryValue.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * This is a class library defining an arbitrary value.
 * The data value is typed at runtime and stored in memory.
 *
 */
#ifndef ArbitraryValue_H
#define ArbitraryValue_H

#define ARBITRARY_VALUE_LOCK_TIMEOUT 3600



/**
 *
 */
#include <cstdlib>
#include <cstring>
#include <stdlib.h>
#include <string>
#include <stdlib.h>
#include <iostream>
#include <cmath>
#include <limits>
#include "projects/common/SimpleLock/src/SimpleLock.h"
#include "projects/common/types/unsigned_int.h"
#include "projects/common/flags/flags.h"
#include "projects/common/exceptions/exceptions.h"
#include "projects/common/types/ValueTypes.h"

using namespace std;

/**
 * Note that we encode boolean true/false
 * as 0x1O and 0xFF because 0x00 (NULL)
 * has a special meaning in our system.
 */
const char v_true = '1';
const char v_false = '0';

class ArbitraryValue {
public:
    /**
     * @name default constructor
     * @brief Class constructor (default)
     */
    ArbitraryValue();

    /**
     * @name Class Constructor (bool)
     * @brief initialize boolean object
     *
     * @param v bool
     */
    ArbitraryValue(bool v);

    /**
     * @name Class Constructor (double)
     * @brief initialize double object
     *
     * @param v double
     */
    ArbitraryValue(double v);

    /**
     * @name Class Constructor (float)
     * @brief initialize float object
     *
     * @param v float
     */
    ArbitraryValue(float v);

    /**
     * @name Class Constructor (int)
     * @brief initialize int object
     *
     * @param v int
     */
    ArbitraryValue(int v);

    /**
     * @name constructor(string)
     * @brief initialize string object
     *
     * @param v string
     */
    ArbitraryValue(string v);

    /**
     * @name Class Constructor (uint)
     * @brief initialize uint object
     *
     * @param v uint
     */
    ArbitraryValue(uint v);

    /**
     * Class destructor.
     *   - Safely erase our memory
     */
    ~ArbitraryValue();

    /**
     * @name getType
     * @brief return the ValueType of the given object
     *
     * @return ValueTypes
     */
    ValueTypes inline getType();

    /**
     * @name getBool()
     * @brief decode the data and return the given datatype.
     *        If a given type is requested (e.g. string but the)
     *        object is a different type (e.g. bool) perform a
     *        type conversion if possible or throw an exception.
     *
     * @warning will throw runtime_error exception if strongTyping is enforced and violation occurs.
     *
     * @param strongTyping bool (default: false)
     * @return bool
     */
    bool getBool(bool strongTyping);

    /**
     * @name getDouble()
     * @brief decode the data and return the given datatype.
     *        If a given type is requested (e.g. string but the)
     *        object is a different type (e.g. bool) perform a
     *        type conversion if possible or throw an exception.
     *
     * @warning will throw runtime_error exception if strongTyping is enforced and violation occurs.
     *
     * @param strongTyping bool (default: false)
     * @return double
     */
    double getDouble(bool strongTyping);

    /**
     * @name getFloat()
     * @brief decode the data and return the given datatype.
     *        If a given type is requested (e.g. string but the)
     *        object is a different type (e.g. bool) perform a
     *        type conversion if possible or throw an exception.
     *
     * @warning will throw runtime_error exception if strongTyping is enforced and violation occurs.
     *
     * @param strongTyping bool (default: false)
     * @return float
     */
    float getFloat(bool strongTyping);

    /**
     * @name getInt()
     * @brief decode the data and return the given datatype.
     *        If a given type is requested (e.g. string but the)
     *        object is a different type (e.g. bool) perform a
     *        type conversion if possible or throw an exception.
     *
     * @warning will throw runtime_error exception if strongTyping is enforced and violation occurs.
     *
     * @param strongTyping bool (default: false)
     * @return int
     */
    int getInt(bool strongTyping);

    /**
     * @name getString()
     * @brief decode the data and return the given datatype.
     *        If a given type is requested (e.g. string but the)
     *        object is a different type (e.g. bool) perform a
     *        type conversion if possible or throw an exception.
     *
     * @warning will throw runtime_error exception if strongTyping is enforced and violation occurs.
     *
     * @param strongTyping bool (default: false)
     * @return string
     */
    string getString(bool strongTyping);

    /**
     * @name getUint()
     * @brief decode the data and return the given datatype.
     *        If a given type is requested (e.g. string but the)
     *        object is a different type (e.g. bool) perform a
     *        type conversion if possible or throw an exception.
     *
     * @warning will throw runtime_error exception if strongTyping is enforced and violation occurs.
     *
     * @param strongTyping bool (default: false)
     * @return uint
     */
    uint getUint(bool strongTyping);

    /**
     * @name isType
     * @brief return true/false whether or not the class is of type t.
     *
     * @param ValueTypes t
     * @return bool
     */
    bool isType(ValueTypes t);

    /**
     *  @name setBool()
     *  @brief Stores the given value (v) as byte data
     *
     * @param bool *v
     * @return bool
     */
    bool setBool(bool v);

    /**
    *  @name setDouble
    *  @brief Stores the given value (v) as byte data
    *
    * @param double v
    * @return bool
    */
    bool setDouble(double v);

    /**
     * @name setFloat
     * @brief Stores the given value (v) as byte data
     *
     * @param float v
     * @return bool
     */
    bool setFloat(float v);

    /**
     * @name setInt
     * @brief Stores the given value (v) as byte data
     *
     * @param int v
     * @return bool
     */
    bool setInt(int v);

    /**
     * @name setString(*v))
     * @brief Stores the given value (v) (NULL))
     *
     * @param string v
     * @return bool
     */
    bool setString(string v);

    /**
     * @name setUint
     * @brief Stores the given value (v) as byte data
     *
     * @param uint *v
     * @return bool
     */
    bool setUint(uint v);

    /**
     * @name typeCheck(t,setIfNotMatched)
     * @brief Perform type check.  If not type t, then set class to
     *        type t and return t if setIfNotMatched is true.
     *        But if setIfNotMatched is false and no match, throw an
     *        exception.
     *
     * @param ValueTypes t
     * @param bool setIfNotMatched (default False)
     * @return ValueTypes
     */
    ValueTypes typeCheck(ValueTypes t, bool setIfNotMatched = false);

    /**
     * @name typeCheck(t,dt)
     * @brief Check to ensure current state is type t or type nt.
     *        Otherwise throw exception.
     *
     * @param ValueTypes t
     * @param ValueTypes nt
     * @return ValueTypes
     */
    ValueTypes typeCheck(ValueTypes t, ValueTypes nt);

    /**
     * @name Assignment operator (=) boolean
     * @param v bool
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(bool v);

    /**
     * @name Assignment operator (=) boolean
     * @param *v bool pointer
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(bool *v);

    /**
     * @name Assignment operator (=) double-precision floating-point number
     * @param v double
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(double v);

    /**
     * @name Assignment operator (=) double-precision floating-point number
     * @param *v double pointer
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(double *v);

    /**
     * @name Assignment operator (=) single-precision floating-point number
     * @param v float
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(float v);

    /**
     * @name Assignment operator (=) single-precision floating-point number
     * @param *v float pointer
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(float *v);

    /**
     * @name Assignment operator(=) integer
     * @param v int
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(int v);

    /**
     * @name Assignment operator(=) integer
     * @param *v int pointer
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(int *v);

    /**
     * @name Assignment operator (=) string
     * @param v string
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(string v);

    /**
     * @name Assignment operator (=) string
     * @param *v string pointer
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(string *v);

    /**
     * @name Assignment operator (=) ArbitraryValue
     * @param v ArbitraryValue
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(uint v);

    /**
     * @name Assignment operator (=) ArbitraryValue
     * @param v ArbitraryValue pointer
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(uint *v);

    /**
     * @name Assignment operator (=) ArbitraryValue
     * @param v ArbitraryValue
     * @return ArbitraryValue
     */
    ArbitraryValue operator=(ArbitraryValue v);

    /**
     * @name Assignment operator (=) ArbitraryValue
     * @param *v ArbitraryValue pointer
     * @return ArbitraryValue pointer
     */
    ArbitraryValue *operator=(ArbitraryValue *v);

private:
    /**
     * @name sz
     * @brief memory buffer length for internal state.
     */
    uint sz;
    /**
     * @name *value
     * @brief pointer to the first byte of arbitrary stored data.
     *
     */
    void *value;
    /**
     * @name type
     * @brief ValueTypes type specifier for the stored data.
     */
    ValueTypes type;
    /**
     * @name SimpleLock
     * @brief locking mechanism to ensure safety.
     */
    SimpleLock *lock;

    /**
     * @name nullCheck
     * @brief throw exception if value is null.
     *
     * @param t ValueTypes
     */
    void nullCheck();

    /**
     * @name safeEraseMemory()
     * @brief overwrite the value memory buffer of the given
     *        size, then ensure the buffer is null-terminated.
     *
     * @param *p void pointer
     * @param sz int
     */
    void safeEraseMemory(void *p, int sz);

    /**
     * @name memoryCleanup()
     * @brief free any allocated memory object
     * @note use before assigning a new value to avoid leaks.
     */
    inline void memoryCleanup();
};

#include "nullCheck.cpp"
#include "getType.cpp"
#include "isType.cpp"
#include "typeCheck.cpp"
#include "constructors.cpp"
#include "destructors.cpp"
#include "setBool.cpp"
#include "setDouble.cpp"
#include "setFloat.cpp"
#include "setInt.cpp"
#include "setString.cpp"
#include "setUint.cpp"
#include "getBool.cpp"
#include "getDouble.cpp"
#include "getFloat.cpp"
#include "getInt.cpp"
#include "getString.cpp"
#include "getUint.cpp"
#include "operatorAssign.cpp"


#endif /* ArbitraryValue_H */