#
# copyModulesLocal.cmake
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
function(copyModulesLocal)
    debug("copyModulesLocal() start")
    debug("Copy cmake modules from local modules directory")
    debug("source     : ${CMAKE_ROOT}/Modules/")
    debug("destination: ${CMAKE_SOURCE_DIR}/cmake/Modules/")

    set(SOURCE_DIRECTORY "${CMAKE_ROOT}/Modules/")
    set(TARGET_DIRECTORY "${CMAKE_SOURCE_DIR}/cmake/Modules/")

    deleteIfExists("${TARGET_DIRECTORY}")
    file(MAKE_DIRECTORY "${TARGET_DIRECTORY}")
    ok("Created empty ${TARGET_DIRECTORY}")

    file(COPY "${SOURCE_DIRECTORY}" DESTINATION "${TARGET_DIRECTORY}")
    debug("${SOURCE_DIRECTORY} copied to ${TARGET_DIRECTORY}")

    debug("Copy custom make modules into local modules directory")
    set(SOURCE_DIRECTORY "${CMAKE_SOURCE_DIR}/cmake/CustomModules/")
    file(COPY "${SOURCE_DIRECTORY}" DESTINATION "${TARGET_DIRECTORY}")
    debug("${SOURCE_DIRECTORY} copied")

    set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_SOURCE_DIR}/cmake/Modules/")
    info("Update CMAKE Modules path to ${CMAKE_SOURCE_DIR}/cmake/Modules")
    ok("copyModulesLocal() done")
    return(PROPAGATE CMAKE_MODULE_PATH)
endfunction()