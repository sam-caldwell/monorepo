/**
 * @name ArbitraryKvList/src/GreaterThan.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name GreaterThan
 * @brief Evaluate if the left-hand (lhs) key is greater than right hand (rhs)
 *
 * @param LHS node pointer
 * @param RHS node pointer
 * @return bool
 */
bool GreaterThan(ArbitraryKvNode *LHS, ArbitraryKvNode *RHS) {
    string LHS_KEY = (LHS) ? LHS->key() : "";
    string RHS_KEY = (RHS) ? RHS->key() : "";
    return LHS_KEY > RHS_KEY;
}
