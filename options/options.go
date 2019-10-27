package options

import "flag"

const (
	defaultClasspath = "."
	defaultBootJar   = "./jar/rt.jar"
	//defaultBootJar   = "./jdk/build/libs/jdk.jar"
)

type Options struct {
	Classpath   string
	BootJarPath string
	HelpFlag    bool
	Version     bool
}

func Parse() (Options, []string) {
	options := Options{}
	flag.StringVar(&options.Classpath, "classpath", defaultClasspath, "Specifies a list of directories, JAR files, and ZIP archives to search for class files.")
	flag.StringVar(&options.Classpath, "cp", defaultClasspath, "Specifies a list of directories, JAR files, and ZIP archives to search for class files.")
	flag.StringVar(&options.BootJarPath, "bootJar", defaultBootJar, "rt jar path.")
	flag.BoolVar(&options.HelpFlag, "help", false, "Displays usage information and exit.")
	flag.BoolVar(&options.Version, "version", false, "jvm version.")
	flag.Parse()

	return options, flag.Args()
}

func PrintDefaults() {
	flag.PrintDefaults()
}
