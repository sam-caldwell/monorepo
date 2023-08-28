#
# initBuildDirectory.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
set(CMAKE_BUILD_ROOT ${CMAKE_BINARY_DIR}/build)
function(initBuildDirectory)
    debug("initBuildDirectory() running")
    deleteIfExists(${CMAKE_BINARY_ROOT})
    file(MAKE_DIRECTORY ${CMAKE_BUILD_ROOT})
    ok("initBuildDirectory() ok")
endfunction()