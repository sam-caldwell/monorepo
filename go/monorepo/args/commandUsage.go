package args

const CommandUsage = `

	monorepo -h | --help    : show usage (this message)
	monorepo -v | --version : show current version

	--- GENERAL COMMAND FORMAT ---

		monorepo [General Options] [commandGroup] [command] [parameters]
	
			General Options:
			-------------------------------------------------
			--debug : print verbose debug messages
			--color : color output with ANSI escape codes
			--noop  : perform command in 'no operation' mode 
					  for planning/testing
	
	--- GENERAL TASK SHELL COMMAND ---

		monorepo exec <string> [options]

			This command will execute the given string using the provided
			options and return the exit code from the command.  Output will
			be written to stdout.

			Shell Options:
			-------------------------------------------------
			--working-dir  : Specify the path from which the
							 command will be run.

			--shell <name> : Specify the shell (bash, zsh,
							 PowerShell) to use to execute.
	
	--- CONFIGURATION COMMANDS ---

		monorepo config <object> <command> [parameters]

			Configure the monorepo and its tooling.

				Objects:
				-------------------------------------------------------
				cri			: supported container runtime
				host		: build host configurations
				hypervisor 	: supported hypervisors
				language 	: supported programming languages
				platform	: supported platforms (for build hosts)
				projects	: source code projects

				Commands:
				-------------------------------------------------------
				create <name>	: Create the specified object
				check <name>	: Verify the specified object
				delete <name>	: Delete the specified object
				disable <name>	: Mark the object as disabled
				enable <name>	: Mark the object as enabled
				list			: List all objects
				show <name>		: Show the specified object configuration.


				Options:
				-------------------------------------------------------
				General Options
					--enabled 			: Set the object as enabled.

					--platform <name> 	: Add a supported platform 
										  for the object.

					--description <desc>: Describe the object.

				Host-Create Options
					--connect <method>  : define how to connect to
										  the named build host
						methods:
							local : local command execution
							ssh   : key-based ssh to run commands

					--host 				: IP address or FQDN and port
										  with which to connect to the
 										  build host.

					--username <user>	: The username used to connect
										  to the build host.  If not
										  specified, current user is
										  used.

					--hypervisor <name>	: Specifies a hypervisor to
										  expect on the build host.

					--cri <name>		: Specifies a CRI to
										  expect on the build host.

				Language-Create Options
					--directory <path> 	: The path where source files
										  will be stored.

				Platform-Create Options
					--opsys_family <regex>  : regex to restrict platform
 											  to a given operating system
											  family.
					--opsys_version <regex> : regex to restrict platform
											  to a given operating system
											  version.
					--cpu_arch <regex>		: regex to restrict platform
											  to a given CPU.

				Project-Create Options
					--author <string> 		 : The name of the author.
					--maintainer <string> 	 : The name of the maintainer.
					--output <artifact type> : The type of artifact a
											   project produces.
						artifact_types:
							executable : a binary executable
							source     : reusable source code
					--

	--- PROJECT-ACTION COMMANDS ---

		monorepo <action> <all | project {name}>

			Execute a project action (e.g. build, lint, test) against 
			one or more projects.

		Actions
		----------------------------------------------------------------
		build : Build one or more projects into their respective artifacts.
		lint  : Lint sources for one or more projects
		test  : Run tests for one or more projects

		


`
