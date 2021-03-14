package main

import (
	"fmt"
	"runtime"
)
type HostInfo struct {
	KernelFamily string
}

type Environment struct {
	host HostInfo
}

func main() {
	// Load config
	environment := probeEnvironment()

	fmt.Println(environment.host.KernelFamily)
	// Parse local database to find commands

	// Run them

}

func probeEnvironment() Environment {
	result := Environment{}
	result.host.KernelFamily = runtime.GOOS
	return result
}

func determineDependencyResolution() {

}