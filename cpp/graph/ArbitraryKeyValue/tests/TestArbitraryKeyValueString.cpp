/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKeyValue/src/ArbitraryKeyValue.h"

class TestArbitraryKeyValueString : public TestBase {
public:
    TestArbitraryKeyValueString(string n) { name = n; }

    /**
     * Given a boolean value (a), a name "testKey" and an ArbitraryKeyValue,
     * - Expect the name is properly set.
     * - Expect the value (a) is properly set.
     * - Expect the ArbitraryKeyValue is of type Bool.
     *
     * @param a string
     * @return bool
     */
    bool test_func(string key, string a) {
        debug(name + "::test_func(): starting...");
        ArbitraryKeyValue v(key, a);
        if(!expect(v.getType() == String, "Expect Type String")) return false;
        if(!expect(v.key()==key, "Expect key match "+key)) return false;
        bool result = v.getString() == a;
        debug(name + "::test_func(): done (result:" + to_string(result) + ")...");
        return result;
    }

    bool main() {
        return expect(
                test_func("key", "test"),
                "Given a name and value('1'), expect ArbitraryKeyValue to properly configure");
    }
};
