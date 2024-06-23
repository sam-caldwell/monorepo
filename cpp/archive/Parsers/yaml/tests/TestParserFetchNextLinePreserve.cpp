/**
 * @name Parsers/tests/TestParserFetchNextLine.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../../Tester/Tester/main.h"
#include "../src/fetchNextLine.cpp"
#include "../../../exceptions/exceptions.h"

class TestParserFetchNextLinePreserve : public TestBase {
private:
    const string GoodYamlFile = "data_good.yaml";
    Parser *p;
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserFetchNextLinePreserve(string n) {
        name = n;
        p = new Parser(GoodYamlFile, Yaml);
    }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserFetchNextLinePreserve() {
        if (p) delete p;
    }

    bool test_fetch_line(string s, int lno, bool getNextLine) {
        debug(name + "::test_fetch_line() getNextLine:" + to_string(getNextLine));
        if(!getNextLine)
            p->preserveLine();
        else
            p->nextLine();
        return expect(p->fetchNextLine() == getNextLine,
                      "expect (" + to_string(p->getLine()) + ") fetchNextLine() ok") &&
               expect(p->getLine() == lno,
                      "expect line number " + to_string(p->getLine()) + ":" + to_string(lno));
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        const string line1 = "---"; //comments on line 1-3 ignored
        const string line2 = "sectionA:  ";
        const string line3 = "  - item1";
        const string line4 = "  - item2";
        const string line5 = "  - item3";
        const string line6 = "  - item4";
        return expect(test_fetch_line(line1, 4, true), "Expect test_fetch_line() ok(1)") &&
               expect(test_fetch_line(line1, 4, false), "Expect test_fetch_line() ok(2)") &&

               expect(test_fetch_line(line2, 5, true), "Expect test_fetch_line() ok(3)") &&
               expect(test_fetch_line(line3, 6, true), "Expect test_fetch_line() ok(4)") &&
               expect(test_fetch_line(line3, 6, false), "Expect test_fetch_line() ok(5)") &&
               expect(test_fetch_line(line4, 7, true), "Expect test_fetch_line() ok(6)") &&
               expect(test_fetch_line(line4, 7, false), "Expect test_fetch_line() ok(7)");
    }
};
