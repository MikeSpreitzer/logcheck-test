package main

import (
	"flag"
	"context"

	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(flag.CommandLine)
	flag.Parse()
	ctx := context.Background()
	logger := klog.FromContext(ctx)
	logger.Info("My message", "Fubar", "baz")
	logger.Info("My message", "fubar", "baz")
}
