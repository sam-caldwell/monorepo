/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 */
#include <iostream>
#include "../src/ArbitraryKvList.h"

#define PI 3.141529

class TestLinkListGetValue : public TestBase {
private:
    ArbitraryKvList *root;
public:
    TestLinkListGetValue(string n) { name = n; }

    ~TestLinkListGetValue() { delete root; }

    bool test_setup() {
        debug(name + "::test_setup() starting");
        root = new ArbitraryKvList(true);
        return expect(root->insert("bool_true", true), "Expect we can insert boolean true") &&
               expect(root->insert("bool_false", false), "Expect we can insert boolean false") &&
               expect(root->insert("double_pi", (double) (PI)), "Expect we can insert double") &&
               expect(root->insert("float_pi", (float) (PI)), "Expect we can insert float") &&
               expect(root->insert("int_key", (int) (-3)), "Expect we can insert int") &&
               expect(root->insert("string_key", (string)"string_value"), "Expect we can insert string node.") &&
               expect(root->insert("uint_key", (uint) (1)), "Expect we can insert uint node.");
    }

    bool test_exists() {
        debug(name + "::test_exists() starting");
        debug("list of keys:" + root->dumpKeys());
        return expect(root->exists("bool_true"), "expect bool_true exists") &&
               expect(root->exists("bool_false"), "expect bool_false exists") &&
               expect(root->exists("double_pi"), "expect double_pi exists") &&
               expect(root->exists("float_pi"), "expect float_pi exists") &&
               expect(root->exists("int_key"), "Expect int_key exists") &&
               expect(root->exists("string_key"), "Expect string_key exists") &&
               expect(root->exists("uint_key"), "Expect uint_key exists");
    }

    bool test_getters() {
        debug(name + "::test_getters() starting");
        debug("list of keys:" + root->dumpKeys());
        return expect(root->getBool("bool_true"), ">>Expect getBool true will return true.") &&
               expect(!root->getBool("bool_false"), ">>Expect getBool false will return false.") &&
               expect(root->getDouble("double_pi") == (double) (PI), ">>Expect getDouble true will return PI.") &&
               expect(root->getFloat("float_pi") == (float) (PI), ">>Expect getFloat true will return PI.") &&
               expect(root->getInt("int_key") == -3, ">>Expect getInt true will return -3.") &&
               expect(root->getString("string_key") == "string_value",
                      ">>Expect getString will return string value. Found:"+root->getString("string_key")) &&
               expect(root->getUint("uint_key") == 1, ">>Expect getUint will return 1.");

    }

    bool main() {
        debug(name + "::main() starting...");
        return expect(test_setup(), "Expect we can create initialize a default state.") &&
               expect(test_exists(), "Expect that inserted nodes exist.") &&
               expect(test_getters(), "Expect the getter get<Type>() methods return expected values..");
    }
};