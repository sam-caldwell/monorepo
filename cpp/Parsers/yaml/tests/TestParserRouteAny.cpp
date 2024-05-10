/**
 * @name Parsers/tests/TestParserRouteAny.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Parser/src/Parser.h"
#include "projects/common/exceptions/exceptions.h"

class TestParserRouteAny : public TestBase {
private:
    const string GoodMultiLineTag = "multiLine.yaml";
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserRouteAny(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserRouteAny() {}

    bool test_tag_multiline() {
        Parser p(GoodMultiLineTag, Yaml);
        p.nextLine();

        return expect(p.fetchNextLine(), "Expect fetchNextLine() ok(1)") &&
               expect(p.routeAny(), "Expect multi-line tag ok(2).") &&
               expect(p.getType() == MultiLineTag, "Expect token type MultiLineTag(3)") &&
               expect(!p.fetchNextLine(), "Expect fetchNextLine() ok(4)") &&
               expect(p.routeAny(), "Expect multi-line value ok(5)") &&
               expect(p.getType() == MultiLineTag, "Expect token type MultiLineTag(6)");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_tag_multiline(), "Expect test_tag_multiline() ok");
    }
};