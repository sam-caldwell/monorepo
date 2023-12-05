/**
 * @name types/tests/TestAdd.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

#include <iostream>
#include "projects/common/types/ConfigStateMap.h"

class TestAdd : public TestBase {
private:
    ConfigStateMap *cfg;
public:
    /**
     * @name class constructor
     * @brief setup the test.
     * @param n
     */
    TestAdd(string n) {
        name = n;
        cfg = new ConfigStateMap;
    }

    /**
     * @name class destructor
     * @brief destroy the test
     */
    ~TestAdd() {
        if (cfg) delete cfg;
    }

    /**
     * @name test_add
     * @brief add a set of records with various inputs and defaults to the ConfigStateMap (cfg) property.
     *
     * @return bool (pass/fail)
     */
    bool test_add() {
        return expect(
                cfg->add("PropertyA", "internal1", Bool, false, &ConfigStateMapNullValidator) &&
                cfg->add("PropertyB", "internal2", Bool, true, false, &ConfigStateMapNullValidator) &&
                cfg->add("PropertyC", "internal3", Double) &&
                cfg->add("PropertyD", "internal4", Float) &&
                cfg->add("PropertyE", "internal5", Int) &&
                cfg->add("PropertyF", "internal6", String) &&
                cfg->add("PropertyG", "internal7", Uint), "Expect we can create multiple map nodes.");
    }

    /**
     * @name test_add_duplicates_not_allowed()
     * @brief Verify that a duplicate record added to ConfigStateMap(cfg) will throw a runtime_error exception.
     *
     * @return bool (pass/fail)
     */
    bool test_add_duplicates_not_allowed() {
        try {
            return expect(cfg->add("A", "a1", Bool), "we can add an initial node.") &&
                   expect(cfg->add("A", "a1", Bool), "Expect we cannot add duplicate lhs values");
        } catch (runtime_error &e) {
            return expect(true, "Expect we cannot add duplicate lhs values: " + string(e.what()));
        }
    }

    /**
     * @name test_count
     * @brief verify the record count in ConfigStateMap matches the given input count (c).
     *
     * @param c int
     * @return bool (pass/fail)
     */
    bool test_count(int c) {
        return expect(cfg->count() == c, "Expect count(" + to_string(c) + ") nodes");
    }

    /**
     * @name test_exists_true
     * @brief test the exists() method with a condition that expects true result.
     *
     * @return bool (pass/fail)
     */
    bool test_exists_true() {
        return expect(cfg->exists("PropertyA"), "Expect that 'PropertyA' exists");
    }

    /**
     * @param test_exists_false
     * @brief test the exists() method with a condition that expects false result
     *
     * @return bool (pass/fail)
     */
    bool test_exists_false() {
        return expect(!cfg->exists("BadName"), "Expect that 'BadName' does not exist and null is returned.");
    }

    /**
     * @name test_rhs_exists_true
     * @brief test whether a known rhs value exists as expected.
     *
     * @return bool (pass/fail)
     */
    bool test_rhs_exists_true() {
        return expect(cfg->rhsExists("internal1"), "Expect that rhs 'internal1' exists");
    }

    /**
     * @name test_rhs_exists_false
     * @brief test whether a bogus rhs value will cause rhsExists() to return false as expected.
     *
     * @return bool (pass/fail)
     */
    bool test_rhs_exists_false() {
        return expect(!cfg->rhsExists("BadName"), "Expect that rhs 'BadName' does not exist");
    }

    /**
     * @name test_get_required
     * @brief return true/false result for the getRequired() method based on prior inputs (known conditions)
     *
     * @return bool (pass/fail)
     */
    bool test_get_required() {
        return expect(!cfg->getRequired("PropertyA"), "Expect PropertyA is required (false)") &&
               expect(cfg->getRequired("PropertyB"), "Expect PropertyB is required (true)");
    }

    /**
     * @name test_get_required_missing
     * @brief expect an out_of_range exception when getRequired() method is called with a non-existent lhs value.
     *
     * @return bool (pass/fail)
     */
    bool test_get_required_missing() {
        try {
            return expect(cfg->getRequired("BadToken") == Bool,
                          "Expect getRequired() exception on non-existing token. Oops");
        } catch (out_of_range &e) {
            return expect(true, "Expect getRequired() exception on non-existing token. ok");
        }
    }

    /**
     * @name test_get_type
     * @brief expect that given a known ConfigStateMap record the getType() method will return the corresponding
     *        ValueTypes value.
     *
     * @return bool (pass/fail)
     */
    bool test_get_type() {
        return expect(cfg->getType("PropertyA") == Bool, "Expect .getType() will return Bool") &&
               expect(cfg->getType("PropertyB") == Bool, "Expect .getType() will return Bool") &&
               expect(cfg->getType("PropertyC") == Double, "Expect .getType() will return Double") &&
               expect(cfg->getType("PropertyD") == Float, "Expect .getType() will return Float") &&
               expect(cfg->getType("PropertyE") == Int, "Expect .getType() will return Int") &&
               expect(cfg->getType("PropertyF") == String, "Expect .getType() will return String") &&
               expect(cfg->getType("PropertyG") == Uint, "Expect .getType() will return Uint");
    }

    /**
     * @name test_get_type_missing
     * @brief expect that an out_of_range exception will be thrown if a non-existent key name is passed to getType()
     *
     * @return bool (pass/fail)
     */
    bool test_get_type_missing() {
        try {
            return expect(cfg->getType("BadToken") == Bool, "Expect getType() exception on non-existing token. Oops");
        } catch (out_of_range &e) {
            return expect(true, "Expect getType() exception on non-existing token. ok");
        }
    }

    /**
     * @name test_get_rhs
     * @brief Expect that given a known ConfigStatusMap, calling getRhsString() for a given key will return the
     *        associated rhs string value.
     *
     * @return bool (pass/fail)
     */
    bool test_get_rhs() {
        return expect(cfg->getRhsString("PropertyA") == "internal1", "Expect getRhsString will return 'internal1'");
    }

    /**
     * @name test_get_rhs_missing
     * @brief Expect that an out_of_range exception will be thrown if getRhsString() is called with a non-existent
     *        node key.
     *
     * @return bool (pass/fail)
     */
    bool test_get_rhs_missing() {
        try {
            return expect(cfg->getRhsString("BadToken") == "internal1",
                          "Expect getRhsString() exception on non-existing token. Oops");
        } catch (out_of_range &e) {
            return expect(true, "Expect getRhsString() exception on non-existing token. ok");
        }
    }

    /**
     * @name test_run_validate
     * @brief Given a known ConfigStateMap, call the validate() method with a valid input value and expect that
     *        the value will be validated against the configured validator function in the ConfigStateMap.  If the
     *        value is valid, expect validate() will return true, otherwise it should return false.
     *
     * @return bool (pass/fail)
     */
    bool test_run_validate() {
        return expect(cfg->validate("PropertyA", "Happy"), "Expect A Validates with noop") &&
               expect(cfg->validate("PropertyB", "Happy"), "Expect B validates with simple test") &&
               expect(cfg->validate("PropertyC", "Joy! Joy!"), "Expect C fails with simple test.");
    }

    /**
     * @name test_run_validate_missing
     * @brief Given a known ConfigStateMap, calling the validate() method with an invalid lhs key should cause
     *        an exception.
     * @return
     */
    bool test_run_validate_missing() {
        try {
            return expect(cfg->validate("fluffyMcNotHere", string("foobar")), "Expect exception on non-existent key");
        } catch (out_of_range) {
            return expect(true, "Expect exception on non-existent key");
        }
    }

    /**
     * @name main
     * @brief test coordination
     *
     * @return bool (pass/fail)
     */
    bool main() {
        debug(name + "::main()");
        return expect(test_count(0), "Expect test_count 0 OK") &&
               expect(test_add(), "Expect test_add() ok") &&
               expect(test_count(7), "Expect test_count 7 OK") &&
               expect(test_add_duplicates_not_allowed(), "Expect test_add_duplicates_not_allowed OK") &&
               expect(test_count(8), "Expect test_count 8 OK") &&
               expect(test_exists_true(), "Expect test_exists_true OK") &&
               expect(test_exists_false(), "Expect test_exists_false OK") &&
               expect(test_rhs_exists_true(), "Expect test_rhs_exists_true OK") &&
               expect(test_exists_false(), "Expect test_exists_false OK") &&
               expect(test_get_required(), "Expect test_get_required OK") &&
               expect(test_get_required_missing(), "Expect test_get_required_missing OK") &&
               expect(test_get_type(), "Expect test_get_type OK") &&
               expect(test_get_type_missing(), "Expect test_get_type_missing() OK") &&
               expect(test_get_rhs(), "Expect test_get_rhs OK") &&
               expect(test_get_rhs_missing(), "Expect test_get_rhs_missing OK") &&
               expect(test_run_validate(), "Expect test_run_validate OK") &&
               expect(test_run_validate_missing(), "Expect test_run_validate_missing OK");
    }

};/*end of class*/