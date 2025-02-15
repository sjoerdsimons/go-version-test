package main

import (
	"github.com/sjoerdsimons/go-version-test"
	"fmt"
	"runtime/debug"
)

var Version string

func main() {
	gvt.SayHi();
	fmt.Printf("Hi\n");

	info, ok := debug.ReadBuildInfo();
	if !ok {
		fmt.Printf("No BuildInfo\n");
	} else {
		fmt.Printf("=== BuildInfo ===\n")
		fmt.Printf("Main Path: %v\n", info.Main.Path);
		fmt.Printf("Main Version: %v\n", info.Main.Version);
		fmt.Printf("Main Sum: %v\n", info.Main.Sum);
		for _, d := range info.Deps {
			fmt.Printf("Dep: %v\n", d);
		}
		for _, s := range info.Settings {
			fmt.Printf("Setting: %v\n", s);
		}
		fmt.Printf("======\n")
	}
	fmt.Printf("Version symbol: %v\n", Version);

	// Determine best version
	DeterminedVersion := "unknown"
	// Explicit set version first
	if len(Version) > 0 {
		DeterminedVersion = Version
	} else {
		// Try vcs version first as it will only be set on a local build
		var revision *string;
		var modified *string;
		for _, s := range info.Settings {
			if s.Key == "vcs.revision" {
				revision = &s.Value
			}
			if s.Key == "vcs.modified" {
				modified = &s.Value
			}
		}
		if revision != nil {
			DeterminedVersion = *revision;
			if modified != nil && *modified == "true" {
				DeterminedVersion += "-dirty"
			}
		} else {
			DeterminedVersion = info.Main.Version
		}
	}
	fmt.Printf("Determined Version: %v\n", DeterminedVersion);
}
