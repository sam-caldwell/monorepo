/**
 * @name ArbitraryValue/tests/TestArbitraryValueSetStringGetString.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <ctime>
#include <iostream>
#include <unistd.h>
#include <iostream>
#include "projects/graph/ArbitraryValue/src/ArbitraryValue.h"
#include "projects/common/types/ValueTypes.h"

#define CHARS "`~1!2@3#4$5%6^7&8*9(0)-_=+qQwWeErRtTyYuUiIoOpP[{]}\\|aAsSdDfFgGhHjJkKlL;:'\"zZxXcCvVbBnNmM,<.>/? "

class TestArbitraryValueSetStringGetString : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueSetStringGetString(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueSetStringGetString() {}

    /**
     * @name test_fuzz_length
     * @brief continually increase string length with random strings and
     *        expect input to be returned correctly.
     *
     * @param numCycles int (number of passes)
     * @param length int (length of the random string)
     * @return bool
     */
    bool test_fuzz_length(int numCycles, int endLength, int length = 100) {
        for (int cycles = 0; cycles < numCycles; cycles++) {
            string word = [](int sz) -> string {
                const char keyspace[73] = {'a', 'b', 'c', 'd', 'e', 'f', 'g',
                                           'h', 'i', 'j', 'k', 'l', 'm', 'n',
                                           'o', 'p', 'q', 'r', 's', 't', 'u',
                                           'v', 'w', 'x', 'y', 'z', '0', '1',
                                           '2', '3', '4', '5', '6', '7', '8',
                                           '9', '0', '~', '!', '@', '#', '$',
                                           '%', ',', '^', ' ', '&', '*', '(',
                                           ')', '_', '+', '-', '=', '{', '}',
                                           '|', '[', ']', '\'', '"', ':', ';',
                                           '<', '>', '?', ',', '.', '/', '\\',
                                           0x00, 0x12, 0x29};
                string result = "";
                for (int i = 0; i < sz; i++)
                    result += keyspace[rand() % (sizeof(keyspace) - 1)];
                return result;
            }(length);
            ArbitraryValue t;
            t.setString(word);
            if (!expect(t.getString() == word, "Expected random string #" + to_string(cycles) + " ok"))
                return false;
        }
        return true;
    }

    /**
     * @name test_strong_type
     * @brief Test strongTyping enforcement where we expect exception.
     *
     * @param v string
     * @return bool
     */
    bool test_strong_type(string v) {
        debug(name + "::test_strong_type(" + v + ") start");
        ArbitraryValue t;
        t.setString(v);
        try {
            t.getString(true);
            return true;
        } catch (runtime_error &e) {
            return false;
        }
    }

    /**
     * @name test_func
     * @brief Given an ArbitraryValue class instance (t) with a null (default state),
     *        we expect we can use the appropriate getter/setter methods to set and
     *        retrieve a given value of a given type.
     *
     * @param v string
     * @return bool (result of test)
     */
    bool test_func(string v) {
        ArbitraryValue t;
        return expect(t.getType() == Null, "Expect default type Null") &&
               expect(t.getString() == "", "Expect default (loose-typed) value ''.") &&
               expect(t.setString(v), "Expect setter ok") &&
               expect(t.getType() == String, "Expect type String") &&
               expect(t.getString() == v, "Expect string value returned (loose-typing)") &&
               expect(t.getString(true) == v, "Expect string value returned (strong-typing)");
    }

/**
 * @name main()
 * @brief coordinate test execution.
 *
 * @return bool
 */
    bool main() {
        return expect(test_func("true"), "Given string 'true' expect matching result.") &&
               expect(test_func("false"), "Given string 'false' expect matching result.") &&
               expect(test_func(""), "Given string '' expect matching result.") &&
               expect(test_func(CHARS), "Expect matching response for string containing all printable chars.") &&
               expect(test_fuzz_length(100000, 100000), "random char strings ok");
    }

};