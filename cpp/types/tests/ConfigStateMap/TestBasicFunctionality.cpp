/**
 * @name types/tests/TestBasicFunctionality.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/types/ConfigStateMap.h"

bool test_validator(string input) {
    return input == "Happy";
}

class TestBasicFunctionality : public TestBase {
private:
    ConfigStateMap *p;
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestBasicFunctionality(string n) {
        name = n;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestBasicFunctionality() {
        if (p) delete p;
    }

    /**
     * @name test_initialize_class
     * @brief just make sure we can instantiate ConfigStateMap without error
     *
     * @return bool (pass/fail)
     */
    bool test_initialize_class() {
        debug(name + "::test_initialize_class()");
        ConfigStateMap o;
        return true;
    }

    /**
     * @name test_add
     * @brief verify the ConfigStateMap::add() method works minimally.
     *
     * @return bool (pass/fail)
     */

    bool test_add() {
        debug(name + "::test_add_erase()");
        ConfigStateMap o;
        return expect(
                o.add("A", "a1", Bool, false, "An error", ConfigStateMapNullValidator) &&
                o.add("B", "b2", Bool) &&
                o.add("C", "c3", Double) &&
                o.add("D", "d4", Float) &&
                o.add("E", "e5", Int) &&
                o.add("F", "f6", String) &&
                o.add("G", "g7", Uint), "Expect we can create multiple map nodes.");
    }

    /**
     * @name test_erase
     * @brief verify ConfigStateMap::erase() works
     *
     * @return bool (pass/fail)
     */
    bool test_erase() {
        debug(name + "::test_add_erase()");
        ConfigStateMap o;
        return expect(
                o.add("A", "a1", Bool) &&
                o.add("B", "b2", Bool) &&
                o.add("C", "c3", Double) &&
                o.erase("A") &&
                o.erase("B") &&
                o.erase("C"), "Expect we can create and delete multiple map nodes.") &&
               expect(!o.erase("D"), "Expect that deleting non-existent node will return false");
    }

    /**
     * @name test_exists
     * @brief verify ConfigStateMap::exists() works
     *
     * @return bool (pass/fail)
     */
    bool test_exists() {
        debug(name + "::test_exists()");
        ConfigStateMap o;
        return expect(o.add("A", "a1", Bool), "Expect we can add 'A'") &&
               expect(o.exists("A"), "Expect that 'A' exists is true") &&
               expect(!o.exists("B"), "Expect that 'B' does not exist");
    }

    /**
     * @name test_rhs_exists
     * @brief verify ConfigStateMap::rhsExists() works
     *
     * @return bool (pass/fail)
     */
    bool test_rhs_exists() {
        debug(name + "::test_rhs_exists()");
        ConfigStateMap o;
        return expect(o.add("A", "a1", Bool), "Expect we can add 'A'") &&
               expect(o.rhsExists("a1"), "Expect that rhs 'a1' exists is true") &&
               expect(!o.rhsExists("b2"), "Expect that rhs 'b2' does not exist");
    }

    /**
     * @name test_get_required
     * @brief verify ConfigStateMap::getRequired() will work
     *
     * @return bool (pass/fail)
     */
    bool test_get_required() {
        debug(name + "::test_get_required");
        ConfigStateMap o;
        return expect(o.add("A", "a1", Int, true), "expect A is required.") &&
               expect(o.add("B", "b2", Int, false), "expect B is not required") &&
               expect(o.getRequired("A"), "Expect A is required (true)") &&
               expect(!o.getRequired("B"), "Expect B is required (false)");
    }

    /**
     * @name test_get_required_exception
     * @brief confirm that ConfigStateMap::getRequired will throw out_of_range if key does not exist.
     * @return bool (pass/fail)
     */
    bool test_get_required_exception() {
        debug(name + "::test_get_required_exception");
        ConfigStateMap o;
        try {
            o.getRequired("C");
            return false; //Expected exception.
        } catch (out_of_range) {
            return true;
        }
    }
    /**
     * @name test_get_type
     * @brief verify ConfigStateMap::getType() works
     *
     * @return bool (pass/fail)
     */
    bool test_get_type() {
        debug(name + "::test_get_type()");
        ConfigStateMap o;
        return expect(o.add("A", "a1", Bool), "Expect we can add 'A' as Bool") &&
               expect(o.getType("A") == Bool, "Expect .get() will return Bool");
    }

    /**
     * @name test_get_type_exception
     * @brief verify ConfigStateMap::getType() works
     *
     * @return bool (pass/fail)
     */
    bool test_get_type_exception() {
        debug(name + "::test_get_type_exception()");
        ConfigStateMap o;
        if (!expect(o.add("A", "a1", Bool), "Expect we can add 'A' as Bool")) return false; //Expect true
        try {
            o.getType("B"); //Expect an exception
            return false;
        } catch (out_of_range) {
            return true;
        }
    }

    /**
     * @name test_get_rhs
     * @brief verify ConfigStateMap::getRhsString() works
     *
     * @return bool (pass/fail)
     */
    bool test_get_rhs() {
        debug(name + "::test_get_rhs()");
        ConfigStateMap o;
        return expect(o.add("A", "a1", Int), "Expect we can add lhs A mapped to rhs '1'") &&
               expect(o.getRhsString("A") == "a1", "Expect getRhsString will return '1'");
    }

    /**
     * @name test_get_rhs_exception
     * @brief verify ConfigStateMap::getRhsString() works
     *
     * @return bool (pass/fail)
     */
    bool test_get_rhs_exception() {
        debug(name + "::test_get_rhs_exception()");
        ConfigStateMap o;
        if (!expect(o.add("A", "a1", Bool), "Expect we can add 'A' as Bool")) return false; //Expect true
        try {
            o.getRhsString("B"); //Expect an exception
            return false;
        } catch (out_of_range) {
            return true;
        }
    }

    /**
     * @name test_run_validate
     * @brief verify ConfigStateMap::validate() works
     *
     * @return bool (pass/fail)
     */
    bool test_run_validate() {
        debug(name + "::test_run_validate()");
        ConfigStateMap o;
        return expect(o.add("A", "a1", String, false, "errorA", ConfigStateMapNullValidator), "Expect add A OK") &&
               expect(o.add("B", "b2", String, false, "errorB", test_validator),
                      "Expect add B OK") &&
               expect(o.add("C", "c3", String, false, "errorC", test_validator),
                      "Expect add C OK") &&
               expect(o.add("emptyError", "ee", String, false),
                      "Expect empty error ok.") &&
               expect(o.validate("A", "Happy"), "Expect A Validates with noop") &&
               expect(o.validate("B", "Happy"), "Expect B validates with simple test") &&
               expect(o.validate("C", "Joy! Joy!"), "Expect C fails with simple test.");
    }

    /**
     * @name test_run_validate_missing
     * @brief verify ConfigStateMap::validate() works
     *
     * @return bool (pass/fail)
     */
    bool test_run_validate_missing() {
        debug(name + "::test_run_validate_missing() start");
        ConfigStateMap cfg;
        return expect(cfg.add("A", "a1", Bool), "Expect we can add 'A' as Bool") &&
               expect(cfg.count() == 1, "Expect count is 1.  Actual:" + to_string(cfg.count())) &&
               expect(cfg.validate("fluffyMcNotHere", string("foobar")), "Expect false when key does not exist");
    }

    /**
     * @name test_count
     * @brief verify ConfigStateMap::test_count() works
     *
     * @return bool (pass/fail)
     */
    bool test_count() {
        debug(name + "::test_count()");
        ConfigStateMap o;
        return expect(o.add("A", "a1", Bool), "Expect we can add 'A' as Bool") &&
               expect(o.count() == 1, "Expect count is 1") &&
               expect(o.add("B", "b2", Bool), "Expect add node ok") &&
               expect(o.count() == 2, "Expect count is 2") &&
               expect(o.add("C", "c3", Bool), "Expect add node ok") &&
               expect(o.count() == 3, "Expect count is 3") &&
               expect(o.add("D", "d4", Bool), "Expect add node ok") &&
               expect(o.count() == 4, "Expect count is 4") &&
               expect(o.add("E", "e5", Bool), "Expect add node ok") &&
               expect(o.count() == 5, "Expect count is 5") &&
               expect(o.add("F", "f6", Bool), "Expect add node ok") &&
               expect(o.count() == 6, "Expect count is 6") &&
               expect(o.add("G", "g7", Bool), "Expect add node ok") &&
               expect(o.count() == 7, "Expect count is 7") &&
               expect(o.add("H", "h8", Bool), "Expect add node ok") &&
               expect(o.count() == 8, "Expect count is 8");
    }

    /**
     * @name test_initialization_with_existing_map
     * @brief verify ConfigStateMap::ConfigStateMap(&map) works
     *
     * @return bool (pass/fail)
     */
    bool test_initialization_with_existing_map() {
        debug(name + "::test_initialization_with_existing_map()");
        ConfigStateMap o;
        return expect(
                o.add("A", "a1", Bool) &&
                o.add("B", "b2", Bool) &&
                o.add("C", "c3", Double), "Expect we can add three records to map o") &&
               expect(p = new ConfigStateMap(&o), "Expect new class is initialized without error.") &&
               expect(o.count() == 3, "Expect 3 nodes in map o") &&
               expect(p->count() == 3, "Expect 3 nodes in map p");
    }

    /**
     * @name test_getLhsBegin
     * @brief verify ConfigStateMap::getLhsBegin() works
     *
     * @return bool (pass/fail)
     */
    bool test_getLhsBegin() {
        debug(name + "::test_getLhsBegin()");
        ConfigStateMap o;
        o.add("A", "a1", Bool) &&
        o.add("B", "b2", Bool) &&
        o.add("C", "c3", Double);
        auto a = o.getLhsBegin();
        return expect(a->first == "A", "Expect 'A' as the beginning of the map. Found:" + a->first);
    }

    /**
     * @name test_getLhsEnd
     * @brief verify ConfigStateMap::getLhsEnd() works
     *
     * @return bool (pass/fail)
     */
    bool test_getLhsEnd() {
        debug(name + "::test_getLhsBegin()");
        ConfigStateMap o;
        o.add("A", "a1", Bool);
        o.add("B", "b2", Bool);
        o.add("C", "c3", Bool);
        o.add("D", "d4", Bool);
        o.add("E", "e5", Bool);
        o.add("F", "f6", Bool);
        o.add("G", "g7", Bool);
        string s = "";
        for (auto i = o.getLhsBegin(); i != o.getLhsEnd(); ++i) s += i->first;
        return expect(s == "ABCDEFG", "expect we can enumerate our set.");
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_initialize_class(), "Expect test_initialize_class OK (map created)") &&
               expect(test_add(), "Expect test_add OK") &&
               expect(test_erase(), "Expect test_erase OK") &&
               expect(test_exists(), "Expect test_exists OK") &&
               expect(test_rhs_exists(), "Expect test_rhs_exists() OK") &&
               expect(test_count(), "Expect count OK") &&
               expect(test_get_required(), "Expect test_get_required OK") &&
               expect(test_get_required_exception(), "Expect test_get_required_exception OK") &&
               expect(test_get_type(), "expect test_get_type OK") &&
               expect(test_get_type_exception(), "expect test_get_type_exception OK") &&
               expect(test_get_rhs(), "expect test_get_rhs OK") &&
               expect(test_get_rhs_exception(), "expect test_get_rhs_exception throws exception on invalid lhs") &&
               expect(test_initialization_with_existing_map(), "expect test_initialization_with_existing_map ok") &&
               expect(test_getLhsBegin(), "expect test_getLhsBegin() OK") &&
               expect(test_getLhsEnd(), "Expect test_getLhsEnd() ok") &&
               expect(test_run_validate(), "expect test_run_validate OK") &&
               expect(test_run_validate_missing(),
                      "expect test_run_validate_missing returns false when lhs not found.");
    }

};/*end of class*/