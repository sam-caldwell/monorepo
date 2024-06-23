cmake_minimum_required(VERSION 3.29)

set(CMAKE_INCLUDE_IF_EXISTS True)

function(include_if_exists dir)
    if(EXISTS ${CMAKE_SOURCE_DIR}/${dir}/CMakeLists.txt)
        add_subdirectory(${CMAKE_SOURCE_DIR}/${dir})
        logOK("Directory '${dir}' added")
    else()
        logFAIL("Directory '${dir}' does not exist or has no CMakeLists.txt")
    endif()
endfunction()