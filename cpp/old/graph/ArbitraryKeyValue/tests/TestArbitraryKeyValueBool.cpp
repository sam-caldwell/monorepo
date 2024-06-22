/**
 * @name tests/TestArbitraryKeyValueBool.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @brief test an arbitrary key value object's boolean getter
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryKeyValue.h"

class TestArbitraryKeyValueBool : public TestBase {
public:
    TestArbitraryKeyValueBool(string n) { name = n; }

    /**
     * @brief Given a boolean value (a), a name "testKey" and an ArbitraryKeyValue,
     * @note Expect the name is properly set.
     * @note Expect the value (a) is properly set.
     * @note Expect the ArbitraryKeyValue is of type Bool.
     *
     * @param a bool
     * @return bool
     */
    bool test_func(string key, bool a) {
        debug(name + "::test_func(): starting...");
        ArbitraryKeyValue v(key, a);
        if(!expect(v.getType() == Bool, "Expect Type Bool")) return false;
        if(!expect(v.key()==key, "Expect key match "+key)) return false;
        bool result = v.getBool() == a;
        debug(name + "::test_func(): done (result:" + to_string(result) + ")...");
        return result;
    }

    bool main() {
        return expect(test_func("key", true),
                      "Given a name and value(true), expect ArbitraryKeyValue to properly configure") ||
               expect(test_func("key", false),
                      "Given a name and value(false), expect ArbitraryKeyValue to properly configure");
    }
};
