/**
 * @name Parsers/tests/TestParserTokenIterator.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "projects/Tester/TestBase/main.h"
#include "projects/common/Parser/src/Parser.h"
#include "projects/common/exceptions/exceptions.h"

class TestParserTokenIterator : public TestBase {
private:
    const string GoodYaml = "data_good.yaml";
    Parser *p;
public:
    /**
     * @name Class constructor
     * @brief initialize test class
     *
     * @param string n : test name
     */
    TestParserTokenIterator(string n) {
        name = n;
        p = new Parser(GoodYaml, Yaml);
    }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserTokenIterator() {
        if (p) delete p;
    }

    bool test_parse() {
        p->parse();
        return true;
    }

    bool test_initIterator() {
        debug("test_initIterator(): reset iterator");
        p->initIterator();
        return true;
    }

    bool test_iterator_loop() {
        debug("test_iterator_loop(): test_iterator_loop start");
        try {
            while (!p->endIterator()) {
                pair <string, string> token = p->next();
                debug("token(key):'" + token.first + "', token(value):'" + token.second + "'");
            }
            debug("get_value(): test_iterator_loop ok");
            return true;
        } catch (out_of_range &e) {
            return expect(false, "expect no exception. " + string(e.what()));
        }
    }

    string get_value() {
        try {
            pair <string, string> token = p->next();
            string v = token.first + ":" + token.second;
            return v;
        } catch (exception &e) {
            error(string(e.what()));
            return "";
        }
    }

    bool test_iterator_elements() {
        const string multilineA = "This is where we are gonna go wild # but we should never see # a comment in this multiline.";
        const string multilineB = "Another multiline.";
        p->initIterator();
        return expect(get_value().compare("myKey:value") == 0, "line 01 ok") &&
               expect(get_value().compare("secA.[0]:item0") == 0, "line 02 ok") &&
               expect(get_value().compare("secA.[1]:item1") == 0, "line 03 ok") &&
               expect(get_value().compare("secA.[2]:item2") == 0, "line 04 ok") &&
               expect(get_value().compare("secA.[3]:item3") == 0, "line 05 ok") &&
               expect(get_value().compare("secB.[0].key0:value0") == 0, "line 06 ok") &&
               expect(get_value().compare("secB.[1].key1:value1") == 0, "line 07 ok") &&
               expect(get_value().compare("secB.[2].key2:value2") == 0, "line 08 ok") &&
               expect(get_value().compare("secC.obj1.propA:0") == 0, "line 09 ok") &&
               expect(get_value().compare("secC.obj1.propB:1") == 0, "line 10 ok") &&
               expect(get_value().compare("secC.obj1.propC:2") == 0, "line 11 ok") &&
               expect(get_value().compare("secC.obj2.propD:0") == 0, "line 12 ok") &&
               expect(get_value().compare("secC.obj2.propE:1") == 0, "line 13 ok") &&
               expect(get_value().compare("secC.obj2.propF:2") == 0, "line 14 ok") &&
               expect(get_value().compare("secD.objD1.objD11.objD111.propG:0") == 0, "line 15 ok") &&
               expect(get_value().compare("secD.objD1.objD11.objD111.propH:1") == 0, "line 16 ok") &&
               expect(get_value().compare("secD.objD1.objD11.objD112.propI:0") == 0, "line 17 ok") &&
               expect(get_value().compare("secD.objD1.objD11.objD112.propJ:1") == 0, "line 18 ok") &&
               expect(get_value().compare("secE.multilineTextA:" + multilineA) == 0, "line 19 ok") &&
               expect(get_value().compare("secE.multilineTextB:" + multilineB) == 0, "line 20 ok");
    }

    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        return expect(p->count() == 0, "Expect count()==20 ok.  count:" + to_string(p->count())) &&
               expect(test_parse(), "expect parse() ok.") &&
               expect(test_initIterator(), "expect test_initIterator() ok") &&
               expect(test_iterator_loop(), "expect test_iterator_loop() ok") &&
               expect(test_iterator_elements(), "expect test_iterator_elements ok");
    }
};
