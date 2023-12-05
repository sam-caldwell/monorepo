/**
 * @name Parser/tests/TestParserRouteKeyValue.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Parser/src/Parser.h"
#include "projects/common/exceptions/exceptions.h"

class TestParserRouteKeyValue : public TestBase {
private:
    const string GoodYamlFile = "data_good.yaml";
    const string BadYamlFile = "data_bad.yaml";
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserRouteKeyValue(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserRouteKeyValue() {}

    bool test_find_routeKeyValue() {
        Parser p(GoodYamlFile, Yaml);
        p.nextLine();

        return expect(p.getType() == Header, "Expect header as first expected token type") &&
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeHeader(), "Expect routeHeader() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected") &&
               expect(p.routeKeyValueTag(), "Expect routeKeyValueTag() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_find_routeKeyValue(), "Expect test_find_routeKeyValue() ok");
    }
};