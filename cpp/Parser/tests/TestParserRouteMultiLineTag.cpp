/**
 * @name Parser/tests/TestParserRouteMultiLineTag.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Parser/src/Parser.h"
#include "projects/common/exceptions/exceptions.h"

class TestParserRouteMultiLineTag : public TestBase {
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
    TestParserRouteMultiLineTag(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserRouteMultiLineTag() {}

    bool test_find_routeMultiLineTag() {
        Parser p(GoodYamlFile, Yaml);
        p.nextLine();

        return expect(p.getType() == Header, "Expect header as first expected token type") &&
               expect(p.fetchNextLine(), "01. Expect fetchNextLine() ok") && // line 4
               expect(p.routeHeader(), "Expect routeHeader() ok") &&

               expect(p.fetchNextLine(), "02. Expect fetchNextLine() ok") && // line 5
               expect(p.getLine() == 5, "Expect line position verified") &&
               expect(p.fetchNextLine(), "03. Expect fetchNextLine() ok") && // line 6
               expect(p.fetchNextLine(), "04. Expect fetchNextLine() ok") && // line 7
               expect(p.fetchNextLine(), "05. Expect fetchNextLine() ok") && // line 8
               expect(p.fetchNextLine(), "06. Expect fetchNextLine() ok") && // line 9
               expect(p.fetchNextLine(), "07. Expect fetchNextLine() ok") && // line 10
               expect(p.fetchNextLine(), "08. Expect fetchNextLine() ok") && // line 11
               expect(p.fetchNextLine(), "09. Expect fetchNextLine() ok") && // line 13
               expect(p.fetchNextLine(), "10. Expect fetchNextLine() ok") && // line 14
               expect(p.fetchNextLine(), "11. Expect fetchNextLine() ok") && // line 15
               expect(p.fetchNextLine(), "12. Expect fetchNextLine() ok") && // line 17
               expect(p.fetchNextLine(), "13. Expect fetchNextLine() ok") && // line 19
               expect(p.fetchNextLine(), "14. Expect fetchNextLine() ok") && // line 20
               expect(p.fetchNextLine(), "15. Expect fetchNextLine() ok") && // line 21
               expect(p.fetchNextLine(), "16. Expect fetchNextLine() ok") && // line 22
               expect(p.fetchNextLine(), "17. Expect fetchNextLine() ok") && // line 23
               expect(p.fetchNextLine(), "18. Expect fetchNextLine() ok") && // line 24
               expect(p.fetchNextLine(), "17. Expect fetchNextLine() ok") && // line 25
               expect(p.fetchNextLine(), "18. Expect fetchNextLine() ok") && // line 26
               expect(p.fetchNextLine(), "19. Expect fetchNextLine() ok") && // line 28
               expect(p.fetchNextLine(), "20. Expect fetchNextLine() ok") && // line 29
               expect(p.fetchNextLine(), "21. Expect fetchNextLine() ok") && // line 31
               expect(p.fetchNextLine(), "22. Expect fetchNextLine() ok") && // line 33
               expect(p.fetchNextLine(), "23. Expect fetchNextLine() ok") && // line 35
               expect(p.fetchNextLine(), "24. Expect fetchNextLine() ok") && // line 37
               expect(p.fetchNextLine(), "25. Expect fetchNextLine() ok") && // line 38
               expect(p.fetchNextLine(), "26. Expect fetchNextLine() ok") && // line 39
               expect(p.fetchNextLine(), "27. Expect fetchNextLine() ok") && // line 40
               expect(p.fetchNextLine(), "28. Expect fetchNextLine() ok") && // line 41
               //
               //multilineTextA: |  # A comment about our multiline.
               //
               expect(p.fetchNextLine(), "29. Expect fetchNextLine() ok") && // line 42
               expect(p.getLine() == 42, "    Expect line position 42 verified:" + to_string(p.getLine())) &&

               expect(p.fetchNextLine(), "30. Expect fetchNextLine() ok") && // line 43
               expect(p.fetchNextLine(), "30. Expect fetchNextLine() ok") && // line 44
               expect(p.fetchNextLine(), "30. Expect fetchNextLine() ok") && // line 45
               expect(p.getLine() == 45, "    Expect line position 45 verified:" + to_string(p.getLine())) &&
               //
               // multilineTextB: |
               //
               expect(p.fetchNextLine(), "31. Expect fetchNextLine() ok") && // line 46
               expect(p.getLine() == 46, "    Expect line position 46 verified:" + to_string(p.getLine())) &&
               expect(p.fetchNextLine(), "32. Expect fetchNextLine() ok") && // line 47
               expect(p.getLine() == 47, "    Expect line position 47 verified:" + to_string(p.getLine()));
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(test_find_routeMultiLineTag(), "Expect test_find_routeMultiLineTag() ok");
    }
};