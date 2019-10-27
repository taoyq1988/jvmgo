package main

import (
	"fmt"
	"github.com/taoyq1988/jvmgo/classpath"
	"github.com/taoyq1988/jvmgo/interpret"
	_ "github.com/taoyq1988/jvmgo/native"
	"github.com/taoyq1988/jvmgo/options"
	"github.com/taoyq1988/jvmgo/rtda/heap"
	"os"
	"strings"
)

const (
	version = "0.0.1"
)

func main() {
	opts, args := options.Parse()
	if opts.HelpFlag {
		printUsage()
	} else if opts.Version {
		printVersion()
	} else {
		startJVM(opts, args[0], args[1:])
	}
}

func printUsage() {
	fmt.Printf("usage: %s [-options] class [args...]\n", os.Args[0])
	options.PrintDefaults()
}

func printVersion() {
	fmt.Println(version)
}

func startJVM(opts options.Options, mainClassName string, args []string) {
	cp := classpath.NewClassPath(opts.BootJarPath, opts.Classpath)
	heap.InitBootLoader(cp, true)
	classLoader := heap.BootLoader()
	className := strings.Replace(mainClassName, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetStaticMethod("main", "([Ljava/lang/String;)V")
	interpret.Interpret(mainMethod)
}
