// Package version allows for the git SHA to be stored when being build via ldflags.
//
// Usage
//
// When building any packages that import version, pass the build/install cmd ldflags
// like so:
//    go install -ldflags "-X doge/version.commit GIT_GHA ...
//
// If you would like to make use of the flag either use InitVersion or interfact directly via Flag.
package version

import (
	"flag"
	"fmt"
	"os"
)

var (
	commit      string = ""
	versionFlag        = flag.Bool("version", false, "Print the version and exit.")
)

func init() {
	flag.BoolVar(versionFlag, "v", false, "Print the version and exit.")
}

// InitVersion must be run after flag.Parse(). It will print the version to
// stdout and then exit immediately. Deferred functions will not be run.
func InitVersion() {
	if *versionFlag {
		fmt.Println(GetCommit())
		os.Exit(0)
	}
}

func GetCommit() string {
	return commit
}

// Flag exposes the value of the versionFlag. Should only be used after flag.Parse().
func Flag() bool {
	return *versionFlag
}
