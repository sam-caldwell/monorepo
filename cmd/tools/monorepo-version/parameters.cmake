#
# go/main.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
function(configure_project)
    set(PROJECT_ENABLED, True)
    #
    # For non-tool binaries...
    #
    set(TARGET_DIRECTORY ${CMAKE_BUILD_ROOT})
    return(PROPAGATE PROJECT_ENABLED)
    return(PROPAGATE TARGET_DIRECTORY)
endfunction()