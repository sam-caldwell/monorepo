/**
 * @name Parser/tests/TestParserStripComment.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Parser/src/Parser.h"
#include "projects/common/exceptions/exceptions.h"

class TestParserStripComment : public TestBase {
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserStripComment(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserStripComment() {}

    bool fetch_strip_comments(bool stripOk, string *line, string expected) {
        debug(name + "::test_fetch_next_line(): ln='" + *line + "'");

        if (stripComment(line) != stripOk) return false;

        debug("line: '" + *line +"'");

        return expect(line->compare(expected) == 0,
                      "stripComment('" + *line + "','" + expected + "','" + to_string(line->compare(expected)) + "')");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        string ln1 = "# file contains 33 objects (not including comments)";
        string ln2 = "# comment line A";
        string ln3 = "# comment line B";
        string ln4 = "---";
        string ln5 = "sectionA:  # this is a list";
        string ln6 = "aValidEntry: entry";
        string ln7 = "  - item  #foo";
        string ln8 = "  - item";
        return expect(fetch_strip_comments(true, &ln1, ""), "Expect fetch_strip_comments() ok") &&
               expect(fetch_strip_comments(true, &ln2, ""), "Expect fetch_strip_comments() ok") &&
               expect(fetch_strip_comments(true, &ln3, ""), "Expect fetch_strip_comments() ok") &&
               expect(fetch_strip_comments(false, &ln4, "---"), "Expect fetch_strip_comments() ok") &&
               expect(fetch_strip_comments(true, &ln5, "sectionA:"), "Expect fetch_strip_comments() ok") &&
               expect(fetch_strip_comments(false, &ln6, "aValidEntry: entry"), "Expect fetch_strip_comments() ok") &&
               expect(fetch_strip_comments(true, &ln7, "- item"), "Expect fetch_strip_comments() ok(7)") &&
               expect(fetch_strip_comments(false, &ln8, "  - item"), "Expect fetch_strip_comments() ok(8)");
    }
};
