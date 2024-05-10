/**
 * @name Parsers/tests/TestParserRouteSequenceItem.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Parser/src/Parser.h"
#include "projects/common/exceptions/exceptions.h"

class TestParserRouteSequenceItem : public TestBase {
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
    TestParserRouteSequenceItem(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserRouteSequenceItem() {}

    bool test() {
        Parser p(GoodYamlFile, Yaml);
        p.nextLine();

        return expect(p.getType() == Header, "00. Expect header as first expected token type") &&
               expect(p.fetchNextLine(), "01. Expect fetchNextLine() ok line:" + to_string(p.getLine())) &&
               expect(p.routeHeader(), "01. Expect routeHeader() ok") &&
               expect(p.getType() == Any, "01. Expect the expected token will shift to Any after header is detected") &&

               expect(p.fetchNextLine(), "02. Expect fetchNextLine() ok line:" + to_string(p.getLine())) &&
               expect(p.routeTagMap(), "02. Expect routeTagMap() ok") &&
               expect(p.getType() == Any, "02. Expect the expected token will shift to Any after header is detected") &&

               expect(p.fetchNextLine(), "03. Expect fetchNextLine() ok line:" + to_string(p.getLine())) && //line 6
               expect(p.routeTagMap(), "03. Expect routeTagMap() ok") && //myKey: value
               expect(p.getType() == Any, "03. Expect Any type ok") &&

               expect(p.fetchNextLine(), "04. Expect fetchNextLine() ok line:" + to_string(p.getLine())) &&
               expect(p.routeSequenceItem(), "04. Expect routeSequenceItem() ok") && //item1
               expect(p.getType() == SequenceItem, "04. Expect sequenceItem type ok") &&

               expect(p.fetchNextLine(), "05. Expect fetchNextLine() ok line:" + to_string(p.getLine())) &&
               expect(p.routeSequenceItem(), "05. Expect routeSequenceItem() ok") && //item2
               expect(p.getType() == SequenceItem, "05. Expect sequenceItem type ok") &&

               expect(p.fetchNextLine(), "06. Expect fetchNextLine() ok line:" + to_string(p.getLine())) &&
               expect(p.routeSequenceItem(), "06. Expect routeSequenceItem() ok") && //item3
               expect(p.getType() == SequenceItem, "06. Expect sequenceItem type ok") &&

               expect(p.fetchNextLine(), "06. Expect fetchNextLine() ok line:" + to_string(p.getLine())) &&
               expect(p.routeSequenceItem(), "06. Expect routeSequenceItem() ok") && //item3
               expect(p.getType() == SequenceItem, "06. Expect sequenceItem type ok") &&

               expect(p.fetchNextLine(), "07. Expect fetchNextLine() ok line:" + to_string(p.getLine())) &&
               expect(p.getLine() == 11, "07. Expect to be at line 11 of source file. line:" + to_string(p.getLine()));
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