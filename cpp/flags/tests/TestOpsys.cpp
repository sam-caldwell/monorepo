/**
 * @name flags/tests/TestOpsys.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/common/flags/flags.h"

class TestOpsys : public TestBase {
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestOpsys(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestOpsys() {/*noop*/}

    /**
     * @name test_opsys
     * @brief test the operating system values
     *
     * @params a - int
     * @params b - int
     * @return bool
     */
    bool test_opsys(int a, int b) { return a == b; }

    /**
     * @name test_detected_opsys
     * @brief test the detected opsys is one of the expected values.
     *
     * @return bool
     */
    bool test_detected_opsys() {
        return (OPSYS == OPSYS_WINDOWS) ||
               (OPSYS == OPSYS_LINUX) ||
               (OPSYS == OPSYS_MACOS);
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_opsys(OPSYS_WINDOWS, 0), "Expect test_arch ARCH_X32(0) ok") &&
               expect(test_opsys(OPSYS_LINUX, 1), "Expect test_arch ARCH_x64(1) ok") &&
               expect(test_opsys(OPSYS_MACOS, 2), "Expect test_arch ARCH_ARM32(2) ok") &&
               expect(test_detected_opsys(), "Expect detected opsys ok");
    }
};