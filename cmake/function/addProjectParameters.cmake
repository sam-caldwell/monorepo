#
# addProjectParameters.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#

function(addProjectParameters PROJECT_PARAMETERS )
    if (EXISTS ${PROJECT_PARAMETERS})
        include(${PROJECT_PARAMETERS})
    else()
        set(TARGET_DIRECTORY ${CMAKE_BUILD_ROOT})
        info("  project parameters: <not found>")
    endif ()
    return(PROPAGATE TARGET_DIRECTORY)
endfunction()
