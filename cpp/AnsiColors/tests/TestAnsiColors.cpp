/**
 * @name AnsiColors/tests/TestBasics.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include "projects/Tester/TestBase/main.h"

#include <iostream>
#include "projects/AnsiColors/AnsiColors.h"

class TestAnsiColors : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestAnsiColors(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestAnsiColors() {}

    /**
     * @name test_MAX_CLIENT_TIMEOUT
     * @brief validate that MAX_CLIENT_TIMEOUT is set as expected
     */
    bool test(string a, string b) {
        return a.compare(b) == 0;
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test(COLOR_BLACK, "\033[1;30m"), "COLOR_BLACK ok") &&
               expect(test(COLOR_RED, "\033[1;31m"), "COLOR_RED ok") &&
               expect(test(COLOR_GREEN, "\033[1;32m"), "COLOR_GREEN ok") &&
               expect(test(COLOR_YELLOW, "\033[1;33m"), "COLOR_YELLOW ok") &&
               expect(test(COLOR_BLUE, "\033[1;34m"), "COLOR_BLUE ok") &&
               expect(test(COLOR_MAGENTA, "\033[1;35m"), "COLOR_MAGENTA ok") &&
               expect(test(COLOR_CYAN, "\033[1;36m"), "COLOR_CYAN ok") &&
               expect(test(COLOR_WHITE, "\033[1;37m"), "COLOR_WHITE ok") &&
               expect(test(COLOR_GRAY, "\033[38;5;245m"), "COLOR_GRAY ok") &&
               expect(test(COLOR_DEFAULT, "\033[1;39m"), "COLOR_DEFAULT ok") &&
               expect(test(COLOR_RESET, "\033[0m"), "COLOR_RESET ok");

    }

};