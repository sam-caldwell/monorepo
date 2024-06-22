/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 * An exception which will be thrown by Channel objects if the
 * channel is closed.
 */
#ifndef closedChannelException_H
#define closedChannelException_H

#include <iostream>
#include "base.h"

using namespace std;

class closedChannelException : public BaseException {
public:
    explicit closedChannelException() {};

    explicit closedChannelException(string msg) {};
};

#endif /**closedChannelException_H*/