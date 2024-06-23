cmake_minimum_required(VERSION 3.29)
project(monorepo)

set(DEBUG TRUE)
set(CMAKE_BUILD_DIR "${CMAKE_BINARY_DIR}/build")
set(CPP_ENABLED TRUE)
#set(GO_ENABLED TRUE)


#
# --------------------------------
# Setup the reusable CMAKE tooling
# --------------------------------
#

include("CMake/colors.cmake")
include("CMake/logger.cmake")
include("CMake/require.cmake")
include("CMake/include_if_exists.cmake")
include("CMake/create_build_dir.cmake")
include("CMake/cpp_add_project.cmake")
include("CMake/cpp_project_tests.cmake")

#
# -----------------------
# Create our build target
# -----------------------
#

create_build_dir()

#
# ---------------------------------------
# include language-specific project files
# ---------------------------------------
#

if (DEFINED CPP_ENABLED)
    include_if_exists(cpp)
endif ()
if (DEFINED GOLANG_ENABLED)
    include_if_exists(go)
endif ()
# Add more project types here
