/**
 * @name exceptions/tests/LockTimeoutExpired.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 */
#ifndef LOCK_TIMEOUT_EXPIRED_H
#define LOCK_TIMEOUT_EXPIRED_H

#include <stdio.h>
#include <stdlib.h>
#include <string>

using namespace std;

class LockTimeoutExpired : public BaseException {
public:
    LockTimeoutExpired(){};
};

#endif //LOCK_TIMEOUT_EXPIRED_H
