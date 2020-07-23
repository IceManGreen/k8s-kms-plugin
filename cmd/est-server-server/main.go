// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/loads"
	flag "github.com/spf13/pflag"
	"github.com/thalescpl-io/k8s-kms-plugin/pkg/est/restapi"
	"github.com/thalescpl-io/k8s-kms-plugin/pkg/est/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	var server *restapi.Server // make sure init is called

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprint(os.Stderr, "  est-server-server [OPTIONS]\n\n")

		title := "est server"
		fmt.Fprint(os.Stderr, title+"\n\n")
		desc := "RFC 7030 (EST) server implementation"
		if desc != "" {
			fmt.Fprintf(os.Stderr, desc+"\n\n")
		}
		fmt.Fprintln(os.Stderr, flag.CommandLine.FlagUsages())
	}
	// parse the CLI flags
	flag.Parse()

	api := operations.NewEstServerAPI(swaggerSpec)
	// get server with flag values filled out
	server = restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
