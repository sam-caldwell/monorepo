/**
 * @name ArbitraryValue/tests/TestArbitraryValueType.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <string>
#include <iostream>
#include "../../../../types/ValueTypes.h"


class TestArbitraryValueType : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestArbitraryValueType(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestArbitraryValueType() {}

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        ValueTypes a[7] = {
                Bool,
                Double,
                Float,
                Int,
                Null,
                String,
                Uint
        };
        bool outcome = true;
        for (int i = 0; (i < 7) && outcome; i++) {
            ValueTypes b = a[i];
            outcome = expect(b == a[i], "Type mismatch " + to_string(i));
        }
        return outcome;
    }
};
