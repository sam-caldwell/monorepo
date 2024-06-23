cmake_minimum_required(VERSION 3.29)

function(cpp_project_tests name root)

    set(TEST_CPP "${root}/${name}/test.cpp")

    logDebug("cpp_project_tests: ${name}: '${MAIN_CPP}'")

    if (EXISTS ${TEST_CPP})
        add_executable("${name}_test" "${TEST_CPP}")
    else ()
        logFatal("${TEST_CPP} not found")
    endif ()
endfunction()
