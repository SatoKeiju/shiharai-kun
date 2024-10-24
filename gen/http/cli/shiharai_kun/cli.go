// Code generated by goa v3.19.1, DO NOT EDIT.
//
// shiharai-kun HTTP client CLI support package
//
// Command:
// $ goa gen github.com/SatoKeiju/shiharai-kun/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	invoicesc "github.com/SatoKeiju/shiharai-kun/gen/http/invoices/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `invoices fetch
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` invoices fetch --user-id "Qui itaque molestiae doloribus saepe expedita dolorum."` + "\n" +
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
) (goa.Endpoint, any, error) {
	var (
		invoicesFlags = flag.NewFlagSet("invoices", flag.ContinueOnError)

		invoicesFetchFlags      = flag.NewFlagSet("fetch", flag.ExitOnError)
		invoicesFetchUserIDFlag = invoicesFetchFlags.String("user-id", "REQUIRED", "")
	)
	invoicesFlags.Usage = invoicesUsage
	invoicesFetchFlags.Usage = invoicesFetchUsage

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
		case "invoices":
			svcf = invoicesFlags
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
		case "invoices":
			switch epn {
			case "fetch":
				epf = invoicesFetchFlags

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
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "invoices":
			c := invoicesc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "fetch":
				endpoint = c.Fetch()
				data, err = invoicesc.BuildFetchPayload(*invoicesFetchUserIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// invoicesUsage displays the usage of the invoices command and its subcommands.
func invoicesUsage() {
	fmt.Fprintf(os.Stderr, `invoice service
Usage:
    %[1]s [globalflags] invoices COMMAND [flags]

COMMAND:
    fetch: Fetch implements fetch.

Additional help:
    %[1]s invoices COMMAND --help
`, os.Args[0])
}
func invoicesFetchUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] invoices fetch -user-id STRING

Fetch implements fetch.
    -user-id STRING: 

Example:
    %[1]s invoices fetch --user-id "Qui itaque molestiae doloribus saepe expedita dolorum."
`, os.Args[0])
}