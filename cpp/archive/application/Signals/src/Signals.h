/**
 * @name Signals/src/Signals.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 *
 * @brief see https://man7.org/linux/man-pages/man7/signal.7.html
 */
#ifndef Signals_H
#define Signals_H

#include <map>
#include <csignal>
#include "../../../../flags/flags.h"

#if OPSYS == WINDOWS
/*  */#error "Windows is not supported"
#endif

#if !((CPU_ARCH == ARCH_X32) || (CPU_ARCH == ARCH_X64) || (CPU_ARCH == ARCH_ARM32) || (CPU_ARCH == ARCH_ARM64))
/*  */#error "Only x86 or ARM are supported"
#endif


namespace SignalHandler {
    /**
     * @typedef SignalHandler::Func(int)
     * @brief This is a function pointer to a signal handler function.
     */
    typedef void (*Func)(int);

    /**
     * @typedef SignalHandler::Table
     * @brief This is a table mapping interrupt signals (int) to
     *        signal handler functions (SignalHandler::Func).
     */
    typedef map<int, Func> Table;

    /**
     * @name SignalHandler::init
     * @brief This function will iterate over a signal handler map
     *        table (SignalHandler::Table) and initialize the operating
     *        system signal handler to invoke the mapped handler for a
     *        given signal.
     *
     * @param signals
     */
    void init(SignalHandler::Table *signals) {
        for (auto it = signals->begin(); it != signals->end(); ++it)
            signal(it->first, it->second);
    }
}
#endif /* Signals_H */