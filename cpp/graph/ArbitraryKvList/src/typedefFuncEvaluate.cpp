/**
 * @name ArbitraryKvList/src/typedefFuncEvaluate.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name FuncEvaluate (typedef)
 * @brief This is the function pointer pattern for
 *        LessThan() and GreaterThan(), etc. to
 *        compare two pointers.
 */
typedef bool (*FuncEvaluate)(ArbitraryKvNode *, ArbitraryKvNode *);