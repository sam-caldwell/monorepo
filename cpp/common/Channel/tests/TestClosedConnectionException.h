/**
 * @name Channel/tests/TestClosedConnectionException.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include <thread>
#include "../src/Channel.h"

using namespace std;

class TestClosedConnectionException : public TestBase {
public:
    TestClosedConnectionException(string n) { name = n; }

    ~TestClosedConnectionException() {}

    bool test_setup() {
        debug(name + "::test_setup() starting...");
        Channel<int> data;
        data.push(1);
        data.pop();
        try {
            data.close();
            data.push(1); //expect error.
            return false;
        } catch (closedChannelException) {
            return true;
        }
    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect push to closed channel will throw exception.");
    }
};