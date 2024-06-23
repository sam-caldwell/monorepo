cmake_minimum_required(VERSION 3.29)

set(CMAKE_COLORS True)

# Function to print a message in red text
function(Red text)
    if (UNIX OR CYGWIN)
        execute_process(COMMAND echo -e "\\033[31m${text}\\033[0m")
    elseif (WIN32)
        # Windows may not support ANSI codes by default in cmd.exe, but works in some terminals like PowerShell or MinGW
        message("${text}")
    else ()
        message("${text}")
    endif ()
endfunction()

# Function to print a message in green text
function(Green text)
    if (UNIX OR CYGWIN)
        execute_process(COMMAND echo -e "\\033[32m${text}\\033[0m")
    elseif (WIN32)
        message("${text}")
    else ()
        message("${text}")
    endif ()
endfunction()

# Function to print a message in yellow text
function(Yellow text)
    if (UNIX OR CYGWIN)
        execute_process(COMMAND echo -e "\\033[33m${text}\\033[0m")
    elseif (WIN32)
        message("${text}")
    else ()
        message("${text}")
    endif ()
endfunction()

# Function to print a message in blue text
function(Blue text)
    if (UNIX OR CYGWIN)
        execute_process(COMMAND echo -e "\\033[34m${text}\\033[0m")
    elseif (WIN32)
        message("${text}")
    else ()
        message("${text}")
    endif ()
endfunction()

function(colors_test)
    Red("Red: Enabled")
    Yellow("Yellow: Enabled")
    Green("Green: Enabled")
    Blue("Blue: Enabled")
endfunction()

if (DEFINED DEBUG)
    Blue("COLORS: Enabled")
endif ()