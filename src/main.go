package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

type HostInfo struct {
	KernelFamily string
}

type Environment struct {
	Host            HostInfo
	PackageManagers []string
	Targets         []string
	Euid            int
}

type DatabaseEntry struct {
	TargetName           string
	InstallationCommands map[string]map[string]string
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
	flag.Parse()
	result.Targets = flag.Args()

	result.Host.KernelFamily = runtime.GOOS

	for _, PkgMgr := range []string{"apt", "yum", "brew", "apk"} {
		if _, err := exec.LookPath(PkgMgr); err == nil {
			result.PackageManagers = append(result.PackageManagers, PkgMgr)
		}
	}

	result.Euid = os.Geteuid()
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
