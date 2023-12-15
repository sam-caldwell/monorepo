#!/usr/bin/env python3
from argparse import ArgumentParser
from re import compile
from subprocess import run

EXIT_SUCCESS = 0

EXIT_ERROR_LOCAL_GIT_COMMAND = 1
EXIT_ERROR_SSH_GIT_COMMAND = 2
EXIT_ERROR_GET_SERVER_BAD_INPUT = 3
EXIT_ERROR_GET_SERVER_EXCEPTION = 4
EXIT_ERROR_SET_SERVER_INVALID = 5
EXIT_ERROR_SET_SERVER_EXCEPTION = 6
EXIT_ERROR_AUTH_REPO_EXCEPTION = 7
EXIT_ERROR_CREATE_REPO_INVALID = 8
EXIT_ERROR_CREATE_REPO_EXCEPTION = 9
EXIT_ERROR_DELETE_REPO_INVALID = 10
EXIT_ERROR_DELETE_REPO_EXCEPTION = 11
EXIT_ERROR_LIST_REPOS_EXCEPTION = 12
EXIT_ERROR_PROXY_REPO_EXCEPTION = 13
EXIT_ERROR_SSH_INVALID_KEY = 14

EXIT_UNDEFINED_ERROR = 253
EXIT_UNSPECIFIED_ERROR = 254
EXIT_INTERNAL_ERROR_INVALID_CMD = 255

REQUIRE_NO_PARAMETERS = {}
PREFERRED_SERVER_KEY = "core.preferredGitserver"

CMD_AUTHORIZE = "authorize"
CMD_AUTHORIZED = "authorized"
CMD_CREATE = "create"
CMD_DELETE = "delete"
CMD_LIST = "list"
CMD_PROXY = "proxy"
CMD_RENAME = "rename"
CMD_USE = "use"

help_text = """
Usage:
    git authorize <ssh_key> [--debug]
    git authorized [--debug]
    git create <repo> [--debug]
    git delete <repo> [--debug]
    git list [--debug]
    git proxy <git ssh repo url> [--debug]
    git rename <old_repo_name> <new_repo_name> [--debug]
    git use <server> [--debug]
"""


class GitServer(object):
    """
        Class for Interacting with Git (local and remote)
    """
    """
        __valid_repo_name_regex:
            A regular expression used to evaluate the validity
            of a repository name used with our git-sqldbtest and the
            git-server
    """
    __valid_repo_name_regex = "^[a-zA-Z][a-zA-Z.\/-_/0-9]+[a-zA-Z0-9]$"
    """
        __valid_sshkey_regex:
            A regular expression used to evaluate the validity of
            an ssh key used with ssh for git-server and git-sqldbtest.
    """
    __valid_sshkey_regex = "^(" + \
                           "(" + \
                           "ssh-rsa|" + \
                           "ssh-dss|" + \
                           "ecdsa-sha2-nistp256|" + \
                           "ecdsa-sha2-nistp384|" + \
                           "ecdsa-sha2-nistp521|" + \
                           "ssh-ed25519" + \
                           ") " + \
                           "AAAA[0-9A-Za-z+/]+[=]{0,3}" + \
                           "( " + \
                           "[^@]+@[^@]+" + \
                           ")?" + \
                           ")$"
    """
        __arg_error_codes:
            A set of error codes for invalid arguments
            encountered when running GitServer.parameter_check()
    """
    __arg_error_codes = {
        "prohibited_repo": 100,
        "prohibited_source": 101,
        "prohibited_destination": 102,
        "prohibited_server": 103,
        "required_sshkey": 104,
        "required_repo": 105,
        "required_source": 106,
        "required_destination": 107,
        "required_server": 107
    }
    """
        __disallowed_git_servers:
            A list of git servers we will not allow to be used
            as preferred (local) git servers for use with the
            git-sqldbtest.
    """
    __disallowed_git_servers = [
        'github.com',
        'bitbucket.com',
        'gitlabs.com'
        'visualstudio.com',
        'vsts.com',
    ]

    def __init__(self) -> None:
        """
            class constructor.
            Processes the commandline arguments.

            :return: None
        """
        parser = ArgumentParser(description="Git Tools CommandLine")
        parser.add_argument(
            "--command",
            type=str,
            required=True,
            choices=[
                CMD_AUTHORIZE,
                CMD_AUTHORIZED,
                CMD_CREATE,
                CMD_DELETE,
                CMD_LIST,
                CMD_PROXY,
                CMD_RENAME,
                CMD_USE
            ],
            help="Git Tools Command")

        parser.add_argument(
            "--repo",
            type=str,
            required=False,
            default="",
            help="specify a repository name")

        parser.add_argument(
            "--sshkey",
            type=str,
            required=False,
            default="",
            help="specify an ssh public key")

        parser.add_argument(
            "--source",
            type=str,
            required=False,
            default="",
            help="specify a source repository name"
        )

        parser.add_argument(
            "--destination",
            type=str,
            required=False,
            default="",
            help="specify a destination repository name"
        )

        parser.add_argument(
            "--server",
            type=str,
            required=False,
            default="",
            help="specify a git server")

        parser.add_argument(
            "--global",
            dest="scope",
            required=False,
            action="store_true",
            default=False,
            help="search global or local .gitconfig")

        parser.add_argument(
            "--debug",
            required=False,
            default=False,
            action="store_true",
            help="enable debug messages"
        )

        self.args = parser.parse_args()
        self.debug("Commandline arguments processed.")

    def debug(self, msg: str) -> None:
        """
            print debug messages

            :param msg: str
            :return: None
        """
        if self.args.debug:
            print(f"[DEBUG]: {msg}")

    @staticmethod
    def __valid_server_name(server_name: str) -> bool:
        """
            return whether the server_name is valid
            and not disallowed.

            :param server_name: str
            :return: bool
        """
        # ToDo: validate name against a regex
        for disallowed_server in GitServer.__disallowed_git_servers:
            if disallowed_server in server_name:
                return False
        return True

    def __valid_repo_name(self, repo: str) -> bool:
        """
            Determine if the input string is a valid repository name

            :param repo: str
            :return: bool
        """
        self.debug(f"Validating repo '{repo}'")
        pattern = compile(GitServer.__valid_repo_name_regex)
        self.debug(f"Valid? {pattern.match(repo)}")
        return pattern.match(repo) is not None

    def __valid_sshkey_name(self, key: str) -> bool:
        """
            Determine if the ssh key is valid

            :param key: str
            :return: bool
        """
        self.debug(f"Validating ssh key '{key}'")
        pattern = compile(GitServer.__valid_sshkey_regex)
        self.debug(f"Valid? {pattern.match(key)}")
        return pattern.match(key) is not None

    @staticmethod
    def __global_flag(global_scope: bool) -> str:
        """
            return the --global flag if s is true.

            :param global_scope: bool
            :return: str
        """
        if global_scope:
            return "--global"
        else:
            return ""

    def runner(self, command: str) -> (int, str):
        """
            Execute a command and return the captured output.

            :param command: str
            :return: int(exit_code), str(stdout)
        """
        try:
            self.debug(f"command(runner): {command}")
            result = run(command,
                         shell=True,
                         check=False,
                         capture_output=True)
            return result.returncode, result.stdout.decode().strip()
        except Exception as e:
            return EXIT_ERROR_LOCAL_GIT_COMMAND, f"{e}"

    def ssh_runner(self,
                   server: str,
                   command: str) -> (int, str):
        """
            Execute an ssh command against the remote git server.

            :param server: str
            :param command: str
            :return: int(exit_code), str(stdout)
        """
        try:
            cmd = "ssh -o 'StrictHostKeyChecking no' " + \
                  f"git@{server} {command}"
            self.debug(f"command(ssh_runner): {cmd}")
            return self.runner(cmd)
        except Exception as e:
            return EXIT_ERROR_SSH_GIT_COMMAND, f"{e}"

    def __get_server(self, this_scope: bool = False) -> (int, str):
        """
            return the preferred git server

            :param this_scope: bool
            :return: int(exit_code), str(stdout)
        """
        try:
            cmd = "git config " + \
                  f"{self.__global_flag(this_scope).strip()} " + \
                  f"--get {PREFERRED_SERVER_KEY}"
            exit_code, git_server = self.runner(cmd)
            self.debug(f"git_server[{exit_code}]:'{git_server}'")
            if (exit_code != 0) or (git_server == ""):
                return EXIT_ERROR_GET_SERVER_BAD_INPUT, \
                    f"{PREFERRED_SERVER_KEY} not defined." \
                    f"Use 'git {CMD_USE} <private_server>' first."
            else:
                return EXIT_SUCCESS, git_server
        except Exception as e:
            if this_scope:
                return EXIT_ERROR_GET_SERVER_EXCEPTION, \
                    f"(getting preferred server) {e}"
            else:
                # Retry with global scope because we were not
                # in a git repository we could configure locally.
                return self.__get_server(this_scope=True)

    def __set_server(self, server_name: str,
                     this_scope: bool = False) -> (int, str):
        """
            set the preferred git server

            :param server_name: str
            :param this_scope: bool
            :return: int(exit_code), str(stdout)
        """
        try:
            self.debug(f"setting preferred server:'{server_name}'")
            if self.__valid_server_name(server_name):
                cmd = "git config " + \
                      f"{self.__global_flag(this_scope)} " + \
                      f"{PREFERRED_SERVER_KEY} " + \
                      f"{server_name}"
                exit_code, stdout = self.runner(cmd)
                if exit_code != 0:
                    if this_scope:
                        return exit_code, stdout
                    else:
                        return self.__set_server(server_name=server_name,
                                                 this_scope=True)
                else:
                    return exit_code, stdout

            else:
                return EXIT_ERROR_SET_SERVER_INVALID, \
                    f"{server_name} is invalid or is one of several " \
                    "disallowed git servers that cannot be used with Git Tools"
        except Exception as e:
            if this_scope:
                return EXIT_ERROR_SET_SERVER_EXCEPTION, \
                    f"(setting preferred server) {e}"
            else:
                self.debug("escalating scope")
                return self.__set_server(server_name=server_name,
                                         this_scope=True)

    def authorize(self, ssh_key: str,
                  search_scope: bool = False) -> (int, str):
        """
            Add a new authorized ssh key to the preferred git server.

            :param ssh_key: str
            :param search_scope: bool (default: False
            :return: int (exit_code), str (list of repos)
        """
        server = ""
        try:
            self.debug("authorize() starting...")
            exit_code, stdout = self.__get_server(search_scope)
            if exit_code != 0:
                self.debug(f"could not find preferred server"
                           f"[{exit_code}]: '{stdout}'")
                return exit_code, stdout
            server = stdout
            cmd = f"authorize --sshkey '{ssh_key}'"
            exit_code, stdout = self.ssh_runner(server=server,
                                                command=cmd)
            if exit_code == 255:
                return 255, "connection failed (unauthorized)"
            else:
                return exit_code, stdout

        except Exception as e:
            return EXIT_ERROR_AUTH_REPO_EXCEPTION, \
                f"could not create repository ({ssh_key}) " \
                f"on '{server}'. {e}"

    def authorized(self, search_scope: bool = False) -> (int, str):
        """
            Return an enumerated list of authorized keys for the
            preferred server.

            :param search_scope: bool (default: False
            :return: int (exit_code), str (list of repos)
        """
        exit_code, stdout = self.__get_server(search_scope)
        if exit_code != 0:
            return exit_code, f"(authorized): {stdout}"
        server = stdout
        exit_code, stdout = self.ssh_runner(server=server,
                                            command=CMD_AUTHORIZED)
        if exit_code == 1:
            return 1, f"Please use 'git authorize --sshkey <key>' " \
                      f"first. {stdout}"
        elif exit_code == 255:
            return 255, "connection failed (unauthorized)"
        else:
            return exit_code, stdout

    def create_repository(self, repo: str,
                          search_scope: bool = False) -> (int, str):
        """
            Create a new repository on the preferred server.

            :param repo: str
            :param search_scope: bool (default: false)
            :return: int (exit_code), str (list of repos)
        """
        server = ""
        try:
            self.debug(f"Get Preferred server to create '{repo}'")
            exit_code, stdout = self.__get_server(search_scope)
            if exit_code == 0:
                self.debug(f"Preferred server found for create '{repo}'")
            else:
                self.debug(f"could not find preferred server"
                           f"[{exit_code}]: '{stdout}'")
                return exit_code, stdout
            server = stdout
            self.debug(f"Validate repo name ({repo}) "
                       f"for create on '{server}'")
            if self.__valid_repo_name(repo):
                self.debug(f"repo name is valid: '{repo}'")
                cmd = f"create --repo {repo}"
                return self.ssh_runner(server=server, command=cmd)
            else:
                self.debug(f"repo name is not valid: '{repo}'")
                return EXIT_ERROR_CREATE_REPO_INVALID, f"{repo} is not valid"
        except Exception as e:
            return EXIT_ERROR_CREATE_REPO_EXCEPTION, \
                f"could not create repository ({repo}) on '{server}'. {e}"

    def delete_repository(self, repo: str,
                          search_scope: bool = False) -> (int, str):
        """
            Delete a new repository on the preferred server.

            :param repo: str
            :param search_scope: bool (default: false)
            :return: int (exit_code), str (list of repos)
        """
        server = ""
        try:
            self.debug(f"Get Preferred server to delete '{repo}'")
            exit_code, stdout = self.__get_server(search_scope)
            if exit_code == 0:
                self.debug(f"Preferred server found for delete '{repo}'")
            else:
                self.debug(f"could not find preferred server"
                           f"[{exit_code}]: '{stdout}'")
                return exit_code, stdout
            server = stdout
            self.debug(f"Validate repo name ({repo}) "
                       f"for delete on '{server}'")
            if self.__valid_repo_name(repo):
                self.debug(f"repo name is valid: '{repo}'")
                cmd = f"delete {repo}"
                return self.ssh_runner(server=server, command=cmd)
            else:
                self.debug(f"repo name is not valid: '{repo}'")
                return EXIT_ERROR_DELETE_REPO_INVALID, f"{repo} is not valid"
        except Exception as e:
            return EXIT_ERROR_DELETE_REPO_EXCEPTION, \
                f"could not delete repository ({repo}) on '{server}'. {e}"

    def list_repositories(self, search_scope: bool = False) -> (int, str):
        """
            List the repositories in the preferred server (if set)

            :param search_scope: bool (default: false)
            :return: int (exit_code), str (list of repos)
        """
        try:
            exit_code, stdout = self.__get_server(search_scope)
            if exit_code != 0:
                return exit_code, f"(list): {stdout}"
            server = stdout
            exit_code, stdout = self.ssh_runner(server=server,
                                                command=CMD_LIST)
            if exit_code == 0:
                repo_list = ""
                header = "repositories on {server}"
                base_path = "/git/repos/"
                name_width = 0
                path_width = 0
                repos = {}
                for name in stdout.split('\n'):
                    path = f"{base_path}{name}"
                    repos[name] = {
                        "name": name + " " * (name_width - len(name)),
                        "path": path + " " * (path_width - len(path)),
                    }
                    if len(name) > name_width:
                        name_width = len(name)
                    if len(repos[name]["path"]) > path_width:
                        path_width = len(repos[name]["path"])
                width = len(header)
                repo_list = ""

                for name, value in repos.items():
                    line = f"| {value['name']} | {value['path']} |"
                    if len(line) > width:
                        width = len(line)
                    repo_list += line + "\n"

                self.debug(f"width: {width}")
                header = "|" + header + " " * (width - len(header) - 2) + "|\n"
                separator = "+" + "-" * (width - 2) + "+\n"
                return exit_code, f"{separator}" \
                                  f"{header}" \
                                  f"{separator}" \
                                  f"{repo_list}" \
                                  f"{separator}"
            else:
                return exit_code, stdout
        except Exception as e:
            return EXIT_ERROR_LIST_REPOS_EXCEPTION, \
                f"Error: could not list repositories. {e}"

    def proxy(self, repo: str, search_scope: bool = False) -> (int, str):
        """
            Clone a repository from a third-party server to the git-server
            as a proxy, which will allow for a chained interaction from the
            user/client to the git-server then up to the third-party remote
            server.

            :param repo:
            :param search_scope:
            :return: int (exit_code), str (list of repos)
        """
        server = ""
        try:
            self.debug(f"Get Preferred server to proxy '{repo}'")
            exit_code, stdout = self.__get_server(search_scope)
            if exit_code == 0:
                self.debug(f"Preferred server found for proxy '{repo}'")
            else:
                self.debug(f"could not find preferred server"
                           f"[{exit_code}]: '{stdout}'")
                return exit_code, stdout
            server = stdout
            cmd = f"proxy {repo}"
            return self.ssh_runner(server=server, command=cmd)
        except Exception as e:
            return EXIT_ERROR_PROXY_REPO_EXCEPTION, \
                f"could not proxy repository ({repo}) on '{server}'. {e}"

    def rename(self, source_repo: str, destination_repo: str,
               search_scope: bool = False) -> (int, str):
        """
            Move/rename a source_repo to a destination repo within
            a given search_scope.

            :param source_repo: str
            :param destination_repo: str
            :param search_scope: bool
            :return: int (exit_code), str (list of repos)
        """
        try:
            exit_code, stdout = self.__get_server(search_scope)
            if exit_code != 0:
                return exit_code, stdout
            server = stdout
            cmd = f"rename {source_repo} {destination_repo}"
            return self.ssh_runner(server=server, command=cmd)
        except Exception as e:
            return EXIT_ERROR_LIST_REPOS_EXCEPTION, \
                f"Error: could not move/rename repository. " \
                f"{source_repo} to {destination_repo} " \
                f"error: {e}"

    def use(self, server_name: str,
            this_scope: bool = False) -> (int, str):
        """
            Configure the current preferred git server.

            :param server_name: str
            :param this_scope: bool (default False
            :return: int(exit_code), str(stdout)
        """
        exit_code, stdout = self.__set_server(server_name=server_name,
                                              this_scope=this_scope)
        if exit_code == 0:
            return exit_code, \
                f"Set preferred server ('{server_name}'): ok"
        else:
            return exit_code, \
                f"Set preferred server ('{server_name}'): failed"

    @staticmethod
    def show_usage(error: str,
                   ret_code: int = EXIT_UNDEFINED_ERROR) -> int:
        """
            Show the program usage on error.

            :param error: str
            :param ret_code: int (default: EXIT_UNDEFINED_ERROR)
            :return: int
        """
        if error.strip() == "":
            error = "unspecified error"
        print(f"Error: '{error}'\n\n{help_text}\n")
        return ret_code

    def prohibited(self,
                   param_name: str,
                   param_value: str,
                   ret_code: int = EXIT_SUCCESS) -> int:
        """
            return exit_code and message if prohibited
            pattern is not empty

            :param param_name: str
            :param param_value: str
            :param ret_code: int (default: EXIT_SUCCESS)
            :return: int (exit_code), str (list of repos)
        """
        if param_value == "":
            return 0
        else:
            return self.show_usage(f"{param_name} should not be provided "
                                   f"for {self.args.command}", ret_code)

    def required(self,
                 param_name: str,
                 param_value: str,
                 ret_code: int = EXIT_SUCCESS) -> int:
        """
            return exit_code and message if required
            parameter is missing

            :param param_name: str
            :param param_value: str
            :param ret_code: int (default: EXIT_SUCCESS)
            :return: int (exit_code), str (list of repos)
        """
        if param_value == "":
            return self.show_usage(f"{param_name} must be provided "
                                   f"for {self.args.command}", ret_code)
        else:
            return 0

    def parameter_check(self, required: dict,
                        prohibited: dict) -> int:
        """
            check the set of required parameters

            :param required: dict - set of required parameters.
            :param prohibited: dict - set of prohibited parameters.
            :return: int (exit_code)
        """
        check_config = {
            "prohibited": {"fn": self.prohibited, "data": prohibited},
            "required": {"fn": self.required, "data": required}
        }
        for prefix, check in check_config.items():
            check_function = check.get("fn")
            for p, v in check.get("data", {}).items():
                error_vector = f"{prefix}_{p}"
                code = check_function(p,
                                      v,
                                      GitServer.__arg_error_codes.get(
                                          error_vector,
                                          EXIT_UNSPECIFIED_ERROR))
                if code == EXIT_SUCCESS:
                    self.debug(f"\t\t{p:30}: ok")
                else:
                    self.debug(f"\t\t{p:30}: fail ({code})")
                    return code
        self.debug("parameter_check() complete")
        return 0

    def cmd_authorize(self) -> int:
        """
            git authorize <ssh_public_key>
                authorize a new ssh public key for the
                current preferred git server.
        """
        exit_code = self.parameter_check(
            required={
                "sshkey": self.args.sshkey.strip()
            },
            prohibited={
                "repo": self.args.repo.strip(),
                "source": self.args.source.strip(),
                "destination": self.args.destination.strip(),
                "server": self.args.server.strip()
            })
        if exit_code != EXIT_SUCCESS:
            return exit_code
        if self.__valid_sshkey_name(self.args.sshkey):
            self.debug("Key is valid")
        else:
            print("Invalid SSH Key")
            return EXIT_ERROR_SSH_INVALID_KEY
        exit_code, stdout = self.authorize(ssh_key=self.args.sshkey.strip(),
                                           search_scope=self.args.scope)
        if exit_code != 0:
            return self.show_usage(stdout, exit_code)
        else:
            print(f"{stdout}\n")
            return exit_code

    def cmd_authorized(self) -> int:
        """
            git authorized
                list the authorized ssh keys on the preferred server
        """
        self.debug("cmd_authorized() starting...")
        exit_code = self.parameter_check(
            required=REQUIRE_NO_PARAMETERS,
            prohibited={
                "repo": self.args.repo.strip(),
                "source": self.args.source.strip(),
                "destination": self.args.destination.strip(),
                "server": self.args.server.strip(),
                "sshkey": self.args.sshkey.strip()
            })
        if exit_code != EXIT_SUCCESS:
            return exit_code
        exit_code, stdout = self.authorized(self.args.scope)
        if exit_code != 0:
            return self.show_usage(stdout, exit_code)
        else:
            print(f"{stdout}\n")
            return exit_code

    def cmd_create(self) -> int:
        """
            get create <repo>
                -- create or delete a repository on the preferred git server,
                   where scope is defined or where the preferred git server
                   lists repositories on a locally defined server and if not
                   defined it will list the repositories on the globally
                   defined preferred server.
        """
        exit_code = self.parameter_check(
            required={
                "repo": self.args.repo.strip(),
            },
            prohibited={
                "server": self.args.server.strip(),
                "source": self.args.source.strip(),
                "destination": self.args.destination.strip(),
                "sshkey": self.args.sshkey.strip()
            })
        if exit_code != EXIT_SUCCESS:
            return exit_code

        exit_code, stdout = self.create_repository(
            repo=self.args.repo,
            search_scope=self.args.scope)
        if exit_code != 0:
            return self.show_usage(stdout, exit_code)
        else:
            print(stdout)
            return exit_code

    def cmd_delete(self) -> int:
        """
             get delete <repo>
                 -- create or delete a repository on the preferred git server,
                    where scope is defined or where the preferred git server
                    lists repositories on a locally defined server and if not
                    defined it will list the repositories on the globally
                    defined preferred server.
         """
        exit_code = self.parameter_check(
            required={
                "repo": self.args.repo.strip(),
            },
            prohibited={
                "server": self.args.server.strip(),
                "source": self.args.source.strip(),
                "destination": self.args.destination.strip(),
                "sshkey": self.args.sshkey.strip()
            })
        if exit_code != EXIT_SUCCESS:
            return exit_code

        exit_code, stdout = self.delete_repository(
            repo=self.args.repo,
            search_scope=self.args.scope)
        if exit_code != 0:
            return self.show_usage(stdout, exit_code)
        else:
            print(stdout)
            return exit_code

    def cmd_list(self) -> int:
        """
            git list
                lists the repositories on the preferred git server, where
                scope is defined or where the preferred git server lists
                repositories on a locally defined server and if not
                defined it will list the repositories on the globally
                defined preferred server.
        """
        self.debug("cmd_list() starting...")
        exit_code = self.parameter_check(
            required=REQUIRE_NO_PARAMETERS,
            prohibited={
                "repo": self.args.repo.strip(),
                "server": self.args.server.strip(),
                "source": self.args.source.strip(),
                "destination": self.args.destination.strip(),
                "sshkey": self.args.sshkey.strip()
            })
        self.debug(f"cmd_list() input validation: {exit_code}")
        if exit_code != EXIT_SUCCESS:
            return exit_code

        exit_code, stdout = self.list_repositories(self.args.scope)
        self.debug(f"cmd_list() list_repositories() has returned {exit_code}")
        if exit_code != 0:
            return self.show_usage(stdout, exit_code)
        else:
            print(stdout)
            return exit_code

    def cmd_proxy(self) -> int:
        """
            Clone a given repository via ssh into the preferred
            server
        """
        exit_code = self.parameter_check(
            required={
                "repo": self.args.repo.strip(),
            },
            prohibited={
                "server": self.args.server.strip(),
                "source": self.args.source.strip(),
                "destination": self.args.destination.strip(),
                "sshkey": self.args.sshkey.strip()
            })
        if exit_code != EXIT_SUCCESS:
            return exit_code
        #
        # ToDo: Call the authorize command.
        #

    def cmd_rename(self) -> int:
        """
            get rename <source_repo> <destination_repo>
                -- rename/move a source_repo to the destination_repo.
                -- create or delete a repository on the preferred git server,
                   where scope is defined or where the preferred git server
                   lists repositories on a locally defined server and if not
                   defined it will list the repositories on the globally
                   defined preferred server.
        """
        exit_code = self.parameter_check(
            required={
                "destination": self.args.destination.strip(),
                "source": self.args.source.strip()
            },
            prohibited={
                "repo": self.args.repo.strip(),
                "server": self.args.server.strip(),
                "sshkey": self.args.sshkey.strip()
            })
        if exit_code != EXIT_SUCCESS:
            return exit_code

        exit_code, stdout = self.rename(
            source_repo=self.args.source.strip(),
            destination_repo=self.args.destination.strip(),
            search_scope=self.args.scope)

        if exit_code != 0:
            return self.show_usage(stdout, exit_code)
        return exit_code

    def cmd_use(self) -> int:
        """
            git use <server>
                -- define the preferred git server for the specified scope
                -- where scope is not defined, local scope will be used if
                   the current directory is a git repo.
                -- if no scope is defined in the current directory, scope
                   will fall back to global scope.
        """
        exit_code = self.parameter_check(
            required={
                "server": self.args.server.strip(),
            },
            prohibited={
                "repo": self.args.repo.strip(),
                "destination": self.args.destination.strip(),
                "source": self.args.source.strip(),
                "sshkey": self.args.sshkey.strip()
            })
        if exit_code != EXIT_SUCCESS:
            return exit_code

        exit_code, stdout = self.use(server_name=self.args.server,
                                     this_scope=self.args.scope)

        if exit_code == 0:
            print(stdout)
            return exit_code
        else:
            return self.show_usage(stdout, exit_code)

    def execute(self) -> int:
        """
            execute the git commands defined in command-line arguments.
            :return: int (exit code)
        """
        self.debug(f"execute() is starting... "
                   f"(command: {self.args.command})")
        vector_table = {
            CMD_AUTHORIZE: self.cmd_authorize,
            CMD_AUTHORIZED: self.cmd_authorized,
            CMD_CREATE: self.cmd_create,
            CMD_DELETE: self.cmd_delete,
            CMD_LIST: self.cmd_list,
            CMD_PROXY: self.cmd_proxy,
            CMD_RENAME: self.cmd_rename,
            CMD_USE: self.cmd_use
        }
        self.debug(f"is command in vector_table? "
                   f"{self.args.command in vector_table}")
        command = vector_table.get(self.args.command, None)
        self.debug("command is setup")
        if command is None:
            return self.show_usage("Internal programming error "
                                   f"(command: {self.args.command})",
                                   EXIT_INTERNAL_ERROR_INVALID_CMD)
        else:
            return command()


if __name__ == "__main__":
    exit(GitServer().execute())
