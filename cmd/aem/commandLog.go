package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/pborman/getopt/v2"
)

func newLogCommand() logCommand {
	return logCommand{
		follow:           false,
		projectStructure: new(projectStructure),
		utility:          new(utility),
		name:             configDefaultInstance,
	}
}

type logCommand struct {
	follow           bool
	projectStructure *projectStructure
	utility          *utility
	name             string
}

func (s *logCommand) Execute(args []string) {
	s.getOpt(args)
	logFile := s.projectStructure.getLogFileLocation(s.utility.getInstanceByName(s.name))

	t, err := tail.TailFile(logFile, tail.Config{Follow: s.follow, MustExist: true})
	exitFatal(err, "Could not read error log. %s", logFile)
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}

func (s *logCommand) getOpt(args []string) {
	getopt.FlagLong(&s.follow, "follow", 'f', "Follow log file. Show new lines if they come in.")
	getopt.FlagLong(&s.name, "name", 'n', "Instance to start. (default: "+configDefaultInstance+")")
	getopt.CommandLine.Parse(args)
}
