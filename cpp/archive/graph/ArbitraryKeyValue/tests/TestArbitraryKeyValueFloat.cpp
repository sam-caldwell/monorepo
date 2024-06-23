/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "../src/ArbitraryKeyValue.h"

class TestArbitraryKeyValueFloat : public TestBase {
public:
    TestArbitraryKeyValueFloat(string n) { name = n; }

    /**
     * Given a boolean value (a), a name "testKey" and an ArbitraryKeyValue,
     * - Expect the name is properly set.
     * - Expect the value (a) is properly set.
     * - Expect the ArbitraryKeyValue is of type Bool.
     *
     * @param a float
     * @return bool
     */
    bool test_func(string key, float a) {
        debug(name + "::test_func(): starting...");
        ArbitraryKeyValue v(key, a);
        if(!expect(v.getType() == Float, "Expect Type Float")) return false;
        if(!expect(v.key()==key, "Expect key match "+key)) return false;
        bool result = v.getFloat() == a;
        debug(name + "::test_func(): done (result:" + to_string(result) + ")...");
        return result;
    }

    bool main() {
        return expect(
                test_func("key", 1.0),
                "Given a name and value(1.0), expect ArbitraryKeyValue to properly configure") ||
               expect(
                       test_func("key", 0.0),
                       "Given a name and value(0.0), expect ArbitraryKeyValue to properly configure") ||
               expect(
                       test_func("key", -1.0),
                       "Given a name and value(-1.0), expect ArbitraryKeyValue to properly configure");
    }
};
