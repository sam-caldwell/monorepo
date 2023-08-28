#
# deleteIfExists.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
function(deleteIfExists TARGET_DIRECTORY)
    set(CMAKE_DEBUG True)
    if (EXISTS "${TARGET_DIRECTORY}")
        debug("Deleting : '${TARGET_DIRECTORY}'")
        execute_process(COMMAND ${CMAKE_COMMAND} -E remove_directory "${TARGET_DIRECTORY}")
        ok("Deleted '${TARGET_DIRECTORY}'")
    endif ()
endfunction()