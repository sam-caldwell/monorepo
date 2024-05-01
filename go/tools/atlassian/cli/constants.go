package cli

const (
    JiraClientUsage = `
jira create {ticket,project} <name> --descriptor <filename.yaml> [--api-key <string>]

jira read {ticket, project} <name> [--api-key <string>]

jira update {ticket,project  --descriptor <filename.yaml>

jira delete {ticket, project} <name> [--api-key <string>]

jira list ticket <jql> [--api-key <string>]

jira list projects --offset <int> --count <int> [--api-key <string>]
`

    HelpText = "Use -h | --help for usage"
)
