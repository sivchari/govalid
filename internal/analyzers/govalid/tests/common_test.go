package tests

import "flag"

var update bool

func init() {
	flag.BoolVar(&update, "update", false, "update golden files")
}
