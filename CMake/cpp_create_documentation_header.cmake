cmake_minimum_required(VERSION 3.29)

function(cpp_create_documentation_header readmeFile)

    set(DOC_LINES
            "C++ Project Document Root"
            "========================="
            "(c) Sam Caldwell.  See LICENSE.txt."
            ""
            "## Objectives"
            "  This README file provides a root for all C++ project documentation"
            "  each project should be referenced (and linked if appropriate)"
            ""
            "## Projects")

    file(WRITE ${readmeFile})
    foreach (LINE ${DOC_LINES})
        file(APPEND ${readmeFile} "${LINE}\n")
    endforeach ()
endfunction()

function(cpp_create_documentation_line readmeFile project_path project_name)
    logDebug("Checking project_path: ${project_path}")
    if (EXISTS cpp/${project_path}/README.md)
        file(APPEND ${readmeFile} "* [${project_name}](${project_path}/README.md)\n")
    else()
        logWARN("project (${project_name}) has no README.md (${readmeFile}) file")
        file(APPEND ${readmeFile} "* ${project_name}\n")
    endif ()
endfunction()

function(cpp_create_documentation_footer readmeFile)
    file(APPEND ${readmeFile} "\n\n")
    string(TIMESTAMP TIMESTAMP "%Y-%m-%d %H:%M:%S")
    file(APPEND ${readmeFile} "\nGenerated at ${TIMESTAMP}\n")
endfunction()


