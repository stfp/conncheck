# Quick http2 checker

A simple program making grpc requests over http2 to various APIs.


## Kentik

This uses one of our simplest APIs. You need a valid user email and API token.
By default we hit the US saas cluster.

In the following example we run into a suspected issue where responses get mangled, likely by a proxy in the chain.

```
$ export KENTIK_EMAIL=user@domain.com
$ export KENTIK_TOKEN=000000000000000
$ export KENTIK_ROOT=grpc.api.kentik.com:443
$ ./conncheck
kentik api: could not list labels: rpc error: code = Internal desc = server closed the stream without sending trailers
```


## Google datastore

This uses the google GCP datastore API. 
You need a GCP project and credentials in json format which you can get from the gcp console.
https://cloud.google.com/docs/authentication/application-default-credentials#GAC

Could also just use `gcloud auth application-default login` when testing locally to automatically use your user credentials.a

```
$ export GCP_PROJECT=myproject
$ export GOOGLE_APPLICATION_CREDENTIALS=creds.json
$ ./conncheck
gcp datastore api: could not set entity: rpc error: code = Unknown desc =
```

In the example above we run into a suspected issue where responses get mangled, likely by a proxy in the chain.
