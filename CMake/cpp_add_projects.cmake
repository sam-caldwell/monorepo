cmake_minimum_required(VERSION 3.29)

require("CMake/cpp_project_artifact.cmake")
#require("CMake/cpp_project_tests.cmake")

function(cpp_add_project name path)
    logDebug("add project '${name}'")
    project(${name})


    if (EXISTS ${path})
        cpp_project_artifact(${name} ${path})
        cpp_project_tests(${name} ${path})
        logOk("project (${name}) added")
    else ()
        logFATAL("project (${name}) not found: ${path}")
    endif ()
endfunction()
