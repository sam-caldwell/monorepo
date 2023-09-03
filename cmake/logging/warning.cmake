#
# warning.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Logging function.
#

function(warning msg)
    include(cmake/logging/colors.cmake)
    message("${Red}     # [WARNING]:${msg}${ColourReset}")
endfunction()
