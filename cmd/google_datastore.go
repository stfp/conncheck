package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"crypto/tls"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"os"
	"time"
)

type dataStoreEntity struct {
	Value string
}

// testGoogleDatastoreAPI performs a couple of get/put operations against the google cloud datastores API
// using http2 grpc.
//
// GCP's automagic credentials logic applies.
// For example using the SDK run gcloud auth application-default login
// Or set the GOOGLE_APPLICATION_CREDENTIALS env var to the name of a file containing GCP credentials.
// See https://cloud.google.com/docs/authentication/application-default-credentials#GAC
func testGoogleDatastoreAPI(info, success, fail io.Writer) {
	project := os.Getenv("GCP_PROJECT")
	if project == "" {
		fmt.Fprintln(fail, "gcp datastore api: GCP_PROJECT not set, skipping")
		return
	}

	fmt.Fprintln(info, "\n\n\n***** testing with gcp datastore API ****\n\n\n")

	ctx := context.Background()

	tc := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	dsClient, err := datastore.NewClient(ctx,
		project,
		option.WithGRPCDialOption(grpc.WithTransportCredentials(tc)),
	)
	if err != nil {
		fmt.Fprintln(fail, "gcp datastore api: could not create client: %v", err)
		return
	}
	defer dsClient.Close()

	k := datastore.NameKey("test", "stefan-testing", nil)
	e := new(dataStoreEntity)
	if err := dsClient.Get(ctx, k, e); err != nil {
		fmt.Fprintf(fail, "gcp datastore api: could not get entity: %v (continuing anyway)\n", err)
	}
	e.Value = time.Now().String()
	if _, err := dsClient.Put(ctx, k, e); err != nil {
		fmt.Fprintf(fail, "gcp datastore api: could not set entity: %v\n", err)
		return
	}

	fmt.Fprintln(success, "gcp datastore api: SUCCESS")
}
