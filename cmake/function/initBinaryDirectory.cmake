#
# initBinaryDirectory.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
set(CMAKE_BINARY_ROOT ${CMAKE_BINARY_DIR}/bin)
function(initBinaryDirectory)
    debug("initBinaryDirectory() running")
    deleteIfExists(${CMAKE_BINARY_ROOT})
    file(MAKE_DIRECTORY ${CMAKE_BINARY_ROOT})
    ok("initBinaryDirectory() ok")
endfunction()