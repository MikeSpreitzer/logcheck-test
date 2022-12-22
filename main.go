package main

import "flag"
import "k8s.io/klog/v2"

func main() {
	klog.InitFlags(flag.CommandLine)
}
