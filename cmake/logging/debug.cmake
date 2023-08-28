#
# debug.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Logging functions.
#

function(debug msg)
    if (CMAKE_DEBUG)
        include(cmake/logging/colors.cmake)
        message("${Yellow}     # [DEBUG]:${msg}${ColourReset}")
    endif ()
endfunction()
