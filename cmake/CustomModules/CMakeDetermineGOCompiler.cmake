# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This module is part of github.com/sam-caldwell/monorepo/ and is released under
# the MIT LICENSE (See https://github.com/sam-caldwell/monorepo/LICENSE.txt.  Nothing
# in this file should be interpreted as interfering or conflicting with the license or
# rights conveyed by the wider CMAKE project under its respective license.
#
# To edit this file, edit monorepo/cmake/CustomModules/CMakeDetermineGOCompiler.cmake.
# The version in the cmake Modules/ directory will be overwritten when cmake is run.
#

# 1. This file determines and configures the Golang compiler.
# 2. This determination may be forced by setting CMAKE_GO_COMPILER.
# 3. This module will configure--
# 3.a. CMAKE_GO_COMPILER

if(NOT CMAKE_GO_COMPILER)
  if(NOT $ENV{GO_COMPILER} STREQUAL "")
    get_filename_component(CMAKE_GO_COMPILER_INIT $ENV{GO_COMPILER} PROGRAM PROGRAM_ARGS CMAKE_GO_FLAGS_ENV_INIT)

    if(CMAKE_GO_FLAGS_ENV_INIT)
      set(CMAKE_GO_COMPILER_ARG1 "${CMAKE_GO_FLAGS_ENV_INIT}" CACHE STRING "First argument to Go compiler")
    endif()

    if(NOT EXISTS ${CMAKE_GO_COMPILER_INIT})
      message(SEND_ERROR "Could not find compiler set in environment variable GO_COMPILER:\n$ENV{GO_COMPILER}.")
    endif()

  endif()

  set(Go_BIN_PATH
    $ENV{GOPATH}
    $ENV{GOROOT}
    $ENV{GOROOT}/../bin
    $ENV{GO_COMPILER}
    /usr/bin
    /usr/local/bin
    )

  if(CMAKE_GO_COMPILER_INIT)
    set(CMAKE_GO_COMPILER ${CMAKE_GO_COMPILER_INIT} CACHE PATH "Go Compiler")
  else()
    find_program(CMAKE_GO_COMPILER
      NAMES go
      PATHS ${Go_BIN_PATH}
    )
    EXEC_PROGRAM(${CMAKE_GO_COMPILER} ARGS version OUTPUT_VARIABLE GOLANG_VERSION)
    STRING(REGEX MATCH "go[0-9]+.[0-9]+.[0-9]+[ /A-Za-z0-9]*" VERSION "${GOLANG_VERSION}")
    message("-- The Golang compiler identification is ${VERSION}")
    message("-- Check for working Golang compiler: ${CMAKE_GO_COMPILER}")
  endif()

endif()

mark_as_advanced(CMAKE_GO_COMPILER)

configure_file(${CMAKE_CURRENT_SOURCE_DIR}/cmake/CMakeGoCompiler.cmake.in
  ${CMAKE_PLATFORM_INFO_DIR}/CMakeGoCompiler.cmake @ONLY)

set(CMAKE_GO_COMPILER_ENV_VAR "GO_COMPILER")