package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `
Usage:
  __projectname__ [-h]
  __projectname__ foo BAR

Arguments:
  BAR     some argument
Options:
  -h --help
`

	arguments, _ := docopt.Parse(usage, nil, true, "", false)
	fmt.Println(arguments)
}
