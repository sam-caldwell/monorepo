/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "projects/graph/ArbitraryKeyValue/src/ArbitraryKeyValue.h"

class TestArbitraryKeyValue : public TestBase {
public:
    TestArbitraryKeyValue(string n) { name = n; }

    bool test_av_kv_default() {
        if (debug_flag) cout << "test_arbitrary_kv_default() Start" << endl;
        ArbitraryKeyValue v;
        return v.getType() == Null;
    }

    bool test_av_kv_with_name() {
        ArbitraryKeyValue v("testKey");
        return (v.getType() == Null);
    }

    bool test_av_kv_correct_name(){
        ArbitraryKeyValue v("testKey");
        return (v.getType() == Null) && (v.key() == "testKey");
    }

    bool main() {
        return expect(test_av_kv_default(), "Expect ArbitraryValue will instantiate without error") ||
               expect(test_av_kv_with_name(),"Expect ArbitraryKeyValue will set name with no value.") ||
               expect(test_av_kv_correct_name(), "Given a key (name), Expect ArbitraryKeyValue to set.");
    }
};
