cmake_minimum_required(VERSION 3.29)

function(create_documentation_line root project_list)
    set(COUNTER 0 PARENT_SCOPE)
    set(README_FILE ${root}/README.md)
    foreach (TEST_CPP_FILE ${CPP_FILES})

        get_filename_component(PROJECT_PATH ${TEST_CPP_FILE} DIRECTORY)
        get_filename_component(PROJECT_NAME ${PROJECT_DIR} NAME)
        set(PROJECT_README, "${root}/${PROJECT_PATH}/README.md")

        logDebug("root:           '${root}'")
        logDebug("PROJECT_PATH:   '${PROJECT_PATH}'")
        logDebug("PROJECT_NAME:   '${PROJECT_NAME}'")

        logDebug("ROOT_README:    '${ROOT_README}'")
        logDebug("PROJECT_README: '${PROJECT_README}'")

        if (EXISTS ${PROJECT_README})
            file(APPEND "${README_FILE}" "* [${PROJECT_NAME}](${PROJECT_README})\n")
        else ()

            logWARN("project (${PROJECT_NAME}) has no README.md (${PROJECT_README}) file")

            file(APPEND "${README_FILE}" "* ${PROJECT_NAME}\n")

        endif ()
        math(EXPR COUNTER "${COUNTER} + 1")
        set(COUNTER ${COUNTER} PARENT_SCOPE)
    endforeach ()
    if(${COUNTER} EQUAL 0)
        file(APPEND ${README_FILE} "* ${PROJECT_NAME}")
    endif()
endfunction()
