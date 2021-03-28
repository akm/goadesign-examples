// Code generated by goa v3.3.1, DO NOT EDIT.
//
// session HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/cookies/design -o
// $(GOPATH)/src/goa.design/examples/cookies

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	sessionc "goa.design/examples/cookies/gen/http/session/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `session (create-session|use-session)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` session create-session --body '{
      "name": "Ut quas sit at illum tempore ad."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		sessionFlags = flag.NewFlagSet("session", flag.ContinueOnError)

		sessionCreateSessionFlags    = flag.NewFlagSet("create-session", flag.ExitOnError)
		sessionCreateSessionBodyFlag = sessionCreateSessionFlags.String("body", "REQUIRED", "")

		sessionUseSessionFlags         = flag.NewFlagSet("use-session", flag.ExitOnError)
		sessionUseSessionSessionIDFlag = sessionUseSessionFlags.String("session-id", "REQUIRED", "")
	)
	sessionFlags.Usage = sessionUsage
	sessionCreateSessionFlags.Usage = sessionCreateSessionUsage
	sessionUseSessionFlags.Usage = sessionUseSessionUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "session":
			svcf = sessionFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "session":
			switch epn {
			case "create-session":
				epf = sessionCreateSessionFlags

			case "use-session":
				epf = sessionUseSessionFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "session":
			c := sessionc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create-session":
				endpoint = c.CreateSession()
				data, err = sessionc.BuildCreateSessionPayload(*sessionCreateSessionBodyFlag)
			case "use-session":
				endpoint = c.UseSession()
				data, err = sessionc.BuildUseSessionPayload(*sessionUseSessionSessionIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// sessionUsage displays the usage of the session command and its subcommands.
func sessionUsage() {
	fmt.Fprintf(os.Stderr, `The session service illustrates managing user sessions with cookies.
Usage:
    %s [globalflags] session COMMAND [flags]

COMMAND:
    create-session: CreateSession implements create_session.
    use-session: UseSession implements use_session.

Additional help:
    %s session COMMAND --help
`, os.Args[0], os.Args[0])
}
func sessionCreateSessionUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] session create-session -body JSON

CreateSession implements create_session.
    -body JSON: 

Example:
    `+os.Args[0]+` session create-session --body '{
      "name": "Ut quas sit at illum tempore ad."
   }'
`, os.Args[0])
}

func sessionUseSessionUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] session use-session -session-id STRING

UseSession implements use_session.
    -session-id STRING: 

Example:
    `+os.Args[0]+` session use-session --session-id "Qui voluptatem dolores."
`, os.Args[0])
}
