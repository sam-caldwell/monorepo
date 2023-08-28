# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This module is part of github.com/sam-caldwell/monorepo/ and is released under
# the MIT LICENSE (See https://github.com/sam-caldwell/monorepo/LICENSE.txt.  Nothing
# in this file should be interpreted as interfering or conflicting with the license or
# rights conveyed by the wider CMAKE project under its respective license.
#
# To edit this file, edit monorepo/cmake/CustomModules/CMakeGOInformation.cmake.
# The version in the cmake Modules/ directory will be overwritten when cmake is run.
#

# This file sets the basic flags for Golang in cmake and loads available platform
# information where possible.

include(CMakeLanguageInformation)

set(CGO_RESOLVE_BLACKLIST
        pthread
        rt
        gcov
        systemd
)

set(CGO_CFLAGS_BLACKLIST
        "-Werror"
        "-Wall"
        "-Wextra"
        "-Wold-style-definition"
        "-fdiagnostics-color=always"
        "-Wformat-nonliteral"
        "-Wformat=2"
)


set(CMAKE_BASE_NAME)
get_filename_component(CMAKE_BASE_NAME "${CMAKE_GO_COMPILER}" NAME_WE)

# load any compiler-wrapper specific information
if (CMAKE_GO_COMPILER_WRAPPER)
    __cmake_include_compiler_wrapper(GO)
endif ()

# Remove this when all compiler info is removed from system files.
if (NOT _INCLUDED_FILE)
    include(Platform/${CMAKE_SYSTEM_NAME} OPTIONAL)
endif ()

if(CMAKE_GO_SIZEOF_DATA_PTR)
    foreach(f ${CMAKE_GO_ABI_FILES})
        include(${f})
    endforeach()
    unset(CMAKE_GO_ABI_FILES)
endif()

if(CMAKE_USER_MAKE_RULES_OVERRIDE)
    include(${CMAKE_USER_MAKE_RULES_OVERRIDE} RESULT_VARIABLE _override)
    set(CMAKE_USER_MAKE_RULES_OVERRIDE "${_override}")
endif()

if(CMAKE_USER_MAKE_RULES_OVERRIDE_GO)
    include(${CMAKE_USER_MAKE_RULES_OVERRIDE_GO} RESULT_VARIABLE _override)
    set(CMAKE_USER_MAKE_RULES_OVERRIDE_GO "${_override}")
endif()

if(NOT CMAKE_GO_COMPILE_OPTIONS_PIC)
    set(CMAKE_GO_COMPILE_OPTIONS_PIC ${CMAKE_C_COMPILE_OPTIONS_PIC})
endif()

if(NOT CMAKE_GO_COMPILE_OPTIONS_PIE)
    set(CMAKE_GO_COMPILE_OPTIONS_PIE ${CMAKE_C_COMPILE_OPTIONS_PIE})
endif()

if(NOT CMAKE_GO_LINK_OPTIONS_PIE)
    set(CMAKE_GO_LINK_OPTIONS_PIE ${CMAKE_C_LINK_OPTIONS_PIE})
endif()

if(NOT CMAKE_GO_LINK_OPTIONS_NO_PIE)
    set(CMAKE_GO_LINK_OPTIONS_NO_PIE ${CMAKE_C_LINK_OPTIONS_NO_PIE})
endif()

if(NOT CMAKE_GO_COMPILE_OPTIONS_DLL)
    set(CMAKE_GO_COMPILE_OPTIONS_DLL ${CMAKE_C_COMPILE_OPTIONS_DLL})
endif()

if(NOT DEFINED CMAKE_SHARED_LIBRARY_CREATE_GO_FLAGS)
    set(CMAKE_SHARED_LIBRARY_CREATE_GO_FLAGS ${CMAKE_SHARED_LIBRARY_CREATE_C_FLAGS})
endif()

if(NOT DEFINED CMAKE_SHARED_LIBRARY_GO_FLAGS)
    set(CMAKE_SHARED_LIBRARY_GO_FLAGS ${CMAKE_SHARED_LIBRARY_C_FLAGS})
endif()

if(NOT DEFINED CMAKE_SHARED_LIBRARY_LINK_GO_FLAGS)
    set(CMAKE_SHARED_LIBRARY_LINK_GO_FLAGS ${CMAKE_SHARED_LIBRARY_LINK_C_FLAGS})
endif()

if(NOT DEFINED CMAKE_SHARED_LIBRARY_RUNTIME_GO_FLAG)
    set(CMAKE_SHARED_LIBRARY_RUNTIME_GO_FLAG ${CMAKE_SHARED_LIBRARY_RUNTIME_C_FLAG})
endif()

if(NOT DEFINED CMAKE_SHARED_LIBRARY_RUNTIME_GO_FLAG_SEP)
    set(CMAKE_SHARED_LIBRARY_RUNTIME_GO_FLAG_SEP ${CMAKE_SHARED_LIBRARY_RUNTIME_C_FLAG_SEP})
endif()

if(NOT DEFINED CMAKE_SHARED_LIBRARY_RPATH_LINK_GO_FLAG)
    set(CMAKE_SHARED_LIBRARY_RPATH_LINK_GO_FLAG ${CMAKE_SHARED_LIBRARY_RPATH_LINK_C_FLAG})
endif()

if(NOT DEFINED CMAKE_EXE_EXPORTS_GO_FLAG)
    set(CMAKE_EXE_EXPORTS_GO_FLAG ${CMAKE_EXE_EXPORTS_C_FLAG})
endif()

if(NOT DEFINED CMAKE_SHARED_LIBRARY_SONAME_GO_FLAG)
    set(CMAKE_SHARED_LIBRARY_SONAME_GO_FLAG ${CMAKE_SHARED_LIBRARY_SONAME_C_FLAG})
endif()

if(NOT CMAKE_MODULE_EXISTS)
    set(CMAKE_SHARED_MODULE_GO_FLAGS ${CMAKE_SHARED_LIBRARY_GO_FLAGS})
    set(CMAKE_SHARED_MODULE_CREATE_GO_FLAGS ${CMAKE_SHARED_LIBRARY_CREATE_GO_FLAGS})
endif()

if(NOT DEFINED CMAKE_SHARED_MODULE_CREATE_GO_FLAGS)
    set(CMAKE_SHARED_MODULE_CREATE_GO_FLAGS ${CMAKE_SHARED_MODULE_CREATE_C_FLAGS})
endif()

if(NOT DEFINED CMAKE_SHARED_MODULE_GO_FLAGS)
    set(CMAKE_SHARED_MODULE_GO_FLAGS ${CMAKE_SHARED_MODULE_C_FLAGS})
endif()

if(NOT DEFINED CMAKE_EXECUTABLE_RUNTIME_GO_FLAG)
    set(CMAKE_EXECUTABLE_RUNTIME_GO_FLAG ${CMAKE_SHARED_LIBRARY_RUNTIME_GO_FLAG})
endif()

if(NOT DEFINED CMAKE_EXECUTABLE_RUNTIME_GO_FLAG_SEP)
    set(CMAKE_EXECUTABLE_RUNTIME_GO_FLAG_SEP ${CMAKE_SHARED_LIBRARY_RUNTIME_GO_FLAG_SEP})
endif()

if(NOT DEFINED CMAKE_EXECUTABLE_RPATH_LINK_GO_FLAG)
    set(CMAKE_EXECUTABLE_RPATH_LINK_GO_FLAG ${CMAKE_SHARED_LIBRARY_RPATH_LINK_GO_FLAG})
endif()

if(NOT DEFINED CMAKE_SHARED_LIBRARY_LINK_GO_WITH_RUNTIME_PATH)
    set(CMAKE_SHARED_LIBRARY_LINK_GO_WITH_RUNTIME_PATH ${CMAKE_SHARED_LIBRARY_LINK_C_WITH_RUNTIME_PATH})
endif()

if(NOT CMAKE_INCLUDE_FLAG_GO)
    set(CMAKE_INCLUDE_FLAG_GO ${CMAKE_INCLUDE_FLAG_C})
endif()

if(CMAKE_EXECUTABLE_FORMAT STREQUAL "ELF")
    if(NOT DEFINED CMAKE_GO_LINK_WHAT_YOU_USE_FLAG)
        set(CMAKE_GO_LINK_WHAT_YOU_USE_FLAG "LINKER:--no-as-needed")
    endif()
    if(NOT DEFINED CMAKE_LINK_WHAT_YOU_USE_CHECK)
        set(CMAKE_LINK_WHAT_YOU_USE_CHECK ldd -u -r)
    endif()
endif()

set(CMAKE_VERBOSE_MAKEFILE FALSE CACHE BOOL "If this value is on, makefiles will be generated without the .SILENT directive, and all commands will be echoed to the console during the make.  This is useful for debugging only. With Visual Studio IDE projects all commands are done without /nologo.")

set(CMAKE_GO_FLAGS_INIT "$ENV{FFLAGS} ${CMAKE_GO_FLAGS_INIT}")

cmake_initialize_per_config_variable(CMAKE_GO_FLAGS "Flags used by the GO compiler")

if(NOT CMAKE_GO_COMPILER_LAUNCHER AND DEFINED ENV{CMAKE_GO_COMPILER_LAUNCHER})
    set(CMAKE_GO_COMPILER_LAUNCHER "$ENV{CMAKE_GO_COMPILER_LAUNCHER}"
            CACHE STRING "Compiler launcher for GO.")
endif()

include(CMakeCommonLanguageInclude)


# create a GO shared library
if(NOT CMAKE_GO_CREATE_SHARED_LIBRARY)
    set(CMAKE_GO_CREATE_SHARED_LIBRARY
            "<CMAKE_GO_COMPILER> <CMAKE_SHARED_LIBRARY_GO_FLAGS> <LANGUAGE_COMPILE_FLAGS> <LINK_FLAGS> <CMAKE_SHARED_LIBRARY_CREATE_GO_FLAGS> <SONAME_FLAG><TARGET_SONAME> -o <TARGET> <OBJECTS> <LINK_LIBRARIES>")
endif()

# create a GO shared module just copy the shared library rule
if(NOT CMAKE_GO_CREATE_SHARED_MODULE)
    set(CMAKE_GO_CREATE_SHARED_MODULE ${CMAKE_GO_CREATE_SHARED_LIBRARY})
endif()

# Create a static archive incrementally for large object file counts.
# If CMAKE_GO_CREATE_STATIC_LIBRARY is set it will override these.
if(NOT DEFINED CMAKE_GO_ARCHIVE_CREATE)
    set(CMAKE_GO_ARCHIVE_CREATE "<CMAKE_AR> qc <TARGET> <LINK_FLAGS> <OBJECTS>")
endif()
if(NOT DEFINED CMAKE_GO_ARCHIVE_APPEND)
    set(CMAKE_GO_ARCHIVE_APPEND "<CMAKE_AR> q <TARGET> <LINK_FLAGS> <OBJECTS>")
endif()
if(NOT DEFINED CMAKE_GO_ARCHIVE_FINISH)
    set(CMAKE_GO_ARCHIVE_FINISH "<CMAKE_RANLIB> <TARGET>")
endif()

# compile a GO file into an object file
# (put -o after -c to workaround bug in at least one mpif77 wrapper)
if(NOT CMAKE_GO_COMPILE_OBJECT)
    set(CMAKE_GO_COMPILE_OBJECT
            "<CMAKE_GO_COMPILER> <DEFINES> <INCLUDES> <FLAGS> -c <SOURCE> -o <OBJECT>")
endif()

# link a GO program
if(NOT CMAKE_GO_LINK_EXECUTABLE)
    set(CMAKE_GO_LINK_EXECUTABLE
            "<CMAKE_GO_COMPILER> <CMAKE_GO_LINK_FLAGS> <LINK_FLAGS> <FLAGS> <OBJECTS> -o <TARGET> <LINK_LIBRARIES>")
endif()

if(CMAKE_GO_STANDARD_LIBRARIES_INIT)
    set(CMAKE_GO_STANDARD_LIBRARIES "${CMAKE_GO_STANDARD_LIBRARIES_INIT}"
            CACHE STRING "Libraries linked by default with all GO applications.")
    mark_as_advanced(CMAKE_GO_STANDARD_LIBRARIES)
endif()

# set this variable so we can avoid loading this more than once.
set(CMAKE_GO_INFORMATION_LOADED 1)
