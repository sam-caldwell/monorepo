/**
 * @name exceptions/classes/MissingInputs.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef MissingInputs_H
#define MissingInputs_H

#include "base.h"
#include <iostream>
#include <string>

using namespace std;

class MissingInputs : public BaseException {
public:
    MissingInputs(string msg) : BaseException(msg) {}
};

#endif /* MissingInputs_H */