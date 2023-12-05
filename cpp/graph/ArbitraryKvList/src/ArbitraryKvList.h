/**
 * @name ArbitraryKvList/src/ArbitraryKvList.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * This file defines a double-linked link list
 * container.
 *
 * WE SHOULD NEVER DIRECTLY EXPOSE OUR DATA NODES!
 */
#ifndef ARBITRARY_KV_LINK_LIST_H
#define ARBITRARY_KV_LINK_LIST_H

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

#include <regex>
#include <list>
#include <stdexcept>

#include "projects/common/SimpleLock/src/SimpleLock.h"
#include "projects/common/types/unsigned_int.h"
#include "projects/common/types/SortOrder.h"
#include "projects/graph/ArbitraryKvNode/src/ArbitraryKvNode.h"

using namespace std;

class ArbitraryKvList {
public:
    /**
     * @name Default Class constructor.
     * @brief Initialize class (SimpleLock mechanism, root pointer, count, and set unique false).
     */
    ArbitraryKvList();

    /**
     * @name Class Constructor
     * @brief Initialize class (SimpleLock mechanism, root pointer,
     *        count, set unique and moveLeftOnInsert flag).
     *
     * @param unique bool.
     * @param moveLeftOnInsert bool flag to move root pointer left on insert.
     * @param strongTyping bool flag to enforce strong typing.
     */
    ArbitraryKvList(bool unique, bool moveLeftOnInsert, bool strongTyping);

    /**
     * Class destructor.
     * - decrement counter
     * - delete root, causing cascading node delete.
     */
    ~ArbitraryKvList();

    /**
     * @name count
     * @brief Return the total count of nodes.
     *
     * @return int (count)
     */
    int count();

    /**
     * @name count [matches]
     * @brief Count the number of nodes with the given key.
     *
     * @param key string
     * @return int (count of matches)
     */
    int count(string key);

    /**
     * @name dumpKeys
     * @brief  Dump List keys as a comma-delimited list.
     *
     * @return string
     */
    string dumpKeys();

    /**
     * @name exists
     * @brief Determine if a node with the given key exists.
     *        To return the contents of the node, use 'search()'
     *
     * @param criteria string
     * @return operation outcome (success/fail)
     */
    bool exists(string criteria);

    /**
     * @name getBool
     * @brief Search for and return the given node as boolean.
     *
     * @warning will throw out_of_range exception if no node exists.
     *
     * @param key string
     * @return bool
     */
    bool getBool(string key);

    /**
     * @name getBool
     * @brief return boolean value of current node.
     *
     * @warning will throw out_of_range exception if no node exists.
     *
     * @return bool
     */
    bool getBool();

    /**
     * @name getDouble
     * @brief Return the current node (double).
     *        Root node will be left on this node.
     *
     * @warning will throw NotFound() exception if key not found
     *
     * @param string key
     * @return double
     */
    double getDouble(string key);

    /**
     * @name getFloat
     * @brief Return the current node (float).
     *        Root node will be left on this node.
     *
     * @warning will throw NotFound() exception if key not found
     *
     * @param string key
     * @return float
     */
    float getFloat(string key);

    /**
     * @name getInt
     * @brief Return the current node (int).
     *        Root node will be left on this node.
     *
     * @warning will throw NotFound() exception if key not found
     *
     * @param string key
     * @return int
     */
    int getInt(string key);

    /**
     * @name getString
     * @brief Return the current node (string).
     *        Root node will be left on this node.
     *
     * @warning will throw NotFound() exception if key not found
     *
     * @param string key
     * @return string
     */
    string getString(string key);

    /**
     * @name getType
     * @brief returns the ValueTypes type of the named node.
     *
     * @warning will throw NotFound() exception if key not found
     *
     * @param string key
     * @return ValueTypes
     */
    inline ValueTypes getType(string key);

    /**
     * @name getType
     * @brief returns the ValueTypes type of the current node.
     *
     * @warning will throw out_of_range exception if no node exists.
     *
     * @return ValueTypes
     */
    inline ValueTypes getType();

    /**
     * @name getUint
     * @brief Return the current node (uint).
     *        Root node will be left on this node.
     *
     * @warning will throw NotFound() exception if key not found
     *
     * @param string key
     * @return uint
     */
    uint getUint(string key);

    /**
     * @name insert [bool]
     * @brief Insert new node to immediate left of root.
     *        Check uniqueness flag and guarantee unique
     *        key if enabled.
     *
     * @param key string
     * @param v bool
     * @return operational outcome (success/fail)
     */
    bool insert(string key, bool v);

    /**
     * @name insert [double]
     * @brief Insert new node to immediate left of root.
     *        Check uniqueness flag and guarantee unique
     *        key if enabled.
     *
     * @param key string
     * @param v double
     * @return operational outcome (success/fail)
     */
    bool insert(string key, double v);

    /**
     * @name insert [float]
     * @brief Insert new node to immediate left of root.
     *        Check uniqueness flag and guarantee unique
     *        key if enabled.
     *
     * @param key string
     * @param v float
     * @return operational outcome (success/fail)
     */
    bool insert(string key, float v);

    /**
     * @name insert [int]
     * @brief Insert new node to immediate left of root.
     *        Check uniqueness flag and guarantee unique
     *        key if enabled.
     *
     * @param key string
     * @param v int
     * @return operational outcome (success/fail)
     */
    bool insert(string key, int v);

    /**
     * @name insert [string]
     * @brief Insert new node to immediate left of root.
     *        Check uniqueness flag and guarantee unique
     *        key if enabled.
     *
     * @param key string
     * @param v string
     * @return operational outcome (success/fail)
     */
    bool insert(string key, string v);

    /**
     * @name insert [uint]
     * @brief Insert new node to immediate left of root.
     *        Check uniqueness flag and guarantee unique
     *        key if enabled.
     *
     * @param key uint
     * @param v bool
     * @return operational outcome (success/fail)
     */
    bool insert(string key, uint v);

    /**
     * @name key
     * @brief Get the key value of the current (root) node.
     *
     * @return string
     */
    string key();

    /**
     * @name moveLeft
     * @brief Move one node left
     *
     * @return bool
     */
    inline bool moveLeft();

    /**
     * @name moveRight
     * @brief Move one node right
     *
     * @return bool
     */
    inline bool moveRight();

    /**
     * @name remove
     * @brief Delete the current node.
     *
     * @return bool
     */
    bool remove();

    /**
     * @name remove [keys]
     * @brief Delete a node with the given key.
     *        This will delete all matching keys.
     *
     * @param key string
     * @return bool
     */
    bool remove(string key);

    /**
     * @name resetLeft
     * @brief Move root pointer to the left-most node.
     */
    void resetLeft();

    /**
     * @name resetRight
     * @brief Move root pointer to the right-most node.
     */
    void resetRight();

    /**
     * @name search
     * @brief Search for a given key and leave root pointing
     *        to the given node.  If we do not find the criteria,
     *        then root will stay where it was.
     *
     * @param criteria string
     * @return bool (true: found, false: not found)
     */
    bool search(string criteria);

    /**
     * @name sort
     * @brief Sort the list ascending or descending.
     *
     * @param order SortOrder
     * @return bool
     */
    bool sort(SortOrder order);

    /**
     * @name swap
     * @brief Swap two nodes (by key)
     *
     * @param key1 string
     * @param key2 string
     * @return bool
     */
    bool swap(string key1, string key2);

    /**
     * @name unique [getter]
     * @brief get the current unique-ness
     *
     * @return bool (unique flag)
     */
    inline bool unique();

    /**
     * @name unique [setter]
     * @brief set the current unique-ness
     *
     * @param unique bool guarantee unique keys
     * @return bool (unique flag)
     */
    inline bool unique(bool unique);

private:
    /**
     * root pointer to the linked list.
     */
    ArbitraryKvNode *root;

    /**
     * node count
     */
    int _count;

    /**
     * flag to enforce uniqueness
     */
    bool _unique;

    /**
     * flag to enforce strong typing
     */
    bool _strongTyping;

    /**
     * flag to move left after each insert
     */
    bool _moveLeftOnInsert;

    /**
     * read/write SimpleLock
     */
    SimpleLock *lock;

    /**
     * @name getNodePointer
     * @brief Get Node pointer for given key
     *
     * @param key string
     * @return ArbitraryKvNode* pointer (null if key does not exist).
     */
    ArbitraryKvNode *getNodePointer(string key);
};

#include "LessThan.cpp"
#include "GreaterThan.cpp"
#include "swapPointers.cpp"
#include "typedefFuncEvaluate.cpp"
#include "Constructors.cpp"
#include "Destructor.cpp"
#include "UniqueGetter.cpp"
#include "UniqueSetter.cpp"
#include "Key.cpp"
#include "InsertBool.cpp"
#include "InsertDouble.cpp"
#include "InsertFloat.cpp"
#include "InsertInt.cpp"
#include "InsertString.cpp"
#include "InsertUint.cpp"
#include "ResetLeft.cpp"
#include "ResetRight.cpp"
#include "Exists.cpp"
#include "Count.cpp"
#include "CountMatches.cpp"
#include "Search.cpp"
#include "GetBool.cpp"
#include "GetDouble.cpp"
#include "GetFloat.cpp"
#include "GetInt.cpp"
#include "GetString.cpp"
#include "GetUint.cpp"
#include "DumpKeys.cpp"
#include "RemoveKeys.cpp"
#include "Remove.cpp"
#include "GetNodePointer.cpp"
#include "Swap.cpp"
#include "MoveLeft.cpp"
#include "MoveRight.cpp"
#include "Sort.cpp"

#endif //ARBITRARY_KV_LINK_LIST_H