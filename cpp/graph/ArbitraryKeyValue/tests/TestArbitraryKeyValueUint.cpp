/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKeyValue/src/ArbitraryKeyValue.h"

class TestArbitraryKeyValueUint : public TestBase {
public:
    TestArbitraryKeyValueUint(string n) { name = n; }

    /**
     * Given a boolean value (a), a name "testKey" and an ArbitraryKeyValue,
     * - Expect the name is properly set.
     * - Expect the value (a) is properly set.
     * - Expect the ArbitraryKeyValue is of type Bool.
     *
     * @param a uint
     * @return bool
     */
    bool test_func(string key, uint a) {
        debug(name + "::test_func(): starting...");
        ArbitraryKeyValue v(key, a);
        if(!expect(v.getType() == Uint, "Expect Type Uint")) return false;
        if(!expect(v.key()==key, "Expect key match "+key)) return false;
        bool result = v.getUint() == a;
        debug(name + "::test_func(): done (result:" + to_string(result) + ")...");
        return result;
    }

    bool main() {
        return expect(
                test_func("key", 1),
                "Given a name and value(1), expect ArbitraryKeyValue to properly configure") ||
               expect(
                       test_func("key", 0),
                       "Given a name and value(0), expect ArbitraryKeyValue to properly configure") ||
               expect(
                       test_func("key", -1),
                       "Given a name and value(-1), expect ArbitraryKeyValue to properly configure");
    }
};
