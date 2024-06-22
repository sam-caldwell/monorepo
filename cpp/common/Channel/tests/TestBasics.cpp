/**
 * @name common/Channel/tests/TestBasics.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <thread>
#include "../src/Channel.h"

using namespace std;

class TestBasics : public TestBase {

    bool status;

    Channel<int> data;

private:
public:

    TestBasics(string n) { name = n; }

    ~TestBasics() {}

    bool test_setup() {
        debug(name + "::test_setup() starting...");

        int expected = 1;
        int actual;

        if (!expect(data.empty(), "Expect channel to be empty at initialization.")) return false;

        if (!expect(data.size() == 0, "Expect size == 0 at initialization")) return false;

        debug("test closed method");

        if (!expect(!data.closed(), "Expect channel not closed")) return false;

        debug("pushing data onto channel");

        data.push(expected);

        debug("data has been pushed to channel");

        if (!expect(data.size() == 1, "Expect channel size ==1 after first push")) return false;

        debug("data being popped from channel");

        actual = data.pop();

        debug("popped value from channel (" + to_string(actual) + ")");

        if (!expect(data.size() == 0, "Expect size == 0 after pop first node.")) return false;

        return expect(actual == expected, "expect input to match output."
                                          " actual:" + to_string(actual) +
                                          " expected:" + to_string(expected));
    }

    bool test_size() {
        while (!data.empty()) { data.pop(); } /*Make sure the channel is clear*/
        for (int i = 1; i < 65535; i++) {
            data.push(i);
            if (!expect(data.size() == i, "Expect size matches index i:" + to_string(i)))
                return false;
        }
        return true;
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect some basic tests work without error.") &&
               expect(test_size(), "Expect size test to work without error.");
    }
};