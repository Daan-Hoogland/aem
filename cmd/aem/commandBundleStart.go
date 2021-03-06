package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func newBundleStartCommand() bundleStartCommand {
	return bundleStartCommand{
		name:    configDefaultInstance,
		http:    new(httpRequests),
		utility: new(utility),
		bundle:  "",
	}
}

type bundleStartCommand struct {
	name    string
	http    *httpRequests
	utility *utility
	bundle  string
}

func (c *bundleStartCommand) Execute(args []string) {
	u := utility{}
	c.getOpt(args)

	instance := u.getInstanceByName(c.name)
	bundlePicker := newBundlePicker()
	bundles := make([]bundle, 0)

	if len(c.bundle) > 0 {
		bundles = append(bundles, bundle{SymbolicName: c.bundle})
	} else {
		bundles = bundlePicker.picker(instance)
	}

	for _, bundle := range bundles {
		fmt.Printf("Starting %s\n", bundle.Name)
		resp := c.http.bundleStopStart(instance, bundle, BundleStatusStart)
		fmt.Printf("%s (%s) | Status %s -> %s\n", bundle.Name, bundle.SymbolicName, bundle.State, bundleRawState[resp.StateRaw])
	}

}

func (c *bundleStartCommand) getOpt(args []string) {
	getopt.FlagLong(&c.name, "name", 'n', "Name of instance to list bundles from from (default: "+configDefaultInstance+")")
	getopt.FlagLong(&c.bundle, "bundle", 'b', "bundle to start (Symbolic name)")
	getopt.CommandLine.Parse(args)
}
