package application

const (
	commandUsage = `
%s

Syntax:
    %s [mode]

modes:
    Run as client mode:

        client --host <hostname> --port <portNumber> --apiKey <base64 key string>

    Run as server mode:

    server --host <hostname> --port <portNumber> --apiKey <base64 key string>

Options:
    -h | --help displays this message
`
)
