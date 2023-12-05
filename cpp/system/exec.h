/**
 * @name common/system/exec.h
 * @copyright (c) 2022 Sam Caldwell.  All Rights Reserved.
 * @author Sam Caldwell <mail@samcaldwell.net>
 */
#ifndef Common_System_exec_h
#define Common_System_exec_h

#include <iostream>
#include <stdio.h>
#include <stdlib.h>

namespace exec {
    /**
     * @name exec::run
     * @brief a pretty wrapper for executing things at the commandline (shell)
     *        and capturing the stdout/stderr output.
     * @param cmd - string
     * @param captureStderr - bool
     * @return string - output
     */
    string run(string cmd, bool captureStderr = false) {

        FILE *stream;
        const int bufferSz = 256;
        char buffer[bufferSz];
        string output;

        if (captureStderr)
            cmd.append(" 2>&1"); //Make sure we capture stderr and stdout

        stream = popen(cmd.c_str(), "r");
        if (stream) {
            while (!feof(stream))
                if (fgets(buffer, bufferSz, stream) != NULL)
                    output.append(buffer);
            pclose(stream);
        }
        return output;
    }
}

#endif /* Common_System_exec_h */