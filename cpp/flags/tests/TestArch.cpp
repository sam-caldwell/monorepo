/**
 * @name flags/tests/TestArch.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/common/flags/flags.h"

class TestArch : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestArch(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestArch() {/*noop*/}

    /**
     * @name test_arch
     * @brief test the architecture values
     *
     * @params a - int
     * @params b - int
     * @return bool
     */
    bool test_arch(int a, int b) { return a == b; }

    /**
     * @name test_detected_arch
     * @brief test the detected architecture is one of the expected values.
     *
     * @return bool
     */
    bool test_detected_arch() {
        return (CPU_ARCH == ARCH_X32) ||
               (CPU_ARCH == ARCH_X64) ||
               (CPU_ARCH == ARCH_ARM32) ||
               (CPU_ARCH == ARCH_ARM64);
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_arch(ARCH_X32, 0), "Expect test_arch ARCH_X32(0) ok") &&
               expect(test_arch(ARCH_X64, 1), "Expect test_arch ARCH_x64(1) ok") &&
               expect(test_arch(ARCH_ARM32, 2), "Expect test_arch ARCH_ARM32(2) ok") &&
               expect(test_arch(ARCH_ARM64, 3), "Expect test_arch ARCH_ARM64(3) ok") &&
               expect(test_detected_arch(), "Expect detected arch ok");
    }

};