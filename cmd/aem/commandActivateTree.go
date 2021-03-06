package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

func newActivateTreeCommand() commandActivateTree {
	return commandActivateTree{
		Name:             configDefaultInstance,
		Path:             "",
		utility:          new(utility),
		i:                new(instance),
		projectStructure: new(projectStructure),
		http:             new(httpRequests),
	}
}

type commandActivateTree struct {
	Name             string
	Path             string
	utility          *utility
	i                *instance
	projectStructure *projectStructure
	http             *httpRequests
}

func (c *commandActivateTree) Execute(args []string) {
	c.getOpt(args)
	instance := c.i.getByName(c.Name)
	if c.http.activateTree(instance, c.Path) {
		fmt.Printf("Tree activated.\n")
	} else {
		fmt.Printf("Error while activating tree.\n")
	}
}

func (c *commandActivateTree) getOpt(args []string) {
	getopt.FlagLong(&c.Name, "instance", 'i', "Activate Tree on instance (Default: "+configDefaultInstance+")")
	getopt.FlagLong(&c.Path, "path", 'p', "Path to activate")
	getopt.CommandLine.Parse(args)
}
