#
# info.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Logging function.
#

function(info msg)
    include(cmake/logging/colors.cmake)
    message("${Blue}     # [INFO]:${msg}${ColourReset}")
endfunction()
