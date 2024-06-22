/**
 * @name CommandLineParser/tests/TestStringToSet.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <set>
#include "../formatter.h"

class TestStringToSet : public TestBase {
private:
    const string data = "A,B,C,D,E,F";
public:
    /**
     * @name class constructor
     * @brief initialize the test
     *
     * @param n string
     */
    TestStringToSet(string n) {
        debug(n + "::TestParse()");
        name = n;
    }

    /**
     * @name class destructor
     * @brief clean up the test.
     */
    ~TestStringToSet() {
        debug(name + "::~TestParse()");
    }

    /**
     * @name test_deserialize
     * @brief test the initial state post-construction.
     *
     * @return bool
     */
    bool test_deserialize() {
        set<string> s = stringToSet(data);
        return expect(s.find("A")!=s.end(), "Expect A") &&
               expect(s.find("B")!=s.end(), "Expect B") &&
               expect(s.find("C")!=s.end(), "Expect C") &&
               expect(s.find("D")!=s.end(), "Expect D") &&
               expect(s.find("E")!=s.end(), "Expect E") &&
               expect(s.find("F")!=s.end(), "Expect F");
    }

    /**
     * @name main
     * @brief coordinate tests.
     *
     * @return bool
     */
    bool main() {
        return expect(test_deserialize(), "expect test_serialize() ok");
    }
};