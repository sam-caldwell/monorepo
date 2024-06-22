/**
 * @name projects/common/exceptions/classes/UnexpectedType.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef UNEXPECTED_TYPE_ERROR_H
#define UNEXPECTED_TYPE_ERROR_H

#include <stdio.h>
#include "../../../types/ValueTypes.h"
#include "base.h"
#include <stdlib.h>

class UnexpectedType : public BaseException {
public:
    UnexpectedType() {};

    UnexpectedType(ValueTypes t) {
        message = (char *) (("UnexpectedType():" + to_string((uint)(t))).c_str());
    }
};

#endif //UNEXPECTED_TYPE_ERROR_H
