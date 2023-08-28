#
# debugLine.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Logging functions.
#

function(debugLine)
    string(REPEAT "-" 60 OUTPUT)
    debug("${OUTPUT}")
endfunction()
