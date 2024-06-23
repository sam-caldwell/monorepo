/**
 * @name ArbitraryKvBtree/tests/TestRemoveAllNodes.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#include <iostream>
#include "../src/ArbitraryKvBtree.h"

class TestRemoveBasic : public TestBase {
public:
    TestRemoveBasic(string n) {
        name = n;
    }

    ~TestRemoveBasic() {/*noop*/}

    bool test_remove_single_node() {
        debug(name + "::test_remove_single_node()");
        ArbitraryKvBtree tree(true); // unique is turned on...
        return expect(tree.insert("nodeM", true), "We create a single node.") &&
               expect(tree.remove("nodeM", false), "We can delete the single node.") &&
               expect(!tree.exists("nodeM"), ">>Expect nodeM does not exist");
    }

    bool test_remove_single_node_subtree() {
        debug(name + "::test_remove_single_node_subtree()");
        ArbitraryKvBtree tree(true); // unique is turned on...
        return expect(tree.insert("nodeM", true), "We create a single node.") &&
               expect(tree.remove("nodeM", true), "We can delete the single node.") &&
               expect(!tree.exists("nodeM"), ">>Expect nodeM does not exist");
    }


    bool test_remove_subtree_root() {
        debug(name + "::test_remove_subtree_root()");
        ArbitraryKvBtree tree(true); // unique is turned on...
        return expect(
                (tree.insert("nodeM", true) &&
                 tree.insert("nodeG", true) &&
                 tree.insert("nodeD", true) &&
                 tree.insert("nodeB", true) &&
                 tree.insert("nodeA", true) &&
                 tree.insert("nodeC", true) &&
                 tree.insert("nodeF", true) &&
                 tree.insert("nodeE", true) &&
                 tree.insert("nodeJ", true) &&
                 tree.insert("nodeI", true) &&
                 tree.insert("nodeK", true) &&
                 tree.insert("nodeL", true) &&
                 tree.insert("nodeT", true) &&
                 tree.insert("nodeQ", true) &&
                 tree.insert("nodeW", true) &&
                 tree.insert("nodeY", true) &&
                 tree.insert("nodeX", true) &&
                 tree.insert("nodeZ", true) &&
                 tree.insert("nodeU", true) &&
                 tree.insert("nodeV", true) &&
                 tree.insert("nodeO", true) &&
                 tree.insert("nodeN", true) &&
                 tree.insert("nodeP", true) &&
                 tree.insert("nodeS", true) &&
                 tree.insert("nodeR", true)
                ), ">>Expect we can insert a 25 nodes") &&
               expect(tree.count() == 25, "Expect 25 nodes Found:" + to_string(tree.count())) &&
               expect(tree.remove("nodeM", true), ">>expect to delete root (node4) and down") &&
               expect(!tree.exists("nodeM"), ">>Expect nodeM does not exist") &&
               expect(!tree.exists("nodeR"), ">>Expect nodeR does not exist") &&
               expect(!tree.exists("nodeA"), ">>Expect nodeA does not exist") &&
               expect(tree.count() == 0, "Expect 0 nodes Found:" + to_string(tree.count()));
    }

    bool test_remove_subtree_nonroot() {
        debug(name + "::test_remove_subtree_nonroot()");
        ArbitraryKvBtree tree(true); // unique is turned on...
        return expect(
                (tree.insert("nodeM", true) &&
                 tree.insert("nodeG", true) &&
                 tree.insert("nodeD", true) &&
                 tree.insert("nodeB", true) &&
                 tree.insert("nodeA", true) &&
                 tree.insert("nodeC", true) &&
                 tree.insert("nodeF", true) &&
                 tree.insert("nodeE", true) &&
                 tree.insert("nodeJ", true) &&
                 tree.insert("nodeI", true) &&
                 tree.insert("nodeK", true) &&
                 tree.insert("nodeL", true) &&
                 tree.insert("nodeT", true) &&
                 tree.insert("nodeQ", true) &&
                 tree.insert("nodeW", true) &&
                 tree.insert("nodeY", true) &&
                 tree.insert("nodeX", true) &&
                 tree.insert("nodeZ", true) &&
                 tree.insert("nodeU", true) &&
                 tree.insert("nodeV", true) &&
                 tree.insert("nodeO", true) &&
                 tree.insert("nodeN", true) &&
                 tree.insert("nodeP", true) &&
                 tree.insert("nodeS", true) &&
                 tree.insert("nodeR", true)
                ), ">>Expect we can insert a 25 nodes") &&
               expect(tree.count() == 25, "Expect 25 nodes Found:" + to_string(tree.count())) &&
               expect(tree.remove("nodeB", true), ">>expect to delete nodeB and down") &&
                expect(!tree.exists("nodeB"), ">>Expect nodeB does not exist") &&
               expect(!tree.exists("nodeA"), ">>Expect nodeA does not exist") &&
               expect(!tree.exists("nodeC"), ">>Expect nodeC does not exist") &&
               expect(tree.count() == 22, "Expect 22 nodes Found:" + to_string(tree.count()));
    }

    bool test_remove_root_node() {
        debug(name + "::test_remove_root_node()");
        ArbitraryKvBtree tree(true); // unique is turned on...
        return expect(
                (tree.insert("nodeM", true) &&
                 tree.insert("nodeG", true) &&
                 tree.insert("nodeD", true) &&
                 tree.insert("nodeB", true) &&
                 tree.insert("nodeA", true) &&
                 tree.insert("nodeC", true) &&
                 tree.insert("nodeF", true) &&
                 tree.insert("nodeE", true) &&
                 tree.insert("nodeJ", true) &&
                 tree.insert("nodeI", true) &&
                 tree.insert("nodeK", true) &&
                 tree.insert("nodeL", true) &&
                 tree.insert("nodeT", true) &&
                 tree.insert("nodeQ", true) &&
                 tree.insert("nodeW", true) &&
                 tree.insert("nodeY", true) &&
                 tree.insert("nodeX", true) &&
                 tree.insert("nodeZ", true) &&
                 tree.insert("nodeU", true) &&
                 tree.insert("nodeV", true) &&
                 tree.insert("nodeO", true) &&
                 tree.insert("nodeN", true) &&
                 tree.insert("nodeP", true) &&
                 tree.insert("nodeS", true) &&
                 tree.insert("nodeR", true)
                ), ">>Expect we can insert a 25 nodes") &&
               expect(tree.count() == 25, "Expect 25 nodes Found:" + to_string(tree.count())) &&
               expect(tree.remove("nodeM", false), ">>expect to delete only nodeM") &&
               expect(tree.count() == 24, "Expect 24 nodes Found:" + to_string(tree.count())) &&
               expect(!tree.exists("nodeM"), ">>Expect nodeM does not exist") &&
               expect(tree.exists("nodeR"), ">>Expect nodeR does exist") &&
               expect(tree.exists("nodeL"), ">>Expect nodeL does exist");
    }

    bool test_remove_middle_node() {
        debug(name + "::test_remove_middle_node()");
        ArbitraryKvBtree tree(true); // unique is turned on...
        return expect(
                (tree.insert("nodeM", true) &&
                 tree.insert("nodeG", true) &&
                 tree.insert("nodeD", true) &&
                 tree.insert("nodeB", true) &&
                 tree.insert("nodeA", true) &&
                 tree.insert("nodeC", true) &&
                 tree.insert("nodeF", true) &&
                 tree.insert("nodeE", true) &&
                 tree.insert("nodeJ", true) &&
                 tree.insert("nodeI", true) &&
                 tree.insert("nodeK", true) &&
                 tree.insert("nodeL", true) &&
                 tree.insert("nodeT", true) &&
                 tree.insert("nodeQ", true) &&
                 tree.insert("nodeW", true) &&
                 tree.insert("nodeY", true) &&
                 tree.insert("nodeX", true) &&
                 tree.insert("nodeZ", true) &&
                 tree.insert("nodeU", true) &&
                 tree.insert("nodeV", true) &&
                 tree.insert("nodeO", true) &&
                 tree.insert("nodeN", true) &&
                 tree.insert("nodeP", true) &&
                 tree.insert("nodeS", true) &&
                 tree.insert("nodeR", true)
                ), ">>Expect we can insert a 25 nodes") &&
               expect(tree.count() == 25, "Expect 25 nodes Found:" + to_string(tree.count())) &&
               expect(tree.remove("nodeB", false), ">>expect to delete only nodeB") &&
               expect(!tree.exists("nodeB"), ">>Expect nodeB does not exist") &&
                expect(tree.count() == 24, "Expect 24 nodes Found:" + to_string(tree.count())) &&
               expect(tree.exists("nodeA"), ">>Expect nodeA does exist") &&
               expect(tree.exists("nodeL"), ">>Expect nodeL does exist") &&
               expect(tree.exists("nodeG"), ">>Expect nodeG does exist") &&
               expect(tree.exists("nodeF"), ">>Expect nodeF does exist") &&
               expect(tree.exists("nodeE"), ">>Expect nodeE does exist") &&
               expect(tree.exists("nodeC"), ">>Expect nodeC does exist") ||
               debug(tree.drawTree());
    }

    bool test_remove_trailing_node() {
        debug(name + "::test_remove_trailing_node()");
        ArbitraryKvBtree tree(true); // unique is turned on...
        return expect(
                (tree.insert("nodeM", true) &&
                 tree.insert("nodeG", true) &&
                 tree.insert("nodeD", true) &&
                 tree.insert("nodeB", true) &&
                 tree.insert("nodeA", true) &&
                 tree.insert("nodeC", true) &&
                 tree.insert("nodeF", true) &&
                 tree.insert("nodeE", true) &&
                 tree.insert("nodeJ", true) &&
                 tree.insert("nodeI", true) &&
                 tree.insert("nodeK", true) &&
                 tree.insert("nodeL", true) &&
                 tree.insert("nodeT", true) &&
                 tree.insert("nodeQ", true) &&
                 tree.insert("nodeW", true) &&
                 tree.insert("nodeY", true) &&
                 tree.insert("nodeX", true) &&
                 tree.insert("nodeZ", true) &&
                 tree.insert("nodeU", true) &&
                 tree.insert("nodeV", true) &&
                 tree.insert("nodeO", true) &&
                 tree.insert("nodeN", true) &&
                 tree.insert("nodeP", true) &&
                 tree.insert("nodeS", true) &&
                 tree.insert("nodeR", true)
                ), ">>Expect we can insert a 25 nodes") &&
               expect(tree.count() == 25, "Expect 25 nodes Found:" + to_string(tree.count())) &&
               expect(tree.remove("nodeC", false), ">>expect to delete only nodeC") &&
               expect(!tree.exists("nodeC"), ">>Expect nodeC does not exist") &&
               expect(tree.count() == 24, "Expect 24 nodes Found:" + to_string(tree.count())) &&
               expect(tree.exists("nodeA"), ">>Expect nodeA does exist") &&
               expect(tree.exists("nodeL"), ">>Expect nodeL does exist") &&
               expect(tree.exists("nodeG"), ">>Expect nodeG does exist") &&
               expect(tree.exists("nodeF"), ">>Expect nodeF does exist") &&
               expect(tree.exists("nodeE"), ">>Expect nodeE does exist") ||
               debug(tree.drawTree());
    }


    bool main() {
        debug(name + "::main() starting");
        return expect(test_remove_single_node(), "Expect create/delete a single node (removeSubtree:true)") &&
               expect(test_remove_single_node_subtree(), "Expect create/delete a single node (removeSubtree:false)") &&
               expect(test_remove_subtree_root(), "Expect(1) we can remove an entire subtree from root") &&
               expect(test_remove_subtree_nonroot(), "Expect(2)we can remove an entire subtree from child node.") &&
               expect(test_remove_root_node(), "Expect we can remove a single root node.") &&
               expect(test_remove_middle_node(), "Expect we can remove a single middle node.") &&
               expect(test_remove_trailing_node(), "Expect we can remove a single trailing node.");
    }
};/*end of class*/
