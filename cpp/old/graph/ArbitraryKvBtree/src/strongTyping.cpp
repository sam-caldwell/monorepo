/**
 * @name ArbitraryKvBtree/src/strongTyping.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef strongTyping_h
#define strongTyping_h

#include <stdexcept>
#include "../../../../exceptions/exceptions.h"

/**
 * @name strongTyping [getter]
 * @brief strongTyping property getter
 *
 * @return  bool
 */
inline bool ArbitraryKvBtree::strongTyping() {
    return _strongTyping;
}

/**
 * @name strongTyping [setter]
 * @brief strongTyping property setter
 *
 * @param v bool
 * @return bool
 */
inline bool ArbitraryKvBtree::strongTyping(bool v) {
    return _strongTyping = v;
}

#endif /* strongTyping_h */