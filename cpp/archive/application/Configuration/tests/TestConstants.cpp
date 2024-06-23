/**
 * @name Configuration/tests/TestConstants.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../src/Configuration.h"

class TestConstants : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestConstants(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestConstants() {}

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(ConfigLogLevel.compare("log.level")==0, "log.level ok") &&
               expect(ConfigLogOutput.compare("log.output")==0, "log.output ok") &&
               expect(ConfigLogTags.compare("log.tags")==0, "log.tags ok") &&
               expect(ConfigLogDestination.compare("log.destination")==0, "log.destination ok");
    }
};
