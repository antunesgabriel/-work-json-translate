package main

import (
	"github.com/antunesgabriel/work-json-translate/cmd"
)

func main() {
	cmd.MakeEmptyJson("./users.json", "new.json", "output")
}
