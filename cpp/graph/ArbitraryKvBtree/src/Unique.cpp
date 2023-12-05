/**
 * @name ArbitraryKvBtree/src/Unique.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef unique_h
#define unique_h

#include <stdexcept>
#include "projects/common/exceptions/exceptions.h"

/**
 * @name unique [getter]
 * @brief Unique property getter
 *
 * @return count int.
 */
inline int ArbitraryKvBtree::unique() {
    return _unique;
}

/**
 * @name unique [setter]
 * @brief Unique property setter
 *
 * @param v bool
 * @return bool
 */
inline bool ArbitraryKvBtree::unique(bool v) {
    /**
     * Note that unique setter relies upon duplicates()
     * which could potentially block for a stats-cache
     * refresh if the cache has been invalidated by
     * node deletions, etc.
     */
    if (!_unique && (duplicates() > 0)) {
        throw std::runtime_error("duplicate nodes prevent unique flag enable.");
    }
    return _unique = v;
}

#endif /*unique_h*/