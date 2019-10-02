package main

import (
	"encoding/hex"
	"fmt"
	"github.com/taoyq1988/jvmgo/classpath"
	"github.com/taoyq1988/jvmgo/options"
	"os"
)

const (
	version = "0.0.1"
)

func main() {
	opts, args := options.Parse()
	fmt.Println(args)
	if opts.HelpFlag {
		printUsage()
	} else if opts.Version {
		printVersion()
	} else if opts.Classpath != "" {
		cp := classpath.NewClassPath(opts.BootJarPath, opts.Classpath)
		d, _ := cp.ReadClass("java.lang.Object")
		fmt.Println(hex.EncodeToString(d))
	}
}

func printUsage() {
	fmt.Printf("usage: %s [-options] class [args...]\n", os.Args[0])
	options.PrintDefaults()
}

func printVersion() {
	fmt.Println(version)
}
