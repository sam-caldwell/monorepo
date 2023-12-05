/**
 * @name ArbitraryKvList/src/LessThan.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name LessThan
 * @brief Evaluate if the left-hand (lhs) key is less than right hand (rhs)
 *
 * @param *LHS node pointer
 * @param *RHS node pointer
 * @return bool
 */
bool LessThan(ArbitraryKvNode *LHS, ArbitraryKvNode *RHS) {
    string LHS_KEY = (LHS) ? LHS->key() : "";
    string RHS_KEY = (RHS) ? RHS->key() : "";
    return LHS_KEY < RHS_KEY;
}
