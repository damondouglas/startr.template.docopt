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

  var commands map[string]func(map[string]interface{})
  commands = make(map[string]func(map[string]interface{}))
  commands["foo"] = foo

	arguments, _ := docopt.Parse(usage, nil, true, "", false)
  for key, f := range commands {
    if arguments[key] == true {
      f(arguments)
    }
  }
}

func foo(arguments map[string]interface{}) {
  fmt.Println(arguments)
}
