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

debug("CMakeDetermineGOCompiler.cmake: start")

include(${CMAKE_ROOT}/Modules/CMakeDetermineCompiler.cmake)

if( NOT CMAKE_GO_COMPILER )
    debug("CMAKE_GO_COMPILER not set...")
    find_program(CMAKE_GO_COMPILER "go" "/")
    get_filename_component(go_bin_dir ${CMAKE_GO_COMPILER} DIRECTORY)
    get_filename_component(GOROOT ${go_bin_dir} DIRECTORY)
endif ()

set(GLOBAL PROPERTY GOROOT "${CMAKE_GO_COMPILER}")
set(GLOBAL PROPERTY GOPATH "${CMAKE_BINARY_DIR}")
set(ENV{GOPATH} "${CMAKE_BINARY_DIR}")

set(ENV{CGO_ENABLED} 0)
set(ENV{CGO_LDFLAGS} "")
set(ENV{CGO_CFLAGS} "")

ok("CMAKE_GO_COMPILER: '${CMAKE_GO_COMPILER}'")
ok("           GOROOT: '${GOROOT}'")

# configure variables set in this file for fast reload later on
configure_file(${CMAKE_BINARY_DIR}/cmake/Modules/CMakeGOCompiler.cmake.in
        ${CMAKE_PLATFORM_INFO_DIR}/CMakeGOCompiler.cmake
        @ONLY
)
#set(CMAKE_GO_COMPILER_ENV_VAR "GO")


ok("CMakeDetermineGOCompiler.cmake: done")

