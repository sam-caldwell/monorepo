/**
 * @name Configuration/src/Configuration.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @brief This file is the start of the configuration library.
 *
 * @details The Configuration class will allow an application to
 *          be configured using command-line (CLI) arguments, a
 *          configuration file (of some supported format) or
 *          environment variables.  Configuration parameters
 *          are stored as key-value pairs using the Arbitrary
 *          Value Binary Tree project (ArbitraryKvBtree).
 */
#ifndef Configuration_H
#define Configuration_H



#include "../../../../types/ValueTypes.h"
#include "../../../../types/unsigned_int.h"
#include "../../../../types/src/ConfigStateMap/ConfigStateMap.h"
#include "../../../../exceptions/exceptions.h"
#include "constants.h"
#include "../../../../graph/ArbitraryKvBtree/src/ArbitraryKvBtree.h"

using namespace std;

class Configuration {
public:
    /**
     * @name Default Class constructor.
     * @brief Initialize state for blank state with static arbitrary value binary tree.
     * @details Configuration is a base class which can be inherited by both
     *          the EnvironmentVariableParser, CommandLineParser and
     *          ConfigurationFileParser as well as the consuming application.
     *          This class has a static tree property, shared across the various
     *          consumers.
     */
    Configuration();

    /**
     * @name Class destructor
     * @brief tear down the system.
     */
    ~Configuration();

    /**
     * @name count
     * @brief return the number of configuration data objects
     *
     * @return int
     */
    inline int count();

    /**
     * @name getBool
     * @brief return the string value of a given node identified by key.
     *
     * @param key string
     * @return bool
     */
    inline bool getBool(string key);

    /**
     * @name getDouble
     * @brief return the double-precision floating-point
     *        value of a given node identified by a key.
     *
     * @param key string
     * @return double
     */
    inline double getDouble(string key);

    /**
     * @name getFloat
     * @brief return the single-precision floating-point
     *        value of a given node identified by key.
     *
     * @param key string
     * @return float
     */
    inline float getFloat(string key);

    /**
     * @name getInt
     * @brief return the integer value of a given node identified by key.
     *
     * @param key string
     * @return int
     */
    inline int getInt(string key);

    /**
     * @name getString
     * @brief return the string value of a given node identified by key.
     *
     * @param key string
     * @return string
     */
    inline string getString(string key);

    /**
     * @name getType
     * @brief return the ValueTypes of the given node.
     *
     * @param key string
     * @return ValueTypes
     */
    inline ValueTypes getType(string key);

    /**
     * @name getUint
     * @brief return the unsigned-integer (uint) value of a given node identified by key.
     *
     * @param key string
     * @return uint
     */
    inline uint getUint(string key);

    /**
     * @name insert[bool]
     * @brief insert new node with a boolean value
     *
     * @param key string
     * @param n bool
     * @return bool
     */
    inline bool insert(string key, bool n);

    /**
     * @name insert[double]
     * @brief insert new node with a double-precision floating-point value.
     *
     * @param key string
     * @param n double
     * @return bool
     */
    inline bool insert(string key, double n);

    /**
     * @name insert[float]
     * @brief insert new node with a single-precision floating-point number
     *
     * @param key string
     * @param n float
     * @return bool
     */
    inline bool insert(string key, float n);

    /**
     * @name insert[int]
     * @brief insert new node with an integer
     *
     * @param key string
     * @param n int
     * @return bool
     */
    inline bool insert(string key, int n);

    /**
     * @name insert[string]
     * @brief insert new node with a string
     *
     * @param key string
     * @param n string
     * @return bool
     */
    inline bool insert(string key, string n);

    /**
     * @name insert[uint]
     * @brief insert new node with an unsigned integer (uint)
     *
     * @param key string
     * @param n uint
     * @return bool
     */
    inline bool insert(string key, uint n);

    /**
     * @name loadData
     * @brief Given the ConfigStateMap (by reference), the source name (lhsName) in this
     *        ConfigStateMap, used by the caller to get the data object (value), the target
     *        name (rhsName) to which we will store the loaded (type-aware) value, lookup and
     *        validate the string value then translate it into the appropriate ArbitraryKvBtree
     *        object in the Configuration binary tree structure.
     *
     * @throws runtime_error exception on any type, validation or storage issues.
     *
     * @param *map ConfigStateMap pointer
     * @param lhsName string
     * @param value string
     *
     * @return bool (operation outcome)
     */
    inline bool loadData(ConfigStateMap *map, string lhsName, string value);

    /**
     * @name parseBool
     * @brief Parse a given input (v) and return the boolean equivalent.
     *        Where an empty value "" is encountered, return the default value.
     * @param v string
     * @param defaultValue bool
     * @return bool
     */
    bool parseBool(string v, bool defaultValue);

    /**
     * @name parseDouble
     * @brief Parse a given input (v) and return the double-precision floating point
     *        representation.  If an empty value "" is encountered, return the default
     *        value.
     *
     * @param v string
     * @param defaultValue double
     * @return double.
     */
    double parseDouble(string v, double defaultValue);

    /**
     * @name parseFloat
     * @brief Parse a given input (v) and return the single-precision floating point
     *        representation.  If an empty value "" is encountered, return the default
     *        value.
     *
     * @param v string
     * @param defaultValue float
     * @return double.
     */
    float parseFloat(string v, float defaultValue);

    /**
     * @name parseInt
     * @brief Parse a given input (v) and return the integer representation.
     *        If an empty value "" is encountered, return the default value.
     *
     * @param v string
     * @param defaultValue int
     * @return int.
     */
    int parseInt(string v, int defaultValue);

    /**
     * @name parseUint
     * @brief Parse a given input (v) and return the unsigned integer representation.
     *        If an empty value "" is encountered, return the default value.
     *
     * @warning If no conversion could be performed, an invalid_argument exception is thrown.
     *
     * @param v string
     * @param defaultValue uint
     * @return uint.
     */
    int parseUint(string v, uint defaultValue);

    /**
     * @name remove
     * @brief remove node using key
     *
     * @param s string
     * @return bool
    */
    inline bool remove(string key);

    /**
     * @name showKeys
     * @brief print all the keys in the Configuration object
     *
     * @return string
     */
    inline string showKeys(){
        return (tree)?tree->drawTree():"";
    }

private:
    /**
     * The *tree pointer references a binary tree of ArbitraryValue nodes
     * used to store a set of configuration objects.
     */
    ArbitraryKvBtree *tree;

}; /*end of Configuration*/

#include "Constructors.cpp"
#include "count.cpp"
#include "Destructors.cpp"
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
#include "loadData.cpp"
#include "parseBool.cpp"
#include "parseDouble.cpp"
#include "parseFloat.cpp"
#include "parseInt.cpp"
#include "parseUint.cpp"
#include "remove.cpp"


#endif /* Configuration_H */