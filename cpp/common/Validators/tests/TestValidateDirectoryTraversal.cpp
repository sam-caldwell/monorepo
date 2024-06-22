/**
 * @name Validators/tests/TestValidateDirectoryTraversal.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "../../networking/TcpCommon/src/PipeResponse/main.h"
#include "../validators.h"

class TestValidateDirectoryTraversal : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestValidateDirectoryTraversal(string n) { name = n; }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestValidateDirectoryTraversal() {}

    /**
     * @name test_directoryTraversal
     * @brief test the directoryTraversal() function
     *
     * @return bool
     */
    bool test_directoryTraversal(string s, bool expected) {
        return expect(directoryTraversal(&s) == expected, "expect directory traversal:'" + s + "'");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_directoryTraversal("..", true), "scenario 1.1 ok") &&
               expect(test_directoryTraversal("%2e%2e", true), "scenario 1.2 ok") &&
               expect(test_directoryTraversal("../..", true), "scenario 1.3 ok") &&
               expect(test_directoryTraversal("/a/b/../..", true), "scenario 1.4 ok") &&
               expect(test_directoryTraversal("%2e%2e/%2e%2e", true), "scenario() 1.5 ok") &&
               expect(test_directoryTraversal("/a/b/%2e%2e/%2e%2e", true), "scenario() 1.6 ok") &&
               expect(test_directoryTraversal("/a/b/../../foo.txt", true), "scenario 1.7 ok");
    }

};