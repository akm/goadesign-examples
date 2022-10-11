// Code generated by goa v3.10.0, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/error/design -o error

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	calcc "goa.design/examples/error/gen/http/calc/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `calc divide
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` calc divide --body '{
      "dividend": 6322633713974661021,
      "divisor": 3793862871819669726
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
		calcFlags = flag.NewFlagSet("calc", flag.ContinueOnError)

		calcDivideFlags    = flag.NewFlagSet("divide", flag.ExitOnError)
		calcDivideBodyFlag = calcDivideFlags.String("body", "REQUIRED", "")
	)
	calcFlags.Usage = calcUsage
	calcDivideFlags.Usage = calcDivideUsage

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
		case "calc":
			svcf = calcFlags
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
		case "calc":
			switch epn {
			case "divide":
				epf = calcDivideFlags

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
		case "calc":
			c := calcc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "divide":
				endpoint = c.Divide()
				data, err = calcc.BuildDividePayload(*calcDivideBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// calcUsage displays the usage of the calc command and its subcommands.
func calcUsage() {
	fmt.Fprintf(os.Stderr, `Service is the calc service interface.
Usage:
    %[1]s [globalflags] calc COMMAND [flags]

COMMAND:
    divide: Divide implements divide.

Additional help:
    %[1]s calc COMMAND --help
`, os.Args[0])
}
func calcDivideUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc divide -body JSON

Divide implements divide.
    -body JSON: 

Example:
    %[1]s calc divide --body '{
      "dividend": 6322633713974661021,
      "divisor": 3793862871819669726
   }'
`, os.Args[0])
}
