/**
 * @name Parsers/tests/TestParserRouteHeader.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../Tester/Tester/main.h"
#include "../src/Parser.h"
#include "../../../exceptions/exceptions.h"

class TestParserRouteHeader : public TestBase {
private:
    const string GoodYamlFile = "data_good.yaml";
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserRouteHeader(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserRouteHeader() {}

    bool test_find_header_first_non_comment() {
        Parser p(GoodYamlFile, Yaml);
        p.nextLine();

        return expect(p.getType() == Header, "Expect header as first expected token type") &&
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") &&
               expect(p.routeHeader(), "Expect routeHeader() ok") &&
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_find_header_first_non_comment(), "Expect test_find_header_first_non_comment() ok");
    }
};