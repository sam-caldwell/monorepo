/**
 * @name CommandLineParser/tests/TestSetToString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/common/formatter/formatter.h"

class TestSetToString : public TestBase {
private:
    set <string> *data;
public:
    /**
     * @name class constructor
     * @brief initialize the test
     *
     * @param n string
     */
    TestSetToString(string n) {
        debug(n + "::TestParse()");
        name = n;
        data = new set<string>;
        data->insert("A");
        data->insert("B");
        data->insert("C");
        data->insert("D");
        data->insert("E");
    }

    /**
     * @name class destructor
     * @brief clean up the test.
     */
    ~TestSetToString() {
        debug(name + "::~TestParse()");
        if (data) delete data;
    }

    /**
     * @name test_initial_state
     * @brief test the initial state post-construction.
     *
     * @return bool
     */
    bool test_serialize() {
        return expect(setToString(data) == "\"A\",\"B\",\"C\",\"D\",\"E\"", "expect serialize ok");
    }

    /**
     * @name main
     * @brief coordinate tests.
     *
     * @return bool
     */
    bool main() {
        return expect(test_serialize(), "expect test_serialize() ok");
    }
};