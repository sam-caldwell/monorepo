/**
 * @name networking/TcpCommon/tests/TestPipeResponse.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "../src/PipeResponse/main.h"

#include <iostream>
#include "../src/PipeResponse/main.h"

class TestPipeResponse : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestPipeResponse(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestPipeResponse() {}

    /**
     * @name test_default
     *
     * @return bool
     */
    bool test_default() {
        networking::PipeResponse p;
        return expect(p.message().compare("") == 0, "expect default constructor yields empty('') message.") &&
               expect(!p.isSuccessful(), "expect default constructor yields default false from isSuccessful()");
    }

    bool test_pipe_response(bool sf, string m) {
        networking::PipeResponse p(sf, m);
        return expect(p.isSuccessful() == sf, "expect isSuccessful() ok") &&
               expect(p.message().compare(m) == 0, "expect message ok");
    }

    /**
     * @name test_message
     *
     * @return bool
     */
    bool test_message() {
        networking::PipeResponse p;
        return expect(p.failure("foo").message().compare("foo") == 0, "expect failure message ok") &&
               expect(p.success("bar").message().compare("bar") == 0, "expect success message ok");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_default(), "expect test_default() ok") &&
               expect(test_pipe_response(true, "foo"), "expect test_pipe_response(true, 'foo1') ok") &&
               expect(test_pipe_response(false, "bar"), "expect test_pipe_response(false, 'bar2') ok") &&
               expect(test_message(), "expect test_message() ok");
    }

};