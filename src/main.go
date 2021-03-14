package main

import (
	"fmt"
	"flag"
	"runtime"
	"io/ioutil"
	"os"
	"encoding/json"
)
type HostInfo struct {
	KernelFamily string
}

type Environment struct {
	Host 			HostInfo
	Targets			[]string
}

type DatabaseEntry struct {
	TargetName 			   string
	InstallationCommands   map[string][]string
}

type Database struct {
	Entries []DatabaseEntry
}

func main() {
	// Load config
	environment := probeEnvironment()

	fmt.Println(environment.Host.KernelFamily)
	fmt.Print(environment.Targets)
	// Parse local database to find commands
	generateCommandList(environment)

	// Run them

}

func probeEnvironment() Environment {
	result := Environment{}
	result.Host.KernelFamily = runtime.GOOS

	flag.Parse()
	result.Targets = flag.Args()
	return result
}

func generateCommandList(environment Environment) {
	RawDBData, err := ioutil.ReadFile("gimmedb.json")
	db := Database{}
	AssertNo(err)
	err = json.Unmarshal(RawDBData, &db)
	AssertNo(err)
	for _, Target := range environment.Targets {
		fmt.Println(Target)
		var satisfied bool = false
		for _, DBEntry := range db.Entries {
			if DBEntry.TargetName == Target {
				satisfied = true
			}
		}
		if !satisfied {
			fmt.Printf("Could not figure out how to install %s", Target)
			os.Exit(1)
		}
	}
}

func AssertNo(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}