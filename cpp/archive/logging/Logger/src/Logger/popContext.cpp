/**
 * @name Logger/src/Logger/popContext.cpp
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Logger_Logger_popContext_cpp
#define Logger_Logger_popContext_cpp

/**
 * @name popContext
 * @brief remove the current context (which will destroy this context).
 *
 * @return bool
 */
bool Logger::popContext() {
    if (context.size() == 1) {
        //Do not allow deletion of the root context
        throw out_of_range("logger ContextId 0 cannot be removed.");
    } else {
        context.pop();
    }
    return true;
}

#endif /* Logger_Logger_popContext_cpp */
