package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sonm-io/marketplace/cli"
)

// build info
var (
	// AppVersion provides version of the application. Typically name of git branch. Set by linker.
	AppVersion string
	// GoVersion provides version of go the binary was compiled with. Set by linker.
	GoVersion string
	// BuildDate contains date of the build. Set by linker.
	BuildDate string
	// GitRev provides exact git revision of the source the binary was built from. Set by linker.
	GitRev string
	// GitLog provides exact git log. Set by linker.
	GitLog string
)

// flags
var (
	listenAddr    = flag.String("addr", ":9090", "SONM Marketplace service listen addr")
	showVersion   = flag.Bool("version", false, "Show SONM Marketplace version and exit")
	showBuildInfo = flag.Bool("build-info", false, "Display build info and exit")
)

func main() {
	flag.Parse()
	fillBuildInfo()

	if *showBuildInfo {
		printBuildInfo()
		return
	}

	if *showVersion {
		fmt.Printf("SONM Marketplace build version %q\n", AppVersion)
		return
	}

	log.Printf("Starting SONM Marketplace service at %s...\r\n", *listenAddr)
	app := cli.NewApp(cli.WithListenAddr(*listenAddr))
	if err := app.Run(); err != nil {
		log.Fatalf("Cannot start SONM Marketplace service: %s\r\n", err)
	}
}

func fillBuildInfo() {
	if AppVersion == "" {
		AppVersion = "dev"
	}
}

func printBuildInfo() {
	fmt.Printf("SONM Marketplace  build version %q\n", AppVersion)
	fmt.Printf("Built at %s with compiler %q\n", BuildDate, GoVersion)
	fmt.Printf("From git rev %s\n\n", GitRev)
	fmt.Printf("From git commit %s\n\n", GitLog)
}
