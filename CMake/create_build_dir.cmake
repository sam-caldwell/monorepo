cmake_minimum_required(VERSION 3.29)

set(CMAKE_CREATE_BUILD_DIR True)

# Remove build directory if it exists then recreate it
function(create_build_dir)
    logDebug("Build dir '${CMAKE_BUILD_DIR}'")
    if(EXISTS ${CMAKE_BUILD_DIR})
        logDebug("Deleting '${CMAKE_BUILD_DIR}'")
        file(REMOVE_RECURSE ${CMAKE_BUILD_DIR})
        logOK("Existing directory deleted (${CMAKE_BUILD_DIR})")
    else()
        logOK("Build Dir does not exist")
    endif()
    file(MAKE_DIRECTORY ${CMAKE_BUILD_DIR})
    logOK("Build Dir created (${CMAKE_BUILD_DIR})")
endfunction()
