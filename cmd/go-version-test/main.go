package main

import (
	"github.com/sjoerdsimons/go-version-test"
	"fmt"
	"runtime/debug"
)


func main() {
	gvt.SayHi();
	fmt.Printf("Hi\n");

	info, ok := debug.ReadBuildInfo();
	if !ok {
		fmt.Printf("No BuildInfo\n");
	} else {
		fmt.Printf("=== BuildInfo ===\n")
		fmt.Printf("Main: %v\n", info.Main);
		for _, d := range info.Deps {
			fmt.Printf("Dep: %v\n", d);
		}
		for _, s := range info.Settings {
			fmt.Printf("Setting: %v\n", s);
		}
	}
}
