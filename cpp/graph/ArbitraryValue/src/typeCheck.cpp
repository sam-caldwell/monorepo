/**
 * @name ArbitraryValue/src/typeCheck.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */

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
ValueTypes ArbitraryValue::typeCheck(ValueTypes t, bool setIfNotMatched) {
    if (isType(t) || setIfNotMatched) {
        type = t;
        return t;
    }
    throw UnexpectedType(t);
}

/**
 * @name typeCheck(t,dt)
 * @brief Check to ensure current state is type t or type nt.
 *        Otherwise throw exception.
 *
 * @param ValueTypes t
 * @param ValueTypes nt
 * @return ValueTypes
 */
ValueTypes inline ArbitraryValue::typeCheck(ValueTypes t, ValueTypes dt) {
    if (isType(t) || isType(dt))
        return t;
    else
        throw UnexpectedType(t);
}
