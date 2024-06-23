cmake_minimum_required(VERSION 3.29)

function(cpp_add_projects path)
    logDebug("cpp_add_projects() start")

    file(GLOB_RECURSE CPP_FILES RELATIVE ${path} "*/test.cpp")

    foreach (TEST_CPP_FILE ${CPP_FILES})
        if (TEST_CPP_FILE MATCHES "archive/")
            yellow("cpp_add_projects() skip archive subdirectory")
            continue()
        endif ()
        logDebug("cpp_add_projects() iterate over TEST_CPP_FILE: ${TEST_CPP_FILE}")
        get_filename_component(PROJECT_DIR ${TEST_CPP_FILE} DIRECTORY)
        logDebug("cpp_add_projects() iterate over PROJECT_DIR: ${PROJECT_DIR}")
        get_filename_component(PROJECT_NAME ${PROJECT_DIR} NAME)
        logDebug("cpp_add_projects() iterate over PROJECT_NAME: ${PROJECT_NAME}")

        if (EXISTS ${path}/project.disabled)
            yellow("${PROJECT_NAME}: disabled")
            continue()
        endif ()

        set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BUILD_DIR}/${PROJECT_NAME})

        project(${PROJECT_NAME})
        logDebug("add project '${PROJECT_NAME}' (${path}/${TEST_CPP_FILE})")
        add_executable(${PROJECT_NAME}_test ${path}/${TEST_CPP_FILE})
        set_target_properties(${PROJECT_NAME}_test PROPERTIES
                RUNTIME_OUTPUT_DIRECTORY ${CMAKE_RUNTIME_OUTPUT_DIRECTORY})

        add_custom_target(test_${PROJECT_NAME}
                COMMAND ${PROJECT_NAME}_test
                DEPENDS ${PROJECT_NAME}_test
                WORKING_DIRECTORY ${CMAKE_BUILD_DIR}/${PROJECT_NAME}
                COMMENT "Running color tests"
        )

        set(MAIN_CPP_FILE ${PROJECT_DIR}/main.cpp)
        logDebug("       main '${PROJECT_NAME}' (${path}/${MAIN_CPP_FILE})")
        if (EXISTS ${path}/${MAIN_CPP_FILE})
            add_executable(${PROJECT_NAME} ${path}/${MAIN_CPP_FILE})
            set_target_properties(${PROJECT_NAME} PROPERTIES
                    RUNTIME_OUTPUT_DIRECTORY ${CMAKE_RUNTIME_OUTPUT_DIRECTORY})
        endif ()
    endforeach ()
    logDebug("cpp_add_projects() terminate")
endfunction()
