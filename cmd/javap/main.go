package main

import (
	"fmt"
	"github.com/taoyq1988/jvmgo/classfile"
	"github.com/taoyq1988/jvmgo/classpath"
	"github.com/taoyq1988/jvmgo/options"
)

func main() {
	opts, args := options.Parse()
	if opts.HelpFlag || len(args) == 0 {
		printUsage()
	}
	printClassInfo(opts, args[0])
}

func printUsage() {
	fmt.Println("usage: javap [-options] class [args...]")
}

func printClassInfo(opts options.Options, className string) {
	cp := classpath.NewClassPath(opts.BootJarPath, opts.Classpath)
	data, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	class, err := classfile.Parse(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("class:", class.GetClassName(), ", access flag:", class.AccessFlags)
	fmt.Println("super class:", class.GetSuperClass())
	fmt.Println("interfaces:", class.GetInterfaces())
}
