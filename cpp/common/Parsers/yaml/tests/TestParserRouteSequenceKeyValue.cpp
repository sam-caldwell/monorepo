/**
 * @name Parsers/tests/TestParserRouteSequenceKeyValue.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../../Tester/Tester/main.h"
#include "../src/Parser.h"
#include "../../../exceptions/exceptions.h"

class TestParserRouteSequenceKeyValue : public TestBase {
private:
    const string GoodYamlFile = "data_good.yaml";
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserRouteSequenceKeyValue(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserRouteSequenceKeyValue() {}

    bool test() {
        Parser p(GoodYamlFile, Yaml);
        p.nextLine();

        return expect(p.getType() == Header, "Expect header as first expected token type") &&
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") && //line 04
               expect(p.routeHeader(), "Expect routeHeader() ok") &&

               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") && //line 05
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") && //line 06 (secA)
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") && //line 07
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") && //line 08
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") && //line 09
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") && //line 10 (- item3)
               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") && //line 11 (secB)

               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") && //line 13
               expect(p.getType() == Any, "Expect the expected token will shift to Any after header is detected 01") &&
               expect(p.routeSequenceKeyValue(), "Expect routeSequenceKeyValue() ok") &&
               expect(p.getType() == SequenceKeyValue,
                      "Expect the expected token will shift to Any after header is detected 02") &&

               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") && //line 14
               expect(p.getType() == SequenceKeyValue,
                      "Expect the expected token will shift to Any after header is detected 03") &&
               expect(p.routeSequenceKeyValue(), "Expect routeSequenceKeyValue() ok") &&
               expect(p.getType() == SequenceKeyValue,
                      "Expect the expected token will shift to Any after header is detected 04") &&

               expect(p.fetchNextLine(), "Expect fetchNextLine() ok") && //line 15
               expect(p.getType() == SequenceKeyValue,
                      "Expect the expected token will shift to Any after header is detected 05") &&
               expect(p.routeSequenceKeyValue(), "Expect routeSequenceKeyValue() ok") &&
               expect(p.getType() == SequenceKeyValue,
                      "Expect the expected token will shift to Any after header is detected 06") &&
               expect(p.getLine() == 15, "Expect line 15 ok");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test(), "Expect test() ok");
    }
};