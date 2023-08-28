#
# error.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Logging function.
#

function(error msg)
    include(cmake/logging/colors.cmake)
    message(ERROR "${Red}     # [ERROR]:${msg}${ColourReset}")
endfunction()
