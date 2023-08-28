#
# initBuildDirectory.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
set(CMAKE_BUILD_ROOT ${CMAKE_BINARY_DIR}/build)
function(initBuildDirectory)
    debug("initBuildDirectory() running")
    if (EXISTS ${CMAKE_BUILD_ROOT})
        debug("Cleaning: ${CMAKE_BUILD_ROOT}")
        execute_process(COMMAND ${CMAKE_COMMAND} -E remove_directory ${CMAKE_BUILD_ROOT})
    endif ()
    file(MAKE_DIRECTORY ${CMAKE_BUILD_ROOT})
    debug("initBuildDirectory() ok")
endfunction()