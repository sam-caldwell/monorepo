/**
 * @name ArbitraryValue/src/getUint.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */


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
uint ArbitraryValue::getUint(bool strongTyping = false) {
    lock->wait()->on();
    uint result;
    if (strongTyping && (type != Uint)) {
        lock->off();
        throw runtime_error("type mismatch");
    }
    if (value) {
        if (strongTyping && (type != Uint)) {
            throw runtime_error("type mismatch");
        }
        switch (type) {
            case Null: {
                lock->off();
                throw runtime_error("INTERNAL_ERROR: Null type objects cannot have data");
                break;
            }
            case Double: {
                double *f;
                f = (double *) value;
                if (*f < 0.0)
                    throw BoundsCheckError();
                else
                    result = (uint) trunc(*f);
                break;
            }
            case Float: {
                float *f;
                f = (float *) value;
                if (*f < 0.0)
                    throw BoundsCheckError();
                else
                    result = (uint) trunc(*f);
                break;
            }
            case Int: {
                result = *((int *) value);
                if (result < 0)
                    throw BoundsCheckError();
                break;
            }
            case Uint: {
                result = *((uint *) value);
                break;
            }
            case Bool: {
                bool f = (*((char *) value) == v_true);
                if (f)
                    result = 1;
                else
                    result = 0;
                break;
            }
            case String: {
                char *p;
                string t;
                p = (char *) value;
                for (int i = 0; i < sz; i++)
                    t.append(1, *(p + i));
                int i;
                i = atoi(t.c_str());
                if (i < 0)
                    throw BoundsCheckError();
                result = i;
                break;
            }
            default:
                lock->off();
                throw UnexpectedType(type);
        }
    }else{
        result =0;
    }
    lock->off();
    return result;
}
