cmake_minimum_required(VERSION 3.29)

function(go_test)
    function(go_test)
        # Set the working directory to the ./go directory
        set(go_directory "${CMAKE_SOURCE_DIR}/go")

        # Execute the go test command
        execute_process(
                COMMAND go test -v ./...
                WORKING_DIRECTORY ${go_directory}
                RESULT_VARIABLE result
        )

        # Check the result
        if(result EQUAL 0)
            LogOK("Go tests passed successfully.")
        else()
            LogFail("Go tests failed with error code ${result}.")
        endif()
    endfunction()
endfunction()

logDebug("golang projects are included")