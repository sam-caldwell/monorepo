#
# debugList.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Logging functions.
#

function(debugList HEADER , LIST)
    debug("--- ${HEADER} ---")
    if (CMAKE_DEBUG)
        foreach (item ${LIST})
            debug(" - ${item} ")
        endforeach ()
    endif()
    debug("-------------------")
endfunction()
