cmake_minimum_required(VERSION 3.29)

function(create_documentation_header root)
    set(README_FILE ${root}/README.md)
    set(DOC_LINES
            "C++ Project Document Root"
            "========================="
            "(c) Sam Caldwell.  See LICENSE.txt.\n"
            ""
            "## Objectives"
            "  This README file provides a root for all C++ project documentation"
            "  each project should be referenced (and linked if appropriate)"
            "\n"
            "## Projects")

    file(WRITE ${README_FILE})

    foreach (LINE ${DOC_LINES})

        file(APPEND ${README_FILE} "${LINE}\n")

    endforeach ()
    logDebug("documentation header written")
endfunction()
