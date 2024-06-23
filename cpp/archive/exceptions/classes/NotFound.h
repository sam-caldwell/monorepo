/**
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 *
 */
#ifndef NOT_FOUND_H
#define NOT_FOUND_H

#include <stdio.h>
#include <stdlib.h>
#include <string>

using namespace std;

class NotFound : public BaseException {
public:
    NotFound() {};
};

#endif //NOT_FOUND_H