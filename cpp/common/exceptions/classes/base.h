/**
 * @name projects/exceptions/src/base.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef EXCEPTIONS_BASE_H
#define EXCEPTIONS_BASE_H

#include "exception"
#include <iostream>
#include <cstring>
#include <strings.h>

using namespace std;
/**
 * @name BaseException
 * @brief This is the base class for our exceptions.
 *        We will overload this with specific exceptions
 *        either in this project (reusable) or in specific
 *        project.
 */
class BaseException : public exception {
protected:

    string message;

public:

    BaseException() {};

    BaseException(string msg) {
        message = msg;
    }

    char *what() const throw() {
        char *p;
        char buf[message.length() + 1];
        strcpy(buf, message.c_str());
        p=&buf[0];
        return p;
    }
};

#endif //EXCEPTIONS_BASE_H