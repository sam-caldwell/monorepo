/**
 * @name types/src/SortOrderToString.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */

/**
 * @name SortOrderToString
 * @brief Convert a sort order (enum) to string.
 * @param o SortOrder
 * @return string
 */
string SortOrderToString(SortOrder o) {
    switch (o) {
        case Ascending:
            return "ascending";
            break;
        case Descending:
            return "descending";
            break;
        default:
            return "undefined";
            break;
    }
}