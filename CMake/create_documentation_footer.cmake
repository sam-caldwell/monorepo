cmake_minimum_required(VERSION 3.29)

function(create_documentation_footer path)

    set(README_FILE ${path}/README.md)

    file(APPEND ${README_FILE} "\n")

    string(TIMESTAMP TIMESTAMP "%Y-%m-%d %H:%M:%S")

    file(APPEND ${README_FILE} "\nGenerated at ${TIMESTAMP}\n")

    logDebug("documentation footer written")

endfunction()
