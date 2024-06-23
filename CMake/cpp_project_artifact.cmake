cmake_minimum_required(VERSION 3.29)

function(cpp_project_artifact name root)

    set(MAIN_CPP "${root}/${name}/main.cpp")

    logDebug("cpp_project_artifact: ${name}: '${MAIN_CPP}'")

    if (EXISTS ${MAIN_CPP})
        add_executable("${name}" "${MAIN_CPP}")
    else ()
        logFatal("${MAIN_CPP} not found")
    endif ()
endfunction()
