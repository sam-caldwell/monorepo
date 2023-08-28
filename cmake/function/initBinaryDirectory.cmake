#
# initBinaryDirectory.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
set(CMAKE_BINARY_ROOT ${CMAKE_BINARY_DIR}/bin)
function(initBinaryDirectory)
    debug("initBinaryDirectory() running")
    if (EXISTS ${CMAKE_BINARY_ROOT})
        debug("Cleaning: ${CMAKE_BINARY_ROOT}")
        execute_process(COMMAND ${CMAKE_COMMAND} -E remove_directory ${CMAKE_BINARY_ROOT})
    endif ()
    file(MAKE_DIRECTORY ${CMAKE_BINARY_ROOT})
    debug("initBinaryDirectory() ok")
endfunction()