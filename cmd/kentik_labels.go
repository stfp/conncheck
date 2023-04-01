package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/kentik/api-artifacts/go/kentik/label/v202210"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
	"os"
)

// testKentikAPI makes a simple http2 grpc call to the labels Kentik API using credentials
// specified in KENTIK_EMAIL and KENTIK_TOKEN environment variables.
// KENTIK_ROOT defaults to grpc.api.kentik.com:443 (our US saas instance) and can be overriden to test
// against other kentik clusters.
func testKentikAPI(info, success, fail io.Writer) {

	email := os.Getenv("KENTIK_EMAIL")
	token := os.Getenv("KENTIK_TOKEN")
	if email == "" || token == "" {
		fmt.Fprintln(fail, "kentik api: KENTIK_EMAIL/KENTIK_TOKEN are not set, skipping test")
		return
	}

	fmt.Fprintln(info, "\n\n\n***** testing with kentik API ****\n\n\n")

	target, ok := os.LookupEnv("KENTIK_ROOT")
	if !ok {
		target = "grpc.api.kentik.com:443"
	}

	ctx := context.Background()

	tc := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	conn, err := grpc.DialContext(ctx, target, grpc.WithTransportCredentials(tc))
	if err != nil {
		panic(err)
	}

	client := label.NewLabelServiceClient(conn)

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(
		"X-CH-Auth-Email", email,
		"X-CH-Auth-API-Token", os.Getenv("KENTIK_TOKEN"),
	))

	resp, err := client.ListLabels(ctx, &label.ListLabelsRequest{})
	if err != nil {
		fmt.Fprintf(fail, "kentik api: could not list labels: %v\n", err)
		return
	}

	fmt.Fprintf(info, "kentik api: got %d labels\n", len(resp.Labels))
	for i, l := range resp.Labels {
		fmt.Fprintf(fail, "kentik api: label %d: %s\n", i, l.String())
	}

	fmt.Fprintln(success, "kentik api: SUCCESS")
}
