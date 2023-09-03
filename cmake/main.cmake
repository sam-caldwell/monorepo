#
# configureGolang.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file is the single point of including functions
# for use in the CMakeLists.txt file.
#
include(cmake/logging/debug.cmake)
include(cmake/logging/debugList.cmake)
include(cmake/logging/debugLine.cmake)
include(cmake/logging/error.cmake)
include(cmake/logging/info.cmake)
include(cmake/logging/ok.cmake)

include(cmake/function/deleteIfExists.cmake)
include(cmake/function/initBinaryDirectory.cmake)
include(cmake/function/initBuildDirectory.cmake)
include(cmake/function/showMonorepoVersion.cmake)
include(cmake/function/addProjectParameters.cmake)
include(cmake/function/copyModulesLocal.cmake)

include(cmake/os/apple.cmake)
include(cmake/os/unix.cmake)
include(cmake/os/windows.cmake)

include(cpp/cmake/configure.cmake)
include(go/cmake/buildTargets.cmake)

include(cpp/cmake/cppBuildTargets.cmake)
include(go/cmake/buildTargets.cmake)
