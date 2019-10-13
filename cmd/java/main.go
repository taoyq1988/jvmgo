package main

import (
	"encoding/hex"
	"fmt"
	"github.com/taoyq1988/jvmgo/classfile"
	"github.com/taoyq1988/jvmgo/classpath"
	"github.com/taoyq1988/jvmgo/interpret"
	"github.com/taoyq1988/jvmgo/options"
	"os"
	"strings"
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

func startJVM(opts options.Options, mainClass string, args []string) {
	cp := classpath.NewClassPath(opts.BootJarPath, opts.Classpath)
	className := strings.Replace(mainClass, ".", "/", -1)
	cf := loadClass(className, cp)
	mainInfo := getMainMethodInfo(cf)
	interpret.Interpret(&mainInfo)
}

func loadClass(className string, classpath *classpath.ClassPath) *classfile.Classfile {
	classData, _ := classpath.ReadClass(className)
	fmt.Println("code", hex.EncodeToString(classData))
	cf, _ := classfile.Parse(classData)
	return cf
}

func getMainMethodInfo(cf *classfile.Classfile) classfile.MemberInfo {
	for _, m := range cf.Methods {
		mName := cf.GetConstantInfo(m.NameIndex).(string)
		fmt.Println(mName)
		if mName == "main" {
			fmt.Println("get main method")
			return m
		}
	}
	return classfile.MemberInfo{}
}
