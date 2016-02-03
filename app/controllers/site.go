
package controllers

import (
	"os"
)

var CLONES_DIR = os.Getenv("GOPATH") + "/src/github.com/pedromorgan/revelframework.com/workspace"

var SiteSections []string = []string{"manual", "tutorial", "samples"}

