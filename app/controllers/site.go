
package controllers

import (
	"os"
)

var CLONES_DIR = os.Getenv("GOPATH") + "/src/github.com/pedromorgan/revel-www/workspace"

var SiteSections []string = []string{"manual", "tutorial", "samples"}

