#
# ok.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Logging function.
#

function(ok msg)
    include(cmake/logging/colors.cmake)
    message("${Green}     # [OK]:${msg}${ColourReset}")
endfunction()
