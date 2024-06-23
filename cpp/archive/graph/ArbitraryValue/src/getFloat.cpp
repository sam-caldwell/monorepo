/**
 * @name ArbitraryValue/src/getFloat.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * @brief This is a class library defining an arbitrary value.
 *        The data value is typed at runtime and stored in memory.
 */


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
float ArbitraryValue::getFloat(bool strongTyping = false) {
    lock->wait()->on();
    float result;
    if (strongTyping && (type != Float)) {
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
                result = (float) (*f);
                break;
            }
            case Float: {
                result = *((float *) value);
                break;
            }
            case Int: {
                int *f;
                f = (int *) value;
                result = (float) (*f);
                break;
            }
            case Uint: {
                uint *f;
                f = (uint *) value;
                result = (float) (*f);
                break;
            }
            case Bool: {
                bool f = (*((char *) value) == v_true);
                if (f)
                    result = 1.0;
                else
                    result = 0.0;
                break;
            }
            case String: {
                char *p;
                string t;
                p = (char *) value;
                for (int i = 0; i < sz; i++)
                    t.append(1, *(p + i));
                result = atof(t.c_str());
                break;
            }
            default:
                lock->off();
                throw UnexpectedType(type);
        }
    }else{
        result = 0.0;
    }
    lock->off();
    return result;
}
