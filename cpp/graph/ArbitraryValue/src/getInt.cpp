/**
 * @name ArbitraryValue/src/getInt.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */


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
int ArbitraryValue::getInt(bool strongTyping = false) {
    lock->wait()->on();
    int result;
    if (strongTyping && (type != Int)) {
        lock->off();
        throw runtime_error("type mismatch");
    }
    if (value) {
        switch (type) {
            case Null: {
                lock->off();
                throw runtime_error("INTERNAL_ERROR: Null type objects cannot have data");
                break;
            }
            case Double: {
                double *f;
                f = (double *) value;
                if ((*f > numeric_limits<int>::max()) || (*f < numeric_limits<int>::min()))
                    throw BoundsCheckError();
                result = (int) trunc(*f);
                break;
            }
            case Float: {
                float *f;
                f = (float *) value;
                if ((*f > numeric_limits<int>::max()) || (*f < numeric_limits<int>::min()))
                    throw BoundsCheckError();
                result = (int) trunc(*f);
                break;
            }
            case Int: {
                result = *((int *) value);
                break;
            }
            case Uint: {
                uint *f;
                f = (uint *) value;
                result = (int) (*f);
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
                result = atoi(t.c_str());
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
