package main

import (
	"github.com/sam-caldwell/monorepo/go/exit"
)

const CommandUsage = `

	monorepo -h | --help    : show usage (this message)
	monorepo -v | --version : show current version

	--- GENERAL COMMAND FORMAT ---

		monorepo [GeneralOptions] [commandGroup] [command] [parameters]
	
			GeneralOptions:
			-------------------------------------------------
			--debug : print verbose debug messages
			--color : color output with ANSI escape codes
			--noop  : perform command in 'no operation' mode 
					  for planning/testing
	
	--- PROJECT MANAGEMENT COMMANDS ---
	
		monorepo project list [--language <lang>]
			
			Print a list of all projects in the monorepo.
		
			If the --language argument is given, filter for projects 
			using the specified language.
		
		monorepo project show <projectName>
		
			Show the current RESOLVED configuration for the named
			project.
	
		monorepo project <enable|disable> <feature> <project>
	
			Enable or disable specified feature flags for a given project
		
		monorepo project delete <projectName>
		
			Delete the named project.
			
		monorepo project create [Project Options]
		
			Create a new project using the given options
	
				Project Options:
                ------------------------------------------------------------
				--description <string>       : describe the project
				--author <string>            : name the project author
				--output <executable|source> : define what the project
											   will produce (executable
											   or reusable source code)
				--enable <feature>           : Enable a feature flag
											   (linter, build, tests)
				--addPlatform <platform>     : Add a supported platform
											   for the project.

	--- BUILD HOST COMMANDS ---
		
		monorepo host list

            List the currently configured build hosts

        monorepo host show <name>

			Print the current build host configuration.

		monorepo host delete <name>

			Delete the build host configuration.

		monorepo host <enable|disable> <name>

			Enable or disable the build host configuration.

		monorepo host check <name>

			Verify connectivity to the build host.

		monorepo host create <name> [Host options]

			Create a new build host configuration using the options.

				Host Options:
                ------------------------------------------------------------
				--host <string>     : The IP address or FQDN and port
										for the target build host.
				--enabled           : enable the build host
				--username <name>   : The username used for connectivity.
									  If not specified, the current user
									  is used.  Assumes all authentication
									  is key-based.
				--platform <name>   : Specify the target host platform.
				--hypervisor <name> : Define what hypervisor is expected 
									  on remote host.
				--cri <name>        : Identify the container runtime to be
									  used with the platform (e.g. docker
									  or Kubernetes).

	--- CONTAINER RUNTIME INTERFACE (CRI) COMMANDS ---

		monorepo cri list

			list the supported container-runtimes.

		monorepo cri <enable|disable> <name>

			Enable or disable the build cri configuration.

		monorepo cri show <name>

			print the container-runtime configuration.

		monorepo cri delete <name>

			delete the current container runtime configuration.

		monorepo cri create <name> [CRI Options]

				CRI Options:
                ------------------------------------------------------------
				--enabled           : Enable the CRI. (Default: disabled)
				--platform <name>   : Specify the target host platform.

	--- HYPERVISOR COMMANDS ---

		monorepo hypervisor list

			list the supported hypervisors.

		monorepo hypervisor show <name>

			print the hypervisor configuration.

		monorepo hypervisor <enable|disable> <name>

			Enable or disable the build hypervisor configuration.

		monorepo hypervisor delete <name>

			delete the current hypervisor configuration.

		monorepo hypervisor create <name> [CRI Options]

				CRI Options:
                ------------------------------------------------------------
				--enabled         : Enable the config. (Default: disabled)
				--platform <name> : Specify the target host platform.

	--- LANGUAGE COMMANDS ---

		monorepo language list

			list the supported languages.

		monorepo language show <name>

			print the language configuration.

		monorepo language <enable|disable> <name>

			Enable or disable the build language configuration.

		monorepo language delete <name>

			delete the current hypervisor configuration.

		monorepo language create <name> [Language Options]

				Language Options:
                ------------------------------------------------------------
				--enabled          : Enable the config. (Default: disabled)
				--platform <name>  : Specify the target host platform.
				--directory <path> : The path where the source code 
									 will exist.

	--- PLATFORM COMMANDS ---

		monorepo platform list

			list the supported platforms.

		monorepo platform show <name>

			print the platform configuration.

		monorepo platform delete <name>

			delete the current platform configuration.

		monorepo platform <enable|disable> <name>

			Enable or disable the build platform configuration.

		monorepo platform create <name> [CRI Options]

				Platform Options:
                ------------------------------------------------------------
				--description           : describe the platform
				--enabled               : Enable config. (Default: disabled)
				--connect <local|ssh>   : Connection method (local or ssh)
				--opsys_family <regex>  : Regex for a supported opsys family. 
				--opsys_version <regex> : Regex for a supported opsys version.
				--cpu_arch <regex>      : A regex for a cpu architecture.

`

func main() {
	exit.IfHelpRequested(CommandUsage)
	exit.IfVersionRequested()
}
