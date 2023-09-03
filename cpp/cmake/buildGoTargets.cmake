#
# buildGoTargets.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#

function(buildGoTargets source_directory)
    debug("cppBuildTargets: start")
    #
    # Define executables
    #
    file(GLOB_RECURSE HEADER_FILES "${source_directory}/*.h")
    file(GLOB_RECURSE CPP_FILES "${source_directory}/main.go")
    #
    # Assume <project_name>/<program_name>/main.cpp
    #
    debug("Looping through source files...")
    foreach (source_code_file ${CPP_FILES})
        get_filename_component(project_directory ${source_code_file} DIRECTORY)
        get_filename_component(program_name ${project_directory} NAME_WE)
        set(PROJECT_PARAMETERS ${project_directory}/parameters.cmake)
        addProjectParameters(${PROJECT_PARAMETERS})
        debug("       program_name: ${program_name}")
        if(PROJECT_ENABLED)
            debug("            source: ${source_code_file}")
            debug(" project_directory: ${project_directory}")
            debug("  build parameters: ${PROJECT_PARAMETERS}")
            debug("   target directory: ${TARGET_DIRECTORY}")
            info("Project (${program_name}): enabled")
            add_executable("${program_name}" "${source_code_file}" ${HEADER_FILES})
            set_target_properties("${program_name}"
                    PROPERTIES RUNTIME_OUTPUT_DIRECTORY ${TARGET_DIRECTORY}
                    LINKER_LANGUAGE CXX)
        else()
            info("Project (${program_name}): disabled")
        endif ()
        debugLine()
    endforeach ()
    debug("cppBuildTargets: done")
    debugLine()
endfunction()