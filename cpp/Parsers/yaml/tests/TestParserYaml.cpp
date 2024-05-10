/**
 * @name Parsers/tests/TestParserYaml.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include "../../../Tester/Tester/main.h"
#include "../src/Parser.h"
#include "../../../exceptions/exceptions.h"

class TestParserYaml : public TestBase {
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
    TestParserYaml(string n) { name = n; }

    /**
     * @name Class destructor
     * @brief clean up
     */
    ~TestParserYaml() {}


    /**
     * @name main()
     * @brief coordinate test execution.
     *
     * @return bool
     */
    bool main() {
        debug(name + "::testParseComments()");
        Parser p(GoodYamlFile);
        p.parse();
        try {
            return expect(p.getToken("myKey") == "value", "expect myKey ok") &&
                   expect(p.getToken("secA.[0]") == "item0", "expect secA.[0] == item0 ok") &&
                   expect(p.getToken("secA.[1]") == "item1", "expect secA.[1] == item1 ok") &&
                   expect(p.getToken("secA.[2]") == "item2", "expect secA.[2] == item2 ok") &&
                   expect(p.getToken("secA.[3]") == "item3", "expect secA.[3] == item3 ok") &&

                   expect(p.getToken("secB.[0].key0") == "value0", "expect secB.[0].key0 == value0 ok") &&
                   expect(p.getToken("secB.[1].key1") == "value1", "expect secB.[1].key1 == value1 ok") &&
                   expect(p.getToken("secB.[2].key2") == "value2", "expect secB.[2].key2 == value2 ok") &&

                   expect(p.getToken("secC.obj1.propA") == "0", "expect secC.obj1.propA ok") &&
                   expect(p.getToken("secC.obj1.propB") == "1", "expect secC.obj1.propB ok") &&
                   expect(p.getToken("secC.obj1.propC") == "2", "expect secC.obj1.propC ok") &&

                   expect(p.getToken("secC.obj2.propD") == "0", "expect secC.obj2.propD ok") &&
                   expect(p.getToken("secC.obj2.propE") == "1", "expect secC.obj2.propE ok") &&
                   expect(p.getToken("secC.obj2.propF") == "2", "expect secC.obj2.propF ok") &&

                   expect(p.getToken("secD.objD1.objD11.objD111.propG") == "0",
                          "expect secD.objD1.objD11.objD111.propG ok") &&
                   expect(p.getToken("secD.objD1.objD11.objD111.propH") == "1",
                          "expect secD.objD1.objD11.objD111.propH ok") &&

                   expect(p.getToken("secD.objD1.objD11.objD112.propI") == "0",
                          "expect secD.objD1.objD11.objD112.propI ok") &&
                   expect(p.getToken("secD.objD1.objD11.objD112.propJ") == "1",
                          "expect secD.objD1.objD11.objD112.propJ ok") &&

                   expect(p.getToken("secE.multilineTextA") ==
                          "This is where we are gonna go wild # but we should never see # a comment in this multiline.",
                          "expect secE.multilineTextA ok.") &&
                   expect(p.getToken("secE.multilineTextB") == "Another multiline.",
                          "expect secE.multilineTextB ok. value:" + p.getToken("secE.multilineTextB"));
        } catch (out_of_range &e) {
            return expect(false, "Expect no exceptions.  error:" + string(e.what()));
        }
    }
};
