/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKeyValue/src/ArbitraryKeyValue.h"

class TestArbitraryKeyValueDouble : public TestBase {
public:
    TestArbitraryKeyValueDouble(string n) { name = n; }

    /**
     * Given a boolean value (a), a name "testKey" and an ArbitraryKeyValue,
     * - Expect the name is properly set.
     * - Expect the value (a) is properly set.
     * - Expect the ArbitraryKeyValue is of type Bool.
     *
     * @param a double
     * @return bool
     */
    bool test_func(string key, double a) {
        debug(name + "::test_func(): starting...");
        ArbitraryKeyValue v(key, a);
        if(!expect(v.getType() == Double, "Expect Type Double")) return false;
        if(!expect(v.key()==key, "Expect key match "+key)) return false;
        bool result = v.getDouble() == a;
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
