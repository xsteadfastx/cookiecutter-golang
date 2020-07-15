package main

import (
	log "github.com/sirupsen/logrus"
    "{{ cookiecutter.repo }}/cmd"
)
var version string = "development"

func main() {
	log.SetLevel(log.DebugLevel)
	cmd.Run(version)
}
