cmake_minimum_required(VERSION 3.29)
project(monorepo)

enable_testing()

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
include("CMake/create_documentation_header.cmake")
include("CMake/create_documentation_line.cmake")
include("CMake/create_documentation_footer.cmake")
include("CMake/add_go_projects.cmake")
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

if (DEFINED GOLANG_ENABLED)
#    add_go_projects(${CMAKE_SOURCE_DIR}/go)
endif ()
