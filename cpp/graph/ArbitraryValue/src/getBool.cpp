/**
 * @name ArbitraryValue/src/getBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */


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
bool ArbitraryValue::getBool(bool strongTyping = false) {
    bool response;
    lock->wait()->on();
    if (strongTyping && (type != Bool)) {
        lock->off();
        throw runtime_error("type mismatch");
    }
    if (value) {
        switch (type) {
            case Null:
                lock->off();
                throw runtime_error("INTERNAL_ERROR: Null type objects cannot have data");
                break;
            case Bool:
                response = (value) ? (*((char *) value) == v_true) : false;
                break;
            case Double:
                response = (*((double *) value) != 0.0);
                break;
            case Float:
                response = (*((float *) value) != 0.0);
                break;
            case Int:
                response = (*((int *) value) != 0);
                break;
            case Uint:
                response = (*((uint *) value) > 0);
                break;
            case String: {
                char *p;
                string t;
                p = (char *) value;
                for (int i = 0; i < sz; i++)
                    t.append(1, *(p + i));
                if (t == "true")
                    response = true;
                else
                    response = false;
                break;
            }
            default:
                lock->off();
                throw UnexpectedType(type);
        }
    } else {
        response = false;
    }
    lock->off();
    return response;
}
