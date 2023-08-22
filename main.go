/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/dzeleniak/connartist/cmd"
	"github.com/dzeleniak/connartist/lib/registry"
)

var (
	globalRegistry *registry.Registry
)
func init() {
	r, err := registry.InitRegistry("registry", ".");
	if err != nil {
		panic(err);
	}

	globalRegistry = r;
}

func main() {
	defer globalRegistry.Cleanup();
	cmd.Execute()
}
