package main

import (
	"github.com/fatih/color"
	"google.golang.org/grpc/grpclog"
)

func main() {
	l := grpclog.NewLoggerV2(color.New(color.FgHiBlack), color.New(color.FgHiYellow), color.New(color.FgHiRed))
	grpclog.SetLoggerV2(l)

	info, success, fail := color.New(color.FgHiWhite), color.New(color.FgHiGreen), color.New(color.FgHiRed)

	testKentikAPI(info, success, fail)
	testGoogleDatastoreAPI(info, success, fail)
}
