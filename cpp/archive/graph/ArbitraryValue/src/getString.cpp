/**
 * @name ArbitraryValue/src/getString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */


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
string ArbitraryValue::getString(bool strongTyping = false) {
    lock->wait()->on();
    string result;
    if (strongTyping && (type != String)) {
        lock->off();
        throw runtime_error("type mismatch");
    }
    if (value) {
        char *p;
        p = (char *) value;
        switch (type) {
            case Null: {
                lock->off();
                throw runtime_error("INTERNAL_ERROR: Null type objects cannot have data");
                break;
            }
            case String: {
                for (int i = 0; i < sz; i++)
                    result.append(1, *(p + i));
                break;
            }
            case Bool: {
                result = (*p == v_true) ? "true" : "false";
                break;
            }
            case Double: {
                double *f;
                f = (double *) value;
                result = to_string(*f);
                break;
            }
            case Float: {
                float *f;
                f = (float *) value;
                result = to_string(*f);
                break;
            }
            case Int: {
                int *f;
                f = (int *) value;
                result = to_string(*f);
                break;
            }
            case Uint: {
                uint *f;
                f = (uint *) value;
                result = to_string(*f);
                break;
            }
            default:
                lock->off();
                throw UnexpectedType(type);
        }
    } else {
        result = ""; // NULL string returns as ""
    }
    lock->off();
    return result;
}
