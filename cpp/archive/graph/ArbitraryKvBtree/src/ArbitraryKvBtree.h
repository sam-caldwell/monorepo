/**
 * @name ArbitraryKvBtree/src/ArbitraryKvBtree.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @brief This file creates a single-link binary tree container.
 *
 * @note We MAY allow duplicates if unique==false.
 * @note duplicate nodes will evaluate to the right side of a subtree.
 *
 * @warning WE SHOULD NEVER DIRECTLY EXPOSE OUR DATA NODES!
 */
#ifndef ARBITRARY_KV_BTREE_H
#define ARBITRARY_KV_BTREE_H

/**
 * DEFAULT_BALANCER_THRESHOLD
 *      Determines the threshold beyond which a tree
 *      will be rebalanced.  When a tree's balance
 *      score exceeds this threshold, a re-balance
 *      operation will start.  This is an expensive
 *      operation as it will block the entire tree
 *      until completed.
 */
#define DEFAULT_BALANCER_THRESHOLD 10
#define DEFAULT_CALC_STATS_BACKOFF 100ms

/**
 * Use the following preprocessor directive to disable
 * debugger functionality.
 */
#ifdef TESTER_DEBUGGING
#undef TESTER_DEBUGGING
#endif /**TESTER_DEBUGGING*/
/**
 * Use the following preprocessor directive to enable
 * debugger functionality.
 */
//

#include <chrono>
#include <thread>
#include <stdexcept>

#include "../../../../SimpleLock/src/SimpleLock.h"
#include "../../../../types/unsigned_int.h"
#include "../../../../exceptions/exceptions.h"
#include "../../ArbitraryKvBtNode/src/ArbitraryKvBtNode.h"

using namespace std;

class ArbitraryKvBtree {
public:
    /**
     * @name Default Class constructor.
     * @brief initialize state (enforce unique)
     * @note Ensure root is null
     * @note Ensure count is zero (0)
     * @note Unique is false.
     */
    ArbitraryKvBtree();

    /**
     * @name Class constructor
     * @brief Initialize state (unique configurable)
     * @note Ensure root is null
     * @note Ensure count is zero (0)
     * @note Set unique flag
     * @note Set balanceThreshold value
     *
     * @param unique bool.
     * @param strongTyping bool (default false)
     */
    ArbitraryKvBtree(bool unique, bool strongTyping);

    /**
     * @name Class destructor.
     * @note decrement counter
     * @note delete root, causing cascading node delete.
     */
    ~ArbitraryKvBtree();

    /**
     * @name balance()
     * @brief Return the tree's balance score.
     *
     * @return int
     */
    inline int balance();

    /**
     * @name count()
     * @brief Count the nodes.
     *
     * @return int
     */
    inline int count();

    /**
     * @name count(criteria)
     * @brief Count the number of nodes with the given key.
     *
     * @param key string
     * @return int (count of matches)
     */
    inline int count(string criteria);

    /**
     * @name drawTree()
     * @brief Draw Binary Tree.
     *
     * @return string
     */
    inline string drawTree();

    /**
     * @name duplicates()
     * @brief Return duplicate count.
     *
     * @return int
     */
    inline int duplicates();

    /**
     * @name exists()
     * @brief Return whether a given node exists.  This will NOT
     * affect the node where root is pointing in the end.
     *
     * @param criteria string
     * @return bool
     */
    inline bool exists(string criteria);

    /**
     * @name getBool
     * @brief Return a boolean value for a matching criteria.
     *
     * @param criteria string
     * @return bool
     */
    inline bool getBool(string criteria);

    /**
     * @name getDobule
     * @brief Return a double value for a matching criteria.
     *
     * @param criteria string
     * @return double
     */
    inline double getDouble(string criteria);

    /**
     * @name getFloat
     * @brief Return a float value for a matching criteria.
     *
     * @param criteria string
     * @return float
     */
    inline float getFloat(string criteria);

    /**
     * @name getInt
     * @brief Return a int value for a matching criteria.
     *
     * @param criteria string
     * @return int
     */
    inline int getInt(string criteria);

    /**
     * @name getString
     * @brief Return a string value for a matching criteria.
     *
     * @param criteria string
     * @return string
     */
    inline string getString(string criteria);

    /**
     * @name getType
     * @brief Return the ValueTypes value of a given node.
     *
     * @param criteria string
     * @return ValueTypes
     */
    inline ValueTypes getType(string criteria);

    /**
     * @name getUint
     * @brief Return a uint value for a matching criteria.
     *
     * @param criteria string
     * @return uint
     */
    inline uint getUint(string criteria);

    /**
      * @name insert[bool]
      * @brief Insert boolean node.
      * @param key string
      * @param value bool
      * @return bool
      */
    inline bool insert(string key, bool v);

    /**
     * Insert boolean node.
     * @param key string
     * @param value double
     * @return bool
     */
    inline bool insert(string key, double v);

    /**
     * Insert boolean node.
     * @param key string
     * @param value float
     * @return bool
     */
    inline bool insert(string key, float v);

    /**
     * Insert boolean node.
     * @param key string
     * @param value int
     * @return bool
     */
    inline bool insert(string key, int v);

    /**
     * Insert boolean node.
     * @param key string
     * @param value string
     * @return bool
     */
    inline bool insert(string key, string v);

    /**
     * Insert boolean node.
     * @param key string
     * @param value uint
     * @return bool
     */
    inline bool insert(string key, uint v);

    /**
     * @name isLocked
     * @brief return tree SimpleLock state.
     * @return bool.
     */
    inline bool isLocked();

    /**
     * Remove a node with a given criteria
     *
     * @param criteria string
     * @param removeSubtree bool (default false)
     * @return bool
     */
    bool remove(string criteria, bool removeSubtree);

    /**
     * @name strongTyping [getter]
     * @brief strongTyping property getter
     *
     * @return count int.
     */
    inline bool strongTyping();

    /**
     * @name strongTyping [setter]
     * @brief strongTyping property setter
     *
     * @param v bool
     * @return bool
     */
    inline bool strongTyping(bool v);

    /**
     * Unique property getter
     * @return count int.
     */
    inline int unique();

    /**
     * Unique property setter
     * @param v bool
     * @return bool
     */
    inline bool unique(bool v);

protected:

    /**
     * Recursively calculate balance score
     * for the tree.  left: -1, right: +1
     *
     * @param *n ArbitraryKvBtNode pointer
     * @param d int
     * @return int
     */
    int balanceRecursiveCount(ArbitraryKvBtNode *n, int d);

    /**
     * Calculate stats cache (if invalid)
     */
    void calcStats();

    /**
     * Recursively Count Nodes.
     * @param *n ArbitraryKvBtNode pointer
     * @return int
     */
    int countRecursiveCount(ArbitraryKvBtNode *n);

    /**
     * @name deleteNode
     * @brief delete a single node and manage pointer references thereto.
     *
     * @param *n ArbitraryKvBtNode pointer
     * @return bool (operation outcome)
     */
    bool deleteNode(ArbitraryKvBtNode *n);

    /**
     * @name deleteSubtree
     * @brief Delete an entire subtree from some arbitrary root.
     *
     * @param *n ArbitraryKvBtNode pointer
     * @return bool (operation outcome)
     */
    bool deleteSubtree(ArbitraryKvBtNode *n);

    /**
     * Recursively calculate duplicate count.
     *
     * @param *n ArbitraryKvBtNode pointer
     * @return int
     */
    int duplicateRecursiveCount(ArbitraryKvBtNode *n);

    /**
     * Return whether a given node exists.  This will NOT
     * affect the node where root is pointing in the end.
     *
     * @param criteria string
     * @param lockOperation bool
     * @return bool
     */
    ArbitraryKvBtNode *findNode(string criteria, bool lockOperation);

    /**
     * Find a given node in the tree and return a pointer to
     * its node object (if found) or NULL (if not found), yielding
     * to any locks that may be in place.
     *
     * @param criteria string
     * @return ArbitraryKvBtNode pointer
     */
    ArbitraryKvBtNode *GetNode(string criteria);

    /**
     * @name insertNode
     * @brief Insert a given node (n) into the binary tree (root).
     *
     * @param *n ArbitraryKvBtNode pointer
     * @return bool
     */
    bool insertNode(ArbitraryKvBtNode *n);

    /**
     * Re-balance a subtree into the root tree.
     *
     * @param balance int
     * @param *root ArbitraryKvBtNode pointer
     * @param *lhs ArbitraryKvBtNode pointer
     * @param *rhs ArbitraryKvBtNode pointer
     * @param *n ArbitraryKvBtNode
     * @return bool
     */
    inline bool reBalance(int balance,
                          ArbitraryKvBtNode *root,
                          ArbitraryKvBtNode *lhs,
                          ArbitraryKvBtNode *rhs) {
        //ToDo: Implement re-balancer
        return true;
    };

    /**
     * Remove a given ArbitraryKvBtNode object.
     * Given removeSubtree true, we delete the subtrees.
     * Otherwise we will re-balance.
     *
     * @param *n ArbitraryKvBtNode pointer
     * @param removeSubtree bool
     * @return bool
     */
    bool removeNode(ArbitraryKvBtNode *n, bool removeSubtree);

    /**
     * Wrapper for recursive node counter used by calcStats
     * to multi-thread the operation.
     * @param *n ArbitraryKvBtNode pointer
     */
    void threadCountNodes(ArbitraryKvBtNode *n);

    /**
     * Wrapper for recursive tree-balance counter used by calcStats
     * to multi-thread the operation.
     * @param *n ArbitraryKvBtNode pointer
     */
    void threadCountBalance(ArbitraryKvBtNode *n);

    /**
     * Wrapper for recursive duplicate-node counter used by calcStats
     * to multi-thread the operation.
     * @param *n ArbitraryKvBtNode pointer
     */
    void threadCountDuplicates(ArbitraryKvBtNode *n);

private:
    /**
     * Flags
     */
    bool _unique;            // Flag which prohibits duplicate nodes
    bool _strongTyping;      // Flag to indicate strong typing should be enforced.
    bool _statsOk;           // Flag to indicate states are being recalculated (when false).
    /**
     * Counters (stats)
     */
    int _nodeCount;       //Count of all nodes in a tree.
    int _balanceCount;    //Count of the tree balance (-1 left, +1 right)
    int _duplicateCount;  //Count of all duplicate nodes.
    /**
     * Locks
     */
    SimpleLock *lock;
    SimpleLock *_lockNodeCount;      // used to SimpleLock count during recount
    SimpleLock *_lockBalanceCount;   // used to SimpleLock count during recount
    SimpleLock *_lockDuplicateCount; // used to SimpleLock count during recount
    /**
     * Tree pointer
     */
    ArbitraryKvBtNode *root;    // This should be the immutable pointer to the root of the binary tree.
}/**end of class*/;

#include "Constructors.cpp"
#include "Destructors.cpp"
#include "balance.cpp"
#include "calcStats.cpp"
#include "count.cpp"
#include "countMatches.cpp"
#include "duplicates.cpp"
#include "strongTyping.cpp"
#include "Unique.cpp"
#include "isLocked.cpp"
#include "Exists.cpp"
#include "findNode.cpp"
#include "getBool.cpp"
#include "getDouble.cpp"
#include "getFloat.cpp"
#include "getInt.cpp"
#include "getString.cpp"
#include "getType.cpp"
#include "getUint.cpp"
#include "insertBool.cpp"
#include "insertDouble.cpp"
#include "insertFloat.cpp"
#include "insertInt.cpp"
#include "insertString.cpp"
#include "insertUint.cpp"
#include "drawTree.cpp"
#include "Remove.cpp"


#endif //ARBITRARY_KV_BTREE_H