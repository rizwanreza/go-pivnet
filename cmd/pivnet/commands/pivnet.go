package commands

import (
	"fmt"
	"os"

	"github.com/pivotal-cf-experimental/go-pivnet"
	"github.com/pivotal-cf-experimental/go-pivnet/cmd/pivnet/version"
	"github.com/pivotal-golang/lager"
)

const (
	printAsTable = "table"
	printAsJSON  = "json"
	printAsYAML  = "yaml"
)

type PivnetCommand struct {
	Version func() `short:"v" long:"version" description:"Print the version of Pivnet and exit"`

	PrintAs string `long:"print-as" description:"Format to print as" default:"table" choice:"table" choice:"json" choice:"yaml"`

	APIToken string `long:"api-token" description:"Pivnet API token"`
	Endpoint string `long:"endpoint" description:"Pivnet API Endpoint"`

	EULAs      EULAsCommand      `command:"eulas" description:"List eulas"`
	AcceptEULA AcceptEULACommand `command:"accept-eula" description:"Accepts eula"`
	Product    ProductCommand    `command:"product" description:"Show product"`
}

var Pivnet PivnetCommand

func init() {
	Pivnet.Version = func() {
		fmt.Println(version.Version)
		os.Exit(0)
	}

	if Pivnet.Endpoint == "" {
		Pivnet.Endpoint = pivnet.Endpoint
	}
}

func NewClient() pivnet.Client {
	useragent := fmt.Sprintf(
		"go-pivnet/%s",
		version.Version,
	)

	pivnetClient := pivnet.NewClient(
		pivnet.ClientConfig{
			Token:     Pivnet.APIToken,
			Endpoint:  Pivnet.Endpoint,
			UserAgent: useragent,
		},
		lager.NewLogger("pivnet CLI"),
	)

	return pivnetClient
}